package models

// Guru represents the guru model
type Guru struct {
	ID     uint   `gorm:"primaryKey"`
	UserID uint   `gorm:"unique"`
	NIP    string `gorm:"type:varchar(20);unique"`
	User   User   `gorm:"foreignKey:UserID"`
}
