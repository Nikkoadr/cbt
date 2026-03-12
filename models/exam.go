package models

import (
	"time"

	"gorm.io/gorm"
)

// Exam represents the exam model
type Exam struct {
	gorm.Model
	Mapel     string
	KelasID   uint
	Tanggal   time.Time
	JamMulai  time.Time
	JamSelesai time.Time
	Durasi    int
	Status    string
	Kelas     Kelas `gorm:"foreignKey:KelasID"`
}
