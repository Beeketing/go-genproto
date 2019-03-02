// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/feed_service.proto

package services // import "github.com/Beeketing/go-genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/Beeketing/protobuf/ptypes/wrappers"
import resources "github.com/Beeketing/go-genproto/googleapis/ads/googleads/v0/resources"
import _ "github.com/Beeketing/go-genproto/googleapis/api/annotations"
import status "github.com/Beeketing/go-genproto/googleapis/rpc/status"
import field_mask "github.com/Beeketing/go-genproto/protobuf/field_mask"

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

// Request message for [FeedService.GetFeed][google.ads.googleads.v0.services.FeedService.GetFeed].
type GetFeedRequest struct {
	// The resource name of the feed to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedRequest) Reset()         { *m = GetFeedRequest{} }
func (m *GetFeedRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedRequest) ProtoMessage()    {}
func (*GetFeedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_service_dff6ab5365713504, []int{0}
}
func (m *GetFeedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedRequest.Unmarshal(m, b)
}
func (m *GetFeedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedRequest.Marshal(b, m, deterministic)
}
func (dst *GetFeedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedRequest.Merge(dst, src)
}
func (m *GetFeedRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedRequest.Size(m)
}
func (m *GetFeedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedRequest proto.InternalMessageInfo

func (m *GetFeedRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [FeedService.MutateFeeds][google.ads.googleads.v0.services.FeedService.MutateFeeds].
type MutateFeedsRequest struct {
	// The ID of the customer whose feeds are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual feeds.
	Operations []*FeedOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateFeedsRequest) Reset()         { *m = MutateFeedsRequest{} }
func (m *MutateFeedsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateFeedsRequest) ProtoMessage()    {}
func (*MutateFeedsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_service_dff6ab5365713504, []int{1}
}
func (m *MutateFeedsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedsRequest.Unmarshal(m, b)
}
func (m *MutateFeedsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedsRequest.Marshal(b, m, deterministic)
}
func (dst *MutateFeedsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedsRequest.Merge(dst, src)
}
func (m *MutateFeedsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateFeedsRequest.Size(m)
}
func (m *MutateFeedsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedsRequest proto.InternalMessageInfo

func (m *MutateFeedsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateFeedsRequest) GetOperations() []*FeedOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateFeedsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateFeedsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on an feed.
type FeedOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*FeedOperation_Create
	//	*FeedOperation_Update
	//	*FeedOperation_Remove
	Operation            isFeedOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *FeedOperation) Reset()         { *m = FeedOperation{} }
func (m *FeedOperation) String() string { return proto.CompactTextString(m) }
func (*FeedOperation) ProtoMessage()    {}
func (*FeedOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_service_dff6ab5365713504, []int{2}
}
func (m *FeedOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedOperation.Unmarshal(m, b)
}
func (m *FeedOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedOperation.Marshal(b, m, deterministic)
}
func (dst *FeedOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedOperation.Merge(dst, src)
}
func (m *FeedOperation) XXX_Size() int {
	return xxx_messageInfo_FeedOperation.Size(m)
}
func (m *FeedOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedOperation.DiscardUnknown(m)
}

var xxx_messageInfo_FeedOperation proto.InternalMessageInfo

func (m *FeedOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isFeedOperation_Operation interface {
	isFeedOperation_Operation()
}

type FeedOperation_Create struct {
	Create *resources.Feed `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type FeedOperation_Update struct {
	Update *resources.Feed `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type FeedOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*FeedOperation_Create) isFeedOperation_Operation() {}

func (*FeedOperation_Update) isFeedOperation_Operation() {}

func (*FeedOperation_Remove) isFeedOperation_Operation() {}

func (m *FeedOperation) GetOperation() isFeedOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *FeedOperation) GetCreate() *resources.Feed {
	if x, ok := m.GetOperation().(*FeedOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *FeedOperation) GetUpdate() *resources.Feed {
	if x, ok := m.GetOperation().(*FeedOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *FeedOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*FeedOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FeedOperation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FeedOperation_OneofMarshaler, _FeedOperation_OneofUnmarshaler, _FeedOperation_OneofSizer, []interface{}{
		(*FeedOperation_Create)(nil),
		(*FeedOperation_Update)(nil),
		(*FeedOperation_Remove)(nil),
	}
}

func _FeedOperation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FeedOperation)
	// operation
	switch x := m.Operation.(type) {
	case *FeedOperation_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *FeedOperation_Update:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *FeedOperation_Remove:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Remove)
	case nil:
	default:
		return fmt.Errorf("FeedOperation.Operation has unexpected type %T", x)
	}
	return nil
}

func _FeedOperation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FeedOperation)
	switch tag {
	case 1: // operation.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.Feed)
		err := b.DecodeMessage(msg)
		m.Operation = &FeedOperation_Create{msg}
		return true, err
	case 2: // operation.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.Feed)
		err := b.DecodeMessage(msg)
		m.Operation = &FeedOperation_Update{msg}
		return true, err
	case 3: // operation.remove
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operation = &FeedOperation_Remove{x}
		return true, err
	default:
		return false, nil
	}
}

