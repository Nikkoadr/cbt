package models

import "time"

// SesiUjian represents the sesi_ujian model
type SesiUjian struct {
	ID           uint      `gorm:"primaryKey"`
	UjianID      uint
	SiswaID      uint
	WaktuMulai   time.Time
	WaktuSelesai *time.Time
	IsSelesai    bool      `gorm:"type:tinyint(1);default:0"`
	Skor         float64   `gorm:"type:decimal(5,2);default:0.00"`
	Ujian        Ujian     `gorm:"foreignKey:UjianID"`
	Siswa        Siswa     `gorm:"foreignKey:SiswaID"`
}
