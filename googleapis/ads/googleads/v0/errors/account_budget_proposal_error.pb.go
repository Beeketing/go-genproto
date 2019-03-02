// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/account_budget_proposal_error.proto

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

// Enum describing possible account budget proposal errors.
type AccountBudgetProposalErrorEnum_AccountBudgetProposalError int32

const (
	// Enum unspecified.
	AccountBudgetProposalErrorEnum_UNSPECIFIED AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 0
	// The received error code is not known in this version.
	AccountBudgetProposalErrorEnum_UNKNOWN AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 1
	// The field mask must be empty for create/end/remove proposals.
	AccountBudgetProposalErrorEnum_FIELD_MASK_NOT_ALLOWED AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 2
	// The field cannot be set because of the proposal type.
	AccountBudgetProposalErrorEnum_IMMUTABLE_FIELD AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 3
	// The field is required because of the proposal type.
	AccountBudgetProposalErrorEnum_REQUIRED_FIELD_MISSING AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 4
	// Proposals that have been approved cannot be cancelled.
	AccountBudgetProposalErrorEnum_CANNOT_CANCEL_APPROVED_PROPOSAL AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 5
	// Budgets that haven't been approved cannot be removed.
	AccountBudgetProposalErrorEnum_CANNOT_REMOVE_UNAPPROVED_BUDGET AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 6
	// Budgets that are currently running cannot be removed.
	AccountBudgetProposalErrorEnum_CANNOT_REMOVE_RUNNING_BUDGET AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 7
	// Budgets that haven't been approved cannot be truncated.
	AccountBudgetProposalErrorEnum_CANNOT_END_UNAPPROVED_BUDGET AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 8
	// Only budgets that are currently running can be truncated.
	AccountBudgetProposalErrorEnum_CANNOT_END_INACTIVE_BUDGET AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 9
	// All budgets must have names.
	AccountBudgetProposalErrorEnum_BUDGET_NAME_REQUIRED AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 10
	// Expired budgets cannot be edited after a sufficient amount of time has
	// passed.
	AccountBudgetProposalErrorEnum_CANNOT_UPDATE_OLD_BUDGET AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 11
	// It is not permissible a propose a new budget that ends in the past.
	AccountBudgetProposalErrorEnum_CANNOT_END_IN_PAST AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 12
	// An expired budget cannot be extended to overlap with the running budget.
	AccountBudgetProposalErrorEnum_CANNOT_EXTEND_END_TIME AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 13
	// A purchase order number is required.
	AccountBudgetProposalErrorEnum_PURCHASE_ORDER_NUMBER_REQUIRED AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 14
	// Budgets that have a pending update cannot be updated.
	AccountBudgetProposalErrorEnum_PENDING_UPDATE_PROPOSAL_EXISTS AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 15
	// Cannot propose more than one budget when the corresponding billing setup
	// hasn't been approved.
	AccountBudgetProposalErrorEnum_MULTIPLE_BUDGETS_NOT_ALLOWED_FOR_UNAPPROVED_BILLING_SETUP AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 16
	// Cannot update the start time of a budget that has already started.
	AccountBudgetProposalErrorEnum_CANNOT_UPDATE_START_TIME_FOR_STARTED_BUDGET AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 17
	// Cannot update the spending limit of a budget with an amount lower than
	// what has already been spent.
	AccountBudgetProposalErrorEnum_SPENDING_LIMIT_LOWER_THAN_ACCRUED_COST_NOT_ALLOWED AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 18
	// Cannot propose a budget update without actually changing any fields.
	AccountBudgetProposalErrorEnum_UPDATE_IS_NO_OP AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 19
	// The end time must come after the start time.
	AccountBudgetProposalErrorEnum_END_TIME_MUST_FOLLOW_START_TIME AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 20
	// The budget's date range must fall within the date range of its billing
	// setup.
	AccountBudgetProposalErrorEnum_BUDGET_DATE_RANGE_INCOMPATIBLE_WITH_BILLING_SETUP AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 21
	// The user is not authorized to mutate budgets for the given billing setup.
	AccountBudgetProposalErrorEnum_NOT_AUTHORIZED AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 22
	// Mutates are not allowed for the given billing setup.
	AccountBudgetProposalErrorEnum_INVALID_BILLING_SETUP AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 23
)

