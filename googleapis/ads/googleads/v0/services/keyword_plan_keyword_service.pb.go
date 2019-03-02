// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/keyword_plan_keyword_service.proto

package services // import "github.com/Beeketing/genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/Beeketing/protobuf/ptypes/wrappers"
import resources "github.com/Beeketing/genproto/googleapis/ads/googleads/v0/resources"
import _ "github.com/Beeketing/genproto/googleapis/api/annotations"
import status "github.com/Beeketing/genproto/googleapis/rpc/status"
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

// Request message for [KeywordPlanKeywordService.GetKeywordPlanKeyword][google.ads.googleads.v0.services.KeywordPlanKeywordService.GetKeywordPlanKeyword].
type GetKeywordPlanKeywordRequest struct {
	// The resource name of the ad group keyword to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKeywordPlanKeywordRequest) Reset()         { *m = GetKeywordPlanKeywordRequest{} }
func (m *GetKeywordPlanKeywordRequest) String() string { return proto.CompactTextString(m) }
func (*GetKeywordPlanKeywordRequest) ProtoMessage()    {}
func (*GetKeywordPlanKeywordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_keyword_service_9676619440e5ca2d, []int{0}
}
func (m *GetKeywordPlanKeywordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKeywordPlanKeywordRequest.Unmarshal(m, b)
}
func (m *GetKeywordPlanKeywordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKeywordPlanKeywordRequest.Marshal(b, m, deterministic)
}
func (dst *GetKeywordPlanKeywordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKeywordPlanKeywordRequest.Merge(dst, src)
}
func (m *GetKeywordPlanKeywordRequest) XXX_Size() int {
	return xxx_messageInfo_GetKeywordPlanKeywordRequest.Size(m)
}
func (m *GetKeywordPlanKeywordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKeywordPlanKeywordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetKeywordPlanKeywordRequest proto.InternalMessageInfo

func (m *GetKeywordPlanKeywordRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [KeywordPlanKeywordService.MutateKeywordPlanKeywords][google.ads.googleads.v0.services.KeywordPlanKeywordService.MutateKeywordPlanKeywords].
type MutateKeywordPlanKeywordsRequest struct {
	// The ID of the customer whose Keyword Plan keywords are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual Keyword Plan keywords.
	Operations []*KeywordPlanKeywordOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateKeywordPlanKeywordsRequest) Reset()         { *m = MutateKeywordPlanKeywordsRequest{} }
func (m *MutateKeywordPlanKeywordsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateKeywordPlanKeywordsRequest) ProtoMessage()    {}
func (*MutateKeywordPlanKeywordsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_keyword_service_9676619440e5ca2d, []int{1}
}
func (m *MutateKeywordPlanKeywordsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateKeywordPlanKeywordsRequest.Unmarshal(m, b)
}
func (m *MutateKeywordPlanKeywordsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateKeywordPlanKeywordsRequest.Marshal(b, m, deterministic)
}
func (dst *MutateKeywordPlanKeywordsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateKeywordPlanKeywordsRequest.Merge(dst, src)
}
func (m *MutateKeywordPlanKeywordsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateKeywordPlanKeywordsRequest.Size(m)
}
func (m *MutateKeywordPlanKeywordsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateKeywordPlanKeywordsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateKeywordPlanKeywordsRequest proto.InternalMessageInfo

func (m *MutateKeywordPlanKeywordsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateKeywordPlanKeywordsRequest) GetOperations() []*KeywordPlanKeywordOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateKeywordPlanKeywordsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateKeywordPlanKeywordsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on a Keyword Plan keyword.
type KeywordPlanKeywordOperation struct {
	// The FieldMask that determines which resource fields are modified in an
	// update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*KeywordPlanKeywordOperation_Create
	//	*KeywordPlanKeywordOperation_Update
	//	*KeywordPlanKeywordOperation_Remove
	Operation            isKeywordPlanKeywordOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *KeywordPlanKeywordOperation) Reset()         { *m = KeywordPlanKeywordOperation{} }
func (m *KeywordPlanKeywordOperation) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanKeywordOperation) ProtoMessage()    {}
func (*KeywordPlanKeywordOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_keyword_service_9676619440e5ca2d, []int{2}
}
func (m *KeywordPlanKeywordOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanKeywordOperation.Unmarshal(m, b)
}
func (m *KeywordPlanKeywordOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanKeywordOperation.Marshal(b, m, deterministic)
}
func (dst *KeywordPlanKeywordOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanKeywordOperation.Merge(dst, src)
}
func (m *KeywordPlanKeywordOperation) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanKeywordOperation.Size(m)
}
func (m *KeywordPlanKeywordOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanKeywordOperation.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanKeywordOperation proto.InternalMessageInfo

func (m *KeywordPlanKeywordOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isKeywordPlanKeywordOperation_Operation interface {
	isKeywordPlanKeywordOperation_Operation()
}

type KeywordPlanKeywordOperation_Create struct {
	Create *resources.KeywordPlanKeyword `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type KeywordPlanKeywordOperation_Update struct {
	Update *resources.KeywordPlanKeyword `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type KeywordPlanKeywordOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*KeywordPlanKeywordOperation_Create) isKeywordPlanKeywordOperation_Operation() {}

func (*KeywordPlanKeywordOperation_Update) isKeywordPlanKeywordOperation_Operation() {}

func (*KeywordPlanKeywordOperation_Remove) isKeywordPlanKeywordOperation_Operation() {}

func (m *KeywordPlanKeywordOperation) GetOperation() isKeywordPlanKeywordOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *KeywordPlanKeywordOperation) GetCreate() *resources.KeywordPlanKeyword {
	if x, ok := m.GetOperation().(*KeywordPlanKeywordOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *KeywordPlanKeywordOperation) GetUpdate() *resources.KeywordPlanKeyword {
	if x, ok := m.GetOperation().(*KeywordPlanKeywordOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *KeywordPlanKeywordOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*KeywordPlanKeywordOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*KeywordPlanKeywordOperation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _KeywordPlanKeywordOperation_OneofMarshaler, _KeywordPlanKeywordOperation_OneofUnmarshaler, _KeywordPlanKeywordOperation_OneofSizer, []interface{}{
		(*KeywordPlanKeywordOperation_Create)(nil),
		(*KeywordPlanKeywordOperation_Update)(nil),
		(*KeywordPlanKeywordOperation_Remove)(nil),
	}
}

func _KeywordPlanKeywordOperation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*KeywordPlanKeywordOperation)
	// operation
	switch x := m.Operation.(type) {
	case *KeywordPlanKeywordOperation_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *KeywordPlanKeywordOperation_Update:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *KeywordPlanKeywordOperation_Remove:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Remove)
	case nil:
	default:
		return fmt.Errorf("KeywordPlanKeywordOperation.Operation has unexpected type %T", x)
	}
	return nil
}

func _KeywordPlanKeywordOperation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*KeywordPlanKeywordOperation)
	switch tag {
	case 1: // operation.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.KeywordPlanKeyword)
		err := b.DecodeMessage(msg)
		m.Operation = &KeywordPlanKeywordOperation_Create{msg}
		return true, err
	case 2: // operation.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.KeywordPlanKeyword)
		err := b.DecodeMessage(msg)
		m.Operation = &KeywordPlanKeywordOperation_Update{msg}
		return true, err
	case 3: // operation.remove
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operation = &KeywordPlanKeywordOperation_Remove{x}
		return true, err
	default:
		return false, nil
	}
}

