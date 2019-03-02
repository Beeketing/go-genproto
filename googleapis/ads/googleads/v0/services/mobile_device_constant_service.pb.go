// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/mobile_device_constant_service.proto

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

// Request message for [MobileDeviceConstantService.GetMobileDeviceConstant][google.ads.googleads.v0.services.MobileDeviceConstantService.GetMobileDeviceConstant].
type GetMobileDeviceConstantRequest struct {
	// Resource name of the mobile device to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMobileDeviceConstantRequest) Reset()         { *m = GetMobileDeviceConstantRequest{} }
func (m *GetMobileDeviceConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetMobileDeviceConstantRequest) ProtoMessage()    {}
func (*GetMobileDeviceConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_mobile_device_constant_service_09a983f12dae70fa, []int{0}
}
func (m *GetMobileDeviceConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMobileDeviceConstantRequest.Unmarshal(m, b)
}
func (m *GetMobileDeviceConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMobileDeviceConstantRequest.Marshal(b, m, deterministic)
}
func (dst *GetMobileDeviceConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMobileDeviceConstantRequest.Merge(dst, src)
}
func (m *GetMobileDeviceConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetMobileDeviceConstantRequest.Size(m)
}
func (m *GetMobileDeviceConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMobileDeviceConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMobileDeviceConstantRequest proto.InternalMessageInfo

func (m *GetMobileDeviceConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetMobileDeviceConstantRequest)(nil), "google.ads.googleads.v0.services.GetMobileDeviceConstantRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MobileDeviceConstantServiceClient is the client API for MobileDeviceConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MobileDeviceConstantServiceClient interface {
	// Returns the requested mobile device constant in full detail.
	GetMobileDeviceConstant(ctx context.Context, in *GetMobileDeviceConstantRequest, opts ...grpc.CallOption) (*resources.MobileDeviceConstant, error)
}

type mobileDeviceConstantServiceClient struct {
	cc *grpc.ClientConn
}

func NewMobileDeviceConstantServiceClient(cc *grpc.ClientConn) MobileDeviceConstantServiceClient {
	return &mobileDeviceConstantServiceClient{cc}
}

func (c *mobileDeviceConstantServiceClient) GetMobileDeviceConstant(ctx context.Context, in *GetMobileDeviceConstantRequest, opts ...grpc.CallOption) (*resources.MobileDeviceConstant, error) {
	out := new(resources.MobileDeviceConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.MobileDeviceConstantService/GetMobileDeviceConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MobileDeviceConstantServiceServer is the server API for MobileDeviceConstantService service.
type MobileDeviceConstantServiceServer interface {
	// Returns the requested mobile device constant in full detail.
	GetMobileDeviceConstant(context.Context, *GetMobileDeviceConstantRequest) (*resources.MobileDeviceConstant, error)
}

func RegisterMobileDeviceConstantServiceServer(s *grpc.Server, srv MobileDeviceConstantServiceServer) {
	s.RegisterService(&_MobileDeviceConstantService_serviceDesc, srv)
}

func _MobileDeviceConstantService_GetMobileDeviceConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMobileDeviceConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileDeviceConstantServiceServer).GetMobileDeviceConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.MobileDeviceConstantService/GetMobileDeviceConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileDeviceConstantServiceServer).GetMobileDeviceConstant(ctx, req.(*GetMobileDeviceConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MobileDeviceConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.MobileDeviceConstantService",
	HandlerType: (*MobileDeviceConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMobileDeviceConstant",
			Handler:    _MobileDeviceConstantService_GetMobileDeviceConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/mobile_device_constant_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/mobile_device_constant_service.proto", fileDescriptor_mobile_device_constant_service_09a983f12dae70fa)
}

var fileDescriptor_mobile_device_constant_service_09a983f12dae70fa = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xbf, 0x4a, 0xc3, 0x40,
	0x18, 0x27, 0x11, 0x04, 0x83, 0x2e, 0x59, 0x94, 0x2a, 0x12, 0x6a, 0x07, 0x51, 0xbc, 0x0b, 0x76,
	0x10, 0x4e, 0x14, 0x53, 0x2d, 0x75, 0x51, 0x4a, 0x85, 0x0e, 0x12, 0x08, 0xd7, 0xe6, 0x08, 0x81,
	0xe6, 0xae, 0xe6, 0xbb, 0x76, 0x11, 0x17, 0x17, 0x1f, 0xc0, 0x37, 0x70, 0xf4, 0x4d, 0x74, 0xf5,
	0x15, 0x5c, 0xf4, 0x29, 0x24, 0xb9, 0x5c, 0x40, 0x48, 0xda, 0xed, 0x97, 0xbb, 0xdf, 0x9f, 0xef,
	0x7e, 0x5f, 0xac, 0x6e, 0x24, 0x44, 0x34, 0x61, 0x98, 0x86, 0x80, 0x15, 0xcc, 0xd0, 0xdc, 0xc5,
	0xc0, 0xd2, 0x79, 0x3c, 0x66, 0x80, 0x13, 0x31, 0x8a, 0x27, 0x2c, 0x08, 0x59, 0xf6, 0x19, 0x8c,
	0x05, 0x07, 0x49, 0xb9, 0x0c, 0x8a, 0x7b, 0x34, 0x4d, 0x85, 0x14, 0xb6, 0xa3, 0xb4, 0x88, 0x86,
	0x80, 0x4a, 0x1b, 0x34, 0x77, 0x91, 0xb6, 0x69, 0x9c, 0xd7, 0x05, 0xa5, 0x0c, 0xc4, 0x2c, 0xad,
	0x4f, 0x52, 0x09, 0x8d, 0x1d, 0xad, 0x9f, 0xc6, 0x98, 0x72, 0x2e, 0x24, 0x95, 0xb1, 0xe0, 0xa0,
	0x6e, 0x9b, 0x5d, 0x6b, 0xb7, 0xc7, 0xe4, 0x4d, 0x6e, 0x70, 0x95, 0xeb, 0x2f, 0x0b, 0xf9, 0x80,
	0x3d, 0xcc, 0x18, 0x48, 0x7b, 0xcf, 0xda, 0xd0, 0x49, 0x01, 0xa7, 0x09, 0xdb, 0x32, 0x1c, 0x63,
	0x7f, 0x6d, 0xb0, 0xae, 0x0f, 0x6f, 0x69, 0xc2, 0x8e, 0x7f, 0x0c, 0x6b, 0xbb, 0xca, 0xe4, 0x4e,
	0xbd, 0xc2, 0xfe, 0x30, 0xac, 0xcd, 0x9a, 0x1c, 0xfb, 0x02, 0x2d, 0xeb, 0x00, 0x2d, 0x1e, 0xb1,
	0x71, 0x52, 0xeb, 0x50, 0x76, 0x84, 0xaa, 0xf4, 0xcd, 0xf6, 0xf3, 0xd7, 0xf7, 0xab, 0x79, 0x64,
	0x1f, 0x66, 0x7d, 0x3e, 0xfe, 0x7b, 0xe6, 0x59, 0x52, 0x21, 0x00, 0x7c, 0xf0, 0xd4, 0x79, 0x31,
	0xad, 0xd6, 0x58, 0x24, 0x4b, 0xa7, 0xee, 0x38, 0x0b, 0x1a, 0xe9, 0x67, 0xed, 0xf7, 0x8d, 0xfb,
	0xeb, 0xc2, 0x25, 0x12, 0x13, 0xca, 0x23, 0x24, 0xd2, 0x08, 0x47, 0x8c, 0xe7, 0xbb, 0xd1, 0xdb,
	0x9e, 0xc6, 0x50, 0xff, 0x97, 0x9d, 0x6a, 0xf0, 0x66, 0xae, 0xf4, 0x3c, 0xef, 0xdd, 0x74, 0x7a,
	0xca, 0xd0, 0x0b, 0x01, 0x29, 0x98, 0xa1, 0xa1, 0x8b, 0x8a, 0x60, 0xf8, 0xd4, 0x14, 0xdf, 0x0b,
	0xc1, 0x2f, 0x29, 0xfe, 0xd0, 0xf5, 0x35, 0xe5, 0xd7, 0x6c, 0xa9, 0x73, 0x42, 0xbc, 0x10, 0x08,
	0x29, 0x49, 0x84, 0x0c, 0x5d, 0x42, 0x34, 0x6d, 0xb4, 0x9a, 0xcf, 0xd9, 0xfe, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0xcd, 0xc0, 0x26, 0x9e, 0x0c, 0x03, 0x00, 0x00,
}
