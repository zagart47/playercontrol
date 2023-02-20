// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: proto/player.proto

package pb

import (
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

type PlayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlayRequest) Reset() {
	*x = PlayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayRequest) ProtoMessage() {}

func (x *PlayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayRequest.ProtoReflect.Descriptor instead.
func (*PlayRequest) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{0}
}

type PlayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlayResponse) Reset() {
	*x = PlayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayResponse) ProtoMessage() {}

func (x *PlayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayResponse.ProtoReflect.Descriptor instead.
func (*PlayResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{1}
}

type PauseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PauseRequest) Reset() {
	*x = PauseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PauseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PauseRequest) ProtoMessage() {}

func (x *PauseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PauseRequest.ProtoReflect.Descriptor instead.
func (*PauseRequest) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{2}
}

type PauseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PauseResponse) Reset() {
	*x = PauseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PauseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PauseResponse) ProtoMessage() {}

func (x *PauseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PauseResponse.ProtoReflect.Descriptor instead.
func (*PauseResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{3}
}

type NextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NextRequest) Reset() {
	*x = NextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextRequest) ProtoMessage() {}

func (x *NextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextRequest.ProtoReflect.Descriptor instead.
func (*NextRequest) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{4}
}

type NextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NextResponse) Reset() {
	*x = NextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextResponse) ProtoMessage() {}

func (x *NextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextResponse.ProtoReflect.Descriptor instead.
func (*NextResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{5}
}

type PrevRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PrevRequest) Reset() {
	*x = PrevRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrevRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrevRequest) ProtoMessage() {}

func (x *PrevRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrevRequest.ProtoReflect.Descriptor instead.
func (*PrevRequest) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{6}
}

type PrevResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PrevResponse) Reset() {
	*x = PrevResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrevResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrevResponse) ProtoMessage() {}

func (x *PrevResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrevResponse.ProtoReflect.Descriptor instead.
func (*PrevResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{7}
}

type AddSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *AddSongRequest) Reset() {
	*x = AddSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSongRequest) ProtoMessage() {}

func (x *AddSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSongRequest.ProtoReflect.Descriptor instead.
func (*AddSongRequest) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{8}
}

func (x *AddSongRequest) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type AddSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddSongResponse) Reset() {
	*x = AddSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSongResponse) ProtoMessage() {}

func (x *AddSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSongResponse.ProtoReflect.Descriptor instead.
func (*AddSongResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{9}
}

type DeleteSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSongRequest) Reset() {
	*x = DeleteSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSongRequest) ProtoMessage() {}

func (x *DeleteSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSongRequest.ProtoReflect.Descriptor instead.
func (*DeleteSongRequest) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{10}
}

type DeleteSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSongResponse) Reset() {
	*x = DeleteSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSongResponse) ProtoMessage() {}

func (x *DeleteSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSongResponse.ProtoReflect.Descriptor instead.
func (*DeleteSongResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{11}
}

type UpdateSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateSongRequest) Reset() {
	*x = UpdateSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSongRequest) ProtoMessage() {}

func (x *UpdateSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSongRequest.ProtoReflect.Descriptor instead.
func (*UpdateSongRequest) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{12}
}

type UpdateSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateSongResponse) Reset() {
	*x = UpdateSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSongResponse) ProtoMessage() {}

func (x *UpdateSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSongResponse.ProtoReflect.Descriptor instead.
func (*UpdateSongResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{13}
}

var File_proto_player_proto protoreflect.FileDescriptor

var file_proto_player_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x50,
	0x6c, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x6c,
	0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x61,
	0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x50, 0x61,
	0x75, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x4e,
	0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x4e, 0x65,
	0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x72,
	0x65, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x72, 0x65,
	0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x0a, 0x0e, 0x41, 0x64, 0x64,
	0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x22, 0x11, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x13, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa6, 0x03, 0x0a, 0x0d, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x04,
	0x50, 0x6c, 0x61, 0x79, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x34, 0x0a, 0x05, 0x50, 0x61, 0x75, 0x73, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x04, 0x4e, 0x65, 0x78, 0x74, 0x12, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x04, 0x50, 0x72, 0x65, 0x76,
	0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x65, 0x76, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x65,
	0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x64, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x43, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_player_proto_rawDescOnce sync.Once
	file_proto_player_proto_rawDescData = file_proto_player_proto_rawDesc
)

