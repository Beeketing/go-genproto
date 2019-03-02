// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/build/v1/build_status.proto

package build // import "github.com/Beeketing/genproto/googleapis/devtools/build/v1"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/Beeketing/protobuf/ptypes/any"
import _ "github.com/Beeketing/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The end result of the Build.
type BuildStatus_Result int32

const (
	// Unspecified or unknown.
	BuildStatus_UNKNOWN_STATUS BuildStatus_Result = 0
	// Build was successful and tests (if requested) all pass.
	BuildStatus_COMMAND_SUCCEEDED BuildStatus_Result = 1
	// Build error and/or test failure.
	BuildStatus_COMMAND_FAILED BuildStatus_Result = 2
	// Unable to obtain a result due to input provided by the user.
	BuildStatus_USER_ERROR BuildStatus_Result = 3
	// Unable to obtain a result due to a failure within the build system.
	BuildStatus_SYSTEM_ERROR BuildStatus_Result = 4
	// Build required too many resources, such as build tool RAM.
	BuildStatus_RESOURCE_EXHAUSTED BuildStatus_Result = 5
	// An invocation attempt time exceeded its deadline.
	BuildStatus_INVOCATION_DEADLINE_EXCEEDED BuildStatus_Result = 6
	// Build request time exceeded the request_deadline
	BuildStatus_REQUEST_DEADLINE_EXCEEDED BuildStatus_Result = 8
	// The build was cancelled by a call to CancelBuild.
	BuildStatus_CANCELLED BuildStatus_Result = 7
)

var BuildStatus_Result_name = map[int32]string{
	0: "UNKNOWN_STATUS",
	1: "COMMAND_SUCCEEDED",
	2: "COMMAND_FAILED",
	3: "USER_ERROR",
	4: "SYSTEM_ERROR",
	5: "RESOURCE_EXHAUSTED",
	6: "INVOCATION_DEADLINE_EXCEEDED",
	8: "REQUEST_DEADLINE_EXCEEDED",
	7: "CANCELLED",
}
var BuildStatus_Result_value = map[string]int32{
	"UNKNOWN_STATUS":               0,
	"COMMAND_SUCCEEDED":            1,
	"COMMAND_FAILED":               2,
	"USER_ERROR":                   3,
	"SYSTEM_ERROR":                 4,
	"RESOURCE_EXHAUSTED":           5,
	"INVOCATION_DEADLINE_EXCEEDED": 6,
	"REQUEST_DEADLINE_EXCEEDED":    8,
	"CANCELLED":                    7,
}

func (x BuildStatus_Result) String() string {
	return proto.EnumName(BuildStatus_Result_name, int32(x))
}
func (BuildStatus_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_build_status_ac8c345d98bfecf7, []int{0, 0}
}

