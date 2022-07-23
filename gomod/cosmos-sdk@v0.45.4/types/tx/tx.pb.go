


package tx

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	types1 "github.com/cosmos/cosmos-sdk/crypto/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types2 "github.com/cosmos/cosmos-sdk/types"
	signing "github.com/cosmos/cosmos-sdk/types/tx/signing"
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


type Tx struct {
	
	Body *TxBody `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	
	
	AuthInfo *AuthInfo `protobuf:"bytes,2,opt,name=auth_info,json=authInfo,proto3" json:"auth_info,omitempty"`
	
	
	
	Signatures [][]byte `protobuf:"bytes,3,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (m *Tx) Reset()         { *m = Tx{} }
func (m *Tx) String() string { return proto.CompactTextString(m) }
func (*Tx) ProtoMessage()    {}
func (*Tx) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d1575ffde80842, []int{0}
}
func (m *Tx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Tx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tx.Merge(m, src)
}
func (m *Tx) XXX_Size() int {
	return m.Size()
}
func (m *Tx) XXX_DiscardUnknown() {
	xxx_messageInfo_Tx.DiscardUnknown(m)
}

var xxx_messageInfo_Tx proto.InternalMessageInfo

func (m *Tx) GetBody() *TxBody {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Tx) GetAuthInfo() *AuthInfo {
	if m != nil {
		return m.AuthInfo
	}
	return nil
}

func (m *Tx) GetSignatures() [][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}






type TxRaw struct {
	
	
	BodyBytes []byte `protobuf:"bytes,1,opt,name=body_bytes,json=bodyBytes,proto3" json:"body_bytes,omitempty"`
	
	
	AuthInfoBytes []byte `protobuf:"bytes,2,opt,name=auth_info_bytes,json=authInfoBytes,proto3" json:"auth_info_bytes,omitempty"`
	
	
	
	Signatures [][]byte `protobuf:"bytes,3,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (m *TxRaw) Reset()         { *m = TxRaw{} }
func (m *TxRaw) String() string { return proto.CompactTextString(m) }
func (*TxRaw) ProtoMessage()    {}
func (*TxRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d1575ffde80842, []int{1}
}
func (m *TxRaw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxRaw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxRaw.Merge(m, src)
}
func (m *TxRaw) XXX_Size() int {
	return m.Size()
}
func (m *TxRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_TxRaw.DiscardUnknown(m)
}

var xxx_messageInfo_TxRaw proto.InternalMessageInfo

func (m *TxRaw) GetBodyBytes() []byte {
	if m != nil {
		return m.BodyBytes
	}
	return nil
}

func (m *TxRaw) GetAuthInfoBytes() []byte {
	if m != nil {
		return m.AuthInfoBytes
	}
	return nil
}

func (m *TxRaw) GetSignatures() [][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}


type SignDoc struct {
	
	
	BodyBytes []byte `protobuf:"bytes,1,opt,name=body_bytes,json=bodyBytes,proto3" json:"body_bytes,omitempty"`
	
	
	AuthInfoBytes []byte `protobuf:"bytes,2,opt,name=auth_info_bytes,json=authInfoBytes,proto3" json:"auth_info_bytes,omitempty"`
	
	
	
	ChainId string `protobuf:"bytes,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	
	AccountNumber uint64 `protobuf:"varint,4,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
}

func (m *SignDoc) Reset()         { *m = SignDoc{} }
func (m *SignDoc) String() string { return proto.CompactTextString(m) }
func (*SignDoc) ProtoMessage()    {}
func (*SignDoc) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d1575ffde80842, []int{2}
}
func (m *SignDoc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignDoc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignDoc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignDoc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignDoc.Merge(m, src)
}
func (m *SignDoc) XXX_Size() int {
	return m.Size()
}
func (m *SignDoc) XXX_DiscardUnknown() {
	xxx_messageInfo_SignDoc.DiscardUnknown(m)
}

var xxx_messageInfo_SignDoc proto.InternalMessageInfo

func (m *SignDoc) GetBodyBytes() []byte {
	if m != nil {
		return m.BodyBytes
	}
	return nil
}

func (m *SignDoc) GetAuthInfoBytes() []byte {
	if m != nil {
		return m.AuthInfoBytes
	}
	return nil
}

func (m *SignDoc) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *SignDoc) GetAccountNumber() uint64 {
	if m != nil {
		return m.AccountNumber
	}
	return 0
}


type TxBody struct {
	
	
	
	
	
	
	
	Messages []*types.Any `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	
	
	
	Memo string `protobuf:"bytes,2,opt,name=memo,proto3" json:"memo,omitempty"`
	
	
	TimeoutHeight uint64 `protobuf:"varint,3,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height,omitempty"`
	
	
	
	ExtensionOptions []*types.Any `protobuf:"bytes,1023,rep,name=extension_options,json=extensionOptions,proto3" json:"extension_options,omitempty"`
	
	
	
	NonCriticalExtensionOptions []*types.Any `protobuf:"bytes,2047,rep,name=non_critical_extension_options,json=nonCriticalExtensionOptions,proto3" json:"non_critical_extension_options,omitempty"`
}

