


package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)


var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf





const _ = proto.GoGoProtoPackageIsVersion3


type Params struct {

	IndexNumHeight int64 `protobuf:"varint,1,opt,name=index_num_height,json=indexNumHeight,proto3" json:"index_num_height,omitempty"`

	RedeemFeeHeight int64 `protobuf:"varint,2,opt,name=redeem_fee_height,json=redeemFeeHeight,proto3" json:"redeem_fee_height,omitempty"`

	RedeemFee github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=redeem_fee,json=redeemFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"redeem_fee"`

	MinDelegate github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=min_delegate,json=minDelegate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_delegate"`

	Validity int64 `protobuf:"varint,5,opt,name=validity,proto3" json:"validity,omitempty"`

	BonusCycle int64 `protobuf:"varint,6,opt,name=bonus_cycle,json=bonusCycle,proto3" json:"bonus_cycle,omitempty"`

	BonusHalve int64 `protobuf:"varint,7,opt,name=bonus_halve,json=bonusHalve,proto3" json:"bonus_halve,omitempty"`

	Bonus                github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=bonus,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"bonus"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{0}
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

func (m *Params) GetIndexNumHeight() int64 {
	if m != nil {
		return m.IndexNumHeight
	}
	return 0
}

func (m *Params) GetRedeemFeeHeight() int64 {
	if m != nil {
		return m.RedeemFeeHeight
	}
	return 0
}

func (m *Params) GetValidity() int64 {
	if m != nil {
		return m.Validity
	}
	return 0
}

func (m *Params) GetBonusCycle() int64 {
	if m != nil {
		return m.BonusCycle
	}
	return 0
}

func (m *Params) GetBonusHalve() int64 {
	if m != nil {
		return m.BonusHalve
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "freemasonry.comm.v1.Params")
}

func init() { proto.RegisterFile("gateway.proto", fileDescriptor_f1a937782ebbded5) }

var fileDescriptor_f1a937782ebbded5 = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x4f, 0x4e, 0xc2, 0x40,
	0x14, 0xc6, 0xad, 0x08, 0xc2, 0xe0, 0x1f, 0xac, 0x2e, 0x1a, 0x36, 0x25, 0x2e, 0x48, 0x63, 0x62,
	0x1b, 0xe3, 0xce, 0x25, 0x12, 0x83, 0x0b, 0x8d, 0xb2, 0x74, 0xd3, 0x4c, 0xa7, 0x8f, 0x76, 0x42,
	0x67, 0x86, 0x74, 0xa6, 0x48, 0x6f, 0xe0, 0xd2, 0x23, 0xc0, 0x25, 0x3c, 0x83, 0x67, 0x70, 0xc1,
	0xda, 0x63, 0x98, 0x4e, 0x11, 0x74, 0xa9, 0xab, 0xf6, 0x7d, 0xdf, 0x6f, 0xbe, 0xbc, 0x97, 0x0f,
	0xed, 0x47, 0x58, 0xc1, 0x33, 0xce, 0xdd, 0x49, 0x2a, 0x94, 0x30, 0x8f, 0x47, 0x29, 0x00, 0xc3,
	0x52, 0xf0, 0x34, 0x77, 0x89, 0x60, 0xcc, 0x9d, 0x5e, 0xb4, 0x4f, 0x22, 0x11, 0x09, 0xed, 0x7b,
	0xc5, 0x5f, 0x89, 0x9e, 0xbe, 0x55, 0x50, 0xed, 0x01, 0xa7, 0x98, 0x49, 0xd3, 0x41, 0x2d, 0xca,
	0x43, 0x98, 0xf9, 0x3c, 0x63, 0x7e, 0x0c, 0x34, 0x8a, 0x95, 0x65, 0x74, 0x0c, 0xa7, 0x32, 0x3c,
	0xd0, 0xfa, 0x7d, 0xc6, 0x06, 0x5a, 0x35, 0xcf, 0xd0, 0x51, 0x0a, 0x21, 0x00, 0xf3, 0x47, 0x00,
	0xdf, 0xe8, 0xb6, 0x46, 0x0f, 0x4b, 0xe3, 0x06, 0x60, 0xc5, 0xde, 0x21, 0xb4, 0x61, 0xad, 0x4a,
	0xc7, 0x70, 0x1a, 0x3d, 0xf7, 0x7d, 0x69, 0x6f, 0x7d, 0x2c, 0xed, 0x6e, 0x44, 0x55, 0x9c, 0x05,
	0xc5, 0x8a, 0x1e, 0x11, 0x92, 0x09, 0xb9, 0xfa, 0x9c, 0xcb, 0x70, 0xec, 0xa9, 0x7c, 0x02, 0xd2,
	0xed, 0x03, 0x19, 0x36, 0xd6, 0xa1, 0xe6, 0x23, 0xda, 0x63, 0x94, 0xfb, 0x21, 0x24, 0x50, 0xdc,
	0x6c, 0xed, 0xfc, 0x39, 0xf0, 0x96, 0xab, 0x61, 0x93, 0x51, 0xde, 0x5f, 0x45, 0x98, 0x6d, 0x54,
	0x9f, 0xe2, 0x84, 0x86, 0x54, 0xe5, 0x56, 0x55, 0x1f, 0xb1, 0x9e, 0x4d, 0x1b, 0x35, 0x03, 0xc1,
	0x33, 0xe9, 0x93, 0x9c, 0x24, 0x60, 0xd5, 0xb4, 0x8d, 0xb4, 0x74, 0x5d, 0x28, 0x1b, 0x20, 0xc6,
	0xc9, 0x14, 0xac, 0xdd, 0x1f, 0xc0, 0xa0, 0x50, 0xcc, 0x3e, 0xaa, 0xea, 0xc9, 0xaa, 0xff, 0x6b,
	0xd3, 0xf2, 0xf1, 0x55, 0xeb, 0x73, 0x6e, 0x1b, 0x2f, 0x0b, 0xdb, 0x78, 0x5d, 0xd8, 0xc6, 0x7c,
	0x61, 0x1b, 0x3d, 0xe7, 0xa9, 0xfb, 0xab, 0x65, 0xe2, 0x05, 0x89, 0x20, 0x63, 0x12, 0x63, 0xca,
	0xbd, 0x99, 0x57, 0xb4, 0x5e, 0xc6, 0x04, 0x35, 0xdd, 0xf4, 0xe5, 0x57, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x9b, 0x77, 0xcc, 0x04, 0x25, 0x02, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.IndexNumHeight != that1.IndexNumHeight {
		return false
	}
	if this.RedeemFeeHeight != that1.RedeemFeeHeight {
		return false
	}
	if !this.RedeemFee.Equal(that1.RedeemFee) {
		return false
	}
	if !this.MinDelegate.Equal(that1.MinDelegate) {
		return false
	}
	if this.Validity != that1.Validity {
		return false
	}
	if this.BonusCycle != that1.BonusCycle {
		return false
	}
	if this.BonusHalve != that1.BonusHalve {
		return false
	}
	if !this.Bonus.Equal(that1.Bonus) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
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
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	{
		size := m.Bonus.Size()
		i -= size
		if _, err := m.Bonus.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGateway(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42

	if m.BonusHalve != 0 {
		i = encodeVarintGateway(dAtA, i, uint64(m.BonusHalve))
		i--
		dAtA[i] = 0x38
	}
	if m.BonusCycle != 0 {
		i = encodeVarintGateway(dAtA, i, uint64(m.BonusCycle))
		i--
		dAtA[i] = 0x30
	}
	if m.Validity != 0 {
		i = encodeVarintGateway(dAtA, i, uint64(m.Validity))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.MinDelegate.Size()
		i -= size
		if _, err := m.MinDelegate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGateway(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22

	{
		size := m.RedeemFee.Size()
		i -= size
		if _, err := m.RedeemFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGateway(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a

	if m.RedeemFeeHeight != 0 {
		i = encodeVarintGateway(dAtA, i, uint64(m.RedeemFeeHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.IndexNumHeight != 0 {
		i = encodeVarintGateway(dAtA, i, uint64(m.IndexNumHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGateway(dAtA []byte, offset int, v uint64) int {
	offset -= sovGateway(v)
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
	if m.IndexNumHeight != 0 {
		n += 1 + sovGateway(uint64(m.IndexNumHeight))
	}
	if m.RedeemFeeHeight != 0 {
		n += 1 + sovGateway(uint64(m.RedeemFeeHeight))
	}
	l = m.RedeemFee.Size()
	if l > 0 {
		n += 1 + l + sovGateway(uint64(l))
	}
	l = m.MinDelegate.Size()
	if l > 0 {
		n += 1 + l + sovGateway(uint64(l))
	}
	if m.Validity != 0 {
		n += 1 + sovGateway(uint64(m.Validity))
	}
	if m.BonusCycle != 0 {
		n += 1 + sovGateway(uint64(m.BonusCycle))
	}
	if m.BonusHalve != 0 {
		n += 1 + sovGateway(uint64(m.BonusHalve))
	}
	l = m.Bonus.Size()
	if l > 0 {
		n += 1 + l + sovGateway(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGateway(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGateway(x uint64) (n int) {
	return sovGateway(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGateway
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexNumHeight", wireType)
			}
			m.IndexNumHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateway
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IndexNumHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedeemFeeHeight", wireType)
			}
			m.RedeemFeeHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateway
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RedeemFeeHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedeemFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateway
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
				return ErrInvalidLengthGateway
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGateway
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RedeemFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDelegate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateway
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
				return ErrInvalidLengthGateway
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGateway
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinDelegate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validity", wireType)
			}
			m.Validity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateway
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Validity |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BonusCycle", wireType)
			}
			m.BonusCycle = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateway
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BonusCycle |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BonusHalve", wireType)
			}
			m.BonusHalve = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateway
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BonusHalve |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bonus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateway
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
				return ErrInvalidLengthGateway
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGateway
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Bonus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGateway(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGateway
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
func skipGateway(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGateway
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
					return 0, ErrIntOverflowGateway
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
					return 0, ErrIntOverflowGateway
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
				return 0, ErrInvalidLengthGateway
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGateway
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGateway
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGateway        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGateway          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGateway = fmt.Errorf("proto: unexpected end of group")
)
