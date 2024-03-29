// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: tntchain/tntchain/extend.proto

package types

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ExtendClient is the client API for Extend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExtendClient interface {
	// Transfer v1beta1版本的钱包转账.
	Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error)
	// AddKey v1beta1版本的创建钱包.
	AddKey(ctx context.Context, in *AddKeyRequest, opts ...grpc.CallOption) (*AddKeyResponse, error)
	// GetTransactionList v1beta1版本的获取交易列表
	GetTransactionList(ctx context.Context, in *GetTransactionListRequest, opts ...grpc.CallOption) (*GetTransactionListResponse, error)
}

type extendClient struct {
	cc grpc.ClientConnInterface
}

func NewExtendClient(cc grpc.ClientConnInterface) ExtendClient {
	return &extendClient{cc}
}

func (c *extendClient) Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error) {
	out := new(TransferResponse)
	err := c.cc.Invoke(ctx, "/tntchain.tntchain.Extend/Transfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extendClient) AddKey(ctx context.Context, in *AddKeyRequest, opts ...grpc.CallOption) (*AddKeyResponse, error) {
	out := new(AddKeyResponse)
	err := c.cc.Invoke(ctx, "/tntchain.tntchain.Extend/AddKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extendClient) GetTransactionList(ctx context.Context, in *GetTransactionListRequest, opts ...grpc.CallOption) (*GetTransactionListResponse, error) {
	out := new(GetTransactionListResponse)
	err := c.cc.Invoke(ctx, "/tntchain.tntchain.Extend/GetTransactionList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExtendServer is the server API for Extend service.
// All implementations must embed UnimplementedExtendServer
// for forward compatibility
type ExtendServer interface {
	// Transfer v1beta1版本的钱包转账.
	Transfer(context.Context, *TransferRequest) (*TransferResponse, error)
	// AddKey v1beta1版本的创建钱包.
	AddKey(context.Context, *AddKeyRequest) (*AddKeyResponse, error)
	// GetTransactionList v1beta1版本的获取交易列表
	GetTransactionList(context.Context, *GetTransactionListRequest) (*GetTransactionListResponse, error)
	mustEmbedUnimplementedExtendServer()
}

// UnimplementedExtendServer must be embedded to have forward compatible implementations.
type UnimplementedExtendServer struct {
}

func (UnimplementedExtendServer) Transfer(context.Context, *TransferRequest) (*TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (UnimplementedExtendServer) AddKey(context.Context, *AddKeyRequest) (*AddKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddKey not implemented")
}
func (UnimplementedExtendServer) GetTransactionList(context.Context, *GetTransactionListRequest) (*GetTransactionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionList not implemented")
}
func (UnimplementedExtendServer) mustEmbedUnimplementedExtendServer() {}

// UnsafeExtendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExtendServer will
// result in compilation errors.
type UnsafeExtendServer interface {
	mustEmbedUnimplementedExtendServer()
}

func RegisterExtendServer(s grpc.ServiceRegistrar, srv ExtendServer) {
	s.RegisterService(&Extend_ServiceDesc, srv)
}

func _Extend_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtendServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tntchain.tntchain.Extend/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtendServer).Transfer(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Extend_AddKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtendServer).AddKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tntchain.tntchain.Extend/AddKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtendServer).AddKey(ctx, req.(*AddKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Extend_GetTransactionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtendServer).GetTransactionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tntchain.tntchain.Extend/GetTransactionList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtendServer).GetTransactionList(ctx, req.(*GetTransactionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Extend_ServiceDesc is the grpc.ServiceDesc for Extend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Extend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tntchain.tntchain.Extend",
	HandlerType: (*ExtendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transfer",
			Handler:    _Extend_Transfer_Handler,
		},
		{
			MethodName: "AddKey",
			Handler:    _Extend_AddKey_Handler,
		},
		{
			MethodName: "GetTransactionList",
			Handler:    _Extend_GetTransactionList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tntchain/tntchain/extend.proto",
}
