// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/user_interest.proto

package resources // import "github.com/Beeketing/genproto/googleapis/ads/googleads/v0/resources"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/Beeketing/protobuf/ptypes/wrappers"
import common "github.com/Beeketing/genproto/googleapis/ads/googleads/v0/common"
import enums "github.com/Beeketing/genproto/googleapis/ads/googleads/v0/enums"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A user interest: a particular interest-based vertical to be targeted.
type UserInterest struct {
	// The resource name of the user interest.
	// User interest resource names have the form:
	//
	// `customers/{customer_id}/userInterests/{user_interest_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Taxonomy type of the user interest.
	TaxonomyType enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType `protobuf:"varint,2,opt,name=taxonomy_type,json=taxonomyType,proto3,enum=google.ads.googleads.v0.enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType" json:"taxonomy_type,omitempty"`
	// The ID of the user interest.
	UserInterestId *wrappers.Int64Value `protobuf:"bytes,3,opt,name=user_interest_id,json=userInterestId,proto3" json:"user_interest_id,omitempty"`
	// The name of the user interest.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The parent of the user interest.
	UserInterestParent *wrappers.StringValue `protobuf:"bytes,5,opt,name=user_interest_parent,json=userInterestParent,proto3" json:"user_interest_parent,omitempty"`
	// True if the user interest is launched to all channels and locales.
	LaunchedToAll *wrappers.BoolValue `protobuf:"bytes,6,opt,name=launched_to_all,json=launchedToAll,proto3" json:"launched_to_all,omitempty"`
	// Availability information of the user interest.
	Availabilities       []*common.CriterionCategoryAvailability `protobuf:"bytes,7,rep,name=availabilities,proto3" json:"availabilities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *UserInterest) Reset()         { *m = UserInterest{} }
func (m *UserInterest) String() string { return proto.CompactTextString(m) }
func (*UserInterest) ProtoMessage()    {}
func (*UserInterest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_interest_3d8339d911d0d13a, []int{0}
}
func (m *UserInterest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInterest.Unmarshal(m, b)
}
func (m *UserInterest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInterest.Marshal(b, m, deterministic)
}
func (dst *UserInterest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInterest.Merge(dst, src)
}
func (m *UserInterest) XXX_Size() int {
	return xxx_messageInfo_UserInterest.Size(m)
}
func (m *UserInterest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInterest.DiscardUnknown(m)
}

var xxx_messageInfo_UserInterest proto.InternalMessageInfo

func (m *UserInterest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *UserInterest) GetTaxonomyType() enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType {
	if m != nil {
		return m.TaxonomyType
	}
	return enums.UserInterestTaxonomyTypeEnum_UNSPECIFIED
}

func (m *UserInterest) GetUserInterestId() *wrappers.Int64Value {
	if m != nil {
		return m.UserInterestId
	}
	return nil
}

func (m *UserInterest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *UserInterest) GetUserInterestParent() *wrappers.StringValue {
	if m != nil {
		return m.UserInterestParent
	}
	return nil
}

func (m *UserInterest) GetLaunchedToAll() *wrappers.BoolValue {
	if m != nil {
		return m.LaunchedToAll
	}
	return nil
}

func (m *UserInterest) GetAvailabilities() []*common.CriterionCategoryAvailability {
	if m != nil {
		return m.Availabilities
	}
	return nil
}

func init() {
	proto.RegisterType((*UserInterest)(nil), "google.ads.googleads.v0.resources.UserInterest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/user_interest.proto", fileDescriptor_user_interest_3d8339d911d0d13a)
}

var fileDescriptor_user_interest_3d8339d911d0d13a = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x49, 0xb7, 0xae, 0x38, 0xdb, 0x56, 0x0d, 0x5e, 0x84, 0x2a, 0xd2, 0x55, 0x16, 0x7a,
	0x35, 0x09, 0xf5, 0xcf, 0x45, 0x44, 0x24, 0x5d, 0x97, 0xa5, 0x5e, 0x2c, 0xa5, 0xd6, 0x5e, 0x48,
	0x21, 0x4c, 0x93, 0x63, 0x1c, 0x98, 0xcc, 0x84, 0x99, 0x49, 0x35, 0xaf, 0xb3, 0x97, 0x3e, 0x8a,
	0x8f, 0xe2, 0x43, 0x88, 0x34, 0x93, 0x84, 0x74, 0x97, 0xaa, 0x77, 0x67, 0x3a, 0xdf, 0xf7, 0x3b,
	0xe7, 0x74, 0xbe, 0xa0, 0x57, 0x89, 0x10, 0x09, 0x03, 0x97, 0xc4, 0xca, 0x35, 0xe5, 0xae, 0xda,
	0x7a, 0xae, 0x04, 0x25, 0x72, 0x19, 0x81, 0x72, 0x73, 0x05, 0x32, 0xa4, 0x5c, 0x83, 0x04, 0xa5,
	0x71, 0x26, 0x85, 0x16, 0xf6, 0xa9, 0xd1, 0x62, 0x12, 0x2b, 0xdc, 0xd8, 0xf0, 0xd6, 0xc3, 0x8d,
	0x6d, 0xf8, 0xfe, 0x10, 0x39, 0x12, 0x69, 0x2a, 0xb8, 0x1b, 0x49, 0xaa, 0x41, 0x52, 0xc1, 0xc3,
	0x88, 0x68, 0x48, 0x84, 0x2c, 0x42, 0xb2, 0x25, 0x94, 0x91, 0x0d, 0x65, 0x54, 0x17, 0xa6, 0xd1,
	0xf0, 0xdd, 0x21, 0x0a, 0xf0, 0x3c, 0xbd, 0x31, 0x5b, 0xa8, 0xc9, 0x77, 0xc1, 0x45, 0x5a, 0x84,
	0xba, 0xc8, 0xa0, 0x02, 0x3c, 0xad, 0x00, 0xe5, 0x69, 0x93, 0x7f, 0x71, 0xbf, 0x49, 0x92, 0x65,
	0x20, 0x95, 0xb9, 0x7f, 0x76, 0xdd, 0x45, 0xbd, 0x4f, 0x0a, 0xe4, 0xac, 0x82, 0xd8, 0xcf, 0x51,
	0xbf, 0x5e, 0x22, 0xe4, 0x24, 0x05, 0xc7, 0x1a, 0x59, 0xe3, 0x7b, 0x8b, 0x5e, 0xfd, 0xe3, 0x15,
	0x49, 0xc1, 0x2e, 0x50, 0x7f, 0xaf, 0x99, 0xd3, 0x19, 0x59, 0xe3, 0xc1, 0x64, 0x89, 0x0f, 0xfd,
	0x2f, 0xe5, 0xb8, 0xb8, 0xdd, 0x68, 0x59, 0xf9, 0x97, 0x45, 0x06, 0x17, 0x3c, 0x4f, 0x0f, 0x5e,
	0x2e, 0x7a, 0xba, 0x75, 0xb2, 0x2f, 0xd0, 0x83, 0xfd, 0xad, 0x69, 0xec, 0x1c, 0x8d, 0xac, 0xf1,
	0xc9, 0xe4, 0x71, 0xdd, 0xbd, 0xde, 0x15, 0xcf, 0xb8, 0x7e, 0xfd, 0x72, 0x45, 0x58, 0x0e, 0x8b,
	0x41, 0xde, 0xc2, 0xcf, 0x62, 0xdb, 0x43, 0xdd, 0x72, 0xbb, 0x6e, 0x69, 0x7d, 0x72, 0xcb, 0xfa,
	0x51, 0x4b, 0xca, 0x13, 0xe3, 0x2d, 0x95, 0xf6, 0x15, 0x7a, 0xb4, 0xdf, 0x38, 0x23, 0x12, 0xb8,
	0x76, 0xee, 0xfc, 0x07, 0xc1, 0x6e, 0x77, 0x9f, 0x97, 0x3e, 0x7b, 0x8a, 0xee, 0x33, 0x92, 0xf3,
	0xe8, 0x2b, 0xc4, 0xa1, 0x16, 0x21, 0x61, 0xcc, 0x39, 0x2e, 0x51, 0xc3, 0x5b, 0xa8, 0xa9, 0x10,
	0xcc, 0x80, 0xfa, 0xb5, 0x65, 0x29, 0x02, 0xc6, 0x6c, 0x40, 0x83, 0x56, 0x68, 0x28, 0x28, 0xe7,
	0xee, 0xe8, 0x68, 0x7c, 0x32, 0x79, 0x7b, 0xf0, 0x21, 0x4c, 0xfa, 0xf0, 0x79, 0x9d, 0xbe, 0xf3,
	0x2a, 0x7c, 0x41, 0x2b, 0x7b, 0x8b, 0x1b, 0xd0, 0xe9, 0x6f, 0x0b, 0x9d, 0x45, 0x22, 0xc5, 0xff,
	0x4c, 0xfd, 0xf4, 0x61, 0xfb, 0x15, 0xe7, 0xbb, 0xf9, 0xe7, 0xd6, 0xe7, 0x0f, 0x95, 0x2f, 0x11,
	0x8c, 0xf0, 0x04, 0x0b, 0x99, 0xb8, 0x09, 0xf0, 0x72, 0xbb, 0x3a, 0xd4, 0x19, 0x55, 0x7f, 0xf9,
	0x06, 0xdf, 0x34, 0xd5, 0x75, 0xe7, 0xe8, 0x32, 0x08, 0x7e, 0x74, 0x4e, 0x2f, 0x0d, 0x32, 0x88,
	0x15, 0x36, 0xe5, 0xae, 0x5a, 0x79, 0x78, 0x51, 0x2b, 0x7f, 0xd6, 0x9a, 0x75, 0x10, 0xab, 0x75,
	0xa3, 0x59, 0xaf, 0xbc, 0x75, 0xa3, 0xf9, 0xd5, 0x39, 0x33, 0x17, 0xbe, 0x1f, 0xc4, 0xca, 0xf7,
	0x1b, 0x95, 0xef, 0xaf, 0x3c, 0xdf, 0x6f, 0x74, 0x9b, 0xe3, 0x72, 0xd8, 0x17, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xa2, 0xa1, 0xae, 0x6c, 0x2f, 0x04, 0x00, 0x00,
}
