// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/webpage_condition_operand.proto

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

// The webpage condition operand in webpage criterion.
type WebpageConditionOperandEnum_WebpageConditionOperand int32

const (
	// Not specified.
	WebpageConditionOperandEnum_UNSPECIFIED WebpageConditionOperandEnum_WebpageConditionOperand = 0
	// Used for return value only. Represents value unknown in this version.
	WebpageConditionOperandEnum_UNKNOWN WebpageConditionOperandEnum_WebpageConditionOperand = 1
	// Operand denoting a webpage URL targeting condition.
	WebpageConditionOperandEnum_URL WebpageConditionOperandEnum_WebpageConditionOperand = 2
	// Operand denoting a webpage category targeting condition.
	WebpageConditionOperandEnum_CATEGORY WebpageConditionOperandEnum_WebpageConditionOperand = 3
	// Operand denoting a webpage title targeting condition.
	WebpageConditionOperandEnum_PAGE_TITLE WebpageConditionOperandEnum_WebpageConditionOperand = 4
	// Operand denoting a webpage content targeting condition.
	WebpageConditionOperandEnum_PAGE_CONTENT WebpageConditionOperandEnum_WebpageConditionOperand = 5
	// Operand denoting a webpage custom label targeting condition.
	WebpageConditionOperandEnum_CUSTOM_LABEL WebpageConditionOperandEnum_WebpageConditionOperand = 6
)

var WebpageConditionOperandEnum_WebpageConditionOperand_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "URL",
	3: "CATEGORY",
	4: "PAGE_TITLE",
	5: "PAGE_CONTENT",
	6: "CUSTOM_LABEL",
}
var WebpageConditionOperandEnum_WebpageConditionOperand_value = map[string]int32{
	"UNSPECIFIED":  0,
	"UNKNOWN":      1,
	"URL":          2,
	"CATEGORY":     3,
	"PAGE_TITLE":   4,
	"PAGE_CONTENT": 5,
	"CUSTOM_LABEL": 6,
}

func (x WebpageConditionOperandEnum_WebpageConditionOperand) String() string {
	return proto.EnumName(WebpageConditionOperandEnum_WebpageConditionOperand_name, int32(x))
}
func (WebpageConditionOperandEnum_WebpageConditionOperand) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_webpage_condition_operand_c33b7f446ef8416d, []int{0, 0}
}

// Container for enum describing webpage condition operand in webpage criterion.
type WebpageConditionOperandEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebpageConditionOperandEnum) Reset()         { *m = WebpageConditionOperandEnum{} }
func (m *WebpageConditionOperandEnum) String() string { return proto.CompactTextString(m) }
func (*WebpageConditionOperandEnum) ProtoMessage()    {}
func (*WebpageConditionOperandEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_webpage_condition_operand_c33b7f446ef8416d, []int{0}
}
func (m *WebpageConditionOperandEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebpageConditionOperandEnum.Unmarshal(m, b)
}
func (m *WebpageConditionOperandEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebpageConditionOperandEnum.Marshal(b, m, deterministic)
}
func (dst *WebpageConditionOperandEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebpageConditionOperandEnum.Merge(dst, src)
}
func (m *WebpageConditionOperandEnum) XXX_Size() int {
	return xxx_messageInfo_WebpageConditionOperandEnum.Size(m)
}
func (m *WebpageConditionOperandEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_WebpageConditionOperandEnum.DiscardUnknown(m)
}

var xxx_messageInfo_WebpageConditionOperandEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*WebpageConditionOperandEnum)(nil), "google.ads.googleads.v0.enums.WebpageConditionOperandEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.WebpageConditionOperandEnum_WebpageConditionOperand", WebpageConditionOperandEnum_WebpageConditionOperand_name, WebpageConditionOperandEnum_WebpageConditionOperand_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/webpage_condition_operand.proto", fileDescriptor_webpage_condition_operand_c33b7f446ef8416d)
}

