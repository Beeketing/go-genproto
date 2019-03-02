// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/search_term_view.proto

package resources // import "github.com/Beeketing/go-genproto/googleapis/ads/googleads/v0/resources"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/Beeketing/protobuf/ptypes/wrappers"
import enums "github.com/Beeketing/go-genproto/googleapis/ads/googleads/v0/enums"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A search term view with metrics aggregated by search term at the ad group
// level.
type SearchTermView struct {
	// The resource name of the search term view.
	// Search term view resource names have the form:
	//
	// `customers/{customer_id}/searchTermViews/{campaign_id}_{ad_group_id}_
	// {URL-base64 search term}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The search term.
	SearchTerm *wrappers.StringValue `protobuf:"bytes,2,opt,name=search_term,json=searchTerm,proto3" json:"search_term,omitempty"`
	// The ad group the search term served in.
	AdGroup *wrappers.StringValue `protobuf:"bytes,3,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// Indicates whether the search term is currently one of your
	// targeted or excluded keywords.
	Status               enums.SearchTermTargetingStatusEnum_SearchTermTargetingStatus `protobuf:"varint,4,opt,name=status,proto3,enum=google.ads.googleads.v0.enums.SearchTermTargetingStatusEnum_SearchTermTargetingStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                      `json:"-"`
	XXX_unrecognized     []byte                                                        `json:"-"`
	XXX_sizecache        int32                                                         `json:"-"`
}

func (m *SearchTermView) Reset()         { *m = SearchTermView{} }
func (m *SearchTermView) String() string { return proto.CompactTextString(m) }
func (*SearchTermView) ProtoMessage()    {}
func (*SearchTermView) Descriptor() ([]byte, []int) {
	return fileDescriptor_search_term_view_3cb0ee904f1d648d, []int{0}
}
func (m *SearchTermView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchTermView.Unmarshal(m, b)
}
func (m *SearchTermView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchTermView.Marshal(b, m, deterministic)
}
func (dst *SearchTermView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchTermView.Merge(dst, src)
}
func (m *SearchTermView) XXX_Size() int {
	return xxx_messageInfo_SearchTermView.Size(m)
}
func (m *SearchTermView) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchTermView.DiscardUnknown(m)
}

var xxx_messageInfo_SearchTermView proto.InternalMessageInfo

func (m *SearchTermView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *SearchTermView) GetSearchTerm() *wrappers.StringValue {
	if m != nil {
		return m.SearchTerm
	}
	return nil
}

func (m *SearchTermView) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *SearchTermView) GetStatus() enums.SearchTermTargetingStatusEnum_SearchTermTargetingStatus {
	if m != nil {
		return m.Status
	}
	return enums.SearchTermTargetingStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*SearchTermView)(nil), "google.ads.googleads.v0.resources.SearchTermView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/search_term_view.proto", fileDescriptor_search_term_view_3cb0ee904f1d648d)
}

