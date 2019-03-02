// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/call_conversion_reporting_state.proto

package enums // import "github.com/Beeketing/go-genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Possible data types for a call conversion action state.
type CallConversionReportingStateEnum_CallConversionReportingState int32

const (
	// Not specified.
	CallConversionReportingStateEnum_UNSPECIFIED CallConversionReportingStateEnum_CallConversionReportingState = 0
	// Used for return value only. Represents value unknown in this version.
	CallConversionReportingStateEnum_UNKNOWN CallConversionReportingStateEnum_CallConversionReportingState = 1
	// Call conversion action is disabled.
	CallConversionReportingStateEnum_DISABLED CallConversionReportingStateEnum_CallConversionReportingState = 2
	// Call conversion action will use call conversion type set at the
	// account level.
	CallConversionReportingStateEnum_USE_ACCOUNT_LEVEL_CALL_CONVERSION_ACTION CallConversionReportingStateEnum_CallConversionReportingState = 3
	// Call conversion action will use call conversion type set at the resource
	// (call only ads/call extensions) level.
	CallConversionReportingStateEnum_USE_RESOURCE_LEVEL_CALL_CONVERSION_ACTION CallConversionReportingStateEnum_CallConversionReportingState = 4
)

var CallConversionReportingStateEnum_CallConversionReportingState_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "DISABLED",
	3: "USE_ACCOUNT_LEVEL_CALL_CONVERSION_ACTION",
	4: "USE_RESOURCE_LEVEL_CALL_CONVERSION_ACTION",
}
var CallConversionReportingStateEnum_CallConversionReportingState_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"DISABLED":    2,
	"USE_ACCOUNT_LEVEL_CALL_CONVERSION_ACTION":  3,
	"USE_RESOURCE_LEVEL_CALL_CONVERSION_ACTION": 4,
}

func (x CallConversionReportingStateEnum_CallConversionReportingState) String() string {
	return proto.EnumName(CallConversionReportingStateEnum_CallConversionReportingState_name, int32(x))
}
func (CallConversionReportingStateEnum_CallConversionReportingState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_call_conversion_reporting_state_5e7c8cb103fba176, []int{0, 0}
}

// Container for enum describing possible data types for call conversion
// reporting state.
type CallConversionReportingStateEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallConversionReportingStateEnum) Reset()         { *m = CallConversionReportingStateEnum{} }
func (m *CallConversionReportingStateEnum) String() string { return proto.CompactTextString(m) }
func (*CallConversionReportingStateEnum) ProtoMessage()    {}
func (*CallConversionReportingStateEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_call_conversion_reporting_state_5e7c8cb103fba176, []int{0}
}
func (m *CallConversionReportingStateEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallConversionReportingStateEnum.Unmarshal(m, b)
}
func (m *CallConversionReportingStateEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallConversionReportingStateEnum.Marshal(b, m, deterministic)
}
func (dst *CallConversionReportingStateEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallConversionReportingStateEnum.Merge(dst, src)
}
func (m *CallConversionReportingStateEnum) XXX_Size() int {
	return xxx_messageInfo_CallConversionReportingStateEnum.Size(m)
}
func (m *CallConversionReportingStateEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CallConversionReportingStateEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CallConversionReportingStateEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CallConversionReportingStateEnum)(nil), "google.ads.googleads.v0.enums.CallConversionReportingStateEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.CallConversionReportingStateEnum_CallConversionReportingState", CallConversionReportingStateEnum_CallConversionReportingState_name, CallConversionReportingStateEnum_CallConversionReportingState_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/call_conversion_reporting_state.proto", fileDescriptor_call_conversion_reporting_state_5e7c8cb103fba176)
}

