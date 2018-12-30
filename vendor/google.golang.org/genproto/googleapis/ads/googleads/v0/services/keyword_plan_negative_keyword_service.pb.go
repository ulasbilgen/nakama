// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/keyword_plan_negative_keyword_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Request message for
// [KeywordPlanNegativeKeywordService.GetKeywordPlanNegativeKeyword][google.ads.googleads.v0.services.KeywordPlanNegativeKeywordService.GetKeywordPlanNegativeKeyword].
type GetKeywordPlanNegativeKeywordRequest struct {
	// The resource name of the plan to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKeywordPlanNegativeKeywordRequest) Reset()         { *m = GetKeywordPlanNegativeKeywordRequest{} }
func (m *GetKeywordPlanNegativeKeywordRequest) String() string { return proto.CompactTextString(m) }
func (*GetKeywordPlanNegativeKeywordRequest) ProtoMessage()    {}
func (*GetKeywordPlanNegativeKeywordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_negative_keyword_service_8357f05c658450a2, []int{0}
}
func (m *GetKeywordPlanNegativeKeywordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKeywordPlanNegativeKeywordRequest.Unmarshal(m, b)
}
func (m *GetKeywordPlanNegativeKeywordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKeywordPlanNegativeKeywordRequest.Marshal(b, m, deterministic)
}
func (dst *GetKeywordPlanNegativeKeywordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKeywordPlanNegativeKeywordRequest.Merge(dst, src)
}
func (m *GetKeywordPlanNegativeKeywordRequest) XXX_Size() int {
	return xxx_messageInfo_GetKeywordPlanNegativeKeywordRequest.Size(m)
}
func (m *GetKeywordPlanNegativeKeywordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKeywordPlanNegativeKeywordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetKeywordPlanNegativeKeywordRequest proto.InternalMessageInfo

func (m *GetKeywordPlanNegativeKeywordRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [KeywordPlanNegativeKeywordService.MutateKeywordPlanNegativeKeywords][google.ads.googleads.v0.services.KeywordPlanNegativeKeywordService.MutateKeywordPlanNegativeKeywords].
type MutateKeywordPlanNegativeKeywordsRequest struct {
	// The ID of the customer whose negative keywords are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual Keyword Plan negative
	// keywords.
	Operations           []*KeywordPlanNegativeKeywordOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *MutateKeywordPlanNegativeKeywordsRequest) Reset() {
	*m = MutateKeywordPlanNegativeKeywordsRequest{}
}
func (m *MutateKeywordPlanNegativeKeywordsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateKeywordPlanNegativeKeywordsRequest) ProtoMessage()    {}
func (*MutateKeywordPlanNegativeKeywordsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_negative_keyword_service_8357f05c658450a2, []int{1}
}
func (m *MutateKeywordPlanNegativeKeywordsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordsRequest.Unmarshal(m, b)
}
func (m *MutateKeywordPlanNegativeKeywordsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordsRequest.Marshal(b, m, deterministic)
}
func (dst *MutateKeywordPlanNegativeKeywordsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateKeywordPlanNegativeKeywordsRequest.Merge(dst, src)
}
func (m *MutateKeywordPlanNegativeKeywordsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordsRequest.Size(m)
}
func (m *MutateKeywordPlanNegativeKeywordsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateKeywordPlanNegativeKeywordsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateKeywordPlanNegativeKeywordsRequest proto.InternalMessageInfo

func (m *MutateKeywordPlanNegativeKeywordsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateKeywordPlanNegativeKeywordsRequest) GetOperations() []*KeywordPlanNegativeKeywordOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

// A single operation (create, update, remove) on a Keyword Plan negative
// keyword.
type KeywordPlanNegativeKeywordOperation struct {
	// The FieldMask that determines which resource fields are modified in an
	// update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*KeywordPlanNegativeKeywordOperation_Create
	//	*KeywordPlanNegativeKeywordOperation_Update
	//	*KeywordPlanNegativeKeywordOperation_Remove
	Operation            isKeywordPlanNegativeKeywordOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *KeywordPlanNegativeKeywordOperation) Reset()         { *m = KeywordPlanNegativeKeywordOperation{} }
func (m *KeywordPlanNegativeKeywordOperation) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanNegativeKeywordOperation) ProtoMessage()    {}
func (*KeywordPlanNegativeKeywordOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_negative_keyword_service_8357f05c658450a2, []int{2}
}
func (m *KeywordPlanNegativeKeywordOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanNegativeKeywordOperation.Unmarshal(m, b)
}
func (m *KeywordPlanNegativeKeywordOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanNegativeKeywordOperation.Marshal(b, m, deterministic)
}
func (dst *KeywordPlanNegativeKeywordOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanNegativeKeywordOperation.Merge(dst, src)
}
func (m *KeywordPlanNegativeKeywordOperation) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanNegativeKeywordOperation.Size(m)
}
func (m *KeywordPlanNegativeKeywordOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanNegativeKeywordOperation.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanNegativeKeywordOperation proto.InternalMessageInfo

func (m *KeywordPlanNegativeKeywordOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isKeywordPlanNegativeKeywordOperation_Operation interface {
	isKeywordPlanNegativeKeywordOperation_Operation()
}

type KeywordPlanNegativeKeywordOperation_Create struct {
	Create *resources.KeywordPlanNegativeKeyword `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type KeywordPlanNegativeKeywordOperation_Update struct {
	Update *resources.KeywordPlanNegativeKeyword `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type KeywordPlanNegativeKeywordOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*KeywordPlanNegativeKeywordOperation_Create) isKeywordPlanNegativeKeywordOperation_Operation() {}

func (*KeywordPlanNegativeKeywordOperation_Update) isKeywordPlanNegativeKeywordOperation_Operation() {}

func (*KeywordPlanNegativeKeywordOperation_Remove) isKeywordPlanNegativeKeywordOperation_Operation() {}

func (m *KeywordPlanNegativeKeywordOperation) GetOperation() isKeywordPlanNegativeKeywordOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *KeywordPlanNegativeKeywordOperation) GetCreate() *resources.KeywordPlanNegativeKeyword {
	if x, ok := m.GetOperation().(*KeywordPlanNegativeKeywordOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *KeywordPlanNegativeKeywordOperation) GetUpdate() *resources.KeywordPlanNegativeKeyword {
	if x, ok := m.GetOperation().(*KeywordPlanNegativeKeywordOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *KeywordPlanNegativeKeywordOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*KeywordPlanNegativeKeywordOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*KeywordPlanNegativeKeywordOperation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _KeywordPlanNegativeKeywordOperation_OneofMarshaler, _KeywordPlanNegativeKeywordOperation_OneofUnmarshaler, _KeywordPlanNegativeKeywordOperation_OneofSizer, []interface{}{
		(*KeywordPlanNegativeKeywordOperation_Create)(nil),
		(*KeywordPlanNegativeKeywordOperation_Update)(nil),
		(*KeywordPlanNegativeKeywordOperation_Remove)(nil),
	}
}

func _KeywordPlanNegativeKeywordOperation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*KeywordPlanNegativeKeywordOperation)
	// operation
	switch x := m.Operation.(type) {
	case *KeywordPlanNegativeKeywordOperation_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *KeywordPlanNegativeKeywordOperation_Update:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *KeywordPlanNegativeKeywordOperation_Remove:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Remove)
	case nil:
	default:
		return fmt.Errorf("KeywordPlanNegativeKeywordOperation.Operation has unexpected type %T", x)
	}
	return nil
}

func _KeywordPlanNegativeKeywordOperation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*KeywordPlanNegativeKeywordOperation)
	switch tag {
	case 1: // operation.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.KeywordPlanNegativeKeyword)
		err := b.DecodeMessage(msg)
		m.Operation = &KeywordPlanNegativeKeywordOperation_Create{msg}
		return true, err
	case 2: // operation.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.KeywordPlanNegativeKeyword)
		err := b.DecodeMessage(msg)
		m.Operation = &KeywordPlanNegativeKeywordOperation_Update{msg}
		return true, err
	case 3: // operation.remove
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operation = &KeywordPlanNegativeKeywordOperation_Remove{x}
		return true, err
	default:
		return false, nil
	}
}

