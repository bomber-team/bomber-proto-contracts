// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: contracts/rest_contracts/script.proto

package rest_contracts

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type RestScript struct {
	SchemeId             string               `protobuf:"bytes,1,opt,name=schemeId,proto3" json:"schemeId,omitempty"`
	Address              string               `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	RequestMethod        string               `protobuf:"bytes,3,opt,name=requestMethod,proto3" json:"requestMethod,omitempty"`
	Config               *ConfigurationScript `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RestScript) Reset()         { *m = RestScript{} }
func (m *RestScript) String() string { return proto.CompactTextString(m) }
func (*RestScript) ProtoMessage()    {}
func (*RestScript) Descriptor() ([]byte, []int) {
	return fileDescriptor_91f6f3ec99002da4, []int{0}
}
func (m *RestScript) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RestScript) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RestScript.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RestScript) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestScript.Merge(m, src)
}
func (m *RestScript) XXX_Size() int {
	return m.Size()
}
func (m *RestScript) XXX_DiscardUnknown() {
	xxx_messageInfo_RestScript.DiscardUnknown(m)
}

var xxx_messageInfo_RestScript proto.InternalMessageInfo

func (m *RestScript) GetSchemeId() string {
	if m != nil {
		return m.SchemeId
	}
	return ""
}

func (m *RestScript) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RestScript) GetRequestMethod() string {
	if m != nil {
		return m.RequestMethod
	}
	return ""
}

func (m *RestScript) GetConfig() *ConfigurationScript {
	if m != nil {
		return m.Config
	}
	return nil
}

