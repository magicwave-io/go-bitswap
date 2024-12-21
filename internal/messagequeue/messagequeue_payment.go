package messagequeue

import (
	"time"

	bsmsg "github.com/ipfs/go-bitswap/message"
	"paidpiper.com/payment-gateway/paymentmanager"
)

func WithMessageCutoff(mq *MessageQueue, sendMessageCutoff int) *MessageQueue {
	mq.sendMessageCutoff = sendMessageCutoff
	return mq
}

func WithPayment(mq *MessageQueue) *MessageQueueWithPayment {
	return &MessageQueueWithPayment{
		MessageQueue: *mq,
	}
}

type MessageQueueWithPayment struct {
	MessageQueue
}

// Bitswap ++

func (mq *MessageQueueWithPayment) SendPaymentDataMessage(data paymentmanager.PaymentData) {
	mq.msg.(bsmsg.PaymentBitSwapMessage).SetPaymentData(data)
	select {

	case mq.outgoingWork <- time.Now():

	default:
	}
}