func _KeywordPlanNegativeKeywordOperation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*KeywordPlanNegativeKeywordOperation)
	// operation
	switch x := m.Operation.(type) {
	case *KeywordPlanNegativeKeywordOperation_Create:
		s := proto.Size(x.Create)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *KeywordPlanNegativeKeywordOperation_Update:
		s := proto.Size(x.Update)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *KeywordPlanNegativeKeywordOperation_Remove:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Remove)))
		n += len(x.Remove)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for a Keyword Plan negative keyword mutate.
type MutateKeywordPlanNegativeKeywordsResponse struct {
	// All results for the mutate.
	Results              []*MutateKeywordPlanNegativeKeywordResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *MutateKeywordPlanNegativeKeywordsResponse) Reset() {
	*m = MutateKeywordPlanNegativeKeywordsResponse{}
}
func (m *MutateKeywordPlanNegativeKeywordsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateKeywordPlanNegativeKeywordsResponse) ProtoMessage()    {}
func (*MutateKeywordPlanNegativeKeywordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_negative_keyword_service_8357f05c658450a2, []int{3}
}
func (m *MutateKeywordPlanNegativeKeywordsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordsResponse.Unmarshal(m, b)
}
func (m *MutateKeywordPlanNegativeKeywordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordsResponse.Marshal(b, m, deterministic)
}
func (dst *MutateKeywordPlanNegativeKeywordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateKeywordPlanNegativeKeywordsResponse.Merge(dst, src)
}
func (m *MutateKeywordPlanNegativeKeywordsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordsResponse.Size(m)
}
func (m *MutateKeywordPlanNegativeKeywordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateKeywordPlanNegativeKeywordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateKeywordPlanNegativeKeywordsResponse proto.InternalMessageInfo

func (m *MutateKeywordPlanNegativeKeywordsResponse) GetResults() []*MutateKeywordPlanNegativeKeywordResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the Keyword Plan negative keyword mutate.
type MutateKeywordPlanNegativeKeywordResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateKeywordPlanNegativeKeywordResult) Reset() {
	*m = MutateKeywordPlanNegativeKeywordResult{}
}
func (m *MutateKeywordPlanNegativeKeywordResult) String() string { return proto.CompactTextString(m) }
func (*MutateKeywordPlanNegativeKeywordResult) ProtoMessage()    {}
func (*MutateKeywordPlanNegativeKeywordResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_negative_keyword_service_8357f05c658450a2, []int{4}
}
func (m *MutateKeywordPlanNegativeKeywordResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordResult.Unmarshal(m, b)
}
func (m *MutateKeywordPlanNegativeKeywordResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordResult.Marshal(b, m, deterministic)
}
func (dst *MutateKeywordPlanNegativeKeywordResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateKeywordPlanNegativeKeywordResult.Merge(dst, src)
}
func (m *MutateKeywordPlanNegativeKeywordResult) XXX_Size() int {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordResult.Size(m)
}
func (m *MutateKeywordPlanNegativeKeywordResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateKeywordPlanNegativeKeywordResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateKeywordPlanNegativeKeywordResult proto.InternalMessageInfo

func (m *MutateKeywordPlanNegativeKeywordResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetKeywordPlanNegativeKeywordRequest)(nil), "google.ads.googleads.v0.services.GetKeywordPlanNegativeKeywordRequest")
	proto.RegisterType((*MutateKeywordPlanNegativeKeywordsRequest)(nil), "google.ads.googleads.v0.services.MutateKeywordPlanNegativeKeywordsRequest")
	proto.RegisterType((*KeywordPlanNegativeKeywordOperation)(nil), "google.ads.googleads.v0.services.KeywordPlanNegativeKeywordOperation")
	proto.RegisterType((*MutateKeywordPlanNegativeKeywordsResponse)(nil), "google.ads.googleads.v0.services.MutateKeywordPlanNegativeKeywordsResponse")
	proto.RegisterType((*MutateKeywordPlanNegativeKeywordResult)(nil), "google.ads.googleads.v0.services.MutateKeywordPlanNegativeKeywordResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeywordPlanNegativeKeywordServiceClient is the client API for KeywordPlanNegativeKeywordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordPlanNegativeKeywordServiceClient interface {
	// Returns the requested plan in full detail.
	GetKeywordPlanNegativeKeyword(ctx context.Context, in *GetKeywordPlanNegativeKeywordRequest, opts ...grpc.CallOption) (*resources.KeywordPlanNegativeKeyword, error)
	// Creates, updates, or removes Keyword Plan negative keywords. Operation
	// statuses are returned.
	MutateKeywordPlanNegativeKeywords(ctx context.Context, in *MutateKeywordPlanNegativeKeywordsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanNegativeKeywordsResponse, error)
}

