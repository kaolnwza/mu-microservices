// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: storage.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProfileServiceClient is the client API for ProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileServiceClient interface {
	GetProfileImage(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*ProfileResponse, error)
	NewProfileImage(ctx context.Context, in *NewProfileRequest, opts ...grpc.CallOption) (*NewProfileResponse, error)
}

type profileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileServiceClient(cc grpc.ClientConnInterface) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) GetProfileImage(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*ProfileResponse, error) {
	out := new(ProfileResponse)
	err := c.cc.Invoke(ctx, "/storage.ProfileService/GetProfileImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) NewProfileImage(ctx context.Context, in *NewProfileRequest, opts ...grpc.CallOption) (*NewProfileResponse, error) {
	out := new(NewProfileResponse)
	err := c.cc.Invoke(ctx, "/storage.ProfileService/NewProfileImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServiceServer is the server API for ProfileService service.
// All implementations must embed UnimplementedProfileServiceServer
// for forward compatibility
type ProfileServiceServer interface {
	GetProfileImage(context.Context, *ProfileRequest) (*ProfileResponse, error)
	NewProfileImage(context.Context, *NewProfileRequest) (*NewProfileResponse, error)
	mustEmbedUnimplementedProfileServiceServer()
}

// UnimplementedProfileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProfileServiceServer struct {
}

func (UnimplementedProfileServiceServer) GetProfileImage(context.Context, *ProfileRequest) (*ProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileImage not implemented")
}
func (UnimplementedProfileServiceServer) NewProfileImage(context.Context, *NewProfileRequest) (*NewProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewProfileImage not implemented")
}
func (UnimplementedProfileServiceServer) mustEmbedUnimplementedProfileServiceServer() {}

// UnsafeProfileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServiceServer will
// result in compilation errors.
type UnsafeProfileServiceServer interface {
	mustEmbedUnimplementedProfileServiceServer()
}

func RegisterProfileServiceServer(s grpc.ServiceRegistrar, srv ProfileServiceServer) {
	s.RegisterService(&ProfileService_ServiceDesc, srv)
}

func _ProfileService_GetProfileImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetProfileImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.ProfileService/GetProfileImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetProfileImage(ctx, req.(*ProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_NewProfileImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).NewProfileImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.ProfileService/NewProfileImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).NewProfileImage(ctx, req.(*NewProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfileService_ServiceDesc is the grpc.ServiceDesc for ProfileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage.ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfileImage",
			Handler:    _ProfileService_GetProfileImage_Handler,
		},
		{
			MethodName: "NewProfileImage",
			Handler:    _ProfileService_NewProfileImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage.proto",
}

// HoroSvcServiceClient is the client API for HoroSvcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HoroSvcServiceClient interface {
	NewHoroServiceImage(ctx context.Context, opts ...grpc.CallOption) (HoroSvcService_NewHoroServiceImageClient, error)
	GetHoroServiceImage(ctx context.Context, in *GetHoroServiceRequest, opts ...grpc.CallOption) (HoroSvcService_GetHoroServiceImageClient, error)
}

type horoSvcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHoroSvcServiceClient(cc grpc.ClientConnInterface) HoroSvcServiceClient {
	return &horoSvcServiceClient{cc}
}

func (c *horoSvcServiceClient) NewHoroServiceImage(ctx context.Context, opts ...grpc.CallOption) (HoroSvcService_NewHoroServiceImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &HoroSvcService_ServiceDesc.Streams[0], "/storage.HoroSvcService/NewHoroServiceImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &horoSvcServiceNewHoroServiceImageClient{stream}
	return x, nil
}

type HoroSvcService_NewHoroServiceImageClient interface {
	Send(*NewHoroServiceRequest) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type horoSvcServiceNewHoroServiceImageClient struct {
	grpc.ClientStream
}

func (x *horoSvcServiceNewHoroServiceImageClient) Send(m *NewHoroServiceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *horoSvcServiceNewHoroServiceImageClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *horoSvcServiceClient) GetHoroServiceImage(ctx context.Context, in *GetHoroServiceRequest, opts ...grpc.CallOption) (HoroSvcService_GetHoroServiceImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &HoroSvcService_ServiceDesc.Streams[1], "/storage.HoroSvcService/GetHoroServiceImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &horoSvcServiceGetHoroServiceImageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HoroSvcService_GetHoroServiceImageClient interface {
	Recv() (*HoroServiceImageResponse, error)
	grpc.ClientStream
}

type horoSvcServiceGetHoroServiceImageClient struct {
	grpc.ClientStream
}

func (x *horoSvcServiceGetHoroServiceImageClient) Recv() (*HoroServiceImageResponse, error) {
	m := new(HoroServiceImageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HoroSvcServiceServer is the server API for HoroSvcService service.
// All implementations must embed UnimplementedHoroSvcServiceServer
// for forward compatibility
type HoroSvcServiceServer interface {
	NewHoroServiceImage(HoroSvcService_NewHoroServiceImageServer) error
	GetHoroServiceImage(*GetHoroServiceRequest, HoroSvcService_GetHoroServiceImageServer) error
	mustEmbedUnimplementedHoroSvcServiceServer()
}

// UnimplementedHoroSvcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHoroSvcServiceServer struct {
}

func (UnimplementedHoroSvcServiceServer) NewHoroServiceImage(HoroSvcService_NewHoroServiceImageServer) error {
	return status.Errorf(codes.Unimplemented, "method NewHoroServiceImage not implemented")
}
func (UnimplementedHoroSvcServiceServer) GetHoroServiceImage(*GetHoroServiceRequest, HoroSvcService_GetHoroServiceImageServer) error {
	return status.Errorf(codes.Unimplemented, "method GetHoroServiceImage not implemented")
}
func (UnimplementedHoroSvcServiceServer) mustEmbedUnimplementedHoroSvcServiceServer() {}

// UnsafeHoroSvcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HoroSvcServiceServer will
// result in compilation errors.
type UnsafeHoroSvcServiceServer interface {
	mustEmbedUnimplementedHoroSvcServiceServer()
}

func RegisterHoroSvcServiceServer(s grpc.ServiceRegistrar, srv HoroSvcServiceServer) {
	s.RegisterService(&HoroSvcService_ServiceDesc, srv)
}

func _HoroSvcService_NewHoroServiceImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HoroSvcServiceServer).NewHoroServiceImage(&horoSvcServiceNewHoroServiceImageServer{stream})
}

type HoroSvcService_NewHoroServiceImageServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*NewHoroServiceRequest, error)
	grpc.ServerStream
}

