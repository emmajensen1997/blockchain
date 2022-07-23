


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


type Params struct {
	CommunityTax        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=community_tax,json=communityTax,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"community_tax" yaml:"community_tax"`
	BaseProposerReward  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=base_proposer_reward,json=baseProposerReward,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"base_proposer_reward" yaml:"base_proposer_reward"`
	BonusProposerReward github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=bonus_proposer_reward,json=bonusProposerReward,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"bonus_proposer_reward" yaml:"bonus_proposer_reward"`
	WithdrawAddrEnabled bool                                   `protobuf:"varint,4,opt,name=withdraw_addr_enabled,json=withdrawAddrEnabled,proto3" json:"withdraw_addr_enabled,omitempty" yaml:"withdraw_addr_enabled"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd78a31ea281a992, []int{0}
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

func (m *Params) GetWithdrawAddrEnabled() bool {
	if m != nil {
		return m.WithdrawAddrEnabled
	}
	return false
}













type ValidatorHistoricalRewards struct {
	CumulativeRewardRatio github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=cumulative_reward_ratio,json=cumulativeRewardRatio,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"cumulative_reward_ratio" yaml:"cumulative_reward_ratio"`
	ReferenceCount        uint32                                      `protobuf:"varint,2,opt,name=reference_count,json=referenceCount,proto3" json:"reference_count,omitempty" yaml:"reference_count"`
}

func (m *ValidatorHistoricalRewards) Reset()         { *m = ValidatorHistoricalRewards{} }
func (m *ValidatorHistoricalRewards) String() string { return proto.CompactTextString(m) }
func (*ValidatorHistoricalRewards) ProtoMessage()    {}
func (*ValidatorHistoricalRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd78a31ea281a992, []int{1}
}
func (m *ValidatorHistoricalRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorHistoricalRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorHistoricalRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorHistoricalRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorHistoricalRewards.Merge(m, src)
}
func (m *ValidatorHistoricalRewards) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorHistoricalRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorHistoricalRewards.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorHistoricalRewards proto.InternalMessageInfo

func (m *ValidatorHistoricalRewards) GetCumulativeRewardRatio() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.CumulativeRewardRatio
	}
	return nil
}

func (m *ValidatorHistoricalRewards) GetReferenceCount() uint32 {
	if m != nil {
		return m.ReferenceCount
	}
	return 0
}




type ValidatorCurrentRewards struct {
	Rewards github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=rewards,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"rewards"`
	Period  uint64                                      `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
}

func (m *ValidatorCurrentRewards) Reset()         { *m = ValidatorCurrentRewards{} }
func (m *ValidatorCurrentRewards) String() string { return proto.CompactTextString(m) }
func (*ValidatorCurrentRewards) ProtoMessage()    {}
func (*ValidatorCurrentRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd78a31ea281a992, []int{2}
}
func (m *ValidatorCurrentRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorCurrentRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorCurrentRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorCurrentRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorCurrentRewards.Merge(m, src)
}
func (m *ValidatorCurrentRewards) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorCurrentRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorCurrentRewards.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorCurrentRewards proto.InternalMessageInfo

func (m *ValidatorCurrentRewards) GetRewards() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Rewards
	}
	return nil
}

func (m *ValidatorCurrentRewards) GetPeriod() uint64 {
	if m != nil {
		return m.Period
	}
	return 0
}



type ValidatorAccumulatedCommission struct {
	Commission github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=commission,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"commission"`
}

func (m *ValidatorAccumulatedCommission) Reset()         { *m = ValidatorAccumulatedCommission{} }
func (m *ValidatorAccumulatedCommission) String() string { return proto.CompactTextString(m) }
func (*ValidatorAccumulatedCommission) ProtoMessage()    {}
func (*ValidatorAccumulatedCommission) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd78a31ea281a992, []int{3}
}
func (m *ValidatorAccumulatedCommission) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorAccumulatedCommission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorAccumulatedCommission.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorAccumulatedCommission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorAccumulatedCommission.Merge(m, src)
}
func (m *ValidatorAccumulatedCommission) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorAccumulatedCommission) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorAccumulatedCommission.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorAccumulatedCommission proto.InternalMessageInfo

func (m *ValidatorAccumulatedCommission) GetCommission() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Commission
	}
	return nil
}



type ValidatorOutstandingRewards struct {
	Rewards github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=rewards,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"rewards" yaml:"rewards"`
}

func (m *ValidatorOutstandingRewards) Reset()         { *m = ValidatorOutstandingRewards{} }
func (m *ValidatorOutstandingRewards) String() string { return proto.CompactTextString(m) }
func (*ValidatorOutstandingRewards) ProtoMessage()    {}
func (*ValidatorOutstandingRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd78a31ea281a992, []int{4}
}
func (m *ValidatorOutstandingRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorOutstandingRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorOutstandingRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorOutstandingRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorOutstandingRewards.Merge(m, src)
}
func (m *ValidatorOutstandingRewards) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorOutstandingRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorOutstandingRewards.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorOutstandingRewards proto.InternalMessageInfo

func (m *ValidatorOutstandingRewards) GetRewards() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Rewards
	}
	return nil
}





type ValidatorSlashEvent struct {
	ValidatorPeriod uint64                                 `protobuf:"varint,1,opt,name=validator_period,json=validatorPeriod,proto3" json:"validator_period,omitempty" yaml:"validator_period"`
	Fraction        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=fraction,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"fraction"`
}

