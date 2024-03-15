package model

import "github.com/google/uuid"

// import "github.com/google/uuid"

type CreateMultipleChoiceOnboarding struct {
	Question string `json:"question"`
	Options  []struct {
		Description string `json:"description"`
		Language    string `json:"-"`
	} `json:"options"`
}

type UserAnswerOnBoarding struct {
	ID         uuid.UUID `json:"id"`
	QuestionID uint      `json:"question_id" binding:"required"`
	Answer     uint      `json:"answer" binding:"required"`
}

type RecommendLanguage struct {
	ID                uuid.UUID `json:"id"`
	RecommendLanguage string    `json:"recomended"`
}
