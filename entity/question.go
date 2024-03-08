package entity

import "github.com/google/uuid"

type Question struct {
	ID             uint      `json:"id" gorm:"primary_key;autoIncrement"`
	QuizID         uint      `json:"quiz_id" gorm:"index"`
	LanguageCodeID uuid.UUID `json:"language_code_id"`
	Question       string    `json:"question" gorm:"type:text"`
	Options        []Option  `json:"options" gorm:"foreignkey:QuestionID"`
	CorrectAnswer  string    `json:"correct_answer"`
	Score          int64     `json:"score"`
}