func (m *ValidatorSlashEvent) Reset()         { *m = ValidatorSlashEvent{} }
func (m *ValidatorSlashEvent) String() string { return proto.CompactTextString(m) }
func (*ValidatorSlashEvent) ProtoMessage()    {}
func (*ValidatorSlashEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd78a31ea281a992, []int{5}
}
func (m *ValidatorSlashEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSlashEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSlashEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSlashEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSlashEvent.Merge(m, src)
}
func (m *ValidatorSlashEvent) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSlashEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSlashEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSlashEvent proto.InternalMessageInfo

func (m *ValidatorSlashEvent) GetValidatorPeriod() uint64 {
	if m != nil {
		return m.ValidatorPeriod
	}
	return 0
}


type ValidatorSlashEvents struct {
	ValidatorSlashEvents []ValidatorSlashEvent `protobuf:"bytes,1,rep,name=validator_slash_events,json=validatorSlashEvents,proto3" json:"validator_slash_events" yaml:"validator_slash_events"`
}

func (m *ValidatorSlashEvents) Reset()      { *m = ValidatorSlashEvents{} }
func (*ValidatorSlashEvents) ProtoMessage() {}
func (*ValidatorSlashEvents) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd78a31ea281a992, []int{6}
}
func (m *ValidatorSlashEvents) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSlashEvents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSlashEvents.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSlashEvents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSlashEvents.Merge(m, src)
}
func (m *ValidatorSlashEvents) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSlashEvents) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSlashEvents.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSlashEvents proto.InternalMessageInfo

func (m *ValidatorSlashEvents) GetValidatorSlashEvents() []ValidatorSlashEvent {
	if m != nil {
		return m.ValidatorSlashEvents
	}
	return nil
}


type FeePool struct {
	CommunityPool github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=community_pool,json=communityPool,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"community_pool" yaml:"community_pool"`
}

func (m *FeePool) Reset()         { *m = FeePool{} }
func (m *FeePool) String() string { return proto.CompactTextString(m) }
func (*FeePool) ProtoMessage()    {}
func (*FeePool) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd78a31ea281a992, []int{7}
}
func (m *FeePool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeePool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeePool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeePool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeePool.Merge(m, src)
}
func (m *FeePool) XXX_Size() int {
	return m.Size()
}
func (m *FeePool) XXX_DiscardUnknown() {
	xxx_messageInfo_FeePool.DiscardUnknown(m)
}

var xxx_messageInfo_FeePool proto.InternalMessageInfo

func (m *FeePool) GetCommunityPool() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.CommunityPool
	}
	return nil
}




type CommunityPoolSpendProposal struct {
	Title       string                                   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                                   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Recipient   string                                   `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Amount      github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *CommunityPoolSpendProposal) Reset()      { *m = CommunityPoolSpendProposal{} }
