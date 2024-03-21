package repository

import (
	"CodegreeWebbs/entity"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IPayment interface {
	Save(payment *entity.Payment) error
	UpdatePaymentStatus(orderID string) error
	DeletOrderID(orderID string) error
	CreateFreeTrial(freetrial *entity.FreeTrial) error
	CheckTransaction(userid uuid.UUID) error
	CheckTrial(userid uuid.UUID) error
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

func (repo *PaymentRepo) CreateFreeTrial(freetrial *entity.FreeTrial) error {
	if err := repo.db.Debug().Create(freetrial).Error; err != nil {
		return err
	}
	return nil
}

func (repo *PaymentRepo) CheckTransaction(userid uuid.UUID) error {
	var payment entity.Payment
	if err := repo.db.Where("user_id = ?", userid).First(&payment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("you are not subscribed")
		}
		return err
	}
	if payment.Status != "paid" {
		return errors.New("your transaction is not paid yet")
	}

	if payment.CreatedAt.AddDate(0, 1, 0).Before(time.Now()) {
		return errors.New("your subscription has ended")

	}

	return nil
}

func (repo *PaymentRepo) CheckTrial(userid uuid.UUID) error {
	var payment entity.FreeTrial
	if err := repo.db.Where("user_id = ?", userid).First(&payment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("you are not subscribed")
		}
		return err
	}

	if payment.CreatedAt.AddDate(0, 1, 0).Before(time.Now()) {
		return errors.New("your subscription has ended")

	}

	return nil
}
