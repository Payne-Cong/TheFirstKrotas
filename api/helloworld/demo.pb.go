// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: helloworld/demo.proto

package helloworld

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_SUCCESS Status = 0
	Status_FAIL    Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "SUCCESS",
		1: "FAIL",
	}
	Status_value = map[string]int32{
		"SUCCESS": 0,
		"FAIL":    1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_helloworld_demo_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_helloworld_demo_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_helloworld_demo_proto_rawDescGZIP(), []int{0}
}

type CreateDemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *CreateDemoRequest) Reset() {
	*x = CreateDemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDemoRequest) ProtoMessage() {}

func (x *CreateDemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDemoRequest.ProtoReflect.Descriptor instead.
func (*CreateDemoRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_demo_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDemoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateDemoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDemoRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type CreateDemoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=api.helloworld.Status" json:"status,omitempty"`
}

func (x *CreateDemoReply) Reset() {
	*x = CreateDemoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDemoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDemoReply) ProtoMessage() {}

func (x *CreateDemoReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDemoReply.ProtoReflect.Descriptor instead.
func (*CreateDemoReply) Descriptor() ([]byte, []int) {
	return file_helloworld_demo_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDemoReply) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_SUCCESS
}

type UpdateDemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateDemoRequest) Reset() {
	*x = UpdateDemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_demo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDemoRequest) ProtoMessage() {}

func (x *UpdateDemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_demo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDemoRequest.ProtoReflect.Descriptor instead.
func (*UpdateDemoRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_demo_proto_rawDescGZIP(), []int{2}
}

type UpdateDemoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateDemoReply) Reset() {
	*x = UpdateDemoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_demo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDemoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDemoReply) ProtoMessage() {}

func (x *UpdateDemoReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_demo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDemoReply.ProtoReflect.Descriptor instead.
func (*UpdateDemoReply) Descriptor() ([]byte, []int) {
	return file_helloworld_demo_proto_rawDescGZIP(), []int{3}
}

type DeleteDemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDemoRequest) Reset() {
	*x = DeleteDemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_demo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDemoRequest) ProtoMessage() {}

func (x *DeleteDemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_demo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDemoRequest.ProtoReflect.Descriptor instead.
func (*DeleteDemoRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_demo_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteDemoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteDemoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=api.helloworld.Status" json:"status,omitempty"`
}

func (x *DeleteDemoReply) Reset() {
	*x = DeleteDemoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_demo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDemoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDemoReply) ProtoMessage() {}

func (x *DeleteDemoReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_demo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDemoReply.ProtoReflect.Descriptor instead.
func (*DeleteDemoReply) Descriptor() ([]byte, []int) {
	return file_helloworld_demo_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteDemoReply) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_SUCCESS
}

type GetDemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDemoRequest) Reset() {
	*x = GetDemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_demo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDemoRequest) ProtoMessage() {}

