// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/tap/v3alpha/tap.proto

package envoy_service_tap_v3alpha

import (
	context "context"
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/data/tap/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type StreamTapsRequest struct {
	Identifier           *StreamTapsRequest_Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	TraceId              uint64                        `protobuf:"varint,2,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	Trace                *v3alpha.TraceWrapper         `protobuf:"bytes,3,opt,name=trace,proto3" json:"trace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *StreamTapsRequest) Reset()         { *m = StreamTapsRequest{} }
func (m *StreamTapsRequest) String() string { return proto.CompactTextString(m) }
func (*StreamTapsRequest) ProtoMessage()    {}
func (*StreamTapsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8837c21a05e91cb, []int{0}
}

func (m *StreamTapsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTapsRequest.Unmarshal(m, b)
}
func (m *StreamTapsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTapsRequest.Marshal(b, m, deterministic)
}
func (m *StreamTapsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTapsRequest.Merge(m, src)
}
func (m *StreamTapsRequest) XXX_Size() int {
	return xxx_messageInfo_StreamTapsRequest.Size(m)
}
func (m *StreamTapsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTapsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTapsRequest proto.InternalMessageInfo

func (m *StreamTapsRequest) GetIdentifier() *StreamTapsRequest_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *StreamTapsRequest) GetTraceId() uint64 {
	if m != nil {
		return m.TraceId
	}
	return 0
}

func (m *StreamTapsRequest) GetTrace() *v3alpha.TraceWrapper {
	if m != nil {
		return m.Trace
	}
	return nil
}

type StreamTapsRequest_Identifier struct {
	Node                 *core.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	TapId                string     `protobuf:"bytes,2,opt,name=tap_id,json=tapId,proto3" json:"tap_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StreamTapsRequest_Identifier) Reset()         { *m = StreamTapsRequest_Identifier{} }
func (m *StreamTapsRequest_Identifier) String() string { return proto.CompactTextString(m) }
func (*StreamTapsRequest_Identifier) ProtoMessage()    {}
func (*StreamTapsRequest_Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8837c21a05e91cb, []int{0, 0}
}

func (m *StreamTapsRequest_Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTapsRequest_Identifier.Unmarshal(m, b)
}
func (m *StreamTapsRequest_Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTapsRequest_Identifier.Marshal(b, m, deterministic)
}
func (m *StreamTapsRequest_Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTapsRequest_Identifier.Merge(m, src)
}
func (m *StreamTapsRequest_Identifier) XXX_Size() int {
	return xxx_messageInfo_StreamTapsRequest_Identifier.Size(m)
}
func (m *StreamTapsRequest_Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTapsRequest_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTapsRequest_Identifier proto.InternalMessageInfo

func (m *StreamTapsRequest_Identifier) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *StreamTapsRequest_Identifier) GetTapId() string {
	if m != nil {
		return m.TapId
	}
	return ""
}

type StreamTapsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamTapsResponse) Reset()         { *m = StreamTapsResponse{} }
func (m *StreamTapsResponse) String() string { return proto.CompactTextString(m) }
func (*StreamTapsResponse) ProtoMessage()    {}
func (*StreamTapsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8837c21a05e91cb, []int{1}
}

func (m *StreamTapsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTapsResponse.Unmarshal(m, b)
}
func (m *StreamTapsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTapsResponse.Marshal(b, m, deterministic)
}
func (m *StreamTapsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTapsResponse.Merge(m, src)
}
func (m *StreamTapsResponse) XXX_Size() int {
	return xxx_messageInfo_StreamTapsResponse.Size(m)
}
func (m *StreamTapsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTapsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTapsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StreamTapsRequest)(nil), "envoy.service.tap.v3alpha.StreamTapsRequest")
	proto.RegisterType((*StreamTapsRequest_Identifier)(nil), "envoy.service.tap.v3alpha.StreamTapsRequest.Identifier")
	proto.RegisterType((*StreamTapsResponse)(nil), "envoy.service.tap.v3alpha.StreamTapsResponse")
}

func init() {
	proto.RegisterFile("envoy/service/tap/v3alpha/tap.proto", fileDescriptor_f8837c21a05e91cb)
}

var fileDescriptor_f8837c21a05e91cb = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0x02, 0x41,
	0x18, 0xc6, 0x1b, 0x53, 0xb3, 0x09, 0xa2, 0x86, 0x22, 0x5d, 0x3a, 0x98, 0x09, 0x79, 0xa8, 0x5d,
	0xd0, 0x43, 0xe1, 0x71, 0x6f, 0x5e, 0x42, 0xd6, 0x05, 0x8f, 0xf2, 0xea, 0xbc, 0xd1, 0x90, 0xee,
	0x4c, 0xb3, 0xd3, 0x96, 0xa7, 0xae, 0xd1, 0xe7, 0xe8, 0x53, 0x76, 0x8a, 0xdd, 0x59, 0xff, 0x11,
	0x42, 0xdd, 0xde, 0x61, 0x7e, 0xcf, 0xf3, 0xbc, 0x7f, 0xe8, 0x25, 0x46, 0x89, 0x9c, 0x7b, 0x31,
	0xea, 0x44, 0x4c, 0xd0, 0x33, 0xa0, 0xbc, 0xa4, 0x03, 0x53, 0xf5, 0x08, 0x69, 0xed, 0x2a, 0x2d,
	0x8d, 0x64, 0xb5, 0x0c, 0x72, 0x73, 0xc8, 0x4d, 0x3f, 0x72, 0xc8, 0xb9, 0xb0, 0x7a, 0x50, 0x62,
	0xa9, 0x9b, 0x48, 0x8d, 0xde, 0x18, 0x62, 0xb4, 0x6a, 0xa7, 0x69, 0x11, 0x0e, 0x06, 0x36, 0xfc,
	0x5f, 0x35, 0x28, 0x85, 0x3a, 0xa7, 0xce, 0x12, 0x98, 0x0a, 0x0e, 0x06, 0xbd, 0x45, 0x61, 0x3f,
	0x1a, 0x5f, 0x05, 0x7a, 0x3c, 0x30, 0x1a, 0x61, 0x16, 0x82, 0x8a, 0x03, 0x7c, 0x7e, 0xc1, 0xd8,
	0xb0, 0x21, 0xa5, 0x82, 0x63, 0x64, 0xc4, 0x83, 0x40, 0x5d, 0x25, 0x75, 0xd2, 0x3a, 0x68, 0xdf,
	0xba, 0x5b, 0xfb, 0x74, 0x7f, 0x39, 0xb8, 0xbd, 0xa5, 0x3c, 0x58, 0xb3, 0x62, 0x35, 0x5a, 0x31,
	0x1a, 0x26, 0x38, 0x12, 0xbc, 0x5a, 0xa8, 0x93, 0x56, 0x31, 0xd8, 0xcb, 0xde, 0x3d, 0xce, 0xba,
	0xb4, 0x94, 0x95, 0xd5, 0xdd, 0x2c, 0xae, 0x99, 0xc7, 0xa5, 0x83, 0x6d, 0x64, 0x85, 0x29, 0x34,
	0xb4, 0xd3, 0x05, 0x56, 0xe2, 0x8c, 0x28, 0x5d, 0x05, 0xb2, 0x2e, 0x2d, 0x46, 0x92, 0x63, 0xde,
	0xf7, 0x79, 0x6e, 0x04, 0x4a, 0x2c, 0x3d, 0xd2, 0x25, 0xba, 0xf7, 0x92, 0xa3, 0x5f, 0xf9, 0xf6,
	0x4b, 0x9f, 0xa4, 0x70, 0x44, 0x82, 0x4c, 0xc3, 0x4e, 0x69, 0xd9, 0x80, 0x5a, 0xb4, 0xb7, 0x1f,
	0x94, 0x0c, 0xa8, 0x1e, 0x6f, 0x9c, 0x50, 0xb6, 0x3e, 0x63, 0xac, 0x64, 0x14, 0x63, 0xfb, 0x9d,
	0x1e, 0x86, 0xa0, 0x06, 0x22, 0x7a, 0x1a, 0xd8, 0xa5, 0xb0, 0x19, 0xa5, 0x2b, 0x8e, 0x5d, 0xff,
	0x67, 0x65, 0xce, 0xcd, 0x1f, 0x69, 0x1b, 0xde, 0xd8, 0x69, 0x11, 0xff, 0x8e, 0x5e, 0x09, 0x69,
	0x65, 0x4a, 0xcb, 0xb7, 0xf9, 0x76, 0x07, 0xbf, 0x12, 0x82, 0xea, 0xa7, 0x27, 0xef, 0x93, 0x0f,
	0x42, 0xc6, 0xe5, 0xec, 0xfc, 0x9d, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xda, 0x21, 0x27,
	0xa2, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TapSinkServiceClient is the client API for TapSinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TapSinkServiceClient interface {
	StreamTaps(ctx context.Context, opts ...grpc.CallOption) (TapSinkService_StreamTapsClient, error)
}

type tapSinkServiceClient struct {
	cc *grpc.ClientConn
}

func NewTapSinkServiceClient(cc *grpc.ClientConn) TapSinkServiceClient {
	return &tapSinkServiceClient{cc}
}

func (c *tapSinkServiceClient) StreamTaps(ctx context.Context, opts ...grpc.CallOption) (TapSinkService_StreamTapsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TapSinkService_serviceDesc.Streams[0], "/envoy.service.tap.v3alpha.TapSinkService/StreamTaps", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapSinkServiceStreamTapsClient{stream}
	return x, nil
}

type TapSinkService_StreamTapsClient interface {
	Send(*StreamTapsRequest) error
	CloseAndRecv() (*StreamTapsResponse, error)
	grpc.ClientStream
}

type tapSinkServiceStreamTapsClient struct {
	grpc.ClientStream
}

func (x *tapSinkServiceStreamTapsClient) Send(m *StreamTapsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tapSinkServiceStreamTapsClient) CloseAndRecv() (*StreamTapsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamTapsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TapSinkServiceServer is the server API for TapSinkService service.
type TapSinkServiceServer interface {
	StreamTaps(TapSinkService_StreamTapsServer) error
}

// UnimplementedTapSinkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTapSinkServiceServer struct {
}

func (*UnimplementedTapSinkServiceServer) StreamTaps(srv TapSinkService_StreamTapsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTaps not implemented")
}

func RegisterTapSinkServiceServer(s *grpc.Server, srv TapSinkServiceServer) {
	s.RegisterService(&_TapSinkService_serviceDesc, srv)
}

func _TapSinkService_StreamTaps_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TapSinkServiceServer).StreamTaps(&tapSinkServiceStreamTapsServer{stream})
}

type TapSinkService_StreamTapsServer interface {
	SendAndClose(*StreamTapsResponse) error
	Recv() (*StreamTapsRequest, error)
	grpc.ServerStream
}

type tapSinkServiceStreamTapsServer struct {
	grpc.ServerStream
}

func (x *tapSinkServiceStreamTapsServer) SendAndClose(m *StreamTapsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tapSinkServiceStreamTapsServer) Recv() (*StreamTapsRequest, error) {
	m := new(StreamTapsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TapSinkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.tap.v3alpha.TapSinkService",
	HandlerType: (*TapSinkServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTaps",
			Handler:       _TapSinkService_StreamTaps_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/tap/v3alpha/tap.proto",
}
