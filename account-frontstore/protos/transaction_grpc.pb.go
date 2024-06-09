// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: protos/transaction.proto

package protos

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

// TransactionPointsClient is the client API for TransactionPoints service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionPointsClient interface {
	AddTransactionPoints(ctx context.Context, in *TransactionDetails, opts ...grpc.CallOption) (*AddPointsResponse, error)
	UseTransactionPoints(ctx context.Context, in *TransactionDetails, opts ...grpc.CallOption) (*UsePointsResponse, error)
	GetTransactionPoints(ctx context.Context, in *GetPointsRequest, opts ...grpc.CallOption) (*GetPointsResponse, error)
}

type transactionPointsClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionPointsClient(cc grpc.ClientConnInterface) TransactionPointsClient {
	return &transactionPointsClient{cc}
}

func (c *transactionPointsClient) AddTransactionPoints(ctx context.Context, in *TransactionDetails, opts ...grpc.CallOption) (*AddPointsResponse, error) {
	out := new(AddPointsResponse)
	err := c.cc.Invoke(ctx, "/protos.TransactionPoints/AddTransactionPoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionPointsClient) UseTransactionPoints(ctx context.Context, in *TransactionDetails, opts ...grpc.CallOption) (*UsePointsResponse, error) {
	out := new(UsePointsResponse)
	err := c.cc.Invoke(ctx, "/protos.TransactionPoints/UseTransactionPoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionPointsClient) GetTransactionPoints(ctx context.Context, in *GetPointsRequest, opts ...grpc.CallOption) (*GetPointsResponse, error) {
	out := new(GetPointsResponse)
	err := c.cc.Invoke(ctx, "/protos.TransactionPoints/GetTransactionPoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionPointsServer is the server API for TransactionPoints service.
// All implementations must embed UnimplementedTransactionPointsServer
// for forward compatibility
type TransactionPointsServer interface {
	AddTransactionPoints(context.Context, *TransactionDetails) (*AddPointsResponse, error)
	UseTransactionPoints(context.Context, *TransactionDetails) (*UsePointsResponse, error)
	GetTransactionPoints(context.Context, *GetPointsRequest) (*GetPointsResponse, error)
	mustEmbedUnimplementedTransactionPointsServer()
}

// UnimplementedTransactionPointsServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionPointsServer struct {
}

func (UnimplementedTransactionPointsServer) AddTransactionPoints(context.Context, *TransactionDetails) (*AddPointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTransactionPoints not implemented")
}
func (UnimplementedTransactionPointsServer) UseTransactionPoints(context.Context, *TransactionDetails) (*UsePointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UseTransactionPoints not implemented")
}
func (UnimplementedTransactionPointsServer) GetTransactionPoints(context.Context, *GetPointsRequest) (*GetPointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionPoints not implemented")
}
func (UnimplementedTransactionPointsServer) mustEmbedUnimplementedTransactionPointsServer() {}

// UnsafeTransactionPointsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionPointsServer will
// result in compilation errors.
type UnsafeTransactionPointsServer interface {
	mustEmbedUnimplementedTransactionPointsServer()
}

func RegisterTransactionPointsServer(s grpc.ServiceRegistrar, srv TransactionPointsServer) {
	s.RegisterService(&TransactionPoints_ServiceDesc, srv)
}

func _TransactionPoints_AddTransactionPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionPointsServer).AddTransactionPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TransactionPoints/AddTransactionPoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionPointsServer).AddTransactionPoints(ctx, req.(*TransactionDetails))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionPoints_UseTransactionPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionPointsServer).UseTransactionPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TransactionPoints/UseTransactionPoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionPointsServer).UseTransactionPoints(ctx, req.(*TransactionDetails))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionPoints_GetTransactionPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionPointsServer).GetTransactionPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TransactionPoints/GetTransactionPoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionPointsServer).GetTransactionPoints(ctx, req.(*GetPointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionPoints_ServiceDesc is the grpc.ServiceDesc for TransactionPoints service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionPoints_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.TransactionPoints",
	HandlerType: (*TransactionPointsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTransactionPoints",
			Handler:    _TransactionPoints_AddTransactionPoints_Handler,
		},
		{
			MethodName: "UseTransactionPoints",
			Handler:    _TransactionPoints_UseTransactionPoints_Handler,
		},
		{
			MethodName: "GetTransactionPoints",
			Handler:    _TransactionPoints_GetTransactionPoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/transaction.proto",
}