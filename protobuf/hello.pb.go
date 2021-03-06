// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/hello.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type String struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dd845233a77a6ff, []int{0}
}

func (m *String) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_String.Unmarshal(m, b)
}
func (m *String) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_String.Marshal(b, m, deterministic)
}
func (m *String) XXX_Merge(src proto.Message) {
	xxx_messageInfo_String.Merge(m, src)
}
func (m *String) XXX_Size() int {
	return xxx_messageInfo_String.Size(m)
}
func (m *String) XXX_DiscardUnknown() {
	xxx_messageInfo_String.DiscardUnknown(m)
}

var xxx_messageInfo_String proto.InternalMessageInfo

func (m *String) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*String)(nil), "proto.String")
}

func init() { proto.RegisterFile("protobuf/hello.proto", fileDescriptor_8dd845233a77a6ff) }

var fileDescriptor_8dd845233a77a6ff = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x03, 0x73, 0x85, 0x58, 0xc1,
	0x94, 0x92, 0x1c, 0x17, 0x5b, 0x70, 0x49, 0x51, 0x66, 0x5e, 0xba, 0x90, 0x08, 0x17, 0x6b, 0x59,
	0x62, 0x4e, 0x69, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84, 0x63, 0x94, 0xc4, 0xc5,
	0xe3, 0x01, 0xd2, 0x15, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0xa4, 0xca, 0xc5, 0x0a, 0xe6,
	0x0b, 0xf1, 0x42, 0xcc, 0xd1, 0x83, 0xe8, 0x96, 0x42, 0xe5, 0x0a, 0x69, 0x73, 0xb1, 0x3b, 0x67,
	0x24, 0xe6, 0xe5, 0xa5, 0xe6, 0xe0, 0x57, 0xa8, 0xc1, 0x68, 0xc0, 0x68, 0x94, 0xca, 0xc5, 0x1b,
	0x50, 0x9a, 0x54, 0x5c, 0x9a, 0x04, 0xb3, 0x44, 0x9d, 0x8b, 0xbd, 0xa0, 0x34, 0x29, 0x27, 0xb3,
	0x38, 0x83, 0xa0, 0x35, 0x9c, 0xc1, 0xa5, 0x49, 0xc5, 0xc9, 0x45, 0x99, 0x49, 0xa9, 0xf8, 0x95,
	0x1a, 0x30, 0x26, 0xb1, 0x81, 0xf9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x78, 0xe2,
	0x73, 0x10, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	Hello(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error)
	// 使用stream关键字支持流
	Channel(ctx context.Context, opts ...grpc.CallOption) (HelloService_ChannelClient, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) Hello(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/proto.HelloService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) Channel(ctx context.Context, opts ...grpc.CallOption) (HelloService_ChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloService_serviceDesc.Streams[0], "/proto.HelloService/Channel", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceChannelClient{stream}
	return x, nil
}

type HelloService_ChannelClient interface {
	Send(*String) error
	Recv() (*String, error)
	grpc.ClientStream
}

type helloServiceChannelClient struct {
	grpc.ClientStream
}

func (x *helloServiceChannelClient) Send(m *String) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceChannelClient) Recv() (*String, error) {
	m := new(String)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	Hello(context.Context, *String) (*String, error)
	// 使用stream关键字支持流
	Channel(HelloService_ChannelServer) error
}

// UnimplementedHelloServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (*UnimplementedHelloServiceServer) Hello(ctx context.Context, req *String) (*String, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (*UnimplementedHelloServiceServer) Channel(srv HelloService_ChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method Channel not implemented")
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HelloService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Hello(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_Channel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).Channel(&helloServiceChannelServer{stream})
}

type HelloService_ChannelServer interface {
	Send(*String) error
	Recv() (*String, error)
	grpc.ServerStream
}

type helloServiceChannelServer struct {
	grpc.ServerStream
}

func (x *helloServiceChannelServer) Send(m *String) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceChannelServer) Recv() (*String, error) {
	m := new(String)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _HelloService_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Channel",
			Handler:       _HelloService_Channel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protobuf/hello.proto",
}

// PubsubServiceClient is the client API for PubsubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PubsubServiceClient interface {
	Publish(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error)
	Subscribe(ctx context.Context, in *String, opts ...grpc.CallOption) (PubsubService_SubscribeClient, error)
}

type pubsubServiceClient struct {
	cc *grpc.ClientConn
}

func NewPubsubServiceClient(cc *grpc.ClientConn) PubsubServiceClient {
	return &pubsubServiceClient{cc}
}

func (c *pubsubServiceClient) Publish(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/proto.PubsubService/publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubsubServiceClient) Subscribe(ctx context.Context, in *String, opts ...grpc.CallOption) (PubsubService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PubsubService_serviceDesc.Streams[0], "/proto.PubsubService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &pubsubServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PubsubService_SubscribeClient interface {
	Recv() (*String, error)
	grpc.ClientStream
}

type pubsubServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *pubsubServiceSubscribeClient) Recv() (*String, error) {
	m := new(String)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PubsubServiceServer is the server API for PubsubService service.
type PubsubServiceServer interface {
	Publish(context.Context, *String) (*String, error)
	Subscribe(*String, PubsubService_SubscribeServer) error
}

// UnimplementedPubsubServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPubsubServiceServer struct {
}

func (*UnimplementedPubsubServiceServer) Publish(ctx context.Context, req *String) (*String, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (*UnimplementedPubsubServiceServer) Subscribe(req *String, srv PubsubService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterPubsubServiceServer(s *grpc.Server, srv PubsubServiceServer) {
	s.RegisterService(&_PubsubService_serviceDesc, srv)
}

func _PubsubService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubsubServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PubsubService/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubsubServiceServer).Publish(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubsubService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(String)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PubsubServiceServer).Subscribe(m, &pubsubServiceSubscribeServer{stream})
}

type PubsubService_SubscribeServer interface {
	Send(*String) error
	grpc.ServerStream
}

type pubsubServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *pubsubServiceSubscribeServer) Send(m *String) error {
	return x.ServerStream.SendMsg(m)
}

var _PubsubService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PubsubService",
	HandlerType: (*PubsubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "publish",
			Handler:    _PubsubService_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _PubsubService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protobuf/hello.proto",
}
