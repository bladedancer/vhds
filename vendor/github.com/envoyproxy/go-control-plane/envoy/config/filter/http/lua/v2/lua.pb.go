// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/lua/v2/lua.proto

package envoy_config_filter_http_lua_v2

import (
	fmt "fmt"
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

type Lua struct {
	InlineCode           string   `protobuf:"bytes,1,opt,name=inline_code,json=inlineCode,proto3" json:"inline_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Lua) Reset()         { *m = Lua{} }
func (m *Lua) String() string { return proto.CompactTextString(m) }
func (*Lua) ProtoMessage()    {}
func (*Lua) Descriptor() ([]byte, []int) {
	return fileDescriptor_f59dca3e63e33613, []int{0}
}

func (m *Lua) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Lua.Unmarshal(m, b)
}
func (m *Lua) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Lua.Marshal(b, m, deterministic)
}
func (m *Lua) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lua.Merge(m, src)
}
func (m *Lua) XXX_Size() int {
	return xxx_messageInfo_Lua.Size(m)
}
func (m *Lua) XXX_DiscardUnknown() {
	xxx_messageInfo_Lua.DiscardUnknown(m)
}

var xxx_messageInfo_Lua proto.InternalMessageInfo

func (m *Lua) GetInlineCode() string {
	if m != nil {
		return m.InlineCode
	}
	return ""
}

func init() {
	proto.RegisterType((*Lua)(nil), "envoy.config.filter.http.lua.v2.Lua")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/lua/v2/lua.proto", fileDescriptor_f59dca3e63e33613)
}

var fileDescriptor_f59dca3e63e33613 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x28, 0x29, 0x29, 0xd0, 0xcf, 0x29, 0x4d, 0xd4, 0x2f, 0x33, 0x02, 0x51, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0xf2, 0x60, 0xa5, 0x7a, 0x10, 0xa5, 0x7a, 0x10, 0xa5, 0x7a, 0x20, 0xa5,
	0x7a, 0x20, 0x35, 0x65, 0x46, 0x52, 0xe2, 0x65, 0x89, 0x39, 0x99, 0x29, 0x89, 0x25, 0xa9, 0xfa,
	0x30, 0x06, 0x44, 0xa7, 0x92, 0x3e, 0x17, 0xb3, 0x4f, 0x69, 0xa2, 0x90, 0x06, 0x17, 0x77, 0x66,
	0x5e, 0x4e, 0x66, 0x5e, 0x6a, 0x7c, 0x72, 0x7e, 0x4a, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xa7,
	0x13, 0xfb, 0x2f, 0x27, 0x96, 0x22, 0x26, 0x05, 0xc6, 0x20, 0x2e, 0x88, 0x9c, 0x73, 0x7e, 0x4a,
	0xaa, 0x93, 0x35, 0x97, 0x6e, 0x66, 0xbe, 0x1e, 0xd8, 0xbe, 0x82, 0xa2, 0xfc, 0x8a, 0x4a, 0x3d,
	0x02, 0x56, 0x3b, 0x71, 0xf8, 0x94, 0x26, 0x06, 0x80, 0xec, 0x0a, 0x60, 0x4c, 0x62, 0x03, 0x5b,
	0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x29, 0x57, 0x20, 0xdb, 0x00, 0x00, 0x00,
}
