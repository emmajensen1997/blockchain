


package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)


var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf





const _ = proto.GoGoProtoPackageIsVersion3


type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5efd02cbc06efdc9, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo


type QueryParamsResponse struct {

	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5efd02cbc06efdc9, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}



type QueryValidatorOutstandingRewardsRequest struct {

	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
}

func (m *QueryValidatorOutstandingRewardsRequest) Reset() {
	*m = QueryValidatorOutstandingRewardsRequest{}
}
func (m *QueryValidatorOutstandingRewardsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryValidatorOutstandingRewardsRequest) ProtoMessage()    {}
func (*QueryValidatorOutstandingRewardsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5efd02cbc06efdc9, []int{2}
}
func (m *QueryValidatorOutstandingRewardsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorOutstandingRewardsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorOutstandingRewardsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorOutstandingRewardsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorOutstandingRewardsRequest.Merge(m, src)
}
func (m *QueryValidatorOutstandingRewardsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorOutstandingRewardsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorOutstandingRewardsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorOutstandingRewardsRequest proto.InternalMessageInfo

func (m *QueryValidatorOutstandingRewardsRequest) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}



type QueryValidatorOutstandingRewardsResponse struct {
	Rewards ValidatorOutstandingRewards `protobuf:"bytes,1,opt,name=rewards,proto3" json:"rewards"`
}

