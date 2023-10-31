// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: remote.proto

package remote_pb

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

// ///////////////////////
// Remote Storage related
// ///////////////////////
type RemoteConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type                            string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name                            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	S3AccessKey                     string `protobuf:"bytes,4,opt,name=s3_access_key,json=s3AccessKey,proto3" json:"s3_access_key,omitempty"`
	S3SecretKey                     string `protobuf:"bytes,5,opt,name=s3_secret_key,json=s3SecretKey,proto3" json:"s3_secret_key,omitempty"`
	S3Region                        string `protobuf:"bytes,6,opt,name=s3_region,json=s3Region,proto3" json:"s3_region,omitempty"`
	S3Endpoint                      string `protobuf:"bytes,7,opt,name=s3_endpoint,json=s3Endpoint,proto3" json:"s3_endpoint,omitempty"`
	S3StorageClass                  string `protobuf:"bytes,8,opt,name=s3_storage_class,json=s3StorageClass,proto3" json:"s3_storage_class,omitempty"`
	S3ForcePathStyle                bool   `protobuf:"varint,9,opt,name=s3_force_path_style,json=s3ForcePathStyle,proto3" json:"s3_force_path_style,omitempty"`
	S3SupportTagging                bool   `protobuf:"varint,13,opt,name=s3_support_tagging,json=s3SupportTagging,proto3" json:"s3_support_tagging,omitempty"`
	S3V4Signature                   bool   `protobuf:"varint,11,opt,name=s3_v4_signature,json=s3V4Signature,proto3" json:"s3_v4_signature,omitempty"`
	GcsGoogleApplicationCredentials string `protobuf:"bytes,10,opt,name=gcs_google_application_credentials,json=gcsGoogleApplicationCredentials,proto3" json:"gcs_google_application_credentials,omitempty"`
	GcsProjectId                    string `protobuf:"bytes,12,opt,name=gcs_project_id,json=gcsProjectId,proto3" json:"gcs_project_id,omitempty"`
	AzureAccountName                string `protobuf:"bytes,15,opt,name=azure_account_name,json=azureAccountName,proto3" json:"azure_account_name,omitempty"`
	AzureAccountKey                 string `protobuf:"bytes,16,opt,name=azure_account_key,json=azureAccountKey,proto3" json:"azure_account_key,omitempty"`
	BackblazeKeyId                  string `protobuf:"bytes,20,opt,name=backblaze_key_id,json=backblazeKeyId,proto3" json:"backblaze_key_id,omitempty"`
	BackblazeApplicationKey         string `protobuf:"bytes,21,opt,name=backblaze_application_key,json=backblazeApplicationKey,proto3" json:"backblaze_application_key,omitempty"`
	BackblazeEndpoint               string `protobuf:"bytes,22,opt,name=backblaze_endpoint,json=backblazeEndpoint,proto3" json:"backblaze_endpoint,omitempty"`
	BackblazeRegion                 string `protobuf:"bytes,23,opt,name=backblaze_region,json=backblazeRegion,proto3" json:"backblaze_region,omitempty"`
	AliyunAccessKey                 string `protobuf:"bytes,25,opt,name=aliyun_access_key,json=aliyunAccessKey,proto3" json:"aliyun_access_key,omitempty"`
	AliyunSecretKey                 string `protobuf:"bytes,26,opt,name=aliyun_secret_key,json=aliyunSecretKey,proto3" json:"aliyun_secret_key,omitempty"`
	AliyunEndpoint                  string `protobuf:"bytes,27,opt,name=aliyun_endpoint,json=aliyunEndpoint,proto3" json:"aliyun_endpoint,omitempty"`
	AliyunRegion                    string `protobuf:"bytes,28,opt,name=aliyun_region,json=aliyunRegion,proto3" json:"aliyun_region,omitempty"`
	TencentSecretId                 string `protobuf:"bytes,30,opt,name=tencent_secret_id,json=tencentSecretId,proto3" json:"tencent_secret_id,omitempty"`
	TencentSecretKey                string `protobuf:"bytes,31,opt,name=tencent_secret_key,json=tencentSecretKey,proto3" json:"tencent_secret_key,omitempty"`
	TencentEndpoint                 string `protobuf:"bytes,32,opt,name=tencent_endpoint,json=tencentEndpoint,proto3" json:"tencent_endpoint,omitempty"`
	BaiduAccessKey                  string `protobuf:"bytes,35,opt,name=baidu_access_key,json=baiduAccessKey,proto3" json:"baidu_access_key,omitempty"`
	BaiduSecretKey                  string `protobuf:"bytes,36,opt,name=baidu_secret_key,json=baiduSecretKey,proto3" json:"baidu_secret_key,omitempty"`
	BaiduEndpoint                   string `protobuf:"bytes,37,opt,name=baidu_endpoint,json=baiduEndpoint,proto3" json:"baidu_endpoint,omitempty"`
	BaiduRegion                     string `protobuf:"bytes,38,opt,name=baidu_region,json=baiduRegion,proto3" json:"baidu_region,omitempty"`
	WasabiAccessKey                 string `protobuf:"bytes,40,opt,name=wasabi_access_key,json=wasabiAccessKey,proto3" json:"wasabi_access_key,omitempty"`
	WasabiSecretKey                 string `protobuf:"bytes,41,opt,name=wasabi_secret_key,json=wasabiSecretKey,proto3" json:"wasabi_secret_key,omitempty"`
	WasabiEndpoint                  string `protobuf:"bytes,42,opt,name=wasabi_endpoint,json=wasabiEndpoint,proto3" json:"wasabi_endpoint,omitempty"`
	WasabiRegion                    string `protobuf:"bytes,43,opt,name=wasabi_region,json=wasabiRegion,proto3" json:"wasabi_region,omitempty"`
	FilebaseAccessKey               string `protobuf:"bytes,60,opt,name=filebase_access_key,json=filebaseAccessKey,proto3" json:"filebase_access_key,omitempty"`
	FilebaseSecretKey               string `protobuf:"bytes,61,opt,name=filebase_secret_key,json=filebaseSecretKey,proto3" json:"filebase_secret_key,omitempty"`
	FilebaseEndpoint                string `protobuf:"bytes,62,opt,name=filebase_endpoint,json=filebaseEndpoint,proto3" json:"filebase_endpoint,omitempty"`
	StorjAccessKey                  string `protobuf:"bytes,65,opt,name=storj_access_key,json=storjAccessKey,proto3" json:"storj_access_key,omitempty"`
	StorjSecretKey                  string `protobuf:"bytes,66,opt,name=storj_secret_key,json=storjSecretKey,proto3" json:"storj_secret_key,omitempty"`
	StorjEndpoint                   string `protobuf:"bytes,67,opt,name=storj_endpoint,json=storjEndpoint,proto3" json:"storj_endpoint,omitempty"`
	ContaboAccessKey                string `protobuf:"bytes,68,opt,name=contabo_access_key,json=contaboAccessKey,proto3" json:"contabo_access_key,omitempty"`
	ContaboSecretKey                string `protobuf:"bytes,69,opt,name=contabo_secret_key,json=contaboSecretKey,proto3" json:"contabo_secret_key,omitempty"`
	ContaboEndpoint                 string `protobuf:"bytes,70,opt,name=contabo_endpoint,json=contaboEndpoint,proto3" json:"contabo_endpoint,omitempty"`
	ContaboRegion                   string `protobuf:"bytes,71,opt,name=contabo_region,json=contaboRegion,proto3" json:"contabo_region,omitempty"`
}

