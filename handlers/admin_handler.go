package handlers

import (
	"net/http"
	"strconv"

	"cbt/models"
	"cbt/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AdminHandler holds the database connection
type AdminHandler struct {
	DB *gorm.DB
}

// GetUjianToken: Admin melihat token untuk ujian tertentu
func (h *AdminHandler) GetUjianToken(c *gin.Context) {
	ujianIDStr := c.Param("id")
	ujianID, err := strconv.Atoi(ujianIDStr)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid Ujian ID")
		return
	}

	var ujian models.Ujian
	if err := h.DB.First(&ujian, uint(ujianID)).Error; err != nil {
		utils.SendError(c, http.StatusNotFound, "Ujian tidak ditemukan")
		return
	}

	utils.SendSuccess(c, "Token Ujian", gin.H{"token": ujian.Token})
}

// KoreksiUjian: Admin menghitung dan menyimpan skor untuk sesi ujian yang sudah selesai
func (h *AdminHandler) KoreksiUjian(c *gin.Context) {
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

	// Pastikan ujian memang sudah diselesaikan oleh siswa sebelum dikoreksi
	if !sesi.IsSelesai {
		utils.SendError(c, http.StatusConflict, "Ujian ini belum diselesaikan oleh siswa")
		return
	}

	var jawabanSiswa []models.JawabanSiswa
	if err := h.DB.Where("sesi_ujian_id = ?", sesiID).Find(&jawabanSiswa).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil jawaban siswa")
		return
	}

	// Logika perhitungan skor
	var jawabanBenar int
	for _, jawaban := range jawabanSiswa {
		var soalPilihan models.SoalPilihan
		if jawaban.PilihanID != nil {
			// Periksa apakah pilihan siswa adalah kunci jawaban
			if err := h.DB.Where("id = ? AND adalah_kunci = TRUE", *jawaban.PilihanID).First(&soalPilihan).Error; err == nil {
				jawabanBenar++
			}
		}
	}

	var skor float64
	if len(jawabanSiswa) > 0 {
		skor = (float64(jawabanBenar) / float64(len(jawabanSiswa))) * 100
	}

	// Simpan skor ke dalam sesi ujian
	sesi.Skor = skor
	if err := h.DB.Save(&sesi).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Gagal menyimpan skor ujian")
		return
	}

	utils.SendSuccess(c, "Ujian berhasil dikoreksi dan skor telah disimpan", gin.H{"skor": skor})
}
