// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/common/frequency_cap.proto

package common // import "github.com/Beeketing/go-genproto/googleapis/ads/googleads/v0/common"

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

// A rule specifying the maximum number of times an ad (or some set of ads) can
// be shown to a user over a particular time period.
type FrequencyCapEntry struct {
	// The key of a particular frequency cap. There can be no more
	// than one frequency cap with the same key.
	Key *FrequencyCapKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Maximum number of events allowed during the time range by this cap.
	Cap                  *wrappers.Int32Value `protobuf:"bytes,2,opt,name=cap,proto3" json:"cap,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FrequencyCapEntry) Reset()         { *m = FrequencyCapEntry{} }
func (m *FrequencyCapEntry) String() string { return proto.CompactTextString(m) }
func (*FrequencyCapEntry) ProtoMessage()    {}
func (*FrequencyCapEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_frequency_cap_9bbb940e531af356, []int{0}
}
func (m *FrequencyCapEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrequencyCapEntry.Unmarshal(m, b)
}
func (m *FrequencyCapEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrequencyCapEntry.Marshal(b, m, deterministic)
}
func (dst *FrequencyCapEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrequencyCapEntry.Merge(dst, src)
}
func (m *FrequencyCapEntry) XXX_Size() int {
	return xxx_messageInfo_FrequencyCapEntry.Size(m)
}
func (m *FrequencyCapEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_FrequencyCapEntry.DiscardUnknown(m)
}

var xxx_messageInfo_FrequencyCapEntry proto.InternalMessageInfo

func (m *FrequencyCapEntry) GetKey() *FrequencyCapKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *FrequencyCapEntry) GetCap() *wrappers.Int32Value {
	if m != nil {
		return m.Cap
	}
	return nil
}

// A group of fields used as keys for a frequency cap.
// There can be no more than one frequency cap with the same key.
type FrequencyCapKey struct {
	// The level on which the cap is to be applied (e.g. ad group ad, ad group).
	// The cap is applied to all the entities of this level.
	Level enums.FrequencyCapLevelEnum_FrequencyCapLevel `protobuf:"varint,1,opt,name=level,proto3,enum=google.ads.googleads.v0.enums.FrequencyCapLevelEnum_FrequencyCapLevel" json:"level,omitempty"`
	// The type of event that the cap applies to (e.g. impression).
	EventType enums.FrequencyCapEventTypeEnum_FrequencyCapEventType `protobuf:"varint,3,opt,name=event_type,json=eventType,proto3,enum=google.ads.googleads.v0.enums.FrequencyCapEventTypeEnum_FrequencyCapEventType" json:"event_type,omitempty"`
	// Unit of time the cap is defined at (e.g. day, week).
	TimeUnit enums.FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit `protobuf:"varint,2,opt,name=time_unit,json=timeUnit,proto3,enum=google.ads.googleads.v0.enums.FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit" json:"time_unit,omitempty"`
	// Number of time units the cap lasts.
	TimeLength           *wrappers.Int32Value `protobuf:"bytes,4,opt,name=time_length,json=timeLength,proto3" json:"time_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FrequencyCapKey) Reset()         { *m = FrequencyCapKey{} }
func (m *FrequencyCapKey) String() string { return proto.CompactTextString(m) }
func (*FrequencyCapKey) ProtoMessage()    {}
func (*FrequencyCapKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_frequency_cap_9bbb940e531af356, []int{1}
}
func (m *FrequencyCapKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrequencyCapKey.Unmarshal(m, b)
}
func (m *FrequencyCapKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrequencyCapKey.Marshal(b, m, deterministic)
}
func (dst *FrequencyCapKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrequencyCapKey.Merge(dst, src)
}
func (m *FrequencyCapKey) XXX_Size() int {
	return xxx_messageInfo_FrequencyCapKey.Size(m)
}
func (m *FrequencyCapKey) XXX_DiscardUnknown() {
	xxx_messageInfo_FrequencyCapKey.DiscardUnknown(m)
}

var xxx_messageInfo_FrequencyCapKey proto.InternalMessageInfo

func (m *FrequencyCapKey) GetLevel() enums.FrequencyCapLevelEnum_FrequencyCapLevel {
	if m != nil {
		return m.Level
	}
	return enums.FrequencyCapLevelEnum_UNSPECIFIED
}

func (m *FrequencyCapKey) GetEventType() enums.FrequencyCapEventTypeEnum_FrequencyCapEventType {
	if m != nil {
		return m.EventType
	}
	return enums.FrequencyCapEventTypeEnum_UNSPECIFIED
}

func (m *FrequencyCapKey) GetTimeUnit() enums.FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit {
	if m != nil {
		return m.TimeUnit
	}
	return enums.FrequencyCapTimeUnitEnum_UNSPECIFIED
}

func (m *FrequencyCapKey) GetTimeLength() *wrappers.Int32Value {
	if m != nil {
		return m.TimeLength
	}
	return nil
}

func init() {
	proto.RegisterType((*FrequencyCapEntry)(nil), "google.ads.googleads.v0.common.FrequencyCapEntry")
	proto.RegisterType((*FrequencyCapKey)(nil), "google.ads.googleads.v0.common.FrequencyCapKey")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/common/frequency_cap.proto", fileDescriptor_frequency_cap_9bbb940e531af356)
}

var fileDescriptor_frequency_cap_9bbb940e531af356 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0x26, 0x1b, 0x15, 0x3b, 0x85, 0x8a, 0xb9, 0x5a, 0x2a, 0x14, 0xd9, 0x2b, 0x6f, 0x9c, 0x84,
	0xf4, 0x42, 0x48, 0x7b, 0x93, 0xd6, 0x6d, 0x11, 0x17, 0x29, 0xa1, 0xee, 0x85, 0x04, 0x96, 0x69,
	0x72, 0x1a, 0x83, 0x99, 0x1f, 0x33, 0x93, 0x95, 0x3c, 0x80, 0xf8, 0x1e, 0x5e, 0xfa, 0x28, 0x3e,
	0x8a, 0xf8, 0x10, 0x32, 0x33, 0x49, 0x74, 0x5d, 0xb6, 0x25, 0x57, 0xf9, 0x32, 0xe7, 0xfb, 0x39,
	0x39, 0x67, 0x82, 0xc2, 0x82, 0xf3, 0xa2, 0x02, 0x9f, 0xe4, 0xd2, 0xb7, 0x50, 0xa3, 0x75, 0xe0,
	0x67, 0x9c, 0x52, 0xce, 0xfc, 0xdb, 0x1a, 0x3e, 0x37, 0xc0, 0xb2, 0x76, 0x95, 0x11, 0x81, 0x45,
	0xcd, 0x15, 0xf7, 0x8e, 0x2c, 0x11, 0x93, 0x5c, 0xe2, 0x41, 0x83, 0xd7, 0x01, 0xb6, 0x9a, 0xc3,
	0xd3, 0x5d, 0x9e, 0xc0, 0x1a, 0x2a, 0x37, 0x2d, 0x57, 0xb0, 0x06, 0xa6, 0x56, 0xaa, 0x15, 0x60,
	0xdd, 0x0f, 0x5f, 0x8d, 0x51, 0x57, 0xb0, 0x86, 0xaa, 0x13, 0x9e, 0x8c, 0x11, 0xaa, 0x92, 0xc2,
	0xaa, 0x61, 0xa5, 0xea, 0xc4, 0xdd, 0x37, 0xf9, 0xe6, 0xed, 0xa6, 0xb9, 0xf5, 0xbf, 0xd4, 0x44,
	0x08, 0xa8, 0xa5, 0xad, 0xcf, 0xbe, 0x3a, 0xe8, 0xe9, 0x45, 0xef, 0x70, 0x4e, 0xc4, 0x9c, 0xa9,
	0xba, 0xf5, 0x62, 0xe4, 0x7e, 0x82, 0x76, 0xea, 0x3c, 0x77, 0x5e, 0xec, 0x87, 0x3e, 0xbe, 0x7b,
	0x2e, 0xf8, 0x5f, 0xfd, 0x5b, 0x68, 0x13, 0xad, 0xf5, 0x5e, 0x22, 0x37, 0x23, 0x62, 0x3a, 0x31,
	0x16, 0xcf, 0x7a, 0x8b, 0xbe, 0x0d, 0xfc, 0x86, 0xa9, 0xe3, 0x70, 0x49, 0xaa, 0x06, 0x12, 0xcd,
	0x9b, 0x7d, 0x73, 0xd1, 0x93, 0xff, 0x7c, 0xbc, 0x14, 0x3d, 0x34, 0x73, 0x30, 0x7d, 0x1c, 0x84,
	0x17, 0x3b, 0xfb, 0x30, 0x83, 0xd8, 0x68, 0x63, 0xa1, 0x75, 0x73, 0xd6, 0xd0, 0xed, 0xd3, 0xc4,
	0x9a, 0x7a, 0x14, 0xa1, 0xbf, 0x3b, 0x9a, 0xba, 0x26, 0xe2, 0xdd, 0x88, 0x88, 0xb9, 0x16, 0x5f,
	0xb7, 0x02, 0xb6, 0x62, 0x86, 0x4a, 0xb2, 0x07, 0x3d, 0xf4, 0x4a, 0xb4, 0x37, 0xec, 0xc6, 0x4c,
	0xe5, 0x20, 0x5c, 0x8c, 0x48, 0xbb, 0x2e, 0x29, 0xbc, 0x67, 0xa5, 0xda, 0x0a, 0xeb, 0x0b, 0xc9,
	0x63, 0xd5, 0x21, 0xef, 0x14, 0xed, 0x9b, 0xa8, 0x0a, 0x58, 0xa1, 0x3e, 0x4e, 0x1f, 0xdc, 0xbf,
	0x02, 0xa4, 0xf9, 0x0b, 0x43, 0x3f, 0xfb, 0xed, 0xa0, 0x59, 0xc6, 0xe9, 0x3d, 0x4b, 0x3f, 0xdb,
	0xb8, 0x35, 0x57, 0xda, 0xf3, 0xca, 0xf9, 0xf0, 0xba, 0x13, 0x15, 0xbc, 0x22, 0xac, 0xc0, 0xbc,
	0x2e, 0xfc, 0x02, 0x98, 0x49, 0xec, 0xaf, 0xae, 0x28, 0xe5, 0xae, 0x9f, 0xf2, 0xc4, 0x3e, 0xbe,
	0x4f, 0xdc, 0xcb, 0x38, 0xfe, 0x31, 0x39, 0xba, 0xb4, 0x66, 0x71, 0x2e, 0xb1, 0x85, 0x1a, 0x2d,
	0x03, 0x7c, 0x6e, 0x68, 0x3f, 0x7b, 0x42, 0x1a, 0xe7, 0x32, 0x1d, 0x08, 0xe9, 0x32, 0x48, 0x2d,
	0xe1, 0xd7, 0x64, 0x66, 0x4f, 0xa3, 0x28, 0xce, 0x65, 0x14, 0x0d, 0x94, 0x28, 0x5a, 0x06, 0x51,
	0x64, 0x49, 0x37, 0x8f, 0x4c, 0x77, 0xc7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x05, 0x27, 0x7e,
	0x57, 0x31, 0x04, 0x00, 0x00,
}
