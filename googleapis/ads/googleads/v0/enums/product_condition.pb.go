// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/product_condition.proto

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

// Enum describing the condition of a product offer.
type ProductConditionEnum_ProductCondition int32

const (
	// Not specified.
	ProductConditionEnum_UNSPECIFIED ProductConditionEnum_ProductCondition = 0
	// Used for return value only. Represents value unknown in this version.
	ProductConditionEnum_UNKNOWN ProductConditionEnum_ProductCondition = 1
	// The product condition is new.
	ProductConditionEnum_NEW ProductConditionEnum_ProductCondition = 3
	// The product condition is refurbished.
	ProductConditionEnum_REFURBISHED ProductConditionEnum_ProductCondition = 4
	// The product condition is used.
	ProductConditionEnum_USED ProductConditionEnum_ProductCondition = 5
)

var ProductConditionEnum_ProductCondition_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	3: "NEW",
	4: "REFURBISHED",
	5: "USED",
}
var ProductConditionEnum_ProductCondition_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"NEW":         3,
	"REFURBISHED": 4,
	"USED":        5,
}

func (x ProductConditionEnum_ProductCondition) String() string {
	return proto.EnumName(ProductConditionEnum_ProductCondition_name, int32(x))
}
func (ProductConditionEnum_ProductCondition) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_product_condition_b012576f89964505, []int{0, 0}
}

// Condition of a product offer.
type ProductConditionEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductConditionEnum) Reset()         { *m = ProductConditionEnum{} }
func (m *ProductConditionEnum) String() string { return proto.CompactTextString(m) }
func (*ProductConditionEnum) ProtoMessage()    {}
func (*ProductConditionEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_condition_b012576f89964505, []int{0}
}
func (m *ProductConditionEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductConditionEnum.Unmarshal(m, b)
}
func (m *ProductConditionEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductConditionEnum.Marshal(b, m, deterministic)
}
func (dst *ProductConditionEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductConditionEnum.Merge(dst, src)
}
func (m *ProductConditionEnum) XXX_Size() int {
	return xxx_messageInfo_ProductConditionEnum.Size(m)
}
func (m *ProductConditionEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductConditionEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ProductConditionEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProductConditionEnum)(nil), "google.ads.googleads.v0.enums.ProductConditionEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.ProductConditionEnum_ProductCondition", ProductConditionEnum_ProductCondition_name, ProductConditionEnum_ProductCondition_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/product_condition.proto", fileDescriptor_product_condition_b012576f89964505)
}

var fileDescriptor_product_condition_b012576f89964505 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xdf, 0x6a, 0x83, 0x30,
	0x18, 0xc5, 0xd7, 0x3f, 0x5b, 0x47, 0x7a, 0xb1, 0x20, 0xdb, 0x65, 0x2f, 0xda, 0x07, 0x88, 0xc2,
	0xd8, 0x4d, 0x76, 0xa5, 0x6d, 0xda, 0xc9, 0xc0, 0x49, 0x9d, 0x16, 0x86, 0x30, 0x9c, 0x91, 0x20,
	0x68, 0x22, 0x46, 0xfb, 0x40, 0xbb, 0xdc, 0xa3, 0xec, 0x51, 0xfa, 0x14, 0xc3, 0x64, 0x7a, 0x51,
	0xd8, 0x6e, 0xe4, 0xf0, 0x9d, 0xf3, 0x33, 0xdf, 0xf9, 0xc0, 0x03, 0x13, 0x82, 0x15, 0x99, 0x99,
	0x50, 0x69, 0x6a, 0xd9, 0xa9, 0xa3, 0x65, 0x66, 0xbc, 0x2d, 0xa5, 0x59, 0xd5, 0x82, 0xb6, 0x69,
	0xf3, 0x9e, 0x0a, 0x4e, 0xf3, 0x26, 0x17, 0x1c, 0x55, 0xb5, 0x68, 0x84, 0xb1, 0xd0, 0x59, 0x94,
	0x50, 0x89, 0x06, 0x0c, 0x1d, 0x2d, 0xa4, 0xb0, 0x55, 0x01, 0x6e, 0x7d, 0x4d, 0xae, 0x7b, 0x90,
	0xf0, 0xb6, 0x5c, 0xbd, 0x02, 0x78, 0x3e, 0x37, 0x6e, 0xc0, 0x3c, 0xf4, 0x02, 0x9f, 0xac, 0xdd,
	0xad, 0x4b, 0x36, 0xf0, 0xc2, 0x98, 0x83, 0x59, 0xe8, 0x3d, 0x7b, 0x2f, 0x07, 0x0f, 0x8e, 0x8c,
	0x19, 0x98, 0x78, 0xe4, 0x00, 0x27, 0x5d, 0x6c, 0x4f, 0xb6, 0xe1, 0xde, 0x71, 0x83, 0x27, 0xb2,
	0x81, 0x53, 0xe3, 0x1a, 0x4c, 0xc3, 0x80, 0x6c, 0xe0, 0xa5, 0x73, 0x1a, 0x81, 0x65, 0x2a, 0x4a,
	0xf4, 0xef, 0x4e, 0xce, 0xdd, 0xf9, 0xcb, 0x7e, 0xd7, 0xc4, 0x1f, 0xbd, 0x39, 0xbf, 0x1c, 0x13,
	0x45, 0xc2, 0x19, 0x12, 0x35, 0x33, 0x59, 0xc6, 0x55, 0xcf, 0xfe, 0x24, 0x55, 0x2e, 0xff, 0xb8,
	0xd0, 0xa3, 0xfa, 0x7e, 0x8e, 0x27, 0x3b, 0xdb, 0xfe, 0x1a, 0x2f, 0x76, 0xfa, 0x57, 0x36, 0x95,
	0x48, 0xcb, 0x4e, 0x45, 0x16, 0xea, 0xda, 0xcb, 0xef, 0xde, 0x8f, 0x6d, 0x2a, 0xe3, 0xc1, 0x8f,
	0x23, 0x2b, 0x56, 0xfe, 0x69, 0xbc, 0xd4, 0x43, 0x8c, 0x6d, 0x2a, 0x31, 0x1e, 0x12, 0x18, 0x47,
	0x16, 0xc6, 0x2a, 0xf3, 0x71, 0xa5, 0x16, 0xbb, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xed,
	0x72, 0x33, 0xb9, 0x01, 0x00, 0x00,
}
