


package types

import (
	fmt "fmt"
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

type AuthorizationType int32

const (

	AuthorizationType_AUTHORIZATION_TYPE_UNSPECIFIED AuthorizationType = 0

	AuthorizationType_AUTHORIZATION_TYPE_DELEGATE AuthorizationType = 1

	AuthorizationType_AUTHORIZATION_TYPE_UNDELEGATE AuthorizationType = 2

	AuthorizationType_AUTHORIZATION_TYPE_REDELEGATE AuthorizationType = 3
)

var AuthorizationType_name = map[int32]string{
	0: "AUTHORIZATION_TYPE_UNSPECIFIED",
	1: "AUTHORIZATION_TYPE_DELEGATE",
	2: "AUTHORIZATION_TYPE_UNDELEGATE",
	3: "AUTHORIZATION_TYPE_REDELEGATE",
}

var AuthorizationType_value = map[string]int32{
	"AUTHORIZATION_TYPE_UNSPECIFIED": 0,
	"AUTHORIZATION_TYPE_DELEGATE":    1,
	"AUTHORIZATION_TYPE_UNDELEGATE":  2,
	"AUTHORIZATION_TYPE_REDELEGATE":  3,
}

func (x AuthorizationType) String() string {
	return proto.EnumName(AuthorizationType_name, int32(x))
}

func (AuthorizationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d6d8cdbc6f4432f0, []int{0}
}


//

type StakeAuthorization struct {


	MaxTokens *types.Coin `protobuf:"bytes,1,opt,name=max_tokens,json=maxTokens,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"max_tokens,omitempty"`

	//



	Validators isStakeAuthorization_Validators `protobuf_oneof:"validators"`

	AuthorizationType AuthorizationType `protobuf:"varint,4,opt,name=authorization_type,json=authorizationType,proto3,enum=cosmos.staking.v1beta1.AuthorizationType" json:"authorization_type,omitempty"`
}

func (m *StakeAuthorization) Reset()         { *m = StakeAuthorization{} }
func (m *StakeAuthorization) String() string { return proto.CompactTextString(m) }
func (*StakeAuthorization) ProtoMessage()    {}
func (*StakeAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6d8cdbc6f4432f0, []int{0}
}
func (m *StakeAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StakeAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StakeAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StakeAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakeAuthorization.Merge(m, src)
}
func (m *StakeAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *StakeAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_StakeAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_StakeAuthorization proto.InternalMessageInfo

type isStakeAuthorization_Validators interface {
	isStakeAuthorization_Validators()
	MarshalTo([]byte) (int, error)
	Size() int
}

type StakeAuthorization_AllowList struct {
	AllowList *StakeAuthorization_Validators `protobuf:"bytes,2,opt,name=allow_list,json=allowList,proto3,oneof" json:"allow_list,omitempty"`
}
type StakeAuthorization_DenyList struct {
	DenyList *StakeAuthorization_Validators `protobuf:"bytes,3,opt,name=deny_list,json=denyList,proto3,oneof" json:"deny_list,omitempty"`
}

func (*StakeAuthorization_AllowList) isStakeAuthorization_Validators() {}
func (*StakeAuthorization_DenyList) isStakeAuthorization_Validators()  {}

func (m *StakeAuthorization) GetValidators() isStakeAuthorization_Validators {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *StakeAuthorization) GetMaxTokens() *types.Coin {
	if m != nil {
		return m.MaxTokens
	}
	return nil
}

func (m *StakeAuthorization) GetAllowList() *StakeAuthorization_Validators {
	if x, ok := m.GetValidators().(*StakeAuthorization_AllowList); ok {
		return x.AllowList
	}
	return nil
}

func (m *StakeAuthorization) GetDenyList() *StakeAuthorization_Validators {
	if x, ok := m.GetValidators().(*StakeAuthorization_DenyList); ok {
		return x.DenyList
	}
	return nil
}

func (m *StakeAuthorization) GetAuthorizationType() AuthorizationType {
	if m != nil {
		return m.AuthorizationType
	}
	return AuthorizationType_AUTHORIZATION_TYPE_UNSPECIFIED
}


func (*StakeAuthorization) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StakeAuthorization_AllowList)(nil),
		(*StakeAuthorization_DenyList)(nil),
	}
}


type StakeAuthorization_Validators struct {
	Address []string `protobuf:"bytes,1,rep,name=address,proto3" json:"address,omitempty"`
}

func (m *StakeAuthorization_Validators) Reset()         { *m = StakeAuthorization_Validators{} }
func (m *StakeAuthorization_Validators) String() string { return proto.CompactTextString(m) }
func (*StakeAuthorization_Validators) ProtoMessage()    {}
func (*StakeAuthorization_Validators) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6d8cdbc6f4432f0, []int{0, 0}
}
func (m *StakeAuthorization_Validators) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StakeAuthorization_Validators) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StakeAuthorization_Validators.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StakeAuthorization_Validators) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakeAuthorization_Validators.Merge(m, src)
}
func (m *StakeAuthorization_Validators) XXX_Size() int {
	return m.Size()
}
func (m *StakeAuthorization_Validators) XXX_DiscardUnknown() {
	xxx_messageInfo_StakeAuthorization_Validators.DiscardUnknown(m)
}

