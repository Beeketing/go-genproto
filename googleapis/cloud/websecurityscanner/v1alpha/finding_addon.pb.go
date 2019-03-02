// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/websecurityscanner/v1alpha/finding_addon.proto

package websecurityscanner // import "github.com/Beeketing/go-genproto/googleapis/cloud/websecurityscanner/v1alpha"

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

// Information reported for an outdated library.
type OutdatedLibrary struct {
	// The name of the outdated library.
	LibraryName string `protobuf:"bytes,1,opt,name=library_name,json=libraryName,proto3" json:"library_name,omitempty"`
	// The version number.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// URLs to learn more information about the vulnerabilities in the library.
	LearnMoreUrls        []string `protobuf:"bytes,3,rep,name=learn_more_urls,json=learnMoreUrls,proto3" json:"learn_more_urls,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutdatedLibrary) Reset()         { *m = OutdatedLibrary{} }
func (m *OutdatedLibrary) String() string { return proto.CompactTextString(m) }
func (*OutdatedLibrary) ProtoMessage()    {}
func (*OutdatedLibrary) Descriptor() ([]byte, []int) {
	return fileDescriptor_finding_addon_10f44c23dab1483b, []int{0}
}
func (m *OutdatedLibrary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutdatedLibrary.Unmarshal(m, b)
}
func (m *OutdatedLibrary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutdatedLibrary.Marshal(b, m, deterministic)
}
func (dst *OutdatedLibrary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutdatedLibrary.Merge(dst, src)
}
func (m *OutdatedLibrary) XXX_Size() int {
	return xxx_messageInfo_OutdatedLibrary.Size(m)
}
func (m *OutdatedLibrary) XXX_DiscardUnknown() {
	xxx_messageInfo_OutdatedLibrary.DiscardUnknown(m)
}

var xxx_messageInfo_OutdatedLibrary proto.InternalMessageInfo

func (m *OutdatedLibrary) GetLibraryName() string {
	if m != nil {
		return m.LibraryName
	}
	return ""
}

func (m *OutdatedLibrary) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *OutdatedLibrary) GetLearnMoreUrls() []string {
	if m != nil {
		return m.LearnMoreUrls
	}
	return nil
}

// Information regarding any resource causing the vulnerability such
// as JavaScript sources, image, audio files, etc.
type ViolatingResource struct {
	// The MIME type of this resource.
	ContentType string `protobuf:"bytes,1,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// URL of this violating resource.
	ResourceUrl          string   `protobuf:"bytes,2,opt,name=resource_url,json=resourceUrl,proto3" json:"resource_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViolatingResource) Reset()         { *m = ViolatingResource{} }
func (m *ViolatingResource) String() string { return proto.CompactTextString(m) }
func (*ViolatingResource) ProtoMessage()    {}
func (*ViolatingResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_finding_addon_10f44c23dab1483b, []int{1}
}
func (m *ViolatingResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViolatingResource.Unmarshal(m, b)
}
func (m *ViolatingResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViolatingResource.Marshal(b, m, deterministic)
}
func (dst *ViolatingResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViolatingResource.Merge(dst, src)
}
func (m *ViolatingResource) XXX_Size() int {
	return xxx_messageInfo_ViolatingResource.Size(m)
}
func (m *ViolatingResource) XXX_DiscardUnknown() {
	xxx_messageInfo_ViolatingResource.DiscardUnknown(m)
}

var xxx_messageInfo_ViolatingResource proto.InternalMessageInfo

func (m *ViolatingResource) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *ViolatingResource) GetResourceUrl() string {
	if m != nil {
		return m.ResourceUrl
	}
	return ""
}

// Information about vulnerable request parameters.
type VulnerableParameters struct {
	// The vulnerable parameter names.
	ParameterNames       []string `protobuf:"bytes,1,rep,name=parameter_names,json=parameterNames,proto3" json:"parameter_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VulnerableParameters) Reset()         { *m = VulnerableParameters{} }
func (m *VulnerableParameters) String() string { return proto.CompactTextString(m) }
func (*VulnerableParameters) ProtoMessage()    {}
func (*VulnerableParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_finding_addon_10f44c23dab1483b, []int{2}
}
func (m *VulnerableParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VulnerableParameters.Unmarshal(m, b)
}
func (m *VulnerableParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VulnerableParameters.Marshal(b, m, deterministic)
}
func (dst *VulnerableParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VulnerableParameters.Merge(dst, src)
}
func (m *VulnerableParameters) XXX_Size() int {
	return xxx_messageInfo_VulnerableParameters.Size(m)
}
func (m *VulnerableParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_VulnerableParameters.DiscardUnknown(m)
}

var xxx_messageInfo_VulnerableParameters proto.InternalMessageInfo

func (m *VulnerableParameters) GetParameterNames() []string {
	if m != nil {
		return m.ParameterNames
	}
	return nil
}

// Information reported for an XSS.
type Xss struct {
	// Stack traces leading to the point where the XSS occurred.
	StackTraces []string `protobuf:"bytes,1,rep,name=stack_traces,json=stackTraces,proto3" json:"stack_traces,omitempty"`
	// An error message generated by a javascript breakage.
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Xss) Reset()         { *m = Xss{} }
func (m *Xss) String() string { return proto.CompactTextString(m) }
func (*Xss) ProtoMessage()    {}
func (*Xss) Descriptor() ([]byte, []int) {
	return fileDescriptor_finding_addon_10f44c23dab1483b, []int{3}
}
func (m *Xss) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Xss.Unmarshal(m, b)
}
func (m *Xss) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Xss.Marshal(b, m, deterministic)
}
func (dst *Xss) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Xss.Merge(dst, src)
}
func (m *Xss) XXX_Size() int {
	return xxx_messageInfo_Xss.Size(m)
}
func (m *Xss) XXX_DiscardUnknown() {
	xxx_messageInfo_Xss.DiscardUnknown(m)
}

