


package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types1 "github.com/tendermint/tendermint/abci/types"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)


var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf





const _ = proto.GoGoProtoPackageIsVersion3 



type TxResponse struct {
	
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	
	TxHash string `protobuf:"bytes,2,opt,name=txhash,proto3" json:"txhash,omitempty"`
	
	Codespace string `protobuf:"bytes,3,opt,name=codespace,proto3" json:"codespace,omitempty"`
	
	Code uint32 `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	
	Data string `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	
	
	RawLog string `protobuf:"bytes,6,opt,name=raw_log,json=rawLog,proto3" json:"raw_log,omitempty"`
	
	Logs ABCIMessageLogs `protobuf:"bytes,7,rep,name=logs,proto3,castrepeated=ABCIMessageLogs" json:"logs"`
	
	Info string `protobuf:"bytes,8,opt,name=info,proto3" json:"info,omitempty"`
	
	GasWanted int64 `protobuf:"varint,9,opt,name=gas_wanted,json=gasWanted,proto3" json:"gas_wanted,omitempty"`
	
	GasUsed int64 `protobuf:"varint,10,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	
	Tx *types.Any `protobuf:"bytes,11,opt,name=tx,proto3" json:"tx,omitempty"`
	
	
	
	Timestamp string `protobuf:"bytes,12,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	
	
	
	
	//
	
	Events []types1.Event `protobuf:"bytes,13,rep,name=events,proto3" json:"events"`
}

func (m *TxResponse) Reset()      { *m = TxResponse{} }
func (*TxResponse) ProtoMessage() {}
func (*TxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e37629bc7eb0df8, []int{0}
}
func (m *TxResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxResponse.Merge(m, src)
}
func (m *TxResponse) XXX_Size() int {
	return m.Size()
}
func (m *TxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TxResponse proto.InternalMessageInfo


type ABCIMessageLog struct {
	MsgIndex uint32 `protobuf:"varint,1,opt,name=msg_index,json=msgIndex,proto3" json:"msg_index,omitempty"`
	Log      string `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
	
	
	Events StringEvents `protobuf:"bytes,3,rep,name=events,proto3,castrepeated=StringEvents" json:"events"`
}

func (m *ABCIMessageLog) Reset()      { *m = ABCIMessageLog{} }
func (*ABCIMessageLog) ProtoMessage() {}
func (*ABCIMessageLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e37629bc7eb0df8, []int{1}
}
func (m *ABCIMessageLog) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ABCIMessageLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ABCIMessageLog.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ABCIMessageLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ABCIMessageLog.Merge(m, src)
}
func (m *ABCIMessageLog) XXX_Size() int {
	return m.Size()
}
func (m *ABCIMessageLog) XXX_DiscardUnknown() {
	xxx_messageInfo_ABCIMessageLog.DiscardUnknown(m)
}

var xxx_messageInfo_ABCIMessageLog proto.InternalMessageInfo

func (m *ABCIMessageLog) GetMsgIndex() uint32 {
	if m != nil {
		return m.MsgIndex
	}
	return 0
}

func (m *ABCIMessageLog) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func (m *ABCIMessageLog) GetEvents() StringEvents {
	if m != nil {
		return m.Events
	}
	return nil
}



type StringEvent struct {
	Type       string      `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Attributes []Attribute `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes"`
}

func (m *StringEvent) Reset()      { *m = StringEvent{} }
func (*StringEvent) ProtoMessage() {}
func (*StringEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e37629bc7eb0df8, []int{2}
}
func (m *StringEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StringEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StringEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StringEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringEvent.Merge(m, src)
}
func (m *StringEvent) XXX_Size() int {
	return m.Size()
}
func (m *StringEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_StringEvent.DiscardUnknown(m)
}

var xxx_messageInfo_StringEvent proto.InternalMessageInfo

func (m *StringEvent) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *StringEvent) GetAttributes() []Attribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}



type Attribute struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Attribute) Reset()      { *m = Attribute{} }
func (*Attribute) ProtoMessage() {}
func (*Attribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e37629bc7eb0df8, []int{3}
}
func (m *Attribute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Attribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Attribute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Attribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attribute.Merge(m, src)
}
func (m *Attribute) XXX_Size() int {
	return m.Size()
}
func (m *Attribute) XXX_DiscardUnknown() {
	xxx_messageInfo_Attribute.DiscardUnknown(m)
}

var xxx_messageInfo_Attribute proto.InternalMessageInfo

func (m *Attribute) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Attribute) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}