var fileDescriptor_search_term_view_3cb0ee904f1d648d = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x6b, 0xdb, 0x30,
	0x18, 0xc6, 0xce, 0xc8, 0x36, 0x65, 0xcb, 0xc1, 0xbb, 0x98, 0x30, 0x46, 0xb2, 0x11, 0xc8, 0x49,
	0x36, 0xd9, 0x61, 0x43, 0xa5, 0x50, 0x07, 0x4a, 0xa0, 0x87, 0x12, 0x9c, 0xe0, 0x43, 0x31, 0x18,
	0x25, 0x7e, 0xab, 0x1a, 0x62, 0xc9, 0x48, 0x76, 0x72, 0xed, 0x1f, 0xe9, 0xa5, 0xc7, 0xfe, 0x94,
	0xfe, 0x94, 0xfe, 0x8a, 0xe2, 0x2f, 0xb5, 0x39, 0xa4, 0xed, 0xed, 0xb1, 0xdf, 0xe7, 0x43, 0xcf,
	0x2b, 0xa1, 0xff, 0x4c, 0x08, 0xb6, 0x05, 0x87, 0xc6, 0xca, 0xa9, 0x61, 0x89, 0x76, 0xae, 0x23,
	0x41, 0x89, 0x42, 0x6e, 0x40, 0x39, 0x0a, 0xa8, 0xdc, 0xdc, 0x44, 0x39, 0xc8, 0x34, 0xda, 0x25,
	0xb0, 0xc7, 0x99, 0x14, 0xb9, 0xb0, 0x46, 0x35, 0x1d, 0xd3, 0x58, 0x61, 0xad, 0xc4, 0x3b, 0x17,
	0x6b, 0xe5, 0xe0, 0xec, 0x98, 0x39, 0xf0, 0x22, 0x3d, 0x34, 0xce, 0xa9, 0x64, 0x90, 0x27, 0x9c,
	0x45, 0x2a, 0xa7, 0x79, 0xa1, 0xea, 0x90, 0xc1, 0xaf, 0xc6, 0xa1, 0xfa, 0x5a, 0x17, 0xd7, 0xce,
	0x5e, 0xd2, 0x2c, 0x03, 0xd9, 0xcc, 0x7f, 0xdf, 0x99, 0xa8, 0xbf, 0xac, 0x6c, 0x56, 0x20, 0xd3,
	0x20, 0x81, 0xbd, 0xf5, 0x07, 0x7d, 0x6f, 0x4f, 0x10, 0x71, 0x9a, 0x82, 0x6d, 0x0c, 0x8d, 0xc9,
	0x57, 0xff, 0x5b, 0xfb, 0xf3, 0x92, 0xa6, 0x60, 0x9d, 0xa2, 0xde, 0xab, 0x74, 0xdb, 0x1c, 0x1a,
	0x93, 0xde, 0xf4, 0x67, 0xd3, 0x03, 0xb7, 0x69, 0x78, 0x99, 0xcb, 0x84, 0xb3, 0x80, 0x6e, 0x0b,
	0xf0, 0x91, 0xd2, 0x39, 0xd6, 0x3f, 0xf4, 0x85, 0xc6, 0x11, 0x93, 0xa2, 0xc8, 0xec, 0xce, 0x07,
	0xb4, 0x9f, 0x69, 0x3c, 0x2f, 0xc9, 0x16, 0x47, 0xdd, 0xba, 0x9f, 0xfd, 0x69, 0x68, 0x4c, 0xfa,
	0xd3, 0x00, 0x1f, 0xdb, 0x62, 0xb5, 0x22, 0xfc, 0xd2, 0x6d, 0xd5, 0x2e, 0x68, 0x59, 0xe9, 0xcf,
	0x79, 0x91, 0x1e, 0x9f, 0xfa, 0x4d, 0xca, 0xec, 0xd6, 0x44, 0xe3, 0x8d, 0x48, 0xf1, 0xbb, 0x77,
	0x35, 0xfb, 0x71, 0xb8, 0xc6, 0x45, 0x59, 0x63, 0x61, 0x5c, 0x5d, 0x34, 0x4a, 0x26, 0xb6, 0x94,
	0x33, 0x2c, 0x24, 0x73, 0x18, 0xf0, 0xaa, 0x64, 0x7b, 0xa5, 0x59, 0xa2, 0xde, 0x78, 0x3e, 0x27,
	0x1a, 0xdd, 0x9b, 0x9d, 0xb9, 0xe7, 0x3d, 0x98, 0xa3, 0x79, 0x6d, 0xe9, 0xc5, 0x0a, 0xd7, 0xb0,
	0x44, 0x81, 0x8b, 0xfd, 0x96, 0xf9, 0xd8, 0x72, 0x42, 0x2f, 0x56, 0xa1, 0xe6, 0x84, 0x81, 0x1b,
	0x6a, 0xce, 0x93, 0x39, 0xae, 0x07, 0x84, 0x78, 0xb1, 0x22, 0x44, 0xb3, 0x08, 0x09, 0x5c, 0x42,
	0x34, 0x6f, 0xdd, 0xad, 0x0e, 0xfb, 0xf7, 0x39, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x27, 0x40, 0x59,
	0xea, 0x02, 0x00, 0x00,
}
