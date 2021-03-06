// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpcserver/token_storage.proto

/*
Package grpcserver is a generated protocol buffer package.

It is generated from these files:
	grpcserver/token_storage.proto

It has these top-level messages:
	Token
	TokenRequest
	TokenResponse
*/
package grpcserver

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

type Token struct {
	Value     string   `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	UserId    string   `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Lifetime  int32    `protobuf:"varint,3,opt,name=lifetime" json:"lifetime,omitempty"`
	Timestamp int64    `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Scope     []string `protobuf:"bytes,5,rep,name=scope" json:"scope,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Token) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Token) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Token) GetLifetime() int32 {
	if m != nil {
		return m.Lifetime
	}
	return 0
}

func (m *Token) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Token) GetScope() []string {
	if m != nil {
		return m.Scope
	}
	return nil
}

type TokenRequest struct {
	Signature string `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Token     *Token `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
}

func (m *TokenRequest) Reset()                    { *m = TokenRequest{} }
func (m *TokenRequest) String() string            { return proto.CompactTextString(m) }
func (*TokenRequest) ProtoMessage()               {}
func (*TokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TokenRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *TokenRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TokenRequest) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type TokenResponse struct {
	IsOk   bool   `protobuf:"varint,1,opt,name=is_ok,json=isOk" json:"is_ok,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	Token  *Token `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
}

func (m *TokenResponse) Reset()                    { *m = TokenResponse{} }
func (m *TokenResponse) String() string            { return proto.CompactTextString(m) }
func (*TokenResponse) ProtoMessage()               {}
func (*TokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TokenResponse) GetIsOk() bool {
	if m != nil {
		return m.IsOk
	}
	return false
}

func (m *TokenResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TokenResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func init() {
	proto.RegisterType((*Token)(nil), "grpcserver.Token")
	proto.RegisterType((*TokenRequest)(nil), "grpcserver.TokenRequest")
	proto.RegisterType((*TokenResponse)(nil), "grpcserver.TokenResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TokenStorage service

type TokenStorageClient interface {
	AddToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
	DropToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
	ValidateToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
}

type tokenStorageClient struct {
	cc *grpc.ClientConn
}

func NewTokenStorageClient(cc *grpc.ClientConn) TokenStorageClient {
	return &tokenStorageClient{cc}
}

func (c *tokenStorageClient) AddToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := grpc.Invoke(ctx, "/grpcserver.TokenStorage/AddToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenStorageClient) DropToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := grpc.Invoke(ctx, "/grpcserver.TokenStorage/DropToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenStorageClient) ValidateToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := grpc.Invoke(ctx, "/grpcserver.TokenStorage/ValidateToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TokenStorage service

type TokenStorageServer interface {
	AddToken(context.Context, *TokenRequest) (*TokenResponse, error)
	DropToken(context.Context, *TokenRequest) (*TokenResponse, error)
	ValidateToken(context.Context, *TokenRequest) (*TokenResponse, error)
}

func RegisterTokenStorageServer(s *grpc.Server, srv TokenStorageServer) {
	s.RegisterService(&_TokenStorage_serviceDesc, srv)
}

func _TokenStorage_AddToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenStorageServer).AddToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcserver.TokenStorage/AddToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenStorageServer).AddToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenStorage_DropToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenStorageServer).DropToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcserver.TokenStorage/DropToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenStorageServer).DropToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenStorage_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenStorageServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcserver.TokenStorage/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenStorageServer).ValidateToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TokenStorage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcserver.TokenStorage",
	HandlerType: (*TokenStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddToken",
			Handler:    _TokenStorage_AddToken_Handler,
		},
		{
			MethodName: "DropToken",
			Handler:    _TokenStorage_DropToken_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _TokenStorage_ValidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcserver/token_storage.proto",
}

