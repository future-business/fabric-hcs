// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderer/configuration.proto

package orderer

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// State defines the orderer mode of operation, typically for consensus-type migration.
// NORMAL is during normal operation, when consensus-type migration is not, and can not, take place.
// MAINTENANCE is when the consensus-type can be changed.
type ConsensusType_State int32

const (
	ConsensusType_STATE_NORMAL      ConsensusType_State = 0
	ConsensusType_STATE_MAINTENANCE ConsensusType_State = 1
)

var ConsensusType_State_name = map[int32]string{
	0: "STATE_NORMAL",
	1: "STATE_MAINTENANCE",
}

var ConsensusType_State_value = map[string]int32{
	"STATE_NORMAL":      0,
	"STATE_MAINTENANCE": 1,
}

func (x ConsensusType_State) String() string {
	return proto.EnumName(ConsensusType_State_name, int32(x))
}

func (ConsensusType_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bcce68f21316dd30, []int{0, 0}
}

type ConsensusType struct {
	// The consensus type: "solo", "kafka" or "etcdraft".
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Opaque metadata, dependent on the consensus type.
	Metadata []byte `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The state signals the ordering service to go into maintenance mode, typically for consensus-type migration.
	State                ConsensusType_State `protobuf:"varint,3,opt,name=state,proto3,enum=orderer.ConsensusType_State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ConsensusType) Reset()         { *m = ConsensusType{} }
func (m *ConsensusType) String() string { return proto.CompactTextString(m) }
func (*ConsensusType) ProtoMessage()    {}
func (*ConsensusType) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcce68f21316dd30, []int{0}
}

func (m *ConsensusType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusType.Unmarshal(m, b)
}
func (m *ConsensusType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusType.Marshal(b, m, deterministic)
}
func (m *ConsensusType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusType.Merge(m, src)
}
func (m *ConsensusType) XXX_Size() int {
	return xxx_messageInfo_ConsensusType.Size(m)
}
func (m *ConsensusType) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusType.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusType proto.InternalMessageInfo

func (m *ConsensusType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ConsensusType) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ConsensusType) GetState() ConsensusType_State {
	if m != nil {
		return m.State
	}
	return ConsensusType_STATE_NORMAL
}

type BatchSize struct {
	// Simply specified as number of messages for now, in the future
	// we may want to allow this to be specified by size in bytes
	MaxMessageCount uint32 `protobuf:"varint,1,opt,name=max_message_count,json=maxMessageCount,proto3" json:"max_message_count,omitempty"`
	// The byte count of the serialized messages in a batch cannot
	// exceed this value.
	AbsoluteMaxBytes uint32 `protobuf:"varint,2,opt,name=absolute_max_bytes,json=absoluteMaxBytes,proto3" json:"absolute_max_bytes,omitempty"`
	// The byte count of the serialized messages in a batch should not
	// exceed this value.
	PreferredMaxBytes    uint32   `protobuf:"varint,3,opt,name=preferred_max_bytes,json=preferredMaxBytes,proto3" json:"preferred_max_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchSize) Reset()         { *m = BatchSize{} }
func (m *BatchSize) String() string { return proto.CompactTextString(m) }
func (*BatchSize) ProtoMessage()    {}
func (*BatchSize) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcce68f21316dd30, []int{1}
}

func (m *BatchSize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchSize.Unmarshal(m, b)
}
func (m *BatchSize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchSize.Marshal(b, m, deterministic)
}
func (m *BatchSize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchSize.Merge(m, src)
}
func (m *BatchSize) XXX_Size() int {
	return xxx_messageInfo_BatchSize.Size(m)
}
func (m *BatchSize) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchSize.DiscardUnknown(m)
}

var xxx_messageInfo_BatchSize proto.InternalMessageInfo

func (m *BatchSize) GetMaxMessageCount() uint32 {
	if m != nil {
		return m.MaxMessageCount
	}
	return 0
}

func (m *BatchSize) GetAbsoluteMaxBytes() uint32 {
	if m != nil {
		return m.AbsoluteMaxBytes
	}
	return 0
}

func (m *BatchSize) GetPreferredMaxBytes() uint32 {
	if m != nil {
		return m.PreferredMaxBytes
	}
	return 0
}

type BatchTimeout struct {
	// Any duration string parseable by ParseDuration():
	// https://golang.org/pkg/time/#ParseDuration
	Timeout              string   `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchTimeout) Reset()         { *m = BatchTimeout{} }
func (m *BatchTimeout) String() string { return proto.CompactTextString(m) }
func (*BatchTimeout) ProtoMessage()    {}
func (*BatchTimeout) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcce68f21316dd30, []int{2}
}

func (m *BatchTimeout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchTimeout.Unmarshal(m, b)
}
func (m *BatchTimeout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchTimeout.Marshal(b, m, deterministic)
}
func (m *BatchTimeout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchTimeout.Merge(m, src)
}
func (m *BatchTimeout) XXX_Size() int {
	return xxx_messageInfo_BatchTimeout.Size(m)
}
func (m *BatchTimeout) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchTimeout.DiscardUnknown(m)
}