func (*CommunityPoolSpendProposal) ProtoMessage() {}
func (*CommunityPoolSpendProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd78a31ea281a992, []int{8}
}
func (m *CommunityPoolSpendProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommunityPoolSpendProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommunityPoolSpendProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommunityPoolSpendProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommunityPoolSpendProposal.Merge(m, src)
}
func (m *CommunityPoolSpendProposal) XXX_Size() int {
	return m.Size()
}
func (m *CommunityPoolSpendProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CommunityPoolSpendProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CommunityPoolSpendProposal proto.InternalMessageInfo







type DelegatorStartingInfo struct {
	PreviousPeriod uint64                                 `protobuf:"varint,1,opt,name=previous_period,json=previousPeriod,proto3" json:"previous_period,omitempty" yaml:"previous_period"`
	Stake          github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=stake,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"stake" yaml:"stake"`
	Height         uint64                                 `protobuf:"varint,3,opt,name=height,proto3" json:"creation_height" yaml:"creation_height"`
}

func (m *DelegatorStartingInfo) Reset()         { *m = DelegatorStartingInfo{} }
func (m *DelegatorStartingInfo) String() string { return proto.CompactTextString(m) }
func (*DelegatorStartingInfo) ProtoMessage()    {}
func (*DelegatorStartingInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd78a31ea281a992, []int{9}
}
func (m *DelegatorStartingInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelegatorStartingInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelegatorStartingInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelegatorStartingInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegatorStartingInfo.Merge(m, src)
}
func (m *DelegatorStartingInfo) XXX_Size() int {
	return m.Size()
}
func (m *DelegatorStartingInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegatorStartingInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DelegatorStartingInfo proto.InternalMessageInfo

func (m *DelegatorStartingInfo) GetPreviousPeriod() uint64 {
	if m != nil {
		return m.PreviousPeriod
	}
	return 0
}

func (m *DelegatorStartingInfo) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}



type DelegationDelegatorReward struct {
	ValidatorAddress string                                      `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty" yaml:"validator_address"`
	Reward           github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,2,rep,name=reward,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"reward"`
}

func (m *DelegationDelegatorReward) Reset()         { *m = DelegationDelegatorReward{} }
func (m *DelegationDelegatorReward) String() string { return proto.CompactTextString(m) }
func (*DelegationDelegatorReward) ProtoMessage()    {}
func (*DelegationDelegatorReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd78a31ea281a992, []int{10}
}
func (m *DelegationDelegatorReward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelegationDelegatorReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelegationDelegatorReward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelegationDelegatorReward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegationDelegatorReward.Merge(m, src)
}
func (m *DelegationDelegatorReward) XXX_Size() int {
	return m.Size()
}
func (m *DelegationDelegatorReward) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegationDelegatorReward.DiscardUnknown(m)
}

var xxx_messageInfo_DelegationDelegatorReward proto.InternalMessageInfo



type CommunityPoolSpendProposalWithDeposit struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	Recipient   string `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty" yaml:"recipient"`
	Amount      string `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty" yaml:"amount"`
	Deposit     string `protobuf:"bytes,5,opt,name=deposit,proto3" json:"deposit,omitempty" yaml:"deposit"`
}

func (m *CommunityPoolSpendProposalWithDeposit) Reset()         { *m = CommunityPoolSpendProposalWithDeposit{} }
func (m *CommunityPoolSpendProposalWithDeposit) String() string { return proto.CompactTextString(m) }
func (*CommunityPoolSpendProposalWithDeposit) ProtoMessage()    {}
func (*CommunityPoolSpendProposalWithDeposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd78a31ea281a992, []int{11}
}
func (m *CommunityPoolSpendProposalWithDeposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommunityPoolSpendProposalWithDeposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommunityPoolSpendProposalWithDeposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommunityPoolSpendProposalWithDeposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommunityPoolSpendProposalWithDeposit.Merge(m, src)
}
func (m *CommunityPoolSpendProposalWithDeposit) XXX_Size() int {
	return m.Size()
}
func (m *CommunityPoolSpendProposalWithDeposit) XXX_DiscardUnknown() {
	xxx_messageInfo_CommunityPoolSpendProposalWithDeposit.DiscardUnknown(m)
}

var xxx_messageInfo_CommunityPoolSpendProposalWithDeposit proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "cosmos.distribution.v1beta1.Params")
	proto.RegisterType((*ValidatorHistoricalRewards)(nil), "cosmos.distribution.v1beta1.ValidatorHistoricalRewards")
	proto.RegisterType((*ValidatorCurrentRewards)(nil), "cosmos.distribution.v1beta1.ValidatorCurrentRewards")
	proto.RegisterType((*ValidatorAccumulatedCommission)(nil), "cosmos.distribution.v1beta1.ValidatorAccumulatedCommission")
	proto.RegisterType((*ValidatorOutstandingRewards)(nil), "cosmos.distribution.v1beta1.ValidatorOutstandingRewards")
	proto.RegisterType((*ValidatorSlashEvent)(nil), "cosmos.distribution.v1beta1.ValidatorSlashEvent")
	proto.RegisterType((*ValidatorSlashEvents)(nil), "cosmos.distribution.v1beta1.ValidatorSlashEvents")
	proto.RegisterType((*FeePool)(nil), "cosmos.distribution.v1beta1.FeePool")
	proto.RegisterType((*CommunityPoolSpendProposal)(nil), "cosmos.distribution.v1beta1.CommunityPoolSpendProposal")
	proto.RegisterType((*DelegatorStartingInfo)(nil), "cosmos.distribution.v1beta1.DelegatorStartingInfo")
	proto.RegisterType((*DelegationDelegatorReward)(nil), "cosmos.distribution.v1beta1.DelegationDelegatorReward")
	proto.RegisterType((*CommunityPoolSpendProposalWithDeposit)(nil), "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit")
}

func init() {
	proto.RegisterFile("cosmos/distribution/v1beta1/distribution.proto", fileDescriptor_cd78a31ea281a992)
}

