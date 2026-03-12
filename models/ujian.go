package models

import "time"

// Ujian represents the ujian model
type Ujian struct {
	ID           uint      `gorm:"primaryKey"`
	JudulUjian   string    `gorm:"type:varchar(255);not null"`
	GuruID       *uint
	DurasiMenit  int       `gorm:"not null"`
	WaktuMulai   time.Time `gorm:"not null"`
	WaktuSelesai time.Time `gorm:"not null"`
	Token        string    `gorm:"type:varchar(10)"`
	AcakSoal     bool      `gorm:"type:tinyint(1);default:0"`
	Guru         *Guru     `gorm:"foreignKey:GuruID"`
}
