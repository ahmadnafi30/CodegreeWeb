package entity

import "github.com/google/uuid"

type Payment struct {
	Id            int64     `json:"id" gorm:"primaryKey;not null"`
	UserID        uuid.UUID `json:"user_id" gorm:"not null"`
	Amount        int64     `json:"amount" gorm:"not null"`
	Status        bool      `json:"status" gorm:"not null"`
	SnapURL       string    `json:"snap_url"`
	TransactionID string    `json:"transaction_id"`
	PaymentURL    string    `json:"payment_url"`
}
