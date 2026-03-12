package models

// Siswa represents the siswa model
type Siswa struct {
	ID     uint   `gorm:"primaryKey"`
	UserID uint   `gorm:"unique"`
	NISN   string `gorm:"type:varchar(20);unique;not null"`
	User   User   `gorm:"foreignKey:UserID"`
}
