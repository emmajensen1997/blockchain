


package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
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



type QueryCurrentPlanRequest struct {
}

func (m *QueryCurrentPlanRequest) Reset()         { *m = QueryCurrentPlanRequest{} }
func (m *QueryCurrentPlanRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCurrentPlanRequest) ProtoMessage()    {}
func (*QueryCurrentPlanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a334d07ad8374f0, []int{0}
}
func (m *QueryCurrentPlanRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCurrentPlanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCurrentPlanRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCurrentPlanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCurrentPlanRequest.Merge(m, src)
}
func (m *QueryCurrentPlanRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCurrentPlanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCurrentPlanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCurrentPlanRequest proto.InternalMessageInfo



type QueryCurrentPlanResponse struct {

	Plan *Plan `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
}

func (m *QueryCurrentPlanResponse) Reset()         { *m = QueryCurrentPlanResponse{} }
func (m *QueryCurrentPlanResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCurrentPlanResponse) ProtoMessage()    {}
func (*QueryCurrentPlanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a334d07ad8374f0, []int{1}
}
func (m *QueryCurrentPlanResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCurrentPlanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCurrentPlanResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCurrentPlanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCurrentPlanResponse.Merge(m, src)
}
func (m *QueryCurrentPlanResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCurrentPlanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCurrentPlanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCurrentPlanResponse proto.InternalMessageInfo

func (m *QueryCurrentPlanResponse) GetPlan() *Plan {
	if m != nil {
		return m.Plan
	}
	return nil
}



type QueryAppliedPlanRequest struct {

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *QueryAppliedPlanRequest) Reset()         { *m = QueryAppliedPlanRequest{} }
func (m *QueryAppliedPlanRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAppliedPlanRequest) ProtoMessage()    {}
func (*QueryAppliedPlanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a334d07ad8374f0, []int{2}
}
func (m *QueryAppliedPlanRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAppliedPlanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAppliedPlanRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAppliedPlanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAppliedPlanRequest.Merge(m, src)
}
func (m *QueryAppliedPlanRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAppliedPlanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAppliedPlanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAppliedPlanRequest proto.InternalMessageInfo

func (m *QueryAppliedPlanRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}



type QueryAppliedPlanResponse struct {

	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *QueryAppliedPlanResponse) Reset()         { *m = QueryAppliedPlanResponse{} }
func (m *QueryAppliedPlanResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAppliedPlanResponse) ProtoMessage()    {}
func (*QueryAppliedPlanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a334d07ad8374f0, []int{3}
}
func (m *QueryAppliedPlanResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAppliedPlanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAppliedPlanResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAppliedPlanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAppliedPlanResponse.Merge(m, src)
}
func (m *QueryAppliedPlanResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAppliedPlanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAppliedPlanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAppliedPlanResponse proto.InternalMessageInfo

func (m *QueryAppliedPlanResponse) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}



//

type QueryUpgradedConsensusStateRequest struct {


	LastHeight int64 `protobuf:"varint,1,opt,name=last_height,json=lastHeight,proto3" json:"last_height,omitempty"`
}

func (m *QueryUpgradedConsensusStateRequest) Reset()         { *m = QueryUpgradedConsensusStateRequest{} }
func (m *QueryUpgradedConsensusStateRequest) String() string { return proto.CompactTextString(m) }
func (*QueryUpgradedConsensusStateRequest) ProtoMessage()    {}
func (*QueryUpgradedConsensusStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a334d07ad8374f0, []int{4}
}
func (m *QueryUpgradedConsensusStateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUpgradedConsensusStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUpgradedConsensusStateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUpgradedConsensusStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUpgradedConsensusStateRequest.Merge(m, src)
}
func (m *QueryUpgradedConsensusStateRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryUpgradedConsensusStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUpgradedConsensusStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUpgradedConsensusStateRequest proto.InternalMessageInfo

func (m *QueryUpgradedConsensusStateRequest) GetLastHeight() int64 {
	if m != nil {
		return m.LastHeight
	}
	return 0
}



//

type QueryUpgradedConsensusStateResponse struct {

	UpgradedConsensusState []byte `protobuf:"bytes,2,opt,name=upgraded_consensus_state,json=upgradedConsensusState,proto3" json:"upgraded_consensus_state,omitempty"`
}

func (m *QueryUpgradedConsensusStateResponse) Reset()         { *m = QueryUpgradedConsensusStateResponse{} }
func (m *QueryUpgradedConsensusStateResponse) String() string { return proto.CompactTextString(m) }
func (*QueryUpgradedConsensusStateResponse) ProtoMessage()    {}
func (*QueryUpgradedConsensusStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a334d07ad8374f0, []int{5}
}
func (m *QueryUpgradedConsensusStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUpgradedConsensusStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUpgradedConsensusStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUpgradedConsensusStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUpgradedConsensusStateResponse.Merge(m, src)
}
func (m *QueryUpgradedConsensusStateResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryUpgradedConsensusStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUpgradedConsensusStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUpgradedConsensusStateResponse proto.InternalMessageInfo

func (m *QueryUpgradedConsensusStateResponse) GetUpgradedConsensusState() []byte {
	if m != nil {
		return m.UpgradedConsensusState
	}
	return nil
}



//

type QueryModuleVersionsRequest struct {



	ModuleName string `protobuf:"bytes,1,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
}

func (m *QueryModuleVersionsRequest) Reset()         { *m = QueryModuleVersionsRequest{} }
func (m *QueryModuleVersionsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryModuleVersionsRequest) ProtoMessage()    {}
func (*QueryModuleVersionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a334d07ad8374f0, []int{6}
}
func (m *QueryModuleVersionsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryModuleVersionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryModuleVersionsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryModuleVersionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryModuleVersionsRequest.Merge(m, src)
}
func (m *QueryModuleVersionsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryModuleVersionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryModuleVersionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryModuleVersionsRequest proto.InternalMessageInfo

func (m *QueryModuleVersionsRequest) GetModuleName() string {
	if m != nil {
		return m.ModuleName
	}
	return ""
}



//

type QueryModuleVersionsResponse struct {

	ModuleVersions []*ModuleVersion `protobuf:"bytes,1,rep,name=module_versions,json=moduleVersions,proto3" json:"module_versions,omitempty"`
}

func (m *QueryModuleVersionsResponse) Reset()         { *m = QueryModuleVersionsResponse{} }
func (m *QueryModuleVersionsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryModuleVersionsResponse) ProtoMessage()    {}
func (*QueryModuleVersionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a334d07ad8374f0, []int{7}
}
func (m *QueryModuleVersionsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryModuleVersionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryModuleVersionsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryModuleVersionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryModuleVersionsResponse.Merge(m, src)
}
func (m *QueryModuleVersionsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryModuleVersionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryModuleVersionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryModuleVersionsResponse proto.InternalMessageInfo

func (m *QueryModuleVersionsResponse) GetModuleVersions() []*ModuleVersion {
	if m != nil {
		return m.ModuleVersions
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryCurrentPlanRequest)(nil), "cosmos.upgrade.v1beta1.QueryCurrentPlanRequest")
	proto.RegisterType((*QueryCurrentPlanResponse)(nil), "cosmos.upgrade.v1beta1.QueryCurrentPlanResponse")
	proto.RegisterType((*QueryAppliedPlanRequest)(nil), "cosmos.upgrade.v1beta1.QueryAppliedPlanRequest")
	proto.RegisterType((*QueryAppliedPlanResponse)(nil), "cosmos.upgrade.v1beta1.QueryAppliedPlanResponse")
	proto.RegisterType((*QueryUpgradedConsensusStateRequest)(nil), "cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest")
	proto.RegisterType((*QueryUpgradedConsensusStateResponse)(nil), "cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse")
	proto.RegisterType((*QueryModuleVersionsRequest)(nil), "cosmos.upgrade.v1beta1.QueryModuleVersionsRequest")
	proto.RegisterType((*QueryModuleVersionsResponse)(nil), "cosmos.upgrade.v1beta1.QueryModuleVersionsResponse")
}

func init() {
	proto.RegisterFile("cosmos/upgrade/v1beta1/query.proto", fileDescriptor_4a334d07ad8374f0)
}

var fileDescriptor_4a334d07ad8374f0 = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x6f, 0xd3, 0x3c,
	0x1c, 0xae, 0xbb, 0x6e, 0x7a, 0x5f, 0x17, 0x0d, 0xe4, 0x43, 0xc9, 0xc2, 0x14, 0x2a, 0x33, 0xa0,
	0x88, 0x35, 0xde, 0xda, 0x0b, 0x1a, 0x02, 0x01, 0x93, 0x10, 0x43, 0x30, 0x41, 0x11, 0x1c, 0xb8,
	0x54, 0x6e, 0x63, 0xda, 0x88, 0x24, 0xce, 0x62, 0x67, 0x62, 0x9a, 0x76, 0xe1, 0xc4, 0x11, 0x89,
	0x3b, 0x37, 0x2e, 0x7c, 0x12, 0x8e, 0x93, 0xb8, 0x70, 0xd8, 0x01, 0xb5, 0x7c, 0x10, 0x14, 0xc7,
	0x45, 0x29, 0x4d, 0x3a, 0xe0, 0xd4, 0xc6, 0x7e, 0xfe, 0xfd, 0xe2, 0xc7, 0x81, 0xb8, 0xcf, 0x85,
	0xcf, 0x05, 0x89, 0xc3, 0x41, 0x44, 0x1d, 0x46, 0xf6, 0x37, 0x7b, 0x4c, 0xd2, 0x4d, 0xb2, 0x17,
	0xb3, 0xe8, 0xc0, 0x0e, 0x23, 0x2e, 0x39, 0xaa, 0xa5, 0x18, 0x5b, 0x63, 0x6c, 0x8d, 0x31, 0x57,
	0x06, 0x9c, 0x0f, 0x3c, 0x46, 0x14, 0xaa, 0x17, 0xbf, 0x22, 0x34, 0xd0, 0x14, 0x73, 0x55, 0x6f,
	0xd1, 0xd0, 0x25, 0x34, 0x08, 0xb8, 0xa4, 0xd2, 0xe5, 0x81, 0xd0, 0xbb, 0x6b, 0x05, 0xa6, 0x13,
	0x03, 0x85, 0xc2, 0x2b, 0xf0, 0xfc, 0xd3, 0x24, 0xc5, 0x76, 0x1c, 0x45, 0x2c, 0x90, 0x4f, 0x3c,
	0x1a, 0x74, 0xd8, 0x5e, 0xcc, 0x84, 0xc4, 0x8f, 0xa0, 0x31, 0xbb, 0x25, 0x42, 0x1e, 0x08, 0x86,
	0x36, 0x60, 0x25, 0xf4, 0x68, 0x60, 0x80, 0x3a, 0x68, 0x54, 0x5b, 0xab, 0x76, 0x7e, 0x78, 0x5b,
	0x71, 0x14, 0x12, 0x37, 0xb5, 0xd1, 0xdd, 0x30, 0xf4, 0x5c, 0xe6, 0x64, 0x8c, 0x10, 0x82, 0x95,
	0x80, 0xfa, 0x4c, 0x89, 0xfd, 0xdf, 0x51, 0xff, 0x71, 0x4b, 0x9b, 0x4f, 0xc1, 0xb5, 0x79, 0x0d,
	0x2e, 0x0d, 0x99, 0x3b, 0x18, 0x4a, 0xc5, 0x58, 0xe8, 0xe8, 0x27, 0xbc, 0x03, 0xb1, 0xe2, 0x3c,
	0x4f, 0x53, 0x38, 0xdb, 0x09, 0x3a, 0x10, 0xb1, 0x78, 0x26, 0xa9, 0x64, 0x13, 0xb7, 0x8b, 0xb0,
	0xea, 0x51, 0x21, 0xbb, 0x53, 0x12, 0x30, 0x59, 0x7a, 0xa0, 0x56, 0xb6, 0xca, 0x06, 0xc0, 0x2e,
	0xbc, 0x34, 0x57, 0x4a, 0x27, 0xb9, 0x01, 0x0d, 0x3d, 0xb2, 0xd3, 0xed, 0x4f, 0x20, 0x5d, 0x91,
	0x60, 0x8c, 0x72, 0x1d, 0x34, 0xce, 0x74, 0x6a, 0x71, 0xae, 0x42, 0x62, 0xf2, 0xb0, 0xf2, 0x1f,
	0x38, 0x57, 0xc6, 0xb7, 0xa0, 0xa9, 0xac, 0x1e, 0x73, 0x27, 0xf6, 0xd8, 0x0b, 0x16, 0x89, 0xe4,
	0x10, 0x33, 0x69, 0x7d, 0xb5, 0xd1, 0xcd, 0xbc, 0x22, 0x98, 0x2e, 0xed, 0x26, 0x2f, 0xca, 0x87,
	0x17, 0x72, 0xe9, 0x3a, 0xe1, 0x2e, 0x3c, 0xab, 0xf9, 0xfb, 0x7a, 0xcb, 0x00, 0xf5, 0x85, 0x46,
	0xb5, 0x75, 0xb9, 0xe8, 0xcc, 0xa6, 0x84, 0x3a, 0xcb, 0xfe, 0x94, 0x6e, 0xeb, 0x64, 0x11, 0x2e,
	0x2a, 0x3f, 0xf4, 0x11, 0xc0, 0x6a, 0xa6, 0x1a, 0x88, 0x14, 0x09, 0x16, 0xf4, 0xcb, 0xdc, 0xf8,
	0x73, 0x42, 0x3a, 0x0c, 0x5e, 0x7f, 0xfb, 0xf5, 0xc7, 0x87, 0xf2, 0x15, 0xb4, 0x46, 0x0a, 0xba,
	0xdd, 0x4f, 0x49, 0xdd, 0xa4, 0x71, 0xe8, 0x13, 0x80, 0xd5, 0x4c, 0x7d, 0x4e, 0x09, 0x38, 0xdb,
	0xcb, 0x53, 0x02, 0xe6, 0x34, 0x13, 0xb7, 0x55, 0xc0, 0x26, 0xba, 0x5e, 0x14, 0x90, 0xa6, 0x24,
	0x15, 0x90, 0x1c, 0x26, 0x47, 0x7a, 0x84, 0x4e, 0x00, 0xac, 0xe5, 0xf7, 0x0c, 0x6d, 0xcd, 0x4d,
	0x30, 0xb7, 0xe7, 0xe6, 0xcd, 0x7f, 0xe2, 0xea, 0x41, 0x76, 0xd4, 0x20, 0x77, 0xd0, 0x6d, 0x32,
	0xff, 0x2b, 0x32, 0x53, 0x7b, 0x72, 0x98, 0xb9, 0x5c, 0x47, 0xef, 0xca, 0x00, 0x7d, 0x06, 0x70,
	0x79, 0xba, 0x9c, 0xa8, 0x35, 0x37, 0x5a, 0xee, 0x45, 0x30, 0xdb, 0x7f, 0xc5, 0xd1, 0x63, 0x10,
	0x35, 0xc6, 0x35, 0x74, 0xb5, 0x68, 0x8c, 0xdf, 0xee, 0xc6, 0xbd, 0xfb, 0x5f, 0x46, 0x16, 0x38,
	0x1e, 0x59, 0xe0, 0xfb, 0xc8, 0x02, 0xef, 0xc7, 0x56, 0xe9, 0x78, 0x6c, 0x95, 0xbe, 0x8d, 0xad,
	0xd2, 0xcb, 0xf5, 0x81, 0x2b, 0x87, 0x71, 0xcf, 0xee, 0x73, 0x7f, 0x22, 0x96, 0xfe, 0x34, 0x85,
	0xf3, 0x9a, 0xbc, 0xf9, 0xa5, 0x2c, 0x0f, 0x42, 0x26, 0x7a, 0x4b, 0xea, 0xeb, 0xda, 0xfe, 0x19,
	0x00, 0x00, 0xff, 0xff, 0x80, 0xcd, 0x5d, 0xc6, 0xfa, 0x05, 0x00, 0x00,
}


var _ context.Context
var _ grpc.ClientConn



const _ = grpc.SupportPackageIsVersion4


//

type QueryClient interface {

	CurrentPlan(ctx context.Context, in *QueryCurrentPlanRequest, opts ...grpc.CallOption) (*QueryCurrentPlanResponse, error)

	AppliedPlan(ctx context.Context, in *QueryAppliedPlanRequest, opts ...grpc.CallOption) (*QueryAppliedPlanResponse, error)






	UpgradedConsensusState(ctx context.Context, in *QueryUpgradedConsensusStateRequest, opts ...grpc.CallOption) (*QueryUpgradedConsensusStateResponse, error)

	//

	ModuleVersions(ctx context.Context, in *QueryModuleVersionsRequest, opts ...grpc.CallOption) (*QueryModuleVersionsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) CurrentPlan(ctx context.Context, in *QueryCurrentPlanRequest, opts ...grpc.CallOption) (*QueryCurrentPlanResponse, error) {
	out := new(QueryCurrentPlanResponse)
	err := c.cc.Invoke(ctx, "/cosmos.upgrade.v1beta1.Query/CurrentPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AppliedPlan(ctx context.Context, in *QueryAppliedPlanRequest, opts ...grpc.CallOption) (*QueryAppliedPlanResponse, error) {
	out := new(QueryAppliedPlanResponse)
	err := c.cc.Invoke(ctx, "/cosmos.upgrade.v1beta1.Query/AppliedPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


func (c *queryClient) UpgradedConsensusState(ctx context.Context, in *QueryUpgradedConsensusStateRequest, opts ...grpc.CallOption) (*QueryUpgradedConsensusStateResponse, error) {
	out := new(QueryUpgradedConsensusStateResponse)
	err := c.cc.Invoke(ctx, "/cosmos.upgrade.v1beta1.Query/UpgradedConsensusState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ModuleVersions(ctx context.Context, in *QueryModuleVersionsRequest, opts ...grpc.CallOption) (*QueryModuleVersionsResponse, error) {
	out := new(QueryModuleVersionsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.upgrade.v1beta1.Query/ModuleVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


type QueryServer interface {

	CurrentPlan(context.Context, *QueryCurrentPlanRequest) (*QueryCurrentPlanResponse, error)

	AppliedPlan(context.Context, *QueryAppliedPlanRequest) (*QueryAppliedPlanResponse, error)






	UpgradedConsensusState(context.Context, *QueryUpgradedConsensusStateRequest) (*QueryUpgradedConsensusStateResponse, error)

	//

	ModuleVersions(context.Context, *QueryModuleVersionsRequest) (*QueryModuleVersionsResponse, error)
}


type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) CurrentPlan(ctx context.Context, req *QueryCurrentPlanRequest) (*QueryCurrentPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentPlan not implemented")
}
func (*UnimplementedQueryServer) AppliedPlan(ctx context.Context, req *QueryAppliedPlanRequest) (*QueryAppliedPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppliedPlan not implemented")
}
func (*UnimplementedQueryServer) UpgradedConsensusState(ctx context.Context, req *QueryUpgradedConsensusStateRequest) (*QueryUpgradedConsensusStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradedConsensusState not implemented")
}
func (*UnimplementedQueryServer) ModuleVersions(ctx context.Context, req *QueryModuleVersionsRequest) (*QueryModuleVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModuleVersions not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_CurrentPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCurrentPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CurrentPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.upgrade.v1beta1.Query/CurrentPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CurrentPlan(ctx, req.(*QueryCurrentPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AppliedPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAppliedPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AppliedPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.upgrade.v1beta1.Query/AppliedPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AppliedPlan(ctx, req.(*QueryAppliedPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_UpgradedConsensusState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUpgradedConsensusStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).UpgradedConsensusState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.upgrade.v1beta1.Query/UpgradedConsensusState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).UpgradedConsensusState(ctx, req.(*QueryUpgradedConsensusStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ModuleVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryModuleVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ModuleVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.upgrade.v1beta1.Query/ModuleVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ModuleVersions(ctx, req.(*QueryModuleVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.upgrade.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CurrentPlan",
			Handler:    _Query_CurrentPlan_Handler,
		},
		{
			MethodName: "AppliedPlan",
			Handler:    _Query_AppliedPlan_Handler,
		},
		{
			MethodName: "UpgradedConsensusState",
			Handler:    _Query_UpgradedConsensusState_Handler,
		},
		{
			MethodName: "ModuleVersions",
			Handler:    _Query_ModuleVersions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/upgrade/v1beta1/query.proto",
}

func (m *QueryCurrentPlanRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCurrentPlanRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCurrentPlanRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryCurrentPlanResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCurrentPlanResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCurrentPlanResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Plan != nil {
		{
			size, err := m.Plan.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAppliedPlanRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAppliedPlanRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAppliedPlanRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAppliedPlanResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAppliedPlanResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAppliedPlanResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryUpgradedConsensusStateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUpgradedConsensusStateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUpgradedConsensusStateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.LastHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryUpgradedConsensusStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUpgradedConsensusStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUpgradedConsensusStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UpgradedConsensusState) > 0 {
		i -= len(m.UpgradedConsensusState)
		copy(dAtA[i:], m.UpgradedConsensusState)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.UpgradedConsensusState)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *QueryModuleVersionsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryModuleVersionsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryModuleVersionsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ModuleName) > 0 {
		i -= len(m.ModuleName)
		copy(dAtA[i:], m.ModuleName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ModuleName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryModuleVersionsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryModuleVersionsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryModuleVersionsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ModuleVersions) > 0 {
		for iNdEx := len(m.ModuleVersions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ModuleVersions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryCurrentPlanRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryCurrentPlanResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Plan != nil {
		l = m.Plan.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAppliedPlanRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAppliedPlanResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovQuery(uint64(m.Height))
	}
	return n
}

func (m *QueryUpgradedConsensusStateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LastHeight != 0 {
		n += 1 + sovQuery(uint64(m.LastHeight))
	}
	return n
}

func (m *QueryUpgradedConsensusStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UpgradedConsensusState)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryModuleVersionsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ModuleName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryModuleVersionsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ModuleVersions) > 0 {
		for _, e := range m.ModuleVersions {
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
func (m *QueryCurrentPlanRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCurrentPlanRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCurrentPlanRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryCurrentPlanResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCurrentPlanResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCurrentPlanResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plan", wireType)
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
			if m.Plan == nil {
				m.Plan = &Plan{}
			}
			if err := m.Plan.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAppliedPlanRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAppliedPlanRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAppliedPlanRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
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
func (m *QueryAppliedPlanResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAppliedPlanResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAppliedPlanResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
func (m *QueryUpgradedConsensusStateRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryUpgradedConsensusStateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUpgradedConsensusStateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastHeight", wireType)
			}
			m.LastHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastHeight |= int64(b&0x7F) << shift
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
func (m *QueryUpgradedConsensusStateResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryUpgradedConsensusStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUpgradedConsensusStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpgradedConsensusState", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpgradedConsensusState = append(m.UpgradedConsensusState[:0], dAtA[iNdEx:postIndex]...)
			if m.UpgradedConsensusState == nil {
				m.UpgradedConsensusState = []byte{}
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
func (m *QueryModuleVersionsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryModuleVersionsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryModuleVersionsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleName", wireType)
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
			m.ModuleName = string(dAtA[iNdEx:postIndex])
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
func (m *QueryModuleVersionsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryModuleVersionsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryModuleVersionsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleVersions", wireType)
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
			m.ModuleVersions = append(m.ModuleVersions, &ModuleVersion{})
			if err := m.ModuleVersions[len(m.ModuleVersions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