func _FeedOperation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FeedOperation)
	// operation
	switch x := m.Operation.(type) {
	case *FeedOperation_Create:
		s := proto.Size(x.Create)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *FeedOperation_Update:
		s := proto.Size(x.Update)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *FeedOperation_Remove:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Remove)))
		n += len(x.Remove)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for an feed mutate.
type MutateFeedsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateFeedResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MutateFeedsResponse) Reset()         { *m = MutateFeedsResponse{} }
func (m *MutateFeedsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateFeedsResponse) ProtoMessage()    {}
func (*MutateFeedsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_service_dff6ab5365713504, []int{3}
}
func (m *MutateFeedsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedsResponse.Unmarshal(m, b)
}
func (m *MutateFeedsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedsResponse.Marshal(b, m, deterministic)
}
func (dst *MutateFeedsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedsResponse.Merge(dst, src)
}
func (m *MutateFeedsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateFeedsResponse.Size(m)
}
func (m *MutateFeedsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedsResponse proto.InternalMessageInfo

func (m *MutateFeedsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateFeedsResponse) GetResults() []*MutateFeedResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the feed mutate.
type MutateFeedResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateFeedResult) Reset()         { *m = MutateFeedResult{} }
func (m *MutateFeedResult) String() string { return proto.CompactTextString(m) }
func (*MutateFeedResult) ProtoMessage()    {}
func (*MutateFeedResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_service_dff6ab5365713504, []int{4}
}
func (m *MutateFeedResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedResult.Unmarshal(m, b)
}
func (m *MutateFeedResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedResult.Marshal(b, m, deterministic)
}
func (dst *MutateFeedResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedResult.Merge(dst, src)
}
func (m *MutateFeedResult) XXX_Size() int {
	return xxx_messageInfo_MutateFeedResult.Size(m)
}
func (m *MutateFeedResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedResult proto.InternalMessageInfo

func (m *MutateFeedResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFeedRequest)(nil), "google.ads.googleads.v0.services.GetFeedRequest")
	proto.RegisterType((*MutateFeedsRequest)(nil), "google.ads.googleads.v0.services.MutateFeedsRequest")
	proto.RegisterType((*FeedOperation)(nil), "google.ads.googleads.v0.services.FeedOperation")
	proto.RegisterType((*MutateFeedsResponse)(nil), "google.ads.googleads.v0.services.MutateFeedsResponse")
	proto.RegisterType((*MutateFeedResult)(nil), "google.ads.googleads.v0.services.MutateFeedResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FeedServiceClient is the client API for FeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedServiceClient interface {
	// Returns the requested feed in full detail.
	GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*resources.Feed, error)
	// Creates, updates, or removes feeds. Operation statuses are
	// returned.
	MutateFeeds(ctx context.Context, in *MutateFeedsRequest, opts ...grpc.CallOption) (*MutateFeedsResponse, error)
}

type feedServiceClient struct {
	cc *grpc.ClientConn
}

func NewFeedServiceClient(cc *grpc.ClientConn) FeedServiceClient {
	return &feedServiceClient{cc}
}

