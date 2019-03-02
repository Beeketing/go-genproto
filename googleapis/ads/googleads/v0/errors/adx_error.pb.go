// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/adx_error.proto

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

// Enum describing possible adx errors.
type AdxErrorEnum_AdxError int32

const (
	// Enum unspecified.
	AdxErrorEnum_UNSPECIFIED AdxErrorEnum_AdxError = 0
	// The received error code is not known in this version.
	AdxErrorEnum_UNKNOWN AdxErrorEnum_AdxError = 1
	// Attempt to use non-AdX feature by AdX customer.
	AdxErrorEnum_UNSUPPORTED_FEATURE AdxErrorEnum_AdxError = 2
)

var AdxErrorEnum_AdxError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "UNSUPPORTED_FEATURE",
}
var AdxErrorEnum_AdxError_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"UNSUPPORTED_FEATURE": 2,
}

func (x AdxErrorEnum_AdxError) String() string {
	return proto.EnumName(AdxErrorEnum_AdxError_name, int32(x))
}
func (AdxErrorEnum_AdxError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_adx_error_d89de8e4329878ad, []int{0, 0}
}

// Container for enum describing possible adx errors.
type AdxErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdxErrorEnum) Reset()         { *m = AdxErrorEnum{} }
func (m *AdxErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AdxErrorEnum) ProtoMessage()    {}
func (*AdxErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_adx_error_d89de8e4329878ad, []int{0}
}
func (m *AdxErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdxErrorEnum.Unmarshal(m, b)
}
func (m *AdxErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdxErrorEnum.Marshal(b, m, deterministic)
}
func (dst *AdxErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdxErrorEnum.Merge(dst, src)
}
func (m *AdxErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AdxErrorEnum.Size(m)
}
func (m *AdxErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdxErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdxErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdxErrorEnum)(nil), "google.ads.googleads.v0.errors.AdxErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.AdxErrorEnum_AdxError", AdxErrorEnum_AdxError_name, AdxErrorEnum_AdxError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/adx_error.proto", fileDescriptor_adx_error_d89de8e4329878ad)
}

var fileDescriptor_adx_error_d89de8e4329878ad = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x6d, 0x05, 0x95, 0x4c, 0xb1, 0xd4, 0x83, 0xb7, 0x1d, 0xfa, 0x00, 0x5f, 0x03, 0xde,
	0xe2, 0x29, 0xb3, 0xd9, 0x18, 0x42, 0x16, 0xb7, 0xa5, 0x82, 0x14, 0x46, 0x35, 0x25, 0x08, 0x5b,
	0x33, 0x12, 0x1d, 0x7b, 0x1e, 0x8f, 0x3e, 0x8a, 0x4f, 0x22, 0x3e, 0x85, 0xb4, 0xb1, 0xbd, 0xb9,
	0x53, 0x7e, 0x5f, 0xf8, 0x25, 0xff, 0xef, 0x8f, 0x40, 0x1b, 0xa3, 0xd7, 0x55, 0x5a, 0x2a, 0x97,
	0x7a, 0x6c, 0x68, 0x87, 0xd3, 0xca, 0x5a, 0x63, 0x5d, 0x5a, 0xaa, 0xfd, 0xaa, 0x45, 0xd8, 0x5a,
	0xf3, 0x66, 0xe2, 0xa1, 0x97, 0xa0, 0x54, 0x0e, 0x7a, 0x1f, 0x76, 0x18, 0xbc, 0x9f, 0x3c, 0xa0,
	0x73, 0xaa, 0xf6, 0xac, 0x19, 0x58, 0xfd, 0xbe, 0x49, 0x28, 0x3a, 0xeb, 0xe6, 0xf8, 0x12, 0x0d,
	0x24, 0x5f, 0x08, 0x76, 0x37, 0x1d, 0x4f, 0x59, 0x16, 0x1d, 0xc5, 0x03, 0x74, 0x2a, 0xf9, 0x3d,
	0x9f, 0x3d, 0xf2, 0x28, 0x88, 0xaf, 0xd1, 0x95, 0xe4, 0x0b, 0x29, 0xc4, 0x6c, 0xbe, 0x64, 0xd9,
	0x6a, 0xcc, 0xe8, 0x52, 0xce, 0x59, 0x14, 0x8e, 0xbe, 0x03, 0x94, 0xbc, 0x98, 0x0d, 0x1c, 0x4e,
	0x1e, 0x5d, 0x74, 0x39, 0xa2, 0x59, 0x54, 0x04, 0x4f, 0xd9, 0xdf, 0x03, 0x6d, 0xd6, 0x65, 0xad,
	0xc1, 0x58, 0x9d, 0xea, 0xaa, 0x6e, 0x6b, 0x74, 0x55, 0xb7, 0xaf, 0xee, 0xbf, 0xe6, 0xb7, 0xfe,
	0xf8, 0x08, 0x8f, 0x27, 0x94, 0x7e, 0x86, 0xc3, 0x89, 0xff, 0x8c, 0x2a, 0x07, 0x1e, 0x1b, 0xca,
	0x31, 0xb4, 0x91, 0xee, 0xab, 0x13, 0x0a, 0xaa, 0x5c, 0xd1, 0x0b, 0x45, 0x8e, 0x0b, 0x2f, 0xfc,
	0x84, 0x89, 0xbf, 0x25, 0x84, 0x2a, 0x47, 0x48, 0xaf, 0x10, 0x92, 0x63, 0x42, 0xbc, 0xf4, 0x7c,
	0xd2, 0x6e, 0x77, 0xf3, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xce, 0xd3, 0xa2, 0xb3, 0x96, 0x01, 0x00,
	0x00,
}