func (m *TxBody) Reset()         { *m = TxBody{} }
func (m *TxBody) String() string { return proto.CompactTextString(m) }
func (*TxBody) ProtoMessage()    {}
func (*TxBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d1575ffde80842, []int{3}
}
func (m *TxBody) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxBody.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxBody.Merge(m, src)
}
func (m *TxBody) XXX_Size() int {
	return m.Size()
}
func (m *TxBody) XXX_DiscardUnknown() {
	xxx_messageInfo_TxBody.DiscardUnknown(m)
}

var xxx_messageInfo_TxBody proto.InternalMessageInfo

func (m *TxBody) GetMessages() []*types.Any {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *TxBody) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *TxBody) GetTimeoutHeight() uint64 {
	if m != nil {
		return m.TimeoutHeight
	}
	return 0
}

func (m *TxBody) GetExtensionOptions() []*types.Any {
	if m != nil {
		return m.ExtensionOptions
	}
	return nil
}

func (m *TxBody) GetNonCriticalExtensionOptions() []*types.Any {
	if m != nil {
		return m.NonCriticalExtensionOptions
	}
	return nil
}



type AuthInfo struct {
	
	
	
	
	SignerInfos []*SignerInfo `protobuf:"bytes,1,rep,name=signer_infos,json=signerInfos,proto3" json:"signer_infos,omitempty"`
	
	
	
	
	Fee *Fee `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (m *AuthInfo) Reset()         { *m = AuthInfo{} }
func (m *AuthInfo) String() string { return proto.CompactTextString(m) }
func (*AuthInfo) ProtoMessage()    {}
func (*AuthInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d1575ffde80842, []int{4}
}
func (m *AuthInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthInfo.Merge(m, src)
}
func (m *AuthInfo) XXX_Size() int {
	return m.Size()
}
func (m *AuthInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AuthInfo proto.InternalMessageInfo

func (m *AuthInfo) GetSignerInfos() []*SignerInfo {
	if m != nil {
		return m.SignerInfos
	}
	return nil
}

func (m *AuthInfo) GetFee() *Fee {
	if m != nil {
		return m.Fee
	}
	return nil
}



type SignerInfo struct {
	
	
	
	PublicKey *types.Any `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	
	
	ModeInfo *ModeInfo `protobuf:"bytes,2,opt,name=mode_info,json=modeInfo,proto3" json:"mode_info,omitempty"`
	
	
	
	Sequence uint64 `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *SignerInfo) Reset()         { *m = SignerInfo{} }
func (m *SignerInfo) String() string { return proto.CompactTextString(m) }
func (*SignerInfo) ProtoMessage()    {}
func (*SignerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d1575ffde80842, []int{5}
}
func (m *SignerInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignerInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignerInfo.Merge(m, src)
}
func (m *SignerInfo) XXX_Size() int {
	return m.Size()
}
func (m *SignerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SignerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SignerInfo proto.InternalMessageInfo

func (m *SignerInfo) GetPublicKey() *types.Any {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *SignerInfo) GetModeInfo() *ModeInfo {
	if m != nil {
		return m.ModeInfo
	}
	return nil
}

func (m *SignerInfo) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}


type ModeInfo struct {
	
	
	//
	
	
	
	Sum isModeInfo_Sum `protobuf_oneof:"sum"`
}

func (m *ModeInfo) Reset()         { *m = ModeInfo{} }
func (m *ModeInfo) String() string { return proto.CompactTextString(m) }
func (*ModeInfo) ProtoMessage()    {}
func (*ModeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d1575ffde80842, []int{6}
}
func (m *ModeInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ModeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ModeInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ModeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModeInfo.Merge(m, src)
}
func (m *ModeInfo) XXX_Size() int {
	return m.Size()
}
func (m *ModeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ModeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ModeInfo proto.InternalMessageInfo

type isModeInfo_Sum interface {
	isModeInfo_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ModeInfo_Single_ struct {
	Single *ModeInfo_Single `protobuf:"bytes,1,opt,name=single,proto3,oneof" json:"single,omitempty"`
}
type ModeInfo_Multi_ struct {
	Multi *ModeInfo_Multi `protobuf:"bytes,2,opt,name=multi,proto3,oneof" json:"multi,omitempty"`
}

func (*ModeInfo_Single_) isModeInfo_Sum() {}
func (*ModeInfo_Multi_) isModeInfo_Sum()  {}

func (m *ModeInfo) GetSum() isModeInfo_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *ModeInfo) GetSingle() *ModeInfo_Single {
	if x, ok := m.GetSum().(*ModeInfo_Single_); ok {
		return x.Single
	}
	return nil
}

func (m *ModeInfo) GetMulti() *ModeInfo_Multi {
	if x, ok := m.GetSum().(*ModeInfo_Multi_); ok {
		return x.Multi
	}
	return nil
}


func (*ModeInfo) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ModeInfo_Single_)(nil),
		(*ModeInfo_Multi_)(nil),
	}
}




type ModeInfo_Single struct {
	
	Mode signing.SignMode `protobuf:"varint,1,opt,name=mode,proto3,enum=cosmos.tx.signing.v1beta1.SignMode" json:"mode,omitempty"`
}

func (m *ModeInfo_Single) Reset()         { *m = ModeInfo_Single{} }
func (m *ModeInfo_Single) String() string { return proto.CompactTextString(m) }
func (*ModeInfo_Single) ProtoMessage()    {}
func (*ModeInfo_Single) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d1575ffde80842, []int{6, 0}
}
func (m *ModeInfo_Single) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ModeInfo_Single) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ModeInfo_Single.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ModeInfo_Single) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModeInfo_Single.Merge(m, src)
}
func (m *ModeInfo_Single) XXX_Size() int {
	return m.Size()
}
func (m *ModeInfo_Single) XXX_DiscardUnknown() {
	xxx_messageInfo_ModeInfo_Single.DiscardUnknown(m)
}

var xxx_messageInfo_ModeInfo_Single proto.InternalMessageInfo

func (m *ModeInfo_Single) GetMode() signing.SignMode {
	if m != nil {
		return m.Mode
	}
	return signing.SignMode_SIGN_MODE_UNSPECIFIED
}


type ModeInfo_Multi struct {
	
	Bitarray *types1.CompactBitArray `protobuf:"bytes,1,opt,name=bitarray,proto3" json:"bitarray,omitempty"`
	
	
	ModeInfos []*ModeInfo `protobuf:"bytes,2,rep,name=mode_infos,json=modeInfos,proto3" json:"mode_infos,omitempty"`
}

func (m *ModeInfo_Multi) Reset()         { *m = ModeInfo_Multi{} }
func (m *ModeInfo_Multi) String() string { return proto.CompactTextString(m) }
func (*ModeInfo_Multi) ProtoMessage()    {}
func (*ModeInfo_Multi) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d1575ffde80842, []int{6, 1}
}
func (m *ModeInfo_Multi) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ModeInfo_Multi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ModeInfo_Multi.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ModeInfo_Multi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModeInfo_Multi.Merge(m, src)
}
func (m *ModeInfo_Multi) XXX_Size() int {
	return m.Size()
}
func (m *ModeInfo_Multi) XXX_DiscardUnknown() {
	xxx_messageInfo_ModeInfo_Multi.DiscardUnknown(m)
}

var xxx_messageInfo_ModeInfo_Multi proto.InternalMessageInfo

func (m *ModeInfo_Multi) GetBitarray() *types1.CompactBitArray {
	if m != nil {
		return m.Bitarray
	}
	return nil
}

func (m *ModeInfo_Multi) GetModeInfos() []*ModeInfo {
	if m != nil {
		return m.ModeInfos
	}
	return nil
}




type Fee struct {
	
	Amount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	
	
	GasLimit uint64 `protobuf:"varint,2,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	
	
	
	Payer string `protobuf:"bytes,3,opt,name=payer,proto3" json:"payer,omitempty"`
	
	
	
	Granter string `protobuf:"bytes,4,opt,name=granter,proto3" json:"granter,omitempty"`
}

