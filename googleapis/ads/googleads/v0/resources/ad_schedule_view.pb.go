// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/ad_schedule_view.proto

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

// An ad schedule view summarizes the performance of campaigns by
// AdSchedule criteria.
type AdScheduleView struct {
	// The resource name of the ad schedule view.
	// AdSchedule view resource names have the form:
	//
	// `customers/{customer_id}/adScheduleViews/{campaign_id}_{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdScheduleView) Reset()         { *m = AdScheduleView{} }
func (m *AdScheduleView) String() string { return proto.CompactTextString(m) }
func (*AdScheduleView) ProtoMessage()    {}
func (*AdScheduleView) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_schedule_view_4605c4ad10395f1c, []int{0}
}
func (m *AdScheduleView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdScheduleView.Unmarshal(m, b)
}
func (m *AdScheduleView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdScheduleView.Marshal(b, m, deterministic)
}
func (dst *AdScheduleView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdScheduleView.Merge(dst, src)
}
func (m *AdScheduleView) XXX_Size() int {
	return xxx_messageInfo_AdScheduleView.Size(m)
}
func (m *AdScheduleView) XXX_DiscardUnknown() {
	xxx_messageInfo_AdScheduleView.DiscardUnknown(m)
}

var xxx_messageInfo_AdScheduleView proto.InternalMessageInfo

func (m *AdScheduleView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*AdScheduleView)(nil), "google.ads.googleads.v0.resources.AdScheduleView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/ad_schedule_view.proto", fileDescriptor_ad_schedule_view_4605c4ad10395f1c)
}

var fileDescriptor_ad_schedule_view_4605c4ad10395f1c = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x48, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xa2, 0xd4,
	0xe2, 0xfc, 0xd2, 0xa2, 0xe4, 0xd4, 0x62, 0xfd, 0xc4, 0x94, 0xf8, 0xe2, 0xe4, 0x8c, 0xd4, 0x94,
	0xd2, 0x9c, 0xd4, 0xf8, 0xb2, 0xcc, 0xd4, 0x72, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x45,
	0x88, 0x72, 0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x4e, 0xbd, 0x32, 0x03, 0x3d, 0xb8, 0x4e, 0x25,
	0x53, 0x2e, 0x3e, 0xc7, 0x94, 0x60, 0xa8, 0xde, 0xb0, 0xcc, 0xd4, 0x72, 0x21, 0x65, 0x2e, 0x5e,
	0x98, 0x74, 0x7c, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x0f, 0x4c,
	0xd0, 0x2f, 0x31, 0x37, 0xd5, 0xa9, 0x81, 0x89, 0x4b, 0x35, 0x39, 0x3f, 0x57, 0x8f, 0xa0, 0x05,
	0x4e, 0xc2, 0xa8, 0xc6, 0x07, 0x80, 0x1c, 0x16, 0xc0, 0x18, 0xe5, 0x05, 0xd5, 0x99, 0x9e, 0x9f,
	0x93, 0x98, 0x97, 0xae, 0x97, 0x5f, 0x94, 0xae, 0x9f, 0x9e, 0x9a, 0x07, 0x76, 0x36, 0xcc, 0x93,
	0x05, 0x99, 0xc5, 0x78, 0xfc, 0x6c, 0x0d, 0x67, 0x2d, 0x62, 0x62, 0x76, 0x77, 0x74, 0x5c, 0xc5,
	0xa4, 0xe8, 0x0e, 0x31, 0xd2, 0x31, 0xa5, 0x58, 0x0f, 0xc2, 0x04, 0xb1, 0xc2, 0x0c, 0xf4, 0x82,
	0x60, 0x2a, 0x4f, 0xc1, 0xd4, 0xc4, 0x38, 0xa6, 0x14, 0xc7, 0xc0, 0xd5, 0xc4, 0x84, 0x19, 0xc4,
	0xc0, 0xd5, 0xbc, 0x62, 0x52, 0x85, 0x48, 0x58, 0x59, 0x39, 0xa6, 0x14, 0x5b, 0x59, 0xc1, 0x55,
	0x59, 0x59, 0x85, 0x19, 0x58, 0x59, 0xc1, 0xd5, 0x25, 0xb1, 0x81, 0x1d, 0x6b, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x1f, 0x2f, 0x68, 0x08, 0x9f, 0x01, 0x00, 0x00,
}
