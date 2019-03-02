// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/ad_group_feed_error.proto

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

// Enum describing possible ad group feed errors.
type AdGroupFeedErrorEnum_AdGroupFeedError int32

const (
	// Enum unspecified.
	AdGroupFeedErrorEnum_UNSPECIFIED AdGroupFeedErrorEnum_AdGroupFeedError = 0
	// The received error code is not known in this version.
	AdGroupFeedErrorEnum_UNKNOWN AdGroupFeedErrorEnum_AdGroupFeedError = 1
	// An active feed already exists for this ad group and place holder type.
	AdGroupFeedErrorEnum_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE AdGroupFeedErrorEnum_AdGroupFeedError = 2
	// The specified feed is removed.
	AdGroupFeedErrorEnum_CANNOT_CREATE_FOR_REMOVED_FEED AdGroupFeedErrorEnum_AdGroupFeedError = 3
	// The AdGroupFeed already exists. UPDATE operation should be used to modify
	// the existing AdGroupFeed.
	AdGroupFeedErrorEnum_ADGROUP_FEED_ALREADY_EXISTS AdGroupFeedErrorEnum_AdGroupFeedError = 4
	// Cannot operate on removed AdGroupFeed.
	AdGroupFeedErrorEnum_CANNOT_OPERATE_ON_REMOVED_ADGROUP_FEED AdGroupFeedErrorEnum_AdGroupFeedError = 5
	// Invalid placeholder type.
	AdGroupFeedErrorEnum_INVALID_PLACEHOLDER_TYPE AdGroupFeedErrorEnum_AdGroupFeedError = 6
	// Feed mapping for this placeholder type does not exist.
	AdGroupFeedErrorEnum_MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE AdGroupFeedErrorEnum_AdGroupFeedError = 7
	// Location AdGroupFeeds cannot be created unless there is a location
	// CustomerFeed for the specified feed.
	AdGroupFeedErrorEnum_NO_EXISTING_LOCATION_CUSTOMER_FEED AdGroupFeedErrorEnum_AdGroupFeedError = 8
)

var AdGroupFeedErrorEnum_AdGroupFeedError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE",
	3: "CANNOT_CREATE_FOR_REMOVED_FEED",
	4: "ADGROUP_FEED_ALREADY_EXISTS",
	5: "CANNOT_OPERATE_ON_REMOVED_ADGROUP_FEED",
	6: "INVALID_PLACEHOLDER_TYPE",
	7: "MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE",
	8: "NO_EXISTING_LOCATION_CUSTOMER_FEED",
}
var AdGroupFeedErrorEnum_AdGroupFeedError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE": 2,
	"CANNOT_CREATE_FOR_REMOVED_FEED":           3,
	"ADGROUP_FEED_ALREADY_EXISTS":              4,
	"CANNOT_OPERATE_ON_REMOVED_ADGROUP_FEED":   5,
	"INVALID_PLACEHOLDER_TYPE":                 6,
	"MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE": 7,
	"NO_EXISTING_LOCATION_CUSTOMER_FEED":       8,
}

func (x AdGroupFeedErrorEnum_AdGroupFeedError) String() string {
	return proto.EnumName(AdGroupFeedErrorEnum_AdGroupFeedError_name, int32(x))
}
func (AdGroupFeedErrorEnum_AdGroupFeedError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_feed_error_a650797b4dbac82f, []int{0, 0}
}

// Container for enum describing possible ad group feed errors.
type AdGroupFeedErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupFeedErrorEnum) Reset()         { *m = AdGroupFeedErrorEnum{} }
func (m *AdGroupFeedErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupFeedErrorEnum) ProtoMessage()    {}
func (*AdGroupFeedErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_feed_error_a650797b4dbac82f, []int{0}
}
func (m *AdGroupFeedErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupFeedErrorEnum.Unmarshal(m, b)
}
func (m *AdGroupFeedErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupFeedErrorEnum.Marshal(b, m, deterministic)
}
func (dst *AdGroupFeedErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupFeedErrorEnum.Merge(dst, src)
}
func (m *AdGroupFeedErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupFeedErrorEnum.Size(m)
}
func (m *AdGroupFeedErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupFeedErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupFeedErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdGroupFeedErrorEnum)(nil), "google.ads.googleads.v0.errors.AdGroupFeedErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.AdGroupFeedErrorEnum_AdGroupFeedError", AdGroupFeedErrorEnum_AdGroupFeedError_name, AdGroupFeedErrorEnum_AdGroupFeedError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/ad_group_feed_error.proto", fileDescriptor_ad_group_feed_error_a650797b4dbac82f)
}

