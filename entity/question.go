package entity

import "github.com/google/uuid"

type Question struct {
	ID            uint      `json:"id"`
	SubLanguageID uuid.UUID `json:"sub_language_id"`
	Question      string    `json:"question"`
	Options       []Option  `json:"options,omitempty"`
	Score         int64     `json:"score"`
}
