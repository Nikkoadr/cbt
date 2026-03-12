package models

// Kelas represents the kelas model
type Kelas struct {
	ID               uint    `gorm:"primaryKey"`
	TingkatID        uint
	JurusanID        uint
	NomorKelas       string  `gorm:"type:varchar(5)"`
	NamaKelasLengkap string  `gorm:"type:varchar(50)"`
	Tingkat          Tingkat `gorm:"foreignKey:TingkatID"`
	Jurusan          Jurusan `gorm:"foreignKey:JurusanID"`
}