var xxx_messageInfo_Xss proto.InternalMessageInfo

func (m *Xss) GetStackTraces() []string {
	if m != nil {
		return m.StackTraces
	}
	return nil
}

func (m *Xss) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*OutdatedLibrary)(nil), "google.cloud.websecurityscanner.v1alpha.OutdatedLibrary")
	proto.RegisterType((*ViolatingResource)(nil), "google.cloud.websecurityscanner.v1alpha.ViolatingResource")
	proto.RegisterType((*VulnerableParameters)(nil), "google.cloud.websecurityscanner.v1alpha.VulnerableParameters")
	proto.RegisterType((*Xss)(nil), "google.cloud.websecurityscanner.v1alpha.Xss")
}

func init() {
	proto.RegisterFile("google/cloud/websecurityscanner/v1alpha/finding_addon.proto", fileDescriptor_finding_addon_10f44c23dab1483b)
}

var fileDescriptor_finding_addon_10f44c23dab1483b = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6b, 0xd4, 0x50,
	0x10, 0xc6, 0x89, 0x0b, 0x4a, 0xdf, 0xb6, 0x2e, 0x0d, 0x1e, 0x82, 0x78, 0xa8, 0x2b, 0xd8, 0x82,
	0x90, 0x20, 0x1e, 0x7b, 0x10, 0x7b, 0xf0, 0xe4, 0xea, 0xb2, 0xb4, 0xc5, 0x7a, 0x09, 0xb3, 0xc9,
	0xf8, 0x7c, 0xf8, 0x32, 0x13, 0x66, 0x5e, 0x56, 0xf2, 0xc7, 0xf8, 0xbf, 0x4a, 0x5e, 0xb2, 0x7b,
	0xd9, 0x83, 0xbd, 0xe5, 0xfd, 0x66, 0xbe, 0xf9, 0x66, 0xf8, 0x62, 0xae, 0x2d, 0xb3, 0xf5, 0x58,
	0x54, 0x9e, 0xbb, 0xba, 0xf8, 0x83, 0x5b, 0xc5, 0xaa, 0x13, 0x17, 0x7a, 0xad, 0x80, 0x08, 0xa5,
	0xd8, 0xbd, 0x07, 0xdf, 0xfe, 0x82, 0xe2, 0xa7, 0xa3, 0xda, 0x91, 0x2d, 0xa1, 0xae, 0x99, 0xf2,
	0x56, 0x38, 0x70, 0x7a, 0x39, 0x8a, 0xf3, 0x28, 0xce, 0x8f, 0xc5, 0xf9, 0x24, 0x7e, 0xf9, 0x6a,
	0x72, 0x81, 0xd6, 0x15, 0x40, 0xc4, 0x01, 0x82, 0x63, 0xd2, 0x71, 0xcc, 0x72, 0x67, 0x16, 0xdf,
	0xba, 0x50, 0x43, 0xc0, 0xfa, 0x8b, 0xdb, 0x0a, 0x48, 0x9f, 0xbe, 0x36, 0xa7, 0x7e, 0xfc, 0x2c,
	0x09, 0x1a, 0xcc, 0x92, 0x8b, 0xe4, 0xea, 0x64, 0x33, 0x9f, 0xd8, 0x57, 0x68, 0x30, 0xcd, 0xcc,
	0xb3, 0x1d, 0x8a, 0x3a, 0xa6, 0xec, 0x49, 0xac, 0xee, 0x9f, 0xe9, 0x5b, 0xb3, 0xf0, 0x08, 0x42,
	0x65, 0xc3, 0x82, 0x65, 0x27, 0x5e, 0xb3, 0xd9, 0xc5, 0xec, 0xea, 0x64, 0x73, 0x16, 0xf1, 0x8a,
	0x05, 0xef, 0xc4, 0xeb, 0xf2, 0xc1, 0x9c, 0xdf, 0x3b, 0xf6, 0x10, 0x1c, 0xd9, 0x0d, 0x2a, 0x77,
	0x52, 0xe1, 0xe0, 0x5c, 0x31, 0x05, 0xa4, 0x50, 0x86, 0xbe, 0x3d, 0x38, 0x4f, 0xec, 0xb6, 0x6f,
	0x63, 0x8b, 0x4c, 0xed, 0xc3, 0xf4, 0xc9, 0x7e, 0xbe, 0x67, 0x77, 0xe2, 0x97, 0x1f, 0xcd, 0x8b,
	0xfb, 0xce, 0x13, 0x0a, 0x6c, 0x3d, 0xae, 0x41, 0xa0, 0xc1, 0x80, 0xa2, 0xe9, 0xa5, 0x59, 0xb4,
	0xfb, 0x57, 0xbc, 0x4c, 0xb3, 0x24, 0xae, 0xf6, 0xfc, 0x80, 0x87, 0xe3, 0x74, 0xb9, 0x32, 0xb3,
	0xef, 0xaa, 0x83, 0x95, 0x06, 0xa8, 0x7e, 0x97, 0x41, 0xa0, 0x3a, 0x34, 0xcf, 0x23, 0xbb, 0x8d,
	0x28, 0x7d, 0x63, 0xce, 0x50, 0x84, 0xa5, 0x6c, 0x50, 0x15, 0x2c, 0x4e, 0xeb, 0x9c, 0x46, 0xb8,
	0x1a, 0xd9, 0xcd, 0xdf, 0xc4, 0xbc, 0xab, 0xb8, 0xc9, 0x1f, 0x19, 0xd8, 0xcd, 0xf9, 0xe7, 0x31,
	0xee, 0x4f, 0x43, 0xda, 0xeb, 0x21, 0xa5, 0x75, 0xf2, 0xe3, 0x61, 0x52, 0x5b, 0xf6, 0x40, 0x36,
	0x67, 0xb1, 0x85, 0x45, 0x8a, 0x19, 0x16, 0x63, 0x09, 0x5a, 0xa7, 0xff, 0xfd, 0x95, 0xae, 0x8f,
	0x4b, 0xdb, 0xa7, 0x71, 0xca, 0x87, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x61, 0x84, 0x84,
	0x8f, 0x02, 0x00, 0x00,
}
