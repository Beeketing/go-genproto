// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/user_list_error.proto

package errors // import "github.com/Beeketing/go-genproto/googleapis/ads/googleads/v0/errors"

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

// Enum describing possible user list errors.
type UserListErrorEnum_UserListError int32

const (
	// Enum unspecified.
	UserListErrorEnum_UNSPECIFIED UserListErrorEnum_UserListError = 0
	// The received error code is not known in this version.
	UserListErrorEnum_UNKNOWN UserListErrorEnum_UserListError = 1
	// Creating and updating external remarketing user lists is not supported.
	UserListErrorEnum_EXTERNAL_REMARKETING_USER_LIST_MUTATE_NOT_SUPPORTED UserListErrorEnum_UserListError = 2
	// Concrete type of user list is required.
	UserListErrorEnum_CONCRETE_TYPE_REQUIRED UserListErrorEnum_UserListError = 3
	// Creating/updating user list conversion types requires specifying the
	// conversion type Id.
	UserListErrorEnum_CONVERSION_TYPE_ID_REQUIRED UserListErrorEnum_UserListError = 4
	// Remarketing user list cannot have duplicate conversion types.
	UserListErrorEnum_DUPLICATE_CONVERSION_TYPES UserListErrorEnum_UserListError = 5
	// Conversion type is invalid/unknown.
	UserListErrorEnum_INVALID_CONVERSION_TYPE UserListErrorEnum_UserListError = 6
	// User list description is empty or invalid.
	UserListErrorEnum_INVALID_DESCRIPTION UserListErrorEnum_UserListError = 7
	// User list name is empty or invalid.
	UserListErrorEnum_INVALID_NAME UserListErrorEnum_UserListError = 8
	// Type of the UserList does not match.
	UserListErrorEnum_INVALID_TYPE UserListErrorEnum_UserListError = 9
	// Embedded logical user lists are not allowed.
	UserListErrorEnum_CAN_NOT_ADD_LOGICAL_LIST_AS_LOGICAL_LIST_OPERAND UserListErrorEnum_UserListError = 10
	// User list rule operand is invalid.
	UserListErrorEnum_INVALID_USER_LIST_LOGICAL_RULE_OPERAND UserListErrorEnum_UserListError = 11
	// Name is already being used for another user list for the account.
	UserListErrorEnum_NAME_ALREADY_USED UserListErrorEnum_UserListError = 12
	// Name is required when creating a new conversion type.
	UserListErrorEnum_NEW_CONVERSION_TYPE_NAME_REQUIRED UserListErrorEnum_UserListError = 13
	// The given conversion type name has been used.
	UserListErrorEnum_CONVERSION_TYPE_NAME_ALREADY_USED UserListErrorEnum_UserListError = 14
	// Only an owner account may edit a user list.
	UserListErrorEnum_OWNERSHIP_REQUIRED_FOR_SET UserListErrorEnum_UserListError = 15
	// Creating user list without setting type in oneof user_list field, or
	// creating/updating read-only user list types is not allowed.
	UserListErrorEnum_USER_LIST_MUTATE_NOT_SUPPORTED UserListErrorEnum_UserListError = 16
	// Rule is invalid.
	UserListErrorEnum_INVALID_RULE UserListErrorEnum_UserListError = 17
	// The specified date range is empty.
	UserListErrorEnum_INVALID_DATE_RANGE UserListErrorEnum_UserListError = 27
	// A UserList which is privacy sensitive or legal rejected cannot be mutated
	// by external users.
	UserListErrorEnum_CAN_NOT_MUTATE_SENSITIVE_USERLIST UserListErrorEnum_UserListError = 28
	// Maximum number of rulebased user lists a customer can have.
	UserListErrorEnum_MAX_NUM_RULEBASED_USERLISTS UserListErrorEnum_UserListError = 29
	// BasicUserList's billable record field cannot be modified once it is set.
	UserListErrorEnum_CANNOT_MODIFY_BILLABLE_RECORD_COUNT UserListErrorEnum_UserListError = 30
	// crm_based_user_list.app_id field must be set when upload_key_type is
	// MOBILE_ADVERTISING_ID.
	UserListErrorEnum_APP_ID_NOT_SET UserListErrorEnum_UserListError = 31
	// Name of the user list is reserved for system generated lists and cannot
	// be used.
	UserListErrorEnum_USERLIST_NAME_IS_RESERVED_FOR_SYSTEM_LIST UserListErrorEnum_UserListError = 32
	// Advertiser needs to be whitelisted to use remarketing lists created from
	// advertiser uploaded data (e.g., Customer Match lists).
	UserListErrorEnum_ADVERTISER_NOT_WHITELISTED_FOR_USING_UPLOADED_DATA UserListErrorEnum_UserListError = 33
	// The provided rule_type is not supported for the user list.
	UserListErrorEnum_RULE_TYPE_IS_NOT_SUPPORTED UserListErrorEnum_UserListError = 34
	// Similar user list cannot be used as a logical user list operand.
	UserListErrorEnum_CAN_NOT_ADD_A_SIMILAR_USERLIST_AS_LOGICAL_LIST_OPERAND UserListErrorEnum_UserListError = 35
	// Logical user list should not have a mix of CRM based user list and other
	// types of lists in its rules.
	UserListErrorEnum_CAN_NOT_MIX_CRM_BASED_IN_LOGICAL_LIST_WITH_OTHER_LISTS UserListErrorEnum_UserListError = 36
)

