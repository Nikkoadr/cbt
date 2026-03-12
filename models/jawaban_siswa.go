package models

import "time"

// JawabanSiswa represents the jawaban_siswa model
type JawabanSiswa struct {
	ID               uint           `gorm:"primaryKey"`
	PercobaanID      uint
	SoalID           uint
	PilihanID        *uint
	TeksJawaban      *string        `gorm:"type:text"`
	AdalahBenar      bool           `gorm:"type:tinyint(1);default:0"`
	WaktuJawab       time.Time      `gorm:"not null"`
	UjianPercobaan   UjianPercobaan `gorm:"foreignKey:PercobaanID"`
	Soal             Soal           `gorm:"foreignKey:SoalID"`
	SoalPilihan      *SoalPilihan   `gorm:"foreignKey:PilihanID"`
}
