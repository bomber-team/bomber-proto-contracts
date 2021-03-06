// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: contracts/system/status_changer.proto

package system

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

type StatusBomber int32

const (
	StatusBomber_DOWN           StatusBomber = 0
	StatusBomber_UP             StatusBomber = 1
	StatusBomber_WORKING        StatusBomber = 2
	StatusBomber_PREPARING_DATA StatusBomber = 3
)

var StatusBomber_name = map[int32]string{
	0: "DOWN",
	1: "UP",
	2: "WORKING",
	3: "PREPARING_DATA",
}

var StatusBomber_value = map[string]int32{
	"DOWN":           0,
	"UP":             1,
	"WORKING":        2,
	"PREPARING_DATA": 3,
}

func (x StatusBomber) String() string {
	return proto.EnumName(StatusBomber_name, int32(x))
}

func (StatusBomber) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f01e2bcb1a7d6b50, []int{0}
}

// send to main bomber server
type BomberStatusChange struct {
	BomberId             string       `protobuf:"bytes,1,opt,name=bomberId,proto3" json:"bomberId,omitempty"`
	Ip                   string       `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	StatusBomber         StatusBomber `protobuf:"varint,3,opt,name=statusBomber,proto3,enum=contracts.system.StatusBomber" json:"statusBomber,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BomberStatusChange) Reset()         { *m = BomberStatusChange{} }
func (m *BomberStatusChange) String() string { return proto.CompactTextString(m) }
func (*BomberStatusChange) ProtoMessage()    {}
func (*BomberStatusChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_f01e2bcb1a7d6b50, []int{0}
}
func (m *BomberStatusChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BomberStatusChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BomberStatusChange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BomberStatusChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BomberStatusChange.Merge(m, src)
}
func (m *BomberStatusChange) XXX_Size() int {
	return m.Size()
}
func (m *BomberStatusChange) XXX_DiscardUnknown() {
	xxx_messageInfo_BomberStatusChange.DiscardUnknown(m)
}

var xxx_messageInfo_BomberStatusChange proto.InternalMessageInfo

func (m *BomberStatusChange) GetBomberId() string {
	if m != nil {
		return m.BomberId
	}
	return ""
}

func (m *BomberStatusChange) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *BomberStatusChange) GetStatusBomber() StatusBomber {
	if m != nil {
		return m.StatusBomber
	}
	return StatusBomber_DOWN
}

func init() {
	proto.RegisterEnum("contracts.system.StatusBomber", StatusBomber_name, StatusBomber_value)
	proto.RegisterType((*BomberStatusChange)(nil), "contracts.system.BomberStatusChange")
}

func init() {
	proto.RegisterFile("contracts/system/status_changer.proto", fileDescriptor_f01e2bcb1a7d6b50)
}