func (x *RemoteConf) Reset() {
	*x = RemoteConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteConf) ProtoMessage() {}

func (x *RemoteConf) ProtoReflect() protoreflect.Message {
	mi := &file_remote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteConf.ProtoReflect.Descriptor instead.
func (*RemoteConf) Descriptor() ([]byte, []int) {
	return file_remote_proto_rawDescGZIP(), []int{0}
}

func (x *RemoteConf) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RemoteConf) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RemoteConf) GetS3AccessKey() string {
	if x != nil {
		return x.S3AccessKey
	}
	return ""
}

func (x *RemoteConf) GetS3SecretKey() string {
	if x != nil {
		return x.S3SecretKey
	}
	return ""
}

func (x *RemoteConf) GetS3Region() string {
	if x != nil {
		return x.S3Region
	}
	return ""
}

func (x *RemoteConf) GetS3Endpoint() string {
	if x != nil {
		return x.S3Endpoint
	}
	return ""
}

func (x *RemoteConf) GetS3StorageClass() string {
	if x != nil {
		return x.S3StorageClass
	}
	return ""
}

func (x *RemoteConf) GetS3ForcePathStyle() bool {
	if x != nil {
		return x.S3ForcePathStyle
	}
	return false
}

func (x *RemoteConf) GetS3SupportTagging() bool {
	if x != nil {
		return x.S3SupportTagging
	}
	return false
}

