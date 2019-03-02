// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/device.proto

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

// Enumerates Google Ads devices available for targeting.
type DeviceEnum_Device int32

const (
	// Not specified.
	DeviceEnum_UNSPECIFIED DeviceEnum_Device = 0
	// The value is unknown in this version.
	DeviceEnum_UNKNOWN DeviceEnum_Device = 1
	// Mobile devices with full browsers.
	DeviceEnum_MOBILE DeviceEnum_Device = 2
	// Tablets with full browsers.
	DeviceEnum_TABLET DeviceEnum_Device = 3
	// Computers.
	DeviceEnum_DESKTOP DeviceEnum_Device = 4
)

var DeviceEnum_Device_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "MOBILE",
	3: "TABLET",
	4: "DESKTOP",
}
var DeviceEnum_Device_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"MOBILE":      2,
	"TABLET":      3,
	"DESKTOP":     4,
}

func (x DeviceEnum_Device) String() string {
	return proto.EnumName(DeviceEnum_Device_name, int32(x))
}
func (DeviceEnum_Device) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_device_6d726a8a981312d4, []int{0, 0}
}

// Container for enumeration of Google Ads devices available for targeting.
type DeviceEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceEnum) Reset()         { *m = DeviceEnum{} }
func (m *DeviceEnum) String() string { return proto.CompactTextString(m) }
func (*DeviceEnum) ProtoMessage()    {}
func (*DeviceEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_6d726a8a981312d4, []int{0}
}
func (m *DeviceEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceEnum.Unmarshal(m, b)
}
func (m *DeviceEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceEnum.Marshal(b, m, deterministic)
}
func (dst *DeviceEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceEnum.Merge(dst, src)
}
func (m *DeviceEnum) XXX_Size() int {
	return xxx_messageInfo_DeviceEnum.Size(m)
}
func (m *DeviceEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceEnum.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DeviceEnum)(nil), "google.ads.googleads.v0.enums.DeviceEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.DeviceEnum_Device", DeviceEnum_Device_name, DeviceEnum_Device_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/device.proto", fileDescriptor_device_6d726a8a981312d4)
}

var fileDescriptor_device_6d726a8a981312d4 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xc1, 0x4e, 0x83, 0x30,
	0x00, 0x86, 0x1d, 0x33, 0x98, 0x94, 0x83, 0x84, 0xfb, 0x0e, 0xdb, 0xd1, 0x43, 0x69, 0xe2, 0xad,
	0x9e, 0x5a, 0xa9, 0x0b, 0x61, 0x02, 0xc9, 0x18, 0x46, 0xc3, 0x05, 0xd7, 0xa6, 0x59, 0x32, 0xe8,
	0xb2, 0x3a, 0x1e, 0xc8, 0xa3, 0x8f, 0xe2, 0x53, 0x78, 0xf6, 0x29, 0x4c, 0x5b, 0xc7, 0x4d, 0x2f,
	0xe4, 0xa3, 0xff, 0x97, 0xf6, 0xff, 0xc1, 0x8d, 0x54, 0x4a, 0xee, 0x45, 0xdc, 0x72, 0x1d, 0x3b,
	0x34, 0x34, 0xa0, 0x58, 0xf4, 0xa7, 0x4e, 0xc7, 0x5c, 0x0c, 0xbb, 0xad, 0x80, 0x87, 0xa3, 0x7a,
	0x53, 0xd1, 0xcc, 0x09, 0xb0, 0xe5, 0x1a, 0x8e, 0x2e, 0x1c, 0x10, 0xb4, 0xee, 0xe2, 0x19, 0x80,
	0xc4, 0xea, 0xac, 0x3f, 0x75, 0x8b, 0x0c, 0xf8, 0xee, 0x2f, 0xba, 0x06, 0xc1, 0x26, 0x5f, 0x97,
	0xec, 0x3e, 0x7d, 0x48, 0x59, 0x12, 0x5e, 0x44, 0x01, 0xb8, 0xda, 0xe4, 0x59, 0x5e, 0x3c, 0xe5,
	0xe1, 0x24, 0x02, 0xc0, 0x7f, 0x2c, 0x68, 0xba, 0x62, 0xa1, 0x67, 0xb8, 0x22, 0x74, 0xc5, 0xaa,
	0x70, 0x6a, 0xa4, 0x84, 0xad, 0xb3, 0xaa, 0x28, 0xc3, 0x4b, 0xfa, 0x35, 0x01, 0xf3, 0xad, 0xea,
	0xe0, 0xbf, 0x05, 0x68, 0xe0, 0x1e, 0x2c, 0x4d, 0xd9, 0x72, 0xf2, 0x42, 0x7f, 0x6d, 0xa9, 0xf6,
	0x6d, 0x2f, 0xa1, 0x3a, 0xca, 0x58, 0x8a, 0xde, 0x4e, 0x39, 0x4f, 0x3d, 0xec, 0xf4, 0x1f, 0xcb,
	0xef, 0xec, 0xf7, 0xdd, 0x9b, 0x2e, 0x09, 0xf9, 0xf0, 0x66, 0x4b, 0x77, 0x15, 0xe1, 0x1a, 0x3a,
	0x34, 0x54, 0x23, 0x68, 0xa6, 0xea, 0xcf, 0x73, 0xde, 0x10, 0xae, 0x9b, 0x31, 0x6f, 0x6a, 0xd4,
	0xd8, 0xfc, 0xdb, 0x9b, 0xbb, 0x43, 0x8c, 0x09, 0xd7, 0x18, 0x8f, 0x06, 0xc6, 0x35, 0xc2, 0xd8,
	0x3a, 0xaf, 0xbe, 0x2d, 0x76, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xb8, 0xc1, 0xc6, 0x91,
	0x01, 0x00, 0x00,
}
