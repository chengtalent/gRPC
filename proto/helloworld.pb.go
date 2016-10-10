// Code generated by protoc-gen-go.
// source: helloworld.proto
// DO NOT EDIT!

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	helloworld.proto

It has these top-level messages:
	HelloRequest
	HelloReply
	NoParam
	IPList
	CertificateRequest
	CertificateReply
*/
package helloworld

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type NoParam struct {
}

func (m *NoParam) Reset()                    { *m = NoParam{} }
func (m *NoParam) String() string            { return proto.CompactTextString(m) }
func (*NoParam) ProtoMessage()               {}
func (*NoParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type IPList struct {
	Ip []string `protobuf:"bytes,1,rep,name=ip" json:"ip,omitempty"`
}

func (m *IPList) Reset()                    { *m = IPList{} }
func (m *IPList) String() string            { return proto.CompactTextString(m) }
func (*IPList) ProtoMessage()               {}
func (*IPList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type CertificateRequest struct {
	In   []byte `protobuf:"bytes,1,opt,name=in,proto3" json:"in,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *CertificateRequest) Reset()                    { *m = CertificateRequest{} }
func (m *CertificateRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateRequest) ProtoMessage()               {}
func (*CertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type CertificateReply struct {
	In []byte `protobuf:"bytes,1,opt,name=in,proto3" json:"in,omitempty"`
}

func (m *CertificateReply) Reset()                    { *m = CertificateReply{} }
func (m *CertificateReply) String() string            { return proto.CompactTextString(m) }
func (*CertificateReply) ProtoMessage()               {}
func (*CertificateReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
	proto.RegisterType((*NoParam)(nil), "helloworld.NoParam")
	proto.RegisterType((*IPList)(nil), "helloworld.IPList")
	proto.RegisterType((*CertificateRequest)(nil), "helloworld.CertificateRequest")
	proto.RegisterType((*CertificateReply)(nil), "helloworld.CertificateReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Greeter service

type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

// Client API for Whitelist service

type WhitelistClient interface {
	GetWhitelist(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*IPList, error)
}

type whitelistClient struct {
	cc *grpc.ClientConn
}

func NewWhitelistClient(cc *grpc.ClientConn) WhitelistClient {
	return &whitelistClient{cc}
}

func (c *whitelistClient) GetWhitelist(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*IPList, error) {
	out := new(IPList)
	err := grpc.Invoke(ctx, "/helloworld.Whitelist/GetWhitelist", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Whitelist service

type WhitelistServer interface {
	GetWhitelist(context.Context, *NoParam) (*IPList, error)
}

func RegisterWhitelistServer(s *grpc.Server, srv WhitelistServer) {
	s.RegisterService(&_Whitelist_serviceDesc, srv)
}

func _Whitelist_GetWhitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhitelistServer).GetWhitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Whitelist/GetWhitelist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhitelistServer).GetWhitelist(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

var _Whitelist_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Whitelist",
	HandlerType: (*WhitelistServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWhitelist",
			Handler:    _Whitelist_GetWhitelist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

// Client API for CA service

type CAClient interface {
	IssueCertificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateReply, error)
	GetCACertificate(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*CertificateReply, error)
}

type cAClient struct {
	cc *grpc.ClientConn
}

func NewCAClient(cc *grpc.ClientConn) CAClient {
	return &cAClient{cc}
}

func (c *cAClient) IssueCertificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateReply, error) {
	out := new(CertificateReply)
	err := grpc.Invoke(ctx, "/helloworld.CA/IssueCertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cAClient) GetCACertificate(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*CertificateReply, error) {
	out := new(CertificateReply)
	err := grpc.Invoke(ctx, "/helloworld.CA/GetCACertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CA service

type CAServer interface {
	IssueCertificate(context.Context, *CertificateRequest) (*CertificateReply, error)
	GetCACertificate(context.Context, *NoParam) (*CertificateReply, error)
}

func RegisterCAServer(s *grpc.Server, srv CAServer) {
	s.RegisterService(&_CA_serviceDesc, srv)
}

func _CA_IssueCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAServer).IssueCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.CA/IssueCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAServer).IssueCertificate(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CA_GetCACertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAServer).GetCACertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.CA/GetCACertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAServer).GetCACertificate(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

var _CA_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.CA",
	HandlerType: (*CAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueCertificate",
			Handler:    _CA_IssueCertificate_Handler,
		},
		{
			MethodName: "GetCACertificate",
			Handler:    _CA_GetCACertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x9b, 0x28, 0xad, 0x19, 0x8a, 0x86, 0x11, 0x24, 0x54, 0x91, 0xb2, 0x07, 0xe9, 0x29,
	0x48, 0xbd, 0xe8, 0x45, 0x68, 0x0b, 0xc6, 0x82, 0x48, 0x88, 0x87, 0x9e, 0xd7, 0x3a, 0xb6, 0x0b,
	0x9b, 0x66, 0xdd, 0xdd, 0xa2, 0x7d, 0x1c, 0xdf, 0x54, 0x12, 0x1b, 0xbb, 0x6a, 0xc5, 0xdb, 0xce,
	0x9f, 0x6f, 0xf6, 0xf7, 0x0d, 0x03, 0xe1, 0x9c, 0xa4, 0x2c, 0x5e, 0x0b, 0x2d, 0x9f, 0x62, 0xa5,
	0x0b, 0x5b, 0x20, 0x6c, 0x32, 0x8c, 0x41, 0xfb, 0xb6, 0x8c, 0x32, 0x7a, 0x59, 0x92, 0xb1, 0x88,
	0xb0, 0xbb, 0xe0, 0x39, 0x45, 0x5e, 0xd7, 0xeb, 0x05, 0x59, 0xf5, 0x66, 0x67, 0x00, 0xeb, 0x1e,
	0x25, 0x57, 0x18, 0x41, 0x2b, 0x27, 0x63, 0xf8, 0xac, 0x6e, 0xaa, 0x43, 0x16, 0x40, 0xeb, 0xbe,
	0x48, 0xb9, 0xe6, 0x39, 0x8b, 0xa0, 0x39, 0x4e, 0xef, 0x84, 0xb1, 0xb8, 0x0f, 0xbe, 0x50, 0x91,
	0xd7, 0xdd, 0xe9, 0x05, 0x99, 0x2f, 0x14, 0xbb, 0x04, 0x1c, 0x91, 0xb6, 0xe2, 0x59, 0x4c, 0xb9,
	0xa5, 0xfa, 0xdb, 0xb2, 0x6b, 0x51, 0xcd, 0x6b, 0x67, 0xbe, 0x58, 0x7c, 0x61, 0xf8, 0x0e, 0x06,
	0x83, 0xf0, 0x9b, 0xb2, 0x84, 0xf9, 0xa1, 0xeb, 0x8f, 0xa1, 0x95, 0x68, 0x22, 0x4b, 0x1a, 0xaf,
	0x61, 0xef, 0x81, 0xaf, 0x2a, 0x70, 0x8c, 0x62, 0x67, 0x09, 0xae, 0xdf, 0xce, 0xd1, 0x96, 0x8a,
	0x92, 0x2b, 0xd6, 0xe8, 0xdf, 0x40, 0x30, 0x99, 0x0b, 0x4b, 0xb2, 0x74, 0x71, 0x05, 0xed, 0x84,
	0xec, 0x26, 0x3e, 0x74, 0x65, 0x6b, 0xd3, 0x1d, 0x74, 0x93, 0x9f, 0xf6, 0x59, 0xa3, 0xff, 0xee,
	0x81, 0x3f, 0x1a, 0x60, 0x06, 0xe1, 0xd8, 0x98, 0x25, 0x39, 0x16, 0xf0, 0xd4, 0x15, 0xfc, 0xde,
	0x4a, 0xe7, 0xe4, 0xcf, 0x7a, 0x85, 0x88, 0x09, 0x84, 0x09, 0xd9, 0xd1, 0xc0, 0x9d, 0xb9, 0x95,
	0xec, 0x9f, 0x41, 0xc3, 0x73, 0x38, 0x16, 0x45, 0x3c, 0xd3, 0x6a, 0x1a, 0xd3, 0x1b, 0xcf, 0x95,
	0x24, 0xe3, 0x28, 0x86, 0x07, 0xd5, 0x62, 0x26, 0xe5, 0x3b, 0x2d, 0x2f, 0x28, 0xf5, 0x1e, 0x9b,
	0xd5, 0x29, 0x5d, 0x7c, 0x04, 0x00, 0x00, 0xff, 0xff, 0x49, 0x69, 0xdb, 0x9d, 0x5e, 0x02, 0x00,
	0x00,
}
