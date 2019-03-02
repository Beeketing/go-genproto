// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/dialogflow/v2beta1/document.proto

package dialogflow // import "github.com/Beeketing/genproto/googleapis/cloud/dialogflow/v2beta1"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/Beeketing/protobuf/ptypes/empty"
import _ "github.com/Beeketing/genproto/googleapis/api/annotations"
import longrunning "github.com/Beeketing/genproto/googleapis/longrunning"
import _ "github.com/Beeketing/genproto/protobuf/field_mask"

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

// The knowledge type of document content.
type Document_KnowledgeType int32

const (
	// The type is unspecified or arbitrary.
	Document_KNOWLEDGE_TYPE_UNSPECIFIED Document_KnowledgeType = 0
	// The document content contains question and answer pairs as either HTML or
	// CSV. Typical FAQ HTML formats are parsed accurately, but unusual formats
	// may fail to be parsed.
	//
	// CSV must have questions in the first column and answers in the second,
	// with no header. Because of this explicit format, they are always parsed
	// accurately.
	Document_FAQ Document_KnowledgeType = 1
	// Documents for which unstructured text is extracted and used for
	// question answering.
	Document_EXTRACTIVE_QA Document_KnowledgeType = 2
)

var Document_KnowledgeType_name = map[int32]string{
	0: "KNOWLEDGE_TYPE_UNSPECIFIED",
	1: "FAQ",
	2: "EXTRACTIVE_QA",
}
var Document_KnowledgeType_value = map[string]int32{
	"KNOWLEDGE_TYPE_UNSPECIFIED": 0,
	"FAQ":                        1,
	"EXTRACTIVE_QA":              2,
}

func (x Document_KnowledgeType) String() string {
	return proto.EnumName(Document_KnowledgeType_name, int32(x))
}
func (Document_KnowledgeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_document_3131c12790f8bf44, []int{0, 0}
}

// States of the operation.
type KnowledgeOperationMetadata_State int32

const (
	// State unspecified.
	KnowledgeOperationMetadata_STATE_UNSPECIFIED KnowledgeOperationMetadata_State = 0
	// The operation has been created.
	KnowledgeOperationMetadata_PENDING KnowledgeOperationMetadata_State = 1
	// The operation is currently running.
	KnowledgeOperationMetadata_RUNNING KnowledgeOperationMetadata_State = 2
	// The operation is done, either cancelled or completed.
	KnowledgeOperationMetadata_DONE KnowledgeOperationMetadata_State = 3
)

var KnowledgeOperationMetadata_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "PENDING",
	2: "RUNNING",
	3: "DONE",
}
var KnowledgeOperationMetadata_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"PENDING":           1,
	"RUNNING":           2,
	"DONE":              3,
}

func (x KnowledgeOperationMetadata_State) String() string {
	return proto.EnumName(KnowledgeOperationMetadata_State_name, int32(x))
}
func (KnowledgeOperationMetadata_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_document_3131c12790f8bf44, []int{6, 0}
}

// A document resource.
type Document struct {
	// The document resource name.
	// The name must be empty when creating a document.
	// Format: `projects/<Project ID>/knowledgeBases/<Knowledge Base
	// ID>/documents/<Document ID>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The display name of the document. The name must be 1024 bytes or
	// less; otherwise, the creation request fails.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Required. The MIME type of this document.
	MimeType string `protobuf:"bytes,3,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// Required. The knowledge type of document content.
	KnowledgeTypes []Document_KnowledgeType `protobuf:"varint,4,rep,packed,name=knowledge_types,json=knowledgeTypes,proto3,enum=google.cloud.dialogflow.v2beta1.Document_KnowledgeType" json:"knowledge_types,omitempty"`
	// Required. The source of this document.
	//
	// Types that are valid to be assigned to Source:
	//	*Document_ContentUri
	//	*Document_Content
	Source               isDocument_Source `protobuf_oneof:"source"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Document) Reset()         { *m = Document{} }