func (m *Fee) Reset()         { *m = Fee{} }
func (m *Fee) String() string { return proto.CompactTextString(m) }
func (*Fee) ProtoMessage()    {}
func (*Fee) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d1575ffde80842, []int{7}
}
func (m *Fee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fee.Merge(m, src)
}
func (m *Fee) XXX_Size() int {
	return m.Size()
}
func (m *Fee) XXX_DiscardUnknown() {
	xxx_messageInfo_Fee.DiscardUnknown(m)
}

var xxx_messageInfo_Fee proto.InternalMessageInfo

func (m *Fee) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *Fee) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *Fee) GetPayer() string {
	if m != nil {
		return m.Payer
	}
	return ""
}

func (m *Fee) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func init() {
	proto.RegisterType((*Tx)(nil), "cosmos.tx.v1beta1.Tx")
	proto.RegisterType((*TxRaw)(nil), "cosmos.tx.v1beta1.TxRaw")
	proto.RegisterType((*SignDoc)(nil), "cosmos.tx.v1beta1.SignDoc")
	proto.RegisterType((*TxBody)(nil), "cosmos.tx.v1beta1.TxBody")
	proto.RegisterType((*AuthInfo)(nil), "cosmos.tx.v1beta1.AuthInfo")
	proto.RegisterType((*SignerInfo)(nil), "cosmos.tx.v1beta1.SignerInfo")
	proto.RegisterType((*ModeInfo)(nil), "cosmos.tx.v1beta1.ModeInfo")
	proto.RegisterType((*ModeInfo_Single)(nil), "cosmos.tx.v1beta1.ModeInfo.Single")
	proto.RegisterType((*ModeInfo_Multi)(nil), "cosmos.tx.v1beta1.ModeInfo.Multi")
	proto.RegisterType((*Fee)(nil), "cosmos.tx.v1beta1.Fee")
}

