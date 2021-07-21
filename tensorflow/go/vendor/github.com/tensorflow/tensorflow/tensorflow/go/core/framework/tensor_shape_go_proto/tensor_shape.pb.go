// Protocol buffer representing the shape of tensors.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: tensorflow/core/framework/tensor_shape.proto

package tensor_shape_go_proto

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

// Dimensions of a tensor.
type TensorShapeProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Dimensions of the tensor, such as {"input", 30}, {"output", 40}
	// for a 30 x 40 2D tensor.  If an entry has size -1, this
	// corresponds to a dimension of unknown size. The names are
	// optional.
	//
	// The order of entries in "dim" matters: It indicates the layout of the
	// values in the tensor in-memory representation.
	//
	// The first entry in "dim" is the outermost dimension used to layout the
	// values, the last entry is the innermost dimension.  This matches the
	// in-memory layout of RowMajor Eigen tensors.
	//
	// If "dim.size()" > 0, "unknown_rank" must be false.
	Dim []*TensorShapeProto_Dim `protobuf:"bytes,2,rep,name=dim,proto3" json:"dim,omitempty"`
	// If true, the number of dimensions in the shape is unknown.
	//
	// If true, "dim.size()" must be 0.
	UnknownRank bool `protobuf:"varint,3,opt,name=unknown_rank,json=unknownRank,proto3" json:"unknown_rank,omitempty"`
}

func (x *TensorShapeProto) Reset() {
	*x = TensorShapeProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_tensor_shape_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TensorShapeProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TensorShapeProto) ProtoMessage() {}

func (x *TensorShapeProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_tensor_shape_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TensorShapeProto.ProtoReflect.Descriptor instead.
func (*TensorShapeProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_tensor_shape_proto_rawDescGZIP(), []int{0}
}

func (x *TensorShapeProto) GetDim() []*TensorShapeProto_Dim {
	if x != nil {
		return x.Dim
	}
	return nil
}

func (x *TensorShapeProto) GetUnknownRank() bool {
	if x != nil {
		return x.UnknownRank
	}
	return false
}

// One dimension of the tensor.
type TensorShapeProto_Dim struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Size of the tensor in that dimension.
	// This value must be >= -1, but values of -1 are reserved for "unknown"
	// shapes (values of -1 mean "unknown" dimension).  Certain wrappers
	// that work with TensorShapeProto may fail at runtime when deserializing
	// a TensorShapeProto containing a dim value of -1.
	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// Optional name of the tensor dimension.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TensorShapeProto_Dim) Reset() {
	*x = TensorShapeProto_Dim{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_tensor_shape_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TensorShapeProto_Dim) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TensorShapeProto_Dim) ProtoMessage() {}

func (x *TensorShapeProto_Dim) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_tensor_shape_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TensorShapeProto_Dim.ProtoReflect.Descriptor instead.
func (*TensorShapeProto_Dim) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_tensor_shape_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TensorShapeProto_Dim) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *TensorShapeProto_Dim) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_tensorflow_core_framework_tensor_shape_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_tensor_shape_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0x98, 0x01, 0x0a, 0x10, 0x54,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x68, 0x61, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x32, 0x0a, 0x03, 0x64, 0x69, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x53, 0x68, 0x61, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x69, 0x6d, 0x52, 0x03,
	0x64, 0x69, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x72,
	0x61, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x75, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x52, 0x61, 0x6e, 0x6b, 0x1a, 0x2d, 0x0a, 0x03, 0x44, 0x69, 0x6d, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x87, 0x01, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x42, 0x11, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x68, 0x61, 0x70, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x73, 0x68,
	0x61, 0x70, 0x65, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_tensor_shape_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_tensor_shape_proto_rawDescData = file_tensorflow_core_framework_tensor_shape_proto_rawDesc
)

func file_tensorflow_core_framework_tensor_shape_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_tensor_shape_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_tensor_shape_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_tensor_shape_proto_rawDescData)
	})
	return file_tensorflow_core_framework_tensor_shape_proto_rawDescData
}

var file_tensorflow_core_framework_tensor_shape_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_core_framework_tensor_shape_proto_goTypes = []interface{}{
	(*TensorShapeProto)(nil),     // 0: tensorflow.TensorShapeProto
	(*TensorShapeProto_Dim)(nil), // 1: tensorflow.TensorShapeProto.Dim
}
var file_tensorflow_core_framework_tensor_shape_proto_depIdxs = []int32{
	1, // 0: tensorflow.TensorShapeProto.dim:type_name -> tensorflow.TensorShapeProto.Dim
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_tensor_shape_proto_init() }
func file_tensorflow_core_framework_tensor_shape_proto_init() {
	if File_tensorflow_core_framework_tensor_shape_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_tensor_shape_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TensorShapeProto); i {
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
		file_tensorflow_core_framework_tensor_shape_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TensorShapeProto_Dim); i {
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
			RawDescriptor: file_tensorflow_core_framework_tensor_shape_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_tensor_shape_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_tensor_shape_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_framework_tensor_shape_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_tensor_shape_proto = out.File
	file_tensorflow_core_framework_tensor_shape_proto_rawDesc = nil
	file_tensorflow_core_framework_tensor_shape_proto_goTypes = nil
	file_tensorflow_core_framework_tensor_shape_proto_depIdxs = nil
}
