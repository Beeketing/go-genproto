// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/common/explorer_auto_optimizer_setting.proto

package common // import "github.com/Beeketing/go-genproto/googleapis/ads/googleads/v0/common"

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

// Settings for the
// <a href="https://support.google.com/google-ads/answer/190596">
// Display Campaign Optimizer</a>, initially termed "Explorer".
type ExplorerAutoOptimizerSetting struct {
	// Indicates whether the optimizer is turned on.
	OptIn                *wrappers.BoolValue `protobuf:"bytes,1,opt,name=opt_in,json=optIn,proto3" json:"opt_in,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ExplorerAutoOptimizerSetting) Reset()         { *m = ExplorerAutoOptimizerSetting{} }
func (m *ExplorerAutoOptimizerSetting) String() string { return proto.CompactTextString(m) }
func (*ExplorerAutoOptimizerSetting) ProtoMessage()    {}
func (*ExplorerAutoOptimizerSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_explorer_auto_optimizer_setting_b7458dc3996d9eae, []int{0}
}
func (m *ExplorerAutoOptimizerSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExplorerAutoOptimizerSetting.Unmarshal(m, b)
}
func (m *ExplorerAutoOptimizerSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExplorerAutoOptimizerSetting.Marshal(b, m, deterministic)
}
func (dst *ExplorerAutoOptimizerSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExplorerAutoOptimizerSetting.Merge(dst, src)
}
func (m *ExplorerAutoOptimizerSetting) XXX_Size() int {
	return xxx_messageInfo_ExplorerAutoOptimizerSetting.Size(m)
}
func (m *ExplorerAutoOptimizerSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_ExplorerAutoOptimizerSetting.DiscardUnknown(m)
}

var xxx_messageInfo_ExplorerAutoOptimizerSetting proto.InternalMessageInfo

func (m *ExplorerAutoOptimizerSetting) GetOptIn() *wrappers.BoolValue {
	if m != nil {
		return m.OptIn
	}
	return nil
}

func init() {
	proto.RegisterType((*ExplorerAutoOptimizerSetting)(nil), "google.ads.googleads.v0.common.ExplorerAutoOptimizerSetting")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/common/explorer_auto_optimizer_setting.proto", fileDescriptor_explorer_auto_optimizer_setting_b7458dc3996d9eae)
}

var fileDescriptor_explorer_auto_optimizer_setting_b7458dc3996d9eae = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x95, 0x20, 0x3a, 0x84, 0xad, 0x13, 0xaa, 0x50, 0x05, 0x99, 0x98, 0xce, 0x01, 0x36,
	0x33, 0x39, 0x80, 0x2a, 0x26, 0x0a, 0x48, 0x19, 0x50, 0xa4, 0xc8, 0x6d, 0x8c, 0x15, 0x29, 0xf1,
	0x59, 0xb6, 0x53, 0x10, 0x8f, 0xc3, 0xc8, 0xa3, 0xf0, 0x28, 0x3c, 0x03, 0x03, 0x4a, 0x9c, 0x64,
	0xa3, 0x53, 0x7e, 0xc5, 0xdf, 0xfd, 0xdf, 0xe9, 0xa2, 0x5b, 0x89, 0x28, 0x6b, 0x41, 0x78, 0x69,
	0x89, 0x8f, 0x5d, 0xda, 0x25, 0x64, 0x8b, 0x4d, 0x83, 0x8a, 0x88, 0x77, 0x5d, 0xa3, 0x11, 0xa6,
	0xe0, 0xad, 0xc3, 0x02, 0xb5, 0xab, 0x9a, 0xea, 0x43, 0x98, 0xc2, 0x0a, 0xe7, 0x2a, 0x25, 0x41,
	0x1b, 0x74, 0x38, 0x5f, 0xfa, 0x51, 0xe0, 0xa5, 0x85, 0xa9, 0x05, 0x76, 0x09, 0xf8, 0x96, 0xc5,
	0xf0, 0x4e, 0x7a, 0x7a, 0xd3, 0xbe, 0x92, 0x37, 0xc3, 0xb5, 0x16, 0xc6, 0xfa, 0xf9, 0xf8, 0x31,
	0x3a, 0xb9, 0x1b, 0x44, 0xac, 0x75, 0xf8, 0x30, 0x6a, 0x9e, 0xbd, 0x65, 0x7e, 0x11, 0xcd, 0x50,
	0xbb, 0xa2, 0x52, 0xc7, 0xc1, 0x69, 0x70, 0x7e, 0x74, 0xb9, 0x18, 0x2c, 0x30, 0x16, 0x42, 0x8a,
	0x58, 0x67, 0xbc, 0x6e, 0xc5, 0xd3, 0x21, 0x6a, 0x77, 0xaf, 0xd2, 0xdf, 0x20, 0x8a, 0xb7, 0xd8,
	0xc0, 0xfe, 0xcd, 0xd2, 0xb3, 0x7d, 0xde, 0x75, 0xd7, 0xbe, 0x0e, 0x5e, 0x86, 0x23, 0x81, 0xc4,
	0x9a, 0x2b, 0x09, 0x68, 0x24, 0x91, 0x42, 0xf5, 0xee, 0xf1, 0x68, 0xba, 0xb2, 0xff, 0xdd, 0xf0,
	0xda, 0x7f, 0x3e, 0xc3, 0x83, 0x15, 0x63, 0x5f, 0xe1, 0x72, 0xe5, 0xcb, 0x58, 0x69, 0xc1, 0xc7,
	0x2e, 0x65, 0x09, 0xdc, 0xf4, 0xd8, 0xf7, 0x08, 0xe4, 0xac, 0xb4, 0xf9, 0x04, 0xe4, 0x59, 0x92,
	0x7b, 0xe0, 0x27, 0x8c, 0xfd, 0x5f, 0x4a, 0x59, 0x69, 0x29, 0x9d, 0x10, 0x4a, 0xb3, 0x84, 0x52,
	0x0f, 0x6d, 0x66, 0xfd, 0x76, 0x57, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x97, 0xb4, 0x35, 0x7a,
	0xe0, 0x01, 0x00, 0x00,
}