func (c *feedServiceClient) GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*resources.Feed, error) {
	out := new(resources.Feed)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.FeedService/GetFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) MutateFeeds(ctx context.Context, in *MutateFeedsRequest, opts ...grpc.CallOption) (*MutateFeedsResponse, error) {
	out := new(MutateFeedsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.FeedService/MutateFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedServiceServer is the server API for FeedService service.
type FeedServiceServer interface {
	// Returns the requested feed in full detail.
	GetFeed(context.Context, *GetFeedRequest) (*resources.Feed, error)
	// Creates, updates, or removes feeds. Operation statuses are
	// returned.
	MutateFeeds(context.Context, *MutateFeedsRequest) (*MutateFeedsResponse, error)
}

func RegisterFeedServiceServer(s *grpc.Server, srv FeedServiceServer) {
	s.RegisterService(&_FeedService_serviceDesc, srv)
}

func _FeedService_GetFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).GetFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.FeedService/GetFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).GetFeed(ctx, req.(*GetFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_MutateFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).MutateFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.FeedService/MutateFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).MutateFeeds(ctx, req.(*MutateFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.FeedService",
	HandlerType: (*FeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeed",
			Handler:    _FeedService_GetFeed_Handler,
		},
		{
			MethodName: "MutateFeeds",
			Handler:    _FeedService_MutateFeeds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/feed_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/feed_service.proto", fileDescriptor_feed_service_dff6ab5365713504)
}

var fileDescriptor_feed_service_dff6ab5365713504 = []byte{
	// 695 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x41, 0x6b, 0x13, 0x4f,
	0x14, 0xff, 0x6f, 0xfa, 0xa7, 0xb5, 0xb3, 0x6d, 0x2d, 0x53, 0xc4, 0x10, 0x44, 0xc3, 0x5a, 0x68,
	0x0d, 0xb2, 0x13, 0x52, 0x8b, 0xb0, 0xa5, 0x87, 0x14, 0x4c, 0x2b, 0x58, 0x5b, 0xb6, 0xd0, 0x83,
	0x04, 0x96, 0x69, 0xf6, 0x25, 0x2c, 0xdd, 0xdd, 0x59, 0x67, 0x66, 0x23, 0xa5, 0xf4, 0xe2, 0x57,
	0x10, 0xbf, 0x80, 0x47, 0x3d, 0xf9, 0x35, 0xbc, 0xea, 0xc5, 0xbb, 0x27, 0xbf, 0x80, 0x78, 0x93,
	0xd9, 0xd9, 0x49, 0x93, 0x42, 0x49, 0x7b, 0x7b, 0xfb, 0xe6, 0xf7, 0xfb, 0xbd, 0xdf, 0xbc, 0xb7,
	0x6f, 0xd0, 0xc6, 0x80, 0xb1, 0x41, 0x0c, 0x84, 0x86, 0x82, 0xe8, 0x50, 0x45, 0xc3, 0x26, 0x11,
	0xc0, 0x87, 0x51, 0x0f, 0x04, 0xe9, 0x03, 0x84, 0x41, 0xf9, 0xe5, 0x66, 0x9c, 0x49, 0x86, 0xeb,
	0x1a, 0xe9, 0xd2, 0x50, 0xb8, 0x23, 0x92, 0x3b, 0x6c, 0xba, 0x86, 0x54, 0x7b, 0x7a, 0x9d, 0x2c,
	0x07, 0xc1, 0x72, 0x6e, 0x74, 0xb5, 0x5e, 0xed, 0x81, 0x41, 0x67, 0x11, 0xa1, 0x69, 0xca, 0x24,
	0x95, 0x11, 0x4b, 0x45, 0x79, 0x5a, 0x56, 0x23, 0xc5, 0xd7, 0x49, 0xde, 0x27, 0xfd, 0x08, 0xe2,
	0x30, 0x48, 0xa8, 0x38, 0x2d, 0x11, 0x0f, 0xaf, 0x22, 0xde, 0x71, 0x9a, 0x65, 0xc0, 0x8d, 0xc2,
	0xfd, 0xf2, 0x9c, 0x67, 0x3d, 0x22, 0x24, 0x95, 0x79, 0x79, 0xe0, 0x6c, 0xa2, 0xa5, 0x5d, 0x90,
	0x1d, 0x80, 0xd0, 0x87, 0xb7, 0x39, 0x08, 0x89, 0x1f, 0xa3, 0x45, 0x63, 0x31, 0x48, 0x69, 0x02,
	0x55, 0xab, 0x6e, 0xad, 0xcf, 0xfb, 0x0b, 0x26, 0xf9, 0x9a, 0x26, 0xe0, 0xfc, 0xb0, 0x10, 0xde,
	0xcf, 0x25, 0x95, 0xa0, 0xa8, 0xc2, 0x70, 0x1f, 0x21, 0xbb, 0x97, 0x0b, 0xc9, 0x12, 0xe0, 0x41,
	0x14, 0x96, 0x4c, 0x64, 0x52, 0x2f, 0x43, 0x7c, 0x80, 0x10, 0xcb, 0x80, 0xeb, 0xdb, 0x55, 0x2b,
	0xf5, 0x99, 0x75, 0xbb, 0x45, 0xdc, 0x69, 0xcd, 0x74, 0x55, 0x91, 0x03, 0xc3, 0xf3, 0xc7, 0x24,
	0xf0, 0x1a, 0xba, 0x9b, 0x51, 0x2e, 0x23, 0x1a, 0x07, 0x7d, 0x1a, 0xc5, 0x39, 0x87, 0xea, 0x4c,
	0xdd, 0x5a, 0xbf, 0xe3, 0x2f, 0x95, 0xe9, 0x8e, 0xce, 0xaa, 0x6b, 0x0d, 0x69, 0x1c, 0x85, 0x54,
	0x42, 0xc0, 0xd2, 0xf8, 0xac, 0xfa, 0x7f, 0x01, 0x5b, 0x30, 0xc9, 0x83, 0x34, 0x3e, 0x73, 0xfe,
	0x5a, 0x68, 0x71, 0xa2, 0x16, 0xde, 0x42, 0x76, 0x9e, 0x15, 0x24, 0xd5, 0xed, 0x82, 0x64, 0xb7,
	0x6a, 0xc6, 0xb1, 0x69, 0xb7, 0xdb, 0x51, 0x03, 0xd9, 0xa7, 0xe2, 0xd4, 0x47, 0x1a, 0xae, 0x62,
	0xdc, 0x46, 0xb3, 0x3d, 0x0e, 0x54, 0xea, 0x1e, 0xda, 0xad, 0xb5, 0x6b, 0x6f, 0x3a, 0xfa, 0x29,
	0x8a, 0xab, 0xee, 0xfd, 0xe7, 0x97, 0x44, 0x25, 0xa1, 0x05, 0xab, 0x95, 0x5b, 0x4b, 0x68, 0x22,
	0xae, 0xa2, 0x59, 0x0e, 0x09, 0x1b, 0xea, 0xce, 0xcc, 0xab, 0x13, 0xfd, 0xbd, 0x63, 0xa3, 0xf9,
	0x51, 0x2b, 0x9d, 0x2f, 0x16, 0x5a, 0x99, 0x18, 0xa9, 0xc8, 0x58, 0x2a, 0x00, 0x77, 0xd0, 0xbd,
	0x2b, 0x1d, 0x0e, 0x80, 0x73, 0xc6, 0x0b, 0x35, 0xbb, 0x85, 0x8d, 0x21, 0x9e, 0xf5, 0xdc, 0xa3,
	0xe2, 0xd7, 0xf2, 0x57, 0x26, 0x7b, 0xff, 0x42, 0xc1, 0xf1, 0x2b, 0x34, 0xc7, 0x41, 0xe4, 0xb1,
	0x34, 0x73, 0x6f, 0x4d, 0x9f, 0xfb, 0xa5, 0x1f, 0xbf, 0xa0, 0xfa, 0x46, 0xc2, 0x79, 0x8e, 0x96,
	0xaf, 0x1e, 0xde, 0xe8, 0xcf, 0x6d, 0xfd, 0xac, 0x20, 0x5b, 0x71, 0x8e, 0x74, 0x0d, 0xfc, 0xd1,
	0x42, 0x73, 0xe5, 0x06, 0xe0, 0xe6, 0x74, 0x47, 0x93, 0xcb, 0x52, 0xbb, 0xe9, 0x38, 0x1c, 0xf2,
	0xfe, 0xfb, 0xaf, 0x0f, 0x95, 0x27, 0x78, 0x4d, 0x3d, 0x01, 0xe7, 0x13, 0x36, 0xb7, 0xcd, 0x7e,
	0x08, 0xd2, 0x28, 0xde, 0x04, 0x41, 0x1a, 0x17, 0xf8, 0xab, 0x85, 0xec, 0xb1, 0x71, 0xe0, 0x67,
	0xb7, 0xe9, 0x96, 0x59, 0xc8, 0xda, 0xe6, 0x2d, 0x59, 0x7a, 0xe6, 0xce, 0x66, 0xe1, 0x96, 0x38,
	0x0d, 0xe5, 0xf6, 0xd2, 0xde, 0xf9, 0xd8, 0x72, 0x6f, 0x37, 0x2e, 0xb4, 0x59, 0x2f, 0x29, 0x04,
	0x3c, 0xab, 0xb1, 0xf3, 0xc7, 0x42, 0xab, 0x3d, 0x96, 0x4c, 0xad, 0xb9, 0xb3, 0x3c, 0x36, 0x81,
	0x43, 0xb5, 0x43, 0x87, 0xd6, 0x9b, 0xbd, 0x92, 0x35, 0x60, 0x31, 0x4d, 0x07, 0x2e, 0xe3, 0x03,
	0x32, 0x80, 0xb4, 0xd8, 0x30, 0xf3, 0x80, 0x66, 0x91, 0xb8, 0xfe, 0x99, 0xde, 0x32, 0xc1, 0xa7,
	0xca, 0xcc, 0x6e, 0xbb, 0xfd, 0xb9, 0x52, 0xdf, 0xd5, 0x82, 0xed, 0x50, 0xb8, 0x3a, 0x54, 0xd1,
	0x71, 0xd3, 0x2d, 0x0b, 0x8b, 0x6f, 0x06, 0xd2, 0x6d, 0x87, 0xa2, 0x3b, 0x82, 0x74, 0x8f, 0x9b,
	0x5d, 0x03, 0xf9, 0x5d, 0x59, 0xd5, 0x79, 0xcf, 0x6b, 0x87, 0xc2, 0xf3, 0x46, 0x20, 0xcf, 0x3b,
	0x6e, 0x7a, 0x9e, 0x81, 0x9d, 0xcc, 0x16, 0x3e, 0x37, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x0d,
	0x61, 0x14, 0xd1, 0x4d, 0x06, 0x00, 0x00,
}