// Status used for both invocation attempt and overall build completion.
type BuildStatus struct {
	// The end result.
	Result BuildStatus_Result `protobuf:"varint,1,opt,name=result,proto3,enum=google.devtools.build.v1.BuildStatus_Result" json:"result,omitempty"`
	// Fine-grained diagnostic information to complement the status.
	Details              *any.Any `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildStatus) Reset()         { *m = BuildStatus{} }
func (m *BuildStatus) String() string { return proto.CompactTextString(m) }
func (*BuildStatus) ProtoMessage()    {}
func (*BuildStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_build_status_ac8c345d98bfecf7, []int{0}
}
func (m *BuildStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildStatus.Unmarshal(m, b)
}
func (m *BuildStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildStatus.Marshal(b, m, deterministic)
}
func (dst *BuildStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildStatus.Merge(dst, src)
}
func (m *BuildStatus) XXX_Size() int {
	return xxx_messageInfo_BuildStatus.Size(m)
}
func (m *BuildStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildStatus.DiscardUnknown(m)
}

var xxx_messageInfo_BuildStatus proto.InternalMessageInfo

func (m *BuildStatus) GetResult() BuildStatus_Result {
	if m != nil {
		return m.Result
	}
	return BuildStatus_UNKNOWN_STATUS
}

func (m *BuildStatus) GetDetails() *any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*BuildStatus)(nil), "google.devtools.build.v1.BuildStatus")
	proto.RegisterEnum("google.devtools.build.v1.BuildStatus_Result", BuildStatus_Result_name, BuildStatus_Result_value)
}

func init() {
	proto.RegisterFile("google/devtools/build/v1/build_status.proto", fileDescriptor_build_status_ac8c345d98bfecf7)
}

var fileDescriptor_build_status_ac8c345d98bfecf7 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x0b, 0xd3, 0x30,
	0x18, 0xc6, 0xcd, 0xd4, 0x4e, 0x33, 0x1d, 0x35, 0xa8, 0x6c, 0x63, 0xc2, 0xd8, 0x69, 0xa0, 0xa4,
	0x6c, 0x1e, 0xc5, 0x43, 0xd6, 0x44, 0x2c, 0x6e, 0xe9, 0x4c, 0x5a, 0xff, 0x5d, 0x4a, 0xe6, 0x6a,
	0x29, 0xd4, 0x66, 0xac, 0xe9, 0x60, 0x1f, 0xd3, 0x93, 0x5f, 0xc5, 0xa3, 0xf4, 0x1f, 0x0c, 0x74,
	0xb7, 0xf4, 0x7d, 0x7e, 0xcf, 0xf3, 0xbe, 0x3c, 0x14, 0xbe, 0x4c, 0xb4, 0x4e, 0xb2, 0xd8, 0x39,
	0xc4, 0x67, 0xa3, 0x75, 0x56, 0x38, 0xfb, 0x32, 0xcd, 0x0e, 0xce, 0x79, 0xd9, 0x3c, 0xa2, 0xc2,
	0x28, 0x53, 0x16, 0xf8, 0x78, 0xd2, 0x46, 0xa3, 0x51, 0x03, 0xe3, 0x0e, 0xc6, 0x35, 0x83, 0xcf,
	0xcb, 0xc9, 0xb4, 0x8d, 0x51, 0xc7, 0xd4, 0x51, 0x79, 0xae, 0x8d, 0x32, 0xa9, 0xce, 0x5b, 0xdf,
	0x64, 0xdc, 0xaa, 0xf5, 0xd7, 0xbe, 0xfc, 0xe1, 0xa8, 0xfc, 0xd2, 0x48, 0xf3, 0xdf, 0x3d, 0x38,
	0x58, 0x57, 0x29, 0xb2, 0x5e, 0x84, 0x28, 0xb4, 0x4e, 0x71, 0x51, 0x66, 0x66, 0x04, 0x66, 0x60,
	0x31, 0x5c, 0xbd, 0xc2, 0xb7, 0x76, 0xe2, 0x2b, 0x1b, 0x16, 0xb5, 0x47, 0xb4, 0x5e, 0x84, 0x61,
	0xff, 0x10, 0x1b, 0x95, 0x66, 0xc5, 0xa8, 0x37, 0x03, 0x8b, 0xc1, 0xea, 0x69, 0x17, 0xd3, 0x9d,
	0x80, 0x49, 0x7e, 0x11, 0x1d, 0x34, 0xff, 0x05, 0xa0, 0xd5, 0x44, 0x20, 0x04, 0x87, 0x21, 0xff,
	0xc0, 0xfd, 0xcf, 0x3c, 0x92, 0x01, 0x09, 0x42, 0x69, 0xdf, 0x41, 0xcf, 0xe0, 0x13, 0xd7, 0xdf,
	0x6e, 0x09, 0xa7, 0x91, 0x0c, 0x5d, 0x97, 0x31, 0xca, 0xa8, 0x0d, 0x2a, 0xb4, 0x1b, 0xbf, 0x23,
	0xde, 0x86, 0x51, 0xbb, 0x87, 0x86, 0x10, 0x86, 0x92, 0x89, 0x88, 0x09, 0xe1, 0x0b, 0xfb, 0x2e,
	0xb2, 0xe1, 0x23, 0xf9, 0x55, 0x06, 0x6c, 0xdb, 0x4e, 0xee, 0xa1, 0xe7, 0x10, 0x09, 0x26, 0xfd,
	0x50, 0xb8, 0x2c, 0x62, 0x5f, 0xde, 0x93, 0x50, 0x06, 0x8c, 0xda, 0xf7, 0xd1, 0x0c, 0x4e, 0x3d,
	0xfe, 0xc9, 0x77, 0x49, 0xe0, 0xf9, 0x3c, 0xa2, 0x8c, 0xd0, 0x8d, 0xc7, 0x2b, 0xa4, 0xdd, 0x67,
	0xa1, 0x17, 0x70, 0x2c, 0xd8, 0xc7, 0x90, 0xc9, 0xe0, 0x3f, 0xf2, 0x03, 0xf4, 0x18, 0x3e, 0x74,
	0x09, 0x77, 0xd9, 0xa6, 0xba, 0xa4, 0xbf, 0x36, 0x70, 0xfa, 0x5d, 0xff, 0xbc, 0x59, 0xdf, 0xda,
	0xbe, 0xea, 0x6f, 0x57, 0xb5, 0xb2, 0x03, 0xdf, 0xde, 0xb6, 0x74, 0xa2, 0x33, 0x95, 0x27, 0x58,
	0x9f, 0x12, 0x27, 0x89, 0xf3, 0xba, 0x33, 0xa7, 0x91, 0xd4, 0x31, 0x2d, 0xfe, 0xfd, 0x59, 0xde,
	0xd4, 0x8f, 0x3f, 0x00, 0xec, 0xad, 0x1a, 0x7e, 0xfd, 0x37, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x3d,
	0xf5, 0x87, 0x58, 0x02, 0x00, 0x00,
}
