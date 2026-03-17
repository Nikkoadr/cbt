package models

import "time"

// Ujian represents the ujian model
type Ujian struct {
	ID               uint                  `gorm:"primaryKey"`
	NamaUjian        string                `gorm:"type:varchar(255);not null"`
	Deskripsi        string                `gorm:"type:text"`
	WaktuMulai       time.Time             `gorm:"not null"`
	Durasi           int                   `gorm:"not null"` // Durasi dalam menit
	Token            string                `gorm:"type:varchar(255);unique;not null"`
	Soal             []Soal                `gorm:"foreignKey:UjianID"`
	UjianPesertaKelas []UjianPesertaKelas `gorm:"foreignKey:UjianID"`
}
