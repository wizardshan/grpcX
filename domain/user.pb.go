// demo/proto/book/book.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: domain/user.proto

// 声明protobuf中的包名

package domain

import (
	user "github.com/wizardshan/grpcX/domain/user"
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

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string         `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	AuthorInfo *user.Nickname `protobuf:"bytes,3,opt,name=authorInfo,proto3" json:"authorInfo,omitempty"` // 需要带package前缀
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_domain_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_domain_user_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthorInfo() *user.Nickname {
	if x != nil {
		return x.AuthorInfo
	}
	return nil
}

var File_domain_user_proto protoreflect.FileDescriptor

var file_domain_user_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x1a, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x69, 0x7a, 0x61, 0x72, 0x64, 0x73, 0x68, 0x61, 0x6e, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x58, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_domain_user_proto_rawDescOnce sync.Once
	file_domain_user_proto_rawDescData = file_domain_user_proto_rawDesc
)

func file_domain_user_proto_rawDescGZIP() []byte {
	file_domain_user_proto_rawDescOnce.Do(func() {
		file_domain_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_user_proto_rawDescData)
	})
	return file_domain_user_proto_rawDescData
}

var file_domain_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_domain_user_proto_goTypes = []any{
	(*Book)(nil),          // 0: domain.Book
	(*user.Nickname)(nil), // 1: user.Nickname
}
var file_domain_user_proto_depIdxs = []int32{
	1, // 0: domain.Book.authorInfo:type_name -> user.Nickname
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_domain_user_proto_init() }
func file_domain_user_proto_init() {
	if File_domain_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_domain_user_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Book); i {
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
			RawDescriptor: file_domain_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_user_proto_goTypes,
		DependencyIndexes: file_domain_user_proto_depIdxs,
		MessageInfos:      file_domain_user_proto_msgTypes,
	}.Build()
	File_domain_user_proto = out.File
	file_domain_user_proto_rawDesc = nil
	file_domain_user_proto_goTypes = nil
	file_domain_user_proto_depIdxs = nil
}