type horoSvcServiceNewHoroServiceImageServer struct {
	grpc.ServerStream
}

func (x *horoSvcServiceNewHoroServiceImageServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *horoSvcServiceNewHoroServiceImageServer) Recv() (*NewHoroServiceRequest, error) {
	m := new(NewHoroServiceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HoroSvcService_GetHoroServiceImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetHoroServiceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HoroSvcServiceServer).GetHoroServiceImage(m, &horoSvcServiceGetHoroServiceImageServer{stream})
}

type HoroSvcService_GetHoroServiceImageServer interface {
	Send(*HoroServiceImageResponse) error
	grpc.ServerStream
}

type horoSvcServiceGetHoroServiceImageServer struct {
	grpc.ServerStream
}

func (x *horoSvcServiceGetHoroServiceImageServer) Send(m *HoroServiceImageResponse) error {
	return x.ServerStream.SendMsg(m)
}

// HoroSvcService_ServiceDesc is the grpc.ServiceDesc for HoroSvcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HoroSvcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage.HoroSvcService",
	HandlerType: (*HoroSvcServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NewHoroServiceImage",
			Handler:       _HoroSvcService_NewHoroServiceImage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetHoroServiceImage",
			Handler:       _HoroSvcService_GetHoroServiceImage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "storage.proto",
}