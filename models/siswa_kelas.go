package models

// SiswaKelas represents the siswa_kelas model
type SiswaKelas struct {
	SiswaID uint `gorm:"primaryKey"`
	KelasID uint `gorm:"primaryKey"`
	Siswa   Siswa `gorm:"foreignKey:SiswaID"`
	Kelas   Kelas `gorm:"foreignKey:KelasID"`
}
