// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/customer_client_link.proto

package resources // import "github.com/Beeketing/genproto/googleapis/ads/googleads/v0/resources"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/Beeketing/protobuf/ptypes/wrappers"
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

// Represents customer client link relationship.
type CustomerClientLink struct {
	// Name of the resource.
	// CustomerClientLink resource names have the form:
	//
	// `customers/{customer_id}/customerClientLinks/{client_customer_id}_{manager_link_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The client customer linked to this customer.
	ClientCustomer *wrappers.StringValue `protobuf:"bytes,3,opt,name=client_customer,json=clientCustomer,proto3" json:"client_customer,omitempty"`
	// This is uniquely identifies a customer client link. Read only.
	ManagerLinkId *wrappers.Int64Value `protobuf:"bytes,4,opt,name=manager_link_id,json=managerLinkId,proto3" json:"manager_link_id,omitempty"`
	// This is the status of the link between client and manager.
	Status enums.ManagerLinkStatusEnum_ManagerLinkStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v0.enums.ManagerLinkStatusEnum_ManagerLinkStatus" json:"status,omitempty"`
	// The visibility of the link. Users can choose whether or not to see hidden
	// links in the AdWords UI.
	// Default value is false
	Hidden               *wrappers.BoolValue `protobuf:"bytes,6,opt,name=hidden,proto3" json:"hidden,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CustomerClientLink) Reset()         { *m = CustomerClientLink{} }
func (m *CustomerClientLink) String() string { return proto.CompactTextString(m) }
func (*CustomerClientLink) ProtoMessage()    {}
func (*CustomerClientLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_client_link_269244d846b23fdf, []int{0}
}
func (m *CustomerClientLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerClientLink.Unmarshal(m, b)
}
func (m *CustomerClientLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerClientLink.Marshal(b, m, deterministic)
}
func (dst *CustomerClientLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerClientLink.Merge(dst, src)
}
func (m *CustomerClientLink) XXX_Size() int {
	return xxx_messageInfo_CustomerClientLink.Size(m)
}
func (m *CustomerClientLink) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerClientLink.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerClientLink proto.InternalMessageInfo

func (m *CustomerClientLink) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CustomerClientLink) GetClientCustomer() *wrappers.StringValue {
	if m != nil {
		return m.ClientCustomer
	}
	return nil
}

func (m *CustomerClientLink) GetManagerLinkId() *wrappers.Int64Value {
	if m != nil {
		return m.ManagerLinkId
	}
	return nil
}

func (m *CustomerClientLink) GetStatus() enums.ManagerLinkStatusEnum_ManagerLinkStatus {
	if m != nil {
		return m.Status
	}
	return enums.ManagerLinkStatusEnum_UNSPECIFIED
}

func (m *CustomerClientLink) GetHidden() *wrappers.BoolValue {
	if m != nil {
		return m.Hidden
	}
	return nil
}

func init() {
	proto.RegisterType((*CustomerClientLink)(nil), "google.ads.googleads.v0.resources.CustomerClientLink")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/customer_client_link.proto", fileDescriptor_customer_client_link_269244d846b23fdf)
}

var fileDescriptor_customer_client_link_269244d846b23fdf = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x5d, 0x8a, 0xd4, 0x30,
	0x1c, 0xa7, 0x5d, 0x2d, 0x18, 0xdd, 0x5d, 0xe8, 0x8b, 0x65, 0x14, 0x99, 0x55, 0x16, 0xe6, 0x29,
	0x2d, 0xa3, 0x28, 0x44, 0x5f, 0x3a, 0xc3, 0xba, 0xac, 0xa8, 0x2c, 0x5d, 0xe8, 0x83, 0x14, 0x4b,
	0xb6, 0x89, 0xb1, 0x6c, 0x9b, 0x94, 0xa4, 0x1d, 0x2f, 0xe0, 0x49, 0x7c, 0xf4, 0x28, 0xde, 0xc0,
	0x2b, 0x78, 0x0a, 0x69, 0x3e, 0x2a, 0x52, 0x46, 0xdf, 0x7e, 0x49, 0x7e, 0x5f, 0xff, 0x24, 0xe0,
	0x15, 0x13, 0x82, 0x35, 0x34, 0xc6, 0x44, 0xc5, 0x06, 0x8e, 0x68, 0x97, 0xc4, 0x92, 0x2a, 0x31,
	0xc8, 0x8a, 0xaa, 0xb8, 0x1a, 0x54, 0x2f, 0x5a, 0x2a, 0xcb, 0xaa, 0xa9, 0x29, 0xef, 0xcb, 0xa6,
	0xe6, 0x37, 0xb0, 0x93, 0xa2, 0x17, 0xe1, 0x89, 0x91, 0x40, 0x4c, 0x14, 0x9c, 0xd4, 0x70, 0x97,
	0xc0, 0x49, 0xbd, 0x78, 0xb1, 0x2f, 0x80, 0xf2, 0xa1, 0x55, 0x71, 0x8b, 0x39, 0x66, 0x54, 0x6a,
	0xd3, 0x52, 0xf5, 0xb8, 0x1f, 0x94, 0xf1, 0x5e, 0x3c, 0xb2, 0x42, 0xbd, 0xba, 0x1e, 0x3e, 0xc5,
	0x5f, 0x24, 0xee, 0x3a, 0x2a, 0xed, 0xf9, 0xe3, 0x9f, 0x3e, 0x08, 0xb7, 0xb6, 0xda, 0x56, 0x37,
	0x7b, 0x5b, 0xf3, 0x9b, 0xf0, 0x09, 0x38, 0x74, 0xe1, 0x25, 0xc7, 0x2d, 0x8d, 0xbc, 0xa5, 0xb7,
	0xba, 0x93, 0xdd, 0x73, 0x9b, 0xef, 0x71, 0x4b, 0xc3, 0x33, 0x70, 0x6c, 0x87, 0x71, 0xc3, 0x45,
	0x07, 0x4b, 0x6f, 0x75, 0x77, 0xfd, 0xd0, 0x8e, 0x01, 0x5d, 0x2a, 0xbc, 0xea, 0x65, 0xcd, 0x59,
	0x8e, 0x9b, 0x81, 0x66, 0x47, 0x46, 0xe4, 0x52, 0xc3, 0x2d, 0x38, 0xfe, 0xab, 0x7f, 0x4d, 0xa2,
	0x5b, 0xda, 0xe6, 0xc1, 0xcc, 0xe6, 0x82, 0xf7, 0xcf, 0x9f, 0x19, 0x97, 0x43, 0xab, 0x19, 0xeb,
	0x5e, 0x90, 0xf0, 0x23, 0x08, 0xcc, 0xdc, 0xd1, 0xed, 0xa5, 0xb7, 0x3a, 0x5a, 0xbf, 0x86, 0xfb,
	0x2e, 0x55, 0xdf, 0x18, 0x7c, 0xf7, 0x47, 0x7d, 0xa5, 0x75, 0x67, 0x7c, 0x68, 0xe7, 0xbb, 0x99,
	0x75, 0x0d, 0xd7, 0x20, 0xf8, 0x5c, 0x13, 0x42, 0x79, 0x14, 0xe8, 0x6e, 0x8b, 0x59, 0xb7, 0x8d,
	0x10, 0x8d, 0xa9, 0x66, 0x99, 0x9b, 0xaf, 0x3e, 0x38, 0xad, 0x44, 0x0b, 0xff, 0xfb, 0xbc, 0x9b,
	0xfb, 0xf3, 0x27, 0xb8, 0x1c, 0x7d, 0x2f, 0xbd, 0x0f, 0x6f, 0xac, 0x9a, 0x89, 0x06, 0x73, 0x06,
	0x85, 0x64, 0x31, 0xa3, 0x5c, 0xa7, 0xba, 0x9f, 0xd0, 0xd5, 0xea, 0x1f, 0x3f, 0xef, 0xe5, 0x84,
	0xbe, 0xf9, 0x07, 0xe7, 0x69, 0xfa, 0xdd, 0x3f, 0x39, 0x37, 0x96, 0x29, 0x51, 0xd0, 0xc0, 0x11,
	0xe5, 0x09, 0xcc, 0x1c, 0xf3, 0x87, 0xe3, 0x14, 0x29, 0x51, 0xc5, 0xc4, 0x29, 0xf2, 0xa4, 0x98,
	0x38, 0xbf, 0xfc, 0x53, 0x73, 0x80, 0x50, 0x4a, 0x14, 0x42, 0x13, 0x0b, 0xa1, 0x3c, 0x41, 0x68,
	0xe2, 0x5d, 0x07, 0xba, 0xec, 0xd3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x0c, 0xa0, 0xa7,
	0x25, 0x03, 0x00, 0x00,
}