func file_proto_player_proto_rawDescGZIP() []byte {
	file_proto_player_proto_rawDescOnce.Do(func() {
		file_proto_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_player_proto_rawDescData)
	})
	return file_proto_player_proto_rawDescData
}

var file_proto_player_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_proto_player_proto_goTypes = []interface{}{
	(*PlayRequest)(nil),        // 0: proto.PlayRequest
	(*PlayResponse)(nil),       // 1: proto.PlayResponse
	(*PauseRequest)(nil),       // 2: proto.PauseRequest
	(*PauseResponse)(nil),      // 3: proto.PauseResponse
	(*NextRequest)(nil),        // 4: proto.NextRequest
	(*NextResponse)(nil),       // 5: proto.NextResponse
	(*PrevRequest)(nil),        // 6: proto.PrevRequest
	(*PrevResponse)(nil),       // 7: proto.PrevResponse
	(*AddSongRequest)(nil),     // 8: proto.AddSongRequest
	(*AddSongResponse)(nil),    // 9: proto.AddSongResponse
	(*DeleteSongRequest)(nil),  // 10: proto.DeleteSongRequest
	(*DeleteSongResponse)(nil), // 11: proto.DeleteSongResponse
	(*UpdateSongRequest)(nil),  // 12: proto.UpdateSongRequest
	(*UpdateSongResponse)(nil), // 13: proto.UpdateSongResponse
}
var file_proto_player_proto_depIdxs = []int32{
	0,  // 0: proto.PlayerService.Play:input_type -> proto.PlayRequest
	2,  // 1: proto.PlayerService.Pause:input_type -> proto.PauseRequest
	4,  // 2: proto.PlayerService.Next:input_type -> proto.NextRequest
	6,  // 3: proto.PlayerService.Prev:input_type -> proto.PrevRequest
	8,  // 4: proto.PlayerService.AddSong:input_type -> proto.AddSongRequest
	10, // 5: proto.PlayerService.DeleteSong:input_type -> proto.DeleteSongRequest
	12, // 6: proto.PlayerService.UpdateSong:input_type -> proto.UpdateSongRequest
	1,  // 7: proto.PlayerService.Play:output_type -> proto.PlayResponse
	3,  // 8: proto.PlayerService.Pause:output_type -> proto.PauseResponse
	5,  // 9: proto.PlayerService.Next:output_type -> proto.NextResponse
	7,  // 10: proto.PlayerService.Prev:output_type -> proto.PrevResponse
	9,  // 11: proto.PlayerService.AddSong:output_type -> proto.AddSongResponse
	11, // 12: proto.PlayerService.DeleteSong:output_type -> proto.DeleteSongResponse
	13, // 13: proto.PlayerService.UpdateSong:output_type -> proto.UpdateSongResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_player_proto_init() }
func file_proto_player_proto_init() {
	if File_proto_player_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayRequest); i {
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
		file_proto_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayResponse); i {
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
		file_proto_player_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PauseRequest); i {
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
		file_proto_player_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PauseResponse); i {
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
		file_proto_player_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NextRequest); i {
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
		file_proto_player_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NextResponse); i {
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
		file_proto_player_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrevRequest); i {
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
		file_proto_player_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrevResponse); i {
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
		file_proto_player_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSongRequest); i {
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
		file_proto_player_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSongResponse); i {
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
		file_proto_player_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSongRequest); i {
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
		file_proto_player_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSongResponse); i {
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
		file_proto_player_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSongRequest); i {
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
		file_proto_player_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSongResponse); i {
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
			RawDescriptor: file_proto_player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_player_proto_goTypes,
		DependencyIndexes: file_proto_player_proto_depIdxs,
		MessageInfos:      file_proto_player_proto_msgTypes,
	}.Build()
	File_proto_player_proto = out.File
	file_proto_player_proto_rawDesc = nil
	file_proto_player_proto_goTypes = nil
	file_proto_player_proto_depIdxs = nil
}