type GasInfo struct {
	
	GasWanted uint64 `protobuf:"varint,1,opt,name=gas_wanted,json=gasWanted,proto3" json:"gas_wanted,omitempty" yaml:"gas_wanted"`
	
	GasUsed uint64 `protobuf:"varint,2,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty" yaml:"gas_used"`
}

func (m *GasInfo) Reset()      { *m = GasInfo{} }
func (*GasInfo) ProtoMessage() {}
func (*GasInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e37629bc7eb0df8, []int{4}
}
func (m *GasInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GasInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GasInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GasInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GasInfo.Merge(m, src)
}
func (m *GasInfo) XXX_Size() int {
	return m.Size()
}
func (m *GasInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GasInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GasInfo proto.InternalMessageInfo

func (m *GasInfo) GetGasWanted() uint64 {
	if m != nil {
		return m.GasWanted
	}
	return 0
}

func (m *GasInfo) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}


type Result struct {
	
	
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	
	Log string `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
	
	
	Events []types1.Event `protobuf:"bytes,3,rep,name=events,proto3" json:"events"`
}

func (m *Result) Reset()      { *m = Result{} }
func (*Result) ProtoMessage() {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e37629bc7eb0df8, []int{5}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Result.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return m.Size()
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo



type SimulationResponse struct {
	GasInfo `protobuf:"bytes,1,opt,name=gas_info,json=gasInfo,proto3,embedded=gas_info" json:"gas_info"`
	Result  *Result `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *SimulationResponse) Reset()      { *m = SimulationResponse{} }
func (*SimulationResponse) ProtoMessage() {}
func (*SimulationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e37629bc7eb0df8, []int{6}
}
func (m *SimulationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SimulationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SimulationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SimulationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimulationResponse.Merge(m, src)
}
func (m *SimulationResponse) XXX_Size() int {
	return m.Size()
}
func (m *SimulationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SimulationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SimulationResponse proto.InternalMessageInfo

func (m *SimulationResponse) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}



type MsgData struct {
	MsgType string `protobuf:"bytes,1,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
	Data    []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *MsgData) Reset()      { *m = MsgData{} }
func (*MsgData) ProtoMessage() {}
func (*MsgData) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e37629bc7eb0df8, []int{7}
}
func (m *MsgData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgData.Merge(m, src)
}
func (m *MsgData) XXX_Size() int {
	return m.Size()
}
func (m *MsgData) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgData.DiscardUnknown(m)
}

var xxx_messageInfo_MsgData proto.InternalMessageInfo

func (m *MsgData) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *MsgData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}



type TxMsgData struct {
	Data []*MsgData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (m *TxMsgData) Reset()      { *m = TxMsgData{} }
func (*TxMsgData) ProtoMessage() {}
func (*TxMsgData) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e37629bc7eb0df8, []int{8}
}
func (m *TxMsgData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxMsgData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxMsgData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxMsgData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxMsgData.Merge(m, src)
}
func (m *TxMsgData) XXX_Size() int {
	return m.Size()
}
func (m *TxMsgData) XXX_DiscardUnknown() {
	xxx_messageInfo_TxMsgData.DiscardUnknown(m)
}

var xxx_messageInfo_TxMsgData proto.InternalMessageInfo

func (m *TxMsgData) GetData() []*MsgData {
	if m != nil {
		return m.Data
	}
	return nil
}


type SearchTxsResult struct {
	
	TotalCount uint64 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count" yaml:"total_count"`
	
	Count uint64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	
	PageNumber uint64 `protobuf:"varint,3,opt,name=page_number,json=pageNumber,proto3" json:"page_number" yaml:"page_number"`
	
	PageTotal uint64 `protobuf:"varint,4,opt,name=page_total,json=pageTotal,proto3" json:"page_total" yaml:"page_total"`
	
	Limit uint64 `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	
	Txs []*TxResponse `protobuf:"bytes,6,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (m *SearchTxsResult) Reset()      { *m = SearchTxsResult{} }
func (*SearchTxsResult) ProtoMessage() {}
func (*SearchTxsResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e37629bc7eb0df8, []int{9}
}
func (m *SearchTxsResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SearchTxsResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SearchTxsResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SearchTxsResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchTxsResult.Merge(m, src)
}
func (m *SearchTxsResult) XXX_Size() int {
	return m.Size()
}
func (m *SearchTxsResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchTxsResult.DiscardUnknown(m)
}

var xxx_messageInfo_SearchTxsResult proto.InternalMessageInfo

func (m *SearchTxsResult) GetTotalCount() uint64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *SearchTxsResult) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *SearchTxsResult) GetPageNumber() uint64 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *SearchTxsResult) GetPageTotal() uint64 {
	if m != nil {
		return m.PageTotal
	}
	return 0
}

