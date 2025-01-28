package network_test

import (
	"context"
	"testing"
	"time"

	bsmsg "github.com/ipfs/go-bitswap/message"
	tn "github.com/ipfs/go-bitswap/testnet"
	blocksutil "github.com/ipfs/go-ipfs-blocksutil"

	"github.com/libp2p/go-libp2p-core/peer"
	tnet "github.com/libp2p/go-libp2p-testing/net"
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
)

// Receiver is an interface for receiving messages from the GraphSyncNetwork.
type receiver struct {
	peers           map[peer.ID]struct{}
	messageReceived chan struct{}
	connectionEvent chan struct{}
	lastMessage     bsmsg.BitSwapMessage
	lastSender      peer.ID
}

func (r *receiver) ReceiveMessage(
	ctx context.Context,
	sender peer.ID,
	incoming bsmsg.BitSwapMessage) {
	r.lastSender = sender
	r.lastMessage = incoming
	select {
	case <-ctx.Done():
	case r.messageReceived <- struct{}{}:
	}
}

func (r *receiver) ReceiveError(err error) {
}

func (r *receiver) PeerConnected(p peer.ID) {
	r.peers[p] = struct{}{}
	select {
	case r.connectionEvent <- struct{}{}:
	}
}

func (r *receiver) PeerDisconnected(p peer.ID) {
	delete(r.peers, p)
	select {
	case r.connectionEvent <- struct{}{}:
	}
}
func TestMessageSendAndReceive(t *testing.T) {
	// create network
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	mn := mocknet.New(ctx)
	mr := tn.NewMockServer()
	streamNet, err := tn.StreamNet(ctx, mn, mr)
	if err != nil {
		t.Fatal("Unable to setup network")
	}
	p1 := tnet.RandIdentityOrFatal(t)
	p2 := tnet.RandIdentityOrFatal(t)

	bsnet1 := streamNet.Adapter(p1)
	bsnet2 := streamNet.Adapter(p2)
	r1 := &receiver{
		peers:           make(map[peer.ID]struct{}),
		messageReceived: make(chan struct{}),
		connectionEvent: make(chan struct{}, 1),
	}
	r2 := &receiver{
		peers:           make(map[peer.ID]struct{}),
		messageReceived: make(chan struct{}),
		connectionEvent: make(chan struct{}, 1),
	}
	bsnet1.SetDelegate(r1)
	bsnet2.SetDelegate(r2)

	mn.LinkAll()
	bsnet1.ConnectTo(ctx, p2.ID())
	select {
	case <-ctx.Done():
		t.Fatal("did not connect peer")
	case <-r1.connectionEvent:
	}
	bsnet2.ConnectTo(ctx, p1.ID())
	select {
	case <-ctx.Done():
		t.Fatal("did not connect peer")
	case <-r2.connectionEvent:
	}
	if _, ok := r1.peers[p2.ID()]; !ok {
		t.Fatal("did to connect to correct peer")
	}
	if _, ok := r2.peers[p1.ID()]; !ok {
		t.Fatal("did to connect to correct peer")
	}
	blockGenerator := blocksutil.NewBlockGenerator()
	block1 := blockGenerator.Next()
	block2 := blockGenerator.Next()
	sent := bsmsg.New(false)
	sent.AddEntry(block1.Cid(), 1)
	sent.AddBlock(block2)

	bsnet1.SendMessage(ctx, p2.ID(), sent)

	select {
	case <-ctx.Done():
		t.Fatal("did not receive message sent")
	case <-r2.messageReceived:
	}

	sender := r2.lastSender
	if sender != p1.ID() {
		t.Fatal("received message from wrong node")
	}

	received := r2.lastMessage

	sentWants := sent.Wantlist()
	if len(sentWants) != 1 {
		t.Fatal("Did not add want to sent message")
	}
	sentWant := sentWants[0]
	receivedWants := received.Wantlist()
	if len(receivedWants) != 1 {
		t.Fatal("Did not add want to received message")
	}
	receivedWant := receivedWants[0]
	if receivedWant.Cid != sentWant.Cid ||
		receivedWant.Priority != receivedWant.Priority ||
		receivedWant.Cancel != receivedWant.Cancel {
		t.Fatal("Sent message wants did not match received message wants")
	}
	sentBlocks := sent.Blocks()
	if len(sentBlocks) != 1 {
		t.Fatal("Did not add block to sent message")
	}
	sentBlock := sentBlocks[0]
	receivedBlocks := received.Blocks()
	if len(receivedBlocks) != 1 {
		t.Fatal("Did not add response to received message")
	}
	receivedBlock := receivedBlocks[0]
	if receivedBlock.Cid() != sentBlock.Cid() {
		t.Fatal("Sent message blocks did not match received message blocks")
	}
}

