// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/resultstore/v2/configured_target.proto

package resultstore // import "github.com/Beeketing/go-genproto/googleapis/devtools/resultstore/v2"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/Beeketing/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Each ConfiguredTarget represents data for a given configuration of a given
// target in a given Invocation.
// Every ConfiguredTarget should have at least one Action as a child resource
// before the invocation is finalized. Refer to the Action's documentation for
// more info on this.
type ConfiguredTarget struct {
	// The resource name.  Its format must be:
	// invocations/${INVOCATION_ID}/targets/${TARGET_ID}/configuredTargets/${CONFIG_ID}
	// where ${CONFIG_ID} must match the ID of an existing Configuration under
	// this Invocation.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The resource ID components that identify the ConfiguredTarget. They must
	// match the resource name after proper encoding.
	Id *ConfiguredTarget_Id `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The aggregate status for this configuration of this target. If testing
	// was not requested, set this to the build status (e.g. BUILT or
	// FAILED_TO_BUILD).
	StatusAttributes *StatusAttributes `protobuf:"bytes,3,opt,name=status_attributes,json=statusAttributes,proto3" json:"status_attributes,omitempty"`
	// Captures the start time and duration of this configured target.
	Timing *Timing `protobuf:"bytes,4,opt,name=timing,proto3" json:"timing,omitempty"`
	// Test specific attributes for this ConfiguredTarget.
	TestAttributes *ConfiguredTestAttributes `protobuf:"bytes,6,opt,name=test_attributes,json=testAttributes,proto3" json:"test_attributes,omitempty"`
	// Arbitrary name-value pairs.
	// This is implemented as a multi-map. Multiple properties are allowed with
	// the same key. Properties will be returned in lexicographical order by key.
	Properties []*Property `protobuf:"bytes,7,rep,name=properties,proto3" json:"properties,omitempty"`
	// A list of file references for configured target level files.
	// The file IDs must be unique within this list. Duplicate file IDs will
	// result in an error. Files will be returned in lexicographical order by ID.
	Files                []*File  `protobuf:"bytes,8,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfiguredTarget) Reset()         { *m = ConfiguredTarget{} }
func (m *ConfiguredTarget) String() string { return proto.CompactTextString(m) }
func (*ConfiguredTarget) ProtoMessage()    {}
func (*ConfiguredTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_configured_target_dc7d1ae233b30620, []int{0}
}
func (m *ConfiguredTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfiguredTarget.Unmarshal(m, b)
}
func (m *ConfiguredTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfiguredTarget.Marshal(b, m, deterministic)
}
func (dst *ConfiguredTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfiguredTarget.Merge(dst, src)
}
func (m *ConfiguredTarget) XXX_Size() int {
	return xxx_messageInfo_ConfiguredTarget.Size(m)
}
func (m *ConfiguredTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfiguredTarget.DiscardUnknown(m)
}

var xxx_messageInfo_ConfiguredTarget proto.InternalMessageInfo

func (m *ConfiguredTarget) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ConfiguredTarget) GetId() *ConfiguredTarget_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConfiguredTarget) GetStatusAttributes() *StatusAttributes {
	if m != nil {
		return m.StatusAttributes
	}
	return nil
}

func (m *ConfiguredTarget) GetTiming() *Timing {
	if m != nil {
		return m.Timing
	}
	return nil
}

func (m *ConfiguredTarget) GetTestAttributes() *ConfiguredTestAttributes {
	if m != nil {
		return m.TestAttributes
	}
	return nil
}

func (m *ConfiguredTarget) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *ConfiguredTarget) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

// The resource ID components that identify the ConfiguredTarget.
type ConfiguredTarget_Id struct {
	// The Invocation ID.
	InvocationId string `protobuf:"bytes,1,opt,name=invocation_id,json=invocationId,proto3" json:"invocation_id,omitempty"`
	// The Target ID.
	TargetId string `protobuf:"bytes,2,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	// The Configuration ID.
	ConfigurationId      string   `protobuf:"bytes,3,opt,name=configuration_id,json=configurationId,proto3" json:"configuration_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfiguredTarget_Id) Reset()         { *m = ConfiguredTarget_Id{} }
func (m *ConfiguredTarget_Id) String() string { return proto.CompactTextString(m) }
func (*ConfiguredTarget_Id) ProtoMessage()    {}
func (*ConfiguredTarget_Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_configured_target_dc7d1ae233b30620, []int{0, 0}
}
func (m *ConfiguredTarget_Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfiguredTarget_Id.Unmarshal(m, b)
}
func (m *ConfiguredTarget_Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfiguredTarget_Id.Marshal(b, m, deterministic)
}
func (dst *ConfiguredTarget_Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfiguredTarget_Id.Merge(dst, src)
}
func (m *ConfiguredTarget_Id) XXX_Size() int {
	return xxx_messageInfo_ConfiguredTarget_Id.Size(m)
}
func (m *ConfiguredTarget_Id) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfiguredTarget_Id.DiscardUnknown(m)
}