func init() { proto.RegisterFile("cosmos/tx/v1beta1/tx.proto", fileDescriptor_96d1575ffde80842) }

var fileDescriptor_96d1575ffde80842 = []byte{
	
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0x5e, 0xef, 0x5f, 0xd6, 0x27, 0x49, 0x4b, 0x47, 0x11, 0xda, 0x6c, 0x54, 0x37, 0x18, 0x15,
	0xf6, 0x26, 0x76, 0x9b, 0x5e, 0xf0, 0x23, 0x24, 0xc8, 0x16, 0xaa, 0x54, 0xa5, 0x20, 0x4d, 0x72,
	0xd5, 0x1b, 0x6b, 0xec, 0x9d, 0x78, 0x47, 0x5d, 0xcf, 0x2c, 0x9e, 0x71, 0xb1, 0x1f, 0x02, 0xa9,
	0x42, 0x42, 0xbc, 0x03, 0x2f, 0xc0, 0x2b, 0xf4, 0xb2, 0x97, 0x5c, 0x41, 0x95, 0x3c, 0x08, 0x68,
	0xc6, 0x63, 0x27, 0x82, 0x55, 0x72, 0xc3, 0x95, 0xe7, 0x9c, 0xf9, 0xce, 0x37, 0x9f, 0xcf, 0x1f,
	0x4c, 0x12, 0x21, 0x33, 0x21, 0x43, 0x55, 0x86, 0xaf, 0x1e, 0xc6, 0x54, 0x91, 0x87, 0xa1, 0x2a,
	0x83, 0x55, 0x2e, 0x94, 0x40, 0x77, 0xea, 0xbb, 0x40, 0x95, 0x81, 0xbd, 0x9b, 0xec, 0xa4, 0x22,
	0x15, 0xe6, 0x36, 0xd4, 0xa7, 0x1a, 0x38, 0x39, 0xb0, 0x24, 0x49, 0x5e, 0xad, 0x94, 0x08, 0xb3,
	0x62, 0xa9, 0x98, 0x64, 0x69, 0xcb, 0xd8, 0x38, 0x2c, 0xdc, 0xb3, 0xf0, 0x98, 0x48, 0xda, 0x62,
	0x12, 0xc1, 0xb8, 0xbd, 0xff, 0xf8, 0x52, 0x93, 0x64, 0x29, 0x67, 0xfc, 0x92, 0xc9, 0xda, 0x16,
	0xb8, 0x9b, 0x0a, 0x91, 0x2e, 0x69, 0x68, 0xac, 0xb8, 0x38, 0x0b, 0x09, 0xaf, 0xea, 0x2b, 0xff,
	0x27, 0x07, 0xba, 0xa7, 0x25, 0x3a, 0x80, 0x7e, 0x2c, 0xe6, 0xd5, 0xd8, 0xd9, 0x77, 0xa6, 0x9b,
	0x87, 0xbb, 0xc1, 0x7f, 0xfe, 0x28, 0x38, 0x2d, 0x67, 0x62, 0x5e, 0x61, 0x03, 0x43, 0x9f, 0x82,
	0x4b, 0x0a, 0xb5, 0x88, 0x18, 0x3f, 0x13, 0xe3, 0xae, 0x89, 0xd9, 0x5b, 0x13, 0x73, 0x54, 0xa8,
	0xc5, 0x53, 0x7e, 0x26, 0xf0, 0x88, 0xd8, 0x13, 0xf2, 0x00, 0xb4, 0x36, 0xa2, 0x8a, 0x9c, 0xca,
	0x71, 0x6f, 0xbf, 0x37, 0xdd, 0xc2, 0x57, 0x3c, 0x3e, 0x87, 0xc1, 0x69, 0x89, 0xc9, 0x8f, 0xe8,
	0x2e, 0x80, 0x7e, 0x2a, 0x8a, 0x2b, 0x45, 0xa5, 0xd1, 0xb5, 0x85, 0x5d, 0xed, 0x99, 0x69, 0x07,
	0xfa, 0x08, 0x6e, 0xb7, 0x0a, 0x2c, 0xa6, 0x6b, 0x30, 0xdb, 0xcd, 0x53, 0x35, 0xee, 0xa6, 0xf7,
	0x7e, 0x76, 0x60, 0xe3, 0x84, 0xa5, 0xfc, 0x6b, 0x91, 0xfc, 0x5f, 0x4f, 0xee, 0xc2, 0x28, 0x59,
	0x10, 0xc6, 0x23, 0x36, 0x1f, 0xf7, 0xf6, 0x9d, 0xa9, 0x8b, 0x37, 0x8c, 0xfd, 0x74, 0x8e, 0xee,
	0xc3, 0x2d, 0x92, 0x24, 0xa2, 0xe0, 0x2a, 0xe2, 0x45, 0x16, 0xd3, 0x7c, 0xdc, 0xdf, 0x77, 0xa6,
	0x7d, 0xbc, 0x6d, 0xbd, 0xdf, 0x19, 0xa7, 0xff, 0x4b, 0x17, 0x86, 0x75, 0xbe, 0xd1, 0x03, 0x18,
	0x65, 0x54, 0x4a, 0x92, 0x1a, 0x45, 0xbd, 0xe9, 0xe6, 0xe1, 0x4e, 0x50, 0x57, 0x33, 0x68, 0xaa,
	0x19, 0x1c, 0xf1, 0x0a, 0xb7, 0x28, 0x84, 0xa0, 0x9f, 0xd1, 0xac, 0x2e, 0x8b, 0x8b, 0xcd, 0x59,
	0xbf, 0xab, 0x58, 0x46, 0x45, 0xa1, 0xa2, 0x05, 0x65, 0xe9, 0x42, 0x19, 0x61, 0x7d, 0xbc, 0x6d,
	0xbd, 0xc7, 0xc6, 0x89, 0x66, 0x70, 0x87, 0x96, 0x8a, 0x72, 0xc9, 0x04, 0x8f, 0xc4, 0x4a, 0x31,
	0xc1, 0xe5, 0xf8, 0xef, 0x8d, 0x6b, 0x9e, 0x7d, 0xaf, 0xc5, 0x7f, 0x5f, 0xc3, 0xd1, 0x0b, 0xf0,
	0xb8, 0xe0, 0x51, 0x92, 0x33, 0xc5, 0x12, 0xb2, 0x8c, 0xd6, 0x10, 0xde, 0xbe, 0x86, 0x70, 0x8f,
	0x0b, 0xfe, 0xd8, 0xc6, 0x7e, 0xf3, 0x2f, 0x6e, 0xff, 0x15, 0x8c, 0x9a, 0x96, 0x42, 0x5f, 0xc1,
	0x96, 0x2e, 0x23, 0xcd, 0x4d, 0x3d, 0x9a, 0xe4, 0xdc, 0x5d, 0xd3, 0x85, 0x27, 0x06, 0x66, 0xfa,
	0x70, 0x53, 0xb6, 0x67, 0x89, 0xa6, 0xd0, 0x3b, 0xa3, 0xd4, 0xb6, 0xef, 0xfb, 0x6b, 0x02, 0x9f,
	0x50, 0x8a, 0x35, 0xc4, 0xff, 0xd5, 0x01, 0xb8, 0x64, 0x41, 0x8f, 0x00, 0x56, 0x45, 0xbc, 0x64,
	0x49, 0xf4, 0x92, 0x36, 0x23, 0xb3, 0xfe, 0x6f, 0xdc, 0x1a, 0xf7, 0x8c, 0x9a, 0x91, 0xc9, 0xc4,
	0x9c, 0xde, 0x34, 0x32, 0xcf, 0xc5, 0x9c, 0xd6, 0x23, 0x93, 0xd9, 0x13, 0x9a, 0xc0, 0x48, 0xd2,
	0x1f, 0x0a, 0xca, 0x13, 0x6a, 0xcb, 0xd6, 0xda, 0xfe, 0xbb, 0x2e, 0x8c, 0x9a, 0x10, 0xf4, 0x05,
	0x0c, 0x25, 0xe3, 0xe9, 0x92, 0x5a, 0x4d, 0xfe, 0x35, 0xfc, 0xc1, 0x89, 0x41, 0x1e, 0x77, 0xb0,
	0x8d, 0x41, 0x9f, 0xc1, 0xc0, 0xec, 0x1f, 0x2b, 0xee, 0x83, 0xeb, 0x82, 0x9f, 0x6b, 0xe0, 0x71,
	0x07, 0xd7, 0x11, 0x93, 0x23, 0x18, 0xd6, 0x74, 0xe8, 0x13, 0xe8, 0x6b, 0xdd, 0x46, 0xc0, 0xad,
	0xc3, 0x0f, 0xaf, 0x70, 0x34, 0x1b, 0xe9, 0x6a, 0x55, 0x34, 0x1f, 0x36, 0x01, 0x93, 0xd7, 0x0e,
	0x0c, 0x0c, 0x2b, 0x7a, 0x06, 0xa3, 0x98, 0x29, 0x92, 0xe7, 0xa4, 0xc9, 0x6d, 0xd8, 0xd0, 0xd4,
	0x7b, 0x33, 0x68, 0xd7, 0x64, 0xc3, 0xf5, 0x58, 0x64, 0x2b, 0x92, 0xa8, 0x19, 0x53, 0x47, 0x3a,
	0x0c, 0xb7, 0x04, 0xe8, 0x73, 0x80, 0x36, 0xeb, 0x7a, 0x5c, 0x7b, 0x37, 0xa5, 0xdd, 0x6d, 0xd2,
	0x2e, 0x67, 0x03, 0xe8, 0xc9, 0x22, 0xf3, 0x7f, 0x77, 0xa0, 0xf7, 0x84, 0x52, 0x94, 0xc0, 0x90,
	0x64, 0x7a, 0x48, 0x6d, 0xab, 0xb5, 0x4b, 0x52, 0xaf, 0xe7, 0x2b, 0x52, 0x18, 0x9f, 0x3d, 0x78,
	0xf3, 0xe7, 0xbd, 0xce, 0x6f, 0x7f, 0xdd, 0x9b, 0xa6, 0x4c, 0x2d, 0x8a, 0x38, 0x48, 0x44, 0x16,
	0x36, 0xab, 0xdf, 0x7c, 0x0e, 0xe4, 0xfc, 0x65, 0xa8, 0xaa, 0x15, 0x95, 0x26, 0x40, 0x62, 0x4b,
	0x8d, 0xf6, 0xc0, 0x4d, 0x89, 0x8c, 0x96, 0x2c, 0x63, 0xca, 0x14, 0xa2, 0x8f, 0x47, 0x29, 0x91,
	0xdf, 0x6a, 0x1b, 0xed, 0xc0, 0x60, 0x45, 0x2a, 0x9a, 0xdb, 0xad, 0x52, 0x1b, 0x68, 0x0c, 0x1b,
	0x69, 0x4e, 0xb8, 0xb2, 0xcb, 0xc4, 0xc5, 0x8d, 0x39, 0xfb, 0xf2, 0xcd, 0xb9, 0xe7, 0xbc, 0x3d,
	0xf7, 0x9c, 0x77, 0xe7, 0x9e, 0xf3, 0xfa, 0xc2, 0xeb, 0xbc, 0xbd, 0xf0, 0x3a, 0x7f, 0x5c, 0x78,
	0x9d, 0x17, 0xf7, 0x6f, 0x16, 0x16, 0xaa, 0x32, 0x1e, 0x9a, 0x66, 0x7e, 0xf4, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd4, 0xb7, 0x75, 0x7d, 0xfd, 0x06, 0x00, 0x00,
}

