package entity

import "github.com/google/uuid"

type UserProfile struct {
	ID     uuid.UUID `gorm:"primaryKey"`
	Name   string
	Email  string
	Level  int
	Xp     uint64
	Hearth int
}
