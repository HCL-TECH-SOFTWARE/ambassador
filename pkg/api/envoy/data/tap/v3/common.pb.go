// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/data/tap/v3/common.proto

package envoy_data_tap_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type Body struct {
	Truncated bool `protobuf:"varint,3,opt,name=truncated,proto3" json:"truncated,omitempty"`
	// Types that are valid to be assigned to BodyType:
	//	*Body_AsBytes
	//	*Body_AsString
	BodyType             isBody_BodyType `protobuf_oneof:"body_type"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Body) Reset()         { *m = Body{} }
func (m *Body) String() string { return proto.CompactTextString(m) }
func (*Body) ProtoMessage()    {}
func (*Body) Descriptor() ([]byte, []int) {
	return fileDescriptor_b542295a11de64c9, []int{0}
}

func (m *Body) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Body.Unmarshal(m, b)
}
func (m *Body) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Body.Marshal(b, m, deterministic)
}
func (m *Body) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Body.Merge(m, src)
}
func (m *Body) XXX_Size() int {
	return xxx_messageInfo_Body.Size(m)
}
func (m *Body) XXX_DiscardUnknown() {
	xxx_messageInfo_Body.DiscardUnknown(m)
}

var xxx_messageInfo_Body proto.InternalMessageInfo

func (m *Body) GetTruncated() bool {
	if m != nil {
		return m.Truncated
	}
	return false
}

type isBody_BodyType interface {
	isBody_BodyType()
}

type Body_AsBytes struct {
	AsBytes []byte `protobuf:"bytes,1,opt,name=as_bytes,json=asBytes,proto3,oneof"`
}

type Body_AsString struct {
	AsString string `protobuf:"bytes,2,opt,name=as_string,json=asString,proto3,oneof"`
}

func (*Body_AsBytes) isBody_BodyType() {}

func (*Body_AsString) isBody_BodyType() {}

func (m *Body) GetBodyType() isBody_BodyType {
	if m != nil {
		return m.BodyType
	}
	return nil
}

func (m *Body) GetAsBytes() []byte {
	if x, ok := m.GetBodyType().(*Body_AsBytes); ok {
		return x.AsBytes
	}
	return nil
}

func (m *Body) GetAsString() string {
	if x, ok := m.GetBodyType().(*Body_AsString); ok {
		return x.AsString
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Body) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Body_AsBytes)(nil),
		(*Body_AsString)(nil),
	}
}

func init() {
	proto.RegisterType((*Body)(nil), "envoy.data.tap.v3.Body")
}

func init() { proto.RegisterFile("envoy/data/tap/v3/common.proto", fileDescriptor_b542295a11de64c9) }

var fileDescriptor_b542295a11de64c9 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0xeb, 0x80, 0x4a, 0xe2, 0x32, 0x40, 0xa6, 0xa8, 0x25, 0x25, 0x74, 0xca, 0x64, 0x4b,
	0x64, 0x41, 0x8c, 0x61, 0xe9, 0x58, 0x85, 0x0f, 0x88, 0x5e, 0x9a, 0xa8, 0x44, 0xa2, 0x7e, 0x56,
	0xfc, 0x12, 0xe1, 0x8d, 0x91, 0x95, 0x95, 0x4f, 0x61, 0x47, 0x62, 0xe5, 0x8f, 0x50, 0xdc, 0x01,
	0x89, 0xae, 0xe7, 0xf8, 0xca, 0xef, 0x5e, 0xbe, 0x6c, 0xd4, 0x80, 0x56, 0xd6, 0x40, 0x20, 0x09,
	0xb4, 0x1c, 0x32, 0xb9, 0xc5, 0xfd, 0x1e, 0x95, 0xd0, 0x1d, 0x12, 0x86, 0x97, 0xce, 0x8b, 0xd1,
	0x0b, 0x02, 0x2d, 0x86, 0x6c, 0x1e, 0xf7, 0xb5, 0x06, 0x09, 0x4a, 0x21, 0x01, 0xb5, 0xa8, 0x8c,
	0x34, 0x04, 0xd4, 0x9b, 0x43, 0x62, 0x7e, 0x73, 0xa4, 0x87, 0xa6, 0x33, 0x2d, 0xaa, 0x56, 0xed,
	0x0e, 0x4f, 0x56, 0xef, 0x8c, 0x9f, 0xe6, 0x58, 0xdb, 0xf0, 0x8a, 0x07, 0xd4, 0xf5, 0x6a, 0x0b,
	0xd4, 0xd4, 0xd1, 0x49, 0xc2, 0x52, 0xbf, 0xf8, 0x03, 0xe1, 0x82, 0xfb, 0x60, 0xca, 0xca, 0x52,
	0x63, 0x22, 0x96, 0xb0, 0xf4, 0x7c, 0x3d, 0x29, 0xce, 0xc0, 0xe4, 0x23, 0x08, 0x63, 0x1e, 0x80,
	0x29, 0x0d, 0x75, 0xad, 0xda, 0x45, 0x5e, 0xc2, 0xd2, 0x60, 0x3d, 0x29, 0x7c, 0x30, 0x8f, 0x8e,
	0xdc, 0xaf, 0x3e, 0xbe, 0xde, 0x96, 0x31, 0x5f, 0xfc, 0x3f, 0xff, 0x16, 0x9e, 0xf5, 0x13, 0x88,
	0xf1, 0xf7, 0x7c, 0xc6, 0x83, 0x0a, 0x6b, 0x5b, 0x92, 0xd5, 0x4d, 0x7e, 0xf7, 0xf9, 0xfa, 0xfd,
	0x33, 0xf5, 0x2e, 0x3c, 0x7e, 0xdd, 0xa2, 0x70, 0x31, 0xdd, 0xe1, 0x8b, 0x15, 0x47, 0x03, 0xe4,
	0xb3, 0x07, 0xb7, 0xd0, 0x66, 0xec, 0xb2, 0x61, 0xd5, 0xd4, 0x95, 0xca, 0x7e, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x62, 0x3f, 0xc9, 0xa7, 0x4b, 0x01, 0x00, 0x00,
}
