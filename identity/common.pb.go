// Code generated by protoc-gen-go.
// source: common.proto
// DO NOT EDIT!

/*
Package identity is a generated protocol buffer package.

It is generated from these files:
	common.proto
	organization.proto

It has these top-level messages:
	PublicKey
	Signature
	Organization
	SignedCertificate
	CertificateData
	OrganizationDetails
*/
package identity

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

type PublicKey_Algorithm int32

const (
	PublicKey_UNKNOWN PublicKey_Algorithm = 0
	PublicKey_RSA     PublicKey_Algorithm = 1
	PublicKey_ECDSA   PublicKey_Algorithm = 2
)

var PublicKey_Algorithm_name = map[int32]string{
	0: "UNKNOWN",
	1: "RSA",
	2: "ECDSA",
}
var PublicKey_Algorithm_value = map[string]int32{
	"UNKNOWN": 0,
	"RSA":     1,
	"ECDSA":   2,
}

func (x PublicKey_Algorithm) String() string {
	return proto.EnumName(PublicKey_Algorithm_name, int32(x))
}
func (PublicKey_Algorithm) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type PublicKey struct {
	Algorithm PublicKey_Algorithm `protobuf:"varint,1,opt,name=algorithm,enum=identity.proto.PublicKey_Algorithm" json:"algorithm,omitempty"`
	Data      []byte              `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *PublicKey) Reset()                    { *m = PublicKey{} }
func (m *PublicKey) String() string            { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()               {}
func (*PublicKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PublicKey) GetAlgorithm() PublicKey_Algorithm {
	if m != nil {
		return m.Algorithm
	}
	return PublicKey_UNKNOWN
}

func (m *PublicKey) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Signature struct {
	SignerCertId string `protobuf:"bytes,1,opt,name=signer_cert_id,json=signerCertId" json:"signer_cert_id,omitempty"`
	Data         []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Timestamp    uint64 `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *Signature) Reset()                    { *m = Signature{} }
func (m *Signature) String() string            { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()               {}
func (*Signature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Signature) GetSignerCertId() string {
	if m != nil {
		return m.SignerCertId
	}
	return ""
}

func (m *Signature) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Signature) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*PublicKey)(nil), "identity.proto.PublicKey")
	proto.RegisterType((*Signature)(nil), "identity.proto.Signature")
	proto.RegisterEnum("identity.proto.PublicKey_Algorithm", PublicKey_Algorithm_name, PublicKey_Algorithm_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x8f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xdd, 0xb6, 0x5a, 0x77, 0x0c, 0x21, 0xec, 0x29, 0x07, 0xc1, 0x10, 0x3d, 0xe4, 0xa0,
	0x39, 0xe8, 0x13, 0xc4, 0xea, 0x41, 0x0a, 0xb1, 0x24, 0x88, 0xe0, 0xa5, 0x6c, 0x36, 0x4b, 0x3a,
	0x90, 0xcd, 0x86, 0xcd, 0xf4, 0xd0, 0xb7, 0xf0, 0x91, 0x85, 0x94, 0xb6, 0x28, 0xde, 0x66, 0x3e,
	0xbe, 0x9f, 0x99, 0x1f, 0x3c, 0x65, 0x8d, 0xb1, 0x5d, 0xda, 0x3b, 0x4b, 0x56, 0xf8, 0x58, 0xeb,
	0x8e, 0x90, 0x76, 0xfb, 0x3d, 0xfe, 0x66, 0xc0, 0x57, 0xdb, 0xaa, 0x45, 0xb5, 0xd4, 0x3b, 0x91,
	0x01, 0x97, 0x6d, 0x63, 0x1d, 0xd2, 0xc6, 0x84, 0x2c, 0x62, 0x89, 0xff, 0x78, 0x9b, 0xfe, 0x4e,
	0xa4, 0x47, 0x3b, 0xcd, 0x0e, 0x6a, 0x71, 0x4a, 0x09, 0x01, 0xb3, 0x5a, 0x92, 0x0c, 0x27, 0x11,
	0x4b, 0xbc, 0x62, 0x9c, 0xe3, 0x7b, 0xe0, 0x47, 0x57, 0x5c, 0xc1, 0xfc, 0x23, 0x5f, 0xe6, 0xef,
	0x9f, 0x79, 0x70, 0x26, 0xe6, 0x30, 0x2d, 0xca, 0x2c, 0x60, 0x82, 0xc3, 0xf9, 0xeb, 0xe2, 0xa5,
	0xcc, 0x82, 0x49, 0xac, 0x80, 0x97, 0xd8, 0x74, 0x92, 0xb6, 0x4e, 0x8b, 0x3b, 0xf0, 0x07, 0x6c,
	0x3a, 0xed, 0xd6, 0x4a, 0x3b, 0x5a, 0x63, 0x3d, 0xbe, 0xc5, 0x0b, 0x6f, 0x4f, 0x17, 0xda, 0xd1,
	0x5b, 0xfd, 0xdf, 0x51, 0x71, 0x0d, 0x9c, 0xd0, 0xe8, 0x81, 0xa4, 0xe9, 0xc3, 0x69, 0xc4, 0x92,
	0x59, 0x71, 0x02, 0xcf, 0x0f, 0x70, 0xa3, 0xac, 0x49, 0xfb, 0x01, 0xfb, 0x0d, 0x36, 0xad, 0xad,
	0x64, 0xfb, 0xa7, 0xe8, 0x8a, 0x7d, 0x5d, 0x1e, 0x48, 0x75, 0x31, 0xa2, 0xa7, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xff, 0x51, 0x35, 0xe5, 0x4d, 0x01, 0x00, 0x00,
}
