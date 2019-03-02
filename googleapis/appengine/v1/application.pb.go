// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/appengine/v1/application.proto

package appengine // import "github.com/Beeketing/go-genproto/googleapis/appengine/v1"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/Beeketing/protobuf/ptypes/duration"
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

// An Application resource contains the top-level configuration of an App
// Engine application.
type Application struct {
	// Full path to the Application resource in the API.
	// Example: `apps/myapp`.
	//
	// @OutputOnly
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Identifier of the Application resource. This identifier is equivalent
	// to the project ID of the Google Cloud Platform project where you want to
	// deploy your application.
	// Example: `myapp`.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// HTTP path dispatch rules for requests to the application that do not
	// explicitly target a service or version. Rules are order-dependent.
	//
	// @OutputOnly
	DispatchRules []*UrlDispatchRule `protobuf:"bytes,3,rep,name=dispatch_rules,json=dispatchRules,proto3" json:"dispatch_rules,omitempty"`
	// Google Apps authentication domain that controls which users can access
	// this application.
	//
	// Defaults to open access for any Google Account.
	AuthDomain string `protobuf:"bytes,6,opt,name=auth_domain,json=authDomain,proto3" json:"auth_domain,omitempty"`
	// Location from which this application will be run. Application instances
	// will run out of data centers in the chosen location, which is also where
	// all of the application's end user content is stored.
	//
	// Defaults to `us-central`.
	//
	// Options are:
	//
	// `us-central` - Central US
	//
	// `europe-west` - Western Europe
	//
	// `us-east1` - Eastern US
	LocationId string `protobuf:"bytes,7,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	// Google Cloud Storage bucket that can be used for storing files
	// associated with this application. This bucket is associated with the
	// application and can be used by the gcloud deployment commands.
	//
	// @OutputOnly
	CodeBucket string `protobuf:"bytes,8,opt,name=code_bucket,json=codeBucket,proto3" json:"code_bucket,omitempty"`
	// Cookie expiration policy for this application.
	//
	// @OutputOnly
	DefaultCookieExpiration *duration.Duration `protobuf:"bytes,9,opt,name=default_cookie_expiration,json=defaultCookieExpiration,proto3" json:"default_cookie_expiration,omitempty"`
	// Hostname used to reach this application, as resolved by App Engine.
	//
	// @OutputOnly
	DefaultHostname string `protobuf:"bytes,11,opt,name=default_hostname,json=defaultHostname,proto3" json:"default_hostname,omitempty"`
	// Google Cloud Storage bucket that can be used by this application to store
	// content.
	//
	// @OutputOnly
	DefaultBucket        string   `protobuf:"bytes,12,opt,name=default_bucket,json=defaultBucket,proto3" json:"default_bucket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Application) Reset()         { *m = Application{} }
func (m *Application) String() string { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()    {}
func (*Application) Descriptor() ([]byte, []int) {
	return fileDescriptor_application_d5f9ae7b7e94b936, []int{0}
}
func (m *Application) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Application.Unmarshal(m, b)
}
func (m *Application) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Application.Marshal(b, m, deterministic)
}
func (dst *Application) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Application.Merge(dst, src)
}
func (m *Application) XXX_Size() int {
	return xxx_messageInfo_Application.Size(m)
}
func (m *Application) XXX_DiscardUnknown() {
	xxx_messageInfo_Application.DiscardUnknown(m)
}

var xxx_messageInfo_Application proto.InternalMessageInfo

func (m *Application) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Application) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Application) GetDispatchRules() []*UrlDispatchRule {
	if m != nil {
		return m.DispatchRules
	}
	return nil
}

func (m *Application) GetAuthDomain() string {
	if m != nil {
		return m.AuthDomain
	}
	return ""
}

func (m *Application) GetLocationId() string {
	if m != nil {
		return m.LocationId
	}
	return ""
}

func (m *Application) GetCodeBucket() string {
	if m != nil {
		return m.CodeBucket
	}
	return ""
}

