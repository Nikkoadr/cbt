package models

// TipeSoal mendefinisikan tipe soal
type TipeSoal string

const (
	PilihanGanda TipeSoal = "pg"
	Essay        TipeSoal = "essay"
)

// TingkatKesulitanSoal mendefinisikan tingkat kesulitan soal
type TingkatKesulitanSoal string

const (
	Mudah  TingkatKesulitanSoal = "mudah"
	Sedang TingkatKesulitanSoal = "sedang"
	Sulit  TingkatKesulitanSoal = "sulit"
)

// Soal represents the soal model
type Soal struct {
	ID              uint                 `gorm:"primaryKey"`
	GuruID          *uint
	TeksSoal        string               `gorm:"type:text;not null"`
	TipeSoal        TipeSoal             `gorm:"type:enum('pg','essay');default:'pg'"`
	KategoriSoal    string               `gorm:"type:varchar(100)"`
	TingkatKesulitan TingkatKesulitanSoal `gorm:"type:enum('mudah','sedang','sulit');default:'sedang'"`
	Guru            *Guru                `gorm:"foreignKey:GuruID"`
}
