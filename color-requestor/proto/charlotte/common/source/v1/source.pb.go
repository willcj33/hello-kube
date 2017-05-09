// Code generated by protoc-gen-go.
// source: charlotte/common/source/v1/source.proto
// DO NOT EDIT!

/*
Package charlotte_common_source_v1 is a generated protocol buffer package.

It is generated from these files:
	charlotte/common/source/v1/source.proto

It has these top-level messages:
	Source
*/
package charlotte_common_source_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SourceType int32

const (
	SourceType_GOOGLE_MAIL     SourceType = 0
	SourceType_GOOGLE_CALENDAR SourceType = 1
	SourceType_LINKED_IN       SourceType = 2
	SourceType_FACEBOOK        SourceType = 3
	SourceType_INSTAGRAM       SourceType = 4
)

var SourceType_name = map[int32]string{
	0: "GOOGLE_MAIL",
	1: "GOOGLE_CALENDAR",
	2: "LINKED_IN",
	3: "FACEBOOK",
	4: "INSTAGRAM",
}
var SourceType_value = map[string]int32{
	"GOOGLE_MAIL":     0,
	"GOOGLE_CALENDAR": 1,
	"LINKED_IN":       2,
	"FACEBOOK":        3,
	"INSTAGRAM":       4,
}

func (x SourceType) String() string {
	return proto.EnumName(SourceType_name, int32(x))
}
func (SourceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Source struct {
	Id   string     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name string     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Type SourceType `protobuf:"varint,3,opt,name=type,enum=charlotte.common.source.v1.SourceType" json:"type,omitempty"`
}

func (m *Source) Reset()                    { *m = Source{} }
func (m *Source) String() string            { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()               {}
func (*Source) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Source) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Source) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Source) GetType() SourceType {
	if m != nil {
		return m.Type
	}
	return SourceType_GOOGLE_MAIL
}

func init() {
	proto.RegisterType((*Source)(nil), "charlotte.common.source.v1.Source")
	proto.RegisterEnum("charlotte.common.source.v1.SourceType", SourceType_name, SourceType_value)
}

func init() { proto.RegisterFile("charlotte/common/source/v1/source.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0xce, 0x48, 0x2c,
	0xca, 0xc9, 0x2f, 0x29, 0x49, 0xd5, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2f, 0xce, 0x2f,
	0x2d, 0x4a, 0x4e, 0xd5, 0x2f, 0x33, 0x84, 0xb2, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xa4,
	0xe0, 0x0a, 0xf5, 0x20, 0x0a, 0xf5, 0xa0, 0xd2, 0x65, 0x86, 0x4a, 0x19, 0x5c, 0x6c, 0xc1, 0x60,
	0x8e, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x53, 0x66,
	0x8a, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0x13, 0x58, 0x04, 0xcc, 0x16, 0xb2,
	0xe2, 0x62, 0x29, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x33, 0x52, 0xd3, 0xc3,
	0x6d, 0xb0, 0x1e, 0xc4, 0xd4, 0x90, 0xca, 0x82, 0xd4, 0x20, 0xb0, 0x1e, 0xad, 0x38, 0x2e, 0x2e,
	0x84, 0x98, 0x10, 0x3f, 0x17, 0xb7, 0xbb, 0xbf, 0xbf, 0xbb, 0x8f, 0x6b, 0xbc, 0xaf, 0xa3, 0xa7,
	0x8f, 0x00, 0x83, 0x90, 0x30, 0x17, 0x3f, 0x54, 0xc0, 0xd9, 0xd1, 0xc7, 0xd5, 0xcf, 0xc5, 0x31,
	0x48, 0x80, 0x51, 0x88, 0x97, 0x8b, 0xd3, 0xc7, 0xd3, 0xcf, 0xdb, 0xd5, 0x25, 0xde, 0xd3, 0x4f,
	0x80, 0x49, 0x88, 0x87, 0x8b, 0xc3, 0xcd, 0xd1, 0xd9, 0xd5, 0xc9, 0xdf, 0xdf, 0x5b, 0x80, 0x19,
	0x24, 0xe9, 0xe9, 0x17, 0x1c, 0xe2, 0xe8, 0x1e, 0xe4, 0xe8, 0x2b, 0xc0, 0x92, 0xc4, 0x06, 0xf6,
	0xac, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x83, 0xfe, 0xe5, 0x17, 0x01, 0x00, 0x00,
}