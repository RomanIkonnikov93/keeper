// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: keeper/server/internal/proto/keeper.proto

package keeper

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

const (
	Keeper_AddUser_FullMethodName      = "/keeper.Keeper/AddUser"
	Keeper_CheckUser_FullMethodName    = "/keeper.Keeper/CheckUser"
	Keeper_Add_FullMethodName          = "/keeper.Keeper/Add"
	Keeper_Get_FullMethodName          = "/keeper.Keeper/Get"
	Keeper_GetAllByType_FullMethodName = "/keeper.Keeper/GetAllByType"
	Keeper_UpdateByID_FullMethodName   = "/keeper.Keeper/UpdateByID"
	Keeper_DeleteByID_FullMethodName   = "/keeper.Keeper/DeleteByID"
)

// KeeperClient is the client API for Keeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeeperClient interface {
	AddUser(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Auth, error)
	CheckUser(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Auth, error)
	Add(ctx context.Context, in *Record, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Get(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Record, error)
	GetAllByType(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Record, error)
	UpdateByID(ctx context.Context, in *Record, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteByID(ctx context.Context, in *Record, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type keeperClient struct {
	cc grpc.ClientConnInterface
}

func NewKeeperClient(cc grpc.ClientConnInterface) KeeperClient {
	return &keeperClient{cc}
}

func (c *keeperClient) AddUser(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Auth, error) {
	out := new(Auth)
	err := c.cc.Invoke(ctx, Keeper_AddUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperClient) CheckUser(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Auth, error) {
	out := new(Auth)
	err := c.cc.Invoke(ctx, Keeper_CheckUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperClient) Add(ctx context.Context, in *Record, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Keeper_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperClient) Get(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Record, error) {
	out := new(Record)
	err := c.cc.Invoke(ctx, Keeper_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperClient) GetAllByType(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Record, error) {
	out := new(Record)
	err := c.cc.Invoke(ctx, Keeper_GetAllByType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperClient) UpdateByID(ctx context.Context, in *Record, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Keeper_UpdateByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperClient) DeleteByID(ctx context.Context, in *Record, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Keeper_DeleteByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeeperServer is the server API for Keeper service.
// All implementations must embed UnimplementedKeeperServer
// for forward compatibility
type KeeperServer interface {
	AddUser(context.Context, *Auth) (*Auth, error)
	CheckUser(context.Context, *Auth) (*Auth, error)
	Add(context.Context, *Record) (*emptypb.Empty, error)
	Get(context.Context, *Record) (*Record, error)
	GetAllByType(context.Context, *Record) (*Record, error)
	UpdateByID(context.Context, *Record) (*emptypb.Empty, error)
	DeleteByID(context.Context, *Record) (*emptypb.Empty, error)
	mustEmbedUnimplementedKeeperServer()
}

// UnimplementedKeeperServer must be embedded to have forward compatible implementations.
type UnimplementedKeeperServer struct {
}

func (UnimplementedKeeperServer) AddUser(context.Context, *Auth) (*Auth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedKeeperServer) CheckUser(context.Context, *Auth) (*Auth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUser not implemented")
}
func (UnimplementedKeeperServer) Add(context.Context, *Record) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedKeeperServer) Get(context.Context, *Record) (*Record, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedKeeperServer) GetAllByType(context.Context, *Record) (*Record, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllByType not implemented")
}
func (UnimplementedKeeperServer) UpdateByID(context.Context, *Record) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateByID not implemented")
}
func (UnimplementedKeeperServer) DeleteByID(context.Context, *Record) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByID not implemented")
}
func (UnimplementedKeeperServer) mustEmbedUnimplementedKeeperServer() {}

// UnsafeKeeperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeeperServer will
// result in compilation errors.
type UnsafeKeeperServer interface {
	mustEmbedUnimplementedKeeperServer()
}

func RegisterKeeperServer(s grpc.ServiceRegistrar, srv KeeperServer) {
	s.RegisterService(&Keeper_ServiceDesc, srv)
}

func _Keeper_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_AddUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).AddUser(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keeper_CheckUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).CheckUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_CheckUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).CheckUser(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keeper_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).Add(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keeper_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).Get(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keeper_GetAllByType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).GetAllByType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_GetAllByType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).GetAllByType(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keeper_UpdateByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).UpdateByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_UpdateByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).UpdateByID(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keeper_DeleteByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).DeleteByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_DeleteByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).DeleteByID(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

// Keeper_ServiceDesc is the grpc.ServiceDesc for Keeper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Keeper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keeper.Keeper",
	HandlerType: (*KeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _Keeper_AddUser_Handler,
		},
		{
			MethodName: "CheckUser",
			Handler:    _Keeper_CheckUser_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Keeper_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Keeper_Get_Handler,
		},
		{
			MethodName: "GetAllByType",
			Handler:    _Keeper_GetAllByType_Handler,
		},
		{
			MethodName: "UpdateByID",
			Handler:    _Keeper_UpdateByID_Handler,
		},
		{
			MethodName: "DeleteByID",
			Handler:    _Keeper_DeleteByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keeper/server/internal/proto/keeper.proto",
}
