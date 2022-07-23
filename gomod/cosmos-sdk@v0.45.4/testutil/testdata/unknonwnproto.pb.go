


package testdata

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	tx "github.com/cosmos/cosmos-sdk/types/tx"
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

type Customer2_City int32

const (
	Customer2_Laos       Customer2_City = 0
	Customer2_LosAngeles Customer2_City = 1
	Customer2_PaloAlto   Customer2_City = 2
	Customer2_Moscow     Customer2_City = 3
	Customer2_Nairobi    Customer2_City = 4
)

var Customer2_City_name = map[int32]string{
	0: "Laos",
	1: "LosAngeles",
	2: "PaloAlto",
	3: "Moscow",
	4: "Nairobi",
}

var Customer2_City_value = map[string]int32{
	"Laos":       0,
	"LosAngeles": 1,
	"PaloAlto":   2,
	"Moscow":     3,
	"Nairobi":    4,
}

func (x Customer2_City) String() string {
	return proto.EnumName(Customer2_City_name, int32(x))
}

func (Customer2_City) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{1, 0}
}

type Customer1 struct {
	Id              int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SubscriptionFee float32 `protobuf:"fixed32,3,opt,name=subscription_fee,json=subscriptionFee,proto3" json:"subscription_fee,omitempty"`
	Payment         string  `protobuf:"bytes,7,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (m *Customer1) Reset()         { *m = Customer1{} }
func (m *Customer1) String() string { return proto.CompactTextString(m) }
func (*Customer1) ProtoMessage()    {}
func (*Customer1) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{0}
}
func (m *Customer1) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Customer1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Customer1.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Customer1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer1.Merge(m, src)
}
func (m *Customer1) XXX_Size() int {
	return m.Size()
}
func (m *Customer1) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer1.DiscardUnknown(m)
}

var xxx_messageInfo_Customer1 proto.InternalMessageInfo

func (m *Customer1) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Customer1) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Customer1) GetSubscriptionFee() float32 {
	if m != nil {
		return m.SubscriptionFee
	}
	return 0
}

func (m *Customer1) GetPayment() string {
	if m != nil {
		return m.Payment
	}
	return ""
}

type Customer2 struct {
	Id            int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Industry      int32          `protobuf:"varint,2,opt,name=industry,proto3" json:"industry,omitempty"`
	Name          string         `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Fewer         float32        `protobuf:"fixed32,4,opt,name=fewer,proto3" json:"fewer,omitempty"`
	Reserved      int64          `protobuf:"varint,1047,opt,name=reserved,proto3" json:"reserved,omitempty"`
	City          Customer2_City `protobuf:"varint,6,opt,name=city,proto3,enum=testdata.Customer2_City" json:"city,omitempty"`
	Miscellaneous *types.Any     `protobuf:"bytes,10,opt,name=miscellaneous,proto3" json:"miscellaneous,omitempty"`
}

func (m *Customer2) Reset()         { *m = Customer2{} }
func (m *Customer2) String() string { return proto.CompactTextString(m) }
func (*Customer2) ProtoMessage()    {}
func (*Customer2) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{1}
}
func (m *Customer2) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Customer2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Customer2.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Customer2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer2.Merge(m, src)
}
func (m *Customer2) XXX_Size() int {
	return m.Size()
}
func (m *Customer2) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer2.DiscardUnknown(m)
}

var xxx_messageInfo_Customer2 proto.InternalMessageInfo

func (m *Customer2) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Customer2) GetIndustry() int32 {
	if m != nil {
		return m.Industry
	}
	return 0
}

func (m *Customer2) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Customer2) GetFewer() float32 {
	if m != nil {
		return m.Fewer
	}
	return 0
}

func (m *Customer2) GetReserved() int64 {
	if m != nil {
		return m.Reserved
	}
	return 0
}

func (m *Customer2) GetCity() Customer2_City {
	if m != nil {
		return m.City
	}
	return Customer2_Laos
}

func (m *Customer2) GetMiscellaneous() *types.Any {
	if m != nil {
		return m.Miscellaneous
	}
	return nil
}

type Nested4A struct {
	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Nested4A) Reset()         { *m = Nested4A{} }
func (m *Nested4A) String() string { return proto.CompactTextString(m) }
func (*Nested4A) ProtoMessage()    {}
func (*Nested4A) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{2}
}
func (m *Nested4A) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Nested4A) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Nested4A.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Nested4A) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nested4A.Merge(m, src)
}
func (m *Nested4A) XXX_Size() int {
	return m.Size()
}
func (m *Nested4A) XXX_DiscardUnknown() {
	xxx_messageInfo_Nested4A.DiscardUnknown(m)
}

var xxx_messageInfo_Nested4A proto.InternalMessageInfo

func (m *Nested4A) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Nested4A) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Nested3A struct {
	Id    int32               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	A4    []*Nested4A         `protobuf:"bytes,4,rep,name=a4,proto3" json:"a4,omitempty"`
	Index map[int64]*Nested4A `protobuf:"bytes,5,rep,name=index,proto3" json:"index,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Nested3A) Reset()         { *m = Nested3A{} }
func (m *Nested3A) String() string { return proto.CompactTextString(m) }
func (*Nested3A) ProtoMessage()    {}
func (*Nested3A) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{3}
}
func (m *Nested3A) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Nested3A) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Nested3A.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Nested3A) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nested3A.Merge(m, src)
}
func (m *Nested3A) XXX_Size() int {
	return m.Size()
}
func (m *Nested3A) XXX_DiscardUnknown() {
	xxx_messageInfo_Nested3A.DiscardUnknown(m)
}

var xxx_messageInfo_Nested3A proto.InternalMessageInfo

func (m *Nested3A) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Nested3A) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Nested3A) GetA4() []*Nested4A {
	if m != nil {
		return m.A4
	}
	return nil
}

func (m *Nested3A) GetIndex() map[int64]*Nested4A {
	if m != nil {
		return m.Index
	}
	return nil
}

type Nested2A struct {
	Id     int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Nested *Nested3A `protobuf:"bytes,3,opt,name=nested,proto3" json:"nested,omitempty"`
}

func (m *Nested2A) Reset()         { *m = Nested2A{} }
func (m *Nested2A) String() string { return proto.CompactTextString(m) }
func (*Nested2A) ProtoMessage()    {}
func (*Nested2A) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{4}
}
func (m *Nested2A) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Nested2A) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Nested2A.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Nested2A) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nested2A.Merge(m, src)
}
func (m *Nested2A) XXX_Size() int {
	return m.Size()
}
func (m *Nested2A) XXX_DiscardUnknown() {
	xxx_messageInfo_Nested2A.DiscardUnknown(m)
}

var xxx_messageInfo_Nested2A proto.InternalMessageInfo

func (m *Nested2A) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Nested2A) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Nested2A) GetNested() *Nested3A {
	if m != nil {
		return m.Nested
	}
	return nil
}

type Nested1A struct {
	Id     int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Nested *Nested2A `protobuf:"bytes,2,opt,name=nested,proto3" json:"nested,omitempty"`
}

func (m *Nested1A) Reset()         { *m = Nested1A{} }
func (m *Nested1A) String() string { return proto.CompactTextString(m) }
func (*Nested1A) ProtoMessage()    {}
func (*Nested1A) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{5}
}
func (m *Nested1A) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Nested1A) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Nested1A.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Nested1A) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nested1A.Merge(m, src)
}
func (m *Nested1A) XXX_Size() int {
	return m.Size()
}
func (m *Nested1A) XXX_DiscardUnknown() {
	xxx_messageInfo_Nested1A.DiscardUnknown(m)
}

var xxx_messageInfo_Nested1A proto.InternalMessageInfo

func (m *Nested1A) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Nested1A) GetNested() *Nested2A {
	if m != nil {
		return m.Nested
	}
	return nil
}

type Nested4B struct {
	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Nested4B) Reset()         { *m = Nested4B{} }
func (m *Nested4B) String() string { return proto.CompactTextString(m) }
func (*Nested4B) ProtoMessage()    {}
func (*Nested4B) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{6}
}
func (m *Nested4B) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Nested4B) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Nested4B.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Nested4B) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nested4B.Merge(m, src)
}
func (m *Nested4B) XXX_Size() int {
	return m.Size()
}
func (m *Nested4B) XXX_DiscardUnknown() {
	xxx_messageInfo_Nested4B.DiscardUnknown(m)
}

var xxx_messageInfo_Nested4B proto.InternalMessageInfo

func (m *Nested4B) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Nested4B) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Nested4B) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Nested3B struct {
	Id   int32       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Age  int32       `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Name string      `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	B4   []*Nested4B `protobuf:"bytes,4,rep,name=b4,proto3" json:"b4,omitempty"`
}

func (m *Nested3B) Reset()         { *m = Nested3B{} }
func (m *Nested3B) String() string { return proto.CompactTextString(m) }
func (*Nested3B) ProtoMessage()    {}
func (*Nested3B) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{7}
}
func (m *Nested3B) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Nested3B) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Nested3B.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Nested3B) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nested3B.Merge(m, src)
}
func (m *Nested3B) XXX_Size() int {
	return m.Size()
}
func (m *Nested3B) XXX_DiscardUnknown() {
	xxx_messageInfo_Nested3B.DiscardUnknown(m)
}

var xxx_messageInfo_Nested3B proto.InternalMessageInfo

func (m *Nested3B) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Nested3B) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Nested3B) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Nested3B) GetB4() []*Nested4B {
	if m != nil {
		return m.B4
	}
	return nil
}

type Nested2B struct {
	Id     int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Fee    float64   `protobuf:"fixed64,2,opt,name=fee,proto3" json:"fee,omitempty"`
	Nested *Nested3B `protobuf:"bytes,3,opt,name=nested,proto3" json:"nested,omitempty"`
	Route  string    `protobuf:"bytes,4,opt,name=route,proto3" json:"route,omitempty"`
}

func (m *Nested2B) Reset()         { *m = Nested2B{} }
func (m *Nested2B) String() string { return proto.CompactTextString(m) }
func (*Nested2B) ProtoMessage()    {}
func (*Nested2B) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{8}
}
func (m *Nested2B) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Nested2B) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Nested2B.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Nested2B) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nested2B.Merge(m, src)
}
func (m *Nested2B) XXX_Size() int {
	return m.Size()
}
func (m *Nested2B) XXX_DiscardUnknown() {
	xxx_messageInfo_Nested2B.DiscardUnknown(m)
}

var xxx_messageInfo_Nested2B proto.InternalMessageInfo

func (m *Nested2B) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Nested2B) GetFee() float64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *Nested2B) GetNested() *Nested3B {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *Nested2B) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

type Nested1B struct {
	Id     int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Nested *Nested2B `protobuf:"bytes,2,opt,name=nested,proto3" json:"nested,omitempty"`
	Age    int32     `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
}

func (m *Nested1B) Reset()         { *m = Nested1B{} }
func (m *Nested1B) String() string { return proto.CompactTextString(m) }
func (*Nested1B) ProtoMessage()    {}
func (*Nested1B) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{9}
}
func (m *Nested1B) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Nested1B) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Nested1B.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Nested1B) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nested1B.Merge(m, src)
}
func (m *Nested1B) XXX_Size() int {
	return m.Size()
}
func (m *Nested1B) XXX_DiscardUnknown() {
	xxx_messageInfo_Nested1B.DiscardUnknown(m)
}

var xxx_messageInfo_Nested1B proto.InternalMessageInfo

func (m *Nested1B) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Nested1B) GetNested() *Nested2B {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *Nested1B) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type Customer3 struct {
	Id          int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Sf          float32 `protobuf:"fixed32,3,opt,name=sf,proto3" json:"sf,omitempty"`
	Surcharge   float32 `protobuf:"fixed32,4,opt,name=surcharge,proto3" json:"surcharge,omitempty"`
	Destination string  `protobuf:"bytes,5,opt,name=destination,proto3" json:"destination,omitempty"`
	
	
	
	Payment  isCustomer3_Payment `protobuf_oneof:"payment"`
	Original *Customer1          `protobuf:"bytes,9,opt,name=original,proto3" json:"original,omitempty"`
}

func (m *Customer3) Reset()         { *m = Customer3{} }
func (m *Customer3) String() string { return proto.CompactTextString(m) }
func (*Customer3) ProtoMessage()    {}
func (*Customer3) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{10}
}
func (m *Customer3) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Customer3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Customer3.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Customer3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer3.Merge(m, src)
}
func (m *Customer3) XXX_Size() int {
	return m.Size()
}
func (m *Customer3) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer3.DiscardUnknown(m)
}

var xxx_messageInfo_Customer3 proto.InternalMessageInfo

type isCustomer3_Payment interface {
	isCustomer3_Payment()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Customer3_CreditCardNo struct {
	CreditCardNo string `protobuf:"bytes,7,opt,name=credit_card_no,json=creditCardNo,proto3,oneof" json:"credit_card_no,omitempty"`
}
type Customer3_ChequeNo struct {
	ChequeNo string `protobuf:"bytes,8,opt,name=cheque_no,json=chequeNo,proto3,oneof" json:"cheque_no,omitempty"`
}

func (*Customer3_CreditCardNo) isCustomer3_Payment() {}
func (*Customer3_ChequeNo) isCustomer3_Payment()     {}

func (m *Customer3) GetPayment() isCustomer3_Payment {
	if m != nil {
		return m.Payment
	}
	return nil
}

func (m *Customer3) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Customer3) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Customer3) GetSf() float32 {
	if m != nil {
		return m.Sf
	}
	return 0
}

func (m *Customer3) GetSurcharge() float32 {
	if m != nil {
		return m.Surcharge
	}
	return 0
}

func (m *Customer3) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *Customer3) GetCreditCardNo() string {
	if x, ok := m.GetPayment().(*Customer3_CreditCardNo); ok {
		return x.CreditCardNo
	}
	return ""
}

func (m *Customer3) GetChequeNo() string {
	if x, ok := m.GetPayment().(*Customer3_ChequeNo); ok {
		return x.ChequeNo
	}
	return ""
}

func (m *Customer3) GetOriginal() *Customer1 {
	if m != nil {
		return m.Original
	}
	return nil
}


func (*Customer3) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Customer3_CreditCardNo)(nil),
		(*Customer3_ChequeNo)(nil),
	}
}

