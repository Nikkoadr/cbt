package models

// Admin represents the admin model
type Admin struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint `gorm:"unique"`
	User   User `gorm:"foreignKey:UserID"`
}
