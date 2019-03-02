// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/dialogflow/v2beta1/webhook.proto

package dialogflow // import "github.com/Beeketing/go-genproto/googleapis/cloud/dialogflow/v2beta1"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/Beeketing/protobuf/ptypes/struct"
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
	// Alternative query results from KnowledgeService.
	AlternativeQueryResults []*QueryResult `protobuf:"bytes,5,rep,name=alternative_query_results,json=alternativeQueryResults,proto3" json:"alternative_query_results,omitempty"`
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
	return fileDescriptor_webhook_d5d0c4a657643b71, []int{0}
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

func (m *WebhookRequest) GetAlternativeQueryResults() []*QueryResult {
	if m != nil {
		return m.AlternativeQueryResults
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
	FollowupEventInput *EventInput `protobuf:"bytes,6,opt,name=followup_event_input,json=followupEventInput,proto3" json:"followup_event_input,omitempty"`
	// Optional. Indicates that this intent ends an interaction. Some integrations
	// (e.g., Actions on Google or Dialogflow phone gateway) use this information
	// to close interaction with an end user. Default is false.
	EndInteraction       bool     `protobuf:"varint,8,opt,name=end_interaction,json=endInteraction,proto3" json:"end_interaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebhookResponse) Reset()         { *m = WebhookResponse{} }
func (m *WebhookResponse) String() string { return proto.CompactTextString(m) }
func (*WebhookResponse) ProtoMessage()    {}
func (*WebhookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_webhook_d5d0c4a657643b71, []int{1}
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

func (m *WebhookResponse) GetEndInteraction() bool {
	if m != nil {
		return m.EndInteraction
	}
	return false
}

// Represents the contents of the original request that was passed to
// the `[Streaming]DetectIntent` call.
type OriginalDetectIntentRequest struct {
	// The source of this request, e.g., `google`, `facebook`, `slack`. It is set
	// by Dialogflow-owned servers.
	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// Optional. The version of the protocol used for this request.
	// This field is AoG-specific.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Optional. This field is set to the value of `QueryParameters.payload` field
	// passed in the request.
	//
	// This field is used for the telephony gateway. It should have a
	// structure similar to this JSON message:
	// <pre>{
	//  "telephony": {
	//    "caller_id": "+18558363987"
	//  }
	// }</pre>
	// Note: The caller ID field (`caller_id`) will be in
	// [E.164 format](https://en.wikipedia.org/wiki/E.164) and is only supported
	// for Enterprise Edition and not for Standard Edition agents. When the
	// telephony gateway is used with a standard tier agent the `caller_id` field
	// above will have a value of `REDACTED_IN_STANDARD_TIER_AGENT`.
	Payload              *_struct.Struct `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *OriginalDetectIntentRequest) Reset()         { *m = OriginalDetectIntentRequest{} }
func (m *OriginalDetectIntentRequest) String() string { return proto.CompactTextString(m) }
func (*OriginalDetectIntentRequest) ProtoMessage()    {}
func (*OriginalDetectIntentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_webhook_d5d0c4a657643b71, []int{2}
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

func (m *OriginalDetectIntentRequest) GetVersion() string {
	if m != nil {
		return m.Version
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
	proto.RegisterType((*WebhookRequest)(nil), "google.cloud.dialogflow.v2beta1.WebhookRequest")
	proto.RegisterType((*WebhookResponse)(nil), "google.cloud.dialogflow.v2beta1.WebhookResponse")
	proto.RegisterType((*OriginalDetectIntentRequest)(nil), "google.cloud.dialogflow.v2beta1.OriginalDetectIntentRequest")
}

func init() {
	proto.RegisterFile("google/cloud/dialogflow/v2beta1/webhook.proto", fileDescriptor_webhook_d5d0c4a657643b71)
}

var fileDescriptor_webhook_d5d0c4a657643b71 = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0x95, 0x5d, 0xe8, 0x1f, 0x6f, 0xd5, 0x45, 0xa6, 0xa2, 0xa1, 0x45, 0x74, 0x55, 0x0e,
	0x2c, 0xa2, 0x24, 0x6a, 0x39, 0xc2, 0xa9, 0x2d, 0xa0, 0x3d, 0xa0, 0xb6, 0x01, 0x81, 0x84, 0x84,
	0x22, 0x6f, 0x32, 0x9b, 0x46, 0xb8, 0x9e, 0x34, 0xb6, 0x77, 0xdb, 0x23, 0x1c, 0x79, 0x0c, 0x8e,
	0x7d, 0x10, 0x9e, 0x89, 0x23, 0x8a, 0xed, 0xb0, 0xe1, 0x50, 0x52, 0x8e, 0x33, 0x9e, 0xef, 0xe7,
	0xf1, 0x97, 0x99, 0x90, 0x67, 0x19, 0x62, 0xc6, 0x21, 0x4c, 0x38, 0xea, 0x34, 0x4c, 0x73, 0xc6,
	0x31, 0x9b, 0x70, 0x9c, 0x85, 0xd3, 0xbd, 0x31, 0x28, 0xb6, 0x1b, 0xce, 0x60, 0x7c, 0x8a, 0xf8,
	0x25, 0x28, 0x4a, 0x54, 0x48, 0xb7, 0x6c, 0x79, 0x60, 0xca, 0x83, 0x79, 0x79, 0xe0, 0xca, 0x37,
	0x1e, 0x38, 0x1e, 0x2b, 0xf2, 0x90, 0x09, 0x81, 0x8a, 0xa9, 0x1c, 0x85, 0xb4, 0xf2, 0x8d, 0xd6,
	0xdb, 0x12, 0x14, 0x0a, 0x2e, 0x94, 0x2b, 0xdf, 0x69, 0x2b, 0xcf, 0x85, 0x02, 0xa1, 0x6e, 0x0a,
	0x97, 0x20, 0x65, 0x8e, 0xc2, 0x95, 0xd7, 0x9d, 0x9a, 0x68, 0xac, 0x27, 0xa1, 0x54, 0xa5, 0x4e,
	0x1c, 0x6c, 0xfb, 0x7b, 0x97, 0xac, 0x7e, 0xb4, 0x4f, 0x8f, 0xe0, 0x5c, 0x83, 0x54, 0xd4, 0x27,
	0x8b, 0x8e, 0xe0, 0xdf, 0x1a, 0x78, 0xc3, 0xe5, 0xa8, 0x0e, 0xe9, 0x16, 0xe9, 0x95, 0x20, 0x0b,
	0x14, 0x12, 0xe2, 0x3c, 0xf5, 0x3d, 0x73, 0x4a, 0xea, 0xd4, 0x28, 0xa5, 0x47, 0x64, 0xe5, 0x5c,
	0x43, 0x79, 0x19, 0x97, 0x20, 0x35, 0x57, 0x7e, 0x67, 0xe0, 0x0d, 0x7b, 0x7b, 0x3b, 0x41, 0x8b,
	0x9b, 0xc1, 0x49, 0x25, 0x8a, 0x8c, 0x26, 0xea, 0x9d, 0xcf, 0x03, 0x7a, 0x4a, 0xee, 0x33, 0xae,
	0xa0, 0x14, 0x4c, 0xe5, 0x53, 0x88, 0x9b, 0x70, 0xe9, 0xdf, 0x1e, 0x74, 0xff, 0x9b, 0xbe, 0xde,
	0xc0, 0x35, 0xf2, 0x92, 0x7e, 0xf5, 0xc8, 0x43, 0x2c, 0xf3, 0x2c, 0x17, 0x8c, 0xc7, 0x29, 0x28,
	0x48, 0x54, 0x6c, 0x6d, 0x8f, 0x4b, 0x6b, 0x8c, 0xdf, 0x35, 0xaf, 0x79, 0xd9, 0x7a, 0xdf, 0x91,
	0xc3, 0x1c, 0x1a, 0xca, 0xc8, 0x40, 0x9c, 0xb9, 0xd1, 0x26, 0x5e, 0x7f, 0xb8, 0xfd, 0xb3, 0x4b,
	0xfa, 0x7f, 0x3e, 0x86, 0x35, 0x95, 0x3e, 0x21, 0x77, 0x26, 0x9a, 0x4f, 0x72, 0xce, 0xcf, 0xaa,
	0x5e, 0xaa, 0xa9, 0x71, 0xc6, 0xf7, 0x1b, 0xf9, 0xf7, 0x70, 0xa1, 0xe8, 0x98, 0xac, 0x35, 0x4b,
	0xcf, 0x40, 0x4a, 0x96, 0x81, 0xf4, 0x3b, 0xc6, 0xa7, 0xb0, 0xb5, 0x6f, 0xdb, 0x4c, 0xf0, 0xd6,
	0xea, 0xa2, 0xbb, 0x0d, 0x98, 0xcb, 0x49, 0x7a, 0x8f, 0x2c, 0x48, 0xd4, 0x65, 0x02, 0xc6, 0x8d,
	0xe5, 0xc8, 0x45, 0x74, 0x97, 0x2c, 0x16, 0xec, 0x92, 0x23, 0x4b, 0xcd, 0xd0, 0xf4, 0xf6, 0xd6,
	0xeb, 0xeb, 0xea, 0xb9, 0x0b, 0xde, 0x99, 0xb9, 0x8b, 0xea, 0x3a, 0x7a, 0x42, 0xfa, 0xa8, 0x55,
	0xa1, 0x55, 0xec, 0xb6, 0xa1, 0xfe, 0xa2, 0xc3, 0xd6, 0x4e, 0x0f, 0xac, 0x20, 0x5a, 0xb5, 0x00,
	0x17, 0x4a, 0xfa, 0x99, 0xac, 0x4d, 0x90, 0x73, 0x9c, 0xe9, 0x22, 0x86, 0x69, 0x65, 0x42, 0x2e,
	0x0a, 0xad, 0xfc, 0x05, 0xd3, 0xd2, 0xd3, 0x56, 0xee, 0xab, 0x4a, 0x33, 0xaa, 0x24, 0x11, 0xad,
	0x41, 0xf3, 0x1c, 0x7d, 0x4c, 0xfa, 0x20, 0x52, 0x33, 0x16, 0x25, 0x4b, 0xaa, 0x85, 0xf7, 0x97,
	0x06, 0xde, 0x70, 0x29, 0x5a, 0x05, 0x91, 0x8e, 0xe6, 0xd9, 0xed, 0x6f, 0x1e, 0xd9, 0xfc, 0xc7,
	0x14, 0x34, 0x5c, 0xf4, 0xfe, 0x72, 0xd1, 0x27, 0x8b, 0x53, 0x28, 0xcd, 0xea, 0x75, 0xec, 0xea,
	0xb9, 0xb0, 0xe9, 0x6f, 0xf7, 0x66, 0xfe, 0xee, 0x5f, 0x79, 0xe4, 0x51, 0x82, 0x67, 0x6d, 0x8f,
	0xde, 0x5f, 0x71, 0x23, 0x77, 0x5c, 0x81, 0x8e, 0xbd, 0x4f, 0x23, 0x27, 0xc8, 0x90, 0x33, 0x91,
	0x05, 0x58, 0x66, 0x61, 0x06, 0xc2, 0x5c, 0x13, 0xda, 0x23, 0x56, 0xe4, 0xf2, 0xda, 0xdf, 0xcf,
	0x8b, 0x79, 0xea, 0x97, 0xe7, 0xfd, 0xe8, 0x74, 0x0e, 0x5f, 0x5f, 0x75, 0xb6, 0xde, 0x58, 0xe6,
	0x81, 0x69, 0xe2, 0x70, 0xde, 0xc4, 0x07, 0x2b, 0x1a, 0x2f, 0x18, 0xfe, 0xf3, 0xdf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x42, 0x28, 0x28, 0xef, 0xa8, 0x05, 0x00, 0x00,
}
