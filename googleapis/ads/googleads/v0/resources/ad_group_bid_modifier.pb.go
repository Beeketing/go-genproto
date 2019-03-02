// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/ad_group_bid_modifier.proto

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

// Represents an ad group bid modifier.
type AdGroupBidModifier struct {
	// The resource name of the ad group bid modifier.
	// Ad group bid modifier resource names have the form:
	//
	// `customers/{customer_id}/adGroupBidModifiers/{ad_group_id}_{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ad group to which this criterion belongs.
	AdGroup *wrappers.StringValue `protobuf:"bytes,2,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// The ID of the criterion to bid modify.
	//
	// This field is ignored for mutates.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,3,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// The modifier for the bid when the criterion matches. The modifier must be
	// in the range: 0.1 - 10.0. The range is 1.0 - 6.0 for PreferredContent.
	// Use 0 to opt out of a Device type.
	BidModifier *wrappers.DoubleValue `protobuf:"bytes,4,opt,name=bid_modifier,json=bidModifier,proto3" json:"bid_modifier,omitempty"`
	// The base ad group from which this draft/trial adgroup bid modifier was
	// created. If ad_group is a base ad group then this field will be equal to
	// ad_group. If the ad group was created in the draft or trial and has no
	// corresponding base ad group, then this field will be null.
	// This field is readonly.
	BaseAdGroup *wrappers.StringValue `protobuf:"bytes,9,opt,name=base_ad_group,json=baseAdGroup,proto3" json:"base_ad_group,omitempty"`
	// Bid modifier source.
	BidModifierSource enums.BidModifierSourceEnum_BidModifierSource `protobuf:"varint,10,opt,name=bid_modifier_source,json=bidModifierSource,proto3,enum=google.ads.googleads.v0.enums.BidModifierSourceEnum_BidModifierSource" json:"bid_modifier_source,omitempty"`
	// The criterion of this ad group bid modifier.
	//
	// Types that are valid to be assigned to Criterion:
	//	*AdGroupBidModifier_HotelDateSelectionType
	//	*AdGroupBidModifier_HotelAdvanceBookingWindow
	//	*AdGroupBidModifier_HotelLengthOfStay
	//	*AdGroupBidModifier_HotelCheckInDay
	//	*AdGroupBidModifier_Device
	//	*AdGroupBidModifier_PreferredContent
	Criterion            isAdGroupBidModifier_Criterion `protobuf_oneof:"criterion"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *AdGroupBidModifier) Reset()         { *m = AdGroupBidModifier{} }
func (m *AdGroupBidModifier) String() string { return proto.CompactTextString(m) }
func (*AdGroupBidModifier) ProtoMessage()    {}
func (*AdGroupBidModifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_bid_modifier_019268ab03739a54, []int{0}
}
func (m *AdGroupBidModifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupBidModifier.Unmarshal(m, b)
}
func (m *AdGroupBidModifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupBidModifier.Marshal(b, m, deterministic)
}
func (dst *AdGroupBidModifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupBidModifier.Merge(dst, src)
}
func (m *AdGroupBidModifier) XXX_Size() int {
	return xxx_messageInfo_AdGroupBidModifier.Size(m)
}
func (m *AdGroupBidModifier) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupBidModifier.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupBidModifier proto.InternalMessageInfo

func (m *AdGroupBidModifier) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupBidModifier) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *AdGroupBidModifier) GetCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CriterionId
	}
	return nil
}

func (m *AdGroupBidModifier) GetBidModifier() *wrappers.DoubleValue {
	if m != nil {
		return m.BidModifier
	}
	return nil
}

func (m *AdGroupBidModifier) GetBaseAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.BaseAdGroup
	}
	return nil
}

func (m *AdGroupBidModifier) GetBidModifierSource() enums.BidModifierSourceEnum_BidModifierSource {
	if m != nil {
		return m.BidModifierSource
	}
	return enums.BidModifierSourceEnum_UNSPECIFIED
}

type isAdGroupBidModifier_Criterion interface {
	isAdGroupBidModifier_Criterion()
}

