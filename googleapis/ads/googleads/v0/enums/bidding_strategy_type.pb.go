// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/bidding_strategy_type.proto

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

// Enum describing possible bidding strategy types.
type BiddingStrategyTypeEnum_BiddingStrategyType int32

const (
	// Not specified.
	BiddingStrategyTypeEnum_UNSPECIFIED BiddingStrategyTypeEnum_BiddingStrategyType = 0
	// Used for return value only. Represents value unknown in this version.
	BiddingStrategyTypeEnum_UNKNOWN BiddingStrategyTypeEnum_BiddingStrategyType = 1
	// Enhanced CPC is a bidding strategy that raises bids for clicks
	// that seem more likely to lead to a conversion and lowers
	// them for clicks where they seem less likely.
	BiddingStrategyTypeEnum_ENHANCED_CPC BiddingStrategyTypeEnum_BiddingStrategyType = 2
	// Manual click based bidding where user pays per click.
	BiddingStrategyTypeEnum_MANUAL_CPC BiddingStrategyTypeEnum_BiddingStrategyType = 3
	// Manual impression based bidding
	// where user pays per thousand impressions.
	BiddingStrategyTypeEnum_MANUAL_CPM BiddingStrategyTypeEnum_BiddingStrategyType = 4
	// A bidding strategy that pays a configurable amount per video view.
	BiddingStrategyTypeEnum_MANUAL_CPV BiddingStrategyTypeEnum_BiddingStrategyType = 13
	// A bidding strategy that automatically maximizes number of conversions
	// given a daily budget.
	BiddingStrategyTypeEnum_MAXIMIZE_CONVERSIONS BiddingStrategyTypeEnum_BiddingStrategyType = 10
	// An automated bidding strategy that automatically sets bids to maximize
	// revenue while spending your budget.
	BiddingStrategyTypeEnum_MAXIMIZE_CONVERSION_VALUE BiddingStrategyTypeEnum_BiddingStrategyType = 11
	// Page-One Promoted bidding scheme, which sets max cpc bids to
	// target impressions on page one or page one promoted slots on google.com.
	BiddingStrategyTypeEnum_PAGE_ONE_PROMOTED BiddingStrategyTypeEnum_BiddingStrategyType = 5
	// Percent Cpc is bidding strategy where bids are a fraction of the
	// advertised price for some good or service.
	BiddingStrategyTypeEnum_PERCENT_CPC BiddingStrategyTypeEnum_BiddingStrategyType = 12
	// Target CPA is an automated bid strategy that sets bids
	// to help get as many conversions as possible
	// at the target cost-per-acquisition (CPA) you set.
	BiddingStrategyTypeEnum_TARGET_CPA BiddingStrategyTypeEnum_BiddingStrategyType = 6
	// Target CPM is an automated bid strategy that sets bids to help get
	// as many impressions as possible at the target cost per one thousand
	// impressions (CPM) you set.
	BiddingStrategyTypeEnum_TARGET_CPM BiddingStrategyTypeEnum_BiddingStrategyType = 14
	// Target Outrank Share is an automated bidding strategy that sets bids
	// based on the target fraction of auctions where the advertiser
	// should outrank a specific competitor.
	BiddingStrategyTypeEnum_TARGET_OUTRANK_SHARE BiddingStrategyTypeEnum_BiddingStrategyType = 7
	// Target ROAS is an automated bidding strategy
	// that helps you maximize revenue while averaging
	// a specific target Return On Average Spend (ROAS).
	BiddingStrategyTypeEnum_TARGET_ROAS BiddingStrategyTypeEnum_BiddingStrategyType = 8
	// Target Spend is an automated bid strategy that sets your bids
	// to help get as many clicks as possible within your budget.
	BiddingStrategyTypeEnum_TARGET_SPEND BiddingStrategyTypeEnum_BiddingStrategyType = 9
)