var xxx_messageInfo_ConfiguredTarget_Id proto.InternalMessageInfo

func (m *ConfiguredTarget_Id) GetInvocationId() string {
	if m != nil {
		return m.InvocationId
	}
	return ""
}

func (m *ConfiguredTarget_Id) GetTargetId() string {
	if m != nil {
		return m.TargetId
	}
	return ""
}

func (m *ConfiguredTarget_Id) GetConfigurationId() string {
	if m != nil {
		return m.ConfigurationId
	}
	return ""
}

// Attributes that apply only to test actions under this configured target.
type ConfiguredTestAttributes struct {
	// Total number of test runs. For example, in bazel this is specified with
	// --runs_per_test. Zero if runs_per_test is not used.
	TotalRunCount int32 `protobuf:"varint,2,opt,name=total_run_count,json=totalRunCount,proto3" json:"total_run_count,omitempty"`
	// Total number of test shards. Zero if shard count was not specified.
	TotalShardCount int32 `protobuf:"varint,3,opt,name=total_shard_count,json=totalShardCount,proto3" json:"total_shard_count,omitempty"`
	// How long test is allowed to run.
	TimeoutDuration      *duration.Duration `protobuf:"bytes,5,opt,name=timeout_duration,json=timeoutDuration,proto3" json:"timeout_duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ConfiguredTestAttributes) Reset()         { *m = ConfiguredTestAttributes{} }
func (m *ConfiguredTestAttributes) String() string { return proto.CompactTextString(m) }
func (*ConfiguredTestAttributes) ProtoMessage()    {}
func (*ConfiguredTestAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_configured_target_dc7d1ae233b30620, []int{1}
}
func (m *ConfiguredTestAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfiguredTestAttributes.Unmarshal(m, b)
}
func (m *ConfiguredTestAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfiguredTestAttributes.Marshal(b, m, deterministic)
}
func (dst *ConfiguredTestAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfiguredTestAttributes.Merge(dst, src)
}
func (m *ConfiguredTestAttributes) XXX_Size() int {
	return xxx_messageInfo_ConfiguredTestAttributes.Size(m)
}
func (m *ConfiguredTestAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfiguredTestAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ConfiguredTestAttributes proto.InternalMessageInfo

func (m *ConfiguredTestAttributes) GetTotalRunCount() int32 {
	if m != nil {
		return m.TotalRunCount
	}
	return 0
}

func (m *ConfiguredTestAttributes) GetTotalShardCount() int32 {
	if m != nil {
		return m.TotalShardCount
	}
	return 0
}

func (m *ConfiguredTestAttributes) GetTimeoutDuration() *duration.Duration {
	if m != nil {
		return m.TimeoutDuration
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfiguredTarget)(nil), "google.devtools.resultstore.v2.ConfiguredTarget")
	proto.RegisterType((*ConfiguredTarget_Id)(nil), "google.devtools.resultstore.v2.ConfiguredTarget.Id")
	proto.RegisterType((*ConfiguredTestAttributes)(nil), "google.devtools.resultstore.v2.ConfiguredTestAttributes")
}

func init() {
	proto.RegisterFile("google/devtools/resultstore/v2/configured_target.proto", fileDescriptor_configured_target_dc7d1ae233b30620)
}

var fileDescriptor_configured_target_dc7d1ae233b30620 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xdf, 0x8a, 0x13, 0x31,
	0x14, 0xc6, 0xe9, 0x5f, 0xb7, 0x67, 0x5d, 0xdb, 0xcd, 0xd5, 0x58, 0x41, 0x4a, 0x95, 0xa5, 0xab,
	0x30, 0x23, 0x5d, 0x10, 0x51, 0x10, 0xb4, 0x8b, 0x38, 0x77, 0x4b, 0x76, 0xaf, 0x04, 0x19, 0xd2,
	0x26, 0x1d, 0x03, 0x33, 0x49, 0x37, 0x39, 0x29, 0xf8, 0x5e, 0x3e, 0x93, 0xcf, 0x21, 0x4d, 0x66,
	0x6a, 0x77, 0x61, 0x1d, 0xbc, 0x9b, 0x7c, 0xf3, 0x7d, 0xbf, 0x9c, 0xe4, 0x9c, 0xc0, 0xdb, 0x5c,
	0xeb, 0xbc, 0x10, 0x09, 0x17, 0x5b, 0xd4, 0xba, 0xb0, 0x89, 0x11, 0xd6, 0x15, 0x68, 0x51, 0x1b,
	0x91, 0x6c, 0xe7, 0xc9, 0x4a, 0xab, 0xb5, 0xcc, 0x9d, 0x11, 0x3c, 0x43, 0x66, 0x72, 0x81, 0xf1,
	0xc6, 0x68, 0xd4, 0xe4, 0x79, 0xc8, 0xc5, 0x75, 0x2e, 0x3e, 0xc8, 0xc5, 0xdb, 0xf9, 0xf8, 0x75,
	0x23, 0xb7, 0x2c, 0xb5, 0x0a, 0xb0, 0xf1, 0x79, 0x83, 0x79, 0x2d, 0x0b, 0x51, 0x59, 0xab, 0x7d,
	0x13, 0xbf, 0x5a, 0xba, 0x75, 0xc2, 0x9d, 0x61, 0x28, 0x6b, 0xd4, 0xf4, 0x77, 0x17, 0x46, 0x8b,
	0x7d, 0xcd, 0x37, 0xbe, 0x64, 0x42, 0xa0, 0xab, 0x58, 0x29, 0xa2, 0xd6, 0xa4, 0x35, 0x1b, 0x50,
	0xff, 0x4d, 0x16, 0xd0, 0x96, 0x3c, 0x6a, 0x4f, 0x5a, 0xb3, 0xe3, 0xf9, 0x45, 0xfc, 0xef, 0xd3,
	0xc4, 0xf7, 0x89, 0x71, 0xca, 0x69, 0x5b, 0x72, 0xf2, 0x1d, 0x4e, 0x2d, 0x32, 0x74, 0x36, 0x63,
	0x88, 0x46, 0x2e, 0x1d, 0x0a, 0x1b, 0x75, 0x3c, 0xf3, 0x4d, 0x13, 0xf3, 0xda, 0x07, 0x3f, 0xed,
	0x73, 0x74, 0x64, 0xef, 0x29, 0xe4, 0x23, 0xf4, 0x51, 0x96, 0x52, 0xe5, 0x51, 0xd7, 0x33, 0xcf,
	0x9a, 0x98, 0x37, 0xde, 0x4d, 0xab, 0x14, 0x61, 0x30, 0x44, 0x61, 0xf1, 0xb0, 0xb8, 0xbe, 0x07,
	0xbd, 0xfb, 0x8f, 0x03, 0x0b, 0x8b, 0x07, 0x45, 0x3e, 0xc1, 0x3b, 0x6b, 0xf2, 0x15, 0x60, 0x63,
	0xf4, 0x46, 0x18, 0x94, 0xc2, 0x46, 0x8f, 0x26, 0x9d, 0xd9, 0xf1, 0x7c, 0xd6, 0x44, 0xbf, 0x0a,
	0x89, 0x9f, 0xf4, 0x20, 0x4b, 0xde, 0x43, 0x6f, 0xd7, 0x67, 0x1b, 0x1d, 0x79, 0xc8, 0xcb, 0x26,
	0xc8, 0x17, 0x59, 0x08, 0x1a, 0x22, 0xe3, 0x5b, 0x68, 0xa7, 0x9c, 0xbc, 0x80, 0x13, 0xa9, 0xb6,
	0x7a, 0xe5, 0xe7, 0x21, 0x93, 0xbc, 0xea, 0xf7, 0xe3, 0xbf, 0x62, 0xca, 0xc9, 0x33, 0x18, 0x84,
	0x41, 0xce, 0xaa, 0xf6, 0x0f, 0xe8, 0x51, 0x10, 0x52, 0x4e, 0xce, 0x61, 0x54, 0x0f, 0xfc, 0x1e,
	0xd2, 0xf1, 0x9e, 0xe1, 0x1d, 0x3d, 0xe5, 0xd3, 0x5f, 0x2d, 0x88, 0x1e, 0xba, 0x25, 0x72, 0x06,
	0x43, 0xd4, 0xc8, 0x8a, 0xcc, 0x38, 0x95, 0xad, 0xb4, 0x53, 0xe8, 0xb7, 0xea, 0xd1, 0x13, 0x2f,
	0x53, 0xa7, 0x16, 0x3b, 0x91, 0xbc, 0x82, 0xd3, 0xe0, 0xb3, 0x3f, 0x98, 0xe1, 0x95, 0xb3, 0xe3,
	0x9d, 0x01, 0x70, 0xbd, 0xd3, 0x83, 0xf7, 0x12, 0x46, 0x28, 0x4b, 0xa1, 0x1d, 0x66, 0xf5, 0xcc,
	0x47, 0x3d, 0xdf, 0xcd, 0xa7, 0xf5, 0x55, 0xd5, 0x8f, 0x22, 0xbe, 0xac, 0x0c, 0x74, 0x58, 0x45,
	0x6a, 0xe1, 0xf3, 0x2d, 0x4c, 0x57, 0xba, 0x6c, 0xb8, 0xdb, 0xab, 0xd6, 0xb7, 0xb4, 0x72, 0xe4,
	0xba, 0x60, 0x2a, 0x8f, 0xb5, 0xc9, 0x93, 0x5c, 0x28, 0xbf, 0x41, 0x12, 0x7e, 0xb1, 0x8d, 0xb4,
	0x0f, 0xbd, 0xd8, 0x0f, 0x07, 0xcb, 0x65, 0xdf, 0xa7, 0x2e, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xe2, 0xc4, 0xd5, 0x45, 0x6b, 0x04, 0x00, 0x00,
}
