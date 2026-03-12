package models

import "gorm.io/gorm"

// Kelas represents the kelas model
type Kelas struct {
	gorm.Model
	NamaKelas string `gorm:"unique"`
}
