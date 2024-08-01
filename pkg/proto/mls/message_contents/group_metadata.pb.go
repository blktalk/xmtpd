// Group immutable metadata

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: mls/message_contents/group_metadata.proto

package message_contents

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

// Defines the type of conversation
type ConversationType int32

const (
	ConversationType_CONVERSATION_TYPE_UNSPECIFIED ConversationType = 0
	ConversationType_CONVERSATION_TYPE_GROUP       ConversationType = 1
	ConversationType_CONVERSATION_TYPE_DM          ConversationType = 2
	ConversationType_CONVERSATION_TYPE_SYNC        ConversationType = 3
)

// Enum value maps for ConversationType.
var (
	ConversationType_name = map[int32]string{
		0: "CONVERSATION_TYPE_UNSPECIFIED",
		1: "CONVERSATION_TYPE_GROUP",
		2: "CONVERSATION_TYPE_DM",
		3: "CONVERSATION_TYPE_SYNC",
	}
	ConversationType_value = map[string]int32{
		"CONVERSATION_TYPE_UNSPECIFIED": 0,
		"CONVERSATION_TYPE_GROUP":       1,
		"CONVERSATION_TYPE_DM":          2,
		"CONVERSATION_TYPE_SYNC":        3,
	}
)

func (x ConversationType) Enum() *ConversationType {
	p := new(ConversationType)
	*p = x
	return p
}

func (x ConversationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConversationType) Descriptor() protoreflect.EnumDescriptor {
	return file_mls_message_contents_group_metadata_proto_enumTypes[0].Descriptor()
}

func (ConversationType) Type() protoreflect.EnumType {
	return &file_mls_message_contents_group_metadata_proto_enumTypes[0]
}

func (x ConversationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConversationType.Descriptor instead.
func (ConversationType) EnumDescriptor() ([]byte, []int) {
	return file_mls_message_contents_group_metadata_proto_rawDescGZIP(), []int{0}
}

// Parent message for group metadata
type GroupMetadataV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConversationType ConversationType `protobuf:"varint,1,opt,name=conversation_type,json=conversationType,proto3,enum=xmtp.mls.message_contents.ConversationType" json:"conversation_type,omitempty"`
	// This will be removed soon
	CreatorAccountAddress string `protobuf:"bytes,2,opt,name=creator_account_address,json=creatorAccountAddress,proto3" json:"creator_account_address,omitempty"`
	CreatorInboxId        string `protobuf:"bytes,3,opt,name=creator_inbox_id,json=creatorInboxId,proto3" json:"creator_inbox_id,omitempty"`
	// Should only be present for CONVERSATION_TYPE_DM
	DmMembers *DmMembers `protobuf:"bytes,4,opt,name=dm_members,json=dmMembers,proto3,oneof" json:"dm_members,omitempty"`
}

func (x *GroupMetadataV1) Reset() {
	*x = GroupMetadataV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mls_message_contents_group_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupMetadataV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMetadataV1) ProtoMessage() {}

func (x *GroupMetadataV1) ProtoReflect() protoreflect.Message {
	mi := &file_mls_message_contents_group_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMetadataV1.ProtoReflect.Descriptor instead.
func (*GroupMetadataV1) Descriptor() ([]byte, []int) {
	return file_mls_message_contents_group_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *GroupMetadataV1) GetConversationType() ConversationType {
	if x != nil {
		return x.ConversationType
	}
	return ConversationType_CONVERSATION_TYPE_UNSPECIFIED
}

func (x *GroupMetadataV1) GetCreatorAccountAddress() string {
	if x != nil {
		return x.CreatorAccountAddress
	}
	return ""
}

func (x *GroupMetadataV1) GetCreatorInboxId() string {
	if x != nil {
		return x.CreatorInboxId
	}
	return ""
}

func (x *GroupMetadataV1) GetDmMembers() *DmMembers {
	if x != nil {
		return x.DmMembers
	}
	return nil
}

// Wrapper around an Inbox Id
type Inbox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InboxId string `protobuf:"bytes,1,opt,name=inbox_id,json=inboxId,proto3" json:"inbox_id,omitempty"`
}

func (x *Inbox) Reset() {
	*x = Inbox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mls_message_contents_group_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inbox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inbox) ProtoMessage() {}

func (x *Inbox) ProtoReflect() protoreflect.Message {
	mi := &file_mls_message_contents_group_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inbox.ProtoReflect.Descriptor instead.
func (*Inbox) Descriptor() ([]byte, []int) {
	return file_mls_message_contents_group_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *Inbox) GetInboxId() string {
	if x != nil {
		return x.InboxId
	}
	return ""
}

// Ordering does not matter here
type DmMembers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DmMemberOne *Inbox `protobuf:"bytes,1,opt,name=dm_member_one,json=dmMemberOne,proto3" json:"dm_member_one,omitempty"`
	DmMemberTwo *Inbox `protobuf:"bytes,2,opt,name=dm_member_two,json=dmMemberTwo,proto3" json:"dm_member_two,omitempty"`
}

func (x *DmMembers) Reset() {
	*x = DmMembers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mls_message_contents_group_metadata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DmMembers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DmMembers) ProtoMessage() {}