var UserListErrorEnum_UserListError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "EXTERNAL_REMARKETING_USER_LIST_MUTATE_NOT_SUPPORTED",
	3:  "CONCRETE_TYPE_REQUIRED",
	4:  "CONVERSION_TYPE_ID_REQUIRED",
	5:  "DUPLICATE_CONVERSION_TYPES",
	6:  "INVALID_CONVERSION_TYPE",
	7:  "INVALID_DESCRIPTION",
	8:  "INVALID_NAME",
	9:  "INVALID_TYPE",
	10: "CAN_NOT_ADD_LOGICAL_LIST_AS_LOGICAL_LIST_OPERAND",
	11: "INVALID_USER_LIST_LOGICAL_RULE_OPERAND",
	12: "NAME_ALREADY_USED",
	13: "NEW_CONVERSION_TYPE_NAME_REQUIRED",
	14: "CONVERSION_TYPE_NAME_ALREADY_USED",
	15: "OWNERSHIP_REQUIRED_FOR_SET",
	16: "USER_LIST_MUTATE_NOT_SUPPORTED",
	17: "INVALID_RULE",
	27: "INVALID_DATE_RANGE",
	28: "CAN_NOT_MUTATE_SENSITIVE_USERLIST",
	29: "MAX_NUM_RULEBASED_USERLISTS",
	30: "CANNOT_MODIFY_BILLABLE_RECORD_COUNT",
	31: "APP_ID_NOT_SET",
	32: "USERLIST_NAME_IS_RESERVED_FOR_SYSTEM_LIST",
	33: "ADVERTISER_NOT_WHITELISTED_FOR_USING_UPLOADED_DATA",
	34: "RULE_TYPE_IS_NOT_SUPPORTED",
	35: "CAN_NOT_ADD_A_SIMILAR_USERLIST_AS_LOGICAL_LIST_OPERAND",
	36: "CAN_NOT_MIX_CRM_BASED_IN_LOGICAL_LIST_WITH_OTHER_LISTS",
}
var UserListErrorEnum_UserListError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"EXTERNAL_REMARKETING_USER_LIST_MUTATE_NOT_SUPPORTED":    2,
	"CONCRETE_TYPE_REQUIRED":                                 3,
	"CONVERSION_TYPE_ID_REQUIRED":                            4,
	"DUPLICATE_CONVERSION_TYPES":                             5,
	"INVALID_CONVERSION_TYPE":                                6,
	"INVALID_DESCRIPTION":                                    7,
	"INVALID_NAME":                                           8,
	"INVALID_TYPE":                                           9,
	"CAN_NOT_ADD_LOGICAL_LIST_AS_LOGICAL_LIST_OPERAND":       10,
	"INVALID_USER_LIST_LOGICAL_RULE_OPERAND":                 11,
	"NAME_ALREADY_USED":                                      12,
	"NEW_CONVERSION_TYPE_NAME_REQUIRED":                      13,
	"CONVERSION_TYPE_NAME_ALREADY_USED":                      14,
	"OWNERSHIP_REQUIRED_FOR_SET":                             15,
	"USER_LIST_MUTATE_NOT_SUPPORTED":                         16,
	"INVALID_RULE":                                           17,
	"INVALID_DATE_RANGE":                                     27,
	"CAN_NOT_MUTATE_SENSITIVE_USERLIST":                      28,
	"MAX_NUM_RULEBASED_USERLISTS":                            29,
	"CANNOT_MODIFY_BILLABLE_RECORD_COUNT":                    30,
	"APP_ID_NOT_SET":                                         31,
	"USERLIST_NAME_IS_RESERVED_FOR_SYSTEM_LIST":              32,
	"ADVERTISER_NOT_WHITELISTED_FOR_USING_UPLOADED_DATA":     33,
	"RULE_TYPE_IS_NOT_SUPPORTED":                             34,
	"CAN_NOT_ADD_A_SIMILAR_USERLIST_AS_LOGICAL_LIST_OPERAND": 35,
	"CAN_NOT_MIX_CRM_BASED_IN_LOGICAL_LIST_WITH_OTHER_LISTS": 36,
}

