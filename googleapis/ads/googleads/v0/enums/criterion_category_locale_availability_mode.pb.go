// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/criterion_category_locale_availability_mode.proto

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

// Enum containing the possible CriterionCategoryLocaleAvailabilityMode.
type CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode int32

const (
	// Not specified.
	CriterionCategoryLocaleAvailabilityModeEnum_UNSPECIFIED CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode = 0
	// Used for return value only. Represents value unknown in this version.
	CriterionCategoryLocaleAvailabilityModeEnum_UNKNOWN CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode = 1
	// The category is available to campaigns of all locales.
	CriterionCategoryLocaleAvailabilityModeEnum_ALL_LOCALES CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode = 2
	// The category is available to campaigns within a list of countries,
	// regardless of language.
	CriterionCategoryLocaleAvailabilityModeEnum_COUNTRY_AND_ALL_LANGUAGES CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode = 3
	// The category is available to campaigns within a list of languages,
	// regardless of country.
	CriterionCategoryLocaleAvailabilityModeEnum_LANGUAGE_AND_ALL_COUNTRIES CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode = 4
	// The category is available to campaigns within a list of country, language
	// pairs.
	CriterionCategoryLocaleAvailabilityModeEnum_COUNTRY_AND_LANGUAGE CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode = 5
)

var CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ALL_LOCALES",
	3: "COUNTRY_AND_ALL_LANGUAGES",
	4: "LANGUAGE_AND_ALL_COUNTRIES",
	5: "COUNTRY_AND_LANGUAGE",
}
var CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode_value = map[string]int32{
	"UNSPECIFIED":                0,
	"UNKNOWN":                    1,
	"ALL_LOCALES":                2,
	"COUNTRY_AND_ALL_LANGUAGES":  3,
	"LANGUAGE_AND_ALL_COUNTRIES": 4,
	"COUNTRY_AND_LANGUAGE":       5,
}

func (x CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode) String() string {
	return proto.EnumName(CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode_name, int32(x))
}
func (CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_criterion_category_locale_availability_mode_a48a18156046d7ad, []int{0, 0}
}

// Describes locale availabilty mode for a criterion availability - whether
// it's available globally, or a particular country with all languages, or a
// particular language with all countries, or a country-language pair.
type CriterionCategoryLocaleAvailabilityModeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CriterionCategoryLocaleAvailabilityModeEnum) Reset() {
	*m = CriterionCategoryLocaleAvailabilityModeEnum{}
}
func (m *CriterionCategoryLocaleAvailabilityModeEnum) String() string {
	return proto.CompactTextString(m)
}
func (*CriterionCategoryLocaleAvailabilityModeEnum) ProtoMessage() {}
func (*CriterionCategoryLocaleAvailabilityModeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_criterion_category_locale_availability_mode_a48a18156046d7ad, []int{0}
}
func (m *CriterionCategoryLocaleAvailabilityModeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriterionCategoryLocaleAvailabilityModeEnum.Unmarshal(m, b)
}
func (m *CriterionCategoryLocaleAvailabilityModeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriterionCategoryLocaleAvailabilityModeEnum.Marshal(b, m, deterministic)
}
func (dst *CriterionCategoryLocaleAvailabilityModeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriterionCategoryLocaleAvailabilityModeEnum.Merge(dst, src)
}
func (m *CriterionCategoryLocaleAvailabilityModeEnum) XXX_Size() int {
	return xxx_messageInfo_CriterionCategoryLocaleAvailabilityModeEnum.Size(m)
}
func (m *CriterionCategoryLocaleAvailabilityModeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CriterionCategoryLocaleAvailabilityModeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CriterionCategoryLocaleAvailabilityModeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CriterionCategoryLocaleAvailabilityModeEnum)(nil), "google.ads.googleads.v0.enums.CriterionCategoryLocaleAvailabilityModeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode", CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode_name, CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/criterion_category_locale_availability_mode.proto", fileDescriptor_criterion_category_locale_availability_mode_a48a18156046d7ad)
}

