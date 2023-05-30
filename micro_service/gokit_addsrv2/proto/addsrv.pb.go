// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: addsrv.proto

package proto

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

// Sum方法的请求参数
type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int64 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int64 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addsrv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_addsrv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_addsrv_proto_rawDescGZIP(), []int{0}
}

func (x *SumRequest) GetA() int64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *SumRequest) GetB() int64 {
	if x != nil {
		return x.B
	}
	return 0
}

// Sum方法的响应
type SumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V   int64  `protobuf:"varint,1,opt,name=v,proto3" json:"v,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addsrv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_addsrv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_addsrv_proto_rawDescGZIP(), []int{1}
}

func (x *SumResponse) GetV() int64 {
	if x != nil {
		return x.V
	}
	return 0
}

func (x *SumResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

// Concat方法的请求参数
type ConcatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B string `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *ConcatRequest) Reset() {
	*x = ConcatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addsrv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConcatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConcatRequest) ProtoMessage() {}

func (x *ConcatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_addsrv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConcatRequest.ProtoReflect.Descriptor instead.
func (*ConcatRequest) Descriptor() ([]byte, []int) {
	return file_addsrv_proto_rawDescGZIP(), []int{2}
}

func (x *ConcatRequest) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *ConcatRequest) GetB() string {
	if x != nil {
		return x.B
	}
	return ""
}

// Concat方法的响应
type ConcatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V   string `protobuf:"bytes,1,opt,name=v,proto3" json:"v,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *ConcatResponse) Reset() {
	*x = ConcatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addsrv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConcatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConcatResponse) ProtoMessage() {}

func (x *ConcatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_addsrv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConcatResponse.ProtoReflect.Descriptor instead.
func (*ConcatResponse) Descriptor() ([]byte, []int) {
	return file_addsrv_proto_rawDescGZIP(), []int{3}
}

func (x *ConcatResponse) GetV() string {
	if x != nil {
		return x.V
	}
	return ""
}

func (x *ConcatResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_addsrv_proto protoreflect.FileDescriptor

var file_addsrv_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x64, 0x64, 0x73, 0x72, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x28, 0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x61, 0x12, 0x0c,
	0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x62, 0x22, 0x2d, 0x0a, 0x0b,
	0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x76,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x76, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x2b, 0x0a, 0x0d, 0x43,
	0x6f, 0x6e, 0x63, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x62, 0x22, 0x30, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x63,
	0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x76, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x32, 0x62, 0x0a, 0x03, 0x41, 0x64,
	0x64, 0x12, 0x28, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x43,
	0x6f, 0x6e, 0x63, 0x61, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f,
	0x6e, 0x63, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x15,
	0x5a, 0x13, 0x67, 0x6f, 0x6b, 0x69, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x73, 0x72, 0x76, 0x32, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_addsrv_proto_rawDescOnce sync.Once
	file_addsrv_proto_rawDescData = file_addsrv_proto_rawDesc
)

func file_addsrv_proto_rawDescGZIP() []byte {
	file_addsrv_proto_rawDescOnce.Do(func() {
		file_addsrv_proto_rawDescData = protoimpl.X.CompressGZIP(file_addsrv_proto_rawDescData)
	})
	return file_addsrv_proto_rawDescData
}

var file_addsrv_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_addsrv_proto_goTypes = []interface{}{
	(*SumRequest)(nil),     // 0: pb.SumRequest
	(*SumResponse)(nil),    // 1: pb.SumResponse
	(*ConcatRequest)(nil),  // 2: pb.ConcatRequest
	(*ConcatResponse)(nil), // 3: pb.ConcatResponse
}
var file_addsrv_proto_depIdxs = []int32{
	0, // 0: pb.Add.Sum:input_type -> pb.SumRequest
	2, // 1: pb.Add.Concat:input_type -> pb.ConcatRequest
	1, // 2: pb.Add.Sum:output_type -> pb.SumResponse
	3, // 3: pb.Add.Concat:output_type -> pb.ConcatResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_addsrv_proto_init() }
func file_addsrv_proto_init() {
	if File_addsrv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_addsrv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
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
		file_addsrv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumResponse); i {
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
		file_addsrv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConcatRequest); i {
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
		file_addsrv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConcatResponse); i {
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
			RawDescriptor: file_addsrv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_addsrv_proto_goTypes,
		DependencyIndexes: file_addsrv_proto_depIdxs,
		MessageInfos:      file_addsrv_proto_msgTypes,
	}.Build()
	File_addsrv_proto = out.File
	file_addsrv_proto_rawDesc = nil
	file_addsrv_proto_goTypes = nil
	file_addsrv_proto_depIdxs = nil
}