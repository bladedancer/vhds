// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/v3alpha/http.proto

package envoy_type_v3alpha

import (
	fmt "fmt"
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

type CodecClientType int32

const (
	CodecClientType_HTTP1 CodecClientType = 0
	CodecClientType_HTTP2 CodecClientType = 1
	CodecClientType_HTTP3 CodecClientType = 2
)

var CodecClientType_name = map[int32]string{
	0: "HTTP1",
	1: "HTTP2",
	2: "HTTP3",
}

var CodecClientType_value = map[string]int32{
	"HTTP1": 0,
	"HTTP2": 1,
	"HTTP3": 2,
}

func (x CodecClientType) String() string {
	return proto.EnumName(CodecClientType_name, int32(x))
}

func (CodecClientType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_387cc5085da723e7, []int{0}
}

func init() {
	proto.RegisterEnum("envoy.type.v3alpha.CodecClientType", CodecClientType_name, CodecClientType_value)
}

func init() { proto.RegisterFile("envoy/type/v3alpha/http.proto", fileDescriptor_387cc5085da723e7) }

var fileDescriptor_387cc5085da723e7 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2f, 0x33, 0x4e, 0xcc, 0x29, 0xc8, 0x48, 0xd4, 0xcf,
	0x28, 0x29, 0x29, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x02, 0x4b, 0xeb, 0x81, 0xa4,
	0xf5, 0xa0, 0xd2, 0x5a, 0x46, 0x5c, 0xfc, 0xce, 0xf9, 0x29, 0xa9, 0xc9, 0xce, 0x39, 0x99, 0xa9,
	0x79, 0x25, 0x21, 0x95, 0x05, 0xa9, 0x42, 0x9c, 0x5c, 0xac, 0x1e, 0x21, 0x21, 0x01, 0x86, 0x02,
	0x0c, 0x30, 0xa6, 0x91, 0x00, 0x23, 0x8c, 0x69, 0x2c, 0xc0, 0xe4, 0xa4, 0xcf, 0xa5, 0x90, 0x99,
	0xaf, 0x07, 0x36, 0xac, 0xa0, 0x28, 0xbf, 0xa2, 0x52, 0x0f, 0xd3, 0x5c, 0x27, 0x4e, 0x8f, 0x92,
	0x92, 0x82, 0x00, 0x90, 0xb5, 0x01, 0x8c, 0x49, 0x6c, 0x60, 0xfb, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xaa, 0xdb, 0x88, 0x15, 0xa0, 0x00, 0x00, 0x00,
}
