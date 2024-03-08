package entity

import "github.com/google/uuid"

type Onboarding struct {
	ID                    uint     `json:"id" gorm:"primary_key;autoIncrement"`
	Question              string   `json:"question" gorm:"type:text"`
	Options               []Option `json:"options" gorm:"foreignkey:QuestionID"`
	Pointrange            uint
	RecommendedLanguageID uuid.UUID `json:"-"`
}
