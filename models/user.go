package models

import (
	"time"
)

// PeranUser mendefinisikan peran pengguna
type PeranUser string

const (
	PeranAdmin PeranUser = "admin"
	PeranGuru  PeranUser = "guru"
	PeranSiswa PeranUser = "siswa"
)

// User represents the user model
type User struct {
	ID          uint      `gorm:"primaryKey"`
	Username    string    `gorm:"type:varchar(50);unique;not null"`
	Password    string    `gorm:"type:varchar(255);not null"`
	NamaLengkap string    `gorm:"type:varchar(100);not null"`
	Peran       PeranUser `gorm:"type:enum('admin','guru','siswa');not null"`
	IsAktif     bool      `gorm:"type:tinyint(1);default:1"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:current_timestamp"`
}
