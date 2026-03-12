package models

// UjianSoal represents the ujian_soal model
type UjianSoal struct {
	ID          uint  `gorm:"primaryKey"`
	UjianID     uint
	SoalID      uint
	UrutanTampil int
	Ujian       Ujian `gorm:"foreignKey:UjianID"`
	Soal        Soal  `gorm:"foreignKey:SoalID"`
}
