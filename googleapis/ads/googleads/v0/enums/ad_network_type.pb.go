// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/ad_network_type.proto

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

// Enumerates Google Ads network types.
type AdNetworkTypeEnum_AdNetworkType int32

const (
	// Not specified.
	AdNetworkTypeEnum_UNSPECIFIED AdNetworkTypeEnum_AdNetworkType = 0
	// The value is unknown in this version.
	AdNetworkTypeEnum_UNKNOWN AdNetworkTypeEnum_AdNetworkType = 1
	// Google search.
	AdNetworkTypeEnum_SEARCH AdNetworkTypeEnum_AdNetworkType = 2
	// Search partners.
	AdNetworkTypeEnum_SEARCH_PARTNERS AdNetworkTypeEnum_AdNetworkType = 3
	// Display Network.
	AdNetworkTypeEnum_CONTENT AdNetworkTypeEnum_AdNetworkType = 4
	// YouTube Search.
	AdNetworkTypeEnum_YOUTUBE_SEARCH AdNetworkTypeEnum_AdNetworkType = 5
	// YouTube Videos
	AdNetworkTypeEnum_YOUTUBE_WATCH AdNetworkTypeEnum_AdNetworkType = 6
	// Cross-network.
	AdNetworkTypeEnum_MIXED AdNetworkTypeEnum_AdNetworkType = 7
)

var AdNetworkTypeEnum_AdNetworkType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SEARCH",
	3: "SEARCH_PARTNERS",
	4: "CONTENT",
	5: "YOUTUBE_SEARCH",
	6: "YOUTUBE_WATCH",
	7: "MIXED",
}
var AdNetworkTypeEnum_AdNetworkType_value = map[string]int32{
	"UNSPECIFIED":     0,
	"UNKNOWN":         1,
	"SEARCH":          2,
	"SEARCH_PARTNERS": 3,
	"CONTENT":         4,
	"YOUTUBE_SEARCH":  5,
	"YOUTUBE_WATCH":   6,
	"MIXED":           7,
}

func (x AdNetworkTypeEnum_AdNetworkType) String() string {
	return proto.EnumName(AdNetworkTypeEnum_AdNetworkType_name, int32(x))
}
func (AdNetworkTypeEnum_AdNetworkType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad_network_type_4963ee02d7690a17, []int{0, 0}
}

// Container for enumeration of Google Ads network types.
type AdNetworkTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdNetworkTypeEnum) Reset()         { *m = AdNetworkTypeEnum{} }
func (m *AdNetworkTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AdNetworkTypeEnum) ProtoMessage()    {}
func (*AdNetworkTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_network_type_4963ee02d7690a17, []int{0}
}
func (m *AdNetworkTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdNetworkTypeEnum.Unmarshal(m, b)
}
func (m *AdNetworkTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdNetworkTypeEnum.Marshal(b, m, deterministic)
}
func (dst *AdNetworkTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdNetworkTypeEnum.Merge(dst, src)
}
func (m *AdNetworkTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AdNetworkTypeEnum.Size(m)
}
func (m *AdNetworkTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdNetworkTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdNetworkTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdNetworkTypeEnum)(nil), "google.ads.googleads.v0.enums.AdNetworkTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.AdNetworkTypeEnum_AdNetworkType", AdNetworkTypeEnum_AdNetworkType_name, AdNetworkTypeEnum_AdNetworkType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/ad_network_type.proto", fileDescriptor_ad_network_type_4963ee02d7690a17)
}

