// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: default.proto

package cmdpb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DefaultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reqid int32 `protobuf:"varint,1,opt,name=reqid,proto3" json:"reqid,omitempty"`
}

func (x *DefaultRequest) Reset() {
	*x = DefaultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_default_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultRequest) ProtoMessage() {}

func (x *DefaultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_default_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefaultRequest.ProtoReflect.Descriptor instead.
func (*DefaultRequest) Descriptor() ([]byte, []int) {
	return file_default_proto_rawDescGZIP(), []int{0}
}

func (x *DefaultRequest) GetReqid() int32 {
	if x != nil {
		return x.Reqid
	}
	return 0
}

type DefaultRequestMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DefaultRequestMsg) Reset() {
	*x = DefaultRequestMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_default_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultRequestMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultRequestMsg) ProtoMessage() {}

func (x *DefaultRequestMsg) ProtoReflect() protoreflect.Message {
	mi := &file_default_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefaultRequestMsg.ProtoReflect.Descriptor instead.
func (*DefaultRequestMsg) Descriptor() ([]byte, []int) {
	return file_default_proto_rawDescGZIP(), []int{1}
}

func (x *DefaultRequestMsg) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DefaultResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DefaultResp) Reset() {
	*x = DefaultResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_default_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultResp) ProtoMessage() {}

func (x *DefaultResp) ProtoReflect() protoreflect.Message {
	mi := &file_default_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefaultResp.ProtoReflect.Descriptor instead.
func (*DefaultResp) Descriptor() ([]byte, []int) {
	return file_default_proto_rawDescGZIP(), []int{2}
}

func (x *DefaultResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DefaultRequestIDMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reqid   int32  `protobuf:"varint,1,opt,name=reqid,proto3" json:"reqid,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DefaultRequestIDMsg) Reset() {
	*x = DefaultRequestIDMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_default_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultRequestIDMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultRequestIDMsg) ProtoMessage() {}

func (x *DefaultRequestIDMsg) ProtoReflect() protoreflect.Message {
	mi := &file_default_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefaultRequestIDMsg.ProtoReflect.Descriptor instead.
func (*DefaultRequestIDMsg) Descriptor() ([]byte, []int) {
	return file_default_proto_rawDescGZIP(), []int{3}
}

func (x *DefaultRequestIDMsg) GetReqid() int32 {
	if x != nil {
		return x.Reqid
	}
	return 0
}

func (x *DefaultRequestIDMsg) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_default_proto protoreflect.FileDescriptor

var file_default_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x63, 0x6d, 0x64, 0x70, 0x62, 0x22, 0x26, 0x0a, 0x0e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x71, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x65, 0x71, 0x69, 0x64, 0x22, 0x2d,
	0x0a, 0x11, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x27, 0x0a,
	0x0b, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x45, 0x0a, 0x13, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x72, 0x65, 0x71, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x65,
	0x71, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x3b, 0x63, 0x6d, 0x64, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_default_proto_rawDescOnce sync.Once
	file_default_proto_rawDescData = file_default_proto_rawDesc
)

func file_default_proto_rawDescGZIP() []byte {
	file_default_proto_rawDescOnce.Do(func() {
		file_default_proto_rawDescData = protoimpl.X.CompressGZIP(file_default_proto_rawDescData)
	})
	return file_default_proto_rawDescData
}

var file_default_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_default_proto_goTypes = []interface{}{
	(*DefaultRequest)(nil),      // 0: cmdpb.DefaultRequest
	(*DefaultRequestMsg)(nil),   // 1: cmdpb.DefaultRequestMsg
	(*DefaultResp)(nil),         // 2: cmdpb.DefaultResp
	(*DefaultRequestIDMsg)(nil), // 3: cmdpb.DefaultRequestIDMsg
}
var file_default_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_default_proto_init() }
func file_default_proto_init() {
	if File_default_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_default_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefaultRequest); i {
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
		file_default_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefaultRequestMsg); i {
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
		file_default_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefaultResp); i {
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
		file_default_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefaultRequestIDMsg); i {
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
			RawDescriptor: file_default_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_default_proto_goTypes,
		DependencyIndexes: file_default_proto_depIdxs,
		MessageInfos:      file_default_proto_msgTypes,
	}.Build()
	File_default_proto = out.File
	file_default_proto_rawDesc = nil
	file_default_proto_goTypes = nil
	file_default_proto_depIdxs = nil
}
