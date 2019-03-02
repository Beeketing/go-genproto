// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/resultstore/v2/resultstore_file_download.proto

package resultstore // import "github.com/Beeketing/genproto/googleapis/devtools/resultstore/v2"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
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

// Request object for GetFile
type GetFileRequest struct {
	// This corresponds to the uri field in the File message.
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// The offset for the first byte to return in the read, relative to the start
	// of the resource.
	//
	// A `read_offset` that is negative or greater than the size of the resource
	// will cause an `OUT_OF_RANGE` error.
	ReadOffset int64 `protobuf:"varint,2,opt,name=read_offset,json=readOffset,proto3" json:"read_offset,omitempty"`
	// The maximum number of `data` bytes the server is allowed to return in the
	// sum of all `ReadResponse` messages. A `read_limit` of zero indicates that
	// there is no limit, and a negative `read_limit` will cause an error.
	//
	// If the stream returns fewer bytes than allowed by the `read_limit` and no
	// error occurred, the stream includes all data from the `read_offset` to the
	// end of the resource.
	ReadLimit            int64    `protobuf:"varint,3,opt,name=read_limit,json=readLimit,proto3" json:"read_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFileRequest) Reset()         { *m = GetFileRequest{} }
func (m *GetFileRequest) String() string { return proto.CompactTextString(m) }
func (*GetFileRequest) ProtoMessage()    {}
func (*GetFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_resultstore_file_download_258c94f8298cac64, []int{0}
}
func (m *GetFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFileRequest.Unmarshal(m, b)
}
func (m *GetFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFileRequest.Marshal(b, m, deterministic)
}
func (dst *GetFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFileRequest.Merge(dst, src)
}
func (m *GetFileRequest) XXX_Size() int {
	return xxx_messageInfo_GetFileRequest.Size(m)
}
func (m *GetFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFileRequest proto.InternalMessageInfo

func (m *GetFileRequest) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *GetFileRequest) GetReadOffset() int64 {
	if m != nil {
		return m.ReadOffset
	}
	return 0
}

func (m *GetFileRequest) GetReadLimit() int64 {
	if m != nil {
		return m.ReadLimit
	}
	return 0
}

// Response object for GetFile
type GetFileResponse struct {
	// The file data.
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFileResponse) Reset()         { *m = GetFileResponse{} }
func (m *GetFileResponse) String() string { return proto.CompactTextString(m) }
func (*GetFileResponse) ProtoMessage()    {}
func (*GetFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_resultstore_file_download_258c94f8298cac64, []int{1}
}
func (m *GetFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFileResponse.Unmarshal(m, b)
}
func (m *GetFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFileResponse.Marshal(b, m, deterministic)
}
func (dst *GetFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFileResponse.Merge(dst, src)
}
func (m *GetFileResponse) XXX_Size() int {
	return xxx_messageInfo_GetFileResponse.Size(m)
}
func (m *GetFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFileResponse proto.InternalMessageInfo

func (m *GetFileResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Request object for GetFileTail
type GetFileTailRequest struct {
	// This corresponds to the uri field in the File message.
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// The offset for the first byte to return in the read, relative to the end
	// of the resource.
	//
	// A `read_offset` that is negative or greater than the size of the resource
	// will cause an `OUT_OF_RANGE` error.
	ReadOffset int64 `protobuf:"varint,2,opt,name=read_offset,json=readOffset,proto3" json:"read_offset,omitempty"`
	// The maximum number of `data` bytes the server is allowed to return. The
	// server will return bytes starting from the tail of the file.
	//
	// A `read_limit` of zero indicates that there is no limit, and a negative
	// `read_limit` will cause an error.
	ReadLimit            int64    `protobuf:"varint,3,opt,name=read_limit,json=readLimit,proto3" json:"read_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFileTailRequest) Reset()         { *m = GetFileTailRequest{} }
func (m *GetFileTailRequest) String() string { return proto.CompactTextString(m) }
func (*GetFileTailRequest) ProtoMessage()    {}
func (*GetFileTailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_resultstore_file_download_258c94f8298cac64, []int{2}
}
func (m *GetFileTailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFileTailRequest.Unmarshal(m, b)
}
func (m *GetFileTailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFileTailRequest.Marshal(b, m, deterministic)
}
func (dst *GetFileTailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFileTailRequest.Merge(dst, src)
}
func (m *GetFileTailRequest) XXX_Size() int {
	return xxx_messageInfo_GetFileTailRequest.Size(m)
}
func (m *GetFileTailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFileTailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFileTailRequest proto.InternalMessageInfo

func (m *GetFileTailRequest) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *GetFileTailRequest) GetReadOffset() int64 {
	if m != nil {
		return m.ReadOffset
	}
	return 0
}

func (m *GetFileTailRequest) GetReadLimit() int64 {
	if m != nil {
		return m.ReadLimit
	}
	return 0
}

// Response object for GetFileTail
type GetFileTailResponse struct {
	// The file data, encoded with UTF-8.
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFileTailResponse) Reset()         { *m = GetFileTailResponse{} }
func (m *GetFileTailResponse) String() string { return proto.CompactTextString(m) }
func (*GetFileTailResponse) ProtoMessage()    {}
func (*GetFileTailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_resultstore_file_download_258c94f8298cac64, []int{3}
}
func (m *GetFileTailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFileTailResponse.Unmarshal(m, b)
}
func (m *GetFileTailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFileTailResponse.Marshal(b, m, deterministic)
}
func (dst *GetFileTailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFileTailResponse.Merge(dst, src)
}
func (m *GetFileTailResponse) XXX_Size() int {
	return xxx_messageInfo_GetFileTailResponse.Size(m)
}
func (m *GetFileTailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFileTailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFileTailResponse proto.InternalMessageInfo

func (m *GetFileTailResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetFileRequest)(nil), "google.devtools.resultstore.v2.GetFileRequest")
	proto.RegisterType((*GetFileResponse)(nil), "google.devtools.resultstore.v2.GetFileResponse")
	proto.RegisterType((*GetFileTailRequest)(nil), "google.devtools.resultstore.v2.GetFileTailRequest")
	proto.RegisterType((*GetFileTailResponse)(nil), "google.devtools.resultstore.v2.GetFileTailResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResultStoreFileDownloadClient is the client API for ResultStoreFileDownload service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResultStoreFileDownloadClient interface {
	// Retrieves the File with the given uri.
	// returns a stream of bytes to be stitched together in order.
	//
	// An error will be reported in the following cases:
	// - If the File is not found.
	// - If the given File uri is badly formatted.
	GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (ResultStoreFileDownload_GetFileClient, error)
	// Retrieves the tail of a File with the given uri.
	//
	// An error will be reported in the following cases:
	// - If the File is not found.
	// - If the given File uri is badly formatted.
	GetFileTail(ctx context.Context, in *GetFileTailRequest, opts ...grpc.CallOption) (*GetFileTailResponse, error)
}

type resultStoreFileDownloadClient struct {
	cc *grpc.ClientConn
}

func NewResultStoreFileDownloadClient(cc *grpc.ClientConn) ResultStoreFileDownloadClient {
	return &resultStoreFileDownloadClient{cc}
}

func (c *resultStoreFileDownloadClient) GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (ResultStoreFileDownload_GetFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ResultStoreFileDownload_serviceDesc.Streams[0], "/google.devtools.resultstore.v2.ResultStoreFileDownload/GetFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &resultStoreFileDownloadGetFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ResultStoreFileDownload_GetFileClient interface {
	Recv() (*GetFileResponse, error)
	grpc.ClientStream
}

type resultStoreFileDownloadGetFileClient struct {
	grpc.ClientStream
}

func (x *resultStoreFileDownloadGetFileClient) Recv() (*GetFileResponse, error) {
	m := new(GetFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *resultStoreFileDownloadClient) GetFileTail(ctx context.Context, in *GetFileTailRequest, opts ...grpc.CallOption) (*GetFileTailResponse, error) {
	out := new(GetFileTailResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.resultstore.v2.ResultStoreFileDownload/GetFileTail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResultStoreFileDownloadServer is the server API for ResultStoreFileDownload service.
type ResultStoreFileDownloadServer interface {
	// Retrieves the File with the given uri.
	// returns a stream of bytes to be stitched together in order.
	//
	// An error will be reported in the following cases:
	// - If the File is not found.
	// - If the given File uri is badly formatted.
	GetFile(*GetFileRequest, ResultStoreFileDownload_GetFileServer) error
	// Retrieves the tail of a File with the given uri.
	//
	// An error will be reported in the following cases:
	// - If the File is not found.
	// - If the given File uri is badly formatted.
	GetFileTail(context.Context, *GetFileTailRequest) (*GetFileTailResponse, error)
}

func RegisterResultStoreFileDownloadServer(s *grpc.Server, srv ResultStoreFileDownloadServer) {
	s.RegisterService(&_ResultStoreFileDownload_serviceDesc, srv)
}

func _ResultStoreFileDownload_GetFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ResultStoreFileDownloadServer).GetFile(m, &resultStoreFileDownloadGetFileServer{stream})
}

type ResultStoreFileDownload_GetFileServer interface {
	Send(*GetFileResponse) error
	grpc.ServerStream
}

type resultStoreFileDownloadGetFileServer struct {
	grpc.ServerStream
}

func (x *resultStoreFileDownloadGetFileServer) Send(m *GetFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ResultStoreFileDownload_GetFileTail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileTailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultStoreFileDownloadServer).GetFileTail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.resultstore.v2.ResultStoreFileDownload/GetFileTail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultStoreFileDownloadServer).GetFileTail(ctx, req.(*GetFileTailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResultStoreFileDownload_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.resultstore.v2.ResultStoreFileDownload",
	HandlerType: (*ResultStoreFileDownloadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFileTail",
			Handler:    _ResultStoreFileDownload_GetFileTail_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetFile",
			Handler:       _ResultStoreFileDownload_GetFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "google/devtools/resultstore/v2/resultstore_file_download.proto",
}

func init() {
	proto.RegisterFile("google/devtools/resultstore/v2/resultstore_file_download.proto", fileDescriptor_resultstore_file_download_258c94f8298cac64)
}

var fileDescriptor_resultstore_file_download_258c94f8298cac64 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xc1, 0x4a, 0xe3, 0x40,
	0x18, 0x26, 0xed, 0xb2, 0x4b, 0xa7, 0xcb, 0x6e, 0x99, 0x65, 0x69, 0x28, 0xdb, 0xb5, 0x04, 0x84,
	0xda, 0xc3, 0x8c, 0xa4, 0x47, 0xd1, 0x83, 0x88, 0x22, 0x08, 0x4a, 0xf4, 0xe4, 0x25, 0x4c, 0xcd,
	0x24, 0x0c, 0x4c, 0xf3, 0xa7, 0x99, 0x49, 0x3d, 0x48, 0x2f, 0x1e, 0x7c, 0x01, 0xf1, 0x65, 0x7c,
	0x0d, 0x5f, 0xc1, 0x07, 0x91, 0x99, 0x44, 0x49, 0x0f, 0xd5, 0x7a, 0xf0, 0x36, 0xff, 0x37, 0xff,
	0xf7, 0xfd, 0x5f, 0xbe, 0xfc, 0x83, 0xf6, 0x12, 0x80, 0x44, 0x72, 0x1a, 0xf1, 0xb9, 0x06, 0x90,
	0x8a, 0xe6, 0x5c, 0x15, 0x52, 0x2b, 0x0d, 0x39, 0xa7, 0x73, 0xbf, 0x5e, 0x86, 0xb1, 0x90, 0x3c,
	0x8c, 0xe0, 0x3a, 0x95, 0xc0, 0x22, 0x92, 0xe5, 0xa0, 0x01, 0xff, 0x2f, 0xf9, 0xe4, 0x95, 0x4f,
	0x6a, 0x04, 0x32, 0xf7, 0x7b, 0xff, 0x2a, 0x7d, 0x96, 0x09, 0xca, 0xd2, 0x14, 0x34, 0xd3, 0x02,
	0x52, 0x55, 0xb2, 0xbd, 0x09, 0xfa, 0x75, 0xc4, 0xf5, 0xa1, 0x90, 0x3c, 0xe0, 0xb3, 0x82, 0x2b,
	0x8d, 0x3b, 0xa8, 0x59, 0xe4, 0xc2, 0x75, 0x06, 0xce, 0xb0, 0x15, 0x98, 0x23, 0xde, 0x40, 0xed,
	0x9c, 0xb3, 0x28, 0x84, 0x38, 0x56, 0x5c, 0xbb, 0x8d, 0x81, 0x33, 0x6c, 0x06, 0xc8, 0x40, 0xa7,
	0x16, 0xc1, 0x7d, 0x64, 0xab, 0x50, 0x8a, 0xa9, 0xd0, 0x6e, 0xd3, 0xde, 0xb7, 0x0c, 0x72, 0x62,
	0x00, 0x6f, 0x13, 0xfd, 0x7e, 0x9b, 0xa1, 0x32, 0x48, 0x15, 0xc7, 0x18, 0x7d, 0x8b, 0x98, 0x66,
	0x76, 0xca, 0xcf, 0xc0, 0x9e, 0xbd, 0x18, 0xe1, 0xaa, 0xed, 0x82, 0x09, 0xf9, 0x75, 0x76, 0xb6,
	0xd0, 0x9f, 0xa5, 0x39, 0xab, 0x2d, 0xf9, 0x8f, 0x0d, 0xd4, 0x0d, 0x6c, 0x9c, 0xe7, 0x26, 0x4e,
	0xc3, 0x39, 0xa8, 0xd2, 0xc7, 0x77, 0x0e, 0xfa, 0x51, 0xe9, 0x60, 0x42, 0xde, 0xff, 0x09, 0x64,
	0x39, 0xe3, 0x1e, 0x5d, 0xbb, 0xbf, 0x34, 0xe7, 0xb9, 0xb7, 0x4f, 0xcf, 0xf7, 0x0d, 0x8c, 0x3b,
	0x66, 0x23, 0x6e, 0x8a, 0x5c, 0xec, 0x9a, 0x55, 0xa0, 0xa3, 0xc5, 0xb6, 0x83, 0x1f, 0x1c, 0xd4,
	0xae, 0x7d, 0x10, 0xf6, 0xd7, 0x14, 0xaf, 0xa5, 0xdc, 0x1b, 0x7f, 0x8a, 0x53, 0x99, 0xea, 0x5b,
	0x53, 0x5d, 0xfc, 0x77, 0xd9, 0x94, 0x66, 0x42, 0xd2, 0xd1, 0x62, 0x7f, 0x86, 0xbc, 0x2b, 0x98,
	0x7e, 0x20, 0x7c, 0xe6, 0x5c, 0x1e, 0x57, 0x1d, 0x09, 0x48, 0x96, 0x26, 0x04, 0xf2, 0x84, 0x26,
	0x3c, 0xb5, 0xeb, 0x49, 0xcb, 0x2b, 0x96, 0x09, 0xb5, 0xea, 0x7d, 0xec, 0xd4, 0xca, 0xc9, 0x77,
	0xcb, 0x1a, 0xbf, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x3d, 0x87, 0x5d, 0x54, 0x03, 0x00, 0x00,
}