func (m *Document) String() string { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()    {}
func (*Document) Descriptor() ([]byte, []int) {
	return fileDescriptor_document_3131c12790f8bf44, []int{0}
}
func (m *Document) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Document.Unmarshal(m, b)
}
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Document.Marshal(b, m, deterministic)
}
func (dst *Document) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document.Merge(dst, src)
}
func (m *Document) XXX_Size() int {
	return xxx_messageInfo_Document.Size(m)
}
func (m *Document) XXX_DiscardUnknown() {
	xxx_messageInfo_Document.DiscardUnknown(m)
}

var xxx_messageInfo_Document proto.InternalMessageInfo

func (m *Document) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Document) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Document) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *Document) GetKnowledgeTypes() []Document_KnowledgeType {
	if m != nil {
		return m.KnowledgeTypes
	}
	return nil
}

type isDocument_Source interface {
	isDocument_Source()
}

type Document_ContentUri struct {
	ContentUri string `protobuf:"bytes,5,opt,name=content_uri,json=contentUri,proto3,oneof"`
}

type Document_Content struct {
	Content string `protobuf:"bytes,6,opt,name=content,proto3,oneof"`
}

func (*Document_ContentUri) isDocument_Source() {}

func (*Document_Content) isDocument_Source() {}

func (m *Document) GetSource() isDocument_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Document) GetContentUri() string {
	if x, ok := m.GetSource().(*Document_ContentUri); ok {
		return x.ContentUri
	}
	return ""
}

func (m *Document) GetContent() string {
	if x, ok := m.GetSource().(*Document_Content); ok {
		return x.Content
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Document) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Document_OneofMarshaler, _Document_OneofUnmarshaler, _Document_OneofSizer, []interface{}{
		(*Document_ContentUri)(nil),
		(*Document_Content)(nil),
	}
}

func _Document_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Document)
	// source
	switch x := m.Source.(type) {
	case *Document_ContentUri:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ContentUri)
	case *Document_Content:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Content)
	case nil:
	default:
		return fmt.Errorf("Document.Source has unexpected type %T", x)
	}
	return nil
}

func _Document_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Document)
	switch tag {
	case 5: // source.content_uri
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Source = &Document_ContentUri{x}
		return true, err
	case 6: // source.content
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Source = &Document_Content{x}
		return true, err
	default:
		return false, nil
	}
}

func _Document_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Document)
	// source
	switch x := m.Source.(type) {
	case *Document_ContentUri:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.ContentUri)))
		n += len(x.ContentUri)
	case *Document_Content:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Content)))
		n += len(x.Content)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Request message for [Documents.ListDocuments][google.cloud.dialogflow.v2beta1.Documents.ListDocuments].
