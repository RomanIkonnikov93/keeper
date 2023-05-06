// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: keeper/server/internal/proto/keeper.proto

package keeper

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserLogin    string `protobuf:"bytes,1,opt,name=user_login,json=userLogin,proto3" json:"user_login,omitempty"`
	UserPassword string `protobuf:"bytes,2,opt,name=user_password,json=userPassword,proto3" json:"user_password,omitempty"`
	UserId       string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserToken    string `protobuf:"bytes,4,opt,name=user_token,json=userToken,proto3" json:"user_token,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keeper_server_internal_proto_keeper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_keeper_server_internal_proto_keeper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_keeper_server_internal_proto_keeper_proto_rawDescGZIP(), []int{0}
}

func (x *Auth) GetUserLogin() string {
	if x != nil {
		return x.UserLogin
	}
	return ""
}

func (x *Auth) GetUserPassword() string {
	if x != nil {
		return x.UserPassword
	}
	return ""
}

func (x *Auth) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Auth) GetUserToken() string {
	if x != nil {
		return x.UserToken
	}
	return ""
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordID    int32  `protobuf:"varint,1,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
	RecordType  string `protobuf:"bytes,2,opt,name=record_type,json=recordType,proto3" json:"record_type,omitempty"`
	UserID      string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Metadata    string `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Login       string `protobuf:"bytes,6,opt,name=login,proto3" json:"login,omitempty"`
	Password    string `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
	Card        string `protobuf:"bytes,8,opt,name=card,proto3" json:"card,omitempty"`
	File        []byte `protobuf:"bytes,9,opt,name=file,proto3" json:"file,omitempty"`
	CreatedAt   string `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keeper_server_internal_proto_keeper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_keeper_server_internal_proto_keeper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_keeper_server_internal_proto_keeper_proto_rawDescGZIP(), []int{1}
}

func (x *Record) GetRecordId() int32 {
	if x != nil {
		return x.RecordID
	}
	return 0
}

func (x *Record) GetRecordType() string {
	if x != nil {
		return x.RecordType
	}
	return ""
}

func (x *Record) GetUserId() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Record) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Record) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *Record) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *Record) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Record) GetCard() string {
	if x != nil {
		return x.Card
	}
	return ""
}

func (x *Record) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *Record) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Note []*Record `protobuf:"bytes,1,rep,name=note,proto3" json:"note,omitempty"`
}

func (x *List) Reset() {
	*x = List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keeper_server_internal_proto_keeper_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*List) ProtoMessage() {}

func (x *List) ProtoReflect() protoreflect.Message {
	mi := &file_keeper_server_internal_proto_keeper_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use List.ProtoReflect.Descriptor instead.
func (*List) Descriptor() ([]byte, []int) {
	return file_keeper_server_internal_proto_keeper_proto_rawDescGZIP(), []int{2}
}

func (x *List) GetNote() []*Record {
	if x != nil {
		return x.Note
	}
	return nil
}

var File_keeper_server_internal_proto_keeper_proto protoreflect.FileDescriptor

var file_keeper_server_internal_proto_keeper_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b,
	0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x82, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x75, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x96, 0x02, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x61, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2a,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x32, 0xf0, 0x02, 0x0a, 0x06, 0x4b,
	0x65, 0x65, 0x70, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0c, 0x2e, 0x6b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x0c, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x12, 0x27, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0c, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x1a, 0x0c, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x12, 0x33,
	0x0a, 0x09, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x2e, 0x6b, 0x65,
	0x65, 0x70, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x2b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x0e, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x1a, 0x0e, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x33, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x0c, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0e, 0x2e, 0x6b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x3a, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0e, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x24, 0x5a,
	0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x6d, 0x61,
	0x6e, 0x49, 0x6b, 0x6f, 0x6e, 0x6e, 0x69, 0x6b, 0x6f, 0x76, 0x39, 0x33, 0x2f, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_keeper_server_internal_proto_keeper_proto_rawDescOnce sync.Once
	file_keeper_server_internal_proto_keeper_proto_rawDescData = file_keeper_server_internal_proto_keeper_proto_rawDesc
)

func file_keeper_server_internal_proto_keeper_proto_rawDescGZIP() []byte {
	file_keeper_server_internal_proto_keeper_proto_rawDescOnce.Do(func() {
		file_keeper_server_internal_proto_keeper_proto_rawDescData = protoimpl.X.CompressGZIP(file_keeper_server_internal_proto_keeper_proto_rawDescData)
	})
	return file_keeper_server_internal_proto_keeper_proto_rawDescData
}

var file_keeper_server_internal_proto_keeper_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_keeper_server_internal_proto_keeper_proto_goTypes = []interface{}{
	(*Auth)(nil),          // 0: keeper.Auth
	(*Record)(nil),        // 1: keeper.Record
	(*List)(nil),          // 2: keeper.List
	(*emptypb.Empty)(nil), // 3: google.protobuf.Empty
}
var file_keeper_server_internal_proto_keeper_proto_depIdxs = []int32{
	1, // 0: keeper.List.note:type_name -> keeper.Record
	0, // 1: keeper.Keeper.RegistrationUser:input_type -> keeper.Auth
	0, // 2: keeper.Keeper.LoginUser:input_type -> keeper.Auth
	1, // 3: keeper.Keeper.AddRecord:input_type -> keeper.Record
	1, // 4: keeper.Keeper.GetRecordByID:input_type -> keeper.Record
	1, // 5: keeper.Keeper.GetAllRecordsByType:input_type -> keeper.Record
	1, // 6: keeper.Keeper.UpdateRecordByID:input_type -> keeper.Record
	1, // 7: keeper.Keeper.DeleteRecordByID:input_type -> keeper.Record
	0, // 8: keeper.Keeper.RegistrationUser:output_type -> keeper.Auth
	0, // 9: keeper.Keeper.LoginUser:output_type -> keeper.Auth
	3, // 10: keeper.Keeper.AddRecord:output_type -> google.protobuf.Empty
	1, // 11: keeper.Keeper.GetRecordByID:output_type -> keeper.Record
	2, // 12: keeper.Keeper.GetAllRecordsByType:output_type -> keeper.List
	3, // 13: keeper.Keeper.UpdateRecordByID:output_type -> google.protobuf.Empty
	3, // 14: keeper.Keeper.DeleteRecordByID:output_type -> google.protobuf.Empty
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_keeper_server_internal_proto_keeper_proto_init() }
func file_keeper_server_internal_proto_keeper_proto_init() {
	if File_keeper_server_internal_proto_keeper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_keeper_server_internal_proto_keeper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_keeper_server_internal_proto_keeper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_keeper_server_internal_proto_keeper_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*List); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_keeper_server_internal_proto_keeper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_keeper_server_internal_proto_keeper_proto_goTypes,
		DependencyIndexes: file_keeper_server_internal_proto_keeper_proto_depIdxs,
		MessageInfos:      file_keeper_server_internal_proto_keeper_proto_msgTypes,
	}.Build()
	File_keeper_server_internal_proto_keeper_proto = out.File
	file_keeper_server_internal_proto_keeper_proto_rawDesc = nil
	file_keeper_server_internal_proto_keeper_proto_goTypes = nil
	file_keeper_server_internal_proto_keeper_proto_depIdxs = nil
}
