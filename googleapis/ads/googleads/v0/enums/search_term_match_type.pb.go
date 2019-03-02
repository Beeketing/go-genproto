// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/search_term_match_type.proto

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

// Possible match types for a keyword triggering an ad, including variants.
type SearchTermMatchTypeEnum_SearchTermMatchType int32

const (
	// Not specified.
	SearchTermMatchTypeEnum_UNSPECIFIED SearchTermMatchTypeEnum_SearchTermMatchType = 0
	// Used for return value only. Represents value unknown in this version.
	SearchTermMatchTypeEnum_UNKNOWN SearchTermMatchTypeEnum_SearchTermMatchType = 1
	// Broad match.
	SearchTermMatchTypeEnum_BROAD SearchTermMatchTypeEnum_SearchTermMatchType = 2
	// Exact match.
	SearchTermMatchTypeEnum_EXACT SearchTermMatchTypeEnum_SearchTermMatchType = 3
	// Phrase match.
	SearchTermMatchTypeEnum_PHRASE SearchTermMatchTypeEnum_SearchTermMatchType = 4
	// Exact match (close variant).
	SearchTermMatchTypeEnum_NEAR_EXACT SearchTermMatchTypeEnum_SearchTermMatchType = 5
	// Phrase match (close variant).
	SearchTermMatchTypeEnum_NEAR_PHRASE SearchTermMatchTypeEnum_SearchTermMatchType = 6
)

var SearchTermMatchTypeEnum_SearchTermMatchType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "BROAD",
	3: "EXACT",
	4: "PHRASE",
	5: "NEAR_EXACT",
	6: "NEAR_PHRASE",
}
var SearchTermMatchTypeEnum_SearchTermMatchType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"BROAD":       2,
	"EXACT":       3,
	"PHRASE":      4,
	"NEAR_EXACT":  5,
	"NEAR_PHRASE": 6,
}

func (x SearchTermMatchTypeEnum_SearchTermMatchType) String() string {
	return proto.EnumName(SearchTermMatchTypeEnum_SearchTermMatchType_name, int32(x))
}
func (SearchTermMatchTypeEnum_SearchTermMatchType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_search_term_match_type_3e483cda240776af, []int{0, 0}
}

// Container for enum describing match types for a keyword triggering an ad.
type SearchTermMatchTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchTermMatchTypeEnum) Reset()         { *m = SearchTermMatchTypeEnum{} }
func (m *SearchTermMatchTypeEnum) String() string { return proto.CompactTextString(m) }
func (*SearchTermMatchTypeEnum) ProtoMessage()    {}
func (*SearchTermMatchTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_search_term_match_type_3e483cda240776af, []int{0}
}
func (m *SearchTermMatchTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchTermMatchTypeEnum.Unmarshal(m, b)
}
func (m *SearchTermMatchTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchTermMatchTypeEnum.Marshal(b, m, deterministic)
}
func (dst *SearchTermMatchTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchTermMatchTypeEnum.Merge(dst, src)
}
func (m *SearchTermMatchTypeEnum) XXX_Size() int {
	return xxx_messageInfo_SearchTermMatchTypeEnum.Size(m)
}
func (m *SearchTermMatchTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchTermMatchTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SearchTermMatchTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SearchTermMatchTypeEnum)(nil), "google.ads.googleads.v0.enums.SearchTermMatchTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.SearchTermMatchTypeEnum_SearchTermMatchType", SearchTermMatchTypeEnum_SearchTermMatchType_name, SearchTermMatchTypeEnum_SearchTermMatchType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/search_term_match_type.proto", fileDescriptor_search_term_match_type_3e483cda240776af)
}

