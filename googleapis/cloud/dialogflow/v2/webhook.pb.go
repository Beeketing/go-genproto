// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/dialogflow/v2/webhook.proto

package dialogflow // import "github.com/Beeketing/genproto/googleapis/cloud/dialogflow/v2"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/Beeketing/protobuf/ptypes/struct"
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

// The request message for a webhook call.
type WebhookRequest struct {
	// The unique identifier of detectIntent request session.
	// Can be used to identify end-user inside webhook implementation.
	// Format: `projects/<Project ID>/agent/sessions/<Session ID>`.
	Session string `protobuf:"bytes,4,opt,name=session,proto3" json:"session,omitempty"`
	// The unique identifier of the response. Contains the same value as
	// `[Streaming]DetectIntentResponse.response_id`.
	ResponseId string `protobuf:"bytes,1,opt,name=response_id,json=responseId,proto3" json:"response_id,omitempty"`
	// The result of the conversational query or event processing. Contains the
	// same value as `[Streaming]DetectIntentResponse.query_result`.
	QueryResult *QueryResult `protobuf:"bytes,2,opt,name=query_result,json=queryResult,proto3" json:"query_result,omitempty"`
	// Optional. The contents of the original request that was passed to
	// `[Streaming]DetectIntent` call.
	OriginalDetectIntentRequest *OriginalDetectIntentRequest `protobuf:"bytes,3,opt,name=original_detect_intent_request,json=originalDetectIntentRequest,proto3" json:"original_detect_intent_request,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}                     `json:"-"`
	XXX_unrecognized            []byte                       `json:"-"`
	XXX_sizecache               int32                        `json:"-"`
}

func (m *WebhookRequest) Reset()         { *m = WebhookRequest{} }
func (m *WebhookRequest) String() string { return proto.CompactTextString(m) }
func (*WebhookRequest) ProtoMessage()    {}
func (*WebhookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_webhook_d01678681840ffff, []int{0}
}
func (m *WebhookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebhookRequest.Unmarshal(m, b)
}
func (m *WebhookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebhookRequest.Marshal(b, m, deterministic)
}
func (dst *WebhookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebhookRequest.Merge(dst, src)
}
func (m *WebhookRequest) XXX_Size() int {
	return xxx_messageInfo_WebhookRequest.Size(m)
}
func (m *WebhookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WebhookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WebhookRequest proto.InternalMessageInfo

func (m *WebhookRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func (m *WebhookRequest) GetResponseId() string {
	if m != nil {
		return m.ResponseId
	}
	return ""
}

func (m *WebhookRequest) GetQueryResult() *QueryResult {
	if m != nil {
		return m.QueryResult
	}
	return nil
}

func (m *WebhookRequest) GetOriginalDetectIntentRequest() *OriginalDetectIntentRequest {
	if m != nil {
		return m.OriginalDetectIntentRequest
	}
	return nil
}

// The response message for a webhook call.
type WebhookResponse struct {
	// Optional. The text to be shown on the screen. This value is passed directly
	// to `QueryResult.fulfillment_text`.
	FulfillmentText string `protobuf:"bytes,1,opt,name=fulfillment_text,json=fulfillmentText,proto3" json:"fulfillment_text,omitempty"`
	// Optional. The collection of rich messages to present to the user. This
	// value is passed directly to `QueryResult.fulfillment_messages`.
	FulfillmentMessages []*Intent_Message `protobuf:"bytes,2,rep,name=fulfillment_messages,json=fulfillmentMessages,proto3" json:"fulfillment_messages,omitempty"`
	// Optional. This value is passed directly to `QueryResult.webhook_source`.
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	// Optional. This value is passed directly to `QueryResult.webhook_payload`.
	// See the related `fulfillment_messages[i].payload field`, which may be used
	// as an alternative to this field.
	//
	// This field can be used for Actions on Google responses.
	// It should have a structure similar to the JSON message shown here. For more
	// information, see
	// [Actions on Google Webhook
	// Format](https://developers.google.com/actions/dialogflow/webhook)
	// <pre>{
	//   "google": {
	//     "expectUserResponse": true,
	//     "richResponse": {
	//       "items": [
	//         {
	//           "simpleResponse": {
	//             "textToSpeech": "this is a simple response"
	//           }
	//         }
	//       ]
	//     }
	//   }
	// }</pre>
	Payload *_struct.Struct `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	// Optional. The collection of output contexts. This value is passed directly
	// to `QueryResult.output_contexts`.
	OutputContexts []*Context `protobuf:"bytes,5,rep,name=output_contexts,json=outputContexts,proto3" json:"output_contexts,omitempty"`
	// Optional. Makes the platform immediately invoke another `DetectIntent` call
	// internally with the specified event as input.
	FollowupEventInput   *EventInput `protobuf:"bytes,6,opt,name=followup_event_input,json=followupEventInput,proto3" json:"followup_event_input,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *WebhookResponse) Reset()         { *m = WebhookResponse{} }
func (m *WebhookResponse) String() string { return proto.CompactTextString(m) }
func (*WebhookResponse) ProtoMessage()    {}
func (*WebhookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_webhook_d01678681840ffff, []int{1}
}
func (m *WebhookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebhookResponse.Unmarshal(m, b)
}
func (m *WebhookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebhookResponse.Marshal(b, m, deterministic)
}
func (dst *WebhookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebhookResponse.Merge(dst, src)
}
func (m *WebhookResponse) XXX_Size() int {
	return xxx_messageInfo_WebhookResponse.Size(m)
}
func (m *WebhookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WebhookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WebhookResponse proto.InternalMessageInfo

func (m *WebhookResponse) GetFulfillmentText() string {
	if m != nil {
		return m.FulfillmentText
	}
	return ""
}

func (m *WebhookResponse) GetFulfillmentMessages() []*Intent_Message {
	if m != nil {
		return m.FulfillmentMessages
	}
	return nil
}

func (m *WebhookResponse) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *WebhookResponse) GetPayload() *_struct.Struct {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *WebhookResponse) GetOutputContexts() []*Context {
	if m != nil {
		return m.OutputContexts
	}
	return nil
}

func (m *WebhookResponse) GetFollowupEventInput() *EventInput {
	if m != nil {
		return m.FollowupEventInput
	}
	return nil
}

// Represents the contents of the original request that was passed to
// the `[Streaming]DetectIntent` call.
type OriginalDetectIntentRequest struct {
	// The source of this request, e.g., `google`, `facebook`, `slack`. It is set
	// by Dialogflow-owned servers.
	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// Optional. This field is set to the value of `QueryParameters.payload` field
	// passed in the request.
	Payload              *_struct.Struct `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *OriginalDetectIntentRequest) Reset()         { *m = OriginalDetectIntentRequest{} }