func (x *RemoteConf) GetS3V4Signature() bool {
	if x != nil {
		return x.S3V4Signature
	}
	return false
}

func (x *RemoteConf) GetGcsGoogleApplicationCredentials() string {
	if x != nil {
		return x.GcsGoogleApplicationCredentials
	}
	return ""
}

func (x *RemoteConf) GetGcsProjectId() string {
	if x != nil {
		return x.GcsProjectId
	}
	return ""
}

func (x *RemoteConf) GetAzureAccountName() string {
	if x != nil {
		return x.AzureAccountName
	}
	return ""
}

func (x *RemoteConf) GetAzureAccountKey() string {
	if x != nil {
		return x.AzureAccountKey
	}
	return ""
}

func (x *RemoteConf) GetBackblazeKeyId() string {
	if x != nil {
		return x.BackblazeKeyId
	}
	return ""
}

func (x *RemoteConf) GetBackblazeApplicationKey() string {
	if x != nil {
		return x.BackblazeApplicationKey
	}
	return ""
}

func (x *RemoteConf) GetBackblazeEndpoint() string {
	if x != nil {
		return x.BackblazeEndpoint
	}
	return ""
}

func (x *RemoteConf) GetBackblazeRegion() string {
	if x != nil {
		return x.BackblazeRegion
	}
	return ""
}

func (x *RemoteConf) GetAliyunAccessKey() string {
	if x != nil {
		return x.AliyunAccessKey
	}
	return ""
}

func (x *RemoteConf) GetAliyunSecretKey() string {
	if x != nil {
		return x.AliyunSecretKey
	}
	return ""
}

func (x *RemoteConf) GetAliyunEndpoint() string {
	if x != nil {
		return x.AliyunEndpoint
	}
	return ""
}

func (x *RemoteConf) GetAliyunRegion() string {
	if x != nil {
		return x.AliyunRegion
	}
	return ""
}

func (x *RemoteConf) GetTencentSecretId() string {
	if x != nil {
		return x.TencentSecretId
	}
	return ""
}

func (x *RemoteConf) GetTencentSecretKey() string {
	if x != nil {
		return x.TencentSecretKey
	}
	return ""
}

func (x *RemoteConf) GetTencentEndpoint() string {
	if x != nil {
		return x.TencentEndpoint
	}
	return ""
}

func (x *RemoteConf) GetBaiduAccessKey() string {
	if x != nil {
		return x.BaiduAccessKey
	}
	return ""
}

func (x *RemoteConf) GetBaiduSecretKey() string {
	if x != nil {
		return x.BaiduSecretKey
	}
	return ""
}

func (x *RemoteConf) GetBaiduEndpoint() string {
	if x != nil {
		return x.BaiduEndpoint
	}
	return ""
}

func (x *RemoteConf) GetBaiduRegion() string {
	if x != nil {
		return x.BaiduRegion
	}
	return ""
}

func (x *RemoteConf) GetWasabiAccessKey() string {
	if x != nil {
		return x.WasabiAccessKey
	}
	return ""
}

func (x *RemoteConf) GetWasabiSecretKey() string {
	if x != nil {
		return x.WasabiSecretKey
	}
	return ""
}

func (x *RemoteConf) GetWasabiEndpoint() string {
	if x != nil {
		return x.WasabiEndpoint
	}
	return ""
}

func (x *RemoteConf) GetWasabiRegion() string {
	if x != nil {
		return x.WasabiRegion
	}
	return ""
}

func (x *RemoteConf) GetFilebaseAccessKey() string {
	if x != nil {
		return x.FilebaseAccessKey
	}
	return ""
}

func (x *RemoteConf) GetFilebaseSecretKey() string {
	if x != nil {
		return x.FilebaseSecretKey
	}
	return ""
}

func (x *RemoteConf) GetFilebaseEndpoint() string {
	if x != nil {
		return x.FilebaseEndpoint
	}
	return ""
}

func (x *RemoteConf) GetStorjAccessKey() string {
	if x != nil {
		return x.StorjAccessKey
	}
	return ""
}

func (x *RemoteConf) GetStorjSecretKey() string {
	if x != nil {
		return x.StorjSecretKey
	}
	return ""
}

func (x *RemoteConf) GetStorjEndpoint() string {
	if x != nil {
		return x.StorjEndpoint
	}
	return ""
}

func (x *RemoteConf) GetContaboAccessKey() string {
	if x != nil {
		return x.ContaboAccessKey
	}
	return ""
}

