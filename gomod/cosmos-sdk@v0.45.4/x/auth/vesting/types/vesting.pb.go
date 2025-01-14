


package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/x/auth/types"
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



type BaseVestingAccount struct {
	*types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty"`
	OriginalVesting    github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=original_vesting,json=originalVesting,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"original_vesting" yaml:"original_vesting"`
	DelegatedFree      github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=delegated_free,json=delegatedFree,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"delegated_free" yaml:"delegated_free"`
	DelegatedVesting   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=delegated_vesting,json=delegatedVesting,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"delegated_vesting" yaml:"delegated_vesting"`
	EndTime            int64                                    `protobuf:"varint,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty" yaml:"end_time"`
}

func (m *BaseVestingAccount) Reset()      { *m = BaseVestingAccount{} }
func (*BaseVestingAccount) ProtoMessage() {}
func (*BaseVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_89e80273ca606d6e, []int{0}
}
func (m *BaseVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseVestingAccount.Merge(m, src)
}
func (m *BaseVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *BaseVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_BaseVestingAccount proto.InternalMessageInfo



type ContinuousVestingAccount struct {
	*BaseVestingAccount `protobuf:"bytes,1,opt,name=base_vesting_account,json=baseVestingAccount,proto3,embedded=base_vesting_account" json:"base_vesting_account,omitempty"`
	StartTime           int64 `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty" yaml:"start_time"`
}

func (m *ContinuousVestingAccount) Reset()      { *m = ContinuousVestingAccount{} }
func (*ContinuousVestingAccount) ProtoMessage() {}
func (*ContinuousVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_89e80273ca606d6e, []int{1}
}
func (m *ContinuousVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContinuousVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContinuousVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContinuousVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContinuousVestingAccount.Merge(m, src)
}
func (m *ContinuousVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *ContinuousVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ContinuousVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ContinuousVestingAccount proto.InternalMessageInfo




type DelayedVestingAccount struct {
	*BaseVestingAccount `protobuf:"bytes,1,opt,name=base_vesting_account,json=baseVestingAccount,proto3,embedded=base_vesting_account" json:"base_vesting_account,omitempty"`
}

func (m *DelayedVestingAccount) Reset()      { *m = DelayedVestingAccount{} }
func (*DelayedVestingAccount) ProtoMessage() {}
func (*DelayedVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_89e80273ca606d6e, []int{2}
}
func (m *DelayedVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelayedVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelayedVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelayedVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelayedVestingAccount.Merge(m, src)
}
func (m *DelayedVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *DelayedVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_DelayedVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_DelayedVestingAccount proto.InternalMessageInfo


type Period struct {
	Length int64                                    `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	Amount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *Period) Reset()      { *m = Period{} }
func (*Period) ProtoMessage() {}
func (*Period) Descriptor() ([]byte, []int) {
	return fileDescriptor_89e80273ca606d6e, []int{3}
}
func (m *Period) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Period) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Period.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Period) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Period.Merge(m, src)
}
func (m *Period) XXX_Size() int {
	return m.Size()
}
func (m *Period) XXX_DiscardUnknown() {
	xxx_messageInfo_Period.DiscardUnknown(m)
}

var xxx_messageInfo_Period proto.InternalMessageInfo

func (m *Period) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Period) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}



type PeriodicVestingAccount struct {
	*BaseVestingAccount `protobuf:"bytes,1,opt,name=base_vesting_account,json=baseVestingAccount,proto3,embedded=base_vesting_account" json:"base_vesting_account,omitempty"`
	StartTime           int64    `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty" yaml:"start_time"`
	VestingPeriods      []Period `protobuf:"bytes,3,rep,name=vesting_periods,json=vestingPeriods,proto3" json:"vesting_periods" yaml:"vesting_periods"`
}

func (m *PeriodicVestingAccount) Reset()      { *m = PeriodicVestingAccount{} }
func (*PeriodicVestingAccount) ProtoMessage() {}
func (*PeriodicVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_89e80273ca606d6e, []int{4}
}
func (m *PeriodicVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PeriodicVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PeriodicVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PeriodicVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeriodicVestingAccount.Merge(m, src)
}
func (m *PeriodicVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *PeriodicVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_PeriodicVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_PeriodicVestingAccount proto.InternalMessageInfo




//

type PermanentLockedAccount struct {
	*BaseVestingAccount `protobuf:"bytes,1,opt,name=base_vesting_account,json=baseVestingAccount,proto3,embedded=base_vesting_account" json:"base_vesting_account,omitempty"`
}

func (m *PermanentLockedAccount) Reset()      { *m = PermanentLockedAccount{} }
func (*PermanentLockedAccount) ProtoMessage() {}
func (*PermanentLockedAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_89e80273ca606d6e, []int{5}
}
func (m *PermanentLockedAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PermanentLockedAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PermanentLockedAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PermanentLockedAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermanentLockedAccount.Merge(m, src)
}
func (m *PermanentLockedAccount) XXX_Size() int {
	return m.Size()
}
func (m *PermanentLockedAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_PermanentLockedAccount.DiscardUnknown(m)
}

var xxx_messageInfo_PermanentLockedAccount proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BaseVestingAccount)(nil), "cosmos.vesting.v1beta1.BaseVestingAccount")
	proto.RegisterType((*ContinuousVestingAccount)(nil), "cosmos.vesting.v1beta1.ContinuousVestingAccount")
	proto.RegisterType((*DelayedVestingAccount)(nil), "cosmos.vesting.v1beta1.DelayedVestingAccount")
	proto.RegisterType((*Period)(nil), "cosmos.vesting.v1beta1.Period")
	proto.RegisterType((*PeriodicVestingAccount)(nil), "cosmos.vesting.v1beta1.PeriodicVestingAccount")
	proto.RegisterType((*PermanentLockedAccount)(nil), "cosmos.vesting.v1beta1.PermanentLockedAccount")
}

