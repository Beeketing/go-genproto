// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/query_error.proto

package errors // import "github.com/Beeketing/genproto/googleapis/ads/googleads/v0/errors"

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

// Enum describing possible query errors.
type QueryErrorEnum_QueryError int32

const (
	// Name unspecified.
	QueryErrorEnum_UNSPECIFIED QueryErrorEnum_QueryError = 0
	// The received error code is not known in this version.
	QueryErrorEnum_UNKNOWN QueryErrorEnum_QueryError = 1
	// Returned if all other query error reasons are not applicable.
	QueryErrorEnum_QUERY_ERROR QueryErrorEnum_QueryError = 50
	// A condition used in the query references an invalid enum constant.
	QueryErrorEnum_BAD_ENUM_CONSTANT QueryErrorEnum_QueryError = 18
	// Query contains an invalid escape sequence.
	QueryErrorEnum_BAD_ESCAPE_SEQUENCE QueryErrorEnum_QueryError = 7
	// Field name is invalid.
	QueryErrorEnum_BAD_FIELD_NAME QueryErrorEnum_QueryError = 12
	// Limit value is invalid (i.e. not a number)
	QueryErrorEnum_BAD_LIMIT_VALUE QueryErrorEnum_QueryError = 15
	// Encountered number can not be parsed.
	QueryErrorEnum_BAD_NUMBER QueryErrorEnum_QueryError = 5
	// Invalid operator encountered.
	QueryErrorEnum_BAD_OPERATOR QueryErrorEnum_QueryError = 3
	// Invalid resource type was specified in the FROM clause.
	QueryErrorEnum_BAD_RESOURCE_TYPE_IN_FROM_CLAUSE QueryErrorEnum_QueryError = 45
	// Non-ASCII symbol encountered outside of strings.
	QueryErrorEnum_BAD_SYMBOL QueryErrorEnum_QueryError = 2
	// Value is invalid.
	QueryErrorEnum_BAD_VALUE QueryErrorEnum_QueryError = 4
	// Date filters fail to restrict date to a range smaller than 31 days.
	// Applicable if the query is segmented by date.
	QueryErrorEnum_DATE_RANGE_TOO_WIDE QueryErrorEnum_QueryError = 36
	// Expected AND between values with BETWEEN operator.
	QueryErrorEnum_EXPECTED_AND QueryErrorEnum_QueryError = 30
	// Expecting ORDER BY to have BY.
	QueryErrorEnum_EXPECTED_BY QueryErrorEnum_QueryError = 14
	// There was no dimension field selected.
	QueryErrorEnum_EXPECTED_DIMENSION_FIELD_IN_SELECT_CLAUSE QueryErrorEnum_QueryError = 37
	// Missing filters on date related fields.
	QueryErrorEnum_EXPECTED_FILTERS_ON_DATE_RANGE QueryErrorEnum_QueryError = 55
	// Missing FROM clause.
	QueryErrorEnum_EXPECTED_FROM QueryErrorEnum_QueryError = 44
	// The operator used in the conditions requires the value to be a list.
	QueryErrorEnum_EXPECTED_LIST QueryErrorEnum_QueryError = 41
	// Fields used in WHERE or ORDER BY clauses are missing from the SELECT
	// clause.
	QueryErrorEnum_EXPECTED_REFERENCED_FIELD_IN_SELECT_CLAUSE QueryErrorEnum_QueryError = 16
	// SELECT is missing at the beginning of query.
	QueryErrorEnum_EXPECTED_SELECT QueryErrorEnum_QueryError = 13
	// A list was passed as a value to a condition whose operator expects a
	// single value.
	QueryErrorEnum_EXPECTED_SINGLE_VALUE QueryErrorEnum_QueryError = 42
	// Missing one or both values with BETWEEN operator.
	QueryErrorEnum_EXPECTED_VALUE_WITH_BETWEEN_OPERATOR QueryErrorEnum_QueryError = 29
	// Invalid date format. Expected 'YYYY-MM-DD'.
	QueryErrorEnum_INVALID_DATE_FORMAT QueryErrorEnum_QueryError = 38
	// Value passed was not a string when it should have been. I.e., it was a
	// number or unquoted literal.
	QueryErrorEnum_INVALID_STRING_VALUE QueryErrorEnum_QueryError = 57
	// A String value passed to the BETWEEN operator does not parse as a date.
	QueryErrorEnum_INVALID_VALUE_WITH_BETWEEN_OPERATOR QueryErrorEnum_QueryError = 26
	// The value passed to the DURING operator is not a Date range literal
	QueryErrorEnum_INVALID_VALUE_WITH_DURING_OPERATOR QueryErrorEnum_QueryError = 22
	// A non-string value was passed to the LIKE operator.
	QueryErrorEnum_INVALID_VALUE_WITH_LIKE_OPERATOR QueryErrorEnum_QueryError = 56
	// An operator was provided that is inapplicable to the field being
	// filtered.
	QueryErrorEnum_OPERATOR_FIELD_MISMATCH QueryErrorEnum_QueryError = 35
	// A Condition was found with an empty list.
	QueryErrorEnum_PROHIBITED_EMPTY_LIST_IN_CONDITION QueryErrorEnum_QueryError = 28
	// A condition used in the query references an unsupported enum constant.
	QueryErrorEnum_PROHIBITED_ENUM_CONSTANT QueryErrorEnum_QueryError = 54
	// Fields that are not allowed to be selected together were included in
	// the SELECT clause.
	QueryErrorEnum_PROHIBITED_FIELD_COMBINATION_IN_SELECT_CLAUSE QueryErrorEnum_QueryError = 31
	// A field that is not orderable was included in the ORDER BY clause.
	QueryErrorEnum_PROHIBITED_FIELD_IN_ORDER_BY_CLAUSE QueryErrorEnum_QueryError = 40
	// A field that is not selectable was included in the SELECT clause.
	QueryErrorEnum_PROHIBITED_FIELD_IN_SELECT_CLAUSE QueryErrorEnum_QueryError = 23
	// A field that is not filterable was included in the WHERE clause.
	QueryErrorEnum_PROHIBITED_FIELD_IN_WHERE_CLAUSE QueryErrorEnum_QueryError = 24
	// Resource type specified in the FROM clause is not supported by this
	// service.
	QueryErrorEnum_PROHIBITED_RESOURCE_TYPE_IN_FROM_CLAUSE QueryErrorEnum_QueryError = 43
	// A field that comes from an incompatible resource was included in the
	// SELECT clause.
	QueryErrorEnum_PROHIBITED_RESOURCE_TYPE_IN_SELECT_CLAUSE QueryErrorEnum_QueryError = 48
	// A field that comes from an incompatible resource was included in the
	// WHERE clause.
	QueryErrorEnum_PROHIBITED_RESOURCE_TYPE_IN_WHERE_CLAUSE QueryErrorEnum_QueryError = 58
	// A metric incompatible with the main resource or other selected
	// segmenting resources was included in the SELECT or WHERE clause.
	QueryErrorEnum_PROHIBITED_METRIC_IN_SELECT_OR_WHERE_CLAUSE QueryErrorEnum_QueryError = 49
	// A segment incompatible with the main resource or other selected
	// segmenting resources was included in the SELECT or WHERE clause.
	QueryErrorEnum_PROHIBITED_SEGMENT_IN_SELECT_OR_WHERE_CLAUSE QueryErrorEnum_QueryError = 51
	// A segment in the SELECT clause is incompatible with a metric in the
	// SELECT or WHERE clause.
	QueryErrorEnum_PROHIBITED_SEGMENT_WITH_METRIC_IN_SELECT_OR_WHERE_CLAUSE QueryErrorEnum_QueryError = 53
	// The value passed to the limit clause is too low.
	QueryErrorEnum_LIMIT_VALUE_TOO_LOW QueryErrorEnum_QueryError = 25
	// Query has a string containing a newline character.
	QueryErrorEnum_PROHIBITED_NEWLINE_IN_STRING QueryErrorEnum_QueryError = 8
	// List contains values of different types.
	QueryErrorEnum_PROHIBITED_VALUE_COMBINATION_IN_LIST QueryErrorEnum_QueryError = 10
	// The values passed to the BETWEEN operator are not of the same type.
	QueryErrorEnum_PROHIBITED_VALUE_COMBINATION_WITH_BETWEEN_OPERATOR QueryErrorEnum_QueryError = 21
	// Query contains unterminated string.
	QueryErrorEnum_STRING_NOT_TERMINATED QueryErrorEnum_QueryError = 6
	// Too many segments are specified in SELECT clause.
	QueryErrorEnum_TOO_MANY_SEGMENTS QueryErrorEnum_QueryError = 34
	// Query is incomplete and cannot be parsed.
	QueryErrorEnum_UNEXPECTED_END_OF_QUERY QueryErrorEnum_QueryError = 9
	// FROM clause cannot be specified in this query.
	QueryErrorEnum_UNEXPECTED_FROM_CLAUSE QueryErrorEnum_QueryError = 47
	// Query contains one or more unrecognized fields.
	QueryErrorEnum_UNRECOGNIZED_FIELD QueryErrorEnum_QueryError = 32
	// Query has an unexpected extra part.
	QueryErrorEnum_UNEXPECTED_INPUT QueryErrorEnum_QueryError = 11
	// Metrics cannot be requested for a manager account. To retrieve metrics,
	// issue separate requests against each client account under the manager
	// account.
	QueryErrorEnum_REQUESTED_METRICS_FOR_MANAGER QueryErrorEnum_QueryError = 59
)