type TestVersion1 struct {
	X int64           `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	A *TestVersion1   `protobuf:"bytes,2,opt,name=a,proto3" json:"a,omitempty"`
	B *TestVersion1   `protobuf:"bytes,3,opt,name=b,proto3" json:"b,omitempty"`
	C []*TestVersion1 `protobuf:"bytes,4,rep,name=c,proto3" json:"c,omitempty"`
	D []TestVersion1  `protobuf:"bytes,5,rep,name=d,proto3" json:"d"`
	
	
	
	Sum isTestVersion1_Sum `protobuf_oneof:"sum"`
	G   *types.Any         `protobuf:"bytes,8,opt,name=g,proto3" json:"g,omitempty"`
	H   []*TestVersion1    `protobuf:"bytes,9,rep,name=h,proto3" json:"h,omitempty"`
	
	
	*Customer1 `protobuf:"bytes,12,opt,name=k,proto3,embedded=k" json:"k,omitempty"`
}

func (m *TestVersion1) Reset()         { *m = TestVersion1{} }
func (m *TestVersion1) String() string { return proto.CompactTextString(m) }
func (*TestVersion1) ProtoMessage()    {}
func (*TestVersion1) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{11}
}
func (m *TestVersion1) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestVersion1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestVersion1.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestVersion1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestVersion1.Merge(m, src)
}
func (m *TestVersion1) XXX_Size() int {
	return m.Size()
}
func (m *TestVersion1) XXX_DiscardUnknown() {
	xxx_messageInfo_TestVersion1.DiscardUnknown(m)
}

var xxx_messageInfo_TestVersion1 proto.InternalMessageInfo

type isTestVersion1_Sum interface {
	isTestVersion1_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type TestVersion1_E struct {
	E int32 `protobuf:"varint,6,opt,name=e,proto3,oneof" json:"e,omitempty"`
}
type TestVersion1_F struct {
	F *TestVersion1 `protobuf:"bytes,7,opt,name=f,proto3,oneof" json:"f,omitempty"`
}

func (*TestVersion1_E) isTestVersion1_Sum() {}
func (*TestVersion1_F) isTestVersion1_Sum() {}

func (m *TestVersion1) GetSum() isTestVersion1_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *TestVersion1) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *TestVersion1) GetA() *TestVersion1 {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *TestVersion1) GetB() *TestVersion1 {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *TestVersion1) GetC() []*TestVersion1 {
	if m != nil {
		return m.C
	}
	return nil
}

func (m *TestVersion1) GetD() []TestVersion1 {
	if m != nil {
		return m.D
	}
	return nil
}

func (m *TestVersion1) GetE() int32 {
	if x, ok := m.GetSum().(*TestVersion1_E); ok {
		return x.E
	}
	return 0
}

func (m *TestVersion1) GetF() *TestVersion1 {
	if x, ok := m.GetSum().(*TestVersion1_F); ok {
		return x.F
	}
	return nil
}

func (m *TestVersion1) GetG() *types.Any {
	if m != nil {
		return m.G
	}
	return nil
}

func (m *TestVersion1) GetH() []*TestVersion1 {
	if m != nil {
		return m.H
	}
	return nil
}


func (*TestVersion1) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestVersion1_E)(nil),
		(*TestVersion1_F)(nil),
	}
}

type TestVersion2 struct {
	X int64           `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	A *TestVersion2   `protobuf:"bytes,2,opt,name=a,proto3" json:"a,omitempty"`
	B *TestVersion2   `protobuf:"bytes,3,opt,name=b,proto3" json:"b,omitempty"`
	C []*TestVersion2 `protobuf:"bytes,4,rep,name=c,proto3" json:"c,omitempty"`
	D []*TestVersion2 `protobuf:"bytes,5,rep,name=d,proto3" json:"d,omitempty"`
	
	
	
	Sum isTestVersion2_Sum `protobuf_oneof:"sum"`
	G   *types.Any         `protobuf:"bytes,8,opt,name=g,proto3" json:"g,omitempty"`
	H   []*TestVersion1    `protobuf:"bytes,9,rep,name=h,proto3" json:"h,omitempty"`
	
	
	*Customer1 `protobuf:"bytes,12,opt,name=k,proto3,embedded=k" json:"k,omitempty"`
	NewField   uint64 `protobuf:"varint,25,opt,name=new_field,json=newField,proto3" json:"new_field,omitempty"`
}

func (m *TestVersion2) Reset()         { *m = TestVersion2{} }
func (m *TestVersion2) String() string { return proto.CompactTextString(m) }
func (*TestVersion2) ProtoMessage()    {}
func (*TestVersion2) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{12}
}
func (m *TestVersion2) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestVersion2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestVersion2.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestVersion2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestVersion2.Merge(m, src)
}
func (m *TestVersion2) XXX_Size() int {
	return m.Size()
}
func (m *TestVersion2) XXX_DiscardUnknown() {
	xxx_messageInfo_TestVersion2.DiscardUnknown(m)
}

var xxx_messageInfo_TestVersion2 proto.InternalMessageInfo

type isTestVersion2_Sum interface {
	isTestVersion2_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type TestVersion2_E struct {
	E int32 `protobuf:"varint,6,opt,name=e,proto3,oneof" json:"e,omitempty"`
}
type TestVersion2_F struct {
	F *TestVersion2 `protobuf:"bytes,7,opt,name=f,proto3,oneof" json:"f,omitempty"`
}

func (*TestVersion2_E) isTestVersion2_Sum() {}
func (*TestVersion2_F) isTestVersion2_Sum() {}

func (m *TestVersion2) GetSum() isTestVersion2_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *TestVersion2) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *TestVersion2) GetA() *TestVersion2 {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *TestVersion2) GetB() *TestVersion2 {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *TestVersion2) GetC() []*TestVersion2 {
	if m != nil {
		return m.C
	}
	return nil
}

func (m *TestVersion2) GetD() []*TestVersion2 {
	if m != nil {
		return m.D
	}
	return nil
}

func (m *TestVersion2) GetE() int32 {
	if x, ok := m.GetSum().(*TestVersion2_E); ok {
		return x.E
	}
	return 0
}

func (m *TestVersion2) GetF() *TestVersion2 {
	if x, ok := m.GetSum().(*TestVersion2_F); ok {
		return x.F
	}
	return nil
}

func (m *TestVersion2) GetG() *types.Any {
	if m != nil {
		return m.G
	}
	return nil
}

func (m *TestVersion2) GetH() []*TestVersion1 {
	if m != nil {
		return m.H
	}
	return nil
}

func (m *TestVersion2) GetNewField() uint64 {
	if m != nil {
		return m.NewField
	}
	return 0
}


func (*TestVersion2) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestVersion2_E)(nil),
		(*TestVersion2_F)(nil),
	}
}

type TestVersion3 struct {
	X int64           `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	A *TestVersion3   `protobuf:"bytes,2,opt,name=a,proto3" json:"a,omitempty"`
	B *TestVersion3   `protobuf:"bytes,3,opt,name=b,proto3" json:"b,omitempty"`
	C []*TestVersion3 `protobuf:"bytes,4,rep,name=c,proto3" json:"c,omitempty"`
	D []*TestVersion3 `protobuf:"bytes,5,rep,name=d,proto3" json:"d,omitempty"`
	
	
	
	Sum isTestVersion3_Sum `protobuf_oneof:"sum"`
	G   *types.Any         `protobuf:"bytes,8,opt,name=g,proto3" json:"g,omitempty"`
	H   []*TestVersion1    `protobuf:"bytes,9,rep,name=h,proto3" json:"h,omitempty"`
	
	
	*Customer1       `protobuf:"bytes,12,opt,name=k,proto3,embedded=k" json:"k,omitempty"`
	NonCriticalField string `protobuf:"bytes,1031,opt,name=non_critical_field,json=nonCriticalField,proto3" json:"non_critical_field,omitempty"`
}

func (m *TestVersion3) Reset()         { *m = TestVersion3{} }
func (m *TestVersion3) String() string { return proto.CompactTextString(m) }
func (*TestVersion3) ProtoMessage()    {}
func (*TestVersion3) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{13}
}
func (m *TestVersion3) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestVersion3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestVersion3.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestVersion3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestVersion3.Merge(m, src)
}
func (m *TestVersion3) XXX_Size() int {
	return m.Size()
}
func (m *TestVersion3) XXX_DiscardUnknown() {
	xxx_messageInfo_TestVersion3.DiscardUnknown(m)
}

var xxx_messageInfo_TestVersion3 proto.InternalMessageInfo

type isTestVersion3_Sum interface {
	isTestVersion3_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type TestVersion3_E struct {
	E int32 `protobuf:"varint,6,opt,name=e,proto3,oneof" json:"e,omitempty"`
}
type TestVersion3_F struct {
	F *TestVersion3 `protobuf:"bytes,7,opt,name=f,proto3,oneof" json:"f,omitempty"`
}

func (*TestVersion3_E) isTestVersion3_Sum() {}
func (*TestVersion3_F) isTestVersion3_Sum() {}

func (m *TestVersion3) GetSum() isTestVersion3_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *TestVersion3) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *TestVersion3) GetA() *TestVersion3 {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *TestVersion3) GetB() *TestVersion3 {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *TestVersion3) GetC() []*TestVersion3 {
	if m != nil {
		return m.C
	}
	return nil
}

func (m *TestVersion3) GetD() []*TestVersion3 {
	if m != nil {
		return m.D
	}
	return nil
}

func (m *TestVersion3) GetE() int32 {
	if x, ok := m.GetSum().(*TestVersion3_E); ok {
		return x.E
	}
	return 0
}

func (m *TestVersion3) GetF() *TestVersion3 {
	if x, ok := m.GetSum().(*TestVersion3_F); ok {
		return x.F
	}
	return nil
}

func (m *TestVersion3) GetG() *types.Any {
	if m != nil {
		return m.G
	}
	return nil
}

func (m *TestVersion3) GetH() []*TestVersion1 {
	if m != nil {
		return m.H
	}
	return nil
}

func (m *TestVersion3) GetNonCriticalField() string {
	if m != nil {
		return m.NonCriticalField
	}
	return ""
}


func (*TestVersion3) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestVersion3_E)(nil),
		(*TestVersion3_F)(nil),
	}
}

type TestVersion3LoneOneOfValue struct {
	X int64           `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	A *TestVersion3   `protobuf:"bytes,2,opt,name=a,proto3" json:"a,omitempty"`
	B *TestVersion3   `protobuf:"bytes,3,opt,name=b,proto3" json:"b,omitempty"`
	C []*TestVersion3 `protobuf:"bytes,4,rep,name=c,proto3" json:"c,omitempty"`
	D []*TestVersion3 `protobuf:"bytes,5,rep,name=d,proto3" json:"d,omitempty"`
	
	
	Sum isTestVersion3LoneOneOfValue_Sum `protobuf_oneof:"sum"`
	G   *types.Any                       `protobuf:"bytes,8,opt,name=g,proto3" json:"g,omitempty"`
	H   []*TestVersion1                  `protobuf:"bytes,9,rep,name=h,proto3" json:"h,omitempty"`
	
	
	*Customer1       `protobuf:"bytes,12,opt,name=k,proto3,embedded=k" json:"k,omitempty"`
	NonCriticalField string `protobuf:"bytes,1031,opt,name=non_critical_field,json=nonCriticalField,proto3" json:"non_critical_field,omitempty"`
}

func (m *TestVersion3LoneOneOfValue) Reset()         { *m = TestVersion3LoneOneOfValue{} }
func (m *TestVersion3LoneOneOfValue) String() string { return proto.CompactTextString(m) }
func (*TestVersion3LoneOneOfValue) ProtoMessage()    {}
func (*TestVersion3LoneOneOfValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{14}
}
func (m *TestVersion3LoneOneOfValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestVersion3LoneOneOfValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestVersion3LoneOneOfValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestVersion3LoneOneOfValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestVersion3LoneOneOfValue.Merge(m, src)
}
func (m *TestVersion3LoneOneOfValue) XXX_Size() int {
	return m.Size()
}
func (m *TestVersion3LoneOneOfValue) XXX_DiscardUnknown() {
	xxx_messageInfo_TestVersion3LoneOneOfValue.DiscardUnknown(m)
}

var xxx_messageInfo_TestVersion3LoneOneOfValue proto.InternalMessageInfo

type isTestVersion3LoneOneOfValue_Sum interface {
	isTestVersion3LoneOneOfValue_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type TestVersion3LoneOneOfValue_E struct {
	E int32 `protobuf:"varint,6,opt,name=e,proto3,oneof" json:"e,omitempty"`
}

func (*TestVersion3LoneOneOfValue_E) isTestVersion3LoneOneOfValue_Sum() {}

func (m *TestVersion3LoneOneOfValue) GetSum() isTestVersion3LoneOneOfValue_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *TestVersion3LoneOneOfValue) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *TestVersion3LoneOneOfValue) GetA() *TestVersion3 {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *TestVersion3LoneOneOfValue) GetB() *TestVersion3 {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *TestVersion3LoneOneOfValue) GetC() []*TestVersion3 {
	if m != nil {
		return m.C
	}
	return nil
}

func (m *TestVersion3LoneOneOfValue) GetD() []*TestVersion3 {
	if m != nil {
		return m.D
	}
	return nil
}

func (m *TestVersion3LoneOneOfValue) GetE() int32 {
	if x, ok := m.GetSum().(*TestVersion3LoneOneOfValue_E); ok {
		return x.E
	}
	return 0
}

func (m *TestVersion3LoneOneOfValue) GetG() *types.Any {
	if m != nil {
		return m.G
	}
	return nil
}

func (m *TestVersion3LoneOneOfValue) GetH() []*TestVersion1 {
	if m != nil {
		return m.H
	}
	return nil
}

func (m *TestVersion3LoneOneOfValue) GetNonCriticalField() string {
	if m != nil {
		return m.NonCriticalField
	}
	return ""
}


func (*TestVersion3LoneOneOfValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestVersion3LoneOneOfValue_E)(nil),
	}
}

type TestVersion3LoneNesting struct {
	X int64           `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	A *TestVersion3   `protobuf:"bytes,2,opt,name=a,proto3" json:"a,omitempty"`
	B *TestVersion3   `protobuf:"bytes,3,opt,name=b,proto3" json:"b,omitempty"`
	C []*TestVersion3 `protobuf:"bytes,4,rep,name=c,proto3" json:"c,omitempty"`
	D []*TestVersion3 `protobuf:"bytes,5,rep,name=d,proto3" json:"d,omitempty"`
	
	
	Sum isTestVersion3LoneNesting_Sum `protobuf_oneof:"sum"`
	G   *types.Any                    `protobuf:"bytes,8,opt,name=g,proto3" json:"g,omitempty"`
	H   []*TestVersion1               `protobuf:"bytes,9,rep,name=h,proto3" json:"h,omitempty"`
	
	
	*Customer1       `protobuf:"bytes,12,opt,name=k,proto3,embedded=k" json:"k,omitempty"`
	NonCriticalField string                          `protobuf:"bytes,1031,opt,name=non_critical_field,json=nonCriticalField,proto3" json:"non_critical_field,omitempty"`
	Inner1           *TestVersion3LoneNesting_Inner1 `protobuf:"bytes,14,opt,name=inner1,proto3" json:"inner1,omitempty"`
	Inner2           *TestVersion3LoneNesting_Inner2 `protobuf:"bytes,15,opt,name=inner2,proto3" json:"inner2,omitempty"`
}

func (m *TestVersion3LoneNesting) Reset()         { *m = TestVersion3LoneNesting{} }
func (m *TestVersion3LoneNesting) String() string { return proto.CompactTextString(m) }
func (*TestVersion3LoneNesting) ProtoMessage()    {}
func (*TestVersion3LoneNesting) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{15}
}
func (m *TestVersion3LoneNesting) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestVersion3LoneNesting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestVersion3LoneNesting.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestVersion3LoneNesting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestVersion3LoneNesting.Merge(m, src)
}
func (m *TestVersion3LoneNesting) XXX_Size() int {
	return m.Size()
}
func (m *TestVersion3LoneNesting) XXX_DiscardUnknown() {
	xxx_messageInfo_TestVersion3LoneNesting.DiscardUnknown(m)
}

var xxx_messageInfo_TestVersion3LoneNesting proto.InternalMessageInfo

type isTestVersion3LoneNesting_Sum interface {
	isTestVersion3LoneNesting_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type TestVersion3LoneNesting_F struct {
	F *TestVersion3LoneNesting `protobuf:"bytes,7,opt,name=f,proto3,oneof" json:"f,omitempty"`
}

func (*TestVersion3LoneNesting_F) isTestVersion3LoneNesting_Sum() {}

func (m *TestVersion3LoneNesting) GetSum() isTestVersion3LoneNesting_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *TestVersion3LoneNesting) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *TestVersion3LoneNesting) GetA() *TestVersion3 {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *TestVersion3LoneNesting) GetB() *TestVersion3 {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *TestVersion3LoneNesting) GetC() []*TestVersion3 {
	if m != nil {
		return m.C
	}
	return nil
}

func (m *TestVersion3LoneNesting) GetD() []*TestVersion3 {
	if m != nil {
		return m.D
	}
	return nil
}

func (m *TestVersion3LoneNesting) GetF() *TestVersion3LoneNesting {
	if x, ok := m.GetSum().(*TestVersion3LoneNesting_F); ok {
		return x.F
	}
	return nil
}

func (m *TestVersion3LoneNesting) GetG() *types.Any {
	if m != nil {
		return m.G
	}
	return nil
}

func (m *TestVersion3LoneNesting) GetH() []*TestVersion1 {
	if m != nil {
		return m.H
	}
	return nil
}

func (m *TestVersion3LoneNesting) GetNonCriticalField() string {
	if m != nil {
		return m.NonCriticalField
	}
	return ""
}

func (m *TestVersion3LoneNesting) GetInner1() *TestVersion3LoneNesting_Inner1 {
	if m != nil {
		return m.Inner1
	}
	return nil
}

func (m *TestVersion3LoneNesting) GetInner2() *TestVersion3LoneNesting_Inner2 {
	if m != nil {
		return m.Inner2
	}
	return nil
}


func (*TestVersion3LoneNesting) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestVersion3LoneNesting_F)(nil),
	}
}

type TestVersion3LoneNesting_Inner1 struct {
	Id    int64                                      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string                                     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Inner *TestVersion3LoneNesting_Inner1_InnerInner `protobuf:"bytes,3,opt,name=inner,proto3" json:"inner,omitempty"`
}

func (m *TestVersion3LoneNesting_Inner1) Reset()         { *m = TestVersion3LoneNesting_Inner1{} }
func (m *TestVersion3LoneNesting_Inner1) String() string { return proto.CompactTextString(m) }
func (*TestVersion3LoneNesting_Inner1) ProtoMessage()    {}
func (*TestVersion3LoneNesting_Inner1) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{15, 0}
}
func (m *TestVersion3LoneNesting_Inner1) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestVersion3LoneNesting_Inner1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestVersion3LoneNesting_Inner1.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestVersion3LoneNesting_Inner1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestVersion3LoneNesting_Inner1.Merge(m, src)
}
func (m *TestVersion3LoneNesting_Inner1) XXX_Size() int {
	return m.Size()
}
func (m *TestVersion3LoneNesting_Inner1) XXX_DiscardUnknown() {
	xxx_messageInfo_TestVersion3LoneNesting_Inner1.DiscardUnknown(m)
}

var xxx_messageInfo_TestVersion3LoneNesting_Inner1 proto.InternalMessageInfo

func (m *TestVersion3LoneNesting_Inner1) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TestVersion3LoneNesting_Inner1) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestVersion3LoneNesting_Inner1) GetInner() *TestVersion3LoneNesting_Inner1_InnerInner {
	if m != nil {
		return m.Inner
	}
	return nil
}

type TestVersion3LoneNesting_Inner1_InnerInner struct {
	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	City string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
}

func (m *TestVersion3LoneNesting_Inner1_InnerInner) Reset() {
	*m = TestVersion3LoneNesting_Inner1_InnerInner{}
}
func (m *TestVersion3LoneNesting_Inner1_InnerInner) String() string {
	return proto.CompactTextString(m)
}
func (*TestVersion3LoneNesting_Inner1_InnerInner) ProtoMessage() {}
func (*TestVersion3LoneNesting_Inner1_InnerInner) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{15, 0, 0}
}
func (m *TestVersion3LoneNesting_Inner1_InnerInner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestVersion3LoneNesting_Inner1_InnerInner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestVersion3LoneNesting_Inner1_InnerInner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestVersion3LoneNesting_Inner1_InnerInner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestVersion3LoneNesting_Inner1_InnerInner.Merge(m, src)
}
func (m *TestVersion3LoneNesting_Inner1_InnerInner) XXX_Size() int {
	return m.Size()
}
func (m *TestVersion3LoneNesting_Inner1_InnerInner) XXX_DiscardUnknown() {
	xxx_messageInfo_TestVersion3LoneNesting_Inner1_InnerInner.DiscardUnknown(m)
}

var xxx_messageInfo_TestVersion3LoneNesting_Inner1_InnerInner proto.InternalMessageInfo

func (m *TestVersion3LoneNesting_Inner1_InnerInner) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TestVersion3LoneNesting_Inner1_InnerInner) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

type TestVersion3LoneNesting_Inner2 struct {
	Id      string                                     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Country string                                     `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Inner   *TestVersion3LoneNesting_Inner2_InnerInner `protobuf:"bytes,3,opt,name=inner,proto3" json:"inner,omitempty"`
}