func init() {
	proto.RegisterFile("cosmos/vesting/v1beta1/vesting.proto", fileDescriptor_89e80273ca606d6e)
}

var fileDescriptor_89e80273ca606d6e = []byte{
	
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x95, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x7d, 0x4d, 0x08, 0xe5, 0x0a, 0xfd, 0x63, 0xda, 0x60, 0x3a, 0xd8, 0x91, 0xc5, 0x10,
	0x21, 0xe1, 0xd0, 0xc2, 0xd4, 0x0d, 0x17, 0x21, 0x55, 0xed, 0x80, 0x2c, 0xc4, 0xc0, 0x12, 0x9d,
	0xed, 0xc3, 0xb1, 0x1a, 0xdf, 0x45, 0xbe, 0x4b, 0x45, 0x3e, 0x00, 0x08, 0xa9, 0x0b, 0x48, 0x0c,
	0x8c, 0x5d, 0x58, 0xf8, 0x10, 0xcc, 0x1d, 0x23, 0x26, 0xa6, 0x80, 0x92, 0x6f, 0x90, 0x4f, 0x80,
	0x7c, 0x77, 0x76, 0x8a, 0x0b, 0x44, 0x65, 0x00, 0x31, 0x25, 0x77, 0xef, 0xfb, 0x3e, 0xf7, 0x7b,
	0x5f, 0x3f, 0xa7, 0x83, 0xb7, 0x02, 0xca, 0x12, 0xca, 0x5a, 0x47, 0x98, 0xf1, 0x98, 0x44, 0xad,
	0xa3, 0x2d, 0x1f, 0x73, 0xb4, 0x95, 0xaf, 0x9d, 0x5e, 0x4a, 0x39, 0xd5, 0xeb, 0x32, 0xcb, 0xc9,
	0x77, 0x55, 0xd6, 0xe6, 0x7a, 0x44, 0x23, 0x2a, 0x52, 0x5a, 0xd9, 0x3f, 0x99, 0xbd, 0x69, 0x2a,
	0x4d, 0x1f, 0x31, 0x5c, 0x08, 0x06, 0x34, 0x26, 0xa5, 0x38, 0xea, 0xf3, 0x4e, 0x11, 0xcf, 0x16,
	0x32, 0x6e, 0x7f, 0xae, 0x42, 0xdd, 0x45, 0x0c, 0x3f, 0x95, 0xa7, 0x3d, 0x08, 0x02, 0xda, 0x27,
	0x5c, 0xdf, 0x83, 0x57, 0x33, 0xc5, 0x36, 0x92, 0x6b, 0x03, 0x34, 0x40, 0x73, 0x69, 0xbb, 0xe1,
	0x28, 0x36, 0x21, 0xa0, 0xd4, 0x9c, 0xac, 0x5c, 0xd5, 0xb9, 0xd5, 0xe1, 0xc8, 0x02, 0xde, 0x92,
	0x3f, 0xdb, 0xd2, 0xdf, 0x02, 0xb8, 0x4a, 0xd3, 0x38, 0x8a, 0x09, 0xea, 0xb6, 0x55, 0x53, 0xc6,
	0x42, 0xa3, 0xd2, 0x5c, 0xda, 0xbe, 0x99, 0xeb, 0x65, 0xf9, 0x85, 0xde, 0x2e, 0x8d, 0x89, 0xbb,
	0x7f, 0x3a, 0xb2, 0xb4, 0xe9, 0xc8, 0xba, 0x31, 0x40, 0x49, 0x77, 0xc7, 0x2e, 0x0b, 0xd8, 0x1f,
	0xbf, 0x5a, 0xcd, 0x28, 0xe6, 0x9d, 0xbe, 0xef, 0x04, 0x34, 0x69, 0xa9, 0x2e, 0xe5, 0xcf, 0x1d,
	0x16, 0x1e, 0xb6, 0xf8, 0xa0, 0x87, 0x99, 0xd0, 0x62, 0xde, 0x4a, 0x5e, 0xae, 0xba, 0xd4, 0x8f,
	0x01, 0x5c, 0x0e, 0x71, 0x17, 0x47, 0x88, 0xe3, 0xb0, 0xfd, 0x3c, 0xc5, 0xd8, 0xa8, 0xcc, 0x23,
	0xda, 0x53, 0x44, 0x1b, 0x92, 0xe8, 0xc7, 0xf2, 0x8b, 0xf1, 0x5c, 0x2b, 0x8a, 0x1f, 0xa5, 0x18,
	0xeb, 0xef, 0x00, 0x5c, 0x9b, 0xc9, 0xe5, 0x23, 0xaa, 0xce, 0x03, 0x3a, 0x50, 0x40, 0x46, 0x19,
	0xe8, 0x8f, 0x66, 0xb4, 0x5a, 0xd4, 0xe7, 0x43, 0x72, 0xe0, 0x22, 0x26, 0x61, 0x9b, 0xc7, 0x09,
	0x36, 0x2e, 0x35, 0x40, 0xb3, 0xe2, 0x5e, 0x9f, 0x8e, 0xac, 0x15, 0x79, 0x5a, 0x1e, 0xb1, 0xbd,
	0xcb, 0x98, 0x84, 0x4f, 0xe2, 0x04, 0xef, 0x2c, 0xbe, 0x3e, 0xb1, 0xb4, 0xf7, 0x27, 0x96, 0x66,
	0x7f, 0x02, 0xd0, 0xd8, 0xa5, 0x84, 0xc7, 0xa4, 0x4f, 0xfb, 0xac, 0x64, 0x2d, 0x1f, 0xae, 0x0b,
	0x6b, 0x29, 0xca, 0x92, 0xc5, 0x6e, 0x3b, 0x3f, 0xb7, 0xbf, 0x73, 0xde, 0xa4, 0xca, 0x6c, 0xba,
	0x7f, 0xde, 0xbe, 0xf7, 0x21, 0x64, 0x1c, 0xa5, 0x5c, 0xc2, 0x2f, 0x08, 0xf8, 0x8d, 0xe9, 0xc8,
	0x5a, 0x93, 0xf0, 0xb3, 0x98, 0xed, 0x5d, 0x11, 0x8b, 0x52, 0x03, 0x2f, 0x01, 0xdc, 0x78, 0x88,
	0xbb, 0x68, 0x50, 0x4c, 0xe3, 0x2f, 0xd2, 0x9f, 0xe1, 0x38, 0x06, 0xb0, 0xf6, 0x18, 0xa7, 0x31,
	0x0d, 0xf5, 0x3a, 0xac, 0x75, 0x31, 0x89, 0x78, 0x47, 0x1c, 0x55, 0xf1, 0xd4, 0x4a, 0x0f, 0x60,
	0x0d, 0x25, 0x02, 0x61, 0xee, 0x9d, 0xba, 0x9b, 0x19, 0xe6, 0x42, 0xa6, 0x50, 0xd2, 0x3b, 0x55,
	0x41, 0xf3, 0x61, 0x01, 0xd6, 0x25, 0x4d, 0x1c, 0xfc, 0x2f, 0x1f, 0x55, 0x8f, 0xe0, 0x4a, 0x0e,
	0xd5, 0x13, 0xec, 0x4c, 0x5d, 0x75, 0xf3, 0x57, 0x50, 0xb2, 0x45, 0xd7, 0x54, 0xd7, 0xab, 0x2e,
	0xe5, 0x4b, 0x22, 0xb6, 0xb7, 0xac, 0x76, 0x64, 0x3a, 0x3b, 0xf3, 0xd5, 0x5e, 0x01, 0x31, 0xa7,
	0x04, 0x11, 0x4c, 0xf8, 0x01, 0x0d, 0x0e, 0x71, 0xf8, 0x4f, 0xec, 0xe3, 0xee, 0x9f, 0x8e, 0x4d,
	0x30, 0x1c, 0x9b, 0xe0, 0xdb, 0xd8, 0x04, 0x6f, 0x26, 0xa6, 0x36, 0x9c, 0x98, 0xda, 0x97, 0x89,
	0xa9, 0x3d, 0xdb, 0xfa, 0xad, 0x05, 0x5e, 0xa8, 0xe7, 0x42, 0xbd, 0x53, 0xc2, 0x11, 0x7e, 0x4d,
	0x3c, 0x18, 0xf7, 0xbe, 0x07, 0x00, 0x00, 0xff, 0xff, 0x80, 0x00, 0xd5, 0x07, 0xc6, 0x06, 0x00,
	0x00,
}