type keywordPlanNegativeKeywordServiceClient struct {
	cc *grpc.ClientConn
}

func NewKeywordPlanNegativeKeywordServiceClient(cc *grpc.ClientConn) KeywordPlanNegativeKeywordServiceClient {
	return &keywordPlanNegativeKeywordServiceClient{cc}
}

func (c *keywordPlanNegativeKeywordServiceClient) GetKeywordPlanNegativeKeyword(ctx context.Context, in *GetKeywordPlanNegativeKeywordRequest, opts ...grpc.CallOption) (*resources.KeywordPlanNegativeKeyword, error) {
	out := new(resources.KeywordPlanNegativeKeyword)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.KeywordPlanNegativeKeywordService/GetKeywordPlanNegativeKeyword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordPlanNegativeKeywordServiceClient) MutateKeywordPlanNegativeKeywords(ctx context.Context, in *MutateKeywordPlanNegativeKeywordsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanNegativeKeywordsResponse, error) {
	out := new(MutateKeywordPlanNegativeKeywordsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.KeywordPlanNegativeKeywordService/MutateKeywordPlanNegativeKeywords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanNegativeKeywordServiceServer is the server API for KeywordPlanNegativeKeywordService service.
type KeywordPlanNegativeKeywordServiceServer interface {
	// Returns the requested plan in full detail.
	GetKeywordPlanNegativeKeyword(context.Context, *GetKeywordPlanNegativeKeywordRequest) (*resources.KeywordPlanNegativeKeyword, error)
	// Creates, updates, or removes Keyword Plan negative keywords. Operation
	// statuses are returned.
	MutateKeywordPlanNegativeKeywords(context.Context, *MutateKeywordPlanNegativeKeywordsRequest) (*MutateKeywordPlanNegativeKeywordsResponse, error)
}

func RegisterKeywordPlanNegativeKeywordServiceServer(s *grpc.Server, srv KeywordPlanNegativeKeywordServiceServer) {
	s.RegisterService(&_KeywordPlanNegativeKeywordService_serviceDesc, srv)
}

func _KeywordPlanNegativeKeywordService_GetKeywordPlanNegativeKeyword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeywordPlanNegativeKeywordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanNegativeKeywordServiceServer).GetKeywordPlanNegativeKeyword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.KeywordPlanNegativeKeywordService/GetKeywordPlanNegativeKeyword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanNegativeKeywordServiceServer).GetKeywordPlanNegativeKeyword(ctx, req.(*GetKeywordPlanNegativeKeywordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeywordPlanNegativeKeywordService_MutateKeywordPlanNegativeKeywords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateKeywordPlanNegativeKeywordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanNegativeKeywordServiceServer).MutateKeywordPlanNegativeKeywords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.KeywordPlanNegativeKeywordService/MutateKeywordPlanNegativeKeywords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanNegativeKeywordServiceServer).MutateKeywordPlanNegativeKeywords(ctx, req.(*MutateKeywordPlanNegativeKeywordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeywordPlanNegativeKeywordService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.KeywordPlanNegativeKeywordService",
	HandlerType: (*KeywordPlanNegativeKeywordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKeywordPlanNegativeKeyword",
			Handler:    _KeywordPlanNegativeKeywordService_GetKeywordPlanNegativeKeyword_Handler,
		},
		{
			MethodName: "MutateKeywordPlanNegativeKeywords",
			Handler:    _KeywordPlanNegativeKeywordService_MutateKeywordPlanNegativeKeywords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/keyword_plan_negative_keyword_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/keyword_plan_negative_keyword_service.proto", fileDescriptor_keyword_plan_negative_keyword_service_8357f05c658450a2)
}

var fileDescriptor_keyword_plan_negative_keyword_service_8357f05c658450a2 = []byte{
	// 617 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0x7f, 0x76, 0xab, 0xfe, 0xd4, 0x35, 0x5c, 0xf6, 0x64, 0x45, 0x20, 0x52, 0xb7, 0xaa,
	0x42, 0x0e, 0x76, 0x14, 0x6e, 0xad, 0x22, 0x48, 0xa4, 0x34, 0x81, 0x90, 0x12, 0x19, 0xa9, 0x48,
	0x28, 0x92, 0xb5, 0x89, 0xa7, 0x96, 0x15, 0xdb, 0x6b, 0xbc, 0xeb, 0x20, 0x54, 0xf5, 0xc2, 0x8d,
	0x13, 0x07, 0x1e, 0x00, 0x89, 0x23, 0x47, 0xae, 0xbc, 0x01, 0x57, 0x5e, 0x01, 0x71, 0xe3, 0x1d,
	0x90, 0xbd, 0xde, 0x50, 0x24, 0x1c, 0x47, 0x6a, 0x6f, 0xe3, 0xf1, 0xf8, 0x33, 0x7f, 0xbe, 0xbb,
	0x63, 0xf4, 0xd4, 0xa3, 0xd4, 0x0b, 0xc0, 0x22, 0x2e, 0xb3, 0x84, 0x99, 0x59, 0xcb, 0x96, 0xc5,
	0x20, 0x59, 0xfa, 0x73, 0x60, 0xd6, 0x02, 0xde, 0xbc, 0xa6, 0x89, 0xeb, 0xc4, 0x01, 0x89, 0x9c,
	0x08, 0x3c, 0xc2, 0xfd, 0x25, 0x38, 0xd2, 0x5b, 0x84, 0x99, 0x71, 0x42, 0x39, 0xc5, 0x75, 0x81,
	0x30, 0x89, 0xcb, 0xcc, 0x15, 0xcd, 0x5c, 0xb6, 0x4c, 0x49, 0xab, 0xf5, 0xcb, 0xf2, 0x25, 0xc0,
	0x68, 0x9a, 0x54, 0x26, 0x14, 0x89, 0x6a, 0x77, 0x24, 0x26, 0xf6, 0x2d, 0x12, 0x45, 0x94, 0x13,
	0xee, 0xd3, 0x88, 0x15, 0x6f, 0x8b, 0x32, 0xac, 0xfc, 0x69, 0x96, 0x9e, 0x5b, 0xe7, 0x3e, 0x04,
	0xae, 0x13, 0x12, 0xb6, 0x10, 0x11, 0xc6, 0x08, 0x1d, 0x0c, 0x80, 0x8f, 0x04, 0x73, 0x12, 0x90,
	0xe8, 0xb4, 0xc8, 0x53, 0xb8, 0x6c, 0x78, 0x95, 0x02, 0xe3, 0x78, 0x1f, 0xdd, 0x96, 0x85, 0x39,
	0x11, 0x09, 0x41, 0x57, 0xea, 0x4a, 0x63, 0xd7, 0xbe, 0x25, 0x9d, 0xa7, 0x24, 0x04, 0xe3, 0x8b,
	0x82, 0x1a, 0xe3, 0x94, 0x13, 0x0e, 0xe5, 0x40, 0x26, 0x89, 0xf7, 0x90, 0x36, 0x4f, 0x19, 0xa7,
	0x21, 0x24, 0x8e, 0xef, 0x16, 0x3c, 0x24, 0x5d, 0x8f, 0x5d, 0x0c, 0x08, 0xd1, 0x18, 0x12, 0xd1,
	0x90, 0xae, 0xd6, 0xb7, 0x1a, 0x5a, 0xbb, 0x6f, 0x56, 0x0d, 0xd6, 0x2c, 0x4f, 0xfd, 0x4c, 0xd2,
	0xec, 0x2b, 0x60, 0xe3, 0xab, 0x8a, 0xf6, 0x37, 0xf8, 0x06, 0x1f, 0x23, 0x2d, 0x8d, 0x5d, 0xc2,
	0x21, 0x1f, 0x9f, 0xbe, 0x5d, 0x57, 0x1a, 0x5a, 0xbb, 0x26, 0xeb, 0x91, 0x13, 0x36, 0x4f, 0xb2,
	0x09, 0x8f, 0x09, 0x5b, 0xd8, 0x48, 0x84, 0x67, 0x36, 0x7e, 0x81, 0x76, 0xe6, 0x09, 0x10, 0x2e,
	0xe6, 0xa6, 0xb5, 0x3b, 0xa5, 0x7d, 0xac, 0xe4, 0x5f, 0xd3, 0xc8, 0xf0, 0x3f, 0xbb, 0xc0, 0x65,
	0x60, 0x91, 0x46, 0x57, 0x6f, 0x08, 0x2c, 0x70, 0x58, 0x47, 0x3b, 0x09, 0x84, 0x74, 0x09, 0xfa,
	0x56, 0xa6, 0x4c, 0xf6, 0x46, 0x3c, 0xf7, 0x34, 0xb4, 0xbb, 0x1a, 0x9f, 0xf1, 0x5e, 0x41, 0xf7,
	0x37, 0x90, 0x9c, 0xc5, 0x34, 0x62, 0x80, 0x67, 0xe8, 0xff, 0x04, 0x58, 0x1a, 0x70, 0xa9, 0xe7,
	0xb0, 0x5a, 0xcf, 0x2a, 0xba, 0x9d, 0x03, 0x6d, 0x09, 0x36, 0xc6, 0xe8, 0x70, 0xb3, 0x4f, 0x36,
	0x3a, 0xd3, 0xed, 0x8f, 0xdb, 0x68, 0xaf, 0x9c, 0xf4, 0x5c, 0x54, 0x89, 0x7f, 0x29, 0xe8, 0xee,
	0xda, 0x7b, 0x84, 0x4f, 0xaa, 0x3b, 0xdd, 0xe4, 0x22, 0xd6, 0xae, 0x27, 0xb0, 0xd1, 0x7f, 0xfb,
	0xfd, 0xc7, 0x07, 0xf5, 0x21, 0xee, 0x64, 0xab, 0xe6, 0xe2, 0xaf, 0xf6, 0x3b, 0xf2, 0xee, 0x31,
	0xab, 0x29, 0x77, 0xcf, 0xbf, 0xd4, 0xb4, 0x9a, 0x97, 0xf8, 0x9d, 0x8a, 0xf6, 0x2a, 0x65, 0xc7,
	0x4f, 0xae, 0xaf, 0xae, 0x5c, 0x17, 0xb5, 0xd1, 0x8d, 0xb0, 0xc4, 0x39, 0x34, 0x46, 0xf9, 0x14,
	0xfa, 0xc6, 0xa3, 0x6c, 0x0a, 0x7f, 0xda, 0xbe, 0xb8, 0xb2, 0x90, 0x3a, 0xcd, 0xcb, 0x75, 0x43,
	0x38, 0x0a, 0xf3, 0x64, 0x47, 0x4a, 0xb3, 0xf7, 0x53, 0x41, 0x07, 0x73, 0x1a, 0x56, 0xd6, 0xd7,
	0x3b, 0xac, 0x3c, 0x47, 0x93, 0x6c, 0x8b, 0x4c, 0x94, 0x97, 0xc3, 0x82, 0xe5, 0xd1, 0x80, 0x44,
	0x9e, 0x49, 0x13, 0xcf, 0xf2, 0x20, 0xca, 0x77, 0x8c, 0xfc, 0x59, 0xc4, 0x3e, 0x2b, 0xff, 0x57,
	0x1d, 0x4b, 0xe3, 0x93, 0xba, 0x35, 0xe8, 0x76, 0x3f, 0xab, 0xf5, 0x81, 0x00, 0x76, 0x5d, 0x66,
	0x0a, 0x33, 0xb3, 0xce, 0x5a, 0x66, 0x91, 0x98, 0x7d, 0x93, 0x21, 0xd3, 0xae, 0xcb, 0xa6, 0xab,
	0x90, 0xe9, 0x59, 0x6b, 0x2a, 0x43, 0x66, 0x3b, 0x79, 0x01, 0x0f, 0x7e, 0x07, 0x00, 0x00, 0xff,
	0xff, 0xe9, 0x0e, 0xc7, 0x43, 0x2b, 0x07, 0x00, 0x00,
}