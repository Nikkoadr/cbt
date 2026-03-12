package models

// SoalPilihan represents the soal_pilihan model
type SoalPilihan struct {
	ID          uint    `gorm:"primaryKey"`
	SoalID      uint
	TeksPilihan string  `gorm:"type:text;not null"`
	AdalahKunci bool    `gorm:"type:tinyint(1);default:0"`
	BobotPilihan float64 `gorm:"type:decimal(3,2);default:0.00"`
	Soal        Soal    `gorm:"foreignKey:SoalID"`
}
