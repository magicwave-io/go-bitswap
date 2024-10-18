package message

import (
	"encoding/binary"
	"fmt"
	"io"

	pb "github.com/ipfs/go-bitswap/message/pb"
	"github.com/ipfs/go-bitswap/wantlist"
	blocks "github.com/ipfs/go-block-format"

	"github.com/ipfs/go-cid"
	pool "github.com/libp2p/go-buffer-pool"
	"github.com/libp2p/go-msgio"

	"github.com/libp2p/go-libp2p-core/network"
)

// BitSwapMessage is the basic interface for interacting building, encoding,
// and decoding messages sent on the BitSwap protocol.
type BitSwapMessage interface {
	// Wantlist returns a slice of unique keys that represent data wanted by
	// the sender.
	Wantlist() []Entry

	// Blocks returns a slice of unique blocks.
	Blocks() []blocks.Block

	// AddEntry adds an entry to the Wantlist.
	AddEntry(key cid.Cid, priority int)

	Cancel(key cid.Cid)

	Empty() bool

	// A full wantlist is an authoritative copy, a 'non-full' wantlist is a patch-set
	Full() bool

	AddBlock(blocks.Block)
	Exportable

	Loggable() map[string]interface{}

	// Bitswap +
	InitiatePayment(paymentRequest string)

	PaymentCommand(commandId string, commandBody []byte, commandType int32, sessionId string)

	PaymentResponse(commandId string, commandReply []byte, sessionId string)

	GetInitiatePayment() *InitiatePayment

	GetPaymentCommand() *PaymentCommand

	GetPaymentResponse() *PaymentResponse
}

// Exportable is an interface for structures than can be
// encoded in a bitswap protobuf.
type Exportable interface {
	ToProtoV0() *pb.Message
	ToProtoV1() *pb.Message
	ToNetV0(w io.Writer) error
	ToNetV1(w io.Writer) error
}

type impl struct {
	full     bool
	wantlist map[cid.Cid]*Entry
	blocks   map[cid.Cid]blocks.Block

	initiatePayment	*InitiatePayment
	paymentCommand	*PaymentCommand
	paymentResponse *PaymentResponse
}

type InitiatePayment struct {
	paymentRequest		string
}

func (i *InitiatePayment) GetPaymentRequest() string {
	return i.paymentRequest
}

type PaymentCommand struct {
	commandId 	string
	commandBody []byte
	commandType int32
	sessionId	string
}

func (i *PaymentCommand) GetCommandId() string {
	return i.commandId
}

func (i *PaymentCommand) GetCommandBody() []byte {
	return i.commandBody
}

func (i *PaymentCommand) GetCommandType() int32 {
	return i.commandType
}

func (i *PaymentCommand) GetSessionId() string  {
	return i.sessionId
}

type PaymentResponse struct {
	commandId		string
	commandReply 	[]byte
	sessionId		string
}

func (i *PaymentResponse) GetCommandId() string {
	return i.commandId
}

func (i *PaymentResponse) GetCommandReply() []byte {
	return i.commandReply
}

func (i *PaymentResponse) GetSessionId() string  {
	return i.sessionId
}

// New returns a new, empty bitswap message
func New(full bool) BitSwapMessage {
	return newMsg(full)
}

func newMsg(full bool) *impl {
	return &impl{
		blocks:   make(map[cid.Cid]blocks.Block),
		wantlist: make(map[cid.Cid]*Entry),
		full:     full,
	}
}

// Entry is an wantlist entry in a Bitswap message (along with whether it's an
// add or cancel).
type Entry struct {
	wantlist.Entry
	Cancel bool
}