var fileDescriptor_webpage_condition_operand_c33b7f446ef8416d = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x6e, 0xe2, 0x30,
	0x18, 0x85, 0x27, 0x61, 0x06, 0x46, 0x06, 0xcd, 0x58, 0xd9, 0x74, 0xd1, 0xb2, 0x80, 0x03, 0x38,
	0x91, 0xba, 0x73, 0xd5, 0x85, 0x93, 0xba, 0x11, 0x6a, 0x9a, 0x44, 0x90, 0x80, 0x5a, 0x45, 0x8a,
	0x02, 0x89, 0x2c, 0x24, 0xb0, 0x23, 0x0c, 0x74, 0xdf, 0xa3, 0xb0, 0xec, 0x51, 0x7a, 0x94, 0xee,
	0x7a, 0x83, 0x2a, 0x76, 0x61, 0x47, 0x37, 0xd1, 0xcb, 0xff, 0xfe, 0xf7, 0xf4, 0xfb, 0x03, 0xb7,
	0x4c, 0x08, 0xb6, 0xaa, 0xec, 0xa2, 0x94, 0xb6, 0x96, 0x8d, 0xda, 0x3b, 0x76, 0xc5, 0x77, 0x6b,
	0x69, 0xbf, 0x54, 0xf3, 0xba, 0x60, 0x55, 0xbe, 0x10, 0xbc, 0x5c, 0x6e, 0x97, 0x82, 0xe7, 0xa2,
	0xae, 0x36, 0x05, 0x2f, 0x51, 0xbd, 0x11, 0x5b, 0x61, 0xf5, 0x75, 0x06, 0x15, 0xa5, 0x44, 0xa7,
	0x38, 0xda, 0x3b, 0x48, 0xc5, 0x87, 0x07, 0x03, 0x5c, 0xce, 0x74, 0x85, 0x77, 0x6c, 0x88, 0x74,
	0x01, 0xe5, 0xbb, 0xf5, 0xf0, 0xd5, 0x00, 0x17, 0x67, 0x7c, 0xeb, 0x3f, 0xe8, 0xa6, 0xe1, 0x24,
	0xa6, 0xde, 0xe8, 0x7e, 0x44, 0xef, 0xe0, 0x2f, 0xab, 0x0b, 0x3a, 0x69, 0xf8, 0x10, 0x46, 0xb3,
	0x10, 0x1a, 0x56, 0x07, 0xb4, 0xd2, 0x71, 0x00, 0x4d, 0xab, 0x07, 0xfe, 0x7a, 0x24, 0xa1, 0x7e,
	0x34, 0x7e, 0x82, 0x2d, 0xeb, 0x1f, 0x00, 0x31, 0xf1, 0x69, 0x9e, 0x8c, 0x92, 0x80, 0xc2, 0xdf,
	0x16, 0x04, 0x3d, 0xf5, 0xef, 0x45, 0x61, 0x42, 0xc3, 0x04, 0xfe, 0x69, 0x26, 0x5e, 0x3a, 0x49,
	0xa2, 0xc7, 0x3c, 0x20, 0x2e, 0x0d, 0x60, 0xdb, 0xfd, 0x34, 0xc0, 0x60, 0x21, 0xd6, 0xe8, 0xc7,
	0xa7, 0xb8, 0x57, 0x67, 0xee, 0x8c, 0x1b, 0x0e, 0xb1, 0xf1, 0xec, 0x7e, 0xc7, 0x99, 0x58, 0x15,
	0x9c, 0x21, 0xb1, 0x61, 0x36, 0xab, 0xb8, 0xa2, 0x74, 0x04, 0x5b, 0x2f, 0xe5, 0x19, 0xce, 0x37,
	0xea, 0x7b, 0x30, 0x5b, 0x3e, 0x21, 0x6f, 0x66, 0xdf, 0xd7, 0x55, 0xa4, 0x94, 0x48, 0xcb, 0x46,
	0x4d, 0x1d, 0xd4, 0x30, 0x93, 0xef, 0x47, 0x3f, 0x23, 0xa5, 0xcc, 0x4e, 0x7e, 0x36, 0x75, 0x32,
	0xe5, 0x7f, 0x98, 0x03, 0x3d, 0xc4, 0x98, 0x94, 0x12, 0xe3, 0xd3, 0x06, 0xc6, 0x53, 0x07, 0x63,
	0xb5, 0x33, 0x6f, 0xab, 0xc3, 0xae, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x74, 0xc1, 0x73,
	0xff, 0x01, 0x00, 0x00,
}