var fileDescriptor_search_term_match_type_3e483cda240776af = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xdf, 0x4a, 0xc3, 0x30,
	0x18, 0xc5, 0x6d, 0xe7, 0x26, 0x66, 0xa0, 0x25, 0x5e, 0xe8, 0xcd, 0x2e, 0xb6, 0x07, 0x48, 0x0b,
	0xde, 0xc5, 0xab, 0x74, 0x8b, 0x73, 0x88, 0x5d, 0xd9, 0x3f, 0x45, 0x0a, 0xa3, 0xae, 0x21, 0x0a,
	0x4b, 0x53, 0x9a, 0x6d, 0xb0, 0xc7, 0xf0, 0x15, 0xbc, 0xf4, 0x51, 0x7c, 0x14, 0xc1, 0x77, 0x90,
	0x24, 0x5b, 0xaf, 0xa6, 0x37, 0xe1, 0x24, 0xe7, 0xfb, 0x7d, 0x9c, 0x1c, 0x80, 0xb9, 0x94, 0x7c,
	0xc9, 0xfc, 0x34, 0x53, 0xbe, 0x95, 0x5a, 0x6d, 0x02, 0x9f, 0xe5, 0x6b, 0xa1, 0x7c, 0xc5, 0xd2,
	0x72, 0xf1, 0x3a, 0x5f, 0xb1, 0x52, 0xcc, 0x45, 0xba, 0xd2, 0x72, 0x5b, 0x30, 0x54, 0x94, 0x72,
	0x25, 0x61, 0xcb, 0x02, 0x28, 0xcd, 0x14, 0xaa, 0x58, 0xb4, 0x09, 0x90, 0x61, 0x3b, 0xef, 0x0e,
	0xb8, 0x1c, 0x1b, 0x7e, 0xc2, 0x4a, 0xf1, 0xa0, 0xe9, 0xc9, 0xb6, 0x60, 0x34, 0x5f, 0x8b, 0xce,
	0x06, 0x5c, 0x1c, 0xb0, 0xe0, 0x39, 0x68, 0x4e, 0xa3, 0x71, 0x4c, 0xbb, 0x83, 0xdb, 0x01, 0xed,
	0x79, 0x47, 0xb0, 0x09, 0x4e, 0xa6, 0xd1, 0x7d, 0x34, 0x7c, 0x8c, 0x3c, 0x07, 0x9e, 0x82, 0x7a,
	0x38, 0x1a, 0x92, 0x9e, 0xe7, 0x6a, 0x49, 0x9f, 0x48, 0x77, 0xe2, 0xd5, 0x20, 0x00, 0x8d, 0xf8,
	0x6e, 0x44, 0xc6, 0xd4, 0x3b, 0x86, 0x67, 0x00, 0x44, 0x94, 0x8c, 0xe6, 0xd6, 0xab, 0xeb, 0x7d,
	0xe6, 0xbe, 0x1b, 0x68, 0x84, 0x3f, 0x0e, 0x68, 0x2f, 0xa4, 0x40, 0xff, 0x26, 0x0f, 0xaf, 0x0e,
	0x64, 0x8b, 0xf5, 0x97, 0x63, 0xe7, 0x39, 0xdc, 0xa1, 0x5c, 0x2e, 0xd3, 0x9c, 0x23, 0x59, 0x72,
	0x9f, 0xb3, 0xdc, 0x14, 0xb2, 0x2f, 0xb0, 0x78, 0x53, 0x7f, 0xf4, 0x79, 0x63, 0xce, 0x0f, 0xb7,
	0xd6, 0x27, 0xe4, 0xd3, 0x6d, 0xf5, 0xed, 0x2a, 0x92, 0x29, 0x64, 0xa5, 0x56, 0xb3, 0x00, 0xe9,
	0x8a, 0xd4, 0xd7, 0xde, 0x4f, 0x48, 0xa6, 0x92, 0xca, 0x4f, 0x66, 0x41, 0x62, 0xfc, 0x6f, 0xb7,
	0x6d, 0x1f, 0x31, 0x26, 0x99, 0xc2, 0xb8, 0x9a, 0xc0, 0x78, 0x16, 0x60, 0x6c, 0x66, 0x5e, 0x1a,
	0x26, 0xd8, 0xf5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0xf7, 0x90, 0xbe, 0xe7, 0x01, 0x00,
	0x00,
}