var fileDescriptor_ad_group_feed_error_a650797b4dbac82f = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x8a, 0xd4, 0x30,
	0x1c, 0xc6, 0x9d, 0xae, 0xee, 0x4a, 0xf6, 0x60, 0x09, 0x0a, 0x82, 0x32, 0x42, 0x0f, 0x8b, 0x88,
	0xa4, 0x05, 0x2f, 0x52, 0x4f, 0xd9, 0xf6, 0x3f, 0xb5, 0xd8, 0x49, 0x42, 0xda, 0xa9, 0xae, 0x0c,
	0x84, 0xd1, 0xd4, 0x22, 0xec, 0x4e, 0x86, 0xc6, 0xdd, 0x07, 0xf2, 0xe8, 0x6b, 0x78, 0xf3, 0x19,
	0x7c, 0x02, 0x4f, 0x3e, 0x82, 0xa4, 0xd9, 0x1d, 0x64, 0x1d, 0xf7, 0xd4, 0x8f, 0x7f, 0x7f, 0xdf,
	0x97, 0x8f, 0xfc, 0x83, 0x5e, 0xf6, 0xc6, 0xf4, 0xa7, 0x5d, 0xbc, 0xd2, 0x36, 0xf6, 0xd2, 0xa9,
	0x8b, 0x24, 0xee, 0x86, 0xc1, 0x0c, 0x36, 0x5e, 0x69, 0xd5, 0x0f, 0xe6, 0x7c, 0xa3, 0x3e, 0x75,
	0x9d, 0x56, 0xe3, 0x90, 0x6c, 0x06, 0xf3, 0xc5, 0xe0, 0xa9, 0xc7, 0xc9, 0x4a, 0x5b, 0xb2, 0x75,
	0x92, 0x8b, 0x84, 0x78, 0x67, 0xf4, 0x33, 0x40, 0xf7, 0xa9, 0x2e, 0x9c, 0x79, 0xd6, 0x75, 0x1a,
	0xdc, 0x14, 0xd6, 0xe7, 0x67, 0xd1, 0xf7, 0x00, 0x85, 0xd7, 0x7f, 0xe0, 0x7b, 0xe8, 0x70, 0xc1,
	0x6a, 0x01, 0x59, 0x39, 0x2b, 0x21, 0x0f, 0x6f, 0xe1, 0x43, 0x74, 0xb0, 0x60, 0x6f, 0x18, 0x7f,
	0xcb, 0xc2, 0x09, 0x7e, 0x8e, 0x9e, 0xce, 0x00, 0x72, 0x45, 0x2b, 0x09, 0x34, 0x3f, 0x51, 0xf0,
	0xae, 0xac, 0x9b, 0x5a, 0xcd, 0xb8, 0x54, 0xa2, 0xa2, 0x19, 0xbc, 0xe6, 0x55, 0x0e, 0x52, 0x35,
	0x27, 0x02, 0xc2, 0x00, 0x47, 0x68, 0x9a, 0x51, 0xc6, 0x78, 0xa3, 0x32, 0x09, 0xb4, 0x81, 0x91,
	0x93, 0x30, 0xe7, 0x2d, 0xe4, 0xca, 0xe5, 0x84, 0x7b, 0xf8, 0x09, 0x7a, 0x44, 0xf3, 0x42, 0xf2,
	0x85, 0x50, 0x3b, 0x92, 0xc3, 0xdb, 0xf8, 0x19, 0x3a, 0xba, 0x0c, 0xe1, 0x02, 0xa4, 0x4b, 0xe1,
	0x6c, 0x1b, 0xf2, 0xb7, 0x35, 0xbc, 0x83, 0x1f, 0xa3, 0x87, 0x25, 0x6b, 0x69, 0x55, 0xe6, 0xff,
	0xd6, 0xd9, 0x77, 0xe5, 0xe7, 0x65, 0x5d, 0x97, 0xac, 0x18, 0xf9, 0x39, 0x15, 0x62, 0xd4, 0xbb,
	0xca, 0x1f, 0xe0, 0x23, 0x14, 0x31, 0xee, 0x6b, 0x38, 0xaa, 0xe2, 0x19, 0x6d, 0x4a, 0xce, 0x54,
	0xb6, 0xa8, 0x1b, 0x3e, 0x07, 0xe9, 0xcf, 0xbc, 0x7b, 0xfc, 0x7b, 0x82, 0xa2, 0x8f, 0xe6, 0x8c,
	0xdc, 0xbc, 0x85, 0xe3, 0x07, 0xd7, 0x6f, 0x5a, 0xb8, 0xe5, 0x89, 0xc9, 0xfb, 0xfc, 0xd2, 0xd8,
	0x9b, 0xd3, 0xd5, 0xba, 0x27, 0x66, 0xe8, 0xe3, 0xbe, 0x5b, 0x8f, 0xab, 0xbd, 0x7a, 0x08, 0x9b,
	0xcf, 0xf6, 0x7f, 0xef, 0xe2, 0x95, 0xff, 0x7c, 0x0d, 0xf6, 0x0a, 0x4a, 0xbf, 0x05, 0xd3, 0xc2,
	0x87, 0x51, 0x6d, 0x89, 0x97, 0x4e, 0xb5, 0x09, 0x19, 0x8f, 0xb4, 0x3f, 0xae, 0x80, 0x25, 0xd5,
	0x76, 0xb9, 0x05, 0x96, 0x6d, 0xb2, 0xf4, 0xc0, 0xaf, 0x20, 0xf2, 0xd3, 0x34, 0xa5, 0xda, 0xa6,
	0xe9, 0x16, 0x49, 0xd3, 0x36, 0x49, 0x53, 0x0f, 0x7d, 0xd8, 0x1f, 0xdb, 0xbd, 0xf8, 0x13, 0x00,
	0x00, 0xff, 0xff, 0x91, 0xc8, 0xfc, 0x23, 0xb4, 0x02, 0x00, 0x00,
}