var AccountBudgetProposalErrorEnum_AccountBudgetProposalError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "FIELD_MASK_NOT_ALLOWED",
	3:  "IMMUTABLE_FIELD",
	4:  "REQUIRED_FIELD_MISSING",
	5:  "CANNOT_CANCEL_APPROVED_PROPOSAL",
	6:  "CANNOT_REMOVE_UNAPPROVED_BUDGET",
	7:  "CANNOT_REMOVE_RUNNING_BUDGET",
	8:  "CANNOT_END_UNAPPROVED_BUDGET",
	9:  "CANNOT_END_INACTIVE_BUDGET",
	10: "BUDGET_NAME_REQUIRED",
	11: "CANNOT_UPDATE_OLD_BUDGET",
	12: "CANNOT_END_IN_PAST",
	13: "CANNOT_EXTEND_END_TIME",
	14: "PURCHASE_ORDER_NUMBER_REQUIRED",
	15: "PENDING_UPDATE_PROPOSAL_EXISTS",
	16: "MULTIPLE_BUDGETS_NOT_ALLOWED_FOR_UNAPPROVED_BILLING_SETUP",
	17: "CANNOT_UPDATE_START_TIME_FOR_STARTED_BUDGET",
	18: "SPENDING_LIMIT_LOWER_THAN_ACCRUED_COST_NOT_ALLOWED",
	19: "UPDATE_IS_NO_OP",
	20: "END_TIME_MUST_FOLLOW_START_TIME",
	21: "BUDGET_DATE_RANGE_INCOMPATIBLE_WITH_BILLING_SETUP",
	22: "NOT_AUTHORIZED",
	23: "INVALID_BILLING_SETUP",
}
var AccountBudgetProposalErrorEnum_AccountBudgetProposalError_value = map[string]int32{
	"UNSPECIFIED":                     0,
	"UNKNOWN":                         1,
	"FIELD_MASK_NOT_ALLOWED":          2,
	"IMMUTABLE_FIELD":                 3,
	"REQUIRED_FIELD_MISSING":          4,
	"CANNOT_CANCEL_APPROVED_PROPOSAL": 5,
	"CANNOT_REMOVE_UNAPPROVED_BUDGET": 6,
	"CANNOT_REMOVE_RUNNING_BUDGET":    7,
	"CANNOT_END_UNAPPROVED_BUDGET":    8,
	"CANNOT_END_INACTIVE_BUDGET":      9,
	"BUDGET_NAME_REQUIRED":            10,
	"CANNOT_UPDATE_OLD_BUDGET":        11,
	"CANNOT_END_IN_PAST":              12,
	"CANNOT_EXTEND_END_TIME":          13,
	"PURCHASE_ORDER_NUMBER_REQUIRED":  14,
	"PENDING_UPDATE_PROPOSAL_EXISTS":  15,
	"MULTIPLE_BUDGETS_NOT_ALLOWED_FOR_UNAPPROVED_BILLING_SETUP": 16,
	"CANNOT_UPDATE_START_TIME_FOR_STARTED_BUDGET":               17,
	"SPENDING_LIMIT_LOWER_THAN_ACCRUED_COST_NOT_ALLOWED":        18,
	"UPDATE_IS_NO_OP":                                   19,
	"END_TIME_MUST_FOLLOW_START_TIME":                   20,
	"BUDGET_DATE_RANGE_INCOMPATIBLE_WITH_BILLING_SETUP": 21,
	"NOT_AUTHORIZED":                                    22,
	"INVALID_BILLING_SETUP":                             23,
}

func (x AccountBudgetProposalErrorEnum_AccountBudgetProposalError) String() string {
	return proto.EnumName(AccountBudgetProposalErrorEnum_AccountBudgetProposalError_name, int32(x))
}
func (AccountBudgetProposalErrorEnum_AccountBudgetProposalError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_account_budget_proposal_error_b048cd2ed2033a8f, []int{0, 0}
}

// Container for enum describing possible account budget proposal errors.
type AccountBudgetProposalErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountBudgetProposalErrorEnum) Reset()         { *m = AccountBudgetProposalErrorEnum{} }
func (m *AccountBudgetProposalErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AccountBudgetProposalErrorEnum) ProtoMessage()    {}
func (*AccountBudgetProposalErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_budget_proposal_error_b048cd2ed2033a8f, []int{0}
}
func (m *AccountBudgetProposalErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountBudgetProposalErrorEnum.Unmarshal(m, b)
}
func (m *AccountBudgetProposalErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountBudgetProposalErrorEnum.Marshal(b, m, deterministic)
}
func (dst *AccountBudgetProposalErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountBudgetProposalErrorEnum.Merge(dst, src)
}
func (m *AccountBudgetProposalErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AccountBudgetProposalErrorEnum.Size(m)
}
func (m *AccountBudgetProposalErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountBudgetProposalErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AccountBudgetProposalErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AccountBudgetProposalErrorEnum)(nil), "google.ads.googleads.v0.errors.AccountBudgetProposalErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.AccountBudgetProposalErrorEnum_AccountBudgetProposalError", AccountBudgetProposalErrorEnum_AccountBudgetProposalError_name, AccountBudgetProposalErrorEnum_AccountBudgetProposalError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/account_budget_proposal_error.proto", fileDescriptor_account_budget_proposal_error_b048cd2ed2033a8f)
}