func (m *SearchTxsResult) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SearchTxsResult) GetTxs() []*TxResponse {
	if m != nil {
		return m.Txs
	}
	return nil
}

func init() {
	proto.RegisterType((*TxResponse)(nil), "cosmos.base.abci.v1beta1.TxResponse")
	proto.RegisterType((*ABCIMessageLog)(nil), "cosmos.base.abci.v1beta1.ABCIMessageLog")
	proto.RegisterType((*StringEvent)(nil), "cosmos.base.abci.v1beta1.StringEvent")
	proto.RegisterType((*Attribute)(nil), "cosmos.base.abci.v1beta1.Attribute")
	proto.RegisterType((*GasInfo)(nil), "cosmos.base.abci.v1beta1.GasInfo")
	proto.RegisterType((*Result)(nil), "cosmos.base.abci.v1beta1.Result")
	proto.RegisterType((*SimulationResponse)(nil), "cosmos.base.abci.v1beta1.SimulationResponse")
	proto.RegisterType((*MsgData)(nil), "cosmos.base.abci.v1beta1.MsgData")
	proto.RegisterType((*TxMsgData)(nil), "cosmos.base.abci.v1beta1.TxMsgData")
	proto.RegisterType((*SearchTxsResult)(nil), "cosmos.base.abci.v1beta1.SearchTxsResult")
}

func init() {
	proto.RegisterFile("cosmos/base/abci/v1beta1/abci.proto", fileDescriptor_4e37629bc7eb0df8)
}

