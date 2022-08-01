// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: iam.proto

package iam_pb

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

type S3ApiConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identities []*Identity `protobuf:"bytes,1,rep,name=identities,proto3" json:"identities,omitempty"`
}

func (x *S3ApiConfiguration) Reset() {
	*x = S3ApiConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3ApiConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3ApiConfiguration) ProtoMessage() {}

func (x *S3ApiConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_iam_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3ApiConfiguration.ProtoReflect.Descriptor instead.
func (*S3ApiConfiguration) Descriptor() ([]byte, []int) {
	return file_iam_proto_rawDescGZIP(), []int{0}
}

func (x *S3ApiConfiguration) GetIdentities() []*Identity {
	if x != nil {
		return x.Identities
	}
	return nil
}

type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Credentials []*Credential `protobuf:"bytes,2,rep,name=credentials,proto3" json:"credentials,omitempty"`
	Actions     []string      `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_iam_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_iam_proto_rawDescGZIP(), []int{1}
}

func (x *Identity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Identity) GetCredentials() []*Credential {
	if x != nil {
		return x.Credentials
	}
	return nil
}

func (x *Identity) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

type Credential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessKey string `protobuf:"bytes,1,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretKey string `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
}

func (x *Credential) Reset() {
	*x = Credential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credential) ProtoMessage() {}

func (x *Credential) ProtoReflect() protoreflect.Message {
	mi := &file_iam_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credential.ProtoReflect.Descriptor instead.
func (*Credential) Descriptor() ([]byte, []int) {
	return file_iam_proto_rawDescGZIP(), []int{2}
}

func (x *Credential) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *Credential) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

var File_iam_proto protoreflect.FileDescriptor

var file_iam_proto_rawDesc = []byte{
	0x0a, 0x09, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x69, 0x61, 0x6d,
	0x5f, 0x70, 0x62, 0x22, 0x46, 0x0a, 0x12, 0x53, 0x33, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x69, 0x61, 0x6d, 0x5f, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x6e, 0x0a, 0x08, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x69, 0x61, 0x6d, 0x5f, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4a, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x32, 0x21, 0x0a, 0x1f, 0x53, 0x65, 0x61, 0x77, 0x65,
	0x65, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x4b, 0x0a, 0x10, 0x73, 0x65,
	0x61, 0x77, 0x65, 0x65, 0x64, 0x66, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x08,
	0x49, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x61, 0x77, 0x65, 0x65, 0x64, 0x66, 0x73, 0x2f, 0x73,
	0x65, 0x61, 0x77, 0x65, 0x65, 0x64, 0x66, 0x73, 0x2f, 0x77, 0x65, 0x65, 0x64, 0x2f, 0x70, 0x62,
	0x2f, 0x69, 0x61, 0x6d, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iam_proto_rawDescOnce sync.Once
	file_iam_proto_rawDescData = file_iam_proto_rawDesc
)

func file_iam_proto_rawDescGZIP() []byte {
	file_iam_proto_rawDescOnce.Do(func() {
		file_iam_proto_rawDescData = protoimpl.X.CompressGZIP(file_iam_proto_rawDescData)
	})
	return file_iam_proto_rawDescData
}

var file_iam_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_iam_proto_goTypes = []interface{}{
	(*S3ApiConfiguration)(nil), // 0: iam_pb.S3ApiConfiguration
	(*Identity)(nil),           // 1: iam_pb.Identity
	(*Credential)(nil),         // 2: iam_pb.Credential
}
var file_iam_proto_depIdxs = []int32{
	1, // 0: iam_pb.S3ApiConfiguration.identities:type_name -> iam_pb.Identity
	2, // 1: iam_pb.Identity.credentials:type_name -> iam_pb.Credential
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_iam_proto_init() }
func file_iam_proto_init() {
	if File_iam_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iam_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3ApiConfiguration); i {
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
		file_iam_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity); i {
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
		file_iam_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credential); i {
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
			RawDescriptor: file_iam_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iam_proto_goTypes,
		DependencyIndexes: file_iam_proto_depIdxs,
		MessageInfos:      file_iam_proto_msgTypes,
	}.Build()
	File_iam_proto = out.File
	file_iam_proto_rawDesc = nil
	file_iam_proto_goTypes = nil
	file_iam_proto_depIdxs = nil
}