func _KeywordPlanKeywordOperation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*KeywordPlanKeywordOperation)
	// operation
	switch x := m.Operation.(type) {
	case *KeywordPlanKeywordOperation_Create:
		s := proto.Size(x.Create)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *KeywordPlanKeywordOperation_Update:
		s := proto.Size(x.Update)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *KeywordPlanKeywordOperation_Remove:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Remove)))
		n += len(x.Remove)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for a Keyword Plan keyword mutate.
type MutateKeywordPlanKeywordsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateKeywordPlanKeywordResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *MutateKeywordPlanKeywordsResponse) Reset()         { *m = MutateKeywordPlanKeywordsResponse{} }
func (m *MutateKeywordPlanKeywordsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateKeywordPlanKeywordsResponse) ProtoMessage()    {}
func (*MutateKeywordPlanKeywordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_keyword_service_9676619440e5ca2d, []int{3}
}
func (m *MutateKeywordPlanKeywordsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateKeywordPlanKeywordsResponse.Unmarshal(m, b)
}
func (m *MutateKeywordPlanKeywordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateKeywordPlanKeywordsResponse.Marshal(b, m, deterministic)
}
func (dst *MutateKeywordPlanKeywordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateKeywordPlanKeywordsResponse.Merge(dst, src)
}
func (m *MutateKeywordPlanKeywordsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateKeywordPlanKeywordsResponse.Size(m)
}
func (m *MutateKeywordPlanKeywordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateKeywordPlanKeywordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateKeywordPlanKeywordsResponse proto.InternalMessageInfo

func (m *MutateKeywordPlanKeywordsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateKeywordPlanKeywordsResponse) GetResults() []*MutateKeywordPlanKeywordResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the Keyword Plan keyword mutate.
type MutateKeywordPlanKeywordResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateKeywordPlanKeywordResult) Reset()         { *m = MutateKeywordPlanKeywordResult{} }
func (m *MutateKeywordPlanKeywordResult) String() string { return proto.CompactTextString(m) }
func (*MutateKeywordPlanKeywordResult) ProtoMessage()    {}
func (*MutateKeywordPlanKeywordResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_keyword_service_9676619440e5ca2d, []int{4}
}
func (m *MutateKeywordPlanKeywordResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateKeywordPlanKeywordResult.Unmarshal(m, b)
}
func (m *MutateKeywordPlanKeywordResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateKeywordPlanKeywordResult.Marshal(b, m, deterministic)
}
func (dst *MutateKeywordPlanKeywordResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateKeywordPlanKeywordResult.Merge(dst, src)
}
func (m *MutateKeywordPlanKeywordResult) XXX_Size() int {
	return xxx_messageInfo_MutateKeywordPlanKeywordResult.Size(m)
}
func (m *MutateKeywordPlanKeywordResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateKeywordPlanKeywordResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateKeywordPlanKeywordResult proto.InternalMessageInfo

func (m *MutateKeywordPlanKeywordResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetKeywordPlanKeywordRequest)(nil), "google.ads.googleads.v0.services.GetKeywordPlanKeywordRequest")
	proto.RegisterType((*MutateKeywordPlanKeywordsRequest)(nil), "google.ads.googleads.v0.services.MutateKeywordPlanKeywordsRequest")
	proto.RegisterType((*KeywordPlanKeywordOperation)(nil), "google.ads.googleads.v0.services.KeywordPlanKeywordOperation")
	proto.RegisterType((*MutateKeywordPlanKeywordsResponse)(nil), "google.ads.googleads.v0.services.MutateKeywordPlanKeywordsResponse")
	proto.RegisterType((*MutateKeywordPlanKeywordResult)(nil), "google.ads.googleads.v0.services.MutateKeywordPlanKeywordResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeywordPlanKeywordServiceClient is the client API for KeywordPlanKeywordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordPlanKeywordServiceClient interface {
	// Returns the requested Keyword Plan keyword in full detail.
	GetKeywordPlanKeyword(ctx context.Context, in *GetKeywordPlanKeywordRequest, opts ...grpc.CallOption) (*resources.KeywordPlanKeyword, error)
	// Creates, updates, or removes Keyword Plan keywords. Operation statuses are
	// returned.
	MutateKeywordPlanKeywords(ctx context.Context, in *MutateKeywordPlanKeywordsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanKeywordsResponse, error)
}