var fileDescriptor_call_conversion_reporting_state_5e7c8cb103fba176 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x6b, 0xb3, 0x30,
	0x1c, 0xc6, 0x5f, 0xed, 0xcb, 0xfb, 0x8e, 0x74, 0x30, 0xf1, 0xbc, 0xc2, 0xda, 0xd3, 0x06, 0x5b,
	0x14, 0x76, 0xcb, 0x4e, 0x31, 0xcd, 0x8a, 0x4c, 0x62, 0xd1, 0xea, 0x60, 0x08, 0xc1, 0x55, 0x91,
	0x82, 0x4d, 0x8a, 0xb1, 0xfd, 0x40, 0xbb, 0x0c, 0xf6, 0x51, 0xc6, 0x3e, 0xc9, 0x3e, 0xc0, 0xce,
	0x43, 0x5d, 0xbd, 0xcd, 0x4b, 0x78, 0xc8, 0xf3, 0xe4, 0x47, 0xfe, 0xcf, 0x1f, 0x90, 0x42, 0xca,
	0xa2, 0xcc, 0xad, 0x34, 0x53, 0x56, 0x27, 0x1b, 0x75, 0xb0, 0xad, 0x5c, 0xec, 0xb7, 0xca, 0x5a,
	0xa7, 0x65, 0xc9, 0xd7, 0x52, 0x1c, 0xf2, 0x4a, 0x6d, 0xa4, 0xe0, 0x55, 0xbe, 0x93, 0x55, 0xbd,
	0x11, 0x05, 0x57, 0x75, 0x5a, 0xe7, 0x70, 0x57, 0xc9, 0x5a, 0x9a, 0x93, 0xee, 0x25, 0x4c, 0x33,
	0x05, 0x7b, 0x08, 0x3c, 0xd8, 0xb0, 0x85, 0xcc, 0x3e, 0x34, 0x70, 0x41, 0xd2, 0xb2, 0x24, 0x3d,
	0x27, 0x38, 0x62, 0xc2, 0x86, 0x42, 0xc5, 0x7e, 0x3b, 0x7b, 0xd5, 0xc0, 0xf9, 0x50, 0xc8, 0x3c,
	0x03, 0xe3, 0x88, 0x85, 0x4b, 0x4a, 0xdc, 0x7b, 0x97, 0xce, 0x8d, 0x3f, 0xe6, 0x18, 0xfc, 0x8f,
	0xd8, 0x03, 0xf3, 0x1f, 0x99, 0xa1, 0x99, 0xa7, 0xe0, 0x64, 0xee, 0x86, 0xd8, 0xf1, 0xe8, 0xdc,
	0xd0, 0xcd, 0x6b, 0x70, 0x19, 0x85, 0x94, 0x63, 0x42, 0xfc, 0x88, 0xad, 0xb8, 0x47, 0x63, 0xea,
	0x71, 0x82, 0x3d, 0x8f, 0x13, 0x9f, 0xc5, 0x34, 0x08, 0x5d, 0x9f, 0x71, 0x4c, 0x56, 0xae, 0xcf,
	0x8c, 0x91, 0x79, 0x03, 0xae, 0x9a, 0x74, 0x40, 0x43, 0x3f, 0x0a, 0x08, 0x1d, 0x8e, 0xff, 0x75,
	0xbe, 0x34, 0x30, 0x5d, 0xcb, 0x2d, 0x1c, 0x1c, 0xda, 0x99, 0x0e, 0x0d, 0xb3, 0x6c, 0x6a, 0x5b,
	0x6a, 0x4f, 0xce, 0x0f, 0xa3, 0x90, 0x65, 0x2a, 0x0a, 0x28, 0xab, 0xc2, 0x2a, 0x72, 0xd1, 0x96,
	0x7a, 0xdc, 0xc6, 0x6e, 0xa3, 0x7e, 0x59, 0xce, 0x5d, 0x7b, 0xbe, 0xe8, 0xa3, 0x05, 0xc6, 0x6f,
	0xfa, 0x64, 0xd1, 0xa1, 0x70, 0xa6, 0x60, 0x27, 0x1b, 0x15, 0xdb, 0xb0, 0x69, 0x57, 0xbd, 0x1f,
	0xfd, 0x04, 0x67, 0x2a, 0xe9, 0xfd, 0x24, 0xb6, 0x93, 0xd6, 0xff, 0xd4, 0xa7, 0xdd, 0x25, 0x42,
	0x38, 0x53, 0x08, 0xf5, 0x09, 0x84, 0x62, 0x1b, 0xa1, 0x36, 0xf3, 0xfc, 0xaf, 0xfd, 0xd8, 0xed,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x50, 0x89, 0xd5, 0x06, 0x34, 0x02, 0x00, 0x00,
}
