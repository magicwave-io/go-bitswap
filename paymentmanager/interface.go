package paymentmanager

import (
	"context"
)

type ClientHandler interface {
	ProcessCommand(commandId string, commandType int32, commandBody []byte, nodeId string, sessionId string) error
	ProcessPayment(paymentRequest string, nodeId string)
	ValidatePayment(paymentRequest string) (uint32,error)
	CreatePaymentInfo(amount uint32) (string, error)
	GetTransaction(sessionId string) (*PaymentTransaction, error)
	ProcessResponse(commandId string, responseBody []byte, nodeId string, sessionId string)
}

type CallbackHandler interface {
	Start()
	Shutdown(ctx context.Context)
}