type keywordPlanKeywordServiceClient struct {
	cc *grpc.ClientConn
}

func NewKeywordPlanKeywordServiceClient(cc *grpc.ClientConn) KeywordPlanKeywordServiceClient {
	return &keywordPlanKeywordServiceClient{cc}
}

func (c *keywordPlanKeywordServiceClient) GetKeywordPlanKeyword(ctx context.Context, in *GetKeywordPlanKeywordRequest, opts ...grpc.CallOption) (*resources.KeywordPlanKeyword, error) {
	out := new(resources.KeywordPlanKeyword)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.KeywordPlanKeywordService/GetKeywordPlanKeyword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordPlanKeywordServiceClient) MutateKeywordPlanKeywords(ctx context.Context, in *MutateKeywordPlanKeywordsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanKeywordsResponse, error) {
	out := new(MutateKeywordPlanKeywordsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.KeywordPlanKeywordService/MutateKeywordPlanKeywords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanKeywordServiceServer is the server API for KeywordPlanKeywordService service.
type KeywordPlanKeywordServiceServer interface {
	// Returns the requested Keyword Plan keyword in full detail.
	GetKeywordPlanKeyword(context.Context, *GetKeywordPlanKeywordRequest) (*resources.KeywordPlanKeyword, error)
	// Creates, updates, or removes Keyword Plan keywords. Operation statuses are
	// returned.
	MutateKeywordPlanKeywords(context.Context, *MutateKeywordPlanKeywordsRequest) (*MutateKeywordPlanKeywordsResponse, error)
}

func RegisterKeywordPlanKeywordServiceServer(s *grpc.Server, srv KeywordPlanKeywordServiceServer) {
	s.RegisterService(&_KeywordPlanKeywordService_serviceDesc, srv)
}

func _KeywordPlanKeywordService_GetKeywordPlanKeyword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeywordPlanKeywordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanKeywordServiceServer).GetKeywordPlanKeyword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.KeywordPlanKeywordService/GetKeywordPlanKeyword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanKeywordServiceServer).GetKeywordPlanKeyword(ctx, req.(*GetKeywordPlanKeywordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeywordPlanKeywordService_MutateKeywordPlanKeywords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateKeywordPlanKeywordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanKeywordServiceServer).MutateKeywordPlanKeywords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.KeywordPlanKeywordService/MutateKeywordPlanKeywords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanKeywordServiceServer).MutateKeywordPlanKeywords(ctx, req.(*MutateKeywordPlanKeywordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeywordPlanKeywordService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.KeywordPlanKeywordService",
	HandlerType: (*KeywordPlanKeywordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKeywordPlanKeyword",
			Handler:    _KeywordPlanKeywordService_GetKeywordPlanKeyword_Handler,
		},
		{
			MethodName: "MutateKeywordPlanKeywords",
			Handler:    _KeywordPlanKeywordService_MutateKeywordPlanKeywords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/keyword_plan_keyword_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/keyword_plan_keyword_service.proto", fileDescriptor_keyword_plan_keyword_service_9676619440e5ca2d)
}

var fileDescriptor_keyword_plan_keyword_service_9676619440e5ca2d = []byte{
	// 721 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xdd, 0x6a, 0xd4, 0x4e,
	0x14, 0xc0, 0xff, 0xc9, 0xfe, 0xa9, 0x76, 0x52, 0x15, 0x46, 0x8a, 0xdb, 0xb5, 0xd4, 0x35, 0x16,
	0x2c, 0x7b, 0x91, 0x2c, 0x2b, 0x45, 0x49, 0x5d, 0x71, 0x77, 0xe9, 0x87, 0x48, 0x6d, 0x49, 0xa1,
	0x17, 0x65, 0x25, 0x4c, 0x37, 0xd3, 0x25, 0x34, 0xc9, 0xc4, 0x99, 0xc9, 0x96, 0x52, 0x7a, 0x23,
	0xf8, 0x04, 0xbe, 0x81, 0xde, 0xf9, 0x22, 0x82, 0xe0, 0x95, 0x17, 0xbe, 0x80, 0x37, 0x7a, 0xe5,
	0x23, 0x48, 0x32, 0x33, 0x6b, 0xbf, 0xb2, 0x2b, 0xed, 0xdd, 0x99, 0x33, 0x27, 0xbf, 0xf3, 0x39,
	0x27, 0xa0, 0xd3, 0x27, 0xa4, 0x1f, 0x62, 0x1b, 0xf9, 0xcc, 0x16, 0x62, 0x26, 0x0d, 0xea, 0x36,
	0xc3, 0x74, 0x10, 0xf4, 0x30, 0xb3, 0xf7, 0xf1, 0xe1, 0x01, 0xa1, 0xbe, 0x97, 0x84, 0x28, 0xf6,
	0xd4, 0x41, 0xde, 0x5a, 0x09, 0x25, 0x9c, 0xc0, 0xaa, 0xf8, 0xd2, 0x42, 0x3e, 0xb3, 0x86, 0x10,
	0x6b, 0x50, 0xb7, 0x14, 0xa4, 0xf2, 0xb4, 0xc8, 0x0d, 0xc5, 0x8c, 0xa4, 0xb4, 0xc8, 0x8f, 0xe0,
	0x57, 0x66, 0xd5, 0xd7, 0x49, 0x60, 0xa3, 0x38, 0x26, 0x1c, 0xf1, 0x80, 0xc4, 0x4c, 0xde, 0x4a,
	0xef, 0x76, 0x7e, 0xda, 0x4d, 0xf7, 0xec, 0xbd, 0x00, 0x87, 0xbe, 0x17, 0x21, 0xb6, 0x2f, 0x2d,
	0xe6, 0xce, 0x5a, 0x1c, 0x50, 0x94, 0x24, 0x98, 0x2a, 0xc2, 0x1d, 0x79, 0x4f, 0x93, 0x9e, 0xcd,
	0x38, 0xe2, 0xa9, 0xbc, 0x30, 0x3b, 0x60, 0x76, 0x15, 0xf3, 0x97, 0x22, 0x98, 0xcd, 0x10, 0xc5,
	0x52, 0x74, 0xf1, 0x9b, 0x14, 0x33, 0x0e, 0x1f, 0x80, 0x1b, 0x2a, 0x01, 0x2f, 0x46, 0x11, 0x2e,
	0x6b, 0x55, 0x6d, 0x61, 0xd2, 0x9d, 0x52, 0xca, 0x57, 0x28, 0xc2, 0xe6, 0x6f, 0x0d, 0x54, 0xd7,
	0x53, 0x8e, 0x38, 0x3e, 0x0f, 0x62, 0x8a, 0x74, 0x0f, 0x18, 0xbd, 0x94, 0x71, 0x12, 0x61, 0xea,
	0x05, 0xbe, 0xe4, 0x00, 0xa5, 0x7a, 0xe1, 0xc3, 0xd7, 0x00, 0x90, 0x04, 0x53, 0x91, 0x79, 0x59,
	0xaf, 0x96, 0x16, 0x8c, 0x46, 0xd3, 0x1a, 0x57, 0x78, 0xeb, 0xbc, 0xcb, 0x0d, 0x45, 0x71, 0x4f,
	0x00, 0xe1, 0x43, 0x70, 0x2b, 0x41, 0x94, 0x07, 0x28, 0xf4, 0xf6, 0x50, 0x10, 0xa6, 0x14, 0x97,
	0x4b, 0x55, 0x6d, 0xe1, 0xba, 0x7b, 0x53, 0xaa, 0x57, 0x84, 0x36, 0x4b, 0x79, 0x80, 0xc2, 0xc0,
	0x47, 0x1c, 0x7b, 0x24, 0x0e, 0x0f, 0xcb, 0xff, 0xe7, 0x66, 0x53, 0x4a, 0xb9, 0x11, 0x87, 0x87,
	0xe6, 0x47, 0x1d, 0xdc, 0x1d, 0xe1, 0x19, 0x2e, 0x01, 0x23, 0x4d, 0x72, 0x44, 0xd6, 0xa5, 0x1c,
	0x61, 0x34, 0x2a, 0x2a, 0x1b, 0xd5, 0x26, 0x6b, 0x25, 0x6b, 0xe4, 0x3a, 0x62, 0xfb, 0x2e, 0x10,
	0xe6, 0x99, 0x0c, 0x37, 0xc0, 0x44, 0x8f, 0x62, 0xc4, 0x45, 0xb5, 0x8d, 0xc6, 0x62, 0x61, 0x15,
	0x86, 0xc3, 0x75, 0x41, 0x19, 0xd6, 0xfe, 0x73, 0x25, 0x26, 0x03, 0x0a, 0x7c, 0x59, 0xbf, 0x22,
	0x50, 0x60, 0x60, 0x19, 0x4c, 0x50, 0x1c, 0x91, 0x81, 0xa8, 0xe1, 0x64, 0x76, 0x23, 0xce, 0x6d,
	0x03, 0x4c, 0x0e, 0x8b, 0x6e, 0x7e, 0xd6, 0xc0, 0xfd, 0x11, 0x83, 0xc1, 0x12, 0x12, 0x33, 0x0c,
	0x57, 0xc0, 0xf4, 0x99, 0xce, 0x78, 0x98, 0x52, 0x42, 0x73, 0xb6, 0xd1, 0x80, 0x2a, 0x58, 0x9a,
	0xf4, 0xac, 0xad, 0x7c, 0x78, 0xdd, 0xdb, 0xa7, 0x7b, 0xb6, 0x9c, 0x99, 0xc3, 0x1d, 0x70, 0x8d,
	0x62, 0x96, 0x86, 0x5c, 0x4d, 0xcf, 0xf3, 0xf1, 0xd3, 0x53, 0x14, 0x9d, 0x9b, 0x83, 0x5c, 0x05,
	0x34, 0x97, 0xc1, 0xdc, 0x68, 0xd3, 0x7f, 0x7a, 0x29, 0x8d, 0xef, 0x25, 0x30, 0x73, 0x9e, 0xb0,
	0x25, 0xa2, 0x81, 0x5f, 0x35, 0x30, 0x7d, 0xe1, 0x6b, 0x84, 0xcf, 0xc6, 0x67, 0x32, 0xea, 0x19,
	0x57, 0x2e, 0xd7, 0x70, 0xb3, 0xf9, 0xf6, 0xdb, 0x8f, 0xf7, 0xfa, 0x63, 0xb8, 0x98, 0x2d, 0xb2,
	0xa3, 0x53, 0xe9, 0x35, 0xd5, 0xcb, 0x65, 0x76, 0x4d, 0x6d, 0xb6, 0x93, 0xdd, 0xb5, 0x6b, 0xc7,
	0xf0, 0xa7, 0x06, 0x66, 0x0a, 0xdb, 0x0f, 0xdb, 0x97, 0xef, 0x8e, 0x5a, 0x2a, 0x95, 0xce, 0x95,
	0x18, 0x62, 0xfe, 0xcc, 0x4e, 0x9e, 0x65, 0xd3, 0x7c, 0x92, 0x65, 0xf9, 0x37, 0xad, 0xa3, 0x13,
	0xeb, 0xaa, 0x59, 0x3b, 0xbe, 0x28, 0x49, 0x27, 0xca, 0xe1, 0x8e, 0x56, 0x6b, 0xbf, 0xd3, 0xc1,
	0x7c, 0x8f, 0x44, 0x63, 0xe3, 0x69, 0xcf, 0x15, 0xf6, 0x7f, 0x33, 0xdb, 0x0a, 0x9b, 0xda, 0xce,
	0x9a, 0x64, 0xf4, 0x49, 0x88, 0xe2, 0xbe, 0x45, 0x68, 0xdf, 0xee, 0xe3, 0x38, 0xdf, 0x19, 0xea,
	0xd7, 0x92, 0x04, 0xac, 0xf8, 0x87, 0xb6, 0xa4, 0x84, 0x0f, 0x7a, 0x69, 0xb5, 0xd5, 0xfa, 0xa4,
	0x57, 0x57, 0x05, 0xb0, 0xe5, 0x33, 0x4b, 0x88, 0x99, 0xb4, 0x5d, 0xb7, 0xa4, 0x63, 0xf6, 0x45,
	0x99, 0x74, 0x5b, 0x3e, 0xeb, 0x0e, 0x4d, 0xba, 0xdb, 0xf5, 0xae, 0x32, 0xf9, 0xa5, 0xcf, 0x0b,
	0xbd, 0xe3, 0xb4, 0x7c, 0xe6, 0x38, 0x43, 0x23, 0xc7, 0xd9, 0xae, 0x3b, 0x8e, 0x32, 0xdb, 0x9d,
	0xc8, 0xe3, 0x7c, 0xf4, 0x27, 0x00, 0x00, 0xff, 0xff, 0xba, 0x01, 0x70, 0x40, 0x77, 0x07, 0x00,
	0x00,
}