func (x *DmMembers) ProtoReflect() protoreflect.Message {
	mi := &file_mls_message_contents_group_metadata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DmMembers.ProtoReflect.Descriptor instead.
func (*DmMembers) Descriptor() ([]byte, []int) {
	return file_mls_message_contents_group_metadata_proto_rawDescGZIP(), []int{2}
}

func (x *DmMembers) GetDmMemberOne() *Inbox {
	if x != nil {
		return x.DmMemberOne
	}
	return nil
}

func (x *DmMembers) GetDmMemberTwo() *Inbox {
	if x != nil {
		return x.DmMemberTwo
	}
	return nil
}

var File_mls_message_contents_group_metadata_proto protoreflect.FileDescriptor

var file_mls_message_contents_group_metadata_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6d, 0x6c, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x78, 0x6d, 0x74,
	0x70, 0x2e, 0x6d, 0x6c, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xa6, 0x02, 0x0a, 0x0f, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x56, 0x31, 0x12, 0x58, 0x0a, 0x11, 0x63, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x6c, 0x73,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x28, 0x0a, 0x10,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49,
	0x6e, 0x62, 0x6f, 0x78, 0x49, 0x64, 0x12, 0x48, 0x0a, 0x0a, 0x64, 0x6d, 0x5f, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x78, 0x6d, 0x74,
	0x70, 0x2e, 0x6d, 0x6c, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x48, 0x00, 0x52, 0x09, 0x64, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x88, 0x01, 0x01,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x64, 0x6d, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22,
	0x22, 0x0a, 0x05, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6e, 0x62, 0x6f,
	0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x62, 0x6f,
	0x78, 0x49, 0x64, 0x22, 0x97, 0x01, 0x0a, 0x09, 0x44, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x12, 0x44, 0x0a, 0x0d, 0x64, 0x6d, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f,
	0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e,
	0x6d, 0x6c, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x52, 0x0b, 0x64, 0x6d, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x4f, 0x6e, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x64, 0x6d, 0x5f, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x74, 0x77, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x6c, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x49, 0x6e, 0x62, 0x6f, 0x78,
	0x52, 0x0b, 0x64, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x77, 0x6f, 0x2a, 0x88, 0x01,
	0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50,
	0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x4d, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16,
	0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x10, 0x03, 0x42, 0xeb, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d,
	0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x6c, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x12, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x6d, 0x74,
	0x70, 0x2f, 0x78, 0x6d, 0x74, 0x70, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6d, 0x6c, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0xa2, 0x02, 0x03, 0x58, 0x4d, 0x4d, 0xaa, 0x02, 0x18, 0x58,
	0x6d, 0x74, 0x70, 0x2e, 0x4d, 0x6c, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0xca, 0x02, 0x18, 0x58, 0x6d, 0x74, 0x70, 0x5c, 0x4d,
	0x6c, 0x73, 0x5c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0xe2, 0x02, 0x24, 0x58, 0x6d, 0x74, 0x70, 0x5c, 0x4d, 0x6c, 0x73, 0x5c, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x58, 0x6d, 0x74, 0x70,
	0x3a, 0x3a, 0x4d, 0x6c, 0x73, 0x3a, 0x3a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mls_message_contents_group_metadata_proto_rawDescOnce sync.Once
	file_mls_message_contents_group_metadata_proto_rawDescData = file_mls_message_contents_group_metadata_proto_rawDesc
)

func file_mls_message_contents_group_metadata_proto_rawDescGZIP() []byte {
	file_mls_message_contents_group_metadata_proto_rawDescOnce.Do(func() {
		file_mls_message_contents_group_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_mls_message_contents_group_metadata_proto_rawDescData)
	})
	return file_mls_message_contents_group_metadata_proto_rawDescData
}

var file_mls_message_contents_group_metadata_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mls_message_contents_group_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mls_message_contents_group_metadata_proto_goTypes = []any{
	(ConversationType)(0),   // 0: xmtp.mls.message_contents.ConversationType
	(*GroupMetadataV1)(nil), // 1: xmtp.mls.message_contents.GroupMetadataV1
	(*Inbox)(nil),           // 2: xmtp.mls.message_contents.Inbox
	(*DmMembers)(nil),       // 3: xmtp.mls.message_contents.DmMembers
}
var file_mls_message_contents_group_metadata_proto_depIdxs = []int32{
	0, // 0: xmtp.mls.message_contents.GroupMetadataV1.conversation_type:type_name -> xmtp.mls.message_contents.ConversationType
	3, // 1: xmtp.mls.message_contents.GroupMetadataV1.dm_members:type_name -> xmtp.mls.message_contents.DmMembers
	2, // 2: xmtp.mls.message_contents.DmMembers.dm_member_one:type_name -> xmtp.mls.message_contents.Inbox
	2, // 3: xmtp.mls.message_contents.DmMembers.dm_member_two:type_name -> xmtp.mls.message_contents.Inbox
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_mls_message_contents_group_metadata_proto_init() }
func file_mls_message_contents_group_metadata_proto_init() {
	if File_mls_message_contents_group_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mls_message_contents_group_metadata_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GroupMetadataV1); i {
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
		file_mls_message_contents_group_metadata_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Inbox); i {
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
		file_mls_message_contents_group_metadata_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DmMembers); i {
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
	file_mls_message_contents_group_metadata_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mls_message_contents_group_metadata_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mls_message_contents_group_metadata_proto_goTypes,
		DependencyIndexes: file_mls_message_contents_group_metadata_proto_depIdxs,
		EnumInfos:         file_mls_message_contents_group_metadata_proto_enumTypes,
		MessageInfos:      file_mls_message_contents_group_metadata_proto_msgTypes,
	}.Build()
	File_mls_message_contents_group_metadata_proto = out.File
	file_mls_message_contents_group_metadata_proto_rawDesc = nil
	file_mls_message_contents_group_metadata_proto_goTypes = nil
	file_mls_message_contents_group_metadata_proto_depIdxs = nil
}
