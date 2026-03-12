package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"cbt/models"
	"cbt/utils"
)

// ExamHandler holds the database connection
type ExamHandler struct {
	DB *gorm.DB
}

// GetTodayExams returns today's exams for the logged in user's class
func (h *ExamHandler) GetTodayExams(c *gin.Context) {
	kelasID, _ := c.Get("kelas_id")

	var exams []models.Exam
	today := time.Now()
	year, month, day := today.Date()

	if err := h.DB.Where("kelas_id = ? AND DATE(tanggal) = ?", kelasID, time.Date(year, month, day, 0, 0, 0, 0, today.Location())).Find(&exams).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to fetch exams")
		return
	}

	utils.SendSuccess(c, "Today's exams", exams)
}

// GetExamDetails returns the details of a specific exam
func (h *ExamHandler) GetExamDetails(c *gin.Context) {
	var exam models.Exam
	if err := h.DB.First(&exam, c.Param("id")).Error; err != nil {
		utils.SendError(c, http.StatusNotFound, "Exam not found")
		return
	}

	utils.SendSuccess(c, "Exam details", exam)
}

// CheckExamAccess checks if an exam can be accessed
func (h *ExamHandler) CheckExamAccess(c *gin.Context) {
	var exam models.Exam
	if err := h.DB.First(&exam, c.Param("id")).Error; err != nil {
		utils.SendError(c, http.StatusNotFound, "Exam not found")
		return
	}

	if time.Now().Before(exam.JamMulai) {
		utils.SendError(c, http.StatusForbidden, "Ujian belum dimulai")
	} else {
		utils.SendSuccess(c, "Ujian dapat diakses", nil)
	}
}
