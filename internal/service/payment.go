package service

import (
	"CodegreeWebbs/internal/repository"
)

type SPayment interface {
}

type PaymentService struct {
	PaymentRepo repository.IPayment
}

func NewPaymentService(paymentRepo repository.IPayment) SPayment {
	return &PaymentService{PaymentRepo: paymentRepo}
}