func TestRealProviding(t *testing.T) {
	// create network
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	mn := mocknet.New(ctx)
	routing := tn.NewDHTServer()
	streamNet, err := tn.StreamNet(ctx, mn, routing)
	if err != nil {
		t.Fatal("Unable to setup network")
	}
	p1 := tnet.RandIdentityOrFatal(t)
	p2 := tnet.RandIdentityOrFatal(t)
	p3 := tnet.RandIdentityOrFatal(t)

	bsnet1 := streamNet.Adapter(p1)
	bsnet2 := streamNet.Adapter(p2)
	bsnet3 := streamNet.Adapter(p3)

	r1 := &receiver{
		peers:           make(map[peer.ID]struct{}),
		messageReceived: make(chan struct{}),
		connectionEvent: make(chan struct{}, 1),
	}
	r2 := &receiver{
		peers:           make(map[peer.ID]struct{}),
		messageReceived: make(chan struct{}),
		connectionEvent: make(chan struct{}, 1),
	}
	r3 := &receiver{
		peers:           make(map[peer.ID]struct{}),
		messageReceived: make(chan struct{}),
		connectionEvent: make(chan struct{}, 1),
	}
	bsnet1.SetDelegate(r1)
	bsnet2.SetDelegate(r2)
	bsnet3.SetDelegate(r3)

	mn.LinkAll()
	bsnet1.ConnectTo(ctx, p2.ID())
	select {
	case <-ctx.Done():
		t.Fatal("did not connect peer")
	case <-r1.connectionEvent:
	}
	select {
	case <-ctx.Done():
		t.Fatal("did not connect peer")
	case <-r2.connectionEvent:
	}
	bsnet2.ConnectTo(ctx, p3.ID())
	select {
	case <-ctx.Done():
		t.Fatal("did not connect peer")
	case <-r2.connectionEvent:
	}
	select {
	case <-ctx.Done():
		t.Fatal("did not connect peer")
	case <-r3.connectionEvent:
	}
	if _, ok := r1.peers[p2.ID()]; !ok {
		t.Fatal("did to connect to correct peer")
	}
	if _, ok := r2.peers[p1.ID()]; !ok {
		t.Fatal("did to connect to correct peer")
	}
	if _, ok := r2.peers[p3.ID()]; !ok {
		t.Fatal("did to connect to correct peer")
	}
	if _, ok := r3.peers[p2.ID()]; !ok {
		t.Fatal("did to connect to correct peer")
	}
	blockGenerator := blocksutil.NewBlockGenerator()
	block1 := blockGenerator.Next()
	// provide on node 3
	// DHT can setup routing table in seperate thread,
	// meaning it's not finished even bitswap peers are connected
	// so attempt a few times with delay
provideLoop:
	for i := 0; i < 5; i++ {
		err = bsnet3.Provide(ctx, block1.Cid())
		if err == nil {
			break provideLoop
		}
		time.Sleep(10 * time.Millisecond)
	}
	if err != nil {
		t.Fatalf("Unable to provide node: %s", err)
	}
	incomingProvders := bsnet1.FindProvidersAsync(ctx, block1.Cid(), 1)
	select {
	case <-ctx.Done():
		t.Fatal("Unable to find provided block")
	case prov := <-incomingProvders:
		if prov != p3.ID() {
			t.Fatal("Found incorrect provider")
		}
	}
}
