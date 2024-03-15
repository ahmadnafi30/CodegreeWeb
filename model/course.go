package model

import "github.com/google/uuid"

type CreateCourse struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	LanguageCode uuid.UUID `json:"language_code"`
}
