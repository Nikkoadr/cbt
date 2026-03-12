package models

// UjianPesertaKelas represents the ujian_peserta_kelas model
type UjianPesertaKelas struct {
	UjianID uint `gorm:"primaryKey"`
	KelasID uint `gorm:"primaryKey"`
	Ujian   Ujian `gorm:"foreignKey:UjianID"`
	Kelas   Kelas `gorm:"foreignKey:KelasID"`
}
