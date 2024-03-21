package repository

import (
	"CodegreeWebbs/entity"

	"gorm.io/gorm"
)

type IPayment interface {
	Save(payment *entity.Payment) error
	UpdatePaymentStatus(orderID string) error
	DeletOrderID(orderID string) error
}

type PaymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) IPayment {
	return &PaymentRepo{db: db}
}

func (repo *PaymentRepo) Save(payment *entity.Payment) error {
	err := repo.db.Debug().Create(payment).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *PaymentRepo) UpdatePaymentStatus(orderID string) error {
	var payment entity.Payment
	if err := repo.db.Where("order_id = ?", orderID).First(&payment).Error; err != nil {
		return err
	}
	payment.Status = "paid"
	if err := repo.db.Save(&payment).Error; err != nil {
		return err
	}
	return nil
}

func (repo *PaymentRepo) DeletOrderID(orderID string) error {
	var payment entity.Payment
	if err := repo.db.Where("order_id = ?", orderID).Delete(&payment).Error; err != nil {
		return err
	}
	return nil
}
