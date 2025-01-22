// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user

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

// LoginServiceClient is the client API for LoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginServiceClient interface {
	GetUserInfo(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type loginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginServiceClient(cc grpc.ClientConnInterface) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) GetUserInfo(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/LoginService/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServiceServer is the server API for LoginService service.
// All implementations should embed UnimplementedLoginServiceServer
// for forward compatibility
type LoginServiceServer interface {
	GetUserInfo(context.Context, *LoginRequest) (*LoginResponse, error)
}

// UnimplementedLoginServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLoginServiceServer struct {
}

func (UnimplementedLoginServiceServer) GetUserInfo(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}

// UnsafeLoginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginServiceServer will
// result in compilation errors.
type UnsafeLoginServiceServer interface {
	mustEmbedUnimplementedLoginServiceServer()
}

func RegisterLoginServiceServer(s grpc.ServiceRegistrar, srv LoginServiceServer) {
	s.RegisterService(&LoginService_ServiceDesc, srv)
}

func _LoginService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LoginService/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).GetUserInfo(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LoginService_ServiceDesc is the grpc.ServiceDesc for LoginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _LoginService_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}
