// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/feed_origin.proto

package enums // import "github.com/Beeketing/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Possible values for a feed origin.
type FeedOriginEnum_FeedOrigin int32

const (
	// Not specified.
	FeedOriginEnum_UNSPECIFIED FeedOriginEnum_FeedOrigin = 0
	// Used for return value only. Represents value unknown in this version.
	FeedOriginEnum_UNKNOWN FeedOriginEnum_FeedOrigin = 1
	// The FeedAttributes for this Feed are managed by the
	// user. Users can add FeedAttributes to this Feed.
	FeedOriginEnum_USER FeedOriginEnum_FeedOrigin = 2
	// The FeedAttributes for an GOOGLE Feed are created by Google. A feed of
	// this type is maintained by Google and will have the correct attributes
	// for the placeholder type of the feed.
	FeedOriginEnum_GOOGLE FeedOriginEnum_FeedOrigin = 3
)

var FeedOriginEnum_FeedOrigin_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "USER",
	3: "GOOGLE",
}
var FeedOriginEnum_FeedOrigin_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"USER":        2,
	"GOOGLE":      3,
}

func (x FeedOriginEnum_FeedOrigin) String() string {
	return proto.EnumName(FeedOriginEnum_FeedOrigin_name, int32(x))
}
func (FeedOriginEnum_FeedOrigin) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_feed_origin_ebd6dc41326e27b0, []int{0, 0}
}

// Container for enum describing possible values for a feed origin.
type FeedOriginEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedOriginEnum) Reset()         { *m = FeedOriginEnum{} }
func (m *FeedOriginEnum) String() string { return proto.CompactTextString(m) }
func (*FeedOriginEnum) ProtoMessage()    {}
func (*FeedOriginEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_origin_ebd6dc41326e27b0, []int{0}
}
func (m *FeedOriginEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedOriginEnum.Unmarshal(m, b)
}
func (m *FeedOriginEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedOriginEnum.Marshal(b, m, deterministic)
}
func (dst *FeedOriginEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedOriginEnum.Merge(dst, src)
}
func (m *FeedOriginEnum) XXX_Size() int {
	return xxx_messageInfo_FeedOriginEnum.Size(m)
}
func (m *FeedOriginEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedOriginEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedOriginEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FeedOriginEnum)(nil), "google.ads.googleads.v0.enums.FeedOriginEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.FeedOriginEnum_FeedOrigin", FeedOriginEnum_FeedOrigin_name, FeedOriginEnum_FeedOrigin_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/feed_origin.proto", fileDescriptor_feed_origin_ebd6dc41326e27b0)
}

var fileDescriptor_feed_origin_ebd6dc41326e27b0 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4f, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0x86, 0x32, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc, 0xd2,
	0xdc, 0x62, 0xfd, 0xb4, 0xd4, 0xd4, 0x94, 0xf8, 0xfc, 0xa2, 0xcc, 0xf4, 0xcc, 0x3c, 0xbd, 0x82,
	0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x2a, 0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x06, 0xbd,
	0x32, 0x03, 0x3d, 0xb0, 0x06, 0xa5, 0x20, 0x2e, 0x3e, 0xb7, 0xd4, 0xd4, 0x14, 0x7f, 0xb0, 0x16,
	0xd7, 0xbc, 0xd2, 0x5c, 0x25, 0x07, 0x2e, 0x2e, 0x84, 0x88, 0x10, 0x3f, 0x17, 0x77, 0xa8, 0x5f,
	0x70, 0x80, 0xab, 0xb3, 0xa7, 0x9b, 0xa7, 0xab, 0x8b, 0x00, 0x83, 0x10, 0x37, 0x17, 0x7b, 0xa8,
	0x9f, 0xb7, 0x9f, 0x7f, 0xb8, 0x9f, 0x00, 0xa3, 0x10, 0x07, 0x17, 0x4b, 0x68, 0xb0, 0x6b, 0x90,
	0x00, 0x93, 0x10, 0x17, 0x17, 0x9b, 0xbb, 0xbf, 0xbf, 0xbb, 0x8f, 0xab, 0x00, 0xb3, 0xd3, 0x13,
	0x46, 0x2e, 0xc5, 0xe4, 0xfc, 0x5c, 0x3d, 0xbc, 0x36, 0x3b, 0xf1, 0x23, 0x6c, 0x09, 0x00, 0xb9,
	0x34, 0x80, 0x31, 0xca, 0x09, 0xaa, 0x23, 0x3d, 0x3f, 0x27, 0x31, 0x2f, 0x5d, 0x2f, 0xbf, 0x28,
	0x5d, 0x3f, 0x3d, 0x35, 0x0f, 0xec, 0x0f, 0x98, 0x67, 0x0b, 0x32, 0x8b, 0x71, 0xf8, 0xdd, 0x1a,
	0x4c, 0x2e, 0x62, 0x62, 0x76, 0x77, 0x74, 0x5c, 0xc5, 0x24, 0xeb, 0x0e, 0x31, 0xca, 0x31, 0xa5,
	0x58, 0x0f, 0xc2, 0x04, 0xb1, 0xc2, 0x0c, 0xf4, 0x40, 0x7e, 0x2c, 0x3e, 0x05, 0x93, 0x8f, 0x71,
	0x4c, 0x29, 0x8e, 0x81, 0xcb, 0xc7, 0x84, 0x19, 0xc4, 0x80, 0xe5, 0x5f, 0x31, 0x29, 0x42, 0x04,
	0xad, 0xac, 0x1c, 0x53, 0x8a, 0xad, 0xac, 0xe0, 0x2a, 0xac, 0xac, 0xc2, 0x0c, 0xac, 0xac, 0xc0,
	0x6a, 0x92, 0xd8, 0xc0, 0x0e, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x4a, 0x39, 0xbc,
	0x93, 0x01, 0x00, 0x00,
}
