// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/user_list_rule_type.proto

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

// Enum describing possible user list rule types.
type UserListRuleTypeEnum_UserListRuleType int32

const (
	// Not specified.
	UserListRuleTypeEnum_UNSPECIFIED UserListRuleTypeEnum_UserListRuleType = 0
	// Used for return value only. Represents value unknown in this version.
	UserListRuleTypeEnum_UNKNOWN UserListRuleTypeEnum_UserListRuleType = 1
	// Conjunctive normal form.
	UserListRuleTypeEnum_AND_OF_ORS UserListRuleTypeEnum_UserListRuleType = 2
	// Disjunctive normal form.
	UserListRuleTypeEnum_OR_OF_ANDS UserListRuleTypeEnum_UserListRuleType = 3
)

var UserListRuleTypeEnum_UserListRuleType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "AND_OF_ORS",
	3: "OR_OF_ANDS",
}
var UserListRuleTypeEnum_UserListRuleType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"AND_OF_ORS":  2,
	"OR_OF_ANDS":  3,
}

func (x UserListRuleTypeEnum_UserListRuleType) String() string {
	return proto.EnumName(UserListRuleTypeEnum_UserListRuleType_name, int32(x))
}
func (UserListRuleTypeEnum_UserListRuleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_list_rule_type_1306ab40864c33aa, []int{0, 0}
}

// Rule based user list rule type.
type UserListRuleTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListRuleTypeEnum) Reset()         { *m = UserListRuleTypeEnum{} }
func (m *UserListRuleTypeEnum) String() string { return proto.CompactTextString(m) }
func (*UserListRuleTypeEnum) ProtoMessage()    {}
func (*UserListRuleTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_list_rule_type_1306ab40864c33aa, []int{0}
}
func (m *UserListRuleTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListRuleTypeEnum.Unmarshal(m, b)
}
func (m *UserListRuleTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListRuleTypeEnum.Marshal(b, m, deterministic)
}
func (dst *UserListRuleTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListRuleTypeEnum.Merge(dst, src)
}
func (m *UserListRuleTypeEnum) XXX_Size() int {
	return xxx_messageInfo_UserListRuleTypeEnum.Size(m)
}
func (m *UserListRuleTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListRuleTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UserListRuleTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UserListRuleTypeEnum)(nil), "google.ads.googleads.v0.enums.UserListRuleTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.UserListRuleTypeEnum_UserListRuleType", UserListRuleTypeEnum_UserListRuleType_name, UserListRuleTypeEnum_UserListRuleType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/user_list_rule_type.proto", fileDescriptor_user_list_rule_type_1306ab40864c33aa)
}

var fileDescriptor_user_list_rule_type_1306ab40864c33aa = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xb6, 0x29, 0x28, 0x6c, 0x41, 0x43, 0xd0, 0x63, 0x0f, 0xed, 0x03, 0x6c, 0x02, 0x1e, 0x84,
	0xf5, 0xb4, 0xb5, 0x3f, 0x14, 0x65, 0x13, 0x12, 0x13, 0x41, 0x02, 0x21, 0x9a, 0x65, 0x0d, 0x24,
	0xd9, 0x90, 0x49, 0x0a, 0x7d, 0x1d, 0x8f, 0x3e, 0x8a, 0x8f, 0xd2, 0xa7, 0x90, 0xdd, 0xd8, 0x1c,
	0x0a, 0x7a, 0x59, 0xbe, 0x99, 0xef, 0xfb, 0x76, 0xe6, 0x1b, 0x74, 0x27, 0xa4, 0x14, 0x05, 0xb7,
	0xd3, 0x0c, 0xec, 0x1e, 0x2a, 0xb4, 0x73, 0x6c, 0x5e, 0x75, 0x25, 0xd8, 0x1d, 0xf0, 0x26, 0x29,
	0x72, 0x68, 0x93, 0xa6, 0x2b, 0x78, 0xd2, 0xee, 0x6b, 0x8e, 0xeb, 0x46, 0xb6, 0xd2, 0x9a, 0xf6,
	0x6a, 0x9c, 0x66, 0x80, 0x07, 0x23, 0xde, 0x39, 0x58, 0x1b, 0xe7, 0x1f, 0xe8, 0x3a, 0x04, 0xde,
	0x3c, 0xe5, 0xd0, 0xfa, 0x5d, 0xc1, 0x9f, 0xf7, 0x35, 0x5f, 0x55, 0x5d, 0x39, 0xf7, 0x90, 0x79,
	0xda, 0xb7, 0xae, 0xd0, 0x24, 0x64, 0x81, 0xb7, 0x7a, 0xd8, 0xae, 0xb7, 0xab, 0xa5, 0x79, 0x66,
	0x4d, 0xd0, 0x45, 0xc8, 0x1e, 0x99, 0xfb, 0xc2, 0xcc, 0x91, 0x75, 0x89, 0x10, 0x65, 0xcb, 0xc4,
	0x5d, 0x27, 0xae, 0x1f, 0x98, 0x86, 0xaa, 0x5d, 0x5f, 0x95, 0x94, 0x2d, 0x03, 0x73, 0xbc, 0x38,
	0x8c, 0xd0, 0xec, 0x5d, 0x96, 0xf8, 0xdf, 0x7d, 0x16, 0x37, 0xa7, 0x53, 0x3d, 0x95, 0xc2, 0x1b,
	0xbd, 0x2e, 0x7e, 0x7d, 0x42, 0x16, 0x69, 0x25, 0xb0, 0x6c, 0x84, 0x2d, 0x78, 0xa5, 0x33, 0x1e,
	0x0f, 0x52, 0xe7, 0xf0, 0xc7, 0x7d, 0xee, 0xf5, 0xfb, 0x69, 0x8c, 0x37, 0x94, 0x7e, 0x19, 0xd3,
	0x4d, 0xff, 0x15, 0xcd, 0x00, 0xf7, 0x50, 0xa1, 0xc8, 0xc1, 0x2a, 0x39, 0x7c, 0x1f, 0xf9, 0x98,
	0x66, 0x10, 0x0f, 0x7c, 0x1c, 0x39, 0xb1, 0xe6, 0x0f, 0xc6, 0xac, 0x6f, 0x12, 0x42, 0x33, 0x20,
	0x64, 0x50, 0x10, 0x12, 0x39, 0x84, 0x68, 0xcd, 0xdb, 0xb9, 0x5e, 0xec, 0xf6, 0x27, 0x00, 0x00,
	0xff, 0xff, 0x2f, 0x25, 0xe1, 0xc9, 0xb7, 0x01, 0x00, 0x00,
}
