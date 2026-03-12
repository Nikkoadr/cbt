package models

// NilaiAkhir represents the nilai_akhir model
type NilaiAkhir struct {
	ID        uint    `gorm:"primaryKey"`
	UjianID   uint
	SiswaID   uint
	Nilai     float64 `gorm:"type:decimal(5,2);default:0.00"`
	SudahLulus bool    `gorm:"type:tinyint(1);default:0"`
	Ujian     Ujian   `gorm:"foreignKey:UjianID"`
	Siswa     Siswa   `gorm:"foreignKey:SiswaID"`
}
