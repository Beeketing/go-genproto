// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/genomics/v1/operations.proto

package genomics // import "github.com/Beeketing/go-genproto/googleapis/genomics/v1"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/Beeketing/protobuf/ptypes/any"
import timestamp "github.com/Beeketing/protobuf/ptypes/timestamp"
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

// Metadata describing an [Operation][google.longrunning.Operation].
type OperationMetadata struct {
	// The Google Cloud Project in which the job is scoped.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The time at which the job was submitted to the Genomics service.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The time at which the job began to run.
	StartTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The time at which the job stopped running.
	EndTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// The original request that started the operation. Note that this will be in
	// current version of the API. If the operation was started with v1beta2 API
	// and a GetOperation is performed on v1 API, a v1 request will be returned.
	Request *any.Any `protobuf:"bytes,5,opt,name=request,proto3" json:"request,omitempty"`
	// Optional event messages that were generated during the job's execution.
	// This also contains any warnings that were generated during import
	// or export.
	Events []*OperationEvent `protobuf:"bytes,6,rep,name=events,proto3" json:"events,omitempty"`
	// This field is deprecated. Use `labels` instead. Optionally provided by the
	// caller when submitting the request that creates the operation.
	ClientId string `protobuf:"bytes,7,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// Runtime metadata on this Operation.
	RuntimeMetadata *any.Any `protobuf:"bytes,8,opt,name=runtime_metadata,json=runtimeMetadata,proto3" json:"runtime_metadata,omitempty"`
	// Optionally provided by the caller when submitting the request that creates
	// the operation.
	Labels               map[string]string `protobuf:"bytes,9,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *OperationMetadata) Reset()         { *m = OperationMetadata{} }
func (m *OperationMetadata) String() string { return proto.CompactTextString(m) }
func (*OperationMetadata) ProtoMessage()    {}
func (*OperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_operations_4f155d6eb213ff75, []int{0}
}
func (m *OperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationMetadata.Unmarshal(m, b)
}
func (m *OperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationMetadata.Marshal(b, m, deterministic)
}
func (dst *OperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationMetadata.Merge(dst, src)
}
func (m *OperationMetadata) XXX_Size() int {
	return xxx_messageInfo_OperationMetadata.Size(m)
}
func (m *OperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_OperationMetadata proto.InternalMessageInfo

func (m *OperationMetadata) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *OperationMetadata) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *OperationMetadata) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *OperationMetadata) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *OperationMetadata) GetRequest() *any.Any {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *OperationMetadata) GetEvents() []*OperationEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *OperationMetadata) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *OperationMetadata) GetRuntimeMetadata() *any.Any {
	if m != nil {
		return m.RuntimeMetadata
	}
	return nil
}

func (m *OperationMetadata) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// An event that occurred during an [Operation][google.longrunning.Operation].
type OperationEvent struct {
	// Optional time of when event started.
	StartTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Optional time of when event finished. An event can have a start time and no
	// finish time. If an event has a finish time, there must be a start time.
	EndTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Required description of event.
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationEvent) Reset()         { *m = OperationEvent{} }
func (m *OperationEvent) String() string { return proto.CompactTextString(m) }
func (*OperationEvent) ProtoMessage()    {}
func (*OperationEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_operations_4f155d6eb213ff75, []int{1}
}
func (m *OperationEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationEvent.Unmarshal(m, b)
}
func (m *OperationEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationEvent.Marshal(b, m, deterministic)
}
func (dst *OperationEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationEvent.Merge(dst, src)
}
func (m *OperationEvent) XXX_Size() int {
	return xxx_messageInfo_OperationEvent.Size(m)
}
func (m *OperationEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationEvent.DiscardUnknown(m)
}

var xxx_messageInfo_OperationEvent proto.InternalMessageInfo

func (m *OperationEvent) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *OperationEvent) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *OperationEvent) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*OperationMetadata)(nil), "google.genomics.v1.OperationMetadata")
	proto.RegisterMapType((map[string]string)(nil), "google.genomics.v1.OperationMetadata.LabelsEntry")
	proto.RegisterType((*OperationEvent)(nil), "google.genomics.v1.OperationEvent")
}

func init() {
	proto.RegisterFile("google/genomics/v1/operations.proto", fileDescriptor_operations_4f155d6eb213ff75)
}

