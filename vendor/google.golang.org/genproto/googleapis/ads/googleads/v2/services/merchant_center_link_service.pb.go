// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/merchant_center_link_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [MerchantCenterLinkService.ListMerchantCenterLinks][google.ads.googleads.v2.services.MerchantCenterLinkService.ListMerchantCenterLinks].
type ListMerchantCenterLinksRequest struct {
	// The ID of the customer onto which to apply the Merchant Center link list
	// operation.
	CustomerId           string   `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMerchantCenterLinksRequest) Reset()         { *m = ListMerchantCenterLinksRequest{} }
func (m *ListMerchantCenterLinksRequest) String() string { return proto.CompactTextString(m) }
func (*ListMerchantCenterLinksRequest) ProtoMessage()    {}
func (*ListMerchantCenterLinksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd0dc4f0ea434cdd, []int{0}
}

func (m *ListMerchantCenterLinksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMerchantCenterLinksRequest.Unmarshal(m, b)
}
func (m *ListMerchantCenterLinksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMerchantCenterLinksRequest.Marshal(b, m, deterministic)
}
func (m *ListMerchantCenterLinksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMerchantCenterLinksRequest.Merge(m, src)
}
func (m *ListMerchantCenterLinksRequest) XXX_Size() int {
	return xxx_messageInfo_ListMerchantCenterLinksRequest.Size(m)
}
func (m *ListMerchantCenterLinksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMerchantCenterLinksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMerchantCenterLinksRequest proto.InternalMessageInfo

func (m *ListMerchantCenterLinksRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

// Response message for [MerchantCenterLinkService.ListMerchantCenterLinks][google.ads.googleads.v2.services.MerchantCenterLinkService.ListMerchantCenterLinks].
type ListMerchantCenterLinksResponse struct {
	// Merchant Center links available for the requested customer
	MerchantCenterLinks  []*resources.MerchantCenterLink `protobuf:"bytes,1,rep,name=merchant_center_links,json=merchantCenterLinks,proto3" json:"merchant_center_links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ListMerchantCenterLinksResponse) Reset()         { *m = ListMerchantCenterLinksResponse{} }
func (m *ListMerchantCenterLinksResponse) String() string { return proto.CompactTextString(m) }
func (*ListMerchantCenterLinksResponse) ProtoMessage()    {}
func (*ListMerchantCenterLinksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd0dc4f0ea434cdd, []int{1}
}

func (m *ListMerchantCenterLinksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMerchantCenterLinksResponse.Unmarshal(m, b)
}
func (m *ListMerchantCenterLinksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMerchantCenterLinksResponse.Marshal(b, m, deterministic)
}
func (m *ListMerchantCenterLinksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMerchantCenterLinksResponse.Merge(m, src)
}
func (m *ListMerchantCenterLinksResponse) XXX_Size() int {
	return xxx_messageInfo_ListMerchantCenterLinksResponse.Size(m)
}
func (m *ListMerchantCenterLinksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMerchantCenterLinksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMerchantCenterLinksResponse proto.InternalMessageInfo

func (m *ListMerchantCenterLinksResponse) GetMerchantCenterLinks() []*resources.MerchantCenterLink {
	if m != nil {
		return m.MerchantCenterLinks
	}
	return nil
}

// Request message for [MerchantCenterLinkService.GetMerchantCenterLink][google.ads.googleads.v2.services.MerchantCenterLinkService.GetMerchantCenterLink].
type GetMerchantCenterLinkRequest struct {
	// Resource name of the Merchant Center link.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMerchantCenterLinkRequest) Reset()         { *m = GetMerchantCenterLinkRequest{} }
func (m *GetMerchantCenterLinkRequest) String() string { return proto.CompactTextString(m) }
func (*GetMerchantCenterLinkRequest) ProtoMessage()    {}
func (*GetMerchantCenterLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd0dc4f0ea434cdd, []int{2}
}

func (m *GetMerchantCenterLinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMerchantCenterLinkRequest.Unmarshal(m, b)
}
func (m *GetMerchantCenterLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMerchantCenterLinkRequest.Marshal(b, m, deterministic)
}
func (m *GetMerchantCenterLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMerchantCenterLinkRequest.Merge(m, src)
}
func (m *GetMerchantCenterLinkRequest) XXX_Size() int {
	return xxx_messageInfo_GetMerchantCenterLinkRequest.Size(m)
}
func (m *GetMerchantCenterLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMerchantCenterLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMerchantCenterLinkRequest proto.InternalMessageInfo

func (m *GetMerchantCenterLinkRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [MerchantCenterLinkService.MutateMerchantCenterLink][google.ads.googleads.v2.services.MerchantCenterLinkService.MutateMerchantCenterLink].
type MutateMerchantCenterLinkRequest struct {
	// The ID of the customer being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The operation to perform on the link
	Operation            *MerchantCenterLinkOperation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *MutateMerchantCenterLinkRequest) Reset()         { *m = MutateMerchantCenterLinkRequest{} }
func (m *MutateMerchantCenterLinkRequest) String() string { return proto.CompactTextString(m) }
func (*MutateMerchantCenterLinkRequest) ProtoMessage()    {}
func (*MutateMerchantCenterLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd0dc4f0ea434cdd, []int{3}
}

func (m *MutateMerchantCenterLinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateMerchantCenterLinkRequest.Unmarshal(m, b)
}
func (m *MutateMerchantCenterLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateMerchantCenterLinkRequest.Marshal(b, m, deterministic)
}
func (m *MutateMerchantCenterLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateMerchantCenterLinkRequest.Merge(m, src)
}
func (m *MutateMerchantCenterLinkRequest) XXX_Size() int {
	return xxx_messageInfo_MutateMerchantCenterLinkRequest.Size(m)
}
func (m *MutateMerchantCenterLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateMerchantCenterLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateMerchantCenterLinkRequest proto.InternalMessageInfo

func (m *MutateMerchantCenterLinkRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateMerchantCenterLinkRequest) GetOperation() *MerchantCenterLinkOperation {
	if m != nil {
		return m.Operation
	}
	return nil
}

// A single update on a Merchant Center link.
type MerchantCenterLinkOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The operation to perform
	//
	// Types that are valid to be assigned to Operation:
	//	*MerchantCenterLinkOperation_Update
	//	*MerchantCenterLinkOperation_Remove
	Operation            isMerchantCenterLinkOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *MerchantCenterLinkOperation) Reset()         { *m = MerchantCenterLinkOperation{} }
func (m *MerchantCenterLinkOperation) String() string { return proto.CompactTextString(m) }
func (*MerchantCenterLinkOperation) ProtoMessage()    {}
func (*MerchantCenterLinkOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd0dc4f0ea434cdd, []int{4}
}

func (m *MerchantCenterLinkOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantCenterLinkOperation.Unmarshal(m, b)
}
func (m *MerchantCenterLinkOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantCenterLinkOperation.Marshal(b, m, deterministic)
}
func (m *MerchantCenterLinkOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantCenterLinkOperation.Merge(m, src)
}
func (m *MerchantCenterLinkOperation) XXX_Size() int {
	return xxx_messageInfo_MerchantCenterLinkOperation.Size(m)
}
func (m *MerchantCenterLinkOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantCenterLinkOperation.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantCenterLinkOperation proto.InternalMessageInfo

func (m *MerchantCenterLinkOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isMerchantCenterLinkOperation_Operation interface {
	isMerchantCenterLinkOperation_Operation()
}

type MerchantCenterLinkOperation_Update struct {
	Update *resources.MerchantCenterLink `protobuf:"bytes,1,opt,name=update,proto3,oneof"`
}

type MerchantCenterLinkOperation_Remove struct {
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*MerchantCenterLinkOperation_Update) isMerchantCenterLinkOperation_Operation() {}

func (*MerchantCenterLinkOperation_Remove) isMerchantCenterLinkOperation_Operation() {}

func (m *MerchantCenterLinkOperation) GetOperation() isMerchantCenterLinkOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *MerchantCenterLinkOperation) GetUpdate() *resources.MerchantCenterLink {
	if x, ok := m.GetOperation().(*MerchantCenterLinkOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *MerchantCenterLinkOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*MerchantCenterLinkOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MerchantCenterLinkOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MerchantCenterLinkOperation_Update)(nil),
		(*MerchantCenterLinkOperation_Remove)(nil),
	}
}

// Response message for Merchant Center link mutate.
type MutateMerchantCenterLinkResponse struct {
	// Result for the mutate.
	Result               *MutateMerchantCenterLinkResult `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *MutateMerchantCenterLinkResponse) Reset()         { *m = MutateMerchantCenterLinkResponse{} }
func (m *MutateMerchantCenterLinkResponse) String() string { return proto.CompactTextString(m) }
func (*MutateMerchantCenterLinkResponse) ProtoMessage()    {}
func (*MutateMerchantCenterLinkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd0dc4f0ea434cdd, []int{5}
}

func (m *MutateMerchantCenterLinkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateMerchantCenterLinkResponse.Unmarshal(m, b)
}
func (m *MutateMerchantCenterLinkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateMerchantCenterLinkResponse.Marshal(b, m, deterministic)
}
func (m *MutateMerchantCenterLinkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateMerchantCenterLinkResponse.Merge(m, src)
}
func (m *MutateMerchantCenterLinkResponse) XXX_Size() int {
	return xxx_messageInfo_MutateMerchantCenterLinkResponse.Size(m)
}
func (m *MutateMerchantCenterLinkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateMerchantCenterLinkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateMerchantCenterLinkResponse proto.InternalMessageInfo

func (m *MutateMerchantCenterLinkResponse) GetResult() *MutateMerchantCenterLinkResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// The result for the Merchant Center link mutate.
type MutateMerchantCenterLinkResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateMerchantCenterLinkResult) Reset()         { *m = MutateMerchantCenterLinkResult{} }
func (m *MutateMerchantCenterLinkResult) String() string { return proto.CompactTextString(m) }
func (*MutateMerchantCenterLinkResult) ProtoMessage()    {}
func (*MutateMerchantCenterLinkResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd0dc4f0ea434cdd, []int{6}
}

func (m *MutateMerchantCenterLinkResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateMerchantCenterLinkResult.Unmarshal(m, b)
}
func (m *MutateMerchantCenterLinkResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateMerchantCenterLinkResult.Marshal(b, m, deterministic)
}
func (m *MutateMerchantCenterLinkResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateMerchantCenterLinkResult.Merge(m, src)
}
func (m *MutateMerchantCenterLinkResult) XXX_Size() int {
	return xxx_messageInfo_MutateMerchantCenterLinkResult.Size(m)
}
func (m *MutateMerchantCenterLinkResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateMerchantCenterLinkResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateMerchantCenterLinkResult proto.InternalMessageInfo

func (m *MutateMerchantCenterLinkResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*ListMerchantCenterLinksRequest)(nil), "google.ads.googleads.v2.services.ListMerchantCenterLinksRequest")
	proto.RegisterType((*ListMerchantCenterLinksResponse)(nil), "google.ads.googleads.v2.services.ListMerchantCenterLinksResponse")
	proto.RegisterType((*GetMerchantCenterLinkRequest)(nil), "google.ads.googleads.v2.services.GetMerchantCenterLinkRequest")
	proto.RegisterType((*MutateMerchantCenterLinkRequest)(nil), "google.ads.googleads.v2.services.MutateMerchantCenterLinkRequest")
	proto.RegisterType((*MerchantCenterLinkOperation)(nil), "google.ads.googleads.v2.services.MerchantCenterLinkOperation")
	proto.RegisterType((*MutateMerchantCenterLinkResponse)(nil), "google.ads.googleads.v2.services.MutateMerchantCenterLinkResponse")
	proto.RegisterType((*MutateMerchantCenterLinkResult)(nil), "google.ads.googleads.v2.services.MutateMerchantCenterLinkResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/merchant_center_link_service.proto", fileDescriptor_dd0dc4f0ea434cdd)
}

var fileDescriptor_dd0dc4f0ea434cdd = []byte{
	// 697 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x4f, 0x6b, 0xd4, 0x4e,
	0x18, 0xfe, 0x65, 0xcb, 0xaf, 0xd0, 0x59, 0xbd, 0x8c, 0x94, 0xc6, 0x6d, 0x69, 0x97, 0xd8, 0x43,
	0xd9, 0xc3, 0x04, 0x53, 0x8a, 0x9a, 0xba, 0x62, 0x76, 0xd1, 0x56, 0x68, 0x6d, 0x89, 0x50, 0x44,
	0x17, 0x96, 0x34, 0x99, 0xae, 0xa1, 0x49, 0x66, 0xcd, 0x4c, 0xf6, 0x52, 0x7b, 0x11, 0xbc, 0x79,
	0xf3, 0x03, 0x08, 0x1e, 0xfd, 0x28, 0x85, 0x5e, 0xf4, 0xe4, 0x5d, 0x3c, 0xf8, 0x29, 0x24, 0xf3,
	0x67, 0xb7, 0xd2, 0x64, 0xb7, 0x76, 0x6f, 0x6f, 0x66, 0xde, 0x79, 0x9e, 0x79, 0xde, 0xf7, 0x79,
	0x27, 0xa0, 0xdd, 0x23, 0xa4, 0x17, 0x61, 0xd3, 0x0b, 0xa8, 0x29, 0xc2, 0x3c, 0x1a, 0x58, 0x26,
	0xc5, 0xe9, 0x20, 0xf4, 0x31, 0x35, 0x63, 0x9c, 0xfa, 0x6f, 0xbc, 0x84, 0x75, 0x7d, 0x9c, 0x30,
	0x9c, 0x76, 0xa3, 0x30, 0x39, 0xee, 0xca, 0x5d, 0xd4, 0x4f, 0x09, 0x23, 0xb0, 0x2e, 0x4e, 0x22,
	0x2f, 0xa0, 0x68, 0x08, 0x82, 0x06, 0x16, 0x52, 0x20, 0xb5, 0x87, 0x65, 0x34, 0x29, 0xa6, 0x24,
	0x4b, 0xcb, 0x78, 0x04, 0x7e, 0x6d, 0x49, 0x9d, 0xee, 0x87, 0xa6, 0x97, 0x24, 0x84, 0x79, 0x2c,
	0x24, 0x09, 0x95, 0xbb, 0x92, 0xdd, 0xe4, 0x5f, 0x87, 0xd9, 0x91, 0x79, 0x14, 0xe2, 0x28, 0xe8,
	0xc6, 0x1e, 0x55, 0xe7, 0x17, 0x2e, 0x9c, 0xf7, 0xa3, 0x10, 0x27, 0x4c, 0x6c, 0x18, 0x0e, 0x58,
	0xde, 0x09, 0x29, 0xdb, 0x95, 0xd4, 0x6d, 0xce, 0xbc, 0x13, 0x26, 0xc7, 0xd4, 0xc5, 0x6f, 0x33,
	0x4c, 0x19, 0x5c, 0x01, 0x55, 0x3f, 0xa3, 0x8c, 0xc4, 0x38, 0xed, 0x86, 0x81, 0xae, 0xd5, 0xb5,
	0xb5, 0x39, 0x17, 0xa8, 0xa5, 0x67, 0x81, 0xf1, 0x51, 0x03, 0x2b, 0xa5, 0x18, 0xb4, 0x4f, 0x12,
	0x8a, 0x61, 0x08, 0xe6, 0x8b, 0xd4, 0x51, 0x5d, 0xab, 0xcf, 0xac, 0x55, 0xad, 0x0d, 0x54, 0x56,
	0xbf, 0x61, 0x75, 0xd0, 0x65, 0x78, 0xf7, 0x56, 0x7c, 0x99, 0xd2, 0x68, 0x83, 0xa5, 0x2d, 0x5c,
	0x70, 0x19, 0xa5, 0xe7, 0x0e, 0xb8, 0xa9, 0x40, 0xbb, 0x89, 0x17, 0x63, 0xa9, 0xe8, 0x86, 0x5a,
	0x7c, 0xee, 0xc5, 0xd8, 0xf8, 0xac, 0x81, 0x95, 0xdd, 0x8c, 0x79, 0x0c, 0x97, 0x03, 0x4d, 0x2a,
	0x0c, 0x7c, 0x0d, 0xe6, 0x48, 0x1f, 0xa7, 0xbc, 0x55, 0x7a, 0xa5, 0xae, 0xad, 0x55, 0xad, 0x26,
	0x9a, 0x64, 0x94, 0x02, 0x9d, 0x7b, 0x0a, 0xc4, 0x1d, 0xe1, 0x19, 0xdf, 0x34, 0xb0, 0x38, 0x26,
	0x15, 0x6e, 0x82, 0x6a, 0xd6, 0x0f, 0x3c, 0x86, 0xb9, 0x0d, 0xf4, 0x19, 0x4e, 0x5f, 0x53, 0xf4,
	0xca, 0x29, 0xe8, 0x69, 0xee, 0x94, 0x5d, 0x8f, 0x1e, 0xbb, 0x40, 0xa4, 0xe7, 0x31, 0xdc, 0x03,
	0xb3, 0xe2, 0x8b, 0xab, 0xba, 0x6e, 0x7f, 0xb6, 0xff, 0x73, 0x25, 0x0c, 0xd4, 0xc1, 0x6c, 0x8a,
	0x63, 0x32, 0xc0, 0xbc, 0x0e, 0x73, 0xf9, 0x8e, 0xf8, 0x6e, 0x55, 0x2f, 0x14, 0xc9, 0x78, 0x07,
	0xea, 0xe5, 0x55, 0x97, 0x56, 0x7a, 0x99, 0x43, 0xd1, 0x2c, 0x62, 0xb2, 0xa4, 0x8f, 0xaf, 0x50,
	0xd2, 0x72, 0xcc, 0x2c, 0x62, 0xae, 0xc4, 0x33, 0x9e, 0x80, 0xe5, 0xf1, 0x99, 0x57, 0xf2, 0x8e,
	0x75, 0xfe, 0x3f, 0xb8, 0x7d, 0x19, 0xe1, 0x85, 0xb8, 0x0c, 0xfc, 0xa1, 0x81, 0x85, 0x92, 0x69,
	0x81, 0x57, 0x90, 0x32, 0x7e, 0x58, 0x6b, 0xce, 0x14, 0x08, 0xa2, 0xbe, 0xc6, 0x83, 0xf7, 0xdf,
	0x7f, 0x7e, 0xaa, 0xac, 0xc3, 0xbb, 0xf9, 0xe3, 0xa4, 0xdc, 0x4c, 0xcd, 0x93, 0x0b, 0x5e, 0x6f,
	0x36, 0x4e, 0xcd, 0x82, 0xd1, 0x83, 0xe7, 0x1a, 0x98, 0x2f, 0x9c, 0x3d, 0xf8, 0x68, 0xf2, 0xbd,
	0xc6, 0x0d, 0x6d, 0xed, 0x7a, 0x06, 0x34, 0x9a, 0x5c, 0xcb, 0x3d, 0xb8, 0x91, 0x6b, 0x39, 0xf9,
	0xab, 0x75, 0xcd, 0x91, 0xb4, 0x46, 0x91, 0x18, 0xb3, 0x71, 0x0a, 0x7f, 0x69, 0x40, 0x2f, 0x73,
	0x04, 0x74, 0xa6, 0xf1, 0x9d, 0x50, 0xd5, 0x9a, 0xca, 0xba, 0xa2, 0x5d, 0x6d, 0x2e, 0xb1, 0x69,
	0xdc, 0xff, 0xe7, 0x76, 0xd9, 0x31, 0xc7, 0xb6, 0xb5, 0x46, 0x6d, 0xf1, 0xcc, 0xd1, 0x47, 0xfc,
	0x32, 0xea, 0x87, 0x14, 0xf9, 0x24, 0x6e, 0x7d, 0xa8, 0x80, 0x55, 0x9f, 0xc4, 0x13, 0xef, 0xda,
	0x5a, 0x2e, 0x75, 0xfd, 0x7e, 0xfe, 0xdc, 0xec, 0x6b, 0xaf, 0xb6, 0x25, 0x46, 0x8f, 0x44, 0x5e,
	0xd2, 0x43, 0x24, 0xed, 0x99, 0x3d, 0x9c, 0xf0, 0xc7, 0xc8, 0x1c, 0xb1, 0x96, 0xff, 0x8a, 0x37,
	0x55, 0xf0, 0xa5, 0x32, 0xb3, 0xe5, 0x38, 0x5f, 0x2b, 0xf5, 0x2d, 0x01, 0xe8, 0x04, 0x14, 0x89,
	0x30, 0x8f, 0x0e, 0x2c, 0x24, 0x89, 0xe9, 0x99, 0x4a, 0xe9, 0x38, 0x01, 0xed, 0x0c, 0x53, 0x3a,
	0x07, 0x56, 0x47, 0xa5, 0xfc, 0xae, 0xac, 0x8a, 0x75, 0xdb, 0x76, 0x02, 0x6a, 0xdb, 0xc3, 0x24,
	0xdb, 0x3e, 0xb0, 0x6c, 0x5b, 0xa5, 0x1d, 0xce, 0xf2, 0x7b, 0xae, 0xff, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x2b, 0xf0, 0x3b, 0xd9, 0x31, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MerchantCenterLinkServiceClient is the client API for MerchantCenterLinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MerchantCenterLinkServiceClient interface {
	// Returns Merchant Center links available for this customer.
	ListMerchantCenterLinks(ctx context.Context, in *ListMerchantCenterLinksRequest, opts ...grpc.CallOption) (*ListMerchantCenterLinksResponse, error)
	// Returns the Merchant Center link in full detail.
	GetMerchantCenterLink(ctx context.Context, in *GetMerchantCenterLinkRequest, opts ...grpc.CallOption) (*resources.MerchantCenterLink, error)
	// Updates status or removes a Merchant Center link.
	MutateMerchantCenterLink(ctx context.Context, in *MutateMerchantCenterLinkRequest, opts ...grpc.CallOption) (*MutateMerchantCenterLinkResponse, error)
}

type merchantCenterLinkServiceClient struct {
	cc *grpc.ClientConn
}

func NewMerchantCenterLinkServiceClient(cc *grpc.ClientConn) MerchantCenterLinkServiceClient {
	return &merchantCenterLinkServiceClient{cc}
}

func (c *merchantCenterLinkServiceClient) ListMerchantCenterLinks(ctx context.Context, in *ListMerchantCenterLinksRequest, opts ...grpc.CallOption) (*ListMerchantCenterLinksResponse, error) {
	out := new(ListMerchantCenterLinksResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.MerchantCenterLinkService/ListMerchantCenterLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantCenterLinkServiceClient) GetMerchantCenterLink(ctx context.Context, in *GetMerchantCenterLinkRequest, opts ...grpc.CallOption) (*resources.MerchantCenterLink, error) {
	out := new(resources.MerchantCenterLink)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.MerchantCenterLinkService/GetMerchantCenterLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantCenterLinkServiceClient) MutateMerchantCenterLink(ctx context.Context, in *MutateMerchantCenterLinkRequest, opts ...grpc.CallOption) (*MutateMerchantCenterLinkResponse, error) {
	out := new(MutateMerchantCenterLinkResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.MerchantCenterLinkService/MutateMerchantCenterLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MerchantCenterLinkServiceServer is the server API for MerchantCenterLinkService service.
type MerchantCenterLinkServiceServer interface {
	// Returns Merchant Center links available for this customer.
	ListMerchantCenterLinks(context.Context, *ListMerchantCenterLinksRequest) (*ListMerchantCenterLinksResponse, error)
	// Returns the Merchant Center link in full detail.
	GetMerchantCenterLink(context.Context, *GetMerchantCenterLinkRequest) (*resources.MerchantCenterLink, error)
	// Updates status or removes a Merchant Center link.
	MutateMerchantCenterLink(context.Context, *MutateMerchantCenterLinkRequest) (*MutateMerchantCenterLinkResponse, error)
}

// UnimplementedMerchantCenterLinkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMerchantCenterLinkServiceServer struct {
}

func (*UnimplementedMerchantCenterLinkServiceServer) ListMerchantCenterLinks(ctx context.Context, req *ListMerchantCenterLinksRequest) (*ListMerchantCenterLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMerchantCenterLinks not implemented")
}
func (*UnimplementedMerchantCenterLinkServiceServer) GetMerchantCenterLink(ctx context.Context, req *GetMerchantCenterLinkRequest) (*resources.MerchantCenterLink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMerchantCenterLink not implemented")
}
func (*UnimplementedMerchantCenterLinkServiceServer) MutateMerchantCenterLink(ctx context.Context, req *MutateMerchantCenterLinkRequest) (*MutateMerchantCenterLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateMerchantCenterLink not implemented")
}

func RegisterMerchantCenterLinkServiceServer(s *grpc.Server, srv MerchantCenterLinkServiceServer) {
	s.RegisterService(&_MerchantCenterLinkService_serviceDesc, srv)
}

func _MerchantCenterLinkService_ListMerchantCenterLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMerchantCenterLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantCenterLinkServiceServer).ListMerchantCenterLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.MerchantCenterLinkService/ListMerchantCenterLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantCenterLinkServiceServer).ListMerchantCenterLinks(ctx, req.(*ListMerchantCenterLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantCenterLinkService_GetMerchantCenterLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMerchantCenterLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantCenterLinkServiceServer).GetMerchantCenterLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.MerchantCenterLinkService/GetMerchantCenterLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantCenterLinkServiceServer).GetMerchantCenterLink(ctx, req.(*GetMerchantCenterLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantCenterLinkService_MutateMerchantCenterLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateMerchantCenterLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantCenterLinkServiceServer).MutateMerchantCenterLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.MerchantCenterLinkService/MutateMerchantCenterLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantCenterLinkServiceServer).MutateMerchantCenterLink(ctx, req.(*MutateMerchantCenterLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MerchantCenterLinkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.MerchantCenterLinkService",
	HandlerType: (*MerchantCenterLinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMerchantCenterLinks",
			Handler:    _MerchantCenterLinkService_ListMerchantCenterLinks_Handler,
		},
		{
			MethodName: "GetMerchantCenterLink",
			Handler:    _MerchantCenterLinkService_GetMerchantCenterLink_Handler,
		},
		{
			MethodName: "MutateMerchantCenterLink",
			Handler:    _MerchantCenterLinkService_MutateMerchantCenterLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/merchant_center_link_service.proto",
}
