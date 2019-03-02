// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/frequency_cap_event_type.proto

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

// The type of event that the cap applies to (e.g. impression).
type FrequencyCapEventTypeEnum_FrequencyCapEventType int32

const (
	// Not specified.
	FrequencyCapEventTypeEnum_UNSPECIFIED FrequencyCapEventTypeEnum_FrequencyCapEventType = 0
	// Used for return value only. Represents value unknown in this version.
	FrequencyCapEventTypeEnum_UNKNOWN FrequencyCapEventTypeEnum_FrequencyCapEventType = 1
	// The cap applies on ad impressions.
	FrequencyCapEventTypeEnum_IMPRESSION FrequencyCapEventTypeEnum_FrequencyCapEventType = 2
	// The cap applies on video ad views.
	FrequencyCapEventTypeEnum_VIDEO_VIEW FrequencyCapEventTypeEnum_FrequencyCapEventType = 3
)

var FrequencyCapEventTypeEnum_FrequencyCapEventType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "IMPRESSION",
	3: "VIDEO_VIEW",
}
var FrequencyCapEventTypeEnum_FrequencyCapEventType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"IMPRESSION":  2,
	"VIDEO_VIEW":  3,
}

func (x FrequencyCapEventTypeEnum_FrequencyCapEventType) String() string {
	return proto.EnumName(FrequencyCapEventTypeEnum_FrequencyCapEventType_name, int32(x))
}
func (FrequencyCapEventTypeEnum_FrequencyCapEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_frequency_cap_event_type_1e16b8cee22d230e, []int{0, 0}
}

// Container for enum describing the type of event that the cap applies to.
type FrequencyCapEventTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrequencyCapEventTypeEnum) Reset()         { *m = FrequencyCapEventTypeEnum{} }
func (m *FrequencyCapEventTypeEnum) String() string { return proto.CompactTextString(m) }
func (*FrequencyCapEventTypeEnum) ProtoMessage()    {}
func (*FrequencyCapEventTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_frequency_cap_event_type_1e16b8cee22d230e, []int{0}
}
func (m *FrequencyCapEventTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrequencyCapEventTypeEnum.Unmarshal(m, b)
}
func (m *FrequencyCapEventTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrequencyCapEventTypeEnum.Marshal(b, m, deterministic)
}
func (dst *FrequencyCapEventTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrequencyCapEventTypeEnum.Merge(dst, src)
}
func (m *FrequencyCapEventTypeEnum) XXX_Size() int {
	return xxx_messageInfo_FrequencyCapEventTypeEnum.Size(m)
}
func (m *FrequencyCapEventTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FrequencyCapEventTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FrequencyCapEventTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FrequencyCapEventTypeEnum)(nil), "google.ads.googleads.v0.enums.FrequencyCapEventTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.FrequencyCapEventTypeEnum_FrequencyCapEventType", FrequencyCapEventTypeEnum_FrequencyCapEventType_name, FrequencyCapEventTypeEnum_FrequencyCapEventType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/frequency_cap_event_type.proto", fileDescriptor_frequency_cap_event_type_1e16b8cee22d230e)
}

var fileDescriptor_frequency_cap_event_type_1e16b8cee22d230e = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0xb4, 0xa2, 0xd4, 0xc2, 0xd2, 0xd4, 0xbc, 0xe4, 0xca, 0xf8, 0xe4, 0xc4,
	0x82, 0xf8, 0xd4, 0xb2, 0xd4, 0xbc, 0x92, 0xf8, 0x92, 0xca, 0x82, 0x54, 0xbd, 0x82, 0xa2, 0xfc,
	0x92, 0x7c, 0x21, 0x59, 0x88, 0x16, 0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x6e, 0xbd, 0x32, 0x03,
	0x3d, 0xb0, 0x6e, 0xa5, 0x22, 0x2e, 0x49, 0x37, 0x98, 0x01, 0xce, 0x89, 0x05, 0xae, 0x20, 0xed,
	0x21, 0x95, 0x05, 0xa9, 0xae, 0x79, 0xa5, 0xb9, 0x4a, 0xa1, 0x5c, 0xa2, 0x58, 0x25, 0x85, 0xf8,
	0xb9, 0xb8, 0x43, 0xfd, 0x82, 0x03, 0x5c, 0x9d, 0x3d, 0xdd, 0x3c, 0x5d, 0x5d, 0x04, 0x18, 0x84,
	0xb8, 0xb9, 0xd8, 0x43, 0xfd, 0xbc, 0xfd, 0xfc, 0xc3, 0xfd, 0x04, 0x18, 0x85, 0xf8, 0xb8, 0xb8,
	0x3c, 0x7d, 0x03, 0x82, 0x5c, 0x83, 0x83, 0x3d, 0xfd, 0xfd, 0x04, 0x98, 0x40, 0xfc, 0x30, 0x4f,
	0x17, 0x57, 0xff, 0xf8, 0x30, 0x4f, 0xd7, 0x70, 0x01, 0x66, 0xa7, 0xf7, 0x8c, 0x5c, 0x8a, 0xc9,
	0xf9, 0xb9, 0x7a, 0x78, 0x5d, 0xe6, 0x24, 0x85, 0xd5, 0xea, 0x00, 0x90, 0xa7, 0x02, 0x18, 0xa3,
	0x9c, 0xa0, 0x9a, 0xd3, 0xf3, 0x73, 0x12, 0xf3, 0xd2, 0xf5, 0xf2, 0x8b, 0xd2, 0xf5, 0xd3, 0x53,
	0xf3, 0xc0, 0x5e, 0x86, 0x05, 0x52, 0x41, 0x66, 0x31, 0x8e, 0x30, 0xb3, 0x06, 0x93, 0x8b, 0x98,
	0x98, 0xdd, 0x1d, 0x1d, 0x57, 0x31, 0xc9, 0xba, 0x43, 0x8c, 0x72, 0x4c, 0x29, 0xd6, 0x83, 0x30,
	0x41, 0xac, 0x30, 0x03, 0x3d, 0x50, 0x18, 0x14, 0x9f, 0x82, 0xc9, 0xc7, 0x38, 0xa6, 0x14, 0xc7,
	0xc0, 0xe5, 0x63, 0xc2, 0x0c, 0x62, 0xc0, 0xf2, 0xaf, 0x98, 0x14, 0x21, 0x82, 0x56, 0x56, 0x8e,
	0x29, 0xc5, 0x56, 0x56, 0x70, 0x15, 0x56, 0x56, 0x61, 0x06, 0x56, 0x56, 0x60, 0x35, 0x49, 0x6c,
	0x60, 0x87, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x91, 0x6f, 0x0a, 0xa0, 0xcb, 0x01, 0x00,
	0x00,
}
