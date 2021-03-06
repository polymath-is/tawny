// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Presence.proto

package tawny

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PresenceSubscribeResponse_Type int32

const (
	PresenceSubscribeResponse_FULL    PresenceSubscribeResponse_Type = 0
	PresenceSubscribeResponse_PARTIAL PresenceSubscribeResponse_Type = 1
)

var PresenceSubscribeResponse_Type_name = map[int32]string{
	0: "FULL",
	1: "PARTIAL",
}

var PresenceSubscribeResponse_Type_value = map[string]int32{
	"FULL":    0,
	"PARTIAL": 1,
}

func (x PresenceSubscribeResponse_Type) String() string {
	return proto.EnumName(PresenceSubscribeResponse_Type_name, int32(x))
}

func (PresenceSubscribeResponse_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77021aee44002059, []int{2, 0}
}

type PresenceSubscribeInput struct {
	Channel              string   `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Topic                string   `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PresenceSubscribeInput) Reset()         { *m = PresenceSubscribeInput{} }
func (m *PresenceSubscribeInput) String() string { return proto.CompactTextString(m) }
func (*PresenceSubscribeInput) ProtoMessage()    {}
func (*PresenceSubscribeInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_77021aee44002059, []int{0}
}

func (m *PresenceSubscribeInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PresenceSubscribeInput.Unmarshal(m, b)
}
func (m *PresenceSubscribeInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PresenceSubscribeInput.Marshal(b, m, deterministic)
}
func (m *PresenceSubscribeInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PresenceSubscribeInput.Merge(m, src)
}
func (m *PresenceSubscribeInput) XXX_Size() int {
	return xxx_messageInfo_PresenceSubscribeInput.Size(m)
}
func (m *PresenceSubscribeInput) XXX_DiscardUnknown() {
	xxx_messageInfo_PresenceSubscribeInput.DiscardUnknown(m)
}

var xxx_messageInfo_PresenceSubscribeInput proto.InternalMessageInfo

func (m *PresenceSubscribeInput) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *PresenceSubscribeInput) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

