// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	app.proto

It has these top-level messages:
	StoreReq
	StoreResp
	GetReq
	GetResp
*/
package rpc

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

type StoreReq struct {
	Key string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Val string `protobuf:"bytes,2,opt,name=Val" json:"Val,omitempty"`
}

func (m *StoreReq) Reset()                    { *m = StoreReq{} }
func (m *StoreReq) String() string            { return proto.CompactTextString(m) }
func (*StoreReq) ProtoMessage()               {}
func (*StoreReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StoreReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StoreReq) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

type StoreResp struct {
}

func (m *StoreResp) Reset()                    { *m = StoreResp{} }
func (m *StoreResp) String() string            { return proto.CompactTextString(m) }
func (*StoreResp) ProtoMessage()               {}
func (*StoreResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GetReq struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetReq) Reset()                    { *m = GetReq{} }
func (m *GetReq) String() string            { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()               {}
func (*GetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetResp struct {
	Val string `protobuf:"bytes,1,opt,name=val" json:"val,omitempty"`
}

func (m *GetResp) Reset()                    { *m = GetResp{} }
func (m *GetResp) String() string            { return proto.CompactTextString(m) }
func (*GetResp) ProtoMessage()               {}
func (*GetResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetResp) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

func init() {
	proto.RegisterType((*StoreReq)(nil), "rpc.StoreReq")
	proto.RegisterType((*StoreResp)(nil), "rpc.StoreResp")
	proto.RegisterType((*GetReq)(nil), "rpc.GetReq")
	proto.RegisterType((*GetResp)(nil), "rpc.GetResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Cache service

type CacheClient interface {
	Store(ctx context.Context, in *StoreReq, opts ...grpc.CallOption) (*StoreResp, error)
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error)
}

type cacheClient struct {
	cc *grpc.ClientConn
}

func NewCacheClient(cc *grpc.ClientConn) CacheClient {
	return &cacheClient{cc}
}

func (c *cacheClient) Store(ctx context.Context, in *StoreReq, opts ...grpc.CallOption) (*StoreResp, error) {
	out := new(StoreResp)
	err := grpc.Invoke(ctx, "/rpc.Cache/Store", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	out := new(GetResp)
	err := grpc.Invoke(ctx, "/rpc.Cache/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cache service

type CacheServer interface {
	Store(context.Context, *StoreReq) (*StoreResp, error)
	Get(context.Context, *GetReq) (*GetResp, error)
}

func RegisterCacheServer(s *grpc.Server, srv CacheServer) {
	s.RegisterService(&_Cache_serviceDesc, srv)
}

func _Cache_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Cache/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Store(ctx, req.(*StoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Cache/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cache_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Cache",
	HandlerType: (*CacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _Cache_Store_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Cache_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app.proto",
}

func init() { proto.RegisterFile("app.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0x28, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x56, 0xd2, 0xe3, 0xe2, 0x08, 0x2e,
	0xc9, 0x2f, 0x4a, 0x0d, 0x4a, 0x2d, 0x14, 0x12, 0xe0, 0x62, 0xf6, 0x4e, 0xad, 0x94, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x41, 0x22, 0x61, 0x89, 0x39, 0x12, 0x4c, 0x10, 0x91, 0xb0,
	0xc4, 0x1c, 0x25, 0x6e, 0x2e, 0x4e, 0xa8, 0xfa, 0xe2, 0x02, 0x25, 0x29, 0x2e, 0x36, 0xf7, 0xd4,
	0x12, 0xa8, 0xd6, 0x6c, 0x84, 0xd6, 0xec, 0xd4, 0x4a, 0x25, 0x69, 0x2e, 0x76, 0xb0, 0x5c, 0x71,
	0x01, 0x48, 0xb2, 0x2c, 0x31, 0x07, 0x26, 0x59, 0x96, 0x98, 0x63, 0x14, 0xca, 0xc5, 0xea, 0x9c,
	0x98, 0x9c, 0x91, 0x2a, 0xa4, 0xc1, 0xc5, 0x0a, 0x36, 0x4e, 0x88, 0x57, 0xaf, 0xa8, 0x20, 0x59,
	0x0f, 0xe6, 0x14, 0x29, 0x3e, 0x64, 0x6e, 0x71, 0x81, 0x12, 0x83, 0x90, 0x12, 0x17, 0xb3, 0x7b,
	0x6a, 0x89, 0x10, 0x37, 0x58, 0x02, 0x62, 0xab, 0x14, 0x0f, 0x82, 0x03, 0x52, 0x93, 0xc4, 0x06,
	0xf6, 0x98, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x19, 0x5f, 0x3c, 0xb9, 0xe5, 0x00, 0x00, 0x00,
}