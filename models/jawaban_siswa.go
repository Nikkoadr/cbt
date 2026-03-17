package models

import "time"

// JawabanSiswa represents the jawaban_siswa model
type JawabanSiswa struct {
	ID           uint         `gorm:"primaryKey"`
	SesiUjianID  uint
	SoalID       uint
	PilihanID    *uint
	TeksJawaban  *string      `gorm:"type:text"`
	AdalahBenar  bool         `gorm:"type:tinyint(1);default:0"`
	WaktuJawab   time.Time    `gorm:"not null"`
	SesiUjian    SesiUjian    `gorm:"foreignKey:SesiUjianID"`
	Soal         Soal         `gorm:"foreignKey:SoalID"`
	SoalPilihan  *SoalPilihan `gorm:"foreignKey:PilihanID"`
}