type Presence struct {
	State                []byte   `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Presence) Reset()         { *m = Presence{} }
func (m *Presence) String() string { return proto.CompactTextString(m) }
func (*Presence) ProtoMessage()    {}
func (*Presence) Descriptor() ([]byte, []int) {
	return fileDescriptor_77021aee44002059, []int{1}
}

func (m *Presence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Presence.Unmarshal(m, b)
}
func (m *Presence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Presence.Marshal(b, m, deterministic)
}
func (m *Presence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Presence.Merge(m, src)
}
func (m *Presence) XXX_Size() int {
	return xxx_messageInfo_Presence.Size(m)
}
func (m *Presence) XXX_DiscardUnknown() {
	xxx_messageInfo_Presence.DiscardUnknown(m)
}

var xxx_messageInfo_Presence proto.InternalMessageInfo

func (m *Presence) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *Presence) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type PresenceSubscribeResponse struct {
	Presences            []*Presence                    `protobuf:"bytes,1,rep,name=presences,proto3" json:"presences,omitempty"`
	UpdateType           PresenceSubscribeResponse_Type `protobuf:"varint,2,opt,name=update_type,json=updateType,proto3,enum=tawny.PresenceSubscribeResponse_Type" json:"update_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *PresenceSubscribeResponse) Reset()         { *m = PresenceSubscribeResponse{} }
func (m *PresenceSubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*PresenceSubscribeResponse) ProtoMessage()    {}
func (*PresenceSubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77021aee44002059, []int{2}
}

func (m *PresenceSubscribeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PresenceSubscribeResponse.Unmarshal(m, b)
}
func (m *PresenceSubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PresenceSubscribeResponse.Marshal(b, m, deterministic)
}
func (m *PresenceSubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PresenceSubscribeResponse.Merge(m, src)
}
func (m *PresenceSubscribeResponse) XXX_Size() int {
	return xxx_messageInfo_PresenceSubscribeResponse.Size(m)
}
func (m *PresenceSubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PresenceSubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PresenceSubscribeResponse proto.InternalMessageInfo

func (m *PresenceSubscribeResponse) GetPresences() []*Presence {
	if m != nil {
		return m.Presences
	}
	return nil
}

func (m *PresenceSubscribeResponse) GetUpdateType() PresenceSubscribeResponse_Type {
	if m != nil {
		return m.UpdateType
	}
	return PresenceSubscribeResponse_FULL
}

type HeartBeatInput struct {
	State                []byte   `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Channel              string   `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
	Topic                string   `protobuf:"bytes,4,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartBeatInput) Reset()         { *m = HeartBeatInput{} }
func (m *HeartBeatInput) String() string { return proto.CompactTextString(m) }
func (*HeartBeatInput) ProtoMessage()    {}
func (*HeartBeatInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_77021aee44002059, []int{3}
}

func (m *HeartBeatInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartBeatInput.Unmarshal(m, b)
}
func (m *HeartBeatInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartBeatInput.Marshal(b, m, deterministic)
}
func (m *HeartBeatInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartBeatInput.Merge(m, src)
}
func (m *HeartBeatInput) XXX_Size() int {
	return xxx_messageInfo_HeartBeatInput.Size(m)
}
func (m *HeartBeatInput) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartBeatInput.DiscardUnknown(m)
}

var xxx_messageInfo_HeartBeatInput proto.InternalMessageInfo

func (m *HeartBeatInput) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *HeartBeatInput) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *HeartBeatInput) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *HeartBeatInput) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func init() {
	proto.RegisterEnum("tawny.PresenceSubscribeResponse_Type", PresenceSubscribeResponse_Type_name, PresenceSubscribeResponse_Type_value)
	proto.RegisterType((*PresenceSubscribeInput)(nil), "tawny.PresenceSubscribeInput")
	proto.RegisterType((*Presence)(nil), "tawny.Presence")
	proto.RegisterType((*PresenceSubscribeResponse)(nil), "tawny.PresenceSubscribeResponse")
	proto.RegisterType((*HeartBeatInput)(nil), "tawny.HeartBeatInput")
}

func init() {
	proto.RegisterFile("Presence.proto", fileDescriptor_77021aee44002059)
}

var fileDescriptor_77021aee44002059 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0x86, 0xd9, 0x0f, 0xf8, 0x80, 0xc1, 0x00, 0x4e, 0x94, 0x54, 0x0c, 0x09, 0xd9, 0xc4, 0x84,
	0x8b, 0x8b, 0xa9, 0x57, 0x2f, 0x98, 0x48, 0x20, 0xe1, 0x40, 0x2a, 0x26, 0xde, 0x4c, 0x5b, 0x47,
	0x20, 0x62, 0xbb, 0x69, 0xb7, 0x9a, 0xfe, 0x24, 0x8f, 0xfe, 0x43, 0xd3, 0x2d, 0x45, 0xb4, 0x21,
	0x7a, 0xea, 0xcc, 0x74, 0x9e, 0x9d, 0x79, 0xdf, 0x0c, 0x34, 0x66, 0x01, 0x85, 0xe4, 0xb9, 0x24,
	0x64, 0xe0, 0x2b, 0x1f, 0xcb, 0xca, 0x7e, 0xf3, 0xe2, 0xce, 0xe9, 0xc2, 0xf7, 0x17, 0x6b, 0x1a,
	0xe8, 0xa2, 0x13, 0x3d, 0x0d, 0xe8, 0x45, 0xaa, 0x38, 0xed, 0xe1, 0x63, 0x68, 0x67, 0xd4, 0x6d,
	0xe4, 0x84, 0x6e, 0xb0, 0x72, 0x68, 0xe2, 0xc9, 0x48, 0xa1, 0x01, 0x15, 0x77, 0x69, 0x7b, 0x1e,
	0xad, 0x0d, 0xd6, 0x63, 0xfd, 0x9a, 0x95, 0xa5, 0x78, 0x04, 0x65, 0xe5, 0xcb, 0x95, 0x6b, 0xfc,
	0xd3, 0xf5, 0x34, 0xe1, 0x26, 0x54, 0xb3, 0x97, 0x92, 0x8e, 0x50, 0xd9, 0x8a, 0x34, 0x79, 0x60,
	0xa5, 0x09, 0xb6, 0xa0, 0xf8, 0x4c, 0xf1, 0x86, 0x4a, 0x42, 0xfe, 0xc1, 0xe0, 0x24, 0x37, 0xde,
	0xa2, 0x50, 0xfa, 0x5e, 0x48, 0x78, 0x0e, 0x35, 0xb9, 0xf9, 0x19, 0x1a, 0xac, 0x57, 0xec, 0xd7,
	0xcd, 0xa6, 0xd0, 0x9a, 0x44, 0x06, 0x59, 0x5f, 0x1d, 0x38, 0x82, 0x7a, 0x24, 0x1f, 0x6d, 0x45,
	0x0f, 0x2a, 0x96, 0xa4, 0xc7, 0x34, 0xcc, 0xb3, 0x1f, 0x40, 0x6e, 0x8a, 0x98, 0xc7, 0x92, 0x2c,
	0x48, 0xc9, 0x24, 0xe6, 0x5d, 0x28, 0x25, 0x5f, 0xac, 0x42, 0x69, 0x74, 0x37, 0x9d, 0xb6, 0x0a,
	0x58, 0x87, 0xca, 0x6c, 0x68, 0xcd, 0x27, 0xc3, 0x69, 0x8b, 0xf1, 0x25, 0x34, 0xc6, 0x64, 0x07,
	0xea, 0x9a, 0x6c, 0x95, 0x3a, 0xf5, 0x47, 0xb5, 0xbb, 0x8e, 0x16, 0xf7, 0x38, 0x5a, 0xda, 0x71,
	0xd4, 0x7c, 0x67, 0xd0, 0xdc, 0xee, 0x4d, 0xc1, 0xeb, 0xca, 0x25, 0xbc, 0x87, 0xc3, 0x9c, 0x14,
	0xec, 0xee, 0x13, 0xa9, 0xf7, 0xeb, 0xf4, 0x7e, 0xf3, 0x80, 0x17, 0x2e, 0x18, 0x5e, 0x41, 0x6d,
	0xab, 0x0b, 0x8f, 0x37, 0xc8, 0x77, 0xa5, 0x9d, 0xb6, 0x48, 0x6f, 0x49, 0x64, 0xb7, 0x24, 0x6e,
	0x92, 0x5b, 0xe2, 0x05, 0xe7, 0xbf, 0xae, 0x5c, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x35, 0x6e,
	0xce, 0x23, 0x84, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PresenceServiceClient is the client API for PresenceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PresenceServiceClient interface {
	//* Subscribe to a channel/topic to get presence update
	PresenceSubscribe(ctx context.Context, in *PresenceSubscribeInput, opts ...grpc.CallOption) (PresenceService_PresenceSubscribeClient, error)
	//* Heartbeat to a channel
	HeartBeat(ctx context.Context, in *HeartBeatInput, opts ...grpc.CallOption) (*empty.Empty, error)
}

type presenceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPresenceServiceClient(cc grpc.ClientConnInterface) PresenceServiceClient {
	return &presenceServiceClient{cc}
}

func (c *presenceServiceClient) PresenceSubscribe(ctx context.Context, in *PresenceSubscribeInput, opts ...grpc.CallOption) (PresenceService_PresenceSubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PresenceService_serviceDesc.Streams[0], "/tawny.PresenceService/PresenceSubscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &presenceServicePresenceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PresenceService_PresenceSubscribeClient interface {
	Recv() (*PresenceSubscribeResponse, error)
	grpc.ClientStream
}

type presenceServicePresenceSubscribeClient struct {
	grpc.ClientStream
}

func (x *presenceServicePresenceSubscribeClient) Recv() (*PresenceSubscribeResponse, error) {
	m := new(PresenceSubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *presenceServiceClient) HeartBeat(ctx context.Context, in *HeartBeatInput, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/tawny.PresenceService/HeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PresenceServiceServer is the server API for PresenceService service.
type PresenceServiceServer interface {
	//* Subscribe to a channel/topic to get presence update
	PresenceSubscribe(*PresenceSubscribeInput, PresenceService_PresenceSubscribeServer) error
	//* Heartbeat to a channel
	HeartBeat(context.Context, *HeartBeatInput) (*empty.Empty, error)
}

// UnimplementedPresenceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPresenceServiceServer struct {
}

func (*UnimplementedPresenceServiceServer) PresenceSubscribe(req *PresenceSubscribeInput, srv PresenceService_PresenceSubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method PresenceSubscribe not implemented")
}
func (*UnimplementedPresenceServiceServer) HeartBeat(ctx context.Context, req *HeartBeatInput) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}

func RegisterPresenceServiceServer(s *grpc.Server, srv PresenceServiceServer) {
	s.RegisterService(&_PresenceService_serviceDesc, srv)
}

func _PresenceService_PresenceSubscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PresenceSubscribeInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PresenceServiceServer).PresenceSubscribe(m, &presenceServicePresenceSubscribeServer{stream})
}

type PresenceService_PresenceSubscribeServer interface {
	Send(*PresenceSubscribeResponse) error
	grpc.ServerStream
}

type presenceServicePresenceSubscribeServer struct {
	grpc.ServerStream
}

func (x *presenceServicePresenceSubscribeServer) Send(m *PresenceSubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PresenceService_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenceServiceServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tawny.PresenceService/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenceServiceServer).HeartBeat(ctx, req.(*HeartBeatInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _PresenceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tawny.PresenceService",
	HandlerType: (*PresenceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HeartBeat",
			Handler:    _PresenceService_HeartBeat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PresenceSubscribe",
			Handler:       _PresenceService_PresenceSubscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "Presence.proto",
}
