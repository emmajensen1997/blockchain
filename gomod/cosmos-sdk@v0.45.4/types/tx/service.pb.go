


package tx

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	types1 "github.com/tendermint/tendermint/proto/tendermint/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)


var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf





const _ = proto.GoGoProtoPackageIsVersion3


type OrderBy int32

const (

	OrderBy_ORDER_BY_UNSPECIFIED OrderBy = 0

	OrderBy_ORDER_BY_ASC OrderBy = 1

	OrderBy_ORDER_BY_DESC OrderBy = 2
)

var OrderBy_name = map[int32]string{
	0: "ORDER_BY_UNSPECIFIED",
	1: "ORDER_BY_ASC",
	2: "ORDER_BY_DESC",
}

var OrderBy_value = map[string]int32{
	"ORDER_BY_UNSPECIFIED": 0,
	"ORDER_BY_ASC":         1,
	"ORDER_BY_DESC":        2,
}

func (x OrderBy) String() string {
	return proto.EnumName(OrderBy_name, int32(x))
}

func (OrderBy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e0b00a618705eca7, []int{0}
}


type BroadcastMode int32

const (

	BroadcastMode_BROADCAST_MODE_UNSPECIFIED BroadcastMode = 0


	BroadcastMode_BROADCAST_MODE_BLOCK BroadcastMode = 1


	BroadcastMode_BROADCAST_MODE_SYNC BroadcastMode = 2


	BroadcastMode_BROADCAST_MODE_ASYNC BroadcastMode = 3
)

var BroadcastMode_name = map[int32]string{
	0: "BROADCAST_MODE_UNSPECIFIED",
	1: "BROADCAST_MODE_BLOCK",
	2: "BROADCAST_MODE_SYNC",
	3: "BROADCAST_MODE_ASYNC",
}

var BroadcastMode_value = map[string]int32{
	"BROADCAST_MODE_UNSPECIFIED": 0,
	"BROADCAST_MODE_BLOCK":       1,
	"BROADCAST_MODE_SYNC":        2,
	"BROADCAST_MODE_ASYNC":       3,
}

func (x BroadcastMode) String() string {
	return proto.EnumName(BroadcastMode_name, int32(x))
}

func (BroadcastMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e0b00a618705eca7, []int{1}
}



type GetTxsEventRequest struct {

	Events []string `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`

	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	OrderBy    OrderBy            `protobuf:"varint,3,opt,name=order_by,json=orderBy,proto3,enum=cosmos.tx.v1beta1.OrderBy" json:"order_by,omitempty"`
}

