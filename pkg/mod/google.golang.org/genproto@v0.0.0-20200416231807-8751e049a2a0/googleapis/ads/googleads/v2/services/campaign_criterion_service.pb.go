// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/campaign_criterion_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for [CampaignCriterionService.GetCampaignCriterion][google.ads.googleads.v2.services.CampaignCriterionService.GetCampaignCriterion].
type GetCampaignCriterionRequest struct {
	// Required. The resource name of the criterion to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignCriterionRequest) Reset()         { *m = GetCampaignCriterionRequest{} }
func (m *GetCampaignCriterionRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignCriterionRequest) ProtoMessage()    {}
func (*GetCampaignCriterionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdc082e26c646a0b, []int{0}
}

func (m *GetCampaignCriterionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignCriterionRequest.Unmarshal(m, b)
}
func (m *GetCampaignCriterionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignCriterionRequest.Marshal(b, m, deterministic)
}
func (m *GetCampaignCriterionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignCriterionRequest.Merge(m, src)
}
func (m *GetCampaignCriterionRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignCriterionRequest.Size(m)
}
func (m *GetCampaignCriterionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignCriterionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignCriterionRequest proto.InternalMessageInfo

func (m *GetCampaignCriterionRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CampaignCriterionService.MutateCampaignCriteria][google.ads.googleads.v2.services.CampaignCriterionService.MutateCampaignCriteria].
type MutateCampaignCriteriaRequest struct {
	// Required. The ID of the customer whose criteria are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual criteria.
	Operations []*CampaignCriterionOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignCriteriaRequest) Reset()         { *m = MutateCampaignCriteriaRequest{} }
func (m *MutateCampaignCriteriaRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignCriteriaRequest) ProtoMessage()    {}
func (*MutateCampaignCriteriaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdc082e26c646a0b, []int{1}
}

func (m *MutateCampaignCriteriaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignCriteriaRequest.Unmarshal(m, b)
}
func (m *MutateCampaignCriteriaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignCriteriaRequest.Marshal(b, m, deterministic)
}
func (m *MutateCampaignCriteriaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignCriteriaRequest.Merge(m, src)
}
func (m *MutateCampaignCriteriaRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignCriteriaRequest.Size(m)
}
func (m *MutateCampaignCriteriaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignCriteriaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignCriteriaRequest proto.InternalMessageInfo

func (m *MutateCampaignCriteriaRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCampaignCriteriaRequest) GetOperations() []*CampaignCriterionOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCampaignCriteriaRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCampaignCriteriaRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on a campaign criterion.
type CampaignCriterionOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CampaignCriterionOperation_Create
	//	*CampaignCriterionOperation_Update
	//	*CampaignCriterionOperation_Remove
	Operation            isCampaignCriterionOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *CampaignCriterionOperation) Reset()         { *m = CampaignCriterionOperation{} }
func (m *CampaignCriterionOperation) String() string { return proto.CompactTextString(m) }
func (*CampaignCriterionOperation) ProtoMessage()    {}
func (*CampaignCriterionOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdc082e26c646a0b, []int{2}
}

func (m *CampaignCriterionOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignCriterionOperation.Unmarshal(m, b)
}
func (m *CampaignCriterionOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignCriterionOperation.Marshal(b, m, deterministic)
}
func (m *CampaignCriterionOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignCriterionOperation.Merge(m, src)
}
func (m *CampaignCriterionOperation) XXX_Size() int {
	return xxx_messageInfo_CampaignCriterionOperation.Size(m)
}
func (m *CampaignCriterionOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignCriterionOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignCriterionOperation proto.InternalMessageInfo

func (m *CampaignCriterionOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isCampaignCriterionOperation_Operation interface {
	isCampaignCriterionOperation_Operation()
}

type CampaignCriterionOperation_Create struct {
	Create *resources.CampaignCriterion `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CampaignCriterionOperation_Update struct {
	Update *resources.CampaignCriterion `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type CampaignCriterionOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*CampaignCriterionOperation_Create) isCampaignCriterionOperation_Operation() {}

func (*CampaignCriterionOperation_Update) isCampaignCriterionOperation_Operation() {}

func (*CampaignCriterionOperation_Remove) isCampaignCriterionOperation_Operation() {}

func (m *CampaignCriterionOperation) GetOperation() isCampaignCriterionOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CampaignCriterionOperation) GetCreate() *resources.CampaignCriterion {
	if x, ok := m.GetOperation().(*CampaignCriterionOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CampaignCriterionOperation) GetUpdate() *resources.CampaignCriterion {
	if x, ok := m.GetOperation().(*CampaignCriterionOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *CampaignCriterionOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CampaignCriterionOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignCriterionOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignCriterionOperation_Create)(nil),
		(*CampaignCriterionOperation_Update)(nil),
		(*CampaignCriterionOperation_Remove)(nil),
	}
}

// Response message for campaign criterion mutate.
type MutateCampaignCriteriaResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCampaignCriterionResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *MutateCampaignCriteriaResponse) Reset()         { *m = MutateCampaignCriteriaResponse{} }
func (m *MutateCampaignCriteriaResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignCriteriaResponse) ProtoMessage()    {}
func (*MutateCampaignCriteriaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdc082e26c646a0b, []int{3}
}

func (m *MutateCampaignCriteriaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignCriteriaResponse.Unmarshal(m, b)
}
func (m *MutateCampaignCriteriaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignCriteriaResponse.Marshal(b, m, deterministic)
}
func (m *MutateCampaignCriteriaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignCriteriaResponse.Merge(m, src)
}
func (m *MutateCampaignCriteriaResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignCriteriaResponse.Size(m)
}
func (m *MutateCampaignCriteriaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignCriteriaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignCriteriaResponse proto.InternalMessageInfo

func (m *MutateCampaignCriteriaResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCampaignCriteriaResponse) GetResults() []*MutateCampaignCriterionResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the criterion mutate.
type MutateCampaignCriterionResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignCriterionResult) Reset()         { *m = MutateCampaignCriterionResult{} }
func (m *MutateCampaignCriterionResult) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignCriterionResult) ProtoMessage()    {}
func (*MutateCampaignCriterionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdc082e26c646a0b, []int{4}
}

func (m *MutateCampaignCriterionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignCriterionResult.Unmarshal(m, b)
}
func (m *MutateCampaignCriterionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignCriterionResult.Marshal(b, m, deterministic)
}
func (m *MutateCampaignCriterionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignCriterionResult.Merge(m, src)
}
func (m *MutateCampaignCriterionResult) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignCriterionResult.Size(m)
}
func (m *MutateCampaignCriterionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignCriterionResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignCriterionResult proto.InternalMessageInfo

func (m *MutateCampaignCriterionResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignCriterionRequest)(nil), "google.ads.googleads.v2.services.GetCampaignCriterionRequest")
	proto.RegisterType((*MutateCampaignCriteriaRequest)(nil), "google.ads.googleads.v2.services.MutateCampaignCriteriaRequest")
	proto.RegisterType((*CampaignCriterionOperation)(nil), "google.ads.googleads.v2.services.CampaignCriterionOperation")
	proto.RegisterType((*MutateCampaignCriteriaResponse)(nil), "google.ads.googleads.v2.services.MutateCampaignCriteriaResponse")
	proto.RegisterType((*MutateCampaignCriterionResult)(nil), "google.ads.googleads.v2.services.MutateCampaignCriterionResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/campaign_criterion_service.proto", fileDescriptor_cdc082e26c646a0b)
}