func (m *QueryValidatorOutstandingRewardsResponse) Reset() {
	*m = QueryValidatorOutstandingRewardsResponse{}
}
func (m *QueryValidatorOutstandingRewardsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryValidatorOutstandingRewardsResponse) ProtoMessage()    {}
func (*QueryValidatorOutstandingRewardsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5efd02cbc06efdc9, []int{3}
}
func (m *QueryValidatorOutstandingRewardsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorOutstandingRewardsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorOutstandingRewardsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorOutstandingRewardsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorOutstandingRewardsResponse.Merge(m, src)
}
func (m *QueryValidatorOutstandingRewardsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorOutstandingRewardsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorOutstandingRewardsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorOutstandingRewardsResponse proto.InternalMessageInfo

func (m *QueryValidatorOutstandingRewardsResponse) GetRewards() ValidatorOutstandingRewards {
	if m != nil {
		return m.Rewards
	}
	return ValidatorOutstandingRewards{}
}



type QueryValidatorCommissionRequest struct {

	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
}

func (m *QueryValidatorCommissionRequest) Reset()         { *m = QueryValidatorCommissionRequest{} }
func (m *QueryValidatorCommissionRequest) String() string { return proto.CompactTextString(m) }
func (*QueryValidatorCommissionRequest) ProtoMessage()    {}
func (*QueryValidatorCommissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5efd02cbc06efdc9, []int{4}
}
func (m *QueryValidatorCommissionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorCommissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorCommissionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorCommissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorCommissionRequest.Merge(m, src)
}
func (m *QueryValidatorCommissionRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorCommissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorCommissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorCommissionRequest proto.InternalMessageInfo

func (m *QueryValidatorCommissionRequest) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}



type QueryValidatorCommissionResponse struct {

	Commission ValidatorAccumulatedCommission `protobuf:"bytes,1,opt,name=commission,proto3" json:"commission"`
}

func (m *QueryValidatorCommissionResponse) Reset()         { *m = QueryValidatorCommissionResponse{} }
func (m *QueryValidatorCommissionResponse) String() string { return proto.CompactTextString(m) }
func (*QueryValidatorCommissionResponse) ProtoMessage()    {}
func (*QueryValidatorCommissionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5efd02cbc06efdc9, []int{5}
}
func (m *QueryValidatorCommissionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorCommissionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorCommissionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorCommissionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorCommissionResponse.Merge(m, src)
}
func (m *QueryValidatorCommissionResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorCommissionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorCommissionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorCommissionResponse proto.InternalMessageInfo

func (m *QueryValidatorCommissionResponse) GetCommission() ValidatorAccumulatedCommission {
	if m != nil {
		return m.Commission
	}
	return ValidatorAccumulatedCommission{}
}



type QueryValidatorSlashesRequest struct {

	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`

	StartingHeight uint64 `protobuf:"varint,2,opt,name=starting_height,json=startingHeight,proto3" json:"starting_height,omitempty"`

	EndingHeight uint64 `protobuf:"varint,3,opt,name=ending_height,json=endingHeight,proto3" json:"ending_height,omitempty"`

	Pagination *query.PageRequest `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryValidatorSlashesRequest) Reset()         { *m = QueryValidatorSlashesRequest{} }
func (m *QueryValidatorSlashesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryValidatorSlashesRequest) ProtoMessage()    {}
func (*QueryValidatorSlashesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5efd02cbc06efdc9, []int{6}
}
func (m *QueryValidatorSlashesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorSlashesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorSlashesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorSlashesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorSlashesRequest.Merge(m, src)
}
func (m *QueryValidatorSlashesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorSlashesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorSlashesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorSlashesRequest proto.InternalMessageInfo



type QueryValidatorSlashesResponse struct {

	Slashes []ValidatorSlashEvent `protobuf:"bytes,1,rep,name=slashes,proto3" json:"slashes"`

	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryValidatorSlashesResponse) Reset()         { *m = QueryValidatorSlashesResponse{} }
func (m *QueryValidatorSlashesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryValidatorSlashesResponse) ProtoMessage()    {}
func (*QueryValidatorSlashesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5efd02cbc06efdc9, []int{7}
}
func (m *QueryValidatorSlashesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorSlashesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorSlashesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorSlashesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorSlashesResponse.Merge(m, src)
}
func (m *QueryValidatorSlashesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorSlashesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorSlashesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorSlashesResponse proto.InternalMessageInfo

func (m *QueryValidatorSlashesResponse) GetSlashes() []ValidatorSlashEvent {
	if m != nil {
		return m.Slashes
	}
	return nil
}

func (m *QueryValidatorSlashesResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}



type QueryDelegationRewardsRequest struct {

	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`

	ValidatorAddress string `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
}

func (m *QueryDelegationRewardsRequest) Reset()         { *m = QueryDelegationRewardsRequest{} }
func (m *QueryDelegationRewardsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDelegationRewardsRequest) ProtoMessage()    {}
func (*QueryDelegationRewardsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5efd02cbc06efdc9, []int{8}
}
func (m *QueryDelegationRewardsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDelegationRewardsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDelegationRewardsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDelegationRewardsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDelegationRewardsRequest.Merge(m, src)
}
func (m *QueryDelegationRewardsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDelegationRewardsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDelegationRewardsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDelegationRewardsRequest proto.InternalMessageInfo



type QueryDelegationRewardsResponse struct {

	Rewards github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=rewards,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"rewards"`
}

func (m *QueryDelegationRewardsResponse) Reset()         { *m = QueryDelegationRewardsResponse{} }
func (m *QueryDelegationRewardsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDelegationRewardsResponse) ProtoMessage()    {}
func (*QueryDelegationRewardsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5efd02cbc06efdc9, []int{9}
}
func (m *QueryDelegationRewardsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDelegationRewardsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDelegationRewardsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDelegationRewardsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDelegationRewardsResponse.Merge(m, src)
}
func (m *QueryDelegationRewardsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDelegationRewardsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDelegationRewardsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDelegationRewardsResponse proto.InternalMessageInfo

func (m *QueryDelegationRewardsResponse) GetRewards() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Rewards
	}
	return nil
}



type QueryDelegationTotalRewardsRequest struct {

	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
}

func (m *QueryDelegationTotalRewardsRequest) Reset()         { *m = QueryDelegationTotalRewardsRequest{} }
func (m *QueryDelegationTotalRewardsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDelegationTotalRewardsRequest) ProtoMessage()    {}
func (*QueryDelegationTotalRewardsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5efd02cbc06efdc9, []int{10}
}
func (m *QueryDelegationTotalRewardsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDelegationTotalRewardsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDelegationTotalRewardsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDelegationTotalRewardsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDelegationTotalRewardsRequest.Merge(m, src)
}
func (m *QueryDelegationTotalRewardsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDelegationTotalRewardsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDelegationTotalRewardsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDelegationTotalRewardsRequest proto.InternalMessageInfo



type QueryDelegationTotalRewardsResponse struct {

	Rewards []DelegationDelegatorReward `protobuf:"bytes,1,rep,name=rewards,proto3" json:"rewards"`

	Total github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,2,rep,name=total,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"total"`
}

func (m *QueryDelegationTotalRewardsResponse) Reset()         { *m = QueryDelegationTotalRewardsResponse{} }
func (m *QueryDelegationTotalRewardsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDelegationTotalRewardsResponse) ProtoMessage()    {}
func (*QueryDelegationTotalRewardsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5efd02cbc06efdc9, []int{11}
}
func (m *QueryDelegationTotalRewardsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDelegationTotalRewardsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDelegationTotalRewardsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDelegationTotalRewardsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDelegationTotalRewardsResponse.Merge(m, src)
}
func (m *QueryDelegationTotalRewardsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDelegationTotalRewardsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDelegationTotalRewardsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDelegationTotalRewardsResponse proto.InternalMessageInfo

func (m *QueryDelegationTotalRewardsResponse) GetRewards() []DelegationDelegatorReward {
	if m != nil {
		return m.Rewards
	}
	return nil
}

func (m *QueryDelegationTotalRewardsResponse) GetTotal() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Total
	}
	return nil
}



type QueryDelegatorValidatorsRequest struct {

	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
}

func (m *QueryDelegatorValidatorsRequest) Reset()         { *m = QueryDelegatorValidatorsRequest{} }
func (m *QueryDelegatorValidatorsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDelegatorValidatorsRequest) ProtoMessage()    {}
func (*QueryDelegatorValidatorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5efd02cbc06efdc9, []int{12}
}
func (m *QueryDelegatorValidatorsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDelegatorValidatorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDelegatorValidatorsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDelegatorValidatorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDelegatorValidatorsRequest.Merge(m, src)
}
func (m *QueryDelegatorValidatorsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDelegatorValidatorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDelegatorValidatorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDelegatorValidatorsRequest proto.InternalMessageInfo



type QueryDelegatorValidatorsResponse struct {

	Validators []string `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators,omitempty"`
}

func (m *QueryDelegatorValidatorsResponse) Reset()         { *m = QueryDelegatorValidatorsResponse{} }
func (m *QueryDelegatorValidatorsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDelegatorValidatorsResponse) ProtoMessage()    {}
func (*QueryDelegatorValidatorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5efd02cbc06efdc9, []int{13}
}
func (m *QueryDelegatorValidatorsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDelegatorValidatorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDelegatorValidatorsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDelegatorValidatorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDelegatorValidatorsResponse.Merge(m, src)
}
func (m *QueryDelegatorValidatorsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDelegatorValidatorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDelegatorValidatorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDelegatorValidatorsResponse proto.InternalMessageInfo



type QueryDelegatorWithdrawAddressRequest struct {

	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
}

func (m *QueryDelegatorWithdrawAddressRequest) Reset()         { *m = QueryDelegatorWithdrawAddressRequest{} }
func (m *QueryDelegatorWithdrawAddressRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDelegatorWithdrawAddressRequest) ProtoMessage()    {}
func (*QueryDelegatorWithdrawAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5efd02cbc06efdc9, []int{14}
}
func (m *QueryDelegatorWithdrawAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDelegatorWithdrawAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDelegatorWithdrawAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDelegatorWithdrawAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDelegatorWithdrawAddressRequest.Merge(m, src)
}
func (m *QueryDelegatorWithdrawAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDelegatorWithdrawAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDelegatorWithdrawAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDelegatorWithdrawAddressRequest proto.InternalMessageInfo



type QueryDelegatorWithdrawAddressResponse struct {

	WithdrawAddress string `protobuf:"bytes,1,opt,name=withdraw_address,json=withdrawAddress,proto3" json:"withdraw_address,omitempty"`
}

func (m *QueryDelegatorWithdrawAddressResponse) Reset()         { *m = QueryDelegatorWithdrawAddressResponse{} }
func (m *QueryDelegatorWithdrawAddressResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDelegatorWithdrawAddressResponse) ProtoMessage()    {}
func (*QueryDelegatorWithdrawAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5efd02cbc06efdc9, []int{15}
}
func (m *QueryDelegatorWithdrawAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDelegatorWithdrawAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDelegatorWithdrawAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDelegatorWithdrawAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDelegatorWithdrawAddressResponse.Merge(m, src)
}
func (m *QueryDelegatorWithdrawAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDelegatorWithdrawAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDelegatorWithdrawAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDelegatorWithdrawAddressResponse proto.InternalMessageInfo



type QueryCommunityPoolRequest struct {
}

func (m *QueryCommunityPoolRequest) Reset()         { *m = QueryCommunityPoolRequest{} }
func (m *QueryCommunityPoolRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCommunityPoolRequest) ProtoMessage()    {}
func (*QueryCommunityPoolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5efd02cbc06efdc9, []int{16}
}
func (m *QueryCommunityPoolRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCommunityPoolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCommunityPoolRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCommunityPoolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCommunityPoolRequest.Merge(m, src)
}
func (m *QueryCommunityPoolRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCommunityPoolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCommunityPoolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCommunityPoolRequest proto.InternalMessageInfo



type QueryCommunityPoolResponse struct {

	Pool github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=pool,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"pool"`
}

func (m *QueryCommunityPoolResponse) Reset()         { *m = QueryCommunityPoolResponse{} }
func (m *QueryCommunityPoolResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCommunityPoolResponse) ProtoMessage()    {}
func (*QueryCommunityPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5efd02cbc06efdc9, []int{17}
}
func (m *QueryCommunityPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCommunityPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCommunityPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCommunityPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCommunityPoolResponse.Merge(m, src)
}
func (m *QueryCommunityPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCommunityPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCommunityPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCommunityPoolResponse proto.InternalMessageInfo

func (m *QueryCommunityPoolResponse) GetPool() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Pool
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "cosmos.distribution.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "cosmos.distribution.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QueryValidatorOutstandingRewardsRequest)(nil), "cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest")
	proto.RegisterType((*QueryValidatorOutstandingRewardsResponse)(nil), "cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse")
	proto.RegisterType((*QueryValidatorCommissionRequest)(nil), "cosmos.distribution.v1beta1.QueryValidatorCommissionRequest")
	proto.RegisterType((*QueryValidatorCommissionResponse)(nil), "cosmos.distribution.v1beta1.QueryValidatorCommissionResponse")
	proto.RegisterType((*QueryValidatorSlashesRequest)(nil), "cosmos.distribution.v1beta1.QueryValidatorSlashesRequest")
	proto.RegisterType((*QueryValidatorSlashesResponse)(nil), "cosmos.distribution.v1beta1.QueryValidatorSlashesResponse")
	proto.RegisterType((*QueryDelegationRewardsRequest)(nil), "cosmos.distribution.v1beta1.QueryDelegationRewardsRequest")
	proto.RegisterType((*QueryDelegationRewardsResponse)(nil), "cosmos.distribution.v1beta1.QueryDelegationRewardsResponse")
	proto.RegisterType((*QueryDelegationTotalRewardsRequest)(nil), "cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest")
	proto.RegisterType((*QueryDelegationTotalRewardsResponse)(nil), "cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse")
	proto.RegisterType((*QueryDelegatorValidatorsRequest)(nil), "cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest")
	proto.RegisterType((*QueryDelegatorValidatorsResponse)(nil), "cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse")
	proto.RegisterType((*QueryDelegatorWithdrawAddressRequest)(nil), "cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest")
	proto.RegisterType((*QueryDelegatorWithdrawAddressResponse)(nil), "cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse")
	proto.RegisterType((*QueryCommunityPoolRequest)(nil), "cosmos.distribution.v1beta1.QueryCommunityPoolRequest")
	proto.RegisterType((*QueryCommunityPoolResponse)(nil), "cosmos.distribution.v1beta1.QueryCommunityPoolResponse")
}

func init() {
	proto.RegisterFile("cosmos/distribution/v1beta1/query.proto", fileDescriptor_5efd02cbc06efdc9)
}

var fileDescriptor_5efd02cbc06efdc9 = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x98, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xc7, 0x3d, 0x6e, 0xda, 0xd2, 0x57, 0x4a, 0xd2, 0x69, 0x85, 0xcc, 0x26, 0xd8, 0xd1, 0x86,
	0x92, 0x40, 0x54, 0x6f, 0x93, 0x48, 0x05, 0x5a, 0x10, 0xe4, 0x57, 0xa9, 0xd4, 0x2a, 0x4d, 0x4d,
	0x95, 0x84, 0x5f, 0x8a, 0x26, 0xde, 0xd1, 0x7a, 0x55, 0x7b, 0xc7, 0xdd, 0x19, 0x27, 0x44, 0x55,
	0x2f, 0x04, 0x24, 0x2e, 0x48, 0x48, 0x5c, 0x7a, 0xcc, 0x99, 0x3b, 0x17, 0xfe, 0x82, 0x1e, 0x2b,
	0x21, 0xa1, 0x9e, 0x00, 0x25, 0x08, 0x55, 0x42, 0x9c, 0xb9, 0x22, 0xcf, 0xcc, 0xda, 0xbb, 0xf6,
	0x7a, 0xfd, 0x4b, 0x3d, 0xc5, 0x7a, 0x3b, 0xef, 0x3b, 0xef, 0xf3, 0x66, 0xde, 0xbc, 0xa7, 0xc0,
	0x74, 0x91, 0xf1, 0x0a, 0xe3, 0x96, 0xed, 0x72, 0xe1, 0xbb, 0x3b, 0x35, 0xe1, 0x32, 0xcf, 0xda,
	0x9d, 0xdb, 0xa1, 0x82, 0xcc, 0x59, 0x0f, 0x6a, 0xd4, 0xdf, 0xcf, 0x57, 0x7d, 0x26, 0x18, 0x1e,
	0x57, 0x0b, 0xf3, 0xe1, 0x85, 0x79, 0xbd, 0xd0, 0x78, 0x5b, 0xab, 0xec, 0x10, 0x4e, 0x95, 0x57,
	0x43, 0xa3, 0x4a, 0x1c, 0xd7, 0x23, 0x72, 0xb5, 0x14, 0x32, 0x2e, 0x3a, 0xcc, 0x61, 0xf2, 0xa7,
	0x55, 0xff, 0xa5, 0xad, 0x13, 0x0e, 0x63, 0x4e, 0x99, 0x5a, 0xa4, 0xea, 0x5a, 0xc4, 0xf3, 0x98,
	0x90, 0x2e, 0x5c, 0x7f, 0xcd, 0x86, 0xf5, 0x03, 0xe5, 0x22, 0x73, 0x03, 0xcd, 0x7c, 0x12, 0x45,
	0x24, 0x62, 0xb9, 0xde, 0xbc, 0x08, 0xf8, 0x6e, 0x3d, 0xca, 0x75, 0xe2, 0x93, 0x0a, 0x2f, 0xd0,
	0x07, 0x35, 0xca, 0x85, 0xb9, 0x05, 0x17, 0x22, 0x56, 0x5e, 0x65, 0x1e, 0xa7, 0x78, 0x11, 0x4e,
	0x55, 0xa5, 0x25, 0x83, 0x26, 0xd1, 0xcc, 0xd9, 0xf9, 0xa9, 0x7c, 0x42, 0x2a, 0xf2, 0xca, 0x79,
	0x69, 0xe4, 0xc9, 0xef, 0xb9, 0x54, 0x41, 0x3b, 0x9a, 0x1b, 0x30, 0x2d, 0x95, 0x37, 0x48, 0xd9,
	0xb5, 0x89, 0x60, 0xfe, 0x9d, 0x9a, 0xe0, 0x82, 0x78, 0xb6, 0xeb, 0x39, 0x05, 0xba, 0x47, 0x7c,
	0x3b, 0x08, 0x02, 0xcf, 0xc2, 0xf9, 0xdd, 0x60, 0xd5, 0x36, 0xb1, 0x6d, 0x9f, 0x72, 0xb5, 0xf1,
	0x99, 0xc2, 0x58, 0xe3, 0xc3, 0xa2, 0xb2, 0x9b, 0xdf, 0x20, 0x98, 0xe9, 0x2e, 0xac, 0x39, 0xb6,
	0xe0, 0xb4, 0xaf, 0x4c, 0x1a, 0xe4, 0xdd, 0x44, 0x90, 0x04, 0x49, 0x4d, 0x17, 0xc8, 0x99, 0x6b,
	0x90, 0x8b, 0x46, 0xb1, 0xcc, 0x2a, 0x15, 0x97, 0x73, 0x97, 0x79, 0x03, 0x61, 0x7d, 0x8b, 0x60,
	0xb2, 0xb3, 0xa0, 0xc6, 0x21, 0x00, 0xc5, 0x86, 0x55, 0x13, 0x5d, 0xef, 0x8d, 0x68, 0xb1, 0x58,
	0xac, 0x55, 0x6a, 0x65, 0x22, 0xa8, 0xdd, 0x14, 0xd6, 0x50, 0x21, 0x51, 0xf3, 0x1f, 0x04, 0x13,
	0xd1, 0x38, 0x3e, 0x29, 0x13, 0x5e, 0xa2, 0x03, 0x1d, 0x16, 0x9e, 0x86, 0x51, 0x2e, 0x88, 0x2f,
	0x5c, 0xcf, 0xd9, 0x2e, 0x51, 0xd7, 0x29, 0x89, 0x4c, 0x7a, 0x12, 0xcd, 0x8c, 0x14, 0x5e, 0x09,
	0xcc, 0x37, 0xa5, 0x15, 0x4f, 0xc1, 0x39, 0x2a, 0xd3, 0x1d, 0x2c, 0x3b, 0x21, 0x97, 0xbd, 0xac,
	0x8c, 0x7a, 0xd1, 0x0d, 0x80, 0x66, 0x69, 0x65, 0x46, 0x24, 0xfe, 0x9b, 0x01, 0x7e, 0xbd, 0x4e,
	0xf2, 0xaa, 0x7a, 0x9b, 0xf7, 0xd2, 0xa1, 0x3a, 0xec, 0x42, 0xc8, 0xf3, 0xda, 0x4b, 0xdf, 0x1d,
	0xe6, 0x52, 0x8f, 0x0f, 0x73, 0xc8, 0xfc, 0x05, 0xc1, 0xeb, 0x1d, 0x68, 0x75, 0xca, 0xd7, 0xe1,
	0x34, 0x57, 0xa6, 0x0c, 0x9a, 0x3c, 0x31, 0x73, 0x76, 0xfe, 0x4a, 0x6f, 0xf9, 0x96, 0x3a, 0xab,
	0xbb, 0xd4, 0x13, 0xc1, 0xcd, 0xd1, 0x32, 0xf8, 0xe3, 0x08, 0x45, 0x5a, 0x52, 0x4c, 0x77, 0xa5,
	0x50, 0xe1, 0x84, 0x31, 0xcc, 0x83, 0x20, 0xf8, 0x15, 0x5a, 0xa6, 0x8e, 0xb4, 0xb5, 0x17, 0x96,
	0xad, 0xbe, 0xb5, 0x9f, 0x55, 0xe3, 0x43, 0x70, 0x56, 0xb1, 0x07, 0x9b, 0x8e, 0x3f, 0x58, 0x95,
	0xc2, 0xe7, 0x87, 0xb9, 0x94, 0xf9, 0x3d, 0x82, 0x6c, 0xa7, 0x28, 0x74, 0x0e, 0xef, 0x87, 0xab,
	0xb0, 0x9e, 0xc3, 0x89, 0x08, 0x6e, 0x00, 0xba, 0x42, 0x8b, 0xcb, 0xcc, 0xf5, 0x96, 0x16, 0xea,
	0xf9, 0xfa, 0xe9, 0x8f, 0xdc, 0xac, 0xe3, 0x8a, 0x52, 0x6d, 0x27, 0x5f, 0x64, 0x15, 0x4b, 0x3f,
	0x76, 0xea, 0xcf, 0x65, 0x6e, 0xdf, 0xb7, 0xc4, 0x7e, 0x95, 0xf2, 0xc0, 0x87, 0x37, 0x0b, 0xf3,
	0x73, 0x30, 0x5b, 0xc2, 0xb9, 0xc7, 0x04, 0x29, 0x0f, 0x91, 0x99, 0x10, 0xec, 0xdf, 0x08, 0xa6,
	0x12, 0xd5, 0x35, 0xf1, 0x46, 0x2b, 0xf1, 0xd5, 0xc4, 0x5b, 0xd3, 0x54, 0x5b, 0x09, 0xf6, 0x56,
	0x8a, 0x2d, 0xaf, 0x0e, 0x76, 0xe0, 0xa4, 0xa8, 0xef, 0x97, 0x49, 0xbf, 0xa8, 0x3c, 0x2a, 0x7d,
	0x73, 0x4b, 0x3f, 0x6f, 0x8d, 0x78, 0x1a, 0x17, 0x7b, 0xd8, 0x14, 0xde, 0xd6, 0xef, 0x5c, 0xac,
	0xb2, 0x4e, 0x5f, 0x16, 0xa0, 0x71, 0xe3, 0x54, 0x06, 0xcf, 0x14, 0x42, 0x96, 0x90, 0xda, 0x97,
	0xf0, 0x46, 0x54, 0x6d, 0xd3, 0x15, 0x25, 0xdb, 0x27, 0x7b, 0x7a, 0xe3, 0x21, 0x83, 0xfd, 0x02,
	0x2e, 0x75, 0x91, 0xd7, 0x11, 0xbf, 0x05, 0x63, 0x7b, 0xfa, 0x53, 0x8b, 0xfc, 0xe8, 0x5e, 0xd4,
	0x25, 0xa4, 0x3e, 0x0e, 0xaf, 0x49, 0xf5, 0xfa, 0x83, 0x5c, 0xf3, 0x5c, 0xb1, 0xbf, 0xce, 0x58,
	0x39, 0xe8, 0xcc, 0x07, 0x08, 0x8c, 0xb8, 0xaf, 0x7a, 0x43, 0x0a, 0x23, 0x55, 0xc6, 0xca, 0x2f,
	0xae, 0xa0, 0xa4, 0xfc, 0xfc, 0xb3, 0x51, 0x38, 0x29, 0xa3, 0xc0, 0x8f, 0x11, 0x9c, 0x52, 0x8d,
	0x1e, 0x5b, 0x89, 0x97, 0xb9, 0x7d, 0xca, 0x30, 0xae, 0xf4, 0xee, 0xa0, 0xf0, 0xcc, 0xd9, 0xaf,
	0x7f, 0xfd, 0xeb, 0xc7, 0xf4, 0x25, 0x3c, 0x65, 0x25, 0x8d, 0x39, 0x6a, 0xd4, 0xc0, 0x07, 0x69,
	0x18, 0x4f, 0x68, 0xdd, 0x78, 0xa5, 0xfb, 0xf6, 0xdd, 0xa7, 0x14, 0x63, 0x75, 0x48, 0x15, 0x4d,
	0xb6, 0x29, 0xc9, 0xee, 0xe2, 0x3b, 0x89, 0x64, 0xcd, 0xcb, 0x6e, 0x3d, 0x6c, 0x7b, 0x95, 0x1f,
	0x59, 0xac, 0xa9, 0xbf, 0x1d, 0xbc, 0x0d, 0x47, 0x08, 0x2e, 0xc4, 0x0c, 0x0f, 0xf8, 0xfd, 0x3e,
	0xe2, 0x6e, 0x1b, 0x62, 0x8c, 0x0f, 0x06, 0xf4, 0xd6, 0xb4, 0x6b, 0x92, 0xf6, 0x26, 0xbe, 0x31,
	0x0c, 0x6d, 0x73, 0x3c, 0xc1, 0xbf, 0x21, 0x18, 0x6b, 0xed, 0xd5, 0xf8, 0xbd, 0x3e, 0x62, 0x8c,
	0x4e, 0x33, 0xc6, 0xb5, 0x41, 0x5c, 0x35, 0xdb, 0x2d, 0xc9, 0xb6, 0x8a, 0x97, 0x87, 0x61, 0x0b,
	0xa6, 0x82, 0x7f, 0x11, 0x9c, 0x6f, 0xeb, 0xa0, 0xb8, 0x87, 0xf0, 0x3a, 0x35, 0x7f, 0xe3, 0xfa,
	0x40, 0xbe, 0x9a, 0x6d, 0x5b, 0xb2, 0x7d, 0x8a, 0x37, 0x13, 0xd9, 0x1a, 0x2f, 0x27, 0xb7, 0x1e,
	0xb6, 0x3d, 0xaf, 0x8f, 0x2c, 0x7d, 0x33, 0xe3, 0xb8, 0xf1, 0x73, 0x04, 0xaf, 0xc6, 0x37, 0x51,
	0xfc, 0x61, 0x3f, 0x81, 0xc7, 0x34, 0x77, 0xe3, 0xa3, 0xc1, 0x05, 0xfa, 0x3a, 0xda, 0xde, 0xf0,
	0x65, 0x61, 0xc6, 0x74, 0xbb, 0x5e, 0x0a, 0xb3, 0x73, 0xfb, 0xed, 0xa5, 0x30, 0x13, 0x5a, 0x6c,
	0x8f, 0x85, 0xd9, 0x85, 0xb0, 0x79, 0xb7, 0xf1, 0x7f, 0x08, 0x32, 0x9d, 0xba, 0x24, 0x5e, 0xec,
	0x23, 0xd6, 0xf8, 0x06, 0x6e, 0x2c, 0x0d, 0x23, 0xa1, 0x99, 0xef, 0x49, 0xe6, 0x35, 0x7c, 0x7b,
	0x18, 0xe6, 0xd6, 0x36, 0x8f, 0x7f, 0x46, 0x70, 0x2e, 0xd2, 0xa3, 0xf1, 0xd5, 0xee, 0xb1, 0xc6,
	0xb5, 0x7c, 0xe3, 0x9d, 0xbe, 0xfd, 0x34, 0xd8, 0x82, 0x04, 0xbb, 0x8c, 0x67, 0x13, 0xc1, 0x8a,
	0x81, 0xef, 0x76, 0xbd, 0xb5, 0x2f, 0xdd, 0x7a, 0x72, 0x94, 0x45, 0x4f, 0x8f, 0xb2, 0xe8, 0xcf,
	0xa3, 0x2c, 0xfa, 0xe1, 0x38, 0x9b, 0x7a, 0x7a, 0x9c, 0x4d, 0x3d, 0x3b, 0xce, 0xa6, 0x3e, 0x9b,
	0x4b, 0x9c, 0x13, 0xbe, 0x8a, 0xaa, 0xcb, 0xb1, 0x61, 0xe7, 0x94, 0xfc, 0x27, 0xc3, 0xc2, 0xff,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x68, 0x0b, 0x91, 0xae, 0x5c, 0x11, 0x00, 0x00,
}


var _ context.Context
var _ grpc.ClientConn



const _ = grpc.SupportPackageIsVersion4


//

type QueryClient interface {

	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)

	ValidatorOutstandingRewards(ctx context.Context, in *QueryValidatorOutstandingRewardsRequest, opts ...grpc.CallOption) (*QueryValidatorOutstandingRewardsResponse, error)

	ValidatorCommission(ctx context.Context, in *QueryValidatorCommissionRequest, opts ...grpc.CallOption) (*QueryValidatorCommissionResponse, error)

	ValidatorSlashes(ctx context.Context, in *QueryValidatorSlashesRequest, opts ...grpc.CallOption) (*QueryValidatorSlashesResponse, error)

	DelegationRewards(ctx context.Context, in *QueryDelegationRewardsRequest, opts ...grpc.CallOption) (*QueryDelegationRewardsResponse, error)


	DelegationTotalRewards(ctx context.Context, in *QueryDelegationTotalRewardsRequest, opts ...grpc.CallOption) (*QueryDelegationTotalRewardsResponse, error)

	DelegatorValidators(ctx context.Context, in *QueryDelegatorValidatorsRequest, opts ...grpc.CallOption) (*QueryDelegatorValidatorsResponse, error)

	DelegatorWithdrawAddress(ctx context.Context, in *QueryDelegatorWithdrawAddressRequest, opts ...grpc.CallOption) (*QueryDelegatorWithdrawAddressResponse, error)

	CommunityPool(ctx context.Context, in *QueryCommunityPoolRequest, opts ...grpc.CallOption) (*QueryCommunityPoolResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ValidatorOutstandingRewards(ctx context.Context, in *QueryValidatorOutstandingRewardsRequest, opts ...grpc.CallOption) (*QueryValidatorOutstandingRewardsResponse, error) {
	out := new(QueryValidatorOutstandingRewardsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/ValidatorOutstandingRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ValidatorCommission(ctx context.Context, in *QueryValidatorCommissionRequest, opts ...grpc.CallOption) (*QueryValidatorCommissionResponse, error) {
	out := new(QueryValidatorCommissionResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/ValidatorCommission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ValidatorSlashes(ctx context.Context, in *QueryValidatorSlashesRequest, opts ...grpc.CallOption) (*QueryValidatorSlashesResponse, error) {
	out := new(QueryValidatorSlashesResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/ValidatorSlashes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegationRewards(ctx context.Context, in *QueryDelegationRewardsRequest, opts ...grpc.CallOption) (*QueryDelegationRewardsResponse, error) {
	out := new(QueryDelegationRewardsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/DelegationRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegationTotalRewards(ctx context.Context, in *QueryDelegationTotalRewardsRequest, opts ...grpc.CallOption) (*QueryDelegationTotalRewardsResponse, error) {
	out := new(QueryDelegationTotalRewardsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/DelegationTotalRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegatorValidators(ctx context.Context, in *QueryDelegatorValidatorsRequest, opts ...grpc.CallOption) (*QueryDelegatorValidatorsResponse, error) {
	out := new(QueryDelegatorValidatorsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/DelegatorValidators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegatorWithdrawAddress(ctx context.Context, in *QueryDelegatorWithdrawAddressRequest, opts ...grpc.CallOption) (*QueryDelegatorWithdrawAddressResponse, error) {
	out := new(QueryDelegatorWithdrawAddressResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/DelegatorWithdrawAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CommunityPool(ctx context.Context, in *QueryCommunityPoolRequest, opts ...grpc.CallOption) (*QueryCommunityPoolResponse, error) {
	out := new(QueryCommunityPoolResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Query/CommunityPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


type QueryServer interface {

	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)

	ValidatorOutstandingRewards(context.Context, *QueryValidatorOutstandingRewardsRequest) (*QueryValidatorOutstandingRewardsResponse, error)

	ValidatorCommission(context.Context, *QueryValidatorCommissionRequest) (*QueryValidatorCommissionResponse, error)

	ValidatorSlashes(context.Context, *QueryValidatorSlashesRequest) (*QueryValidatorSlashesResponse, error)

	DelegationRewards(context.Context, *QueryDelegationRewardsRequest) (*QueryDelegationRewardsResponse, error)


	DelegationTotalRewards(context.Context, *QueryDelegationTotalRewardsRequest) (*QueryDelegationTotalRewardsResponse, error)

	DelegatorValidators(context.Context, *QueryDelegatorValidatorsRequest) (*QueryDelegatorValidatorsResponse, error)

	DelegatorWithdrawAddress(context.Context, *QueryDelegatorWithdrawAddressRequest) (*QueryDelegatorWithdrawAddressResponse, error)

	CommunityPool(context.Context, *QueryCommunityPoolRequest) (*QueryCommunityPoolResponse, error)
}


type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) ValidatorOutstandingRewards(ctx context.Context, req *QueryValidatorOutstandingRewardsRequest) (*QueryValidatorOutstandingRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorOutstandingRewards not implemented")
}
func (*UnimplementedQueryServer) ValidatorCommission(ctx context.Context, req *QueryValidatorCommissionRequest) (*QueryValidatorCommissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorCommission not implemented")
}
func (*UnimplementedQueryServer) ValidatorSlashes(ctx context.Context, req *QueryValidatorSlashesRequest) (*QueryValidatorSlashesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorSlashes not implemented")
}
func (*UnimplementedQueryServer) DelegationRewards(ctx context.Context, req *QueryDelegationRewardsRequest) (*QueryDelegationRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegationRewards not implemented")
}
func (*UnimplementedQueryServer) DelegationTotalRewards(ctx context.Context, req *QueryDelegationTotalRewardsRequest) (*QueryDelegationTotalRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegationTotalRewards not implemented")
}
func (*UnimplementedQueryServer) DelegatorValidators(ctx context.Context, req *QueryDelegatorValidatorsRequest) (*QueryDelegatorValidatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegatorValidators not implemented")
}
func (*UnimplementedQueryServer) DelegatorWithdrawAddress(ctx context.Context, req *QueryDelegatorWithdrawAddressRequest) (*QueryDelegatorWithdrawAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegatorWithdrawAddress not implemented")
}
func (*UnimplementedQueryServer) CommunityPool(ctx context.Context, req *QueryCommunityPoolRequest) (*QueryCommunityPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommunityPool not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ValidatorOutstandingRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorOutstandingRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorOutstandingRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/ValidatorOutstandingRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorOutstandingRewards(ctx, req.(*QueryValidatorOutstandingRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ValidatorCommission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorCommissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorCommission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/ValidatorCommission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorCommission(ctx, req.(*QueryValidatorCommissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ValidatorSlashes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorSlashesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorSlashes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/ValidatorSlashes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorSlashes(ctx, req.(*QueryValidatorSlashesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegationRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegationRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegationRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/DelegationRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegationRewards(ctx, req.(*QueryDelegationRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegationTotalRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegationTotalRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegationTotalRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/DelegationTotalRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegationTotalRewards(ctx, req.(*QueryDelegationTotalRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegatorValidators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatorValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegatorValidators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/DelegatorValidators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatorValidators(ctx, req.(*QueryDelegatorValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegatorWithdrawAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatorWithdrawAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegatorWithdrawAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/DelegatorWithdrawAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatorWithdrawAddress(ctx, req.(*QueryDelegatorWithdrawAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CommunityPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCommunityPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CommunityPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/CommunityPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CommunityPool(ctx, req.(*QueryCommunityPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.distribution.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ValidatorOutstandingRewards",
			Handler:    _Query_ValidatorOutstandingRewards_Handler,
		},
		{
			MethodName: "ValidatorCommission",
			Handler:    _Query_ValidatorCommission_Handler,
		},
		{
			MethodName: "ValidatorSlashes",
			Handler:    _Query_ValidatorSlashes_Handler,
		},
		{
			MethodName: "DelegationRewards",
			Handler:    _Query_DelegationRewards_Handler,
		},
		{
			MethodName: "DelegationTotalRewards",
			Handler:    _Query_DelegationTotalRewards_Handler,
		},
		{
			MethodName: "DelegatorValidators",
			Handler:    _Query_DelegatorValidators_Handler,
		},
		{
			MethodName: "DelegatorWithdrawAddress",
			Handler:    _Query_DelegatorWithdrawAddress_Handler,
		},
		{
			MethodName: "CommunityPool",
			Handler:    _Query_CommunityPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/distribution/v1beta1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryValidatorOutstandingRewardsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorOutstandingRewardsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorOutstandingRewardsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryValidatorOutstandingRewardsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorOutstandingRewardsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorOutstandingRewardsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryValidatorCommissionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorCommissionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorCommissionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryValidatorCommissionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorCommissionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorCommissionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Commission.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryValidatorSlashesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorSlashesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorSlashesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.EndingHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.EndingHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.StartingHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.StartingHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryValidatorSlashesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorSlashesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorSlashesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Slashes) > 0 {
		for iNdEx := len(m.Slashes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Slashes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryDelegationRewardsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDelegationRewardsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDelegationRewardsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDelegationRewardsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDelegationRewardsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDelegationRewardsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryDelegationTotalRewardsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDelegationTotalRewardsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDelegationTotalRewardsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDelegationTotalRewardsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDelegationTotalRewardsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDelegationTotalRewardsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Total) > 0 {
		for iNdEx := len(m.Total) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Total[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryDelegatorValidatorsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDelegatorValidatorsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDelegatorValidatorsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDelegatorValidatorsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDelegatorValidatorsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDelegatorValidatorsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Validators[iNdEx])
			copy(dAtA[i:], m.Validators[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Validators[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryDelegatorWithdrawAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDelegatorWithdrawAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDelegatorWithdrawAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDelegatorWithdrawAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDelegatorWithdrawAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDelegatorWithdrawAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WithdrawAddress) > 0 {
		i -= len(m.WithdrawAddress)
		copy(dAtA[i:], m.WithdrawAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.WithdrawAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCommunityPoolRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCommunityPoolRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCommunityPoolRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryCommunityPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCommunityPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCommunityPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Pool) > 0 {
		for iNdEx := len(m.Pool) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pool[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryValidatorOutstandingRewardsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryValidatorOutstandingRewardsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Rewards.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryValidatorCommissionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryValidatorCommissionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Commission.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryValidatorSlashesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.StartingHeight != 0 {
		n += 1 + sovQuery(uint64(m.StartingHeight))
	}
	if m.EndingHeight != 0 {
		n += 1 + sovQuery(uint64(m.EndingHeight))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryValidatorSlashesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Slashes) > 0 {
		for _, e := range m.Slashes {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDelegationRewardsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDelegationRewardsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryDelegationTotalRewardsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDelegationTotalRewardsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if len(m.Total) > 0 {
		for _, e := range m.Total {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryDelegatorValidatorsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDelegatorValidatorsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for _, s := range m.Validators {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryDelegatorWithdrawAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDelegatorWithdrawAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.WithdrawAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCommunityPoolRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryCommunityPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Pool) > 0 {
		for _, e := range m.Pool {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryValidatorOutstandingRewardsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryValidatorOutstandingRewardsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryValidatorOutstandingRewardsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryValidatorOutstandingRewardsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryValidatorOutstandingRewardsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryValidatorOutstandingRewardsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
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
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryValidatorCommissionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryValidatorCommissionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryValidatorCommissionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryValidatorCommissionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryValidatorCommissionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryValidatorCommissionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commission", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Commission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryValidatorSlashesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryValidatorSlashesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryValidatorSlashesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartingHeight", wireType)
			}
			m.StartingHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartingHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndingHeight", wireType)
			}
			m.EndingHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndingHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryValidatorSlashesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryValidatorSlashesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryValidatorSlashesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slashes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Slashes = append(m.Slashes, ValidatorSlashEvent{})
			if err := m.Slashes[len(m.Slashes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryDelegationRewardsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDelegationRewardsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDelegationRewardsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
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
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryDelegationRewardsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDelegationRewardsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDelegationRewardsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
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
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryDelegationTotalRewardsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDelegationTotalRewardsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDelegationTotalRewardsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryDelegationTotalRewardsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDelegationTotalRewardsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDelegationTotalRewardsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rewards = append(m.Rewards, DelegationDelegatorReward{})
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Total = append(m.Total, types.DecCoin{})
			if err := m.Total[len(m.Total)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryDelegatorValidatorsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDelegatorValidatorsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDelegatorValidatorsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryDelegatorValidatorsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDelegatorValidatorsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDelegatorValidatorsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryDelegatorWithdrawAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDelegatorWithdrawAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDelegatorWithdrawAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryDelegatorWithdrawAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDelegatorWithdrawAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDelegatorWithdrawAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCommunityPoolRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCommunityPoolRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCommunityPoolRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCommunityPoolResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCommunityPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCommunityPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pool = append(m.Pool, types.DecCoin{})
			if err := m.Pool[len(m.Pool)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
