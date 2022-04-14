// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: epochstorage/endpoint.proto

package types

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

type Endpoint struct {
	IPPORT      string `protobuf:"bytes,1,opt,name=iPPORT,proto3" json:"iPPORT,omitempty"`
	UseType     string `protobuf:"bytes,2,opt,name=useType,proto3" json:"useType,omitempty"`
	Geolocation uint64 `protobuf:"varint,3,opt,name=geolocation,proto3" json:"geolocation,omitempty"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5b1ebaa0f5cf898, []int{0}
}
func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return m.Size()
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetIPPORT() string {
	if m != nil {
		return m.IPPORT
	}
	return ""
}

func (m *Endpoint) GetUseType() string {
	if m != nil {
		return m.UseType
	}
	return ""
}

func (m *Endpoint) GetGeolocation() uint64 {
	if m != nil {
		return m.Geolocation
	}
	return 0
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "lavanet.lava.epochstorage.Endpoint")
}

func init() { proto.RegisterFile("epochstorage/endpoint.proto", fileDescriptor_c5b1ebaa0f5cf898) }

var fileDescriptor_c5b1ebaa0f5cf898 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0x2d, 0xc8, 0x4f,
	0xce, 0x28, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f, 0xd5, 0x4f, 0xcd, 0x4b, 0x29, 0xc8, 0xcf, 0xcc,
	0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xcc, 0x49, 0x2c, 0x4b, 0xcc, 0x4b, 0x2d,
	0xd1, 0x03, 0xd1, 0x7a, 0xc8, 0x2a, 0x95, 0xe2, 0xb8, 0x38, 0x5c, 0xa1, 0x8a, 0x85, 0xc4, 0xb8,
	0xd8, 0x32, 0x03, 0x02, 0xfc, 0x83, 0x42, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xa0, 0x3c,
	0x21, 0x09, 0x2e, 0xf6, 0xd2, 0xe2, 0xd4, 0x90, 0xca, 0x82, 0x54, 0x09, 0x26, 0xb0, 0x04, 0x8c,
	0x2b, 0xa4, 0xc0, 0xc5, 0x9d, 0x9e, 0x9a, 0x9f, 0x93, 0x9f, 0x9c, 0x58, 0x92, 0x99, 0x9f, 0x27,
	0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x12, 0x84, 0x2c, 0xe4, 0xe4, 0x76, 0xe2, 0x91, 0x1c, 0xe3, 0x85,
	0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3,
	0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x3a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9,
	0xfa, 0x50, 0xf7, 0x81, 0x69, 0xfd, 0x0a, 0x7d, 0x14, 0xbf, 0x94, 0x54, 0x16, 0xa4, 0x16, 0x27,
	0xb1, 0x81, 0x7d, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xfd, 0x5e, 0x7a, 0xe8, 0x00,
	0x00, 0x00,
}

func (m *Endpoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Endpoint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Endpoint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Geolocation != 0 {
		i = encodeVarintEndpoint(dAtA, i, uint64(m.Geolocation))
		i--
		dAtA[i] = 0x18
	}
	if len(m.UseType) > 0 {
		i -= len(m.UseType)
		copy(dAtA[i:], m.UseType)
		i = encodeVarintEndpoint(dAtA, i, uint64(len(m.UseType)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.IPPORT) > 0 {
		i -= len(m.IPPORT)
		copy(dAtA[i:], m.IPPORT)
		i = encodeVarintEndpoint(dAtA, i, uint64(len(m.IPPORT)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEndpoint(dAtA []byte, offset int, v uint64) int {
	offset -= sovEndpoint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Endpoint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IPPORT)
	if l > 0 {
		n += 1 + l + sovEndpoint(uint64(l))
	}
	l = len(m.UseType)
	if l > 0 {
		n += 1 + l + sovEndpoint(uint64(l))
	}
	if m.Geolocation != 0 {
		n += 1 + sovEndpoint(uint64(m.Geolocation))
	}
	return n
}

func sovEndpoint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEndpoint(x uint64) (n int) {
	return sovEndpoint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Endpoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEndpoint
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
			return fmt.Errorf("proto: Endpoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Endpoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IPPORT", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEndpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IPPORT = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UseType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEndpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UseType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Geolocation", wireType)
			}
			m.Geolocation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Geolocation |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEndpoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEndpoint
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEndpoint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEndpoint
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
					return 0, ErrIntOverflowEndpoint
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
					return 0, ErrIntOverflowEndpoint
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
				return 0, ErrInvalidLengthEndpoint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEndpoint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEndpoint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEndpoint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEndpoint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEndpoint = fmt.Errorf("proto: unexpected end of group")
)
