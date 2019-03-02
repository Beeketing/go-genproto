// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/ad_group_audience_view.proto

package resources // import "github.com/Beeketing/genproto/googleapis/ads/googleads/v0/resources"

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

// An ad group audience view.
// Includes performance data from interests and remarketing lists for Display
// Network and YouTube Network ads, and remarketing lists for search ads (RLSA),
// aggregated at the audience level.
type AdGroupAudienceView struct {
	// The resource name of the ad group audience view.
	// Ad group audience view resource names have the form:
	//
	// `customers/{customer_id}/adGroupAudienceViews/{ad_group_id}_{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupAudienceView) Reset()         { *m = AdGroupAudienceView{} }
func (m *AdGroupAudienceView) String() string { return proto.CompactTextString(m) }
func (*AdGroupAudienceView) ProtoMessage()    {}
func (*AdGroupAudienceView) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_audience_view_dc60046c027f7c13, []int{0}
}
func (m *AdGroupAudienceView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAudienceView.Unmarshal(m, b)
}
func (m *AdGroupAudienceView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAudienceView.Marshal(b, m, deterministic)
}
func (dst *AdGroupAudienceView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAudienceView.Merge(dst, src)
}
func (m *AdGroupAudienceView) XXX_Size() int {
	return xxx_messageInfo_AdGroupAudienceView.Size(m)
}
func (m *AdGroupAudienceView) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAudienceView.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAudienceView proto.InternalMessageInfo

func (m *AdGroupAudienceView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*AdGroupAudienceView)(nil), "google.ads.googleads.v0.resources.AdGroupAudienceView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/ad_group_audience_view.proto", fileDescriptor_ad_group_audience_view_dc60046c027f7c13)
}

var fileDescriptor_ad_group_audience_view_dc60046c027f7c13 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x69, 0x05, 0xc1, 0xa2, 0x97, 0xf5, 0xb2, 0x47, 0x57, 0x59, 0xf0, 0x94, 0x04, 0xbc,
	0x8d, 0x20, 0x64, 0x2f, 0x05, 0x0f, 0xb2, 0xec, 0xa1, 0x07, 0x29, 0x94, 0xd8, 0x0c, 0xa1, 0xb0,
	0xed, 0x94, 0xc4, 0x76, 0x9f, 0xc0, 0x17, 0xf1, 0xe8, 0xa3, 0xf8, 0x28, 0x3e, 0x85, 0x74, 0xd3,
	0xe4, 0x24, 0xee, 0xed, 0xa3, 0xfd, 0xfe, 0x7f, 0x66, 0x92, 0x3d, 0x19, 0x22, 0xb3, 0x47, 0xae,
	0xb4, 0xe3, 0x1e, 0x27, 0x1a, 0x05, 0xb7, 0xe8, 0x68, 0xb0, 0x35, 0x3a, 0xae, 0x74, 0x65, 0x2c,
	0x0d, 0x7d, 0xa5, 0x06, 0xdd, 0x60, 0x57, 0x63, 0x35, 0x36, 0x78, 0x60, 0xbd, 0xa5, 0x77, 0x5a,
	0xac, 0x7c, 0x88, 0x29, 0xed, 0x58, 0xcc, 0xb3, 0x51, 0xb0, 0x98, 0xbf, 0x85, 0xec, 0x5a, 0xea,
	0x7c, 0x6a, 0x90, 0x73, 0x41, 0xd1, 0xe0, 0x61, 0x71, 0x97, 0x5d, 0x05, 0xa7, 0xea, 0x54, 0x8b,
	0xcb, 0xe4, 0x26, 0xb9, 0xbf, 0xd8, 0x5d, 0x86, 0x8f, 0x2f, 0xaa, 0xc5, 0xcd, 0x47, 0x9a, 0xad,
	0x6b, 0x6a, 0xd9, 0xc9, 0x29, 0x9b, 0xe5, 0x1f, 0x33, 0xb6, 0xd3, 0x8a, 0xdb, 0xe4, 0xf5, 0x79,
	0x8e, 0x1b, 0xda, 0xab, 0xce, 0x30, 0xb2, 0x86, 0x1b, 0xec, 0x8e, 0x07, 0x84, 0xa3, 0xfb, 0xc6,
	0xfd, 0xf3, 0x06, 0x8f, 0x91, 0x3e, 0xd3, 0xb3, 0x5c, 0xca, 0xaf, 0x74, 0x95, 0xfb, 0x4a, 0xa9,
	0x1d, 0xf3, 0x38, 0x51, 0x21, 0xd8, 0x2e, 0x98, 0xdf, 0xc1, 0x29, 0xa5, 0x76, 0x65, 0x74, 0xca,
	0x42, 0x94, 0xd1, 0xf9, 0x49, 0xd7, 0xfe, 0x07, 0x80, 0xd4, 0x0e, 0x20, 0x5a, 0x00, 0x85, 0x00,
	0x88, 0xde, 0xdb, 0xf9, 0x71, 0xd9, 0x87, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0x7a, 0xb3,
	0xc2, 0xaf, 0x01, 0x00, 0x00,
}
