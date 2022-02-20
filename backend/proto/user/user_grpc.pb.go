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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// User
	GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	DeleteUserById(ctx context.Context, in *DeleteUserByIdRequest, opts ...grpc.CallOption) (*DeleteUserByIdResponse, error)
	// GroupUser
	GetGroupUserByUserIdAndGroupId(ctx context.Context, in *GetGroupUserByUserIdAndGroupIdRequest, opts ...grpc.CallOption) (*GetGroupUserByUserIdAndGroupIdResponse, error)
	CreateGroupUser(ctx context.Context, in *CreateGroupUserRequest, opts ...grpc.CallOption) (*CreateGroupUserResponse, error)
	DeleteGroupUserByUserId(ctx context.Context, in *DeleteGroupUserByUserIdRequest, opts ...grpc.CallOption) (*DeleteGroupUserByUserIdResponse, error)
	DeleteGroupUserByGroupId(ctx context.Context, in *DeleteGroupUserByGroupIdRequest, opts ...grpc.CallOption) (*DeleteGroupUserByGroupIdResponse, error)
	// Group
	GetGroupById(ctx context.Context, in *GetGroupByIdRequest, opts ...grpc.CallOption) (*GetGroupByIdResponse, error)
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error)
	UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*UpdateGroupResponse, error)
	DeleteGroupById(ctx context.Context, in *DeleteGroupByIdRequest, opts ...grpc.CallOption) (*DeleteGroupByIdResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error) {
	out := new(GetUserByIdResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUserById(ctx context.Context, in *DeleteUserByIdRequest, opts ...grpc.CallOption) (*DeleteUserByIdResponse, error) {
	out := new(DeleteUserByIdResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/DeleteUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetGroupUserByUserIdAndGroupId(ctx context.Context, in *GetGroupUserByUserIdAndGroupIdRequest, opts ...grpc.CallOption) (*GetGroupUserByUserIdAndGroupIdResponse, error) {
	out := new(GetGroupUserByUserIdAndGroupIdResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetGroupUserByUserIdAndGroupId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateGroupUser(ctx context.Context, in *CreateGroupUserRequest, opts ...grpc.CallOption) (*CreateGroupUserResponse, error) {
	out := new(CreateGroupUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/CreateGroupUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteGroupUserByUserId(ctx context.Context, in *DeleteGroupUserByUserIdRequest, opts ...grpc.CallOption) (*DeleteGroupUserByUserIdResponse, error) {
	out := new(DeleteGroupUserByUserIdResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/DeleteGroupUserByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteGroupUserByGroupId(ctx context.Context, in *DeleteGroupUserByGroupIdRequest, opts ...grpc.CallOption) (*DeleteGroupUserByGroupIdResponse, error) {
	out := new(DeleteGroupUserByGroupIdResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/DeleteGroupUserByGroupId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetGroupById(ctx context.Context, in *GetGroupByIdRequest, opts ...grpc.CallOption) (*GetGroupByIdResponse, error) {
	out := new(GetGroupByIdResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetGroupById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error) {
	out := new(CreateGroupResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*UpdateGroupResponse, error) {
	out := new(UpdateGroupResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteGroupById(ctx context.Context, in *DeleteGroupByIdRequest, opts ...grpc.CallOption) (*DeleteGroupByIdResponse, error) {
	out := new(DeleteGroupByIdResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/DeleteGroupById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// User
	GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUserById(context.Context, *DeleteUserByIdRequest) (*DeleteUserByIdResponse, error)
	// GroupUser
	GetGroupUserByUserIdAndGroupId(context.Context, *GetGroupUserByUserIdAndGroupIdRequest) (*GetGroupUserByUserIdAndGroupIdResponse, error)
	CreateGroupUser(context.Context, *CreateGroupUserRequest) (*CreateGroupUserResponse, error)
	DeleteGroupUserByUserId(context.Context, *DeleteGroupUserByUserIdRequest) (*DeleteGroupUserByUserIdResponse, error)
	DeleteGroupUserByGroupId(context.Context, *DeleteGroupUserByGroupIdRequest) (*DeleteGroupUserByGroupIdResponse, error)
	// Group
	GetGroupById(context.Context, *GetGroupByIdRequest) (*GetGroupByIdResponse, error)
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error)
	UpdateGroup(context.Context, *UpdateGroupRequest) (*UpdateGroupResponse, error)
	DeleteGroupById(context.Context, *DeleteGroupByIdRequest) (*DeleteGroupByIdResponse, error)
}

// UnimplementedUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) DeleteUserById(context.Context, *DeleteUserByIdRequest) (*DeleteUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserById not implemented")
}
func (UnimplementedUserServiceServer) GetGroupUserByUserIdAndGroupId(context.Context, *GetGroupUserByUserIdAndGroupIdRequest) (*GetGroupUserByUserIdAndGroupIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupUserByUserIdAndGroupId not implemented")
}
func (UnimplementedUserServiceServer) CreateGroupUser(context.Context, *CreateGroupUserRequest) (*CreateGroupUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroupUser not implemented")
}
func (UnimplementedUserServiceServer) DeleteGroupUserByUserId(context.Context, *DeleteGroupUserByUserIdRequest) (*DeleteGroupUserByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroupUserByUserId not implemented")
}
func (UnimplementedUserServiceServer) DeleteGroupUserByGroupId(context.Context, *DeleteGroupUserByGroupIdRequest) (*DeleteGroupUserByGroupIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroupUserByGroupId not implemented")
}
func (UnimplementedUserServiceServer) GetGroupById(context.Context, *GetGroupByIdRequest) (*GetGroupByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupById not implemented")
}
func (UnimplementedUserServiceServer) CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedUserServiceServer) UpdateGroup(context.Context, *UpdateGroupRequest) (*UpdateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}
func (UnimplementedUserServiceServer) DeleteGroupById(context.Context, *DeleteGroupByIdRequest) (*DeleteGroupByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroupById not implemented")
}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserById(ctx, req.(*GetUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/DeleteUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUserById(ctx, req.(*DeleteUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetGroupUserByUserIdAndGroupId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupUserByUserIdAndGroupIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetGroupUserByUserIdAndGroupId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetGroupUserByUserIdAndGroupId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetGroupUserByUserIdAndGroupId(ctx, req.(*GetGroupUserByUserIdAndGroupIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateGroupUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateGroupUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/CreateGroupUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateGroupUser(ctx, req.(*CreateGroupUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteGroupUserByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupUserByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteGroupUserByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/DeleteGroupUserByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteGroupUserByUserId(ctx, req.(*DeleteGroupUserByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteGroupUserByGroupId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupUserByGroupIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteGroupUserByGroupId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/DeleteGroupUserByGroupId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteGroupUserByGroupId(ctx, req.(*DeleteGroupUserByGroupIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetGroupById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetGroupById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetGroupById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetGroupById(ctx, req.(*GetGroupByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateGroup(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteGroupById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteGroupById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/DeleteGroupById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteGroupById(ctx, req.(*DeleteGroupByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserById",
			Handler:    _UserService_GetUserById_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUserById",
			Handler:    _UserService_DeleteUserById_Handler,
		},
		{
			MethodName: "GetGroupUserByUserIdAndGroupId",
			Handler:    _UserService_GetGroupUserByUserIdAndGroupId_Handler,
		},
		{
			MethodName: "CreateGroupUser",
			Handler:    _UserService_CreateGroupUser_Handler,
		},
		{
			MethodName: "DeleteGroupUserByUserId",
			Handler:    _UserService_DeleteGroupUserByUserId_Handler,
		},
		{
			MethodName: "DeleteGroupUserByGroupId",
			Handler:    _UserService_DeleteGroupUserByGroupId_Handler,
		},
		{
			MethodName: "GetGroupById",
			Handler:    _UserService_GetGroupById_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _UserService_CreateGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _UserService_UpdateGroup_Handler,
		},
		{
			MethodName: "DeleteGroupById",
			Handler:    _UserService_DeleteGroupById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}