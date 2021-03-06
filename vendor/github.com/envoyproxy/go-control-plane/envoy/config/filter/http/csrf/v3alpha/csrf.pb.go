// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/csrf/v3alpha/csrf.proto

package envoy_config_filter_http_csrf_v3alpha

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type CsrfPolicy struct {
	FilterEnabled        *core.RuntimeFractionalPercent `protobuf:"bytes,1,opt,name=filter_enabled,json=filterEnabled,proto3" json:"filter_enabled,omitempty"`
	ShadowEnabled        *core.RuntimeFractionalPercent `protobuf:"bytes,2,opt,name=shadow_enabled,json=shadowEnabled,proto3" json:"shadow_enabled,omitempty"`
	AdditionalOrigins    []*v3alpha.StringMatcher       `protobuf:"bytes,3,rep,name=additional_origins,json=additionalOrigins,proto3" json:"additional_origins,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *CsrfPolicy) Reset()         { *m = CsrfPolicy{} }
func (m *CsrfPolicy) String() string { return proto.CompactTextString(m) }
func (*CsrfPolicy) ProtoMessage()    {}
func (*CsrfPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0ea04258cfdb2a5, []int{0}
}

func (m *CsrfPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CsrfPolicy.Unmarshal(m, b)
}
func (m *CsrfPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CsrfPolicy.Marshal(b, m, deterministic)
}
func (m *CsrfPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CsrfPolicy.Merge(m, src)
}
func (m *CsrfPolicy) XXX_Size() int {
	return xxx_messageInfo_CsrfPolicy.Size(m)
}
func (m *CsrfPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_CsrfPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_CsrfPolicy proto.InternalMessageInfo

func (m *CsrfPolicy) GetFilterEnabled() *core.RuntimeFractionalPercent {
	if m != nil {
		return m.FilterEnabled
	}
	return nil
}

func (m *CsrfPolicy) GetShadowEnabled() *core.RuntimeFractionalPercent {
	if m != nil {
		return m.ShadowEnabled
	}
	return nil
}

func (m *CsrfPolicy) GetAdditionalOrigins() []*v3alpha.StringMatcher {
	if m != nil {
		return m.AdditionalOrigins
	}
	return nil
}

func init() {
	proto.RegisterType((*CsrfPolicy)(nil), "envoy.config.filter.http.csrf.v3alpha.CsrfPolicy")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/csrf/v3alpha/csrf.proto", fileDescriptor_a0ea04258cfdb2a5)
}

var fileDescriptor_a0ea04258cfdb2a5 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xc1, 0x4e, 0xf2, 0x40,
	0x14, 0x85, 0x53, 0xc8, 0xff, 0x47, 0x87, 0x40, 0xb4, 0x1b, 0x09, 0x2b, 0x34, 0x31, 0xe2, 0x66,
	0x86, 0xc0, 0x1b, 0xd4, 0xe8, 0xce, 0x48, 0xea, 0x42, 0x77, 0xe4, 0xd2, 0xde, 0xd2, 0x9b, 0x94,
	0x99, 0xc9, 0x74, 0x44, 0xfb, 0x0a, 0x3e, 0x86, 0x8f, 0xe9, 0xca, 0xb4, 0xb7, 0x94, 0xad, 0x71,
	0xd7, 0xce, 0x9c, 0xf3, 0x9d, 0x93, 0x33, 0x62, 0x8e, 0x7a, 0x6f, 0x2a, 0x95, 0x18, 0x9d, 0xd1,
	0x56, 0x65, 0x54, 0x78, 0x74, 0x2a, 0xf7, 0xde, 0xaa, 0xa4, 0x74, 0x99, 0xda, 0x2f, 0xa1, 0xb0,
	0x39, 0x34, 0x3f, 0xd2, 0x3a, 0xe3, 0x4d, 0x78, 0xdd, 0x38, 0x24, 0x3b, 0x24, 0x3b, 0x64, 0xed,
	0x90, 0x8d, 0xa8, 0x75, 0x4c, 0x2e, 0x19, 0x0c, 0x96, 0x8e, 0x10, 0xe3, 0x50, 0x6d, 0xa0, 0x44,
	0x26, 0x4d, 0x6e, 0x58, 0xe2, 0x2b, 0x8b, 0x6a, 0x07, 0x3e, 0xc9, 0xd1, 0x75, 0xda, 0xd2, 0x3b,
	0xd2, 0xdb, 0x56, 0x78, 0xb1, 0x87, 0x82, 0x52, 0xf0, 0xa8, 0x0e, 0x1f, 0x7c, 0x71, 0xf5, 0xd5,
	0x13, 0xe2, 0xae, 0x74, 0xd9, 0xca, 0x14, 0x94, 0x54, 0x21, 0x88, 0x11, 0xf7, 0x59, 0xa3, 0x86,
	0x4d, 0x81, 0xe9, 0x38, 0x98, 0x06, 0xb3, 0xc1, 0x62, 0x2e, 0xb9, 0x33, 0x58, 0x3a, 0xf4, 0x93,
	0x75, 0x19, 0x19, 0xbf, 0x69, 0x4f, 0x3b, 0x7c, 0x70, 0x90, 0x78, 0x32, 0x1a, 0x8a, 0x15, 0xba,
	0x04, 0xb5, 0x8f, 0x4e, 0xbe, 0xa3, 0x7f, 0x9f, 0x41, 0xef, 0x2c, 0x88, 0x87, 0x4c, 0xbc, 0x67,
	0x60, 0xf8, 0x22, 0x46, 0x65, 0x0e, 0xa9, 0x79, 0xef, 0x22, 0x7a, 0x7f, 0x8b, 0x88, 0x87, 0xcc,
	0x39, 0x80, 0x5f, 0x45, 0x08, 0x69, 0x4a, 0xac, 0x59, 0x1b, 0x47, 0x5b, 0xd2, 0xe5, 0xb8, 0x3f,
	0xed, 0xcf, 0x06, 0x8b, 0xdb, 0x16, 0x5e, 0x2f, 0x25, 0xdb, 0xa5, 0xba, 0x94, 0xe7, 0x66, 0xa9,
	0x47, 0x3e, 0x8d, 0xcf, 0x8f, 0x90, 0x27, 0x66, 0x44, 0x91, 0x58, 0x92, 0x61, 0x82, 0x75, 0xe6,
	0xa3, 0x92, 0xbf, 0x7a, 0xc0, 0xe8, 0xb4, 0x19, 0xb6, 0x9e, 0x79, 0x15, 0x6c, 0xfe, 0x37, 0x7b,
	0x2f, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xf1, 0x8d, 0x0a, 0x2f, 0x02, 0x00, 0x00,
}