var fileDescriptor_criterion_category_locale_availability_mode_a48a18156046d7ad = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xf3, 0x30,
	0x1c, 0xc6, 0xdf, 0x76, 0xef, 0xfb, 0x0a, 0xd9, 0xc1, 0x52, 0x3c, 0xa8, 0x30, 0x61, 0xbb, 0x78,
	0x50, 0xd2, 0x82, 0xb7, 0x78, 0xca, 0xba, 0x5a, 0x86, 0xb5, 0x1d, 0xd6, 0x4e, 0x94, 0x42, 0xc9,
	0xda, 0x50, 0x0a, 0x6d, 0x33, 0x9a, 0x6e, 0xb0, 0xa3, 0x5f, 0xc5, 0xa3, 0x37, 0xbf, 0x86, 0xdf,
	0xc2, 0xab, 0x9f, 0x42, 0x9a, 0xae, 0x65, 0x17, 0x65, 0x97, 0xf0, 0x24, 0xcf, 0x93, 0x1f, 0xc9,
	0xf3, 0x07, 0x6e, 0xc2, 0x58, 0x92, 0x51, 0x8d, 0xc4, 0x5c, 0x6b, 0x64, 0xad, 0xd6, 0xba, 0x46,
	0x8b, 0x55, 0xce, 0xb5, 0xa8, 0x4c, 0x2b, 0x5a, 0xa6, 0xac, 0x08, 0x23, 0x52, 0xd1, 0x84, 0x95,
	0x9b, 0x30, 0x63, 0x11, 0xc9, 0x68, 0x48, 0xd6, 0x24, 0xcd, 0xc8, 0x22, 0xcd, 0xd2, 0x6a, 0x13,
	0xe6, 0x2c, 0xa6, 0x70, 0x59, 0xb2, 0x8a, 0xa9, 0x83, 0x86, 0x02, 0x49, 0xcc, 0x61, 0x07, 0x84,
	0x6b, 0x1d, 0x0a, 0xe0, 0xe8, 0x53, 0x02, 0x17, 0x46, 0x0b, 0x35, 0xb6, 0x4c, 0x5b, 0x20, 0xf1,
	0x0e, 0xf1, 0x8e, 0xc5, 0xd4, 0x2c, 0x56, 0xf9, 0xe8, 0x5d, 0x02, 0xe7, 0x7b, 0xe6, 0xd5, 0x43,
	0xd0, 0xf7, 0x1d, 0x6f, 0x66, 0x1a, 0xd3, 0x9b, 0xa9, 0x39, 0x51, 0xfe, 0xa8, 0x7d, 0x70, 0xe0,
	0x3b, 0xb7, 0x8e, 0xfb, 0xe8, 0x28, 0x52, 0xed, 0x62, 0xdb, 0x0e, 0x6d, 0xd7, 0xc0, 0xb6, 0xe9,
	0x29, 0xb2, 0x3a, 0x00, 0x27, 0x86, 0xeb, 0x3b, 0x0f, 0xf7, 0x4f, 0x21, 0x76, 0x26, 0xa1, 0x30,
	0xb1, 0x63, 0xf9, 0xd8, 0x32, 0x3d, 0xa5, 0xa7, 0x9e, 0x81, 0xd3, 0x76, 0xdb, 0xf9, 0x4d, 0x7e,
	0x6a, 0x7a, 0xca, 0x5f, 0xf5, 0x18, 0x1c, 0xed, 0x5e, 0x6f, 0xb3, 0xca, 0xbf, 0xf1, 0x8b, 0x0c,
	0x86, 0x11, 0xcb, 0xe1, 0xaf, 0x4d, 0x8c, 0x2f, 0xf7, 0xfc, 0xd6, 0xac, 0xae, 0x75, 0x26, 0x3d,
	0x8f, 0xb7, 0xb8, 0x84, 0x65, 0xa4, 0x48, 0x20, 0x2b, 0x13, 0x2d, 0xa1, 0x85, 0x28, 0xbd, 0x9d,
	0xdc, 0x32, 0xe5, 0x3f, 0x0c, 0xf2, 0x5a, 0xac, 0xaf, 0x72, 0xcf, 0xc2, 0xf8, 0x4d, 0x1e, 0x58,
	0x0d, 0x0a, 0xc7, 0x1c, 0x36, 0xb2, 0x56, 0x73, 0x1d, 0xd6, 0x95, 0xf3, 0x8f, 0xd6, 0x0f, 0x70,
	0xcc, 0x83, 0xce, 0x0f, 0xe6, 0x7a, 0x20, 0xfc, 0x2f, 0x79, 0xd8, 0x1c, 0x22, 0x84, 0x63, 0x8e,
	0x50, 0x97, 0x40, 0x68, 0xae, 0x23, 0x24, 0x32, 0x8b, 0xff, 0xe2, 0x61, 0x57, 0xdf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xc8, 0x4d, 0xb4, 0x90, 0x60, 0x02, 0x00, 0x00,
}
