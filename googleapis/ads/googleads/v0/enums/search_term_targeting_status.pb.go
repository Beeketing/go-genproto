// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/search_term_targeting_status.proto

package enums // import "github.com/Beeketing/genproto/googleapis/ads/googleads/v0/enums"

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

// Indicates whether the search term is one of your targeted or excluded
// keywords.
type SearchTermTargetingStatusEnum_SearchTermTargetingStatus int32

const (
	// Not specified.
	SearchTermTargetingStatusEnum_UNSPECIFIED SearchTermTargetingStatusEnum_SearchTermTargetingStatus = 0
	// Used for return value only. Represents value unknown in this version.
	SearchTermTargetingStatusEnum_UNKNOWN SearchTermTargetingStatusEnum_SearchTermTargetingStatus = 1
	// Search term is added to targeted keywords.
	SearchTermTargetingStatusEnum_ADDED SearchTermTargetingStatusEnum_SearchTermTargetingStatus = 2
	// Search term matches a negative keyword.
	SearchTermTargetingStatusEnum_EXCLUDED SearchTermTargetingStatusEnum_SearchTermTargetingStatus = 3
	// Search term has been both added and excluded.
	SearchTermTargetingStatusEnum_ADDED_EXCLUDED SearchTermTargetingStatusEnum_SearchTermTargetingStatus = 4
	// Search term is neither targeted nor excluded.
	SearchTermTargetingStatusEnum_NONE SearchTermTargetingStatusEnum_SearchTermTargetingStatus = 5
)

var SearchTermTargetingStatusEnum_SearchTermTargetingStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ADDED",
	3: "EXCLUDED",
	4: "ADDED_EXCLUDED",
	5: "NONE",
}
var SearchTermTargetingStatusEnum_SearchTermTargetingStatus_value = map[string]int32{
	"UNSPECIFIED":    0,
	"UNKNOWN":        1,
	"ADDED":          2,
	"EXCLUDED":       3,
	"ADDED_EXCLUDED": 4,
	"NONE":           5,
}

func (x SearchTermTargetingStatusEnum_SearchTermTargetingStatus) String() string {
	return proto.EnumName(SearchTermTargetingStatusEnum_SearchTermTargetingStatus_name, int32(x))
}
func (SearchTermTargetingStatusEnum_SearchTermTargetingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_search_term_targeting_status_ca848c12c28e24e1, []int{0, 0}
}

// Container for enum indicating whether a search term is one of your targeted
// or excluded keywords.
type SearchTermTargetingStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchTermTargetingStatusEnum) Reset()         { *m = SearchTermTargetingStatusEnum{} }
func (m *SearchTermTargetingStatusEnum) String() string { return proto.CompactTextString(m) }
func (*SearchTermTargetingStatusEnum) ProtoMessage()    {}
func (*SearchTermTargetingStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_search_term_targeting_status_ca848c12c28e24e1, []int{0}
}
func (m *SearchTermTargetingStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchTermTargetingStatusEnum.Unmarshal(m, b)
}
func (m *SearchTermTargetingStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchTermTargetingStatusEnum.Marshal(b, m, deterministic)
}
func (dst *SearchTermTargetingStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchTermTargetingStatusEnum.Merge(dst, src)
}
func (m *SearchTermTargetingStatusEnum) XXX_Size() int {
	return xxx_messageInfo_SearchTermTargetingStatusEnum.Size(m)
}
func (m *SearchTermTargetingStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchTermTargetingStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SearchTermTargetingStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SearchTermTargetingStatusEnum)(nil), "google.ads.googleads.v0.enums.SearchTermTargetingStatusEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.SearchTermTargetingStatusEnum_SearchTermTargetingStatus", SearchTermTargetingStatusEnum_SearchTermTargetingStatus_name, SearchTermTargetingStatusEnum_SearchTermTargetingStatus_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/search_term_targeting_status.proto", fileDescriptor_search_term_targeting_status_ca848c12c28e24e1)
}

