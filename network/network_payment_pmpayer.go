package network

import (
	"context"

	bsmsg "github.com/ipfs/go-bitswap/message"

	"github.com/libp2p/go-libp2p-core/peer"
	bspaym "paidpiper.com/payment-gateway/paymentmanager"
)

type PaynetOption func(*paymentNetwork)

//Configures bitswap to use PPChannel
func PPChannelConfig(commandListenPort int, channelUrl string) PaynetOption {
	return func(bs *paymentNetwork) {
		bs.paym.SetHttpConnection(commandListenPort, channelUrl)
	}
}

type PeerManager interface {
	SendPaymentDataMessage(id peer.ID, data bspaym.PaymentData)
}

func WithPayment(ctx context.Context, network BitSwapNetwork, pm PeerManager, opts ...PaynetOption) BitSwapNetwork {

	pNet := &paymentNetwork{
		BitSwapNetwork: network,

		pm: pm,
	}
	pNet.paym = bspaym.New(ctx, pNet)
	for _, opt := range opts {
		opt(pNet)
	}
	return pNet
}

type paymentNetwork struct {
	BitSwapNetwork
	receiver Receiver
	paym     bspaym.PaymentManager
	pm       PeerManager
}

func (pm *paymentNetwork) SendPaymentDataMessage(id bspaym.PeerID, data bspaym.PaymentData) {
	peerID, err := peer.IDHexDecode(string(id))
	if err != nil {
		log.Fatal("Peer id id no valid")
		return
	}
	pm.pm.SendPaymentDataMessage(peerID, data)
}

// SendMessage sends a BitSwap message to a peer.
func (pm *paymentNetwork) SendMessage(
	ctx context.Context,
	id peer.ID,
	msg bsmsg.BitSwapMessage) error {

	err := pm.BitSwapNetwork.SendMessage(ctx, id, msg)
	peerID := peer.IDHexEncode(id)
	pm.paym.RegisterReceivedBytes(ctx, bspaym.PeerID(peerID), calculateSize(msg))
	return err
}

// SetDelegate registers the Reciver to handle messages received from the
// network.
func (pm *paymentNetwork) SetDelegate(receiver Receiver) {
	pm.receiver = receiver
	pm.BitSwapNetwork.SetDelegate(pm)
	pm.paym.Startup()
}

// Receiver
func (pm *paymentNetwork) ReceiveMessage(ctx context.Context, p peer.ID, incoming bsmsg.BitSwapMessage) {
	pm.receiver.ReceiveMessage(ctx, p, incoming)
	bytesCount := calculateSize(incoming)
	peerID := peer.IDHexEncode(p)
	pm.paym.RegisterReceivedBytes(ctx, bspaym.PeerID(peerID), bytesCount)
	incomPayMsg, ok := incoming.(bsmsg.PaymentBitSwapMessage)
	if !ok {
		return
	}
	paymentData := incomPayMsg.GetPaymentData()
	if paymentData != nil {

		pm.paym.ReceivePaymentDataMessage(ctx, bspaym.PeerID(peerID), paymentData)
	}
}

func (pm *paymentNetwork) ReceiveError(err error) {
	pm.receiver.ReceiveError(err)
}

// Connected/Disconnected warns bitswap about peer connections.
func (pm *paymentNetwork) PeerConnected(id peer.ID) {
	pm.receiver.PeerConnected(id)
}
func (pm *paymentNetwork) PeerDisconnected(id peer.ID) {
	pm.receiver.PeerDisconnected(id)
}

// Receiver end

func calculateSize(incoming bsmsg.BitSwapMessage) int {
	blocks := incoming.Blocks()
	var bytes int = 0
	if len(blocks) > 0 {
		// Do some accounting for each block
		for _, b := range blocks {
			blkLen := len(b.RawData())
			bytes += blkLen
		}
	}
	return bytes
}

// func (pm *PaymentNetwork) ConnectTo(ctx context.Context, id peer.ID) error {
// 	return pm.BitSwapNetwork.ConnectTo(ctx, id)
// }
// func (pm *PaymentNetwork) DisconnectFrom(ctx context.Context, id peer.ID) error {
// 	return pm.BitSwapNetwork.DisconnectFrom(ctx, id)
// }

// func (pm *PaymentNetwork) NewMessageSender(ctx context.Context, id peer.ID, opts *bsnet.MessageSenderOpts) (bsnet.MessageSender, error) {
// 	return pm.BitSwapNetwork.NewMessageSender(ctx, id, opts)
// }

// func (pm *PaymentNetwork) ConnectionManager() connmgr.ConnManager {
// 	return pm.BitSwapNetwork.ConnectionManager()
// }

// func (pm *PaymentNetwork) Stats() bsnet.Stats {
// 	return pm.BitSwapNetwork.Stats()
// }
