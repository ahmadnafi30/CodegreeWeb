package entity

import "github.com/google/uuid"

type UserAnswerOnBoarding struct {
	ID         uint      `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	QuestionID uint      `json:"question_id"`
	Response   string    `json:"response"`
}
