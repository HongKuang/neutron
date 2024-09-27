// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: neutron/dex/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the parameters for the module.
type Params struct {
	FeeTiers              []uint64 `protobuf:"varint,1,rep,packed,name=fee_tiers,json=feeTiers,proto3" json:"fee_tiers,omitempty"`
	Paused                bool     `protobuf:"varint,3,opt,name=paused,proto3" json:"paused"`
	MaxJitsPerBlock       uint64   `protobuf:"varint,4,opt,name=max_jits_per_block,json=maxJitsPerBlock,proto3" json:"max_jits_per_block,omitempty"`
	GoodTilPurgeAllowance uint64   `protobuf:"varint,5,opt,name=good_til_purge_allowance,json=goodTilPurgeAllowance,proto3" json:"good_til_purge_allowance,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_84a6bffcfc21009c, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetFeeTiers() []uint64 {
	if m != nil {
		return m.FeeTiers
	}
	return nil
}

func (m *Params) GetPaused() bool {
	if m != nil {
		return m.Paused
	}
	return false
}

func (m *Params) GetMaxJitsPerBlock() uint64 {
	if m != nil {
		return m.MaxJitsPerBlock
	}
	return 0
}

func (m *Params) GetGoodTilPurgeAllowance() uint64 {
	if m != nil {
		return m.GoodTilPurgeAllowance
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "neutron.dex.Params")
}

func init() { proto.RegisterFile("neutron/dex/params.proto", fileDescriptor_84a6bffcfc21009c) }

var fileDescriptor_84a6bffcfc21009c = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0x31, 0x4b, 0x03, 0x31,
	0x14, 0xc7, 0x2f, 0xf4, 0x2c, 0x35, 0x0e, 0x42, 0x50, 0x08, 0x0a, 0x69, 0xe9, 0x54, 0x90, 0x36,
	0x83, 0x88, 0xe0, 0x66, 0x17, 0xc1, 0xa9, 0x94, 0x4e, 0x2e, 0x21, 0xed, 0xbd, 0xc6, 0xe8, 0x5d,
	0x13, 0x92, 0x9c, 0x9e, 0xdf, 0xc2, 0xd1, 0xd1, 0x2f, 0xe1, 0x77, 0x70, 0xec, 0xe8, 0x24, 0x72,
	0xb7, 0xf9, 0x29, 0xe4, 0xce, 0xeb, 0xf4, 0xfe, 0xef, 0xff, 0x7b, 0xff, 0xe1, 0xfd, 0x31, 0xdd,
	0x40, 0x1e, 0x9c, 0xd9, 0xf0, 0x04, 0x0a, 0x6e, 0xa5, 0x93, 0x99, 0x9f, 0x58, 0x67, 0x82, 0x21,
	0x07, 0x2d, 0x99, 0x24, 0x50, 0x9c, 0x1c, 0x29, 0xa3, 0x4c, 0xe3, 0xf3, 0x5a, 0xfd, 0x9f, 0x0c,
	0x3f, 0x10, 0xee, 0xce, 0x9a, 0x0c, 0x39, 0xc5, 0xfb, 0x6b, 0x00, 0x11, 0x34, 0x38, 0x4f, 0xd1,
	0xa0, 0x33, 0x8a, 0xe7, 0xbd, 0x35, 0xc0, 0xa2, 0xde, 0xc9, 0x10, 0x77, 0xad, 0xcc, 0x3d, 0x24,
	0xb4, 0x33, 0x40, 0xa3, 0xde, 0x14, 0xff, 0x7e, 0xf7, 0x5b, 0x67, 0xde, 0x4e, 0x72, 0x86, 0x49,
	0x26, 0x0b, 0xf1, 0xa0, 0x83, 0x17, 0x16, 0x9c, 0x58, 0xa6, 0x66, 0xf5, 0x48, 0xe3, 0x01, 0x1a,
	0xc5, 0xf3, 0xc3, 0x4c, 0x16, 0xb7, 0x3a, 0xf8, 0x19, 0xb8, 0x69, 0x6d, 0x93, 0x4b, 0x4c, 0x95,
	0x31, 0x89, 0x08, 0x3a, 0x15, 0x36, 0x77, 0x0a, 0x84, 0x4c, 0x53, 0xf3, 0x2c, 0x37, 0x2b, 0xa0,
	0x7b, 0x4d, 0xe4, 0xb8, 0xe6, 0x0b, 0x9d, 0xce, 0x6a, 0x7a, 0xbd, 0x83, 0x57, 0xf1, 0xdb, 0x7b,
	0x3f, 0x9a, 0xde, 0x7c, 0x96, 0x0c, 0x6d, 0x4b, 0x86, 0x7e, 0x4a, 0x86, 0x5e, 0x2b, 0x16, 0x6d,
	0x2b, 0x16, 0x7d, 0x55, 0x2c, 0xba, 0x1b, 0x2b, 0x1d, 0xee, 0xf3, 0xe5, 0x64, 0x65, 0x32, 0xde,
	0xfe, 0x3f, 0x36, 0x4e, 0xed, 0x34, 0x7f, 0xba, 0xe0, 0x45, 0x53, 0x55, 0x78, 0xb1, 0xe0, 0x97,
	0xdd, 0xa6, 0x87, 0xf3, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc6, 0xda, 0x12, 0xf5, 0x46, 0x01,
	0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GoodTilPurgeAllowance != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.GoodTilPurgeAllowance))
		i--
		dAtA[i] = 0x28
	}
	if m.MaxJitsPerBlock != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxJitsPerBlock))
		i--
		dAtA[i] = 0x20
	}
	if m.Paused {
		i--
		if m.Paused {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.FeeTiers) > 0 {
		dAtA2 := make([]byte, len(m.FeeTiers)*10)
		var j1 int
		for _, num := range m.FeeTiers {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintParams(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FeeTiers) > 0 {
		l = 0
		for _, e := range m.FeeTiers {
			l += sovParams(uint64(e))
		}
		n += 1 + sovParams(uint64(l)) + l
	}
	if m.Paused {
		n += 2
	}
	if m.MaxJitsPerBlock != 0 {
		n += 1 + sovParams(uint64(m.MaxJitsPerBlock))
	}
	if m.GoodTilPurgeAllowance != 0 {
		n += 1 + sovParams(uint64(m.GoodTilPurgeAllowance))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowParams
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.FeeTiers = append(m.FeeTiers, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowParams
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthParams
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthParams
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.FeeTiers) == 0 {
					m.FeeTiers = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowParams
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.FeeTiers = append(m.FeeTiers, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeTiers", wireType)
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paused", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.Paused = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxJitsPerBlock", wireType)
			}
			m.MaxJitsPerBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxJitsPerBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoodTilPurgeAllowance", wireType)
			}
			m.GoodTilPurgeAllowance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GoodTilPurgeAllowance |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
