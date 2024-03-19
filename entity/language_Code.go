package entity

// import "github.com/google/uuid"

type LanguageCode struct {
	ID          uint     `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string   `json:"title" gorm:"type:varchar(255);not null"`
	Description string   `json:"description" gorm:"type:text;not null"`
	RecomSubab  []*Recom `json:"recom_subab" gorm:"foreignKey:LanguageCodeID;references:ID"`
}
