// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1/organization_settings.proto

package securitycenter // import "github.com/Beeketing/go-genproto/googleapis/cloud/securitycenter/v1"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/Beeketing/go-genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The mode of inclusion when running Asset Discovery.
// Asset discovery can be limited by explicitly identifying projects to be
// included or excluded. If INCLUDE_ONLY is set, then only those projects
// within the organization and their children are discovered during asset
// discovery. If EXCLUDE is set, then projects that don't match those
// projects are discovered during asset discovery. If neither are set, then
// all projects within the organization are discovered during asset
// discovery.
type OrganizationSettings_AssetDiscoveryConfig_InclusionMode int32

const (
	// Unspecified. Setting the mode with this value will disable
	// inclusion/exclusion filtering for Asset Discovery.
	OrganizationSettings_AssetDiscoveryConfig_INCLUSION_MODE_UNSPECIFIED OrganizationSettings_AssetDiscoveryConfig_InclusionMode = 0
	// Asset Discovery will capture only the resources within the projects
	// specified. All other resources will be ignored.
	OrganizationSettings_AssetDiscoveryConfig_INCLUDE_ONLY OrganizationSettings_AssetDiscoveryConfig_InclusionMode = 1
	// Asset Discovery will ignore all resources under the projects specified.
	// All other resources will be retrieved.
	OrganizationSettings_AssetDiscoveryConfig_EXCLUDE OrganizationSettings_AssetDiscoveryConfig_InclusionMode = 2
)

var OrganizationSettings_AssetDiscoveryConfig_InclusionMode_name = map[int32]string{
	0: "INCLUSION_MODE_UNSPECIFIED",
	1: "INCLUDE_ONLY",
	2: "EXCLUDE",
}
var OrganizationSettings_AssetDiscoveryConfig_InclusionMode_value = map[string]int32{
	"INCLUSION_MODE_UNSPECIFIED": 0,
	"INCLUDE_ONLY":               1,
	"EXCLUDE":                    2,
}

func (x OrganizationSettings_AssetDiscoveryConfig_InclusionMode) String() string {
	return proto.EnumName(OrganizationSettings_AssetDiscoveryConfig_InclusionMode_name, int32(x))
}
func (OrganizationSettings_AssetDiscoveryConfig_InclusionMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_organization_settings_7bd2c80d69d259b3, []int{0, 0, 0}
}

// User specified settings that are attached to the Cloud Security Command
// Center (Cloud SCC) organization.
type OrganizationSettings struct {
	// The relative resource name of the settings. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Example:
	// "organizations/123/organizationSettings".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A flag that indicates if Asset Discovery should be enabled. If the flag is
	// set to `true`, then discovery of assets will occur. If it is set to `false,
	// all historical assets will remain, but discovery of future assets will not
	// occur.
	EnableAssetDiscovery bool `protobuf:"varint,2,opt,name=enable_asset_discovery,json=enableAssetDiscovery,proto3" json:"enable_asset_discovery,omitempty"`
	// The configuration used for Asset Discovery runs.
	AssetDiscoveryConfig *OrganizationSettings_AssetDiscoveryConfig `protobuf:"bytes,3,opt,name=asset_discovery_config,json=assetDiscoveryConfig,proto3" json:"asset_discovery_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *OrganizationSettings) Reset()         { *m = OrganizationSettings{} }
func (m *OrganizationSettings) String() string { return proto.CompactTextString(m) }
func (*OrganizationSettings) ProtoMessage()    {}
func (*OrganizationSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_organization_settings_7bd2c80d69d259b3, []int{0}
}
func (m *OrganizationSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganizationSettings.Unmarshal(m, b)
}
func (m *OrganizationSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganizationSettings.Marshal(b, m, deterministic)
}
func (dst *OrganizationSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizationSettings.Merge(dst, src)
}
func (m *OrganizationSettings) XXX_Size() int {
	return xxx_messageInfo_OrganizationSettings.Size(m)
}
func (m *OrganizationSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizationSettings.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizationSettings proto.InternalMessageInfo

func (m *OrganizationSettings) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OrganizationSettings) GetEnableAssetDiscovery() bool {
	if m != nil {
		return m.EnableAssetDiscovery
	}
	return false
}

func (m *OrganizationSettings) GetAssetDiscoveryConfig() *OrganizationSettings_AssetDiscoveryConfig {
	if m != nil {
		return m.AssetDiscoveryConfig
	}
	return nil
}

// The configuration used for Asset Discovery runs.
type OrganizationSettings_AssetDiscoveryConfig struct {
	// The project ids to use for filtering asset discovery.
	ProjectIds []string `protobuf:"bytes,1,rep,name=project_ids,json=projectIds,proto3" json:"project_ids,omitempty"`
	// The mode to use for filtering asset discovery.
	InclusionMode        OrganizationSettings_AssetDiscoveryConfig_InclusionMode `protobuf:"varint,2,opt,name=inclusion_mode,json=inclusionMode,proto3,enum=google.cloud.securitycenter.v1.OrganizationSettings_AssetDiscoveryConfig_InclusionMode" json:"inclusion_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                `json:"-"`
	XXX_unrecognized     []byte                                                  `json:"-"`
	XXX_sizecache        int32                                                   `json:"-"`
}

