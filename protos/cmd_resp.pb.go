// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.1
// source: protos/cmd_resp.proto

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

type CmdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ValueStr    string       `protobuf:"bytes,3,opt,name=value_str,json=valueStr,proto3" json:"value_str,omitempty"`
	Process     *Process     `protobuf:"bytes,4,opt,name=process,proto3" json:"process,omitempty"`
	ProcessList *ProcessList `protobuf:"bytes,5,opt,name=process_list,json=processList,proto3" json:"process_list,omitempty"`
	Group       *Group       `protobuf:"bytes,6,opt,name=group,proto3" json:"group,omitempty"`
	GroupList   *GroupList   `protobuf:"bytes,7,opt,name=group_list,json=groupList,proto3" json:"group_list,omitempty"`
	Error       string       `protobuf:"bytes,8,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CmdResp) Reset() {
	*x = CmdResp{}
	mi := &file_protos_cmd_resp_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CmdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdResp) ProtoMessage() {}

func (x *CmdResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_cmd_resp_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdResp.ProtoReflect.Descriptor instead.
func (*CmdResp) Descriptor() ([]byte, []int) {
	return file_protos_cmd_resp_proto_rawDescGZIP(), []int{0}
}

func (x *CmdResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CmdResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CmdResp) GetValueStr() string {
	if x != nil {
		return x.ValueStr
	}
	return ""
}

func (x *CmdResp) GetProcess() *Process {
	if x != nil {
		return x.Process
	}
	return nil
}

func (x *CmdResp) GetProcessList() *ProcessList {
	if x != nil {
		return x.ProcessList
	}
	return nil
}

func (x *CmdResp) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *CmdResp) GetGroupList() *GroupList {
	if x != nil {
		return x.GroupList
	}
	return nil
}

func (x *CmdResp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_protos_cmd_resp_proto protoreflect.FileDescriptor

var file_protos_cmd_resp_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6d, 0x64, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x6d, 0x64, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x02, 0x0a, 0x07,
	0x43, 0x6d, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x74, 0x72, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x37, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x0a,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x30, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_cmd_resp_proto_rawDescOnce sync.Once
	file_protos_cmd_resp_proto_rawDescData = file_protos_cmd_resp_proto_rawDesc
)

func file_protos_cmd_resp_proto_rawDescGZIP() []byte {
	file_protos_cmd_resp_proto_rawDescOnce.Do(func() {
		file_protos_cmd_resp_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_cmd_resp_proto_rawDescData)
	})
	return file_protos_cmd_resp_proto_rawDescData
}

var file_protos_cmd_resp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_cmd_resp_proto_goTypes = []any{
	(*CmdResp)(nil),     // 0: cmd_resp.CmdResp
	(*Process)(nil),     // 1: process.Process
	(*ProcessList)(nil), // 2: process.ProcessList
	(*Group)(nil),       // 3: pgroup.Group
	(*GroupList)(nil),   // 4: pgroup.GroupList
}
var file_protos_cmd_resp_proto_depIdxs = []int32{
	1, // 0: cmd_resp.CmdResp.process:type_name -> process.Process
	2, // 1: cmd_resp.CmdResp.process_list:type_name -> process.ProcessList
	3, // 2: cmd_resp.CmdResp.group:type_name -> pgroup.Group
	4, // 3: cmd_resp.CmdResp.group_list:type_name -> pgroup.GroupList
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protos_cmd_resp_proto_init() }
func file_protos_cmd_resp_proto_init() {
	if File_protos_cmd_resp_proto != nil {
		return
	}
	file_protos_process_proto_init()
	file_protos_group_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_cmd_resp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_cmd_resp_proto_goTypes,
		DependencyIndexes: file_protos_cmd_resp_proto_depIdxs,
		MessageInfos:      file_protos_cmd_resp_proto_msgTypes,
	}.Build()
	File_protos_cmd_resp_proto = out.File
	file_protos_cmd_resp_proto_rawDesc = nil
	file_protos_cmd_resp_proto_goTypes = nil
	file_protos_cmd_resp_proto_depIdxs = nil
}