var xxx_messageInfo_StakeAuthorization_Validators proto.InternalMessageInfo

func (m *StakeAuthorization_Validators) GetAddress() []string {
	if m != nil {
		return m.Address
	}
	return nil
}

func init() {
	proto.RegisterEnum("cosmos.staking.v1beta1.AuthorizationType", AuthorizationType_name, AuthorizationType_value)
	proto.RegisterType((*StakeAuthorization)(nil), "cosmos.staking.v1beta1.StakeAuthorization")
	proto.RegisterType((*StakeAuthorization_Validators)(nil), "cosmos.staking.v1beta1.StakeAuthorization.Validators")
}

func init() {
	proto.RegisterFile("cosmos/staking/v1beta1/authz.proto", fileDescriptor_d6d8cdbc6f4432f0)
}

var fileDescriptor_d6d8cdbc6f4432f0 = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x3d, 0x0d, 0x02, 0x72, 0xb8, 0xa8, 0x19, 0x21, 0x94, 0x06, 0x31, 0x2d, 0x59, 0x40,
	0xb8, 0x74, 0xac, 0x16, 0xb1, 0x61, 0x97, 0xb4, 0x2e, 0x8d, 0x54, 0xb5, 0x95, 0xeb, 0x56, 0xd0,
	0x8d, 0x35, 0x89, 0x47, 0xc9, 0x28, 0xb6, 0x27, 0xf2, 0x4c, 0x4a, 0xd2, 0xa7, 0xe0, 0x09, 0x78,
	0x00, 0xd6, 0x3c, 0x04, 0x0b, 0x16, 0x15, 0x2b, 0x76, 0xa0, 0xe4, 0x45, 0x90, 0xc7, 0xae, 0xa1,
	0x34, 0x65, 0xc3, 0xca, 0x97, 0xf3, 0xcd, 0xf7, 0x9f, 0x63, 0x1d, 0x43, 0xbd, 0x2b, 0x55, 0x24,
	0x95, 0xad, 0x34, 0x1b, 0x88, 0xb8, 0x67, 0x9f, 0xac, 0x75, 0xb8, 0x66, 0x6b, 0x36, 0x1b, 0xe9,
	0xfe, 0x29, 0x1d, 0x26, 0x52, 0x4b, 0x7c, 0x3f, 0x63, 0x68, 0xce, 0xd0, 0x9c, 0xa9, 0xdd, 0xeb,
	0xc9, 0x9e, 0x34, 0x88, 0x9d, 0xde, 0x65, 0x74, 0x6d, 0x29, 0xa3, 0xfd, 0xac, 0x90, 0x1f, 0xcd,
	0x4a, 0x24, 0x0f, 0xeb, 0x30, 0xc5, 0x8b, 0xa4, 0xae, 0x14, 0x71, 0x56, 0xaf, 0x7f, 0x2d, 0x01,
	0x3e, 0xd0, 0x6c, 0xc0, 0x9b, 0x23, 0xdd, 0x97, 0x89, 0x38, 0x65, 0x5a, 0xc8, 0x18, 0x73, 0x80,
	0x88, 0x8d, 0x7d, 0x2d, 0x07, 0x3c, 0x56, 0x55, 0xb4, 0x82, 0x1a, 0xb7, 0xd6, 0x97, 0x68, 0x6e,
	0x4e, 0x5d, 0xe7, 0x1d, 0xd1, 0x0d, 0x29, 0xe2, 0xd6, 0xf3, 0x4f, 0x3f, 0x96, 0x9f, 0xf4, 0x84,
	0xee, 0x8f, 0x3a, 0xb4, 0x2b, 0xa3, 0xbc, 0x85, 0xfc, 0xb2, 0xaa, 0x82, 0x81, 0xad, 0x27, 0x43,
	0xae, 0x0c, 0xec, 0x96, 0x23, 0x36, 0xf6, 0x8c, 0x18, 0x1f, 0x01, 0xb0, 0x30, 0x94, 0xef, 0xfd,
	0x50, 0x28, 0x5d, 0x5d, 0x30, 0x31, 0xaf, 0xe8, 0xfc, 0xd9, 0xe9, 0xe5, 0x36, 0xe9, 0x11, 0x0b,
	0x45, 0xc0, 0xb4, 0x4c, 0xd4, 0xb6, 0xe5, 0x96, 0x8d, 0x6a, 0x47, 0x28, 0x8d, 0x3d, 0x28, 0x07,
	0x3c, 0x9e, 0x64, 0xda, 0xd2, 0xff, 0x69, 0x6f, 0xa6, 0x26, 0x63, 0x7d, 0x0b, 0x98, 0xfd, 0xc9,
	0xf9, 0xe9, 0x50, 0xd5, 0x6b, 0x2b, 0xa8, 0x71, 0x77, 0xfd, 0xe9, 0x55, 0xfa, 0x0b, 0x66, 0x6f,
	0x32, 0xe4, 0x6e, 0x85, 0xfd, 0xfd, 0xaa, 0xf6, 0x18, 0xe0, 0x77, 0x26, 0xae, 0xc2, 0x0d, 0x16,
	0x04, 0x09, 0x57, 0xe9, 0x97, 0x2f, 0x35, 0xca, 0xee, 0xf9, 0xe3, 0xeb, 0xca, 0xb7, 0xcf, 0xab,
	0x77, 0x2e, 0x18, 0x5b, 0xb7, 0x01, 0x4e, 0x8a, 0xa3, 0xcf, 0x3e, 0x22, 0xa8, 0x5c, 0x4a, 0xc4,
	0x75, 0x20, 0xcd, 0x43, 0x6f, 0x7b, 0xcf, 0x6d, 0x1f, 0x37, 0xbd, 0xf6, 0xde, 0xae, 0xef, 0xbd,
	0xdb, 0x77, 0xfc, 0xc3, 0xdd, 0x83, 0x7d, 0x67, 0xa3, 0xbd, 0xd5, 0x76, 0x36, 0x17, 0x2d, 0xbc,
	0x0c, 0x0f, 0xe6, 0x30, 0x9b, 0xce, 0x8e, 0xf3, 0xa6, 0xe9, 0x39, 0x8b, 0x08, 0x3f, 0x82, 0x87,
	0x73, 0x25, 0x05, 0xb2, 0x70, 0x05, 0xe2, 0x3a, 0x05, 0x52, 0x6a, 0x6d, 0x7d, 0x99, 0x12, 0x74,
	0x36, 0x25, 0xe8, 0xe7, 0x94, 0xa0, 0x0f, 0x33, 0x62, 0x9d, 0xcd, 0x88, 0xf5, 0x7d, 0x46, 0xac,
	0xe3, 0x17, 0xff, 0xdc, 0x9f, 0x71, 0xf1, 0xbb, 0x98, 0x4d, 0xea, 0x5c, 0x37, 0xeb, 0xfb, 0xf2,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x04, 0xbf, 0x8c, 0x4d, 0x03, 0x00, 0x00,
}

