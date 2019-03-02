// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/list_operation_error.proto

package errors // import "github.com/Beeketing/go-genproto/googleapis/ads/googleads/v0/errors"

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

// Enum describing possible list operation errors.
type ListOperationErrorEnum_ListOperationError int32

const (
	// Enum unspecified.
	ListOperationErrorEnum_UNSPECIFIED ListOperationErrorEnum_ListOperationError = 0
	// The received error code is not known in this version.
	ListOperationErrorEnum_UNKNOWN ListOperationErrorEnum_ListOperationError = 1
	// Field required in value is missing.
	ListOperationErrorEnum_REQUIRED_FIELD_MISSING ListOperationErrorEnum_ListOperationError = 7
	// Duplicate or identical value is sent in multiple list operations.
	ListOperationErrorEnum_DUPLICATE_VALUES ListOperationErrorEnum_ListOperationError = 8
)

var ListOperationErrorEnum_ListOperationError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	7: "REQUIRED_FIELD_MISSING",
	8: "DUPLICATE_VALUES",
}
var ListOperationErrorEnum_ListOperationError_value = map[string]int32{
	"UNSPECIFIED":            0,
	"UNKNOWN":                1,
	"REQUIRED_FIELD_MISSING": 7,
	"DUPLICATE_VALUES":       8,
}

func (x ListOperationErrorEnum_ListOperationError) String() string {
	return proto.EnumName(ListOperationErrorEnum_ListOperationError_name, int32(x))
}
func (ListOperationErrorEnum_ListOperationError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_list_operation_error_c336a4dd73d7c486, []int{0, 0}
}

// Container for enum describing possible list operation errors.
type ListOperationErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOperationErrorEnum) Reset()         { *m = ListOperationErrorEnum{} }
func (m *ListOperationErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ListOperationErrorEnum) ProtoMessage()    {}
func (*ListOperationErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_operation_error_c336a4dd73d7c486, []int{0}
}
func (m *ListOperationErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOperationErrorEnum.Unmarshal(m, b)
}
func (m *ListOperationErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOperationErrorEnum.Marshal(b, m, deterministic)
}
func (dst *ListOperationErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOperationErrorEnum.Merge(dst, src)
}
func (m *ListOperationErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ListOperationErrorEnum.Size(m)
}
func (m *ListOperationErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOperationErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ListOperationErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListOperationErrorEnum)(nil), "google.ads.googleads.v0.errors.ListOperationErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.ListOperationErrorEnum_ListOperationError", ListOperationErrorEnum_ListOperationError_name, ListOperationErrorEnum_ListOperationError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/list_operation_error.proto", fileDescriptor_list_operation_error_c336a4dd73d7c486)
}

var fileDescriptor_list_operation_error_c336a4dd73d7c486 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0xdd, 0x04, 0x27, 0xd9, 0xc1, 0x12, 0x64, 0x82, 0x87, 0x1d, 0xfa, 0x00, 0x69, 0xc1,
	0x93, 0xf1, 0x94, 0xad, 0xd9, 0x08, 0xd6, 0xae, 0xae, 0xb6, 0x82, 0x14, 0x4a, 0x35, 0x25, 0x14,
	0xb6, 0x66, 0x24, 0x75, 0x47, 0x1f, 0xc6, 0xa3, 0x8f, 0xe2, 0xa3, 0x78, 0xf4, 0x09, 0xa4, 0x8d,
	0xed, 0x65, 0xe8, 0x29, 0x1f, 0x5f, 0x7e, 0x5f, 0xf2, 0xfd, 0xff, 0xe0, 0x5a, 0x48, 0x29, 0x36,
	0x85, 0x93, 0x73, 0xed, 0x18, 0xd9, 0xa8, 0xbd, 0xeb, 0x14, 0x4a, 0x49, 0xa5, 0x9d, 0x4d, 0xa9,
	0xeb, 0x4c, 0xee, 0x0a, 0x95, 0xd7, 0xa5, 0xac, 0xb2, 0xd6, 0x45, 0x3b, 0x25, 0x6b, 0x09, 0xa7,
	0x86, 0x47, 0x39, 0xd7, 0xa8, 0x8f, 0xa2, 0xbd, 0x8b, 0x4c, 0xd4, 0x7e, 0x03, 0x13, 0xbf, 0xd4,
	0xf5, 0xaa, 0x0b, 0xd3, 0xc6, 0xa6, 0xd5, 0xeb, 0xd6, 0xe6, 0x00, 0x1e, 0xde, 0xc0, 0x33, 0x30,
	0x8e, 0x83, 0x28, 0xa4, 0x73, 0xb6, 0x60, 0xd4, 0xb3, 0x8e, 0xe0, 0x18, 0x8c, 0xe2, 0xe0, 0x36,
	0x58, 0x3d, 0x06, 0xd6, 0x00, 0x5e, 0x82, 0xc9, 0x9a, 0xde, 0xc7, 0x6c, 0x4d, 0xbd, 0x6c, 0xc1,
	0xa8, 0xef, 0x65, 0x77, 0x2c, 0x8a, 0x58, 0xb0, 0xb4, 0x46, 0xf0, 0x1c, 0x58, 0x5e, 0x1c, 0xfa,
	0x6c, 0x4e, 0x1e, 0x68, 0x96, 0x10, 0x3f, 0xa6, 0x91, 0x75, 0x3a, 0xfb, 0x1e, 0x00, 0xfb, 0x45,
	0x6e, 0xd1, 0xff, 0x35, 0x67, 0x17, 0x87, 0x55, 0xc2, 0x66, 0xbe, 0x70, 0xf0, 0xe4, 0xfd, 0x46,
	0x85, 0xdc, 0xe4, 0x95, 0x40, 0x52, 0x09, 0x47, 0x14, 0x55, 0x3b, 0x7d, 0xb7, 0xac, 0x5d, 0xa9,
	0xff, 0xda, 0xdd, 0x8d, 0x39, 0xde, 0x87, 0xc7, 0x4b, 0x42, 0x3e, 0x86, 0xd3, 0xa5, 0x79, 0x8c,
	0x70, 0x8d, 0x8c, 0x6c, 0x54, 0xe2, 0xa2, 0xf6, 0x4b, 0xfd, 0xd9, 0x01, 0x29, 0xe1, 0x3a, 0xed,
	0x81, 0x34, 0x71, 0x53, 0x03, 0x7c, 0x0d, 0x6d, 0xe3, 0x62, 0x4c, 0xb8, 0xc6, 0xb8, 0x47, 0x30,
	0x4e, 0x5c, 0x8c, 0x0d, 0xf4, 0x7c, 0xd2, 0xb6, 0xbb, 0xfa, 0x09, 0x00, 0x00, 0xff, 0xff, 0xf4,
	0xf2, 0x70, 0x8f, 0xd8, 0x01, 0x00, 0x00,
}
