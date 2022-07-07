// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/user_interest_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request message for [UserInterestService.GetUserInterest][google.ads.googleads.v2.services.UserInterestService.GetUserInterest].
type GetUserInterestRequest struct {
	// Required. Resource name of the UserInterest to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInterestRequest) Reset()         { *m = GetUserInterestRequest{} }
func (m *GetUserInterestRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserInterestRequest) ProtoMessage()    {}
func (*GetUserInterestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9d082b840ef8dbf, []int{0}
}

func (m *GetUserInterestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInterestRequest.Unmarshal(m, b)
}
func (m *GetUserInterestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInterestRequest.Marshal(b, m, deterministic)
}
func (m *GetUserInterestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInterestRequest.Merge(m, src)
}
func (m *GetUserInterestRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserInterestRequest.Size(m)
}
func (m *GetUserInterestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInterestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInterestRequest proto.InternalMessageInfo

func (m *GetUserInterestRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserInterestRequest)(nil), "google.ads.googleads.v2.services.GetUserInterestRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/user_interest_service.proto", fileDescriptor_b9d082b840ef8dbf)
}

var fileDescriptor_b9d082b840ef8dbf = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xbf, 0x8a, 0xd4, 0x40,
	0x18, 0x27, 0x11, 0x04, 0x83, 0x22, 0x44, 0xd0, 0x75, 0x15, 0x5c, 0x8e, 0x13, 0xe5, 0xc0, 0x19,
	0x8c, 0x08, 0x32, 0x6a, 0x31, 0xdb, 0xac, 0x36, 0x72, 0x9c, 0xb8, 0x85, 0x2c, 0x84, 0xb9, 0xe4,
	0x33, 0x0e, 0x24, 0x33, 0xeb, 0x7c, 0x93, 0x34, 0x62, 0xe3, 0x2b, 0xf8, 0x06, 0x96, 0xbe, 0x87,
	0xcd, 0xb5, 0x76, 0x56, 0x16, 0x56, 0x3e, 0xc2, 0x35, 0x4a, 0x76, 0x32, 0xb9, 0xac, 0x77, 0xe1,
	0xba, 0x1f, 0xf3, 0xfb, 0xf3, 0xfd, 0x9b, 0xe8, 0x59, 0xa1, 0x75, 0x51, 0x02, 0x15, 0x39, 0x52,
	0x07, 0x5b, 0xd4, 0x24, 0x14, 0xc1, 0x34, 0x32, 0x03, 0xa4, 0x35, 0x82, 0x49, 0xa5, 0xb2, 0x60,
	0x00, 0x6d, 0xda, 0x3d, 0x93, 0xb5, 0xd1, 0x56, 0xc7, 0x33, 0x67, 0x21, 0x22, 0x47, 0xd2, 0xbb,
	0x49, 0x93, 0x10, 0xef, 0x9e, 0x3e, 0x1e, 0xcb, 0x37, 0x80, 0xba, 0x36, 0xa7, 0x0a, 0xb8, 0xe0,
	0xe9, 0x6d, 0x6f, 0x5b, 0x4b, 0x2a, 0x94, 0xd2, 0x56, 0x58, 0xa9, 0x15, 0x76, 0xec, 0x8d, 0x01,
	0x9b, 0x95, 0x12, 0x94, 0xb7, 0xdd, 0x19, 0x10, 0xef, 0x24, 0x94, 0x79, 0x7a, 0x08, 0xef, 0x45,
	0x23, 0xb5, 0xe9, 0x04, 0x37, 0x07, 0x02, 0xdf, 0x81, 0xa3, 0x76, 0xca, 0xe8, 0xfa, 0x02, 0xec,
	0x1b, 0x04, 0xf3, 0xb2, 0xeb, 0xe5, 0x00, 0x3e, 0xd4, 0x80, 0x36, 0x3e, 0x88, 0xae, 0x78, 0x6d,
	0xaa, 0x44, 0x05, 0x93, 0x60, 0x16, 0xdc, 0xbf, 0x34, 0x7f, 0xf0, 0x8b, 0x87, 0xc7, 0xfc, 0x5e,
	0x74, 0xf7, 0x64, 0xf2, 0x0e, 0xad, 0x25, 0x92, 0x4c, 0x57, 0x74, 0x2b, 0xec, 0xb2, 0xcf, 0x78,
	0x25, 0x2a, 0x48, 0x8e, 0x83, 0xe8, 0xda, 0x90, 0x7e, 0xed, 0x16, 0x16, 0x7f, 0x0f, 0xa2, 0xab,
	0xff, 0xb5, 0x11, 0x3f, 0x21, 0xe7, 0xad, 0x99, 0x9c, 0xdd, 0xf9, 0x94, 0x8e, 0x3a, 0xfb, 0xf5,
	0x93, 0xa1, 0x6f, 0x67, 0xf1, 0x93, 0x6f, 0xcf, 0xfa, 0xf9, 0xc7, 0xef, 0x2f, 0xe1, 0xc3, 0x98,
	0xb6, 0x27, 0xfb, 0xb8, 0xc5, 0x3c, 0xcf, 0x6a, 0xb4, 0xba, 0x02, 0x83, 0x74, 0x6f, 0x73, 0x43,
	0x1f, 0x82, 0x74, 0xef, 0xd3, 0xf4, 0xd6, 0x11, 0x9f, 0x8c, 0xed, 0x65, 0xfe, 0x37, 0x88, 0x76,
	0x33, 0x5d, 0x9d, 0x3b, 0xd6, 0x7c, 0x72, 0xc6, 0x8a, 0xf6, 0xdb, 0x6b, 0xed, 0x07, 0x6f, 0x5f,
	0x74, 0xee, 0x42, 0x97, 0x42, 0x15, 0x44, 0x9b, 0x82, 0x16, 0xa0, 0x36, 0xb7, 0xa4, 0x27, 0xf5,
	0xc6, 0x3f, 0xf6, 0x53, 0x0f, 0xbe, 0x86, 0x17, 0x16, 0x9c, 0x7f, 0x0b, 0x67, 0x0b, 0x17, 0xc8,
	0x73, 0x24, 0x0e, 0xb6, 0x68, 0x99, 0x90, 0xae, 0x30, 0x1e, 0x79, 0xc9, 0x8a, 0xe7, 0xb8, 0xea,
	0x25, 0xab, 0x65, 0xb2, 0xf2, 0x92, 0x3f, 0xe1, 0xae, 0x7b, 0x67, 0x8c, 0xe7, 0xc8, 0x58, 0x2f,
	0x62, 0x6c, 0x99, 0x30, 0xe6, 0x65, 0x87, 0x17, 0x37, 0x7d, 0x3e, 0xfa, 0x17, 0x00, 0x00, 0xff,
	0xff, 0x19, 0x93, 0x48, 0xae, 0x7f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserInterestServiceClient is the client API for UserInterestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserInterestServiceClient interface {
	// Returns the requested user interest in full detail
	GetUserInterest(ctx context.Context, in *GetUserInterestRequest, opts ...grpc.CallOption) (*resources.UserInterest, error)
}

type userInterestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserInterestServiceClient(cc grpc.ClientConnInterface) UserInterestServiceClient {
	return &userInterestServiceClient{cc}
}

func (c *userInterestServiceClient) GetUserInterest(ctx context.Context, in *GetUserInterestRequest, opts ...grpc.CallOption) (*resources.UserInterest, error) {
	out := new(resources.UserInterest)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.UserInterestService/GetUserInterest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInterestServiceServer is the server API for UserInterestService service.
type UserInterestServiceServer interface {
	// Returns the requested user interest in full detail
	GetUserInterest(context.Context, *GetUserInterestRequest) (*resources.UserInterest, error)
}

// UnimplementedUserInterestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserInterestServiceServer struct {
}

func (*UnimplementedUserInterestServiceServer) GetUserInterest(ctx context.Context, req *GetUserInterestRequest) (*resources.UserInterest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInterest not implemented")
}

func RegisterUserInterestServiceServer(s *grpc.Server, srv UserInterestServiceServer) {
	s.RegisterService(&_UserInterestService_serviceDesc, srv)
}

func _UserInterestService_GetUserInterest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInterestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInterestServiceServer).GetUserInterest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.UserInterestService/GetUserInterest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInterestServiceServer).GetUserInterest(ctx, req.(*GetUserInterestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserInterestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.UserInterestService",
	HandlerType: (*UserInterestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInterest",
			Handler:    _UserInterestService_GetUserInterest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/user_interest_service.proto",
}