func (m *StakeAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StakeAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakeAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AuthorizationType != 0 {
		i = encodeVarintAuthz(dAtA, i, uint64(m.AuthorizationType))
		i--
		dAtA[i] = 0x20
	}
	if m.Validators != nil {
		{
			size := m.Validators.Size()
			i -= size
			if _, err := m.Validators.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.MaxTokens != nil {
		{
			size, err := m.MaxTokens.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuthz(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StakeAuthorization_AllowList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakeAuthorization_AllowList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.AllowList != nil {
		{
			size, err := m.AllowList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuthz(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *StakeAuthorization_DenyList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakeAuthorization_DenyList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DenyList != nil {
		{
			size, err := m.DenyList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuthz(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *StakeAuthorization_Validators) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StakeAuthorization_Validators) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakeAuthorization_Validators) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		for iNdEx := len(m.Address) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Address[iNdEx])
			copy(dAtA[i:], m.Address[iNdEx])
			i = encodeVarintAuthz(dAtA, i, uint64(len(m.Address[iNdEx])))
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
func (m *StakeAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxTokens != nil {
		l = m.MaxTokens.Size()
		n += 1 + l + sovAuthz(uint64(l))
	}
	if m.Validators != nil {
		n += m.Validators.Size()
	}
	if m.AuthorizationType != 0 {
		n += 1 + sovAuthz(uint64(m.AuthorizationType))
	}
	return n
}

func (m *StakeAuthorization_AllowList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AllowList != nil {
		l = m.AllowList.Size()
		n += 1 + l + sovAuthz(uint64(l))
	}
	return n
}
func (m *StakeAuthorization_DenyList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DenyList != nil {
		l = m.DenyList.Size()
		n += 1 + l + sovAuthz(uint64(l))
	}
	return n
}
func (m *StakeAuthorization_Validators) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Address) > 0 {
		for _, s := range m.Address {
			l = len(s)
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
func (m *StakeAuthorization) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: StakeAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StakeAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTokens", wireType)
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
			if m.MaxTokens == nil {
				m.MaxTokens = &types.Coin{}
			}
			if err := m.MaxTokens.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowList", wireType)
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
			v := &StakeAuthorization_Validators{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Validators = &StakeAuthorization_AllowList{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenyList", wireType)
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
			v := &StakeAuthorization_Validators{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Validators = &StakeAuthorization_DenyList{v}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizationType", wireType)
			}
			m.AuthorizationType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuthorizationType |= AuthorizationType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *StakeAuthorization_Validators) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Validators: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validators: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address, string(dAtA[iNdEx:postIndex]))
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