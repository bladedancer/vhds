// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/fault/v3alpha/fault.proto

package envoy_config_filter_http_fault_v3alpha

import (
	fmt "fmt"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/route"
	v3alpha1 "github.com/envoyproxy/go-control-plane/envoy/config/filter/fault/v3alpha"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/type/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type FaultAbort struct {
	// Types that are valid to be assigned to ErrorType:
	//	*FaultAbort_HttpStatus
	ErrorType            isFaultAbort_ErrorType     `protobuf_oneof:"error_type"`
	Percentage           *v3alpha.FractionalPercent `protobuf:"bytes,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *FaultAbort) Reset()         { *m = FaultAbort{} }
func (m *FaultAbort) String() string { return proto.CompactTextString(m) }
func (*FaultAbort) ProtoMessage()    {}
func (*FaultAbort) Descriptor() ([]byte, []int) {
	return fileDescriptor_063e20c61b5a3683, []int{0}
}

func (m *FaultAbort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultAbort.Unmarshal(m, b)
}
func (m *FaultAbort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultAbort.Marshal(b, m, deterministic)
}
func (m *FaultAbort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultAbort.Merge(m, src)
}
func (m *FaultAbort) XXX_Size() int {
	return xxx_messageInfo_FaultAbort.Size(m)
}
func (m *FaultAbort) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultAbort.DiscardUnknown(m)
}

var xxx_messageInfo_FaultAbort proto.InternalMessageInfo

type isFaultAbort_ErrorType interface {
	isFaultAbort_ErrorType()
}

type FaultAbort_HttpStatus struct {
	HttpStatus uint32 `protobuf:"varint,2,opt,name=http_status,json=httpStatus,proto3,oneof"`
}

func (*FaultAbort_HttpStatus) isFaultAbort_ErrorType() {}

func (m *FaultAbort) GetErrorType() isFaultAbort_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return nil
}

func (m *FaultAbort) GetHttpStatus() uint32 {
	if x, ok := m.GetErrorType().(*FaultAbort_HttpStatus); ok {
		return x.HttpStatus
	}
	return 0
}

func (m *FaultAbort) GetPercentage() *v3alpha.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FaultAbort) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FaultAbort_HttpStatus)(nil),
	}
}

type HTTPFault struct {
	Delay                           *v3alpha1.FaultDelay     `protobuf:"bytes,1,opt,name=delay,proto3" json:"delay,omitempty"`
	Abort                           *FaultAbort              `protobuf:"bytes,2,opt,name=abort,proto3" json:"abort,omitempty"`
	UpstreamCluster                 string                   `protobuf:"bytes,3,opt,name=upstream_cluster,json=upstreamCluster,proto3" json:"upstream_cluster,omitempty"`
	Headers                         []*route.HeaderMatcher   `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	DownstreamNodes                 []string                 `protobuf:"bytes,5,rep,name=downstream_nodes,json=downstreamNodes,proto3" json:"downstream_nodes,omitempty"`
	MaxActiveFaults                 *wrappers.UInt32Value    `protobuf:"bytes,6,opt,name=max_active_faults,json=maxActiveFaults,proto3" json:"max_active_faults,omitempty"`
	ResponseRateLimit               *v3alpha1.FaultRateLimit `protobuf:"bytes,7,opt,name=response_rate_limit,json=responseRateLimit,proto3" json:"response_rate_limit,omitempty"`
	DelayPercentRuntime             string                   `protobuf:"bytes,8,opt,name=delay_percent_runtime,json=delayPercentRuntime,proto3" json:"delay_percent_runtime,omitempty"`
	AbortPercentRuntime             string                   `protobuf:"bytes,9,opt,name=abort_percent_runtime,json=abortPercentRuntime,proto3" json:"abort_percent_runtime,omitempty"`
	DelayDurationRuntime            string                   `protobuf:"bytes,10,opt,name=delay_duration_runtime,json=delayDurationRuntime,proto3" json:"delay_duration_runtime,omitempty"`
	AbortHttpStatusRuntime          string                   `protobuf:"bytes,11,opt,name=abort_http_status_runtime,json=abortHttpStatusRuntime,proto3" json:"abort_http_status_runtime,omitempty"`
	MaxActiveFaultsRuntime          string                   `protobuf:"bytes,12,opt,name=max_active_faults_runtime,json=maxActiveFaultsRuntime,proto3" json:"max_active_faults_runtime,omitempty"`
	ResponseRateLimitPercentRuntime string                   `protobuf:"bytes,13,opt,name=response_rate_limit_percent_runtime,json=responseRateLimitPercentRuntime,proto3" json:"response_rate_limit_percent_runtime,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                 `json:"-"`
	XXX_unrecognized                []byte                   `json:"-"`
	XXX_sizecache                   int32                    `json:"-"`
}

func (m *HTTPFault) Reset()         { *m = HTTPFault{} }
func (m *HTTPFault) String() string { return proto.CompactTextString(m) }
func (*HTTPFault) ProtoMessage()    {}
func (*HTTPFault) Descriptor() ([]byte, []int) {
	return fileDescriptor_063e20c61b5a3683, []int{1}
}

func (m *HTTPFault) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPFault.Unmarshal(m, b)
}
func (m *HTTPFault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPFault.Marshal(b, m, deterministic)
}
func (m *HTTPFault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPFault.Merge(m, src)
}
func (m *HTTPFault) XXX_Size() int {
	return xxx_messageInfo_HTTPFault.Size(m)
}
func (m *HTTPFault) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPFault.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPFault proto.InternalMessageInfo

func (m *HTTPFault) GetDelay() *v3alpha1.FaultDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func (m *HTTPFault) GetAbort() *FaultAbort {
	if m != nil {
		return m.Abort
	}
	return nil
}

func (m *HTTPFault) GetUpstreamCluster() string {
	if m != nil {
		return m.UpstreamCluster
	}
	return ""
}

func (m *HTTPFault) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPFault) GetDownstreamNodes() []string {
	if m != nil {
		return m.DownstreamNodes
	}
	return nil
}

func (m *HTTPFault) GetMaxActiveFaults() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxActiveFaults
	}
	return nil
}

func (m *HTTPFault) GetResponseRateLimit() *v3alpha1.FaultRateLimit {
	if m != nil {
		return m.ResponseRateLimit
	}
	return nil
}

func (m *HTTPFault) GetDelayPercentRuntime() string {
	if m != nil {
		return m.DelayPercentRuntime
	}
	return ""
}

func (m *HTTPFault) GetAbortPercentRuntime() string {
	if m != nil {
		return m.AbortPercentRuntime
	}
	return ""
}

func (m *HTTPFault) GetDelayDurationRuntime() string {
	if m != nil {
		return m.DelayDurationRuntime
	}
	return ""
}

func (m *HTTPFault) GetAbortHttpStatusRuntime() string {
	if m != nil {
		return m.AbortHttpStatusRuntime
	}
	return ""
}

func (m *HTTPFault) GetMaxActiveFaultsRuntime() string {
	if m != nil {
		return m.MaxActiveFaultsRuntime
	}
	return ""
}

func (m *HTTPFault) GetResponseRateLimitPercentRuntime() string {
	if m != nil {
		return m.ResponseRateLimitPercentRuntime
	}
	return ""
}

func init() {
	proto.RegisterType((*FaultAbort)(nil), "envoy.config.filter.http.fault.v3alpha.FaultAbort")
	proto.RegisterType((*HTTPFault)(nil), "envoy.config.filter.http.fault.v3alpha.HTTPFault")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/fault/v3alpha/fault.proto", fileDescriptor_063e20c61b5a3683)
}

var fileDescriptor_063e20c61b5a3683 = []byte{
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcf, 0x4e, 0xdb, 0x4e,
	0x10, 0xc7, 0x7f, 0x26, 0x10, 0x60, 0xf3, 0x43, 0x80, 0x69, 0xa9, 0x8b, 0xaa, 0x36, 0x02, 0x15,
	0xa5, 0x95, 0x58, 0xab, 0x86, 0x4b, 0x6f, 0x25, 0x50, 0x94, 0x56, 0xb4, 0x8a, 0x5c, 0xda, 0xab,
	0x35, 0x24, 0x93, 0xc4, 0x92, 0xe3, 0x5d, 0xad, 0xd7, 0x81, 0xbc, 0x4b, 0x5f, 0xa6, 0xb7, 0x3e,
	0x42, 0x1f, 0xa5, 0xea, 0xa9, 0xda, 0x59, 0xdb, 0xfc, 0x09, 0x95, 0x72, 0x89, 0xe2, 0x9d, 0xf9,
	0xcc, 0x77, 0xe7, 0x3b, 0x63, 0xb3, 0x00, 0xd3, 0x89, 0x98, 0xfa, 0x3d, 0x91, 0x0e, 0xe2, 0xa1,
	0x3f, 0x88, 0x13, 0x8d, 0xca, 0x1f, 0x69, 0x2d, 0xfd, 0x01, 0xe4, 0x89, 0xf6, 0x27, 0x87, 0x90,
	0xc8, 0x11, 0xd8, 0x27, 0x2e, 0x95, 0xd0, 0xc2, 0xdd, 0x27, 0x86, 0x5b, 0x86, 0x5b, 0x86, 0x1b,
	0x86, 0xdb, 0xac, 0x82, 0xd9, 0xd9, 0xb3, 0xb5, 0x41, 0xc6, 0x55, 0x19, 0x25, 0x72, 0x8d, 0xf6,
	0xd7, 0x16, 0xdb, 0x39, 0x78, 0xe8, 0x02, 0xff, 0xd4, 0xde, 0x69, 0xda, 0x74, 0x3d, 0x95, 0x58,
	0xc5, 0x25, 0xaa, 0x1e, 0xa6, 0x65, 0xc6, 0xf3, 0xa1, 0x10, 0xc3, 0x04, 0x7d, 0x7a, 0xba, 0xcc,
	0x07, 0xfe, 0x95, 0x02, 0x29, 0x51, 0x65, 0x45, 0xfc, 0xc9, 0x04, 0x92, 0xb8, 0x0f, 0x1a, 0xfd,
	0xf2, 0x8f, 0x0d, 0xec, 0x7e, 0x77, 0x18, 0x3b, 0x33, 0x52, 0xc7, 0x97, 0x42, 0x69, 0x97, 0xb3,
	0x86, 0xe9, 0x29, 0xca, 0x34, 0xe8, 0x3c, 0xf3, 0x16, 0x9a, 0x4e, 0x6b, 0xad, 0xdd, 0xf8, 0xd3,
	0x5e, 0x79, 0x5d, 0xdf, 0xf8, 0xb5, 0xd8, 0xfa, 0xe9, 0x74, 0xfe, 0x0b, 0x99, 0xc9, 0xf8, 0x42,
	0x09, 0xee, 0x7b, 0xc6, 0x8a, 0x8b, 0xc0, 0x10, 0xbd, 0x5a, 0xd3, 0x69, 0x35, 0x82, 0x97, 0xdc,
	0x5a, 0x65, 0xae, 0x5b, 0xda, 0xc2, 0xcf, 0x14, 0xf4, 0x74, 0x2c, 0x52, 0x48, 0xba, 0x36, 0x3f,
	0xbc, 0x05, 0xb6, 0x37, 0x19, 0x43, 0xa5, 0x84, 0x8a, 0x0c, 0xe3, 0xd6, 0x7e, 0xb7, 0x9d, 0x8f,
	0x8b, 0x2b, 0xce, 0xc6, 0xc2, 0xee, 0x8f, 0x3a, 0x5b, 0xed, 0x5c, 0x5c, 0x74, 0xe9, 0x8a, 0xee,
	0x09, 0x5b, 0xea, 0x63, 0x02, 0x53, 0xcf, 0x21, 0xa1, 0x03, 0xfe, 0xd0, 0x4c, 0xee, 0x8c, 0x83,
	0x13, 0x78, 0x6a, 0xa0, 0xd0, 0xb2, 0x6e, 0x87, 0x2d, 0x81, 0xe9, 0x95, 0x9a, 0x6b, 0x04, 0x01,
	0x9f, 0x6f, 0xb0, 0xfc, 0xc6, 0xa5, 0xd0, 0x16, 0x70, 0x5f, 0xb1, 0x8d, 0x5c, 0x66, 0x5a, 0x21,
	0x8c, 0xa3, 0x5e, 0x92, 0x67, 0x1a, 0x15, 0x59, 0xb0, 0x1a, 0xae, 0x97, 0xe7, 0x27, 0xf6, 0xd8,
	0x7d, 0xc7, 0x96, 0x47, 0x08, 0x7d, 0x54, 0x99, 0xb7, 0xd8, 0xac, 0xb5, 0x1a, 0xc1, 0x7e, 0x21,
	0x0b, 0x32, 0xae, 0x14, 0xec, 0x86, 0x74, 0x28, 0xef, 0x13, 0xe8, 0xde, 0x08, 0x55, 0x58, 0x62,
	0x46, 0xac, 0x2f, 0xae, 0xd2, 0x42, 0x2e, 0x15, 0x7d, 0xcc, 0xbc, 0xa5, 0x66, 0xcd, 0x88, 0xdd,
	0x9c, 0x7f, 0x36, 0xc7, 0x6e, 0x87, 0x6d, 0x8e, 0xe1, 0x3a, 0x32, 0x86, 0x4f, 0x30, 0xa2, 0x2e,
	0x32, 0xaf, 0x4e, 0xdd, 0x3e, 0xe3, 0x76, 0x51, 0x78, 0xb9, 0x28, 0xfc, 0xeb, 0x87, 0x54, 0x1f,
	0x06, 0xdf, 0x20, 0xc9, 0x31, 0x5c, 0x1f, 0xc3, 0xf5, 0x31, 0x51, 0xd4, 0x6c, 0xe6, 0x02, 0xdb,
	0x52, 0x98, 0x49, 0x91, 0x66, 0x18, 0x29, 0xd0, 0x18, 0x25, 0xf1, 0x38, 0xd6, 0xde, 0x32, 0xd5,
	0x7a, 0x33, 0xaf, 0xfd, 0x21, 0x68, 0x3c, 0x37, 0x60, 0xb8, 0x59, 0x56, 0xab, 0x8e, 0xdc, 0x80,
	0x3d, 0xa6, 0xb9, 0x44, 0xc5, 0x3a, 0x44, 0x2a, 0x4f, 0x75, 0x3c, 0x46, 0x6f, 0x85, 0x9c, 0xdc,
	0xa2, 0x60, 0xb9, 0x33, 0x36, 0x64, 0x18, 0x9a, 0xc0, 0x0c, 0xb3, 0x6a, 0x19, 0x0a, 0xde, 0x63,
	0x8e, 0xd8, 0xb6, 0xd5, 0xe9, 0xe7, 0x0a, 0xcc, 0x26, 0x56, 0x10, 0x23, 0xe8, 0x11, 0x45, 0x4f,
	0x8b, 0x60, 0x49, 0xbd, 0x65, 0x4f, 0xad, 0xd2, 0xad, 0xb7, 0xa2, 0x02, 0x1b, 0x04, 0x6e, 0x53,
	0x42, 0xa7, 0x7a, 0x27, 0x6e, 0xa1, 0x33, 0x53, 0xa8, 0xd0, 0xff, 0x2d, 0x7a, 0xcf, 0xef, 0x12,
	0x3d, 0x67, 0x7b, 0x0f, 0xd8, 0x3e, 0xd3, 0xed, 0x1a, 0x15, 0x79, 0x31, 0xe3, 0xe9, 0xdd, 0xce,
	0xdb, 0xa7, 0xec, 0x28, 0x16, 0x76, 0x56, 0x52, 0x89, 0xeb, 0xe9, 0x9c, 0x0b, 0xdf, 0xb6, 0xdf,
	0x85, 0xae, 0x59, 0x94, 0xae, 0x73, 0x59, 0xa7, 0x8d, 0x39, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0x3b, 0xd5, 0xa3, 0x4d, 0x3c, 0x05, 0x00, 0x00,
}