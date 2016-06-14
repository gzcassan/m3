// Code generated by protoc-gen-go.
// source: rawdata.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	rawdata.proto

It has these top-level messages:
	IndexEntry
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto1.ProtoPackageIsVersion1

type IndexEntry struct {
	Id  string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Idx int32  `protobuf:"varint,2,opt,name=idx" json:"idx,omitempty"`
}

func (m *IndexEntry) Reset()                    { *m = IndexEntry{} }
func (m *IndexEntry) String() string            { return proto1.CompactTextString(m) }
func (*IndexEntry) ProtoMessage()               {}
func (*IndexEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto1.RegisterType((*IndexEntry)(nil), "proto.IndexEntry")
}

var fileDescriptor0 = []byte{
	// 91 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4a, 0x2c, 0x4f,
	0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x7a, 0x5c,
	0x5c, 0x9e, 0x79, 0x29, 0xa9, 0x15, 0xae, 0x79, 0x25, 0x45, 0x95, 0x42, 0x7c, 0x5c, 0x4c, 0x99,
	0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x40, 0x96, 0x90, 0x00, 0x17, 0x73, 0x66, 0x4a,
	0x85, 0x04, 0x13, 0x50, 0x80, 0x35, 0x08, 0xc4, 0x4c, 0x62, 0x03, 0x6b, 0x33, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x11, 0x54, 0x9f, 0xb2, 0x4e, 0x00, 0x00, 0x00,
}