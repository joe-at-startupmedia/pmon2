// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.1
// source: protos/group.proto

package protos

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

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	mi := &file_protos_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_protos_group_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_protos_group_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GroupList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *GroupList) Reset() {
	*x = GroupList{}
	mi := &file_protos_group_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupList) ProtoMessage() {}

func (x *GroupList) ProtoReflect() protoreflect.Message {
	mi := &file_protos_group_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupList.ProtoReflect.Descriptor instead.
func (*GroupList) Descriptor() ([]byte, []int) {
	return file_protos_group_proto_rawDescGZIP(), []int{1}
}

func (x *GroupList) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

var File_protos_group_proto protoreflect.FileDescriptor

var file_protos_group_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x2b, 0x0a, 0x05,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x09, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protos_group_proto_rawDescOnce sync.Once
	file_protos_group_proto_rawDescData = file_protos_group_proto_rawDesc
)

func file_protos_group_proto_rawDescGZIP() []byte {
	file_protos_group_proto_rawDescOnce.Do(func() {
		file_protos_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_group_proto_rawDescData)
	})
	return file_protos_group_proto_rawDescData
}

var file_protos_group_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_group_proto_goTypes = []any{
	(*Group)(nil),     // 0: pgroup.Group
	(*GroupList)(nil), // 1: pgroup.GroupList
}
var file_protos_group_proto_depIdxs = []int32{
	0, // 0: pgroup.GroupList.groups:type_name -> pgroup.Group
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_group_proto_init() }
func file_protos_group_proto_init() {
	if File_protos_group_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_group_proto_goTypes,
		DependencyIndexes: file_protos_group_proto_depIdxs,
		MessageInfos:      file_protos_group_proto_msgTypes,
	}.Build()
	File_protos_group_proto = out.File
	file_protos_group_proto_rawDesc = nil
	file_protos_group_proto_goTypes = nil
	file_protos_group_proto_depIdxs = nil
}