var fileDescriptor_cd78a31ea281a992 = []byte{
	
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xf6, 0x24, 0x8e, 0x93, 0x4e, 0xf3, 0xab, 0x13, 0x27, 0x71, 0x93, 0xe0, 0x8d, 0x46, 0x6a,
	0x15, 0x04, 0x75, 0x9a, 0xf6, 0x82, 0x72, 0x40, 0x8a, 0x9d, 0x44, 0x14, 0x01, 0x8d, 0xb6, 0x01,
	0x24, 0x2e, 0xd6, 0x78, 0x77, 0x62, 0x8f, 0x62, 0xef, 0x2c, 0x33, 0x63, 0x27, 0x39, 0x20, 0x24,
	0x4e, 0x5c, 0x10, 0x20, 0x2e, 0x1c, 0x00, 0xe5, 0xc8, 0xaf, 0x3f, 0xa4, 0xc7, 0xde, 0x40, 0x20,
	0x2d, 0x28, 0x11, 0x12, 0xe2, 0xe8, 0x1b, 0x37, 0xb4, 0x3b, 0xb3, 0xbb, 0xb6, 0x6b, 0xaa, 0xb8,
	0x52, 0x4f, 0xf6, 0x7e, 0xf3, 0xe6, 0xbd, 0xef, 0xbd, 0xf7, 0xed, 0x7b, 0x0b, 0x4b, 0x0e, 0x97,
	0x2d, 0x2e, 0x37, 0x5d, 0x26, 0x95, 0x60, 0xb5, 0xb6, 0x62, 0xdc, 0xdb, 0xec, 0x6c, 0xd5, 0xa8,
	0x22, 0x5b, 0x7d, 0x60, 0xc9, 0x17, 0x5c, 0x71, 0xb4, 0xaa, 0xed, 0x4b, 0x7d, 0x47, 0xc6, 0x7e,
	0x25, 0x5f, 0xe7, 0x75, 0x1e, 0xd9, 0x6d, 0x86, 0xff, 0xf4, 0x95, 0x95, 0xa2, 0x09, 0x51, 0x23,
	0x92, 0x26, 0xae, 0x1d, 0xce, 0x8c, 0x4b, 0xfc, 0xcb, 0x38, 0xcc, 0x1d, 0x10, 0x41, 0x5a, 0x12,
	0x1d, 0xc3, 0x19, 0x87, 0xb7, 0x5a, 0x6d, 0x8f, 0xa9, 0xb3, 0xaa, 0x22, 0xa7, 0x05, 0xb0, 0x0e,
	0x36, 0xae, 0x95, 0xf7, 0x1f, 0x07, 0x56, 0xe6, 0xb7, 0xc0, 0xba, 0x5d, 0x67, 0xaa, 0xd1, 0xae,
	0x95, 0x1c, 0xde, 0xda, 0x34, 0x4e, 0xf5, 0xcf, 0x1d, 0xe9, 0x1e, 0x6f, 0xaa, 0x33, 0x9f, 0xca,
	0xd2, 0x2e, 0x75, 0xba, 0x81, 0x95, 0x3f, 0x23, 0xad, 0xe6, 0x36, 0xee, 0x73, 0x86, 0xed, 0xe9,
	0xe4, 0xf9, 0x90, 0x9c, 0xa2, 0x8f, 0x61, 0x3e, 0xa4, 0x54, 0xf5, 0x05, 0xf7, 0xb9, 0xa4, 0xa2,
	0x2a, 0xe8, 0x09, 0x11, 0x6e, 0x61, 0x2c, 0x8a, 0xf9, 0xf6, 0xc8, 0x31, 0x57, 0x75, 0xcc, 0x61,
	0x3e, 0xb1, 0x8d, 0x42, 0xf8, 0xc0, 0xa0, 0x76, 0x04, 0xa2, 0x4f, 0x00, 0x5c, 0xac, 0x71, 0xaf,
	0x2d, 0x9f, 0xa2, 0x30, 0x1e, 0x51, 0x78, 0x67, 0x64, 0x0a, 0x6b, 0x86, 0xc2, 0x30, 0xa7, 0xd8,
	0x5e, 0x88, 0xf0, 0x01, 0x12, 0x87, 0x70, 0xf1, 0x84, 0xa9, 0x86, 0x2b, 0xc8, 0x49, 0x95, 0xb8,
	0xae, 0xa8, 0x52, 0x8f, 0xd4, 0x9a, 0xd4, 0x2d, 0x64, 0xd7, 0xc1, 0xc6, 0x54, 0x79, 0x3d, 0xf5,
	0x3a, 0xd4, 0x0c, 0xdb, 0x0b, 0x31, 0xbe, 0xe3, 0xba, 0x62, 0x4f, 0xa3, 0xdb, 0xd9, 0xaf, 0xcf,
	0xad, 0x0c, 0xfe, 0x7c, 0x0c, 0xae, 0xbc, 0x47, 0x9a, 0xcc, 0x25, 0x8a, 0x8b, 0x37, 0x98, 0x54,
	0x5c, 0x30, 0x87, 0x34, 0x75, 0x64, 0x89, 0x7e, 0x02, 0x70, 0xd9, 0x69, 0xb7, 0xda, 0x4d, 0xa2,
	0x58, 0x87, 0x1a, 0x9a, 0x55, 0x41, 0x14, 0xe3, 0x05, 0xb0, 0x3e, 0xbe, 0x71, 0xfd, 0xde, 0x9a,
	0x91, 0x67, 0x29, 0xac, 0x5e, 0x2c, 0xb3, 0x30, 0xd7, 0x0a, 0x67, 0x5e, 0xf9, 0xdd, 0xb0, 0x3e,
	0xdd, 0xc0, 0x2a, 0x9a, 0x66, 0x0f, 0x77, 0x85, 0x7f, 0xfc, 0xc3, 0x7a, 0xe5, 0x6a, 0x15, 0x0c,
	0xbd, 0x4a, 0x7b, 0x31, 0x75, 0xa4, 0x99, 0xda, 0xa1, 0x1b, 0x54, 0x81, 0x73, 0x82, 0x1e, 0x51,
	0x41, 0x3d, 0x87, 0x56, 0x1d, 0xde, 0xf6, 0x54, 0xa4, 0x94, 0x99, 0xf2, 0x4a, 0x37, 0xb0, 0x96,
	0x34, 0x85, 0x01, 0x03, 0x6c, 0xcf, 0x26, 0x48, 0x25, 0x02, 0xbe, 0x03, 0x70, 0x39, 0xa9, 0x48,
	0xa5, 0x2d, 0x04, 0xf5, 0x54, 0x5c, 0x8e, 0x63, 0x38, 0xa9, 0x79, 0xcb, 0x2b, 0x65, 0x7f, 0x3f,
	0xcc, 0x7e, 0xd4, 0xdc, 0xe2, 0x08, 0x68, 0x09, 0xe6, 0x7c, 0x2a, 0x18, 0xd7, 0x72, 0xcf, 0xda,
	0xe6, 0x09, 0x7f, 0x05, 0x60, 0x31, 0x21, 0xb8, 0xe3, 0x98, 0x52, 0x50, 0xb7, 0xc2, 0x5b, 0x2d,
	0x26, 0x25, 0xe3, 0x1e, 0xfa, 0x10, 0x42, 0x27, 0x79, 0x7a, 0x71, 0x54, 0x7b, 0x82, 0xe0, 0x6f,
	0x00, 0x5c, 0x4d, 0x58, 0x3d, 0x6c, 0x2b, 0xa9, 0x88, 0xe7, 0x32, 0xaf, 0x1e, 0x97, 0xee, 0xa3,
	0xd1, 0x4a, 0xb7, 0x67, 0x84, 0x33, 0x1b, 0x77, 0x2d, 0xba, 0x8a, 0x9f, 0xb7, 0x98, 0xf8, 0x07,
	0x00, 0x17, 0x12, 0x7a, 0x8f, 0x9a, 0x44, 0x36, 0xf6, 0x3a, 0xd4, 0x53, 0x68, 0x1f, 0xce, 0x77,
	0x62, 0xb8, 0x6a, 0xca, 0x1d, 0x4e, 0xb4, 0x6c, 0x79, 0xb5, 0x1b, 0x58, 0xcb, 0x3a, 0xfa, 0xa0,
	0x05, 0xb6, 0xe7, 0x12, 0xe8, 0x20, 0x42, 0xd0, 0x9b, 0x70, 0xea, 0x48, 0x10, 0x27, 0x9c, 0xb5,
	0x66, 0x3a, 0x95, 0x46, 0x1b, 0x0d, 0x76, 0x72, 0x1f, 0xff, 0x0c, 0x60, 0x7e, 0x08, 0x57, 0x89,
	0x3e, 0x03, 0x70, 0x29, 0xe5, 0x22, 0xc3, 0x93, 0x2a, 0x8d, 0x8e, 0x4c, 0x4d, 0xef, 0x96, 0x9e,
	0x31, 0xfb, 0x4b, 0x43, 0x7c, 0x96, 0x6f, 0x99, 0x3a, 0xbf, 0x34, 0x98, 0x69, 0xaf, 0x77, 0x6c,
	0xe7, 0x3b, 0x43, 0xf8, 0x98, 0x11, 0xf2, 0x2d, 0x80, 0x93, 0xfb, 0x94, 0x1e, 0x70, 0xde, 0x44,
	0x5f, 0x02, 0x38, 0x9b, 0x4e, 0x74, 0x9f, 0xf3, 0xe6, 0x95, 0xba, 0xfd, 0x96, 0x61, 0xb1, 0x38,
	0xb8, 0x13, 0x42, 0x0f, 0x23, 0x37, 0x3d, 0x5d, 0x50, 0x21, 0x27, 0xfc, 0x17, 0x80, 0x2b, 0x95,
	0x5e, 0xe4, 0x91, 0x4f, 0x3d, 0x57, 0xcf, 0x58, 0xd2, 0x44, 0x79, 0x38, 0xa1, 0x98, 0x6a, 0x52,
	0xbd, 0xc8, 0x6c, 0xfd, 0x80, 0xd6, 0xe1, 0x75, 0x97, 0x4a, 0x47, 0x30, 0x3f, 0x6d, 0xa9, 0xdd,
	0x0b, 0xa1, 0x35, 0x78, 0x4d, 0x50, 0x87, 0xf9, 0x8c, 0x7a, 0x4a, 0x6f, 0x03, 0x3b, 0x05, 0x90,
	0x03, 0x73, 0xa4, 0x15, 0x4d, 0xa0, 0x6c, 0x94, 0xff, 0xcd, 0xa1, 0xf9, 0x47, 0xc9, 0xdf, 0x35,
	0xaf, 0xde, 0xc6, 0x15, 0x72, 0xd4, 0x09, 0x1a, 0xd7, 0xdb, 0xd3, 0x9f, 0x9e, 0x5b, 0x99, 0xb0,
	0x07, 0x7f, 0x87, 0x7d, 0xf8, 0x17, 0xc0, 0xc5, 0x5d, 0xda, 0xa4, 0xf5, 0xa8, 0x4d, 0x8a, 0x08,
	0xc5, 0xbc, 0xfa, 0x03, 0xef, 0x28, 0x9a, 0x8b, 0xbe, 0xa0, 0x1d, 0xc6, 0xc3, 0x95, 0xd3, 0xab,
	0xf1, 0x9e, 0xb9, 0x38, 0x60, 0x80, 0xed, 0xd9, 0x18, 0x31, 0x0a, 0x3f, 0x84, 0x13, 0x52, 0x91,
	0x63, 0x6a, 0xe4, 0xfd, 0xfa, 0xc8, 0x9b, 0x6f, 0x5a, 0x07, 0x8a, 0x9c, 0x60, 0x5b, 0x3b, 0x43,
	0x7b, 0x30, 0xd7, 0xa0, 0xac, 0xde, 0xd0, 0x25, 0xcc, 0x96, 0xef, 0xfc, 0x13, 0x58, 0x73, 0x8e,
	0xa0, 0xe1, 0x3c, 0xf7, 0xaa, 0xfa, 0x28, 0x25, 0x39, 0x70, 0x80, 0x6d, 0x73, 0x19, 0xff, 0x0e,
	0xe0, 0x4d, 0x93, 0x3b, 0xe3, 0x5e, 0x52, 0x05, 0xb3, 0x40, 0x1f, 0xc0, 0x1b, 0xa9, 0xb0, 0xc3,
	0xd5, 0x48, 0xa5, 0x34, 0xdf, 0x2d, 0x6b, 0xdd, 0xc0, 0x2a, 0x0c, 0x6a, 0xdf, 0x98, 0x60, 0x3b,
	0x9d, 0x0d, 0x3b, 0x1a, 0x42, 0x0c, 0xe6, 0x92, 0x6f, 0x90, 0x17, 0x34, 0x55, 0x4d, 0x80, 0xed,
	0x29, 0xd3, 0x5d, 0x80, 0xcf, 0xc7, 0xe0, 0xad, 0xff, 0x57, 0xf0, 0xfb, 0x4c, 0x35, 0x76, 0xa9,
	0xcf, 0x25, 0x53, 0xe8, 0x76, 0x9f, 0x98, 0xcb, 0xf3, 0x69, 0xd9, 0x23, 0x18, 0xc7, 0xf2, 0x7e,
	0x6d, 0x88, 0xbc, 0xcb, 0x4b, 0xdd, 0xc0, 0x42, 0xda, 0xba, 0xe7, 0x10, 0xf7, 0xcb, 0xfe, 0xde,
	0x53, 0xb2, 0x2f, 0xe7, 0xbb, 0x81, 0x35, 0x1f, 0xcf, 0x69, 0x73, 0x84, 0x7b, 0x5f, 0x86, 0x97,
	0x7b, 0x5e, 0x86, 0xf0, 0xc2, 0x8d, 0x6e, 0x60, 0xcd, 0xe8, 0x0b, 0x1a, 0xc7, 0xb1, 0xa4, 0xd1,
	0xab, 0x70, 0xd2, 0xd5, 0xb9, 0x14, 0x26, 0x22, 0x5b, 0x94, 0x2e, 0x01, 0x73, 0x80, 0xed, 0xd8,
	0x24, 0x2d, 0x51, 0xf9, 0xe1, 0xf7, 0x17, 0x45, 0xf0, 0xf8, 0xa2, 0x08, 0x9e, 0x5c, 0x14, 0xc1,
	0x9f, 0x17, 0x45, 0xf0, 0xc5, 0x65, 0x31, 0xf3, 0xe4, 0xb2, 0x98, 0xf9, 0xf5, 0xb2, 0x98, 0xf9,
	0x60, 0xeb, 0x99, 0xf5, 0x3f, 0xed, 0xff, 0xb4, 0x8e, 0xda, 0x51, 0xcb, 0x45, 0x5f, 0xbe, 0xf7,
	0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x45, 0x1d, 0xb8, 0x42, 0x7e, 0x0b, 0x00, 0x00,
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
	if !this.CommunityTax.Equal(that1.CommunityTax) {
		return false
	}
	if !this.BaseProposerReward.Equal(that1.BaseProposerReward) {
		return false
	}
	if !this.BonusProposerReward.Equal(that1.BonusProposerReward) {
		return false
	}
	if this.WithdrawAddrEnabled != that1.WithdrawAddrEnabled {
		return false
	}
	return true
}
func (this *ValidatorHistoricalRewards) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ValidatorHistoricalRewards)
	if !ok {
		that2, ok := that.(ValidatorHistoricalRewards)
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
	if len(this.CumulativeRewardRatio) != len(that1.CumulativeRewardRatio) {
		return false
	}
	for i := range this.CumulativeRewardRatio {
		if !this.CumulativeRewardRatio[i].Equal(&that1.CumulativeRewardRatio[i]) {
			return false
		}
	}
	if this.ReferenceCount != that1.ReferenceCount {
		return false
	}
	return true
}
func (this *ValidatorCurrentRewards) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ValidatorCurrentRewards)
	if !ok {
		that2, ok := that.(ValidatorCurrentRewards)
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
	if len(this.Rewards) != len(that1.Rewards) {
		return false
	}
	for i := range this.Rewards {
		if !this.Rewards[i].Equal(&that1.Rewards[i]) {
			return false
		}
	}
	if this.Period != that1.Period {
		return false
	}
	return true
}
func (this *ValidatorAccumulatedCommission) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ValidatorAccumulatedCommission)
	if !ok {
		that2, ok := that.(ValidatorAccumulatedCommission)
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
	if len(this.Commission) != len(that1.Commission) {
		return false
	}
	for i := range this.Commission {
		if !this.Commission[i].Equal(&that1.Commission[i]) {
			return false
		}
	}
	return true
}
func (this *ValidatorOutstandingRewards) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ValidatorOutstandingRewards)
	if !ok {
		that2, ok := that.(ValidatorOutstandingRewards)
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
	if len(this.Rewards) != len(that1.Rewards) {
		return false
	}
	for i := range this.Rewards {
		if !this.Rewards[i].Equal(&that1.Rewards[i]) {
			return false
		}
	}
	return true
}
func (this *ValidatorSlashEvent) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ValidatorSlashEvent)
	if !ok {
		that2, ok := that.(ValidatorSlashEvent)
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
	if this.ValidatorPeriod != that1.ValidatorPeriod {
		return false
	}
	if !this.Fraction.Equal(that1.Fraction) {
		return false
	}
	return true
}
func (this *ValidatorSlashEvents) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ValidatorSlashEvents)
	if !ok {
		that2, ok := that.(ValidatorSlashEvents)
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
	if len(this.ValidatorSlashEvents) != len(that1.ValidatorSlashEvents) {
		return false
	}
	for i := range this.ValidatorSlashEvents {
		if !this.ValidatorSlashEvents[i].Equal(&that1.ValidatorSlashEvents[i]) {
			return false
		}
	}
	return true
}
func (this *FeePool) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FeePool)
	if !ok {
		that2, ok := that.(FeePool)
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
	if len(this.CommunityPool) != len(that1.CommunityPool) {
		return false
	}
	for i := range this.CommunityPool {
		if !this.CommunityPool[i].Equal(&that1.CommunityPool[i]) {
			return false
		}
	}
	return true
}
func (this *DelegatorStartingInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DelegatorStartingInfo)
	if !ok {
		that2, ok := that.(DelegatorStartingInfo)
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
	if this.PreviousPeriod != that1.PreviousPeriod {
		return false
	}
	if !this.Stake.Equal(that1.Stake) {
		return false
	}
	if this.Height != that1.Height {
		return false
	}
	return true
}
func (this *DelegationDelegatorReward) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DelegationDelegatorReward)
	if !ok {
		that2, ok := that.(DelegationDelegatorReward)
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
	if this.ValidatorAddress != that1.ValidatorAddress {
		return false
	}
	if len(this.Reward) != len(that1.Reward) {
		return false
	}
	for i := range this.Reward {
		if !this.Reward[i].Equal(&that1.Reward[i]) {
			return false
		}
	}
	return true
}
func (this *CommunityPoolSpendProposalWithDeposit) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CommunityPoolSpendProposalWithDeposit)
	if !ok {
		that2, ok := that.(CommunityPoolSpendProposalWithDeposit)
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
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.Recipient != that1.Recipient {
		return false
	}
	if this.Amount != that1.Amount {
		return false
	}
	if this.Deposit != that1.Deposit {
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
	if m.WithdrawAddrEnabled {
		i--
		if m.WithdrawAddrEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	{
		size := m.BonusProposerReward.Size()
		i -= size
		if _, err := m.BonusProposerReward.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDistribution(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.BaseProposerReward.Size()
		i -= size
		if _, err := m.BaseProposerReward.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDistribution(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.CommunityTax.Size()
		i -= size
		if _, err := m.CommunityTax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDistribution(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ValidatorHistoricalRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorHistoricalRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorHistoricalRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ReferenceCount != 0 {
		i = encodeVarintDistribution(dAtA, i, uint64(m.ReferenceCount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.CumulativeRewardRatio) > 0 {
		for iNdEx := len(m.CumulativeRewardRatio) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CumulativeRewardRatio[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDistribution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorCurrentRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorCurrentRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorCurrentRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Period != 0 {
		i = encodeVarintDistribution(dAtA, i, uint64(m.Period))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDistribution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorAccumulatedCommission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorAccumulatedCommission) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorAccumulatedCommission) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Commission) > 0 {
		for iNdEx := len(m.Commission) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Commission[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDistribution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorOutstandingRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorOutstandingRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorOutstandingRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDistribution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorSlashEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSlashEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSlashEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Fraction.Size()
		i -= size
		if _, err := m.Fraction.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDistribution(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.ValidatorPeriod != 0 {
		i = encodeVarintDistribution(dAtA, i, uint64(m.ValidatorPeriod))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorSlashEvents) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSlashEvents) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSlashEvents) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintDistribution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *FeePool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeePool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeePool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CommunityPool) > 0 {
		for iNdEx := len(m.CommunityPool) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CommunityPool[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDistribution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CommunityPoolSpendProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommunityPoolSpendProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommunityPoolSpendProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintDistribution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintDistribution(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintDistribution(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintDistribution(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DelegatorStartingInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegatorStartingInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelegatorStartingInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintDistribution(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.Stake.Size()
		i -= size
		if _, err := m.Stake.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDistribution(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.PreviousPeriod != 0 {
		i = encodeVarintDistribution(dAtA, i, uint64(m.PreviousPeriod))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DelegationDelegatorReward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegationDelegatorReward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelegationDelegatorReward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Reward) > 0 {
		for iNdEx := len(m.Reward) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Reward[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDistribution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintDistribution(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CommunityPoolSpendProposalWithDeposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommunityPoolSpendProposalWithDeposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommunityPoolSpendProposalWithDeposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Deposit) > 0 {
		i -= len(m.Deposit)
		copy(dAtA[i:], m.Deposit)
		i = encodeVarintDistribution(dAtA, i, uint64(len(m.Deposit)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintDistribution(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintDistribution(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintDistribution(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintDistribution(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDistribution(dAtA []byte, offset int, v uint64) int {
	offset -= sovDistribution(v)
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
	l = m.CommunityTax.Size()
	n += 1 + l + sovDistribution(uint64(l))
	l = m.BaseProposerReward.Size()
	n += 1 + l + sovDistribution(uint64(l))
	l = m.BonusProposerReward.Size()
	n += 1 + l + sovDistribution(uint64(l))
	if m.WithdrawAddrEnabled {
		n += 2
	}
	return n
}

func (m *ValidatorHistoricalRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CumulativeRewardRatio) > 0 {
		for _, e := range m.CumulativeRewardRatio {
			l = e.Size()
			n += 1 + l + sovDistribution(uint64(l))
		}
	}
	if m.ReferenceCount != 0 {
		n += 1 + sovDistribution(uint64(m.ReferenceCount))
	}
	return n
}

func (m *ValidatorCurrentRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovDistribution(uint64(l))
		}
	}
	if m.Period != 0 {
		n += 1 + sovDistribution(uint64(m.Period))
	}
	return n
}

func (m *ValidatorAccumulatedCommission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Commission) > 0 {
		for _, e := range m.Commission {
			l = e.Size()
			n += 1 + l + sovDistribution(uint64(l))
		}
	}
	return n
}

func (m *ValidatorOutstandingRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovDistribution(uint64(l))
		}
	}
	return n
}

func (m *ValidatorSlashEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ValidatorPeriod != 0 {
		n += 1 + sovDistribution(uint64(m.ValidatorPeriod))
	}
	l = m.Fraction.Size()
	n += 1 + l + sovDistribution(uint64(l))
	return n
}

func (m *ValidatorSlashEvents) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ValidatorSlashEvents) > 0 {
		for _, e := range m.ValidatorSlashEvents {
			l = e.Size()
			n += 1 + l + sovDistribution(uint64(l))
		}
	}
	return n
}

func (m *FeePool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CommunityPool) > 0 {
		for _, e := range m.CommunityPool {
			l = e.Size()
			n += 1 + l + sovDistribution(uint64(l))
		}
	}
	return n
}

func (m *CommunityPoolSpendProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovDistribution(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovDistribution(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovDistribution(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovDistribution(uint64(l))
		}
	}
	return n
}

func (m *DelegatorStartingInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PreviousPeriod != 0 {
		n += 1 + sovDistribution(uint64(m.PreviousPeriod))
	}
	l = m.Stake.Size()
	n += 1 + l + sovDistribution(uint64(l))
	if m.Height != 0 {
		n += 1 + sovDistribution(uint64(m.Height))
	}
	return n
}

func (m *DelegationDelegatorReward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovDistribution(uint64(l))
	}
	if len(m.Reward) > 0 {
		for _, e := range m.Reward {
			l = e.Size()
			n += 1 + l + sovDistribution(uint64(l))
		}
	}
	return n
}

func (m *CommunityPoolSpendProposalWithDeposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovDistribution(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovDistribution(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovDistribution(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovDistribution(uint64(l))
	}
	l = len(m.Deposit)
	if l > 0 {
		n += 1 + l + sovDistribution(uint64(l))
	}
	return n
}

func sovDistribution(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDistribution(x uint64) (n int) {
	return sovDistribution(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDistribution
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommunityTax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommunityTax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseProposerReward", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseProposerReward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BonusProposerReward", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BonusProposerReward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawAddrEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
			m.WithdrawAddrEnabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipDistribution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDistribution
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
func (m *ValidatorHistoricalRewards) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDistribution
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
			return fmt.Errorf("proto: ValidatorHistoricalRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorHistoricalRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CumulativeRewardRatio", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CumulativeRewardRatio = append(m.CumulativeRewardRatio, types.DecCoin{})
			if err := m.CumulativeRewardRatio[len(m.CumulativeRewardRatio)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReferenceCount", wireType)
			}
			m.ReferenceCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReferenceCount |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDistribution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDistribution
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
func (m *ValidatorCurrentRewards) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDistribution
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
			return fmt.Errorf("proto: ValidatorCurrentRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorCurrentRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rewards = append(m.Rewards, types.DecCoin{})
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			m.Period = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
		default:
			iNdEx = preIndex
			skippy, err := skipDistribution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDistribution
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
func (m *ValidatorAccumulatedCommission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDistribution
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
			return fmt.Errorf("proto: ValidatorAccumulatedCommission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorAccumulatedCommission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commission", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Commission = append(m.Commission, types.DecCoin{})
			if err := m.Commission[len(m.Commission)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDistribution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDistribution
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
func (m *ValidatorOutstandingRewards) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDistribution
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
			return fmt.Errorf("proto: ValidatorOutstandingRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorOutstandingRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rewards = append(m.Rewards, types.DecCoin{})
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDistribution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDistribution
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
func (m *ValidatorSlashEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDistribution
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
			return fmt.Errorf("proto: ValidatorSlashEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSlashEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorPeriod", wireType)
			}
			m.ValidatorPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidatorPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fraction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fraction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDistribution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDistribution
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
func (m *ValidatorSlashEvents) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDistribution
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
			return fmt.Errorf("proto: ValidatorSlashEvents: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSlashEvents: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorSlashEvents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorSlashEvents = append(m.ValidatorSlashEvents, ValidatorSlashEvent{})
			if err := m.ValidatorSlashEvents[len(m.ValidatorSlashEvents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDistribution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDistribution
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
func (m *FeePool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDistribution
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
			return fmt.Errorf("proto: FeePool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeePool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommunityPool", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommunityPool = append(m.CommunityPool, types.DecCoin{})
			if err := m.CommunityPool[len(m.CommunityPool)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDistribution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDistribution
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
func (m *CommunityPoolSpendProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDistribution
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
			return fmt.Errorf("proto: CommunityPoolSpendProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommunityPoolSpendProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDistribution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDistribution
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
func (m *DelegatorStartingInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDistribution
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
			return fmt.Errorf("proto: DelegatorStartingInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegatorStartingInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousPeriod", wireType)
			}
			m.PreviousPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PreviousPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stake", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stake.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
		default:
			iNdEx = preIndex
			skippy, err := skipDistribution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDistribution
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
func (m *DelegationDelegatorReward) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDistribution
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
			return fmt.Errorf("proto: DelegationDelegatorReward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegationDelegatorReward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reward", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reward = append(m.Reward, types.DecCoin{})
			if err := m.Reward[len(m.Reward)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDistribution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDistribution
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
func (m *CommunityPoolSpendProposalWithDeposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDistribution
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
			return fmt.Errorf("proto: CommunityPoolSpendProposalWithDeposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommunityPoolSpendProposalWithDeposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deposit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDistribution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDistribution
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
func skipDistribution(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDistribution
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
					return 0, ErrIntOverflowDistribution
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
					return 0, ErrIntOverflowDistribution
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
				return 0, ErrInvalidLengthDistribution
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDistribution
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDistribution
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDistribution        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDistribution          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDistribution = fmt.Errorf("proto: unexpected end of group")
)
