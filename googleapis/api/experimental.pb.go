// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/experimental/experimental.proto

package api // import "github.com/Beeketing/go-genproto/googleapis/api"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/Beeketing/go-genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Experimental service configuration. These configuration options can
// only be used by whitelisted users.
type Experimental struct {
	// Authorization configuration.
	Authorization        *AuthorizationConfig `protobuf:"bytes,8,opt,name=authorization,proto3" json:"authorization,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Experimental) Reset()         { *m = Experimental{} }
func (m *Experimental) String() string { return proto.CompactTextString(m) }
func (*Experimental) ProtoMessage()    {}
func (*Experimental) Descriptor() ([]byte, []int) {
	return fileDescriptor_experimental_6711a67650a507d9, []int{0}
}
func (m *Experimental) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Experimental.Unmarshal(m, b)
}
func (m *Experimental) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Experimental.Marshal(b, m, deterministic)
}
func (dst *Experimental) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Experimental.Merge(dst, src)
}
func (m *Experimental) XXX_Size() int {
	return xxx_messageInfo_Experimental.Size(m)
}
func (m *Experimental) XXX_DiscardUnknown() {
	xxx_messageInfo_Experimental.DiscardUnknown(m)
}

var xxx_messageInfo_Experimental proto.InternalMessageInfo

func (m *Experimental) GetAuthorization() *AuthorizationConfig {
	if m != nil {
		return m.Authorization
	}
	return nil
}

func init() {
	proto.RegisterType((*Experimental)(nil), "google.api.Experimental")
}

func init() {
	proto.RegisterFile("google/api/experimental/experimental.proto", fileDescriptor_experimental_6711a67650a507d9)
}

var fileDescriptor_experimental_6711a67650a507d9 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4a, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xad, 0x28, 0x48, 0x2d, 0xca, 0xcc, 0x4d, 0xcd,
	0x2b, 0x49, 0xcc, 0x41, 0xe1, 0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x71, 0x41, 0xd4, 0xea,
	0x25, 0x16, 0x64, 0x4a, 0xc9, 0x20, 0xe9, 0x4b, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc,
	0xcf, 0x2b, 0x86, 0xa8, 0x94, 0x32, 0xc2, 0x65, 0x6a, 0x62, 0x69, 0x49, 0x46, 0x7e, 0x51, 0x66,
	0x15, 0x58, 0x75, 0x7c, 0x72, 0x7e, 0x5e, 0x5a, 0x66, 0x3a, 0x44, 0x8f, 0x52, 0x28, 0x17, 0x8f,
	0x2b, 0x92, 0x52, 0x21, 0x57, 0x2e, 0x5e, 0x14, 0xd5, 0x12, 0x1c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46,
	0xf2, 0x7a, 0x08, 0x57, 0xe8, 0x39, 0x22, 0x2b, 0x70, 0x06, 0x9b, 0x16, 0x84, 0xaa, 0xcb, 0x29,
	0x9a, 0x8b, 0x2f, 0x39, 0x3f, 0x17, 0x49, 0x93, 0x93, 0x20, 0xb2, 0x35, 0x01, 0x20, 0xbb, 0x03,
	0x18, 0xa3, 0x74, 0xa1, 0x0a, 0xd2, 0xf3, 0x73, 0x12, 0xf3, 0xd2, 0xf5, 0xf2, 0x8b, 0xd2, 0xf5,
	0xd3, 0x53, 0xf3, 0xc0, 0x2e, 0xd3, 0x87, 0x48, 0x25, 0x16, 0x64, 0x16, 0x83, 0x3c, 0x64, 0x9d,
	0x58, 0x90, 0xb9, 0x88, 0x89, 0xc5, 0xdd, 0x31, 0xc0, 0x33, 0x89, 0x0d, 0xac, 0xc0, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0xa0, 0x95, 0x20, 0xe5, 0x46, 0x01, 0x00, 0x00,
}
