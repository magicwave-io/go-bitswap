package paymentmanager

import (
	"github.com/libp2p/go-libp2p-core/peer"
	"testing"
)

type PaymentHandlerMock struct {
	debtRegistry	map[peer.ID]*Debt
	requestedPaymentAmount	int
	paymentRequest	string
}

func (p *PaymentHandlerMock) GetDebt(id peer.ID) *Debt {
	return p.debtRegistry[id]
}

func (p *PaymentHandlerMock) CallProcessCommand(commandType int32, commandBody string) (string, error) {
	panic("implement me")
}

func (p *PaymentHandlerMock) CallProcessPayment(paymentRequest string, requestReference string) {
	panic("implement me")
}

func (p *PaymentHandlerMock) CreatePaymentInfo(amount int) (string, error) {
	p.requestedPaymentAmount = amount

	return p.paymentRequest, nil
}

func (p *PaymentHandlerMock) CallProcessResponse(commandId string, responseBody string, nodeId string) {
	panic("implement me")
}

type PeerHandlerMock struct {
	paymentRequests map[peer.ID]string
}

func (p *PeerHandlerMock) InitiatePayment(target peer.ID, paymentRequest string) {
	p.paymentRequests[target] = paymentRequest
}

func (p *PeerHandlerMock) PaymentCommand(target peer.ID, commandId string, commandBody []byte, commandType int32, sessionId string) {
	panic("implement me")
}

func (p *PeerHandlerMock) PaymentResponse(target peer.ID, commandId string, commandReply []byte, sessionId string) {
	panic("implement me")
}

func (p *PeerHandlerMock) PaymentStatusResponse(target peer.ID, sessionId string, status bool) {
	panic("implement me")
}

func TestRequirePayment(t *testing.T) {

	msg := requirePayment {
		target: "TargetId",
		msgSize: 1024 * 1024,
	}

	debt := Debt{
		id:               "TargetId",
		validationQueue:  nil,
		requestedAmount:  0,
		transferredBytes: 49 * 1024 * 1024, // 52 Mega transferred
		receivedBytes:    0,
	}

	paymentMock := &PaymentHandlerMock{
		debtRegistry: map[peer.ID]*Debt{
			"TargetId": &debt,
		},
		requestedPaymentAmount: 0,
		paymentRequest:         "sampleRequestInJson",
	}

	peerMock := &PeerHandlerMock{
		paymentRequests: map[peer.ID]string{},
	}

	msg.handle(paymentMock, peerMock, nil)

	expectedAmount := 50 * 1024 * 1024

	if paymentMock.requestedPaymentAmount != expectedAmount {
		t.Errorf("Invalid amount")
	}

	if peerMock.paymentRequests["TargetId"] != "sampleRequestInJson" {
		t.Errorf("Invalid request")
	}

	if debt.transferredBytes != 0 {
		t.Errorf("Transferred bytes count not zero")
	}

	if debt.requestedAmount != expectedAmount {
		t.Errorf("Debt amount not equal to expected")
	}
}
