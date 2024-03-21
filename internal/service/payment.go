package service

import (
	"CodegreeWebbs/entity"
	"CodegreeWebbs/internal/repository"

	// "CodegreeWebbs/pkg/response"

	// "encoding/json"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"golang.org/x/net/context"
)

type SPayment interface {
	CreatePayment(payment *entity.Payment) error
	UpdatePaymentStatus(data string) error
	VerifyPayment(ctx context.Context, orderId string) (bool, error)
	CreateTrial(freetrial *entity.FreeTrial) error
	CheckAccess(userID uuid.UUID) error
}

type PaymentService struct {
	PaymentRepo    repository.IPayment
	MidtransClient *midtrans.HttpClient
}

func NewPaymentService(paymentRepo repository.IPayment) SPayment {
	return &PaymentService{PaymentRepo: paymentRepo}
}

func (service *PaymentService) CreatePayment(payment *entity.Payment) error {
	err := service.PaymentRepo.Save(payment)
	if err != nil {
		return err
	}
	return nil
}

func (s *PaymentService) UpdatePaymentStatus(data string) error {
	orderID := data

	err := s.PaymentRepo.UpdatePaymentStatus(orderID)
	if err != nil {
		return err
	}

	return nil
}
func (s *PaymentService) VerifyPayment(ctx context.Context, orderId string) (bool, error) {

	var client coreapi.Client

	client.New("SB-Mid-client-viLy_yj40DPmwY0C", midtrans.Sandbox)

	// Parse JSON request body and use it to set json to payload

	// Check transaction status with Midtrans using orderId
	transactionStatusResp, e := client.CheckTransaction(orderId)
	if e != nil {
		return false, e
	} else {
		if transactionStatusResp != nil {
			// Do set transaction status based on response from check transaction status
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					// TODO set transaction status on your database to 'challenge'
					// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
				} else if transactionStatusResp.FraudStatus == "accept" {
					// TODO set transaction status on your database to 'success'
					return true, nil
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				s.UpdatePaymentStatus(orderId)
				return true, nil
			} else if transactionStatusResp.TransactionStatus == "deny" {
				// TODO you can ignore 'deny', because most of the time it allows payment retries
				// and later can become success
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				// TODO set transaction status on your databaase to 'failure'
			} else if transactionStatusResp.TransactionStatus == "pending" {
				// TODO set transaction status on your databaase to 'pending' / waiting payment
			}
		}
	}
	return false, nil
}
func (s *PaymentService) CreateTrial(freetrial *entity.FreeTrial) error {
	err := s.PaymentRepo.CreateFreeTrial(freetrial)
	if err != nil {
		return err
	}
	return nil
}

func (s *PaymentService) CheckAccess(userID uuid.UUID) error {
	err := s.PaymentRepo.CheckTransaction(userID)
	if err != nil {
		return err
	}

	err = s.PaymentRepo.CheckTrial(userID)
	if err != nil {
		return err
	}

	return nil
}
