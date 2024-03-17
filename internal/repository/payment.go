package repository

import (
	// "CodegreeWebbs/entity"

	// "github.com/google/uuid"
	"gorm.io/gorm"
)

type IPayment interface {
}

type PaymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) IPayment {
	return &PaymentRepo{db: db}
}

// func (repo *PaymentRepo) FindById(userId uuid.UUID) (entity.Payment, error) {
// 	var user entity.User
// 	err := repo.db.Debug().Where("id = ?", userId).First(&user).Error
// 	if err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }
