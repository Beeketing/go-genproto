// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/google_ads_field_category.proto

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

// The category of the artifact.
type GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory int32

const (
	// Unspecified
	GoogleAdsFieldCategoryEnum_UNSPECIFIED GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory = 0
	// Unknown
	GoogleAdsFieldCategoryEnum_UNKNOWN GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory = 1
	// The described artifact is a resource.
	GoogleAdsFieldCategoryEnum_RESOURCE GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory = 2
	// The described artifact is a field and is an attribute of a resource.
	// Including a resource attribute field in a query may segment the query if
	// the resource to which it is attributed segments the resource found in
	// the FROM clause.
	GoogleAdsFieldCategoryEnum_ATTRIBUTE GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory = 3
	// The described artifact is a field and always segments search queries.
	GoogleAdsFieldCategoryEnum_SEGMENT GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory = 5
	// The described artifact is a field and is a metric. It never segments
	// search queries.
	GoogleAdsFieldCategoryEnum_METRIC GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory = 6
)

var GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "RESOURCE",
	3: "ATTRIBUTE",
	5: "SEGMENT",
	6: "METRIC",
}
var GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"RESOURCE":    2,
	"ATTRIBUTE":   3,
	"SEGMENT":     5,
	"METRIC":      6,
}

func (x GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory) String() string {
	return proto.EnumName(GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory_name, int32(x))
}
func (GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_google_ads_field_category_c7f8cd502030f943, []int{0, 0}
}

// Container for enum that determines if the described artifact is a resource
// or a field, and if it is a field, when it segments search queries.
type GoogleAdsFieldCategoryEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoogleAdsFieldCategoryEnum) Reset()         { *m = GoogleAdsFieldCategoryEnum{} }
func (m *GoogleAdsFieldCategoryEnum) String() string { return proto.CompactTextString(m) }
func (*GoogleAdsFieldCategoryEnum) ProtoMessage()    {}
func (*GoogleAdsFieldCategoryEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_google_ads_field_category_c7f8cd502030f943, []int{0}
}
func (m *GoogleAdsFieldCategoryEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoogleAdsFieldCategoryEnum.Unmarshal(m, b)
}
func (m *GoogleAdsFieldCategoryEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoogleAdsFieldCategoryEnum.Marshal(b, m, deterministic)
}
func (dst *GoogleAdsFieldCategoryEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoogleAdsFieldCategoryEnum.Merge(dst, src)
}
func (m *GoogleAdsFieldCategoryEnum) XXX_Size() int {
	return xxx_messageInfo_GoogleAdsFieldCategoryEnum.Size(m)
}
func (m *GoogleAdsFieldCategoryEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_GoogleAdsFieldCategoryEnum.DiscardUnknown(m)
}

var xxx_messageInfo_GoogleAdsFieldCategoryEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GoogleAdsFieldCategoryEnum)(nil), "google.ads.googleads.v0.enums.GoogleAdsFieldCategoryEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory", GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory_name, GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/google_ads_field_category.proto", fileDescriptor_google_ads_field_category_c7f8cd502030f943)
}

var fileDescriptor_google_ads_field_category_c7f8cd502030f943 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x18, 0xb7, 0x1d, 0x4e, 0xcd, 0x14, 0x43, 0x0e, 0x1e, 0x94, 0x1d, 0xb6, 0x07, 0x48, 0x0b, 0xde,
	0x22, 0x1e, 0xd2, 0x9a, 0x8d, 0x22, 0xeb, 0x46, 0xd7, 0x4e, 0x90, 0xc2, 0xa8, 0x4b, 0x0d, 0x83,
	0xae, 0x19, 0xcd, 0x36, 0xf0, 0x15, 0x7c, 0x0c, 0x8f, 0x3e, 0x8a, 0x8f, 0xe2, 0xc9, 0x47, 0x90,
	0xa6, 0x6b, 0x4f, 0xd3, 0x4b, 0xf8, 0xf8, 0xfd, 0xc9, 0xf7, 0xfb, 0x7e, 0xe0, 0x5e, 0x48, 0x29,
	0xb2, 0xd4, 0x4a, 0xb8, 0xb2, 0xaa, 0xb1, 0x9c, 0x76, 0xb6, 0x95, 0xe6, 0xdb, 0x55, 0x0d, 0xcd,
	0x13, 0xae, 0xe6, 0xaf, 0xcb, 0x34, 0xe3, 0xf3, 0x45, 0xb2, 0x49, 0x85, 0x2c, 0xde, 0xf0, 0xba,
	0x90, 0x1b, 0x89, 0xba, 0x95, 0x00, 0x27, 0x5c, 0xe1, 0xc6, 0x8e, 0x77, 0x36, 0xd6, 0xf6, 0xfe,
	0xbb, 0x01, 0xae, 0x87, 0x1a, 0xa6, 0x5c, 0x0d, 0xca, 0x0f, 0xdc, 0xbd, 0x9f, 0xe5, 0xdb, 0x55,
	0x3f, 0x03, 0x57, 0x87, 0x59, 0x74, 0x09, 0x3a, 0x91, 0x3f, 0x9d, 0x30, 0xd7, 0x1b, 0x78, 0xec,
	0x01, 0x1e, 0xa1, 0x0e, 0x38, 0x89, 0xfc, 0x47, 0x7f, 0xfc, 0xe4, 0x43, 0x03, 0x9d, 0x83, 0xd3,
	0x80, 0x4d, 0xc7, 0x51, 0xe0, 0x32, 0x68, 0xa2, 0x0b, 0x70, 0x46, 0xc3, 0x30, 0xf0, 0x9c, 0x28,
	0x64, 0xb0, 0x55, 0x2a, 0xa7, 0x6c, 0x38, 0x62, 0x7e, 0x08, 0x8f, 0x11, 0x00, 0xed, 0x11, 0x0b,
	0x03, 0xcf, 0x85, 0x6d, 0xe7, 0xc7, 0x00, 0xbd, 0x85, 0x5c, 0xe1, 0x7f, 0x23, 0x3b, 0x37, 0x87,
	0x13, 0x4d, 0xca, 0x73, 0x27, 0xc6, 0xb3, 0xb3, 0x77, 0x0b, 0x99, 0x25, 0xb9, 0xc0, 0xb2, 0x10,
	0x96, 0x48, 0x73, 0x5d, 0x46, 0xdd, 0xdf, 0x7a, 0xa9, 0xfe, 0xa8, 0xf3, 0x4e, 0xbf, 0x1f, 0x66,
	0x6b, 0x48, 0xe9, 0xa7, 0xd9, 0xad, 0x36, 0x61, 0xca, 0x15, 0x6e, 0x96, 0xe2, 0x99, 0x8d, 0xcb,
	0x6e, 0xd4, 0x57, 0xcd, 0xc7, 0x94, 0xab, 0xb8, 0xe1, 0xe3, 0x99, 0x1d, 0x6b, 0xfe, 0xdb, 0xec,
	0x55, 0x20, 0x21, 0x94, 0x2b, 0x42, 0x1a, 0x05, 0x21, 0x33, 0x9b, 0x10, 0xad, 0x79, 0x69, 0xeb,
	0x60, 0xb7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x62, 0x46, 0x30, 0xe6, 0x01, 0x00, 0x00,
}
