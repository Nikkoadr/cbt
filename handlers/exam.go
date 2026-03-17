package handlers

import (
	"net/http"
	"strconv"
	"time"

	"cbt/models"
	"cbt/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ExamHandler holds the database connection
type ExamHandler struct {
	DB *gorm.DB
}

// GetDaftarUjian: Siswa melihat daftar ujian yang tersedia untuknya pada hari itu
func (h *ExamHandler) GetDaftarUjian(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		utils.SendError(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	var siswa models.Siswa
	if err := h.DB.Where("user_id = ?", userID).First(&siswa).Error; err != nil {
		utils.SendError(c, http.StatusNotFound, "Siswa not found")
		return
	}

	var siswaKelas models.SiswaKelas
	if err := h.DB.Where("siswa_id = ?", siswa.ID).First(&siswaKelas).Error; err != nil {
		utils.SendError(c, http.StatusNotFound, "Kelas siswa not found")
		return
	}

	var ujianPeserta []models.UjianPesertaKelas
	if err := h.DB.Where("kelas_id = ?", siswaKelas.KelasID).Find(&ujianPeserta).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to get exam participants")
		return
	}

	var daftarUjian []models.Ujian
	today := time.Now().Format("2006-01-02")
	for _, peserta := range ujianPeserta {
		var ujian models.Ujian
		if err := h.DB.Where("id = ? AND DATE(waktu_mulai) = ?", peserta.UjianID, today).First(&ujian).Error; err == nil {
			daftarUjian = append(daftarUjian, ujian)
		}
	}

	utils.SendSuccess(c, "Daftar Ujian Hari Ini", daftarUjian)
}

// MulaiUjian: Siswa memulai ujian dengan token yang valid
func (h *ExamHandler) MulaiUjian(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		utils.SendError(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	ujianIDStr := c.Param("id")
	ujianID, err := strconv.Atoi(ujianIDStr)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid Ujian ID")
		return
	}

	var input struct {
		Token string `json:"token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Token is required")
		return
	}

	// Verifikasi Ujian dan Token
	var ujian models.Ujian
	if err := h.DB.Where("id = ?", ujianID).First(&ujian).Error; err != nil {
		utils.SendError(c, http.StatusNotFound, "Ujian tidak ditemukan")
		return
	}

	if ujian.Token != input.Token {
		utils.SendError(c, http.StatusUnauthorized, "Token ujian tidak valid")
		return
	}

	var siswa models.Siswa
	if err := h.DB.Where("user_id = ?", userID).First(&siswa).Error; err != nil {
		utils.SendError(c, http.StatusNotFound, "Siswa not found")
		return
	}

	sesi := models.SesiUjian{
		SiswaID:    siswa.ID,
		UjianID:    uint(ujianID),
		IsSelesai:  false,
		WaktuMulai: time.Now(),
	}

	if err := h.DB.Create(&sesi).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Gagal memulai ujian")
		return
	}

	utils.SendSuccess(c, "Ujian berhasil dimulai", sesi)
}

// SimpanJawaban: Siswa menyimpan atau memperbarui jawaban untuk satu soal
func (h *ExamHandler) SimpanJawaban(c *gin.Context) {
	sesiIDStr := c.Param("sesiID")
	sesiID, err := strconv.Atoi(sesiIDStr)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid Sesi ID")
		return
	}

	var input struct {
		SoalID    string `json:"soal_id"`
		PilihanID string `json:"pilihan_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	soalID, err := strconv.Atoi(input.SoalID)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid Soal ID")
		return
	}

	pilihanID, err := strconv.Atoi(input.PilihanID)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid Pilihan ID")
		return
	}
	uintPilihanID := uint(pilihanID)

	// Cari jawaban yang sudah ada untuk soal ini di sesi ini (UPSERT)
	jawabanSiswa := models.JawabanSiswa{}
	h.DB.Where("sesi_ujian_id = ? AND soal_id = ?", sesiID, soalID).FirstOrInit(&jawabanSiswa)

	// Update field jawaban
	jawabanSiswa.SesiUjianID = uint(sesiID)
	jawabanSiswa.SoalID = uint(soalID)
	jawabanSiswa.PilihanID = &uintPilihanID
	jawabanSiswa.WaktuJawab = time.Now()

	// Simpan perubahan. GORM akan otomatis melakukan UPDATE jika record sudah ada, atau INSERT jika belum.
	if err := h.DB.Save(&jawabanSiswa).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Gagal menyimpan jawaban")
		return
	}

	utils.SendSuccess(c, "Jawaban berhasil disimpan", nil)
}

// SelesaikanUjian: Siswa menandai ujian sebagai selesai (tanpa perhitungan skor)
func (h *ExamHandler) SelesaikanUjian(c *gin.Context) {
	sesiIDStr := c.Param("sesiID")
	sesiID, err := strconv.Atoi(sesiIDStr)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid Sesi ID")
		return
	}

	var sesi models.SesiUjian
	if err := h.DB.First(&sesi, uint(sesiID)).Error; err != nil {
		utils.SendError(c, http.StatusNotFound, "Sesi ujian tidak ditemukan")
		return
	}

	// Pastikan sesi ini milik siswa yang sedang login (tambahan keamanan)
	userID, _ := c.Get("userID")
	var siswa models.Siswa
	h.DB.Where("user_id = ?", userID).First(&siswa)
	if sesi.SiswaID != siswa.ID {
		utils.SendError(c, http.StatusForbidden, "Anda tidak memiliki akses ke sesi ujian ini")
		return
	}

	now := time.Now()

	// Hanya update kolom yang diperlukan, tanpa menghitung skor
	if err := h.DB.Model(&sesi).Updates(models.SesiUjian{IsSelesai: true, WaktuSelesai: &now}).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Gagal menyelesaikan ujian")
		return
	}

	utils.SendSuccess(c, "Ujian telah selesai. Jawaban Anda telah direkam dan akan segera dinilai.", nil)
}
