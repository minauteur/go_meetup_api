// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        (unknown)
// source: api/record/v1/model.proto

package recpb

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

type GetRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRecordRequest) Reset() {
	*x = GetRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_record_v1_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordRequest) ProtoMessage() {}

func (x *GetRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_record_v1_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordRequest.ProtoReflect.Descriptor instead.
func (*GetRecordRequest) Descriptor() ([]byte, []int) {
	return file_api_record_v1_model_proto_rawDescGZIP(), []int{0}
}

func (x *GetRecordRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Record *GetRecordResponse_Record `protobuf:"bytes,4,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *GetRecordResponse) Reset() {
	*x = GetRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_record_v1_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordResponse) ProtoMessage() {}

func (x *GetRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_record_v1_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordResponse.ProtoReflect.Descriptor instead.
func (*GetRecordResponse) Descriptor() ([]byte, []int) {
	return file_api_record_v1_model_proto_rawDescGZIP(), []int{1}
}

func (x *GetRecordResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetRecordResponse) GetRecord() *GetRecordResponse_Record {
	if x != nil {
		return x.Record
	}
	return nil
}

type GetRecordResponse_Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Public  string `protobuf:"bytes,2,opt,name=public,proto3" json:"public,omitempty"`
	Private string `protobuf:"bytes,3,opt,name=private,proto3" json:"private,omitempty"`
}

func (x *GetRecordResponse_Record) Reset() {
	*x = GetRecordResponse_Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_record_v1_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordResponse_Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordResponse_Record) ProtoMessage() {}

func (x *GetRecordResponse_Record) ProtoReflect() protoreflect.Message {
	mi := &file_api_record_v1_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordResponse_Record.ProtoReflect.Descriptor instead.
func (*GetRecordResponse_Record) Descriptor() ([]byte, []int) {
	return file_api_record_v1_model_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetRecordResponse_Record) GetPublic() string {
	if x != nil {
		return x.Public
	}
	return ""
}

func (x *GetRecordResponse_Record) GetPrivate() string {
	if x != nil {
		return x.Private
	}
	return ""
}

var File_api_record_v1_model_proto protoreflect.FileDescriptor

var file_api_record_v1_model_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa0,
	0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x3f, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x3a, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x69, 0x6e, 0x61, 0x75, 0x74, 0x65, 0x75, 0x72, 0x2f, 0x67, 0x6f, 0x5f, 0x6d, 0x65, 0x65,
	0x74, 0x75, 0x70, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x63, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_record_v1_model_proto_rawDescOnce sync.Once
	file_api_record_v1_model_proto_rawDescData = file_api_record_v1_model_proto_rawDesc
)

func file_api_record_v1_model_proto_rawDescGZIP() []byte {
	file_api_record_v1_model_proto_rawDescOnce.Do(func() {
		file_api_record_v1_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_record_v1_model_proto_rawDescData)
	})
	return file_api_record_v1_model_proto_rawDescData
}

var file_api_record_v1_model_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_record_v1_model_proto_goTypes = []interface{}{
	(*GetRecordRequest)(nil),         // 0: api.record.v1.GetRecordRequest
	(*GetRecordResponse)(nil),        // 1: api.record.v1.GetRecordResponse
	(*GetRecordResponse_Record)(nil), // 2: api.record.v1.GetRecordResponse.Record
}
var file_api_record_v1_model_proto_depIdxs = []int32{
	2, // 0: api.record.v1.GetRecordResponse.record:type_name -> api.record.v1.GetRecordResponse.Record
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_record_v1_model_proto_init() }
func file_api_record_v1_model_proto_init() {
	if File_api_record_v1_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_record_v1_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecordRequest); i {
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
		file_api_record_v1_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecordResponse); i {
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
		file_api_record_v1_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecordResponse_Record); i {
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
			RawDescriptor: file_api_record_v1_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_record_v1_model_proto_goTypes,
		DependencyIndexes: file_api_record_v1_model_proto_depIdxs,
		MessageInfos:      file_api_record_v1_model_proto_msgTypes,
	}.Build()
	File_api_record_v1_model_proto = out.File
	file_api_record_v1_model_proto_rawDesc = nil
	file_api_record_v1_model_proto_goTypes = nil
	file_api_record_v1_model_proto_depIdxs = nil
}
