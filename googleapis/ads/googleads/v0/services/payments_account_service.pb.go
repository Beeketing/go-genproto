// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/payments_account_service.proto

package services // import "github.com/Beeketing/genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "github.com/Beeketing/genproto/googleapis/ads/googleads/v0/resources"
import _ "github.com/Beeketing/genproto/googleapis/api/annotations"

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

// Request message for fetching all accessible Payments accounts.
type ListPaymentsAccountsRequest struct {
	// The ID of the customer to apply the PaymentsAccount list operation to.
	CustomerId           string   `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPaymentsAccountsRequest) Reset()         { *m = ListPaymentsAccountsRequest{} }
func (m *ListPaymentsAccountsRequest) String() string { return proto.CompactTextString(m) }
func (*ListPaymentsAccountsRequest) ProtoMessage()    {}
func (*ListPaymentsAccountsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_payments_account_service_e8b479dad89ec6e6, []int{0}
}
func (m *ListPaymentsAccountsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPaymentsAccountsRequest.Unmarshal(m, b)
}
func (m *ListPaymentsAccountsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPaymentsAccountsRequest.Marshal(b, m, deterministic)
}
func (dst *ListPaymentsAccountsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPaymentsAccountsRequest.Merge(dst, src)
}
func (m *ListPaymentsAccountsRequest) XXX_Size() int {
	return xxx_messageInfo_ListPaymentsAccountsRequest.Size(m)
}
func (m *ListPaymentsAccountsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPaymentsAccountsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPaymentsAccountsRequest proto.InternalMessageInfo

func (m *ListPaymentsAccountsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

// Response message for [PaymentsAccountService.ListPaymentsAccounts][google.ads.googleads.v0.services.PaymentsAccountService.ListPaymentsAccounts].
type ListPaymentsAccountsResponse struct {
	// The list of accessible Payments accounts.
	PaymentsAccounts     []*resources.PaymentsAccount `protobuf:"bytes,1,rep,name=payments_accounts,json=paymentsAccounts,proto3" json:"payments_accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ListPaymentsAccountsResponse) Reset()         { *m = ListPaymentsAccountsResponse{} }
func (m *ListPaymentsAccountsResponse) String() string { return proto.CompactTextString(m) }
func (*ListPaymentsAccountsResponse) ProtoMessage()    {}
func (*ListPaymentsAccountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_payments_account_service_e8b479dad89ec6e6, []int{1}
}
func (m *ListPaymentsAccountsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPaymentsAccountsResponse.Unmarshal(m, b)
}
func (m *ListPaymentsAccountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPaymentsAccountsResponse.Marshal(b, m, deterministic)
}
func (dst *ListPaymentsAccountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPaymentsAccountsResponse.Merge(dst, src)
}
func (m *ListPaymentsAccountsResponse) XXX_Size() int {
	return xxx_messageInfo_ListPaymentsAccountsResponse.Size(m)
}
func (m *ListPaymentsAccountsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPaymentsAccountsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPaymentsAccountsResponse proto.InternalMessageInfo

func (m *ListPaymentsAccountsResponse) GetPaymentsAccounts() []*resources.PaymentsAccount {
	if m != nil {
		return m.PaymentsAccounts
	}
	return nil
}

func init() {
	proto.RegisterType((*ListPaymentsAccountsRequest)(nil), "google.ads.googleads.v0.services.ListPaymentsAccountsRequest")
	proto.RegisterType((*ListPaymentsAccountsResponse)(nil), "google.ads.googleads.v0.services.ListPaymentsAccountsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PaymentsAccountServiceClient is the client API for PaymentsAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaymentsAccountServiceClient interface {
	// Returns all Payments accounts associated with all managers
	// between the login customer ID and specified serving customer in the
	// hierarchy, inclusive.
	ListPaymentsAccounts(ctx context.Context, in *ListPaymentsAccountsRequest, opts ...grpc.CallOption) (*ListPaymentsAccountsResponse, error)
}

type paymentsAccountServiceClient struct {
	cc *grpc.ClientConn
}

func NewPaymentsAccountServiceClient(cc *grpc.ClientConn) PaymentsAccountServiceClient {
	return &paymentsAccountServiceClient{cc}
}

func (c *paymentsAccountServiceClient) ListPaymentsAccounts(ctx context.Context, in *ListPaymentsAccountsRequest, opts ...grpc.CallOption) (*ListPaymentsAccountsResponse, error) {
	out := new(ListPaymentsAccountsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.PaymentsAccountService/ListPaymentsAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentsAccountServiceServer is the server API for PaymentsAccountService service.
type PaymentsAccountServiceServer interface {
	// Returns all Payments accounts associated with all managers
	// between the login customer ID and specified serving customer in the
	// hierarchy, inclusive.
	ListPaymentsAccounts(context.Context, *ListPaymentsAccountsRequest) (*ListPaymentsAccountsResponse, error)
}

func RegisterPaymentsAccountServiceServer(s *grpc.Server, srv PaymentsAccountServiceServer) {
	s.RegisterService(&_PaymentsAccountService_serviceDesc, srv)
}

func _PaymentsAccountService_ListPaymentsAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPaymentsAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsAccountServiceServer).ListPaymentsAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.PaymentsAccountService/ListPaymentsAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsAccountServiceServer).ListPaymentsAccounts(ctx, req.(*ListPaymentsAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaymentsAccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.PaymentsAccountService",
	HandlerType: (*PaymentsAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPaymentsAccounts",
			Handler:    _PaymentsAccountService_ListPaymentsAccounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/payments_account_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/payments_account_service.proto", fileDescriptor_payments_account_service_e8b479dad89ec6e6)
}

var fileDescriptor_payments_account_service_e8b479dad89ec6e6 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcf, 0x6a, 0xdb, 0x30,
	0x18, 0xc7, 0x0e, 0x0c, 0xa6, 0x5c, 0x36, 0x33, 0x46, 0x48, 0x02, 0x33, 0x21, 0x87, 0xb0, 0x83,
	0x64, 0x32, 0x18, 0x43, 0x23, 0x19, 0xce, 0x25, 0x1b, 0xec, 0x10, 0x32, 0xc8, 0x61, 0x18, 0x8c,
	0x6b, 0x0b, 0x63, 0x48, 0x24, 0xd7, 0x9f, 0x1c, 0x28, 0xa5, 0x14, 0xf2, 0x0a, 0x7d, 0x83, 0x1e,
	0xfb, 0x28, 0x85, 0x9e, 0xfa, 0x0a, 0x3d, 0x94, 0x3e, 0x45, 0x71, 0x64, 0x99, 0x62, 0x9c, 0x06,
	0x7a, 0xfb, 0x21, 0xfd, 0xfe, 0xe8, 0xfb, 0xe9, 0x43, 0xbf, 0x62, 0x21, 0xe2, 0x35, 0x23, 0x41,
	0x04, 0x44, 0xc1, 0x02, 0x6d, 0x1d, 0x02, 0x2c, 0xdb, 0x26, 0x21, 0x03, 0x92, 0x06, 0x67, 0x1b,
	0xc6, 0x25, 0xf8, 0x41, 0x18, 0x8a, 0x9c, 0x4b, 0xbf, 0xbc, 0xc1, 0x69, 0x26, 0xa4, 0xb0, 0x6c,
	0xa5, 0xc2, 0x41, 0x04, 0xb8, 0x32, 0xc0, 0x5b, 0x07, 0x6b, 0x83, 0xee, 0x8f, 0x43, 0x11, 0x19,
	0x03, 0x91, 0x67, 0x4d, 0x19, 0xca, 0xbb, 0xdb, 0xd7, 0xca, 0x34, 0x21, 0x01, 0xe7, 0x42, 0x06,
	0x32, 0x11, 0x1c, 0xd4, 0xed, 0x60, 0x8a, 0x7a, 0x7f, 0x13, 0x90, 0x8b, 0x52, 0xeb, 0x2a, 0x29,
	0x2c, 0xd9, 0x69, 0xce, 0x40, 0x5a, 0x5f, 0x50, 0x3b, 0xcc, 0x41, 0x8a, 0x0d, 0xcb, 0xfc, 0x24,
	0xea, 0x18, 0xb6, 0x31, 0x7a, 0xbf, 0x44, 0xfa, 0xe8, 0x4f, 0x34, 0xb8, 0x44, 0xfd, 0x66, 0x3d,
	0xa4, 0x82, 0x03, 0xb3, 0x7c, 0xf4, 0xb1, 0xfe, 0x2e, 0xe8, 0x18, 0x76, 0x6b, 0xd4, 0x1e, 0x8f,
	0xf1, 0xa1, 0xa9, 0xab, 0x99, 0x70, 0xcd, 0x77, 0xf9, 0x21, 0xad, 0x05, 0x8d, 0x1f, 0x0d, 0xf4,
	0xb9, 0xc6, 0xfa, 0xa7, 0x4a, 0xb3, 0xee, 0x0c, 0xf4, 0xa9, 0xe9, 0x71, 0xd6, 0x04, 0x1f, 0xeb,
	0x1b, 0xbf, 0x52, 0x4a, 0x77, 0xfa, 0x56, 0xb9, 0xea, 0x64, 0xf0, 0x7d, 0x77, 0xff, 0x70, 0x65,
	0x3a, 0x16, 0x2e, 0xfe, 0x4f, 0x77, 0x09, 0xe4, 0xfc, 0x45, 0xd3, 0x93, 0xaf, 0x17, 0xa4, 0x3e,
	0xea, 0x6c, 0x67, 0xa2, 0x61, 0x28, 0x36, 0x47, 0xd3, 0x67, 0xbd, 0xe6, 0x42, 0x16, 0xc5, 0x8f,
	0x2f, 0x8c, 0xff, 0xbf, 0x4b, 0x83, 0x58, 0xac, 0x03, 0x1e, 0x63, 0x91, 0xc5, 0x24, 0x66, 0x7c,
	0xbf, 0x0f, 0x7a, 0xb7, 0xd2, 0x04, 0x0e, 0x6f, 0xf3, 0x4f, 0x0d, 0xae, 0xcd, 0xd6, 0xdc, 0x75,
	0x6f, 0x4c, 0x7b, 0xae, 0x0c, 0xdd, 0x08, 0xb0, 0x82, 0x05, 0x5a, 0x39, 0xb8, 0x0c, 0x86, 0x5b,
	0x4d, 0xf1, 0xdc, 0x08, 0xbc, 0x8a, 0xe2, 0xad, 0x1c, 0x4f, 0x53, 0x9e, 0xcc, 0xa1, 0x3a, 0xa7,
	0xd4, 0x8d, 0x80, 0xd2, 0x8a, 0x44, 0xe9, 0xca, 0xa1, 0x54, 0xd3, 0x4e, 0xde, 0xed, 0xdf, 0xf9,
	0xed, 0x39, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xc2, 0x8a, 0xd2, 0x74, 0x03, 0x00, 0x00,
}