var fileDescriptor_account_budget_proposal_error_b048cd2ed2033a8f = []byte{
	// 670 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xcb, 0x6e, 0x13, 0x3d,
	0x14, 0xc7, 0xbf, 0xa6, 0x1f, 0x29, 0xb8, 0xd0, 0x1a, 0xf7, 0x42, 0xa9, 0xaa, 0x14, 0x85, 0x25,
	0xd2, 0x24, 0x80, 0x40, 0x62, 0x10, 0x0b, 0x67, 0xc6, 0x49, 0xac, 0xce, 0xd8, 0xc6, 0x97, 0xb4,
	0xaa, 0x22, 0x59, 0x69, 0x13, 0x45, 0x48, 0x6d, 0x26, 0xca, 0xb4, 0x7d, 0x20, 0x96, 0xec, 0x78,
	0x0d, 0xd6, 0x3c, 0x05, 0x8f, 0xc0, 0x0a, 0x79, 0x2e, 0x69, 0x52, 0x51, 0x16, 0x51, 0xce, 0xcc,
	0xf9, 0x9d, 0xff, 0xb9, 0xd8, 0x67, 0x40, 0x6b, 0x9c, 0x24, 0xe3, 0x8b, 0x51, 0x63, 0x30, 0x4c,
	0x1b, 0xb9, 0xe9, 0xac, 0x9b, 0x66, 0x63, 0x34, 0x9b, 0x25, 0xb3, 0xb4, 0x31, 0x38, 0x3f, 0x4f,
	0xae, 0x27, 0x57, 0xf6, 0xec, 0x7a, 0x38, 0x1e, 0x5d, 0xd9, 0xe9, 0x2c, 0x99, 0x26, 0xe9, 0xe0,
	0xc2, 0x66, 0x6e, 0x6f, 0x3a, 0x4b, 0xae, 0x12, 0x54, 0xcb, 0x03, 0xbd, 0xc1, 0x30, 0xf5, 0xe6,
	0x1a, 0xde, 0x4d, 0xd3, 0xcb, 0x35, 0xea, 0x3f, 0xab, 0xa0, 0x86, 0x73, 0x9d, 0x56, 0x26, 0x23,
	0x0a, 0x15, 0xe2, 0xfc, 0x64, 0x72, 0x7d, 0x59, 0xff, 0x5e, 0x05, 0xfb, 0xf7, 0x23, 0x68, 0x13,
	0xac, 0x1b, 0xa6, 0x04, 0x09, 0x68, 0x9b, 0x92, 0x10, 0xfe, 0x87, 0xd6, 0xc1, 0x9a, 0x61, 0x47,
	0x8c, 0x1f, 0x33, 0xb8, 0x82, 0xf6, 0xc1, 0x6e, 0x9b, 0x92, 0x28, 0xb4, 0x31, 0x56, 0x47, 0x96,
	0x71, 0x6d, 0x71, 0x14, 0xf1, 0x63, 0x12, 0xc2, 0x0a, 0xda, 0x02, 0x9b, 0x34, 0x8e, 0x8d, 0xc6,
	0xad, 0x88, 0xd8, 0x8c, 0x82, 0xab, 0x2e, 0x40, 0x92, 0xcf, 0x86, 0x4a, 0x12, 0xda, 0x22, 0x92,
	0x2a, 0x45, 0x59, 0x07, 0xfe, 0x8f, 0x5e, 0x82, 0xc3, 0x00, 0x33, 0x27, 0x12, 0x60, 0x16, 0x90,
	0xc8, 0x62, 0x21, 0x24, 0xef, 0x91, 0xd0, 0x0a, 0xc9, 0x05, 0x57, 0x38, 0x82, 0x0f, 0x16, 0x20,
	0x49, 0x62, 0xde, 0x23, 0xd6, 0xb0, 0x39, 0xd6, 0x32, 0x61, 0x87, 0x68, 0x58, 0x45, 0x2f, 0xc0,
	0xc1, 0x32, 0x24, 0x0d, 0x63, 0x94, 0x75, 0x4a, 0x62, 0x6d, 0x81, 0x20, 0x2c, 0xfc, 0x8b, 0xc6,
	0x43, 0x54, 0x03, 0xfb, 0x0b, 0x04, 0x65, 0x38, 0xd0, 0xb4, 0x47, 0x4a, 0xff, 0x23, 0xb4, 0x07,
	0xb6, 0x73, 0xdb, 0x32, 0x1c, 0x13, 0x5b, 0x76, 0x05, 0x01, 0x3a, 0x00, 0x7b, 0x45, 0xa4, 0x11,
	0x21, 0xd6, 0xc4, 0xf2, 0x68, 0xae, 0xbb, 0x8e, 0x76, 0x01, 0x5a, 0xd2, 0xb5, 0x02, 0x2b, 0x0d,
	0x1f, 0xbb, 0xc9, 0x94, 0xef, 0x4f, 0xb4, 0x73, 0xb9, 0x9f, 0xa6, 0x31, 0x81, 0x4f, 0x50, 0x1d,
	0xd4, 0x84, 0x91, 0x41, 0x17, 0x2b, 0x62, 0xb9, 0x0c, 0x89, 0xb4, 0xcc, 0xc4, 0x2d, 0x22, 0x6f,
	0xb3, 0x6e, 0x64, 0x0c, 0x61, 0xa1, 0xeb, 0xb2, 0x48, 0x5b, 0x4e, 0xcd, 0x92, 0x13, 0xaa, 0xb4,
	0x82, 0x9b, 0xe8, 0x13, 0xf8, 0x10, 0x9b, 0x48, 0x53, 0x11, 0x95, 0x8d, 0xa8, 0xc5, 0x43, 0xb3,
	0x6d, 0x2e, 0x97, 0x66, 0x41, 0xa3, 0xc8, 0xe9, 0x29, 0xa2, 0x8d, 0x80, 0x10, 0x35, 0xc0, 0xab,
	0xe5, 0xc6, 0x94, 0xc6, 0x52, 0x67, 0x45, 0x66, 0xa1, 0xd9, 0xe3, 0xed, 0x0c, 0x9f, 0xa2, 0xf7,
	0xe0, 0x8d, 0x2a, 0x8b, 0x8a, 0x68, 0x4c, 0xb5, 0x75, 0x79, 0xa4, 0xd5, 0x5d, 0xcc, 0x2c, 0x0e,
	0x02, 0x69, 0x48, 0x68, 0x03, 0xae, 0xf4, 0xd2, 0xd5, 0x41, 0xee, 0xea, 0x14, 0x19, 0xa8, 0x2b,
	0xd0, 0x72, 0x01, 0xb7, 0xdc, 0xc9, 0x97, 0x23, 0xb1, 0xb1, 0x51, 0xda, 0xb6, 0xb9, 0x0b, 0x58,
	0x28, 0x02, 0x6e, 0xa3, 0x77, 0xe0, 0x75, 0x71, 0x2a, 0x59, 0xb8, 0xc4, 0xac, 0x43, 0x2c, 0x65,
	0x01, 0x8f, 0x05, 0xd6, 0xd4, 0xdd, 0xc4, 0x63, 0xaa, 0xbb, 0x77, 0x3a, 0xdb, 0x41, 0x08, 0x6c,
	0x64, 0x15, 0x18, 0xdd, 0xe5, 0x92, 0x9e, 0x92, 0x10, 0xee, 0xa2, 0xe7, 0x60, 0x87, 0xb2, 0x1e,
	0x8e, 0xe8, 0xdd, 0x41, 0x3c, 0x6b, 0xfd, 0x5e, 0x01, 0xf5, 0xf3, 0xe4, 0xd2, 0xfb, 0xf7, 0xf6,
	0xb5, 0x0e, 0xef, 0xdf, 0x2b, 0xe1, 0xd6, 0x57, 0xac, 0x9c, 0x86, 0x85, 0xc4, 0x38, 0xb9, 0x18,
	0x4c, 0xc6, 0x5e, 0x32, 0x1b, 0x37, 0xc6, 0xa3, 0x49, 0xb6, 0xdc, 0xe5, 0x47, 0x61, 0xfa, 0x25,
	0xbd, 0xef, 0x1b, 0xf1, 0x31, 0xff, 0xfb, 0x5a, 0x59, 0xed, 0x60, 0xfc, 0xad, 0x52, 0xeb, 0xe4,
	0x62, 0x78, 0x98, 0x7a, 0xb9, 0xe9, 0xac, 0x5e, 0xd3, 0xcb, 0x52, 0xa6, 0x3f, 0x4a, 0xa0, 0x8f,
	0x87, 0x69, 0x7f, 0x0e, 0xf4, 0x7b, 0xcd, 0x7e, 0x0e, 0xfc, 0xaa, 0xd4, 0xf3, 0xb7, 0xbe, 0x8f,
	0x87, 0xa9, 0xef, 0xcf, 0x11, 0xdf, 0xef, 0x35, 0x7d, 0x3f, 0x87, 0xce, 0xaa, 0x59, 0x75, 0x6f,
	0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x29, 0x85, 0x0d, 0xc0, 0x04, 0x00, 0x00,
}
