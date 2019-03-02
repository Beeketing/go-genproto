// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/servicecontrol/v1/log_entry.proto

package servicecontrol // import "github.com/Beeketing/go-genproto/googleapis/api/servicecontrol/v1"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/Beeketing/protobuf/ptypes/any"
import _struct "github.com/Beeketing/protobuf/ptypes/struct"
import timestamp "github.com/Beeketing/protobuf/ptypes/timestamp"
import _ "github.com/Beeketing/go-genproto/googleapis/api/annotations"
import _type "github.com/Beeketing/go-genproto/googleapis/logging/type"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// An individual log entry.
type LogEntry struct {
	// Required. The log to which this log entry belongs. Examples: `"syslog"`,
	// `"book_log"`.
	Name string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	// The time the event described by the log entry occurred. If
	// omitted, defaults to operation start time.
	Timestamp *timestamp.Timestamp `protobuf:"bytes,11,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The severity of the log entry. The default value is
	// `LogSeverity.DEFAULT`.
	Severity _type.LogSeverity `protobuf:"varint,12,opt,name=severity,proto3,enum=google.logging.type.LogSeverity" json:"severity,omitempty"`
	// A unique ID for the log entry used for deduplication. If omitted,
	// the implementation will generate one based on operation_id.
	InsertId string `protobuf:"bytes,4,opt,name=insert_id,json=insertId,proto3" json:"insert_id,omitempty"`
	// A set of user-defined (key, value) data that provides additional
	// information about the log entry.
	Labels map[string]string `protobuf:"bytes,13,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The log entry payload, which can be one of multiple types.
	//
	// Types that are valid to be assigned to Payload:
	//	*LogEntry_ProtoPayload
	//	*LogEntry_TextPayload
	//	*LogEntry_StructPayload
	Payload              isLogEntry_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *LogEntry) Reset()         { *m = LogEntry{} }
func (m *LogEntry) String() string { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()    {}
func (*LogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_log_entry_1cf8a0cea23f865a, []int{0}
}
func (m *LogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntry.Unmarshal(m, b)
}
func (m *LogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntry.Marshal(b, m, deterministic)
}
func (dst *LogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntry.Merge(dst, src)
}
func (m *LogEntry) XXX_Size() int {
	return xxx_messageInfo_LogEntry.Size(m)
}
func (m *LogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntry proto.InternalMessageInfo

func (m *LogEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogEntry) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *LogEntry) GetSeverity() _type.LogSeverity {
	if m != nil {
		return m.Severity
	}
	return _type.LogSeverity_DEFAULT
}

func (m *LogEntry) GetInsertId() string {
	if m != nil {
		return m.InsertId
	}
	return ""
}

func (m *LogEntry) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type isLogEntry_Payload interface {
	isLogEntry_Payload()
}

type LogEntry_ProtoPayload struct {
	ProtoPayload *any.Any `protobuf:"bytes,2,opt,name=proto_payload,json=protoPayload,proto3,oneof"`
}

type LogEntry_TextPayload struct {
	TextPayload string `protobuf:"bytes,3,opt,name=text_payload,json=textPayload,proto3,oneof"`
}

type LogEntry_StructPayload struct {
	StructPayload *_struct.Struct `protobuf:"bytes,6,opt,name=struct_payload,json=structPayload,proto3,oneof"`
}

func (*LogEntry_ProtoPayload) isLogEntry_Payload() {}

func (*LogEntry_TextPayload) isLogEntry_Payload() {}

func (*LogEntry_StructPayload) isLogEntry_Payload() {}

func (m *LogEntry) GetPayload() isLogEntry_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *LogEntry) GetProtoPayload() *any.Any {
	if x, ok := m.GetPayload().(*LogEntry_ProtoPayload); ok {
		return x.ProtoPayload
	}
	return nil
}

func (m *LogEntry) GetTextPayload() string {
	if x, ok := m.GetPayload().(*LogEntry_TextPayload); ok {
		return x.TextPayload
	}
	return ""
}

