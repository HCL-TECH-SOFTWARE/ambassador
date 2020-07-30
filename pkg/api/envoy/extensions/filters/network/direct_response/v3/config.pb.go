// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/network/direct_response/v3/config.proto

package envoy_extensions_filters_network_direct_response_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/config/core/v3"
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

type Config struct {
	Response             *v3.DataSource `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec2699568213c18, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetResponse() *v3.DataSource {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "envoy.extensions.filters.network.direct_response.v3.Config")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/direct_response/v3/config.proto", fileDescriptor_5ec2699568213c18)
}

var fileDescriptor_5ec2699568213c18 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0xd0, 0xb1, 0x4a, 0x43, 0x31,
	0x14, 0x06, 0x60, 0x6e, 0x87, 0x8b, 0xa4, 0x8b, 0x74, 0x92, 0x82, 0x5a, 0x9d, 0x9c, 0x4e, 0xa0,
	0x41, 0x04, 0xe9, 0xa0, 0xd5, 0x07, 0x28, 0xba, 0xb9, 0x48, 0x7a, 0x7b, 0x2a, 0x41, 0xc9, 0xb9,
	0x24, 0xe7, 0xc6, 0xde, 0xcd, 0xc1, 0xc1, 0x67, 0xf0, 0x51, 0xdc, 0x05, 0x57, 0xdf, 0x48, 0x6e,
	0x12, 0x15, 0x15, 0x97, 0xce, 0xe1, 0x7c, 0xff, 0xff, 0x47, 0x9c, 0xa0, 0x0d, 0xd4, 0x4a, 0x5c,
	0x31, 0x5a, 0x6f, 0xc8, 0x7a, 0xb9, 0x34, 0x77, 0x8c, 0xce, 0x4b, 0x8b, 0x7c, 0x4f, 0xee, 0x56,
	0x2e, 0x8c, 0xc3, 0x8a, 0xaf, 0x1d, 0xfa, 0x9a, 0xac, 0x47, 0x19, 0x94, 0xac, 0xc8, 0x2e, 0xcd,
	0x0d, 0xd4, 0x8e, 0x98, 0x06, 0x2a, 0x0a, 0xf0, 0x2d, 0x40, 0x16, 0x20, 0x0b, 0xf0, 0x4b, 0x80,
	0xa0, 0x86, 0xbb, 0x29, 0x36, 0x41, 0xb2, 0x22, 0x17, 0xdd, 0xb9, 0xf6, 0x98, 0xd4, 0xe1, 0x76,
	0xb3, 0xa8, 0xb5, 0xd4, 0xd6, 0x12, 0x6b, 0x8e, 0xbd, 0x3c, 0x6b, 0x6e, 0x7c, 0x7e, 0xde, 0xfb,
	0xf3, 0x1c, 0xd0, 0x75, 0xe9, 0xc6, 0xe6, 0x5e, 0xfb, 0x8f, 0x85, 0x28, 0xcf, 0xa2, 0x3f, 0x98,
	0x88, 0x8d, 0xcf, 0xf0, 0xad, 0x62, 0x54, 0x1c, 0xf4, 0xc7, 0x23, 0x48, 0xad, 0xf3, 0x92, 0xae,
	0x00, 0x04, 0x05, 0xe7, 0x9a, 0xf5, 0x25, 0x35, 0xae, 0xc2, 0x8b, 0xaf, 0x8b, 0xe3, 0xc9, 0xf3,
	0xeb, 0xd3, 0xce, 0x91, 0x38, 0xfc, 0x71, 0x91, 0x36, 0xfe, 0x3f, 0x71, 0x0c, 0x29, 0x7b, 0x7a,
	0xf5, 0xf2, 0xf0, 0xf6, 0x5e, 0xf6, 0x36, 0x7b, 0xe2, 0xd4, 0x50, 0x4a, 0xad, 0x1d, 0xad, 0x5a,
	0x58, 0xe3, 0xdb, 0xa6, 0xfd, 0x84, 0xce, 0xba, 0x81, 0xb3, 0x62, 0x5e, 0xc6, 0xa5, 0xea, 0x23,
	0x00, 0x00, 0xff, 0xff, 0x97, 0x12, 0xd5, 0x47, 0xc5, 0x01, 0x00, 0x00,
}