func (x UserListErrorEnum_UserListError) String() string {
	return proto.EnumName(UserListErrorEnum_UserListError_name, int32(x))
}
func (UserListErrorEnum_UserListError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_list_error_e79c5f0b379df020, []int{0, 0}
}

// Container for enum describing possible user list errors.
type UserListErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListErrorEnum) Reset()         { *m = UserListErrorEnum{} }
func (m *UserListErrorEnum) String() string { return proto.CompactTextString(m) }
func (*UserListErrorEnum) ProtoMessage()    {}
func (*UserListErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_list_error_e79c5f0b379df020, []int{0}
}
func (m *UserListErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListErrorEnum.Unmarshal(m, b)
}
func (m *UserListErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListErrorEnum.Marshal(b, m, deterministic)
}
func (dst *UserListErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListErrorEnum.Merge(dst, src)
}
func (m *UserListErrorEnum) XXX_Size() int {
	return xxx_messageInfo_UserListErrorEnum.Size(m)
}
func (m *UserListErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UserListErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UserListErrorEnum)(nil), "google.ads.googleads.v0.errors.UserListErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.UserListErrorEnum_UserListError", UserListErrorEnum_UserListError_name, UserListErrorEnum_UserListError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/user_list_error.proto", fileDescriptor_user_list_error_e79c5f0b379df020)
}

