// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: contracts/rest_contracts/result.proto

package rest_contracts

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

type BomberResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FormId    string      `protobuf:"bytes,1,opt,name=formId,proto3" json:"formId,omitempty"`
	Responses []*Response `protobuf:"bytes,2,rep,name=responses,proto3" json:"responses,omitempty"`
}

func (x *BomberResult) Reset() {
	*x = BomberResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contracts_rest_contracts_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BomberResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BomberResult) ProtoMessage() {}

func (x *BomberResult) ProtoReflect() protoreflect.Message {
	mi := &file_contracts_rest_contracts_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BomberResult.ProtoReflect.Descriptor instead.
func (*BomberResult) Descriptor() ([]byte, []int) {
	return file_contracts_rest_contracts_result_proto_rawDescGZIP(), []int{0}
}

func (x *BomberResult) GetFormId() string {
	if x != nil {
		return x.FormId
	}
	return ""
}

func (x *BomberResult) GetResponses() []*Response {
	if x != nil {
		return x.Responses
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Time int64 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"` // in ms
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contracts_rest_contracts_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_contracts_rest_contracts_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_contracts_rest_contracts_result_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

var File_contracts_rest_contracts_result_proto protoreflect.FileDescriptor

var file_contracts_rest_contracts_result_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x73, 0x22, 0x68, 0x0a, 0x0c, 0x42, 0x6f, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x09, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x32, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x42,
	0x60, 0x0a, 0x19, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x6f, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x74, 0x65,
	0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x5a, 0x43, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6f, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x74, 0x65, 0x61, 0x6d, 0x2f, 0x62, 0x6f, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contracts_rest_contracts_result_proto_rawDescOnce sync.Once
	file_contracts_rest_contracts_result_proto_rawDescData = file_contracts_rest_contracts_result_proto_rawDesc
)

func file_contracts_rest_contracts_result_proto_rawDescGZIP() []byte {
	file_contracts_rest_contracts_result_proto_rawDescOnce.Do(func() {
		file_contracts_rest_contracts_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_contracts_rest_contracts_result_proto_rawDescData)
	})
	return file_contracts_rest_contracts_result_proto_rawDescData
}

var file_contracts_rest_contracts_result_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_contracts_rest_contracts_result_proto_goTypes = []interface{}{
	(*BomberResult)(nil), // 0: contracts.rest_contracts.BomberResult
	(*Response)(nil),     // 1: contracts.rest_contracts.Response
}
var file_contracts_rest_contracts_result_proto_depIdxs = []int32{
	1, // 0: contracts.rest_contracts.BomberResult.responses:type_name -> contracts.rest_contracts.Response
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_contracts_rest_contracts_result_proto_init() }
func file_contracts_rest_contracts_result_proto_init() {
	if File_contracts_rest_contracts_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contracts_rest_contracts_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BomberResult); i {
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
		file_contracts_rest_contracts_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_contracts_rest_contracts_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contracts_rest_contracts_result_proto_goTypes,
		DependencyIndexes: file_contracts_rest_contracts_result_proto_depIdxs,
		MessageInfos:      file_contracts_rest_contracts_result_proto_msgTypes,
	}.Build()
	File_contracts_rest_contracts_result_proto = out.File
	file_contracts_rest_contracts_result_proto_rawDesc = nil
	file_contracts_rest_contracts_result_proto_goTypes = nil
	file_contracts_rest_contracts_result_proto_depIdxs = nil
}
