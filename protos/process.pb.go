// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.1
// source: protos/process.proto

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

type Process struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt    string   `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    string   `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Pid          uint32   `protobuf:"varint,4,opt,name=pid,proto3" json:"pid,omitempty"`
	Log          string   `protobuf:"bytes,5,opt,name=log,proto3" json:"log,omitempty"`
	Name         string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	ProcessFile  string   `protobuf:"bytes,7,opt,name=process_file,json=processFile,proto3" json:"process_file,omitempty"`
	Args         string   `protobuf:"bytes,8,opt,name=args,proto3" json:"args,omitempty"`
	EnvVars      string   `protobuf:"bytes,9,opt,name=env_vars,json=envVars,proto3" json:"env_vars,omitempty"`
	Status       string   `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	AutoRestart  bool     `protobuf:"varint,11,opt,name=auto_restart,json=autoRestart,proto3" json:"auto_restart,omitempty"`
	Uid          uint32   `protobuf:"varint,12,opt,name=uid,proto3" json:"uid,omitempty"`
	Username     string   `protobuf:"bytes,13,opt,name=username,proto3" json:"username,omitempty"`
	Gid          uint32   `protobuf:"varint,14,opt,name=gid,proto3" json:"gid,omitempty"`
	RestartCount uint32   `protobuf:"varint,15,opt,name=restart_count,json=restartCount,proto3" json:"restart_count,omitempty"`
	MemoryUsage  string   `protobuf:"bytes,16,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`
	CpuUsage     string   `protobuf:"bytes,17,opt,name=cpu_usage,json=cpuUsage,proto3" json:"cpu_usage,omitempty"`
	Dependencies string   `protobuf:"bytes,18,opt,name=dependencies,proto3" json:"dependencies,omitempty"`
	Groups       []*Group `protobuf:"bytes,19,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *Process) Reset() {
	*x = Process{}
	mi := &file_protos_process_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Process) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Process) ProtoMessage() {}

func (x *Process) ProtoReflect() protoreflect.Message {
	mi := &file_protos_process_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Process.ProtoReflect.Descriptor instead.
func (*Process) Descriptor() ([]byte, []int) {
	return file_protos_process_proto_rawDescGZIP(), []int{0}
}

func (x *Process) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Process) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Process) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Process) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Process) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

func (x *Process) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Process) GetProcessFile() string {
	if x != nil {
		return x.ProcessFile
	}
	return ""
}

func (x *Process) GetArgs() string {
	if x != nil {
		return x.Args
	}
	return ""
}

func (x *Process) GetEnvVars() string {
	if x != nil {
		return x.EnvVars
	}
	return ""
}

func (x *Process) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Process) GetAutoRestart() bool {
	if x != nil {
		return x.AutoRestart
	}
	return false
}

func (x *Process) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Process) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Process) GetGid() uint32 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *Process) GetRestartCount() uint32 {
	if x != nil {
		return x.RestartCount
	}
	return 0
}

func (x *Process) GetMemoryUsage() string {
	if x != nil {
		return x.MemoryUsage
	}
	return ""
}

func (x *Process) GetCpuUsage() string {
	if x != nil {
		return x.CpuUsage
	}
	return ""
}

func (x *Process) GetDependencies() string {
	if x != nil {
		return x.Dependencies
	}
	return ""
}

func (x *Process) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

type ProcessList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Processes []*Process `protobuf:"bytes,1,rep,name=processes,proto3" json:"processes,omitempty"`
}

func (x *ProcessList) Reset() {
	*x = ProcessList{}
	mi := &file_protos_process_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessList) ProtoMessage() {}

func (x *ProcessList) ProtoReflect() protoreflect.Message {
	mi := &file_protos_process_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessList.ProtoReflect.Descriptor instead.
func (*ProcessList) Descriptor() ([]byte, []int) {
	return file_protos_process_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessList) GetProcesses() []*Process {
	if x != nil {
		return x.Processes
	}
	return nil
}

var File_protos_process_proto protoreflect.FileDescriptor

var file_protos_process_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x1a,
	0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x04, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x65, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x67,
	0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x70,
	0x75, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e,
	0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64,
	0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x22, 0x3d, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x2e, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_process_proto_rawDescOnce sync.Once
	file_protos_process_proto_rawDescData = file_protos_process_proto_rawDesc
)

func file_protos_process_proto_rawDescGZIP() []byte {
	file_protos_process_proto_rawDescOnce.Do(func() {
		file_protos_process_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_process_proto_rawDescData)
	})
	return file_protos_process_proto_rawDescData
}

var file_protos_process_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_process_proto_goTypes = []any{
	(*Process)(nil),     // 0: process.Process
	(*ProcessList)(nil), // 1: process.ProcessList
	(*Group)(nil),       // 2: pgroup.Group
}
var file_protos_process_proto_depIdxs = []int32{
	2, // 0: process.Process.groups:type_name -> pgroup.Group
	0, // 1: process.ProcessList.processes:type_name -> process.Process
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_process_proto_init() }
func file_protos_process_proto_init() {
	if File_protos_process_proto != nil {
		return
	}
	file_protos_group_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_process_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_process_proto_goTypes,
		DependencyIndexes: file_protos_process_proto_depIdxs,
		MessageInfos:      file_protos_process_proto_msgTypes,
	}.Build()
	File_protos_process_proto = out.File
	file_protos_process_proto_rawDesc = nil
	file_protos_process_proto_goTypes = nil
	file_protos_process_proto_depIdxs = nil
}