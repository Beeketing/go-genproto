// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/change_status.proto

package resources // import "github.com/Beeketing/go-genproto/googleapis/ads/googleads/v0/resources"

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

// Describes the status of returned resource.
type ChangeStatus struct {
	// The resource name of the change status.
	// Change status resource names have the form:
	//
	// `customers/{customer_id}/changeStatus/{change_status_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Time at which the most recent change has occurred on this resource.
	LastChangeDateTime *wrappers.StringValue `protobuf:"bytes,3,opt,name=last_change_date_time,json=lastChangeDateTime,proto3" json:"last_change_date_time,omitempty"`
	// Represents the type of the changed resource. This dictates what fields
	// will be set. For example, for AD_GROUP, campaign and ad_group fields will
	// be set.
	ResourceType enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType `protobuf:"varint,4,opt,name=resource_type,json=resourceType,proto3,enum=google.ads.googleads.v0.enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType" json:"resource_type,omitempty"`
	// The Campaign affected by this change.
	Campaign *wrappers.StringValue `protobuf:"bytes,5,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// The AdGroup affected by this change.
	AdGroup *wrappers.StringValue `protobuf:"bytes,6,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// Represents the status of the changed resource.
	ResourceStatus enums.ChangeStatusOperationEnum_ChangeStatusOperation `protobuf:"varint,8,opt,name=resource_status,json=resourceStatus,proto3,enum=google.ads.googleads.v0.enums.ChangeStatusOperationEnum_ChangeStatusOperation" json:"resource_status,omitempty"`
	// The AdGroupAd affected by this change.
	AdGroupAd *wrappers.StringValue `protobuf:"bytes,9,opt,name=ad_group_ad,json=adGroupAd,proto3" json:"ad_group_ad,omitempty"`
	// The AdGroupCriterion affected by this change.
	AdGroupCriterion *wrappers.StringValue `protobuf:"bytes,10,opt,name=ad_group_criterion,json=adGroupCriterion,proto3" json:"ad_group_criterion,omitempty"`
	// The CampaignCriterion affected by this change.
	CampaignCriterion *wrappers.StringValue `protobuf:"bytes,11,opt,name=campaign_criterion,json=campaignCriterion,proto3" json:"campaign_criterion,omitempty"`
	// The Feed affected by this change.
	Feed *wrappers.StringValue `protobuf:"bytes,12,opt,name=feed,proto3" json:"feed,omitempty"`
	// The FeedItem affected by this change.
	FeedItem             *wrappers.StringValue `protobuf:"bytes,13,opt,name=feed_item,json=feedItem,proto3" json:"feed_item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ChangeStatus) Reset()         { *m = ChangeStatus{} }
func (m *ChangeStatus) String() string { return proto.CompactTextString(m) }
func (*ChangeStatus) ProtoMessage()    {}
func (*ChangeStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_change_status_558457b70433db0d, []int{0}
}
func (m *ChangeStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeStatus.Unmarshal(m, b)
}
func (m *ChangeStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeStatus.Marshal(b, m, deterministic)
}
func (dst *ChangeStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeStatus.Merge(dst, src)
}
func (m *ChangeStatus) XXX_Size() int {
	return xxx_messageInfo_ChangeStatus.Size(m)
}
func (m *ChangeStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeStatus proto.InternalMessageInfo

func (m *ChangeStatus) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ChangeStatus) GetLastChangeDateTime() *wrappers.StringValue {
	if m != nil {
		return m.LastChangeDateTime
	}
	return nil
}

func (m *ChangeStatus) GetResourceType() enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType {
	if m != nil {
		return m.ResourceType
	}
	return enums.ChangeStatusResourceTypeEnum_UNSPECIFIED
}

func (m *ChangeStatus) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *ChangeStatus) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *ChangeStatus) GetResourceStatus() enums.ChangeStatusOperationEnum_ChangeStatusOperation {
	if m != nil {
		return m.ResourceStatus
	}
	return enums.ChangeStatusOperationEnum_UNSPECIFIED
}

func (m *ChangeStatus) GetAdGroupAd() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupAd
	}
	return nil
}

func (m *ChangeStatus) GetAdGroupCriterion() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupCriterion
	}
	return nil
}

func (m *ChangeStatus) GetCampaignCriterion() *wrappers.StringValue {
	if m != nil {
		return m.CampaignCriterion
	}
	return nil
}

func (m *ChangeStatus) GetFeed() *wrappers.StringValue {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *ChangeStatus) GetFeedItem() *wrappers.StringValue {
	if m != nil {
		return m.FeedItem
	}
	return nil
}

func init() {
	proto.RegisterType((*ChangeStatus)(nil), "google.ads.googleads.v0.resources.ChangeStatus")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/change_status.proto", fileDescriptor_change_status_558457b70433db0d)
}

var fileDescriptor_change_status_558457b70433db0d = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xd1, 0x6a, 0xdb, 0x30,
	0x14, 0x86, 0x71, 0xda, 0x65, 0x89, 0x92, 0x76, 0xab, 0x60, 0x60, 0xca, 0x18, 0xe9, 0x46, 0x21,
	0x57, 0xb2, 0xe9, 0x18, 0xdb, 0xdc, 0xc1, 0x70, 0xbb, 0x11, 0xd6, 0x41, 0x5b, 0xd2, 0x90, 0x8b,
	0x11, 0x30, 0x6a, 0x74, 0xea, 0x19, 0x62, 0xcb, 0x48, 0x72, 0x4b, 0x5e, 0x67, 0x97, 0xbb, 0xd9,
	0x7b, 0xec, 0x51, 0xf6, 0x10, 0x63, 0xd8, 0x96, 0xd4, 0xb0, 0x91, 0xd5, 0xbd, 0xca, 0x51, 0xf4,
	0xff, 0xdf, 0x39, 0x47, 0x47, 0x32, 0x7a, 0x15, 0x73, 0x1e, 0x2f, 0xc0, 0xa3, 0x4c, 0x7a, 0x75,
	0x58, 0x46, 0xd7, 0xbe, 0x27, 0x40, 0xf2, 0x42, 0xcc, 0x41, 0x7a, 0xf3, 0xaf, 0x34, 0x8b, 0x21,
	0x92, 0x8a, 0xaa, 0x42, 0x92, 0x5c, 0x70, 0xc5, 0xf1, 0x5e, 0xad, 0x25, 0x94, 0x49, 0x62, 0x6d,
	0xe4, 0xda, 0x27, 0xd6, 0xb6, 0x7b, 0xb8, 0x8e, 0x0c, 0x59, 0x91, 0xfe, 0x45, 0x8d, 0x78, 0x0e,
	0x82, 0xaa, 0x84, 0x67, 0x35, 0x7f, 0xf7, 0xfd, 0x7d, 0xcc, 0x26, 0x67, 0xa4, 0x96, 0x39, 0x68,
	0xc0, 0x33, 0x0d, 0xa8, 0x56, 0x97, 0xc5, 0x95, 0x77, 0x23, 0x68, 0x9e, 0x83, 0xd0, 0x0d, 0x3c,
	0xff, 0xd1, 0x46, 0xfd, 0xe3, 0x8a, 0x72, 0x51, 0x41, 0xf0, 0x0b, 0xb4, 0x65, 0x39, 0x19, 0x4d,
	0xc1, 0x75, 0x06, 0xce, 0xb0, 0x3b, 0xee, 0x9b, 0x3f, 0x4f, 0x69, 0x0a, 0xf8, 0x0c, 0x3d, 0x59,
	0x50, 0xa9, 0x22, 0x9d, 0x9f, 0x51, 0x05, 0x91, 0x4a, 0x52, 0x70, 0x37, 0x06, 0xce, 0xb0, 0x77,
	0xf0, 0x54, 0x9f, 0x05, 0x31, 0x59, 0xc9, 0x85, 0x12, 0x49, 0x16, 0x4f, 0xe9, 0xa2, 0x80, 0x31,
	0x2e, 0xad, 0x75, 0xce, 0x0f, 0x54, 0xc1, 0x24, 0x49, 0x01, 0x2f, 0x57, 0xb2, 0x96, 0xd5, 0xbb,
	0x9b, 0x03, 0x67, 0xb8, 0x7d, 0x30, 0x21, 0xeb, 0xce, 0xb7, 0xea, 0x9f, 0xac, 0x56, 0x3e, 0xd6,
	0xfe, 0xc9, 0x32, 0x87, 0x8f, 0x59, 0x91, 0xae, 0xdd, 0xbc, 0xed, 0xa5, 0x5c, 0xe1, 0x37, 0xa8,
	0x33, 0xa7, 0x69, 0x4e, 0x93, 0x38, 0x73, 0x1f, 0x34, 0x28, 0xdf, 0xaa, 0xf1, 0x6b, 0xd4, 0xa1,
	0x2c, 0x8a, 0x05, 0x2f, 0x72, 0xb7, 0xdd, 0xc0, 0xf9, 0x90, 0xb2, 0x51, 0x29, 0xc6, 0x37, 0xe8,
	0x91, 0xed, 0xb6, 0x9e, 0x9d, 0xdb, 0xa9, 0xfa, 0x3d, 0xbd, 0x47, 0xbf, 0x67, 0xe6, 0xaa, 0xfc,
	0xd3, 0xac, 0xdd, 0x19, 0x6f, 0x9b, 0x34, 0x7a, 0xb8, 0xef, 0x50, 0xcf, 0x54, 0x1c, 0x51, 0xe6,
	0x76, 0x1b, 0x14, 0xdd, 0xd5, 0x45, 0x87, 0x0c, 0x9f, 0x20, 0x6c, 0xdd, 0x73, 0x91, 0x28, 0x10,
	0x09, 0xcf, 0x5c, 0xd4, 0x00, 0xf2, 0x58, 0x43, 0x8e, 0x8d, 0x0b, 0x7f, 0x46, 0xd8, 0x9c, 0xe3,
	0x0a, 0xab, 0xd7, 0x80, 0xb5, 0x63, 0x7c, 0xb7, 0x30, 0x1f, 0x6d, 0x5e, 0x01, 0x30, 0xb7, 0xdf,
	0xc0, 0x5e, 0x29, 0xf1, 0x5b, 0xd4, 0x2d, 0x7f, 0xa3, 0x44, 0x41, 0xea, 0x6e, 0x35, 0x99, 0x7a,
	0x29, 0xff, 0xa4, 0x20, 0x3d, 0xfa, 0xed, 0xa0, 0xfd, 0x39, 0x4f, 0xc9, 0x9d, 0x2f, 0xff, 0x68,
	0x67, 0x75, 0x28, 0xe7, 0x25, 0xf5, 0xdc, 0xf9, 0x72, 0xa2, 0x7d, 0x31, 0x5f, 0xd0, 0x2c, 0x26,
	0x5c, 0xc4, 0x5e, 0x0c, 0x59, 0x95, 0xd3, 0xbc, 0xf0, 0x3c, 0x91, 0xff, 0xf9, 0x0e, 0x1d, 0xda,
	0xe8, 0x5b, 0x6b, 0x63, 0x14, 0x86, 0xdf, 0x5b, 0x7b, 0xa3, 0x1a, 0x19, 0x32, 0x49, 0xea, 0xb0,
	0x8c, 0xa6, 0x3e, 0x31, 0xb7, 0x5e, 0xfe, 0x34, 0x9a, 0x59, 0xc8, 0xe4, 0xcc, 0x6a, 0x66, 0x53,
	0x7f, 0x66, 0x35, 0xbf, 0x5a, 0xfb, 0xf5, 0x46, 0x10, 0x84, 0x4c, 0x06, 0x81, 0x55, 0x05, 0xc1,
	0xd4, 0x0f, 0x02, 0xab, 0xbb, 0x6c, 0x57, 0xc5, 0xbe, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0x21,
	0xc6, 0x81, 0xba, 0x33, 0x05, 0x00, 0x00,
}
