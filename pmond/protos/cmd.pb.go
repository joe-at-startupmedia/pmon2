// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.0
// source: pmond/protos/cmd.proto

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

type Cmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Arg1      string     `protobuf:"bytes,3,opt,name=arg1,proto3" json:"arg1,omitempty"`
	ExecFlags *ExecFlags `protobuf:"bytes,4,opt,name=exec_flags,json=execFlags,proto3" json:"exec_flags,omitempty"`
}

func (x *Cmd) Reset() {
	*x = Cmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pmond_protos_cmd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cmd) ProtoMessage() {}

func (x *Cmd) ProtoReflect() protoreflect.Message {
	mi := &file_pmond_protos_cmd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cmd.ProtoReflect.Descriptor instead.
func (*Cmd) Descriptor() ([]byte, []int) {
	return file_pmond_protos_cmd_proto_rawDescGZIP(), []int{0}
}

func (x *Cmd) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Cmd) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cmd) GetArg1() string {
	if x != nil {
		return x.Arg1
	}
	return ""
}

func (x *Cmd) GetExecFlags() *ExecFlags {
	if x != nil {
		return x.ExecFlags
	}
	return nil
}

var File_pmond_protos_cmd_proto protoreflect.FileDescriptor

var file_pmond_protos_cmd_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6d, 0x6f, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63,
	0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x6d, 0x64, 0x1a, 0x1d, 0x70,
	0x6d, 0x6f, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x65, 0x63,
	0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x03,
	0x43, 0x6d, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x31, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x31, 0x12, 0x34, 0x0a, 0x0a, 0x65,
	0x78, 0x65, 0x63, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x2e, 0x45, 0x78, 0x65,
	0x63, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x09, 0x65, 0x78, 0x65, 0x63, 0x46, 0x6c, 0x61, 0x67,
	0x73, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x6d, 0x6f, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pmond_protos_cmd_proto_rawDescOnce sync.Once
	file_pmond_protos_cmd_proto_rawDescData = file_pmond_protos_cmd_proto_rawDesc
)

func file_pmond_protos_cmd_proto_rawDescGZIP() []byte {
	file_pmond_protos_cmd_proto_rawDescOnce.Do(func() {
		file_pmond_protos_cmd_proto_rawDescData = protoimpl.X.CompressGZIP(file_pmond_protos_cmd_proto_rawDescData)
	})
	return file_pmond_protos_cmd_proto_rawDescData
}

var file_pmond_protos_cmd_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pmond_protos_cmd_proto_goTypes = []interface{}{
	(*Cmd)(nil),       // 0: cmd.Cmd
	(*ExecFlags)(nil), // 1: exec_flags.ExecFlags
}
var file_pmond_protos_cmd_proto_depIdxs = []int32{
	1, // 0: cmd.Cmd.exec_flags:type_name -> exec_flags.ExecFlags
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pmond_protos_cmd_proto_init() }
func file_pmond_protos_cmd_proto_init() {
	if File_pmond_protos_cmd_proto != nil {
		return
	}
	file_pmond_protos_exec_flags_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pmond_protos_cmd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cmd); i {
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
			RawDescriptor: file_pmond_protos_cmd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pmond_protos_cmd_proto_goTypes,
		DependencyIndexes: file_pmond_protos_cmd_proto_depIdxs,
		MessageInfos:      file_pmond_protos_cmd_proto_msgTypes,
	}.Build()
	File_pmond_protos_cmd_proto = out.File
	file_pmond_protos_cmd_proto_rawDesc = nil
	file_pmond_protos_cmd_proto_goTypes = nil
	file_pmond_protos_cmd_proto_depIdxs = nil
}