func (m *Tx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Tx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signatures) > 0 {
		for iNdEx := len(m.Signatures) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signatures[iNdEx])
			copy(dAtA[i:], m.Signatures[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Signatures[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.AuthInfo != nil {
		{
			size, err := m.AuthInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Body != nil {
		{
			size, err := m.Body.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TxRaw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxRaw) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxRaw) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signatures) > 0 {
		for iNdEx := len(m.Signatures) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signatures[iNdEx])
			copy(dAtA[i:], m.Signatures[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Signatures[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.AuthInfoBytes) > 0 {
		i -= len(m.AuthInfoBytes)
		copy(dAtA[i:], m.AuthInfoBytes)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AuthInfoBytes)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BodyBytes) > 0 {
		i -= len(m.BodyBytes)
		copy(dAtA[i:], m.BodyBytes)
		i = encodeVarintTx(dAtA, i, uint64(len(m.BodyBytes)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SignDoc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignDoc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignDoc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AccountNumber != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.AccountNumber))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AuthInfoBytes) > 0 {
		i -= len(m.AuthInfoBytes)
		copy(dAtA[i:], m.AuthInfoBytes)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AuthInfoBytes)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BodyBytes) > 0 {
		i -= len(m.BodyBytes)
		copy(dAtA[i:], m.BodyBytes)
		i = encodeVarintTx(dAtA, i, uint64(len(m.BodyBytes)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TxBody) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxBody) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxBody) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NonCriticalExtensionOptions) > 0 {
		for iNdEx := len(m.NonCriticalExtensionOptions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NonCriticalExtensionOptions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x7f
			i--
			dAtA[i] = 0xfa
		}
	}
	if len(m.ExtensionOptions) > 0 {
		for iNdEx := len(m.ExtensionOptions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExtensionOptions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3f
			i--
			dAtA[i] = 0xfa
		}
	}
	if m.TimeoutHeight != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeoutHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Memo) > 0 {
		i -= len(m.Memo)
		copy(dAtA[i:], m.Memo)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Memo)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Messages) > 0 {
		for iNdEx := len(m.Messages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Messages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *AuthInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Fee != nil {
		{
			size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.SignerInfos) > 0 {
		for iNdEx := len(m.SignerInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SignerInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SignerInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignerInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignerInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x18
	}
	if m.ModeInfo != nil {
		{
			size, err := m.ModeInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.PublicKey != nil {
		{
			size, err := m.PublicKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ModeInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ModeInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModeInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *ModeInfo_Single_) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModeInfo_Single_) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Single != nil {
		{
			size, err := m.Single.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *ModeInfo_Multi_) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModeInfo_Multi_) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Multi != nil {
		{
			size, err := m.Multi.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *ModeInfo_Single) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ModeInfo_Single) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModeInfo_Single) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Mode != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Mode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ModeInfo_Multi) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ModeInfo_Multi) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModeInfo_Multi) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ModeInfos) > 0 {
		for iNdEx := len(m.ModeInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ModeInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Bitarray != nil {
		{
			size, err := m.Bitarray.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Fee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Payer) > 0 {
		i -= len(m.Payer)
		copy(dAtA[i:], m.Payer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Payer)))
		i--
		dAtA[i] = 0x1a
	}
	if m.GasLimit != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GasLimit))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Tx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Body != nil {
		l = m.Body.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.AuthInfo != nil {
		l = m.AuthInfo.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Signatures) > 0 {
		for _, b := range m.Signatures {
			l = len(b)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *TxRaw) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BodyBytes)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.AuthInfoBytes)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Signatures) > 0 {
		for _, b := range m.Signatures {
			l = len(b)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *SignDoc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BodyBytes)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.AuthInfoBytes)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.AccountNumber != 0 {
		n += 1 + sovTx(uint64(m.AccountNumber))
	}
	return n
}