func (x *RemoteConf) GetContaboSecretKey() string {
	if x != nil {
		return x.ContaboSecretKey
	}
	return ""
}

func (x *RemoteConf) GetContaboEndpoint() string {
	if x != nil {
		return x.ContaboEndpoint
	}
	return ""
}

func (x *RemoteConf) GetContaboRegion() string {
	if x != nil {
		return x.ContaboRegion
	}
	return ""
}

type RemoteStorageMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mappings                 map[string]*RemoteStorageLocation `protobuf:"bytes,1,rep,name=mappings,proto3" json:"mappings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	PrimaryBucketStorageName string                            `protobuf:"bytes,2,opt,name=primary_bucket_storage_name,json=primaryBucketStorageName,proto3" json:"primary_bucket_storage_name,omitempty"`
}

func (x *RemoteStorageMapping) Reset() {
	*x = RemoteStorageMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteStorageMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteStorageMapping) ProtoMessage() {}

func (x *RemoteStorageMapping) ProtoReflect() protoreflect.Message {
	mi := &file_remote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteStorageMapping.ProtoReflect.Descriptor instead.
func (*RemoteStorageMapping) Descriptor() ([]byte, []int) {
	return file_remote_proto_rawDescGZIP(), []int{1}
}

func (x *RemoteStorageMapping) GetMappings() map[string]*RemoteStorageLocation {
	if x != nil {
		return x.Mappings
	}
	return nil
}

func (x *RemoteStorageMapping) GetPrimaryBucketStorageName() string {
	if x != nil {
		return x.PrimaryBucketStorageName
	}
	return ""
}

type RemoteStorageLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Bucket string `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Path   string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *RemoteStorageLocation) Reset() {
	*x = RemoteStorageLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteStorageLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteStorageLocation) ProtoMessage() {}

func (x *RemoteStorageLocation) ProtoReflect() protoreflect.Message {
	mi := &file_remote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteStorageLocation.ProtoReflect.Descriptor instead.
func (*RemoteStorageLocation) Descriptor() ([]byte, []int) {
	return file_remote_proto_rawDescGZIP(), []int{2}
}

func (x *RemoteStorageLocation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RemoteStorageLocation) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *RemoteStorageLocation) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_remote_proto protoreflect.FileDescriptor

