// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: pkg/pkginventory/pkg/pkg.proto

package pkg

import (
	api "github.com/kform-dev/pkg-server/pkg/pkginventory/api"
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

type Get struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Get) Reset() {
	*x = Get{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pkginventory_pkg_pkg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Get) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Get) ProtoMessage() {}

func (x *Get) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pkginventory_pkg_pkg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Get.ProtoReflect.Descriptor instead.
func (*Get) Descriptor() ([]byte, []int) {
	return file_pkg_pkginventory_pkg_pkg_proto_rawDescGZIP(), []int{0}
}

type Get_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Package string `protobuf:"bytes,2,opt,name=package,proto3" json:"package,omitempty"`
}

func (x *Get_Request) Reset() {
	*x = Get_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pkginventory_pkg_pkg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Get_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Get_Request) ProtoMessage() {}

func (x *Get_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pkginventory_pkg_pkg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Get_Request.ProtoReflect.Descriptor instead.
func (*Get_Request) Descriptor() ([]byte, []int) {
	return file_pkg_pkginventory_pkg_pkg_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Get_Request) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *Get_Request) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

type Get_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageRevision *api.PackageRevision `protobuf:"bytes,1,opt,name=packageRevision,proto3" json:"packageRevision,omitempty"`
}

func (x *Get_Response) Reset() {
	*x = Get_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pkginventory_pkg_pkg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Get_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Get_Response) ProtoMessage() {}

func (x *Get_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pkginventory_pkg_pkg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Get_Response.ProtoReflect.Descriptor instead.
func (*Get_Response) Descriptor() ([]byte, []int) {
	return file_pkg_pkginventory_pkg_pkg_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Get_Response) GetPackageRevision() *api.PackageRevision {
	if x != nil {
		return x.PackageRevision
	}
	return nil
}

var File_pkg_pkginventory_pkg_pkg_proto protoreflect.FileDescriptor

var file_pkg_pkginventory_pkg_pkg_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6b, 0x67, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x70, 0x6b, 0x67, 0x1a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6b, 0x67, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x1a, 0x3d, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a, 0x4a, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0f, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x3e, 0x0a, 0x10, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x64, 0x65, 0x76,
	0x2f, 0x70, 0x6b, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x6b, 0x67, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pkginventory_pkg_pkg_proto_rawDescOnce sync.Once
	file_pkg_pkginventory_pkg_pkg_proto_rawDescData = file_pkg_pkginventory_pkg_pkg_proto_rawDesc
)

func file_pkg_pkginventory_pkg_pkg_proto_rawDescGZIP() []byte {
	file_pkg_pkginventory_pkg_pkg_proto_rawDescOnce.Do(func() {
		file_pkg_pkginventory_pkg_pkg_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pkginventory_pkg_pkg_proto_rawDescData)
	})
	return file_pkg_pkginventory_pkg_pkg_proto_rawDescData
}

var file_pkg_pkginventory_pkg_pkg_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_pkginventory_pkg_pkg_proto_goTypes = []interface{}{
	(*Get)(nil),                 // 0: pkg.Get
	(*Get_Request)(nil),         // 1: pkg.Get.Request
	(*Get_Response)(nil),        // 2: pkg.Get.Response
	(*api.PackageRevision)(nil), // 3: api.PackageRevision
}
var file_pkg_pkginventory_pkg_pkg_proto_depIdxs = []int32{
	3, // 0: pkg.Get.Response.packageRevision:type_name -> api.PackageRevision
	1, // 1: pkg.PackageResources.Get:input_type -> pkg.Get.Request
	2, // 2: pkg.PackageResources.Get:output_type -> pkg.Get.Response
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_pkginventory_pkg_pkg_proto_init() }
func file_pkg_pkginventory_pkg_pkg_proto_init() {
	if File_pkg_pkginventory_pkg_pkg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pkginventory_pkg_pkg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Get); i {
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
		file_pkg_pkginventory_pkg_pkg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Get_Request); i {
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
		file_pkg_pkginventory_pkg_pkg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Get_Response); i {
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
			RawDescriptor: file_pkg_pkginventory_pkg_pkg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pkginventory_pkg_pkg_proto_goTypes,
		DependencyIndexes: file_pkg_pkginventory_pkg_pkg_proto_depIdxs,
		MessageInfos:      file_pkg_pkginventory_pkg_pkg_proto_msgTypes,
	}.Build()
	File_pkg_pkginventory_pkg_pkg_proto = out.File
	file_pkg_pkginventory_pkg_pkg_proto_rawDesc = nil
	file_pkg_pkginventory_pkg_pkg_proto_goTypes = nil
	file_pkg_pkginventory_pkg_pkg_proto_depIdxs = nil
}