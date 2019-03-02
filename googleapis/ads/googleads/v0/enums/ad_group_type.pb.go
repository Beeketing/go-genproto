// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/ad_group_type.proto

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

// Enum listing the possible types of an ad group.
type AdGroupTypeEnum_AdGroupType int32

const (
	// The type has not been specified.
	AdGroupTypeEnum_UNSPECIFIED AdGroupTypeEnum_AdGroupType = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdGroupTypeEnum_UNKNOWN AdGroupTypeEnum_AdGroupType = 1
	// The default ad group type for Search campaigns.
	AdGroupTypeEnum_SEARCH_STANDARD AdGroupTypeEnum_AdGroupType = 2
	// The default ad group type for Display campaigns.
	AdGroupTypeEnum_DISPLAY_STANDARD AdGroupTypeEnum_AdGroupType = 3
	// The ad group type for Shopping campaigns serving standard product ads.
	AdGroupTypeEnum_SHOPPING_PRODUCT_ADS AdGroupTypeEnum_AdGroupType = 4
	// The default ad group type for Hotel campaigns.
	AdGroupTypeEnum_HOTEL_ADS AdGroupTypeEnum_AdGroupType = 6
	// The type for ad groups in Smart Shopping campaigns.
	AdGroupTypeEnum_SHOPPING_SMART_ADS AdGroupTypeEnum_AdGroupType = 7
	// Short unskippable in-stream video ads.
	AdGroupTypeEnum_VIDEO_BUMPER AdGroupTypeEnum_AdGroupType = 8
	// TrueView (skippable) in-stream video ads.
	AdGroupTypeEnum_VIDEO_TRUE_VIEW_IN_STREAM AdGroupTypeEnum_AdGroupType = 9
	// TrueView in-display video ads.
	AdGroupTypeEnum_VIDEO_TRUE_VIEW_IN_DISPLAY AdGroupTypeEnum_AdGroupType = 10
	// Unskippable in-stream video ads.
	AdGroupTypeEnum_VIDEO_NON_SKIPPABLE_IN_STREAM AdGroupTypeEnum_AdGroupType = 11
	// Outstream video ads.
	AdGroupTypeEnum_VIDEO_OUTSTREAM AdGroupTypeEnum_AdGroupType = 12
)

var AdGroupTypeEnum_AdGroupType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "SEARCH_STANDARD",
	3:  "DISPLAY_STANDARD",
	4:  "SHOPPING_PRODUCT_ADS",
	6:  "HOTEL_ADS",
	7:  "SHOPPING_SMART_ADS",
	8:  "VIDEO_BUMPER",
	9:  "VIDEO_TRUE_VIEW_IN_STREAM",
	10: "VIDEO_TRUE_VIEW_IN_DISPLAY",
	11: "VIDEO_NON_SKIPPABLE_IN_STREAM",
	12: "VIDEO_OUTSTREAM",
}
var AdGroupTypeEnum_AdGroupType_value = map[string]int32{
	"UNSPECIFIED":                   0,
	"UNKNOWN":                       1,
	"SEARCH_STANDARD":               2,
	"DISPLAY_STANDARD":              3,
	"SHOPPING_PRODUCT_ADS":          4,
	"HOTEL_ADS":                     6,
	"SHOPPING_SMART_ADS":            7,
	"VIDEO_BUMPER":                  8,
	"VIDEO_TRUE_VIEW_IN_STREAM":     9,
	"VIDEO_TRUE_VIEW_IN_DISPLAY":    10,
	"VIDEO_NON_SKIPPABLE_IN_STREAM": 11,
	"VIDEO_OUTSTREAM":               12,
}

func (x AdGroupTypeEnum_AdGroupType) String() string {
	return proto.EnumName(AdGroupTypeEnum_AdGroupType_name, int32(x))
}
func (AdGroupTypeEnum_AdGroupType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_type_3047fbf325550cc5, []int{0, 0}
}

// Defines types of an ad group, specific to a particular campaign channel
// type. This type drives validations that restrict which entities can be
// added to the ad group.
type AdGroupTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupTypeEnum) Reset()         { *m = AdGroupTypeEnum{} }
func (m *AdGroupTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupTypeEnum) ProtoMessage()    {}
func (*AdGroupTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_type_3047fbf325550cc5, []int{0}
}
func (m *AdGroupTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupTypeEnum.Unmarshal(m, b)
}
func (m *AdGroupTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupTypeEnum.Marshal(b, m, deterministic)
}
func (dst *AdGroupTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupTypeEnum.Merge(dst, src)
}
func (m *AdGroupTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupTypeEnum.Size(m)
}
func (m *AdGroupTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdGroupTypeEnum)(nil), "google.ads.googleads.v0.enums.AdGroupTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.AdGroupTypeEnum_AdGroupType", AdGroupTypeEnum_AdGroupType_name, AdGroupTypeEnum_AdGroupType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/ad_group_type.proto", fileDescriptor_ad_group_type_3047fbf325550cc5)
}

