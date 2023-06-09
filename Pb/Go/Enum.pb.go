//@proto_enum

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: Enum.proto

package pb

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

type TestEnum1 int32

const (
	TestEnum1_TEST_0 TestEnum1 = 0
	TestEnum1_TEST_1 TestEnum1 = 1
)

// Enum value maps for TestEnum1.
var (
	TestEnum1_name = map[int32]string{
		0: "TEST_0",
		1: "TEST_1",
	}
	TestEnum1_value = map[string]int32{
		"TEST_0": 0,
		"TEST_1": 1,
	}
)

func (x TestEnum1) Enum() *TestEnum1 {
	p := new(TestEnum1)
	*p = x
	return p
}

func (x TestEnum1) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestEnum1) Descriptor() protoreflect.EnumDescriptor {
	return file_Enum_proto_enumTypes[0].Descriptor()
}

func (TestEnum1) Type() protoreflect.EnumType {
	return &file_Enum_proto_enumTypes[0]
}

func (x TestEnum1) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestEnum1.Descriptor instead.
func (TestEnum1) EnumDescriptor() ([]byte, []int) {
	return file_Enum_proto_rawDescGZIP(), []int{0}
}

type TestEnum2 int32

const (
	TestEnum2_TEST_3 TestEnum2 = 0
	TestEnum2_TEST_4 TestEnum2 = 4
)

// Enum value maps for TestEnum2.
var (
	TestEnum2_name = map[int32]string{
		0: "TEST_3",
		4: "TEST_4",
	}
	TestEnum2_value = map[string]int32{
		"TEST_3": 0,
		"TEST_4": 4,
	}
)

func (x TestEnum2) Enum() *TestEnum2 {
	p := new(TestEnum2)
	*p = x
	return p
}

func (x TestEnum2) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestEnum2) Descriptor() protoreflect.EnumDescriptor {
	return file_Enum_proto_enumTypes[1].Descriptor()
}

func (TestEnum2) Type() protoreflect.EnumType {
	return &file_Enum_proto_enumTypes[1]
}

func (x TestEnum2) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestEnum2.Descriptor instead.
func (TestEnum2) EnumDescriptor() ([]byte, []int) {
	return file_Enum_proto_rawDescGZIP(), []int{1}
}

var File_Enum_proto protoreflect.FileDescriptor

var file_Enum_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x2a, 0x23, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x0a, 0x0a,
	0x06, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x30, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x45, 0x53,
	0x54, 0x5f, 0x31, 0x10, 0x01, 0x2a, 0x23, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x75,
	0x6d, 0x32, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x33, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x34, 0x10, 0x04, 0x42, 0x0c, 0x5a, 0x05, 0x2e, 0x2f,
	0x3b, 0x70, 0x62, 0xaa, 0x02, 0x02, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Enum_proto_rawDescOnce sync.Once
	file_Enum_proto_rawDescData = file_Enum_proto_rawDesc
)

func file_Enum_proto_rawDescGZIP() []byte {
	file_Enum_proto_rawDescOnce.Do(func() {
		file_Enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_Enum_proto_rawDescData)
	})
	return file_Enum_proto_rawDescData
}

var file_Enum_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_Enum_proto_goTypes = []interface{}{
	(TestEnum1)(0), // 0: pb.TestEnum1
	(TestEnum2)(0), // 1: pb.TestEnum2
}
var file_Enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Enum_proto_init() }
func file_Enum_proto_init() {
	if File_Enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Enum_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Enum_proto_goTypes,
		DependencyIndexes: file_Enum_proto_depIdxs,
		EnumInfos:         file_Enum_proto_enumTypes,
	}.Build()
	File_Enum_proto = out.File
	file_Enum_proto_rawDesc = nil
	file_Enum_proto_goTypes = nil
	file_Enum_proto_depIdxs = nil
}
