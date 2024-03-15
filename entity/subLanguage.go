package entity

import "github.com/google/uuid"

type SubLanguage struct {
	ID         uint       `json:"id"`
	Title      string     `json:"title"`
	CourseID   uuid.UUID  `json:"course_id"`
	MaterialID uint       `json:"material_id"`
	Material   Material   `json:"material,omitempty"`
	Questions  []Question `json:"questions,omitempty"`
}
