// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/conversion_action.proto

package resources // import "github.com/Beeketing/go-genproto/googleapis/ads/googleads/v0/resources"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/Beeketing/protobuf/ptypes/wrappers"
import common "github.com/Beeketing/go-genproto/googleapis/ads/googleads/v0/common"
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

// A conversion action.
type ConversionAction struct {
	// The resource name of the conversion action.
	// Conversion action resource names have the form:
	//
	// `customers/{customer_id}/conversionActions/{conversion_action_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the conversion action.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the conversion action.
	//
	// This field is required and should not be empty when creating new
	// conversion actions.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The status of this conversion action for conversion event accrual.
	Status enums.ConversionActionStatusEnum_ConversionActionStatus `protobuf:"varint,4,opt,name=status,proto3,enum=google.ads.googleads.v0.enums.ConversionActionStatusEnum_ConversionActionStatus" json:"status,omitempty"`
	// The type of this conversion action.
	Type enums.ConversionActionTypeEnum_ConversionActionType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v0.enums.ConversionActionTypeEnum_ConversionActionType" json:"type,omitempty"`
	// The category of conversions reported for this conversion action.
	Category enums.ConversionActionCategoryEnum_ConversionActionCategory `protobuf:"varint,6,opt,name=category,proto3,enum=google.ads.googleads.v0.enums.ConversionActionCategoryEnum_ConversionActionCategory" json:"category,omitempty"`
	// The resource name of the conversion action owner customer, or null if this
	// is a system-defined conversion action.
	OwnerCustomer *wrappers.StringValue `protobuf:"bytes,7,opt,name=owner_customer,json=ownerCustomer,proto3" json:"owner_customer,omitempty"`
	// Whether this conversion action should be included in the "conversions"
	// metric.
	IncludeInConversionsMetric *wrappers.BoolValue `protobuf:"bytes,8,opt,name=include_in_conversions_metric,json=includeInConversionsMetric,proto3" json:"include_in_conversions_metric,omitempty"`
	// The maximum number of days that may elapse between an interaction
	// (e.g., a click) and a conversion event.
	ClickThroughLookbackWindowDays *wrappers.Int64Value `protobuf:"bytes,9,opt,name=click_through_lookback_window_days,json=clickThroughLookbackWindowDays,proto3" json:"click_through_lookback_window_days,omitempty"`
	// The maximum number of days which may elapse between an impression and a
	// conversion without an interaction.
	ViewThroughLookbackWindowDays *wrappers.Int64Value `protobuf:"bytes,10,opt,name=view_through_lookback_window_days,json=viewThroughLookbackWindowDays,proto3" json:"view_through_lookback_window_days,omitempty"`
	// Settings related to the value for conversion events associated with this
	// conversion action.
	ValueSettings *ConversionAction_ValueSettings `protobuf:"bytes,11,opt,name=value_settings,json=valueSettings,proto3" json:"value_settings,omitempty"`
	// How to count conversion events for the conversion action.
	CountingType enums.ConversionActionCountingTypeEnum_ConversionActionCountingType `protobuf:"varint,12,opt,name=counting_type,json=countingType,proto3,enum=google.ads.googleads.v0.enums.ConversionActionCountingTypeEnum_ConversionActionCountingType" json:"counting_type,omitempty"`
	// Settings related to this conversion action's attribution model.
	AttributionModelSettings *ConversionAction_AttributionModelSettings `protobuf:"bytes,13,opt,name=attribution_model_settings,json=attributionModelSettings,proto3" json:"attribution_model_settings,omitempty"`
	// The snippets used for tracking conversions.
	TagSnippets []*common.TagSnippet `protobuf:"bytes,14,rep,name=tag_snippets,json=tagSnippets,proto3" json:"tag_snippets,omitempty"`
	// The phone call duration in seconds after which a conversion should be
	// reported for this conversion action.
	//
	// The value must be between 0 and 10000, inclusive.
	PhoneCallDurationSeconds *wrappers.Int64Value `protobuf:"bytes,15,opt,name=phone_call_duration_seconds,json=phoneCallDurationSeconds,proto3" json:"phone_call_duration_seconds,omitempty"`
	// App ID for an app conversion action.
	AppId                *wrappers.StringValue `protobuf:"bytes,16,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ConversionAction) Reset()         { *m = ConversionAction{} }
func (m *ConversionAction) String() string { return proto.CompactTextString(m) }
func (*ConversionAction) ProtoMessage()    {}
func (*ConversionAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_action_84c1356f86b03971, []int{0}
}
func (m *ConversionAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionAction.Unmarshal(m, b)
}
func (m *ConversionAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionAction.Marshal(b, m, deterministic)
}
func (dst *ConversionAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionAction.Merge(dst, src)
}
func (m *ConversionAction) XXX_Size() int {
	return xxx_messageInfo_ConversionAction.Size(m)
}
func (m *ConversionAction) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionAction.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionAction proto.InternalMessageInfo

func (m *ConversionAction) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ConversionAction) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConversionAction) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConversionAction) GetStatus() enums.ConversionActionStatusEnum_ConversionActionStatus {
	if m != nil {
		return m.Status
	}
	return enums.ConversionActionStatusEnum_UNSPECIFIED
}

func (m *ConversionAction) GetType() enums.ConversionActionTypeEnum_ConversionActionType {
	if m != nil {
		return m.Type
	}
	return enums.ConversionActionTypeEnum_UNSPECIFIED
}

func (m *ConversionAction) GetCategory() enums.ConversionActionCategoryEnum_ConversionActionCategory {
	if m != nil {
		return m.Category
	}
	return enums.ConversionActionCategoryEnum_UNSPECIFIED
}

func (m *ConversionAction) GetOwnerCustomer() *wrappers.StringValue {
	if m != nil {
		return m.OwnerCustomer
	}
	return nil
}

func (m *ConversionAction) GetIncludeInConversionsMetric() *wrappers.BoolValue {
	if m != nil {
		return m.IncludeInConversionsMetric
	}
	return nil
}

func (m *ConversionAction) GetClickThroughLookbackWindowDays() *wrappers.Int64Value {
	if m != nil {
		return m.ClickThroughLookbackWindowDays
	}
	return nil
}

func (m *ConversionAction) GetViewThroughLookbackWindowDays() *wrappers.Int64Value {
	if m != nil {
		return m.ViewThroughLookbackWindowDays
	}
	return nil
}

func (m *ConversionAction) GetValueSettings() *ConversionAction_ValueSettings {
	if m != nil {
		return m.ValueSettings
	}
	return nil
}

func (m *ConversionAction) GetCountingType() enums.ConversionActionCountingTypeEnum_ConversionActionCountingType {
	if m != nil {
		return m.CountingType
	}
	return enums.ConversionActionCountingTypeEnum_UNSPECIFIED
}

func (m *ConversionAction) GetAttributionModelSettings() *ConversionAction_AttributionModelSettings {
	if m != nil {
		return m.AttributionModelSettings
	}
	return nil
}

func (m *ConversionAction) GetTagSnippets() []*common.TagSnippet {
	if m != nil {
		return m.TagSnippets
	}
	return nil
}

func (m *ConversionAction) GetPhoneCallDurationSeconds() *wrappers.Int64Value {
	if m != nil {
		return m.PhoneCallDurationSeconds
	}
	return nil
}

func (m *ConversionAction) GetAppId() *wrappers.StringValue {
	if m != nil {
		return m.AppId
	}
	return nil
}

// Settings related to this conversion action's attribution model.
type ConversionAction_AttributionModelSettings struct {
	// The attribution model type of this conversion action.
	AttributionModel enums.AttributionModelEnum_AttributionModel `protobuf:"varint,1,opt,name=attribution_model,json=attributionModel,proto3,enum=google.ads.googleads.v0.enums.AttributionModelEnum_AttributionModel" json:"attribution_model,omitempty"`
	// The status of the data-driven attribution model for the conversion
	// action.
	DataDrivenModelStatus enums.DataDrivenModelStatusEnum_DataDrivenModelStatus `protobuf:"varint,2,opt,name=data_driven_model_status,json=dataDrivenModelStatus,proto3,enum=google.ads.googleads.v0.enums.DataDrivenModelStatusEnum_DataDrivenModelStatus" json:"data_driven_model_status,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                                              `json:"-"`
	XXX_unrecognized      []byte                                                `json:"-"`
	XXX_sizecache         int32                                                 `json:"-"`
}