var QueryErrorEnum_QueryError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	50: "QUERY_ERROR",
	18: "BAD_ENUM_CONSTANT",
	7:  "BAD_ESCAPE_SEQUENCE",
	12: "BAD_FIELD_NAME",
	15: "BAD_LIMIT_VALUE",
	5:  "BAD_NUMBER",
	3:  "BAD_OPERATOR",
	45: "BAD_RESOURCE_TYPE_IN_FROM_CLAUSE",
	2:  "BAD_SYMBOL",
	4:  "BAD_VALUE",
	36: "DATE_RANGE_TOO_WIDE",
	30: "EXPECTED_AND",
	14: "EXPECTED_BY",
	37: "EXPECTED_DIMENSION_FIELD_IN_SELECT_CLAUSE",
	55: "EXPECTED_FILTERS_ON_DATE_RANGE",
	44: "EXPECTED_FROM",
	41: "EXPECTED_LIST",
	16: "EXPECTED_REFERENCED_FIELD_IN_SELECT_CLAUSE",
	13: "EXPECTED_SELECT",
	42: "EXPECTED_SINGLE_VALUE",
	29: "EXPECTED_VALUE_WITH_BETWEEN_OPERATOR",
	38: "INVALID_DATE_FORMAT",
	57: "INVALID_STRING_VALUE",
	26: "INVALID_VALUE_WITH_BETWEEN_OPERATOR",
	22: "INVALID_VALUE_WITH_DURING_OPERATOR",
	56: "INVALID_VALUE_WITH_LIKE_OPERATOR",
	35: "OPERATOR_FIELD_MISMATCH",
	28: "PROHIBITED_EMPTY_LIST_IN_CONDITION",
	54: "PROHIBITED_ENUM_CONSTANT",
	31: "PROHIBITED_FIELD_COMBINATION_IN_SELECT_CLAUSE",
	40: "PROHIBITED_FIELD_IN_ORDER_BY_CLAUSE",
	23: "PROHIBITED_FIELD_IN_SELECT_CLAUSE",
	24: "PROHIBITED_FIELD_IN_WHERE_CLAUSE",
	43: "PROHIBITED_RESOURCE_TYPE_IN_FROM_CLAUSE",
	48: "PROHIBITED_RESOURCE_TYPE_IN_SELECT_CLAUSE",
	58: "PROHIBITED_RESOURCE_TYPE_IN_WHERE_CLAUSE",
	49: "PROHIBITED_METRIC_IN_SELECT_OR_WHERE_CLAUSE",
	51: "PROHIBITED_SEGMENT_IN_SELECT_OR_WHERE_CLAUSE",
	53: "PROHIBITED_SEGMENT_WITH_METRIC_IN_SELECT_OR_WHERE_CLAUSE",
	25: "LIMIT_VALUE_TOO_LOW",
	8:  "PROHIBITED_NEWLINE_IN_STRING",
	10: "PROHIBITED_VALUE_COMBINATION_IN_LIST",
	21: "PROHIBITED_VALUE_COMBINATION_WITH_BETWEEN_OPERATOR",
	6:  "STRING_NOT_TERMINATED",
	34: "TOO_MANY_SEGMENTS",
	9:  "UNEXPECTED_END_OF_QUERY",
	47: "UNEXPECTED_FROM_CLAUSE",
	32: "UNRECOGNIZED_FIELD",
	11: "UNEXPECTED_INPUT",
	59: "REQUESTED_METRICS_FOR_MANAGER",
}
var QueryErrorEnum_QueryError_value = map[string]int32{
	"UNSPECIFIED":                      0,
	"UNKNOWN":                          1,
	"QUERY_ERROR":                      50,
	"BAD_ENUM_CONSTANT":                18,
	"BAD_ESCAPE_SEQUENCE":              7,
	"BAD_FIELD_NAME":                   12,
	"BAD_LIMIT_VALUE":                  15,
	"BAD_NUMBER":                       5,
	"BAD_OPERATOR":                     3,
	"BAD_RESOURCE_TYPE_IN_FROM_CLAUSE": 45,
	"BAD_SYMBOL":                       2,
	"BAD_VALUE":                        4,
	"DATE_RANGE_TOO_WIDE":              36,
	"EXPECTED_AND":                     30,
	"EXPECTED_BY":                      14,
	"EXPECTED_DIMENSION_FIELD_IN_SELECT_CLAUSE":                37,
	"EXPECTED_FILTERS_ON_DATE_RANGE":                           55,
	"EXPECTED_FROM":                                            44,
	"EXPECTED_LIST":                                            41,
	"EXPECTED_REFERENCED_FIELD_IN_SELECT_CLAUSE":               16,
	"EXPECTED_SELECT":                                          13,
	"EXPECTED_SINGLE_VALUE":                                    42,
	"EXPECTED_VALUE_WITH_BETWEEN_OPERATOR":                     29,
	"INVALID_DATE_FORMAT":                                      38,
	"INVALID_STRING_VALUE":                                     57,
	"INVALID_VALUE_WITH_BETWEEN_OPERATOR":                      26,
	"INVALID_VALUE_WITH_DURING_OPERATOR":                       22,
	"INVALID_VALUE_WITH_LIKE_OPERATOR":                         56,
	"OPERATOR_FIELD_MISMATCH":                                  35,
	"PROHIBITED_EMPTY_LIST_IN_CONDITION":                       28,
	"PROHIBITED_ENUM_CONSTANT":                                 54,
	"PROHIBITED_FIELD_COMBINATION_IN_SELECT_CLAUSE":            31,
	"PROHIBITED_FIELD_IN_ORDER_BY_CLAUSE":                      40,
	"PROHIBITED_FIELD_IN_SELECT_CLAUSE":                        23,
	"PROHIBITED_FIELD_IN_WHERE_CLAUSE":                         24,
	"PROHIBITED_RESOURCE_TYPE_IN_FROM_CLAUSE":                  43,
	"PROHIBITED_RESOURCE_TYPE_IN_SELECT_CLAUSE":                48,
	"PROHIBITED_RESOURCE_TYPE_IN_WHERE_CLAUSE":                 58,
	"PROHIBITED_METRIC_IN_SELECT_OR_WHERE_CLAUSE":              49,
	"PROHIBITED_SEGMENT_IN_SELECT_OR_WHERE_CLAUSE":             51,
	"PROHIBITED_SEGMENT_WITH_METRIC_IN_SELECT_OR_WHERE_CLAUSE": 53,
	"LIMIT_VALUE_TOO_LOW":                                      25,
	"PROHIBITED_NEWLINE_IN_STRING":                             8,
	"PROHIBITED_VALUE_COMBINATION_IN_LIST":                     10,
	"PROHIBITED_VALUE_COMBINATION_WITH_BETWEEN_OPERATOR":       21,
	"STRING_NOT_TERMINATED":                                    6,
	"TOO_MANY_SEGMENTS":                                        34,
	"UNEXPECTED_END_OF_QUERY":                                  9,
	"UNEXPECTED_FROM_CLAUSE":                                   47,
	"UNRECOGNIZED_FIELD":                                       32,
	"UNEXPECTED_INPUT":                                         11,
	"REQUESTED_METRICS_FOR_MANAGER":                            59,
}

