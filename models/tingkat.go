package models

// Tingkat represents the tingkat model
type Tingkat struct {
	ID         uint   `gorm:"primaryKey"`
	NamaTingkat string `gorm:"type:enum('10','11','12');not null"`
}
