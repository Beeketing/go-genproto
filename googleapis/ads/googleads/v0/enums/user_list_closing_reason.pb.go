// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/user_list_closing_reason.proto

package enums // import "github.com/Beeketing/genproto/googleapis/ads/googleads/v0/enums"

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

// Enum describing possible user list closing reasons.
type UserListClosingReasonEnum_UserListClosingReason int32

const (
	// Not specified.
	UserListClosingReasonEnum_UNSPECIFIED UserListClosingReasonEnum_UserListClosingReason = 0
	// Used for return value only. Represents value unknown in this version.
	UserListClosingReasonEnum_UNKNOWN UserListClosingReasonEnum_UserListClosingReason = 1
	// The userlist was closed because of not being used for over one year.
	UserListClosingReasonEnum_UNUSED UserListClosingReasonEnum_UserListClosingReason = 2
)

var UserListClosingReasonEnum_UserListClosingReason_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "UNUSED",
}
var UserListClosingReasonEnum_UserListClosingReason_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"UNUSED":      2,
}

func (x UserListClosingReasonEnum_UserListClosingReason) String() string {
	return proto.EnumName(UserListClosingReasonEnum_UserListClosingReason_name, int32(x))
}
func (UserListClosingReasonEnum_UserListClosingReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_list_closing_reason_bd7a2f8e86d42e4f, []int{0, 0}
}

// Indicates the reason why the userlist was closed.
// This enum is only used when a list is auto-closed by the system.
type UserListClosingReasonEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListClosingReasonEnum) Reset()         { *m = UserListClosingReasonEnum{} }
func (m *UserListClosingReasonEnum) String() string { return proto.CompactTextString(m) }
func (*UserListClosingReasonEnum) ProtoMessage()    {}
func (*UserListClosingReasonEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_list_closing_reason_bd7a2f8e86d42e4f, []int{0}
}
func (m *UserListClosingReasonEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListClosingReasonEnum.Unmarshal(m, b)
}
func (m *UserListClosingReasonEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListClosingReasonEnum.Marshal(b, m, deterministic)
}
func (dst *UserListClosingReasonEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListClosingReasonEnum.Merge(dst, src)
}
func (m *UserListClosingReasonEnum) XXX_Size() int {
	return xxx_messageInfo_UserListClosingReasonEnum.Size(m)
}
func (m *UserListClosingReasonEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListClosingReasonEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UserListClosingReasonEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UserListClosingReasonEnum)(nil), "google.ads.googleads.v0.enums.UserListClosingReasonEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.UserListClosingReasonEnum_UserListClosingReason", UserListClosingReasonEnum_UserListClosingReason_name, UserListClosingReasonEnum_UserListClosingReason_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/user_list_closing_reason.proto", fileDescriptor_user_list_closing_reason_bd7a2f8e86d42e4f)
}

var fileDescriptor_user_list_closing_reason_bd7a2f8e86d42e4f = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0x5d, 0x85, 0x09, 0xd9, 0xc1, 0x52, 0xf0, 0xa0, 0xb0, 0xc3, 0xf6, 0x00, 0x69, 0xc1,
	0x5b, 0xf4, 0x92, 0x6e, 0x75, 0x0c, 0xa5, 0x16, 0x47, 0x2b, 0x48, 0xb1, 0xd4, 0x25, 0x84, 0x42,
	0x9b, 0x8c, 0xfc, 0xdb, 0x3d, 0x90, 0x47, 0x1f, 0xc5, 0x47, 0xf1, 0xe2, 0x2b, 0x48, 0x13, 0xd7,
	0xd3, 0xf4, 0x12, 0x3e, 0xf2, 0xfd, 0x7f, 0xc9, 0xff, 0xfb, 0xd0, 0xad, 0x50, 0x4a, 0xd4, 0xdc,
	0x2f, 0x19, 0xf8, 0x56, 0xf6, 0x6a, 0x1f, 0xf8, 0x5c, 0x76, 0x0d, 0xf8, 0x1d, 0x70, 0x5d, 0xd4,
	0x15, 0xb4, 0xc5, 0xb6, 0x56, 0x50, 0x49, 0x51, 0x68, 0x5e, 0x82, 0x92, 0x78, 0xa7, 0x55, 0xab,
	0xbc, 0xa9, 0x45, 0x70, 0xc9, 0x00, 0x0f, 0x34, 0xde, 0x07, 0xd8, 0xd0, 0xf3, 0x57, 0x74, 0x99,
	0x02, 0xd7, 0x0f, 0x15, 0xb4, 0x0b, 0x8b, 0x3f, 0x19, 0x3a, 0x92, 0x5d, 0x33, 0xa7, 0xe8, 0xe2,
	0xa8, 0xe9, 0x9d, 0xa3, 0x49, 0x1a, 0x6f, 0x92, 0x68, 0xb1, 0xbe, 0x5b, 0x47, 0x4b, 0xf7, 0xc4,
	0x9b, 0xa0, 0xb3, 0x34, 0xbe, 0x8f, 0x1f, 0x9f, 0x63, 0x77, 0xe4, 0x21, 0x34, 0x4e, 0xe3, 0x74,
	0x13, 0x2d, 0x5d, 0x27, 0xfc, 0x1e, 0xa1, 0xd9, 0x56, 0x35, 0xf8, 0xdf, 0x2d, 0xc2, 0xab, 0xa3,
	0xdf, 0x24, 0x7d, 0x80, 0x64, 0xf4, 0x12, 0xfe, 0xc2, 0x42, 0xd5, 0xa5, 0x14, 0x58, 0x69, 0xe1,
	0x0b, 0x2e, 0x4d, 0xbc, 0x43, 0x21, 0xbb, 0x0a, 0xfe, 0xe8, 0xe7, 0xc6, 0x9c, 0xef, 0xce, 0xe9,
	0x8a, 0xd2, 0x0f, 0x67, 0xba, 0xb2, 0x4f, 0x51, 0x06, 0xd8, 0xca, 0x5e, 0x65, 0x01, 0xee, 0xf3,
	0xc2, 0xe7, 0xc1, 0xcf, 0x29, 0x83, 0x7c, 0xf0, 0xf3, 0x2c, 0xc8, 0x8d, 0xff, 0xe5, 0xcc, 0xec,
	0x25, 0x21, 0x94, 0x01, 0x21, 0xc3, 0x04, 0x21, 0x59, 0x40, 0x88, 0x99, 0x79, 0x1b, 0x9b, 0xc5,
	0xae, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x23, 0x07, 0xab, 0x8d, 0xb7, 0x01, 0x00, 0x00,
}
