package entity

import "github.com/google/uuid"

type LanguageCode struct {
	ID          uuid.UUID  `json:"id" gorm:"type:varchar(50);primary_key"`
	Title       string     `json:"title" gorm:"type:varchar(255);not null;unique"`
	Description string     `json:"description" gorm:"type:text"`
	Questions   []Question `json:"questions,omitempty" gorm:"foreignkey:LanguageCodeID"`
}
