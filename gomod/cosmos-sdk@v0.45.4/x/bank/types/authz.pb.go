


package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	io "io"
	math "math"
	math_bits "math/bits"
)


var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf





const _ = proto.GoGoProtoPackageIsVersion3



//

type SendAuthorization struct {
	SpendLimit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=spend_limit,json=spendLimit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"spend_limit"`
}

func (m *SendAuthorization) Reset()         { *m = SendAuthorization{} }
func (m *SendAuthorization) String() string { return proto.CompactTextString(m) }
func (*SendAuthorization) ProtoMessage()    {}
func (*SendAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d2a37888ea779f, []int{0}
}
func (m *SendAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendAuthorization.Merge(m, src)
}
func (m *SendAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *SendAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_SendAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_SendAuthorization proto.InternalMessageInfo

func (m *SendAuthorization) GetSpendLimit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.SpendLimit
	}
	return nil
}

func init() {
	proto.RegisterType((*SendAuthorization)(nil), "cosmos.bank.v1beta1.SendAuthorization")
}

func init() { proto.RegisterFile("cosmos/bank/v1beta1/authz.proto", fileDescriptor_a4d2a37888ea779f) }

var fileDescriptor_a4d2a37888ea779f = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x4f, 0x4a, 0xcc, 0xcb, 0xd6, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4,
	0x4f, 0x2c, 0x2d, 0xc9, 0xa8, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x86, 0x28, 0xd0,
	0x03, 0x29, 0xd0, 0x83, 0x2a, 0x90, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xcb, 0xeb, 0x83, 0x58,
	0x10, 0xa5, 0x52, 0x92, 0x10, 0xa5, 0xf1, 0x10, 0x09, 0xa8, 0x3e, 0x88, 0x94, 0x1c, 0xdc, 0x9a,
	0xe2, 0x54, 0xb8, 0x35, 0xc9, 0xf9, 0x99, 0x79, 0x10, 0x79, 0xa5, 0x29, 0x8c, 0x5c, 0x82, 0xc1,
	0xa9, 0x79, 0x29, 0x8e, 0xa5, 0x25, 0x19, 0xf9, 0x45, 0x99, 0x55, 0x89, 0x25, 0x99, 0xf9, 0x79,
	0x42, 0x39, 0x5c, 0xdc, 0xc5, 0x05, 0xa9, 0x79, 0x29, 0xf1, 0x39, 0x99, 0xb9, 0x99, 0x25, 0x12,
	0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x92, 0x7a, 0x70, 0x17, 0x15, 0xa7, 0xc2, 0x5c, 0xa4, 0xe7,
	0x9c, 0x9f, 0x99, 0xe7, 0x64, 0x70, 0xe2, 0x9e, 0x3c, 0xc3, 0xaa, 0xfb, 0xf2, 0x1a, 0xe9, 0x99,
	0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0x50, 0x67, 0x40, 0x29, 0xdd, 0xe2, 0x94, 0x6c,
	0xfd, 0x92, 0xca, 0x82, 0xd4, 0x62, 0xb0, 0x86, 0xe2, 0x20, 0x2e, 0xb0, 0xf9, 0x3e, 0x20, 0xe3,
	0xad, 0x04, 0x2f, 0x6d, 0xd1, 0xe5, 0x45, 0x71, 0x80, 0x93, 0xf3, 0x89, 0x47, 0x72, 0x8c, 0x17,
	0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c,
	0x37, 0x1e, 0xcb, 0x31, 0x44, 0x69, 0xe2, 0xb5, 0xa2, 0x02, 0x12, 0x9e, 0x60, 0x9b, 0x92, 0xd8,
	0xc0, 0x5e, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xce, 0xe1, 0x4c, 0x6b, 0x01, 0x00,
	0x00,
}

func (m *SendAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SendAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SpendLimit) > 0 {
		for iNdEx := len(m.SpendLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SpendLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAuthz(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuthz(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuthz(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SendAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SpendLimit) > 0 {
		for _, e := range m.SpendLimit {
			l = e.Size()
			n += 1 + l + sovAuthz(uint64(l))
		}
	}
	return n
}

func sovAuthz(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuthz(x uint64) (n int) {
	return sovAuthz(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SendAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: SendAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpendLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpendLimit = append(m.SpendLimit, types.Coin{})
			if err := m.SpendLimit[len(m.SpendLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func skipAuthz(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
				return 0, ErrInvalidLengthAuthz
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuthz
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuthz
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuthz        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthz          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuthz = fmt.Errorf("proto: unexpected end of group")
)