func (m *Application) GetDefaultCookieExpiration() *duration.Duration {
	if m != nil {
		return m.DefaultCookieExpiration
	}
	return nil
}

func (m *Application) GetDefaultHostname() string {
	if m != nil {
		return m.DefaultHostname
	}
	return ""
}

func (m *Application) GetDefaultBucket() string {
	if m != nil {
		return m.DefaultBucket
	}
	return ""
}

// Rules to match an HTTP request and dispatch that request to a service.
type UrlDispatchRule struct {
	// Domain name to match against. The wildcard "`*`" is supported if
	// specified before a period: "`*.`".
	//
	// Defaults to matching all domains: "`*`".
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// Pathname within the host. Must start with a "`/`". A
	// single "`*`" can be included at the end of the path. The sum
	// of the lengths of the domain and path may not exceed 100
	// characters.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// Resource ID of a service in this application that should
	// serve the matched request. The service must already
	// exist. Example: `default`.
	Service              string   `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UrlDispatchRule) Reset()         { *m = UrlDispatchRule{} }
func (m *UrlDispatchRule) String() string { return proto.CompactTextString(m) }
func (*UrlDispatchRule) ProtoMessage()    {}
func (*UrlDispatchRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_application_d5f9ae7b7e94b936, []int{1}
}
func (m *UrlDispatchRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UrlDispatchRule.Unmarshal(m, b)
}
func (m *UrlDispatchRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UrlDispatchRule.Marshal(b, m, deterministic)
}
func (dst *UrlDispatchRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UrlDispatchRule.Merge(dst, src)
}
func (m *UrlDispatchRule) XXX_Size() int {
	return xxx_messageInfo_UrlDispatchRule.Size(m)
}
func (m *UrlDispatchRule) XXX_DiscardUnknown() {
	xxx_messageInfo_UrlDispatchRule.DiscardUnknown(m)
}

var xxx_messageInfo_UrlDispatchRule proto.InternalMessageInfo

func (m *UrlDispatchRule) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *UrlDispatchRule) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *UrlDispatchRule) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func init() {
	proto.RegisterType((*Application)(nil), "google.appengine.v1.Application")
	proto.RegisterType((*UrlDispatchRule)(nil), "google.appengine.v1.UrlDispatchRule")
}

func init() {
	proto.RegisterFile("google/appengine/v1/application.proto", fileDescriptor_application_d5f9ae7b7e94b936)
}

var fileDescriptor_application_d5f9ae7b7e94b936 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x5f, 0x6b, 0xdb, 0x30,
	0x14, 0xc5, 0x71, 0x3c, 0x92, 0x45, 0x5e, 0xfe, 0xa0, 0xc1, 0xa2, 0x84, 0xb1, 0x85, 0xb0, 0x40,
	0xf6, 0x62, 0x93, 0xec, 0x71, 0x7b, 0x59, 0x9a, 0x42, 0x4b, 0x5f, 0x82, 0x21, 0x14, 0xfa, 0x62,
	0x14, 0x4b, 0xb1, 0x45, 0x14, 0xc9, 0xd8, 0x72, 0xe8, 0x67, 0xe8, 0xa7, 0x2e, 0x96, 0x64, 0x37,
	0x2d, 0x79, 0xd3, 0x3d, 0xfa, 0x1d, 0xdd, 0xeb, 0x73, 0x0d, 0xe6, 0x89, 0x94, 0x09, 0xa7, 0x01,
	0xce, 0x32, 0x2a, 0x12, 0x26, 0x68, 0x70, 0x5e, 0x56, 0x05, 0x67, 0x31, 0x56, 0x4c, 0x0a, 0x3f,
	0xcb, 0xa5, 0x92, 0xf0, 0xab, 0xc1, 0xfc, 0x06, 0xf3, 0xcf, 0xcb, 0xc9, 0xf7, 0xc6, 0xcb, 0x02,
	0x2c, 0x84, 0x54, 0xda, 0x51, 0x18, 0xcb, 0xe4, 0x87, 0xbd, 0xd5, 0xd5, 0xbe, 0x3c, 0x04, 0xa4,
	0xcc, 0x2f, 0x9e, 0x9c, 0xbd, 0xb8, 0xc0, 0xfb, 0xff, 0xd6, 0x08, 0x42, 0xf0, 0x49, 0xe0, 0x13,
	0x45, 0xce, 0xd4, 0x59, 0x74, 0x43, 0x7d, 0x86, 0x7d, 0xd0, 0x62, 0x04, 0xb5, 0xb4, 0xd2, 0x62,
	0x04, 0x3e, 0x80, 0x3e, 0x61, 0x45, 0x86, 0x55, 0x9c, 0x46, 0x79, 0xc9, 0x69, 0x81, 0xdc, 0xa9,
	0xbb, 0xf0, 0x56, 0xbf, 0xfc, 0x2b, 0xf3, 0xf9, 0xbb, 0x9c, 0x6f, 0x2c, 0x1d, 0x96, 0x9c, 0x86,
	0x3d, 0x72, 0x51, 0x15, 0xf0, 0x27, 0xf0, 0x70, 0xa9, 0xd2, 0x88, 0xc8, 0x13, 0x66, 0x02, 0xb5,
	0x75, 0x17, 0x50, 0x49, 0x1b, 0xad, 0x54, 0x00, 0x97, 0x66, 0xba, 0x88, 0x11, 0xd4, 0x31, 0x40,
	0x2d, 0xdd, 0x93, 0x0a, 0x88, 0x25, 0xa1, 0xd1, 0xbe, 0x8c, 0x8f, 0x54, 0xa1, 0xcf, 0x06, 0xa8,
	0xa4, 0xb5, 0x56, 0xe0, 0x0e, 0x8c, 0x09, 0x3d, 0xe0, 0x92, 0xab, 0x28, 0x96, 0xf2, 0xc8, 0x68,
	0x44, 0x9f, 0x33, 0x66, 0x62, 0x40, 0xdd, 0xa9, 0xb3, 0xf0, 0x56, 0xe3, 0x7a, 0xf4, 0x3a, 0x27,
	0x7f, 0x63, 0x73, 0x0a, 0x47, 0xd6, 0x7b, 0xa3, 0xad, 0xb7, 0x8d, 0x13, 0xfe, 0x06, 0xc3, 0xfa,
	0xd9, 0x54, 0x16, 0x4a, 0xc7, 0xe6, 0xe9, 0xe6, 0x03, 0xab, 0xdf, 0x59, 0x19, 0xce, 0x41, 0xbf,
	0x46, 0xed, 0x94, 0x5f, 0x34, 0xd8, 0xb3, 0xaa, 0x19, 0x74, 0xf6, 0x08, 0x06, 0x1f, 0xd2, 0x82,
	0xdf, 0x40, 0xdb, 0x26, 0x63, 0x36, 0x62, 0xab, 0x6a, 0x4f, 0x19, 0x56, 0xa9, 0xdd, 0x8a, 0x3e,
	0x43, 0x04, 0x3a, 0x05, 0xcd, 0xcf, 0x2c, 0xa6, 0xc8, 0xd5, 0x72, 0x5d, 0xae, 0x8f, 0x60, 0x14,
	0xcb, 0xd3, 0xb5, 0xf5, 0xac, 0x87, 0x17, 0xdb, 0xdf, 0x56, 0x1f, 0xbf, 0x75, 0x9e, 0xfe, 0x59,
	0x30, 0x91, 0x1c, 0x8b, 0xc4, 0x97, 0x79, 0x12, 0x24, 0x54, 0xe8, 0x68, 0x02, 0x73, 0x85, 0x33,
	0x56, 0xbc, 0xfb, 0x5b, 0xff, 0x36, 0xc5, 0xbe, 0xad, 0xc1, 0x3f, 0xaf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x7a, 0x51, 0x2e, 0x3c, 0xd5, 0x02, 0x00, 0x00,
}
