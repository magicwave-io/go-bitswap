// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: message.proto

package bitswap_message_pb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Message_BlockPresenceType int32

const (
	Message_Have     Message_BlockPresenceType = 0
	Message_DontHave Message_BlockPresenceType = 1
)

var Message_BlockPresenceType_name = map[int32]string{
	0: "Have",
	1: "DontHave",
}

var Message_BlockPresenceType_value = map[string]int32{
	"Have":     0,
	"DontHave": 1,
}

func (x Message_BlockPresenceType) String() string {
	return proto.EnumName(Message_BlockPresenceType_name, int32(x))
}

func (Message_BlockPresenceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0}
}

type Message_Wantlist_WantType int32

const (
	Message_Wantlist_Block Message_Wantlist_WantType = 0
	Message_Wantlist_Have  Message_Wantlist_WantType = 1
)

var Message_Wantlist_WantType_name = map[int32]string{
	0: "Block",
	1: "Have",
}

var Message_Wantlist_WantType_value = map[string]int32{
	"Block": 0,
	"Have":  1,
}

func (x Message_Wantlist_WantType) String() string {
	return proto.EnumName(Message_Wantlist_WantType_name, int32(x))
}

func (Message_Wantlist_WantType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0, 0}
}

type Message struct {
	Wantlist       Message_Wantlist        `protobuf:"bytes,1,opt,name=wantlist,proto3" json:"wantlist"`
	Blocks         [][]byte                `protobuf:"bytes,2,rep,name=blocks,proto3" json:"blocks,omitempty"`
	Payload        []Message_Block         `protobuf:"bytes,3,rep,name=payload,proto3" json:"payload"`
	BlockPresences []Message_BlockPresence `protobuf:"bytes,4,rep,name=blockPresences,proto3" json:"blockPresences"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetWantlist() Message_Wantlist {
	if m != nil {
		return m.Wantlist
	}
	return Message_Wantlist{}
}

func (m *Message) GetBlocks() [][]byte {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *Message) GetPayload() []Message_Block {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetBlockPresences() []Message_BlockPresence {
	if m != nil {
		return m.BlockPresences
	}
	return nil
}

type Message_Wantlist struct {
	Entries []Message_Wantlist_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries"`
	Full    bool                     `protobuf:"varint,2,opt,name=full,proto3" json:"full,omitempty"`
}

func (m *Message_Wantlist) Reset()         { *m = Message_Wantlist{} }
func (m *Message_Wantlist) String() string { return proto.CompactTextString(m) }
func (*Message_Wantlist) ProtoMessage()    {}
func (*Message_Wantlist) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0}
}
func (m *Message_Wantlist) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message_Wantlist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message_Wantlist.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message_Wantlist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_Wantlist.Merge(m, src)
}
func (m *Message_Wantlist) XXX_Size() int {
	return m.Size()
}
func (m *Message_Wantlist) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_Wantlist.DiscardUnknown(m)
}

var xxx_messageInfo_Message_Wantlist proto.InternalMessageInfo

func (m *Message_Wantlist) GetEntries() []Message_Wantlist_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *Message_Wantlist) GetFull() bool {
	if m != nil {
		return m.Full
	}
	return false
}

type Message_Wantlist_Entry struct {
	Block        []byte                    `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Priority     int32                     `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	Cancel       bool                      `protobuf:"varint,3,opt,name=cancel,proto3" json:"cancel,omitempty"`
	WantType     Message_Wantlist_WantType `protobuf:"varint,4,opt,name=want_type,json=wantType,proto3,enum=bitswap.message.pb.Message_Wantlist_WantType" json:"want_type,omitempty"`
	SendDontHave bool                      `protobuf:"varint,5,opt,name=send_dont_have,json=sendDontHave,proto3" json:"send_dont_have,omitempty"`
}

func (m *Message_Wantlist_Entry) Reset()         { *m = Message_Wantlist_Entry{} }
func (m *Message_Wantlist_Entry) String() string { return proto.CompactTextString(m) }
func (*Message_Wantlist_Entry) ProtoMessage()    {}
func (*Message_Wantlist_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0, 0}
}
func (m *Message_Wantlist_Entry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message_Wantlist_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message_Wantlist_Entry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message_Wantlist_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_Wantlist_Entry.Merge(m, src)
}
func (m *Message_Wantlist_Entry) XXX_Size() int {
	return m.Size()
}
func (m *Message_Wantlist_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_Wantlist_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Message_Wantlist_Entry proto.InternalMessageInfo

func (m *Message_Wantlist_Entry) GetBlock() []byte {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *Message_Wantlist_Entry) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Message_Wantlist_Entry) GetCancel() bool {
	if m != nil {
		return m.Cancel
	}
	return false
}