func (x QueryErrorEnum_QueryError) String() string {
	return proto.EnumName(QueryErrorEnum_QueryError_name, int32(x))
}
func (QueryErrorEnum_QueryError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_query_error_435ab5df2f507e98, []int{0, 0}
}

// Container for enum describing possible query errors.
type QueryErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryErrorEnum) Reset()         { *m = QueryErrorEnum{} }
func (m *QueryErrorEnum) String() string { return proto.CompactTextString(m) }
func (*QueryErrorEnum) ProtoMessage()    {}
func (*QueryErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_error_435ab5df2f507e98, []int{0}
}
func (m *QueryErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryErrorEnum.Unmarshal(m, b)
}
func (m *QueryErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryErrorEnum.Marshal(b, m, deterministic)
}
func (dst *QueryErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryErrorEnum.Merge(dst, src)
}
func (m *QueryErrorEnum) XXX_Size() int {
	return xxx_messageInfo_QueryErrorEnum.Size(m)
}
func (m *QueryErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_QueryErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryErrorEnum)(nil), "google.ads.googleads.v0.errors.QueryErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.QueryErrorEnum_QueryError", QueryErrorEnum_QueryError_name, QueryErrorEnum_QueryError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/query_error.proto", fileDescriptor_query_error_435ab5df2f507e98)
}