var fileDescriptor_user_list_error_e79c5f0b379df020 = []byte{
	// 723 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xff, 0x6e, 0xeb, 0x34,
	0x14, 0x66, 0xbd, 0x70, 0x0b, 0xde, 0xbd, 0x9b, 0x67, 0xc4, 0x26, 0x6d, 0xd0, 0xb1, 0x8e, 0x1f,
	0x02, 0x89, 0xb4, 0x62, 0xd3, 0x90, 0xc2, 0x5f, 0x6e, 0x7c, 0xd6, 0x5a, 0x4b, 0x9c, 0x60, 0x3b,
	0xe9, 0x8a, 0x2a, 0x59, 0x83, 0x56, 0xd5, 0xa4, 0x6d, 0x99, 0x9a, 0x6d, 0x0f, 0xc4, 0x9f, 0x3c,
	0x0a, 0x0f, 0xc0, 0x43, 0x20, 0x78, 0x07, 0x64, 0xa7, 0xc9, 0xda, 0x8a, 0x71, 0xff, 0xaa, 0x7b,
	0xce, 0xf7, 0x7d, 0xf6, 0xf9, 0xce, 0xc9, 0x41, 0xa7, 0xb3, 0x3c, 0x9f, 0xdd, 0x4c, 0x3b, 0x57,
	0x93, 0xa2, 0x53, 0x1e, 0xed, 0xe9, 0xa9, 0xdb, 0x99, 0xce, 0xe7, 0xf9, 0xbc, 0xe8, 0x3c, 0x16,
	0xd3, 0xb9, 0xb9, 0xb9, 0x2e, 0x1e, 0x8c, 0x0b, 0x78, 0xf7, 0xf3, 0xfc, 0x21, 0x27, 0xad, 0x12,
	0xea, 0x5d, 0x4d, 0x0a, 0xaf, 0x66, 0x79, 0x4f, 0x5d, 0xaf, 0x64, 0xb5, 0xff, 0x6e, 0xa2, 0x9d,
	0xb4, 0x98, 0xce, 0xc3, 0xeb, 0xe2, 0x01, 0x6c, 0x08, 0xee, 0x1e, 0x6f, 0xdb, 0x7f, 0x36, 0xd1,
	0xdb, 0x95, 0x28, 0xd9, 0x46, 0x9b, 0xa9, 0x50, 0x09, 0x04, 0xfc, 0x9c, 0x03, 0xc3, 0xef, 0x91,
	0x4d, 0xd4, 0x4c, 0xc5, 0x85, 0x88, 0x87, 0x02, 0x6f, 0x90, 0x1f, 0xd0, 0x09, 0x5c, 0x6a, 0x90,
	0x82, 0x86, 0x46, 0x42, 0x44, 0xe5, 0x05, 0x68, 0x2e, 0xfa, 0x26, 0x55, 0x20, 0x4d, 0xc8, 0x95,
	0x36, 0x51, 0xaa, 0xa9, 0x06, 0x23, 0x62, 0x6d, 0x54, 0x9a, 0x24, 0xb1, 0xd4, 0xc0, 0x70, 0x83,
	0xec, 0xa3, 0xdd, 0x20, 0x16, 0x81, 0x04, 0x0d, 0x46, 0x8f, 0x12, 0x30, 0x12, 0x7e, 0x4a, 0xb9,
	0x04, 0x86, 0x5f, 0x91, 0x43, 0x74, 0x10, 0xc4, 0x22, 0x03, 0xa9, 0x78, 0x2c, 0xca, 0x2c, 0x67,
	0xcf, 0x80, 0xf7, 0x49, 0x0b, 0xed, 0xb3, 0x34, 0x09, 0x79, 0x60, 0x95, 0xd7, 0xa0, 0x0a, 0x7f,
	0x40, 0x0e, 0xd0, 0x1e, 0x17, 0x19, 0x0d, 0x39, 0x5b, 0xcf, 0xe2, 0xd7, 0x64, 0x0f, 0x7d, 0x5c,
	0x25, 0x19, 0xa8, 0x40, 0xf2, 0x44, 0xf3, 0x58, 0xe0, 0x26, 0xc1, 0xe8, 0x4d, 0x95, 0x10, 0x34,
	0x02, 0xfc, 0xe1, 0x72, 0xc4, 0x91, 0x3f, 0x22, 0xa7, 0xa8, 0x1b, 0x50, 0xe1, 0xaa, 0xa1, 0x8c,
	0x99, 0x30, 0xee, 0xf3, 0x80, 0x86, 0x65, 0xa5, 0x54, 0xad, 0xfe, 0x8f, 0x13, 0x90, 0x54, 0x30,
	0x8c, 0xc8, 0xb7, 0xe8, 0xab, 0x4a, 0xe7, 0xd9, 0x98, 0x0a, 0x2b, 0xd3, 0x10, 0x6a, 0xec, 0x26,
	0xf9, 0x04, 0xed, 0xd8, 0xdb, 0x0d, 0x0d, 0x25, 0x50, 0x36, 0xb2, 0x04, 0x86, 0xdf, 0x90, 0x2f,
	0xd1, 0x91, 0x80, 0xe1, 0x7a, 0x39, 0xee, 0xa1, 0xcf, 0xce, 0xbc, 0xb5, 0xb0, 0xff, 0x84, 0xac,
	0xa8, 0x6d, 0x59, 0x03, 0xe3, 0xa1, 0x00, 0xa9, 0x06, 0x3c, 0xa9, 0xe9, 0xe6, 0x3c, 0x96, 0x46,
	0x81, 0xc6, 0xdb, 0xa4, 0x8d, 0x5a, 0xef, 0xe8, 0x20, 0x5e, 0x36, 0xc7, 0x96, 0x80, 0x77, 0xc8,
	0x2e, 0x22, 0xb5, 0xb3, 0x96, 0x21, 0xa9, 0xe8, 0x03, 0x3e, 0x70, 0x8f, 0x5a, 0x98, 0xb6, 0xd0,
	0x52, 0x20, 0x14, 0xd7, 0x3c, 0x03, 0xe7, 0x87, 0xbd, 0x05, 0x7f, 0x6a, 0xdb, 0x1e, 0xd1, 0x4b,
	0x23, 0xd2, 0xc8, 0x09, 0xf6, 0xa8, 0x02, 0x56, 0xe7, 0x15, 0xfe, 0x8c, 0x7c, 0x8d, 0x8e, 0x03,
	0x2a, 0x9c, 0x4c, 0xcc, 0xf8, 0xf9, 0xc8, 0xf4, 0x78, 0x18, 0xd2, 0x5e, 0x68, 0x1d, 0x08, 0x62,
	0x69, 0x9b, 0x9d, 0x0a, 0x8d, 0x5b, 0x84, 0xa0, 0x2d, 0x9a, 0x24, 0x76, 0x68, 0xdc, 0xa3, 0x41,
	0xe3, 0x43, 0xf2, 0x1d, 0xfa, 0xa6, 0xd2, 0x2a, 0x2d, 0xe1, 0xca, 0x48, 0x50, 0x20, 0xb3, 0xaa,
	0xf2, 0x91, 0xd2, 0x10, 0xb9, 0x92, 0xf1, 0xe7, 0xe4, 0x0c, 0x7d, 0x4f, 0x59, 0x06, 0x52, 0x73,
	0xeb, 0x83, 0x95, 0x19, 0x0e, 0xb8, 0x06, 0x9b, 0x5d, 0xe0, 0x53, 0xe5, 0x06, 0x3d, 0x09, 0x63,
	0xca, 0xc0, 0x55, 0x4c, 0xf1, 0x91, 0x75, 0xd6, 0x35, 0xb4, 0x9c, 0x5a, 0xb5, 0xe6, 0x5a, 0x9b,
	0xf8, 0xe8, 0x6c, 0x79, 0x80, 0xa8, 0x51, 0x3c, 0xe2, 0x21, 0x95, 0x75, 0xa1, 0x2f, 0x8e, 0xd1,
	0xf1, 0x32, 0x37, 0xe2, 0x97, 0x26, 0x90, 0x91, 0x29, 0x4d, 0xe2, 0x62, 0x95, 0x32, 0xe4, 0x7a,
	0x60, 0x62, 0x3d, 0x58, 0x74, 0x50, 0xe1, 0x2f, 0x7a, 0xff, 0x6c, 0xa0, 0xf6, 0xaf, 0xf9, 0xad,
	0xf7, 0xff, 0x5b, 0xa1, 0x47, 0x56, 0x3e, 0xfe, 0xc4, 0x6e, 0x92, 0x64, 0xe3, 0x67, 0xb6, 0x60,
	0xcd, 0xf2, 0x9b, 0xab, 0xbb, 0x99, 0x97, 0xcf, 0x67, 0x9d, 0xd9, 0xf4, 0xce, 0xed, 0x99, 0x6a,
	0x23, 0xdd, 0x5f, 0x17, 0x2f, 0x2d, 0xa8, 0x1f, 0xcb, 0x9f, 0xdf, 0x1a, 0xaf, 0xfa, 0x94, 0xfe,
	0xde, 0x68, 0xf5, 0x4b, 0x31, 0x3a, 0x29, 0xbc, 0xf2, 0x68, 0x4f, 0x59, 0xd7, 0x73, 0x57, 0x16,
	0x7f, 0x54, 0x80, 0x31, 0x9d, 0x14, 0xe3, 0x1a, 0x30, 0xce, 0xba, 0xe3, 0x12, 0xf0, 0x57, 0xa3,
	0x5d, 0x46, 0x7d, 0x9f, 0x4e, 0x0a, 0xdf, 0xaf, 0x21, 0xbe, 0x9f, 0x75, 0x7d, 0xbf, 0x04, 0xfd,
	0xf2, 0xda, 0xbd, 0xee, 0xe4, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x65, 0x0a, 0x53, 0x90, 0x3d,
	0x05, 0x00, 0x00,
}