func init() { proto.RegisterFile("grpcserver/token_storage.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4d, 0x4b, 0x3b, 0x31,
	0x10, 0xc6, 0xff, 0x69, 0x9b, 0xfe, 0xdb, 0xd1, 0x1e, 0x8c, 0xa2, 0xb1, 0x88, 0x2c, 0x7b, 0xb1,
	0xa7, 0x0a, 0xf5, 0x13, 0x54, 0x44, 0xf0, 0x24, 0xac, 0xe2, 0xb5, 0xc4, 0xee, 0x58, 0x42, 0x5f,
	0x12, 0x33, 0xd9, 0x7e, 0x04, 0xbf, 0xa5, 0xdf, 0x45, 0x92, 0xb4, 0x2e, 0xbe, 0x1c, 0x84, 0x9e,
	0x96, 0x67, 0x66, 0x33, 0xbf, 0x67, 0x9e, 0x04, 0xce, 0x67, 0xce, 0x4e, 0x09, 0xdd, 0x1a, 0xdd,
	0xa5, 0x37, 0x73, 0x5c, 0x4d, 0xc8, 0x1b, 0xa7, 0x66, 0x38, 0xb4, 0xce, 0x78, 0x23, 0xa0, 0xee,
	0xe7, 0x6f, 0x0c, 0xf8, 0x63, 0xf8, 0x47, 0x1c, 0x01, 0x5f, 0xab, 0x45, 0x85, 0x92, 0x65, 0x6c,
	0xd0, 0x2d, 0x92, 0x10, 0x27, 0xf0, 0xbf, 0x22, 0x74, 0x13, 0x5d, 0xca, 0x46, 0xac, 0xb7, 0x83,
	0xbc, 0x2b, 0x45, 0x1f, 0x3a, 0x0b, 0xfd, 0x82, 0x5e, 0x2f, 0x51, 0x36, 0x33, 0x36, 0xe0, 0xc5,
	0xa7, 0x16, 0x67, 0xd0, 0x0d, 0x5f, 0xf2, 0x6a, 0x69, 0x65, 0x2b, 0x63, 0x83, 0x66, 0x51, 0x17,
	0x02, 0x88, 0xa6, 0xc6, 0xa2, 0xe4, 0x59, 0x33, 0x80, 0xa2, 0xc8, 0x09, 0xf6, 0xa3, 0x8f, 0x02,
	0x5f, 0x2b, 0x24, 0x1f, 0x66, 0x90, 0x9e, 0xad, 0x94, 0xaf, 0xdc, 0xd6, 0x52, 0x5d, 0xf8, 0x4a,
	0x68, 0x7c, 0x27, 0x5c, 0x00, 0x8f, 0x7b, 0x47, 0x63, 0x7b, 0xa3, 0x83, 0x61, 0xbd, 0xf0, 0x30,
	0x41, 0x52, 0x3f, 0x47, 0xe8, 0x6d, 0xa0, 0x64, 0xcd, 0x8a, 0x50, 0x1c, 0x02, 0xd7, 0x34, 0x31,
	0xf3, 0x48, 0xec, 0x14, 0x2d, 0x4d, 0xf7, 0x73, 0x71, 0x0c, 0x6d, 0xf2, 0xca, 0x57, 0xb4, 0x8d,
	0x20, 0xa9, 0x3f, 0x63, 0x46, 0xef, 0x6c, 0xb3, 0xdc, 0x43, 0xba, 0x07, 0x31, 0x86, 0xce, 0xb8,
	0x2c, 0x53, 0xee, 0xf2, 0xe7, 0xb1, 0x14, 0x41, 0xff, 0xf4, 0x97, 0x4e, 0xf2, 0x99, 0xff, 0x13,
	0xd7, 0xd0, 0xbd, 0x71, 0xc6, 0xee, 0x34, 0xe3, 0x16, 0x7a, 0x4f, 0x6a, 0xa1, 0x4b, 0xe5, 0x71,
	0x97, 0x39, 0xcf, 0xed, 0xf8, 0xae, 0xae, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x81, 0x0a, 0xce,
	0x90, 0x79, 0x02, 0x00, 0x00,
}
