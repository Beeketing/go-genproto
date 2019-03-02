// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/websecurityscanner/v1alpha/scan_run.proto

package websecurityscanner // import "github.com/Beeketing/genproto/googleapis/cloud/websecurityscanner/v1alpha"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/Beeketing/protobuf/ptypes/timestamp"
import _ "github.com/Beeketing/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Types of ScanRun execution state.
type ScanRun_ExecutionState int32

const (
	// Represents an invalid state caused by internal server error. This value
	// should never be returned.
	ScanRun_EXECUTION_STATE_UNSPECIFIED ScanRun_ExecutionState = 0
	// The scan is waiting in the queue.
	ScanRun_QUEUED ScanRun_ExecutionState = 1
	// The scan is in progress.
	ScanRun_SCANNING ScanRun_ExecutionState = 2
	// The scan is either finished or stopped by user.
	ScanRun_FINISHED ScanRun_ExecutionState = 3
)

var ScanRun_ExecutionState_name = map[int32]string{
	0: "EXECUTION_STATE_UNSPECIFIED",
	1: "QUEUED",
	2: "SCANNING",
	3: "FINISHED",
}
var ScanRun_ExecutionState_value = map[string]int32{
	"EXECUTION_STATE_UNSPECIFIED": 0,
	"QUEUED":                      1,
	"SCANNING":                    2,
	"FINISHED":                    3,
}

func (x ScanRun_ExecutionState) String() string {
	return proto.EnumName(ScanRun_ExecutionState_name, int32(x))
}
func (ScanRun_ExecutionState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_scan_run_8ce9e7c54bb44c79, []int{0, 0}
}

// Types of ScanRun result state.
type ScanRun_ResultState int32

const (
	// Default value. This value is returned when the ScanRun is not yet
	// finished.
	ScanRun_RESULT_STATE_UNSPECIFIED ScanRun_ResultState = 0
	// The scan finished without errors.
	ScanRun_SUCCESS ScanRun_ResultState = 1
	// The scan finished with errors.
	ScanRun_ERROR ScanRun_ResultState = 2
	// The scan was terminated by user.
	ScanRun_KILLED ScanRun_ResultState = 3
)

var ScanRun_ResultState_name = map[int32]string{
	0: "RESULT_STATE_UNSPECIFIED",
	1: "SUCCESS",
	2: "ERROR",
	3: "KILLED",
}
var ScanRun_ResultState_value = map[string]int32{
	"RESULT_STATE_UNSPECIFIED": 0,
	"SUCCESS":                  1,
	"ERROR":                    2,
	"KILLED":                   3,
}

func (x ScanRun_ResultState) String() string {
	return proto.EnumName(ScanRun_ResultState_name, int32(x))
}
func (ScanRun_ResultState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_scan_run_8ce9e7c54bb44c79, []int{0, 1}
}