func (m *TestVersion3LoneNesting_Inner2) Reset()         { *m = TestVersion3LoneNesting_Inner2{} }
func (m *TestVersion3LoneNesting_Inner2) String() string { return proto.CompactTextString(m) }
func (*TestVersion3LoneNesting_Inner2) ProtoMessage()    {}
func (*TestVersion3LoneNesting_Inner2) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{15, 1}
}
func (m *TestVersion3LoneNesting_Inner2) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestVersion3LoneNesting_Inner2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestVersion3LoneNesting_Inner2.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestVersion3LoneNesting_Inner2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestVersion3LoneNesting_Inner2.Merge(m, src)
}
func (m *TestVersion3LoneNesting_Inner2) XXX_Size() int {
	return m.Size()
}
func (m *TestVersion3LoneNesting_Inner2) XXX_DiscardUnknown() {
	xxx_messageInfo_TestVersion3LoneNesting_Inner2.DiscardUnknown(m)
}

var xxx_messageInfo_TestVersion3LoneNesting_Inner2 proto.InternalMessageInfo

func (m *TestVersion3LoneNesting_Inner2) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TestVersion3LoneNesting_Inner2) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *TestVersion3LoneNesting_Inner2) GetInner() *TestVersion3LoneNesting_Inner2_InnerInner {
	if m != nil {
		return m.Inner
	}
	return nil
}

type TestVersion3LoneNesting_Inner2_InnerInner struct {
	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	City string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
}

func (m *TestVersion3LoneNesting_Inner2_InnerInner) Reset() {
	*m = TestVersion3LoneNesting_Inner2_InnerInner{}
}
func (m *TestVersion3LoneNesting_Inner2_InnerInner) String() string {
	return proto.CompactTextString(m)
}
func (*TestVersion3LoneNesting_Inner2_InnerInner) ProtoMessage() {}
func (*TestVersion3LoneNesting_Inner2_InnerInner) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{15, 1, 0}
}
func (m *TestVersion3LoneNesting_Inner2_InnerInner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestVersion3LoneNesting_Inner2_InnerInner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestVersion3LoneNesting_Inner2_InnerInner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestVersion3LoneNesting_Inner2_InnerInner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestVersion3LoneNesting_Inner2_InnerInner.Merge(m, src)
}
func (m *TestVersion3LoneNesting_Inner2_InnerInner) XXX_Size() int {
	return m.Size()
}
func (m *TestVersion3LoneNesting_Inner2_InnerInner) XXX_DiscardUnknown() {
	xxx_messageInfo_TestVersion3LoneNesting_Inner2_InnerInner.DiscardUnknown(m)
}

var xxx_messageInfo_TestVersion3LoneNesting_Inner2_InnerInner proto.InternalMessageInfo

func (m *TestVersion3LoneNesting_Inner2_InnerInner) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TestVersion3LoneNesting_Inner2_InnerInner) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

type TestVersion4LoneNesting struct {
	X int64           `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	A *TestVersion3   `protobuf:"bytes,2,opt,name=a,proto3" json:"a,omitempty"`
	B *TestVersion3   `protobuf:"bytes,3,opt,name=b,proto3" json:"b,omitempty"`
	C []*TestVersion3 `protobuf:"bytes,4,rep,name=c,proto3" json:"c,omitempty"`
	D []*TestVersion3 `protobuf:"bytes,5,rep,name=d,proto3" json:"d,omitempty"`
	
	
	Sum isTestVersion4LoneNesting_Sum `protobuf_oneof:"sum"`
	G   *types.Any                    `protobuf:"bytes,8,opt,name=g,proto3" json:"g,omitempty"`
	H   []*TestVersion1               `protobuf:"bytes,9,rep,name=h,proto3" json:"h,omitempty"`
	
	
	*Customer1       `protobuf:"bytes,12,opt,name=k,proto3,embedded=k" json:"k,omitempty"`
	NonCriticalField string                          `protobuf:"bytes,1031,opt,name=non_critical_field,json=nonCriticalField,proto3" json:"non_critical_field,omitempty"`
	Inner1           *TestVersion4LoneNesting_Inner1 `protobuf:"bytes,14,opt,name=inner1,proto3" json:"inner1,omitempty"`
	Inner2           *TestVersion4LoneNesting_Inner2 `protobuf:"bytes,15,opt,name=inner2,proto3" json:"inner2,omitempty"`
}

func (m *TestVersion4LoneNesting) Reset()         { *m = TestVersion4LoneNesting{} }
func (m *TestVersion4LoneNesting) String() string { return proto.CompactTextString(m) }
func (*TestVersion4LoneNesting) ProtoMessage()    {}
func (*TestVersion4LoneNesting) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{16}
}
func (m *TestVersion4LoneNesting) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestVersion4LoneNesting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestVersion4LoneNesting.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestVersion4LoneNesting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestVersion4LoneNesting.Merge(m, src)
}
func (m *TestVersion4LoneNesting) XXX_Size() int {
	return m.Size()
}
func (m *TestVersion4LoneNesting) XXX_DiscardUnknown() {
	xxx_messageInfo_TestVersion4LoneNesting.DiscardUnknown(m)
}

var xxx_messageInfo_TestVersion4LoneNesting proto.InternalMessageInfo

type isTestVersion4LoneNesting_Sum interface {
	isTestVersion4LoneNesting_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type TestVersion4LoneNesting_F struct {
	F *TestVersion3LoneNesting `protobuf:"bytes,7,opt,name=f,proto3,oneof" json:"f,omitempty"`
}

func (*TestVersion4LoneNesting_F) isTestVersion4LoneNesting_Sum() {}

func (m *TestVersion4LoneNesting) GetSum() isTestVersion4LoneNesting_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *TestVersion4LoneNesting) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *TestVersion4LoneNesting) GetA() *TestVersion3 {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *TestVersion4LoneNesting) GetB() *TestVersion3 {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *TestVersion4LoneNesting) GetC() []*TestVersion3 {
	if m != nil {
		return m.C
	}
	return nil
}

func (m *TestVersion4LoneNesting) GetD() []*TestVersion3 {
	if m != nil {
		return m.D
	}
	return nil
}

func (m *TestVersion4LoneNesting) GetF() *TestVersion3LoneNesting {
	if x, ok := m.GetSum().(*TestVersion4LoneNesting_F); ok {
		return x.F
	}
	return nil
}

func (m *TestVersion4LoneNesting) GetG() *types.Any {
	if m != nil {
		return m.G
	}
	return nil
}

func (m *TestVersion4LoneNesting) GetH() []*TestVersion1 {
	if m != nil {
		return m.H
	}
	return nil
}

func (m *TestVersion4LoneNesting) GetNonCriticalField() string {
	if m != nil {
		return m.NonCriticalField
	}
	return ""
}

func (m *TestVersion4LoneNesting) GetInner1() *TestVersion4LoneNesting_Inner1 {
	if m != nil {
		return m.Inner1
	}
	return nil
}

func (m *TestVersion4LoneNesting) GetInner2() *TestVersion4LoneNesting_Inner2 {
	if m != nil {
		return m.Inner2
	}
	return nil
}


func (*TestVersion4LoneNesting) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestVersion4LoneNesting_F)(nil),
	}
}

type TestVersion4LoneNesting_Inner1 struct {
	Id    int64                                      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string                                     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Inner *TestVersion4LoneNesting_Inner1_InnerInner `protobuf:"bytes,3,opt,name=inner,proto3" json:"inner,omitempty"`
}

func (m *TestVersion4LoneNesting_Inner1) Reset()         { *m = TestVersion4LoneNesting_Inner1{} }
func (m *TestVersion4LoneNesting_Inner1) String() string { return proto.CompactTextString(m) }
func (*TestVersion4LoneNesting_Inner1) ProtoMessage()    {}
func (*TestVersion4LoneNesting_Inner1) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{16, 0}
}
func (m *TestVersion4LoneNesting_Inner1) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestVersion4LoneNesting_Inner1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestVersion4LoneNesting_Inner1.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestVersion4LoneNesting_Inner1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestVersion4LoneNesting_Inner1.Merge(m, src)
}
func (m *TestVersion4LoneNesting_Inner1) XXX_Size() int {
	return m.Size()
}
func (m *TestVersion4LoneNesting_Inner1) XXX_DiscardUnknown() {
	xxx_messageInfo_TestVersion4LoneNesting_Inner1.DiscardUnknown(m)
}

var xxx_messageInfo_TestVersion4LoneNesting_Inner1 proto.InternalMessageInfo

func (m *TestVersion4LoneNesting_Inner1) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TestVersion4LoneNesting_Inner1) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestVersion4LoneNesting_Inner1) GetInner() *TestVersion4LoneNesting_Inner1_InnerInner {
	if m != nil {
		return m.Inner
	}
	return nil
}

type TestVersion4LoneNesting_Inner1_InnerInner struct {
	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	City string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
}

func (m *TestVersion4LoneNesting_Inner1_InnerInner) Reset() {
	*m = TestVersion4LoneNesting_Inner1_InnerInner{}
}
func (m *TestVersion4LoneNesting_Inner1_InnerInner) String() string {
	return proto.CompactTextString(m)
}
func (*TestVersion4LoneNesting_Inner1_InnerInner) ProtoMessage() {}
func (*TestVersion4LoneNesting_Inner1_InnerInner) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{16, 0, 0}
}
func (m *TestVersion4LoneNesting_Inner1_InnerInner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestVersion4LoneNesting_Inner1_InnerInner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestVersion4LoneNesting_Inner1_InnerInner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestVersion4LoneNesting_Inner1_InnerInner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestVersion4LoneNesting_Inner1_InnerInner.Merge(m, src)
}
func (m *TestVersion4LoneNesting_Inner1_InnerInner) XXX_Size() int {
	return m.Size()
}
func (m *TestVersion4LoneNesting_Inner1_InnerInner) XXX_DiscardUnknown() {
	xxx_messageInfo_TestVersion4LoneNesting_Inner1_InnerInner.DiscardUnknown(m)
}

var xxx_messageInfo_TestVersion4LoneNesting_Inner1_InnerInner proto.InternalMessageInfo

func (m *TestVersion4LoneNesting_Inner1_InnerInner) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TestVersion4LoneNesting_Inner1_InnerInner) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

type TestVersion4LoneNesting_Inner2 struct {
	Id      string                                     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Country string                                     `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Inner   *TestVersion4LoneNesting_Inner2_InnerInner `protobuf:"bytes,3,opt,name=inner,proto3" json:"inner,omitempty"`
}

func (m *TestVersion4LoneNesting_Inner2) Reset()         { *m = TestVersion4LoneNesting_Inner2{} }
func (m *TestVersion4LoneNesting_Inner2) String() string { return proto.CompactTextString(m) }
func (*TestVersion4LoneNesting_Inner2) ProtoMessage()    {}
func (*TestVersion4LoneNesting_Inner2) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{16, 1}
}
func (m *TestVersion4LoneNesting_Inner2) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestVersion4LoneNesting_Inner2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestVersion4LoneNesting_Inner2.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestVersion4LoneNesting_Inner2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestVersion4LoneNesting_Inner2.Merge(m, src)
}
func (m *TestVersion4LoneNesting_Inner2) XXX_Size() int {
	return m.Size()
}
func (m *TestVersion4LoneNesting_Inner2) XXX_DiscardUnknown() {
	xxx_messageInfo_TestVersion4LoneNesting_Inner2.DiscardUnknown(m)
}

var xxx_messageInfo_TestVersion4LoneNesting_Inner2 proto.InternalMessageInfo

func (m *TestVersion4LoneNesting_Inner2) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TestVersion4LoneNesting_Inner2) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *TestVersion4LoneNesting_Inner2) GetInner() *TestVersion4LoneNesting_Inner2_InnerInner {
	if m != nil {
		return m.Inner
	}
	return nil
}

type TestVersion4LoneNesting_Inner2_InnerInner struct {
	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *TestVersion4LoneNesting_Inner2_InnerInner) Reset() {
	*m = TestVersion4LoneNesting_Inner2_InnerInner{}
}
func (m *TestVersion4LoneNesting_Inner2_InnerInner) String() string {
	return proto.CompactTextString(m)
}
func (*TestVersion4LoneNesting_Inner2_InnerInner) ProtoMessage() {}
func (*TestVersion4LoneNesting_Inner2_InnerInner) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{16, 1, 0}
}
func (m *TestVersion4LoneNesting_Inner2_InnerInner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestVersion4LoneNesting_Inner2_InnerInner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestVersion4LoneNesting_Inner2_InnerInner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestVersion4LoneNesting_Inner2_InnerInner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestVersion4LoneNesting_Inner2_InnerInner.Merge(m, src)
}
func (m *TestVersion4LoneNesting_Inner2_InnerInner) XXX_Size() int {
	return m.Size()
}
func (m *TestVersion4LoneNesting_Inner2_InnerInner) XXX_DiscardUnknown() {
	xxx_messageInfo_TestVersion4LoneNesting_Inner2_InnerInner.DiscardUnknown(m)
}

var xxx_messageInfo_TestVersion4LoneNesting_Inner2_InnerInner proto.InternalMessageInfo