var fileDescriptor_f01e2bcb1a7d6b50 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xce, 0xcf, 0x2b,
	0x29, 0x4a, 0x4c, 0x2e, 0x29, 0xd6, 0x2f, 0xae, 0x2c, 0x2e, 0x49, 0xcd, 0xd5, 0x2f, 0x2e, 0x49,
	0x2c, 0x29, 0x2d, 0x8e, 0x4f, 0xce, 0x48, 0xcc, 0x4b, 0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x80, 0x2b, 0xd3, 0x83, 0x28, 0x53, 0x6a, 0x61, 0xe4, 0x12, 0x72, 0xca, 0xcf,
	0x4d, 0x4a, 0x2d, 0x0a, 0x06, 0x6b, 0x70, 0x06, 0xab, 0x17, 0x92, 0xe2, 0xe2, 0x48, 0x02, 0x8b,
	0x7a, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x42, 0x7c, 0x5c, 0x4c, 0x99,
	0x05, 0x12, 0x4c, 0x60, 0x51, 0xa6, 0xcc, 0x02, 0x21, 0x27, 0x2e, 0x1e, 0x88, 0x65, 0x10, 0x73,
	0x24, 0x98, 0x15, 0x18, 0x35, 0xf8, 0x8c, 0xe4, 0xf4, 0xd0, 0xed, 0xd2, 0x0b, 0x46, 0x52, 0x15,
	0x84, 0xa2, 0x47, 0xcb, 0x91, 0x8b, 0x07, 0x59, 0x56, 0x88, 0x83, 0x8b, 0xc5, 0xc5, 0x3f, 0xdc,
	0x4f, 0x80, 0x41, 0x88, 0x8d, 0x8b, 0x29, 0x34, 0x40, 0x80, 0x51, 0x88, 0x9b, 0x8b, 0x3d, 0xdc,
	0x3f, 0xc8, 0xdb, 0xd3, 0xcf, 0x5d, 0x80, 0x49, 0x48, 0x88, 0x8b, 0x2f, 0x20, 0xc8, 0x35, 0xc0,
	0x31, 0xc8, 0xd3, 0xcf, 0x3d, 0xde, 0xc5, 0x31, 0xc4, 0x51, 0x80, 0xd9, 0x29, 0xe5, 0xc4, 0x23,
	0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0xe4, 0x92, 0xcc, 0x2f, 0x4a, 0xd7,
	0x83, 0x38, 0x59, 0xaf, 0x24, 0x35, 0x31, 0x17, 0xe1, 0x9a, 0x28, 0xeb, 0xf4, 0xcc, 0x92, 0x8c,
	0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x88, 0x8a, 0x78, 0x90, 0x0a, 0x18, 0x1b, 0x1c, 0x52,
	0xf1, 0x88, 0xf0, 0x4c, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0x87, 0x06, 0x6b, 0x12, 0x1b, 0x58, 0xda,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xf7, 0xfb, 0x00, 0x71, 0x01, 0x00, 0x00,
}

func (m *BomberStatusChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BomberStatusChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BomberStatusChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.StatusBomber != 0 {
		i = encodeVarintStatusChanger(dAtA, i, uint64(m.StatusBomber))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Ip) > 0 {
		i -= len(m.Ip)
		copy(dAtA[i:], m.Ip)
		i = encodeVarintStatusChanger(dAtA, i, uint64(len(m.Ip)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BomberId) > 0 {
		i -= len(m.BomberId)
		copy(dAtA[i:], m.BomberId)
		i = encodeVarintStatusChanger(dAtA, i, uint64(len(m.BomberId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStatusChanger(dAtA []byte, offset int, v uint64) int {
	offset -= sovStatusChanger(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BomberStatusChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BomberId)
	if l > 0 {
		n += 1 + l + sovStatusChanger(uint64(l))
	}
	l = len(m.Ip)
	if l > 0 {
		n += 1 + l + sovStatusChanger(uint64(l))
	}
	if m.StatusBomber != 0 {
		n += 1 + sovStatusChanger(uint64(m.StatusBomber))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovStatusChanger(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStatusChanger(x uint64) (n int) {
	return sovStatusChanger(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BomberStatusChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatusChanger
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
			return fmt.Errorf("proto: BomberStatusChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BomberStatusChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BomberId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatusChanger
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
				return ErrInvalidLengthStatusChanger
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatusChanger
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BomberId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatusChanger
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
				return ErrInvalidLengthStatusChanger
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatusChanger
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ip = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusBomber", wireType)
			}
			m.StatusBomber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatusChanger
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StatusBomber |= StatusBomber(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStatusChanger(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatusChanger
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStatusChanger
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
func skipStatusChanger(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStatusChanger
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
					return 0, ErrIntOverflowStatusChanger
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
					return 0, ErrIntOverflowStatusChanger
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
				return 0, ErrInvalidLengthStatusChanger
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStatusChanger
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStatusChanger
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStatusChanger        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStatusChanger          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStatusChanger = fmt.Errorf("proto: unexpected end of group")
)