func (m *OrganizationSettings_AssetDiscoveryConfig) Reset() {
	*m = OrganizationSettings_AssetDiscoveryConfig{}
}
func (m *OrganizationSettings_AssetDiscoveryConfig) String() string { return proto.CompactTextString(m) }
func (*OrganizationSettings_AssetDiscoveryConfig) ProtoMessage()    {}
func (*OrganizationSettings_AssetDiscoveryConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_organization_settings_7bd2c80d69d259b3, []int{0, 0}
}
func (m *OrganizationSettings_AssetDiscoveryConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig.Unmarshal(m, b)
}
func (m *OrganizationSettings_AssetDiscoveryConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig.Marshal(b, m, deterministic)
}
func (dst *OrganizationSettings_AssetDiscoveryConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig.Merge(dst, src)
}
func (m *OrganizationSettings_AssetDiscoveryConfig) XXX_Size() int {
	return xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig.Size(m)
}
func (m *OrganizationSettings_AssetDiscoveryConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig proto.InternalMessageInfo

func (m *OrganizationSettings_AssetDiscoveryConfig) GetProjectIds() []string {
	if m != nil {
		return m.ProjectIds
	}
	return nil
}

func (m *OrganizationSettings_AssetDiscoveryConfig) GetInclusionMode() OrganizationSettings_AssetDiscoveryConfig_InclusionMode {
	if m != nil {
		return m.InclusionMode
	}
	return OrganizationSettings_AssetDiscoveryConfig_INCLUSION_MODE_UNSPECIFIED
}

func init() {
	proto.RegisterType((*OrganizationSettings)(nil), "google.cloud.securitycenter.v1.OrganizationSettings")
	proto.RegisterType((*OrganizationSettings_AssetDiscoveryConfig)(nil), "google.cloud.securitycenter.v1.OrganizationSettings.AssetDiscoveryConfig")
	proto.RegisterEnum("google.cloud.securitycenter.v1.OrganizationSettings_AssetDiscoveryConfig_InclusionMode", OrganizationSettings_AssetDiscoveryConfig_InclusionMode_name, OrganizationSettings_AssetDiscoveryConfig_InclusionMode_value)
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1/organization_settings.proto", fileDescriptor_organization_settings_7bd2c80d69d259b3)
}

var fileDescriptor_organization_settings_7bd2c80d69d259b3 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x6b, 0xdb, 0x30,
	0x14, 0xc7, 0xe7, 0x24, 0xed, 0x52, 0x65, 0x2d, 0x46, 0x84, 0x12, 0xc2, 0xe8, 0x4c, 0x4e, 0x3e,
	0xc9, 0xb4, 0xdb, 0x69, 0x3b, 0x6d, 0xb1, 0x07, 0x2e, 0xa9, 0x5d, 0x1c, 0xca, 0x7e, 0x5c, 0x84,
	0x2a, 0x6b, 0x42, 0xc3, 0xd1, 0x33, 0x96, 0x12, 0xe8, 0x0e, 0xdb, 0x71, 0x7f, 0xf2, 0xae, 0xa3,
	0x8a, 0x0b, 0xf1, 0xc8, 0xb6, 0x4b, 0x6f, 0xd2, 0xf7, 0xfb, 0xbe, 0x1f, 0xbd, 0x27, 0x1e, 0x7a,
	0x2d, 0x01, 0x64, 0x25, 0x22, 0x5e, 0xc1, 0xba, 0x8c, 0x8c, 0xe0, 0xeb, 0x46, 0xd9, 0x3b, 0x2e,
	0xb4, 0x15, 0x4d, 0xb4, 0x39, 0x8f, 0xa0, 0x91, 0x4c, 0xab, 0x6f, 0xcc, 0x2a, 0xd0, 0xd4, 0x08,
	0x6b, 0x95, 0x96, 0x86, 0xd4, 0x0d, 0x58, 0xc0, 0x67, 0xdb, 0x2c, 0x71, 0x59, 0xd2, 0xcd, 0x92,
	0xcd, 0xf9, 0xf4, 0x79, 0xcb, 0x66, 0xb5, 0x8a, 0x98, 0xd6, 0x60, 0x1d, 0xa5, 0x4d, 0xcf, 0x7e,
	0xf5, 0xd1, 0x38, 0xdf, 0xa1, 0x2f, 0x5b, 0x38, 0xc6, 0x68, 0xa0, 0xd9, 0x4a, 0x4c, 0xbc, 0xc0,
	0x0b, 0x8f, 0x0a, 0x77, 0xc6, 0xaf, 0xd0, 0xa9, 0xd0, 0xec, 0xb6, 0x12, 0x94, 0x19, 0x23, 0x2c,
	0x2d, 0x95, 0xe1, 0xb0, 0x11, 0xcd, 0xdd, 0xa4, 0x17, 0x78, 0xe1, 0xb0, 0x18, 0x6f, 0xdd, 0xb7,
	0xf7, 0x66, 0xfc, 0xe0, 0xe1, 0x1f, 0xe8, 0xf4, 0x8f, 0x72, 0xca, 0x41, 0x7f, 0x51, 0x72, 0xd2,
	0x0f, 0xbc, 0x70, 0x74, 0x91, 0x92, 0x7f, 0x4f, 0x40, 0xf6, 0xf5, 0x47, 0xba, 0x8f, 0xcc, 0x1d,
	0xb0, 0x18, 0xb3, 0x3d, 0xea, 0xf4, 0x67, 0x0f, 0x8d, 0xf7, 0x95, 0xe3, 0x17, 0x68, 0x54, 0x37,
	0xf0, 0x55, 0x70, 0x4b, 0x55, 0x69, 0x26, 0x5e, 0xd0, 0x0f, 0x8f, 0x0a, 0xd4, 0x4a, 0x69, 0x69,
	0xf0, 0x77, 0x74, 0xa2, 0x34, 0xaf, 0xd6, 0xe6, 0xfe, 0xdf, 0x57, 0x50, 0x0a, 0x37, 0xe8, 0xc9,
	0xc5, 0x87, 0x47, 0x6b, 0x99, 0xa4, 0x0f, 0xfc, 0x2b, 0x28, 0x45, 0x71, 0xac, 0x76, 0xaf, 0xb3,
	0x0c, 0x1d, 0x77, 0x7c, 0x7c, 0x86, 0xa6, 0x69, 0x36, 0x5f, 0xdc, 0x2c, 0xd3, 0x3c, 0xa3, 0x57,
	0x79, 0x9c, 0xd0, 0x9b, 0x6c, 0x79, 0x9d, 0xcc, 0xd3, 0xf7, 0x69, 0x12, 0xfb, 0x4f, 0xb0, 0x8f,
	0x9e, 0x39, 0x3f, 0x4e, 0x68, 0x9e, 0x2d, 0x3e, 0xf9, 0x1e, 0x1e, 0xa1, 0xa7, 0xc9, 0x47, 0xa7,
	0xf8, 0xbd, 0xcb, 0xc1, 0x70, 0xe0, 0x1f, 0x5c, 0x0e, 0x86, 0x07, 0xfe, 0xe1, 0x3b, 0x8b, 0x66,
	0x1c, 0x56, 0xff, 0x19, 0xe4, 0xda, 0xfb, 0xbc, 0x68, 0x2b, 0x24, 0x54, 0x4c, 0x4b, 0x02, 0x8d,
	0x8c, 0xa4, 0xd0, 0x6e, 0x7b, 0xa2, 0xad, 0xc5, 0x6a, 0x65, 0xfe, 0xb6, 0xba, 0x6f, 0xba, 0xca,
	0xed, 0xa1, 0x0b, 0xbe, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x98, 0xf4, 0x7f, 0x14, 0xf2, 0x02,
	0x00, 0x00,
}