func (m *TestVersion4LoneNesting_Inner2_InnerInner) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TestVersion4LoneNesting_Inner2_InnerInner) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type TestVersionFD1 struct {
	X int64         `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	A *TestVersion1 `protobuf:"bytes,2,opt,name=a,proto3" json:"a,omitempty"`
	
	
	
	Sum isTestVersionFD1_Sum `protobuf_oneof:"sum"`
	G   *types.Any           `protobuf:"bytes,8,opt,name=g,proto3" json:"g,omitempty"`
	H   []*TestVersion1      `protobuf:"bytes,9,rep,name=h,proto3" json:"h,omitempty"`
}

func (m *TestVersionFD1) Reset()         { *m = TestVersionFD1{} }
func (m *TestVersionFD1) String() string { return proto.CompactTextString(m) }
func (*TestVersionFD1) ProtoMessage()    {}
func (*TestVersionFD1) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{17}
}
func (m *TestVersionFD1) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestVersionFD1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestVersionFD1.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestVersionFD1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestVersionFD1.Merge(m, src)
}
func (m *TestVersionFD1) XXX_Size() int {
	return m.Size()
}
func (m *TestVersionFD1) XXX_DiscardUnknown() {
	xxx_messageInfo_TestVersionFD1.DiscardUnknown(m)
}

var xxx_messageInfo_TestVersionFD1 proto.InternalMessageInfo

type isTestVersionFD1_Sum interface {
	isTestVersionFD1_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type TestVersionFD1_E struct {
	E int32 `protobuf:"varint,6,opt,name=e,proto3,oneof" json:"e,omitempty"`
}
type TestVersionFD1_F struct {
	F *TestVersion1 `protobuf:"bytes,7,opt,name=f,proto3,oneof" json:"f,omitempty"`
}

func (*TestVersionFD1_E) isTestVersionFD1_Sum() {}
func (*TestVersionFD1_F) isTestVersionFD1_Sum() {}

func (m *TestVersionFD1) GetSum() isTestVersionFD1_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *TestVersionFD1) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *TestVersionFD1) GetA() *TestVersion1 {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *TestVersionFD1) GetE() int32 {
	if x, ok := m.GetSum().(*TestVersionFD1_E); ok {
		return x.E
	}
	return 0
}

func (m *TestVersionFD1) GetF() *TestVersion1 {
	if x, ok := m.GetSum().(*TestVersionFD1_F); ok {
		return x.F
	}
	return nil
}

func (m *TestVersionFD1) GetG() *types.Any {
	if m != nil {
		return m.G
	}
	return nil
}

func (m *TestVersionFD1) GetH() []*TestVersion1 {
	if m != nil {
		return m.H
	}
	return nil
}


func (*TestVersionFD1) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestVersionFD1_E)(nil),
		(*TestVersionFD1_F)(nil),
	}
}

type TestVersionFD1WithExtraAny struct {
	X int64         `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	A *TestVersion1 `protobuf:"bytes,2,opt,name=a,proto3" json:"a,omitempty"`
	
	
	
	Sum isTestVersionFD1WithExtraAny_Sum `protobuf_oneof:"sum"`
	G   *AnyWithExtra                    `protobuf:"bytes,8,opt,name=g,proto3" json:"g,omitempty"`
	H   []*TestVersion1                  `protobuf:"bytes,9,rep,name=h,proto3" json:"h,omitempty"`
}

func (m *TestVersionFD1WithExtraAny) Reset()         { *m = TestVersionFD1WithExtraAny{} }
func (m *TestVersionFD1WithExtraAny) String() string { return proto.CompactTextString(m) }
func (*TestVersionFD1WithExtraAny) ProtoMessage()    {}
func (*TestVersionFD1WithExtraAny) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{18}
}
func (m *TestVersionFD1WithExtraAny) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestVersionFD1WithExtraAny) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestVersionFD1WithExtraAny.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestVersionFD1WithExtraAny) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestVersionFD1WithExtraAny.Merge(m, src)
}
func (m *TestVersionFD1WithExtraAny) XXX_Size() int {
	return m.Size()
}
func (m *TestVersionFD1WithExtraAny) XXX_DiscardUnknown() {
	xxx_messageInfo_TestVersionFD1WithExtraAny.DiscardUnknown(m)
}

var xxx_messageInfo_TestVersionFD1WithExtraAny proto.InternalMessageInfo

type isTestVersionFD1WithExtraAny_Sum interface {
	isTestVersionFD1WithExtraAny_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type TestVersionFD1WithExtraAny_E struct {
	E int32 `protobuf:"varint,6,opt,name=e,proto3,oneof" json:"e,omitempty"`
}
type TestVersionFD1WithExtraAny_F struct {
	F *TestVersion1 `protobuf:"bytes,7,opt,name=f,proto3,oneof" json:"f,omitempty"`
}

func (*TestVersionFD1WithExtraAny_E) isTestVersionFD1WithExtraAny_Sum() {}
func (*TestVersionFD1WithExtraAny_F) isTestVersionFD1WithExtraAny_Sum() {}

func (m *TestVersionFD1WithExtraAny) GetSum() isTestVersionFD1WithExtraAny_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *TestVersionFD1WithExtraAny) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *TestVersionFD1WithExtraAny) GetA() *TestVersion1 {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *TestVersionFD1WithExtraAny) GetE() int32 {
	if x, ok := m.GetSum().(*TestVersionFD1WithExtraAny_E); ok {
		return x.E
	}
	return 0
}

func (m *TestVersionFD1WithExtraAny) GetF() *TestVersion1 {
	if x, ok := m.GetSum().(*TestVersionFD1WithExtraAny_F); ok {
		return x.F
	}
	return nil
}

func (m *TestVersionFD1WithExtraAny) GetG() *AnyWithExtra {
	if m != nil {
		return m.G
	}
	return nil
}

func (m *TestVersionFD1WithExtraAny) GetH() []*TestVersion1 {
	if m != nil {
		return m.H
	}
	return nil
}


func (*TestVersionFD1WithExtraAny) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestVersionFD1WithExtraAny_E)(nil),
		(*TestVersionFD1WithExtraAny_F)(nil),
	}
}

type AnyWithExtra struct {
	*types.Any `protobuf:"bytes,1,opt,name=a,proto3,embedded=a" json:"a,omitempty"`
	B          int64 `protobuf:"varint,3,opt,name=b,proto3" json:"b,omitempty"`
	C          int64 `protobuf:"varint,4,opt,name=c,proto3" json:"c,omitempty"`
}

func (m *AnyWithExtra) Reset()         { *m = AnyWithExtra{} }
func (m *AnyWithExtra) String() string { return proto.CompactTextString(m) }
func (*AnyWithExtra) ProtoMessage()    {}
func (*AnyWithExtra) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{19}
}
func (m *AnyWithExtra) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AnyWithExtra) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AnyWithExtra.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AnyWithExtra) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnyWithExtra.Merge(m, src)
}
func (m *AnyWithExtra) XXX_Size() int {
	return m.Size()
}
func (m *AnyWithExtra) XXX_DiscardUnknown() {
	xxx_messageInfo_AnyWithExtra.DiscardUnknown(m)
}

var xxx_messageInfo_AnyWithExtra proto.InternalMessageInfo

func (m *AnyWithExtra) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

func (m *AnyWithExtra) GetC() int64 {
	if m != nil {
		return m.C
	}
	return 0
}

type TestUpdatedTxRaw struct {
	BodyBytes     []byte   `protobuf:"bytes,1,opt,name=body_bytes,json=bodyBytes,proto3" json:"body_bytes,omitempty"`
	AuthInfoBytes []byte   `protobuf:"bytes,2,opt,name=auth_info_bytes,json=authInfoBytes,proto3" json:"auth_info_bytes,omitempty"`
	Signatures    [][]byte `protobuf:"bytes,3,rep,name=signatures,proto3" json:"signatures,omitempty"`
	NewField_5    []byte   `protobuf:"bytes,5,opt,name=new_field_5,json=newField5,proto3" json:"new_field_5,omitempty"`
	NewField_1024 []byte   `protobuf:"bytes,1024,opt,name=new_field_1024,json=newField1024,proto3" json:"new_field_1024,omitempty"`
}

func (m *TestUpdatedTxRaw) Reset()         { *m = TestUpdatedTxRaw{} }
func (m *TestUpdatedTxRaw) String() string { return proto.CompactTextString(m) }
func (*TestUpdatedTxRaw) ProtoMessage()    {}
func (*TestUpdatedTxRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{20}
}
func (m *TestUpdatedTxRaw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestUpdatedTxRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestUpdatedTxRaw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestUpdatedTxRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestUpdatedTxRaw.Merge(m, src)
}
func (m *TestUpdatedTxRaw) XXX_Size() int {
	return m.Size()
}
func (m *TestUpdatedTxRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_TestUpdatedTxRaw.DiscardUnknown(m)
}

var xxx_messageInfo_TestUpdatedTxRaw proto.InternalMessageInfo

func (m *TestUpdatedTxRaw) GetBodyBytes() []byte {
	if m != nil {
		return m.BodyBytes
	}
	return nil
}

func (m *TestUpdatedTxRaw) GetAuthInfoBytes() []byte {
	if m != nil {
		return m.AuthInfoBytes
	}
	return nil
}

func (m *TestUpdatedTxRaw) GetSignatures() [][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *TestUpdatedTxRaw) GetNewField_5() []byte {
	if m != nil {
		return m.NewField_5
	}
	return nil
}

func (m *TestUpdatedTxRaw) GetNewField_1024() []byte {
	if m != nil {
		return m.NewField_1024
	}
	return nil
}

type TestUpdatedTxBody struct {
	Messages                     []*types.Any `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	Memo                         string       `protobuf:"bytes,2,opt,name=memo,proto3" json:"memo,omitempty"`
	TimeoutHeight                int64        `protobuf:"varint,3,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height,omitempty"`
	SomeNewField                 uint64       `protobuf:"varint,4,opt,name=some_new_field,json=someNewField,proto3" json:"some_new_field,omitempty"`
	SomeNewFieldNonCriticalField string       `protobuf:"bytes,1050,opt,name=some_new_field_non_critical_field,json=someNewFieldNonCriticalField,proto3" json:"some_new_field_non_critical_field,omitempty"`
	ExtensionOptions             []*types.Any `protobuf:"bytes,1023,rep,name=extension_options,json=extensionOptions,proto3" json:"extension_options,omitempty"`
	NonCriticalExtensionOptions  []*types.Any `protobuf:"bytes,2047,rep,name=non_critical_extension_options,json=nonCriticalExtensionOptions,proto3" json:"non_critical_extension_options,omitempty"`
}

func (m *TestUpdatedTxBody) Reset()         { *m = TestUpdatedTxBody{} }
func (m *TestUpdatedTxBody) String() string { return proto.CompactTextString(m) }
func (*TestUpdatedTxBody) ProtoMessage()    {}
func (*TestUpdatedTxBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{21}
}
func (m *TestUpdatedTxBody) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestUpdatedTxBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestUpdatedTxBody.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestUpdatedTxBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestUpdatedTxBody.Merge(m, src)
}
func (m *TestUpdatedTxBody) XXX_Size() int {
	return m.Size()
}
func (m *TestUpdatedTxBody) XXX_DiscardUnknown() {
	xxx_messageInfo_TestUpdatedTxBody.DiscardUnknown(m)
}

var xxx_messageInfo_TestUpdatedTxBody proto.InternalMessageInfo

func (m *TestUpdatedTxBody) GetMessages() []*types.Any {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *TestUpdatedTxBody) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *TestUpdatedTxBody) GetTimeoutHeight() int64 {
	if m != nil {
		return m.TimeoutHeight
	}
	return 0
}

func (m *TestUpdatedTxBody) GetSomeNewField() uint64 {
	if m != nil {
		return m.SomeNewField
	}
	return 0
}

func (m *TestUpdatedTxBody) GetSomeNewFieldNonCriticalField() string {
	if m != nil {
		return m.SomeNewFieldNonCriticalField
	}
	return ""
}

func (m *TestUpdatedTxBody) GetExtensionOptions() []*types.Any {
	if m != nil {
		return m.ExtensionOptions
	}
	return nil
}

func (m *TestUpdatedTxBody) GetNonCriticalExtensionOptions() []*types.Any {
	if m != nil {
		return m.NonCriticalExtensionOptions
	}
	return nil
}

type TestUpdatedAuthInfo struct {
	SignerInfos   []*tx.SignerInfo `protobuf:"bytes,1,rep,name=signer_infos,json=signerInfos,proto3" json:"signer_infos,omitempty"`
	Fee           *tx.Fee          `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee,omitempty"`
	NewField_3    []byte           `protobuf:"bytes,3,opt,name=new_field_3,json=newField3,proto3" json:"new_field_3,omitempty"`
	NewField_1024 []byte           `protobuf:"bytes,1024,opt,name=new_field_1024,json=newField1024,proto3" json:"new_field_1024,omitempty"`
}

func (m *TestUpdatedAuthInfo) Reset()         { *m = TestUpdatedAuthInfo{} }
func (m *TestUpdatedAuthInfo) String() string { return proto.CompactTextString(m) }
func (*TestUpdatedAuthInfo) ProtoMessage()    {}
func (*TestUpdatedAuthInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{22}
}
func (m *TestUpdatedAuthInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestUpdatedAuthInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestUpdatedAuthInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestUpdatedAuthInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestUpdatedAuthInfo.Merge(m, src)
}
func (m *TestUpdatedAuthInfo) XXX_Size() int {
	return m.Size()
}
func (m *TestUpdatedAuthInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TestUpdatedAuthInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TestUpdatedAuthInfo proto.InternalMessageInfo

func (m *TestUpdatedAuthInfo) GetSignerInfos() []*tx.SignerInfo {
	if m != nil {
		return m.SignerInfos
	}
	return nil
}

func (m *TestUpdatedAuthInfo) GetFee() *tx.Fee {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (m *TestUpdatedAuthInfo) GetNewField_3() []byte {
	if m != nil {
		return m.NewField_3
	}
	return nil
}

func (m *TestUpdatedAuthInfo) GetNewField_1024() []byte {
	if m != nil {
		return m.NewField_1024
	}
	return nil
}

type TestRepeatedUints struct {
	Nums []uint64 `protobuf:"varint,1,rep,packed,name=nums,proto3" json:"nums,omitempty"`
}

func (m *TestRepeatedUints) Reset()         { *m = TestRepeatedUints{} }
func (m *TestRepeatedUints) String() string { return proto.CompactTextString(m) }
func (*TestRepeatedUints) ProtoMessage()    {}
func (*TestRepeatedUints) Descriptor() ([]byte, []int) {
	return fileDescriptor_448ea787339d1228, []int{23}
}
func (m *TestRepeatedUints) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestRepeatedUints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestRepeatedUints.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestRepeatedUints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRepeatedUints.Merge(m, src)
}
func (m *TestRepeatedUints) XXX_Size() int {
	return m.Size()
}
func (m *TestRepeatedUints) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRepeatedUints.DiscardUnknown(m)
}

var xxx_messageInfo_TestRepeatedUints proto.InternalMessageInfo

func (m *TestRepeatedUints) GetNums() []uint64 {
	if m != nil {
		return m.Nums
	}
	return nil
}

