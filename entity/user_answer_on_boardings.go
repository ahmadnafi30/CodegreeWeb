package entity

import "github.com/google/uuid"

type UserAnswerOnBoarding struct {
	UserID     uuid.UUID `gorm:"primaryKey"`
	QuestionID uint      `json:"question_id" binding:"required"`
	Answer     uint      `json:"answer" binding:"required"`
}
