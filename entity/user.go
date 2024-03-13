package entity

import (
	"time"

	"github.com/google/uuid"
	// "github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null;unique"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null;"`
	RoleID    uint      `json:"role_id" gorm:"not null"`
	Level     int       `json:"level"`
	Xp        uint64    `json:"xp"`
	Hearth    int       `json:"hearth"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
