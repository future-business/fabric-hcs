// Code generated by protoc-gen-go.
// source: peer/events.proto
// DO NOT EDIT!

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EventType int32

const (
	EventType_REGISTER  EventType = 0
	EventType_BLOCK     EventType = 1
	EventType_CHAINCODE EventType = 2
	EventType_REJECTION EventType = 3
)

var EventType_name = map[int32]string{
	0: "REGISTER",
	1: "BLOCK",
	2: "CHAINCODE",
	3: "REJECTION",
}
var EventType_value = map[string]int32{
	"REGISTER":  0,
	"BLOCK":     1,
	"CHAINCODE": 2,
	"REJECTION": 3,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}
func (EventType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

// ChaincodeReg is used for registering chaincode Interests
// when EventType is CHAINCODE
type ChaincodeReg struct {
	ChaincodeID string `protobuf:"bytes,1,opt,name=chaincodeID" json:"chaincodeID,omitempty"`
	EventName   string `protobuf:"bytes,2,opt,name=eventName" json:"eventName,omitempty"`
}

func (m *ChaincodeReg) Reset()                    { *m = ChaincodeReg{} }
func (m *ChaincodeReg) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeReg) ProtoMessage()               {}
func (*ChaincodeReg) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type Interest struct {
	EventType EventType `protobuf:"varint,1,opt,name=eventType,enum=protos.EventType" json:"eventType,omitempty"`
	// Ideally we should just have the following oneof for different
	// Reg types and get rid of EventType. But this is an API change
	// Additional Reg types may add messages specific to their type
	// to the oneof.
	//
	// Types that are valid to be assigned to RegInfo:
	//	*Interest_ChaincodeRegInfo
	RegInfo isInterest_RegInfo `protobuf_oneof:"RegInfo"`
}

func (m *Interest) Reset()                    { *m = Interest{} }
func (m *Interest) String() string            { return proto.CompactTextString(m) }
func (*Interest) ProtoMessage()               {}
func (*Interest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

type isInterest_RegInfo interface {
	isInterest_RegInfo()
}

type Interest_ChaincodeRegInfo struct {
	ChaincodeRegInfo *ChaincodeReg `protobuf:"bytes,2,opt,name=chaincodeRegInfo,oneof"`
}

func (*Interest_ChaincodeRegInfo) isInterest_RegInfo() {}

func (m *Interest) GetRegInfo() isInterest_RegInfo {
	if m != nil {
		return m.RegInfo
	}
	return nil
}

func (m *Interest) GetChaincodeRegInfo() *ChaincodeReg {
	if x, ok := m.GetRegInfo().(*Interest_ChaincodeRegInfo); ok {
		return x.ChaincodeRegInfo
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Interest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Interest_OneofMarshaler, _Interest_OneofUnmarshaler, _Interest_OneofSizer, []interface{}{
		(*Interest_ChaincodeRegInfo)(nil),
	}
}

func _Interest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Interest)
	// RegInfo
	switch x := m.RegInfo.(type) {
	case *Interest_ChaincodeRegInfo:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ChaincodeRegInfo); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Interest.RegInfo has unexpected type %T", x)
	}
	return nil
}

func _Interest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Interest)
	switch tag {
	case 2: // RegInfo.chaincodeRegInfo
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ChaincodeReg)
		err := b.DecodeMessage(msg)
		m.RegInfo = &Interest_ChaincodeRegInfo{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Interest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Interest)
	// RegInfo
	switch x := m.RegInfo.(type) {
	case *Interest_ChaincodeRegInfo:
		s := proto.Size(x.ChaincodeRegInfo)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// ---------- consumer events ---------
// Register is sent by consumers for registering events
// string type - "register"
type Register struct {
	Events []*Interest `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
}

func (m *Register) Reset()                    { *m = Register{} }
func (m *Register) String() string            { return proto.CompactTextString(m) }
func (*Register) ProtoMessage()               {}
func (*Register) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *Register) GetEvents() []*Interest {
	if m != nil {
		return m.Events
	}
	return nil
}

// Rejection is sent by consumers for erroneous transaction rejection events
// string type - "rejection"
type Rejection struct {
	Tx       *Transaction `protobuf:"bytes,1,opt,name=tx" json:"tx,omitempty"`
	ErrorMsg string       `protobuf:"bytes,2,opt,name=errorMsg" json:"errorMsg,omitempty"`
}

func (m *Rejection) Reset()                    { *m = Rejection{} }
func (m *Rejection) String() string            { return proto.CompactTextString(m) }
func (*Rejection) ProtoMessage()               {}
func (*Rejection) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *Rejection) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

// ---------- producer events ---------
type Unregister struct {
	Events []*Interest `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
}

func (m *Unregister) Reset()                    { *m = Unregister{} }
func (m *Unregister) String() string            { return proto.CompactTextString(m) }
func (*Unregister) ProtoMessage()               {}
func (*Unregister) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *Unregister) GetEvents() []*Interest {
	if m != nil {
		return m.Events
	}
	return nil
}

