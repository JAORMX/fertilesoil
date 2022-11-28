// Code generated by entproto. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: entpb/entpb.proto

package entpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetDirectoryRequest_View int32

const (
	GetDirectoryRequest_VIEW_UNSPECIFIED GetDirectoryRequest_View = 0
	GetDirectoryRequest_BASIC            GetDirectoryRequest_View = 1
	GetDirectoryRequest_WITH_EDGE_IDS    GetDirectoryRequest_View = 2
)

// Enum value maps for GetDirectoryRequest_View.
var (
	GetDirectoryRequest_View_name = map[int32]string{
		0: "VIEW_UNSPECIFIED",
		1: "BASIC",
		2: "WITH_EDGE_IDS",
	}
	GetDirectoryRequest_View_value = map[string]int32{
		"VIEW_UNSPECIFIED": 0,
		"BASIC":            1,
		"WITH_EDGE_IDS":    2,
	}
)

func (x GetDirectoryRequest_View) Enum() *GetDirectoryRequest_View {
	p := new(GetDirectoryRequest_View)
	*p = x
	return p
}

func (x GetDirectoryRequest_View) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetDirectoryRequest_View) Descriptor() protoreflect.EnumDescriptor {
	return file_entpb_entpb_proto_enumTypes[0].Descriptor()
}

func (GetDirectoryRequest_View) Type() protoreflect.EnumType {
	return &file_entpb_entpb_proto_enumTypes[0]
}

func (x GetDirectoryRequest_View) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetDirectoryRequest_View.Descriptor instead.
func (GetDirectoryRequest_View) EnumDescriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{2, 0}
}

type ListDirectoryRequest_View int32

const (
	ListDirectoryRequest_VIEW_UNSPECIFIED ListDirectoryRequest_View = 0
	ListDirectoryRequest_BASIC            ListDirectoryRequest_View = 1
	ListDirectoryRequest_WITH_EDGE_IDS    ListDirectoryRequest_View = 2
)

// Enum value maps for ListDirectoryRequest_View.
var (
	ListDirectoryRequest_View_name = map[int32]string{
		0: "VIEW_UNSPECIFIED",
		1: "BASIC",
		2: "WITH_EDGE_IDS",
	}
	ListDirectoryRequest_View_value = map[string]int32{
		"VIEW_UNSPECIFIED": 0,
		"BASIC":            1,
		"WITH_EDGE_IDS":    2,
	}
)

func (x ListDirectoryRequest_View) Enum() *ListDirectoryRequest_View {
	p := new(ListDirectoryRequest_View)
	*p = x
	return p
}

func (x ListDirectoryRequest_View) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListDirectoryRequest_View) Descriptor() protoreflect.EnumDescriptor {
	return file_entpb_entpb_proto_enumTypes[1].Descriptor()
}

func (ListDirectoryRequest_View) Type() protoreflect.EnumType {
	return &file_entpb_entpb_proto_enumTypes[1]
}

func (x ListDirectoryRequest_View) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListDirectoryRequest_View.Descriptor instead.
func (ListDirectoryRequest_View) EnumDescriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{5, 0}
}

type Directory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Metadata  *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	CreatedAt *timestamppb.Timestamp  `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp  `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *timestamppb.Timestamp  `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	IsRoot    bool                    `protobuf:"varint,7,opt,name=is_root,json=isRoot,proto3" json:"is_root,omitempty"`
	Children  []*Directory            `protobuf:"bytes,8,rep,name=children,proto3" json:"children,omitempty"`
	Parent    *Directory              `protobuf:"bytes,9,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (x *Directory) Reset() {
	*x = Directory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Directory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Directory) ProtoMessage() {}

func (x *Directory) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Directory.ProtoReflect.Descriptor instead.
func (*Directory) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{0}
}

func (x *Directory) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Directory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Directory) GetMetadata() *wrapperspb.StringValue {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Directory) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Directory) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Directory) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Directory) GetIsRoot() bool {
	if x != nil {
		return x.IsRoot
	}
	return false
}

func (x *Directory) GetChildren() []*Directory {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *Directory) GetParent() *Directory {
	if x != nil {
		return x.Parent
	}
	return nil
}

type CreateDirectoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Directory *Directory `protobuf:"bytes,1,opt,name=directory,proto3" json:"directory,omitempty"`
}

func (x *CreateDirectoryRequest) Reset() {
	*x = CreateDirectoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDirectoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDirectoryRequest) ProtoMessage() {}

func (x *CreateDirectoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDirectoryRequest.ProtoReflect.Descriptor instead.
func (*CreateDirectoryRequest) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDirectoryRequest) GetDirectory() *Directory {
	if x != nil {
		return x.Directory
	}
	return nil
}

type GetDirectoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64                    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	View GetDirectoryRequest_View `protobuf:"varint,2,opt,name=view,proto3,enum=entpb.GetDirectoryRequest_View" json:"view,omitempty"`
}

func (x *GetDirectoryRequest) Reset() {
	*x = GetDirectoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDirectoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDirectoryRequest) ProtoMessage() {}

func (x *GetDirectoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDirectoryRequest.ProtoReflect.Descriptor instead.
func (*GetDirectoryRequest) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{2}
}

func (x *GetDirectoryRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetDirectoryRequest) GetView() GetDirectoryRequest_View {
	if x != nil {
		return x.View
	}
	return GetDirectoryRequest_VIEW_UNSPECIFIED
}

type UpdateDirectoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Directory *Directory `protobuf:"bytes,1,opt,name=directory,proto3" json:"directory,omitempty"`
}

func (x *UpdateDirectoryRequest) Reset() {
	*x = UpdateDirectoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDirectoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDirectoryRequest) ProtoMessage() {}

func (x *UpdateDirectoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDirectoryRequest.ProtoReflect.Descriptor instead.
func (*UpdateDirectoryRequest) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateDirectoryRequest) GetDirectory() *Directory {
	if x != nil {
		return x.Directory
	}
	return nil
}

type DeleteDirectoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDirectoryRequest) Reset() {
	*x = DeleteDirectoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDirectoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDirectoryRequest) ProtoMessage() {}

func (x *DeleteDirectoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDirectoryRequest.ProtoReflect.Descriptor instead.
func (*DeleteDirectoryRequest) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteDirectoryRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListDirectoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize  int32                     `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string                    `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	View      ListDirectoryRequest_View `protobuf:"varint,3,opt,name=view,proto3,enum=entpb.ListDirectoryRequest_View" json:"view,omitempty"`
}

func (x *ListDirectoryRequest) Reset() {
	*x = ListDirectoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDirectoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDirectoryRequest) ProtoMessage() {}

func (x *ListDirectoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDirectoryRequest.ProtoReflect.Descriptor instead.
func (*ListDirectoryRequest) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{5}
}

func (x *ListDirectoryRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListDirectoryRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListDirectoryRequest) GetView() ListDirectoryRequest_View {
	if x != nil {
		return x.View
	}
	return ListDirectoryRequest_VIEW_UNSPECIFIED
}

type ListDirectoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DirectoryList []*Directory `protobuf:"bytes,1,rep,name=directory_list,json=directoryList,proto3" json:"directory_list,omitempty"`
	NextPageToken string       `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListDirectoryResponse) Reset() {
	*x = ListDirectoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDirectoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDirectoryResponse) ProtoMessage() {}

func (x *ListDirectoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDirectoryResponse.ProtoReflect.Descriptor instead.
func (*ListDirectoryResponse) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{6}
}

func (x *ListDirectoryResponse) GetDirectoryList() []*Directory {
	if x != nil {
		return x.DirectoryList
	}
	return nil
}

func (x *ListDirectoryResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_entpb_entpb_proto protoreflect.FileDescriptor

var file_entpb_entpb_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x03, 0x0a, 0x09, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x72, 0x6f, 0x6f, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x2c,
	0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x22, 0x48, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2e, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x22, 0x96, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x22, 0x3a, 0x0a,
	0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x14, 0x0a, 0x10, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x42,
	0x41, 0x53, 0x49, 0x43, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x45,
	0x44, 0x47, 0x45, 0x5f, 0x49, 0x44, 0x53, 0x10, 0x02, 0x22, 0x48, 0x0a, 0x16, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xc4, 0x01,
	0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x34, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x20, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x69,
	0x65, 0x77, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x22, 0x3a, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77,
	0x12, 0x14, 0x0a, 0x10, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x41, 0x53, 0x49, 0x43, 0x10,
	0x01, 0x12, 0x11, 0x0a, 0x0d, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x45, 0x44, 0x47, 0x45, 0x5f, 0x49,
	0x44, 0x53, 0x10, 0x02, 0x22, 0x78, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a,
	0x0e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x0d, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xc1,
	0x02, 0x0a, 0x10, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x33,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x39, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x3f,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x41, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4a, 0x41, 0x4f, 0x52, 0x4d, 0x58, 0x2f, 0x66, 0x65, 0x72, 0x74, 0x69, 0x6c, 0x65, 0x73,
	0x6f, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entpb_entpb_proto_rawDescOnce sync.Once
	file_entpb_entpb_proto_rawDescData = file_entpb_entpb_proto_rawDesc
)

func file_entpb_entpb_proto_rawDescGZIP() []byte {
	file_entpb_entpb_proto_rawDescOnce.Do(func() {
		file_entpb_entpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_entpb_entpb_proto_rawDescData)
	})
	return file_entpb_entpb_proto_rawDescData
}

var file_entpb_entpb_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_entpb_entpb_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_entpb_entpb_proto_goTypes = []interface{}{
	(GetDirectoryRequest_View)(0),  // 0: entpb.GetDirectoryRequest.View
	(ListDirectoryRequest_View)(0), // 1: entpb.ListDirectoryRequest.View
	(*Directory)(nil),              // 2: entpb.Directory
	(*CreateDirectoryRequest)(nil), // 3: entpb.CreateDirectoryRequest
	(*GetDirectoryRequest)(nil),    // 4: entpb.GetDirectoryRequest
	(*UpdateDirectoryRequest)(nil), // 5: entpb.UpdateDirectoryRequest
	(*DeleteDirectoryRequest)(nil), // 6: entpb.DeleteDirectoryRequest
	(*ListDirectoryRequest)(nil),   // 7: entpb.ListDirectoryRequest
	(*ListDirectoryResponse)(nil),  // 8: entpb.ListDirectoryResponse
	(*wrapperspb.StringValue)(nil), // 9: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),  // 10: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),          // 11: google.protobuf.Empty
}
var file_entpb_entpb_proto_depIdxs = []int32{
	9,  // 0: entpb.Directory.metadata:type_name -> google.protobuf.StringValue
	10, // 1: entpb.Directory.created_at:type_name -> google.protobuf.Timestamp
	10, // 2: entpb.Directory.updated_at:type_name -> google.protobuf.Timestamp
	10, // 3: entpb.Directory.deleted_at:type_name -> google.protobuf.Timestamp
	2,  // 4: entpb.Directory.children:type_name -> entpb.Directory
	2,  // 5: entpb.Directory.parent:type_name -> entpb.Directory
	2,  // 6: entpb.CreateDirectoryRequest.directory:type_name -> entpb.Directory
	0,  // 7: entpb.GetDirectoryRequest.view:type_name -> entpb.GetDirectoryRequest.View
	2,  // 8: entpb.UpdateDirectoryRequest.directory:type_name -> entpb.Directory
	1,  // 9: entpb.ListDirectoryRequest.view:type_name -> entpb.ListDirectoryRequest.View
	2,  // 10: entpb.ListDirectoryResponse.directory_list:type_name -> entpb.Directory
	3,  // 11: entpb.DirectoryService.Create:input_type -> entpb.CreateDirectoryRequest
	4,  // 12: entpb.DirectoryService.Get:input_type -> entpb.GetDirectoryRequest
	5,  // 13: entpb.DirectoryService.Update:input_type -> entpb.UpdateDirectoryRequest
	6,  // 14: entpb.DirectoryService.Delete:input_type -> entpb.DeleteDirectoryRequest
	7,  // 15: entpb.DirectoryService.List:input_type -> entpb.ListDirectoryRequest
	2,  // 16: entpb.DirectoryService.Create:output_type -> entpb.Directory
	2,  // 17: entpb.DirectoryService.Get:output_type -> entpb.Directory
	2,  // 18: entpb.DirectoryService.Update:output_type -> entpb.Directory
	11, // 19: entpb.DirectoryService.Delete:output_type -> google.protobuf.Empty
	8,  // 20: entpb.DirectoryService.List:output_type -> entpb.ListDirectoryResponse
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_entpb_entpb_proto_init() }
func file_entpb_entpb_proto_init() {
	if File_entpb_entpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entpb_entpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Directory); i {
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
		file_entpb_entpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDirectoryRequest); i {
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
		file_entpb_entpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDirectoryRequest); i {
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
		file_entpb_entpb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDirectoryRequest); i {
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
		file_entpb_entpb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDirectoryRequest); i {
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
		file_entpb_entpb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDirectoryRequest); i {
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
		file_entpb_entpb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDirectoryResponse); i {
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
			RawDescriptor: file_entpb_entpb_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_entpb_entpb_proto_goTypes,
		DependencyIndexes: file_entpb_entpb_proto_depIdxs,
		EnumInfos:         file_entpb_entpb_proto_enumTypes,
		MessageInfos:      file_entpb_entpb_proto_msgTypes,
	}.Build()
	File_entpb_entpb_proto = out.File
	file_entpb_entpb_proto_rawDesc = nil
	file_entpb_entpb_proto_goTypes = nil
	file_entpb_entpb_proto_depIdxs = nil
}