func (m *OriginalDetectIntentRequest) String() string { return proto.CompactTextString(m) }
func (*OriginalDetectIntentRequest) ProtoMessage()    {}
func (*OriginalDetectIntentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_webhook_d01678681840ffff, []int{2}
}
func (m *OriginalDetectIntentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OriginalDetectIntentRequest.Unmarshal(m, b)
}
func (m *OriginalDetectIntentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OriginalDetectIntentRequest.Marshal(b, m, deterministic)
}
func (dst *OriginalDetectIntentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OriginalDetectIntentRequest.Merge(dst, src)
}
func (m *OriginalDetectIntentRequest) XXX_Size() int {
	return xxx_messageInfo_OriginalDetectIntentRequest.Size(m)
}
func (m *OriginalDetectIntentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OriginalDetectIntentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OriginalDetectIntentRequest proto.InternalMessageInfo

func (m *OriginalDetectIntentRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *OriginalDetectIntentRequest) GetPayload() *_struct.Struct {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*WebhookRequest)(nil), "google.cloud.dialogflow.v2.WebhookRequest")
	proto.RegisterType((*WebhookResponse)(nil), "google.cloud.dialogflow.v2.WebhookResponse")
	proto.RegisterType((*OriginalDetectIntentRequest)(nil), "google.cloud.dialogflow.v2.OriginalDetectIntentRequest")
}

