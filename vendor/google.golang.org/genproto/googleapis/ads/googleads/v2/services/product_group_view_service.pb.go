// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/product_group_view_service.proto

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

// Request message for [ProductGroupViewService.GetProductGroupView][google.ads.googleads.v2.services.ProductGroupViewService.GetProductGroupView].
type GetProductGroupViewRequest struct {
	// The resource name of the product group view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductGroupViewRequest) Reset()         { *m = GetProductGroupViewRequest{} }
func (m *GetProductGroupViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductGroupViewRequest) ProtoMessage()    {}
func (*GetProductGroupViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd387eda7160e6d, []int{0}
}

func (m *GetProductGroupViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductGroupViewRequest.Unmarshal(m, b)
}
func (m *GetProductGroupViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductGroupViewRequest.Marshal(b, m, deterministic)
}
func (m *GetProductGroupViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductGroupViewRequest.Merge(m, src)
}
func (m *GetProductGroupViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductGroupViewRequest.Size(m)
}
func (m *GetProductGroupViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductGroupViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductGroupViewRequest proto.InternalMessageInfo

func (m *GetProductGroupViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetProductGroupViewRequest)(nil), "google.ads.googleads.v2.services.GetProductGroupViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/product_group_view_service.proto", fileDescriptor_bdd387eda7160e6d)
}

var fileDescriptor_bdd387eda7160e6d = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xbf, 0x4a, 0xfb, 0x40,
	0x1c, 0x27, 0xf9, 0xc1, 0x0f, 0x0c, 0xba, 0xc4, 0xa1, 0x25, 0x76, 0x28, 0xb5, 0x83, 0x74, 0xb8,
	0x83, 0x14, 0x97, 0xab, 0x0e, 0xd7, 0x25, 0x4e, 0x52, 0x2a, 0x64, 0x90, 0x40, 0x88, 0xc9, 0x11,
	0x02, 0x4d, 0x2e, 0xde, 0x5d, 0xd2, 0x41, 0x5c, 0xf4, 0x11, 0x7c, 0x03, 0x47, 0xdf, 0xc3, 0xa5,
	0xab, 0xaf, 0xe0, 0xe4, 0x43, 0x88, 0xa4, 0x97, 0x4b, 0xb5, 0x34, 0x74, 0xfb, 0x70, 0xdf, 0xcf,
	0x9f, 0xfb, 0x7e, 0xee, 0x0c, 0x1c, 0x53, 0x1a, 0x2f, 0x08, 0x0c, 0x22, 0x0e, 0x25, 0xac, 0x50,
	0x69, 0x43, 0x4e, 0x58, 0x99, 0x84, 0x84, 0xc3, 0x9c, 0xd1, 0xa8, 0x08, 0x85, 0x1f, 0x33, 0x5a,
	0xe4, 0x7e, 0x99, 0x90, 0xa5, 0x5f, 0xcf, 0x40, 0xce, 0xa8, 0xa0, 0x66, 0x5f, 0xea, 0x40, 0x10,
	0x71, 0xd0, 0x58, 0x80, 0xd2, 0x06, 0xca, 0xc2, 0x42, 0x6d, 0x21, 0x8c, 0x70, 0x5a, 0xb0, 0xdd,
	0x29, 0xd2, 0xdd, 0xea, 0x29, 0x6d, 0x9e, 0xc0, 0x20, 0xcb, 0xa8, 0x08, 0x44, 0x42, 0x33, 0x5e,
	0x4f, 0x3b, 0xbf, 0xa6, 0xe1, 0x22, 0x21, 0x99, 0x90, 0x83, 0x01, 0x36, 0x2c, 0x87, 0x88, 0x99,
	0x74, 0x75, 0x2a, 0x53, 0x37, 0x21, 0xcb, 0x39, 0xb9, 0x2f, 0x08, 0x17, 0xe6, 0xa9, 0x71, 0xa4,
	0xa2, 0xfd, 0x2c, 0x48, 0x49, 0x57, 0xeb, 0x6b, 0x67, 0x07, 0xf3, 0x43, 0x75, 0x78, 0x1d, 0xa4,
	0xc4, 0xfe, 0xd6, 0x8c, 0xce, 0xb6, 0xc1, 0x8d, 0x5c, 0xc9, 0x7c, 0xd7, 0x8c, 0xe3, 0x1d, 0xfe,
	0xe6, 0x05, 0xd8, 0x57, 0x06, 0x68, 0xbf, 0x96, 0x35, 0x6e, 0x55, 0x37, 0x45, 0x81, 0x6d, 0xed,
	0x60, 0xf2, 0xf4, 0xf1, 0xf9, 0xa2, 0x9f, 0x9b, 0xe3, 0xaa, 0xd0, 0x87, 0x3f, 0x6b, 0x5d, 0x86,
	0x05, 0x17, 0x34, 0x25, 0x8c, 0xc3, 0x91, 0x6a, 0xb8, 0x11, 0x72, 0x38, 0x7a, 0xb4, 0x4e, 0x56,
	0xb8, 0xbb, 0x09, 0xaa, 0x51, 0x9e, 0x70, 0x10, 0xd2, 0x74, 0xfa, 0xac, 0x1b, 0xc3, 0x90, 0xa6,
	0x7b, 0x57, 0x9a, 0xf6, 0x5a, 0x6a, 0x9a, 0x55, 0x4f, 0x31, 0xd3, 0x6e, 0xaf, 0x6a, 0x87, 0x98,
	0x2e, 0x82, 0x2c, 0x06, 0x94, 0xc5, 0x30, 0x26, 0xd9, 0xfa, 0xa1, 0xe0, 0x26, 0xb3, 0xfd, 0x0f,
	0x4e, 0x14, 0x78, 0xd5, 0xff, 0x39, 0x18, 0xbf, 0xe9, 0x7d, 0x47, 0x1a, 0xe2, 0x88, 0x03, 0x09,
	0x2b, 0xe4, 0xda, 0xa0, 0x0e, 0xe6, 0x2b, 0x45, 0xf1, 0x70, 0xc4, 0xbd, 0x86, 0xe2, 0xb9, 0xb6,
	0xa7, 0x28, 0x5f, 0xfa, 0x50, 0x9e, 0x23, 0x84, 0x23, 0x8e, 0x50, 0x43, 0x42, 0xc8, 0xb5, 0x11,
	0x52, 0xb4, 0xbb, 0xff, 0xeb, 0x7b, 0x8e, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x79, 0x4a, 0x6b,
	0xc0, 0x2a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductGroupViewServiceClient is the client API for ProductGroupViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductGroupViewServiceClient interface {
	// Returns the requested product group view in full detail.
	GetProductGroupView(ctx context.Context, in *GetProductGroupViewRequest, opts ...grpc.CallOption) (*resources.ProductGroupView, error)
}

type productGroupViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductGroupViewServiceClient(cc *grpc.ClientConn) ProductGroupViewServiceClient {
	return &productGroupViewServiceClient{cc}
}

func (c *productGroupViewServiceClient) GetProductGroupView(ctx context.Context, in *GetProductGroupViewRequest, opts ...grpc.CallOption) (*resources.ProductGroupView, error) {
	out := new(resources.ProductGroupView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.ProductGroupViewService/GetProductGroupView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductGroupViewServiceServer is the server API for ProductGroupViewService service.
type ProductGroupViewServiceServer interface {
	// Returns the requested product group view in full detail.
	GetProductGroupView(context.Context, *GetProductGroupViewRequest) (*resources.ProductGroupView, error)
}

// UnimplementedProductGroupViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductGroupViewServiceServer struct {
}

func (*UnimplementedProductGroupViewServiceServer) GetProductGroupView(ctx context.Context, req *GetProductGroupViewRequest) (*resources.ProductGroupView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductGroupView not implemented")
}

func RegisterProductGroupViewServiceServer(s *grpc.Server, srv ProductGroupViewServiceServer) {
	s.RegisterService(&_ProductGroupViewService_serviceDesc, srv)
}

func _ProductGroupViewService_GetProductGroupView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductGroupViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductGroupViewServiceServer).GetProductGroupView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.ProductGroupViewService/GetProductGroupView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductGroupViewServiceServer).GetProductGroupView(ctx, req.(*GetProductGroupViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductGroupViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.ProductGroupViewService",
	HandlerType: (*ProductGroupViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductGroupView",
			Handler:    _ProductGroupViewService_GetProductGroupView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/product_group_view_service.proto",
}
