package service

import (
	"CodegreeWebbs/entity"
	"CodegreeWebbs/internal/repository"
	"os"

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

	serverkey := os.Getenv("MIDTRANS_SERVER_KEY")
	client.New(serverkey, midtrans.Sandbox)

	// Check transaction status with Midtrans using orderId
	transactionStatusResp, err := client.CheckTransaction(orderId)
	if err != nil {
		return false, err
	}

	if transactionStatusResp != nil {
		switch transactionStatusResp.TransactionStatus {
		case "capture":
			if transactionStatusResp.FraudStatus == "challenge" {
				// TODO: Set transaction status on your database to 'challenge'
				// e.g.: 'Payment status challenged. Please take action on your Merchant Administration Portal'
			} else if transactionStatusResp.FraudStatus == "accept" {
				// TODO: Set transaction status on your database to 'success'
				s.UpdatePaymentStatus(orderId)
				return true, nil
			}
		case "settlement":
			// Set transaction status on your database to 'success'
			s.UpdatePaymentStatus(orderId)
			return true, nil
		case "deny":
			// TODO: You can ignore 'deny', because most of the time it allows payment retries
			// and later can become success
		case "cancel", "expire":
			// TODO: Set transaction status on your database to 'failure'
		case "pending":
			// TODO: Set transaction status on your database to 'pending' / waiting payment
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
