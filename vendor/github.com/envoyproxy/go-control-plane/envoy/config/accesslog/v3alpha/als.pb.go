// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/accesslog/v3alpha/als.proto

package envoy_config_accesslog_v3alpha

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type HttpGrpcAccessLogConfig struct {
	CommonConfig                    *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	AdditionalRequestHeadersToLog   []string                   `protobuf:"bytes,2,rep,name=additional_request_headers_to_log,json=additionalRequestHeadersToLog,proto3" json:"additional_request_headers_to_log,omitempty"`
	AdditionalResponseHeadersToLog  []string                   `protobuf:"bytes,3,rep,name=additional_response_headers_to_log,json=additionalResponseHeadersToLog,proto3" json:"additional_response_headers_to_log,omitempty"`
	AdditionalResponseTrailersToLog []string                   `protobuf:"bytes,4,rep,name=additional_response_trailers_to_log,json=additionalResponseTrailersToLog,proto3" json:"additional_response_trailers_to_log,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                   `json:"-"`
	XXX_unrecognized                []byte                     `json:"-"`
	XXX_sizecache                   int32                      `json:"-"`
}

func (m *HttpGrpcAccessLogConfig) Reset()         { *m = HttpGrpcAccessLogConfig{} }
func (m *HttpGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*HttpGrpcAccessLogConfig) ProtoMessage()    {}
func (*HttpGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_938e3858beb7bdc4, []int{0}
}

func (m *HttpGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpGrpcAccessLogConfig.Unmarshal(m, b)
}
func (m *HttpGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpGrpcAccessLogConfig.Marshal(b, m, deterministic)
}
func (m *HttpGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpGrpcAccessLogConfig.Merge(m, src)
}
func (m *HttpGrpcAccessLogConfig) XXX_Size() int {
	return xxx_messageInfo_HttpGrpcAccessLogConfig.Size(m)
}
func (m *HttpGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HttpGrpcAccessLogConfig proto.InternalMessageInfo

func (m *HttpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalRequestHeadersToLog() []string {
	if m != nil {
		return m.AdditionalRequestHeadersToLog
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalResponseHeadersToLog() []string {
	if m != nil {
		return m.AdditionalResponseHeadersToLog
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalResponseTrailersToLog() []string {
	if m != nil {
		return m.AdditionalResponseTrailersToLog
	}
	return nil
}

type TcpGrpcAccessLogConfig struct {
	CommonConfig         *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *TcpGrpcAccessLogConfig) Reset()         { *m = TcpGrpcAccessLogConfig{} }
func (m *TcpGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*TcpGrpcAccessLogConfig) ProtoMessage()    {}
func (*TcpGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_938e3858beb7bdc4, []int{1}
}

func (m *TcpGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpGrpcAccessLogConfig.Unmarshal(m, b)
}
func (m *TcpGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpGrpcAccessLogConfig.Marshal(b, m, deterministic)
}
func (m *TcpGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpGrpcAccessLogConfig.Merge(m, src)
}
func (m *TcpGrpcAccessLogConfig) XXX_Size() int {
	return xxx_messageInfo_TcpGrpcAccessLogConfig.Size(m)
}
func (m *TcpGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TcpGrpcAccessLogConfig proto.InternalMessageInfo

func (m *TcpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

type CommonGrpcAccessLogConfig struct {
	LogName                 string                `protobuf:"bytes,1,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	GrpcService             *core.GrpcService     `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	BufferFlushInterval     *duration.Duration    `protobuf:"bytes,3,opt,name=buffer_flush_interval,json=bufferFlushInterval,proto3" json:"buffer_flush_interval,omitempty"`
	BufferSizeBytes         *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=buffer_size_bytes,json=bufferSizeBytes,proto3" json:"buffer_size_bytes,omitempty"`
	FilterStateObjectsToLog []string              `protobuf:"bytes,5,rep,name=filter_state_objects_to_log,json=filterStateObjectsToLog,proto3" json:"filter_state_objects_to_log,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}              `json:"-"`
	XXX_unrecognized        []byte                `json:"-"`
	XXX_sizecache           int32                 `json:"-"`
}

func (m *CommonGrpcAccessLogConfig) Reset()         { *m = CommonGrpcAccessLogConfig{} }
func (m *CommonGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*CommonGrpcAccessLogConfig) ProtoMessage()    {}
func (*CommonGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_938e3858beb7bdc4, []int{2}
}

func (m *CommonGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonGrpcAccessLogConfig.Unmarshal(m, b)
}
func (m *CommonGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonGrpcAccessLogConfig.Marshal(b, m, deterministic)
}
func (m *CommonGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonGrpcAccessLogConfig.Merge(m, src)
}
func (m *CommonGrpcAccessLogConfig) XXX_Size() int {
	return xxx_messageInfo_CommonGrpcAccessLogConfig.Size(m)
}
func (m *CommonGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommonGrpcAccessLogConfig proto.InternalMessageInfo

func (m *CommonGrpcAccessLogConfig) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

func (m *CommonGrpcAccessLogConfig) GetGrpcService() *core.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func (m *CommonGrpcAccessLogConfig) GetBufferFlushInterval() *duration.Duration {
	if m != nil {
		return m.BufferFlushInterval
	}
	return nil
}

func (m *CommonGrpcAccessLogConfig) GetBufferSizeBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.BufferSizeBytes
	}
	return nil
}

func (m *CommonGrpcAccessLogConfig) GetFilterStateObjectsToLog() []string {
	if m != nil {
		return m.FilterStateObjectsToLog
	}
	return nil
}

func init() {
	proto.RegisterType((*HttpGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v3alpha.HttpGrpcAccessLogConfig")
	proto.RegisterType((*TcpGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v3alpha.TcpGrpcAccessLogConfig")
	proto.RegisterType((*CommonGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v3alpha.CommonGrpcAccessLogConfig")
}

func init() {
	proto.RegisterFile("envoy/config/accesslog/v3alpha/als.proto", fileDescriptor_938e3858beb7bdc4)
}

var fileDescriptor_938e3858beb7bdc4 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0x41, 0x6f, 0xd3, 0x3c,
	0x18, 0xc7, 0xdf, 0xa4, 0xdd, 0xbb, 0xce, 0x1b, 0x02, 0x82, 0xa0, 0x5d, 0x81, 0x52, 0xba, 0x4b,
	0x41, 0x28, 0x91, 0xd6, 0x13, 0x88, 0xcb, 0x32, 0x04, 0x1d, 0xaa, 0xa0, 0x4a, 0x0b, 0x1c, 0x23,
	0x37, 0x7d, 0x92, 0x1a, 0xb9, 0x79, 0x82, 0xed, 0x14, 0xba, 0x23, 0x47, 0x3e, 0x0a, 0x5f, 0x8c,
	0x3b, 0xc7, 0x9d, 0x50, 0xec, 0xac, 0xeb, 0xd8, 0x06, 0x47, 0x6e, 0x6d, 0xfc, 0x7b, 0x7e, 0xfe,
	0xdb, 0xfe, 0x93, 0x2e, 0xa4, 0x0b, 0x5c, 0x7a, 0x11, 0xa6, 0x31, 0x4b, 0x3c, 0x1a, 0x45, 0x20,
	0x25, 0xc7, 0xc4, 0x5b, 0xf4, 0x28, 0xcf, 0x66, 0xd4, 0xa3, 0x5c, 0xba, 0x99, 0x40, 0x85, 0x4e,
	0x4b, 0x93, 0xae, 0x21, 0xdd, 0x15, 0xe9, 0x96, 0x64, 0xf3, 0x91, 0x31, 0xd1, 0x8c, 0xad, 0x86,
	0x23, 0x14, 0xe0, 0x25, 0x22, 0x8b, 0x42, 0x09, 0x62, 0xc1, 0x22, 0x30, 0xaa, 0x66, 0x2b, 0x41,
	0x4c, 0x38, 0x78, 0xfa, 0xdf, 0x24, 0x8f, 0xbd, 0x69, 0x2e, 0xa8, 0x62, 0x98, 0x5e, 0xb5, 0xfe,
	0x59, 0xd0, 0x2c, 0x03, 0x51, 0x46, 0x69, 0xd6, 0x17, 0x94, 0xb3, 0x29, 0x55, 0xe0, 0x9d, 0xfe,
	0x30, 0x0b, 0x9d, 0x1f, 0x36, 0xa9, 0xf7, 0x95, 0xca, 0x5e, 0x89, 0x2c, 0x3a, 0xd0, 0x09, 0x07,
	0x98, 0x1c, 0xea, 0xc4, 0xce, 0x8c, 0x5c, 0x8b, 0x70, 0x3e, 0xc7, 0x34, 0x34, 0x47, 0x68, 0x58,
	0x6d, 0xab, 0xbb, 0xbd, 0xff, 0xd4, 0xfd, 0xf3, 0xb9, 0xdc, 0x43, 0x3d, 0x74, 0x89, 0xd1, 0xaf,
	0x9d, 0xf8, 0x1b, 0xdf, 0x2c, 0xfb, 0x86, 0x15, 0xec, 0x18, 0x73, 0xb9, 0x53, 0x9f, 0x3c, 0xa4,
	0xd3, 0x29, 0x2b, 0x0e, 0x44, 0x79, 0x28, 0xe0, 0x53, 0x0e, 0x52, 0x85, 0x33, 0xa0, 0x53, 0x10,
	0x32, 0x54, 0x18, 0x72, 0x4c, 0x1a, 0x76, 0xbb, 0xd2, 0xdd, 0x0a, 0xee, 0x9f, 0x81, 0x81, 0xe1,
	0xfa, 0x06, 0x1b, 0xe3, 0x00, 0x13, 0xe7, 0x35, 0xe9, 0x9c, 0x33, 0xc9, 0x0c, 0x53, 0x09, 0xbf,
	0xab, 0x2a, 0x5a, 0xd5, 0x5a, 0x57, 0x19, 0xf0, 0x9c, 0x6b, 0x40, 0xf6, 0x2e, 0x73, 0x29, 0x41,
	0x19, 0x5f, 0x93, 0x55, 0xb5, 0xec, 0xc1, 0x45, 0xd9, 0xb8, 0x04, 0xb5, 0xad, 0xf3, 0xd5, 0x22,
	0x77, 0xc6, 0xd1, 0xbf, 0xbd, 0xe8, 0xce, 0x4f, 0x9b, 0xec, 0x5e, 0x39, 0xe5, 0x74, 0x48, 0x8d,
	0x63, 0x12, 0xa6, 0x74, 0x0e, 0x3a, 0xc2, 0x96, 0xbf, 0x79, 0xe2, 0x57, 0x85, 0xdd, 0xb6, 0x82,
	0x4d, 0x8e, 0xc9, 0x1b, 0x3a, 0x07, 0x67, 0x48, 0x76, 0xd6, 0xfb, 0xd9, 0xb0, 0x75, 0xd4, 0xbd,
	0x32, 0x2a, 0xcd, 0xd8, 0x2a, 0x5d, 0xd1, 0x65, 0xb7, 0xd8, 0x66, 0x64, 0xd0, 0xb5, 0x50, 0xdb,
	0xc9, 0xd9, 0x67, 0xe7, 0x03, 0xb9, 0x3d, 0xc9, 0xe3, 0x18, 0x44, 0x18, 0xf3, 0x5c, 0xce, 0x42,
	0x96, 0x2a, 0x10, 0x0b, 0xca, 0x1b, 0x15, 0xad, 0xde, 0x75, 0x4d, 0xb7, 0xdd, 0xd3, 0x6e, 0xbb,
	0x2f, 0xca, 0xee, 0x6b, 0xe1, 0x77, 0xcb, 0x7e, 0xfc, 0x5f, 0x70, 0xcb, 0x18, 0x5e, 0x16, 0x82,
	0xa3, 0x72, 0xde, 0xe9, 0x93, 0x9b, 0xa5, 0x58, 0xb2, 0x63, 0x08, 0x27, 0x4b, 0x05, 0xb2, 0x51,
	0xd5, 0xd2, 0x7b, 0x17, 0xa4, 0xef, 0x8e, 0x52, 0xd5, 0xdb, 0x7f, 0x4f, 0x79, 0x0e, 0xc1, 0x75,
	0x33, 0x36, 0x62, 0xc7, 0xe0, 0x17, 0x43, 0xce, 0x73, 0x72, 0x37, 0x66, 0x5c, 0x15, 0x26, 0x45,
	0x15, 0x84, 0x38, 0xf9, 0x08, 0x91, 0x5a, 0x35, 0x60, 0x43, 0x37, 0xa0, 0x6e, 0x90, 0x51, 0x41,
	0xbc, 0x35, 0x80, 0x7e, 0x79, 0xff, 0x19, 0x79, 0xc2, 0xd0, 0x5c, 0x50, 0x26, 0xf0, 0xcb, 0xf2,
	0x2f, 0xcf, 0xea, 0xd7, 0x0e, 0xb8, 0x1c, 0x16, 0xb9, 0x86, 0xd6, 0xe4, 0x7f, 0x1d, 0xb0, 0xf7,
	0x2b, 0x00, 0x00, 0xff, 0xff, 0x29, 0x92, 0x44, 0x19, 0x76, 0x04, 0x00, 0x00,
}
