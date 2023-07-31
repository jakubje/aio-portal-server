// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: service_aio_portal.proto

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

const (
	AioPortal_CreateUser_FullMethodName = "/pb.AioPortal/CreateUser"
	AioPortal_LoginUser_FullMethodName  = "/pb.AioPortal/LoginUser"
	AioPortal_UpdateUser_FullMethodName = "/pb.AioPortal/UpdateUser"
)

// AioPortalClient is the client API for AioPortal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AioPortalClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
}

type aioPortalClient struct {
	cc grpc.ClientConnInterface
}

func NewAioPortalClient(cc grpc.ClientConnInterface) AioPortalClient {
	return &aioPortalClient{cc}
}

func (c *aioPortalClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, AioPortal_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, AioPortal_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, AioPortal_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AioPortalServer is the server API for AioPortal service.
// All implementations must embed UnimplementedAioPortalServer
// for forward compatibility
type AioPortalServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	mustEmbedUnimplementedAioPortalServer()
}

// UnimplementedAioPortalServer must be embedded to have forward compatible implementations.
type UnimplementedAioPortalServer struct {
}

func (UnimplementedAioPortalServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedAioPortalServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedAioPortalServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedAioPortalServer) mustEmbedUnimplementedAioPortalServer() {}

// UnsafeAioPortalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AioPortalServer will
// result in compilation errors.
type UnsafeAioPortalServer interface {
	mustEmbedUnimplementedAioPortalServer()
}

func RegisterAioPortalServer(s grpc.ServiceRegistrar, srv AioPortalServer) {
	s.RegisterService(&AioPortal_ServiceDesc, srv)
}

func _AioPortal_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AioPortal_ServiceDesc is the grpc.ServiceDesc for AioPortal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AioPortal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AioPortal",
	HandlerType: (*AioPortalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _AioPortal_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _AioPortal_LoginUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _AioPortal_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_aio_portal.proto",
}