var fileDescriptor_cdc082e26c646a0b = []byte{
	// 790 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x4d, 0x6f, 0xe3, 0x44,
	0x18, 0xc6, 0x4e, 0x55, 0xe8, 0xa4, 0x05, 0x69, 0x80, 0x62, 0x52, 0x0a, 0x91, 0xa9, 0x44, 0x15,
	0x55, 0xb6, 0x64, 0xca, 0xc5, 0xa5, 0x02, 0xa7, 0xd0, 0x16, 0xa4, 0x7e, 0x28, 0x95, 0x8a, 0x8a,
	0x22, 0x59, 0x53, 0x7b, 0x1a, 0xac, 0xda, 0x1e, 0x33, 0x33, 0x8e, 0x54, 0x55, 0xbd, 0x20, 0x4e,
	0x5c, 0xf9, 0x03, 0x88, 0x23, 0x7f, 0x83, 0x5b, 0xaf, 0x7b, 0x58, 0xa9, 0xa7, 0x1e, 0xf6, 0xb4,
	0x3f, 0xa0, 0x87, 0x3d, 0xad, 0xec, 0x99, 0xc9, 0x47, 0x13, 0x6f, 0xa4, 0xee, 0xed, 0xcd, 0xbc,
	0x8f, 0x9f, 0xe7, 0xfd, 0x0e, 0xf0, 0x7a, 0x84, 0xf4, 0x62, 0x6c, 0xa3, 0x90, 0xd9, 0xc2, 0x2c,
	0xac, 0xbe, 0x63, 0x33, 0x4c, 0xfb, 0x51, 0x80, 0x99, 0x1d, 0xa0, 0x24, 0x43, 0x51, 0x2f, 0xf5,
	0x03, 0x1a, 0x71, 0x4c, 0x23, 0x92, 0xfa, 0xd2, 0x67, 0x65, 0x94, 0x70, 0x02, 0x9b, 0xe2, 0x3b,
	0x0b, 0x85, 0xcc, 0x1a, 0x50, 0x58, 0x7d, 0xc7, 0x52, 0x14, 0x0d, 0xb7, 0x4a, 0x84, 0x62, 0x46,
	0x72, 0x3a, 0x5d, 0x45, 0xb0, 0x37, 0x3e, 0x53, 0xdf, 0x66, 0x91, 0x8d, 0xd2, 0x94, 0x70, 0xc4,
	0x23, 0x92, 0x32, 0xe9, 0xfd, 0x64, 0xc4, 0x1b, 0xc4, 0x11, 0x4e, 0xb9, 0x74, 0x7c, 0x31, 0xe2,
	0xb8, 0x88, 0x70, 0x1c, 0xfa, 0xe7, 0xf8, 0x37, 0xd4, 0x8f, 0x08, 0x95, 0x80, 0x4f, 0x47, 0x00,
	0x2a, 0x0c, 0xe9, 0x92, 0x09, 0xd9, 0xe5, 0xaf, 0xf3, 0xfc, 0x42, 0x12, 0x24, 0x88, 0x5d, 0x3e,
	0x92, 0xa5, 0x59, 0x60, 0x33, 0x8e, 0x78, 0x2e, 0xe3, 0x31, 0xfb, 0x60, 0x65, 0x0f, 0xf3, 0x1d,
	0x99, 0xcc, 0x8e, 0xca, 0xa5, 0x83, 0x7f, 0xcf, 0x31, 0xe3, 0xf0, 0x17, 0xb0, 0xa4, 0xb4, 0xfc,
	0x14, 0x25, 0xd8, 0xd0, 0x9a, 0xda, 0xfa, 0x42, 0xdb, 0xb9, 0xf7, 0xf4, 0x57, 0xde, 0x06, 0x68,
	0x0d, 0xcb, 0x27, 0xad, 0x2c, 0x62, 0x56, 0x40, 0x12, 0x7b, 0x92, 0x71, 0x51, 0x11, 0x1d, 0xa2,
	0x04, 0x9b, 0x0f, 0x1a, 0x58, 0x3d, 0xc8, 0x39, 0xe2, 0xf8, 0x11, 0x12, 0x29, 0xe9, 0x35, 0x50,
	0x0f, 0x72, 0xc6, 0x49, 0x82, 0xa9, 0x1f, 0x85, 0x52, 0xb8, 0x76, 0xef, 0xe9, 0x1d, 0xa0, 0xde,
	0x7f, 0x0a, 0x21, 0x02, 0x80, 0x64, 0x98, 0x8a, 0x1a, 0x1b, 0x7a, 0xb3, 0xb6, 0x5e, 0x77, 0xbe,
	0xb5, 0x66, 0x35, 0xd8, 0x9a, 0x08, 0xef, 0x48, 0x91, 0x48, 0x89, 0x21, 0x29, 0xfc, 0x0a, 0x7c,
	0x90, 0x21, 0xca, 0x23, 0x14, 0xfb, 0x17, 0x28, 0x8a, 0x73, 0x8a, 0x8d, 0x5a, 0x53, 0x5b, 0x7f,
	0xaf, 0xf3, 0xbe, 0x7c, 0xde, 0x15, 0xaf, 0xf0, 0x4b, 0xb0, 0xd4, 0x47, 0x71, 0x14, 0x22, 0x8e,
	0x7d, 0x92, 0xc6, 0x57, 0xc6, 0x5c, 0x09, 0x5b, 0x54, 0x8f, 0x47, 0x69, 0x7c, 0x65, 0xfe, 0xa3,
	0x83, 0x46, 0xb5, 0x3a, 0xdc, 0x02, 0xf5, 0x3c, 0x2b, 0x19, 0x8a, 0xee, 0x95, 0x0c, 0x75, 0xa7,
	0xa1, 0x12, 0x52, 0x0d, 0xb6, 0x76, 0x8b, 0x06, 0x1f, 0x20, 0x76, 0xd9, 0x01, 0x02, 0x5e, 0xd8,
	0xf0, 0x10, 0xcc, 0x07, 0x14, 0x23, 0x2e, 0xda, 0x54, 0x77, 0x36, 0x2b, 0x0b, 0x31, 0x98, 0xe3,
	0xc9, 0x4a, 0xec, 0xbf, 0xd3, 0x91, 0x2c, 0x05, 0x9f, 0x60, 0x37, 0xf4, 0xb7, 0xe3, 0x13, 0x2c,
	0xd0, 0x00, 0xf3, 0x14, 0x27, 0xa4, 0x2f, 0x0a, 0xb8, 0x50, 0x78, 0xc4, 0xef, 0x76, 0x1d, 0x2c,
	0x0c, 0x2a, 0x6e, 0xfe, 0xaf, 0x81, 0xcf, 0xab, 0x66, 0x83, 0x65, 0x24, 0x65, 0x18, 0xee, 0x82,
	0x8f, 0x1f, 0xf5, 0xc4, 0xc7, 0x94, 0x12, 0x5a, 0x12, 0xd7, 0x1d, 0xa8, 0x02, 0xa5, 0x59, 0x60,
	0x9d, 0x94, 0xf3, 0xde, 0xf9, 0x70, 0xbc, 0x5b, 0x3f, 0x16, 0x70, 0x78, 0x06, 0xde, 0xa5, 0x98,
	0xe5, 0x31, 0x57, 0xb3, 0xf3, 0xdd, 0xec, 0xd9, 0x99, 0x1a, 0x5a, 0xb1, 0x32, 0x05, 0x4f, 0x47,
	0xf1, 0x99, 0x3f, 0x54, 0x0c, 0xb8, 0x42, 0x16, 0xe3, 0x32, 0x65, 0xb7, 0xc6, 0xf7, 0xc4, 0xf9,
	0x6b, 0x0e, 0x18, 0x13, 0x04, 0x27, 0x22, 0x14, 0xf8, 0x5c, 0x03, 0x1f, 0x4d, 0xdb, 0x5e, 0xb8,
	0x3d, 0x3b, 0x8b, 0x37, 0x6c, 0x7d, 0xe3, 0x49, 0x7d, 0x36, 0x7f, 0xbe, 0xf3, 0xc6, 0x13, 0xfa,
	0xe3, 0xd9, 0x8b, 0xbf, 0xf5, 0x4d, 0xe8, 0x14, 0x87, 0xf3, 0x7a, 0xcc, 0xb3, 0xad, 0x36, 0x98,
	0xd9, 0xad, 0xc1, 0x25, 0x55, 0x4d, 0xb6, 0x5b, 0x37, 0xf0, 0x41, 0x03, 0xcb, 0xd3, 0x47, 0x00,
	0x3e, 0xb5, 0x43, 0xea, 0xb0, 0x34, 0xbe, 0x7f, 0x3a, 0x81, 0x98, 0x3e, 0xf3, 0xec, 0xce, 0x5b,
	0x1e, 0xb9, 0x4d, 0x1b, 0xc3, 0x63, 0x51, 0xa6, 0xec, 0x9a, 0xdf, 0x14, 0x29, 0x0f, 0x73, 0xbc,
	0x1e, 0x01, 0x6f, 0xb7, 0x6e, 0x26, 0x32, 0x76, 0x93, 0x52, 0xcf, 0xd5, 0x5a, 0x8d, 0x95, 0x5b,
	0xcf, 0xa8, 0x3a, 0xaa, 0xed, 0x3f, 0x75, 0xb0, 0x16, 0x90, 0x64, 0x66, 0xfc, 0xed, 0xd5, 0xaa,
	0x91, 0x39, 0x2e, 0x0e, 0xc8, 0xb1, 0xf6, 0xeb, 0xbe, 0xa4, 0xe8, 0x91, 0x18, 0xa5, 0x3d, 0x8b,
	0xd0, 0x9e, 0xdd, 0xc3, 0x69, 0x79, 0x5e, 0xec, 0xa1, 0x68, 0xf5, 0x9f, 0xec, 0x96, 0x32, 0xfe,
	0xd5, 0x6b, 0x7b, 0x9e, 0xf7, 0x9f, 0xde, 0xdc, 0x13, 0x84, 0x5e, 0xc8, 0x2c, 0x61, 0x16, 0xd6,
	0xa9, 0x63, 0x49, 0x61, 0x76, 0xab, 0x20, 0x5d, 0x2f, 0x64, 0xdd, 0x01, 0xa4, 0x7b, 0xea, 0x74,
	0x15, 0xe4, 0xa5, 0xbe, 0x26, 0xde, 0x5d, 0xd7, 0x0b, 0x99, 0xeb, 0x0e, 0x40, 0xae, 0x7b, 0xea,
	0xb8, 0xae, 0x82, 0x9d, 0xcf, 0x97, 0x71, 0x7e, 0xfd, 0x3a, 0x00, 0x00, 0xff, 0xff, 0x12, 0xe5,
	0xff, 0x09, 0x0b, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CampaignCriterionServiceClient is the client API for CampaignCriterionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignCriterionServiceClient interface {
	// Returns the requested criterion in full detail.
	GetCampaignCriterion(ctx context.Context, in *GetCampaignCriterionRequest, opts ...grpc.CallOption) (*resources.CampaignCriterion, error)
	// Creates, updates, or removes criteria. Operation statuses are returned.
	MutateCampaignCriteria(ctx context.Context, in *MutateCampaignCriteriaRequest, opts ...grpc.CallOption) (*MutateCampaignCriteriaResponse, error)
}

type campaignCriterionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignCriterionServiceClient(cc grpc.ClientConnInterface) CampaignCriterionServiceClient {
	return &campaignCriterionServiceClient{cc}
}

func (c *campaignCriterionServiceClient) GetCampaignCriterion(ctx context.Context, in *GetCampaignCriterionRequest, opts ...grpc.CallOption) (*resources.CampaignCriterion, error) {
	out := new(resources.CampaignCriterion)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.CampaignCriterionService/GetCampaignCriterion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignCriterionServiceClient) MutateCampaignCriteria(ctx context.Context, in *MutateCampaignCriteriaRequest, opts ...grpc.CallOption) (*MutateCampaignCriteriaResponse, error) {
	out := new(MutateCampaignCriteriaResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.CampaignCriterionService/MutateCampaignCriteria", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignCriterionServiceServer is the server API for CampaignCriterionService service.
type CampaignCriterionServiceServer interface {
	// Returns the requested criterion in full detail.
	GetCampaignCriterion(context.Context, *GetCampaignCriterionRequest) (*resources.CampaignCriterion, error)
	// Creates, updates, or removes criteria. Operation statuses are returned.
	MutateCampaignCriteria(context.Context, *MutateCampaignCriteriaRequest) (*MutateCampaignCriteriaResponse, error)
}

// UnimplementedCampaignCriterionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCampaignCriterionServiceServer struct {
}

func (*UnimplementedCampaignCriterionServiceServer) GetCampaignCriterion(ctx context.Context, req *GetCampaignCriterionRequest) (*resources.CampaignCriterion, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetCampaignCriterion not implemented")
}
func (*UnimplementedCampaignCriterionServiceServer) MutateCampaignCriteria(ctx context.Context, req *MutateCampaignCriteriaRequest) (*MutateCampaignCriteriaResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateCampaignCriteria not implemented")
}

func RegisterCampaignCriterionServiceServer(s *grpc.Server, srv CampaignCriterionServiceServer) {
	s.RegisterService(&_CampaignCriterionService_serviceDesc, srv)
}

func _CampaignCriterionService_GetCampaignCriterion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignCriterionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignCriterionServiceServer).GetCampaignCriterion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.CampaignCriterionService/GetCampaignCriterion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignCriterionServiceServer).GetCampaignCriterion(ctx, req.(*GetCampaignCriterionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignCriterionService_MutateCampaignCriteria_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignCriteriaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignCriterionServiceServer).MutateCampaignCriteria(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.CampaignCriterionService/MutateCampaignCriteria",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignCriterionServiceServer).MutateCampaignCriteria(ctx, req.(*MutateCampaignCriteriaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignCriterionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.CampaignCriterionService",
	HandlerType: (*CampaignCriterionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignCriterion",
			Handler:    _CampaignCriterionService_GetCampaignCriterion_Handler,
		},
		{
			MethodName: "MutateCampaignCriteria",
			Handler:    _CampaignCriterionService_MutateCampaignCriteria_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/campaign_criterion_service.proto",
}