var fileDescriptor_operations_4f155d6eb213ff75 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xe5, 0x76, 0x6b, 0x9b, 0x17, 0x89, 0x0d, 0x6b, 0x42, 0xa1, 0x80, 0xa8, 0xca, 0xa5,
	0x27, 0x47, 0x1d, 0x42, 0x62, 0xdd, 0x01, 0x31, 0x69, 0x87, 0x4a, 0x20, 0xa6, 0x88, 0x13, 0x97,
	0xca, 0x4d, 0x1e, 0x51, 0x46, 0x62, 0x07, 0xdb, 0xad, 0xd4, 0xef, 0xc3, 0x17, 0xe0, 0xdb, 0x71,
	0x44, 0xb1, 0x9d, 0x2a, 0x6c, 0x68, 0x45, 0xdc, 0xec, 0xf7, 0xfe, 0x3f, 0xfb, 0x9f, 0xf7, 0x8f,
	0xe1, 0x55, 0x2e, 0x65, 0x5e, 0x62, 0x9c, 0xa3, 0x90, 0x55, 0x91, 0xea, 0x78, 0x3b, 0x8f, 0x65,
	0x8d, 0x8a, 0x9b, 0x42, 0x0a, 0xcd, 0x6a, 0x25, 0x8d, 0xa4, 0xd4, 0x89, 0x58, 0x2b, 0x62, 0xdb,
	0xf9, 0xf8, 0xb9, 0x07, 0x79, 0x5d, 0xc4, 0x5c, 0x08, 0x69, 0xba, 0xc4, 0xf8, 0xa9, 0xef, 0xda,
	0xdd, 0x7a, 0xf3, 0x35, 0xe6, 0x62, 0xe7, 0x5b, 0x2f, 0xef, 0xb6, 0x4c, 0x51, 0xa1, 0x36, 0xbc,
	0xaa, 0x9d, 0x60, 0xfa, 0xf3, 0x08, 0x1e, 0x7f, 0x6a, 0x2d, 0x7c, 0x44, 0xc3, 0x33, 0x6e, 0x38,
	0x7d, 0x01, 0x50, 0x2b, 0x79, 0x8b, 0xa9, 0x59, 0x15, 0x59, 0x44, 0x26, 0x64, 0x16, 0x24, 0x81,
	0xaf, 0x2c, 0x33, 0x7a, 0x09, 0x61, 0xaa, 0x90, 0x1b, 0x5c, 0x35, 0xc7, 0x45, 0xbd, 0x09, 0x99,
	0x85, 0xe7, 0x63, 0xe6, 0x8d, 0xb7, 0x77, 0xb1, 0xcf, 0xed, 0x5d, 0x09, 0x38, 0x79, 0x53, 0xa0,
	0x17, 0x00, 0xda, 0x70, 0x65, 0x1c, 0xdb, 0x3f, 0xc8, 0x06, 0x56, 0x6d, 0xd1, 0x37, 0x30, 0x42,
	0x91, 0x39, 0xf0, 0xe8, 0x20, 0x38, 0x44, 0x91, 0x59, 0x8c, 0xc1, 0x50, 0xe1, 0xf7, 0x0d, 0x6a,
	0x13, 0x1d, 0x5b, 0xea, 0xec, 0x1e, 0xf5, 0x5e, 0xec, 0x92, 0x56, 0x44, 0x17, 0x30, 0xc0, 0x2d,
	0x0a, 0xa3, 0xa3, 0xc1, 0xa4, 0x3f, 0x0b, 0xcf, 0xa7, 0xec, 0x7e, 0x24, 0x6c, 0x3f, 0xb4, 0xeb,
	0x46, 0x9a, 0x78, 0x82, 0x3e, 0x83, 0x20, 0x2d, 0x0b, 0x14, 0x76, 0x70, 0x43, 0x3b, 0xb8, 0x91,
	0x2b, 0x2c, 0x33, 0xfa, 0x0e, 0x4e, 0xd5, 0x46, 0x34, 0xf6, 0x57, 0x95, 0x1f, 0x75, 0x34, 0x7a,
	0xc0, 0xd1, 0x89, 0x57, 0xef, 0x73, 0x59, 0xc2, 0xa0, 0xe4, 0x6b, 0x2c, 0x75, 0x14, 0x58, 0x67,
	0xf3, 0x07, 0x9d, 0xb5, 0x18, 0xfb, 0x60, 0x99, 0x6b, 0x61, 0xd4, 0x2e, 0xf1, 0x07, 0x8c, 0x2f,
	0x20, 0xec, 0x94, 0xe9, 0x29, 0xf4, 0xbf, 0xe1, 0xce, 0x47, 0xdd, 0x2c, 0xe9, 0x19, 0x1c, 0x6f,
	0x79, 0xb9, 0x71, 0xf1, 0x06, 0x89, 0xdb, 0x2c, 0x7a, 0x6f, 0xc9, 0xf4, 0x07, 0x81, 0x47, 0x7f,
	0x7e, 0xfe, 0x9d, 0x50, 0xc9, 0xff, 0x86, 0xda, 0xfb, 0xf7, 0x50, 0x27, 0x10, 0x66, 0xa8, 0x53,
	0x55, 0xd4, 0x8d, 0x0b, 0xfb, 0x1f, 0x05, 0x49, 0xb7, 0x74, 0x75, 0x0b, 0x4f, 0x52, 0x59, 0xfd,
	0x65, 0x42, 0x57, 0x27, 0x7b, 0xf7, 0xfa, 0xa6, 0xb9, 0xe2, 0x86, 0x7c, 0x59, 0xb4, 0x32, 0x59,
	0x72, 0x91, 0x33, 0xa9, 0xf2, 0xe6, 0x95, 0x5a, 0x03, 0xb1, 0x6b, 0xf1, 0xba, 0xd0, 0xdd, 0x97,
	0x7b, 0xd9, 0xae, 0x7f, 0x11, 0xb2, 0x1e, 0x58, 0xe5, 0xeb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x6a, 0xf6, 0xa8, 0x9a, 0xe2, 0x03, 0x00, 0x00,
}
