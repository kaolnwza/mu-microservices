// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: wallet.proto

package pb

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

// WalletServiceClient is the client API for WalletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletServiceClient interface {
	CreateUserWallet(ctx context.Context, in *UserWalletRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetUserWallet(ctx context.Context, in *UserWalletRequest, opts ...grpc.CallOption) (*WalletResponse, error)
	IncreaseUserWallet(ctx context.Context, in *UpdateWalletRequest, opts ...grpc.CallOption) (*WalletResponse, error)
	DecreaseUserWallet(ctx context.Context, in *UpdateWalletRequest, opts ...grpc.CallOption) (*WalletResponse, error)
}

type walletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletServiceClient(cc grpc.ClientConnInterface) WalletServiceClient {
	return &walletServiceClient{cc}
}

func (c *walletServiceClient) CreateUserWallet(ctx context.Context, in *UserWalletRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/wallet.WalletService/CreateUserWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetUserWallet(ctx context.Context, in *UserWalletRequest, opts ...grpc.CallOption) (*WalletResponse, error) {
	out := new(WalletResponse)
	err := c.cc.Invoke(ctx, "/wallet.WalletService/GetUserWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) IncreaseUserWallet(ctx context.Context, in *UpdateWalletRequest, opts ...grpc.CallOption) (*WalletResponse, error) {
	out := new(WalletResponse)
	err := c.cc.Invoke(ctx, "/wallet.WalletService/IncreaseUserWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) DecreaseUserWallet(ctx context.Context, in *UpdateWalletRequest, opts ...grpc.CallOption) (*WalletResponse, error) {
	out := new(WalletResponse)
	err := c.cc.Invoke(ctx, "/wallet.WalletService/DecreaseUserWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServiceServer is the server API for WalletService service.
// All implementations must embed UnimplementedWalletServiceServer
// for forward compatibility
type WalletServiceServer interface {
	CreateUserWallet(context.Context, *UserWalletRequest) (*EmptyResponse, error)
	GetUserWallet(context.Context, *UserWalletRequest) (*WalletResponse, error)
	IncreaseUserWallet(context.Context, *UpdateWalletRequest) (*WalletResponse, error)
	DecreaseUserWallet(context.Context, *UpdateWalletRequest) (*WalletResponse, error)
	mustEmbedUnimplementedWalletServiceServer()
}

// UnimplementedWalletServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWalletServiceServer struct {
}

func (UnimplementedWalletServiceServer) CreateUserWallet(context.Context, *UserWalletRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserWallet not implemented")
}
func (UnimplementedWalletServiceServer) GetUserWallet(context.Context, *UserWalletRequest) (*WalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserWallet not implemented")
}
func (UnimplementedWalletServiceServer) IncreaseUserWallet(context.Context, *UpdateWalletRequest) (*WalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncreaseUserWallet not implemented")
}
func (UnimplementedWalletServiceServer) DecreaseUserWallet(context.Context, *UpdateWalletRequest) (*WalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecreaseUserWallet not implemented")
}
func (UnimplementedWalletServiceServer) mustEmbedUnimplementedWalletServiceServer() {}

// UnsafeWalletServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServiceServer will
// result in compilation errors.
type UnsafeWalletServiceServer interface {
	mustEmbedUnimplementedWalletServiceServer()
}

func RegisterWalletServiceServer(s grpc.ServiceRegistrar, srv WalletServiceServer) {
	s.RegisterService(&WalletService_ServiceDesc, srv)
}

func _WalletService_CreateUserWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).CreateUserWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.WalletService/CreateUserWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).CreateUserWallet(ctx, req.(*UserWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetUserWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetUserWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.WalletService/GetUserWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetUserWallet(ctx, req.(*UserWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_IncreaseUserWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).IncreaseUserWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.WalletService/IncreaseUserWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).IncreaseUserWallet(ctx, req.(*UpdateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_DecreaseUserWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).DecreaseUserWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.WalletService/DecreaseUserWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).DecreaseUserWallet(ctx, req.(*UpdateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletService_ServiceDesc is the grpc.ServiceDesc for WalletService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wallet.WalletService",
	HandlerType: (*WalletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserWallet",
			Handler:    _WalletService_CreateUserWallet_Handler,
		},
		{
			MethodName: "GetUserWallet",
			Handler:    _WalletService_GetUserWallet_Handler,
		},
		{
			MethodName: "IncreaseUserWallet",
			Handler:    _WalletService_IncreaseUserWallet_Handler,
		},
		{
			MethodName: "DecreaseUserWallet",
			Handler:    _WalletService_DecreaseUserWallet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet.proto",
}