var fileDescriptor_query_error_435ab5df2f507e98 = []byte{
	// 930 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xef, 0x72, 0xd3, 0x46,
	0x10, 0x6f, 0x42, 0x0b, 0x65, 0xf3, 0xef, 0x38, 0xc8, 0x1f, 0x68, 0x48, 0x83, 0x09, 0x10, 0x48,
	0x22, 0x1b, 0x98, 0x52, 0x2a, 0xfa, 0xe5, 0x24, 0xad, 0x9d, 0x1b, 0xa4, 0x3b, 0xe5, 0x74, 0xb2,
	0x6b, 0x26, 0x33, 0x37, 0x69, 0x9d, 0xf1, 0x74, 0x06, 0x62, 0x6a, 0x17, 0x66, 0xfa, 0x04, 0x7d,
	0x8f, 0x7e, 0xec, 0xa3, 0xf4, 0x51, 0xf8, 0xdc, 0x07, 0xe8, 0x9c, 0x64, 0xc9, 0x72, 0x31, 0xe6,
	0x93, 0xee, 0x7e, 0xfb, 0xfb, 0xed, 0xee, 0xad, 0x76, 0xef, 0xa0, 0xd1, 0x1f, 0x0c, 0xfa, 0xaf,
	0xcf, 0xeb, 0x67, 0xbd, 0x51, 0x3d, 0x5f, 0xda, 0xd5, 0xfb, 0x46, 0xfd, 0x7c, 0x38, 0x1c, 0x0c,
	0x47, 0xf5, 0xdf, 0xde, 0x9d, 0x0f, 0xff, 0x30, 0xd9, 0xc6, 0x79, 0x3b, 0x1c, 0xfc, 0x3e, 0xa0,
	0x3b, 0x39, 0xcd, 0x39, 0xeb, 0x8d, 0x9c, 0x52, 0xe1, 0xbc, 0x6f, 0x38, 0xb9, 0xa2, 0xf6, 0xe7,
	0x0a, 0xac, 0x9e, 0x58, 0x15, 0xda, 0x3d, 0x5e, 0xbc, 0x7b, 0x53, 0xfb, 0x77, 0x19, 0x60, 0x02,
	0xd1, 0x35, 0x58, 0x4a, 0x45, 0x12, 0xa3, 0xcf, 0x9b, 0x1c, 0x03, 0xf2, 0x05, 0x5d, 0x82, 0x2b,
	0xa9, 0x78, 0x29, 0x64, 0x47, 0x90, 0x05, 0x6b, 0x3d, 0x49, 0x51, 0x75, 0x0d, 0x2a, 0x25, 0x15,
	0x79, 0x42, 0xd7, 0xe1, 0x9a, 0xc7, 0x02, 0x83, 0x22, 0x8d, 0x8c, 0x2f, 0x45, 0xa2, 0x99, 0xd0,
	0x84, 0xd2, 0x4d, 0xb8, 0x9e, 0xc1, 0x89, 0xcf, 0x62, 0x34, 0x09, 0x9e, 0xa4, 0x28, 0x7c, 0x24,
	0x57, 0x28, 0x85, 0x55, 0x6b, 0x68, 0x72, 0x0c, 0x03, 0x23, 0x58, 0x84, 0x64, 0x99, 0x5e, 0x87,
	0x35, 0x8b, 0x85, 0x3c, 0xe2, 0xda, 0xb4, 0x59, 0x98, 0x22, 0x59, 0xa3, 0xab, 0x00, 0x16, 0x14,
	0x69, 0xe4, 0xa1, 0x22, 0x5f, 0x51, 0x02, 0xcb, 0x76, 0x2f, 0x63, 0x54, 0x4c, 0x4b, 0x45, 0x2e,
	0xd1, 0x3d, 0xd8, 0xb5, 0x88, 0xc2, 0x44, 0xa6, 0xca, 0x47, 0xa3, 0xbb, 0x31, 0x1a, 0x2e, 0x4c,
	0x53, 0xc9, 0xc8, 0xf8, 0x21, 0x4b, 0x13, 0x24, 0x47, 0x85, 0x9f, 0xa4, 0x1b, 0x79, 0x32, 0x24,
	0x8b, 0x74, 0x05, 0xae, 0xda, 0x7d, 0x1e, 0xe6, 0x4b, 0x9b, 0x68, 0xc0, 0x34, 0x1a, 0xc5, 0x44,
	0x0b, 0x8d, 0x96, 0xd2, 0x74, 0x78, 0x80, 0x64, 0xcf, 0xc6, 0xc3, 0x9f, 0x62, 0xf4, 0x35, 0x06,
	0x86, 0x89, 0x80, 0xec, 0xd8, 0xb3, 0x97, 0x88, 0xd7, 0x25, 0xab, 0xf4, 0x08, 0x1e, 0x96, 0x40,
	0xc0, 0x23, 0x14, 0x09, 0x97, 0x62, 0x7c, 0x34, 0x2e, 0x4c, 0x82, 0x21, 0xfa, 0xba, 0xc8, 0xe4,
	0x1e, 0xad, 0xc1, 0x4e, 0x49, 0x6f, 0xf2, 0x50, 0xa3, 0x4a, 0x8c, 0x14, 0x66, 0x12, 0x9e, 0x7c,
	0x4f, 0xaf, 0xc1, 0xca, 0x84, 0xa3, 0x64, 0x44, 0x0e, 0xa7, 0xa0, 0x90, 0x27, 0x9a, 0x3c, 0xa4,
	0x0e, 0x3c, 0x2a, 0x21, 0x85, 0x4d, 0x54, 0xb6, 0xb8, 0xc1, 0xa7, 0x22, 0x13, 0x5b, 0xe0, 0x92,
	0x9f, 0xdb, 0xc8, 0x0a, 0xbd, 0x09, 0xeb, 0x13, 0x90, 0x8b, 0x56, 0x88, 0xe3, 0xa2, 0x3c, 0xa2,
	0xfb, 0xb0, 0x57, 0x9a, 0x32, 0xcc, 0x74, 0xb8, 0x3e, 0x36, 0x1e, 0xea, 0x0e, 0xa2, 0x98, 0xfc,
	0x83, 0xdb, 0xb6, 0x7c, 0x5c, 0xb4, 0x59, 0xc8, 0x83, 0xfc, 0x1c, 0x4d, 0xa9, 0x22, 0xa6, 0xc9,
	0x7d, 0xba, 0x05, 0x37, 0x0a, 0x43, 0xa2, 0x15, 0x17, 0xad, 0xb1, 0xf3, 0x1f, 0xe8, 0x03, 0xb8,
	0x5b, 0x58, 0xe6, 0xf9, 0xbe, 0x45, 0xef, 0x43, 0x6d, 0x06, 0x31, 0x48, 0x33, 0x6f, 0x25, 0x6f,
	0xc3, 0xf6, 0xc1, 0x0c, 0x5e, 0xc8, 0x5f, 0xe2, 0x84, 0xf5, 0x9c, 0x7e, 0x03, 0x9b, 0xc5, 0x6e,
	0x5c, 0xa8, 0x88, 0x27, 0x11, 0xd3, 0xfe, 0x31, 0xb9, 0x6b, 0x43, 0xc5, 0x4a, 0x1e, 0x73, 0x8f,
	0xdb, 0x23, 0x63, 0x14, 0xeb, 0x6e, 0x56, 0x6b, 0x5b, 0x4c, 0x5f, 0x8a, 0x80, 0x6b, 0x2e, 0x05,
	0xd9, 0xa6, 0xdb, 0xb0, 0x55, 0xe5, 0x4d, 0x35, 0xfd, 0x33, 0xfa, 0x18, 0x8e, 0x2a, 0xd6, 0x3c,
	0x88, 0x2f, 0x23, 0x8f, 0x0b, 0x66, 0xf5, 0x1f, 0xff, 0x99, 0x6f, 0x6d, 0x31, 0x3e, 0x92, 0x70,
	0x61, 0xa4, 0x0a, 0x50, 0x19, 0xaf, 0x5b, 0x10, 0xf7, 0xe9, 0x3d, 0xb8, 0x33, 0x8b, 0x38, 0xed,
	0x6f, 0xd3, 0xd6, 0x62, 0x16, 0xad, 0x73, 0x8c, 0x0a, 0x0b, 0xd6, 0x16, 0x3d, 0x80, 0x07, 0x15,
	0xd6, 0xdc, 0x01, 0x3a, 0xb0, 0x5d, 0x3e, 0x8f, 0x3c, 0x9d, 0x41, 0x83, 0x1e, 0xc2, 0xfe, 0x3c,
	0xfa, 0x54, 0x26, 0x2e, 0xad, 0xc3, 0x41, 0x85, 0x1d, 0xa1, 0x56, 0xdc, 0xaf, 0x78, 0x95, 0x6a,
	0x5a, 0xf0, 0x98, 0x36, 0xe0, 0xb0, 0x22, 0x48, 0xb0, 0x15, 0xa1, 0xd0, 0x73, 0x14, 0x4f, 0xe9,
	0x8f, 0xf0, 0x7c, 0x86, 0x22, 0xeb, 0x91, 0xcf, 0xc6, 0xfb, 0xce, 0x36, 0x78, 0xe5, 0x5e, 0xca,
	0x2e, 0x88, 0x50, 0x76, 0xc8, 0x4d, 0xba, 0x0b, 0xdb, 0x15, 0xb7, 0x02, 0x3b, 0x21, 0x17, 0x79,
	0x41, 0xb2, 0x76, 0x27, 0x5f, 0xdb, 0x29, 0xaa, 0x30, 0x72, 0xfd, 0xff, 0xda, 0x21, 0x9b, 0x67,
	0xa0, 0xcf, 0xe0, 0xc9, 0x5c, 0xe6, 0xec, 0x09, 0x59, 0xb7, 0x23, 0x3c, 0x1e, 0x2e, 0x21, 0xb5,
	0xd1, 0xa8, 0x22, 0xab, 0xc0, 0x80, 0x5c, 0xb6, 0xf7, 0xb2, 0xcd, 0x35, 0x62, 0xa2, 0x5b, 0x9c,
	0x39, 0x21, 0x35, 0x3b, 0x05, 0xa9, 0x28, 0x67, 0x1b, 0x45, 0x60, 0x64, 0xd3, 0x64, 0x37, 0x3a,
	0xb9, 0x4a, 0x6f, 0xc1, 0x46, 0xc5, 0x58, 0xed, 0x82, 0x3a, 0xdd, 0x00, 0x9a, 0x0a, 0x85, 0xbe,
	0x6c, 0x09, 0xfe, 0xaa, 0x68, 0x2d, 0xb2, 0x4b, 0x6f, 0x00, 0xa9, 0x68, 0xb8, 0x88, 0x53, 0x4d,
	0x96, 0xe8, 0x1d, 0xb8, 0xad, 0xec, 0x9d, 0x9f, 0x4c, 0xfe, 0x6a, 0x62, 0xef, 0x06, 0x9b, 0x0f,
	0x6b, 0xa1, 0x22, 0x2f, 0xbc, 0x0f, 0x0b, 0x50, 0xfb, 0x65, 0xf0, 0xc6, 0x99, 0xff, 0x60, 0x79,
	0x6b, 0x93, 0xa7, 0x29, 0xb6, 0x2f, 0x5c, 0xbc, 0xf0, 0x2a, 0x18, 0x4b, 0xfa, 0x83, 0xd7, 0x67,
	0x17, 0x7d, 0x67, 0x30, 0xec, 0xd7, 0xfb, 0xe7, 0x17, 0xd9, 0xfb, 0x57, 0xbc, 0x92, 0x6f, 0x7f,
	0x1d, 0x7d, 0xea, 0xd1, 0x7c, 0x91, 0x7f, 0xfe, 0x5a, 0xbc, 0xd4, 0x62, 0xec, 0xef, 0xc5, 0x9d,
	0x56, 0xee, 0x8c, 0xf5, 0x46, 0x4e, 0xbe, 0xb4, 0xab, 0x76, 0xc3, 0xc9, 0x42, 0x8e, 0xfe, 0x29,
	0x08, 0xa7, 0xac, 0x37, 0x3a, 0x2d, 0x09, 0xa7, 0xed, 0xc6, 0x69, 0x4e, 0xf8, 0xb0, 0x58, 0xcb,
	0x51, 0xd7, 0x65, 0xbd, 0x91, 0xeb, 0x96, 0x14, 0xd7, 0x6d, 0x37, 0x5c, 0x37, 0x27, 0xfd, 0x7c,
	0x39, 0xcb, 0xee, 0xe9, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xc5, 0x95, 0xe1, 0xd1, 0x07,
	0x00, 0x00,
}