func (m *Message_Wantlist_Entry) GetWantType() Message_Wantlist_WantType {
	if m != nil {
		return m.WantType
	}
	return Message_Wantlist_Block
}

func (m *Message_Wantlist_Entry) GetSendDontHave() bool {
	if m != nil {
		return m.SendDontHave
	}
	return false
}

type Message_Block struct {
	Prefix []byte `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Message_Block) Reset()         { *m = Message_Block{} }
func (m *Message_Block) String() string { return proto.CompactTextString(m) }
func (*Message_Block) ProtoMessage()    {}
func (*Message_Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 1}
}
func (m *Message_Block) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message_Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message_Block.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message_Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_Block.Merge(m, src)
}
func (m *Message_Block) XXX_Size() int {
	return m.Size()
}
func (m *Message_Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Message_Block proto.InternalMessageInfo

func (m *Message_Block) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *Message_Block) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Message_BlockPresence struct {
	Cid  []byte                    `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	Type Message_BlockPresenceType `protobuf:"varint,2,opt,name=type,proto3,enum=bitswap.message.pb.Message_BlockPresenceType" json:"type,omitempty"`
}

func (m *Message_BlockPresence) Reset()         { *m = Message_BlockPresence{} }
func (m *Message_BlockPresence) String() string { return proto.CompactTextString(m) }
func (*Message_BlockPresence) ProtoMessage()    {}
func (*Message_BlockPresence) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 2}
}
func (m *Message_BlockPresence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message_BlockPresence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message_BlockPresence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message_BlockPresence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_BlockPresence.Merge(m, src)
}
func (m *Message_BlockPresence) XXX_Size() int {
	return m.Size()
}
func (m *Message_BlockPresence) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_BlockPresence.DiscardUnknown(m)
}

var xxx_messageInfo_Message_BlockPresence proto.InternalMessageInfo

func (m *Message_BlockPresence) GetCid() []byte {
	if m != nil {
		return m.Cid
	}
	return nil
}

func (m *Message_BlockPresence) GetType() Message_BlockPresenceType {
	if m != nil {
		return m.Type
	}
	return Message_Have
}