func init() {
	proto.RegisterFile("google/cloud/dialogflow/v2/webhook.proto", fileDescriptor_webhook_d01678681840ffff)
}

var fileDescriptor_webhook_d01678681840ffff = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x95, 0x14, 0x3a, 0xcd, 0x9d, 0x56, 0x64, 0x26, 0x88, 0x3a, 0x34, 0xaa, 0x22, 0xb1,
	0xc2, 0x21, 0x11, 0xe1, 0xc0, 0x81, 0xdb, 0x56, 0x40, 0x45, 0x20, 0x46, 0x40, 0x80, 0x90, 0x50,
	0x94, 0x26, 0x6e, 0x66, 0xe1, 0xfa, 0xa5, 0xb1, 0xdd, 0x6e, 0x12, 0x7f, 0x03, 0x37, 0x4e, 0xdc,
	0x38, 0xf2, 0x17, 0x72, 0x44, 0xb1, 0x1d, 0xd2, 0x03, 0x8d, 0x76, 0x7c, 0xf6, 0xe7, 0xdf, 0xfb,
	0xde, 0xf7, 0x12, 0x34, 0xce, 0x01, 0x72, 0x46, 0x82, 0x94, 0x81, 0xca, 0x82, 0x8c, 0x26, 0x0c,
	0xf2, 0x39, 0x83, 0x75, 0xb0, 0x0a, 0x83, 0x35, 0x99, 0x9d, 0x03, 0x7c, 0xf5, 0x8b, 0x12, 0x24,
	0xe0, 0x81, 0x51, 0xfa, 0x5a, 0xe9, 0x37, 0x4a, 0x7f, 0x15, 0x0e, 0xee, 0x58, 0x4a, 0x52, 0xd0,
	0x20, 0xe1, 0x1c, 0x64, 0x22, 0x29, 0x70, 0x61, 0x5e, 0x0e, 0xda, 0x7a, 0xa4, 0xc0, 0x25, 0xb9,
	0x90, 0x56, 0x79, 0xdc, 0xa2, 0xa4, 0x5c, 0x12, 0x2e, 0xaf, 0x80, 0x14, 0x44, 0x08, 0x0a, 0xdc,
	0x2a, 0x6b, 0x6b, 0xba, 0x9a, 0xa9, 0x79, 0x20, 0x64, 0xa9, 0x52, 0xcb, 0x19, 0xfd, 0x70, 0xd1,
	0xfe, 0x47, 0x33, 0x66, 0x44, 0x96, 0x8a, 0x08, 0x89, 0x3d, 0xb4, 0x63, 0x09, 0xde, 0xb5, 0xa1,
	0x33, 0xde, 0x8d, 0xea, 0x12, 0xdf, 0x45, 0xbd, 0x92, 0x88, 0x02, 0xb8, 0x20, 0x31, 0xcd, 0x3c,
	0x47, 0xdf, 0xa2, 0xfa, 0x68, 0x9a, 0xe1, 0x97, 0x68, 0x6f, 0xa9, 0x48, 0x79, 0x19, 0x97, 0x44,
	0x28, 0x26, 0x3d, 0x77, 0xe8, 0x8c, 0x7b, 0xe1, 0xb1, 0xbf, 0x3d, 0x39, 0xff, 0x6d, 0xa5, 0x8f,
	0xb4, 0x3c, 0xea, 0x2d, 0x9b, 0x02, 0x7f, 0x43, 0x47, 0x50, 0xd2, 0x9c, 0xf2, 0x84, 0xc5, 0x19,
	0x91, 0x24, 0x95, 0xb1, 0x49, 0x20, 0x2e, 0x8d, 0x51, 0xaf, 0xa3, 0xe9, 0x4f, 0xda, 0xe8, 0x6f,
	0x2c, 0x61, 0xa2, 0x01, 0x53, 0xfd, 0xde, 0xce, 0x19, 0x1d, 0xc2, 0xf6, 0xcb, 0xd1, 0xf7, 0x0e,
	0xea, 0xff, 0xcb, 0xc5, 0xcc, 0x87, 0x1f, 0xa0, 0x1b, 0x73, 0xc5, 0xe6, 0x94, 0xb1, 0x45, 0x65,
	0xa3, 0x5a, 0x9b, 0xcd, 0xa0, 0xbf, 0x71, 0xfe, 0x9e, 0x5c, 0x48, 0xfc, 0x05, 0x1d, 0x6c, 0x4a,
	0x17, 0x44, 0x88, 0x24, 0x27, 0xc2, 0x73, 0x87, 0x9d, 0x71, 0x2f, 0x7c, 0xd8, 0x66, 0xd9, 0xf8,
	0xf0, 0x5f, 0x9b, 0x27, 0xd1, 0xcd, 0x0d, 0x8e, 0x3d, 0x13, 0xf8, 0x16, 0xea, 0x0a, 0x50, 0x65,
	0x4a, 0x74, 0x06, 0xbb, 0x91, 0xad, 0xf0, 0x23, 0xb4, 0x53, 0x24, 0x97, 0x0c, 0x92, 0x4c, 0xaf,
	0xae, 0x17, 0xde, 0xae, 0x3b, 0xd5, 0xdb, 0xf7, 0xdf, 0xe9, 0xed, 0x47, 0xb5, 0x0e, 0xbf, 0x42,
	0x7d, 0x50, 0xb2, 0x50, 0x32, 0xb6, 0x5f, 0xa2, 0xf0, 0xae, 0x6b, 0x93, 0xf7, 0xda, 0x4c, 0x9e,
	0x1a, 0x6d, 0xb4, 0x6f, 0xde, 0xda, 0x52, 0xe0, 0x4f, 0xe8, 0x60, 0x0e, 0x8c, 0xc1, 0x5a, 0x15,
	0x31, 0x59, 0x55, 0xa3, 0x53, 0x5e, 0x28, 0xe9, 0x75, 0xb5, 0x9b, 0xfb, 0x6d, 0xc8, 0x67, 0x95,
	0x7c, 0x5a, 0xa9, 0x23, 0x5c, 0x33, 0x9a, 0xb3, 0xd1, 0x39, 0x3a, 0x6c, 0x59, 0xe6, 0x46, 0x22,
	0xce, 0xb6, 0x44, 0x3a, 0x57, 0x4b, 0xe4, 0xe4, 0xa7, 0x83, 0x8e, 0x52, 0x58, 0xb4, 0x78, 0x3d,
	0xd9, 0xb3, 0x9f, 0xc6, 0x59, 0xc5, 0x38, 0x73, 0x3e, 0x4f, 0xac, 0x36, 0x07, 0x96, 0xf0, 0xdc,
	0x87, 0x32, 0x0f, 0x72, 0xc2, 0x75, 0x87, 0xc0, 0x5c, 0x25, 0x05, 0x15, 0xff, 0xfb, 0x59, 0x9f,
	0x36, 0xd5, 0x1f, 0xc7, 0xf9, 0xe5, 0xba, 0x93, 0xe7, 0xbf, 0xdd, 0xc1, 0x0b, 0x83, 0x3b, 0xd5,
	0xad, 0x27, 0x4d, 0xeb, 0x0f, 0xe1, 0xac, 0xab, 0xa9, 0x8f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x75, 0x98, 0x13, 0x07, 0xb8, 0x04, 0x00, 0x00,
}