var BiddingStrategyTypeEnum_BiddingStrategyType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "ENHANCED_CPC",
	3:  "MANUAL_CPC",
	4:  "MANUAL_CPM",
	13: "MANUAL_CPV",
	10: "MAXIMIZE_CONVERSIONS",
	11: "MAXIMIZE_CONVERSION_VALUE",
	5:  "PAGE_ONE_PROMOTED",
	12: "PERCENT_CPC",
	6:  "TARGET_CPA",
	14: "TARGET_CPM",
	7:  "TARGET_OUTRANK_SHARE",
	8:  "TARGET_ROAS",
	9:  "TARGET_SPEND",
}
var BiddingStrategyTypeEnum_BiddingStrategyType_value = map[string]int32{
	"UNSPECIFIED":               0,
	"UNKNOWN":                   1,
	"ENHANCED_CPC":              2,
	"MANUAL_CPC":                3,
	"MANUAL_CPM":                4,
	"MANUAL_CPV":                13,
	"MAXIMIZE_CONVERSIONS":      10,
	"MAXIMIZE_CONVERSION_VALUE": 11,
	"PAGE_ONE_PROMOTED":         5,
	"PERCENT_CPC":               12,
	"TARGET_CPA":                6,
	"TARGET_CPM":                14,
	"TARGET_OUTRANK_SHARE":      7,
	"TARGET_ROAS":               8,
	"TARGET_SPEND":              9,
}

func (x BiddingStrategyTypeEnum_BiddingStrategyType) String() string {
	return proto.EnumName(BiddingStrategyTypeEnum_BiddingStrategyType_name, int32(x))
}
func (BiddingStrategyTypeEnum_BiddingStrategyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bidding_strategy_type_16e7934fd7b18f2d, []int{0, 0}
}

// Container for enum describing possible bidding strategy types.
type BiddingStrategyTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BiddingStrategyTypeEnum) Reset()         { *m = BiddingStrategyTypeEnum{} }
func (m *BiddingStrategyTypeEnum) String() string { return proto.CompactTextString(m) }
func (*BiddingStrategyTypeEnum) ProtoMessage()    {}
func (*BiddingStrategyTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_strategy_type_16e7934fd7b18f2d, []int{0}
}
func (m *BiddingStrategyTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiddingStrategyTypeEnum.Unmarshal(m, b)
}
func (m *BiddingStrategyTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiddingStrategyTypeEnum.Marshal(b, m, deterministic)
}
func (dst *BiddingStrategyTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiddingStrategyTypeEnum.Merge(dst, src)
}
func (m *BiddingStrategyTypeEnum) XXX_Size() int {
	return xxx_messageInfo_BiddingStrategyTypeEnum.Size(m)
}
func (m *BiddingStrategyTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BiddingStrategyTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BiddingStrategyTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BiddingStrategyTypeEnum)(nil), "google.ads.googleads.v0.enums.BiddingStrategyTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.BiddingStrategyTypeEnum_BiddingStrategyType", BiddingStrategyTypeEnum_BiddingStrategyType_name, BiddingStrategyTypeEnum_BiddingStrategyType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/bidding_strategy_type.proto", fileDescriptor_bidding_strategy_type_16e7934fd7b18f2d)
}