var xxx_messageInfo_BatchTimeout proto.InternalMessageInfo

func (m *BatchTimeout) GetTimeout() string {
	if m != nil {
		return m.Timeout
	}
	return ""
}

// Carries a list of bootstrap brokers, i.e. this is not the exclusive set of
// brokers an ordering service
type KafkaBrokers struct {
	// Each broker here should be identified using the (IP|host):port notation,
	// e.g. 127.0.0.1:7050, or localhost:7050 are valid entries
	Brokers              []string `protobuf:"bytes,1,rep,name=brokers,proto3" json:"brokers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaBrokers) Reset()         { *m = KafkaBrokers{} }
func (m *KafkaBrokers) String() string { return proto.CompactTextString(m) }
func (*KafkaBrokers) ProtoMessage()    {}
func (*KafkaBrokers) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcce68f21316dd30, []int{3}
}

func (m *KafkaBrokers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaBrokers.Unmarshal(m, b)
}
func (m *KafkaBrokers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaBrokers.Marshal(b, m, deterministic)
}
func (m *KafkaBrokers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaBrokers.Merge(m, src)
}
func (m *KafkaBrokers) XXX_Size() int {
	return xxx_messageInfo_KafkaBrokers.Size(m)
}
func (m *KafkaBrokers) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaBrokers.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaBrokers proto.InternalMessageInfo

func (m *KafkaBrokers) GetBrokers() []string {
	if m != nil {
		return m.Brokers
	}
	return nil
}

// ChannelRestrictions is the mssage which conveys restrictions on channel creation for an orderer
type ChannelRestrictions struct {
	MaxCount             uint64   `protobuf:"varint,1,opt,name=max_count,json=maxCount,proto3" json:"max_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelRestrictions) Reset()         { *m = ChannelRestrictions{} }
func (m *ChannelRestrictions) String() string { return proto.CompactTextString(m) }
func (*ChannelRestrictions) ProtoMessage()    {}
func (*ChannelRestrictions) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcce68f21316dd30, []int{4}
}

func (m *ChannelRestrictions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelRestrictions.Unmarshal(m, b)
}
func (m *ChannelRestrictions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelRestrictions.Marshal(b, m, deterministic)
}
func (m *ChannelRestrictions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelRestrictions.Merge(m, src)
}
func (m *ChannelRestrictions) XXX_Size() int {
	return xxx_messageInfo_ChannelRestrictions.Size(m)
}
func (m *ChannelRestrictions) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelRestrictions.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelRestrictions proto.InternalMessageInfo

func (m *ChannelRestrictions) GetMaxCount() uint64 {
	if m != nil {
		return m.MaxCount
	}
	return 0
}