func init() {
	proto.RegisterEnum("testdata.Customer2_City", Customer2_City_name, Customer2_City_value)
	proto.RegisterType((*Customer1)(nil), "testdata.Customer1")
	proto.RegisterType((*Customer2)(nil), "testdata.Customer2")
	proto.RegisterType((*Nested4A)(nil), "testdata.Nested4A")
	proto.RegisterType((*Nested3A)(nil), "testdata.Nested3A")
	proto.RegisterMapType((map[int64]*Nested4A)(nil), "testdata.Nested3A.IndexEntry")
	proto.RegisterType((*Nested2A)(nil), "testdata.Nested2A")
	proto.RegisterType((*Nested1A)(nil), "testdata.Nested1A")
	proto.RegisterType((*Nested4B)(nil), "testdata.Nested4B")
	proto.RegisterType((*Nested3B)(nil), "testdata.Nested3B")
	proto.RegisterType((*Nested2B)(nil), "testdata.Nested2B")
	proto.RegisterType((*Nested1B)(nil), "testdata.Nested1B")
	proto.RegisterType((*Customer3)(nil), "testdata.Customer3")
	proto.RegisterType((*TestVersion1)(nil), "testdata.TestVersion1")
	proto.RegisterType((*TestVersion2)(nil), "testdata.TestVersion2")
	proto.RegisterType((*TestVersion3)(nil), "testdata.TestVersion3")
	proto.RegisterType((*TestVersion3LoneOneOfValue)(nil), "testdata.TestVersion3LoneOneOfValue")
	proto.RegisterType((*TestVersion3LoneNesting)(nil), "testdata.TestVersion3LoneNesting")
	proto.RegisterType((*TestVersion3LoneNesting_Inner1)(nil), "testdata.TestVersion3LoneNesting.Inner1")
	proto.RegisterType((*TestVersion3LoneNesting_Inner1_InnerInner)(nil), "testdata.TestVersion3LoneNesting.Inner1.InnerInner")
	proto.RegisterType((*TestVersion3LoneNesting_Inner2)(nil), "testdata.TestVersion3LoneNesting.Inner2")
	proto.RegisterType((*TestVersion3LoneNesting_Inner2_InnerInner)(nil), "testdata.TestVersion3LoneNesting.Inner2.InnerInner")
	proto.RegisterType((*TestVersion4LoneNesting)(nil), "testdata.TestVersion4LoneNesting")
	proto.RegisterType((*TestVersion4LoneNesting_Inner1)(nil), "testdata.TestVersion4LoneNesting.Inner1")
	proto.RegisterType((*TestVersion4LoneNesting_Inner1_InnerInner)(nil), "testdata.TestVersion4LoneNesting.Inner1.InnerInner")
	proto.RegisterType((*TestVersion4LoneNesting_Inner2)(nil), "testdata.TestVersion4LoneNesting.Inner2")
	proto.RegisterType((*TestVersion4LoneNesting_Inner2_InnerInner)(nil), "testdata.TestVersion4LoneNesting.Inner2.InnerInner")
	proto.RegisterType((*TestVersionFD1)(nil), "testdata.TestVersionFD1")
	proto.RegisterType((*TestVersionFD1WithExtraAny)(nil), "testdata.TestVersionFD1WithExtraAny")
	proto.RegisterType((*AnyWithExtra)(nil), "testdata.AnyWithExtra")
	proto.RegisterType((*TestUpdatedTxRaw)(nil), "testdata.TestUpdatedTxRaw")
	proto.RegisterType((*TestUpdatedTxBody)(nil), "testdata.TestUpdatedTxBody")
	proto.RegisterType((*TestUpdatedAuthInfo)(nil), "testdata.TestUpdatedAuthInfo")
	proto.RegisterType((*TestRepeatedUints)(nil), "testdata.TestRepeatedUints")
}

func init() { proto.RegisterFile("unknonwnproto.proto", fileDescriptor_448ea787339d1228) }

var fileDescriptor_448ea787339d1228 = []byte{
	
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x59, 0x4f, 0x6f, 0x1b, 0xc7,
	0x15, 0xd7, 0x70, 0x49, 0x89, 0x7c, 0xa2, 0x69, 0x66, 0x6c, 0xb4, 0x1b, 0x3a, 0x66, 0x98, 0x85,
	0xeb, 0xb0, 0x41, 0x43, 0x9a, 0x4b, 0x06, 0x28, 0x72, 0x32, 0xe9, 0x58, 0x95, 0x01, 0x57, 0x2e,
	0xa6, 0x4e, 0x5a, 0xf8, 0x42, 0x2c, 0xb9, 0x43, 0x72, 0x21, 0x72, 0x46, 0xdd, 0x99, 0xb5, 0xc8,
	0x5b, 0xd1, 0x1e, 0x7a, 0xcd, 0xa5, 0x28, 0xd0, 0x6f, 0xd0, 0x53, 0x91, 0x6f, 0xd0, 0xa3, 0x2f,
	0x05, 0x7c, 0x29, 0x50, 0xa0, 0x40, 0x50, 0xd8, 0xd7, 0x7e, 0x83, 0xa2, 0x48, 0x31, 0xb3, 0x7f,
	0xb8, 0x94, 0x44, 0x85, 0x52, 0xda, 0x18, 0x02, 0x72, 0x11, 0x67, 0xde, 0xfe, 0xe6, 0xcd, 0x7b,
	0xbf, 0xf7, 0x67, 0x77, 0x46, 0x70, 0x23, 0x60, 0x87, 0x8c, 0xb3, 0x63, 0x76, 0xe4, 0x73, 0xc9,
	0x1b, 0xfa, 0x2f, 0xce, 0x4b, 0x2a, 0xa4, 0xeb, 0x48, 0xa7, 0x72, 0x73, 0xcc, 0xc7, 0x5c, 0x0b,
	0x9b, 0x6a, 0x14, 0x3e, 0xaf, 0xbc, 0x3d, 0xe6, 0x7c, 0x3c, 0xa5, 0x4d, 0x3d, 0x1b, 0x04, 0xa3,
	0xa6, 0xc3, 0x16, 0xd1, 0xa3, 0xca, 0x90, 0x8b, 0x19, 0x17, 0x4d, 0x39, 0x6f, 0x3e, 0x6f, 0x0d,
	0xa8, 0x74, 0x5a, 0x4d, 0x39, 0x0f, 0x9f, 0x59, 0x12, 0x0a, 0x0f, 0x02, 0x21, 0xf9, 0x8c, 0xfa,
	0x2d, 0x5c, 0x82, 0x8c, 0xe7, 0x9a, 0xa8, 0x86, 0xea, 0x39, 0x92, 0xf1, 0x5c, 0x8c, 0x21, 0xcb,
	0x9c, 0x19, 0x35, 0x33, 0x35, 0x54, 0x2f, 0x10, 0x3d, 0xc6, 0x3f, 0x84, 0xb2, 0x08, 0x06, 0x62,
	0xe8, 0x7b, 0x47, 0xd2, 0xe3, 0xac, 0x3f, 0xa2, 0xd4, 0x34, 0x6a, 0xa8, 0x9e, 0x21, 0xd7, 0xd3,
	0xf2, 0x3d, 0x4a, 0xb1, 0x09, 0x3b, 0x47, 0xce, 0x62, 0x46, 0x99, 0x34, 0x77, 0xb4, 0x86, 0x78,
	0x6a, 0x7d, 0x91, 0x59, 0x6e, 0x6b, 0x9f, 0xda, 0xb6, 0x02, 0x79, 0x8f, 0xb9, 0x81, 0x90, 0xfe,
	0x42, 0x6f, 0x9d, 0x23, 0xc9, 0x3c, 0x31, 0xc9, 0x48, 0x99, 0x74, 0x13, 0x72, 0x23, 0x7a, 0x4c,
	0x7d, 0x33, 0xab, 0xed, 0x08, 0x27, 0xf8, 0x16, 0xe4, 0x7d, 0x2a, 0xa8, 0xff, 0x9c, 0xba, 0xe6,
	0x1f, 0xf2, 0x35, 0x54, 0x37, 0x48, 0x22, 0xc0, 0x3f, 0x82, 0xec, 0xd0, 0x93, 0x0b, 0x73, 0xbb,
	0x86, 0xea, 0x25, 0xdb, 0x6c, 0xc4, 0xe4, 0x36, 0x12, 0xab, 0x1a, 0x0f, 0x3c, 0xb9, 0x20, 0x1a,
	0x85, 0x3f, 0x86, 0x6b, 0x33, 0x4f, 0x0c, 0xe9, 0x74, 0xea, 0x30, 0xca, 0x03, 0x61, 0x42, 0x0d,
	0xd5, 0x77, 0xed, 0x9b, 0x8d, 0x90, 0xf3, 0x46, 0xcc, 0x79, 0xa3, 0xcb, 0x16, 0x64, 0x15, 0x6a,
	0xfd, 0x04, 0xb2, 0x4a, 0x13, 0xce, 0x43, 0xf6, 0xb1, 0xc3, 0x45, 0x79, 0x0b, 0x97, 0x00, 0x1e,
	0x73, 0xd1, 0x65, 0x63, 0x3a, 0xa5, 0xa2, 0x8c, 0x70, 0x11, 0xf2, 0x3f, 0x73, 0xa6, 0xbc, 0x3b,
	0x95, 0xbc, 0x9c, 0xc1, 0x00, 0xdb, 0x3f, 0xe5, 0x62, 0xc8, 0x8f, 0xcb, 0x06, 0xde, 0x85, 0x9d,
	0x03, 0xc7, 0xf3, 0xf9, 0xc0, 0x2b, 0x67, 0xad, 0x06, 0xe4, 0x0f, 0xa8, 0x90, 0xd4, 0xed, 0x74,
	0x37, 0x09, 0x94, 0xf5, 0x37, 0x14, 0x2f, 0x68, 0x6f, 0xb4, 0x00, 0x5b, 0x90, 0x71, 0x3a, 0x66,
	0xb6, 0x66, 0xd4, 0x77, 0x6d, 0xbc, 0x64, 0x24, 0xde, 0x94, 0x64, 0x9c, 0x0e, 0x6e, 0x43, 0xce,
	0x63, 0x2e, 0x9d, 0x9b, 0x39, 0x0d, 0xbb, 0x7d, 0x12, 0xd6, 0xee, 0x36, 0x1e, 0xa9, 0xe7, 0x0f,
	0x99, 0xf4, 0x17, 0x24, 0xc4, 0x56, 0x1e, 0x03, 0x2c, 0x85, 0xb8, 0x0c, 0xc6, 0x21, 0x5d, 0x68,
	0x5b, 0x0c, 0xa2, 0x86, 0xb8, 0x0e, 0xb9, 0xe7, 0xce, 0x34, 0x08, 0xad, 0x39, 0x7b, 0xef, 0x10,
	0xf0, 0x71, 0xe6, 0xc7, 0xc8, 0x7a, 0x16, 0xbb, 0x65, 0x6f, 0xe6, 0xd6, 0x07, 0xb0, 0xcd, 0x34,
	0x5e, 0xe7, 0xcc, 0x19, 0xea, 0xdb, 0x5d, 0x12, 0x21, 0xac, 0xbd, 0x58, 0x77, 0xeb, 0xb4, 0xee,
	0xa5, 0x9e, 0x35, 0x66, 0xda, 0x4b, 0x3d, 0xf7, 0x93, 0x58, 0xf5, 0x4e, 0xe9, 0x29, 0x83, 0xe1,
	0x8c, 0x69, 0x94, 0xd8, 0x6a, 0x78, 0x56, 0x4e, 0x5b, 0x6e, 0x12, 0xbc, 0x4b, 0x6a, 0x50, 0xe1,
	0x1c, 0xac, 0x0f, 0x67, 0x8f, 0x64, 0x06, 0x1d, 0x8b, 0x25, 0x5c, 0x9e, 0xb9, 0x8b, 0xaa, 0x6d,
	0xb5, 0x0b, 0x22, 0x6a, 0xb8, 0x01, 0x93, 0xbd, 0x98, 0x01, 0x55, 0x93, 0x3e, 0x0f, 0x24, 0xd5,
	0x35, 0x59, 0x20, 0xe1, 0xc4, 0xfa, 0x65, 0xc2, 0x6f, 0xef, 0x12, 0xfc, 0x2e, 0xb5, 0x47, 0x0c,
	0x18, 0x09, 0x03, 0xd6, 0x6f, 0x52, 0x1d, 0xa5, 0xbd, 0x51, 0x5e, 0x94, 0x20, 0x23, 0x46, 0x51,
	0xeb, 0xca, 0x88, 0x11, 0x7e, 0x07, 0x0a, 0x22, 0xf0, 0x87, 0x13, 0xc7, 0x1f, 0xd3, 0xa8, 0x93,
	0x2c, 0x05, 0xb8, 0x06, 0xbb, 0x2e, 0x15, 0xd2, 0x63, 0x8e, 0xea, 0x6e, 0x66, 0x4e, 0x2b, 0x4a,
	0x8b, 0xf0, 0x5d, 0x28, 0x0d, 0x7d, 0xea, 0x7a, 0xb2, 0x3f, 0x74, 0x7c, 0xb7, 0xcf, 0x78, 0xd8,
	0xf4, 0xf6, 0xb7, 0x48, 0x31, 0x94, 0x3f, 0x70, 0x7c, 0xf7, 0x80, 0xe3, 0xdb, 0x50, 0x18, 0x4e,
	0xe8, 0xaf, 0x02, 0xaa, 0x20, 0xf9, 0x08, 0x92, 0x0f, 0x45, 0x07, 0x1c, 0x37, 0x21, 0xcf, 0x7d,
	0x6f, 0xec, 0x31, 0x67, 0x6a, 0x16, 0x34, 0x11, 0x37, 0x4e, 0x77, 0xa7, 0x16, 0x49, 0x40, 0xbd,
	0x42, 0xd2, 0x65, 0xad, 0x7f, 0x65, 0xa0, 0xf8, 0x94, 0x0a, 0xf9, 0x19, 0xf5, 0x85, 0xc7, 0x59,
	0x0b, 0x17, 0x01, 0xcd, 0xa3, 0x4a, 0x43, 0x73, 0x7c, 0x07, 0x90, 0x13, 0x91, 0xfb, 0xbd, 0xa5,
	0xce, 0xf4, 0x02, 0x82, 0x1c, 0x85, 0x1a, 0x44, 0x01, 0x5e, 0x8b, 0x1a, 0x28, 0xd4, 0x30, 0x4a,
	0xae, 0xb5, 0xa8, 0x21, 0xfe, 0x00, 0x90, 0x1b, 0xb5, 0x8a, 0x35, 0xa8, 0x5e, 0xf6, 0xc5, 0x97,
	0xef, 0x6e, 0x11, 0xe4, 0xe2, 0x12, 0x20, 0xaa, 0xfb, 0x71, 0x6e, 0x7f, 0x8b, 0x20, 0x8a, 0xef,
	0x02, 0x1a, 0x69, 0x0a, 0xd7, 0xae, 0x55, 0xb8, 0x11, 0xb6, 0x00, 0x8d, 0x35, 0x8f, 0xeb, 0x1a,
	0x32, 0x1a, 0x2b, 0x6b, 0x27, 0x66, 0xe1, 0x7c, 0x6b, 0x27, 0xf8, 0x7d, 0x40, 0x87, 0x66, 0x71,
	0x2d, 0xe7, 0xbd, 0xec, 0xcb, 0x2f, 0xdf, 0x45, 0x04, 0x1d, 0xf6, 0x72, 0x60, 0x88, 0x60, 0x66,
	0xfd, 0xd6, 0x58, 0xa1, 0xdb, 0xbe, 0x28, 0xdd, 0xf6, 0x46, 0x74, 0xdb, 0x1b, 0xd1, 0x6d, 0x2b,
	0xba, 0xef, 0x7c, 0x1d, 0xdd, 0xf6, 0xa5, 0x88, 0xb6, 0xdf, 0x14, 0xd1, 0xf8, 0x16, 0x14, 0x18,
	0x3d, 0xee, 0x8f, 0x3c, 0x3a, 0x75, 0xcd, 0xb7, 0x6b, 0xa8, 0x9e, 0x25, 0x79, 0x46, 0x8f, 0xf7,
	0xd4, 0x3c, 0x8e, 0xc2, 0xef, 0x57, 0xa3, 0xd0, 0xbe, 0x68, 0x14, 0xda, 0x1b, 0x45, 0xa1, 0xbd,
	0x51, 0x14, 0xda, 0x1b, 0x45, 0xa1, 0x7d, 0xa9, 0x28, 0xb4, 0xdf, 0x58, 0x14, 0x3e, 0x04, 0xcc,
	0x38, 0xeb, 0x0f, 0x7d, 0x4f, 0x7a, 0x43, 0x67, 0x1a, 0x85, 0xe3, 0x77, 0xba, 0x77, 0x91, 0x32,
	0xe3, 0xec, 0x41, 0xf4, 0x64, 0x25, 0x2e, 0xff, 0xce, 0x40, 0x25, 0x6d, 0xfe, 0x63, 0xce, 0xe8,
	0x13, 0x46, 0x9f, 0x8c, 0x3e, 0x53, 0xaf, 0xf2, 0x2b, 0x1a, 0xa5, 0x2b, 0xc3, 0xfe, 0x7f, 0xb6,
	0xe1, 0xfb, 0x27, 0xd9, 0x3f, 0xd0, 0x6f, 0xab, 0xf1, 0x15, 0xa1, 0xbe, 0xb5, 0x2c, 0x88, 0xf7,
	0xce, 0x46, 0xa5, 0x7c, 0xba, 0x22, 0xb5, 0x81, 0xef, 0xc3, 0xb6, 0xc7, 0x18, 0xf5, 0x5b, 0x66,
	0x49, 0x2b, 0xaf, 0x7f, 0xad, 0x67, 0x8d, 0x47, 0x1a, 0x4f, 0xa2, 0x75, 0x89, 0x06, 0xdb, 0xbc,
	0x7e, 0x21, 0x0d, 0x76, 0xa4, 0xc1, 0xae, 0xfc, 0x09, 0xc1, 0x76, 0xa8, 0x34, 0xf5, 0x9d, 0x64,
	0xac, 0xfd, 0x4e, 0x7a, 0xa4, 0x3e, 0xf9, 0x19, 0xf5, 0xa3, 0xe8, 0xb7, 0x37, 0xb5, 0x38, 0xfc,
	0xd1, 0x7f, 0x48, 0xa8, 0xa1, 0x72, 0x4f, 0x1d, 0x04, 0x62, 0x61, 0x6a, 0xf3, 0x42, 0xbc, 0xb9,
	0x3e, 0x93, 0x45, 0x9b, 0xab, 0x71, 0xe5, 0xcf, 0xb1, 0xad, 0xf6, 0x29, 0xb8, 0x09, 0x3b, 0x43,
	0x1e, 0xb0, 0xf8, 0x90, 0x58, 0x20, 0xf1, 0xf4, 0xb2, 0x16, 0xdb, 0xff, 0x0b, 0x8b, 0xe3, 0xfa,
	0xfb, 0x6a, 0xb5, 0xfe, 0x3a, 0xdf, 0xd5, 0xdf, 0x15, 0xaa, 0xbf, 0xce, 0x37, 0xae, 0xbf, 0xce,
	0xb7, 0x5c, 0x7f, 0x9d, 0x6f, 0x54, 0x7f, 0xc6, 0xda, 0xfa, 0xfb, 0xe2, 0xff, 0x56, 0x7f, 0x9d,
	0x8d, 0xea, 0xcf, 0x3e, 0xb7, 0xfe, 0x6e, 0xa6, 0x2f, 0x0e, 0x8c, 0xe8, 0x92, 0x20, 0xae, 0xc0,
	0xbf, 0x22, 0x28, 0xa5, 0xf6, 0xdb, 0xfb, 0xe4, 0x72, 0xc7, 0xa1, 0x37, 0x7e, 0x2c, 0x89, 0xfd,
	0xf9, 0x07, 0x5a, 0xf9, 0x9e, 0xda, 0xfb, 0xa4, 0xf5, 0x0b, 0x4f, 0x4e, 0x1e, 0xce, 0xa5, 0xef,
	0x74, 0xd9, 0xe2, 0x5b, 0xf5, 0xed, 0xce, 0xd2, 0xb7, 0x14, 0xae, 0xcb, 0x16, 0x89, 0x45, 0x17,
	0xf6, 0xee, 0x29, 0x14, 0xd3, 0xeb, 0x71, 0x5d, 0x39, 0x80, 0xd6, 0xd3, 0x17, 0x77, 0x00, 0x47,
	0x39, 0x1e, 0x76, 0x46, 0x43, 0x75, 0xc0, 0x62, 0xd8, 0x01, 0xf5, 0x6c, 0x68, 0xfd, 0x05, 0x41,
	0x59, 0x6d, 0xf8, 0xe9, 0x91, 0xeb, 0x48, 0xea, 0x3e, 0x9d, 0x13, 0xe7, 0x18, 0xdf, 0x06, 0x18,
	0x70, 0x77, 0xd1, 0x1f, 0x2c, 0x24, 0x15, 0x7a, 0x8f, 0x22, 0x29, 0x28, 0x49, 0x4f, 0x09, 0xf0,
	0x5d, 0xb8, 0xee, 0x04, 0x72, 0xd2, 0xf7, 0xd8, 0x88, 0x47, 0x98, 0x8c, 0xc6, 0x5c, 0x53, 0xe2,
	0x47, 0x6c, 0xc4, 0x43, 0x5c, 0x15, 0x40, 0x78, 0x63, 0xe6, 0xc8, 0xc0, 0xa7, 0xc2, 0x34, 0x6a,
	0x46, 0xbd, 0x48, 0x52, 0x12, 0x5c, 0x85, 0xdd, 0xe4, 0xec, 0xd2, 0xff, 0x48, 0xdf, 0x18, 0x14,
	0x49, 0x21, 0x3e, 0xbd, 0x7c, 0x84, 0x7f, 0x00, 0xa5, 0xe5, 0xf3, 0xd6, 0x3d, 0xbb, 0x63, 0xfe,
	0x3a, 0xaf, 0x31, 0xc5, 0x18, 0xa3, 0x84, 0xd6, 0xe7, 0x06, 0xbc, 0xb5, 0xe2, 0x42, 0x8f, 0xbb,
	0x0b, 0x7c, 0x0f, 0xf2, 0x33, 0x2a, 0x84, 0x33, 0xd6, 0x1e, 0x18, 0x6b, 0x93, 0x2c, 0x41, 0xa9,
	0xea, 0x9e, 0xd1, 0x19, 0x8f, 0xab, 0x5b, 0x8d, 0x95, 0x09, 0xd2, 0x9b, 0x51, 0x1e, 0xc8, 0xfe,
	0x84, 0x7a, 0xe3, 0x89, 0x8c, 0x78, 0xbc, 0x16, 0x49, 0xf7, 0xb5, 0x10, 0xdf, 0x81, 0x92, 0xe0,
	0x33, 0xda, 0x5f, 0x1e, 0xc5, 0xb2, 0xfa, 0x28, 0x56, 0x54, 0xd2, 0x83, 0xc8, 0x58, 0xbc, 0x0f,
	0xef, 0xad, 0xa2, 0xfa, 0x67, 0x34, 0xe6, 0x3f, 0x86, 0x8d, 0xf9, 0x9d, 0xf4, 0xca, 0x83, 0x93,
	0x4d, 0xba, 0x07, 0x6f, 0xd1, 0xb9, 0xa4, 0x4c, 0xe5, 0x48, 0x9f, 0xeb, 0xeb, 0x64, 0x61, 0x7e,
	0xb5, 0x73, 0x8e, 0x9b, 0xe5, 0x04, 0xff, 0x24, 0x84, 0xe3, 0x67, 0x50, 0x5d, 0xd9, 0xfe, 0x0c,
	0x85, 0xd7, 0xcf, 0x51, 0x78, 0x2b, 0xf5, 0xe6, 0x78, 0x78, 0x42, 0xb7, 0xf5, 0x02, 0xc1, 0x8d,
	0x54, 0x48, 0xba, 0x51, 0x5a, 0xe0, 0xfb, 0x50, 0x54, 0xf1, 0xa7, 0xbe, 0xce, 0x9d, 0x38, 0x30,
	0xb7, 0x1b, 0xe1, 0xf5, 0x7b, 0x43, 0xce, 0x1b, 0xd1, 0xf5, 0x7b, 0xe3, 0xe7, 0x1a, 0xa6, 0x16,
	0x91, 0x5d, 0x91, 0x8c, 0x05, 0xae, 0x2f, 0xef, 0xdc, 0x54, 0xd1, 0x9c, 0x5e, 0xb8, 0x47, 0x69,
	0x78, 0x17, 0xb7, 0x92, 0x5d, 0x6d, 0x1d, 0xb7, 0x54, 0x76, 0xb5, 0x37, 0xcd, 0xae, 0xf7, 0xc3,
	0xe4, 0x22, 0xf4, 0x88, 0x2a, 0x57, 0x3e, 0xf5, 0x98, 0xd4, 0xa9, 0xc2, 0x82, 0x59, 0x68, 0x7f,
	0x96, 0xe8, 0x71, 0x6f, 0xff, 0xc5, 0xab, 0x2a, 0x7a, 0xf9, 0xaa, 0x8a, 0xfe, 0xf9, 0xaa, 0x8a,
	0x3e, 0x7f, 0x5d, 0xdd, 0x7a, 0xf9, 0xba, 0xba, 0xf5, 0xf7, 0xd7, 0xd5, 0xad, 0x67, 0x8d, 0xb1,
	0x27, 0x27, 0xc1, 0xa0, 0x31, 0xe4, 0xb3, 0x66, 0xf4, 0x8f, 0x86, 0xf0, 0xe7, 0x43, 0xe1, 0x1e,
	0x36, 0x55, 0xdd, 0x07, 0xd2, 0x9b, 0x36, 0xe3, 0x06, 0x30, 0xd8, 0xd6, 0x44, 0xb7, 0xff, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0xa6, 0x69, 0x10, 0x33, 0xe6, 0x18, 0x00, 0x00,
}