func (m *ConversionAction_AttributionModelSettings) Reset() {
	*m = ConversionAction_AttributionModelSettings{}
}
func (m *ConversionAction_AttributionModelSettings) String() string { return proto.CompactTextString(m) }
func (*ConversionAction_AttributionModelSettings) ProtoMessage()    {}
func (*ConversionAction_AttributionModelSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_action_84c1356f86b03971, []int{0, 0}
}
func (m *ConversionAction_AttributionModelSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionAction_AttributionModelSettings.Unmarshal(m, b)
}
func (m *ConversionAction_AttributionModelSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionAction_AttributionModelSettings.Marshal(b, m, deterministic)
}
func (dst *ConversionAction_AttributionModelSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionAction_AttributionModelSettings.Merge(dst, src)
}
func (m *ConversionAction_AttributionModelSettings) XXX_Size() int {
	return xxx_messageInfo_ConversionAction_AttributionModelSettings.Size(m)
}
func (m *ConversionAction_AttributionModelSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionAction_AttributionModelSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionAction_AttributionModelSettings proto.InternalMessageInfo

func (m *ConversionAction_AttributionModelSettings) GetAttributionModel() enums.AttributionModelEnum_AttributionModel {
	if m != nil {
		return m.AttributionModel
	}
	return enums.AttributionModelEnum_UNSPECIFIED
}

func (m *ConversionAction_AttributionModelSettings) GetDataDrivenModelStatus() enums.DataDrivenModelStatusEnum_DataDrivenModelStatus {
	if m != nil {
		return m.DataDrivenModelStatus
	}
	return enums.DataDrivenModelStatusEnum_UNSPECIFIED
}

// Settings related to the value for conversion events associated with this
// conversion action.
type ConversionAction_ValueSettings struct {
	// The value to use when conversion events for this conversion action are
	// sent with an invalid, disallowed or missing value, or when
	// this conversion action is configured to always use the default value.
	DefaultValue *wrappers.DoubleValue `protobuf:"bytes,1,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	// The currency code to use when conversion events for this conversion
	// action are sent with an invalid or missing currency code, or when this
	// conversion action is configured to always use the default value.
	DefaultCurrencyCode *wrappers.StringValue `protobuf:"bytes,2,opt,name=default_currency_code,json=defaultCurrencyCode,proto3" json:"default_currency_code,omitempty"`
	// Controls whether the default value and default currency code are used in
	// place of the value and currency code specified in conversion events for
	// this conversion action.
	AlwaysUseDefaultValue *wrappers.BoolValue `protobuf:"bytes,3,opt,name=always_use_default_value,json=alwaysUseDefaultValue,proto3" json:"always_use_default_value,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}            `json:"-"`
	XXX_unrecognized      []byte              `json:"-"`
	XXX_sizecache         int32               `json:"-"`
}

func (m *ConversionAction_ValueSettings) Reset()         { *m = ConversionAction_ValueSettings{} }
func (m *ConversionAction_ValueSettings) String() string { return proto.CompactTextString(m) }
func (*ConversionAction_ValueSettings) ProtoMessage()    {}
func (*ConversionAction_ValueSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_action_84c1356f86b03971, []int{0, 1}
}
func (m *ConversionAction_ValueSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionAction_ValueSettings.Unmarshal(m, b)
}
func (m *ConversionAction_ValueSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionAction_ValueSettings.Marshal(b, m, deterministic)
}
func (dst *ConversionAction_ValueSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionAction_ValueSettings.Merge(dst, src)
}
func (m *ConversionAction_ValueSettings) XXX_Size() int {
	return xxx_messageInfo_ConversionAction_ValueSettings.Size(m)
}
func (m *ConversionAction_ValueSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionAction_ValueSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionAction_ValueSettings proto.InternalMessageInfo

func (m *ConversionAction_ValueSettings) GetDefaultValue() *wrappers.DoubleValue {
	if m != nil {
		return m.DefaultValue
	}
	return nil
}

func (m *ConversionAction_ValueSettings) GetDefaultCurrencyCode() *wrappers.StringValue {
	if m != nil {
		return m.DefaultCurrencyCode
	}
	return nil
}

func (m *ConversionAction_ValueSettings) GetAlwaysUseDefaultValue() *wrappers.BoolValue {
	if m != nil {
		return m.AlwaysUseDefaultValue
	}
	return nil
}

func init() {
	proto.RegisterType((*ConversionAction)(nil), "google.ads.googleads.v0.resources.ConversionAction")
	proto.RegisterType((*ConversionAction_AttributionModelSettings)(nil), "google.ads.googleads.v0.resources.ConversionAction.AttributionModelSettings")
	proto.RegisterType((*ConversionAction_ValueSettings)(nil), "google.ads.googleads.v0.resources.ConversionAction.ValueSettings")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/conversion_action.proto", fileDescriptor_conversion_action_84c1356f86b03971)
}

var fileDescriptor_conversion_action_84c1356f86b03971 = []byte{
	// 955 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xcf, 0x6e, 0x23, 0x35,
	0x18, 0x57, 0xd2, 0xdd, 0xb2, 0xeb, 0x26, 0xd9, 0x62, 0x54, 0x69, 0x94, 0x65, 0x57, 0xed, 0xa2,
	0x95, 0x2a, 0x90, 0x26, 0x51, 0x16, 0x90, 0x08, 0x08, 0x29, 0x4d, 0xd0, 0xaa, 0xa8, 0x5d, 0x55,
	0x93, 0x52, 0xa4, 0x55, 0x90, 0x71, 0x6c, 0xef, 0x64, 0xd4, 0x19, 0x7b, 0xb0, 0x3d, 0x89, 0x72,
	0x84, 0x0b, 0x12, 0x2f, 0xc0, 0x9d, 0x23, 0x6f, 0xc0, 0x2b, 0xf0, 0x28, 0xbc, 0x01, 0x37, 0x34,
	0x1e, 0xcf, 0x34, 0x4d, 0x32, 0x4d, 0xc3, 0xcd, 0x1e, 0xff, 0xfe, 0xf8, 0xfb, 0xfc, 0xf9, 0xf3,
	0x80, 0x2f, 0x7c, 0x21, 0xfc, 0x90, 0xb5, 0x30, 0x55, 0xad, 0x6c, 0x98, 0x8e, 0xa6, 0xed, 0x96,
	0x64, 0x4a, 0x24, 0x92, 0x30, 0xd5, 0x22, 0x82, 0x4f, 0x99, 0x54, 0x81, 0xe0, 0x08, 0x13, 0x1d,
	0x08, 0xee, 0xc6, 0x52, 0x68, 0x01, 0x8f, 0x32, 0xbc, 0x8b, 0xa9, 0x72, 0x0b, 0xaa, 0x3b, 0x6d,
	0xbb, 0x05, 0xb5, 0xd9, 0x2e, 0x53, 0x27, 0x22, 0x8a, 0x04, 0x6f, 0x69, 0xec, 0x23, 0xc5, 0x83,
	0x38, 0x66, 0x3a, 0x13, 0x6d, 0x7e, 0x56, 0xc6, 0x60, 0x3c, 0x89, 0x54, 0x0b, 0x6b, 0x2d, 0x83,
	0x71, 0x92, 0xee, 0x02, 0x45, 0x82, 0xb2, 0xd0, 0xd2, 0xbe, 0xbe, 0x9b, 0xb6, 0x12, 0x02, 0x22,
	0x58, 0x33, 0x5f, 0xc8, 0xb9, 0xe5, 0xf7, 0xb7, 0xe6, 0x8b, 0x84, 0xeb, 0x80, 0xfb, 0x48, 0xcf,
	0x63, 0x66, 0x45, 0xbe, 0xda, 0x56, 0x44, 0x69, 0xac, 0x13, 0x65, 0xd9, 0xdd, 0x6d, 0xd9, 0xf7,
	0x77, 0xa6, 0x58, 0x63, 0x44, 0x65, 0x30, 0x65, 0x36, 0x6b, 0xb7, 0x9d, 0x9f, 0x5b, 0xb6, 0x99,
	0x8d, 0x93, 0x77, 0xad, 0x99, 0xc4, 0x71, 0xcc, 0xa4, 0x5d, 0x7f, 0xf1, 0xd7, 0x13, 0xb0, 0xdf,
	0x2f, 0xec, 0x7b, 0xc6, 0x1d, 0x7e, 0x04, 0xea, 0xf9, 0x39, 0x23, 0x8e, 0x23, 0xe6, 0x54, 0x0e,
	0x2b, 0xc7, 0x8f, 0xbd, 0x5a, 0xfe, 0xf1, 0x0d, 0x8e, 0x18, 0xfc, 0x04, 0x54, 0x03, 0xea, 0x54,
	0x0f, 0x2b, 0xc7, 0x7b, 0x9d, 0xa7, 0xb6, 0x48, 0xdc, 0xdc, 0xc6, 0x3d, 0xe5, 0xfa, 0xf3, 0x4f,
	0xaf, 0x70, 0x98, 0x30, 0xaf, 0x1a, 0x50, 0xd8, 0x06, 0x0f, 0x8c, 0xd0, 0x8e, 0x81, 0x7f, 0xb8,
	0x02, 0x1f, 0x6a, 0x19, 0x70, 0x3f, 0xc3, 0x1b, 0x24, 0x9c, 0x80, 0xdd, 0x2c, 0x10, 0xe7, 0xc1,
	0x61, 0xe5, 0xb8, 0xd1, 0xb9, 0x70, 0xcb, 0x4a, 0xd2, 0xe4, 0xc1, 0x5d, 0x0e, 0x62, 0x68, 0xc8,
	0xdf, 0xf0, 0x24, 0x2a, 0x59, 0xf2, 0xac, 0x3e, 0xfc, 0x11, 0x3c, 0x48, 0xd3, 0xed, 0x3c, 0x34,
	0x3e, 0x67, 0x5b, 0xfa, 0x5c, 0xce, 0x63, 0xb6, 0xd6, 0x25, 0x5d, 0xf0, 0x8c, 0x32, 0x8c, 0xc1,
	0xa3, 0xbc, 0x26, 0x9d, 0x5d, 0xe3, 0x72, 0xb9, 0xa5, 0x4b, 0xdf, 0xd2, 0xd7, 0x3a, 0xe5, 0x8b,
	0x5e, 0xe1, 0x02, 0xfb, 0xa0, 0x21, 0x66, 0x9c, 0x49, 0x44, 0x12, 0xa5, 0x45, 0xc4, 0xa4, 0xf3,
	0xde, 0x3d, 0x32, 0x5f, 0x37, 0x9c, 0xbe, 0xa5, 0xc0, 0x1f, 0xc0, 0xb3, 0x80, 0x93, 0x30, 0xa1,
	0x0c, 0x05, 0xe9, 0xad, 0xc8, 0x5d, 0x15, 0x8a, 0x98, 0x96, 0x01, 0x71, 0x1e, 0x19, 0xcd, 0xe6,
	0x8a, 0xe6, 0x89, 0x10, 0x61, 0xa6, 0xd8, 0xb4, 0x02, 0xa7, 0xfc, 0x66, 0xd3, 0xea, 0xdc, 0xb0,
	0xa1, 0x0f, 0x5e, 0x90, 0x30, 0x20, 0xd7, 0x48, 0x4f, 0xa4, 0x48, 0xfc, 0x09, 0x0a, 0x85, 0xb8,
	0x1e, 0x63, 0x72, 0x8d, 0x66, 0x01, 0xa7, 0x62, 0x86, 0x28, 0x9e, 0x2b, 0xe7, 0xf1, 0xe6, 0x02,
	0x7b, 0x6e, 0x64, 0x2e, 0x33, 0x95, 0x33, 0x2b, 0xf2, 0xbd, 0xd1, 0x18, 0xe0, 0xb9, 0x82, 0x0c,
	0x1c, 0x4d, 0x03, 0x36, 0xbb, 0xdb, 0x07, 0x6c, 0xf6, 0x79, 0x96, 0xaa, 0x94, 0xdb, 0x4c, 0x40,
	0x63, 0x9a, 0xe2, 0x90, 0x62, 0x3a, 0x6d, 0x1f, 0xca, 0xd9, 0x33, 0x9a, 0x3d, 0x77, 0x63, 0x33,
	0x5d, 0x39, 0x52, 0xd7, 0x38, 0x0e, 0xad, 0x90, 0x57, 0x9f, 0x2e, 0x4e, 0xe1, 0xcf, 0x15, 0x50,
	0xbf, 0xd5, 0xa4, 0x9c, 0x9a, 0xa9, 0xaa, 0xd1, 0xb6, 0x55, 0x65, 0x35, 0x4a, 0x6b, 0x78, 0x11,
	0xe0, 0xd5, 0xc8, 0xc2, 0x0c, 0xfe, 0x56, 0x01, 0xcd, 0x95, 0x8e, 0x7d, 0x13, 0x7a, 0xdd, 0x84,
	0x7e, 0xf6, 0x7f, 0x42, 0xef, 0xdd, 0xa8, 0x9e, 0xa7, 0xa2, 0x45, 0x16, 0x1c, 0x5c, 0xb2, 0x02,
	0xcf, 0x41, 0x6d, 0xe1, 0xb9, 0x51, 0x4e, 0xe3, 0x70, 0xe7, 0x78, 0xaf, 0xf3, 0x71, 0xa9, 0x7b,
	0xf6, 0x44, 0xb9, 0x97, 0xd8, 0x1f, 0x66, 0x14, 0x6f, 0x4f, 0x17, 0x63, 0x05, 0xdf, 0x82, 0xa7,
	0xf1, 0x44, 0x70, 0x86, 0x08, 0x0e, 0x43, 0x44, 0x13, 0x89, 0xb3, 0x96, 0xce, 0x88, 0xe0, 0x54,
	0x39, 0x4f, 0x36, 0x97, 0x8a, 0x63, 0xf8, 0x7d, 0x1c, 0x86, 0x03, 0xcb, 0x1e, 0x66, 0x64, 0xf8,
	0x0a, 0xec, 0xe2, 0x38, 0x46, 0x01, 0x75, 0xf6, 0xef, 0x71, 0x23, 0x1f, 0xe2, 0x38, 0x3e, 0xa5,
	0xcd, 0xdf, 0xab, 0xc0, 0x29, 0x4b, 0x0b, 0xfc, 0x09, 0xbc, 0xbf, 0x72, 0x10, 0xa6, 0x63, 0x37,
	0x3a, 0x83, 0x0d, 0x05, 0xb1, 0xac, 0x69, 0x8a, 0x60, 0xf9, 0xa3, 0xb7, 0xbf, 0x9c, 0x77, 0xf8,
	0x6b, 0x05, 0x38, 0x65, 0x0f, 0x8f, 0x79, 0x12, 0x1a, 0x9d, 0x37, 0x1b, 0xac, 0x07, 0x58, 0xe3,
	0x81, 0x61, 0x67, 0xd1, 0xdc, 0xb4, 0xeb, 0xb5, 0x2b, 0xde, 0x01, 0x5d, 0xf7, 0xb9, 0xf9, 0x6f,
	0x05, 0xd4, 0x6f, 0xdd, 0x15, 0xd8, 0x03, 0x75, 0xca, 0xde, 0xe1, 0x24, 0xd4, 0xc8, 0xdc, 0x1a,
	0x93, 0x8a, 0x75, 0x79, 0x1e, 0x88, 0x64, 0x1c, 0xb2, 0x2c, 0xcf, 0x35, 0x4b, 0x31, 0x33, 0x78,
	0x01, 0x0e, 0x72, 0x09, 0x92, 0x48, 0xc9, 0x38, 0x99, 0x23, 0x22, 0x28, 0xb3, 0xaf, 0xdd, 0xdd,
	0x47, 0xf6, 0x81, 0xa5, 0xf6, 0x2d, 0xb3, 0x2f, 0x28, 0x83, 0x43, 0xe0, 0xe0, 0x70, 0x86, 0xe7,
	0x0a, 0x25, 0x8a, 0xa1, 0xdb, 0xfb, 0xdb, 0xd9, 0xd8, 0x45, 0x0f, 0x32, 0xee, 0x77, 0x8a, 0x0d,
	0x16, 0xb6, 0x79, 0xf2, 0x4b, 0x15, 0xbc, 0x24, 0x22, 0xda, 0x7c, 0xc7, 0x4e, 0x0e, 0x96, 0x2f,
	0xd9, 0x45, 0x6a, 0x72, 0x51, 0x79, 0xfb, 0xad, 0xe5, 0xfa, 0x22, 0xc4, 0xdc, 0x77, 0x85, 0xf4,
	0x5b, 0x3e, 0xe3, 0x66, 0x0b, 0xf9, 0xcf, 0x46, 0x1c, 0xa8, 0x3b, 0xfe, 0x20, 0xbf, 0x2c, 0x46,
	0x7f, 0x54, 0x77, 0x5e, 0xf7, 0x7a, 0x7f, 0x56, 0x8f, 0x5e, 0x67, 0x92, 0x3d, 0xaa, 0xdc, 0x6c,
	0x98, 0x8e, 0xae, 0xda, 0xae, 0x97, 0x23, 0xff, 0xce, 0x31, 0xa3, 0x1e, 0x55, 0xa3, 0x02, 0x33,
	0xba, 0x6a, 0x8f, 0x0a, 0xcc, 0x3f, 0xd5, 0x97, 0xd9, 0x42, 0xb7, 0xdb, 0xa3, 0xaa, 0xdb, 0x2d,
	0x50, 0xdd, 0xee, 0x55, 0xbb, 0xdb, 0x2d, 0x70, 0xe3, 0x5d, 0xb3, 0xd9, 0x57, 0xff, 0x05, 0x00,
	0x00, 0xff, 0xff, 0xca, 0x40, 0xd9, 0xc1, 0xed, 0x0a, 0x00, 0x00,
}