func (m *TxBody) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for _, e := range m.Messages {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TimeoutHeight != 0 {
		n += 1 + sovTx(uint64(m.TimeoutHeight))
	}
	if len(m.ExtensionOptions) > 0 {
		for _, e := range m.ExtensionOptions {
			l = e.Size()
			n += 2 + l + sovTx(uint64(l))
		}
	}
	if len(m.NonCriticalExtensionOptions) > 0 {
		for _, e := range m.NonCriticalExtensionOptions {
			l = e.Size()
			n += 2 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *AuthInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SignerInfos) > 0 {
		for _, e := range m.SignerInfos {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.Fee != nil {
		l = m.Fee.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *SignerInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PublicKey != nil {
		l = m.PublicKey.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.ModeInfo != nil {
		l = m.ModeInfo.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Sequence != 0 {
		n += 1 + sovTx(uint64(m.Sequence))
	}
	return n
}

func (m *ModeInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *ModeInfo_Single_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Single != nil {
		l = m.Single.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}
func (m *ModeInfo_Multi_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Multi != nil {
		l = m.Multi.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}
func (m *ModeInfo_Single) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Mode != 0 {
		n += 1 + sovTx(uint64(m.Mode))
	}
	return n
}

func (m *ModeInfo_Multi) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Bitarray != nil {
		l = m.Bitarray.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.ModeInfos) > 0 {
		for _, e := range m.ModeInfos {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *Fee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.GasLimit != 0 {
		n += 1 + sovTx(uint64(m.GasLimit))
	}
	l = len(m.Payer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Tx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: Tx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Body == nil {
				m.Body = &TxBody{}
			}
			if err := m.Body.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthInfo == nil {
				m.AuthInfo = &AuthInfo{}
			}
			if err := m.AuthInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signatures = append(m.Signatures, make([]byte, postIndex-iNdEx))
			copy(m.Signatures[len(m.Signatures)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *TxRaw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: TxRaw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxRaw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BodyBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BodyBytes = append(m.BodyBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.BodyBytes == nil {
				m.BodyBytes = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthInfoBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthInfoBytes = append(m.AuthInfoBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.AuthInfoBytes == nil {
				m.AuthInfoBytes = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signatures = append(m.Signatures, make([]byte, postIndex-iNdEx))
			copy(m.Signatures[len(m.Signatures)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *SignDoc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: SignDoc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignDoc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BodyBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BodyBytes = append(m.BodyBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.BodyBytes == nil {
				m.BodyBytes = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthInfoBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthInfoBytes = append(m.AuthInfoBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.AuthInfoBytes == nil {
				m.AuthInfoBytes = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountNumber", wireType)
			}
			m.AccountNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccountNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *TxBody) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: TxBody: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxBody: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Messages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Messages = append(m.Messages, &types.Any{})
			if err := m.Messages[len(m.Messages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			m.TimeoutHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 1023:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtensionOptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExtensionOptions = append(m.ExtensionOptions, &types.Any{})
			if err := m.ExtensionOptions[len(m.ExtensionOptions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2047:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NonCriticalExtensionOptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NonCriticalExtensionOptions = append(m.NonCriticalExtensionOptions, &types.Any{})
			if err := m.NonCriticalExtensionOptions[len(m.NonCriticalExtensionOptions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *AuthInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: AuthInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignerInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignerInfos = append(m.SignerInfos, &SignerInfo{})
			if err := m.SignerInfos[len(m.SignerInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fee == nil {
				m.Fee = &Fee{}
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *SignerInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: SignerInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignerInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PublicKey == nil {
				m.PublicKey = &types.Any{}
			}
			if err := m.PublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModeInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ModeInfo == nil {
				m.ModeInfo = &ModeInfo{}
			}
			if err := m.ModeInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *ModeInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: ModeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ModeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Single", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ModeInfo_Single{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &ModeInfo_Single_{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multi", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ModeInfo_Multi{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &ModeInfo_Multi_{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *ModeInfo_Single) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: Single: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Single: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mode", wireType)
			}
			m.Mode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mode |= signing.SignMode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *ModeInfo_Multi) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: Multi: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Multi: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bitarray", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Bitarray == nil {
				m.Bitarray = &types1.CompactBitArray{}
			}
			if err := m.Bitarray.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModeInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModeInfos = append(m.ModeInfos, &ModeInfo{})
			if err := m.ModeInfos[len(m.ModeInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *Fee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: Fee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types2.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLimit", wireType)
			}
			m.GasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