func (m *Customer1) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Customer1) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Customer1) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payment) > 0 {
		i -= len(m.Payment)
		copy(dAtA[i:], m.Payment)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Payment)))
		i--
		dAtA[i] = 0x3a
	}
	if m.SubscriptionFee != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.SubscriptionFee))))
		i--
		dAtA[i] = 0x1d
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Customer2) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Customer2) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Customer2) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Reserved != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.Reserved))
		i--
		dAtA[i] = 0x41
		i--
		dAtA[i] = 0xb8
	}
	if m.Miscellaneous != nil {
		{
			size, err := m.Miscellaneous.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if m.City != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.City))
		i--
		dAtA[i] = 0x30
	}
	if m.Fewer != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Fewer))))
		i--
		dAtA[i] = 0x25
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Industry != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.Industry))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Nested4A) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Nested4A) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Nested4A) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Nested3A) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Nested3A) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Nested3A) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		for k := range m.Index {
			v := m.Index[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.A4) > 0 {
		for iNdEx := len(m.A4) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.A4[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Nested2A) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Nested2A) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Nested2A) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Nested != nil {
		{
			size, err := m.Nested.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Nested1A) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Nested1A) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Nested1A) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Nested != nil {
		{
			size, err := m.Nested.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Nested4B) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Nested4B) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Nested4B) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Age != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.Age))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Nested3B) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Nested3B) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Nested3B) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.B4) > 0 {
		for iNdEx := len(m.B4) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.B4[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Age != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.Age))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Nested2B) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Nested2B) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Nested2B) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Route) > 0 {
		i -= len(m.Route)
		copy(dAtA[i:], m.Route)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Route)))
		i--
		dAtA[i] = 0x22
	}
	if m.Nested != nil {
		{
			size, err := m.Nested.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Fee != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Fee))))
		i--
		dAtA[i] = 0x11
	}
	if m.Id != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Nested1B) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Nested1B) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Nested1B) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Age != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.Age))
		i--
		dAtA[i] = 0x18
	}
	if m.Nested != nil {
		{
			size, err := m.Nested.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Customer3) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Customer3) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Customer3) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Original != nil {
		{
			size, err := m.Original.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.Payment != nil {
		{
			size := m.Payment.Size()
			i -= size
			if _, err := m.Payment.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.Destination) > 0 {
		i -= len(m.Destination)
		copy(dAtA[i:], m.Destination)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Destination)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Surcharge != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Surcharge))))
		i--
		dAtA[i] = 0x25
	}
	if m.Sf != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Sf))))
		i--
		dAtA[i] = 0x1d
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Customer3_CreditCardNo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Customer3_CreditCardNo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.CreditCardNo)
	copy(dAtA[i:], m.CreditCardNo)
	i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.CreditCardNo)))
	i--
	dAtA[i] = 0x3a
	return len(dAtA) - i, nil
}
func (m *Customer3_ChequeNo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Customer3_ChequeNo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.ChequeNo)
	copy(dAtA[i:], m.ChequeNo)
	i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.ChequeNo)))
	i--
	dAtA[i] = 0x42
	return len(dAtA) - i, nil
}
func (m *TestVersion1) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestVersion1) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion1) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Customer1 != nil {
		{
			size, err := m.Customer1.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x62
	}
	if len(m.H) > 0 {
		for iNdEx := len(m.H) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.H[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.G != nil {
		{
			size, err := m.G.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.D) > 0 {
		for iNdEx := len(m.D) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.D[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.C) > 0 {
		for iNdEx := len(m.C) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.C[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.B != nil {
		{
			size, err := m.B.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.A != nil {
		{
			size, err := m.A.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.X != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.X))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TestVersion1_E) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion1_E) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.E))
	i--
	dAtA[i] = 0x30
	return len(dAtA) - i, nil
}
func (m *TestVersion1_F) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion1_F) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.F != nil {
		{
			size, err := m.F.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *TestVersion2) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestVersion2) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion2) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NewField != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.NewField))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xc8
	}
	if m.Customer1 != nil {
		{
			size, err := m.Customer1.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x62
	}
	if len(m.H) > 0 {
		for iNdEx := len(m.H) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.H[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.G != nil {
		{
			size, err := m.G.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.D) > 0 {
		for iNdEx := len(m.D) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.D[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.C) > 0 {
		for iNdEx := len(m.C) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.C[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.B != nil {
		{
			size, err := m.B.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.A != nil {
		{
			size, err := m.A.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.X != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.X))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TestVersion2_E) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion2_E) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.E))
	i--
	dAtA[i] = 0x30
	return len(dAtA) - i, nil
}
func (m *TestVersion2_F) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion2_F) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.F != nil {
		{
			size, err := m.F.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *TestVersion3) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestVersion3) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion3) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NonCriticalField) > 0 {
		i -= len(m.NonCriticalField)
		copy(dAtA[i:], m.NonCriticalField)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.NonCriticalField)))
		i--
		dAtA[i] = 0x40
		i--
		dAtA[i] = 0xba
	}
	if m.Customer1 != nil {
		{
			size, err := m.Customer1.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x62
	}
	if len(m.H) > 0 {
		for iNdEx := len(m.H) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.H[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.G != nil {
		{
			size, err := m.G.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.D) > 0 {
		for iNdEx := len(m.D) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.D[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.C) > 0 {
		for iNdEx := len(m.C) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.C[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.B != nil {
		{
			size, err := m.B.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.A != nil {
		{
			size, err := m.A.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.X != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.X))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TestVersion3_E) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion3_E) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.E))
	i--
	dAtA[i] = 0x30
	return len(dAtA) - i, nil
}
func (m *TestVersion3_F) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion3_F) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.F != nil {
		{
			size, err := m.F.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *TestVersion3LoneOneOfValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestVersion3LoneOneOfValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion3LoneOneOfValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NonCriticalField) > 0 {
		i -= len(m.NonCriticalField)
		copy(dAtA[i:], m.NonCriticalField)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.NonCriticalField)))
		i--
		dAtA[i] = 0x40
		i--
		dAtA[i] = 0xba
	}
	if m.Customer1 != nil {
		{
			size, err := m.Customer1.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x62
	}
	if len(m.H) > 0 {
		for iNdEx := len(m.H) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.H[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.G != nil {
		{
			size, err := m.G.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.D) > 0 {
		for iNdEx := len(m.D) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.D[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.C) > 0 {
		for iNdEx := len(m.C) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.C[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.B != nil {
		{
			size, err := m.B.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.A != nil {
		{
			size, err := m.A.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.X != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.X))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TestVersion3LoneOneOfValue_E) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion3LoneOneOfValue_E) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.E))
	i--
	dAtA[i] = 0x30
	return len(dAtA) - i, nil
}
func (m *TestVersion3LoneNesting) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestVersion3LoneNesting) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion3LoneNesting) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NonCriticalField) > 0 {
		i -= len(m.NonCriticalField)
		copy(dAtA[i:], m.NonCriticalField)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.NonCriticalField)))
		i--
		dAtA[i] = 0x40
		i--
		dAtA[i] = 0xba
	}
	if m.Inner2 != nil {
		{
			size, err := m.Inner2.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x7a
	}
	if m.Inner1 != nil {
		{
			size, err := m.Inner1.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x72
	}
	if m.Customer1 != nil {
		{
			size, err := m.Customer1.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x62
	}
	if len(m.H) > 0 {
		for iNdEx := len(m.H) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.H[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.G != nil {
		{
			size, err := m.G.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.D) > 0 {
		for iNdEx := len(m.D) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.D[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.C) > 0 {
		for iNdEx := len(m.C) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.C[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.B != nil {
		{
			size, err := m.B.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.A != nil {
		{
			size, err := m.A.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.X != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.X))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TestVersion3LoneNesting_F) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion3LoneNesting_F) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.F != nil {
		{
			size, err := m.F.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *TestVersion3LoneNesting_Inner1) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestVersion3LoneNesting_Inner1) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion3LoneNesting_Inner1) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Inner != nil {
		{
			size, err := m.Inner.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TestVersion3LoneNesting_Inner1_InnerInner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestVersion3LoneNesting_Inner1_InnerInner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion3LoneNesting_Inner1_InnerInner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.City) > 0 {
		i -= len(m.City)
		copy(dAtA[i:], m.City)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.City)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TestVersion3LoneNesting_Inner2) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestVersion3LoneNesting_Inner2) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion3LoneNesting_Inner2) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Inner != nil {
		{
			size, err := m.Inner.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Country) > 0 {
		i -= len(m.Country)
		copy(dAtA[i:], m.Country)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Country)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TestVersion3LoneNesting_Inner2_InnerInner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestVersion3LoneNesting_Inner2_InnerInner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion3LoneNesting_Inner2_InnerInner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.City) > 0 {
		i -= len(m.City)
		copy(dAtA[i:], m.City)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.City)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TestVersion4LoneNesting) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestVersion4LoneNesting) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion4LoneNesting) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NonCriticalField) > 0 {
		i -= len(m.NonCriticalField)
		copy(dAtA[i:], m.NonCriticalField)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.NonCriticalField)))
		i--
		dAtA[i] = 0x40
		i--
		dAtA[i] = 0xba
	}
	if m.Inner2 != nil {
		{
			size, err := m.Inner2.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x7a
	}
	if m.Inner1 != nil {
		{
			size, err := m.Inner1.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x72
	}
	if m.Customer1 != nil {
		{
			size, err := m.Customer1.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x62
	}
	if len(m.H) > 0 {
		for iNdEx := len(m.H) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.H[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.G != nil {
		{
			size, err := m.G.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.D) > 0 {
		for iNdEx := len(m.D) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.D[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.C) > 0 {
		for iNdEx := len(m.C) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.C[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.B != nil {
		{
			size, err := m.B.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.A != nil {
		{
			size, err := m.A.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.X != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.X))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TestVersion4LoneNesting_F) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion4LoneNesting_F) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.F != nil {
		{
			size, err := m.F.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *TestVersion4LoneNesting_Inner1) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestVersion4LoneNesting_Inner1) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion4LoneNesting_Inner1) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Inner != nil {
		{
			size, err := m.Inner.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TestVersion4LoneNesting_Inner1_InnerInner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestVersion4LoneNesting_Inner1_InnerInner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion4LoneNesting_Inner1_InnerInner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.City) > 0 {
		i -= len(m.City)
		copy(dAtA[i:], m.City)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.City)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TestVersion4LoneNesting_Inner2) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestVersion4LoneNesting_Inner2) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion4LoneNesting_Inner2) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Inner != nil {
		{
			size, err := m.Inner.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Country) > 0 {
		i -= len(m.Country)
		copy(dAtA[i:], m.Country)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Country)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TestVersion4LoneNesting_Inner2_InnerInner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestVersion4LoneNesting_Inner2_InnerInner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersion4LoneNesting_Inner2_InnerInner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TestVersionFD1) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestVersionFD1) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersionFD1) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.H) > 0 {
		for iNdEx := len(m.H) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.H[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.G != nil {
		{
			size, err := m.G.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.A != nil {
		{
			size, err := m.A.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.X != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.X))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TestVersionFD1_E) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersionFD1_E) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.E))
	i--
	dAtA[i] = 0x30
	return len(dAtA) - i, nil
}
func (m *TestVersionFD1_F) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersionFD1_F) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.F != nil {
		{
			size, err := m.F.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *TestVersionFD1WithExtraAny) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestVersionFD1WithExtraAny) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersionFD1WithExtraAny) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.H) > 0 {
		for iNdEx := len(m.H) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.H[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.G != nil {
		{
			size, err := m.G.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.A != nil {
		{
			size, err := m.A.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.X != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.X))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TestVersionFD1WithExtraAny_E) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersionFD1WithExtraAny_E) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.E))
	i--
	dAtA[i] = 0x30
	return len(dAtA) - i, nil
}
func (m *TestVersionFD1WithExtraAny_F) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestVersionFD1WithExtraAny_F) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.F != nil {
		{
			size, err := m.F.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *AnyWithExtra) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnyWithExtra) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnyWithExtra) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.C != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.C))
		i--
		dAtA[i] = 0x20
	}
	if m.B != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.B))
		i--
		dAtA[i] = 0x18
	}
	if m.Any != nil {
		{
			size, err := m.Any.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TestUpdatedTxRaw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestUpdatedTxRaw) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestUpdatedTxRaw) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewField_1024) > 0 {
		i -= len(m.NewField_1024)
		copy(dAtA[i:], m.NewField_1024)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.NewField_1024)))
		i--
		dAtA[i] = 0x40
		i--
		dAtA[i] = 0x82
	}
	if len(m.NewField_5) > 0 {
		i -= len(m.NewField_5)
		copy(dAtA[i:], m.NewField_5)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.NewField_5)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Signatures) > 0 {
		for iNdEx := len(m.Signatures) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signatures[iNdEx])
			copy(dAtA[i:], m.Signatures[iNdEx])
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Signatures[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.AuthInfoBytes) > 0 {
		i -= len(m.AuthInfoBytes)
		copy(dAtA[i:], m.AuthInfoBytes)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.AuthInfoBytes)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BodyBytes) > 0 {
		i -= len(m.BodyBytes)
		copy(dAtA[i:], m.BodyBytes)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.BodyBytes)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TestUpdatedTxBody) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestUpdatedTxBody) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestUpdatedTxBody) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x7f
			i--
			dAtA[i] = 0xfa
		}
	}
	if len(m.SomeNewFieldNonCriticalField) > 0 {
		i -= len(m.SomeNewFieldNonCriticalField)
		copy(dAtA[i:], m.SomeNewFieldNonCriticalField)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.SomeNewFieldNonCriticalField)))
		i--
		dAtA[i] = 0x41
		i--
		dAtA[i] = 0xd2
	}
	if len(m.ExtensionOptions) > 0 {
		for iNdEx := len(m.ExtensionOptions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExtensionOptions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3f
			i--
			dAtA[i] = 0xfa
		}
	}
	if m.SomeNewField != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.SomeNewField))
		i--
		dAtA[i] = 0x20
	}
	if m.TimeoutHeight != 0 {
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(m.TimeoutHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Memo) > 0 {
		i -= len(m.Memo)
		copy(dAtA[i:], m.Memo)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.Memo)))
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
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *TestUpdatedAuthInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestUpdatedAuthInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestUpdatedAuthInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewField_1024) > 0 {
		i -= len(m.NewField_1024)
		copy(dAtA[i:], m.NewField_1024)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.NewField_1024)))
		i--
		dAtA[i] = 0x40
		i--
		dAtA[i] = 0x82
	}
	if len(m.NewField_3) > 0 {
		i -= len(m.NewField_3)
		copy(dAtA[i:], m.NewField_3)
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(len(m.NewField_3)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Fee != nil {
		{
			size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
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
				i = encodeVarintUnknonwnproto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *TestRepeatedUints) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestRepeatedUints) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestRepeatedUints) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Nums) > 0 {
		dAtA54 := make([]byte, len(m.Nums)*10)
		var j53 int
		for _, num := range m.Nums {
			for num >= 1<<7 {
				dAtA54[j53] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j53++
			}
			dAtA54[j53] = uint8(num)
			j53++
		}
		i -= j53
		copy(dAtA[i:], dAtA54[:j53])
		i = encodeVarintUnknonwnproto(dAtA, i, uint64(j53))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUnknonwnproto(dAtA []byte, offset int, v uint64) int {
	offset -= sovUnknonwnproto(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Customer1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.SubscriptionFee != 0 {
		n += 5
	}
	l = len(m.Payment)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}

func (m *Customer2) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.Id))
	}
	if m.Industry != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.Industry))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.Fewer != 0 {
		n += 5
	}
	if m.City != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.City))
	}
	if m.Miscellaneous != nil {
		l = m.Miscellaneous.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.Reserved != 0 {
		n += 2 + sovUnknonwnproto(uint64(m.Reserved))
	}
	return n
}