var file_remote_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x62, 0x22, 0x9b, 0x0e, 0x0a, 0x0a, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0d, 0x73, 0x33, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x33, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x33, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x33, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x33, 0x5f, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x33, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x33, 0x5f, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x33, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x33, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x73, 0x33, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x12, 0x2d, 0x0a, 0x13, 0x73, 0x33, 0x5f, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x73,
	0x33, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12,
	0x2c, 0x0a, 0x12, 0x73, 0x33, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x61,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x73, 0x33, 0x53,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x61, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a,
	0x0f, 0x73, 0x33, 0x5f, 0x76, 0x34, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x33, 0x56, 0x34, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x4b, 0x0a, 0x22, 0x67, 0x63, 0x73, 0x5f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x1f, 0x67, 0x63, 0x73, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x67, 0x63, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x63, 0x73, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x7a, 0x75, 0x72,
	0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b,
	0x65, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x62, 0x61, 0x63, 0x6b, 0x62, 0x6c, 0x61, 0x7a, 0x65, 0x5f,
	0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x61,
	0x63, 0x6b, 0x62, 0x6c, 0x61, 0x7a, 0x65, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x19,
	0x62, 0x61, 0x63, 0x6b, 0x62, 0x6c, 0x61, 0x7a, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x17, 0x62, 0x61, 0x63, 0x6b, 0x62, 0x6c, 0x61, 0x7a, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x12, 0x62, 0x61, 0x63, 0x6b,
	0x62, 0x6c, 0x61, 0x7a, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x62, 0x61, 0x63, 0x6b, 0x62, 0x6c, 0x61, 0x7a, 0x65, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x62, 0x61, 0x63, 0x6b, 0x62,
	0x6c, 0x61, 0x7a, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x17, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x62, 0x6c, 0x61, 0x7a, 0x65, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x5f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61,
	0x6c, 0x69, 0x79, 0x75, 0x6e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x2a,
	0x0a, 0x11, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x6c, 0x69, 0x79, 0x75,
	0x6e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c,
	0x69, 0x79, 0x75, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x1b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x5f, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x6c, 0x69, 0x79,
	0x75, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x65, 0x6e, 0x63,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b,
	0x65, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x65,
	0x6e, 0x63, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x28, 0x0a,
	0x10, 0x62, 0x61, 0x69, 0x64, 0x75, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x61, 0x69, 0x64, 0x75, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x62, 0x61, 0x69, 0x64, 0x75,
	0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x24, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x62, 0x61, 0x69, 0x64, 0x75, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65,
	0x79, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x61, 0x69, 0x64, 0x75, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x25, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x61, 0x69, 0x64, 0x75,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x61, 0x69, 0x64,
	0x75, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x26, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x62, 0x61, 0x69, 0x64, 0x75, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x77,
	0x61, 0x73, 0x61, 0x62, 0x69, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x77, 0x61, 0x73, 0x61, 0x62, 0x69, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x11, 0x77, 0x61, 0x73, 0x61, 0x62,
	0x69, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x29, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x77, 0x61, 0x73, 0x61, 0x62, 0x69, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x4b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x77, 0x61, 0x73, 0x61, 0x62, 0x69, 0x5f, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x77, 0x61,
	0x73, 0x61, 0x62, 0x69, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x77, 0x61, 0x73, 0x61, 0x62, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x2b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x61, 0x73, 0x61, 0x62, 0x69, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x66, 0x69, 0x6c, 0x65, 0x62, 0x61, 0x73, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65,
	0x79, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x66, 0x69, 0x6c, 0x65, 0x62, 0x61, 0x73, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65,
	0x79, 0x12, 0x2b, 0x0a, 0x11, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x3e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x66, 0x69,
	0x6c, 0x65, 0x62, 0x61, 0x73, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x28,
	0x0a, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x6a, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x41, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x6a, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x74, 0x6f, 0x72,
	0x6a, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x42, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x6a, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b,
	0x65, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x6a, 0x5f, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x43, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72,
	0x6a, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x62, 0x6f, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x44, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x62, 0x6f, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x62, 0x6f, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x45, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x62, 0x6f, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x62, 0x6f,
	0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x62, 0x6f, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x62, 0x6f, 0x5f, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x47, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x62,
	0x6f, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0xff, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x12, 0x49, 0x0a, 0x08, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3d, 0x0a, 0x1b, 0x70,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x18, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x5d, 0x0a, 0x0d, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x36, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x57, 0x0a, 0x15, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x42, 0x50, 0x0a, 0x10, 0x73, 0x65, 0x61, 0x77, 0x65, 0x65, 0x64, 0x66, 0x73, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x65, 0x61, 0x77, 0x65, 0x65, 0x64, 0x66, 0x73, 0x2f, 0x73, 0x65, 0x61, 0x77, 0x65, 0x65, 0x64,
	0x66, 0x73, 0x2f, 0x77, 0x65, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_remote_proto_rawDescOnce sync.Once
	file_remote_proto_rawDescData = file_remote_proto_rawDesc
)

func file_remote_proto_rawDescGZIP() []byte {
	file_remote_proto_rawDescOnce.Do(func() {
		file_remote_proto_rawDescData = protoimpl.X.CompressGZIP(file_remote_proto_rawDescData)
	})
	return file_remote_proto_rawDescData
}

var file_remote_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_remote_proto_goTypes = []interface{}{
	(*RemoteConf)(nil),            // 0: remote_pb.RemoteConf
	(*RemoteStorageMapping)(nil),  // 1: remote_pb.RemoteStorageMapping
	(*RemoteStorageLocation)(nil), // 2: remote_pb.RemoteStorageLocation
	nil,                           // 3: remote_pb.RemoteStorageMapping.MappingsEntry
}
var file_remote_proto_depIdxs = []int32{
	3, // 0: remote_pb.RemoteStorageMapping.mappings:type_name -> remote_pb.RemoteStorageMapping.MappingsEntry
	2, // 1: remote_pb.RemoteStorageMapping.MappingsEntry.value:type_name -> remote_pb.RemoteStorageLocation
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_remote_proto_init() }
func file_remote_proto_init() {
	if File_remote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_remote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteConf); i {
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
		file_remote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteStorageMapping); i {
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
		file_remote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteStorageLocation); i {
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
			RawDescriptor: file_remote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_remote_proto_goTypes,
		DependencyIndexes: file_remote_proto_depIdxs,
		MessageInfos:      file_remote_proto_msgTypes,
	}.Build()
	File_remote_proto = out.File
	file_remote_proto_rawDesc = nil
	file_remote_proto_goTypes = nil
	file_remote_proto_depIdxs = nil
}
