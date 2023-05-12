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
	Keeper_RegistrationUser_FullMethodName    = "/keeper.Keeper/RegistrationUser"
	Keeper_LoginUser_FullMethodName           = "/keeper.Keeper/LoginUser"
	Keeper_AddRecord_FullMethodName           = "/keeper.Keeper/AddRecord"
	Keeper_GetRecordByID_FullMethodName       = "/keeper.Keeper/GetRecordByID"
	Keeper_GetAllRecordsByType_FullMethodName = "/keeper.Keeper/GetAllRecordsByType"
	Keeper_UpdateRecordByID_FullMethodName    = "/keeper.Keeper/UpdateRecordByID"
	Keeper_DeleteRecordByID_FullMethodName    = "/keeper.Keeper/DeleteRecordByID"
	Keeper_Ping_FullMethodName                = "/keeper.Keeper/Ping"
	Keeper_CheckChanges_FullMethodName        = "/keeper.Keeper/CheckChanges"
)

// KeeperClient is the client API for Keeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeeperClient interface {
	RegistrationUser(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Auth, error)
	LoginUser(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Auth, error)
	AddRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetRecordByID(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Record, error)
	GetAllRecordsByType(ctx context.Context, in *Record, opts ...grpc.CallOption) (*List, error)
	UpdateRecordByID(ctx context.Context, in *Record, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteRecordByID(ctx context.Context, in *Record, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CheckChanges(ctx context.Context, in *Record, opts ...grpc.CallOption) (*List, error)
}

type keeperClient struct {
	cc grpc.ClientConnInterface
}

func NewKeeperClient(cc grpc.ClientConnInterface) KeeperClient {
	return &keeperClient{cc}
}

func (c *keeperClient) RegistrationUser(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Auth, error) {
	out := new(Auth)
	err := c.cc.Invoke(ctx, Keeper_RegistrationUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperClient) LoginUser(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Auth, error) {
	out := new(Auth)
	err := c.cc.Invoke(ctx, Keeper_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperClient) AddRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Keeper_AddRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperClient) GetRecordByID(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Record, error) {
	out := new(Record)
	err := c.cc.Invoke(ctx, Keeper_GetRecordByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperClient) GetAllRecordsByType(ctx context.Context, in *Record, opts ...grpc.CallOption) (*List, error) {
	out := new(List)
	err := c.cc.Invoke(ctx, Keeper_GetAllRecordsByType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperClient) UpdateRecordByID(ctx context.Context, in *Record, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Keeper_UpdateRecordByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperClient) DeleteRecordByID(ctx context.Context, in *Record, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Keeper_DeleteRecordByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Keeper_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperClient) CheckChanges(ctx context.Context, in *Record, opts ...grpc.CallOption) (*List, error) {
	out := new(List)
	err := c.cc.Invoke(ctx, Keeper_CheckChanges_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeeperServer is the server API for Keeper service.
// All implementations must embed UnimplementedKeeperServer
// for forward compatibility
type KeeperServer interface {
	RegistrationUser(context.Context, *Auth) (*Auth, error)
	LoginUser(context.Context, *Auth) (*Auth, error)
	AddRecord(context.Context, *Record) (*emptypb.Empty, error)
	GetRecordByID(context.Context, *Record) (*Record, error)
	GetAllRecordsByType(context.Context, *Record) (*List, error)
	UpdateRecordByID(context.Context, *Record) (*emptypb.Empty, error)
	DeleteRecordByID(context.Context, *Record) (*emptypb.Empty, error)
	Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	CheckChanges(context.Context, *Record) (*List, error)
	mustEmbedUnimplementedKeeperServer()
}

// UnimplementedKeeperServer must be embedded to have forward compatible implementations.
type UnimplementedKeeperServer struct {
}

func (UnimplementedKeeperServer) RegistrationUser(context.Context, *Auth) (*Auth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegistrationUser not implemented")
}
func (UnimplementedKeeperServer) LoginUser(context.Context, *Auth) (*Auth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedKeeperServer) AddRecord(context.Context, *Record) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRecord not implemented")
}
func (UnimplementedKeeperServer) GetRecordByID(context.Context, *Record) (*Record, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecordByID not implemented")
}
func (UnimplementedKeeperServer) GetAllRecordsByType(context.Context, *Record) (*List, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRecordsByType not implemented")
}
func (UnimplementedKeeperServer) UpdateRecordByID(context.Context, *Record) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecordByID not implemented")
}
func (UnimplementedKeeperServer) DeleteRecordByID(context.Context, *Record) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecordByID not implemented")
}
func (UnimplementedKeeperServer) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedKeeperServer) CheckChanges(context.Context, *Record) (*List, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckChanges not implemented")
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

func _Keeper_RegistrationUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).RegistrationUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_RegistrationUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).RegistrationUser(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keeper_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).LoginUser(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keeper_AddRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).AddRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_AddRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).AddRecord(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keeper_GetRecordByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).GetRecordByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_GetRecordByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).GetRecordByID(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keeper_GetAllRecordsByType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).GetAllRecordsByType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_GetAllRecordsByType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).GetAllRecordsByType(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keeper_UpdateRecordByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).UpdateRecordByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_UpdateRecordByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).UpdateRecordByID(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keeper_DeleteRecordByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).DeleteRecordByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_DeleteRecordByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).DeleteRecordByID(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keeper_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keeper_CheckChanges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).CheckChanges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_CheckChanges_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).CheckChanges(ctx, req.(*Record))
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
			MethodName: "RegistrationUser",
			Handler:    _Keeper_RegistrationUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _Keeper_LoginUser_Handler,
		},
		{
			MethodName: "AddRecord",
			Handler:    _Keeper_AddRecord_Handler,
		},
		{
			MethodName: "GetRecordByID",
			Handler:    _Keeper_GetRecordByID_Handler,
		},
		{
			MethodName: "GetAllRecordsByType",
			Handler:    _Keeper_GetAllRecordsByType_Handler,
		},
		{
			MethodName: "UpdateRecordByID",
			Handler:    _Keeper_UpdateRecordByID_Handler,
		},
		{
			MethodName: "DeleteRecordByID",
			Handler:    _Keeper_DeleteRecordByID_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Keeper_Ping_Handler,
		},
		{
			MethodName: "CheckChanges",
			Handler:    _Keeper_CheckChanges_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keeper/server/internal/proto/keeper.proto",
}