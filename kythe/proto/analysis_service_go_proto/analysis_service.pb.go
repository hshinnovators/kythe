// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kythe/proto/analysis_service.proto

package analysis_service_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "kythe.io/kythe/proto/analysis_go_proto"
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

func init() {
	proto.RegisterFile("kythe/proto/analysis_service.proto", fileDescriptor_465e5f838439a146)
}

var fileDescriptor_465e5f838439a146 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xca, 0xae, 0x2c, 0xc9,
	0x48, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0xcc, 0x4b, 0xcc, 0xa9, 0x2c, 0xce, 0x2c,
	0x8e, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x03, 0x0b, 0x0b, 0x71, 0x83, 0xd5, 0x40,
	0x38, 0x52, 0x52, 0xd8, 0x34, 0x40, 0xe4, 0x8c, 0xe2, 0xb9, 0x84, 0x9d, 0xf3, 0x73, 0x0b, 0x32,
	0x73, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x1c, 0x41, 0x92, 0x55, 0xa9, 0x45, 0x42, 0x1e, 0x5c, 0xec,
	0x50, 0xb6, 0x90, 0x8c, 0x1e, 0x92, 0x59, 0x7a, 0x8e, 0x50, 0xed, 0x41, 0xa9, 0x85, 0xa5, 0xa9,
	0xc5, 0x25, 0x52, 0xd2, 0x58, 0x65, 0xfd, 0x4b, 0x4b, 0x0a, 0x4a, 0x4b, 0x94, 0x18, 0x0c, 0x18,
	0x8d, 0xfc, 0xb8, 0xf8, 0xdd, 0x32, 0x73, 0x52, 0x5d, 0x12, 0x4b, 0x12, 0x83, 0x21, 0x4e, 0x14,
	0xb2, 0xe6, 0x62, 0x76, 0x4f, 0x2d, 0x11, 0x92, 0x44, 0xd1, 0x0a, 0x52, 0x04, 0x37, 0x55, 0x14,
	0x43, 0x0a, 0xa4, 0x1f, 0x64, 0x9e, 0x93, 0x0d, 0x97, 0x7c, 0x72, 0x7e, 0xae, 0x5e, 0x7a, 0x7e,
	0x7e, 0x7a, 0x4e, 0xaa, 0x5e, 0x4a, 0x6a, 0x59, 0x49, 0x7e, 0x7e, 0x4e, 0x31, 0xb2, 0xfa, 0x28,
	0x49, 0xf4, 0x40, 0x89, 0x4f, 0xcf, 0x8f, 0x07, 0x4b, 0x25, 0xb1, 0x81, 0x29, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x89, 0x92, 0xd1, 0xd4, 0x44, 0x01, 0x00, 0x00,
}
