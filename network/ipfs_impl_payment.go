package network

import (
	"context"
	"fmt"
	"strings"

	bsmsg "github.com/ipfs/go-bitswap/message"
	"github.com/ipfs/go-bitswap/pptools/speedcontrol"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	ma "github.com/multiformats/go-multiaddr"
	boomdata "paidpiper.com/payment-gateway/boom/data"
	boomsvr "paidpiper.com/payment-gateway/boom/server"
	paymmodel "paidpiper.com/payment-gateway/models"
	bspaym "paidpiper.com/payment-gateway/paymentmanager"
)

type implWithPay struct {
	impl
	speedController *speedcontrol.MultiSpeedDetector
	payOption       PPOption

	pm       PeerManager
	paym     bspaym.PaymentManager
	receiver Receiver
	peers    map[peer.ID]struct{}
}
type PPOption struct {
	CommandListenPort int
	ChannelUrl        string
}

func NewPaymentFromIpfsHost(host host.Host, r routing.ContentRouting, payOption PPOption, opts ...NetOpt) BitSwapNetwork {
	bitswapNetwork := NewFromIpfsHost(host, r, opts...)
	implNetwork, ok := bitswapNetwork.(*impl)
	if ok {
		//implNetwork.host.Peerstore().Addrs()
		return &implWithPay{
			impl:            *implNetwork,
			speedController: speedcontrol.NewSpeedDetector(),
			payOption:       payOption,
		}
	}
	panic("not valid implementation")
}

func (bsnet *implWithPay) newPaymentStreamToPeer(ctx context.Context, id peer.ID) (network.Stream, error) {
	stream, err := bsnet.newStreamToPeer(ctx, id)
	if err != nil {
		return nil, err
	}
	return bsnet.speedController.WrapStream(id, stream), nil
}

func (bsnet *implWithPay) SendMessage(
	ctx context.Context,
	p peer.ID,
	outgoing bsmsg.BitSwapMessage) error {
	if s, ok := outgoing.(bsmsg.PaymentBitSwapMessage); ok {
		//fmt.Println("Send ", s.String())
		if s.Empty() {
			fmt.Println("Send Empty")
		}
	}
	s, err := bsnet.newPaymentStreamToPeer(ctx, p)
	if err != nil {
		fmt.Println("Stream not found")
		return err
	}

	if err = bsnet.msgToStream(ctx, s, outgoing, sendMessageTimeout); err != nil {
		_ = s.Reset()
		return err
	}

	err = s.Close()
	// Payment maneger
	peerID := peer.IDHexEncode(p)
	size := calculateSize(outgoing)
	if size > 0 {
		bsnet.paym.RequirePayment(ctx, paymmodel.PeerID(peerID), size)
	}
	// payment manager
	return err
}

type PeerManager interface {
	SendPaymentDataMessage(id peer.ID, data bspaym.PaymentData)
}

func WithPayment(ctx context.Context, network BitSwapNetwork, pm PeerManager, server bspaym.CallbackServer) BitSwapNetwork {
	if n, ok := network.(*implWithPay); ok {
		paym := bspaym.New(ctx, n)
		port := n.payOption.CommandListenPort
		server.SetPort(port)
		paym.SetHttpConnection(port, n.payOption.ChannelUrl, server)
		n.paym = paym
		n.pm = pm
		n.peers = make(map[peer.ID]struct{})
		boomsvr.AddConnectionSource(n)
	}
	return network
}

func (pm *implWithPay) Connections() ([]*boomdata.Connections, error) {
	cns := []*boomdata.Connections{}
	for _, addresses := range pm.PeersAddresses() {
		for _, address := range addresses {
			addr := address.String()
			if strings.HasPrefix(addr, "/onion") {
				cns = append(cns, &boomdata.Connections{
					Hosts: []string{address.String()},
				})
			}
		}
	}
	return cns, nil
}

func (pm *implWithPay) SendPaymentDataMessage(id paymmodel.PeerID, data bspaym.PaymentData) {
	peerID, err := peer.IDHexDecode(string(id))
	if err != nil {
		log.Fatal("Peer id id no valid")
		return
	}
	pm.pm.SendPaymentDataMessage(peerID, data)
}

// SetDelegate registers the Reciver to handle messages received from the
// network.
func (pm *implWithPay) SetDelegate(receiver Receiver) {
	pm.receiver = receiver
	pm.impl.SetDelegate(pm)
	pm.paym.Startup()
}

// Receiver
func (pm *implWithPay) ReceiveMessage(ctx context.Context, p peer.ID, incoming bsmsg.BitSwapMessage) {
	pm.receiver.ReceiveMessage(ctx, p, incoming)
	bytesCount := calculateSize(incoming)
	peerID := peer.IDHexEncode(p)
	pm.paym.RegisterReceivedBytes(ctx, paymmodel.PeerID(peerID), bytesCount)
	incomPayMsg, ok := incoming.(bsmsg.PaymentBitSwapMessage)
	if !ok {
		return
	}
	paymentData := incomPayMsg.GetPaymentData()
	if paymentData != nil {

		pm.paym.ReceivePaymentDataMessage(ctx, paymmodel.PeerID(peerID), paymentData)
	}
}

func (pm *impl) PeersAddresses() map[peer.ID][]ma.Multiaddr {
	return nil
}

func (pm *implWithPay) PeersAddresses() map[peer.ID][]ma.Multiaddr {
	addresses := make(map[peer.ID][]ma.Multiaddr)
	for _, c := range pm.host.Network().Conns() {
		pid := c.RemotePeer()
		addr := c.RemoteMultiaddr()
		addresses[pid] = []ma.Multiaddr{addr}
	}
	return addresses
}

func (pm *implWithPay) ReceiveError(err error) {
	pm.receiver.ReceiveError(err)
}

// Connected/Disconnected warns bitswap about peer connections.
func (pm *implWithPay) PeerConnected(id peer.ID) {
	pm.peers[id] = struct{}{}
	pm.receiver.PeerConnected(id)
}

func (pm *implWithPay) PeerDisconnected(id peer.ID) {
	delete(pm.peers, id)
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