var fileDescriptor_search_term_targeting_status_ca848c12c28e24e1 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4e, 0xeb, 0x30,
	0x18, 0xc4, 0x5f, 0xd2, 0xf6, 0x51, 0x5c, 0x04, 0x91, 0x77, 0x2c, 0x8a, 0xd4, 0x1e, 0xc0, 0x89,
	0xc4, 0xce, 0x6c, 0x70, 0x1b, 0x53, 0x55, 0x20, 0xb7, 0x52, 0xff, 0x80, 0x50, 0xa4, 0xc8, 0x34,
	0x96, 0xa9, 0xd4, 0xc4, 0x95, 0xed, 0xf6, 0x1e, 0x5c, 0x81, 0x25, 0x47, 0xe1, 0x28, 0x6c, 0xb9,
	0x00, 0x8a, 0x43, 0xb3, 0x2b, 0x1b, 0x6b, 0xfc, 0xcd, 0xe7, 0x9f, 0xc6, 0x03, 0x6e, 0xa5, 0x52,
	0x72, 0x23, 0x42, 0x9e, 0x99, 0xb0, 0x92, 0xa5, 0xda, 0x47, 0xa1, 0x28, 0x76, 0xb9, 0x09, 0x8d,
	0xe0, 0x7a, 0xf5, 0x9a, 0x5a, 0xa1, 0xf3, 0xd4, 0x72, 0x2d, 0x85, 0x5d, 0x17, 0x32, 0x35, 0x96,
	0xdb, 0x9d, 0x41, 0x5b, 0xad, 0xac, 0x82, 0xdd, 0xea, 0x19, 0xe2, 0x99, 0x41, 0x35, 0x01, 0xed,
	0x23, 0xe4, 0x08, 0xfd, 0x37, 0x0f, 0x74, 0x67, 0x8e, 0x32, 0x17, 0x3a, 0x9f, 0x1f, 0x18, 0x33,
	0x87, 0xa0, 0xc5, 0x2e, 0xef, 0x6f, 0xc1, 0xe5, 0xd1, 0x05, 0x78, 0x01, 0x3a, 0x0b, 0x36, 0x9b,
	0xd2, 0xe1, 0xf8, 0x6e, 0x4c, 0xe3, 0xe0, 0x1f, 0xec, 0x80, 0x93, 0x05, 0xbb, 0x67, 0x93, 0x47,
	0x16, 0x78, 0xf0, 0x14, 0xb4, 0x48, 0x1c, 0xd3, 0x38, 0xf0, 0xe1, 0x19, 0x68, 0xd3, 0xa7, 0xe1,
	0xc3, 0xa2, 0xbc, 0x35, 0x20, 0x04, 0xe7, 0xce, 0x48, 0xeb, 0x59, 0x13, 0xb6, 0x41, 0x93, 0x4d,
	0x18, 0x0d, 0x5a, 0x83, 0x6f, 0x0f, 0xf4, 0x56, 0x2a, 0x47, 0x7f, 0x26, 0x1f, 0x5c, 0x1d, 0x4d,
	0x35, 0x2d, 0x3f, 0x3e, 0xf5, 0x9e, 0x07, 0xbf, 0x00, 0xa9, 0x36, 0xbc, 0x90, 0x48, 0x69, 0x19,
	0x4a, 0x51, 0xb8, 0x5a, 0x0e, 0x65, 0x6e, 0xd7, 0xe6, 0x48, 0xb7, 0x37, 0xee, 0x7c, 0xf7, 0x1b,
	0x23, 0x42, 0x3e, 0xfc, 0xee, 0xa8, 0x42, 0x91, 0xcc, 0xa0, 0x4a, 0x96, 0x6a, 0x19, 0xa1, 0xb2,
	0x22, 0xf3, 0x79, 0xf0, 0x13, 0x92, 0x99, 0xa4, 0xf6, 0x93, 0x65, 0x94, 0x38, 0xff, 0xcb, 0xef,
	0x55, 0x43, 0x8c, 0x49, 0x66, 0x30, 0xae, 0x37, 0x30, 0x5e, 0x46, 0x18, 0xbb, 0x9d, 0x97, 0xff,
	0x2e, 0xd8, 0xf5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x15, 0x39, 0xfd, 0xf3, 0x01, 0x00,
	0x00,
}
