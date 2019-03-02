// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/monitoring/v3/uptime_service.proto

package monitoring // import "github.com/Beeketing/genproto/googleapis/monitoring/v3"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/Beeketing/protobuf/ptypes/empty"
import _ "github.com/Beeketing/genproto/googleapis/api/annotations"
import field_mask "github.com/Beeketing/genproto/protobuf/field_mask"

import (
	context "golang.org/x/net/context"
	grpc "github.com/Beeketing/grpc-go"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The protocol for the `ListUptimeCheckConfigs` request.
type ListUptimeCheckConfigsRequest struct {
	// The project whose uptime check configurations are listed. The format
	//   is `projects/[PROJECT_ID]`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of results to return in a single response. The server
	// may further constrain the maximum number of results returned in a single
	// page. If the page_size is <=0, the server will decide the number of results
	// to be returned.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// If this field is not empty then it must contain the `nextPageToken` value
	// returned by a previous call to this method.  Using this field causes the
	// method to return more results from the previous method call.
	PageToken            string   `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUptimeCheckConfigsRequest) Reset()         { *m = ListUptimeCheckConfigsRequest{} }
func (m *ListUptimeCheckConfigsRequest) String() string { return proto.CompactTextString(m) }
func (*ListUptimeCheckConfigsRequest) ProtoMessage()    {}
func (*ListUptimeCheckConfigsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_uptime_service_c74f83143a8cb5a4, []int{0}
}
func (m *ListUptimeCheckConfigsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUptimeCheckConfigsRequest.Unmarshal(m, b)
}
func (m *ListUptimeCheckConfigsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUptimeCheckConfigsRequest.Marshal(b, m, deterministic)
}
func (dst *ListUptimeCheckConfigsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUptimeCheckConfigsRequest.Merge(dst, src)
}
func (m *ListUptimeCheckConfigsRequest) XXX_Size() int {
	return xxx_messageInfo_ListUptimeCheckConfigsRequest.Size(m)
}
func (m *ListUptimeCheckConfigsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUptimeCheckConfigsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUptimeCheckConfigsRequest proto.InternalMessageInfo

func (m *ListUptimeCheckConfigsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListUptimeCheckConfigsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListUptimeCheckConfigsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The protocol for the `ListUptimeCheckConfigs` response.
type ListUptimeCheckConfigsResponse struct {
	// The returned uptime check configurations.
	UptimeCheckConfigs []*UptimeCheckConfig `protobuf:"bytes,1,rep,name=uptime_check_configs,json=uptimeCheckConfigs,proto3" json:"uptime_check_configs,omitempty"`
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is empty, it means no further results for the
	// request. To retrieve the next page of results, the value of the
	// next_page_token is passed to the subsequent List method call (in the
	// request message's page_token field).
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// The total number of uptime check configurations for the project,
	// irrespective of any pagination.
	TotalSize            int32    `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUptimeCheckConfigsResponse) Reset()         { *m = ListUptimeCheckConfigsResponse{} }
func (m *ListUptimeCheckConfigsResponse) String() string { return proto.CompactTextString(m) }
func (*ListUptimeCheckConfigsResponse) ProtoMessage()    {}
func (*ListUptimeCheckConfigsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_uptime_service_c74f83143a8cb5a4, []int{1}
}
func (m *ListUptimeCheckConfigsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUptimeCheckConfigsResponse.Unmarshal(m, b)
}
func (m *ListUptimeCheckConfigsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUptimeCheckConfigsResponse.Marshal(b, m, deterministic)
}
func (dst *ListUptimeCheckConfigsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUptimeCheckConfigsResponse.Merge(dst, src)
}
func (m *ListUptimeCheckConfigsResponse) XXX_Size() int {
	return xxx_messageInfo_ListUptimeCheckConfigsResponse.Size(m)
}
func (m *ListUptimeCheckConfigsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUptimeCheckConfigsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUptimeCheckConfigsResponse proto.InternalMessageInfo

func (m *ListUptimeCheckConfigsResponse) GetUptimeCheckConfigs() []*UptimeCheckConfig {
	if m != nil {
		return m.UptimeCheckConfigs
	}
	return nil
}

func (m *ListUptimeCheckConfigsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func (m *ListUptimeCheckConfigsResponse) GetTotalSize() int32 {
	if m != nil {
		return m.TotalSize
	}
	return 0
}

// The protocol for the `GetUptimeCheckConfig` request.
type GetUptimeCheckConfigRequest struct {
	// The uptime check configuration to retrieve. The format
	//   is `projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID]`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUptimeCheckConfigRequest) Reset()         { *m = GetUptimeCheckConfigRequest{} }
func (m *GetUptimeCheckConfigRequest) String() string { return proto.CompactTextString(m) }
func (*GetUptimeCheckConfigRequest) ProtoMessage()    {}
func (*GetUptimeCheckConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_uptime_service_c74f83143a8cb5a4, []int{2}
}
func (m *GetUptimeCheckConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUptimeCheckConfigRequest.Unmarshal(m, b)
}
func (m *GetUptimeCheckConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUptimeCheckConfigRequest.Marshal(b, m, deterministic)
}
func (dst *GetUptimeCheckConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUptimeCheckConfigRequest.Merge(dst, src)
}
func (m *GetUptimeCheckConfigRequest) XXX_Size() int {
	return xxx_messageInfo_GetUptimeCheckConfigRequest.Size(m)
}
func (m *GetUptimeCheckConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUptimeCheckConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUptimeCheckConfigRequest proto.InternalMessageInfo

func (m *GetUptimeCheckConfigRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The protocol for the `CreateUptimeCheckConfig` request.
type CreateUptimeCheckConfigRequest struct {
	// The project in which to create the uptime check. The format
	//   is `projects/[PROJECT_ID]`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The new uptime check configuration.
	UptimeCheckConfig    *UptimeCheckConfig `protobuf:"bytes,2,opt,name=uptime_check_config,json=uptimeCheckConfig,proto3" json:"uptime_check_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateUptimeCheckConfigRequest) Reset()         { *m = CreateUptimeCheckConfigRequest{} }
func (m *CreateUptimeCheckConfigRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUptimeCheckConfigRequest) ProtoMessage()    {}
func (*CreateUptimeCheckConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_uptime_service_c74f83143a8cb5a4, []int{3}
}
func (m *CreateUptimeCheckConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUptimeCheckConfigRequest.Unmarshal(m, b)
}
func (m *CreateUptimeCheckConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUptimeCheckConfigRequest.Marshal(b, m, deterministic)
}
func (dst *CreateUptimeCheckConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUptimeCheckConfigRequest.Merge(dst, src)
}
func (m *CreateUptimeCheckConfigRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUptimeCheckConfigRequest.Size(m)
}
func (m *CreateUptimeCheckConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUptimeCheckConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUptimeCheckConfigRequest proto.InternalMessageInfo

func (m *CreateUptimeCheckConfigRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateUptimeCheckConfigRequest) GetUptimeCheckConfig() *UptimeCheckConfig {
	if m != nil {
		return m.UptimeCheckConfig
	}
	return nil
}

// The protocol for the `UpdateUptimeCheckConfig` request.
type UpdateUptimeCheckConfigRequest struct {
	// Optional. If present, only the listed fields in the current uptime check
	// configuration are updated with values from the new configuration. If this
	// field is empty, then the current configuration is completely replaced with
	// the new configuration.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// Required. If an `"updateMask"` has been specified, this field gives
	// the values for the set of fields mentioned in the `"updateMask"`. If an
	// `"updateMask"` has not been given, this uptime check configuration replaces
	// the current configuration. If a field is mentioned in `"updateMask"` but
	// the corresonding field is omitted in this partial uptime check
	// configuration, it has the effect of deleting/clearing the field from the
	// configuration on the server.
	//
	// The following fields can be updated: `display_name`,
	// `http_check`, `tcp_check`, `timeout`, `content_matchers`, and
	// `selected_regions`.
	UptimeCheckConfig    *UptimeCheckConfig `protobuf:"bytes,3,opt,name=uptime_check_config,json=uptimeCheckConfig,proto3" json:"uptime_check_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UpdateUptimeCheckConfigRequest) Reset()         { *m = UpdateUptimeCheckConfigRequest{} }
func (m *UpdateUptimeCheckConfigRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUptimeCheckConfigRequest) ProtoMessage()    {}
func (*UpdateUptimeCheckConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_uptime_service_c74f83143a8cb5a4, []int{4}
}
func (m *UpdateUptimeCheckConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUptimeCheckConfigRequest.Unmarshal(m, b)
}
func (m *UpdateUptimeCheckConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUptimeCheckConfigRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateUptimeCheckConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUptimeCheckConfigRequest.Merge(dst, src)
}
func (m *UpdateUptimeCheckConfigRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUptimeCheckConfigRequest.Size(m)
}
func (m *UpdateUptimeCheckConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUptimeCheckConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUptimeCheckConfigRequest proto.InternalMessageInfo

func (m *UpdateUptimeCheckConfigRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func (m *UpdateUptimeCheckConfigRequest) GetUptimeCheckConfig() *UptimeCheckConfig {
	if m != nil {
		return m.UptimeCheckConfig
	}
	return nil
}

// The protocol for the `DeleteUptimeCheckConfig` request.
type DeleteUptimeCheckConfigRequest struct {
	// The uptime check configuration to delete. The format
	//   is `projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID]`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUptimeCheckConfigRequest) Reset()         { *m = DeleteUptimeCheckConfigRequest{} }
func (m *DeleteUptimeCheckConfigRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUptimeCheckConfigRequest) ProtoMessage()    {}
func (*DeleteUptimeCheckConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_uptime_service_c74f83143a8cb5a4, []int{5}
}
func (m *DeleteUptimeCheckConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUptimeCheckConfigRequest.Unmarshal(m, b)
}
func (m *DeleteUptimeCheckConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUptimeCheckConfigRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteUptimeCheckConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUptimeCheckConfigRequest.Merge(dst, src)
}
func (m *DeleteUptimeCheckConfigRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUptimeCheckConfigRequest.Size(m)
}
func (m *DeleteUptimeCheckConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUptimeCheckConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUptimeCheckConfigRequest proto.InternalMessageInfo

func (m *DeleteUptimeCheckConfigRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The protocol for the `ListUptimeCheckIps` request.
type ListUptimeCheckIpsRequest struct {
	// The maximum number of results to return in a single response. The server
	// may further constrain the maximum number of results returned in a single
	// page. If the page_size is <=0, the server will decide the number of results
	// to be returned.
	// NOTE: this field is not yet implemented
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// If this field is not empty then it must contain the `nextPageToken` value
	// returned by a previous call to this method.  Using this field causes the
	// method to return more results from the previous method call.
	// NOTE: this field is not yet implemented
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUptimeCheckIpsRequest) Reset()         { *m = ListUptimeCheckIpsRequest{} }
func (m *ListUptimeCheckIpsRequest) String() string { return proto.CompactTextString(m) }
func (*ListUptimeCheckIpsRequest) ProtoMessage()    {}
func (*ListUptimeCheckIpsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_uptime_service_c74f83143a8cb5a4, []int{6}
}
func (m *ListUptimeCheckIpsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUptimeCheckIpsRequest.Unmarshal(m, b)
}
func (m *ListUptimeCheckIpsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUptimeCheckIpsRequest.Marshal(b, m, deterministic)
}
func (dst *ListUptimeCheckIpsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUptimeCheckIpsRequest.Merge(dst, src)
}
func (m *ListUptimeCheckIpsRequest) XXX_Size() int {
	return xxx_messageInfo_ListUptimeCheckIpsRequest.Size(m)
}
func (m *ListUptimeCheckIpsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUptimeCheckIpsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUptimeCheckIpsRequest proto.InternalMessageInfo

func (m *ListUptimeCheckIpsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListUptimeCheckIpsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The protocol for the `ListUptimeCheckIps` response.
type ListUptimeCheckIpsResponse struct {
	// The returned list of IP addresses (including region and location) that the
	// checkers run from.
	UptimeCheckIps []*UptimeCheckIp `protobuf:"bytes,1,rep,name=uptime_check_ips,json=uptimeCheckIps,proto3" json:"uptime_check_ips,omitempty"`
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is empty, it means no further results for the
	// request. To retrieve the next page of results, the value of the
	// next_page_token is passed to the subsequent List method call (in the
	// request message's page_token field).
	// NOTE: this field is not yet implemented
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUptimeCheckIpsResponse) Reset()         { *m = ListUptimeCheckIpsResponse{} }
func (m *ListUptimeCheckIpsResponse) String() string { return proto.CompactTextString(m) }
func (*ListUptimeCheckIpsResponse) ProtoMessage()    {}
func (*ListUptimeCheckIpsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_uptime_service_c74f83143a8cb5a4, []int{7}
}
func (m *ListUptimeCheckIpsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUptimeCheckIpsResponse.Unmarshal(m, b)
}
func (m *ListUptimeCheckIpsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUptimeCheckIpsResponse.Marshal(b, m, deterministic)
}
func (dst *ListUptimeCheckIpsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUptimeCheckIpsResponse.Merge(dst, src)
}
func (m *ListUptimeCheckIpsResponse) XXX_Size() int {
	return xxx_messageInfo_ListUptimeCheckIpsResponse.Size(m)
}
func (m *ListUptimeCheckIpsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUptimeCheckIpsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUptimeCheckIpsResponse proto.InternalMessageInfo

func (m *ListUptimeCheckIpsResponse) GetUptimeCheckIps() []*UptimeCheckIp {
	if m != nil {
		return m.UptimeCheckIps
	}
	return nil
}

func (m *ListUptimeCheckIpsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*ListUptimeCheckConfigsRequest)(nil), "google.monitoring.v3.ListUptimeCheckConfigsRequest")
	proto.RegisterType((*ListUptimeCheckConfigsResponse)(nil), "google.monitoring.v3.ListUptimeCheckConfigsResponse")
	proto.RegisterType((*GetUptimeCheckConfigRequest)(nil), "google.monitoring.v3.GetUptimeCheckConfigRequest")
	proto.RegisterType((*CreateUptimeCheckConfigRequest)(nil), "google.monitoring.v3.CreateUptimeCheckConfigRequest")
	proto.RegisterType((*UpdateUptimeCheckConfigRequest)(nil), "google.monitoring.v3.UpdateUptimeCheckConfigRequest")
	proto.RegisterType((*DeleteUptimeCheckConfigRequest)(nil), "google.monitoring.v3.DeleteUptimeCheckConfigRequest")
	proto.RegisterType((*ListUptimeCheckIpsRequest)(nil), "google.monitoring.v3.ListUptimeCheckIpsRequest")
	proto.RegisterType((*ListUptimeCheckIpsResponse)(nil), "google.monitoring.v3.ListUptimeCheckIpsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UptimeCheckServiceClient is the client API for UptimeCheckService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UptimeCheckServiceClient interface {
	// Lists the existing valid uptime check configurations for the project,
	// leaving out any invalid configurations.
	ListUptimeCheckConfigs(ctx context.Context, in *ListUptimeCheckConfigsRequest, opts ...grpc.CallOption) (*ListUptimeCheckConfigsResponse, error)
	// Gets a single uptime check configuration.
	GetUptimeCheckConfig(ctx context.Context, in *GetUptimeCheckConfigRequest, opts ...grpc.CallOption) (*UptimeCheckConfig, error)
	// Creates a new uptime check configuration.
	CreateUptimeCheckConfig(ctx context.Context, in *CreateUptimeCheckConfigRequest, opts ...grpc.CallOption) (*UptimeCheckConfig, error)
	// Updates an uptime check configuration. You can either replace the entire
	// configuration with a new one or replace only certain fields in the current
	// configuration by specifying the fields to be updated via `"updateMask"`.
	// Returns the updated configuration.
	UpdateUptimeCheckConfig(ctx context.Context, in *UpdateUptimeCheckConfigRequest, opts ...grpc.CallOption) (*UptimeCheckConfig, error)
	// Deletes an uptime check configuration. Note that this method will fail
	// if the uptime check configuration is referenced by an alert policy or
	// other dependent configs that would be rendered invalid by the deletion.
	DeleteUptimeCheckConfig(ctx context.Context, in *DeleteUptimeCheckConfigRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Returns the list of IPs that checkers run from
	ListUptimeCheckIps(ctx context.Context, in *ListUptimeCheckIpsRequest, opts ...grpc.CallOption) (*ListUptimeCheckIpsResponse, error)
}

type uptimeCheckServiceClient struct {
	cc *grpc.ClientConn
}

func NewUptimeCheckServiceClient(cc *grpc.ClientConn) UptimeCheckServiceClient {
	return &uptimeCheckServiceClient{cc}
}

func (c *uptimeCheckServiceClient) ListUptimeCheckConfigs(ctx context.Context, in *ListUptimeCheckConfigsRequest, opts ...grpc.CallOption) (*ListUptimeCheckConfigsResponse, error) {
	out := new(ListUptimeCheckConfigsResponse)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.UptimeCheckService/ListUptimeCheckConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uptimeCheckServiceClient) GetUptimeCheckConfig(ctx context.Context, in *GetUptimeCheckConfigRequest, opts ...grpc.CallOption) (*UptimeCheckConfig, error) {
	out := new(UptimeCheckConfig)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.UptimeCheckService/GetUptimeCheckConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uptimeCheckServiceClient) CreateUptimeCheckConfig(ctx context.Context, in *CreateUptimeCheckConfigRequest, opts ...grpc.CallOption) (*UptimeCheckConfig, error) {
	out := new(UptimeCheckConfig)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.UptimeCheckService/CreateUptimeCheckConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uptimeCheckServiceClient) UpdateUptimeCheckConfig(ctx context.Context, in *UpdateUptimeCheckConfigRequest, opts ...grpc.CallOption) (*UptimeCheckConfig, error) {
	out := new(UptimeCheckConfig)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.UptimeCheckService/UpdateUptimeCheckConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uptimeCheckServiceClient) DeleteUptimeCheckConfig(ctx context.Context, in *DeleteUptimeCheckConfigRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.UptimeCheckService/DeleteUptimeCheckConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uptimeCheckServiceClient) ListUptimeCheckIps(ctx context.Context, in *ListUptimeCheckIpsRequest, opts ...grpc.CallOption) (*ListUptimeCheckIpsResponse, error) {
	out := new(ListUptimeCheckIpsResponse)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.UptimeCheckService/ListUptimeCheckIps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UptimeCheckServiceServer is the server API for UptimeCheckService service.
type UptimeCheckServiceServer interface {
	// Lists the existing valid uptime check configurations for the project,
	// leaving out any invalid configurations.
	ListUptimeCheckConfigs(context.Context, *ListUptimeCheckConfigsRequest) (*ListUptimeCheckConfigsResponse, error)
	// Gets a single uptime check configuration.
	GetUptimeCheckConfig(context.Context, *GetUptimeCheckConfigRequest) (*UptimeCheckConfig, error)
	// Creates a new uptime check configuration.
	CreateUptimeCheckConfig(context.Context, *CreateUptimeCheckConfigRequest) (*UptimeCheckConfig, error)
	// Updates an uptime check configuration. You can either replace the entire
	// configuration with a new one or replace only certain fields in the current
	// configuration by specifying the fields to be updated via `"updateMask"`.
	// Returns the updated configuration.
	UpdateUptimeCheckConfig(context.Context, *UpdateUptimeCheckConfigRequest) (*UptimeCheckConfig, error)
	// Deletes an uptime check configuration. Note that this method will fail
	// if the uptime check configuration is referenced by an alert policy or
	// other dependent configs that would be rendered invalid by the deletion.
	DeleteUptimeCheckConfig(context.Context, *DeleteUptimeCheckConfigRequest) (*empty.Empty, error)
	// Returns the list of IPs that checkers run from
	ListUptimeCheckIps(context.Context, *ListUptimeCheckIpsRequest) (*ListUptimeCheckIpsResponse, error)
}

func RegisterUptimeCheckServiceServer(s *grpc.Server, srv UptimeCheckServiceServer) {
	s.RegisterService(&_UptimeCheckService_serviceDesc, srv)
}

func _UptimeCheckService_ListUptimeCheckConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUptimeCheckConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UptimeCheckServiceServer).ListUptimeCheckConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.UptimeCheckService/ListUptimeCheckConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UptimeCheckServiceServer).ListUptimeCheckConfigs(ctx, req.(*ListUptimeCheckConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UptimeCheckService_GetUptimeCheckConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUptimeCheckConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UptimeCheckServiceServer).GetUptimeCheckConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.UptimeCheckService/GetUptimeCheckConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UptimeCheckServiceServer).GetUptimeCheckConfig(ctx, req.(*GetUptimeCheckConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UptimeCheckService_CreateUptimeCheckConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUptimeCheckConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UptimeCheckServiceServer).CreateUptimeCheckConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.UptimeCheckService/CreateUptimeCheckConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UptimeCheckServiceServer).CreateUptimeCheckConfig(ctx, req.(*CreateUptimeCheckConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UptimeCheckService_UpdateUptimeCheckConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUptimeCheckConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UptimeCheckServiceServer).UpdateUptimeCheckConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.UptimeCheckService/UpdateUptimeCheckConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UptimeCheckServiceServer).UpdateUptimeCheckConfig(ctx, req.(*UpdateUptimeCheckConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UptimeCheckService_DeleteUptimeCheckConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUptimeCheckConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UptimeCheckServiceServer).DeleteUptimeCheckConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.UptimeCheckService/DeleteUptimeCheckConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UptimeCheckServiceServer).DeleteUptimeCheckConfig(ctx, req.(*DeleteUptimeCheckConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UptimeCheckService_ListUptimeCheckIps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUptimeCheckIpsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UptimeCheckServiceServer).ListUptimeCheckIps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.UptimeCheckService/ListUptimeCheckIps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UptimeCheckServiceServer).ListUptimeCheckIps(ctx, req.(*ListUptimeCheckIpsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UptimeCheckService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.monitoring.v3.UptimeCheckService",
	HandlerType: (*UptimeCheckServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUptimeCheckConfigs",
			Handler:    _UptimeCheckService_ListUptimeCheckConfigs_Handler,
		},
		{
			MethodName: "GetUptimeCheckConfig",
			Handler:    _UptimeCheckService_GetUptimeCheckConfig_Handler,
		},
		{
			MethodName: "CreateUptimeCheckConfig",
			Handler:    _UptimeCheckService_CreateUptimeCheckConfig_Handler,
		},
		{
			MethodName: "UpdateUptimeCheckConfig",
			Handler:    _UptimeCheckService_UpdateUptimeCheckConfig_Handler,
		},
		{
			MethodName: "DeleteUptimeCheckConfig",
			Handler:    _UptimeCheckService_DeleteUptimeCheckConfig_Handler,
		},
		{
			MethodName: "ListUptimeCheckIps",
			Handler:    _UptimeCheckService_ListUptimeCheckIps_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/monitoring/v3/uptime_service.proto",
}

func init() {
	proto.RegisterFile("google/monitoring/v3/uptime_service.proto", fileDescriptor_uptime_service_c74f83143a8cb5a4)
}

var fileDescriptor_uptime_service_c74f83143a8cb5a4 = []byte{
	// 747 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xd6, 0x24, 0xbd, 0x55, 0x7b, 0xaa, 0x7b, 0x2f, 0x0c, 0x51, 0x1b, 0x5c, 0x1a, 0x05, 0x23,
	0x41, 0x89, 0x90, 0x4d, 0x93, 0xae, 0xa8, 0xa8, 0x44, 0x03, 0x54, 0x95, 0xa8, 0x54, 0xa5, 0xb4,
	0x15, 0x50, 0x29, 0x72, 0xd3, 0xa9, 0x31, 0x49, 0x3c, 0xc6, 0x33, 0xae, 0xa0, 0xa8, 0x1b, 0xde,
	0x00, 0x75, 0xc9, 0x9e, 0x45, 0x1f, 0x00, 0xd6, 0xb0, 0x41, 0x62, 0x8b, 0x78, 0x03, 0x1e, 0x04,
	0x79, 0x3c, 0x26, 0x7f, 0x63, 0xe3, 0x88, 0x5d, 0x3c, 0xe7, 0xcc, 0x39, 0xdf, 0xf9, 0xfc, 0x9d,
	0x2f, 0x86, 0x9b, 0x36, 0xa5, 0x76, 0x87, 0x98, 0x5d, 0xea, 0x3a, 0x9c, 0xfa, 0x8e, 0x6b, 0x9b,
	0xc7, 0x35, 0x33, 0xf0, 0xb8, 0xd3, 0x25, 0x4d, 0x46, 0xfc, 0x63, 0xa7, 0x45, 0x0c, 0xcf, 0xa7,
	0x9c, 0xe2, 0x42, 0x94, 0x6a, 0xf4, 0x52, 0x8d, 0xe3, 0x9a, 0x76, 0x45, 0x16, 0xb0, 0x3c, 0xc7,
	0xb4, 0x5c, 0x97, 0x72, 0x8b, 0x3b, 0xd4, 0x65, 0xd1, 0x1d, 0xed, 0x6a, 0x4a, 0x79, 0x99, 0x32,
	0x2f, 0x53, 0xc4, 0xd3, 0x41, 0x70, 0x64, 0x92, 0xae, 0xc7, 0x5f, 0xcb, 0x60, 0x79, 0x38, 0x78,
	0xe4, 0x90, 0xce, 0x61, 0xb3, 0x6b, 0xb1, 0x76, 0x94, 0xa1, 0x33, 0x58, 0x78, 0xe4, 0x30, 0xbe,
	0x23, 0x4a, 0xd6, 0x9f, 0x93, 0x56, 0xbb, 0x4e, 0xdd, 0x23, 0xc7, 0x66, 0x0d, 0xf2, 0x32, 0x20,
	0x8c, 0xe3, 0x59, 0x98, 0xf4, 0x2c, 0x9f, 0xb8, 0xbc, 0x88, 0xca, 0x68, 0x71, 0xba, 0x21, 0x9f,
	0xf0, 0x3c, 0x4c, 0x7b, 0x96, 0x4d, 0x9a, 0xcc, 0x39, 0x21, 0xc5, 0x7c, 0x19, 0x2d, 0xfe, 0xd3,
	0x98, 0x0a, 0x0f, 0xb6, 0x9d, 0x13, 0x82, 0x17, 0x00, 0x44, 0x90, 0xd3, 0x36, 0x71, 0x8b, 0x13,
	0xe2, 0xa2, 0x48, 0x7f, 0x1c, 0x1e, 0xe8, 0x5f, 0x10, 0x94, 0x92, 0xba, 0x32, 0x8f, 0xba, 0x8c,
	0xe0, 0x27, 0x50, 0x90, 0x2c, 0xb6, 0xc2, 0x70, 0xb3, 0x15, 0xc5, 0x8b, 0xa8, 0x9c, 0x5f, 0x9c,
	0xa9, 0xde, 0x30, 0x54, 0x64, 0x1a, 0x23, 0xf5, 0x1a, 0x38, 0x18, 0x69, 0x81, 0xaf, 0xc3, 0xff,
	0x2e, 0x79, 0xc5, 0x9b, 0x7d, 0x08, 0x73, 0x02, 0xe1, 0xbf, 0xe1, 0xf1, 0x56, 0x8c, 0x32, 0x1c,
	0x82, 0x53, 0x6e, 0x75, 0xfa, 0x47, 0x9c, 0x16, 0x27, 0xe1, 0x8c, 0xfa, 0x12, 0xcc, 0xaf, 0x93,
	0xd1, 0x11, 0x62, 0xde, 0x30, 0x4c, 0xb8, 0x56, 0x97, 0x48, 0xd6, 0xc4, 0x6f, 0xfd, 0x1d, 0x82,
	0x52, 0xdd, 0x27, 0x16, 0x27, 0x89, 0xd7, 0x92, 0xe8, 0xde, 0x83, 0x4b, 0x0a, 0x3e, 0x04, 0xf0,
	0x31, 0xe8, 0xb8, 0x38, 0x42, 0x87, 0xfe, 0x11, 0x41, 0x69, 0xc7, 0x3b, 0x4c, 0xc3, 0xb4, 0x02,
	0x33, 0x81, 0xc8, 0x10, 0xc2, 0x91, 0x3d, 0xb5, 0xb8, 0x67, 0xac, 0x2d, 0xe3, 0x61, 0xa8, 0xad,
	0x4d, 0x8b, 0xb5, 0x1b, 0x10, 0xa5, 0x87, 0xbf, 0x93, 0x80, 0xe7, 0xff, 0x1a, 0xf8, 0x32, 0x94,
	0xee, 0x93, 0x0e, 0x49, 0xc1, 0xad, 0x7a, 0x05, 0x7b, 0x70, 0x79, 0x48, 0x79, 0x1b, 0xde, 0x6f,
	0xad, 0x0f, 0x68, 0x3a, 0x97, 0xaa, 0xe9, 0xfc, 0xb0, 0xa6, 0xcf, 0x10, 0x68, 0xaa, 0xca, 0x52,
	0xcf, 0x9b, 0x70, 0x61, 0x80, 0x06, 0xc7, 0x8b, 0xb5, 0x7c, 0xed, 0x8f, 0x1c, 0x6c, 0x78, 0x8d,
	0xff, 0x82, 0x81, 0xb2, 0x59, 0x35, 0x5c, 0xfd, 0x3a, 0x05, 0xb8, 0xaf, 0xd2, 0x76, 0xe4, 0x48,
	0xf8, 0x13, 0x82, 0x59, 0xf5, 0x02, 0xe2, 0x9a, 0x1a, 0x4e, 0xaa, 0x49, 0x68, 0xcb, 0xe3, 0x5d,
	0x8a, 0x38, 0xd1, 0xab, 0x6f, 0xbf, 0xff, 0x3c, 0xcb, 0xdd, 0xc2, 0x95, 0xd0, 0xd4, 0xde, 0x44,
	0x42, 0xbf, 0xeb, 0xf9, 0xf4, 0x05, 0x69, 0x71, 0x66, 0x56, 0x4e, 0x4d, 0xc5, 0xf2, 0x7e, 0x40,
	0x50, 0x50, 0xad, 0x1d, 0x5e, 0x52, 0x43, 0x48, 0x59, 0x51, 0x2d, 0xab, 0xfa, 0x86, 0x80, 0x86,
	0x3a, 0xea, 0x83, 0xa9, 0x40, 0x69, 0x56, 0x4e, 0xf1, 0x67, 0x04, 0x73, 0x09, 0xbb, 0x8e, 0x13,
	0xe8, 0x4a, 0xb7, 0x86, 0xec, 0x70, 0xd7, 0x05, 0xdc, 0x7b, 0xfa, 0x18, 0xbc, 0xde, 0x51, 0x2d,
	0x29, 0xfe, 0x81, 0x60, 0x2e, 0xc1, 0x1b, 0x92, 0x66, 0x48, 0xb7, 0x92, 0xec, 0x33, 0x3c, 0x13,
	0x33, 0xec, 0x54, 0x57, 0xc5, 0x0c, 0x0a, 0x70, 0x46, 0xa6, 0xd7, 0xa0, 0x9e, 0xeb, 0x3d, 0x82,
	0xb9, 0x04, 0xef, 0x48, 0x9a, 0x2b, 0xdd, 0x6a, 0xb4, 0xd9, 0x11, 0x37, 0x7c, 0x10, 0xfe, 0x0d,
	0xc7, 0xca, 0xa9, 0x8c, 0xa3, 0x9c, 0x33, 0x04, 0x78, 0xd4, 0x49, 0xb0, 0x99, 0x69, 0xc7, 0x7a,
	0x6e, 0xa6, 0xdd, 0xce, 0x7e, 0x41, 0x2e, 0xa4, 0x26, 0xd0, 0x16, 0x30, 0xee, 0x7d, 0x65, 0xc4,
	0x39, 0x6b, 0xe7, 0x08, 0x8a, 0x2d, 0xda, 0x55, 0xd6, 0x5c, 0x93, 0x1e, 0x23, 0xed, 0x65, 0x2b,
	0xe4, 0x60, 0x0b, 0x3d, 0x5d, 0x95, 0xb9, 0x36, 0xed, 0x58, 0xae, 0x6d, 0x50, 0xdf, 0x36, 0x6d,
	0xe2, 0x0a, 0x86, 0xcc, 0x28, 0x64, 0x79, 0x0e, 0x1b, 0xfc, 0xb8, 0x59, 0xe9, 0x3d, 0x9d, 0xe7,
	0xb4, 0xf5, 0xa8, 0x40, 0xbd, 0x43, 0x83, 0x43, 0x63, 0xb3, 0xd7, 0x72, 0xb7, 0xf6, 0x2d, 0x0e,
	0xee, 0x8b, 0xe0, 0x7e, 0x2f, 0xb8, 0xbf, 0x5b, 0x3b, 0x98, 0x14, 0x4d, 0x6a, 0xbf, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x27, 0xb8, 0x65, 0x92, 0x9f, 0x09, 0x00, 0x00,
}