// Event is used by
//  - consumers (adapters) to send Register
//  - producer to advertise supported types and events
type Event struct {
	// Types that are valid to be assigned to Event:
	//	*Event_Register
	//	*Event_Block
	//	*Event_ChaincodeEvent
	//	*Event_Rejection
	//	*Event_Unregister
	Event isEvent_Event `protobuf_oneof:"Event"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

type isEvent_Event interface {
	isEvent_Event()
}

type Event_Register struct {
	Register *Register `protobuf:"bytes,1,opt,name=register,oneof"`
}
type Event_Block struct {
	Block *Block2 `protobuf:"bytes,2,opt,name=block,oneof"`
}
type Event_ChaincodeEvent struct {
	ChaincodeEvent *ChaincodeEvent `protobuf:"bytes,3,opt,name=chaincodeEvent,oneof"`
}
type Event_Rejection struct {
	Rejection *Rejection `protobuf:"bytes,4,opt,name=rejection,oneof"`
}
type Event_Unregister struct {
	Unregister *Unregister `protobuf:"bytes,5,opt,name=unregister,oneof"`
}

func (*Event_Register) isEvent_Event()       {}
func (*Event_Block) isEvent_Event()          {}
func (*Event_ChaincodeEvent) isEvent_Event() {}
func (*Event_Rejection) isEvent_Event()      {}
func (*Event_Unregister) isEvent_Event()     {}

func (m *Event) GetEvent() isEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Event) GetRegister() *Register {
	if x, ok := m.GetEvent().(*Event_Register); ok {
		return x.Register
	}
	return nil
}

func (m *Event) GetBlock() *Block2 {
	if x, ok := m.GetEvent().(*Event_Block); ok {
		return x.Block
	}
	return nil
}

func (m *Event) GetChaincodeEvent() *ChaincodeEvent {
	if x, ok := m.GetEvent().(*Event_ChaincodeEvent); ok {
		return x.ChaincodeEvent
	}
	return nil
}

func (m *Event) GetRejection() *Rejection {
	if x, ok := m.GetEvent().(*Event_Rejection); ok {
		return x.Rejection
	}
	return nil
}

func (m *Event) GetUnregister() *Unregister {
	if x, ok := m.GetEvent().(*Event_Unregister); ok {
		return x.Unregister
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Event) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Event_OneofMarshaler, _Event_OneofUnmarshaler, _Event_OneofSizer, []interface{}{
		(*Event_Register)(nil),
		(*Event_Block)(nil),
		(*Event_ChaincodeEvent)(nil),
		(*Event_Rejection)(nil),
		(*Event_Unregister)(nil),
	}
}

func _Event_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Event)
	// Event
	switch x := m.Event.(type) {
	case *Event_Register:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Register); err != nil {
			return err
		}
	case *Event_Block:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Block); err != nil {
			return err
		}
	case *Event_ChaincodeEvent:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ChaincodeEvent); err != nil {
			return err
		}
	case *Event_Rejection:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Rejection); err != nil {
			return err
		}
	case *Event_Unregister:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Unregister); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Event.Event has unexpected type %T", x)
	}
	return nil
}

func _Event_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Event)
	switch tag {
	case 1: // Event.register
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Register)
		err := b.DecodeMessage(msg)
		m.Event = &Event_Register{msg}
		return true, err
	case 2: // Event.block
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Block2)
		err := b.DecodeMessage(msg)
		m.Event = &Event_Block{msg}
		return true, err
	case 3: // Event.chaincodeEvent
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ChaincodeEvent)
		err := b.DecodeMessage(msg)
		m.Event = &Event_ChaincodeEvent{msg}
		return true, err
	case 4: // Event.rejection
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Rejection)
		err := b.DecodeMessage(msg)
		m.Event = &Event_Rejection{msg}
		return true, err
	case 5: // Event.unregister
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Unregister)
		err := b.DecodeMessage(msg)
		m.Event = &Event_Unregister{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Event_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Event)
	// Event
	switch x := m.Event.(type) {
	case *Event_Register:
		s := proto.Size(x.Register)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Block:
		s := proto.Size(x.Block)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_ChaincodeEvent:
		s := proto.Size(x.ChaincodeEvent)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Rejection:
		s := proto.Size(x.Rejection)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Unregister:
		s := proto.Size(x.Unregister)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ChaincodeReg)(nil), "protos.ChaincodeReg")
	proto.RegisterType((*Interest)(nil), "protos.Interest")
	proto.RegisterType((*Register)(nil), "protos.Register")
	proto.RegisterType((*Rejection)(nil), "protos.Rejection")
	proto.RegisterType((*Unregister)(nil), "protos.Unregister")
	proto.RegisterType((*Event)(nil), "protos.Event")
	proto.RegisterEnum("protos.EventType", EventType_name, EventType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Events service

type EventsClient interface {
	// event chatting using Event
	Chat(ctx context.Context, opts ...grpc.CallOption) (Events_ChatClient, error)
}

type eventsClient struct {
	cc *grpc.ClientConn
}

func NewEventsClient(cc *grpc.ClientConn) EventsClient {
	return &eventsClient{cc}
}

func (c *eventsClient) Chat(ctx context.Context, opts ...grpc.CallOption) (Events_ChatClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Events_serviceDesc.Streams[0], c.cc, "/protos.Events/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventsChatClient{stream}
	return x, nil
}

type Events_ChatClient interface {
	Send(*Event) error
	Recv() (*Event, error)
	grpc.ClientStream
}

type eventsChatClient struct {
	grpc.ClientStream
}

func (x *eventsChatClient) Send(m *Event) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventsChatClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Events service

type EventsServer interface {
	// event chatting using Event
	Chat(Events_ChatServer) error
}

func RegisterEventsServer(s *grpc.Server, srv EventsServer) {
	s.RegisterService(&_Events_serviceDesc, srv)
}

func _Events_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventsServer).Chat(&eventsChatServer{stream})
}

type Events_ChatServer interface {
	Send(*Event) error
	Recv() (*Event, error)
	grpc.ServerStream
}

type eventsChatServer struct {
	grpc.ServerStream
}

func (x *eventsChatServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventsChatServer) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Events_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Events",
	HandlerType: (*EventsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Chat",
			Handler:       _Events_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor4,
}

func init() { proto.RegisterFile("peer/events.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x51, 0x8f, 0x93, 0x4c,
	0x14, 0x05, 0xba, 0xed, 0xc2, 0xed, 0x6e, 0x43, 0xef, 0xf7, 0x45, 0xb1, 0xd1, 0xa4, 0xc1, 0xc4,
	0xd4, 0x35, 0x29, 0x8a, 0x8d, 0xcf, 0x0a, 0x4b, 0x04, 0x5d, 0xdb, 0x64, 0xac, 0x2f, 0xbe, 0x18,
	0xca, 0xce, 0x52, 0x74, 0x17, 0x9a, 0x61, 0xd6, 0xec, 0xfe, 0x05, 0x5f, 0xfd, 0xc3, 0xa6, 0x03,
	0x03, 0xac, 0xfb, 0xe4, 0x13, 0xcc, 0x3d, 0xe7, 0xdc, 0x7b, 0x38, 0xcc, 0x85, 0xf1, 0x8e, 0x52,
	0xe6, 0xd0, 0x9f, 0x34, 0xe7, 0xe5, 0x7c, 0xc7, 0x0a, 0x5e, 0xe0, 0x40, 0x3c, 0xca, 0xc9, 0x23,
	0x01, 0x25, 0xdb, 0x38, 0xcb, 0x93, 0xe2, 0x9c, 0x0a, 0x4e, 0x45, 0x99, 0x3c, 0x11, 0xd0, 0x45,
	0xbc, 0x61, 0x59, 0xf2, 0x8d, 0xb3, 0x38, 0x2f, 0xe3, 0x84, 0x67, 0x45, 0x5e, 0xc3, 0x0f, 0xbb,
	0xf0, 0xe6, 0xb2, 0x48, 0x7e, 0x54, 0x80, 0xbd, 0x84, 0x23, 0x5f, 0xf6, 0x23, 0x34, 0xc5, 0x29,
	0x0c, 0x9b, 0xfe, 0xd1, 0xa9, 0xa5, 0x4e, 0xd5, 0x99, 0x41, 0xba, 0x25, 0x7c, 0x0c, 0x86, 0x18,
	0xbc, 0x8c, 0xaf, 0xa8, 0xa5, 0x09, 0xbc, 0x2d, 0xd8, 0xbf, 0x54, 0xd0, 0xa3, 0x9c, 0x53, 0x46,
	0x4b, 0x8e, 0x4e, 0x4d, 0x5d, 0xdf, 0xee, 0xa8, 0x68, 0x35, 0x72, 0xc7, 0xd5, 0xdc, 0x72, 0x1e,
	0x48, 0x80, 0xb4, 0x1c, 0xf4, 0xc0, 0x4c, 0x3a, 0x6e, 0xa2, 0xfc, 0xa2, 0x10, 0x23, 0x86, 0xee,
	0xff, 0x52, 0xd7, 0x75, 0x1b, 0x2a, 0xe4, 0x1e, 0xdf, 0x33, 0xe0, 0xb0, 0x7e, 0xb5, 0x17, 0xa0,
	0x13, 0x9a, 0x66, 0x25, 0xa7, 0x0c, 0x67, 0x30, 0xa8, 0x32, 0xb5, 0xd4, 0x69, 0x6f, 0x36, 0x74,
	0x4d, 0xd9, 0x50, 0xba, 0x25, 0x35, 0x6e, 0x9f, 0x81, 0x41, 0xe8, 0x77, 0x2a, 0xe2, 0xc3, 0xa7,
	0xa0, 0xf1, 0x1b, 0xe1, 0x7d, 0xe8, 0xfe, 0x27, 0x25, 0xeb, 0x36, 0x5f, 0xa2, 0xf1, 0x1b, 0x9c,
	0x80, 0x4e, 0x19, 0x2b, 0xd8, 0xa7, 0x32, 0xad, 0x13, 0x69, 0xce, 0xf6, 0x1b, 0x80, 0x2f, 0x39,
	0xfb, 0x77, 0x17, 0xbf, 0x35, 0xe8, 0x8b, 0x8c, 0x70, 0x0e, 0xba, 0xd4, 0xd7, 0x46, 0x1a, 0x95,
	0xfc, 0xba, 0x50, 0x21, 0x0d, 0x07, 0x9f, 0x41, 0x5f, 0xfc, 0xe1, 0x3a, 0xb9, 0x91, 0x24, 0x7b,
	0xfb, 0xa2, 0x1b, 0x2a, 0xa4, 0x82, 0xf1, 0x2d, 0x8c, 0x9a, 0xf0, 0xc4, 0x24, 0xab, 0x27, 0x04,
	0x0f, 0xee, 0x45, 0x2d, 0xd0, 0x50, 0x21, 0x7f, 0xf1, 0xf1, 0x15, 0x18, 0x4c, 0x26, 0x65, 0x1d,
	0x08, 0xf1, 0xb8, 0xb5, 0x56, 0x03, 0xa1, 0x42, 0x5a, 0x16, 0x2e, 0x00, 0xae, 0x9b, 0x38, 0xac,
	0xbe, 0xd0, 0xa0, 0xd4, 0xb4, 0x41, 0x85, 0x0a, 0xe9, 0xf0, 0xbc, 0xc3, 0x3a, 0x8b, 0x13, 0x0f,
	0x8c, 0xe6, 0xe2, 0xe0, 0x11, 0xe8, 0x24, 0x78, 0x1f, 0x7d, 0x5e, 0x07, 0xc4, 0x54, 0xd0, 0x80,
	0xbe, 0x77, 0xb6, 0xf2, 0x3f, 0x9a, 0x2a, 0x1e, 0x83, 0xe1, 0x87, 0xef, 0xa2, 0xa5, 0xbf, 0x3a,
	0x0d, 0x4c, 0x6d, 0x7f, 0x24, 0xc1, 0x87, 0xc0, 0x5f, 0x47, 0xab, 0xa5, 0xd9, 0x73, 0x17, 0x30,
	0x10, 0x3d, 0x4a, 0x3c, 0x81, 0x03, 0x7f, 0x1b, 0x73, 0x3c, 0xbe, 0x73, 0x29, 0x27, 0x77, 0x8f,
	0xb6, 0x32, 0x53, 0x5f, 0xaa, 0xde, 0x8b, 0xaf, 0xcf, 0xd3, 0x8c, 0x6f, 0xaf, 0x37, 0xf3, 0xa4,
	0xb8, 0x72, 0xb6, 0xb7, 0x3b, 0xca, 0x2e, 0xe9, 0x79, 0xda, 0x6c, 0x95, 0x53, 0x69, 0x9c, 0xfd,
	0xa2, 0x6d, 0xaa, 0x85, 0x7d, 0xfd, 0x27, 0x00, 0x00, 0xff, 0xff, 0x44, 0x56, 0x27, 0xe8, 0xcc,
	0x03, 0x00, 0x00,
}