func init() {
	proto.RegisterEnum("bitswap.message.pb.Message_BlockPresenceType", Message_BlockPresenceType_name, Message_BlockPresenceType_value)
	proto.RegisterEnum("bitswap.message.pb.Message_Wantlist_WantType", Message_Wantlist_WantType_name, Message_Wantlist_WantType_value)
	proto.RegisterType((*Message)(nil), "bitswap.message.pb.Message")
	proto.RegisterType((*Message_Wantlist)(nil), "bitswap.message.pb.Message.Wantlist")
	proto.RegisterType((*Message_Wantlist_Entry)(nil), "bitswap.message.pb.Message.Wantlist.Entry")
	proto.RegisterType((*Message_Block)(nil), "bitswap.message.pb.Message.Block")
	proto.RegisterType((*Message_BlockPresence)(nil), "bitswap.message.pb.Message.BlockPresence")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x86, 0xb5, 0xb6, 0x6c, 0x2b, 0x13, 0xc7, 0xb8, 0x4b, 0x29, 0x8b, 0x0e, 0x8a, 0x6b, 0x72,
	0x50, 0x5b, 0xa2, 0x80, 0xf3, 0x04, 0x31, 0x6d, 0x29, 0x81, 0x42, 0x11, 0x85, 0x1c, 0xcd, 0x4a,
	0x5a, 0x3b, 0xa2, 0x8a, 0x56, 0x48, 0xeb, 0xa4, 0x7a, 0x8b, 0xbe, 0x42, 0x9f, 0xa5, 0x97, 0x1c,
	0xd3, 0x5b, 0x4f, 0xa5, 0xd8, 0x2f, 0x52, 0x76, 0xb4, 0x32, 0xa4, 0x81, 0x26, 0xb7, 0xf9, 0x57,
	0xf3, 0x7f, 0x3b, 0xff, 0x2c, 0x82, 0x83, 0x2b, 0x51, 0x55, 0x7c, 0x25, 0x82, 0xa2, 0x94, 0x4a,
	0x52, 0x1a, 0xa5, 0xaa, 0xba, 0xe1, 0x45, 0xb0, 0x3b, 0x8e, 0xdc, 0xe3, 0x55, 0xaa, 0x2e, 0xd7,
	0x51, 0x10, 0xcb, 0xab, 0x93, 0x95, 0x5c, 0xc9, 0x13, 0x6c, 0x8d, 0xd6, 0x4b, 0x54, 0x28, 0xb0,
	0x6a, 0x10, 0xd3, 0xef, 0x7d, 0x18, 0x7c, 0x6c, 0xdc, 0xf4, 0x3d, 0x38, 0x37, 0x3c, 0x57, 0x59,
	0x5a, 0x29, 0x46, 0x26, 0xc4, 0xdf, 0x9f, 0x1d, 0x05, 0x0f, 0x6f, 0x08, 0x4c, 0x7b, 0x70, 0x61,
	0x7a, 0xe7, 0xf6, 0xed, 0xef, 0x43, 0x2b, 0xdc, 0x79, 0xe9, 0x0b, 0xe8, 0x47, 0x99, 0x8c, 0xbf,
	0x54, 0xac, 0x33, 0xe9, 0xfa, 0xc3, 0xd0, 0x28, 0x7a, 0x06, 0x83, 0x82, 0xd7, 0x99, 0xe4, 0x09,
	0xeb, 0x4e, 0xba, 0xfe, 0xfe, 0xec, 0xe5, 0xff, 0xf0, 0x73, 0x6d, 0x32, 0xec, 0xd6, 0x47, 0x2f,
	0x60, 0x84, 0xb0, 0x4f, 0xa5, 0xa8, 0x44, 0x1e, 0x8b, 0x8a, 0xd9, 0x48, 0x7a, 0xf5, 0x28, 0xa9,
	0x75, 0x18, 0xe2, 0x3f, 0x18, 0xf7, 0x67, 0x07, 0x9c, 0x36, 0x10, 0x3d, 0x87, 0x81, 0xc8, 0x55,
	0x99, 0x8a, 0x8a, 0x11, 0xc4, 0xbf, 0x7e, 0xca, 0x1e, 0x82, 0x77, 0xb9, 0x2a, 0xeb, 0x76, 0x62,
	0x03, 0xa0, 0x14, 0xec, 0xe5, 0x3a, 0xcb, 0x58, 0x67, 0x42, 0x7c, 0x27, 0xc4, 0xda, 0xfd, 0x41,
	0xa0, 0x87, 0xcd, 0xf4, 0x39, 0xf4, 0x70, 0x10, 0xdc, 0xf7, 0x30, 0x6c, 0x04, 0x75, 0xc1, 0x29,
	0xca, 0x54, 0x96, 0xa9, 0xaa, 0xd1, 0xd7, 0x0b, 0x77, 0x5a, 0x2f, 0x37, 0xe6, 0x79, 0x2c, 0x32,
	0xd6, 0x45, 0xa2, 0x51, 0xf4, 0x1c, 0xf6, 0xf4, 0x03, 0x2c, 0x54, 0x5d, 0x08, 0x66, 0x4f, 0x88,
	0x3f, 0x9a, 0x1d, 0x3f, 0x69, 0x6a, 0x5d, 0x7c, 0xae, 0x0b, 0xd1, 0x3c, 0xa0, 0xae, 0xe8, 0x11,
	0x8c, 0x2a, 0x91, 0x27, 0x8b, 0x44, 0xe6, 0x6a, 0x71, 0xc9, 0xaf, 0x05, 0xeb, 0xe1, 0x5d, 0x43,
	0x7d, 0xfa, 0x56, 0xe6, 0xea, 0x03, 0xbf, 0x16, 0xd3, 0xc3, 0x66, 0x63, 0xe8, 0xd8, 0x83, 0x1e,
	0x6e, 0x79, 0x6c, 0x51, 0x07, 0x6c, 0xfd, 0x79, 0x4c, 0xdc, 0x53, 0x73, 0xa8, 0x67, 0x2e, 0x4a,
	0xb1, 0x4c, 0xbf, 0x9a, 0x98, 0x46, 0xe9, 0xdd, 0x24, 0x5c, 0x71, 0xcc, 0x38, 0x0c, 0xb1, 0x76,
	0x13, 0x38, 0xb8, 0xf7, 0x5e, 0x74, 0x0c, 0xdd, 0x38, 0x4d, 0x8c, 0x53, 0x97, 0xf4, 0x0c, 0x6c,
	0x4c, 0xd9, 0x79, 0x3c, 0xe5, 0x3d, 0x14, 0xa6, 0x44, 0xeb, 0xf4, 0x0d, 0x3c, 0x7b, 0xf0, 0x69,
	0x37, 0xb9, 0x45, 0x87, 0xe0, 0xb4, 0x31, 0xc7, 0x64, 0xce, 0x6e, 0x37, 0x1e, 0xb9, 0xdb, 0x78,
	0xe4, 0xcf, 0xc6, 0x23, 0xdf, 0xb6, 0x9e, 0x75, 0xb7, 0xf5, 0xac, 0x5f, 0x5b, 0xcf, 0x8a, 0xfa,
	0xf8, 0x13, 0x9d, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x55, 0xec, 0x08, 0x98, 0x03, 0x00,
	0x00,
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BlockPresences) > 0 {
		for iNdEx := len(m.BlockPresences) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BlockPresences[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMessage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Payload) > 0 {
		for iNdEx := len(m.Payload) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Payload[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMessage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Blocks) > 0 {
		for iNdEx := len(m.Blocks) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Blocks[iNdEx])
			copy(dAtA[i:], m.Blocks[iNdEx])
			i = encodeVarintMessage(dAtA, i, uint64(len(m.Blocks[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Wantlist.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Message_Wantlist) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message_Wantlist) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_Wantlist) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Full {
		i--
		if m.Full {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Entries) > 0 {
		for iNdEx := len(m.Entries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Entries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMessage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Message_Wantlist_Entry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message_Wantlist_Entry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_Wantlist_Entry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SendDontHave {
		i--
		if m.SendDontHave {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.WantType != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.WantType))
		i--
		dAtA[i] = 0x20
	}
	if m.Cancel {
		i--
		if m.Cancel {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.Priority != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Priority))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Block) > 0 {
		i -= len(m.Block)
		copy(dAtA[i:], m.Block)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Block)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Message_Block) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message_Block) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_Block) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Prefix) > 0 {
		i -= len(m.Prefix)
		copy(dAtA[i:], m.Prefix)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Prefix)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Message_BlockPresence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message_BlockPresence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_BlockPresence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Cid) > 0 {
		i -= len(m.Cid)
		copy(dAtA[i:], m.Cid)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Cid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Wantlist.Size()
	n += 1 + l + sovMessage(uint64(l))
	if len(m.Blocks) > 0 {
		for _, b := range m.Blocks {
			l = len(b)
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	if len(m.Payload) > 0 {
		for _, e := range m.Payload {
			l = e.Size()
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	if len(m.BlockPresences) > 0 {
		for _, e := range m.BlockPresences {
			l = e.Size()
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	return n
}

func (m *Message_Wantlist) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	if m.Full {
		n += 2
	}
	return n
}

func (m *Message_Wantlist_Entry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Block)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.Priority != 0 {
		n += 1 + sovMessage(uint64(m.Priority))
	}
	if m.Cancel {
		n += 2
	}
	if m.WantType != 0 {
		n += 1 + sovMessage(uint64(m.WantType))
	}
	if m.SendDontHave {
		n += 2
	}
	return n
}

func (m *Message_Block) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Prefix)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func (m *Message_BlockPresence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cid)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovMessage(uint64(m.Type))
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wantlist", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Wantlist.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocks", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blocks = append(m.Blocks, make([]byte, postIndex-iNdEx))
			copy(m.Blocks[len(m.Blocks)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload, Message_Block{})
			if err := m.Payload[len(m.Payload)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockPresences", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockPresences = append(m.BlockPresences, Message_BlockPresence{})
			if err := m.BlockPresences[len(m.BlockPresences)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Message_Wantlist) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Wantlist: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Wantlist: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, Message_Wantlist_Entry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Full", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Full = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Message_Wantlist_Entry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Entry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Block = append(m.Block[:0], dAtA[iNdEx:postIndex]...)
			if m.Block == nil {
				m.Block = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			m.Priority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Priority |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cancel", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Cancel = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WantType", wireType)
			}
			m.WantType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WantType |= Message_Wantlist_WantType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendDontHave", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SendDontHave = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Message_Block) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Block: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Block: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefix", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prefix = append(m.Prefix[:0], dAtA[iNdEx:postIndex]...)
			if m.Prefix == nil {
				m.Prefix = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Message_BlockPresence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockPresence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockPresence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cid = append(m.Cid[:0], dAtA[iNdEx:postIndex]...)
			if m.Cid == nil {
				m.Cid = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Message_BlockPresenceType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthMessage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMessage
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMessage(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthMessage
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMessage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage   = fmt.Errorf("proto: integer overflow")
)