func (m *Nested4A) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}

func (m *Nested3A) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if len(m.A4) > 0 {
		for _, e := range m.A4 {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	if len(m.Index) > 0 {
		for k, v := range m.Index {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovUnknonwnproto(uint64(l))
			}
			mapEntrySize := 1 + sovUnknonwnproto(uint64(k)) + l
			n += mapEntrySize + 1 + sovUnknonwnproto(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *Nested2A) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.Nested != nil {
		l = m.Nested.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}

func (m *Nested1A) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.Id))
	}
	if m.Nested != nil {
		l = m.Nested.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}

func (m *Nested4B) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.Id))
	}
	if m.Age != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.Age))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}

func (m *Nested3B) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.Id))
	}
	if m.Age != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.Age))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if len(m.B4) > 0 {
		for _, e := range m.B4 {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	return n
}

func (m *Nested2B) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.Id))
	}
	if m.Fee != 0 {
		n += 9
	}
	if m.Nested != nil {
		l = m.Nested.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	l = len(m.Route)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}

func (m *Nested1B) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.Id))
	}
	if m.Nested != nil {
		l = m.Nested.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.Age != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.Age))
	}
	return n
}

func (m *Customer3) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.Sf != 0 {
		n += 5
	}
	if m.Surcharge != 0 {
		n += 5
	}
	l = len(m.Destination)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.Payment != nil {
		n += m.Payment.Size()
	}
	if m.Original != nil {
		l = m.Original.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}