var fileDescriptor_4e37629bc7eb0df8 = []byte{
	
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xf6, 0xda, 0xee, 0x3a, 0x7e, 0x4e, 0x08, 0x0c, 0xa1, 0xdd, 0xb4, 0xe0, 0x35, 0x9b, 0x56,
	0xf2, 0x85, 0xb5, 0x9a, 0x06, 0x84, 0x7a, 0x40, 0xd4, 0x85, 0xd2, 0x48, 0x2d, 0x87, 0x89, 0x2b,
	0x24, 0x2e, 0xd6, 0xd8, 0x9e, 0x8e, 0x97, 0x7a, 0x77, 0xac, 0x9d, 0xd9, 0x64, 0x7d, 0xe3, 0xc8,
	0x91, 0x13, 0x07, 0x4e, 0x9c, 0xf9, 0x4b, 0x7a, 0x40, 0x22, 0xc7, 0x1e, 0x90, 0x81, 0xe4, 0xd6,
	0x63, 0xfe, 0x02, 0x34, 0x3f, 0xe2, 0xdd, 0x80, 0x5c, 0x89, 0x93, 0xdf, 0xfb, 0xde, 0x9b, 0x37,
	0xef, 0x7d, 0xef, 0xdb, 0x31, 0xec, 0x8d, 0xb9, 0x88, 0xb9, 0xe8, 0x8d, 0x88, 0xa0, 0x3d, 0x32,
	0x1a, 0x47, 0xbd, 0xe3, 0xbb, 0x23, 0x2a, 0xc9, 0x5d, 0xed, 0x84, 0xf3, 0x94, 0x4b, 0x8e, 0x3c,
	0x93, 0x14, 0xaa, 0xa4, 0x50, 0xe3, 0x36, 0xe9, 0xe6, 0x0e, 0xe3, 0x8c, 0xeb, 0xa4, 0x9e, 0xb2,
	0x4c, 0xfe, 0xcd, 0x5b, 0x92, 0x26, 0x13, 0x9a, 0xc6, 0x51, 0x22, 0x4d, 0x4d, 0xb9, 0x98, 0x53,
	0x61, 0x83, 0xbb, 0x8c, 0x73, 0x36, 0xa3, 0x3d, 0xed, 0x8d, 0xb2, 0xe7, 0x3d, 0x92, 0x2c, 0x4c,
	0x28, 0xf8, 0xad, 0x06, 0x30, 0xc8, 0x31, 0x15, 0x73, 0x9e, 0x08, 0x8a, 0xae, 0x83, 0x3b, 0xa5,
	0x11, 0x9b, 0x4a, 0xcf, 0xe9, 0x38, 0xdd, 0x1a, 0xb6, 0x1e, 0x0a, 0xc0, 0x95, 0xf9, 0x94, 0x88,
	0xa9, 0x57, 0xed, 0x38, 0xdd, 0x66, 0x1f, 0xce, 0x96, 0xbe, 0x3b, 0xc8, 0x1f, 0x13, 0x31, 0xc5,
	0x36, 0x82, 0xde, 0x87, 0xe6, 0x98, 0x4f, 0xa8, 0x98, 0x93, 0x31, 0xf5, 0x6a, 0x2a, 0x0d, 0x17,
	0x00, 0x42, 0x50, 0x57, 0x8e, 0x57, 0xef, 0x38, 0xdd, 0x2d, 0xac, 0x6d, 0x85, 0x4d, 0x88, 0x24,
	0xde, 0x35, 0x9d, 0xac, 0x6d, 0x74, 0x03, 0x1a, 0x29, 0x39, 0x19, 0xce, 0x38, 0xf3, 0x5c, 0x0d,
	0xbb, 0x29, 0x39, 0x79, 0xc2, 0x19, 0x7a, 0x06, 0xf5, 0x19, 0x67, 0xc2, 0x6b, 0x74, 0x6a, 0xdd,
	0xd6, 0x7e, 0x37, 0x5c, 0x47, 0x50, 0xf8, 0xa0, 0xff, 0xf0, 0xf0, 0x29, 0x15, 0x82, 0x30, 0xfa,
	0x84, 0xb3, 0xfe, 0x8d, 0x97, 0x4b, 0xbf, 0xf2, 0xeb, 0x9f, 0xfe, 0xf6, 0x55, 0x5c, 0x60, 0x5d,
	0x4e, 0xf5, 0x10, 0x25, 0xcf, 0xb9, 0xb7, 0x61, 0x7a, 0x50, 0x36, 0xfa, 0x00, 0x80, 0x11, 0x31,
	0x3c, 0x21, 0x89, 0xa4, 0x13, 0xaf, 0xa9, 0x99, 0x68, 0x32, 0x22, 0xbe, 0xd1, 0x00, 0xda, 0x85,
	0x0d, 0x15, 0xce, 0x04, 0x9d, 0x78, 0xa0, 0x83, 0x0d, 0x46, 0xc4, 0x33, 0x41, 0x27, 0xe8, 0x36,
	0x54, 0x65, 0xee, 0xb5, 0x3a, 0x4e, 0xb7, 0xb5, 0xbf, 0x13, 0x1a, 0xda, 0xc3, 0x4b, 0xda, 0xc3,
	0x07, 0xc9, 0x02, 0x57, 0x65, 0xae, 0x98, 0x92, 0x51, 0x4c, 0x85, 0x24, 0xf1, 0xdc, 0xdb, 0x34,
	0x4c, 0xad, 0x00, 0x74, 0x00, 0x2e, 0x3d, 0xa6, 0x89, 0x14, 0xde, 0x96, 0x1e, 0xf5, 0x7a, 0x58,
	0xec, 0xd6, 0x4c, 0xfa, 0xa5, 0x0a, 0xf7, 0xeb, 0x6a, 0x30, 0x6c, 0x73, 0xef, 0xd7, 0x7f, 0xf8,
	0xc5, 0xaf, 0x04, 0x3f, 0x3b, 0xf0, 0xd6, 0xd5, 0x39, 0xd1, 0x2d, 0x68, 0xc6, 0x82, 0x0d, 0xa3,
	0x64, 0x42, 0x73, 0xbd, 0xd5, 0x2d, 0xbc, 0x11, 0x0b, 0x76, 0xa8, 0x7c, 0xf4, 0x36, 0xd4, 0x14,
	0xd3, 0x7a, 0xa9, 0x58, 0x99, 0xe8, 0x68, 0x75, 0x7b, 0x4d, 0xdf, 0x7e, 0x67, 0x3d, 0xd1, 0x47,
	0x32, 0x8d, 0x12, 0x66, 0x9a, 0xd9, 0xb1, 0x2c, 0x6f, 0x96, 0x40, 0x51, 0x34, 0xf7, 0xfd, 0x1f,
	0x1d, 0x27, 0x48, 0xa1, 0x55, 0x8a, 0x2a, 0xe6, 0x95, 0x48, 0x75, 0x4f, 0x4d, 0xac, 0x6d, 0x74,
	0x08, 0x40, 0xa4, 0x4c, 0xa3, 0x51, 0x26, 0xa9, 0xf0, 0xaa, 0xba, 0x83, 0xbd, 0x37, 0xac, 0xfa,
	0x32, 0xd7, 0x92, 0x51, 0x3a, 0x6c, 0xef, 0xbc, 0x07, 0xcd, 0x55, 0x92, 0x9a, 0xf6, 0x05, 0x5d,
	0xd8, 0x0b, 0x95, 0x89, 0x76, 0xe0, 0xda, 0x31, 0x99, 0x65, 0xd4, 0x32, 0x60, 0x9c, 0x80, 0x43,
	0xe3, 0x2b, 0x22, 0x0e, 0x95, 0x14, 0x0e, 0xae, 0x48, 0x41, 0x9d, 0xac, 0xf7, 0xdf, 0xbb, 0x58,
	0xfa, 0xef, 0x2c, 0x48, 0x3c, 0xbb, 0x1f, 0x14, 0xb1, 0xa0, 0xac, 0x90, 0xb0, 0xa4, 0x90, 0xaa,
	0x3e, 0xf3, 0xee, 0xc5, 0xd2, 0xdf, 0x2e, 0xce, 0xa8, 0x48, 0xb0, 0x92, 0x4d, 0xf0, 0x1d, 0xb8,
	0x98, 0x8a, 0x6c, 0x26, 0x57, 0x9f, 0x84, 0xba, 0x69, 0xd3, 0x7e, 0x12, 0xff, 0x5d, 0xd2, 0xc1,
	0xbf, 0x96, 0xf4, 0x7f, 0x24, 0xf2, 0x93, 0x03, 0xe8, 0x28, 0x8a, 0xb3, 0x19, 0x91, 0x11, 0x4f,
	0x56, 0x5f, 0xfe, 0x23, 0xd3, 0xb2, 0xfe, 0x16, 0x1c, 0xad, 0xdf, 0x0f, 0xd7, 0xf3, 0x6e, 0xd9,
	0xe9, 0x6f, 0xa8, 0xfa, 0xa7, 0x4b, 0xdf, 0xd1, 0xa3, 0x68, 0xc2, 0x3e, 0x05, 0x37, 0xd5, 0xa3,
	0xe8, 0x7e, 0x5b, 0xfb, 0x9d, 0xf5, 0x55, 0xcc, 0xc8, 0xd8, 0xe6, 0x07, 0x9f, 0x41, 0xe3, 0xa9,
	0x60, 0x5f, 0xa8, 0x89, 0x77, 0x41, 0x49, 0x74, 0x58, 0x92, 0x47, 0x23, 0x16, 0x6c, 0xa0, 0x14,
	0x72, 0x49, 0x50, 0xb5, 0x20, 0xc8, 0xae, 0xfa, 0x31, 0x34, 0x07, 0xf9, 0x65, 0x85, 0x8f, 0x57,
	0x3c, 0xd6, 0xde, 0x3c, 0x8a, 0x3d, 0x70, 0xa5, 0xd2, 0xef, 0x55, 0xd8, 0x3e, 0xa2, 0x24, 0x1d,
	0x4f, 0x07, 0xb9, 0xb0, 0x8b, 0x79, 0x04, 0x2d, 0xc9, 0x25, 0x99, 0x0d, 0xc7, 0x3c, 0x4b, 0xa4,
	0x55, 0xc2, 0x9d, 0xd7, 0x4b, 0xbf, 0x0c, 0x5f, 0x2c, 0x7d, 0x64, 0x96, 0x5c, 0x02, 0x03, 0x0c,
	0xda, 0x7b, 0xa8, 0x1c, 0xa5, 0x38, 0x53, 0x41, 0xeb, 0x02, 0x1b, 0x47, 0x55, 0x9f, 0x13, 0x46,
	0x87, 0x49, 0x16, 0x8f, 0x68, 0xaa, 0x5f, 0x4f, 0x5b, 0xbd, 0x04, 0x17, 0xd5, 0x4b, 0x60, 0x80,
	0x41, 0x79, 0x5f, 0x6b, 0x07, 0xf5, 0x41, 0x7b, 0x43, 0x7d, 0xa1, 0x7e, 0x6b, 0xeb, 0xfd, 0xbd,
	0xd7, 0x4b, 0xbf, 0x84, 0x16, 0xe2, 0x2d, 0xb0, 0x00, 0x37, 0x95, 0x33, 0x50, 0xb6, 0xea, 0x70,
	0x16, 0xc5, 0x91, 0xd4, 0xcf, 0x72, 0x1d, 0x1b, 0x07, 0x7d, 0x02, 0x35, 0x99, 0x0b, 0xcf, 0xd5,
	0x7c, 0xde, 0x5e, 0xcf, 0x67, 0xf1, 0x67, 0x82, 0xd5, 0x01, 0xc3, 0x68, 0xff, 0xf3, 0x57, 0x7f,
	0xb7, 0x2b, 0x2f, 0xcf, 0xda, 0xce, 0xe9, 0x59, 0xdb, 0xf9, 0xeb, 0xac, 0xed, 0xfc, 0x78, 0xde,
	0xae, 0x9c, 0x9e, 0xb7, 0x2b, 0xaf, 0xce, 0xdb, 0x95, 0x6f, 0x03, 0x16, 0xc9, 0x69, 0x36, 0x0a,
	0xc7, 0x3c, 0xee, 0xd9, 0x3f, 0x47, 0xf3, 0xf3, 0x91, 0x98, 0xbc, 0x30, 0xff, 0x64, 0x23, 0x57,
	0xbf, 0xa2, 0xf7, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x03, 0x85, 0x77, 0x9d, 0x3e, 0x07, 0x00,
	0x00,
}

