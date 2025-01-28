package network

import (
	"context"
	"sync/atomic"

	bsmsg "github.com/ipfs/go-bitswap/message"
	"github.com/ipfs/go-bitswap/speedcontrol"
	"github.com/libp2p/go-libp2p-core/helpers"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
)

type implWithPay struct {
	impl
	speedController *speedcontrol.MultiSpeedDetector
}

func NewPaymentFromIpfsHost(host host.Host, r routing.ContentRouting, opts ...NetOpt) BitSwapNetwork {
	bitswapNetwork := NewFromIpfsHost(host, r, opts...)
	implNetwork, ok := bitswapNetwork.(*impl)
	if ok {
		return &implWithPay{
			impl:            *implNetwork,
			speedController: speedcontrol.NewSpeedDetector(),
		}
	}
	panic("not valid implementation")
}

func (bsnet *implWithPay) newPaymentStreamToPeer(ctx context.Context, id peer.ID) (network.Stream, error) {
	stream, err := bsnet.newPaymentStreamToPeer(ctx, id)
	if err != nil {
		return nil, err
	}
	return bsnet.speedController.WrapStream(id, stream), nil
}

func (bsnet *implWithPay) SendMessage(
	ctx context.Context,
	p peer.ID,
	outgoing bsmsg.BitSwapMessage) error {

	s, err := bsnet.newPaymentStreamToPeer(ctx, p)
	if err != nil {
		return err
	}

	if err = bsnet.msgToStream(ctx, s, outgoing, sendMessageTimeout); err != nil {
		_ = s.Reset()
		return err
	}
	atomic.AddUint64(&bsnet.stats.MessagesSent, 1)

	// TODO(https://github.com/libp2p/go-libp2p-net/issues/28): Avoid this goroutine.
	//nolint
	helpers.MultistreamSemverMatcher()
	go helpers.AwaitEOF(s)
	return s.Close()
}
