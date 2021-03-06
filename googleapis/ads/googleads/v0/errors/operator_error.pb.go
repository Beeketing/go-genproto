// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/operator_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

import proto "github.com/golang/protobuf/proto"
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

// Enum describing possible operator errors.
type OperatorErrorEnum_OperatorError int32

const (
	// Enum unspecified.
	OperatorErrorEnum_UNSPECIFIED OperatorErrorEnum_OperatorError = 0
	// The received error code is not known in this version.
	OperatorErrorEnum_UNKNOWN OperatorErrorEnum_OperatorError = 1
	// Operator not supported.
	OperatorErrorEnum_OPERATOR_NOT_SUPPORTED OperatorErrorEnum_OperatorError = 2
)

var OperatorErrorEnum_OperatorError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "OPERATOR_NOT_SUPPORTED",
}
var OperatorErrorEnum_OperatorError_value = map[string]int32{
	"UNSPECIFIED":            0,
	"UNKNOWN":                1,
	"OPERATOR_NOT_SUPPORTED": 2,
}

func (x OperatorErrorEnum_OperatorError) String() string {
	return proto.EnumName(OperatorErrorEnum_OperatorError_name, int32(x))
}
func (OperatorErrorEnum_OperatorError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_operator_error_fb2f60e4e8472bb4, []int{0, 0}
}

// Container for enum describing possible operator errors.
type OperatorErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperatorErrorEnum) Reset()         { *m = OperatorErrorEnum{} }
func (m *OperatorErrorEnum) String() string { return proto.CompactTextString(m) }
func (*OperatorErrorEnum) ProtoMessage()    {}
func (*OperatorErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_operator_error_fb2f60e4e8472bb4, []int{0}
}
func (m *OperatorErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperatorErrorEnum.Unmarshal(m, b)
}
func (m *OperatorErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperatorErrorEnum.Marshal(b, m, deterministic)
}
func (dst *OperatorErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperatorErrorEnum.Merge(dst, src)
}
func (m *OperatorErrorEnum) XXX_Size() int {
	return xxx_messageInfo_OperatorErrorEnum.Size(m)
}
func (m *OperatorErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_OperatorErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_OperatorErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*OperatorErrorEnum)(nil), "google.ads.googleads.v0.errors.OperatorErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.OperatorErrorEnum_OperatorError", OperatorErrorEnum_OperatorError_name, OperatorErrorEnum_OperatorError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/operator_error.proto", fileDescriptor_operator_error_fb2f60e4e8472bb4)
}

var fileDescriptor_operator_error_fb2f60e4e8472bb4 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xa2,
	0xa2, 0xfc, 0xa2, 0x62, 0xfd, 0xfc, 0x82, 0xd4, 0xa2, 0xc4, 0x92, 0xfc, 0xa2, 0x78, 0x30, 0x5f,
	0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x48, 0x0e, 0xa2, 0x52, 0x2f, 0x31, 0xa5, 0x58, 0x0f, 0xae,
	0x49, 0xaf, 0xcc, 0x40, 0x0f, 0xa2, 0x49, 0x29, 0x8e, 0x4b, 0xd0, 0x1f, 0xaa, 0xcf, 0x15, 0x24,
	0xe2, 0x9a, 0x57, 0x9a, 0xab, 0xe4, 0xc9, 0xc5, 0x8b, 0x22, 0x28, 0xc4, 0xcf, 0xc5, 0x1d, 0xea,
	0x17, 0x1c, 0xe0, 0xea, 0xec, 0xe9, 0xe6, 0xe9, 0xea, 0x22, 0xc0, 0x20, 0xc4, 0xcd, 0xc5, 0x1e,
	0xea, 0xe7, 0xed, 0xe7, 0x1f, 0xee, 0x27, 0xc0, 0x28, 0x24, 0xc5, 0x25, 0xe6, 0x1f, 0xe0, 0x1a,
	0xe4, 0x18, 0xe2, 0x1f, 0x14, 0xef, 0xe7, 0x1f, 0x12, 0x1f, 0x1c, 0x1a, 0x10, 0xe0, 0x1f, 0x14,
	0xe2, 0xea, 0x22, 0xc0, 0xe4, 0xf4, 0x96, 0x91, 0x4b, 0x29, 0x39, 0x3f, 0x57, 0x0f, 0xbf, 0x33,
	0x9c, 0x84, 0x50, 0xec, 0x0b, 0x00, 0x39, 0x3d, 0x80, 0x31, 0xca, 0x05, 0xaa, 0x2b, 0x3d, 0x3f,
	0x27, 0x31, 0x2f, 0x5d, 0x2f, 0xbf, 0x28, 0x5d, 0x3f, 0x3d, 0x35, 0x0f, 0xec, 0x31, 0x58, 0x08,
	0x14, 0x64, 0x16, 0xe3, 0x0a, 0x10, 0x6b, 0x08, 0xb5, 0x88, 0x89, 0xd9, 0xdd, 0xd1, 0x71, 0x15,
	0x93, 0x9c, 0x3b, 0xc4, 0x30, 0xc7, 0x94, 0x62, 0x3d, 0x08, 0x13, 0xc4, 0x0a, 0x33, 0xd0, 0x03,
	0x5b, 0x59, 0x7c, 0x0a, 0xa6, 0x20, 0xc6, 0x31, 0xa5, 0x38, 0x06, 0xae, 0x20, 0x26, 0xcc, 0x20,
	0x06, 0xa2, 0xe0, 0x15, 0x93, 0x12, 0x44, 0xd4, 0xca, 0xca, 0x31, 0xa5, 0xd8, 0xca, 0x0a, 0xae,
	0xc4, 0xca, 0x2a, 0xcc, 0xc0, 0xca, 0x0a, 0xa2, 0x28, 0x89, 0x0d, 0xec, 0x3a, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xc6, 0x9c, 0xb4, 0xbf, 0xad, 0x01, 0x00, 0x00,
}