type Hcs struct {
	// HCS topic ID in the format of shard.realm.num
	TopicId              string   `protobuf:"bytes,1,opt,name=TopicId,proto3" json:"TopicId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hcs) Reset()         { *m = Hcs{} }
func (m *Hcs) String() string { return proto.CompactTextString(m) }
func (*Hcs) ProtoMessage()    {}
func (*Hcs) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcce68f21316dd30, []int{5}
}

func (m *Hcs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hcs.Unmarshal(m, b)
}
func (m *Hcs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hcs.Marshal(b, m, deterministic)
}
func (m *Hcs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hcs.Merge(m, src)
}
func (m *Hcs) XXX_Size() int {
	return xxx_messageInfo_Hcs.Size(m)
}
func (m *Hcs) XXX_DiscardUnknown() {
	xxx_messageInfo_Hcs.DiscardUnknown(m)
}

var xxx_messageInfo_Hcs proto.InternalMessageInfo

func (m *Hcs) GetTopicId() string {
	if m != nil {
		return m.TopicId
	}
	return ""
}

func init() {
	proto.RegisterEnum("orderer.ConsensusType_State", ConsensusType_State_name, ConsensusType_State_value)
	proto.RegisterType((*ConsensusType)(nil), "orderer.ConsensusType")
	proto.RegisterType((*BatchSize)(nil), "orderer.BatchSize")
	proto.RegisterType((*BatchTimeout)(nil), "orderer.BatchTimeout")
	proto.RegisterType((*KafkaBrokers)(nil), "orderer.KafkaBrokers")
	proto.RegisterType((*ChannelRestrictions)(nil), "orderer.ChannelRestrictions")
	proto.RegisterType((*Hcs)(nil), "orderer.Hcs")
}

func init() { proto.RegisterFile("orderer/configuration.proto", fileDescriptor_bcce68f21316dd30) }

var fileDescriptor_bcce68f21316dd30 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xe1, 0x8a, 0xd3, 0x40,
	0x10, 0xc7, 0x8d, 0xbd, 0xf3, 0xae, 0x43, 0xab, 0xed, 0x1e, 0x42, 0xf0, 0x04, 0x4b, 0x40, 0x28,
	0xe2, 0x25, 0x52, 0x9f, 0xa0, 0x2d, 0x05, 0x0f, 0x6d, 0x85, 0x34, 0x1f, 0xc4, 0x2f, 0x65, 0x92,
	0x4c, 0xd3, 0xe5, 0x9a, 0x6c, 0xd8, 0xdd, 0x40, 0xeb, 0x7b, 0xf8, 0x08, 0xbe, 0xa7, 0xec, 0x6e,
	0x72, 0x9e, 0xdf, 0xe6, 0xff, 0x9f, 0xdf, 0x4c, 0x66, 0x32, 0x0b, 0xb7, 0x42, 0xe6, 0x24, 0x49,
	0x46, 0x99, 0xa8, 0xf6, 0xbc, 0x68, 0x24, 0x6a, 0x2e, 0xaa, 0xb0, 0x96, 0x42, 0x0b, 0x76, 0xd5,
	0x26, 0x83, 0x3f, 0x1e, 0x0c, 0x97, 0xa2, 0x52, 0x54, 0xa9, 0x46, 0x25, 0xe7, 0x9a, 0x18, 0x83,
	0x0b, 0x7d, 0xae, 0xc9, 0xf7, 0x26, 0xde, 0xb4, 0x1f, 0xdb, 0x98, 0xbd, 0x81, 0xeb, 0x92, 0x34,
	0xe6, 0xa8, 0xd1, 0x7f, 0x3e, 0xf1, 0xa6, 0x83, 0xf8, 0x51, 0xb3, 0x19, 0x5c, 0x2a, 0x8d, 0x9a,
	0xfc, 0xde, 0xc4, 0x9b, 0xbe, 0x9c, 0xbd, 0x0d, 0xdb, 0xd6, 0xe1, 0x7f, 0x6d, 0xc3, 0xad, 0x61,
	0x62, 0x87, 0x06, 0x9f, 0xe0, 0xd2, 0x6a, 0x36, 0x82, 0xc1, 0x36, 0x99, 0x27, 0xab, 0xdd, 0xe6,
	0x7b, 0xbc, 0x9e, 0x7f, 0x1b, 0x3d, 0x63, 0xaf, 0x61, 0xec, 0x9c, 0xf5, 0xfc, 0x7e, 0x93, 0xac,
	0x36, 0xf3, 0xcd, 0x72, 0x35, 0xf2, 0x82, 0xdf, 0x1e, 0xf4, 0x17, 0xa8, 0xb3, 0xc3, 0x96, 0xff,
	0x22, 0xf6, 0x01, 0xc6, 0x25, 0x9e, 0x76, 0x25, 0x29, 0x85, 0x05, 0xed, 0x32, 0xd1, 0x54, 0xda,
	0x0e, 0x3c, 0x8c, 0x5f, 0x95, 0x78, 0x5a, 0x3b, 0x7f, 0x69, 0x6c, 0xf6, 0x11, 0x18, 0xa6, 0x4a,
	0x1c, 0x1b, 0x4d, 0x3b, 0x53, 0x94, 0x9e, 0x35, 0x29, 0xbb, 0xc5, 0x30, 0x1e, 0x75, 0x99, 0x35,
	0x9e, 0x16, 0xc6, 0x67, 0x21, 0xdc, 0xd4, 0x92, 0xf6, 0x24, 0x25, 0xe5, 0x4f, 0xf0, 0x9e, 0xc5,
	0xc7, 0x8f, 0xa9, 0x8e, 0x0f, 0xa6, 0x30, 0xb0, 0x63, 0x25, 0xbc, 0x24, 0xd1, 0x68, 0xe6, 0xc3,
	0x95, 0x76, 0x61, 0xfb, 0x03, 0x3b, 0x69, 0xc8, 0xaf, 0xb8, 0x7f, 0xc0, 0x85, 0x14, 0x0f, 0x24,
	0x95, 0x21, 0x53, 0x17, 0xfa, 0xde, 0xa4, 0x67, 0xc8, 0x56, 0x06, 0x33, 0xb8, 0x59, 0x1e, 0xb0,
	0xaa, 0xe8, 0x18, 0x93, 0xd2, 0x92, 0x67, 0xe6, 0x70, 0x8a, 0xdd, 0x42, 0xdf, 0x0c, 0xf4, 0x6f,
	0xd9, 0x8b, 0xf8, 0xba, 0xc4, 0x93, 0xdd, 0x32, 0x78, 0x07, 0xbd, 0x2f, 0x99, 0x6d, 0x9a, 0x88,
	0x9a, 0x67, 0xf7, 0x79, 0xf7, 0xf9, 0x56, 0x2e, 0x7e, 0xc0, 0x7b, 0x21, 0x8b, 0xf0, 0x70, 0xae,
	0x49, 0x1e, 0x29, 0x2f, 0x48, 0x86, 0x7b, 0x4c, 0x25, 0xcf, 0xdc, 0x8b, 0x50, 0xdd, 0xd9, 0x7e,
	0x46, 0x05, 0xd7, 0x87, 0x26, 0x0d, 0x33, 0x51, 0x46, 0x4f, 0xe8, 0xc8, 0xd1, 0x77, 0x8e, 0xbe,
	0x2b, 0x44, 0xd4, 0x16, 0xa4, 0x2f, 0xac, 0xf5, 0xf9, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0d,
	0xfc, 0xe3, 0xae, 0x71, 0x02, 0x00, 0x00,
}
