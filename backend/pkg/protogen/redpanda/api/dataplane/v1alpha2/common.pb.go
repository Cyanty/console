// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: redpanda/api/dataplane/v1alpha2/common.proto

package dataplanev1alpha2

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConfigSource int32

const (
	ConfigSource_CONFIG_SOURCE_UNSPECIFIED                   ConfigSource = 0
	ConfigSource_CONFIG_SOURCE_DYNAMIC_TOPIC_CONFIG          ConfigSource = 1
	ConfigSource_CONFIG_SOURCE_DYNAMIC_BROKER_CONFIG         ConfigSource = 2
	ConfigSource_CONFIG_SOURCE_DYNAMIC_DEFAULT_BROKER_CONFIG ConfigSource = 3
	ConfigSource_CONFIG_SOURCE_STATIC_BROKER_CONFIG          ConfigSource = 4
	ConfigSource_CONFIG_SOURCE_DEFAULT_CONFIG                ConfigSource = 5
	ConfigSource_CONFIG_SOURCE_DYNAMIC_BROKER_LOGGER_CONFIG  ConfigSource = 6
)

// Enum value maps for ConfigSource.
var (
	ConfigSource_name = map[int32]string{
		0: "CONFIG_SOURCE_UNSPECIFIED",
		1: "CONFIG_SOURCE_DYNAMIC_TOPIC_CONFIG",
		2: "CONFIG_SOURCE_DYNAMIC_BROKER_CONFIG",
		3: "CONFIG_SOURCE_DYNAMIC_DEFAULT_BROKER_CONFIG",
		4: "CONFIG_SOURCE_STATIC_BROKER_CONFIG",
		5: "CONFIG_SOURCE_DEFAULT_CONFIG",
		6: "CONFIG_SOURCE_DYNAMIC_BROKER_LOGGER_CONFIG",
	}
	ConfigSource_value = map[string]int32{
		"CONFIG_SOURCE_UNSPECIFIED":                   0,
		"CONFIG_SOURCE_DYNAMIC_TOPIC_CONFIG":          1,
		"CONFIG_SOURCE_DYNAMIC_BROKER_CONFIG":         2,
		"CONFIG_SOURCE_DYNAMIC_DEFAULT_BROKER_CONFIG": 3,
		"CONFIG_SOURCE_STATIC_BROKER_CONFIG":          4,
		"CONFIG_SOURCE_DEFAULT_CONFIG":                5,
		"CONFIG_SOURCE_DYNAMIC_BROKER_LOGGER_CONFIG":  6,
	}
)

func (x ConfigSource) Enum() *ConfigSource {
	p := new(ConfigSource)
	*p = x
	return p
}

func (x ConfigSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfigSource) Descriptor() protoreflect.EnumDescriptor {
	return file_redpanda_api_dataplane_v1alpha2_common_proto_enumTypes[0].Descriptor()
}

func (ConfigSource) Type() protoreflect.EnumType {
	return &file_redpanda_api_dataplane_v1alpha2_common_proto_enumTypes[0]
}