func newMessageFromProto(pbm pb.Message) (BitSwapMessage, error) {
	m := newMsg(pbm.Wantlist.Full)
	for _, e := range pbm.Wantlist.Entries {
		c, err := cid.Cast(e.Block)
		if err != nil {
			return nil, fmt.Errorf("incorrectly formatted cid in wantlist: %s", err)
		}
		m.addEntry(c, int(e.Priority), e.Cancel)
	}

	// deprecated
	for _, d := range pbm.Blocks {
		// CIDv0, sha256, protobuf only
		b := blocks.NewBlock(d)
		m.AddBlock(b)
	}
	//

	for _, b := range pbm.GetPayload() {
		pref, err := cid.PrefixFromBytes(b.GetPrefix())
		if err != nil {
			return nil, err
		}

		c, err := pref.Sum(b.GetData())
		if err != nil {
			return nil, err
		}

		blk, err := blocks.NewBlockWithCid(b.GetData(), c)
		if err != nil {
			return nil, err
		}

		m.AddBlock(blk)
	}

	// Bitswap +
	if pbm.PaymentMessage != nil {
		switch paymentMessage := pbm.PaymentMessage.(type) {
		case *pb.Message_PaymentCommand_:
			m.paymentCommand = &PaymentCommand {
				commandId: paymentMessage.PaymentCommand.CommandId,
				commandBody: paymentMessage.PaymentCommand.CommandBody,
				commandType: paymentMessage.PaymentCommand.CommandType,
				sessionId: paymentMessage.PaymentCommand.SessionId,
			}
		case *pb.Message_PaymentResponse_:
			m.paymentResponse = &PaymentResponse {
				commandId: paymentMessage.PaymentResponse.CommandId,
				commandReply: paymentMessage.PaymentResponse.CommandReply,
				sessionId: paymentMessage.PaymentResponse.SessionId,
			}
		case *pb.Message_InitiatePayment_:
			m.initiatePayment = &InitiatePayment {
				paymentRequest: paymentMessage.InitiatePayment.PaymentRequest,
			}
		}
	}

	return m, nil
}

func (m *impl) Full() bool {
	return m.full
}

func (m *impl) Empty() bool {
	return len(m.blocks) == 0 && len(m.wantlist) == 0 && m.initiatePayment == nil && m.paymentCommand == nil && m.paymentResponse == nil
}

func (m *impl) Wantlist() []Entry {
	out := make([]Entry, 0, len(m.wantlist))
	for _, e := range m.wantlist {
		out = append(out, *e)
	}
	return out
}


func (m *impl) GetInitiatePayment() *InitiatePayment {
	return m.initiatePayment
}

func (m *impl) GetPaymentCommand() *PaymentCommand {
	return m.paymentCommand
}

func (m *impl) GetPaymentResponse() *PaymentResponse {
	return m.paymentResponse
}

func (m *impl) InitiatePayment(paymentRequest string) {
	m.initiatePayment = &InitiatePayment{
		paymentRequest,
	}
}

func (m *impl) PaymentCommand(commandId string, commandBody []byte, commandType int32, sessionId string) {
	m.paymentCommand = &PaymentCommand{
		commandId,
		commandBody,
		commandType,
		sessionId,
	}
}

func (m *impl) PaymentResponse(commandId string, commandReply []byte, sessionId string) {
	m.paymentResponse = &PaymentResponse{
		commandId,
		commandReply,
		sessionId,
	}
}

func (m *impl) Blocks() []blocks.Block {
	bs := make([]blocks.Block, 0, len(m.blocks))
	for _, block := range m.blocks {
		bs = append(bs, block)
	}
	return bs
}

func (m *impl) Cancel(k cid.Cid) {
	delete(m.wantlist, k)
	m.addEntry(k, 0, true)
}

func (m *impl) AddEntry(k cid.Cid, priority int) {
	m.addEntry(k, priority, false)
}

func (m *impl) addEntry(c cid.Cid, priority int, cancel bool) {
	e, exists := m.wantlist[c]
	if exists {
		e.Priority = priority
		e.Cancel = cancel
	} else {
		m.wantlist[c] = &Entry{
			Entry: wantlist.Entry{
				Cid:      c,
				Priority: priority,
			},
			Cancel: cancel,
		}
	}
}

func (m *impl) AddBlock(b blocks.Block) {
	m.blocks[b.Cid()] = b
}

// FromNet generates a new BitswapMessage from incoming data on an io.Reader.
func FromNet(r io.Reader) (BitSwapMessage, error) {
	reader := msgio.NewVarintReaderSize(r, network.MessageSizeMax)
	return FromMsgReader(reader)
}

// FromPBReader generates a new Bitswap message from a gogo-protobuf reader
func FromMsgReader(r msgio.Reader) (BitSwapMessage, error) {
	msg, err := r.ReadMsg()
	if err != nil {
		return nil, err
	}

	var pbm pb.Message
	err = pbm.Unmarshal(msg)
	r.ReleaseMsg(msg)
	if err != nil {
		return nil, err
	}

	return newMessageFromProto(pbm)
}

