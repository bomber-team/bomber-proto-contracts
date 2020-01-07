// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contracts/rest_contracts/configuration.proto

package rest_contracts

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

// send from main server to one instance bomber
type Configuration struct {
	MetricAddress        string   `protobuf:"bytes,1,opt,name=metricAddress,proto3" json:"metricAddress,omitempty"`
	ServerAddress        string   `protobuf:"bytes,2,opt,name=serverAddress,proto3" json:"serverAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d591ffb2b411585, []int{0}
}

func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (m *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(m, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetMetricAddress() string {
	if m != nil {
		return m.MetricAddress
	}
	return ""
}

func (m *Configuration) GetServerAddress() string {
	if m != nil {
		return m.ServerAddress
	}
	return ""
}

// send to main server on multicast
type BomberConfig struct {
	BomberId             string   `protobuf:"bytes,1,opt,name=bomberId,proto3" json:"bomberId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BomberConfig) Reset()         { *m = BomberConfig{} }
func (m *BomberConfig) String() string { return proto.CompactTextString(m) }
func (*BomberConfig) ProtoMessage()    {}
func (*BomberConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d591ffb2b411585, []int{1}
}

func (m *BomberConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BomberConfig.Unmarshal(m, b)
}
func (m *BomberConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BomberConfig.Marshal(b, m, deterministic)
}
func (m *BomberConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BomberConfig.Merge(m, src)
}
func (m *BomberConfig) XXX_Size() int {
	return xxx_messageInfo_BomberConfig.Size(m)
}
func (m *BomberConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BomberConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BomberConfig proto.InternalMessageInfo

func (m *BomberConfig) GetBomberId() string {
	if m != nil {
		return m.BomberId
	}
	return ""
}

func init() {
	proto.RegisterType((*Configuration)(nil), "contracts.rest_contracts.Configuration")
	proto.RegisterType((*BomberConfig)(nil), "contracts.rest_contracts.BomberConfig")
}

func init() {
	proto.RegisterFile("contracts/rest_contracts/configuration.proto", fileDescriptor_1d591ffb2b411585)
}

var fileDescriptor_1d591ffb2b411585 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x49, 0xce, 0xcf, 0x2b,
	0x29, 0x4a, 0x4c, 0x2e, 0x29, 0xd6, 0x2f, 0x4a, 0x2d, 0x2e, 0x89, 0x47, 0x70, 0x93, 0xf3, 0xf3,
	0xd2, 0x32, 0xd3, 0x4b, 0x8b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0x24, 0xe0, 0xd2, 0x7a, 0xa8, 0xaa, 0x95, 0xa2, 0xb9, 0x78, 0x9d, 0x91, 0x35, 0x08, 0xa9,
	0x70, 0xf1, 0xe6, 0xa6, 0x96, 0x14, 0x65, 0x26, 0x3b, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x4b,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xa1, 0x0a, 0x82, 0x54, 0x15, 0xa7, 0x16, 0x95, 0xa5, 0x16,
	0xc1, 0x54, 0x31, 0x41, 0x54, 0xa1, 0x08, 0x2a, 0x69, 0x71, 0xf1, 0x38, 0xe5, 0xe7, 0x26, 0xa5,
	0x16, 0x41, 0xac, 0x10, 0x92, 0xe2, 0xe2, 0x48, 0x02, 0xf3, 0x3d, 0x53, 0xa0, 0xc6, 0xc2, 0xf9,
	0x4e, 0x09, 0x5c, 0x92, 0xf9, 0x45, 0xe9, 0x7a, 0x10, 0xbe, 0x5e, 0x49, 0x6a, 0x62, 0xae, 0x1e,
	0xdc, 0x95, 0x51, 0xce, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x10,
	0x15, 0xf1, 0x20, 0x15, 0x30, 0x36, 0xd8, 0x93, 0x48, 0x41, 0x90, 0x9e, 0x9f, 0x93, 0x98, 0x97,
	0x8e, 0x16, 0x30, 0x49, 0x6c, 0x60, 0x65, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0x6f,
	0x6f, 0x4e, 0x3b, 0x01, 0x00, 0x00,
}