var fileDescriptor_ad_network_type_4963ee02d7690a17 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x6e, 0xe2, 0x30,
	0x18, 0x85, 0x27, 0x61, 0x00, 0x8d, 0x11, 0x83, 0xf1, 0xac, 0x59, 0xc0, 0x01, 0x9c, 0x48, 0xec,
	0x3c, 0x2b, 0x27, 0xb8, 0x80, 0xaa, 0x9a, 0x08, 0x12, 0x68, 0xab, 0x48, 0x51, 0x5a, 0x47, 0x51,
	0x55, 0x88, 0x23, 0x0c, 0x54, 0x1c, 0xa2, 0x97, 0x68, 0x77, 0x3d, 0x4a, 0x2f, 0x52, 0xa9, 0xa7,
	0xa8, 0x12, 0x03, 0x12, 0x8b, 0x76, 0x63, 0x3d, 0xfd, 0xef, 0x7d, 0xd6, 0xff, 0x3f, 0xd0, 0x4f,
	0xa5, 0x4c, 0x97, 0x89, 0x15, 0x0b, 0x65, 0x69, 0x59, 0xa8, 0x9d, 0x6d, 0x25, 0xd9, 0x76, 0xa5,
	0xac, 0x58, 0x44, 0x59, 0xb2, 0x79, 0x92, 0xeb, 0xc7, 0x68, 0xb3, 0xcf, 0x13, 0x9c, 0xaf, 0xe5,
	0x46, 0xa2, 0x8e, 0x4e, 0xe2, 0x58, 0x28, 0x7c, 0x82, 0xf0, 0xce, 0xc6, 0x25, 0xd4, 0x7b, 0x35,
	0x40, 0x9b, 0x0a, 0xae, 0x39, 0x7f, 0x9f, 0x27, 0x2c, 0xdb, 0xae, 0x7a, 0xcf, 0x06, 0x68, 0x9e,
	0x4d, 0x51, 0x0b, 0x34, 0x02, 0x3e, 0xf3, 0x98, 0x3b, 0xbe, 0x18, 0xb3, 0x01, 0xfc, 0x85, 0x1a,
	0xa0, 0x1e, 0xf0, 0x4b, 0x3e, 0x59, 0x70, 0x68, 0x20, 0x00, 0x6a, 0x33, 0x46, 0xa7, 0xee, 0x08,
	0x9a, 0xe8, 0x1f, 0x68, 0x69, 0x1d, 0x79, 0x74, 0xea, 0x73, 0x36, 0x9d, 0xc1, 0x4a, 0x91, 0x76,
	0x27, 0xdc, 0x67, 0xdc, 0x87, 0xbf, 0x11, 0x02, 0x7f, 0x6f, 0x26, 0x81, 0x1f, 0x38, 0x2c, 0x3a,
	0x50, 0x55, 0xd4, 0x06, 0xcd, 0xe3, 0x6c, 0x41, 0x7d, 0x77, 0x04, 0x6b, 0xe8, 0x0f, 0xa8, 0x5e,
	0x8d, 0xaf, 0xd9, 0x00, 0xd6, 0x9d, 0x0f, 0x03, 0x74, 0xef, 0xe5, 0x0a, 0xff, 0x78, 0x8b, 0x83,
	0xce, 0x56, 0xf6, 0x8a, 0xf3, 0x3d, 0xe3, 0xd6, 0x39, 0x40, 0xa9, 0x5c, 0xc6, 0x59, 0x8a, 0xe5,
	0x3a, 0xb5, 0xd2, 0x24, 0x2b, 0xcb, 0x39, 0xb6, 0x98, 0x3f, 0xa8, 0x6f, 0x4a, 0xfd, 0x5f, 0xbe,
	0x2f, 0x66, 0x65, 0x48, 0xe9, 0x9b, 0xd9, 0x19, 0xea, 0xaf, 0xa8, 0x50, 0x58, 0xcb, 0x42, 0xcd,
	0x6d, 0x5c, 0x94, 0xa6, 0xde, 0x8f, 0x7e, 0x48, 0x85, 0x0a, 0x4f, 0x7e, 0x38, 0xb7, 0xc3, 0xd2,
	0xff, 0x34, 0xbb, 0x7a, 0x48, 0x08, 0x15, 0x8a, 0x90, 0x53, 0x82, 0x90, 0xb9, 0x4d, 0x48, 0x99,
	0xb9, 0xab, 0x95, 0x8b, 0xf5, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xe7, 0x00, 0x7e, 0xec,
	0x01, 0x00, 0x00,
}
