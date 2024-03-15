package entity

// import "github.com/google/uuid"

type LanguageCode struct {
	ID          uint   `json:"id" gorm:"autoIncrement"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
