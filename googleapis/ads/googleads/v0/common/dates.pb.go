// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/common/dates.proto

package common // import "github.com/Beeketing/genproto/googleapis/ads/googleads/v0/common"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/Beeketing/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A date range.
type DateRange struct {
	// The start date, in yyyy-mm-dd format.
	StartDate *wrappers.StringValue `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// The end date, in yyyy-mm-dd format.
	EndDate              *wrappers.StringValue `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DateRange) Reset()         { *m = DateRange{} }
func (m *DateRange) String() string { return proto.CompactTextString(m) }
func (*DateRange) ProtoMessage()    {}
func (*DateRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_dates_ececfe9883173986, []int{0}
}
func (m *DateRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateRange.Unmarshal(m, b)
}
func (m *DateRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateRange.Marshal(b, m, deterministic)
}
func (dst *DateRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateRange.Merge(dst, src)
}
func (m *DateRange) XXX_Size() int {
	return xxx_messageInfo_DateRange.Size(m)
}
func (m *DateRange) XXX_DiscardUnknown() {
	xxx_messageInfo_DateRange.DiscardUnknown(m)
}

var xxx_messageInfo_DateRange proto.InternalMessageInfo

func (m *DateRange) GetStartDate() *wrappers.StringValue {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *DateRange) GetEndDate() *wrappers.StringValue {
	if m != nil {
		return m.EndDate
	}
	return nil
}

func init() {
	proto.RegisterType((*DateRange)(nil), "google.ads.googleads.v0.common.DateRange")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/common/dates.proto", fileDescriptor_dates_ececfe9883173986)
}

var fileDescriptor_dates_ececfe9883173986 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd0, 0xb1, 0x4a, 0xc4, 0x30,
	0x18, 0x07, 0x70, 0x5a, 0x41, 0xbd, 0xb8, 0x75, 0x12, 0x91, 0x43, 0x3a, 0x89, 0xc3, 0x97, 0xa2,
	0x83, 0x90, 0x9b, 0x7a, 0x1e, 0xdc, 0x7a, 0x9c, 0xd0, 0x41, 0x0a, 0x92, 0xbb, 0xc4, 0x70, 0xd0,
	0x26, 0x25, 0xc9, 0x9d, 0xb3, 0xaf, 0xe2, 0xe8, 0xa3, 0xf8, 0x1e, 0x2e, 0x3e, 0x85, 0x24, 0x5f,
	0xdb, 0x4d, 0x71, 0xea, 0x47, 0xfb, 0xfb, 0xff, 0xd3, 0x7c, 0xe4, 0x46, 0x19, 0xa3, 0x1a, 0x49,
	0xb9, 0x70, 0x14, 0xc7, 0x30, 0x1d, 0x0a, 0xba, 0x35, 0x6d, 0x6b, 0x34, 0x15, 0xdc, 0x4b, 0x07,
	0x9d, 0x35, 0xde, 0x64, 0x53, 0x04, 0xc0, 0x85, 0x83, 0xd1, 0xc2, 0xa1, 0x00, 0xb4, 0x17, 0xfd,
	0x77, 0x1a, 0xf5, 0x66, 0xff, 0x42, 0x5f, 0x2d, 0xef, 0x3a, 0x69, 0xfb, 0x7c, 0xfe, 0x96, 0x90,
	0xc9, 0x82, 0x7b, 0xb9, 0xe6, 0x5a, 0xc9, 0x6c, 0x46, 0x88, 0xf3, 0xdc, 0xfa, 0xe7, 0x70, 0xc4,
	0x79, 0x72, 0x95, 0x5c, 0x9f, 0xdd, 0x5e, 0xf6, 0xbd, 0x30, 0x54, 0xc0, 0xa3, 0xb7, 0x3b, 0xad,
	0x2a, 0xde, 0xec, 0xe5, 0x7a, 0x12, 0x7d, 0x68, 0xc8, 0xee, 0xc9, 0xa9, 0xd4, 0x02, 0xa3, 0xe9,
	0x3f, 0xa2, 0x27, 0x52, 0x8b, 0x10, 0x9c, 0x7f, 0x25, 0x24, 0xdf, 0x9a, 0x16, 0xfe, 0xbe, 0xca,
	0x9c, 0x04, 0xec, 0x56, 0xa1, 0x6a, 0x95, 0x3c, 0x2d, 0x7a, 0xad, 0x4c, 0xc3, 0xb5, 0x02, 0x63,
	0x15, 0x55, 0x52, 0xc7, 0x83, 0x86, 0xa5, 0x75, 0x3b, 0xf7, 0xdb, 0x0e, 0x67, 0xf8, 0x78, 0x4f,
	0x8f, 0x96, 0x65, 0xf9, 0x91, 0x4e, 0x97, 0x58, 0x56, 0x0a, 0x07, 0x38, 0x86, 0xa9, 0x2a, 0xe0,
	0x21, 0xb2, 0xcf, 0x01, 0xd4, 0xa5, 0x70, 0xf5, 0x08, 0xea, 0xaa, 0xa8, 0x11, 0x7c, 0xa7, 0x39,
	0xbe, 0x65, 0xac, 0x14, 0x8e, 0xb1, 0x91, 0x30, 0x56, 0x15, 0x8c, 0x21, 0xda, 0x1c, 0xc7, 0xbf,
	0xbb, 0xfb, 0x09, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xd3, 0xb9, 0xa6, 0xe0, 0x01, 0x00, 0x00,
}