func (x *GetDemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_demo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDemoRequest.ProtoReflect.Descriptor instead.
func (*GetDemoRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_demo_proto_rawDescGZIP(), []int{6}
}

type GetDemoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res string `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *GetDemoReply) Reset() {
	*x = GetDemoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_demo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDemoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDemoReply) ProtoMessage() {}

func (x *GetDemoReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_demo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDemoReply.ProtoReflect.Descriptor instead.
func (*GetDemoReply) Descriptor() ([]byte, []int) {
	return file_helloworld_demo_proto_rawDescGZIP(), []int{7}
}

func (x *GetDemoReply) GetRes() string {
	if x != nil {
		return x.Res
	}
	return ""
}

type ListDemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListDemoRequest) Reset() {
	*x = ListDemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_demo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDemoRequest) ProtoMessage() {}

func (x *ListDemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_demo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDemoRequest.ProtoReflect.Descriptor instead.
func (*ListDemoRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_demo_proto_rawDescGZIP(), []int{8}
}

type ListDemoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListDemoReply) Reset() {
	*x = ListDemoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_demo_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDemoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDemoReply) ProtoMessage() {}

func (x *ListDemoReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_demo_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDemoReply.ProtoReflect.Descriptor instead.
func (*ListDemoReply) Descriptor() ([]byte, []int) {
	return file_helloworld_demo_proto_rawDescGZIP(), []int{9}
}

var File_helloworld_demo_proto protoreflect.FileDescriptor

var file_helloworld_demo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x64, 0x65, 0x6d,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x41, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x11, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x6d,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x20, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x73, 0x22,
	0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x2a, 0x1f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41,
	0x49, 0x4c, 0x10, 0x01, 0x32, 0xa1, 0x03, 0x0a, 0x04, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x50, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x21, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x50, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x21, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x50, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x12,
	0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x57, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x1e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x4a, 0x0a, 0x08,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65,
	0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x38, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x01, 0x5a, 0x24, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x3b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helloworld_demo_proto_rawDescOnce sync.Once
	file_helloworld_demo_proto_rawDescData = file_helloworld_demo_proto_rawDesc
)

func file_helloworld_demo_proto_rawDescGZIP() []byte {
	file_helloworld_demo_proto_rawDescOnce.Do(func() {
		file_helloworld_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_demo_proto_rawDescData)
	})
	return file_helloworld_demo_proto_rawDescData
}

var file_helloworld_demo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_helloworld_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_helloworld_demo_proto_goTypes = []interface{}{
	(Status)(0),               // 0: api.helloworld.Status
	(*CreateDemoRequest)(nil), // 1: api.helloworld.CreateDemoRequest
	(*CreateDemoReply)(nil),   // 2: api.helloworld.CreateDemoReply
	(*UpdateDemoRequest)(nil), // 3: api.helloworld.UpdateDemoRequest
	(*UpdateDemoReply)(nil),   // 4: api.helloworld.UpdateDemoReply
	(*DeleteDemoRequest)(nil), // 5: api.helloworld.DeleteDemoRequest
	(*DeleteDemoReply)(nil),   // 6: api.helloworld.DeleteDemoReply
	(*GetDemoRequest)(nil),    // 7: api.helloworld.GetDemoRequest
	(*GetDemoReply)(nil),      // 8: api.helloworld.GetDemoReply
	(*ListDemoRequest)(nil),   // 9: api.helloworld.ListDemoRequest
	(*ListDemoReply)(nil),     // 10: api.helloworld.ListDemoReply
}
var file_helloworld_demo_proto_depIdxs = []int32{
	0,  // 0: api.helloworld.CreateDemoReply.status:type_name -> api.helloworld.Status
	0,  // 1: api.helloworld.DeleteDemoReply.status:type_name -> api.helloworld.Status
	1,  // 2: api.helloworld.Demo.CreateDemo:input_type -> api.helloworld.CreateDemoRequest
	3,  // 3: api.helloworld.Demo.UpdateDemo:input_type -> api.helloworld.UpdateDemoRequest
	5,  // 4: api.helloworld.Demo.DeleteDemo:input_type -> api.helloworld.DeleteDemoRequest
	7,  // 5: api.helloworld.Demo.GetDemo:input_type -> api.helloworld.GetDemoRequest
	9,  // 6: api.helloworld.Demo.ListDemo:input_type -> api.helloworld.ListDemoRequest
	2,  // 7: api.helloworld.Demo.CreateDemo:output_type -> api.helloworld.CreateDemoReply
	4,  // 8: api.helloworld.Demo.UpdateDemo:output_type -> api.helloworld.UpdateDemoReply
	6,  // 9: api.helloworld.Demo.DeleteDemo:output_type -> api.helloworld.DeleteDemoReply
	8,  // 10: api.helloworld.Demo.GetDemo:output_type -> api.helloworld.GetDemoReply
	10, // 11: api.helloworld.Demo.ListDemo:output_type -> api.helloworld.ListDemoReply
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_helloworld_demo_proto_init() }
func file_helloworld_demo_proto_init() {
	if File_helloworld_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDemoRequest); i {
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
		file_helloworld_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDemoReply); i {
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
		file_helloworld_demo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDemoRequest); i {
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
		file_helloworld_demo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDemoReply); i {
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
		file_helloworld_demo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDemoRequest); i {
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
		file_helloworld_demo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDemoReply); i {
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
		file_helloworld_demo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDemoRequest); i {
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
		file_helloworld_demo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDemoReply); i {
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
		file_helloworld_demo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDemoRequest); i {
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
		file_helloworld_demo_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDemoReply); i {
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
			RawDescriptor: file_helloworld_demo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helloworld_demo_proto_goTypes,
		DependencyIndexes: file_helloworld_demo_proto_depIdxs,
		EnumInfos:         file_helloworld_demo_proto_enumTypes,
		MessageInfos:      file_helloworld_demo_proto_msgTypes,
	}.Build()
	File_helloworld_demo_proto = out.File
	file_helloworld_demo_proto_rawDesc = nil
	file_helloworld_demo_proto_goTypes = nil
	file_helloworld_demo_proto_depIdxs = nil
}