var fileDescriptor_ad_group_type_3047fbf325550cc5 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xa6, 0x19, 0xda, 0xd8, 0xe9, 0x50, 0x2d, 0x33, 0x21, 0x40, 0x2a, 0x52, 0xf7, 0x00, 0x4e,
	0x10, 0x77, 0xe1, 0xca, 0x69, 0x4c, 0x6b, 0xad, 0x75, 0xac, 0x38, 0xc9, 0x04, 0xaa, 0x64, 0x15,
	0x12, 0x45, 0x48, 0x6b, 0x1d, 0xd5, 0xeb, 0xa4, 0xbd, 0x0d, 0xe2, 0x92, 0x6b, 0x9e, 0x82, 0xf7,
	0xe0, 0x86, 0xa7, 0x40, 0x89, 0xbb, 0xd2, 0x0b, 0xd8, 0x4d, 0xf4, 0xe5, 0xfb, 0x39, 0x3a, 0x3e,
	0x1f, 0xbc, 0xa9, 0x8d, 0xa9, 0xaf, 0x2b, 0x7f, 0x59, 0x5a, 0xdf, 0xc1, 0x16, 0xdd, 0x06, 0x7e,
	0xb5, 0xde, 0xae, 0xac, 0xbf, 0x2c, 0x75, 0xbd, 0x31, 0xdb, 0x46, 0xdf, 0xdc, 0x35, 0x15, 0x69,
	0x36, 0xe6, 0xc6, 0xe0, 0xa1, 0xf3, 0x91, 0x65, 0x69, 0xc9, 0x3e, 0x42, 0x6e, 0x03, 0xd2, 0x45,
	0x2e, 0x7e, 0x78, 0x30, 0xa0, 0xe5, 0xa4, 0x4d, 0x65, 0x77, 0x4d, 0xc5, 0xd6, 0xdb, 0xd5, 0xc5,
	0x57, 0x0f, 0xfa, 0x07, 0x1c, 0x1e, 0x40, 0x3f, 0x17, 0x4a, 0xb2, 0x31, 0x7f, 0xcf, 0x59, 0x8c,
	0x1e, 0xe1, 0x3e, 0x9c, 0xe4, 0xe2, 0x52, 0x24, 0x57, 0x02, 0xf5, 0xf0, 0x33, 0x18, 0x28, 0x46,
	0xd3, 0xf1, 0x54, 0xab, 0x8c, 0x8a, 0x98, 0xa6, 0x31, 0xf2, 0xf0, 0x39, 0xa0, 0x98, 0x2b, 0x39,
	0xa3, 0x1f, 0xfe, 0xb2, 0x47, 0xf8, 0x05, 0x9c, 0xab, 0x69, 0x22, 0x25, 0x17, 0x13, 0x2d, 0xd3,
	0x24, 0xce, 0xc7, 0x99, 0xa6, 0xb1, 0x42, 0x8f, 0xf1, 0x53, 0x38, 0x9d, 0x26, 0x19, 0x9b, 0x75,
	0xbf, 0xc7, 0xf8, 0x39, 0xe0, 0xbd, 0x51, 0xcd, 0x69, 0xea, 0x6c, 0x27, 0x18, 0xc1, 0x59, 0xc1,
	0x63, 0x96, 0xe8, 0x28, 0x9f, 0x4b, 0x96, 0xa2, 0x27, 0x78, 0x08, 0x2f, 0x1d, 0x93, 0xa5, 0x39,
	0xd3, 0x05, 0x67, 0x57, 0x9a, 0x0b, 0xad, 0xb2, 0x94, 0xd1, 0x39, 0x3a, 0xc5, 0xaf, 0xe1, 0xd5,
	0x3f, 0xe4, 0xdd, 0x6a, 0x08, 0xf0, 0x08, 0x86, 0x4e, 0x17, 0x89, 0xd0, 0xea, 0x92, 0x4b, 0x49,
	0xa3, 0x19, 0x3b, 0x18, 0xd1, 0x6f, 0xdf, 0xe7, 0x2c, 0x49, 0x9e, 0xed, 0xc8, 0xb3, 0xe8, 0x57,
	0x0f, 0x46, 0x9f, 0xcd, 0x8a, 0x3c, 0x78, 0xdc, 0x08, 0x1d, 0x5c, 0x51, 0xb6, 0x6d, 0xc8, 0xde,
	0xc7, 0x68, 0x17, 0xa9, 0xcd, 0xf5, 0x72, 0x5d, 0x13, 0xb3, 0xa9, 0xfd, 0xba, 0x5a, 0x77, 0x5d,
	0xdd, 0x57, 0xda, 0x7c, 0xb1, 0xff, 0x69, 0xf8, 0x5d, 0xf7, 0xfd, 0xe6, 0x1d, 0x4d, 0x28, 0xfd,
	0xee, 0x0d, 0x27, 0x6e, 0x14, 0x2d, 0x2d, 0x71, 0xb0, 0x45, 0x45, 0x40, 0xda, 0x16, 0xed, 0xcf,
	0x7b, 0x7d, 0x41, 0x4b, 0xbb, 0xd8, 0xeb, 0x8b, 0x22, 0x58, 0x74, 0xfa, 0x6f, 0x6f, 0xe4, 0xc8,
	0x30, 0xa4, 0xa5, 0x0d, 0xc3, 0xbd, 0x23, 0x0c, 0x8b, 0x20, 0x0c, 0x3b, 0xcf, 0xa7, 0xe3, 0x6e,
	0xb1, 0xb7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x63, 0xe5, 0x3b, 0xb8, 0x79, 0x02, 0x00, 0x00,
}
