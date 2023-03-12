// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        (unknown)
// source: api/greeting/v1/model.proto

package greetingpb

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

type GreetingMessage_EntityType int32

const (
	GreetingMessage_ENTITY_TYPE_HUMAN             GreetingMessage_EntityType = 0
	GreetingMessage_ENTITY_TYPE_EXTRA_TERRESTRIAL GreetingMessage_EntityType = 1
)

// Enum value maps for GreetingMessage_EntityType.
var (
	GreetingMessage_EntityType_name = map[int32]string{
		0: "ENTITY_TYPE_HUMAN",
		1: "ENTITY_TYPE_EXTRA_TERRESTRIAL",
	}
	GreetingMessage_EntityType_value = map[string]int32{
		"ENTITY_TYPE_HUMAN":             0,
		"ENTITY_TYPE_EXTRA_TERRESTRIAL": 1,
	}
)

func (x GreetingMessage_EntityType) Enum() *GreetingMessage_EntityType {
	p := new(GreetingMessage_EntityType)
	*p = x
	return p
}

func (x GreetingMessage_EntityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GreetingMessage_EntityType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_greeting_v1_model_proto_enumTypes[0].Descriptor()
}

func (GreetingMessage_EntityType) Type() protoreflect.EnumType {
	return &file_api_greeting_v1_model_proto_enumTypes[0]
}

func (x GreetingMessage_EntityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GreetingMessage_EntityType.Descriptor instead.
func (GreetingMessage_EntityType) EnumDescriptor() ([]byte, []int) {
	return file_api_greeting_v1_model_proto_rawDescGZIP(), []int{0, 0}
}

type GreetingMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Message    string                     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	EntityType GreetingMessage_EntityType `protobuf:"varint,3,opt,name=entity_type,json=entityType,proto3,enum=api.greeting.v1.GreetingMessage_EntityType" json:"entity_type,omitempty"`
}

func (x *GreetingMessage) Reset() {
	*x = GreetingMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_greeting_v1_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetingMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingMessage) ProtoMessage() {}

func (x *GreetingMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_greeting_v1_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingMessage.ProtoReflect.Descriptor instead.
func (*GreetingMessage) Descriptor() ([]byte, []int) {
	return file_api_greeting_v1_model_proto_rawDescGZIP(), []int{0}
}

func (x *GreetingMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GreetingMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GreetingMessage) GetEntityType() GreetingMessage_EntityType {
	if x != nil {
		return x.EntityType
	}
	return GreetingMessage_ENTITY_TYPE_HUMAN
}

type GreetingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GreetingResponse) Reset() {
	*x = GreetingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_greeting_v1_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingResponse) ProtoMessage() {}

func (x *GreetingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_greeting_v1_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingResponse.ProtoReflect.Descriptor instead.
func (*GreetingResponse) Descriptor() ([]byte, []int) {
	return file_api_greeting_v1_model_proto_rawDescGZIP(), []int{1}
}

func (x *GreetingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_greeting_v1_model_proto protoreflect.FileDescriptor

var file_api_greeting_v1_model_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61,
	0x70, 0x69, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x22, 0xd5,
	0x01, 0x0a, 0x0f, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x4c, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x65, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x46,
	0x0a, 0x0a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11,
	0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x55, 0x4d, 0x41,
	0x4e, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x45, 0x58, 0x54, 0x52, 0x41, 0x5f, 0x54, 0x45, 0x52, 0x52, 0x45, 0x53, 0x54,
	0x52, 0x49, 0x41, 0x4c, 0x10, 0x01, 0x22, 0x2c, 0x0a, 0x10, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x6e, 0x61, 0x75, 0x74, 0x65, 0x75, 0x72, 0x2f, 0x67, 0x6f, 0x5f,
	0x6d, 0x65, 0x65, 0x74, 0x75, 0x70, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_greeting_v1_model_proto_rawDescOnce sync.Once
	file_api_greeting_v1_model_proto_rawDescData = file_api_greeting_v1_model_proto_rawDesc
)

func file_api_greeting_v1_model_proto_rawDescGZIP() []byte {
	file_api_greeting_v1_model_proto_rawDescOnce.Do(func() {
		file_api_greeting_v1_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_greeting_v1_model_proto_rawDescData)
	})
	return file_api_greeting_v1_model_proto_rawDescData
}

var file_api_greeting_v1_model_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_greeting_v1_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_greeting_v1_model_proto_goTypes = []interface{}{
	(GreetingMessage_EntityType)(0), // 0: api.greeting.v1.GreetingMessage.EntityType
	(*GreetingMessage)(nil),         // 1: api.greeting.v1.GreetingMessage
	(*GreetingResponse)(nil),        // 2: api.greeting.v1.GreetingResponse
}
var file_api_greeting_v1_model_proto_depIdxs = []int32{
	0, // 0: api.greeting.v1.GreetingMessage.entity_type:type_name -> api.greeting.v1.GreetingMessage.EntityType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_greeting_v1_model_proto_init() }
func file_api_greeting_v1_model_proto_init() {
	if File_api_greeting_v1_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_greeting_v1_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetingMessage); i {
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
		file_api_greeting_v1_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetingResponse); i {
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
			RawDescriptor: file_api_greeting_v1_model_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_greeting_v1_model_proto_goTypes,
		DependencyIndexes: file_api_greeting_v1_model_proto_depIdxs,
		EnumInfos:         file_api_greeting_v1_model_proto_enumTypes,
		MessageInfos:      file_api_greeting_v1_model_proto_msgTypes,
	}.Build()
	File_api_greeting_v1_model_proto = out.File
	file_api_greeting_v1_model_proto_rawDesc = nil
	file_api_greeting_v1_model_proto_goTypes = nil
	file_api_greeting_v1_model_proto_depIdxs = nil
}