func (m *LogEntry) GetStructPayload() *_struct.Struct {
	if x, ok := m.GetPayload().(*LogEntry_StructPayload); ok {
		return x.StructPayload
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LogEntry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LogEntry_OneofMarshaler, _LogEntry_OneofUnmarshaler, _LogEntry_OneofSizer, []interface{}{
		(*LogEntry_ProtoPayload)(nil),
		(*LogEntry_TextPayload)(nil),
		(*LogEntry_StructPayload)(nil),
	}
}

func _LogEntry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LogEntry)
	// payload
	switch x := m.Payload.(type) {
	case *LogEntry_ProtoPayload:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProtoPayload); err != nil {
			return err
		}
	case *LogEntry_TextPayload:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.TextPayload)
	case *LogEntry_StructPayload:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StructPayload); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LogEntry.Payload has unexpected type %T", x)
	}
	return nil
}

func _LogEntry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LogEntry)
	switch tag {
	case 2: // payload.proto_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(any.Any)
		err := b.DecodeMessage(msg)
		m.Payload = &LogEntry_ProtoPayload{msg}
		return true, err
	case 3: // payload.text_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Payload = &LogEntry_TextPayload{x}
		return true, err
	case 6: // payload.struct_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(_struct.Struct)
		err := b.DecodeMessage(msg)
		m.Payload = &LogEntry_StructPayload{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LogEntry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LogEntry)
	// payload
	switch x := m.Payload.(type) {
	case *LogEntry_ProtoPayload:
		s := proto.Size(x.ProtoPayload)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LogEntry_TextPayload:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.TextPayload)))
		n += len(x.TextPayload)
	case *LogEntry_StructPayload:
		s := proto.Size(x.StructPayload)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*LogEntry)(nil), "google.api.servicecontrol.v1.LogEntry")
	proto.RegisterMapType((map[string]string)(nil), "google.api.servicecontrol.v1.LogEntry.LabelsEntry")
}

func init() {
	proto.RegisterFile("google/api/servicecontrol/v1/log_entry.proto", fileDescriptor_log_entry_1cf8a0cea23f865a)
}

