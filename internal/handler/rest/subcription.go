package rest

import (
	// "CodegreeWebbs/entity"
	"CodegreeWebbs/entity"
	"CodegreeWebbs/pkg/response"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	midtrans "github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func (r *Rest) CreatePayment(ctx *gin.Context) {
	userIDRaw, ok := ctx.Get("userID")
	if !ok {
		response.Error(ctx, http.StatusNotFound, "User ID not found.", nil)
		return
	}
	userID, ok := userIDRaw.(uuid.UUID)
	if !ok {
		response.Error(ctx, http.StatusInternalServerError, "Failed to parse User ID.", nil)
		return
	}
	userName := ctx.GetString("name")
	userEmail := ctx.GetString("email")
	clientkey := os.Getenv("MIDTRANS_CLIENT_KEY")
	serverkey := os.Getenv("MIDTRANS_SERVER_KEY")

	midtrans.ClientKey = clientkey
	midtrans.ServerKey = serverkey
	midtrans.Environment = midtrans.Sandbox

	orderID := uuid.New().String()

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: 149000,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: userName,
			Email: userEmail,
		},
	}

	if req.TransactionDetails.OrderID == "" {
		response.Error(ctx, http.StatusBadRequest, "Field OrderID is required", nil)
		return
	}

	snapResp, err := snap.CreateTransaction(req)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "error create transaction", err)
		return
	}

	data := entity.Payment{
		UserID:  userID,
		OrderID: orderID,
		SnapURL: snapResp.RedirectURL,
		Status:  "Pending",
		Amount:  149000,
		// Expiry:  time.Now().AddDate(0, 1, 0),
	}

	if err := r.service.PaymentService.CreatePayment(&data); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to save database", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success make payment", snapResp.RedirectURL)
}
func (r *Rest) PaymentHandlerNotification(ctx *gin.Context) {
	var notificationPayload map[string]interface{}
	if err := ctx.BindJSON(&notificationPayload); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Bad request", err)
		return
	}

	orderID, exists := notificationPayload["order_id"].(string)
	if !exists {
		response.Error(ctx, http.StatusBadRequest, "Order ID not found in payload", nil)
		return
	}

	success, err := r.service.PaymentService.VerifyPayment(ctx, orderID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to verify payment", err)
		return
	}

	if success {
		ctx.JSON(http.StatusOK, gin.H{"message": "Payment verified"})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to verify payment"})
}

func (r *Rest) CreateFreeTrial(ctx *gin.Context) {
	userIDRaw, ok := ctx.Get("userID")
	if !ok {
		response.Error(ctx, http.StatusNotFound, "User ID not found.", nil)
		return
	}
	userID, ok := userIDRaw.(uuid.UUID)
	if !ok {
		response.Error(ctx, http.StatusInternalServerError, "Failed to parse User ID.", nil)
		return
	}
	userName := ctx.GetString("name")
	userEmail := ctx.GetString("email")

	data := entity.FreeTrial{
		UserID: userID,
		Name:   userName,
		Email:  userEmail,
		// Expiry: time.Now().AddDate(0, 1, 0),
	}
	if err := r.service.PaymentService.CreateTrial(&data); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to save database", err)
		return
	}
	response.Success(ctx, http.StatusOK, "success make free trial", nil)
}