type AdGroupBidModifier_HotelDateSelectionType struct {
	HotelDateSelectionType *common.HotelDateSelectionTypeInfo `protobuf:"bytes,5,opt,name=hotel_date_selection_type,json=hotelDateSelectionType,proto3,oneof"`
}

type AdGroupBidModifier_HotelAdvanceBookingWindow struct {
	HotelAdvanceBookingWindow *common.HotelAdvanceBookingWindowInfo `protobuf:"bytes,6,opt,name=hotel_advance_booking_window,json=hotelAdvanceBookingWindow,proto3,oneof"`
}

type AdGroupBidModifier_HotelLengthOfStay struct {
	HotelLengthOfStay *common.HotelLengthOfStayInfo `protobuf:"bytes,7,opt,name=hotel_length_of_stay,json=hotelLengthOfStay,proto3,oneof"`
}

type AdGroupBidModifier_HotelCheckInDay struct {
	HotelCheckInDay *common.HotelCheckInDayInfo `protobuf:"bytes,8,opt,name=hotel_check_in_day,json=hotelCheckInDay,proto3,oneof"`
}

type AdGroupBidModifier_Device struct {
	Device *common.DeviceInfo `protobuf:"bytes,11,opt,name=device,proto3,oneof"`
}

type AdGroupBidModifier_PreferredContent struct {
	PreferredContent *common.PreferredContentInfo `protobuf:"bytes,12,opt,name=preferred_content,json=preferredContent,proto3,oneof"`
}

func (*AdGroupBidModifier_HotelDateSelectionType) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_HotelAdvanceBookingWindow) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_HotelLengthOfStay) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_HotelCheckInDay) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_Device) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_PreferredContent) isAdGroupBidModifier_Criterion() {}

func (m *AdGroupBidModifier) GetCriterion() isAdGroupBidModifier_Criterion {
	if m != nil {
		return m.Criterion
	}
	return nil
}

func (m *AdGroupBidModifier) GetHotelDateSelectionType() *common.HotelDateSelectionTypeInfo {
	if x, ok := m.GetCriterion().(*AdGroupBidModifier_HotelDateSelectionType); ok {
		return x.HotelDateSelectionType
	}
	return nil
}

func (m *AdGroupBidModifier) GetHotelAdvanceBookingWindow() *common.HotelAdvanceBookingWindowInfo {
	if x, ok := m.GetCriterion().(*AdGroupBidModifier_HotelAdvanceBookingWindow); ok {
		return x.HotelAdvanceBookingWindow
	}
	return nil
}

func (m *AdGroupBidModifier) GetHotelLengthOfStay() *common.HotelLengthOfStayInfo {
	if x, ok := m.GetCriterion().(*AdGroupBidModifier_HotelLengthOfStay); ok {
		return x.HotelLengthOfStay
	}
	return nil
}

func (m *AdGroupBidModifier) GetHotelCheckInDay() *common.HotelCheckInDayInfo {
	if x, ok := m.GetCriterion().(*AdGroupBidModifier_HotelCheckInDay); ok {
		return x.HotelCheckInDay
	}
	return nil
}

func (m *AdGroupBidModifier) GetDevice() *common.DeviceInfo {
	if x, ok := m.GetCriterion().(*AdGroupBidModifier_Device); ok {
		return x.Device
	}
	return nil
}

func (m *AdGroupBidModifier) GetPreferredContent() *common.PreferredContentInfo {
	if x, ok := m.GetCriterion().(*AdGroupBidModifier_PreferredContent); ok {
		return x.PreferredContent
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AdGroupBidModifier) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AdGroupBidModifier_OneofMarshaler, _AdGroupBidModifier_OneofUnmarshaler, _AdGroupBidModifier_OneofSizer, []interface{}{
		(*AdGroupBidModifier_HotelDateSelectionType)(nil),
		(*AdGroupBidModifier_HotelAdvanceBookingWindow)(nil),
		(*AdGroupBidModifier_HotelLengthOfStay)(nil),
		(*AdGroupBidModifier_HotelCheckInDay)(nil),
		(*AdGroupBidModifier_Device)(nil),
		(*AdGroupBidModifier_PreferredContent)(nil),
	}
}

