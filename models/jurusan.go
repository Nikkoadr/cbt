package models

// Jurusan represents the jurusan model
type Jurusan struct {
	ID           uint   `gorm:"primaryKey"`
	KodeJurusan  string `gorm:"type:varchar(10);not null"`
	NamaJurusan  string `gorm:"type:varchar(100);not null"`
}
