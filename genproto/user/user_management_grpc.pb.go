// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: user_management.proto

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

// UserManagementClient is the client API for UserManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserManagementClient interface {
	GetUser(ctx context.Context, in *ID, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Void, error)
	DeleteUser(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Void, error)
	GetProfile(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Profile, error)
	UpdateProfile(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*Void, error)
}

type userManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewUserManagementClient(cc grpc.ClientConnInterface) UserManagementClient {
	return &userManagementClient{cc}
}

func (c *userManagementClient) GetUser(ctx context.Context, in *ID, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/protos.UserManagement/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.UserManagement/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) DeleteUser(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.UserManagement/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) GetProfile(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := c.cc.Invoke(ctx, "/protos.UserManagement/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) UpdateProfile(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.UserManagement/UpdateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserManagementServer is the server API for UserManagement service.
// All implementations must embed UnimplementedUserManagementServer
// for forward compatibility
type UserManagementServer interface {
	GetUser(context.Context, *ID) (*User, error)
	UpdateUser(context.Context, *User) (*Void, error)
	DeleteUser(context.Context, *ID) (*Void, error)
	GetProfile(context.Context, *ID) (*Profile, error)
	UpdateProfile(context.Context, *Profile) (*Void, error)
	mustEmbedUnimplementedUserManagementServer()
}

// UnimplementedUserManagementServer must be embedded to have forward compatible implementations.
type UnimplementedUserManagementServer struct {
}

func (UnimplementedUserManagementServer) GetUser(context.Context, *ID) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserManagementServer) UpdateUser(context.Context, *User) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserManagementServer) DeleteUser(context.Context, *ID) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserManagementServer) GetProfile(context.Context, *ID) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedUserManagementServer) UpdateProfile(context.Context, *Profile) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedUserManagementServer) mustEmbedUnimplementedUserManagementServer() {}

// UnsafeUserManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserManagementServer will
// result in compilation errors.
type UnsafeUserManagementServer interface {
	mustEmbedUnimplementedUserManagementServer()
}

func RegisterUserManagementServer(s grpc.ServiceRegistrar, srv UserManagementServer) {
	s.RegisterService(&UserManagement_ServiceDesc, srv)
}

func _UserManagement_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.UserManagement/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).GetUser(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.UserManagement/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).UpdateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.UserManagement/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).DeleteUser(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.UserManagement/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).GetProfile(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Profile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.UserManagement/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).UpdateProfile(ctx, req.(*Profile))
	}
	return interceptor(ctx, in, info, handler)
}

// UserManagement_ServiceDesc is the grpc.ServiceDesc for UserManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.UserManagement",
	HandlerType: (*UserManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserManagement_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserManagement_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserManagement_DeleteUser_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _UserManagement_GetProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _UserManagement_UpdateProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_management.proto",
}