func _AdGroupBidModifier_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AdGroupBidModifier)
	// criterion
	switch x := m.Criterion.(type) {
	case *AdGroupBidModifier_HotelDateSelectionType:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HotelDateSelectionType); err != nil {
			return err
		}
	case *AdGroupBidModifier_HotelAdvanceBookingWindow:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HotelAdvanceBookingWindow); err != nil {
			return err
		}
	case *AdGroupBidModifier_HotelLengthOfStay:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HotelLengthOfStay); err != nil {
			return err
		}
	case *AdGroupBidModifier_HotelCheckInDay:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HotelCheckInDay); err != nil {
			return err
		}
	case *AdGroupBidModifier_Device:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Device); err != nil {
			return err
		}
	case *AdGroupBidModifier_PreferredContent:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PreferredContent); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AdGroupBidModifier.Criterion has unexpected type %T", x)
	}
	return nil
}

func _AdGroupBidModifier_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AdGroupBidModifier)
	switch tag {
	case 5: // criterion.hotel_date_selection_type
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.HotelDateSelectionTypeInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &AdGroupBidModifier_HotelDateSelectionType{msg}
		return true, err
	case 6: // criterion.hotel_advance_booking_window
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.HotelAdvanceBookingWindowInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &AdGroupBidModifier_HotelAdvanceBookingWindow{msg}
		return true, err
	case 7: // criterion.hotel_length_of_stay
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.HotelLengthOfStayInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &AdGroupBidModifier_HotelLengthOfStay{msg}
		return true, err
	case 8: // criterion.hotel_check_in_day
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.HotelCheckInDayInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &AdGroupBidModifier_HotelCheckInDay{msg}
		return true, err
	case 11: // criterion.device
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.DeviceInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &AdGroupBidModifier_Device{msg}
		return true, err
	case 12: // criterion.preferred_content
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.PreferredContentInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &AdGroupBidModifier_PreferredContent{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AdGroupBidModifier_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AdGroupBidModifier)
	// criterion
	switch x := m.Criterion.(type) {
	case *AdGroupBidModifier_HotelDateSelectionType:
		s := proto.Size(x.HotelDateSelectionType)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AdGroupBidModifier_HotelAdvanceBookingWindow:
		s := proto.Size(x.HotelAdvanceBookingWindow)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AdGroupBidModifier_HotelLengthOfStay:
		s := proto.Size(x.HotelLengthOfStay)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AdGroupBidModifier_HotelCheckInDay:
		s := proto.Size(x.HotelCheckInDay)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AdGroupBidModifier_Device:
		s := proto.Size(x.Device)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AdGroupBidModifier_PreferredContent:
		s := proto.Size(x.PreferredContent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*AdGroupBidModifier)(nil), "google.ads.googleads.v0.resources.AdGroupBidModifier")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/ad_group_bid_modifier.proto", fileDescriptor_ad_group_bid_modifier_019268ab03739a54)
}

