package models

import "gorm.io/gorm"

// User represents the user model
type User struct {
	gorm.Model
	NIS      string `gorm:"unique"`
	Nama     string
	KelasID  uint
	Username string `gorm:"unique"`
	Password string
	Kelas    Kelas  `gorm:"foreignKey:KelasID"`
}
