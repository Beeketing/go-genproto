// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/null_error.proto

package errors // import "github.com/Beeketing/genproto/googleapis/ads/googleads/v0/errors"

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

// Enum describing possible null errors.
type NullErrorEnum_NullError int32

const (
	// Enum unspecified.
	NullErrorEnum_UNSPECIFIED NullErrorEnum_NullError = 0
	// The received error code is not known in this version.
	NullErrorEnum_UNKNOWN NullErrorEnum_NullError = 1
	// Specified list/container must not contain any null elements
	NullErrorEnum_NULL_CONTENT NullErrorEnum_NullError = 2
)

var NullErrorEnum_NullError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "NULL_CONTENT",
}
var NullErrorEnum_NullError_value = map[string]int32{
	"UNSPECIFIED":  0,
	"UNKNOWN":      1,
	"NULL_CONTENT": 2,
}

func (x NullErrorEnum_NullError) String() string {
	return proto.EnumName(NullErrorEnum_NullError_name, int32(x))
}
func (NullErrorEnum_NullError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_null_error_573ff81ec6fd9a9e, []int{0, 0}
}

// Container for enum describing possible null errors.
type NullErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NullErrorEnum) Reset()         { *m = NullErrorEnum{} }
func (m *NullErrorEnum) String() string { return proto.CompactTextString(m) }
func (*NullErrorEnum) ProtoMessage()    {}
func (*NullErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_null_error_573ff81ec6fd9a9e, []int{0}
}
func (m *NullErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NullErrorEnum.Unmarshal(m, b)
}
func (m *NullErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NullErrorEnum.Marshal(b, m, deterministic)
}
func (dst *NullErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NullErrorEnum.Merge(dst, src)
}
func (m *NullErrorEnum) XXX_Size() int {
	return xxx_messageInfo_NullErrorEnum.Size(m)
}
func (m *NullErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_NullErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_NullErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NullErrorEnum)(nil), "google.ads.googleads.v0.errors.NullErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.NullErrorEnum_NullError", NullErrorEnum_NullError_name, NullErrorEnum_NullError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/null_error.proto", fileDescriptor_null_error_573ff81ec6fd9a9e)
}

var fileDescriptor_null_error_573ff81ec6fd9a9e = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4f, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0x86, 0x32, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xa2, 0xa2,
	0xfc, 0xa2, 0x62, 0xfd, 0xbc, 0xd2, 0x9c, 0x9c, 0x78, 0x30, 0x5b, 0xaf, 0xa0, 0x28, 0xbf, 0x24,
	0x5f, 0x48, 0x0e, 0xa2, 0x4a, 0x2f, 0x31, 0xa5, 0x58, 0x0f, 0xae, 0x41, 0xaf, 0xcc, 0x40, 0x0f,
	0xa2, 0x41, 0xc9, 0x87, 0x8b, 0xd7, 0xaf, 0x34, 0x27, 0xc7, 0x15, 0xc4, 0x73, 0xcd, 0x2b, 0xcd,
	0x55, 0xb2, 0xe6, 0xe2, 0x84, 0x0b, 0x08, 0xf1, 0x73, 0x71, 0x87, 0xfa, 0x05, 0x07, 0xb8, 0x3a,
	0x7b, 0xba, 0x79, 0xba, 0xba, 0x08, 0x30, 0x08, 0x71, 0x73, 0xb1, 0x87, 0xfa, 0x79, 0xfb, 0xf9,
	0x87, 0xfb, 0x09, 0x30, 0x0a, 0x09, 0x70, 0xf1, 0xf8, 0x85, 0xfa, 0xf8, 0xc4, 0x3b, 0xfb, 0xfb,
	0x85, 0xb8, 0xfa, 0x85, 0x08, 0x30, 0x39, 0xbd, 0x64, 0xe4, 0x52, 0x4a, 0xce, 0xcf, 0xd5, 0xc3,
	0x6f, 0xa9, 0x13, 0x1f, 0xdc, 0x86, 0x00, 0x90, 0x23, 0x03, 0x18, 0xa3, 0x5c, 0xa0, 0x3a, 0xd2,
	0xf3, 0x73, 0x12, 0xf3, 0xd2, 0xf5, 0xf2, 0x8b, 0xd2, 0xf5, 0xd3, 0x53, 0xf3, 0xc0, 0x5e, 0x80,
	0xf9, 0xb3, 0x20, 0xb3, 0x18, 0x97, 0xb7, 0xad, 0x21, 0xd4, 0x22, 0x26, 0x66, 0x77, 0x47, 0xc7,
	0x55, 0x4c, 0x72, 0xee, 0x10, 0xc3, 0x1c, 0x53, 0x8a, 0xf5, 0x20, 0x4c, 0x10, 0x2b, 0xcc, 0x40,
	0x0f, 0x6c, 0x65, 0xf1, 0x29, 0x98, 0x82, 0x18, 0xc7, 0x94, 0xe2, 0x18, 0xb8, 0x82, 0x98, 0x30,
	0x83, 0x18, 0x88, 0x82, 0x57, 0x4c, 0x4a, 0x10, 0x51, 0x2b, 0x2b, 0xc7, 0x94, 0x62, 0x2b, 0x2b,
	0xb8, 0x12, 0x2b, 0xab, 0x30, 0x03, 0x2b, 0x2b, 0x88, 0xa2, 0x24, 0x36, 0xb0, 0xeb, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x87, 0xae, 0x3d, 0x93, 0x01, 0x00, 0x00,
}