var fileDescriptor_ad_group_bid_modifier_019268ab03739a54 = []byte{
	// 679 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x86, 0xbf, 0xa4, 0x1f, 0xfd, 0xd9, 0xa4, 0x40, 0x0d, 0x02, 0x53, 0x2a, 0xd4, 0x82, 0x2a,
	0x55, 0x48, 0xd8, 0x51, 0x5b, 0xa8, 0x64, 0x54, 0x20, 0x69, 0xa0, 0x0d, 0xe2, 0xa7, 0x4a, 0x50,
	0x90, 0x50, 0xa4, 0xd5, 0xda, 0x3b, 0x71, 0xac, 0xc6, 0xbb, 0xd6, 0x7a, 0x9d, 0x28, 0x67, 0x1c,
	0x70, 0x25, 0x1c, 0x72, 0xc2, 0x7d, 0x70, 0x29, 0x5c, 0x05, 0xf2, 0xae, 0x6d, 0x85, 0xb6, 0xa1,
	0x39, 0x9b, 0xcc, 0xce, 0xf3, 0xbe, 0x33, 0x13, 0xef, 0xa2, 0x43, 0x9f, 0x73, 0x7f, 0x08, 0x36,
	0xa1, 0xb1, 0xad, 0xc3, 0x34, 0x1a, 0xd5, 0x6c, 0x01, 0x31, 0x4f, 0x84, 0x07, 0xb1, 0x4d, 0x28,
	0xf6, 0x05, 0x4f, 0x22, 0xec, 0x06, 0x14, 0x87, 0x9c, 0x06, 0xfd, 0x00, 0x84, 0x15, 0x09, 0x2e,
	0xb9, 0xb1, 0xa5, 0x19, 0x8b, 0xd0, 0xd8, 0x2a, 0x70, 0x6b, 0x54, 0xb3, 0x0a, 0x7c, 0xfd, 0xc9,
	0x2c, 0x07, 0x8f, 0x87, 0x21, 0x67, 0xb6, 0x27, 0x02, 0x09, 0x22, 0x20, 0x5a, 0x71, 0xfd, 0x60,
	0x56, 0x39, 0xb0, 0x24, 0x8c, 0xed, 0xe9, 0x1e, 0xb0, 0xb6, 0xc8, 0xc0, 0x07, 0x19, 0xa8, 0x7e,
	0xb9, 0x49, 0xdf, 0x1e, 0x0b, 0x12, 0x45, 0x20, 0x62, 0x7d, 0xfe, 0xf0, 0xe7, 0x32, 0x32, 0xea,
	0xf4, 0x38, 0x9d, 0xa4, 0x11, 0xd0, 0xf7, 0x99, 0x86, 0xf1, 0x08, 0xad, 0xe6, 0xbd, 0x62, 0x46,
	0x42, 0x30, 0x4b, 0x9b, 0xa5, 0x9d, 0x95, 0x76, 0x35, 0x4f, 0x7e, 0x20, 0x21, 0x18, 0x07, 0x68,
	0x39, 0xdf, 0x82, 0x59, 0xde, 0x2c, 0xed, 0x54, 0x76, 0x37, 0xb2, 0x71, 0xad, 0xdc, 0xce, 0xea,
	0x48, 0x11, 0x30, 0xbf, 0x4b, 0x86, 0x09, 0xb4, 0x97, 0x88, 0x36, 0x32, 0x5e, 0xa0, 0x6a, 0x36,
	0x1f, 0x67, 0x38, 0xa0, 0xe6, 0x82, 0x82, 0xef, 0x5f, 0x80, 0x5b, 0x4c, 0x3e, 0xdb, 0xd7, 0x6c,
	0xa5, 0x00, 0x5a, 0xd4, 0x78, 0x89, 0xaa, 0xd3, 0x13, 0x9b, 0xff, 0xcf, 0x30, 0x6f, 0xf2, 0xc4,
	0x1d, 0x42, 0x26, 0xe0, 0x4e, 0x8d, 0xf7, 0x0a, 0xad, 0xba, 0x24, 0x06, 0x5c, 0xb4, 0xbf, 0x32,
	0x47, 0xfb, 0x95, 0x14, 0xc9, 0x76, 0x65, 0x8c, 0xd0, 0xad, 0x4b, 0x96, 0x6e, 0xa2, 0xcd, 0xd2,
	0xce, 0xf5, 0xdd, 0x37, 0xd6, 0xac, 0x0f, 0x40, 0xfd, 0x5d, 0xd6, 0xd4, 0xa6, 0x3b, 0x8a, 0x7b,
	0xcd, 0x92, 0xf0, 0x62, 0xb6, 0xbd, 0xe6, 0x9e, 0x4f, 0x19, 0x63, 0x74, 0x6f, 0xc0, 0x25, 0x0c,
	0x31, 0x25, 0x12, 0x70, 0x0c, 0x43, 0xf0, 0x64, 0xba, 0x45, 0x39, 0x89, 0xc0, 0xbc, 0xa6, 0xa6,
	0x70, 0x66, 0xba, 0xeb, 0x6f, 0xcb, 0x3a, 0x49, 0x05, 0x9a, 0x44, 0x42, 0x27, 0xc7, 0x3f, 0x4d,
	0x22, 0x68, 0xb1, 0x3e, 0x3f, 0xf9, 0xaf, 0x7d, 0x67, 0x70, 0xe9, 0xa9, 0xf1, 0xb5, 0x84, 0x36,
	0xb4, 0x33, 0xa1, 0x23, 0xc2, 0x3c, 0xc0, 0x2e, 0xe7, 0x67, 0x01, 0xf3, 0xf1, 0x38, 0x60, 0x94,
	0x8f, 0xcd, 0x45, 0x65, 0x7e, 0x38, 0x97, 0x79, 0x5d, 0x4b, 0x34, 0xb4, 0xc2, 0x67, 0x25, 0x90,
	0xf9, 0xeb, 0xf1, 0x2e, 0x2b, 0x30, 0x06, 0xe8, 0xb6, 0xee, 0x60, 0x08, 0xcc, 0x97, 0x03, 0xcc,
	0xfb, 0x38, 0x96, 0x64, 0x62, 0x2e, 0x29, 0xe7, 0xa7, 0x73, 0x39, 0xbf, 0x53, 0xe8, 0xc7, 0x7e,
	0x47, 0x92, 0x49, 0xe6, 0xb8, 0x36, 0x38, 0x7f, 0x60, 0xb8, 0xc8, 0xd0, 0x4e, 0xde, 0x00, 0xbc,
	0x33, 0x1c, 0x30, 0x4c, 0xc9, 0xc4, 0x5c, 0x56, 0x3e, 0x7b, 0x73, 0xf9, 0x1c, 0xa5, 0x60, 0x8b,
	0x35, 0x0b, 0x97, 0x1b, 0x83, 0xbf, 0xd3, 0x46, 0x13, 0x2d, 0x52, 0x18, 0x05, 0x1e, 0x98, 0x15,
	0xa5, 0xfb, 0xf8, 0x2a, 0xdd, 0xa6, 0xaa, 0xce, 0xe4, 0x32, 0xd6, 0xf0, 0xd0, 0x5a, 0x24, 0xa0,
	0x0f, 0x42, 0x00, 0xc5, 0x1e, 0x67, 0x12, 0x98, 0x34, 0xab, 0x4a, 0x70, 0xff, 0x2a, 0xc1, 0xd3,
	0x1c, 0x3c, 0xd2, 0x5c, 0x26, 0x7d, 0x33, 0x3a, 0x97, 0x6f, 0x54, 0xd0, 0x4a, 0x71, 0xfd, 0x1a,
	0xdf, 0xca, 0x68, 0xdb, 0xe3, 0xa1, 0x75, 0xe5, 0x1b, 0xd7, 0xb8, 0x7b, 0xf1, 0x61, 0x39, 0x4d,
	0x2f, 0xd6, 0x69, 0xe9, 0xcb, 0xdb, 0x8c, 0xf6, 0xf9, 0x90, 0x30, 0xdf, 0xe2, 0xc2, 0xb7, 0x7d,
	0x60, 0xea, 0xda, 0xe5, 0xef, 0x5b, 0x14, 0xc4, 0xff, 0x78, 0x7f, 0x9f, 0x17, 0xd1, 0xf7, 0xf2,
	0xc2, 0x71, 0xbd, 0xfe, 0xa3, 0xbc, 0x75, 0xac, 0x25, 0xeb, 0x34, 0xb6, 0x74, 0x98, 0x46, 0xdd,
	0x9a, 0xd5, 0xce, 0x2b, 0x7f, 0xe5, 0x35, 0xbd, 0x3a, 0x8d, 0x7b, 0x45, 0x4d, 0xaf, 0x5b, 0xeb,
	0x15, 0x35, 0xbf, 0xcb, 0xdb, 0xfa, 0xc0, 0x71, 0xea, 0x34, 0x76, 0x9c, 0xa2, 0xca, 0x71, 0xba,
	0x35, 0xc7, 0x29, 0xea, 0xdc, 0x45, 0xd5, 0xec, 0xde, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61,
	0x68, 0xa5, 0x12, 0x2b, 0x06, 0x00, 0x00,
}