type ConfigurationScript struct {
	Rps                  int64    `protobuf:"varint,1,opt,name=rps,proto3" json:"rps,omitempty"`
	Time                 int64    `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigurationScript) Reset()         { *m = ConfigurationScript{} }
func (m *ConfigurationScript) String() string { return proto.CompactTextString(m) }
func (*ConfigurationScript) ProtoMessage()    {}
func (*ConfigurationScript) Descriptor() ([]byte, []int) {
	return fileDescriptor_91f6f3ec99002da4, []int{1}
}
func (m *ConfigurationScript) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigurationScript) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigurationScript.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigurationScript) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigurationScript.Merge(m, src)
}
func (m *ConfigurationScript) XXX_Size() int {
	return m.Size()
}
func (m *ConfigurationScript) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigurationScript.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigurationScript proto.InternalMessageInfo

func (m *ConfigurationScript) GetRps() int64 {
	if m != nil {
		return m.Rps
	}
	return 0
}

func (m *ConfigurationScript) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func init() {
	proto.RegisterType((*RestScript)(nil), "contracts.rest_contracts.RestScript")
	proto.RegisterType((*ConfigurationScript)(nil), "contracts.rest_contracts.ConfigurationScript")
}

func init() {
	proto.RegisterFile("contracts/rest_contracts/script.proto", fileDescriptor_91f6f3ec99002da4)
}

var fileDescriptor_91f6f3ec99002da4 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x59, 0x53, 0xaa, 0x8e, 0x08, 0xb2, 0x5e, 0x56, 0x0f, 0xa1, 0x14, 0x85, 0x5e, 0xdc,
	0x82, 0x1e, 0xbd, 0x59, 0x3c, 0x78, 0xf0, 0x12, 0x6f, 0x5e, 0x4a, 0xb2, 0x19, 0x93, 0x40, 0x93,
	0x89, 0xb3, 0x93, 0xe7, 0xf2, 0x35, 0x3c, 0xfa, 0x08, 0x92, 0x27, 0x91, 0x6e, 0x6d, 0x6b, 0x45,
	0x6f, 0xf3, 0xcf, 0xff, 0x2d, 0x7c, 0xcc, 0xc2, 0xa5, 0xa3, 0x46, 0x38, 0x75, 0xe2, 0xa7, 0x8c,
	0x5e, 0xe6, 0xdb, 0xe8, 0x1d, 0x57, 0xad, 0xd8, 0x96, 0x49, 0x48, 0x9b, 0xcd, 0xde, 0xee, 0x62,
	0xe3, 0x37, 0x05, 0x90, 0xa0, 0x97, 0xa7, 0x80, 0xeb, 0x73, 0x38, 0xf0, 0xae, 0xc4, 0x1a, 0x1f,
	0x72, 0xa3, 0x46, 0x6a, 0x72, 0x98, 0x6c, 0xb2, 0x36, 0xb0, 0x9f, 0xe6, 0x39, 0xa3, 0xf7, 0x66,
	0x2f, 0x54, 0xeb, 0xa8, 0x2f, 0xe0, 0x98, 0xf1, 0xb5, 0x43, 0x2f, 0x8f, 0x28, 0x25, 0xe5, 0x26,
	0x0a, 0xfd, 0xee, 0x52, 0xdf, 0xc3, 0xd0, 0x51, 0xf3, 0x52, 0x15, 0x66, 0x30, 0x52, 0x93, 0xa3,
	0xeb, 0x2b, 0xfb, 0x9f, 0x95, 0x9d, 0x05, 0xae, 0xe3, 0x54, 0x2a, 0x6a, 0x56, 0x6a, 0xc9, 0xf7,
	0xe3, 0xf1, 0x2d, 0x9c, 0xfe, 0x51, 0xeb, 0x13, 0x88, 0xb8, 0xf5, 0x41, 0x3a, 0x4a, 0x96, 0xa3,
	0xd6, 0x30, 0x90, 0xaa, 0xc6, 0x20, 0x1b, 0x25, 0x61, 0xbe, 0x5b, 0xbc, 0xf7, 0xb1, 0xfa, 0xe8,
	0x63, 0xf5, 0xd9, 0xc7, 0x0a, 0xce, 0x88, 0x0b, 0x9b, 0x51, 0x9d, 0x21, 0x5b, 0xc1, 0xb4, 0xde,
	0x0a, 0x3d, 0xcf, 0x8a, 0x4a, 0xca, 0x2e, 0xb3, 0x8e, 0xea, 0xe9, 0x8a, 0x98, 0x2f, 0x89, 0xf5,
	0x1c, 0xce, 0xfa, 0xe3, 0xda, 0x05, 0x2d, 0xd2, 0xa6, 0xf8, 0xf5, 0x07, 0xd9, 0x30, 0x60, 0x37,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x68, 0xb9, 0x5f, 0xa8, 0xa6, 0x01, 0x00, 0x00,
}

func (m *RestScript) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RestScript) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RestScript) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Config != nil {
		{
			size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintScript(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.RequestMethod) > 0 {
		i -= len(m.RequestMethod)
		copy(dAtA[i:], m.RequestMethod)
		i = encodeVarintScript(dAtA, i, uint64(len(m.RequestMethod)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintScript(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SchemeId) > 0 {
		i -= len(m.SchemeId)
		copy(dAtA[i:], m.SchemeId)
		i = encodeVarintScript(dAtA, i, uint64(len(m.SchemeId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConfigurationScript) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigurationScript) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfigurationScript) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Time != 0 {
		i = encodeVarintScript(dAtA, i, uint64(m.Time))
		i--
		dAtA[i] = 0x10
	}
	if m.Rps != 0 {
		i = encodeVarintScript(dAtA, i, uint64(m.Rps))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintScript(dAtA []byte, offset int, v uint64) int {
	offset -= sovScript(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RestScript) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SchemeId)
	if l > 0 {
		n += 1 + l + sovScript(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovScript(uint64(l))
	}
	l = len(m.RequestMethod)
	if l > 0 {
		n += 1 + l + sovScript(uint64(l))
	}
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovScript(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ConfigurationScript) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Rps != 0 {
		n += 1 + sovScript(uint64(m.Rps))
	}
	if m.Time != 0 {
		n += 1 + sovScript(uint64(m.Time))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovScript(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozScript(x uint64) (n int) {
	return sovScript(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RestScript) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScript
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RestScript: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RestScript: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScript
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthScript
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScript
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SchemeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScript
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthScript
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScript
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestMethod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScript
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthScript
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScript
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestMethod = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScript
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthScript
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthScript
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &ConfigurationScript{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipScript(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthScript
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthScript
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConfigurationScript) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScript
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConfigurationScript: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigurationScript: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rps", wireType)
			}
			m.Rps = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScript
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rps |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScript
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipScript(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthScript
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthScript
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipScript(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowScript
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowScript
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowScript
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthScript
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupScript
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthScript
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthScript        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowScript          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupScript = fmt.Errorf("proto: unexpected end of group")
)