func (m *TxResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Events) > 0 {
		for iNdEx := len(m.Events) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Events[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAbci(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x6a
		}
	}
	if len(m.Timestamp) > 0 {
		i -= len(m.Timestamp)
		copy(dAtA[i:], m.Timestamp)
		i = encodeVarintAbci(dAtA, i, uint64(len(m.Timestamp)))
		i--
		dAtA[i] = 0x62
	}
	if m.Tx != nil {
		{
			size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAbci(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	if m.GasUsed != 0 {
		i = encodeVarintAbci(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x50
	}
	if m.GasWanted != 0 {
		i = encodeVarintAbci(dAtA, i, uint64(m.GasWanted))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Info) > 0 {
		i -= len(m.Info)
		copy(dAtA[i:], m.Info)
		i = encodeVarintAbci(dAtA, i, uint64(len(m.Info)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Logs) > 0 {
		for iNdEx := len(m.Logs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Logs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAbci(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.RawLog) > 0 {
		i -= len(m.RawLog)
		copy(dAtA[i:], m.RawLog)
		i = encodeVarintAbci(dAtA, i, uint64(len(m.RawLog)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintAbci(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Code != 0 {
		i = encodeVarintAbci(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Codespace) > 0 {
		i -= len(m.Codespace)
		copy(dAtA[i:], m.Codespace)
		i = encodeVarintAbci(dAtA, i, uint64(len(m.Codespace)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintAbci(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Height != 0 {
		i = encodeVarintAbci(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ABCIMessageLog) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ABCIMessageLog) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ABCIMessageLog) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Events) > 0 {
		for iNdEx := len(m.Events) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Events[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAbci(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = encodeVarintAbci(dAtA, i, uint64(len(m.Log)))
		i--
		dAtA[i] = 0x12
	}
	if m.MsgIndex != 0 {
		i = encodeVarintAbci(dAtA, i, uint64(m.MsgIndex))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *StringEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StringEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StringEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Attributes) > 0 {
		for iNdEx := len(m.Attributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAbci(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintAbci(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Attribute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Attribute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Attribute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintAbci(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintAbci(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GasInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GasInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GasInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GasUsed != 0 {
		i = encodeVarintAbci(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x10
	}
	if m.GasWanted != 0 {
		i = encodeVarintAbci(dAtA, i, uint64(m.GasWanted))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Result) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Result) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Result) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Events) > 0 {
		for iNdEx := len(m.Events) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Events[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAbci(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = encodeVarintAbci(dAtA, i, uint64(len(m.Log)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintAbci(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SimulationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SimulationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SimulationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintAbci(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.GasInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAbci(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintAbci(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MsgType) > 0 {
		i -= len(m.MsgType)
		copy(dAtA[i:], m.MsgType)
		i = encodeVarintAbci(dAtA, i, uint64(len(m.MsgType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TxMsgData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxMsgData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxMsgData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		for iNdEx := len(m.Data) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Data[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAbci(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SearchTxsResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SearchTxsResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SearchTxsResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for iNdEx := len(m.Txs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Txs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAbci(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.Limit != 0 {
		i = encodeVarintAbci(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x28
	}
	if m.PageTotal != 0 {
		i = encodeVarintAbci(dAtA, i, uint64(m.PageTotal))
		i--
		dAtA[i] = 0x20
	}
	if m.PageNumber != 0 {
		i = encodeVarintAbci(dAtA, i, uint64(m.PageNumber))
		i--
		dAtA[i] = 0x18
	}
	if m.Count != 0 {
		i = encodeVarintAbci(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x10
	}
	if m.TotalCount != 0 {
		i = encodeVarintAbci(dAtA, i, uint64(m.TotalCount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAbci(dAtA []byte, offset int, v uint64) int {
	offset -= sovAbci(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TxResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovAbci(uint64(m.Height))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovAbci(uint64(l))
	}
	l = len(m.Codespace)
	if l > 0 {
		n += 1 + l + sovAbci(uint64(l))
	}
	if m.Code != 0 {
		n += 1 + sovAbci(uint64(m.Code))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovAbci(uint64(l))
	}
	l = len(m.RawLog)
	if l > 0 {
		n += 1 + l + sovAbci(uint64(l))
	}
	if len(m.Logs) > 0 {
		for _, e := range m.Logs {
			l = e.Size()
			n += 1 + l + sovAbci(uint64(l))
		}
	}
	l = len(m.Info)
	if l > 0 {
		n += 1 + l + sovAbci(uint64(l))
	}
	if m.GasWanted != 0 {
		n += 1 + sovAbci(uint64(m.GasWanted))
	}
	if m.GasUsed != 0 {
		n += 1 + sovAbci(uint64(m.GasUsed))
	}
	if m.Tx != nil {
		l = m.Tx.Size()
		n += 1 + l + sovAbci(uint64(l))
	}
	l = len(m.Timestamp)
	if l > 0 {
		n += 1 + l + sovAbci(uint64(l))
	}
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovAbci(uint64(l))
		}
	}
	return n
}

func (m *ABCIMessageLog) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MsgIndex != 0 {
		n += 1 + sovAbci(uint64(m.MsgIndex))
	}
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovAbci(uint64(l))
	}
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovAbci(uint64(l))
		}
	}
	return n
}

func (m *StringEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovAbci(uint64(l))
	}
	if len(m.Attributes) > 0 {
		for _, e := range m.Attributes {
			l = e.Size()
			n += 1 + l + sovAbci(uint64(l))
		}
	}
	return n
}

func (m *Attribute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovAbci(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovAbci(uint64(l))
	}
	return n
}

func (m *GasInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GasWanted != 0 {
		n += 1 + sovAbci(uint64(m.GasWanted))
	}
	if m.GasUsed != 0 {
		n += 1 + sovAbci(uint64(m.GasUsed))
	}
	return n
}

func (m *Result) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovAbci(uint64(l))
	}
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovAbci(uint64(l))
	}
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovAbci(uint64(l))
		}
	}
	return n
}

func (m *SimulationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.GasInfo.Size()
	n += 1 + l + sovAbci(uint64(l))
	if m.Result != nil {
		l = m.Result.Size()
		n += 1 + l + sovAbci(uint64(l))
	}
	return n
}

func (m *MsgData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MsgType)
	if l > 0 {
		n += 1 + l + sovAbci(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovAbci(uint64(l))
	}
	return n
}

func (m *TxMsgData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovAbci(uint64(l))
		}
	}
	return n
}

func (m *SearchTxsResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TotalCount != 0 {
		n += 1 + sovAbci(uint64(m.TotalCount))
	}
	if m.Count != 0 {
		n += 1 + sovAbci(uint64(m.Count))
	}
	if m.PageNumber != 0 {
		n += 1 + sovAbci(uint64(m.PageNumber))
	}
	if m.PageTotal != 0 {
		n += 1 + sovAbci(uint64(m.PageTotal))
	}
	if m.Limit != 0 {
		n += 1 + sovAbci(uint64(m.Limit))
	}
	if len(m.Txs) > 0 {
		for _, e := range m.Txs {
			l = e.Size()
			n += 1 + l + sovAbci(uint64(l))
		}
	}
	return n
}

func sovAbci(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAbci(x uint64) (n int) {
	return sovAbci(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ABCIMessageLog) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForEvents := "[]StringEvent{"
	for _, f := range this.Events {
		repeatedStringForEvents += strings.Replace(strings.Replace(f.String(), "StringEvent", "StringEvent", 1), `&`, ``, 1) + ","
	}
	repeatedStringForEvents += "}"
	s := strings.Join([]string{`&ABCIMessageLog{`,
		`MsgIndex:` + fmt.Sprintf("%v", this.MsgIndex) + `,`,
		`Log:` + fmt.Sprintf("%v", this.Log) + `,`,
		`Events:` + repeatedStringForEvents + `,`,
		`}`,
	}, "")
	return s
}
func (this *StringEvent) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForAttributes := "[]Attribute{"
	for _, f := range this.Attributes {
		repeatedStringForAttributes += fmt.Sprintf("%v", f) + ","
	}
	repeatedStringForAttributes += "}"
	s := strings.Join([]string{`&StringEvent{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Attributes:` + repeatedStringForAttributes + `,`,
		`}`,
	}, "")
	return s
}
func (this *MsgData) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MsgData{`,
		`MsgType:` + fmt.Sprintf("%v", this.MsgType) + `,`,
		`Data:` + fmt.Sprintf("%v", this.Data) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TxMsgData) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForData := "[]*MsgData{"
	for _, f := range this.Data {
		repeatedStringForData += strings.Replace(f.String(), "MsgData", "MsgData", 1) + ","
	}
	repeatedStringForData += "}"
	s := strings.Join([]string{`&TxMsgData{`,
		`Data:` + repeatedStringForData + `,`,
		`}`,
	}, "")
	return s
}
func (this *SearchTxsResult) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForTxs := "[]*TxResponse{"
	for _, f := range this.Txs {
		repeatedStringForTxs += strings.Replace(fmt.Sprintf("%v", f), "TxResponse", "TxResponse", 1) + ","
	}
	repeatedStringForTxs += "}"
	s := strings.Join([]string{`&SearchTxsResult{`,
		`TotalCount:` + fmt.Sprintf("%v", this.TotalCount) + `,`,
		`Count:` + fmt.Sprintf("%v", this.Count) + `,`,
		`PageNumber:` + fmt.Sprintf("%v", this.PageNumber) + `,`,
		`PageTotal:` + fmt.Sprintf("%v", this.PageTotal) + `,`,
		`Limit:` + fmt.Sprintf("%v", this.Limit) + `,`,
		`Txs:` + repeatedStringForTxs + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAbci(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *TxResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAbci
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
			return fmt.Errorf("proto: TxResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Codespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Codespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawLog", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RawLog = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Logs = append(m.Logs, ABCIMessageLog{})
			if err := m.Logs[len(m.Logs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasWanted", wireType)
			}
			m.GasWanted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasWanted |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tx == nil {
				m.Tx = &types.Any{}
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timestamp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, types1.Event{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAbci(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAbci
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
func (m *ABCIMessageLog) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAbci
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
			return fmt.Errorf("proto: ABCIMessageLog: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ABCIMessageLog: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgIndex", wireType)
			}
			m.MsgIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MsgIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Log = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, StringEvent{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAbci(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAbci
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
func (m *StringEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAbci
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
			return fmt.Errorf("proto: StringEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StringEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = append(m.Attributes, Attribute{})
			if err := m.Attributes[len(m.Attributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAbci(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAbci
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
func (m *Attribute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAbci
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
			return fmt.Errorf("proto: Attribute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Attribute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAbci(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAbci
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
func (m *GasInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAbci
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
			return fmt.Errorf("proto: GasInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GasInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasWanted", wireType)
			}
			m.GasWanted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasWanted |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAbci(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAbci
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
func (m *Result) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAbci
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
			return fmt.Errorf("proto: Result: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Result: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Log = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, types1.Event{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAbci(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAbci
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
func (m *SimulationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAbci
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
			return fmt.Errorf("proto: SimulationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SimulationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
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
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Result == nil {
				m.Result = &Result{}
			}
			if err := m.Result.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAbci(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAbci
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
func (m *MsgData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAbci
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
			return fmt.Errorf("proto: MsgData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAbci(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAbci
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
func (m *TxMsgData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAbci
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
			return fmt.Errorf("proto: TxMsgData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxMsgData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &MsgData{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAbci(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAbci
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
func (m *SearchTxsResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAbci
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
			return fmt.Errorf("proto: SearchTxsResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SearchTxsResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalCount", wireType)
			}
			m.TotalCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PageNumber", wireType)
			}
			m.PageNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PageNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PageTotal", wireType)
			}
			m.PageTotal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PageTotal |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbci
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
				return ErrInvalidLengthAbci
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAbci
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txs = append(m.Txs, &TxResponse{})
			if err := m.Txs[len(m.Txs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAbci(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAbci
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
func skipAbci(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAbci
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
					return 0, ErrIntOverflowAbci
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
					return 0, ErrIntOverflowAbci
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
				return 0, ErrInvalidLengthAbci
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAbci
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAbci
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAbci        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAbci          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAbci = fmt.Errorf("proto: unexpected end of group")
)
