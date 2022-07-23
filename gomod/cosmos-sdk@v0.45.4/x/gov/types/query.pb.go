


package types

import (
	context "context"
	fmt "fmt"
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


type QueryProposalRequest struct {

	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
}

func (m *QueryProposalRequest) Reset()         { *m = QueryProposalRequest{} }
func (m *QueryProposalRequest) String() string { return proto.CompactTextString(m) }
func (*QueryProposalRequest) ProtoMessage()    {}
func (*QueryProposalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35c0d133e91c0a2, []int{0}
}
func (m *QueryProposalRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProposalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProposalRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProposalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProposalRequest.Merge(m, src)
}
func (m *QueryProposalRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryProposalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProposalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProposalRequest proto.InternalMessageInfo

func (m *QueryProposalRequest) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}


type QueryProposalResponse struct {
	Proposal Proposal `protobuf:"bytes,1,opt,name=proposal,proto3" json:"proposal"`
}

func (m *QueryProposalResponse) Reset()         { *m = QueryProposalResponse{} }
func (m *QueryProposalResponse) String() string { return proto.CompactTextString(m) }
func (*QueryProposalResponse) ProtoMessage()    {}
func (*QueryProposalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35c0d133e91c0a2, []int{1}
}
func (m *QueryProposalResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProposalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProposalResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProposalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProposalResponse.Merge(m, src)
}
func (m *QueryProposalResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryProposalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProposalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProposalResponse proto.InternalMessageInfo

func (m *QueryProposalResponse) GetProposal() Proposal {
	if m != nil {
		return m.Proposal
	}
	return Proposal{}
}


type QueryProposalsRequest struct {

	ProposalStatus ProposalStatus `protobuf:"varint,1,opt,name=proposal_status,json=proposalStatus,proto3,enum=cosmos.gov.v1beta1.ProposalStatus" json:"proposal_status,omitempty"`

	Voter string `protobuf:"bytes,2,opt,name=voter,proto3" json:"voter,omitempty"`

	Depositor string `protobuf:"bytes,3,opt,name=depositor,proto3" json:"depositor,omitempty"`

	Pagination *query.PageRequest `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryProposalsRequest) Reset()         { *m = QueryProposalsRequest{} }
func (m *QueryProposalsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryProposalsRequest) ProtoMessage()    {}
func (*QueryProposalsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35c0d133e91c0a2, []int{2}
}
func (m *QueryProposalsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProposalsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProposalsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProposalsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProposalsRequest.Merge(m, src)
}
func (m *QueryProposalsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryProposalsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProposalsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProposalsRequest proto.InternalMessageInfo



type QueryProposalsResponse struct {
	Proposals []Proposal `protobuf:"bytes,1,rep,name=proposals,proto3" json:"proposals"`

	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryProposalsResponse) Reset()         { *m = QueryProposalsResponse{} }
func (m *QueryProposalsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryProposalsResponse) ProtoMessage()    {}
func (*QueryProposalsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35c0d133e91c0a2, []int{3}
}
func (m *QueryProposalsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProposalsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProposalsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProposalsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProposalsResponse.Merge(m, src)
}
func (m *QueryProposalsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryProposalsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProposalsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProposalsResponse proto.InternalMessageInfo

func (m *QueryProposalsResponse) GetProposals() []Proposal {
	if m != nil {
		return m.Proposals
	}
	return nil
}

func (m *QueryProposalsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}


type QueryVoteRequest struct {

	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`

	Voter string `protobuf:"bytes,2,opt,name=voter,proto3" json:"voter,omitempty"`
}

func (m *QueryVoteRequest) Reset()         { *m = QueryVoteRequest{} }
func (m *QueryVoteRequest) String() string { return proto.CompactTextString(m) }
func (*QueryVoteRequest) ProtoMessage()    {}
func (*QueryVoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35c0d133e91c0a2, []int{4}
}
func (m *QueryVoteRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVoteRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVoteRequest.Merge(m, src)
}
func (m *QueryVoteRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryVoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVoteRequest proto.InternalMessageInfo


type QueryVoteResponse struct {

	Vote Vote `protobuf:"bytes,1,opt,name=vote,proto3" json:"vote"`
}

func (m *QueryVoteResponse) Reset()         { *m = QueryVoteResponse{} }
func (m *QueryVoteResponse) String() string { return proto.CompactTextString(m) }
func (*QueryVoteResponse) ProtoMessage()    {}
func (*QueryVoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35c0d133e91c0a2, []int{5}
}
func (m *QueryVoteResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVoteResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVoteResponse.Merge(m, src)
}
func (m *QueryVoteResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryVoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVoteResponse proto.InternalMessageInfo

func (m *QueryVoteResponse) GetVote() Vote {
	if m != nil {
		return m.Vote
	}
	return Vote{}
}


type QueryVotesRequest struct {

	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`

	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryVotesRequest) Reset()         { *m = QueryVotesRequest{} }
func (m *QueryVotesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryVotesRequest) ProtoMessage()    {}
func (*QueryVotesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35c0d133e91c0a2, []int{6}
}
func (m *QueryVotesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVotesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVotesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVotesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVotesRequest.Merge(m, src)
}
func (m *QueryVotesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryVotesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVotesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVotesRequest proto.InternalMessageInfo

func (m *QueryVotesRequest) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}

func (m *QueryVotesRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}


type QueryVotesResponse struct {

	Votes []Vote `protobuf:"bytes,1,rep,name=votes,proto3" json:"votes"`

	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryVotesResponse) Reset()         { *m = QueryVotesResponse{} }
func (m *QueryVotesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryVotesResponse) ProtoMessage()    {}
func (*QueryVotesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35c0d133e91c0a2, []int{7}
}
func (m *QueryVotesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVotesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVotesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVotesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVotesResponse.Merge(m, src)
}
func (m *QueryVotesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryVotesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVotesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVotesResponse proto.InternalMessageInfo

func (m *QueryVotesResponse) GetVotes() []Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *QueryVotesResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}


type QueryParamsRequest struct {


	ParamsType string `protobuf:"bytes,1,opt,name=params_type,json=paramsType,proto3" json:"params_type,omitempty"`
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35c0d133e91c0a2, []int{8}
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

func (m *QueryParamsRequest) GetParamsType() string {
	if m != nil {
		return m.ParamsType
	}
	return ""
}


type QueryParamsResponse struct {

	VotingParams VotingParams `protobuf:"bytes,1,opt,name=voting_params,json=votingParams,proto3" json:"voting_params"`

	DepositParams DepositParams `protobuf:"bytes,2,opt,name=deposit_params,json=depositParams,proto3" json:"deposit_params"`

	TallyParams TallyParams `protobuf:"bytes,3,opt,name=tally_params,json=tallyParams,proto3" json:"tally_params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35c0d133e91c0a2, []int{9}
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

func (m *QueryParamsResponse) GetVotingParams() VotingParams {
	if m != nil {
		return m.VotingParams
	}
	return VotingParams{}
}

func (m *QueryParamsResponse) GetDepositParams() DepositParams {
	if m != nil {
		return m.DepositParams
	}
	return DepositParams{}
}

func (m *QueryParamsResponse) GetTallyParams() TallyParams {
	if m != nil {
		return m.TallyParams
	}
	return TallyParams{}
}


type QueryDepositRequest struct {

	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`

	Depositor string `protobuf:"bytes,2,opt,name=depositor,proto3" json:"depositor,omitempty"`
}

func (m *QueryDepositRequest) Reset()         { *m = QueryDepositRequest{} }
func (m *QueryDepositRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDepositRequest) ProtoMessage()    {}
func (*QueryDepositRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35c0d133e91c0a2, []int{10}
}
func (m *QueryDepositRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDepositRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDepositRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDepositRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDepositRequest.Merge(m, src)
}
func (m *QueryDepositRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDepositRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDepositRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDepositRequest proto.InternalMessageInfo


type QueryDepositResponse struct {

	Deposit Deposit `protobuf:"bytes,1,opt,name=deposit,proto3" json:"deposit"`
}

func (m *QueryDepositResponse) Reset()         { *m = QueryDepositResponse{} }
func (m *QueryDepositResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDepositResponse) ProtoMessage()    {}
func (*QueryDepositResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35c0d133e91c0a2, []int{11}
}
func (m *QueryDepositResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDepositResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDepositResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDepositResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDepositResponse.Merge(m, src)
}
func (m *QueryDepositResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDepositResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDepositResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDepositResponse proto.InternalMessageInfo

func (m *QueryDepositResponse) GetDeposit() Deposit {
	if m != nil {
		return m.Deposit
	}
	return Deposit{}
}


type QueryDepositsRequest struct {

	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`

	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryDepositsRequest) Reset()         { *m = QueryDepositsRequest{} }
func (m *QueryDepositsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDepositsRequest) ProtoMessage()    {}
func (*QueryDepositsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35c0d133e91c0a2, []int{12}
}
func (m *QueryDepositsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDepositsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDepositsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDepositsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDepositsRequest.Merge(m, src)
}
func (m *QueryDepositsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDepositsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDepositsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDepositsRequest proto.InternalMessageInfo

func (m *QueryDepositsRequest) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}

func (m *QueryDepositsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}


type QueryDepositsResponse struct {
	Deposits []Deposit `protobuf:"bytes,1,rep,name=deposits,proto3" json:"deposits"`

	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryDepositsResponse) Reset()         { *m = QueryDepositsResponse{} }
func (m *QueryDepositsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDepositsResponse) ProtoMessage()    {}
func (*QueryDepositsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35c0d133e91c0a2, []int{13}
}
func (m *QueryDepositsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDepositsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDepositsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDepositsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDepositsResponse.Merge(m, src)
}
func (m *QueryDepositsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDepositsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDepositsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDepositsResponse proto.InternalMessageInfo

func (m *QueryDepositsResponse) GetDeposits() []Deposit {
	if m != nil {
		return m.Deposits
	}
	return nil
}

func (m *QueryDepositsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}


type QueryTallyResultRequest struct {

	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
}

func (m *QueryTallyResultRequest) Reset()         { *m = QueryTallyResultRequest{} }
func (m *QueryTallyResultRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTallyResultRequest) ProtoMessage()    {}
func (*QueryTallyResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35c0d133e91c0a2, []int{14}
}
func (m *QueryTallyResultRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTallyResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTallyResultRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTallyResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTallyResultRequest.Merge(m, src)
}
func (m *QueryTallyResultRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryTallyResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTallyResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTallyResultRequest proto.InternalMessageInfo

func (m *QueryTallyResultRequest) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}


type QueryTallyResultResponse struct {

	Tally TallyResult `protobuf:"bytes,1,opt,name=tally,proto3" json:"tally"`
}

func (m *QueryTallyResultResponse) Reset()         { *m = QueryTallyResultResponse{} }
func (m *QueryTallyResultResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTallyResultResponse) ProtoMessage()    {}
func (*QueryTallyResultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35c0d133e91c0a2, []int{15}
}
func (m *QueryTallyResultResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTallyResultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTallyResultResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTallyResultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTallyResultResponse.Merge(m, src)
}
func (m *QueryTallyResultResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryTallyResultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTallyResultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTallyResultResponse proto.InternalMessageInfo

func (m *QueryTallyResultResponse) GetTally() TallyResult {
	if m != nil {
		return m.Tally
	}
	return TallyResult{}
}

func init() {
	proto.RegisterType((*QueryProposalRequest)(nil), "cosmos.gov.v1beta1.QueryProposalRequest")
	proto.RegisterType((*QueryProposalResponse)(nil), "cosmos.gov.v1beta1.QueryProposalResponse")
	proto.RegisterType((*QueryProposalsRequest)(nil), "cosmos.gov.v1beta1.QueryProposalsRequest")
	proto.RegisterType((*QueryProposalsResponse)(nil), "cosmos.gov.v1beta1.QueryProposalsResponse")
	proto.RegisterType((*QueryVoteRequest)(nil), "cosmos.gov.v1beta1.QueryVoteRequest")
	proto.RegisterType((*QueryVoteResponse)(nil), "cosmos.gov.v1beta1.QueryVoteResponse")
	proto.RegisterType((*QueryVotesRequest)(nil), "cosmos.gov.v1beta1.QueryVotesRequest")
	proto.RegisterType((*QueryVotesResponse)(nil), "cosmos.gov.v1beta1.QueryVotesResponse")
	proto.RegisterType((*QueryParamsRequest)(nil), "cosmos.gov.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "cosmos.gov.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QueryDepositRequest)(nil), "cosmos.gov.v1beta1.QueryDepositRequest")
	proto.RegisterType((*QueryDepositResponse)(nil), "cosmos.gov.v1beta1.QueryDepositResponse")
	proto.RegisterType((*QueryDepositsRequest)(nil), "cosmos.gov.v1beta1.QueryDepositsRequest")
	proto.RegisterType((*QueryDepositsResponse)(nil), "cosmos.gov.v1beta1.QueryDepositsResponse")
	proto.RegisterType((*QueryTallyResultRequest)(nil), "cosmos.gov.v1beta1.QueryTallyResultRequest")
	proto.RegisterType((*QueryTallyResultResponse)(nil), "cosmos.gov.v1beta1.QueryTallyResultResponse")
}

func init() { proto.RegisterFile("cosmos/gov/v1beta1/query.proto", fileDescriptor_e35c0d133e91c0a2) }

var fileDescriptor_e35c0d133e91c0a2 = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0x41, 0x6f, 0x1b, 0x55,
	0x10, 0xf6, 0x4b, 0x9c, 0xd6, 0x9e, 0xb4, 0x01, 0x86, 0x00, 0xd6, 0x12, 0xec, 0xb0, 0xa2, 0xad,
	0x49, 0xa9, 0x97, 0x24, 0x05, 0xd4, 0x16, 0x50, 0x89, 0x50, 0x5b, 0x54, 0x09, 0x15, 0xa7, 0x02,
	0x89, 0x03, 0xd1, 0xa6, 0x5e, 0x2d, 0x2b, 0x1c, 0xbf, 0xad, 0xdf, 0xb3, 0x45, 0x14, 0x22, 0x24,
	0x4e, 0x20, 0x2e, 0xa0, 0x22, 0x6e, 0x88, 0x4a, 0x95, 0xf8, 0x2d, 0x3d, 0x56, 0x82, 0x03, 0x27,
	0x84, 0x12, 0x0e, 0x88, 0xdf, 0xc0, 0x01, 0xed, 0x7b, 0xf3, 0xd6, 0xbb, 0xce, 0x3a, 0xbb, 0x29,
	0x55, 0x4f, 0xb1, 0xe7, 0x7d, 0x33, 0xf3, 0x7d, 0x33, 0xf3, 0xe6, 0xc5, 0x50, 0xbf, 0xc5, 0xc5,
	0x16, 0x17, 0x8e, 0xcf, 0x87, 0xce, 0x70, 0x79, 0xd3, 0x93, 0xee, 0xb2, 0x73, 0x7b, 0xe0, 0xf5,
	0xb7, 0x5b, 0x61, 0x9f, 0x4b, 0x8e, 0xa8, 0xcf, 0x5b, 0x3e, 0x1f, 0xb6, 0xe8, 0xdc, 0x5a, 0x22,
	0x9f, 0x4d, 0x57, 0x78, 0x1a, 0x1c, 0xbb, 0x86, 0xae, 0x1f, 0xf4, 0x5c, 0x19, 0xf0, 0x9e, 0xf6,
	0xb7, 0xe6, 0x7d, 0xee, 0x73, 0xf5, 0xd1, 0x89, 0x3e, 0x91, 0x75, 0xc1, 0xe7, 0xdc, 0xef, 0x7a,
	0x8e, 0x1b, 0x06, 0x8e, 0xdb, 0xeb, 0x71, 0xa9, 0x5c, 0x84, 0x39, 0xcd, 0xe0, 0x14, 0xe5, 0x57,
	0xa7, 0xf6, 0x1b, 0x30, 0xff, 0x41, 0x94, 0xf3, 0x46, 0x9f, 0x87, 0x5c, 0xb8, 0xdd, 0xb6, 0x77,
	0x7b, 0xe0, 0x09, 0x89, 0x0d, 0x98, 0x0d, 0xc9, 0xb4, 0x11, 0x74, 0x6a, 0x6c, 0x91, 0x35, 0xcb,
	0x6d, 0x30, 0xa6, 0xf7, 0x3a, 0xf6, 0x47, 0xf0, 0xcc, 0x98, 0xa3, 0x08, 0x79, 0x4f, 0x78, 0xf8,
	0x36, 0x54, 0x0c, 0x4c, 0xb9, 0xcd, 0xae, 0x2c, 0xb4, 0x0e, 0xca, 0x6e, 0x19, 0xbf, 0xb5, 0xf2,
	0xfd, 0x3f, 0x1a, 0xa5, 0x76, 0xec, 0x63, 0xff, 0xc3, 0xc6, 0x22, 0x0b, 0xc3, 0xe9, 0x3a, 0x3c,
	0x11, 0x73, 0x12, 0xd2, 0x95, 0x03, 0xa1, 0x12, 0xcc, 0xad, 0xd8, 0x87, 0x25, 0x58, 0x57, 0xc8,
	0xf6, 0x5c, 0x98, 0xfa, 0x8e, 0xf3, 0x30, 0x33, 0xe4, 0xd2, 0xeb, 0xd7, 0xa6, 0x16, 0x59, 0xb3,
	0xda, 0xd6, 0x5f, 0x70, 0x01, 0xaa, 0x1d, 0x2f, 0xe4, 0x22, 0x90, 0xbc, 0x5f, 0x9b, 0x56, 0x27,
	0x23, 0x03, 0x5e, 0x01, 0x18, 0xb5, 0xa4, 0x56, 0x56, 0xe2, 0x4e, 0x9b, 0xdc, 0x51, 0xff, 0x5a,
	0xba, 0xd9, 0x31, 0x05, 0xd7, 0xf7, 0x88, 0x7c, 0x3b, 0xe1, 0x79, 0xb1, 0xf2, 0xf5, 0xdd, 0x46,
	0xe9, 0xef, 0xbb, 0x8d, 0x92, 0x7d, 0x8f, 0xc1, 0xb3, 0xe3, 0x62, 0xa9, 0x8e, 0x97, 0xa1, 0x6a,
	0x28, 0x47, 0x3a, 0xa7, 0x0b, 0x16, 0x72, 0xe4, 0x84, 0x57, 0x53, 0x74, 0xa7, 0x14, 0xdd, 0x33,
	0xb9, 0x74, 0x75, 0xfa, 0x24, 0x5f, 0x7b, 0x1d, 0x9e, 0x54, 0x24, 0x3f, 0xe4, 0xd2, 0x2b, 0x3a,
	0x20, 0xd9, 0x05, 0x4e, 0x48, 0xbf, 0x0a, 0x4f, 0x25, 0x82, 0x92, 0xe8, 0x15, 0x28, 0x47, 0x38,
	0x1a, 0x9c, 0x5a, 0x96, 0xde, 0x08, 0x4f, 0x5a, 0x15, 0xd6, 0xfe, 0x22, 0x11, 0x48, 0x14, 0xa6,
	0x77, 0x25, 0xa3, 0x38, 0x0f, 0xd1, 0x4b, 0xfb, 0x0e, 0x03, 0x4c, 0xa6, 0x27, 0x21, 0xe7, 0xb5,
	0x7a, 0xd3, 0xb9, 0x3c, 0x25, 0x1a, 0xfc, 0xe8, 0x3a, 0xf6, 0x1a, 0x91, 0xba, 0xe1, 0xf6, 0xdd,
	0xad, 0x54, 0x51, 0x94, 0x61, 0x43, 0x6e, 0x87, 0xba, 0xc8, 0xd5, 0xc8, 0x2d, 0x32, 0xdd, 0xdc,
	0x0e, 0x3d, 0xfb, 0x5f, 0x06, 0x4f, 0xa7, 0xfc, 0x48, 0xcd, 0x75, 0x38, 0x39, 0xe4, 0x32, 0xe8,
	0xf9, 0x1b, 0x1a, 0x4c, 0xfd, 0x59, 0x9c, 0xa0, 0x2a, 0xe8, 0xf9, 0x3a, 0x00, 0xa9, 0x3b, 0x31,
	0x4c, 0xd8, 0xf0, 0x7d, 0x98, 0xa3, 0x2b, 0x65, 0xa2, 0x69, 0xa1, 0x2f, 0x66, 0x45, 0x7b, 0x57,
	0x23, 0x53, 0xe1, 0x4e, 0x76, 0x92, 0x46, 0xbc, 0x06, 0x27, 0xa4, 0xdb, 0xed, 0x6e, 0x9b, 0x68,
	0xd3, 0x2a, 0x5a, 0x23, 0x2b, 0xda, 0xcd, 0x08, 0x97, 0x8a, 0x35, 0x2b, 0x47, 0x26, 0xfb, 0x13,
	0x52, 0x4f, 0x49, 0x0b, 0xcf, 0x52, 0x6a, 0x6b, 0x4c, 0x8d, 0x6d, 0x8d, 0xc4, 0xc8, 0xaf, 0xd3,
	0xb2, 0x8d, 0xe3, 0x53, 0x79, 0x2f, 0xc1, 0x71, 0x82, 0x53, 0x61, 0x9f, 0x3f, 0xa4, 0x14, 0x44,
	0xdc, 0x78, 0xd8, 0x5f, 0xa6, 0x83, 0x3e, 0xfe, 0x1b, 0xf0, 0xb3, 0x59, 0xd8, 0x23, 0x06, 0xa4,
	0xeb, 0x2d, 0xa8, 0x10, 0x4b, 0x73, 0x0f, 0x0a, 0x08, 0x8b, 0x5d, 0x1e, 0xdd, 0x6d, 0xb8, 0x08,
	0xcf, 0x29, 0x82, 0xaa, 0xfd, 0x6d, 0x4f, 0x0c, 0xba, 0xf2, 0x08, 0xef, 0x5c, 0xed, 0xa0, 0x6f,
	0xdc, 0xb7, 0x19, 0x35, 0x3e, 0xd4, 0xb5, 0xc9, 0x23, 0xa7, 0xfd, 0xcc, 0x5d, 0x57, 0x3e, 0x2b,
	0xbf, 0x55, 0x61, 0x46, 0x45, 0xc6, 0x1f, 0x18, 0x54, 0xcc, 0x16, 0xc7, 0x66, 0x56, 0x90, 0xac,
	0x27, 0xda, 0x7a, 0xb9, 0x00, 0x52, 0x13, 0xb5, 0x57, 0xbf, 0xfa, 0xf5, 0xaf, 0x3b, 0x53, 0xe7,
	0xf0, 0xac, 0x93, 0xf1, 0xcf, 0x40, 0xfc, 0x60, 0x38, 0x3b, 0x89, 0x52, 0xec, 0xe2, 0x37, 0x0c,
	0xaa, 0xf1, 0xb3, 0x84, 0xf9, 0xd9, 0xcc, 0xe4, 0x59, 0x4b, 0x45, 0xa0, 0xc4, 0xec, 0x94, 0x62,
	0xd6, 0xc0, 0x17, 0x0e, 0x65, 0x86, 0x3f, 0x32, 0x28, 0x47, 0xeb, 0x12, 0x5f, 0x9a, 0x18, 0x3b,
	0xf1, 0x38, 0x59, 0xa7, 0x72, 0x50, 0x94, 0xfc, 0x1d, 0x95, 0xfc, 0x12, 0x5e, 0x38, 0x42, 0x59,
	0x1c, 0xb5, 0xa9, 0x9d, 0x1d, 0xf5, 0x9c, 0xed, 0xe2, 0xf7, 0x0c, 0x66, 0xd4, 0xe6, 0xc7, 0xc3,
	0x73, 0xc6, 0xc5, 0x39, 0x9d, 0x07, 0x23, 0x6e, 0x17, 0x14, 0xb7, 0x55, 0x5c, 0x3e, 0x32, 0x37,
	0xfc, 0x96, 0xc1, 0x31, 0xda, 0x8d, 0x93, 0xb3, 0xa5, 0x5e, 0x06, 0xeb, 0x4c, 0x2e, 0x8e, 0x68,
	0xbd, 0xaa, 0x68, 0x2d, 0x61, 0x33, 0x93, 0x96, 0xc2, 0x3a, 0x3b, 0x89, 0x47, 0x66, 0x17, 0x7f,
	0x61, 0x70, 0x9c, 0x6e, 0x38, 0x4e, 0x4e, 0x93, 0x5e, 0xb9, 0x56, 0x33, 0x1f, 0x48, 0x84, 0xae,
	0x29, 0x42, 0x6b, 0x78, 0xf9, 0x28, 0x75, 0x32, 0x2b, 0xc6, 0xd9, 0x89, 0xd7, 0xf4, 0x2e, 0xfe,
	0xc4, 0xa0, 0x62, 0x56, 0x18, 0xe6, 0x12, 0x10, 0xf9, 0xd7, 0x70, 0x7c, 0x1f, 0xda, 0x6f, 0x2a,
	0xae, 0xaf, 0xe3, 0xf9, 0x87, 0xe1, 0x8a, 0xf7, 0x18, 0xcc, 0x26, 0xb6, 0x09, 0x9e, 0x9d, 0x98,
	0xf8, 0xe0, 0x9e, 0xb3, 0x5e, 0x29, 0x06, 0xfe, 0x3f, 0xc3, 0xa7, 0xd6, 0xda, 0xda, 0xda, 0xfd,
	0xbd, 0x3a, 0x7b, 0xb0, 0x57, 0x67, 0x7f, 0xee, 0xd5, 0xd9, 0x77, 0xfb, 0xf5, 0xd2, 0x83, 0xfd,
	0x7a, 0xe9, 0xf7, 0xfd, 0x7a, 0xe9, 0xe3, 0xa6, 0x1f, 0xc8, 0x4f, 0x07, 0x9b, 0xad, 0x5b, 0x7c,
	0xcb, 0x84, 0xd5, 0x7f, 0xce, 0x89, 0xce, 0x67, 0xce, 0xe7, 0x2a, 0x47, 0x34, 0x32, 0x62, 0xf3,
	0x98, 0xfa, 0x6d, 0xb2, 0xfa, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x86, 0xe2, 0x1a, 0x4f,
	0x0d, 0x00, 0x00,
}


var _ context.Context
var _ grpc.ClientConn



const _ = grpc.SupportPackageIsVersion4


//

type QueryClient interface {

	Proposal(ctx context.Context, in *QueryProposalRequest, opts ...grpc.CallOption) (*QueryProposalResponse, error)

	Proposals(ctx context.Context, in *QueryProposalsRequest, opts ...grpc.CallOption) (*QueryProposalsResponse, error)

	Vote(ctx context.Context, in *QueryVoteRequest, opts ...grpc.CallOption) (*QueryVoteResponse, error)

	Votes(ctx context.Context, in *QueryVotesRequest, opts ...grpc.CallOption) (*QueryVotesResponse, error)

	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)

	Deposit(ctx context.Context, in *QueryDepositRequest, opts ...grpc.CallOption) (*QueryDepositResponse, error)

	Deposits(ctx context.Context, in *QueryDepositsRequest, opts ...grpc.CallOption) (*QueryDepositsResponse, error)

	TallyResult(ctx context.Context, in *QueryTallyResultRequest, opts ...grpc.CallOption) (*QueryTallyResultResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Proposal(ctx context.Context, in *QueryProposalRequest, opts ...grpc.CallOption) (*QueryProposalResponse, error) {
	out := new(QueryProposalResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.Query/Proposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Proposals(ctx context.Context, in *QueryProposalsRequest, opts ...grpc.CallOption) (*QueryProposalsResponse, error) {
	out := new(QueryProposalsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.Query/Proposals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Vote(ctx context.Context, in *QueryVoteRequest, opts ...grpc.CallOption) (*QueryVoteResponse, error) {
	out := new(QueryVoteResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.Query/Vote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Votes(ctx context.Context, in *QueryVotesRequest, opts ...grpc.CallOption) (*QueryVotesResponse, error) {
	out := new(QueryVotesResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.Query/Votes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Deposit(ctx context.Context, in *QueryDepositRequest, opts ...grpc.CallOption) (*QueryDepositResponse, error) {
	out := new(QueryDepositResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.Query/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Deposits(ctx context.Context, in *QueryDepositsRequest, opts ...grpc.CallOption) (*QueryDepositsResponse, error) {
	out := new(QueryDepositsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.Query/Deposits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TallyResult(ctx context.Context, in *QueryTallyResultRequest, opts ...grpc.CallOption) (*QueryTallyResultResponse, error) {
	out := new(QueryTallyResultResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.Query/TallyResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


type QueryServer interface {

	Proposal(context.Context, *QueryProposalRequest) (*QueryProposalResponse, error)

	Proposals(context.Context, *QueryProposalsRequest) (*QueryProposalsResponse, error)

	Vote(context.Context, *QueryVoteRequest) (*QueryVoteResponse, error)

	Votes(context.Context, *QueryVotesRequest) (*QueryVotesResponse, error)

	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)

	Deposit(context.Context, *QueryDepositRequest) (*QueryDepositResponse, error)

	Deposits(context.Context, *QueryDepositsRequest) (*QueryDepositsResponse, error)

	TallyResult(context.Context, *QueryTallyResultRequest) (*QueryTallyResultResponse, error)
}


type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Proposal(ctx context.Context, req *QueryProposalRequest) (*QueryProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Proposal not implemented")
}
func (*UnimplementedQueryServer) Proposals(ctx context.Context, req *QueryProposalsRequest) (*QueryProposalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Proposals not implemented")
}
func (*UnimplementedQueryServer) Vote(ctx context.Context, req *QueryVoteRequest) (*QueryVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vote not implemented")
}
func (*UnimplementedQueryServer) Votes(ctx context.Context, req *QueryVotesRequest) (*QueryVotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Votes not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Deposit(ctx context.Context, req *QueryDepositRequest) (*QueryDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (*UnimplementedQueryServer) Deposits(ctx context.Context, req *QueryDepositsRequest) (*QueryDepositsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposits not implemented")
}
func (*UnimplementedQueryServer) TallyResult(ctx context.Context, req *QueryTallyResultRequest) (*QueryTallyResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TallyResult not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Proposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Proposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.Query/Proposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Proposal(ctx, req.(*QueryProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Proposals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProposalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Proposals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.Query/Proposals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Proposals(ctx, req.(*QueryProposalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.Query/Vote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Vote(ctx, req.(*QueryVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Votes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Votes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.Query/Votes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Votes(ctx, req.(*QueryVotesRequest))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: "/cosmos.gov.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.Query/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Deposit(ctx, req.(*QueryDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Deposits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDepositsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Deposits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.Query/Deposits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Deposits(ctx, req.(*QueryDepositsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TallyResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTallyResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TallyResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.Query/TallyResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TallyResult(ctx, req.(*QueryTallyResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.gov.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Proposal",
			Handler:    _Query_Proposal_Handler,
		},
		{
			MethodName: "Proposals",
			Handler:    _Query_Proposals_Handler,
		},
		{
			MethodName: "Vote",
			Handler:    _Query_Vote_Handler,
		},
		{
			MethodName: "Votes",
			Handler:    _Query_Votes_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _Query_Deposit_Handler,
		},
		{
			MethodName: "Deposits",
			Handler:    _Query_Deposits_Handler,
		},
		{
			MethodName: "TallyResult",
			Handler:    _Query_TallyResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/gov/v1beta1/query.proto",
}

func (m *QueryProposalRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProposalRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProposalRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProposalId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryProposalResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProposalResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProposalResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Proposal.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryProposalsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProposalsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProposalsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Voter) > 0 {
		i -= len(m.Voter)
		copy(dAtA[i:], m.Voter)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Voter)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProposalStatus != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ProposalStatus))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryProposalsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProposalsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProposalsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Proposals) > 0 {
		for iNdEx := len(m.Proposals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proposals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryVoteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVoteRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVoteRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Voter) > 0 {
		i -= len(m.Voter)
		copy(dAtA[i:], m.Voter)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Voter)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProposalId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryVoteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVoteResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVoteResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Vote.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryVotesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVotesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVotesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if m.ProposalId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryVotesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVotesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVotesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ParamsType) > 0 {
		i -= len(m.ParamsType)
		copy(dAtA[i:], m.ParamsType)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ParamsType)))
		i--
		dAtA[i] = 0xa
	}
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
		size, err := m.TallyParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.DepositParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.VotingParams.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryDepositRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDepositRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDepositRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProposalId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryDepositResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDepositResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDepositResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Deposit.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryDepositsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDepositsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDepositsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if m.ProposalId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryDepositsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDepositsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDepositsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Deposits) > 0 {
		for iNdEx := len(m.Deposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryTallyResultRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTallyResultRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTallyResultRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProposalId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryTallyResultResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTallyResultResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTallyResultResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Tally.MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryProposalRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovQuery(uint64(m.ProposalId))
	}
	return n
}

func (m *QueryProposalResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Proposal.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryProposalsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalStatus != 0 {
		n += 1 + sovQuery(uint64(m.ProposalStatus))
	}
	l = len(m.Voter)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryProposalsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Proposals) > 0 {
		for _, e := range m.Proposals {
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

func (m *QueryVoteRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovQuery(uint64(m.ProposalId))
	}
	l = len(m.Voter)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryVoteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Vote.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryVotesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovQuery(uint64(m.ProposalId))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryVotesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
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

func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ParamsType)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.VotingParams.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.DepositParams.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.TallyParams.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryDepositRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovQuery(uint64(m.ProposalId))
	}
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDepositResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Deposit.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryDepositsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovQuery(uint64(m.ProposalId))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDepositsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Deposits) > 0 {
		for _, e := range m.Deposits {
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

func (m *QueryTallyResultRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovQuery(uint64(m.ProposalId))
	}
	return n
}

func (m *QueryTallyResultResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Tally.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryProposalRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryProposalRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProposalRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *QueryProposalResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryProposalResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProposalResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposal", wireType)
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
			if err := m.Proposal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryProposalsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryProposalsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProposalsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalStatus", wireType)
			}
			m.ProposalStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalStatus |= ProposalStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
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
			m.Voter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
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
			m.Depositor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
func (m *QueryProposalsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryProposalsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProposalsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposals", wireType)
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
			m.Proposals = append(m.Proposals, Proposal{})
			if err := m.Proposals[len(m.Proposals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryVoteRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryVoteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVoteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
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
			m.Voter = string(dAtA[iNdEx:postIndex])
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
func (m *QueryVoteResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryVoteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVoteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
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
			if err := m.Vote.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryVotesRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryVotesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVotesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *QueryVotesResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryVotesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVotesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
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
			m.Votes = append(m.Votes, Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParamsType", wireType)
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
			m.ParamsType = string(dAtA[iNdEx:postIndex])
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
				return fmt.Errorf("proto: wrong wireType = %d for field VotingParams", wireType)
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
			if err := m.VotingParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositParams", wireType)
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
			if err := m.DepositParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TallyParams", wireType)
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
			if err := m.TallyParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryDepositRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDepositRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDepositRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
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
			m.Depositor = string(dAtA[iNdEx:postIndex])
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
func (m *QueryDepositResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDepositResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDepositResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
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
			if err := m.Deposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryDepositsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDepositsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDepositsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *QueryDepositsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDepositsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDepositsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposits", wireType)
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
			m.Deposits = append(m.Deposits, Deposit{})
			if err := m.Deposits[len(m.Deposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryTallyResultRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryTallyResultRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTallyResultRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *QueryTallyResultResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryTallyResultResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTallyResultResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tally", wireType)
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
			if err := m.Tally.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
