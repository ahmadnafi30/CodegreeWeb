package entity

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID            int64     `json:"id" gorm:"primaryKey;not null"`
	UserID        uuid.UUID `json:"user_id" gorm:"not null"`
	Amount        int64     `json:"amount" gorm:"not null"`
	Status        string    `json:"status" gorm:"not null;default:'Pending'"`
	SnapURL       string    `json:"snap_url"`
	TransactionID string    `json:"transaction_id"`
	PaymentURL    string    `json:"payment_url"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	EndAt         time.Time `json:"end_at" gorm:"-"`
}