var fileDescriptor_bidding_strategy_type_16e7934fd7b18f2d = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x18, 0x24, 0x2e, 0xb4, 0xb0, 0x29, 0x65, 0x31, 0x20, 0xca, 0x21, 0x87, 0xf6, 0x01, 0xd6, 0x96,
	0x38, 0x61, 0x4e, 0x5f, 0x9c, 0x8f, 0xd4, 0x6a, 0xbd, 0xb6, 0xfc, 0x07, 0xaa, 0x22, 0xad, 0x52,
	0x6c, 0xad, 0x22, 0x35, 0x76, 0x94, 0x75, 0x2b, 0xe5, 0x75, 0x38, 0x72, 0xe5, 0x2d, 0xb8, 0xf0,
	0x1e, 0x48, 0xbc, 0x03, 0xda, 0xb5, 0x1b, 0x29, 0x52, 0xe0, 0x62, 0xcd, 0xec, 0x7c, 0xb3, 0x1e,
	0x7f, 0x63, 0xf2, 0x41, 0x36, 0x8d, 0xbc, 0xad, 0x9c, 0x79, 0xa9, 0x9c, 0x0e, 0x6a, 0x74, 0xef,
	0x3a, 0x55, 0x7d, 0xb7, 0x54, 0xce, 0xcd, 0xa2, 0x2c, 0x17, 0xb5, 0x14, 0xaa, 0x5d, 0xcf, 0xdb,
	0x4a, 0x6e, 0x44, 0xbb, 0x59, 0x55, 0x6c, 0xb5, 0x6e, 0xda, 0xc6, 0x1e, 0x75, 0xf3, 0x6c, 0x5e,
	0x2a, 0xb6, 0xb5, 0xb2, 0x7b, 0x97, 0x19, 0xeb, 0xf9, 0x2f, 0x8b, 0xbc, 0x1d, 0x77, 0xf6, 0xb4,
	0x77, 0x67, 0x9b, 0x55, 0x85, 0xf5, 0xdd, 0xf2, 0xfc, 0x87, 0x45, 0x5e, 0xed, 0xd1, 0xec, 0x17,
	0x64, 0x98, 0xf3, 0x34, 0x46, 0x3f, 0xf8, 0x14, 0xe0, 0x84, 0x3e, 0xb2, 0x87, 0xe4, 0x28, 0xe7,
	0x97, 0x3c, 0xfa, 0xcc, 0xe9, 0xc0, 0xa6, 0xe4, 0x18, 0xf9, 0x05, 0x70, 0x1f, 0x27, 0xc2, 0x8f,
	0x7d, 0x6a, 0xd9, 0x27, 0x84, 0x84, 0xc0, 0x73, 0xb8, 0x32, 0xfc, 0x60, 0x87, 0x87, 0xf4, 0xf1,
	0x0e, 0x2f, 0xe8, 0x73, 0xfb, 0x94, 0xbc, 0x0e, 0xe1, 0x4b, 0x10, 0x06, 0xd7, 0x28, 0xfc, 0x88,
	0x17, 0x98, 0xa4, 0x41, 0xc4, 0x53, 0x4a, 0xec, 0x11, 0x79, 0xb7, 0x47, 0x11, 0x05, 0x5c, 0xe5,
	0x48, 0x87, 0xf6, 0x1b, 0xf2, 0x32, 0x86, 0x29, 0x8a, 0x88, 0xa3, 0x88, 0x93, 0x28, 0x8c, 0x32,
	0x9c, 0xd0, 0x27, 0x3a, 0x6f, 0x8c, 0x89, 0x8f, 0x3c, 0x33, 0x01, 0x8e, 0xf5, 0x0b, 0x33, 0x48,
	0xa6, 0xa8, 0x39, 0xd0, 0xc3, 0x1d, 0x1e, 0xd2, 0x13, 0x1d, 0xa0, 0xe7, 0x51, 0x9e, 0x25, 0xc0,
	0x2f, 0x45, 0x7a, 0x01, 0x09, 0xd2, 0x23, 0x7d, 0x55, 0xaf, 0x24, 0x11, 0xa4, 0xf4, 0xa9, 0xfe,
	0xda, 0xfe, 0x20, 0x8d, 0x91, 0x4f, 0xe8, 0xb3, 0xf1, 0x9f, 0x01, 0x39, 0xfb, 0xda, 0x2c, 0xd9,
	0x7f, 0xf7, 0x3e, 0x3e, 0xdd, 0xb3, 0xd8, 0x58, 0x17, 0x16, 0x0f, 0xae, 0xc7, 0xbd, 0x55, 0x36,
	0xb7, 0xf3, 0x5a, 0xb2, 0x66, 0x2d, 0x1d, 0x59, 0xd5, 0xa6, 0xce, 0x87, 0xf6, 0x57, 0x0b, 0xf5,
	0x8f, 0x9f, 0xe1, 0xa3, 0x79, 0x7e, 0xb3, 0x0e, 0xa6, 0x00, 0xdf, 0xad, 0xd1, 0xb4, 0xbb, 0x0a,
	0x4a, 0xc5, 0x3a, 0xa8, 0x51, 0xe1, 0x32, 0x5d, 0xb0, 0xfa, 0xf9, 0xa0, 0xcf, 0xa0, 0x54, 0xb3,
	0xad, 0x3e, 0x2b, 0xdc, 0x99, 0xd1, 0x7f, 0x5b, 0x67, 0xdd, 0xa1, 0xe7, 0x41, 0xa9, 0x3c, 0x6f,
	0x3b, 0xe1, 0x79, 0x85, 0xeb, 0x79, 0x66, 0xe6, 0xe6, 0xd0, 0x04, 0x7b, 0xff, 0x37, 0x00, 0x00,
	0xff, 0xff, 0x30, 0x1a, 0x16, 0xc8, 0xa4, 0x02, 0x00, 0x00,
}
