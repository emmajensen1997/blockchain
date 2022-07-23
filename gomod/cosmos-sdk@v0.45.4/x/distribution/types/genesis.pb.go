


package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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




type DelegatorWithdrawInfo struct {

	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty" yaml:"delegator_address"`

	WithdrawAddress string `protobuf:"bytes,2,opt,name=withdraw_address,json=withdrawAddress,proto3" json:"withdraw_address,omitempty" yaml:"withdraw_address"`
}

func (m *DelegatorWithdrawInfo) Reset()         { *m = DelegatorWithdrawInfo{} }
func (m *DelegatorWithdrawInfo) String() string { return proto.CompactTextString(m) }
func (*DelegatorWithdrawInfo) ProtoMessage()    {}
func (*DelegatorWithdrawInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_76eed0f9489db580, []int{0}
}
func (m *DelegatorWithdrawInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelegatorWithdrawInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelegatorWithdrawInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelegatorWithdrawInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegatorWithdrawInfo.Merge(m, src)
}
func (m *DelegatorWithdrawInfo) XXX_Size() int {
	return m.Size()
}
func (m *DelegatorWithdrawInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegatorWithdrawInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DelegatorWithdrawInfo proto.InternalMessageInfo


type ValidatorOutstandingRewardsRecord struct {

	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty" yaml:"validator_address"`

	OutstandingRewards github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,2,rep,name=outstanding_rewards,json=outstandingRewards,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"outstanding_rewards" yaml:"outstanding_rewards"`
}

func (m *ValidatorOutstandingRewardsRecord) Reset()         { *m = ValidatorOutstandingRewardsRecord{} }
func (m *ValidatorOutstandingRewardsRecord) String() string { return proto.CompactTextString(m) }
func (*ValidatorOutstandingRewardsRecord) ProtoMessage()    {}
func (*ValidatorOutstandingRewardsRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_76eed0f9489db580, []int{1}
}
func (m *ValidatorOutstandingRewardsRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorOutstandingRewardsRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorOutstandingRewardsRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorOutstandingRewardsRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorOutstandingRewardsRecord.Merge(m, src)
}
func (m *ValidatorOutstandingRewardsRecord) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorOutstandingRewardsRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorOutstandingRewardsRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorOutstandingRewardsRecord proto.InternalMessageInfo



type ValidatorAccumulatedCommissionRecord struct {

	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty" yaml:"validator_address"`

	Accumulated ValidatorAccumulatedCommission `protobuf:"bytes,2,opt,name=accumulated,proto3" json:"accumulated" yaml:"accumulated"`
}

func (m *ValidatorAccumulatedCommissionRecord) Reset()         { *m = ValidatorAccumulatedCommissionRecord{} }
func (m *ValidatorAccumulatedCommissionRecord) String() string { return proto.CompactTextString(m) }
func (*ValidatorAccumulatedCommissionRecord) ProtoMessage()    {}
func (*ValidatorAccumulatedCommissionRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_76eed0f9489db580, []int{2}
}
func (m *ValidatorAccumulatedCommissionRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorAccumulatedCommissionRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorAccumulatedCommissionRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorAccumulatedCommissionRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorAccumulatedCommissionRecord.Merge(m, src)
}
func (m *ValidatorAccumulatedCommissionRecord) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorAccumulatedCommissionRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorAccumulatedCommissionRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorAccumulatedCommissionRecord proto.InternalMessageInfo



type ValidatorHistoricalRewardsRecord struct {

	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty" yaml:"validator_address"`

	Period uint64 `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`

	Rewards ValidatorHistoricalRewards `protobuf:"bytes,3,opt,name=rewards,proto3" json:"rewards" yaml:"rewards"`
}

func (m *ValidatorHistoricalRewardsRecord) Reset()         { *m = ValidatorHistoricalRewardsRecord{} }
func (m *ValidatorHistoricalRewardsRecord) String() string { return proto.CompactTextString(m) }
func (*ValidatorHistoricalRewardsRecord) ProtoMessage()    {}
func (*ValidatorHistoricalRewardsRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_76eed0f9489db580, []int{3}
}
func (m *ValidatorHistoricalRewardsRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorHistoricalRewardsRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorHistoricalRewardsRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorHistoricalRewardsRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorHistoricalRewardsRecord.Merge(m, src)
}
func (m *ValidatorHistoricalRewardsRecord) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorHistoricalRewardsRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorHistoricalRewardsRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorHistoricalRewardsRecord proto.InternalMessageInfo


type ValidatorCurrentRewardsRecord struct {

	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty" yaml:"validator_address"`

	Rewards ValidatorCurrentRewards `protobuf:"bytes,2,opt,name=rewards,proto3" json:"rewards" yaml:"rewards"`
}

func (m *ValidatorCurrentRewardsRecord) Reset()         { *m = ValidatorCurrentRewardsRecord{} }
func (m *ValidatorCurrentRewardsRecord) String() string { return proto.CompactTextString(m) }
func (*ValidatorCurrentRewardsRecord) ProtoMessage()    {}
func (*ValidatorCurrentRewardsRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_76eed0f9489db580, []int{4}
}
func (m *ValidatorCurrentRewardsRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorCurrentRewardsRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorCurrentRewardsRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorCurrentRewardsRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorCurrentRewardsRecord.Merge(m, src)
}
func (m *ValidatorCurrentRewardsRecord) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorCurrentRewardsRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorCurrentRewardsRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorCurrentRewardsRecord proto.InternalMessageInfo


type DelegatorStartingInfoRecord struct {

	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty" yaml:"delegator_address"`

	ValidatorAddress string `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty" yaml:"validator_address"`

	StartingInfo DelegatorStartingInfo `protobuf:"bytes,3,opt,name=starting_info,json=startingInfo,proto3" json:"starting_info" yaml:"starting_info"`
}

func (m *DelegatorStartingInfoRecord) Reset()         { *m = DelegatorStartingInfoRecord{} }
func (m *DelegatorStartingInfoRecord) String() string { return proto.CompactTextString(m) }
func (*DelegatorStartingInfoRecord) ProtoMessage()    {}
func (*DelegatorStartingInfoRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_76eed0f9489db580, []int{5}
}
func (m *DelegatorStartingInfoRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelegatorStartingInfoRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelegatorStartingInfoRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelegatorStartingInfoRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegatorStartingInfoRecord.Merge(m, src)
}
func (m *DelegatorStartingInfoRecord) XXX_Size() int {
	return m.Size()
}
func (m *DelegatorStartingInfoRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegatorStartingInfoRecord.DiscardUnknown(m)
}

var xxx_messageInfo_DelegatorStartingInfoRecord proto.InternalMessageInfo


type ValidatorSlashEventRecord struct {

	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty" yaml:"validator_address"`

	Height uint64 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`

	Period uint64 `protobuf:"varint,3,opt,name=period,proto3" json:"period,omitempty"`

	ValidatorSlashEvent ValidatorSlashEvent `protobuf:"bytes,4,opt,name=validator_slash_event,json=validatorSlashEvent,proto3" json:"validator_slash_event" yaml:"event"`
}

func (m *ValidatorSlashEventRecord) Reset()         { *m = ValidatorSlashEventRecord{} }
func (m *ValidatorSlashEventRecord) String() string { return proto.CompactTextString(m) }
func (*ValidatorSlashEventRecord) ProtoMessage()    {}
func (*ValidatorSlashEventRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_76eed0f9489db580, []int{6}
}
func (m *ValidatorSlashEventRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSlashEventRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSlashEventRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSlashEventRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSlashEventRecord.Merge(m, src)
}
func (m *ValidatorSlashEventRecord) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSlashEventRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSlashEventRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSlashEventRecord proto.InternalMessageInfo


type GenesisState struct {

	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params" yaml:"params"`

	FeePool FeePool `protobuf:"bytes,2,opt,name=fee_pool,json=feePool,proto3" json:"fee_pool" yaml:"fee_pool"`

	DelegatorWithdrawInfos []DelegatorWithdrawInfo `protobuf:"bytes,3,rep,name=delegator_withdraw_infos,json=delegatorWithdrawInfos,proto3" json:"delegator_withdraw_infos" yaml:"delegator_withdraw_infos"`

	PreviousProposer string `protobuf:"bytes,4,opt,name=previous_proposer,json=previousProposer,proto3" json:"previous_proposer,omitempty" yaml:"previous_proposer"`

	OutstandingRewards []ValidatorOutstandingRewardsRecord `protobuf:"bytes,5,rep,name=outstanding_rewards,json=outstandingRewards,proto3" json:"outstanding_rewards" yaml:"outstanding_rewards"`

	ValidatorAccumulatedCommissions []ValidatorAccumulatedCommissionRecord `protobuf:"bytes,6,rep,name=validator_accumulated_commissions,json=validatorAccumulatedCommissions,proto3" json:"validator_accumulated_commissions" yaml:"validator_accumulated_commissions"`

	ValidatorHistoricalRewards []ValidatorHistoricalRewardsRecord `protobuf:"bytes,7,rep,name=validator_historical_rewards,json=validatorHistoricalRewards,proto3" json:"validator_historical_rewards" yaml:"validator_historical_rewards"`

	ValidatorCurrentRewards []ValidatorCurrentRewardsRecord `protobuf:"bytes,8,rep,name=validator_current_rewards,json=validatorCurrentRewards,proto3" json:"validator_current_rewards" yaml:"validator_current_rewards"`

	DelegatorStartingInfos []DelegatorStartingInfoRecord `protobuf:"bytes,9,rep,name=delegator_starting_infos,json=delegatorStartingInfos,proto3" json:"delegator_starting_infos" yaml:"delegator_starting_infos"`

	ValidatorSlashEvents []ValidatorSlashEventRecord `protobuf:"bytes,10,rep,name=validator_slash_events,json=validatorSlashEvents,proto3" json:"validator_slash_events" yaml:"validator_slash_events"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_76eed0f9489db580, []int{7}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DelegatorWithdrawInfo)(nil), "cosmos.distribution.v1beta1.DelegatorWithdrawInfo")
	proto.RegisterType((*ValidatorOutstandingRewardsRecord)(nil), "cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord")
	proto.RegisterType((*ValidatorAccumulatedCommissionRecord)(nil), "cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord")
	proto.RegisterType((*ValidatorHistoricalRewardsRecord)(nil), "cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord")
	proto.RegisterType((*ValidatorCurrentRewardsRecord)(nil), "cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord")
	proto.RegisterType((*DelegatorStartingInfoRecord)(nil), "cosmos.distribution.v1beta1.DelegatorStartingInfoRecord")
	proto.RegisterType((*ValidatorSlashEventRecord)(nil), "cosmos.distribution.v1beta1.ValidatorSlashEventRecord")
	proto.RegisterType((*GenesisState)(nil), "cosmos.distribution.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("cosmos/distribution/v1beta1/genesis.proto", fileDescriptor_76eed0f9489db580)
}

var fileDescriptor_76eed0f9489db580 = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0xcd, 0x6f, 0x1b, 0x45,
	0x1c, 0xf5, 0x3a, 0x25, 0x49, 0x27, 0x29, 0x0d, 0xdb, 0x7c, 0xb8, 0x4e, 0xea, 0x4d, 0xa7, 0x45,
	0x04, 0x55, 0xac, 0x9b, 0x80, 0x00, 0x05, 0x81, 0x94, 0x4d, 0x29, 0xf4, 0xd4, 0x30, 0x91, 0x00,
	0x71, 0xb1, 0xd6, 0xbb, 0x63, 0x7b, 0x84, 0xbd, 0x63, 0xcd, 0x8c, 0x1d, 0xc2, 0x3f, 0x00, 0x47,
	0x24, 0xc4, 0xa9, 0x1c, 0x72, 0x44, 0x88, 0x63, 0xef, 0x5c, 0x7b, 0xec, 0x91, 0x03, 0x0a, 0x28,
	0xb9, 0x70, 0xce, 0x81, 0x03, 0x27, 0xb4, 0x33, 0xb3, 0x5f, 0xf6, 0xda, 0x38, 0xa1, 0x39, 0x25,
	0x1e, 0xff, 0xf6, 0xbd, 0xf7, 0x7b, 0xf3, 0xfb, 0x58, 0x83, 0xd7, 0x3d, 0xca, 0x3b, 0x94, 0x57,
	0x7d, 0xc2, 0x05, 0x23, 0xf5, 0x9e, 0x20, 0x34, 0xa8, 0xf6, 0x37, 0xeb, 0x58, 0xb8, 0x9b, 0xd5,
	0x26, 0x0e, 0x30, 0x27, 0xdc, 0xee, 0x32, 0x2a, 0xa8, 0xb9, 0xaa, 0x42, 0xed, 0x74, 0xa8, 0xad,
	0x43, 0xcb, 0x8b, 0x4d, 0xda, 0xa4, 0x32, 0xae, 0x1a, 0xfe, 0xa7, 0x1e, 0x29, 0x57, 0x34, 0x7a,
	0xdd, 0xe5, 0x38, 0x46, 0xf5, 0x28, 0x09, 0xf4, 0xf7, 0xf6, 0x38, 0xf6, 0x0c, 0x8f, 0x8c, 0x87,
	0x4f, 0x0d, 0xb0, 0xf4, 0x00, 0xb7, 0x71, 0xd3, 0x15, 0x94, 0x7d, 0x46, 0x44, 0xcb, 0x67, 0xee,
	0xc1, 0xa3, 0xa0, 0x41, 0xcd, 0x47, 0xe0, 0x15, 0x3f, 0xfa, 0xa2, 0xe6, 0xfa, 0x3e, 0xc3, 0x9c,
	0x97, 0x8c, 0x75, 0x63, 0xe3, 0xaa, 0xb3, 0x76, 0x76, 0x6c, 0x95, 0x0e, 0xdd, 0x4e, 0x7b, 0x1b,
	0x0e, 0x85, 0x40, 0xb4, 0x10, 0x9f, 0xed, 0xa8, 0x23, 0xf3, 0x21, 0x58, 0x38, 0xd0, 0xd0, 0x31,
	0x52, 0x51, 0x22, 0xad, 0x9e, 0x1d, 0x5b, 0x2b, 0x0a, 0x69, 0x30, 0x02, 0xa2, 0xeb, 0xd1, 0x91,
	0xc6, 0xd9, 0x9e, 0xfd, 0xf6, 0xc8, 0x2a, 0xfc, 0x75, 0x64, 0x15, 0xe0, 0x93, 0x22, 0xb8, 0xfd,
	0xa9, 0xdb, 0x26, 0x7e, 0x48, 0xf3, 0xb8, 0x27, 0xb8, 0x70, 0x03, 0x9f, 0x04, 0x4d, 0x84, 0x0f,
	0x5c, 0xe6, 0x73, 0x84, 0x3d, 0xca, 0xfc, 0x30, 0x85, 0x7e, 0x14, 0x34, 0x3a, 0x85, 0xa1, 0x10,
	0x88, 0x16, 0xe2, 0xb3, 0x28, 0x85, 0x23, 0x03, 0xdc, 0xa0, 0x09, 0x4f, 0x8d, 0x29, 0xa2, 0x52,
	0x71, 0x7d, 0x6a, 0x63, 0x6e, 0x6b, 0x4d, 0xdb, 0x6e, 0x87, 0xd7, 0x12, 0xdd, 0xa0, 0xfd, 0x00,
	0x7b, 0xbb, 0x94, 0x04, 0xce, 0x27, 0xcf, 0x8e, 0xad, 0xc2, 0xd9, 0xb1, 0x55, 0x56, 0x7c, 0x39,
	0x30, 0xf0, 0xe7, 0x3f, 0xac, 0x7b, 0x4d, 0x22, 0x5a, 0xbd, 0xba, 0xed, 0xd1, 0x4e, 0x55, 0x5f,
	0xa2, 0xfa, 0xf3, 0x06, 0xf7, 0xbf, 0xac, 0x8a, 0xc3, 0x2e, 0xe6, 0x11, 0x22, 0x47, 0x26, 0x1d,
	0xca, 0x39, 0xe5, 0xce, 0xdf, 0x06, 0xb8, 0x1b, 0xbb, 0xb3, 0xe3, 0x79, 0xbd, 0x4e, 0xaf, 0xed,
	0x0a, 0xec, 0xef, 0xd2, 0x4e, 0x87, 0x70, 0x4e, 0x68, 0xf0, 0xe2, 0x0d, 0x3a, 0x04, 0x73, 0x6e,
	0xc2, 0x24, 0xaf, 0x77, 0x6e, 0xeb, 0x3d, 0x7b, 0x4c, 0x85, 0xdb, 0xe3, 0x25, 0x3a, 0x65, 0x6d,
	0x9b, 0xa9, 0x54, 0xa4, 0xd0, 0x21, 0x4a, 0x73, 0xa5, 0x12, 0xff, 0xc7, 0x00, 0xeb, 0x31, 0xea,
	0xc7, 0x84, 0x0b, 0xca, 0x88, 0xe7, 0xb6, 0x2f, 0xad, 0x2a, 0x96, 0xc1, 0x74, 0x17, 0x33, 0x42,
	0x55, 0xbe, 0x57, 0x90, 0xfe, 0x64, 0x12, 0x30, 0x13, 0x15, 0xc8, 0x94, 0x34, 0xe2, 0x9d, 0xc9,
	0x8c, 0x18, 0x92, 0xec, 0x2c, 0x6b, 0x13, 0x5e, 0x56, 0xaa, 0xa2, 0x7a, 0x41, 0x11, 0x7e, 0x2a,
	0xf9, 0xdf, 0x0d, 0x70, 0x2b, 0x46, 0xda, 0xed, 0x31, 0x86, 0x03, 0x71, 0x69, 0x99, 0x37, 0x92,
	0x0c, 0xd5, 0x55, 0xbf, 0x35, 0x59, 0x86, 0x59, 0x5d, 0xe7, 0x49, 0xef, 0x69, 0x11, 0xac, 0xc6,
	0x93, 0x6a, 0x5f, 0xb8, 0x4c, 0x90, 0xa0, 0x19, 0x4e, 0xaa, 0x24, 0xb9, 0x17, 0x35, 0xaf, 0x72,
	0x7d, 0x2a, 0x5e, 0xc8, 0xa7, 0x1e, 0xb8, 0xc6, 0xb5, 0xd6, 0x1a, 0x09, 0x1a, 0x54, 0xd7, 0xc3,
	0xd6, 0x58, 0xb7, 0x72, 0xd3, 0x74, 0xd6, 0xb4, 0x57, 0x8b, 0x8a, 0x3e, 0x03, 0x0b, 0xd1, 0x3c,
	0x4f, 0xc5, 0xa6, 0x6c, 0xfb, 0xb1, 0x08, 0x6e, 0xc6, 0xee, 0xef, 0xb7, 0x5d, 0xde, 0xfa, 0xb0,
	0x2f, 0x2f, 0xe0, 0x12, 0x7a, 0xa1, 0x85, 0x49, 0xb3, 0x25, 0xa2, 0x5e, 0x50, 0x9f, 0x52, 0x3d,
	0x32, 0x95, 0xe9, 0x91, 0xaf, 0xc1, 0x52, 0x82, 0xcb, 0x43, 0x61, 0x35, 0x1c, 0x2a, 0x2b, 0x5d,
	0x91, 0x0e, 0xdd, 0x9f, 0xac, 0x9e, 0x92, 0x8c, 0x9c, 0x45, 0xed, 0xcf, 0xbc, 0x12, 0x2d, 0xc1,
	0x20, 0xba, 0xd1, 0x1f, 0x0e, 0x4d, 0xd9, 0xf3, 0xcd, 0x1c, 0x98, 0xff, 0x48, 0x2d, 0xe5, 0x7d,
	0xe1, 0x0a, 0x6c, 0x22, 0x30, 0xdd, 0x75, 0x99, 0xdb, 0x51, 0x36, 0xcc, 0x6d, 0xdd, 0x19, 0xab,
	0x63, 0x4f, 0x86, 0x3a, 0x4b, 0x9a, 0xfa, 0x9a, 0xa2, 0x56, 0x00, 0x10, 0x69, 0x24, 0xf3, 0x73,
	0x30, 0xdb, 0xc0, 0xb8, 0xd6, 0xa5, 0xb4, 0xad, 0xbb, 0xe5, 0xee, 0x58, 0xd4, 0x87, 0x18, 0xef,
	0x51, 0xda, 0x76, 0x56, 0x34, 0xec, 0x75, 0x05, 0x1b, 0x61, 0x40, 0x34, 0xd3, 0x50, 0x11, 0xe6,
	0x0f, 0x06, 0x28, 0x25, 0x25, 0x1d, 0xaf, 0xd0, 0xb0, 0x24, 0xc2, 0xd1, 0x33, 0x35, 0x79, 0xa9,
	0xa5, 0x77, 0xbf, 0xf3, 0x9a, 0x26, 0xb6, 0x06, 0x9b, 0x26, 0xcb, 0x00, 0xd1, 0xb2, 0x9f, 0xf7,
	0xbc, 0xec, 0xa0, 0x2e, 0xc3, 0x7d, 0x42, 0x7b, 0xbc, 0xd6, 0x65, 0xb4, 0x4b, 0x39, 0x66, 0xf2,
	0x62, 0x33, 0x75, 0x35, 0x14, 0x02, 0xd1, 0x42, 0x74, 0xb6, 0xa7, 0x8f, 0xcc, 0xef, 0x47, 0x6c,
	0xde, 0x97, 0x64, 0x76, 0x1f, 0x4c, 0x56, 0x26, 0xa3, 0x5e, 0x11, 0x1c, 0xf8, 0xdf, 0xbb, 0x39,
	0x6f, 0xd9, 0x9a, 0xbf, 0x1a, 0xe0, 0x76, 0xaa, 0x2d, 0x92, 0x6d, 0x54, 0xf3, 0xe2, 0x0d, 0xc6,
	0x4b, 0xd3, 0x52, 0xe3, 0xce, 0xff, 0xd8, 0x82, 0x5a, 0xe6, 0x7d, 0x2d, 0x73, 0x63, 0xa8, 0x21,
	0xf3, 0x99, 0x21, 0xb2, 0xfa, 0x63, 0x71, 0xb9, 0xf9, 0x8b, 0x01, 0xd6, 0x12, 0x9c, 0x56, 0xbc,
	0x79, 0x62, 0x83, 0x67, 0xa4, 0xf8, 0xf7, 0x2f, 0xb8, 0xb9, 0xb4, 0xf0, 0x7b, 0x5a, 0xf8, 0x9d,
	0x41, 0xe1, 0xc3, 0x84, 0x10, 0x95, 0xfb, 0x23, 0xe1, 0xc2, 0x17, 0xb0, 0x9b, 0xc9, 0xd3, 0x9e,
	0x5a, 0x23, 0xb1, 0xd6, 0x59, 0xa9, 0x75, 0xfb, 0x22, 0x3b, 0x48, 0x0b, 0xdd, 0xd0, 0x42, 0xd7,
	0x07, 0x85, 0x0e, 0x50, 0x41, 0xb4, 0xd2, 0xcf, 0x07, 0x32, 0x9f, 0x64, 0x9a, 0x31, 0x33, 0x9f,
	0x79, 0xe9, 0xaa, 0x54, 0xf8, 0xee, 0xf9, 0xe7, 0xbe, 0xd6, 0x37, 0xb2, 0x25, 0xb3, 0x3c, 0xe9,
	0x96, 0x4c, 0xa3, 0xf0, 0xb0, 0x8f, 0x96, 0x73, 0x07, 0x2e, 0x2f, 0x01, 0xa9, 0xed, 0xed, 0xf3,
	0x4e, 0x5c, 0xad, 0xec, 0x55, 0xad, 0xec, 0xd6, 0xa0, 0x73, 0x69, 0x0e, 0x88, 0x16, 0x73, 0x06,
	0x71, 0x6a, 0xbf, 0x3b, 0x8f, 0x7f, 0x3a, 0xa9, 0x18, 0xcf, 0x4e, 0x2a, 0xc6, 0xf3, 0x93, 0x8a,
	0xf1, 0xe7, 0x49, 0xc5, 0xf8, 0xee, 0xb4, 0x52, 0x78, 0x7e, 0x5a, 0x29, 0xfc, 0x76, 0x5a, 0x29,
	0x7c, 0xb1, 0x39, 0xf6, 0xed, 0xf8, 0xab, 0xec, 0xef, 0x1d, 0xf9, 0xb2, 0x5c, 0x9f, 0x96, 0xbf,
	0x70, 0xde, 0xfc, 0x37, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xee, 0xfe, 0x4a, 0x91, 0x0d, 0x00, 0x00,
}

func (m *DelegatorWithdrawInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegatorWithdrawInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelegatorWithdrawInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WithdrawAddress) > 0 {
		i -= len(m.WithdrawAddress)
		copy(dAtA[i:], m.WithdrawAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.WithdrawAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorOutstandingRewardsRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorOutstandingRewardsRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorOutstandingRewardsRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OutstandingRewards) > 0 {
		for iNdEx := len(m.OutstandingRewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OutstandingRewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorAccumulatedCommissionRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorAccumulatedCommissionRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorAccumulatedCommissionRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Accumulated.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorHistoricalRewardsRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorHistoricalRewardsRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorHistoricalRewardsRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Rewards.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Period != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Period))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorCurrentRewardsRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorCurrentRewardsRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorCurrentRewardsRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Rewards.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DelegatorStartingInfoRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegatorStartingInfoRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelegatorStartingInfoRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.StartingInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorSlashEventRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSlashEventRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSlashEventRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ValidatorSlashEvent.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Period != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Period))
		i--
		dAtA[i] = 0x18
	}
	if m.Height != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorSlashEvents) > 0 {
		for iNdEx := len(m.ValidatorSlashEvents) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValidatorSlashEvents[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.DelegatorStartingInfos) > 0 {
		for iNdEx := len(m.DelegatorStartingInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegatorStartingInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.ValidatorCurrentRewards) > 0 {
		for iNdEx := len(m.ValidatorCurrentRewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValidatorCurrentRewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.ValidatorHistoricalRewards) > 0 {
		for iNdEx := len(m.ValidatorHistoricalRewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValidatorHistoricalRewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.ValidatorAccumulatedCommissions) > 0 {
		for iNdEx := len(m.ValidatorAccumulatedCommissions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValidatorAccumulatedCommissions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.OutstandingRewards) > 0 {
		for iNdEx := len(m.OutstandingRewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OutstandingRewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.PreviousProposer) > 0 {
		i -= len(m.PreviousProposer)
		copy(dAtA[i:], m.PreviousProposer)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PreviousProposer)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.DelegatorWithdrawInfos) > 0 {
		for iNdEx := len(m.DelegatorWithdrawInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegatorWithdrawInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.FeePool.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DelegatorWithdrawInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.WithdrawAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *ValidatorOutstandingRewardsRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.OutstandingRewards) > 0 {
		for _, e := range m.OutstandingRewards {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *ValidatorAccumulatedCommissionRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.Accumulated.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *ValidatorHistoricalRewardsRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Period != 0 {
		n += 1 + sovGenesis(uint64(m.Period))
	}
	l = m.Rewards.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *ValidatorCurrentRewardsRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.Rewards.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *DelegatorStartingInfoRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.StartingInfo.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *ValidatorSlashEventRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovGenesis(uint64(m.Height))
	}
	if m.Period != 0 {
		n += 1 + sovGenesis(uint64(m.Period))
	}
	l = m.ValidatorSlashEvent.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.FeePool.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.DelegatorWithdrawInfos) > 0 {
		for _, e := range m.DelegatorWithdrawInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = len(m.PreviousProposer)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.OutstandingRewards) > 0 {
		for _, e := range m.OutstandingRewards {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ValidatorAccumulatedCommissions) > 0 {
		for _, e := range m.ValidatorAccumulatedCommissions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ValidatorHistoricalRewards) > 0 {
		for _, e := range m.ValidatorHistoricalRewards {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ValidatorCurrentRewards) > 0 {
		for _, e := range m.ValidatorCurrentRewards {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DelegatorStartingInfos) > 0 {
		for _, e := range m.DelegatorStartingInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ValidatorSlashEvents) > 0 {
		for _, e := range m.ValidatorSlashEvents {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DelegatorWithdrawInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: DelegatorWithdrawInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegatorWithdrawInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *ValidatorOutstandingRewardsRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ValidatorOutstandingRewardsRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorOutstandingRewardsRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutstandingRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutstandingRewards = append(m.OutstandingRewards, types.DecCoin{})
			if err := m.OutstandingRewards[len(m.OutstandingRewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *ValidatorAccumulatedCommissionRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ValidatorAccumulatedCommissionRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorAccumulatedCommissionRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accumulated", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Accumulated.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *ValidatorHistoricalRewardsRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ValidatorHistoricalRewardsRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorHistoricalRewardsRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			m.Period = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Period |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Rewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *ValidatorCurrentRewardsRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ValidatorCurrentRewardsRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorCurrentRewardsRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Rewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *DelegatorStartingInfoRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: DelegatorStartingInfoRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegatorStartingInfoRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartingInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StartingInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *ValidatorSlashEventRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ValidatorSlashEventRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSlashEventRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			m.Period = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Period |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorSlashEvent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ValidatorSlashEvent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeePool", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeePool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorWithdrawInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorWithdrawInfos = append(m.DelegatorWithdrawInfos, DelegatorWithdrawInfo{})
			if err := m.DelegatorWithdrawInfos[len(m.DelegatorWithdrawInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousProposer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviousProposer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutstandingRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutstandingRewards = append(m.OutstandingRewards, ValidatorOutstandingRewardsRecord{})
			if err := m.OutstandingRewards[len(m.OutstandingRewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAccumulatedCommissions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAccumulatedCommissions = append(m.ValidatorAccumulatedCommissions, ValidatorAccumulatedCommissionRecord{})
			if err := m.ValidatorAccumulatedCommissions[len(m.ValidatorAccumulatedCommissions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorHistoricalRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorHistoricalRewards = append(m.ValidatorHistoricalRewards, ValidatorHistoricalRewardsRecord{})
			if err := m.ValidatorHistoricalRewards[len(m.ValidatorHistoricalRewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorCurrentRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorCurrentRewards = append(m.ValidatorCurrentRewards, ValidatorCurrentRewardsRecord{})
			if err := m.ValidatorCurrentRewards[len(m.ValidatorCurrentRewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorStartingInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorStartingInfos = append(m.DelegatorStartingInfos, DelegatorStartingInfoRecord{})
			if err := m.DelegatorStartingInfos[len(m.DelegatorStartingInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorSlashEvents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorSlashEvents = append(m.ValidatorSlashEvents, ValidatorSlashEventRecord{})
			if err := m.ValidatorSlashEvents[len(m.ValidatorSlashEvents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
