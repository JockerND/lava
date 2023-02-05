// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: packagemanager/package.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Package struct {
	Index                string     `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Duration             uint64     `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	Epoch                uint64     `protobuf:"varint,3,opt,name=epoch,proto3" json:"epoch,omitempty"`
	Price                types.Coin `protobuf:"bytes,4,opt,name=price,proto3" json:"price"`
	ComputeUnits         uint64     `protobuf:"varint,5,opt,name=computeUnits,proto3" json:"computeUnits,omitempty"`
	ComputeUnitsPerEpoch uint64     `protobuf:"varint,6,opt,name=computeUnitsPerEpoch,proto3" json:"computeUnitsPerEpoch,omitempty"`
	ServicersToPair      uint64     `protobuf:"varint,7,opt,name=servicersToPair,proto3" json:"servicersToPair,omitempty"`
	AllowOveruse         bool       `protobuf:"varint,8,opt,name=allowOveruse,proto3" json:"allowOveruse,omitempty"`
	OveruseRate          uint64     `protobuf:"varint,9,opt,name=overuseRate,proto3" json:"overuseRate,omitempty"`
	Subscriptions        uint64     `protobuf:"varint,10,opt,name=subscriptions,proto3" json:"subscriptions,omitempty"`
	Name                 string     `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Package) Reset()         { *m = Package{} }
func (m *Package) String() string { return proto.CompactTextString(m) }
func (*Package) ProtoMessage()    {}
func (*Package) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5556970f3bd8668, []int{0}
}
func (m *Package) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Package) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Package.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Package) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Package.Merge(m, src)
}
func (m *Package) XXX_Size() int {
	return m.Size()
}
func (m *Package) XXX_DiscardUnknown() {
	xxx_messageInfo_Package.DiscardUnknown(m)
}

var xxx_messageInfo_Package proto.InternalMessageInfo

func (m *Package) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Package) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Package) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *Package) GetPrice() types.Coin {
	if m != nil {
		return m.Price
	}
	return types.Coin{}
}

func (m *Package) GetComputeUnits() uint64 {
	if m != nil {
		return m.ComputeUnits
	}
	return 0
}

func (m *Package) GetComputeUnitsPerEpoch() uint64 {
	if m != nil {
		return m.ComputeUnitsPerEpoch
	}
	return 0
}

func (m *Package) GetServicersToPair() uint64 {
	if m != nil {
		return m.ServicersToPair
	}
	return 0
}

func (m *Package) GetAllowOveruse() bool {
	if m != nil {
		return m.AllowOveruse
	}
	return false
}

func (m *Package) GetOveruseRate() uint64 {
	if m != nil {
		return m.OveruseRate
	}
	return 0
}

func (m *Package) GetSubscriptions() uint64 {
	if m != nil {
		return m.Subscriptions
	}
	return 0
}

func (m *Package) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Package)(nil), "lavanet.lava.packagemanager.Package")
}

func init() { proto.RegisterFile("packagemanager/package.proto", fileDescriptor_a5556970f3bd8668) }

var fileDescriptor_a5556970f3bd8668 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x8d, 0xd9, 0x74, 0xb7, 0xeb, 0x82, 0x90, 0xac, 0x1e, 0x4c, 0x41, 0x21, 0x5a, 0x71, 0xc8,
	0xc9, 0xd6, 0x2e, 0xe2, 0x07, 0x16, 0x21, 0x71, 0xa3, 0x8a, 0xe0, 0xc2, 0xcd, 0xf1, 0x8e, 0x52,
	0x8b, 0xc6, 0x8e, 0x6c, 0x27, 0x94, 0x5f, 0xe0, 0xc4, 0x67, 0xf5, 0xd8, 0x23, 0x27, 0x84, 0xda,
	0x1f, 0x41, 0xb1, 0x2b, 0xd4, 0x54, 0x7b, 0x9a, 0x79, 0x6f, 0xde, 0x8c, 0x9f, 0xc6, 0x83, 0x5f,
	0xb5, 0x42, 0x7e, 0x13, 0x35, 0x34, 0x42, 0x8b, 0x1a, 0x2c, 0x3f, 0x42, 0xd6, 0x5a, 0xe3, 0x0d,
	0x79, 0xb9, 0x16, 0xbd, 0xd0, 0xe0, 0xd9, 0x10, 0xd9, 0x58, 0xba, 0x98, 0xd7, 0xa6, 0x36, 0x41,
	0xc7, 0x87, 0x2c, 0xb6, 0x2c, 0x32, 0x69, 0x5c, 0x63, 0x1c, 0xaf, 0x84, 0x03, 0xde, 0xdf, 0x56,
	0xe0, 0xc5, 0x2d, 0x97, 0x46, 0xe9, 0x58, 0xbf, 0xf9, 0x79, 0x81, 0xaf, 0x96, 0x71, 0x10, 0x99,
	0xe3, 0x89, 0xd2, 0x0f, 0xb0, 0xa1, 0x28, 0x47, 0xc5, 0x75, 0x19, 0x01, 0x59, 0xe0, 0xe9, 0x43,
	0x67, 0x85, 0x57, 0x46, 0xd3, 0x27, 0x39, 0x2a, 0xd2, 0xf2, 0x3f, 0x1e, 0x3a, 0xa0, 0x35, 0x72,
	0x45, 0x2f, 0x42, 0x21, 0x02, 0xf2, 0x0e, 0x4f, 0x5a, 0xab, 0x24, 0xd0, 0x34, 0x47, 0xc5, 0xec,
	0xee, 0x05, 0x8b, 0x1e, 0xd8, 0xe0, 0x81, 0x1d, 0x3d, 0xb0, 0xf7, 0x46, 0xe9, 0xfb, 0x74, 0xfb,
	0xe7, 0x75, 0x52, 0x46, 0x35, 0xb9, 0xc1, 0x4f, 0xa5, 0x69, 0xda, 0xce, 0xc3, 0x17, 0xad, 0xbc,
	0xa3, 0x93, 0x30, 0x73, 0xc4, 0x91, 0x3b, 0x3c, 0x3f, 0xc5, 0x4b, 0xb0, 0x1f, 0xc2, 0xfb, 0x97,
	0x41, 0xfb, 0x68, 0x8d, 0x14, 0xf8, 0xb9, 0x03, 0xdb, 0x2b, 0x09, 0xd6, 0x7d, 0x36, 0x4b, 0xa1,
	0x2c, 0xbd, 0x0a, 0xf2, 0x73, 0x7a, 0x70, 0x20, 0xd6, 0x6b, 0xf3, 0xfd, 0x53, 0x0f, 0xb6, 0x73,
	0x40, 0xa7, 0x39, 0x2a, 0xa6, 0xe5, 0x88, 0x23, 0x39, 0x9e, 0x99, 0x98, 0x96, 0xc2, 0x03, 0xbd,
	0x0e, 0x93, 0x4e, 0x29, 0xf2, 0x06, 0x3f, 0x73, 0x5d, 0xe5, 0xa4, 0x55, 0xed, 0xb0, 0x24, 0x47,
	0x71, 0xd0, 0x8c, 0x49, 0x42, 0x70, 0xaa, 0x45, 0x03, 0x74, 0x16, 0x76, 0x1d, 0xf2, 0xfb, 0x8f,
	0xdb, 0x7d, 0x86, 0x76, 0xfb, 0x0c, 0xfd, 0xdd, 0x67, 0xe8, 0xd7, 0x21, 0x4b, 0x76, 0x87, 0x2c,
	0xf9, 0x7d, 0xc8, 0x92, 0xaf, 0xac, 0x56, 0x7e, 0xd5, 0x55, 0x4c, 0x9a, 0x86, 0x1f, 0x8f, 0x20,
	0x44, 0xbe, 0xe1, 0x67, 0x17, 0xe3, 0x7f, 0xb4, 0xe0, 0xaa, 0xcb, 0xf0, 0xbb, 0x6f, 0xff, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x9f, 0x48, 0x2c, 0x15, 0x50, 0x02, 0x00, 0x00,
}

func (m *Package) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Package) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Package) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPackage(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x5a
	}
	if m.Subscriptions != 0 {
		i = encodeVarintPackage(dAtA, i, uint64(m.Subscriptions))
		i--
		dAtA[i] = 0x50
	}
	if m.OveruseRate != 0 {
		i = encodeVarintPackage(dAtA, i, uint64(m.OveruseRate))
		i--
		dAtA[i] = 0x48
	}
	if m.AllowOveruse {
		i--
		if m.AllowOveruse {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.ServicersToPair != 0 {
		i = encodeVarintPackage(dAtA, i, uint64(m.ServicersToPair))
		i--
		dAtA[i] = 0x38
	}
	if m.ComputeUnitsPerEpoch != 0 {
		i = encodeVarintPackage(dAtA, i, uint64(m.ComputeUnitsPerEpoch))
		i--
		dAtA[i] = 0x30
	}
	if m.ComputeUnits != 0 {
		i = encodeVarintPackage(dAtA, i, uint64(m.ComputeUnits))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPackage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Epoch != 0 {
		i = encodeVarintPackage(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x18
	}
	if m.Duration != 0 {
		i = encodeVarintPackage(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintPackage(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPackage(dAtA []byte, offset int, v uint64) int {
	offset -= sovPackage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Package) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovPackage(uint64(l))
	}
	if m.Duration != 0 {
		n += 1 + sovPackage(uint64(m.Duration))
	}
	if m.Epoch != 0 {
		n += 1 + sovPackage(uint64(m.Epoch))
	}
	l = m.Price.Size()
	n += 1 + l + sovPackage(uint64(l))
	if m.ComputeUnits != 0 {
		n += 1 + sovPackage(uint64(m.ComputeUnits))
	}
	if m.ComputeUnitsPerEpoch != 0 {
		n += 1 + sovPackage(uint64(m.ComputeUnitsPerEpoch))
	}
	if m.ServicersToPair != 0 {
		n += 1 + sovPackage(uint64(m.ServicersToPair))
	}
	if m.AllowOveruse {
		n += 2
	}
	if m.OveruseRate != 0 {
		n += 1 + sovPackage(uint64(m.OveruseRate))
	}
	if m.Subscriptions != 0 {
		n += 1 + sovPackage(uint64(m.Subscriptions))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPackage(uint64(l))
	}
	return n
}

func sovPackage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPackage(x uint64) (n int) {
	return sovPackage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Package) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPackage
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
			return fmt.Errorf("proto: Package: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Package: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackage
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
				return ErrInvalidLengthPackage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPackage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackage
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
				return ErrInvalidLengthPackage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPackage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComputeUnits", wireType)
			}
			m.ComputeUnits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ComputeUnits |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComputeUnitsPerEpoch", wireType)
			}
			m.ComputeUnitsPerEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ComputeUnitsPerEpoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServicersToPair", wireType)
			}
			m.ServicersToPair = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServicersToPair |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowOveruse", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AllowOveruse = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OveruseRate", wireType)
			}
			m.OveruseRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OveruseRate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscriptions", wireType)
			}
			m.Subscriptions = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Subscriptions |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackage
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
				return ErrInvalidLengthPackage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPackage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPackage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPackage
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
func skipPackage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPackage
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
					return 0, ErrIntOverflowPackage
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
					return 0, ErrIntOverflowPackage
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
				return 0, ErrInvalidLengthPackage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPackage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPackage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPackage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPackage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPackage = fmt.Errorf("proto: unexpected end of group")
)