func (m *GetTxsEventRequest) Reset()         { *m = GetTxsEventRequest{} }
func (m *GetTxsEventRequest) String() string { return proto.CompactTextString(m) }
func (*GetTxsEventRequest) ProtoMessage()    {}
func (*GetTxsEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b00a618705eca7, []int{0}
}
func (m *GetTxsEventRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTxsEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTxsEventRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetTxsEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTxsEventRequest.Merge(m, src)
}
func (m *GetTxsEventRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetTxsEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTxsEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTxsEventRequest proto.InternalMessageInfo

func (m *GetTxsEventRequest) GetEvents() []string {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *GetTxsEventRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *GetTxsEventRequest) GetOrderBy() OrderBy {
	if m != nil {
		return m.OrderBy
	}
	return OrderBy_ORDER_BY_UNSPECIFIED
}



type GetTxsEventResponse struct {

	Txs []*Tx `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`

	TxResponses []*types.TxResponse `protobuf:"bytes,2,rep,name=tx_responses,json=txResponses,proto3" json:"tx_responses,omitempty"`

	Pagination *query.PageResponse `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *GetTxsEventResponse) Reset()         { *m = GetTxsEventResponse{} }
func (m *GetTxsEventResponse) String() string { return proto.CompactTextString(m) }
func (*GetTxsEventResponse) ProtoMessage()    {}
func (*GetTxsEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b00a618705eca7, []int{1}
}
func (m *GetTxsEventResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTxsEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTxsEventResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetTxsEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTxsEventResponse.Merge(m, src)
}
func (m *GetTxsEventResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetTxsEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTxsEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTxsEventResponse proto.InternalMessageInfo

func (m *GetTxsEventResponse) GetTxs() []*Tx {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *GetTxsEventResponse) GetTxResponses() []*types.TxResponse {
	if m != nil {
		return m.TxResponses
	}
	return nil
}

func (m *GetTxsEventResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}



type BroadcastTxRequest struct {

	TxBytes []byte        `protobuf:"bytes,1,opt,name=tx_bytes,json=txBytes,proto3" json:"tx_bytes,omitempty"`
	Mode    BroadcastMode `protobuf:"varint,2,opt,name=mode,proto3,enum=cosmos.tx.v1beta1.BroadcastMode" json:"mode,omitempty"`
}

func (m *BroadcastTxRequest) Reset()         { *m = BroadcastTxRequest{} }
func (m *BroadcastTxRequest) String() string { return proto.CompactTextString(m) }
func (*BroadcastTxRequest) ProtoMessage()    {}
func (*BroadcastTxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b00a618705eca7, []int{2}
}
func (m *BroadcastTxRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BroadcastTxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BroadcastTxRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BroadcastTxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadcastTxRequest.Merge(m, src)
}
func (m *BroadcastTxRequest) XXX_Size() int {
	return m.Size()
}
func (m *BroadcastTxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadcastTxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BroadcastTxRequest proto.InternalMessageInfo

func (m *BroadcastTxRequest) GetTxBytes() []byte {
	if m != nil {
		return m.TxBytes
	}
	return nil
}

func (m *BroadcastTxRequest) GetMode() BroadcastMode {
	if m != nil {
		return m.Mode
	}
	return BroadcastMode_BROADCAST_MODE_UNSPECIFIED
}



type BroadcastTxResponse struct {

	TxResponse *types.TxResponse `protobuf:"bytes,1,opt,name=tx_response,json=txResponse,proto3" json:"tx_response,omitempty"`
}

func (m *BroadcastTxResponse) Reset()         { *m = BroadcastTxResponse{} }
func (m *BroadcastTxResponse) String() string { return proto.CompactTextString(m) }
func (*BroadcastTxResponse) ProtoMessage()    {}
func (*BroadcastTxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b00a618705eca7, []int{3}
}
func (m *BroadcastTxResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BroadcastTxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BroadcastTxResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BroadcastTxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadcastTxResponse.Merge(m, src)
}
func (m *BroadcastTxResponse) XXX_Size() int {
	return m.Size()
}
func (m *BroadcastTxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadcastTxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BroadcastTxResponse proto.InternalMessageInfo

func (m *BroadcastTxResponse) GetTxResponse() *types.TxResponse {
	if m != nil {
		return m.TxResponse
	}
	return nil
}



type SimulateRequest struct {


	Tx *Tx `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`

	//

	TxBytes []byte `protobuf:"bytes,2,opt,name=tx_bytes,json=txBytes,proto3" json:"tx_bytes,omitempty"`
}

func (m *SimulateRequest) Reset()         { *m = SimulateRequest{} }
func (m *SimulateRequest) String() string { return proto.CompactTextString(m) }
func (*SimulateRequest) ProtoMessage()    {}
func (*SimulateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b00a618705eca7, []int{4}
}
func (m *SimulateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SimulateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SimulateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SimulateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimulateRequest.Merge(m, src)
}
func (m *SimulateRequest) XXX_Size() int {
	return m.Size()
}
func (m *SimulateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SimulateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SimulateRequest proto.InternalMessageInfo


func (m *SimulateRequest) GetTx() *Tx {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *SimulateRequest) GetTxBytes() []byte {
	if m != nil {
		return m.TxBytes
	}
	return nil
}



type SimulateResponse struct {

	GasInfo *types.GasInfo `protobuf:"bytes,1,opt,name=gas_info,json=gasInfo,proto3" json:"gas_info,omitempty"`

	Result *types.Result `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *SimulateResponse) Reset()         { *m = SimulateResponse{} }
func (m *SimulateResponse) String() string { return proto.CompactTextString(m) }
func (*SimulateResponse) ProtoMessage()    {}
func (*SimulateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b00a618705eca7, []int{5}
}
func (m *SimulateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SimulateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SimulateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SimulateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimulateResponse.Merge(m, src)
}
func (m *SimulateResponse) XXX_Size() int {
	return m.Size()
}
func (m *SimulateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SimulateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SimulateResponse proto.InternalMessageInfo

func (m *SimulateResponse) GetGasInfo() *types.GasInfo {
	if m != nil {
		return m.GasInfo
	}
	return nil
}

func (m *SimulateResponse) GetResult() *types.Result {
	if m != nil {
		return m.Result
	}
	return nil
}



type GetTxRequest struct {

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *GetTxRequest) Reset()         { *m = GetTxRequest{} }
func (m *GetTxRequest) String() string { return proto.CompactTextString(m) }
func (*GetTxRequest) ProtoMessage()    {}
func (*GetTxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b00a618705eca7, []int{6}
}
func (m *GetTxRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTxRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetTxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTxRequest.Merge(m, src)
}
func (m *GetTxRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetTxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTxRequest proto.InternalMessageInfo

func (m *GetTxRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}


type GetTxResponse struct {

	Tx *Tx `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`

	TxResponse *types.TxResponse `protobuf:"bytes,2,opt,name=tx_response,json=txResponse,proto3" json:"tx_response,omitempty"`
}

func (m *GetTxResponse) Reset()         { *m = GetTxResponse{} }
func (m *GetTxResponse) String() string { return proto.CompactTextString(m) }
func (*GetTxResponse) ProtoMessage()    {}
func (*GetTxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b00a618705eca7, []int{7}
}
func (m *GetTxResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTxResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetTxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTxResponse.Merge(m, src)
}
func (m *GetTxResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetTxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTxResponse proto.InternalMessageInfo

func (m *GetTxResponse) GetTx() *Tx {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *GetTxResponse) GetTxResponse() *types.TxResponse {
	if m != nil {
		return m.TxResponse
	}
	return nil
}



//

type GetBlockWithTxsRequest struct {

	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`

	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *GetBlockWithTxsRequest) Reset()         { *m = GetBlockWithTxsRequest{} }
func (m *GetBlockWithTxsRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockWithTxsRequest) ProtoMessage()    {}
func (*GetBlockWithTxsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b00a618705eca7, []int{8}
}
func (m *GetBlockWithTxsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetBlockWithTxsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetBlockWithTxsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetBlockWithTxsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockWithTxsRequest.Merge(m, src)
}
func (m *GetBlockWithTxsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetBlockWithTxsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockWithTxsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockWithTxsRequest proto.InternalMessageInfo

func (m *GetBlockWithTxsRequest) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *GetBlockWithTxsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}


//

type GetBlockWithTxsResponse struct {

	Txs     []*Tx           `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
	BlockId *types1.BlockID `protobuf:"bytes,2,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Block   *types1.Block   `protobuf:"bytes,3,opt,name=block,proto3" json:"block,omitempty"`

	Pagination *query.PageResponse `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *GetBlockWithTxsResponse) Reset()         { *m = GetBlockWithTxsResponse{} }
func (m *GetBlockWithTxsResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlockWithTxsResponse) ProtoMessage()    {}
func (*GetBlockWithTxsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b00a618705eca7, []int{9}
}
func (m *GetBlockWithTxsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetBlockWithTxsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetBlockWithTxsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetBlockWithTxsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockWithTxsResponse.Merge(m, src)
}
func (m *GetBlockWithTxsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetBlockWithTxsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockWithTxsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockWithTxsResponse proto.InternalMessageInfo

func (m *GetBlockWithTxsResponse) GetTxs() []*Tx {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *GetBlockWithTxsResponse) GetBlockId() *types1.BlockID {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *GetBlockWithTxsResponse) GetBlock() *types1.Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *GetBlockWithTxsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterEnum("cosmos.tx.v1beta1.OrderBy", OrderBy_name, OrderBy_value)
	golang_proto.RegisterEnum("cosmos.tx.v1beta1.OrderBy", OrderBy_name, OrderBy_value)
	proto.RegisterEnum("cosmos.tx.v1beta1.BroadcastMode", BroadcastMode_name, BroadcastMode_value)
	golang_proto.RegisterEnum("cosmos.tx.v1beta1.BroadcastMode", BroadcastMode_name, BroadcastMode_value)
	proto.RegisterType((*GetTxsEventRequest)(nil), "cosmos.tx.v1beta1.GetTxsEventRequest")
	golang_proto.RegisterType((*GetTxsEventRequest)(nil), "cosmos.tx.v1beta1.GetTxsEventRequest")
	proto.RegisterType((*GetTxsEventResponse)(nil), "cosmos.tx.v1beta1.GetTxsEventResponse")
	golang_proto.RegisterType((*GetTxsEventResponse)(nil), "cosmos.tx.v1beta1.GetTxsEventResponse")
	proto.RegisterType((*BroadcastTxRequest)(nil), "cosmos.tx.v1beta1.BroadcastTxRequest")
	golang_proto.RegisterType((*BroadcastTxRequest)(nil), "cosmos.tx.v1beta1.BroadcastTxRequest")
	proto.RegisterType((*BroadcastTxResponse)(nil), "cosmos.tx.v1beta1.BroadcastTxResponse")
	golang_proto.RegisterType((*BroadcastTxResponse)(nil), "cosmos.tx.v1beta1.BroadcastTxResponse")
	proto.RegisterType((*SimulateRequest)(nil), "cosmos.tx.v1beta1.SimulateRequest")
	golang_proto.RegisterType((*SimulateRequest)(nil), "cosmos.tx.v1beta1.SimulateRequest")
	proto.RegisterType((*SimulateResponse)(nil), "cosmos.tx.v1beta1.SimulateResponse")
	golang_proto.RegisterType((*SimulateResponse)(nil), "cosmos.tx.v1beta1.SimulateResponse")
	proto.RegisterType((*GetTxRequest)(nil), "cosmos.tx.v1beta1.GetTxRequest")
	golang_proto.RegisterType((*GetTxRequest)(nil), "cosmos.tx.v1beta1.GetTxRequest")
	proto.RegisterType((*GetTxResponse)(nil), "cosmos.tx.v1beta1.GetTxResponse")
	golang_proto.RegisterType((*GetTxResponse)(nil), "cosmos.tx.v1beta1.GetTxResponse")
	proto.RegisterType((*GetBlockWithTxsRequest)(nil), "cosmos.tx.v1beta1.GetBlockWithTxsRequest")
	golang_proto.RegisterType((*GetBlockWithTxsRequest)(nil), "cosmos.tx.v1beta1.GetBlockWithTxsRequest")
	proto.RegisterType((*GetBlockWithTxsResponse)(nil), "cosmos.tx.v1beta1.GetBlockWithTxsResponse")
	golang_proto.RegisterType((*GetBlockWithTxsResponse)(nil), "cosmos.tx.v1beta1.GetBlockWithTxsResponse")
}

func init() { proto.RegisterFile("cosmos/tx/v1beta1/service.proto", fileDescriptor_e0b00a618705eca7) }
func init() {
	golang_proto.RegisterFile("cosmos/tx/v1beta1/service.proto", fileDescriptor_e0b00a618705eca7)
}

var fileDescriptor_e0b00a618705eca7 = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4f, 0x6f, 0x1a, 0x47,
	0x14, 0xf7, 0x2e, 0xb6, 0x21, 0x0f, 0x3b, 0x21, 0x63, 0xd7, 0x26, 0x24, 0xc5, 0x64, 0x53, 0x6c,
	0x07, 0xc9, 0xbb, 0x0a, 0x4d, 0xa5, 0xaa, 0xea, 0xc5, 0xfc, 0x89, 0x8b, 0xda, 0x84, 0x68, 0x70,
	0x15, 0xa5, 0xaa, 0x84, 0x16, 0x98, 0x2c, 0xab, 0x98, 0x1d, 0xbc, 0x33, 0x58, 0x8b, 0x1c, 0xab,
	0x52, 0x8f, 0x3d, 0x55, 0xed, 0xa1, 0x1f, 0xa2, 0x5f, 0xa2, 0xc7, 0x1e, 0x2d, 0xf5, 0xd2, 0x63,
	0x65, 0xf7, 0x03, 0xf4, 0x23, 0x54, 0x3b, 0x3b, 0xc0, 0x82, 0x97, 0x38, 0x8d, 0x7a, 0x81, 0x99,
	0x9d, 0xdf, 0x7b, 0xef, 0x37, 0xbf, 0xd9, 0xdf, 0x9b, 0x85, 0xad, 0x36, 0x65, 0x3d, 0xca, 0x0c,
	0xee, 0x19, 0x27, 0x8f, 0x5a, 0x84, 0x9b, 0x8f, 0x0c, 0x46, 0xdc, 0x13, 0xbb, 0x4d, 0xf4, 0xbe,
	0x4b, 0x39, 0x45, 0xb7, 0x03, 0x80, 0xce, 0x3d, 0x5d, 0x02, 0x32, 0xf7, 0x2c, 0x4a, 0xad, 0x23,
	0x62, 0x98, 0x7d, 0xdb, 0x30, 0x1d, 0x87, 0x72, 0x93, 0xdb, 0xd4, 0x61, 0x41, 0x40, 0xe6, 0x81,
	0xcc, 0xd8, 0x32, 0x19, 0x31, 0xcc, 0x56, 0xdb, 0x1e, 0x27, 0xf6, 0x27, 0x12, 0x94, 0xb9, 0x5a,
	0x96, 0x7b, 0x72, 0x6d, 0xdd, 0xa2, 0x16, 0x15, 0x43, 0xc3, 0x1f, 0xc9, 0xa7, 0x85, 0x70, 0xda,
	0xe3, 0x01, 0x71, 0x87, 0xe3, 0xc8, 0xbe, 0x69, 0xd9, 0x8e, 0xe0, 0x20, 0xb1, 0xf7, 0x38, 0x71,
	0x3a, 0xc4, 0xed, 0xd9, 0x0e, 0x37, 0xf8, 0xb0, 0x4f, 0x98, 0xd1, 0x3a, 0xa2, 0xed, 0xd7, 0x73,
	0x57, 0xc5, 0x6f, 0xb0, 0xaa, 0xfd, 0xaa, 0x00, 0x3a, 0x20, 0xfc, 0xd0, 0x63, 0xd5, 0x13, 0xe2,
	0x70, 0x4c, 0x8e, 0x07, 0x84, 0x71, 0xb4, 0x01, 0xcb, 0xc4, 0x9f, 0xb3, 0xb4, 0x92, 0x8b, 0xed,
	0xde, 0xc0, 0x72, 0x86, 0x9e, 0x00, 0x4c, 0xca, 0xa7, 0xd5, 0x9c, 0xb2, 0x9b, 0x2c, 0x6e, 0xeb,
	0x52, 0x33, 0x9f, 0xab, 0x2e, 0xb8, 0x8e, 0xb4, 0xd3, 0x9f, 0x9b, 0x16, 0x91, 0x39, 0x71, 0x28,
	0x12, 0x7d, 0x02, 0x09, 0xea, 0x76, 0x88, 0xdb, 0x6c, 0x0d, 0xd3, 0xb1, 0x9c, 0xb2, 0x7b, 0xb3,
	0x98, 0xd1, 0xaf, 0x28, 0xaf, 0xd7, 0x7d, 0x48, 0x69, 0x88, 0xe3, 0x34, 0x18, 0x68, 0xe7, 0x0a,
	0xac, 0x4d, 0xb1, 0x65, 0x7d, 0xea, 0x30, 0x82, 0x76, 0x20, 0xc6, 0xbd, 0x80, 0x6b, 0xb2, 0xf8,
	0x41, 0x44, 0xa6, 0x43, 0x0f, 0xfb, 0x08, 0x74, 0x00, 0x2b, 0xdc, 0x6b, 0xba, 0x32, 0x8e, 0xa5,
	0x55, 0x11, 0xf1, 0xd1, 0xd4, 0x0e, 0xc4, 0xb9, 0x85, 0x02, 0x25, 0x18, 0x27, 0xf9, 0x78, 0xec,
	0x27, 0x0a, 0x0b, 0x11, 0x13, 0x42, 0xec, 0x5c, 0x2b, 0x84, 0xcc, 0x14, 0x0a, 0xd5, 0x08, 0xa0,
	0x92, 0x4b, 0xcd, 0x4e, 0xdb, 0x64, 0xdc, 0x2f, 0x16, 0xe8, 0x7f, 0x07, 0x12, 0xdc, 0x6b, 0xb6,
	0x86, 0x9c, 0xf8, 0xbb, 0x52, 0x76, 0x57, 0x70, 0x9c, 0x7b, 0x25, 0x7f, 0x8a, 0x1e, 0xc3, 0x62,
	0x8f, 0x76, 0x88, 0x10, 0xff, 0x66, 0x31, 0x17, 0xb1, 0xd9, 0x71, 0xbe, 0xa7, 0xb4, 0x43, 0xb0,
	0x40, 0x6b, 0xdf, 0xc2, 0xda, 0x54, 0x19, 0x29, 0x5c, 0x15, 0x92, 0x21, 0x3d, 0x44, 0xa9, 0x77,
	0x95, 0x03, 0x26, 0x72, 0x68, 0x2f, 0xe0, 0x56, 0xc3, 0xee, 0x0d, 0x8e, 0x4c, 0x3e, 0x3a, 0x6d,
	0xf4, 0x10, 0x54, 0xee, 0xc9, 0x84, 0xd1, 0x27, 0x52, 0x52, 0xd3, 0x0a, 0x56, 0xb9, 0x37, 0xb5,
	0x59, 0x75, 0x6a, 0xb3, 0xda, 0x0f, 0x0a, 0xa4, 0x26, 0x99, 0x25, 0xe9, 0xcf, 0x21, 0x61, 0x99,
	0xac, 0x69, 0x3b, 0xaf, 0xa8, 0x2c, 0x70, 0x7f, 0x3e, 0xe3, 0x03, 0x93, 0xd5, 0x9c, 0x57, 0x14,
	0xc7, 0xad, 0x60, 0x80, 0x3e, 0x85, 0x65, 0x97, 0xb0, 0xc1, 0x11, 0x97, 0xaf, 0x6f, 0x6e, 0x7e,
	0x2c, 0x16, 0x38, 0x2c, 0xf1, 0x9a, 0x06, 0x2b, 0xe2, 0xe5, 0x1b, 0x6d, 0x11, 0xc1, 0x62, 0xd7,
	0x64, 0x5d, 0xc1, 0xe1, 0x06, 0x16, 0x63, 0xed, 0x0c, 0x56, 0x25, 0x46, 0x92, 0xcd, 0x5f, 0xab,
	0x83, 0xd0, 0x60, 0xe6, 0x20, 0xd4, 0xf7, 0x3c, 0x08, 0x0f, 0x36, 0x0e, 0x08, 0x2f, 0xf9, 0xf6,
	0x7f, 0x61, 0xf3, 0xee, 0xa1, 0xc7, 0x42, 0x8e, 0xee, 0x12, 0xdb, 0xea, 0x72, 0xc1, 0x25, 0x86,
	0xe5, 0xec, 0xff, 0x72, 0xb4, 0xf6, 0x8f, 0x02, 0x9b, 0x57, 0x4a, 0xff, 0x57, 0x7b, 0x3e, 0x86,
	0x84, 0x68, 0x5d, 0x4d, 0xbb, 0x23, 0xa9, 0xdc, 0xd1, 0x27, 0xed, 0x4b, 0x0f, 0x1a, 0x97, 0x28,
	0x51, 0xab, 0xe0, 0xb8, 0x80, 0xd6, 0x3a, 0x68, 0x0f, 0x96, 0xc4, 0x50, 0xda, 0x70, 0x73, 0x4e,
	0x08, 0x0e, 0x50, 0x33, 0xd6, 0x5d, 0x7c, 0x6f, 0xeb, 0x16, 0xbe, 0x80, 0xb8, 0xec, 0x50, 0x28,
	0x0d, 0xeb, 0x75, 0x5c, 0xa9, 0xe2, 0x66, 0xe9, 0x65, 0xf3, 0xeb, 0x67, 0x8d, 0xe7, 0xd5, 0x72,
	0xed, 0x49, 0xad, 0x5a, 0x49, 0x2d, 0xa0, 0x14, 0xac, 0x8c, 0x57, 0xf6, 0x1b, 0xe5, 0x94, 0x82,
	0x6e, 0xc3, 0xea, 0xf8, 0x49, 0xa5, 0xda, 0x28, 0xa7, 0xd4, 0xc2, 0x1b, 0x58, 0x9d, 0x32, 0x2d,
	0xca, 0x42, 0xa6, 0x84, 0xeb, 0xfb, 0x95, 0xf2, 0x7e, 0xe3, 0xb0, 0xf9, 0xb4, 0x5e, 0xa9, 0xce,
	0x64, 0x4d, 0xc3, 0xfa, 0xcc, 0x7a, 0xe9, 0xab, 0x7a, 0xf9, 0xcb, 0x94, 0x82, 0x36, 0x61, 0x6d,
	0x66, 0xa5, 0xf1, 0xf2, 0x59, 0x39, 0xa5, 0x46, 0x84, 0xec, 0x8b, 0x95, 0x58, 0xf1, 0xa7, 0x25,
	0x88, 0x37, 0x82, 0x5b, 0x10, 0x9d, 0x42, 0x62, 0xe4, 0x37, 0xa4, 0x45, 0x9c, 0xd4, 0x8c, 0xcd,
	0x33, 0x0f, 0xde, 0x8a, 0x91, 0x6f, 0xe5, 0xf6, 0xf7, 0x7f, 0xfc, 0xfd, 0xb3, 0x9a, 0xd3, 0xee,
	0x1a, 0x11, 0xd7, 0xaf, 0x04, 0x7f, 0xa6, 0x14, 0xd0, 0x31, 0x2c, 0x09, 0xf3, 0xa0, 0xad, 0x88,
	0xac, 0x61, 0xeb, 0x65, 0x72, 0xf3, 0x01, 0xb2, 0x66, 0x5e, 0xd4, 0xdc, 0x42, 0x1f, 0x1a, 0x51,
	0x77, 0x2f, 0x33, 0x4e, 0x7d, 0xbb, 0x9e, 0xa1, 0xef, 0x20, 0x19, 0xea, 0x8b, 0x28, 0xff, 0xb6,
	0x76, 0x3a, 0x29, 0xbf, 0x7d, 0x1d, 0x4c, 0x92, 0xb8, 0x2f, 0x48, 0xdc, 0xd5, 0x36, 0xa2, 0x49,
	0xf8, 0x7b, 0x7e, 0x03, 0xc9, 0xd0, 0x8d, 0x16, 0x49, 0xe0, 0xea, 0xfd, 0x1c, 0x49, 0x20, 0xe2,
	0x62, 0xd4, 0xb2, 0x82, 0x40, 0x1a, 0xcd, 0x21, 0x80, 0x7e, 0x51, 0xe0, 0xd6, 0x8c, 0x6b, 0xd1,
	0xc3, 0xe8, 0xdc, 0x11, 0x4d, 0x25, 0x53, 0x78, 0x17, 0xa8, 0xa4, 0xb2, 0x27, 0xa8, 0xec, 0xa0,
	0xfc, 0x9c, 0x03, 0x11, 0xe6, 0x34, 0x4e, 0x83, 0xb6, 0x74, 0x56, 0x2a, 0xff, 0x7e, 0x91, 0x55,
	0xce, 0x2f, 0xb2, 0xca, 0x5f, 0x17, 0x59, 0xe5, 0xc7, 0xcb, 0xec, 0xc2, 0x6f, 0x97, 0x59, 0xe5,
	0xfc, 0x32, 0xbb, 0xf0, 0xe7, 0x65, 0x76, 0xe1, 0x9b, 0xbc, 0x65, 0xf3, 0xee, 0xa0, 0xa5, 0xb7,
	0x69, 0x6f, 0x94, 0x2e, 0xf8, 0xdb, 0x63, 0x9d, 0xd7, 0xa3, 0xcf, 0x1c, 0xaf, 0xb5, 0x2c, 0x3e,
	0x72, 0x3e, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0xb5, 0xaa, 0xba, 0x0d, 0xf7, 0x09, 0x00, 0x00,
}


var _ context.Context
var _ grpc.ClientConn



const _ = grpc.SupportPackageIsVersion4


//

type ServiceClient interface {

	Simulate(ctx context.Context, in *SimulateRequest, opts ...grpc.CallOption) (*SimulateResponse, error)

	GetTx(ctx context.Context, in *GetTxRequest, opts ...grpc.CallOption) (*GetTxResponse, error)

	BroadcastTx(ctx context.Context, in *BroadcastTxRequest, opts ...grpc.CallOption) (*BroadcastTxResponse, error)

	GetTxsEvent(ctx context.Context, in *GetTxsEventRequest, opts ...grpc.CallOption) (*GetTxsEventResponse, error)

	//

	GetBlockWithTxs(ctx context.Context, in *GetBlockWithTxsRequest, opts ...grpc.CallOption) (*GetBlockWithTxsResponse, error)
}

type serviceClient struct {
	cc grpc1.ClientConn
}

func NewServiceClient(cc grpc1.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Simulate(ctx context.Context, in *SimulateRequest, opts ...grpc.CallOption) (*SimulateResponse, error) {
	out := new(SimulateResponse)
	err := c.cc.Invoke(ctx, "/cosmos.tx.v1beta1.Service/Simulate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetTx(ctx context.Context, in *GetTxRequest, opts ...grpc.CallOption) (*GetTxResponse, error) {
	out := new(GetTxResponse)
	err := c.cc.Invoke(ctx, "/cosmos.tx.v1beta1.Service/GetTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BroadcastTx(ctx context.Context, in *BroadcastTxRequest, opts ...grpc.CallOption) (*BroadcastTxResponse, error) {
	out := new(BroadcastTxResponse)
	err := c.cc.Invoke(ctx, "/cosmos.tx.v1beta1.Service/BroadcastTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetTxsEvent(ctx context.Context, in *GetTxsEventRequest, opts ...grpc.CallOption) (*GetTxsEventResponse, error) {
	out := new(GetTxsEventResponse)
	err := c.cc.Invoke(ctx, "/cosmos.tx.v1beta1.Service/GetTxsEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetBlockWithTxs(ctx context.Context, in *GetBlockWithTxsRequest, opts ...grpc.CallOption) (*GetBlockWithTxsResponse, error) {
	out := new(GetBlockWithTxsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.tx.v1beta1.Service/GetBlockWithTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


type ServiceServer interface {

	Simulate(context.Context, *SimulateRequest) (*SimulateResponse, error)

	GetTx(context.Context, *GetTxRequest) (*GetTxResponse, error)

	BroadcastTx(context.Context, *BroadcastTxRequest) (*BroadcastTxResponse, error)

	GetTxsEvent(context.Context, *GetTxsEventRequest) (*GetTxsEventResponse, error)

	//

	GetBlockWithTxs(context.Context, *GetBlockWithTxsRequest) (*GetBlockWithTxsResponse, error)
}


type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Simulate(ctx context.Context, req *SimulateRequest) (*SimulateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Simulate not implemented")
}
func (*UnimplementedServiceServer) GetTx(ctx context.Context, req *GetTxRequest) (*GetTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTx not implemented")
}
func (*UnimplementedServiceServer) BroadcastTx(ctx context.Context, req *BroadcastTxRequest) (*BroadcastTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastTx not implemented")
}
func (*UnimplementedServiceServer) GetTxsEvent(ctx context.Context, req *GetTxsEventRequest) (*GetTxsEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxsEvent not implemented")
}
func (*UnimplementedServiceServer) GetBlockWithTxs(ctx context.Context, req *GetBlockWithTxsRequest) (*GetBlockWithTxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockWithTxs not implemented")
}

func RegisterServiceServer(s grpc1.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Simulate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimulateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Simulate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.tx.v1beta1.Service/Simulate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Simulate(ctx, req.(*SimulateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.tx.v1beta1.Service/GetTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetTx(ctx, req.(*GetTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_BroadcastTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).BroadcastTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.tx.v1beta1.Service/BroadcastTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).BroadcastTx(ctx, req.(*BroadcastTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetTxsEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxsEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetTxsEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.tx.v1beta1.Service/GetTxsEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetTxsEvent(ctx, req.(*GetTxsEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetBlockWithTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockWithTxsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetBlockWithTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.tx.v1beta1.Service/GetBlockWithTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetBlockWithTxs(ctx, req.(*GetBlockWithTxsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.tx.v1beta1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Simulate",
			Handler:    _Service_Simulate_Handler,
		},
		{
			MethodName: "GetTx",
			Handler:    _Service_GetTx_Handler,
		},
		{
			MethodName: "BroadcastTx",
			Handler:    _Service_BroadcastTx_Handler,
		},
		{
			MethodName: "GetTxsEvent",
			Handler:    _Service_GetTxsEvent_Handler,
		},
		{
			MethodName: "GetBlockWithTxs",
			Handler:    _Service_GetBlockWithTxs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/tx/v1beta1/service.proto",
}

func (m *GetTxsEventRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTxsEventRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetTxsEventRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OrderBy != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.OrderBy))
		i--
		dAtA[i] = 0x18
	}
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Events) > 0 {
		for iNdEx := len(m.Events) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Events[iNdEx])
			copy(dAtA[i:], m.Events[iNdEx])
			i = encodeVarintService(dAtA, i, uint64(len(m.Events[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetTxsEventResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTxsEventResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetTxsEventResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TxResponses) > 0 {
		for iNdEx := len(m.TxResponses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TxResponses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Txs) > 0 {
		for iNdEx := len(m.Txs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Txs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *BroadcastTxRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BroadcastTxRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BroadcastTxRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Mode != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.Mode))
		i--
		dAtA[i] = 0x10
	}
	if len(m.TxBytes) > 0 {
		i -= len(m.TxBytes)
		copy(dAtA[i:], m.TxBytes)
		i = encodeVarintService(dAtA, i, uint64(len(m.TxBytes)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BroadcastTxResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BroadcastTxResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BroadcastTxResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TxResponse != nil {
		{
			size, err := m.TxResponse.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SimulateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SimulateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SimulateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxBytes) > 0 {
		i -= len(m.TxBytes)
		copy(dAtA[i:], m.TxBytes)
		i = encodeVarintService(dAtA, i, uint64(len(m.TxBytes)))
		i--
		dAtA[i] = 0x12
	}
	if m.Tx != nil {
		{
			size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SimulateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SimulateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SimulateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Result != nil {
		{
			size, err := m.Result.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.GasInfo != nil {
		{
			size, err := m.GasInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetTxRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTxRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetTxRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintService(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetTxResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTxResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetTxResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TxResponse != nil {
		{
			size, err := m.TxResponse.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Tx != nil {
		{
			size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetBlockWithTxsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetBlockWithTxsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetBlockWithTxsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Height != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetBlockWithTxsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetBlockWithTxsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetBlockWithTxsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Block != nil {
		{
			size, err := m.Block.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.BlockId != nil {
		{
			size, err := m.BlockId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Txs) > 0 {
		for iNdEx := len(m.Txs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Txs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	offset -= sovService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetTxsEventRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, s := range m.Events {
			l = len(s)
			n += 1 + l + sovService(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.OrderBy != 0 {
		n += 1 + sovService(uint64(m.OrderBy))
	}
	return n
}

func (m *GetTxsEventResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for _, e := range m.Txs {
			l = e.Size()
			n += 1 + l + sovService(uint64(l))
		}
	}
	if len(m.TxResponses) > 0 {
		for _, e := range m.TxResponses {
			l = e.Size()
			n += 1 + l + sovService(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *BroadcastTxRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxBytes)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	if m.Mode != 0 {
		n += 1 + sovService(uint64(m.Mode))
	}
	return n
}

func (m *BroadcastTxResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TxResponse != nil {
		l = m.TxResponse.Size()
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *SimulateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Tx != nil {
		l = m.Tx.Size()
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.TxBytes)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *SimulateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GasInfo != nil {
		l = m.GasInfo.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.Result != nil {
		l = m.Result.Size()
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *GetTxRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *GetTxResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Tx != nil {
		l = m.Tx.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.TxResponse != nil {
		l = m.TxResponse.Size()
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *GetBlockWithTxsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovService(uint64(m.Height))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *GetBlockWithTxsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for _, e := range m.Txs {
			l = e.Size()
			n += 1 + l + sovService(uint64(l))
		}
	}
	if m.BlockId != nil {
		l = m.BlockId.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.Block != nil {
		l = m.Block.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetTxsEventRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: GetTxsEventRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTxsEventRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderBy", wireType)
			}
			m.OrderBy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderBy |= OrderBy(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *GetTxsEventResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: GetTxsEventResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTxsEventResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txs = append(m.Txs, &Tx{})
			if err := m.Txs[len(m.Txs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxResponses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxResponses = append(m.TxResponses, &types.TxResponse{})
			if err := m.TxResponses[len(m.TxResponses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
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
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *BroadcastTxRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: BroadcastTxRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BroadcastTxRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxBytes = append(m.TxBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.TxBytes == nil {
				m.TxBytes = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mode", wireType)
			}
			m.Mode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mode |= BroadcastMode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *BroadcastTxResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: BroadcastTxResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BroadcastTxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TxResponse == nil {
				m.TxResponse = &types.TxResponse{}
			}
			if err := m.TxResponse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *SimulateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: SimulateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SimulateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tx == nil {
				m.Tx = &Tx{}
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxBytes = append(m.TxBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.TxBytes == nil {
				m.TxBytes = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *SimulateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: SimulateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SimulateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GasInfo == nil {
				m.GasInfo = &types.GasInfo{}
			}
			if err := m.GasInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Result == nil {
				m.Result = &types.Result{}
			}
			if err := m.Result.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *GetTxRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: GetTxRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTxRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *GetTxResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: GetTxResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tx == nil {
				m.Tx = &Tx{}
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TxResponse == nil {
				m.TxResponse = &types.TxResponse{}
			}
			if err := m.TxResponse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *GetBlockWithTxsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: GetBlockWithTxsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetBlockWithTxsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
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
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
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
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *GetBlockWithTxsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: GetBlockWithTxsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetBlockWithTxsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txs = append(m.Txs, &Tx{})
			if err := m.Txs[len(m.Txs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BlockId == nil {
				m.BlockId = &types1.BlockID{}
			}
			if err := m.BlockId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Block == nil {
				m.Block = &types1.Block{}
			}
			if err := m.Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
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
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
				return 0, ErrInvalidLengthService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupService = fmt.Errorf("proto: unexpected end of group")
)