// A ScanRun is a output-only resource representing an actual run of the scan.
type ScanRun struct {
	// Output only.
	// The resource name of the ScanRun. The name follows the format of
	// 'projects/{projectId}/scanConfigs/{scanConfigId}/scanRuns/{scanRunId}'.
	// The ScanRun IDs are generated by the system.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only.
	// The execution state of the ScanRun.
	ExecutionState ScanRun_ExecutionState `protobuf:"varint,2,opt,name=execution_state,json=executionState,proto3,enum=google.cloud.websecurityscanner.v1alpha.ScanRun_ExecutionState" json:"execution_state,omitempty"`
	// Output only.
	// The result state of the ScanRun. This field is only available after the
	// execution state reaches "FINISHED".
	ResultState ScanRun_ResultState `protobuf:"varint,3,opt,name=result_state,json=resultState,proto3,enum=google.cloud.websecurityscanner.v1alpha.ScanRun_ResultState" json:"result_state,omitempty"`
	// Output only.
	// The time at which the ScanRun started.
	StartTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Output only.
	// The time at which the ScanRun reached termination state - that the ScanRun
	// is either finished or stopped by user.
	EndTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Output only.
	// The number of URLs crawled during this ScanRun. If the scan is in progress,
	// the value represents the number of URLs crawled up to now.
	UrlsCrawledCount int64 `protobuf:"varint,6,opt,name=urls_crawled_count,json=urlsCrawledCount,proto3" json:"urls_crawled_count,omitempty"`
	// Output only.
	// The number of URLs tested during this ScanRun. If the scan is in progress,
	// the value represents the number of URLs tested up to now. The number of
	// URLs tested is usually larger than the number URLS crawled because
	// typically a crawled URL is tested with multiple test payloads.
	UrlsTestedCount int64 `protobuf:"varint,7,opt,name=urls_tested_count,json=urlsTestedCount,proto3" json:"urls_tested_count,omitempty"`
	// Output only.
	// Whether the scan run has found any vulnerabilities.
	HasVulnerabilities bool `protobuf:"varint,8,opt,name=has_vulnerabilities,json=hasVulnerabilities,proto3" json:"has_vulnerabilities,omitempty"`
	// Output only.
	// The percentage of total completion ranging from 0 to 100.
	// If the scan is in queue, the value is 0.
	// If the scan is running, the value ranges from 0 to 100.
	// If the scan is finished, the value is 100.
	ProgressPercent      int32    `protobuf:"varint,9,opt,name=progress_percent,json=progressPercent,proto3" json:"progress_percent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanRun) Reset()         { *m = ScanRun{} }
func (m *ScanRun) String() string { return proto.CompactTextString(m) }
func (*ScanRun) ProtoMessage()    {}
func (*ScanRun) Descriptor() ([]byte, []int) {
	return fileDescriptor_scan_run_8ce9e7c54bb44c79, []int{0}
}
func (m *ScanRun) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanRun.Unmarshal(m, b)
}
func (m *ScanRun) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanRun.Marshal(b, m, deterministic)
}
func (dst *ScanRun) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanRun.Merge(dst, src)
}
func (m *ScanRun) XXX_Size() int {
	return xxx_messageInfo_ScanRun.Size(m)
}
func (m *ScanRun) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanRun.DiscardUnknown(m)
}

var xxx_messageInfo_ScanRun proto.InternalMessageInfo

func (m *ScanRun) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ScanRun) GetExecutionState() ScanRun_ExecutionState {
	if m != nil {
		return m.ExecutionState
	}
	return ScanRun_EXECUTION_STATE_UNSPECIFIED
}

func (m *ScanRun) GetResultState() ScanRun_ResultState {
	if m != nil {
		return m.ResultState
	}
	return ScanRun_RESULT_STATE_UNSPECIFIED
}

func (m *ScanRun) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ScanRun) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *ScanRun) GetUrlsCrawledCount() int64 {
	if m != nil {
		return m.UrlsCrawledCount
	}
	return 0
}

func (m *ScanRun) GetUrlsTestedCount() int64 {
	if m != nil {
		return m.UrlsTestedCount
	}
	return 0
}

func (m *ScanRun) GetHasVulnerabilities() bool {
	if m != nil {
		return m.HasVulnerabilities
	}
	return false
}

func (m *ScanRun) GetProgressPercent() int32 {
	if m != nil {
		return m.ProgressPercent
	}
	return 0
}

func init() {
	proto.RegisterType((*ScanRun)(nil), "google.cloud.websecurityscanner.v1alpha.ScanRun")
	proto.RegisterEnum("google.cloud.websecurityscanner.v1alpha.ScanRun_ExecutionState", ScanRun_ExecutionState_name, ScanRun_ExecutionState_value)
	proto.RegisterEnum("google.cloud.websecurityscanner.v1alpha.ScanRun_ResultState", ScanRun_ResultState_name, ScanRun_ResultState_value)
}

func init() {
	proto.RegisterFile("google/cloud/websecurityscanner/v1alpha/scan_run.proto", fileDescriptor_scan_run_8ce9e7c54bb44c79)
}

var fileDescriptor_scan_run_8ce9e7c54bb44c79 = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xdd, 0x6e, 0xd3, 0x3c,
	0x18, 0xfe, 0xb2, 0xad, 0x6b, 0xeb, 0x4e, 0x5b, 0x3e, 0x73, 0x12, 0x8d, 0x49, 0x8b, 0x76, 0x42,
	0xf8, 0x51, 0x22, 0x86, 0x40, 0x42, 0x20, 0xa1, 0x2d, 0xf3, 0x20, 0x62, 0xca, 0x8a, 0x93, 0x20,
	0xc6, 0x49, 0xe4, 0xa6, 0x26, 0x8d, 0x94, 0xda, 0x91, 0xed, 0x6c, 0x70, 0x27, 0x5c, 0x0c, 0x17,
	0x87, 0xec, 0xa4, 0xc0, 0x34, 0xa4, 0x8d, 0xb3, 0x3e, 0xef, 0xf3, 0x57, 0xbd, 0xaf, 0x03, 0x5e,
	0x94, 0x9c, 0x97, 0x35, 0x0d, 0x8a, 0x9a, 0xb7, 0xf3, 0xe0, 0x8a, 0xce, 0x24, 0x2d, 0x5a, 0x51,
	0xa9, 0x6f, 0xb2, 0x20, 0x8c, 0x51, 0x11, 0x5c, 0x3e, 0x25, 0x75, 0xb3, 0x20, 0x81, 0xc6, 0xb9,
	0x68, 0x99, 0xdf, 0x08, 0xae, 0x38, 0x7c, 0xd0, 0xf9, 0x7c, 0xe3, 0xf3, 0x6f, 0xfa, 0xfc, 0xde,
	0xb7, 0xbb, 0xd7, 0x17, 0x90, 0xa6, 0x0a, 0x08, 0x63, 0x5c, 0x11, 0x55, 0x71, 0x26, 0xbb, 0x98,
	0xdd, 0xfd, 0x9e, 0x35, 0x68, 0xd6, 0x7e, 0x09, 0x54, 0xb5, 0xa4, 0x52, 0x91, 0x65, 0xd3, 0x09,
	0x0e, 0x7e, 0x0c, 0xc0, 0x30, 0x29, 0x08, 0xc3, 0x2d, 0x83, 0x10, 0x6c, 0x30, 0xb2, 0xa4, 0x8e,
	0xe5, 0x5a, 0xde, 0x18, 0x9b, 0xdf, 0x70, 0x01, 0x76, 0xe8, 0x57, 0x5a, 0xb4, 0x3a, 0x34, 0x97,
	0x8a, 0x28, 0xea, 0xac, 0xb9, 0x96, 0xb7, 0x7d, 0xf8, 0xc6, 0xbf, 0xe3, 0x3f, 0xf4, 0xfb, 0x78,
	0x1f, 0xad, 0x72, 0x12, 0x1d, 0x83, 0xb7, 0xe9, 0x35, 0x0c, 0x73, 0xb0, 0x25, 0xa8, 0x6c, 0x6b,
	0xd5, 0xd7, 0xac, 0x9b, 0x9a, 0xd7, 0xff, 0x5c, 0x83, 0x4d, 0x48, 0xd7, 0x31, 0x11, 0xbf, 0x01,
	0x7c, 0x09, 0x80, 0x54, 0x44, 0xa8, 0x5c, 0xef, 0xc0, 0xd9, 0x70, 0x2d, 0x6f, 0x72, 0xb8, 0xbb,
	0x8a, 0x5f, 0x2d, 0xc8, 0x4f, 0x57, 0x0b, 0xc2, 0x63, 0xa3, 0xd6, 0x18, 0x3e, 0x07, 0x23, 0xca,
	0xe6, 0x9d, 0x71, 0x70, 0xab, 0x71, 0x48, 0xd9, 0xdc, 0xd8, 0x9e, 0x00, 0xd8, 0x8a, 0x5a, 0xe6,
	0x85, 0x20, 0x57, 0x35, 0x9d, 0xe7, 0x05, 0x6f, 0x99, 0x72, 0x36, 0x5d, 0xcb, 0x5b, 0xc7, 0xb6,
	0x66, 0xc2, 0x8e, 0x08, 0xf5, 0x1c, 0x3e, 0x02, 0xff, 0x1b, 0xb5, 0xa2, 0x52, 0xfd, 0x12, 0x0f,
	0x8d, 0x78, 0x47, 0x13, 0xa9, 0x99, 0x77, 0xda, 0x00, 0xdc, 0x5b, 0x10, 0x99, 0x5f, 0xb6, 0x35,
	0xa3, 0x82, 0xcc, 0xaa, 0xba, 0x52, 0x15, 0x95, 0xce, 0xc8, 0xb5, 0xbc, 0x11, 0x86, 0x0b, 0x22,
	0x3f, 0x5e, 0x67, 0xe0, 0x43, 0x60, 0x37, 0x82, 0x97, 0x82, 0x4a, 0x99, 0x37, 0x54, 0x14, 0x94,
	0x29, 0x67, 0xec, 0x5a, 0xde, 0x00, 0xef, 0xac, 0xe6, 0xd3, 0x6e, 0x7c, 0x70, 0x01, 0xb6, 0xaf,
	0x9f, 0x0a, 0xee, 0x83, 0xfb, 0xe8, 0x13, 0x0a, 0xb3, 0x34, 0x3a, 0x8f, 0xf3, 0x24, 0x3d, 0x4a,
	0x51, 0x9e, 0xc5, 0xc9, 0x14, 0x85, 0xd1, 0x69, 0x84, 0x4e, 0xec, 0xff, 0x20, 0x00, 0x9b, 0x1f,
	0x32, 0x94, 0xa1, 0x13, 0xdb, 0x82, 0x5b, 0x60, 0x94, 0x84, 0x47, 0x71, 0x1c, 0xc5, 0x6f, 0xed,
	0x35, 0x8d, 0x4e, 0xa3, 0x38, 0x4a, 0xde, 0xa1, 0x13, 0x7b, 0xfd, 0xe0, 0x1c, 0x4c, 0xfe, 0x38,
	0x0f, 0xdc, 0x03, 0x0e, 0x46, 0x49, 0x76, 0x96, 0xfe, 0x35, 0x74, 0x02, 0x86, 0x49, 0x16, 0x86,
	0x28, 0x49, 0x6c, 0x0b, 0x8e, 0xc1, 0x00, 0x61, 0x7c, 0x8e, 0xed, 0x35, 0x5d, 0xf6, 0x3e, 0x3a,
	0x3b, 0xd3, 0x81, 0xc7, 0xdf, 0x2d, 0xf0, 0xb8, 0xe0, 0xcb, 0xbb, 0x3e, 0x92, 0xe3, 0xad, 0xfe,
	0x95, 0x4c, 0xf5, 0xd5, 0xa6, 0xd6, 0xe7, 0x8b, 0xde, 0x58, 0xf2, 0x9a, 0xb0, 0xd2, 0xe7, 0xa2,
	0x0c, 0x4a, 0xca, 0xcc, 0x4d, 0x83, 0x8e, 0x22, 0x4d, 0x25, 0x6f, 0xfd, 0x7a, 0x5f, 0xdd, 0xa4,
	0x66, 0x9b, 0x26, 0xe5, 0xd9, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x43, 0xba, 0xc7, 0x02,
	0x04, 0x00, 0x00,
}
