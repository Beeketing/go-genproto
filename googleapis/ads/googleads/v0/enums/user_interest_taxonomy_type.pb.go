// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/user_interest_taxonomy_type.proto

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

// Enum containing the possible UserInterestTaxonomyTypes.
type UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType int32

const (
	// Not specified.
	UserInterestTaxonomyTypeEnum_UNSPECIFIED UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType = 0
	// Used for return value only. Represents value unknown in this version.
	UserInterestTaxonomyTypeEnum_UNKNOWN UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType = 1
	// The affinity for this user interest.
	UserInterestTaxonomyTypeEnum_AFFINITY UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType = 2
	// The market for this user interest.
	UserInterestTaxonomyTypeEnum_IN_MARKET UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType = 3
	// Users known to have installed applications in the specified categories.
	UserInterestTaxonomyTypeEnum_MOBILE_APP_INSTALL_USER UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType = 4
	// The geographical location of the interest-based vertical.
	UserInterestTaxonomyTypeEnum_VERTICAL_GEO UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType = 5
	// User interest criteria for new smart phone users.
	UserInterestTaxonomyTypeEnum_NEW_SMART_PHONE_USER UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType = 6
)

var UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "AFFINITY",
	3: "IN_MARKET",
	4: "MOBILE_APP_INSTALL_USER",
	5: "VERTICAL_GEO",
	6: "NEW_SMART_PHONE_USER",
}
var UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType_value = map[string]int32{
	"UNSPECIFIED":             0,
	"UNKNOWN":                 1,
	"AFFINITY":                2,
	"IN_MARKET":               3,
	"MOBILE_APP_INSTALL_USER": 4,
	"VERTICAL_GEO":            5,
	"NEW_SMART_PHONE_USER":    6,
}

func (x UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType) String() string {
	return proto.EnumName(UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType_name, int32(x))
}
func (UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_interest_taxonomy_type_20639c3e7f8e9149, []int{0, 0}
}

// Message describing a UserInterestTaxonomyType.
type UserInterestTaxonomyTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInterestTaxonomyTypeEnum) Reset()         { *m = UserInterestTaxonomyTypeEnum{} }
func (m *UserInterestTaxonomyTypeEnum) String() string { return proto.CompactTextString(m) }
func (*UserInterestTaxonomyTypeEnum) ProtoMessage()    {}
func (*UserInterestTaxonomyTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_interest_taxonomy_type_20639c3e7f8e9149, []int{0}
}
func (m *UserInterestTaxonomyTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInterestTaxonomyTypeEnum.Unmarshal(m, b)
}
func (m *UserInterestTaxonomyTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInterestTaxonomyTypeEnum.Marshal(b, m, deterministic)
}
func (dst *UserInterestTaxonomyTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInterestTaxonomyTypeEnum.Merge(dst, src)
}
func (m *UserInterestTaxonomyTypeEnum) XXX_Size() int {
	return xxx_messageInfo_UserInterestTaxonomyTypeEnum.Size(m)
}
func (m *UserInterestTaxonomyTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInterestTaxonomyTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UserInterestTaxonomyTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UserInterestTaxonomyTypeEnum)(nil), "google.ads.googleads.v0.enums.UserInterestTaxonomyTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType", UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType_name, UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/user_interest_taxonomy_type.proto", fileDescriptor_user_interest_taxonomy_type_20639c3e7f8e9149)
}

var fileDescriptor_user_interest_taxonomy_type_20639c3e7f8e9149 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x8e, 0xd3, 0x30,
	0x18, 0xc5, 0x49, 0x06, 0x06, 0xf0, 0x0c, 0x22, 0xb2, 0x90, 0x18, 0x09, 0xba, 0x98, 0x1e, 0xc0,
	0x89, 0xc4, 0xce, 0x2c, 0x90, 0x53, 0xdc, 0x62, 0x35, 0x75, 0xa2, 0xfc, 0xab, 0x40, 0x91, 0xac,
	0x40, 0xac, 0xa8, 0x52, 0x13, 0x47, 0x71, 0x5a, 0xd1, 0xcb, 0xb0, 0x60, 0xc9, 0x25, 0xd8, 0x73,
	0x14, 0x96, 0x9c, 0x00, 0x25, 0x69, 0xbb, 0xeb, 0x6c, 0xac, 0x27, 0xbf, 0xf7, 0xfd, 0xe4, 0xef,
	0x19, 0x7c, 0x28, 0x95, 0x2a, 0xb7, 0xd2, 0xce, 0x0b, 0x6d, 0x8f, 0xb2, 0x57, 0x7b, 0xc7, 0x96,
	0xf5, 0xae, 0xd2, 0xf6, 0x4e, 0xcb, 0x56, 0x6c, 0xea, 0x4e, 0xb6, 0x52, 0x77, 0xa2, 0xcb, 0xbf,
	0xab, 0x5a, 0x55, 0x07, 0xd1, 0x1d, 0x1a, 0x89, 0x9a, 0x56, 0x75, 0x0a, 0x4e, 0xc6, 0x29, 0x94,
	0x17, 0x1a, 0x9d, 0x01, 0x68, 0xef, 0xa0, 0x01, 0x30, 0xfd, 0x6d, 0x80, 0xb7, 0x89, 0x96, 0x2d,
	0x3b, 0x32, 0xe2, 0x23, 0x22, 0x3e, 0x34, 0x92, 0xd6, 0xbb, 0x6a, 0xfa, 0xc3, 0x00, 0x77, 0x97,
	0x02, 0xf0, 0x25, 0xb8, 0x49, 0x78, 0x14, 0xd0, 0x19, 0x9b, 0x33, 0xfa, 0xd1, 0x7a, 0x04, 0x6f,
	0xc0, 0xd3, 0x84, 0x2f, 0xb9, 0xbf, 0xe6, 0x96, 0x01, 0x6f, 0xc1, 0x33, 0x32, 0x9f, 0x33, 0xce,
	0xe2, 0xcf, 0x96, 0x09, 0x5f, 0x80, 0xe7, 0x8c, 0x8b, 0x15, 0x09, 0x97, 0x34, 0xb6, 0xae, 0xe0,
	0x1b, 0xf0, 0x7a, 0xe5, 0xbb, 0xcc, 0xa3, 0x82, 0x04, 0x81, 0x60, 0x3c, 0x8a, 0x89, 0xe7, 0x89,
	0x24, 0xa2, 0xa1, 0xf5, 0x18, 0x5a, 0xe0, 0x36, 0xa5, 0x61, 0xcc, 0x66, 0xc4, 0x13, 0x0b, 0xea,
	0x5b, 0x4f, 0xe0, 0x1d, 0x78, 0xc5, 0xe9, 0x5a, 0x44, 0x2b, 0x12, 0xc6, 0x22, 0xf8, 0xe4, 0x73,
	0x3a, 0x66, 0xaf, 0xdd, 0x7f, 0x06, 0xb8, 0xff, 0xa6, 0x2a, 0xf4, 0xe0, 0x9e, 0xee, 0xe4, 0xd2,
	0x0e, 0x41, 0xdf, 0x52, 0x60, 0x7c, 0x71, 0x8f, 0xf3, 0xa5, 0xda, 0xe6, 0x75, 0x89, 0x54, 0x5b,
	0xda, 0xa5, 0xac, 0x87, 0x0e, 0x4f, 0xc5, 0x37, 0x1b, 0x7d, 0xe1, 0x1f, 0xde, 0x0f, 0xe7, 0x4f,
	0xf3, 0x6a, 0x41, 0xc8, 0x2f, 0x73, 0xb2, 0x18, 0x51, 0xa4, 0xd0, 0x68, 0x94, 0xbd, 0x4a, 0x1d,
	0xd4, 0x17, 0xaa, 0xff, 0x9c, 0xfc, 0x8c, 0x14, 0x3a, 0x3b, 0xfb, 0x59, 0xea, 0x64, 0x83, 0xff,
	0xd7, 0xbc, 0x1f, 0x2f, 0x31, 0x26, 0x85, 0xc6, 0xf8, 0x9c, 0xc0, 0x38, 0x75, 0x30, 0x1e, 0x32,
	0x5f, 0xaf, 0x87, 0x87, 0xbd, 0xfb, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x88, 0x14, 0xb0, 0x76, 0x1f,
	0x02, 0x00, 0x00,
}