var fileDescriptor_log_entry_1cf8a0cea23f865a = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x4f, 0x8f, 0xd3, 0x30,
	0x10, 0xc5, 0x9b, 0xed, 0x52, 0x1a, 0xa7, 0x5d, 0x21, 0x6b, 0x25, 0x42, 0xa8, 0x44, 0x04, 0x12,
	0xea, 0x01, 0x39, 0xda, 0x72, 0x59, 0x58, 0x0e, 0x50, 0x09, 0xa9, 0xac, 0x7a, 0x58, 0x65, 0x39,
	0x71, 0xa9, 0xdc, 0xd6, 0x58, 0x16, 0xae, 0x27, 0x4a, 0xdc, 0x88, 0x9c, 0xf8, 0xbc, 0x7c, 0x0b,
	0x94, 0xb1, 0x53, 0xf6, 0x9f, 0xf6, 0xe6, 0xf1, 0xfb, 0xbd, 0x99, 0x79, 0x8e, 0x42, 0xde, 0x49,
	0x00, 0xa9, 0x45, 0xc6, 0x0b, 0x95, 0x55, 0xa2, 0xac, 0xd5, 0x46, 0x6c, 0xc0, 0xd8, 0x12, 0x74,
	0x56, 0x9f, 0x65, 0x1a, 0xe4, 0x4a, 0x18, 0x5b, 0x36, 0xac, 0x28, 0xc1, 0x02, 0x9d, 0x38, 0x9a,
	0xf1, 0x42, 0xb1, 0xdb, 0x34, 0xab, 0xcf, 0x92, 0xc9, 0x8d, 0x5e, 0xdc, 0x18, 0xb0, 0xdc, 0x2a,
	0x30, 0x95, 0xf3, 0x26, 0x6f, 0xbd, 0xaa, 0x41, 0x4a, 0x65, 0x64, 0x66, 0x9b, 0x02, 0x8b, 0x55,
	0x25, 0x6a, 0x51, 0x2a, 0xeb, 0x67, 0x24, 0x2f, 0x3c, 0x87, 0xd5, 0x7a, 0xff, 0x33, 0xe3, 0xa6,
	0x93, 0x26, 0x77, 0xa5, 0xca, 0x96, 0xfb, 0x8d, 0xf5, 0xea, 0xab, 0xbb, 0xaa, 0x55, 0x3b, 0x51,
	0x59, 0xbe, 0x2b, 0x1c, 0xf0, 0xfa, 0x6f, 0x9f, 0x0c, 0x97, 0x20, 0xbf, 0xb6, 0x81, 0x28, 0x25,
	0xc7, 0x86, 0xef, 0x44, 0x4c, 0xd2, 0x60, 0x1a, 0xe6, 0x78, 0xa6, 0xe7, 0x24, 0x3c, 0x78, 0xe2,
	0x28, 0x0d, 0xa6, 0xd1, 0x2c, 0x61, 0x3e, 0x72, 0xd7, 0x95, 0x7d, 0xef, 0x88, 0xfc, 0x3f, 0x4c,
	0x3f, 0x91, 0x61, 0x17, 0x23, 0x1e, 0xa5, 0xc1, 0xf4, 0x64, 0x96, 0x76, 0x46, 0x9f, 0x97, 0xb5,
	0x79, 0xd9, 0x12, 0xe4, 0xb5, 0xe7, 0xf2, 0x83, 0x83, 0xbe, 0x24, 0xa1, 0x32, 0x95, 0x28, 0xed,
	0x4a, 0x6d, 0xe3, 0x63, 0x5c, 0x68, 0xe8, 0x2e, 0xbe, 0x6d, 0xe9, 0x25, 0x19, 0x68, 0xbe, 0x16,
	0xba, 0x8a, 0xc7, 0x69, 0x7f, 0x1a, 0xcd, 0x66, 0xec, 0xb1, 0x8f, 0xc0, 0xba, 0x80, 0x6c, 0x89,
	0x26, 0x3c, 0xe7, 0xbe, 0x03, 0xbd, 0x20, 0x63, 0xcc, 0xb1, 0x2a, 0x78, 0xa3, 0x81, 0x6f, 0xe3,
	0x23, 0x0c, 0x79, 0x7a, 0x2f, 0xe4, 0x17, 0xd3, 0x2c, 0x7a, 0xf9, 0x08, 0xeb, 0x2b, 0xc7, 0xd2,
	0x37, 0x64, 0x64, 0xc5, 0x6f, 0x7b, 0xf0, 0xf6, 0xdb, 0x45, 0x17, 0xbd, 0x3c, 0x6a, 0x6f, 0x3b,
	0xe8, 0x33, 0x39, 0x71, 0x1f, 0xe5, 0x80, 0x0d, 0x70, 0xc4, 0xf3, 0x7b, 0x23, 0xae, 0x11, 0x5b,
	0xf4, 0xf2, 0xb1, 0x33, 0xf8, 0x0e, 0xc9, 0x07, 0x12, 0xdd, 0x58, 0x9d, 0x3e, 0x23, 0xfd, 0x5f,
	0xa2, 0x89, 0x03, 0x7c, 0x95, 0xf6, 0x48, 0x4f, 0xc9, 0x93, 0x9a, 0xeb, 0xbd, 0xc0, 0xe5, 0xc3,
	0xdc, 0x15, 0x1f, 0x8f, 0xce, 0x83, 0x79, 0x48, 0x9e, 0xfa, 0xa9, 0xf3, 0x3f, 0x24, 0xdd, 0xc0,
	0xee, 0xd1, 0xa7, 0x9a, 0x8f, 0xbb, 0xb7, 0xba, 0xc2, 0x98, 0xc1, 0x8f, 0x4b, 0x8f, 0x4b, 0xd0,
	0xdc, 0x48, 0x06, 0xa5, 0xcc, 0xa4, 0x30, 0xb8, 0x71, 0xe6, 0x24, 0x5e, 0xa8, 0xea, 0xe1, 0x7f,
	0xe5, 0xe2, 0xf6, 0xcd, 0x7a, 0x80, 0xb6, 0xf7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x67, 0x50,
	0x6e, 0x13, 0x61, 0x03, 0x00, 0x00,
}
