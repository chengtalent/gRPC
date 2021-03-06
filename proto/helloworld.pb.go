// Code generated by protoc-gen-go.
// source: helloworld.proto
// DO NOT EDIT!

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	helloworld.proto

It has these top-level messages:
	NoParam
	IPList
	CertificateRequest
	CertificateReply
	CertificateData
	SignatureValid
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

type NoParam struct {
}

func (m *NoParam) Reset()                    { *m = NoParam{} }
func (m *NoParam) String() string            { return proto.CompactTextString(m) }
func (*NoParam) ProtoMessage()               {}
func (*NoParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type IPList struct {
	Ip []string `protobuf:"bytes,1,rep,name=ip" json:"ip,omitempty"`
}

func (m *IPList) Reset()                    { *m = IPList{} }
func (m *IPList) String() string            { return proto.CompactTextString(m) }
func (*IPList) ProtoMessage()               {}
func (*IPList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CertificateRequest struct {
	In   []byte `protobuf:"bytes,1,opt,name=in,proto3" json:"in,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *CertificateRequest) Reset()                    { *m = CertificateRequest{} }
func (m *CertificateRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateRequest) ProtoMessage()               {}
func (*CertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type CertificateReply struct {
	In []byte `protobuf:"bytes,1,opt,name=in,proto3" json:"in,omitempty"`
}

func (m *CertificateReply) Reset()                    { *m = CertificateReply{} }
func (m *CertificateReply) String() string            { return proto.CompactTextString(m) }
func (*CertificateReply) ProtoMessage()               {}
func (*CertificateReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type CertificateData struct {
	Cert []byte `protobuf:"bytes,1,opt,name=cert,proto3" json:"cert,omitempty"`
	Root []byte `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
}

func (m *CertificateData) Reset()                    { *m = CertificateData{} }
func (m *CertificateData) String() string            { return proto.CompactTextString(m) }
func (*CertificateData) ProtoMessage()               {}
func (*CertificateData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type SignatureValid struct {
	Valid bool `protobuf:"varint,1,opt,name=valid" json:"valid,omitempty"`
}

func (m *SignatureValid) Reset()                    { *m = SignatureValid{} }
func (m *SignatureValid) String() string            { return proto.CompactTextString(m) }
func (*SignatureValid) ProtoMessage()               {}
func (*SignatureValid) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*NoParam)(nil), "helloworld.NoParam")
	proto.RegisterType((*IPList)(nil), "helloworld.IPList")
	proto.RegisterType((*CertificateRequest)(nil), "helloworld.CertificateRequest")
	proto.RegisterType((*CertificateReply)(nil), "helloworld.CertificateReply")
	proto.RegisterType((*CertificateData)(nil), "helloworld.CertificateData")
	proto.RegisterType((*SignatureValid)(nil), "helloworld.SignatureValid")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

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
	VerifySignature(ctx context.Context, in *CertificateData, opts ...grpc.CallOption) (*SignatureValid, error)
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

func (c *cAClient) VerifySignature(ctx context.Context, in *CertificateData, opts ...grpc.CallOption) (*SignatureValid, error) {
	out := new(SignatureValid)
	err := grpc.Invoke(ctx, "/helloworld.CA/VerifySignature", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CA service

type CAServer interface {
	IssueCertificate(context.Context, *CertificateRequest) (*CertificateReply, error)
	GetCACertificate(context.Context, *NoParam) (*CertificateReply, error)
	VerifySignature(context.Context, *CertificateData) (*SignatureValid, error)
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

func _CA_VerifySignature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAServer).VerifySignature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.CA/VerifySignature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAServer).VerifySignature(ctx, req.(*CertificateData))
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
		{
			MethodName: "VerifySignature",
			Handler:    _CA_VerifySignature_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4f, 0xc2, 0x30,
	0x14, 0xc6, 0xd9, 0x54, 0x74, 0x2f, 0x04, 0x96, 0xea, 0x61, 0x19, 0xc6, 0x90, 0x1e, 0x0c, 0xa7,
	0xc5, 0xe0, 0x45, 0x8e, 0x80, 0x11, 0x49, 0x88, 0x59, 0x66, 0x02, 0xe7, 0x0a, 0x0f, 0x68, 0x52,
	0xd6, 0xd9, 0x15, 0x95, 0xbf, 0xda, 0x7f, 0xc1, 0x74, 0x10, 0x29, 0x2a, 0xf1, 0xf6, 0xf5, 0xf5,
	0xfb, 0xbe, 0xb4, 0xbf, 0x3c, 0xf0, 0x17, 0x28, 0x84, 0x7c, 0x97, 0x4a, 0x4c, 0xa3, 0x4c, 0x49,
	0x2d, 0x09, 0xec, 0x26, 0xd4, 0x83, 0xd3, 0x27, 0x19, 0x33, 0xc5, 0x96, 0x34, 0x80, 0xf2, 0x20,
	0x1e, 0xf2, 0x5c, 0x93, 0x2a, 0xb8, 0x3c, 0x0b, 0x9c, 0xc6, 0x51, 0xd3, 0x4b, 0x5c, 0x9e, 0xd1,
	0x3b, 0x20, 0x3d, 0x54, 0x9a, 0xcf, 0xf8, 0x84, 0x69, 0x4c, 0xf0, 0x75, 0x85, 0x5b, 0x57, 0x1a,
	0x38, 0x0d, 0xa7, 0x59, 0x49, 0x5c, 0x9e, 0x12, 0x02, 0xc7, 0x29, 0x5b, 0x62, 0xe0, 0x36, 0x9c,
	0xa6, 0x97, 0x14, 0x9a, 0x52, 0xf0, 0xf7, 0x92, 0x99, 0x58, 0xff, 0xcc, 0xd1, 0x36, 0xd4, 0x2c,
	0xcf, 0x3d, 0xd3, 0xcc, 0x54, 0x4d, 0x50, 0xe9, 0xad, 0xa9, 0xd0, 0x66, 0xa6, 0xa4, 0xd4, 0x45,
	0x7d, 0x25, 0x29, 0x34, 0xbd, 0x86, 0xea, 0x33, 0x9f, 0xa7, 0x4c, 0xaf, 0x14, 0x8e, 0x98, 0xe0,
	0x53, 0x72, 0x01, 0x27, 0x6f, 0x46, 0x14, 0xd1, 0xb3, 0x64, 0x73, 0x68, 0x3d, 0x80, 0x37, 0x5e,
	0x70, 0x8d, 0xc2, 0xfc, 0xae, 0x0d, 0x95, 0x3e, 0xea, 0xdd, 0xf9, 0x3c, 0xb2, 0x08, 0x6d, 0x61,
	0x84, 0xc4, 0x1e, 0x6e, 0xb0, 0xd0, 0x52, 0xeb, 0xd3, 0x01, 0xb7, 0xd7, 0x21, 0x09, 0xf8, 0x83,
	0x3c, 0x5f, 0xa1, 0xf5, 0x6c, 0x72, 0x65, 0x07, 0x7e, 0xd3, 0x0a, 0x2f, 0x0f, 0xde, 0x67, 0x62,
	0x4d, 0x4b, 0xa4, 0x0f, 0x7e, 0x1f, 0x75, 0xaf, 0x63, 0x77, 0xfe, 0xf9, 0xb2, 0xff, 0x8a, 0x86,
	0x50, 0x1b, 0xa1, 0xe2, 0xb3, 0xf5, 0x37, 0x19, 0x52, 0x3f, 0x10, 0x31, 0xac, 0xc3, 0xd0, 0xbe,
	0xdc, 0xa7, 0x49, 0x4b, 0xdd, 0x1b, 0xa8, 0x73, 0x19, 0xcd, 0x55, 0x36, 0x89, 0xf0, 0x83, 0x2d,
	0x33, 0x81, 0xb9, 0xe5, 0xef, 0xd6, 0x1e, 0x8d, 0x1e, 0x1b, 0x1d, 0x9b, 0xdd, 0x8a, 0x9d, 0x97,
	0x72, 0xb1, 0x64, 0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xb9, 0x6b, 0xb1, 0x78, 0x02,
	0x00, 0x00,
}
