// Code generated by protoc-gen-go.
// source: addressbook.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	addressbook.proto

It has these top-level messages:
	Person
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}
func (Person_PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Person struct {
	Name   string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id     int32                 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Email  string                `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phones []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type Person_PhoneNumber struct {
	Number string           `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type   Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=pb.Person_PhoneType" json:"type,omitempty"`
}

func (m *Person_PhoneNumber) Reset()                    { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()               {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func init() {
	proto.RegisterType((*Person)(nil), "pb.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "pb.Person.PhoneNumber")
	proto.RegisterEnum("pb.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
}

var fileDescriptor0 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4c, 0x49, 0x29,
	0x4a, 0x2d, 0x2e, 0x4e, 0xca, 0xcf, 0xcf, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a,
	0x48, 0x52, 0x3a, 0xc9, 0xc8, 0xc5, 0x16, 0x90, 0x5a, 0x54, 0x9c, 0x9f, 0x27, 0xc4, 0xc3, 0xc5,
	0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x29, 0xc4, 0xc5, 0xc5, 0x94, 0x99,
	0x22, 0xc1, 0x04, 0x64, 0xb3, 0x0a, 0xf1, 0x72, 0xb1, 0xa6, 0xe6, 0x26, 0x66, 0xe6, 0x48, 0x30,
	0x83, 0xa5, 0xd4, 0xb8, 0xd8, 0x0a, 0x32, 0xf2, 0xf3, 0x52, 0x8b, 0x25, 0x58, 0x14, 0x98, 0x35,
	0xb8, 0x8d, 0xc4, 0xf4, 0x0a, 0x92, 0xf4, 0x20, 0x86, 0xe8, 0x05, 0x80, 0x24, 0xfc, 0x4a, 0x73,
	0x93, 0x52, 0x8b, 0xa4, 0x1c, 0xb9, 0xb8, 0x91, 0xb8, 0x42, 0x7c, 0x5c, 0x6c, 0x79, 0x60, 0x16,
	0xd4, 0x06, 0x25, 0x2e, 0x96, 0x92, 0xca, 0x82, 0x54, 0xb0, 0x1d, 0x7c, 0x46, 0x22, 0xe8, 0x86,
	0x84, 0x00, 0xe5, 0x94, 0xb4, 0xb9, 0x38, 0xe1, 0x1c, 0xa0, 0x93, 0xd8, 0x7c, 0xfd, 0x9d, 0x3c,
	0x7d, 0x5c, 0x05, 0x18, 0x84, 0x38, 0xb8, 0x58, 0x3c, 0xfc, 0x7d, 0x5d, 0x05, 0x18, 0x41, 0xac,
	0x70, 0xff, 0x20, 0x6f, 0x01, 0xa6, 0x24, 0x36, 0xb0, 0xb7, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x5a, 0xd6, 0xcd, 0x8b, 0xeb, 0x00, 0x00, 0x00,
}