type ListDocumentsRequest struct {
	// Required. The knowledge base to list all documents for.
	// Format: `projects/<Project ID>/knowledgeBases/<Knowledge Base ID>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. The maximum number of items to return in a single page. By
	// default 10 and at most 100.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The next_page_token value returned from a previous list request.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDocumentsRequest) Reset()         { *m = ListDocumentsRequest{} }
func (m *ListDocumentsRequest) String() string { return proto.CompactTextString(m) }
func (*ListDocumentsRequest) ProtoMessage()    {}
func (*ListDocumentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_document_3131c12790f8bf44, []int{1}
}
func (m *ListDocumentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDocumentsRequest.Unmarshal(m, b)
}
func (m *ListDocumentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDocumentsRequest.Marshal(b, m, deterministic)
}
func (dst *ListDocumentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDocumentsRequest.Merge(dst, src)
}
func (m *ListDocumentsRequest) XXX_Size() int {
	return xxx_messageInfo_ListDocumentsRequest.Size(m)
}
func (m *ListDocumentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDocumentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDocumentsRequest proto.InternalMessageInfo

func (m *ListDocumentsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListDocumentsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListDocumentsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Response message for [Documents.ListDocuments][google.cloud.dialogflow.v2beta1.Documents.ListDocuments].
type ListDocumentsResponse struct {
	// The list of documents.
	Documents []*Document `protobuf:"bytes,1,rep,name=documents,proto3" json:"documents,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDocumentsResponse) Reset()         { *m = ListDocumentsResponse{} }
func (m *ListDocumentsResponse) String() string { return proto.CompactTextString(m) }
func (*ListDocumentsResponse) ProtoMessage()    {}
func (*ListDocumentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_document_3131c12790f8bf44, []int{2}
}
func (m *ListDocumentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDocumentsResponse.Unmarshal(m, b)
}
func (m *ListDocumentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDocumentsResponse.Marshal(b, m, deterministic)
}
func (dst *ListDocumentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDocumentsResponse.Merge(dst, src)
}
func (m *ListDocumentsResponse) XXX_Size() int {
	return xxx_messageInfo_ListDocumentsResponse.Size(m)
}
func (m *ListDocumentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDocumentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDocumentsResponse proto.InternalMessageInfo

func (m *ListDocumentsResponse) GetDocuments() []*Document {
	if m != nil {
		return m.Documents
	}
	return nil
}

func (m *ListDocumentsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request message for [Documents.GetDocument][google.cloud.dialogflow.v2beta1.Documents.GetDocument].
type GetDocumentRequest struct {
	// Required. The name of the document to retrieve.
	// Format `projects/<Project ID>/knowledgeBases/<Knowledge Base
	// ID>/documents/<Document ID>`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDocumentRequest) Reset()         { *m = GetDocumentRequest{} }
func (m *GetDocumentRequest) String() string { return proto.CompactTextString(m) }
func (*GetDocumentRequest) ProtoMessage()    {}
func (*GetDocumentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_document_3131c12790f8bf44, []int{3}
}
func (m *GetDocumentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDocumentRequest.Unmarshal(m, b)
}
func (m *GetDocumentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDocumentRequest.Marshal(b, m, deterministic)
}
func (dst *GetDocumentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDocumentRequest.Merge(dst, src)
}
func (m *GetDocumentRequest) XXX_Size() int {
	return xxx_messageInfo_GetDocumentRequest.Size(m)
}
func (m *GetDocumentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDocumentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDocumentRequest proto.InternalMessageInfo

func (m *GetDocumentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for [Documents.CreateDocument][google.cloud.dialogflow.v2beta1.Documents.CreateDocument].
type CreateDocumentRequest struct {
	// Required. The knoweldge base to create a document for.
	// Format: `projects/<Project ID>/knowledgeBases/<Knowledge Base ID>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The document to create.
	Document             *Document `protobuf:"bytes,2,opt,name=document,proto3" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateDocumentRequest) Reset()         { *m = CreateDocumentRequest{} }
func (m *CreateDocumentRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDocumentRequest) ProtoMessage()    {}
func (*CreateDocumentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_document_3131c12790f8bf44, []int{4}
}
func (m *CreateDocumentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDocumentRequest.Unmarshal(m, b)
}
func (m *CreateDocumentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDocumentRequest.Marshal(b, m, deterministic)
}
func (dst *CreateDocumentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDocumentRequest.Merge(dst, src)
}
func (m *CreateDocumentRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDocumentRequest.Size(m)
}
func (m *CreateDocumentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDocumentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDocumentRequest proto.InternalMessageInfo

func (m *CreateDocumentRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateDocumentRequest) GetDocument() *Document {
	if m != nil {
		return m.Document
	}
	return nil
}

// Request message for [Documents.DeleteDocument][google.cloud.dialogflow.v2beta1.Documents.DeleteDocument].
type DeleteDocumentRequest struct {
	// The name of the document to delete.
	// Format: `projects/<Project ID>/knowledgeBases/<Knowledge Base
	// ID>/documents/<Document ID>`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDocumentRequest) Reset()         { *m = DeleteDocumentRequest{} }
func (m *DeleteDocumentRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDocumentRequest) ProtoMessage()    {}
func (*DeleteDocumentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_document_3131c12790f8bf44, []int{5}
}
func (m *DeleteDocumentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDocumentRequest.Unmarshal(m, b)
}
func (m *DeleteDocumentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDocumentRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteDocumentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDocumentRequest.Merge(dst, src)
}
func (m *DeleteDocumentRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDocumentRequest.Size(m)
}
func (m *DeleteDocumentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDocumentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDocumentRequest proto.InternalMessageInfo

func (m *DeleteDocumentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Metadata in google::longrunning::Operation for Knowledge operations.
type KnowledgeOperationMetadata struct {
	// Required. The current state of this operation.
	State                KnowledgeOperationMetadata_State `protobuf:"varint,1,opt,name=state,proto3,enum=google.cloud.dialogflow.v2beta1.KnowledgeOperationMetadata_State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *KnowledgeOperationMetadata) Reset()         { *m = KnowledgeOperationMetadata{} }
func (m *KnowledgeOperationMetadata) String() string { return proto.CompactTextString(m) }
func (*KnowledgeOperationMetadata) ProtoMessage()    {}
func (*KnowledgeOperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_document_3131c12790f8bf44, []int{6}
}
func (m *KnowledgeOperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KnowledgeOperationMetadata.Unmarshal(m, b)
}
func (m *KnowledgeOperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KnowledgeOperationMetadata.Marshal(b, m, deterministic)
}
func (dst *KnowledgeOperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KnowledgeOperationMetadata.Merge(dst, src)
}
func (m *KnowledgeOperationMetadata) XXX_Size() int {
	return xxx_messageInfo_KnowledgeOperationMetadata.Size(m)
}
func (m *KnowledgeOperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_KnowledgeOperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_KnowledgeOperationMetadata proto.InternalMessageInfo

func (m *KnowledgeOperationMetadata) GetState() KnowledgeOperationMetadata_State {
	if m != nil {
		return m.State
	}
	return KnowledgeOperationMetadata_STATE_UNSPECIFIED
}

func init() {
	proto.RegisterType((*Document)(nil), "google.cloud.dialogflow.v2beta1.Document")
	proto.RegisterType((*ListDocumentsRequest)(nil), "google.cloud.dialogflow.v2beta1.ListDocumentsRequest")
	proto.RegisterType((*ListDocumentsResponse)(nil), "google.cloud.dialogflow.v2beta1.ListDocumentsResponse")
	proto.RegisterType((*GetDocumentRequest)(nil), "google.cloud.dialogflow.v2beta1.GetDocumentRequest")
	proto.RegisterType((*CreateDocumentRequest)(nil), "google.cloud.dialogflow.v2beta1.CreateDocumentRequest")
	proto.RegisterType((*DeleteDocumentRequest)(nil), "google.cloud.dialogflow.v2beta1.DeleteDocumentRequest")
	proto.RegisterType((*KnowledgeOperationMetadata)(nil), "google.cloud.dialogflow.v2beta1.KnowledgeOperationMetadata")
	proto.RegisterEnum("google.cloud.dialogflow.v2beta1.Document_KnowledgeType", Document_KnowledgeType_name, Document_KnowledgeType_value)
	proto.RegisterEnum("google.cloud.dialogflow.v2beta1.KnowledgeOperationMetadata_State", KnowledgeOperationMetadata_State_name, KnowledgeOperationMetadata_State_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DocumentsClient is the client API for Documents service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DocumentsClient interface {
	// Returns the list of all documents of the knowledge base.
	ListDocuments(ctx context.Context, in *ListDocumentsRequest, opts ...grpc.CallOption) (*ListDocumentsResponse, error)
	// Retrieves the specified document.
	GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*Document, error)
	// Creates a new document.
	//
	// Operation <response: [Document][google.cloud.dialogflow.v2beta1.Document],
	//            metadata: [KnowledgeOperationMetadata][google.cloud.dialogflow.v2beta1.KnowledgeOperationMetadata]>
	CreateDocument(ctx context.Context, in *CreateDocumentRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Deletes the specified document.
	//
	// Operation <response: [google.protobuf.Empty][google.protobuf.Empty],
	//            metadata: [KnowledgeOperationMetadata][google.cloud.dialogflow.v2beta1.KnowledgeOperationMetadata]>
	DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
}

type documentsClient struct {
	cc *grpc.ClientConn
}

func NewDocumentsClient(cc *grpc.ClientConn) DocumentsClient {
	return &documentsClient{cc}
}

func (c *documentsClient) ListDocuments(ctx context.Context, in *ListDocumentsRequest, opts ...grpc.CallOption) (*ListDocumentsResponse, error) {
	out := new(ListDocumentsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.Documents/ListDocuments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentsClient) GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*Document, error) {
	out := new(Document)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.Documents/GetDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentsClient) CreateDocument(ctx context.Context, in *CreateDocumentRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.Documents/CreateDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentsClient) DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.Documents/DeleteDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentsServer is the server API for Documents service.
type DocumentsServer interface {
	// Returns the list of all documents of the knowledge base.
	ListDocuments(context.Context, *ListDocumentsRequest) (*ListDocumentsResponse, error)
	// Retrieves the specified document.
	GetDocument(context.Context, *GetDocumentRequest) (*Document, error)
	// Creates a new document.
	//
	// Operation <response: [Document][google.cloud.dialogflow.v2beta1.Document],
	//            metadata: [KnowledgeOperationMetadata][google.cloud.dialogflow.v2beta1.KnowledgeOperationMetadata]>
	CreateDocument(context.Context, *CreateDocumentRequest) (*longrunning.Operation, error)
	// Deletes the specified document.
	//
	// Operation <response: [google.protobuf.Empty][google.protobuf.Empty],
	//            metadata: [KnowledgeOperationMetadata][google.cloud.dialogflow.v2beta1.KnowledgeOperationMetadata]>
	DeleteDocument(context.Context, *DeleteDocumentRequest) (*longrunning.Operation, error)
}

func RegisterDocumentsServer(s *grpc.Server, srv DocumentsServer) {
	s.RegisterService(&_Documents_serviceDesc, srv)
}

func _Documents_ListDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDocumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentsServer).ListDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.Documents/ListDocuments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentsServer).ListDocuments(ctx, req.(*ListDocumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Documents_GetDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentsServer).GetDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.Documents/GetDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentsServer).GetDocument(ctx, req.(*GetDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Documents_CreateDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentsServer).CreateDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.Documents/CreateDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentsServer).CreateDocument(ctx, req.(*CreateDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Documents_DeleteDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentsServer).DeleteDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.Documents/DeleteDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentsServer).DeleteDocument(ctx, req.(*DeleteDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Documents_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.dialogflow.v2beta1.Documents",
	HandlerType: (*DocumentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDocuments",
			Handler:    _Documents_ListDocuments_Handler,
		},
		{
			MethodName: "GetDocument",
			Handler:    _Documents_GetDocument_Handler,
		},
		{
			MethodName: "CreateDocument",
			Handler:    _Documents_CreateDocument_Handler,
		},
		{
			MethodName: "DeleteDocument",
			Handler:    _Documents_DeleteDocument_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/dialogflow/v2beta1/document.proto",
}

func init() {
	proto.RegisterFile("google/cloud/dialogflow/v2beta1/document.proto", fileDescriptor_document_3131c12790f8bf44)
}

var fileDescriptor_document_3131c12790f8bf44 = []byte{
	// 885 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xcd, 0x6e, 0x1b, 0x55,
	0x14, 0xee, 0x1d, 0xe7, 0xc7, 0x39, 0xc6, 0xae, 0x7b, 0x45, 0x90, 0xe5, 0x52, 0x9a, 0x4e, 0x25,
	0x14, 0x8c, 0x34, 0x23, 0x5c, 0x41, 0x25, 0x50, 0x85, 0xec, 0xcc, 0xc4, 0x58, 0x29, 0x13, 0x77,
	0xe2, 0xb4, 0x90, 0xcd, 0x70, 0x63, 0xdf, 0x8c, 0xa6, 0x19, 0xdf, 0x3b, 0xcc, 0xbd, 0x6e, 0x49,
	0x50, 0x25, 0x60, 0xc5, 0x8a, 0x0d, 0xe2, 0x05, 0x58, 0xc2, 0x13, 0xf0, 0x1c, 0x48, 0x3c, 0x01,
	0x2b, 0x56, 0xb0, 0x63, 0x07, 0x9a, 0x5f, 0xc7, 0xae, 0x13, 0xa7, 0xed, 0xce, 0xe7, 0xef, 0x9b,
	0xef, 0x9c, 0xf3, 0x9d, 0xf1, 0x80, 0xe6, 0x72, 0xee, 0xfa, 0x54, 0x1f, 0xf8, 0x7c, 0x3c, 0xd4,
	0x87, 0x1e, 0xf1, 0xb9, 0x7b, 0xe4, 0xf3, 0xa7, 0xfa, 0x93, 0xe6, 0x21, 0x95, 0xe4, 0x3d, 0x7d,
	0xc8, 0x07, 0xe3, 0x11, 0x65, 0x52, 0x0b, 0x42, 0x2e, 0x39, 0xbe, 0x99, 0xe4, 0x6b, 0x71, 0xbe,
	0x36, 0xc9, 0xd7, 0xd2, 0xfc, 0xfa, 0x9b, 0x29, 0x20, 0x09, 0x3c, 0x9d, 0x30, 0xc6, 0x25, 0x91,
	0x1e, 0x67, 0x22, 0x29, 0xaf, 0xdf, 0x4e, 0xa3, 0x3e, 0x67, 0x6e, 0x38, 0x66, 0xcc, 0x63, 0xae,
	0xce, 0x03, 0x1a, 0x4e, 0x25, 0x5d, 0x4f, 0x93, 0x62, 0xeb, 0x70, 0x7c, 0xa4, 0xd3, 0x51, 0x20,
	0x4f, 0xd2, 0xe0, 0xc6, 0x6c, 0xf0, 0xc8, 0xa3, 0xfe, 0xd0, 0x19, 0x11, 0x71, 0x9c, 0x64, 0xa8,
	0x7f, 0x28, 0x50, 0x34, 0x52, 0xd6, 0x18, 0xc3, 0x12, 0x23, 0x23, 0x5a, 0x43, 0x1b, 0x68, 0x73,
	0xcd, 0x8e, 0x7f, 0xe3, 0x5b, 0xf0, 0xda, 0xd0, 0x13, 0x81, 0x4f, 0x4e, 0x9c, 0x38, 0xa6, 0xc4,
	0xb1, 0x52, 0xea, 0xb3, 0xa2, 0x94, 0xeb, 0xb0, 0x36, 0xf2, 0x46, 0xd4, 0x91, 0x27, 0x01, 0xad,
	0x15, 0xe2, 0x78, 0x31, 0x72, 0xf4, 0x4f, 0x02, 0x8a, 0xbf, 0x80, 0xab, 0xc7, 0x8c, 0x3f, 0xf5,
	0xe9, 0xd0, 0x4d, 0x32, 0x44, 0x6d, 0x69, 0xa3, 0xb0, 0x59, 0x69, 0xde, 0xd5, 0x16, 0x4c, 0x47,
	0xcb, 0x78, 0x69, 0x3b, 0x19, 0x40, 0x84, 0x68, 0x57, 0x8e, 0xcf, 0x9a, 0x02, 0xdf, 0x82, 0xd2,
	0x80, 0x33, 0x49, 0x99, 0x74, 0xc6, 0xa1, 0x57, 0x5b, 0x8e, 0x08, 0x7c, 0x72, 0xc5, 0x86, 0xd4,
	0xb9, 0x1f, 0x7a, 0xb8, 0x0e, 0xab, 0xa9, 0x55, 0x5b, 0x49, 0xc3, 0x99, 0x43, 0xdd, 0x81, 0xf2,
	0x14, 0x3e, 0x7e, 0x0b, 0xea, 0x3b, 0xd6, 0xee, 0xa3, 0xfb, 0xa6, 0xd1, 0x31, 0x9d, 0xfe, 0xe7,
	0x3d, 0xd3, 0xd9, 0xb7, 0xf6, 0x7a, 0xe6, 0x56, 0x77, 0xbb, 0x6b, 0x1a, 0xd5, 0x2b, 0x78, 0x15,
	0x0a, 0xdb, 0xad, 0x07, 0x55, 0x84, 0xaf, 0x41, 0xd9, 0xfc, 0xac, 0x6f, 0xb7, 0xb6, 0xfa, 0xdd,
	0x87, 0xa6, 0xf3, 0xa0, 0x55, 0x55, 0xda, 0x45, 0x58, 0x11, 0x7c, 0x1c, 0x0e, 0xa8, 0xfa, 0x18,
	0x5e, 0xbf, 0xef, 0x09, 0x99, 0xf5, 0x20, 0x6c, 0xfa, 0xe5, 0x98, 0x0a, 0x89, 0xdf, 0x80, 0x95,
	0x80, 0x84, 0x11, 0x93, 0x64, 0xca, 0xa9, 0x15, 0x0d, 0x31, 0x20, 0x2e, 0x75, 0x84, 0x77, 0x9a,
	0x0c, 0x79, 0xd9, 0x2e, 0x46, 0x8e, 0x3d, 0xef, 0x94, 0xe2, 0x1b, 0x00, 0x71, 0x50, 0xf2, 0x63,
	0xca, 0xd2, 0x11, 0xc7, 0xe9, 0xfd, 0xc8, 0xa1, 0x7e, 0x8f, 0x60, 0x7d, 0xe6, 0x61, 0x22, 0xe0,
	0x4c, 0x50, 0xdc, 0x81, 0xb5, 0x4c, 0x93, 0xa2, 0x86, 0x36, 0x0a, 0x9b, 0xa5, 0xe6, 0x3b, 0x97,
	0x9e, 0xbb, 0x3d, 0xa9, 0xc5, 0x6f, 0xc3, 0x55, 0x46, 0xbf, 0x92, 0xce, 0x19, 0x1a, 0x89, 0x12,
	0xca, 0x91, 0xbb, 0x97, 0x53, 0xd9, 0x04, 0xdc, 0xa1, 0x39, 0x91, 0xac, 0xe9, 0x39, 0xc2, 0x52,
	0x9f, 0xc0, 0xfa, 0x56, 0x48, 0x89, 0xa4, 0xb3, 0xc9, 0xe7, 0x4d, 0xc8, 0x84, 0x62, 0xc6, 0x27,
	0x7e, 0xf6, 0x0b, 0xb5, 0x92, 0x97, 0xaa, 0xef, 0xc2, 0xba, 0x41, 0x7d, 0xfa, 0xfc, 0x73, 0xe7,
	0x91, 0xfc, 0x0d, 0x41, 0x3d, 0x57, 0xc7, 0x6e, 0x76, 0x7b, 0x9f, 0x52, 0x49, 0x86, 0x44, 0x12,
	0xfc, 0x08, 0x96, 0x85, 0x24, 0x32, 0xa9, 0xa9, 0x34, 0x5b, 0x0b, 0xf9, 0x9c, 0x8f, 0xa5, 0xed,
	0x45, 0x40, 0x76, 0x82, 0xa7, 0xb6, 0x61, 0x39, 0xb6, 0xf1, 0x3a, 0x5c, 0xdb, 0xeb, 0xb7, 0xfa,
	0xb3, 0x1a, 0x2c, 0xc1, 0x6a, 0xcf, 0xb4, 0x8c, 0xae, 0xd5, 0xa9, 0xa2, 0xc8, 0xb0, 0xf7, 0x2d,
	0x2b, 0x32, 0x14, 0x5c, 0x84, 0x25, 0x63, 0xd7, 0x32, 0xab, 0x85, 0xe6, 0x3f, 0xab, 0xb0, 0x96,
	0x2b, 0x02, 0x7f, 0xab, 0x40, 0x79, 0x4a, 0x23, 0xf8, 0xfd, 0x85, 0x6c, 0xe7, 0x09, 0xb8, 0xfe,
	0xc1, 0x8b, 0x96, 0x25, 0x52, 0x54, 0xbf, 0x41, 0xdf, 0xfd, 0xfe, 0xe7, 0x8f, 0xca, 0x29, 0xbe,
	0x9b, 0xbf, 0x2e, 0xbf, 0x4e, 0x36, 0x7b, 0x2f, 0x08, 0xf9, 0x63, 0x3a, 0x90, 0x42, 0x6f, 0xe8,
	0xf9, 0x6d, 0xb7, 0x89, 0xa0, 0x42, 0x6f, 0x3c, 0xcb, 0xdf, 0xa9, 0xe2, 0xe0, 0x63, 0x7c, 0xef,
	0xa2, 0x52, 0xe2, 0x52, 0x26, 0x2f, 0x02, 0xc0, 0x7f, 0x23, 0x28, 0x9d, 0x51, 0x27, 0xbe, 0xb3,
	0xb0, 0x95, 0xe7, 0xb5, 0x5c, 0xbf, 0xbc, 0xe8, 0xe6, 0xb6, 0x1c, 0xc9, 0xea, 0xa2, 0x86, 0x27,
	0x74, 0xf5, 0xc6, 0xb3, 0xe9, 0x96, 0x67, 0x4b, 0xe7, 0x37, 0x3c, 0x05, 0x80, 0xff, 0x43, 0x50,
	0x99, 0x3e, 0x33, 0xbc, 0x78, 0x81, 0x73, 0xef, 0xb2, 0x7e, 0x23, 0xab, 0x3b, 0xf3, 0x7f, 0xa4,
	0xe5, 0x3a, 0x56, 0x7f, 0x4a, 0x9a, 0xfd, 0x01, 0xa9, 0x2f, 0xbb, 0xe0, 0x0f, 0xf3, 0x23, 0x3d,
	0xe8, 0xaa, 0xaf, 0xb6, 0xea, 0x09, 0x14, 0xfe, 0x0b, 0x41, 0x65, 0xfa, 0xe0, 0x2f, 0x31, 0x81,
	0xb9, 0x6f, 0x88, 0x45, 0x13, 0xc8, 0xd6, 0xdd, 0x78, 0xf9, 0x75, 0x37, 0x5e, 0x6d, 0xdd, 0xed,
	0x5f, 0x11, 0xdc, 0x1e, 0xf0, 0xd1, 0xa2, 0xfe, 0xda, 0xe5, 0xac, 0xb5, 0x5e, 0xf4, 0x15, 0xd0,
	0x43, 0x07, 0xdd, 0xb4, 0xc2, 0xe5, 0x3e, 0x61, 0xae, 0xc6, 0x43, 0x57, 0x77, 0x29, 0x8b, 0xbf,
	0x11, 0xf4, 0x24, 0x44, 0x02, 0x4f, 0x9c, 0xfb, 0xe5, 0xf3, 0xd1, 0xc4, 0xf5, 0x2f, 0x42, 0x3f,
	0x2b, 0x8a, 0xb1, 0xfd, 0x8b, 0x72, 0xb3, 0x93, 0x60, 0x6e, 0xc5, 0x2c, 0x8c, 0x09, 0x8b, 0x87,
	0x49, 0xd1, 0xe1, 0x4a, 0x8c, 0x7f, 0xe7, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8d, 0xc6, 0x50,
	0x98, 0x58, 0x09, 0x00, 0x00,
}