func (m *Customer3_CreditCardNo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CreditCardNo)
	n += 1 + l + sovUnknonwnproto(uint64(l))
	return n
}
func (m *Customer3_ChequeNo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChequeNo)
	n += 1 + l + sovUnknonwnproto(uint64(l))
	return n
}
func (m *TestVersion1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.X != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.X))
	}
	if m.A != nil {
		l = m.A.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.B != nil {
		l = m.B.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if len(m.C) > 0 {
		for _, e := range m.C {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	if len(m.D) > 0 {
		for _, e := range m.D {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	if m.G != nil {
		l = m.G.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if len(m.H) > 0 {
		for _, e := range m.H {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	if m.Customer1 != nil {
		l = m.Customer1.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}

func (m *TestVersion1_E) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovUnknonwnproto(uint64(m.E))
	return n
}
func (m *TestVersion1_F) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.F != nil {
		l = m.F.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}
func (m *TestVersion2) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.X != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.X))
	}
	if m.A != nil {
		l = m.A.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.B != nil {
		l = m.B.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if len(m.C) > 0 {
		for _, e := range m.C {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	if len(m.D) > 0 {
		for _, e := range m.D {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	if m.G != nil {
		l = m.G.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if len(m.H) > 0 {
		for _, e := range m.H {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	if m.Customer1 != nil {
		l = m.Customer1.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.NewField != 0 {
		n += 2 + sovUnknonwnproto(uint64(m.NewField))
	}
	return n
}

func (m *TestVersion2_E) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovUnknonwnproto(uint64(m.E))
	return n
}
func (m *TestVersion2_F) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.F != nil {
		l = m.F.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}
func (m *TestVersion3) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.X != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.X))
	}
	if m.A != nil {
		l = m.A.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.B != nil {
		l = m.B.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if len(m.C) > 0 {
		for _, e := range m.C {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	if len(m.D) > 0 {
		for _, e := range m.D {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	if m.G != nil {
		l = m.G.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if len(m.H) > 0 {
		for _, e := range m.H {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	if m.Customer1 != nil {
		l = m.Customer1.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	l = len(m.NonCriticalField)
	if l > 0 {
		n += 2 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}

func (m *TestVersion3_E) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovUnknonwnproto(uint64(m.E))
	return n
}
func (m *TestVersion3_F) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.F != nil {
		l = m.F.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}
func (m *TestVersion3LoneOneOfValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.X != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.X))
	}
	if m.A != nil {
		l = m.A.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.B != nil {
		l = m.B.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if len(m.C) > 0 {
		for _, e := range m.C {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	if len(m.D) > 0 {
		for _, e := range m.D {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	if m.G != nil {
		l = m.G.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if len(m.H) > 0 {
		for _, e := range m.H {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	if m.Customer1 != nil {
		l = m.Customer1.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	l = len(m.NonCriticalField)
	if l > 0 {
		n += 2 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}

func (m *TestVersion3LoneOneOfValue_E) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovUnknonwnproto(uint64(m.E))
	return n
}
func (m *TestVersion3LoneNesting) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.X != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.X))
	}
	if m.A != nil {
		l = m.A.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.B != nil {
		l = m.B.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if len(m.C) > 0 {
		for _, e := range m.C {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	if len(m.D) > 0 {
		for _, e := range m.D {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	if m.G != nil {
		l = m.G.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if len(m.H) > 0 {
		for _, e := range m.H {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	if m.Customer1 != nil {
		l = m.Customer1.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.Inner1 != nil {
		l = m.Inner1.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.Inner2 != nil {
		l = m.Inner2.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	l = len(m.NonCriticalField)
	if l > 0 {
		n += 2 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}

func (m *TestVersion3LoneNesting_F) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.F != nil {
		l = m.F.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}
func (m *TestVersion3LoneNesting_Inner1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.Inner != nil {
		l = m.Inner.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}

func (m *TestVersion3LoneNesting_Inner1_InnerInner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	l = len(m.City)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}

func (m *TestVersion3LoneNesting_Inner2) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	l = len(m.Country)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.Inner != nil {
		l = m.Inner.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}

func (m *TestVersion3LoneNesting_Inner2_InnerInner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	l = len(m.City)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}

func (m *TestVersion4LoneNesting) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.X != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.X))
	}
	if m.A != nil {
		l = m.A.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.B != nil {
		l = m.B.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if len(m.C) > 0 {
		for _, e := range m.C {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	if len(m.D) > 0 {
		for _, e := range m.D {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	if m.G != nil {
		l = m.G.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if len(m.H) > 0 {
		for _, e := range m.H {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	if m.Customer1 != nil {
		l = m.Customer1.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.Inner1 != nil {
		l = m.Inner1.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.Inner2 != nil {
		l = m.Inner2.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	l = len(m.NonCriticalField)
	if l > 0 {
		n += 2 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}

func (m *TestVersion4LoneNesting_F) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.F != nil {
		l = m.F.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}
func (m *TestVersion4LoneNesting_Inner1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.Inner != nil {
		l = m.Inner.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}

func (m *TestVersion4LoneNesting_Inner1_InnerInner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.Id))
	}
	l = len(m.City)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}

func (m *TestVersion4LoneNesting_Inner2) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	l = len(m.Country)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.Inner != nil {
		l = m.Inner.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}

func (m *TestVersion4LoneNesting_Inner2_InnerInner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.Value != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.Value))
	}
	return n
}

func (m *TestVersionFD1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.X != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.X))
	}
	if m.A != nil {
		l = m.A.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	if m.G != nil {
		l = m.G.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if len(m.H) > 0 {
		for _, e := range m.H {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	return n
}

func (m *TestVersionFD1_E) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovUnknonwnproto(uint64(m.E))
	return n
}
func (m *TestVersionFD1_F) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.F != nil {
		l = m.F.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}
func (m *TestVersionFD1WithExtraAny) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.X != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.X))
	}
	if m.A != nil {
		l = m.A.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	if m.G != nil {
		l = m.G.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if len(m.H) > 0 {
		for _, e := range m.H {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	return n
}

func (m *TestVersionFD1WithExtraAny_E) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovUnknonwnproto(uint64(m.E))
	return n
}
func (m *TestVersionFD1WithExtraAny_F) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.F != nil {
		l = m.F.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}
func (m *AnyWithExtra) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Any != nil {
		l = m.Any.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.B != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.B))
	}
	if m.C != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.C))
	}
	return n
}

func (m *TestUpdatedTxRaw) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BodyBytes)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	l = len(m.AuthInfoBytes)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if len(m.Signatures) > 0 {
		for _, b := range m.Signatures {
			l = len(b)
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	l = len(m.NewField_5)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	l = len(m.NewField_1024)
	if l > 0 {
		n += 2 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}

func (m *TestUpdatedTxBody) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for _, e := range m.Messages {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	if m.TimeoutHeight != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.TimeoutHeight))
	}
	if m.SomeNewField != 0 {
		n += 1 + sovUnknonwnproto(uint64(m.SomeNewField))
	}
	if len(m.ExtensionOptions) > 0 {
		for _, e := range m.ExtensionOptions {
			l = e.Size()
			n += 2 + l + sovUnknonwnproto(uint64(l))
		}
	}
	l = len(m.SomeNewFieldNonCriticalField)
	if l > 0 {
		n += 2 + l + sovUnknonwnproto(uint64(l))
	}
	if len(m.NonCriticalExtensionOptions) > 0 {
		for _, e := range m.NonCriticalExtensionOptions {
			l = e.Size()
			n += 2 + l + sovUnknonwnproto(uint64(l))
		}
	}
	return n
}

func (m *TestUpdatedAuthInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SignerInfos) > 0 {
		for _, e := range m.SignerInfos {
			l = e.Size()
			n += 1 + l + sovUnknonwnproto(uint64(l))
		}
	}
	if m.Fee != nil {
		l = m.Fee.Size()
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	l = len(m.NewField_3)
	if l > 0 {
		n += 1 + l + sovUnknonwnproto(uint64(l))
	}
	l = len(m.NewField_1024)
	if l > 0 {
		n += 2 + l + sovUnknonwnproto(uint64(l))
	}
	return n
}

func (m *TestRepeatedUints) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Nums) > 0 {
		l = 0
		for _, e := range m.Nums {
			l += sovUnknonwnproto(uint64(e))
		}
		n += 1 + sovUnknonwnproto(uint64(l)) + l
	}
	return n
}

func sovUnknonwnproto(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUnknonwnproto(x uint64) (n int) {
	return sovUnknonwnproto(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Customer1) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: Customer1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Customer1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscriptionFee", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.SubscriptionFee = float32(math.Float32frombits(v))
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *Customer2) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: Customer2: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Customer2: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Industry", wireType)
			}
			m.Industry = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Industry |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fewer", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Fewer = float32(math.Float32frombits(v))
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field City", wireType)
			}
			m.City = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.City |= Customer2_City(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Miscellaneous", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Miscellaneous == nil {
				m.Miscellaneous = &types.Any{}
			}
			if err := m.Miscellaneous.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1047:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserved", wireType)
			}
			m.Reserved = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Reserved |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *Nested4A) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: Nested4A: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Nested4A: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *Nested3A) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: Nested3A: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Nested3A: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A4", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.A4 = append(m.A4, &Nested4A{})
			if err := m.A4[len(m.A4)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Index == nil {
				m.Index = make(map[int64]*Nested4A)
			}
			var mapkey int64
			var mapvalue *Nested4A
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowUnknonwnproto
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowUnknonwnproto
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowUnknonwnproto
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthUnknonwnproto
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthUnknonwnproto
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Nested4A{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthUnknonwnproto
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Index[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *Nested2A) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: Nested2A: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Nested2A: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nested", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Nested == nil {
				m.Nested = &Nested3A{}
			}
			if err := m.Nested.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *Nested1A) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: Nested1A: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Nested1A: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nested", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Nested == nil {
				m.Nested = &Nested2A{}
			}
			if err := m.Nested.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *Nested4B) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: Nested4B: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Nested4B: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Age", wireType)
			}
			m.Age = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Age |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *Nested3B) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: Nested3B: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Nested3B: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Age", wireType)
			}
			m.Age = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Age |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field B4", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.B4 = append(m.B4, &Nested4B{})
			if err := m.B4[len(m.B4)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *Nested2B) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: Nested2B: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Nested2B: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Fee = float64(math.Float64frombits(v))
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nested", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Nested == nil {
				m.Nested = &Nested3B{}
			}
			if err := m.Nested.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Route", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Route = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *Nested1B) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: Nested1B: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Nested1B: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nested", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Nested == nil {
				m.Nested = &Nested2B{}
			}
			if err := m.Nested.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Age", wireType)
			}
			m.Age = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Age |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *Customer3) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: Customer3: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Customer3: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sf", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Sf = float32(math.Float32frombits(v))
		case 4:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Surcharge", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Surcharge = float32(math.Float32frombits(v))
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Destination", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Destination = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreditCardNo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payment = &Customer3_CreditCardNo{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChequeNo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payment = &Customer3_ChequeNo{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Original", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Original == nil {
				m.Original = &Customer1{}
			}
			if err := m.Original.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *TestVersion1) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: TestVersion1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestVersion1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			m.X = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.X |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.A == nil {
				m.A = &TestVersion1{}
			}
			if err := m.A.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.B == nil {
				m.B = &TestVersion1{}
			}
			if err := m.B.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.C = append(m.C, &TestVersion1{})
			if err := m.C[len(m.C)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field D", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.D = append(m.D, TestVersion1{})
			if err := m.D[len(m.D)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field E", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Sum = &TestVersion1_E{v}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field F", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &TestVersion1{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &TestVersion1_F{v}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field G", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.G == nil {
				m.G = &types.Any{}
			}
			if err := m.G.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field H", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.H = append(m.H, &TestVersion1{})
			if err := m.H[len(m.H)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Customer1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Customer1 == nil {
				m.Customer1 = &Customer1{}
			}
			if err := m.Customer1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *TestVersion2) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: TestVersion2: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestVersion2: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			m.X = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.X |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.A == nil {
				m.A = &TestVersion2{}
			}
			if err := m.A.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.B == nil {
				m.B = &TestVersion2{}
			}
			if err := m.B.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.C = append(m.C, &TestVersion2{})
			if err := m.C[len(m.C)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field D", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.D = append(m.D, &TestVersion2{})
			if err := m.D[len(m.D)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field E", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Sum = &TestVersion2_E{v}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field F", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &TestVersion2{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &TestVersion2_F{v}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field G", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.G == nil {
				m.G = &types.Any{}
			}
			if err := m.G.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field H", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.H = append(m.H, &TestVersion1{})
			if err := m.H[len(m.H)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Customer1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Customer1 == nil {
				m.Customer1 = &Customer1{}
			}
			if err := m.Customer1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 25:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewField", wireType)
			}
			m.NewField = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NewField |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *TestVersion3) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: TestVersion3: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestVersion3: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			m.X = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.X |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.A == nil {
				m.A = &TestVersion3{}
			}
			if err := m.A.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.B == nil {
				m.B = &TestVersion3{}
			}
			if err := m.B.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.C = append(m.C, &TestVersion3{})
			if err := m.C[len(m.C)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field D", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.D = append(m.D, &TestVersion3{})
			if err := m.D[len(m.D)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field E", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Sum = &TestVersion3_E{v}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field F", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &TestVersion3{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &TestVersion3_F{v}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field G", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.G == nil {
				m.G = &types.Any{}
			}
			if err := m.G.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field H", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.H = append(m.H, &TestVersion1{})
			if err := m.H[len(m.H)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Customer1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Customer1 == nil {
				m.Customer1 = &Customer1{}
			}
			if err := m.Customer1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1031:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NonCriticalField", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NonCriticalField = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *TestVersion3LoneOneOfValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: TestVersion3LoneOneOfValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestVersion3LoneOneOfValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			m.X = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.X |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.A == nil {
				m.A = &TestVersion3{}
			}
			if err := m.A.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.B == nil {
				m.B = &TestVersion3{}
			}
			if err := m.B.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.C = append(m.C, &TestVersion3{})
			if err := m.C[len(m.C)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field D", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.D = append(m.D, &TestVersion3{})
			if err := m.D[len(m.D)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field E", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Sum = &TestVersion3LoneOneOfValue_E{v}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field G", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.G == nil {
				m.G = &types.Any{}
			}
			if err := m.G.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field H", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.H = append(m.H, &TestVersion1{})
			if err := m.H[len(m.H)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Customer1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Customer1 == nil {
				m.Customer1 = &Customer1{}
			}
			if err := m.Customer1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1031:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NonCriticalField", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NonCriticalField = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *TestVersion3LoneNesting) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: TestVersion3LoneNesting: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestVersion3LoneNesting: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			m.X = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.X |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.A == nil {
				m.A = &TestVersion3{}
			}
			if err := m.A.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.B == nil {
				m.B = &TestVersion3{}
			}
			if err := m.B.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.C = append(m.C, &TestVersion3{})
			if err := m.C[len(m.C)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field D", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.D = append(m.D, &TestVersion3{})
			if err := m.D[len(m.D)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field F", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &TestVersion3LoneNesting{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &TestVersion3LoneNesting_F{v}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field G", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.G == nil {
				m.G = &types.Any{}
			}
			if err := m.G.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field H", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.H = append(m.H, &TestVersion1{})
			if err := m.H[len(m.H)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Customer1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Customer1 == nil {
				m.Customer1 = &Customer1{}
			}
			if err := m.Customer1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inner1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Inner1 == nil {
				m.Inner1 = &TestVersion3LoneNesting_Inner1{}
			}
			if err := m.Inner1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inner2", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Inner2 == nil {
				m.Inner2 = &TestVersion3LoneNesting_Inner2{}
			}
			if err := m.Inner2.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1031:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NonCriticalField", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NonCriticalField = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *TestVersion3LoneNesting_Inner1) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: Inner1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Inner1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inner", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Inner == nil {
				m.Inner = &TestVersion3LoneNesting_Inner1_InnerInner{}
			}
			if err := m.Inner.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *TestVersion3LoneNesting_Inner1_InnerInner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: InnerInner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InnerInner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field City", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.City = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *TestVersion3LoneNesting_Inner2) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: Inner2: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Inner2: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Country", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Country = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inner", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Inner == nil {
				m.Inner = &TestVersion3LoneNesting_Inner2_InnerInner{}
			}
			if err := m.Inner.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *TestVersion3LoneNesting_Inner2_InnerInner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: InnerInner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InnerInner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field City", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.City = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *TestVersion4LoneNesting) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: TestVersion4LoneNesting: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestVersion4LoneNesting: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			m.X = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.X |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.A == nil {
				m.A = &TestVersion3{}
			}
			if err := m.A.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.B == nil {
				m.B = &TestVersion3{}
			}
			if err := m.B.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.C = append(m.C, &TestVersion3{})
			if err := m.C[len(m.C)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field D", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.D = append(m.D, &TestVersion3{})
			if err := m.D[len(m.D)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field F", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &TestVersion3LoneNesting{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &TestVersion4LoneNesting_F{v}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field G", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.G == nil {
				m.G = &types.Any{}
			}
			if err := m.G.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field H", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.H = append(m.H, &TestVersion1{})
			if err := m.H[len(m.H)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Customer1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Customer1 == nil {
				m.Customer1 = &Customer1{}
			}
			if err := m.Customer1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inner1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Inner1 == nil {
				m.Inner1 = &TestVersion4LoneNesting_Inner1{}
			}
			if err := m.Inner1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inner2", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Inner2 == nil {
				m.Inner2 = &TestVersion4LoneNesting_Inner2{}
			}
			if err := m.Inner2.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1031:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NonCriticalField", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NonCriticalField = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *TestVersion4LoneNesting_Inner1) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: Inner1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Inner1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inner", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Inner == nil {
				m.Inner = &TestVersion4LoneNesting_Inner1_InnerInner{}
			}
			if err := m.Inner.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *TestVersion4LoneNesting_Inner1_InnerInner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: InnerInner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InnerInner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field City", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.City = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *TestVersion4LoneNesting_Inner2) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: Inner2: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Inner2: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Country", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Country = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inner", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Inner == nil {
				m.Inner = &TestVersion4LoneNesting_Inner2_InnerInner{}
			}
			if err := m.Inner.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *TestVersion4LoneNesting_Inner2_InnerInner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: InnerInner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InnerInner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *TestVersionFD1) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: TestVersionFD1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestVersionFD1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			m.X = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.X |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.A == nil {
				m.A = &TestVersion1{}
			}
			if err := m.A.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field E", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Sum = &TestVersionFD1_E{v}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field F", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &TestVersion1{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &TestVersionFD1_F{v}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field G", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.G == nil {
				m.G = &types.Any{}
			}
			if err := m.G.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field H", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.H = append(m.H, &TestVersion1{})
			if err := m.H[len(m.H)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *TestVersionFD1WithExtraAny) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: TestVersionFD1WithExtraAny: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestVersionFD1WithExtraAny: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			m.X = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.X |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.A == nil {
				m.A = &TestVersion1{}
			}
			if err := m.A.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field E", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Sum = &TestVersionFD1WithExtraAny_E{v}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field F", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &TestVersion1{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &TestVersionFD1WithExtraAny_F{v}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field G", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.G == nil {
				m.G = &AnyWithExtra{}
			}
			if err := m.G.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field H", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.H = append(m.H, &TestVersion1{})
			if err := m.H[len(m.H)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *AnyWithExtra) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: AnyWithExtra: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AnyWithExtra: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Any", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Any == nil {
				m.Any = &types.Any{}
			}
			if err := m.Any.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			m.B = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.B |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			m.C = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.C |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *TestUpdatedTxRaw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: TestUpdatedTxRaw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestUpdatedTxRaw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BodyBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
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
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
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
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signatures = append(m.Signatures, make([]byte, postIndex-iNdEx))
			copy(m.Signatures[len(m.Signatures)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewField_5", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewField_5 = append(m.NewField_5[:0], dAtA[iNdEx:postIndex]...)
			if m.NewField_5 == nil {
				m.NewField_5 = []byte{}
			}
			iNdEx = postIndex
		case 1024:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewField_1024", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewField_1024 = append(m.NewField_1024[:0], dAtA[iNdEx:postIndex]...)
			if m.NewField_1024 == nil {
				m.NewField_1024 = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *TestUpdatedTxBody) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: TestUpdatedTxBody: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestUpdatedTxBody: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Messages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
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
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
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
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SomeNewField", wireType)
			}
			m.SomeNewField = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SomeNewField |= uint64(b&0x7F) << shift
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
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExtensionOptions = append(m.ExtensionOptions, &types.Any{})
			if err := m.ExtensionOptions[len(m.ExtensionOptions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1050:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SomeNewFieldNonCriticalField", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SomeNewFieldNonCriticalField = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2047:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NonCriticalExtensionOptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
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
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *TestUpdatedAuthInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: TestUpdatedAuthInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestUpdatedAuthInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignerInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignerInfos = append(m.SignerInfos, &tx.SignerInfo{})
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
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fee == nil {
				m.Fee = &tx.Fee{}
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewField_3", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewField_3 = append(m.NewField_3[:0], dAtA[iNdEx:postIndex]...)
			if m.NewField_3 == nil {
				m.NewField_3 = []byte{}
			}
			iNdEx = postIndex
		case 1024:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewField_1024", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknonwnproto
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
				return ErrInvalidLengthUnknonwnproto
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUnknonwnproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewField_1024 = append(m.NewField_1024[:0], dAtA[iNdEx:postIndex]...)
			if m.NewField_1024 == nil {
				m.NewField_1024 = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func (m *TestRepeatedUints) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknonwnproto
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
			return fmt.Errorf("proto: TestRepeatedUints: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestRepeatedUints: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowUnknonwnproto
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Nums = append(m.Nums, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowUnknonwnproto
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthUnknonwnproto
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthUnknonwnproto
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Nums) == 0 {
					m.Nums = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowUnknonwnproto
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Nums = append(m.Nums, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Nums", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUnknonwnproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknonwnproto
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
func skipUnknonwnproto(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUnknonwnproto
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
					return 0, ErrIntOverflowUnknonwnproto
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
					return 0, ErrIntOverflowUnknonwnproto
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
				return 0, ErrInvalidLengthUnknonwnproto
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUnknonwnproto
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUnknonwnproto
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUnknonwnproto        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUnknonwnproto          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUnknonwnproto = fmt.Errorf("proto: unexpected end of group")
)
