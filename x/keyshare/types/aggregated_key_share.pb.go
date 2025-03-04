// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fairyring/keyshare/aggregated_key_share.proto

package types

import (
	fmt "fmt"
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

type AggregatedKeyShare struct {
	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Data   string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *AggregatedKeyShare) Reset()         { *m = AggregatedKeyShare{} }
func (m *AggregatedKeyShare) String() string { return proto.CompactTextString(m) }
func (*AggregatedKeyShare) ProtoMessage()    {}
func (*AggregatedKeyShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b3cffdc2704870f, []int{0}
}
func (m *AggregatedKeyShare) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AggregatedKeyShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AggregatedKeyShare.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AggregatedKeyShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregatedKeyShare.Merge(m, src)
}
func (m *AggregatedKeyShare) XXX_Size() int {
	return m.Size()
}
func (m *AggregatedKeyShare) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregatedKeyShare.DiscardUnknown(m)
}

var xxx_messageInfo_AggregatedKeyShare proto.InternalMessageInfo

func (m *AggregatedKeyShare) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *AggregatedKeyShare) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*AggregatedKeyShare)(nil), "fairyring.keyshare.AggregatedKeyShare")
}

func init() {
	proto.RegisterFile("fairyring/keyshare/aggregated_key_share.proto", fileDescriptor_9b3cffdc2704870f)
}

var fileDescriptor_9b3cffdc2704870f = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4d, 0x4b, 0xcc, 0x2c,
	0xaa, 0x2c, 0xca, 0xcc, 0x4b, 0xd7, 0xcf, 0x4e, 0xad, 0x2c, 0xce, 0x48, 0x2c, 0x4a, 0xd5, 0x4f,
	0x4c, 0x4f, 0x2f, 0x4a, 0x4d, 0x4f, 0x2c, 0x49, 0x4d, 0x89, 0xcf, 0x4e, 0xad, 0x8c, 0x07, 0x0b,
	0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x09, 0xc1, 0x95, 0xeb, 0xc1, 0x94, 0x2b, 0x39, 0x70,
	0x09, 0x39, 0xc2, 0x75, 0x78, 0xa7, 0x56, 0x06, 0x83, 0x44, 0x85, 0xc4, 0xb8, 0xd8, 0x32, 0x52,
	0x33, 0xd3, 0x33, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x82, 0xa0, 0x3c, 0x21, 0x21, 0x2e,
	0x96, 0x94, 0xc4, 0x92, 0x44, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0xdb, 0xc9, 0xe4,
	0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e,
	0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xa4, 0x10, 0xce, 0xab, 0x40, 0x38,
	0xb0, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x24, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x18, 0xe9, 0x89, 0x36, 0xc3, 0x00, 0x00, 0x00,
}

func (m *AggregatedKeyShare) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AggregatedKeyShare) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AggregatedKeyShare) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintAggregatedKeyShare(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if m.Height != 0 {
		i = encodeVarintAggregatedKeyShare(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAggregatedKeyShare(dAtA []byte, offset int, v uint64) int {
	offset -= sovAggregatedKeyShare(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AggregatedKeyShare) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovAggregatedKeyShare(uint64(m.Height))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovAggregatedKeyShare(uint64(l))
	}
	return n
}

func sovAggregatedKeyShare(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAggregatedKeyShare(x uint64) (n int) {
	return sovAggregatedKeyShare(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AggregatedKeyShare) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAggregatedKeyShare
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
			return fmt.Errorf("proto: AggregatedKeyShare: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AggregatedKeyShare: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregatedKeyShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregatedKeyShare
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
				return ErrInvalidLengthAggregatedKeyShare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAggregatedKeyShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAggregatedKeyShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAggregatedKeyShare
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
func skipAggregatedKeyShare(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAggregatedKeyShare
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
					return 0, ErrIntOverflowAggregatedKeyShare
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
					return 0, ErrIntOverflowAggregatedKeyShare
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
				return 0, ErrInvalidLengthAggregatedKeyShare
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAggregatedKeyShare
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAggregatedKeyShare
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAggregatedKeyShare        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAggregatedKeyShare          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAggregatedKeyShare = fmt.Errorf("proto: unexpected end of group")
)
