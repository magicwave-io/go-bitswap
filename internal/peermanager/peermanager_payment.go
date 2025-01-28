package peermanager

import (
	"context"

	cid "github.com/ipfs/go-cid"
	peer "github.com/libp2p/go-libp2p-core/peer"
	"paidpiper.com/payment-gateway/paymentmanager"
)

type PeerQueueWithPayment interface {
	AddBroadcastWantHaves([]cid.Cid)
	AddWants([]cid.Cid, []cid.Cid)
	AddCancels([]cid.Cid)
	ResponseReceived(ks []cid.Cid)
	Startup()
	Shutdown()
	SendPaymentDataMessage(data paymentmanager.PaymentData)
}

type PaymentPeerManager struct {
	PeerManager
}

// PeerQueueFactory provides a function that will create a PeerQueue.
type PaymentPeerQueueFactory func(ctx context.Context, p peer.ID) PeerQueueWithPayment

func NewPaymentPeerManager(ctx context.Context, createPeerQueue PaymentPeerQueueFactory, self peer.ID) *PaymentPeerManager {
	return &PaymentPeerManager{
		PeerManager: *New(ctx, func(ctx context.Context, p peer.ID) PeerQueue {
			return createPeerQueue(ctx, p)
		}, self),
	}

}

func (pm *PaymentPeerManager) SendPaymentDataMessage(target peer.ID, data paymentmanager.PaymentData) {
	pqi := pm.getOrCreateWithPayment(target)
	pqi.SendPaymentDataMessage(data)
}

func (pm *PaymentPeerManager) getOrCreate(p peer.ID) PeerQueue {
	pq, ok := pm.peerQueues[p]
	if !ok {
		pq = pm.createPeerQueue(pm.ctx, p)
		pq.Startup()
		pm.peerQueues[p] = pq
	}
	return pq
}
func (pm *PaymentPeerManager) getOrCreateWithPayment(p peer.ID) PeerQueueWithPayment {
	item, ok := pm.getOrCreate(p).(PeerQueueWithPayment)
	if !ok {
		panic("not posible")
	}
	return item
}