func (x ConfigSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfigSource.Descriptor instead.
func (ConfigSource) EnumDescriptor() ([]byte, []int) {
	return file_redpanda_api_dataplane_v1alpha2_common_proto_rawDescGZIP(), []int{0}
}

type ConfigType int32

const (
	ConfigType_CONFIG_TYPE_UNSPECIFIED ConfigType = 0
	ConfigType_CONFIG_TYPE_BOOLEAN     ConfigType = 1
	ConfigType_CONFIG_TYPE_STRING      ConfigType = 2
	ConfigType_CONFIG_TYPE_INT         ConfigType = 3
	ConfigType_CONFIG_TYPE_SHORT       ConfigType = 4
	ConfigType_CONFIG_TYPE_LONG        ConfigType = 5
	ConfigType_CONFIG_TYPE_DOUBLE      ConfigType = 6
	ConfigType_CONFIG_TYPE_LIST        ConfigType = 7
	ConfigType_CONFIG_TYPE_CLASS       ConfigType = 8
	ConfigType_CONFIG_TYPE_PASSWORD    ConfigType = 9
)

// Enum value maps for ConfigType.
var (
	ConfigType_name = map[int32]string{
		0: "CONFIG_TYPE_UNSPECIFIED",
		1: "CONFIG_TYPE_BOOLEAN",
		2: "CONFIG_TYPE_STRING",
		3: "CONFIG_TYPE_INT",
		4: "CONFIG_TYPE_SHORT",
		5: "CONFIG_TYPE_LONG",
		6: "CONFIG_TYPE_DOUBLE",
		7: "CONFIG_TYPE_LIST",
		8: "CONFIG_TYPE_CLASS",
		9: "CONFIG_TYPE_PASSWORD",
	}
	ConfigType_value = map[string]int32{
		"CONFIG_TYPE_UNSPECIFIED": 0,
		"CONFIG_TYPE_BOOLEAN":     1,
		"CONFIG_TYPE_STRING":      2,
		"CONFIG_TYPE_INT":         3,
		"CONFIG_TYPE_SHORT":       4,
		"CONFIG_TYPE_LONG":        5,
		"CONFIG_TYPE_DOUBLE":      6,
		"CONFIG_TYPE_LIST":        7,
		"CONFIG_TYPE_CLASS":       8,
		"CONFIG_TYPE_PASSWORD":    9,
	}
)

func (x ConfigType) Enum() *ConfigType {
	p := new(ConfigType)
	*p = x
	return p
}

func (x ConfigType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfigType) Descriptor() protoreflect.EnumDescriptor {
	return file_redpanda_api_dataplane_v1alpha2_common_proto_enumTypes[1].Descriptor()
}

func (ConfigType) Type() protoreflect.EnumType {
	return &file_redpanda_api_dataplane_v1alpha2_common_proto_enumTypes[1]
}

func (x ConfigType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfigType.Descriptor instead.
func (ConfigType) EnumDescriptor() ([]byte, []int) {
	return file_redpanda_api_dataplane_v1alpha2_common_proto_rawDescGZIP(), []int{1}
}

type ConfigAlterOperation int32

const (
	ConfigAlterOperation_CONFIG_ALTER_OPERATION_UNSPECIFIED ConfigAlterOperation = 0
	ConfigAlterOperation_CONFIG_ALTER_OPERATION_SET         ConfigAlterOperation = 1
	ConfigAlterOperation_CONFIG_ALTER_OPERATION_DELETE      ConfigAlterOperation = 2
	ConfigAlterOperation_CONFIG_ALTER_OPERATION_APPEND      ConfigAlterOperation = 3
	ConfigAlterOperation_CONFIG_ALTER_OPERATION_SUBTRACT    ConfigAlterOperation = 4
)

// Enum value maps for ConfigAlterOperation.
var (
	ConfigAlterOperation_name = map[int32]string{
		0: "CONFIG_ALTER_OPERATION_UNSPECIFIED",
		1: "CONFIG_ALTER_OPERATION_SET",
		2: "CONFIG_ALTER_OPERATION_DELETE",
		3: "CONFIG_ALTER_OPERATION_APPEND",
		4: "CONFIG_ALTER_OPERATION_SUBTRACT",
	}
	ConfigAlterOperation_value = map[string]int32{
		"CONFIG_ALTER_OPERATION_UNSPECIFIED": 0,
		"CONFIG_ALTER_OPERATION_SET":         1,
		"CONFIG_ALTER_OPERATION_DELETE":      2,
		"CONFIG_ALTER_OPERATION_APPEND":      3,
		"CONFIG_ALTER_OPERATION_SUBTRACT":    4,
	}
)

func (x ConfigAlterOperation) Enum() *ConfigAlterOperation {
	p := new(ConfigAlterOperation)
	*p = x
	return p
}

func (x ConfigAlterOperation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfigAlterOperation) Descriptor() protoreflect.EnumDescriptor {
	return file_redpanda_api_dataplane_v1alpha2_common_proto_enumTypes[2].Descriptor()
}

func (ConfigAlterOperation) Type() protoreflect.EnumType {
	return &file_redpanda_api_dataplane_v1alpha2_common_proto_enumTypes[2]
}

func (x ConfigAlterOperation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfigAlterOperation.Descriptor instead.
func (ConfigAlterOperation) EnumDescriptor() ([]byte, []int) {
	return file_redpanda_api_dataplane_v1alpha2_common_proto_rawDescGZIP(), []int{2}
}

type ConfigSynonym struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value         *string                `protobuf:"bytes,2,opt,name=value,proto3,oneof" json:"value,omitempty"`
	Source        ConfigSource           `protobuf:"varint,3,opt,name=source,proto3,enum=redpanda.api.dataplane.v1alpha2.ConfigSource" json:"source,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConfigSynonym) Reset() {
	*x = ConfigSynonym{}
	mi := &file_redpanda_api_dataplane_v1alpha2_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfigSynonym) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigSynonym) ProtoMessage() {}

func (x *ConfigSynonym) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_dataplane_v1alpha2_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigSynonym.ProtoReflect.Descriptor instead.
func (*ConfigSynonym) Descriptor() ([]byte, []int) {
	return file_redpanda_api_dataplane_v1alpha2_common_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigSynonym) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConfigSynonym) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

func (x *ConfigSynonym) GetSource() ConfigSource {
	if x != nil {
		return x.Source
	}
	return ConfigSource_CONFIG_SOURCE_UNSPECIFIED
}

var File_redpanda_api_dataplane_v1alpha2_common_proto protoreflect.FileDescriptor

var file_redpanda_api_dataplane_v1alpha2_common_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x22,
	0x8f, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x79, 0x6e, 0x6f, 0x6e, 0x79,
	0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x45, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2d, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x2a, 0xa9, 0x02, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x26, 0x0a, 0x22, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x5f, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x54, 0x4f, 0x50, 0x49, 0x43,
	0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10, 0x01, 0x12, 0x27, 0x0a, 0x23, 0x43, 0x4f, 0x4e,
	0x46, 0x49, 0x47, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x44, 0x59, 0x4e, 0x41, 0x4d,
	0x49, 0x43, 0x5f, 0x42, 0x52, 0x4f, 0x4b, 0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47,
	0x10, 0x02, 0x12, 0x2f, 0x0a, 0x2b, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x44, 0x45, 0x46, 0x41,
	0x55, 0x4c, 0x54, 0x5f, 0x42, 0x52, 0x4f, 0x4b, 0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49,
	0x47, 0x10, 0x03, 0x12, 0x26, 0x0a, 0x22, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x49, 0x43, 0x5f, 0x42, 0x52, 0x4f, 0x4b,
	0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10, 0x04, 0x12, 0x20, 0x0a, 0x1c, 0x43,
	0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x44, 0x45, 0x46,
	0x41, 0x55, 0x4c, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10, 0x05, 0x12, 0x2e, 0x0a,
	0x2a, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x44,
	0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x42, 0x52, 0x4f, 0x4b, 0x45, 0x52, 0x5f, 0x4c, 0x4f,
	0x47, 0x47, 0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10, 0x06, 0x2a, 0xfb, 0x01,
	0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17,
	0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4e,
	0x46, 0x49, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x45, 0x41, 0x4e,
	0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f,
	0x4e, 0x46, 0x49, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x10, 0x03, 0x12,
	0x15, 0x0a, 0x11, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53,
	0x48, 0x4f, 0x52, 0x54, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12,
	0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x4f, 0x55, 0x42,
	0x4c, 0x45, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x07, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f,
	0x4e, 0x46, 0x49, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x10,
	0x08, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x09, 0x2a, 0xc9, 0x01, 0x0a, 0x14,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x22, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x41,
	0x4c, 0x54, 0x45, 0x52, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a,
	0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x41, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x4f, 0x50, 0x45,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d,
	0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x41, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x4f, 0x50, 0x45,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x12,
	0x21, 0x0a, 0x1d, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x41, 0x4c, 0x54, 0x45, 0x52, 0x5f,
	0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x50, 0x50, 0x45, 0x4e, 0x44,
	0x10, 0x03, 0x12, 0x23, 0x0a, 0x1f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x41, 0x4c, 0x54,
	0x45, 0x52, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x55, 0x42,
	0x54, 0x52, 0x41, 0x43, 0x54, 0x10, 0x04, 0x42, 0xba, 0x02, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e,
	0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42,
	0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x67,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x64, 0x70, 0x61,
	0x6e, 0x64, 0x61, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xa2, 0x02, 0x03, 0x52, 0x41, 0x44, 0xaa, 0x02, 0x1f,
	0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xca,
	0x02, 0x1f, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x44,
	0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0xe2, 0x02, 0x2b, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x5c, 0x41, 0x70, 0x69,
	0x5c, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x22, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a,
	0x3a, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_redpanda_api_dataplane_v1alpha2_common_proto_rawDescOnce sync.Once
	file_redpanda_api_dataplane_v1alpha2_common_proto_rawDescData = file_redpanda_api_dataplane_v1alpha2_common_proto_rawDesc
)

func file_redpanda_api_dataplane_v1alpha2_common_proto_rawDescGZIP() []byte {
	file_redpanda_api_dataplane_v1alpha2_common_proto_rawDescOnce.Do(func() {
		file_redpanda_api_dataplane_v1alpha2_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_redpanda_api_dataplane_v1alpha2_common_proto_rawDescData)
	})
	return file_redpanda_api_dataplane_v1alpha2_common_proto_rawDescData
}

var file_redpanda_api_dataplane_v1alpha2_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_redpanda_api_dataplane_v1alpha2_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_redpanda_api_dataplane_v1alpha2_common_proto_goTypes = []any{
	(ConfigSource)(0),         // 0: redpanda.api.dataplane.v1alpha2.ConfigSource
	(ConfigType)(0),           // 1: redpanda.api.dataplane.v1alpha2.ConfigType
	(ConfigAlterOperation)(0), // 2: redpanda.api.dataplane.v1alpha2.ConfigAlterOperation
	(*ConfigSynonym)(nil),     // 3: redpanda.api.dataplane.v1alpha2.ConfigSynonym
}
var file_redpanda_api_dataplane_v1alpha2_common_proto_depIdxs = []int32{
	0, // 0: redpanda.api.dataplane.v1alpha2.ConfigSynonym.source:type_name -> redpanda.api.dataplane.v1alpha2.ConfigSource
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_redpanda_api_dataplane_v1alpha2_common_proto_init() }
func file_redpanda_api_dataplane_v1alpha2_common_proto_init() {
	if File_redpanda_api_dataplane_v1alpha2_common_proto != nil {
		return
	}
	file_redpanda_api_dataplane_v1alpha2_common_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_redpanda_api_dataplane_v1alpha2_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_redpanda_api_dataplane_v1alpha2_common_proto_goTypes,
		DependencyIndexes: file_redpanda_api_dataplane_v1alpha2_common_proto_depIdxs,
		EnumInfos:         file_redpanda_api_dataplane_v1alpha2_common_proto_enumTypes,
		MessageInfos:      file_redpanda_api_dataplane_v1alpha2_common_proto_msgTypes,
	}.Build()
	File_redpanda_api_dataplane_v1alpha2_common_proto = out.File
	file_redpanda_api_dataplane_v1alpha2_common_proto_rawDesc = nil
	file_redpanda_api_dataplane_v1alpha2_common_proto_goTypes = nil
	file_redpanda_api_dataplane_v1alpha2_common_proto_depIdxs = nil
}
