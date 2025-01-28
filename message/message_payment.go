package message

import (
	pb "github.com/ipfs/go-bitswap/message/pb"
	"paidpiper.com/payment-gateway/models"
	"paidpiper.com/payment-gateway/paymentmanager"
)

func FromProto(pbm pb.IsMessage_PaymentMessage) paymentmanager.PaymentData {
	switch v := pbm.(type) {
	case *pb.Message_InitiatePayment_:
		return &paymentmanager.InitiatePayment{
			PaymentRequest: v.InitiatePayment.PaymentRequest,
		}
	case *pb.Message_PaymentCommand_:
		return &paymentmanager.PaymentCommand{
			CommandId:   v.PaymentCommand.CommandId,
			CommandBody: v.PaymentCommand.CommandBody,
			CommandType: models.CommandType(v.PaymentCommand.CommandType),
			SessionId:   v.PaymentCommand.SessionId,
		}
	case *pb.Message_PaymentResponse_:
		return &paymentmanager.PaymentResponse{
			CommandId:    v.PaymentResponse.CommandId,
			CommandReply: v.PaymentResponse.CommandReply,
			SessionId:    v.PaymentResponse.SessionId,
		}
	case *pb.Message_PaymentStatusResponse_:
		return &paymentmanager.PaymentStatusResponse{
			SessionId: v.PaymentStatusResponse.SessionId,
			Status:    v.PaymentStatusResponse.Status,
		}

	}
	return nil
}
func ToProto(pd paymentmanager.PaymentData) pb.IsMessage_PaymentMessage {
	switch m := pd.(type) {
	case *paymentmanager.InitiatePayment:
		return &pb.Message_InitiatePayment_{
			InitiatePayment: &pb.Message_InitiatePayment{
				PaymentRequest: m.PaymentRequest,
			},
		}
	case *paymentmanager.PaymentCommand:
		return &pb.Message_PaymentCommand_{
			PaymentCommand: &pb.Message_PaymentCommand{
				CommandId:   m.CommandId,
				CommandBody: m.CommandBody,
				CommandType: int32(m.CommandType),
				SessionId:   m.SessionId,
			},
		}
	case *paymentmanager.PaymentResponse:
		return &pb.Message_PaymentResponse_{
			PaymentResponse: &pb.Message_PaymentResponse{
				CommandId:    m.CommandId,
				CommandReply: m.CommandReply,
				SessionId:    m.SessionId,
			},
		}
	case *paymentmanager.PaymentStatusResponse:
		return &pb.Message_PaymentStatusResponse_{
			PaymentStatusResponse: &pb.Message_PaymentStatusResponse{
				SessionId: m.SessionId,
				Status:    m.Status,
			},
		}
	}
	return nil
}

// BitSwapMessage is the basic interface for interacting building, encoding,
// and decoding messages sent on the BitSwap protocol.
type PaymentBitSwapMessage interface {
	SetPaymentData(data paymentmanager.PaymentData)
	GetPaymentData() paymentmanager.PaymentData
	HasPayment() bool
}

type implWithPayment struct {
	impl
	paymentData paymentmanager.PaymentData
}

func WithPayment(m *impl) BitSwapMessage {
	return &implWithPayment{
		impl:        *m,
		paymentData: nil,
	}
}

func (m *implWithPayment) Clone() BitSwapMessage {
	implClonse, ok := m.impl.Clone().(*impl)
	if !ok {
		panic("BitSwapMessage impl clone not valid")
	}
	msg := &implWithPayment{
		impl: *implClonse,
	}
	if msg.paymentData != nil {
		msg.paymentData = m.paymentData
	} else {
		msg.paymentData = nil
	}
	return msg
}
func (m *implWithPayment) Reset(full bool) {
	m.impl.Reset(full)
}

// NOTE: should change in proro implementation original name newMessageFromProto
func newMessageWithPaymentFromProto(pbm pb.Message) (BitSwapMessage, error) {
	m := newMsgWithPayment(pbm.Wantlist.Full)

	// Bitswap +
	if pbm.PaymentMessage != nil {
		m.paymentData = FromProto(pbm.PaymentMessage)
	}

	return m, nil
}
func (m *implWithPayment) ToPaymentProto(pbm *pb.Message) {
	if m.paymentData != nil {

		pbm.PaymentMessage = ToProto(m.paymentData)
	}

}
func (m *implWithPayment) ToProtoV0() *pb.Message {
	pbm := m.ToProtoV0()
	// Bitswap +
	m.ToPaymentProto(pbm)
	return pbm
}
func (m *implWithPayment) ToProtoV1() *pb.Message {
	pbm := m.ToProtoV1()
	// Bitswap +
	m.ToPaymentProto(pbm)
	return pbm
}
func (m *implWithPayment) Loggable() map[string]interface{} {
	blocks := make([]string, 0, len(m.blocks))
	for _, v := range m.blocks {
		blocks = append(blocks, v.Cid().String())
	}
	return map[string]interface{}{
		"blocks":      blocks,
		"wants":       m.Wantlist(),
		"paymentData": m.paymentData,
	}
}
func newMsgWithPayment(full bool) *implWithPayment {
	return &implWithPayment{
		impl:        *newMsg(full),
		paymentData: nil,
	}
}

func (m *implWithPayment) Empty() bool {
	return len(m.blocks) == 0 &&
		len(m.wantlist) == 0 &&
		len(m.blockPresences) == 0 &&
		m.paymentData == nil
}

func (m *implWithPayment) HasPayment() bool {
	return m.paymentData != nil
}

func (m *implWithPayment) PaymentDataMessage(data paymentmanager.PaymentData) {
	m.paymentData = data
}
