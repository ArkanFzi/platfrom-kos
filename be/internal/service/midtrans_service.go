package service

import (
	"koskosan-be/internal/models"
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransService interface {
	CreateTransaction(booking *models.Pemesanan, amount float64) (*snap.Response, string, error)
	VerifyNotification(payload map[string]interface{}) (string, string, error)
	CheckTransaction(orderID string) (*coreapi.TransactionStatusResponse, error)
}

type midtransService struct {
	snapClient    snap.Client
	coreApiClient coreapi.Client
}

func NewMidtransService() MidtransService {
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	isProd := os.Getenv("MIDTRANS_IS_PRODUCTION") == "true"
	
	s := snap.Client{}
	s.New(serverKey, midtrans.Sandbox)
	
	c := coreapi.Client{}
	c.New(serverKey, midtrans.Sandbox)

	if isProd {
		s.New(serverKey, midtrans.Production)
		c.New(serverKey, midtrans.Production)
	}

	return &midtransService{
		snapClient:    s,
		coreApiClient: c,
	}
}

func (s *midtransService) CreateTransaction(booking *models.Pemesanan, amount float64) (*snap.Response, string, error) {
	orderID := "ORDER-" + booking.Penyewa.NamaLengkap + "-" + string(rune(booking.ID)) 
	
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: int64(amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: booking.Penyewa.NamaLengkap,
			Email: "customer@example.com", 
		},
	}

	snapResp, err := s.snapClient.CreateTransaction(req)
	if err != nil {
		return nil, "", err
	}

	return snapResp, orderID, nil
}

func (s *midtransService) VerifyNotification(payload map[string]interface{}) (string, string, error) {
	orderID, _ := payload["order_id"].(string)
	transactionStatus, _ := payload["transaction_status"].(string)
	
	return orderID, transactionStatus, nil
}

func (s *midtransService) CheckTransaction(orderID string) (*coreapi.TransactionStatusResponse, error) {
	return s.coreApiClient.CheckTransaction(orderID)
}
