// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: daemon/v1/daemon.proto

package daemonv1

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

type SetSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetSecretRequest) Reset() {
	*x = SetSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_v1_daemon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSecretRequest) ProtoMessage() {}

func (x *SetSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_v1_daemon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSecretRequest.ProtoReflect.Descriptor instead.
func (*SetSecretRequest) Descriptor() ([]byte, []int) {
	return file_daemon_v1_daemon_proto_rawDescGZIP(), []int{0}
}

func (x *SetSecretRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SetSecretRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SetSecretResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetSecretResponse) Reset() {
	*x = SetSecretResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_v1_daemon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSecretResponse) ProtoMessage() {}

func (x *SetSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_v1_daemon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSecretResponse.ProtoReflect.Descriptor instead.
func (*SetSecretResponse) Descriptor() ([]byte, []int) {
	return file_daemon_v1_daemon_proto_rawDescGZIP(), []int{1}
}

type GetSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetSecretRequest) Reset() {
	*x = GetSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_v1_daemon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretRequest) ProtoMessage() {}

func (x *GetSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_v1_daemon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretRequest.ProtoReflect.Descriptor instead.
func (*GetSecretRequest) Descriptor() ([]byte, []int) {
	return file_daemon_v1_daemon_proto_rawDescGZIP(), []int{2}
}

func (x *GetSecretRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetSecretResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value  string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Exists bool   `protobuf:"varint,2,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *GetSecretResponse) Reset() {
	*x = GetSecretResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_v1_daemon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretResponse) ProtoMessage() {}

func (x *GetSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_v1_daemon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretResponse.ProtoReflect.Descriptor instead.
func (*GetSecretResponse) Descriptor() ([]byte, []int) {
	return file_daemon_v1_daemon_proto_rawDescGZIP(), []int{3}
}

func (x *GetSecretResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *GetSecretResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

type GetProcessIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetProcessIdRequest) Reset() {
	*x = GetProcessIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_v1_daemon_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProcessIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProcessIdRequest) ProtoMessage() {}

func (x *GetProcessIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_v1_daemon_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProcessIdRequest.ProtoReflect.Descriptor instead.
func (*GetProcessIdRequest) Descriptor() ([]byte, []int) {
	return file_daemon_v1_daemon_proto_rawDescGZIP(), []int{4}
}

type GetProcessIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid int32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *GetProcessIdResponse) Reset() {
	*x = GetProcessIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_v1_daemon_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProcessIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProcessIdResponse) ProtoMessage() {}

func (x *GetProcessIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_v1_daemon_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProcessIdResponse.ProtoReflect.Descriptor instead.
func (*GetProcessIdResponse) Descriptor() ([]byte, []int) {
	return file_daemon_v1_daemon_proto_rawDescGZIP(), []int{5}
}

func (x *GetProcessIdResponse) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

var File_daemon_v1_daemon_proto protoreflect.FileDescriptor

var file_daemon_v1_daemon_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x65, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x22, 0x3c, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x41,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69,
	0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x73, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x28, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70,
	0x69, 0x64, 0x32, 0xf6, 0x01, 0x0a, 0x0d, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x12, 0x1b, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x64, 0x61,
	0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1e, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x94, 0x01, 0x0a, 0x0d,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x44,
	0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x74, 0x65, 0x6e,
	0x74, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x64, 0x61, 0x65,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x44, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x09, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15,
	0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_daemon_v1_daemon_proto_rawDescOnce sync.Once
	file_daemon_v1_daemon_proto_rawDescData = file_daemon_v1_daemon_proto_rawDesc
)

func file_daemon_v1_daemon_proto_rawDescGZIP() []byte {
	file_daemon_v1_daemon_proto_rawDescOnce.Do(func() {
		file_daemon_v1_daemon_proto_rawDescData = protoimpl.X.CompressGZIP(file_daemon_v1_daemon_proto_rawDescData)
	})
	return file_daemon_v1_daemon_proto_rawDescData
}

var file_daemon_v1_daemon_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_daemon_v1_daemon_proto_goTypes = []interface{}{
	(*SetSecretRequest)(nil),     // 0: daemon.v1.SetSecretRequest
	(*SetSecretResponse)(nil),    // 1: daemon.v1.SetSecretResponse
	(*GetSecretRequest)(nil),     // 2: daemon.v1.GetSecretRequest
	(*GetSecretResponse)(nil),    // 3: daemon.v1.GetSecretResponse
	(*GetProcessIdRequest)(nil),  // 4: daemon.v1.GetProcessIdRequest
	(*GetProcessIdResponse)(nil), // 5: daemon.v1.GetProcessIdResponse
}
var file_daemon_v1_daemon_proto_depIdxs = []int32{
	0, // 0: daemon.v1.DaemonService.SetSecret:input_type -> daemon.v1.SetSecretRequest
	2, // 1: daemon.v1.DaemonService.GetSecret:input_type -> daemon.v1.GetSecretRequest
	4, // 2: daemon.v1.DaemonService.GetProcessId:input_type -> daemon.v1.GetProcessIdRequest
	1, // 3: daemon.v1.DaemonService.SetSecret:output_type -> daemon.v1.SetSecretResponse
	3, // 4: daemon.v1.DaemonService.GetSecret:output_type -> daemon.v1.GetSecretResponse
	5, // 5: daemon.v1.DaemonService.GetProcessId:output_type -> daemon.v1.GetProcessIdResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_daemon_v1_daemon_proto_init() }
func file_daemon_v1_daemon_proto_init() {
	if File_daemon_v1_daemon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_daemon_v1_daemon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetSecretRequest); i {
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
		file_daemon_v1_daemon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetSecretResponse); i {
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
		file_daemon_v1_daemon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretRequest); i {
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
		file_daemon_v1_daemon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretResponse); i {
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
		file_daemon_v1_daemon_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProcessIdRequest); i {
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
		file_daemon_v1_daemon_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProcessIdResponse); i {
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
			RawDescriptor: file_daemon_v1_daemon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_daemon_v1_daemon_proto_goTypes,
		DependencyIndexes: file_daemon_v1_daemon_proto_depIdxs,
		MessageInfos:      file_daemon_v1_daemon_proto_msgTypes,
	}.Build()
	File_daemon_v1_daemon_proto = out.File
	file_daemon_v1_daemon_proto_rawDesc = nil
	file_daemon_v1_daemon_proto_goTypes = nil
	file_daemon_v1_daemon_proto_depIdxs = nil
}