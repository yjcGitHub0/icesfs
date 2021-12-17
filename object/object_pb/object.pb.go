// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: object.proto

package object_pb

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

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullPath string `protobuf:"bytes,1,opt,name=full_path,json=fullPath,proto3" json:"full_path,omitempty"`
	Set      string `protobuf:"bytes,2,opt,name=set_iam,proto3" json:"set_iam,omitempty"`
	Time     int64  `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	Mode     uint32 `protobuf:"varint,4,opt,name=mode,proto3" json:"mode,omitempty"`
	Mine     string `protobuf:"bytes,5,opt,name=mine,proto3" json:"mine,omitempty"`
	Md5      []byte `protobuf:"bytes,6,opt,name=md5,proto3" json:"md5,omitempty"`
	FileSize uint64 `protobuf:"varint,7,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	VolumeId uint64 `protobuf:"varint,8,opt,name=volume_id,json=volumeId,proto3" json:"volume_id,omitempty"`
	Fid      string `protobuf:"bytes,9,opt,name=fid,proto3" json:"fid,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_object_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_object_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_object_proto_rawDescGZIP(), []int{0}
}

func (x *Object) GetFullPath() string {
	if x != nil {
		return x.FullPath
	}
	return ""
}

func (x *Object) GetSet() string {
	if x != nil {
		return x.Set
	}
	return ""
}

func (x *Object) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Object) GetMode() uint32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *Object) GetMine() string {
	if x != nil {
		return x.Mine
	}
	return ""
}

func (x *Object) GetMd5() []byte {
	if x != nil {
		return x.Md5
	}
	return nil
}

func (x *Object) GetFileSize() uint64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *Object) GetVolumeId() uint64 {
	if x != nil {
		return x.VolumeId
	}
	return 0
}

func (x *Object) GetFid() string {
	if x != nil {
		return x.Fid
	}
	return ""
}

var File_object_proto protoreflect.FileDescriptor

var file_object_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x62, 0x22, 0xd1, 0x01, 0x0a, 0x06, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d,
	0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x69, 0x6e, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6d, 0x64,
	0x35, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x66,
	0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x69, 0x64, 0x42, 0x0d, 0x5a,
	0x0b, 0x2e, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_object_proto_rawDescOnce sync.Once
	file_object_proto_rawDescData = file_object_proto_rawDesc
)

func file_object_proto_rawDescGZIP() []byte {
	file_object_proto_rawDescOnce.Do(func() {
		file_object_proto_rawDescData = protoimpl.X.CompressGZIP(file_object_proto_rawDescData)
	})
	return file_object_proto_rawDescData
}

var file_object_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_object_proto_goTypes = []interface{}{
	(*Object)(nil), // 0: object_pb.Object
}
var file_object_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_object_proto_init() }
func file_object_proto_init() {
	if File_object_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_object_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
			RawDescriptor: file_object_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_object_proto_goTypes,
		DependencyIndexes: file_object_proto_depIdxs,
		MessageInfos:      file_object_proto_msgTypes,
	}.Build()
	File_object_proto = out.File
	file_object_proto_rawDesc = nil
	file_object_proto_goTypes = nil
	file_object_proto_depIdxs = nil
}
