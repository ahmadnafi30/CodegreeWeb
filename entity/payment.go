package entity

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID        int64     `json:"id" gorm:"primaryKey;not null"`
	UserID    uuid.UUID `json:"user_id" gorm:"not null"`
	OrderID   string    `json:"order_id"`
	Amount    int64     `json:"amount" gorm:"not null"`
	Status    string    `json:"status" gorm:"not null;default:'Pending'"`
	SnapURL   string    `json:"snap_url"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	// Expiry    time.Time `json:"expiry" gorm:"not null"`
}

type FreeTrial struct {
	ID        int64     `json:"id" gorm:"primaryKey;not null"`
	UserID    uuid.UUID `json:"user_id" gorm:"not null"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	// Expiry    time.Time `json:"expiry" gorm:"not null"`
}
