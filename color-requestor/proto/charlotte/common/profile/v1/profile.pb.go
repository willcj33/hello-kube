// Code generated by protoc-gen-go.
// source: charlotte/common/profile/v1/profile.proto
// DO NOT EDIT!

/*
Package charlotte_common_profile_v1 is a generated protocol buffer package.

It is generated from these files:
	charlotte/common/profile/v1/profile.proto

It has these top-level messages:
	Profile
*/
package charlotte_common_profile_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import charlotte_common_traits_profile_v1 "charlotte/common/traits/profile/v1"
import charlotte_common_source_v1 "charlotte/common/source/v1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Profile struct {
	Id               string                                       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name             string                                       `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	CreatedOn        int64                                        `protobuf:"varint,3,opt,name=createdOn" json:"createdOn,omitempty"`
	SourceIdentifier string                                       `protobuf:"bytes,4,opt,name=sourceIdentifier" json:"sourceIdentifier,omitempty"`
	SourceType       charlotte_common_source_v1.SourceType        `protobuf:"varint,5,opt,name=sourceType,enum=charlotte.common.source.v1.SourceType" json:"sourceType,omitempty"`
	Traits           []*charlotte_common_traits_profile_v1.ITrait `protobuf:"bytes,6,rep,name=traits" json:"traits,omitempty"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Profile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Profile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Profile) GetCreatedOn() int64 {
	if m != nil {
		return m.CreatedOn
	}
	return 0
}

func (m *Profile) GetSourceIdentifier() string {
	if m != nil {
		return m.SourceIdentifier
	}
	return ""
}

func (m *Profile) GetSourceType() charlotte_common_source_v1.SourceType {
	if m != nil {
		return m.SourceType
	}
	return charlotte_common_source_v1.SourceType_GOOGLE_MAIL
}

func (m *Profile) GetTraits() []*charlotte_common_traits_profile_v1.ITrait {
	if m != nil {
		return m.Traits
	}
	return nil
}

func init() {
	proto.RegisterType((*Profile)(nil), "charlotte.common.profile.v1.Profile")
}

func init() { proto.RegisterFile("charlotte/common/profile/v1/profile.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x52, 0x23, 0x1d, 0xa1, 0xc8, 0x9c, 0x96, 0xea, 0x21, 0x78, 0xd0, 0xd8, 0xc3,
	0x86, 0xd4, 0x37, 0xf0, 0x20, 0xf4, 0xa4, 0xc4, 0xbe, 0x40, 0xcc, 0x4e, 0x71, 0xa1, 0xd9, 0x0d,
	0xdb, 0x35, 0xe0, 0xd9, 0x17, 0x17, 0x77, 0xda, 0x1a, 0xd8, 0xdb, 0xf0, 0xed, 0xff, 0x7f, 0x2c,
	0x3f, 0x3c, 0x76, 0x9f, 0xad, 0xdb, 0x5b, 0xef, 0xa9, 0xea, 0x6c, 0xdf, 0x5b, 0x53, 0x0d, 0xce,
	0xee, 0xf4, 0x9e, 0xaa, 0xb1, 0x3e, 0x9d, 0x72, 0x70, 0xd6, 0x5b, 0xbc, 0x39, 0x47, 0x25, 0x47,
	0xe5, 0xe9, 0x7d, 0xac, 0x97, 0x55, 0xe4, 0xf1, 0xae, 0xd5, 0xfe, 0x30, 0xd5, 0xe9, 0x80, 0xd8,
	0xb6, 0x7c, 0x88, 0x0a, 0x07, 0xfb, 0xe5, 0xba, 0x10, 0xe4, 0x8b, 0x83, 0x77, 0x3f, 0x29, 0x5c,
	0xbe, 0xb1, 0x04, 0x17, 0x90, 0x6a, 0x25, 0x92, 0x22, 0x29, 0xe7, 0x4d, 0xaa, 0x15, 0x22, 0xcc,
	0x4c, 0xdb, 0x93, 0x48, 0x03, 0x09, 0x37, 0xde, 0xc2, 0xbc, 0x73, 0xd4, 0x7a, 0x52, 0xaf, 0x46,
	0x64, 0x45, 0x52, 0x66, 0xcd, 0x3f, 0xc0, 0x15, 0x5c, 0xb3, 0x7d, 0xa3, 0xc8, 0x78, 0xbd, 0xd3,
	0xe4, 0xc4, 0x2c, 0xb4, 0x23, 0x8e, 0x2f, 0x00, 0xcc, 0xb6, 0xdf, 0x03, 0x89, 0x8b, 0x22, 0x29,
	0x17, 0xeb, 0x7b, 0x19, 0xad, 0x70, 0xfc, 0xed, 0x58, 0xcb, 0xf7, 0x73, 0xba, 0x99, 0x34, 0xf1,
	0x19, 0x72, 0x1e, 0x43, 0xe4, 0x45, 0x56, 0x5e, 0xad, 0x57, 0xb1, 0x83, 0xdf, 0x27, 0x83, 0xca,
	0xcd, 0xf6, 0x0f, 0x35, 0xc7, 0xe6, 0x47, 0x1e, 0xc6, 0x78, 0xfa, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0xef, 0xf6, 0x41, 0x72, 0xb0, 0x01, 0x00, 0x00,
}