func (m *BaseVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EndTime != 0 {
		i = encodeVarintVesting(dAtA, i, uint64(m.EndTime))
		i--
		dAtA[i] = 0x28
	}
	if len(m.DelegatedVesting) > 0 {
		for iNdEx := len(m.DelegatedVesting) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegatedVesting[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.DelegatedFree) > 0 {
		for iNdEx := len(m.DelegatedFree) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegatedFree[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.OriginalVesting) > 0 {
		for iNdEx := len(m.OriginalVesting) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OriginalVesting[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.BaseAccount != nil {
		{
			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVesting(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ContinuousVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContinuousVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContinuousVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StartTime != 0 {
		i = encodeVarintVesting(dAtA, i, uint64(m.StartTime))
		i--
		dAtA[i] = 0x10
	}
	if m.BaseVestingAccount != nil {
		{
			size, err := m.BaseVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVesting(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DelayedVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelayedVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelayedVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BaseVestingAccount != nil {
		{
			size, err := m.BaseVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVesting(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Period) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Period) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Period) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Length != 0 {
		i = encodeVarintVesting(dAtA, i, uint64(m.Length))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PeriodicVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeriodicVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PeriodicVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VestingPeriods) > 0 {
		for iNdEx := len(m.VestingPeriods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingPeriods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.StartTime != 0 {
		i = encodeVarintVesting(dAtA, i, uint64(m.StartTime))
		i--
		dAtA[i] = 0x10
	}
	if m.BaseVestingAccount != nil {
		{
			size, err := m.BaseVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVesting(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PermanentLockedAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PermanentLockedAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PermanentLockedAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BaseVestingAccount != nil {
		{
			size, err := m.BaseVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVesting(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVesting(dAtA []byte, offset int, v uint64) int {
	offset -= sovVesting(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovVesting(uint64(l))
	}
	if len(m.OriginalVesting) > 0 {
		for _, e := range m.OriginalVesting {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	if len(m.DelegatedFree) > 0 {
		for _, e := range m.DelegatedFree {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	if len(m.DelegatedVesting) > 0 {
		for _, e := range m.DelegatedVesting {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	if m.EndTime != 0 {
		n += 1 + sovVesting(uint64(m.EndTime))
	}
	return n
}

func (m *ContinuousVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseVestingAccount != nil {
		l = m.BaseVestingAccount.Size()
		n += 1 + l + sovVesting(uint64(l))
	}
	if m.StartTime != 0 {
		n += 1 + sovVesting(uint64(m.StartTime))
	}
	return n
}

func (m *DelayedVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseVestingAccount != nil {
		l = m.BaseVestingAccount.Size()
		n += 1 + l + sovVesting(uint64(l))
	}
	return n
}

func (m *Period) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Length != 0 {
		n += 1 + sovVesting(uint64(m.Length))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	return n
}

func (m *PeriodicVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseVestingAccount != nil {
		l = m.BaseVestingAccount.Size()
		n += 1 + l + sovVesting(uint64(l))
	}
	if m.StartTime != 0 {
		n += 1 + sovVesting(uint64(m.StartTime))
	}
	if len(m.VestingPeriods) > 0 {
		for _, e := range m.VestingPeriods {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	return n
}

func (m *PermanentLockedAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseVestingAccount != nil {
		l = m.BaseVestingAccount.Size()
		n += 1 + l + sovVesting(uint64(l))
	}
	return n
}

func sovVesting(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVesting(x uint64) (n int) {
	return sovVesting(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: BaseVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseAccount == nil {
				m.BaseAccount = &types.BaseAccount{}
			}
			if err := m.BaseAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginalVesting", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginalVesting = append(m.OriginalVesting, types1.Coin{})
			if err := m.OriginalVesting[len(m.OriginalVesting)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatedFree", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatedFree = append(m.DelegatedFree, types1.Coin{})
			if err := m.DelegatedFree[len(m.DelegatedFree)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatedVesting", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatedVesting = append(m.DelegatedVesting, types1.Coin{})
			if err := m.DelegatedVesting[len(m.DelegatedVesting)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			m.EndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func (m *ContinuousVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: ContinuousVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContinuousVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseVestingAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseVestingAccount == nil {
				m.BaseVestingAccount = &BaseVestingAccount{}
			}
			if err := m.BaseVestingAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			m.StartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func (m *DelayedVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: DelayedVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelayedVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseVestingAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseVestingAccount == nil {
				m.BaseVestingAccount = &BaseVestingAccount{}
			}
			if err := m.BaseVestingAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func (m *Period) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: Period: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Period: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length", wireType)
			}
			m.Length = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Length |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types1.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func (m *PeriodicVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: PeriodicVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeriodicVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseVestingAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseVestingAccount == nil {
				m.BaseVestingAccount = &BaseVestingAccount{}
			}
			if err := m.BaseVestingAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			m.StartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingPeriods = append(m.VestingPeriods, Period{})
			if err := m.VestingPeriods[len(m.VestingPeriods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func (m *PermanentLockedAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: PermanentLockedAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PermanentLockedAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseVestingAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseVestingAccount == nil {
				m.BaseVestingAccount = &BaseVestingAccount{}
			}
			if err := m.BaseVestingAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func skipVesting(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
				return 0, ErrInvalidLengthVesting
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVesting
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVesting
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVesting        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVesting          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVesting = fmt.Errorf("proto: unexpected end of group")
)