func (m *impl) ToProtoV0() *pb.Message {
	pbm := new(pb.Message)
	pbm.Wantlist.Entries = make([]pb.Message_Wantlist_Entry, 0, len(m.wantlist))
	for _, e := range m.wantlist {
		pbm.Wantlist.Entries = append(pbm.Wantlist.Entries, pb.Message_Wantlist_Entry{
			Block:    e.Cid.Bytes(),
			Priority: int32(e.Priority),
			Cancel:   e.Cancel,
		})
	}
	pbm.Wantlist.Full = m.full

	messageBlocks := m.Blocks()
	pbm.Blocks = make([][]byte, 0, len(messageBlocks))
	for _, b := range messageBlocks {
		pbm.Blocks = append(pbm.Blocks, b.RawData())
	}

	// Bitswap +
	m.ToPaymentProto(pbm)

	return pbm
}

func (m * impl) ToPaymentProto(pbm *pb.Message) {
	if m.initiatePayment != nil {
		pbm.PaymentMessage = &pb.Message_InitiatePayment_ {
			InitiatePayment: &pb.Message_InitiatePayment {
				PaymentRequest: m.initiatePayment.paymentRequest,
			},
		}
	}

	if m.paymentCommand != nil {
		pbm.PaymentMessage = &pb.Message_PaymentCommand_{
			PaymentCommand: &pb.Message_PaymentCommand{
				CommandId: m.paymentCommand.commandId,
				CommandBody: m.paymentCommand.commandBody,
				CommandType: m.paymentCommand.commandType,
				SessionId:   m.paymentCommand.sessionId,
			},
		}
	}

	if m.paymentResponse != nil {
		pbm.PaymentMessage = &pb.Message_PaymentResponse_{
			PaymentResponse: &pb.Message_PaymentResponse{
				CommandId: m.paymentResponse.commandId,
				CommandReply: m.paymentResponse.commandReply,
				SessionId:   m.paymentResponse.sessionId,
			},
		}
	}
}
func (m *impl) ToProtoV1() *pb.Message {
	pbm := new(pb.Message)
	pbm.Wantlist.Entries = make([]pb.Message_Wantlist_Entry, 0, len(m.wantlist))
	for _, e := range m.wantlist {
		pbm.Wantlist.Entries = append(pbm.Wantlist.Entries, pb.Message_Wantlist_Entry{
			Block:    e.Cid.Bytes(),
			Priority: int32(e.Priority),
			Cancel:   e.Cancel,
		})
	}
	pbm.Wantlist.Full = m.full

	messageBlocks := m.Blocks()
	pbm.Payload = make([]pb.Message_Block, 0, len(messageBlocks))
	for _, b := range messageBlocks {
		pbm.Payload = append(pbm.Payload, pb.Message_Block{
			Data:   b.RawData(),
			Prefix: b.Cid().Prefix().Bytes(),
		})
	}

	// Bitswap +
	m.ToPaymentProto(pbm)

	return pbm
}

func (m *impl) ToNetV0(w io.Writer) error {
	return write(w, m.ToProtoV0())
}

func (m *impl) ToNetV1(w io.Writer) error {
	return write(w, m.ToProtoV1())
}

func write(w io.Writer, m *pb.Message) error {
	size := m.Size()

	buf := pool.Get(size + binary.MaxVarintLen64)
	defer pool.Put(buf)

	n := binary.PutUvarint(buf, uint64(size))

	written, err := m.MarshalTo(buf[n:])
	if err != nil {
		return err
	}
	n += written

	_, err = w.Write(buf[:n])
	return err
}

func (m *impl) Loggable() map[string]interface{} {
	messageBlocks := make([]string, 0, len(m.blocks))
	for _, v := range m.blocks {
		messageBlocks = append(messageBlocks, v.Cid().String())
	}

	return map[string]interface{}{
		"blocks": messageBlocks,
		"wants":  m.Wantlist(),
		"initiatePayment": m.initiatePayment,
		"paymentCommand": m.paymentCommand,
		"paymentResponse": m.paymentResponse,
	}
}
