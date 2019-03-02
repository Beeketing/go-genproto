// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/firestore/admin/v1/location.proto

package admin // import "github.com/Beeketing/genproto/googleapis/firestore/admin/v1"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/Beeketing/genproto/googleapis/api/annotations"
import _ "github.com/Beeketing/genproto/googleapis/type/latlng"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The metadata message for [google.cloud.location.Location.metadata][google.cloud.location.Location.metadata].
type LocationMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocationMetadata) Reset()         { *m = LocationMetadata{} }
func (m *LocationMetadata) String() string { return proto.CompactTextString(m) }
func (*LocationMetadata) ProtoMessage()    {}
func (*LocationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_location_d6351944632cbafa, []int{0}
}
func (m *LocationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocationMetadata.Unmarshal(m, b)
}
func (m *LocationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocationMetadata.Marshal(b, m, deterministic)
}
func (dst *LocationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationMetadata.Merge(dst, src)
}
func (m *LocationMetadata) XXX_Size() int {
	return xxx_messageInfo_LocationMetadata.Size(m)
}
func (m *LocationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_LocationMetadata proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LocationMetadata)(nil), "google.firestore.admin.v1.LocationMetadata")
}

func init() {
	proto.RegisterFile("google/firestore/admin/v1/location.proto", fileDescriptor_location_d6351944632cbafa)
}

var fileDescriptor_location_d6351944632cbafa = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0xcb, 0x2c, 0x4a, 0x2d, 0x2e, 0xc9, 0x2f, 0x4a, 0xd5, 0x4f, 0x4c, 0xc9,
	0xcd, 0xcc, 0xd3, 0x2f, 0x33, 0xd4, 0xcf, 0xc9, 0x4f, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x84, 0xa8, 0xd4, 0x83, 0xab, 0xd4, 0x03, 0xab, 0xd4, 0x2b,
	0x33, 0x94, 0x92, 0x81, 0x1a, 0x92, 0x58, 0x90, 0xa9, 0x9f, 0x98, 0x97, 0x97, 0x5f, 0x02, 0xd6,
	0x57, 0x0c, 0xd1, 0x28, 0x25, 0x01, 0x95, 0x2d, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0x49, 0x2c, 0xc9,
	0xc9, 0x4b, 0x87, 0xc8, 0x28, 0x09, 0x71, 0x09, 0xf8, 0x40, 0x2d, 0xf1, 0x4d, 0x2d, 0x49, 0x4c,
	0x49, 0x2c, 0x49, 0x74, 0xda, 0xcd, 0xc8, 0x25, 0x9b, 0x9c, 0x9f, 0xab, 0x87, 0xd3, 0x36, 0x27,
	0x5e, 0x98, 0x9e, 0x00, 0x90, 0x21, 0x01, 0x8c, 0x51, 0x76, 0x50, 0xb5, 0xe9, 0xf9, 0x39, 0x89,
	0x79, 0xe9, 0x7a, 0xf9, 0x45, 0xe9, 0xfa, 0xe9, 0xa9, 0x79, 0x60, 0x2b, 0xf4, 0x21, 0x52, 0x89,
	0x05, 0x99, 0xc5, 0x58, 0xbc, 0x68, 0x0d, 0x66, 0x2c, 0x62, 0x62, 0x71, 0x77, 0x76, 0x0b, 0x5e,
	0xc5, 0x24, 0xef, 0x0e, 0x31, 0xc7, 0x39, 0x27, 0xbf, 0x34, 0x45, 0xcf, 0x0d, 0x6e, 0xb3, 0x23,
	0xd8, 0xe6, 0x30, 0xc3, 0x53, 0x30, 0x15, 0x31, 0x60, 0x15, 0x31, 0x70, 0x15, 0x31, 0x60, 0x15,
	0x31, 0x61, 0x86, 0x49, 0x6c, 0x60, 0x5b, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x00, 0xb6,
	0x56, 0xea, 0x57, 0x01, 0x00, 0x00,
}
