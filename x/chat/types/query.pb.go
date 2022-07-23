


package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)


var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf





const _ = proto.GoGoProtoPackageIsVersion3



type TestQueryRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestQueryRequest) Reset()         { *m = TestQueryRequest{} }
func (m *TestQueryRequest) String() string { return proto.CompactTextString(m) }
func (*TestQueryRequest) ProtoMessage()    {}
func (*TestQueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{0}
}
func (m *TestQueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestQueryRequest.Unmarshal(m, b)
}
func (m *TestQueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestQueryRequest.Marshal(b, m, deterministic)
}
func (m *TestQueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestQueryRequest.Merge(m, src)
}
func (m *TestQueryRequest) XXX_Size() int {
	return xxx_messageInfo_TestQueryRequest.Size(m)
}
func (m *TestQueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestQueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestQueryRequest proto.InternalMessageInfo



type TestQueryResponse struct {
	TestR                string   `protobuf:"bytes,1,opt,name=testR,proto3" json:"testR,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestQueryResponse) Reset()         { *m = TestQueryResponse{} }
func (m *TestQueryResponse) String() string { return proto.CompactTextString(m) }
func (*TestQueryResponse) ProtoMessage()    {}
func (*TestQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{1}
}
func (m *TestQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestQueryResponse.Unmarshal(m, b)
}
func (m *TestQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestQueryResponse.Marshal(b, m, deterministic)
}
func (m *TestQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestQueryResponse.Merge(m, src)
}
func (m *TestQueryResponse) XXX_Size() int {
	return xxx_messageInfo_TestQueryResponse.Size(m)
}
func (m *TestQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestQueryResponse proto.InternalMessageInfo

func (m *TestQueryResponse) GetTestR() string {
	if m != nil {
		return m.TestR
	}
	return ""
}

func init() {
	proto.RegisterType((*TestQueryRequest)(nil), "freemasonry.chat.v1.TestQueryRequest")
	proto.RegisterType((*TestQueryResponse)(nil), "freemasonry.chat.v1.TestQueryResponse")
}

func init() { proto.RegisterFile("query.proto", fileDescriptor_5c6ac9b241082464) }

var fileDescriptor_5c6ac9b241082464 = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xc1, 0x4a, 0xc6, 0x30,
	0x10, 0x84, 0xed, 0xa1, 0x42, 0xe3, 0x45, 0x63, 0x0f, 0x52, 0x3c, 0x48, 0xc1, 0x52, 0x2f, 0x09,
	0xea, 0x1b, 0xf8, 0x06, 0x16, 0x4f, 0xe2, 0x25, 0x0d, 0x6b, 0x5a, 0xac, 0xd9, 0x34, 0xbb, 0x15,
	0xfb, 0xf6, 0xd2, 0x06, 0x44, 0x45, 0xf8, 0x6f, 0xb3, 0xc3, 0x7c, 0x3b, 0xbb, 0xe2, 0x64, 0x5e,
	0x20, 0xae, 0x2a, 0x44, 0x64, 0x94, 0xe7, 0xaf, 0x11, 0xe0, 0xdd, 0x10, 0xfa, 0xb8, 0x2a, 0x3b,
	0x18, 0x56, 0x1f, 0xb7, 0xd5, 0xa5, 0x43, 0x74, 0x13, 0x68, 0x13, 0x46, 0x6d, 0xbc, 0x47, 0x36,
	0x3c, 0xa2, 0xa7, 0x84, 0x54, 0xa5, 0x43, 0x87, 0xbb, 0xd4, 0x9b, 0x4a, 0x6e, 0x2d, 0xc5, 0xe9,
	0x13, 0x10, 0x3f, 0x6e, 0xbb, 0x3b, 0x98, 0x17, 0x20, 0xae, 0x6f, 0xc4, 0xd9, 0x0f, 0x8f, 0x02,
	0x7a, 0x02, 0x59, 0x8a, 0x9c, 0x81, 0xb8, 0xbb, 0xc8, 0xae, 0xb2, 0xb6, 0xe8, 0xd2, 0x70, 0x07,
	0x22, 0xdf, 0x63, 0xf2, 0x45, 0x14, 0xdf, 0x8c, 0xbc, 0x56, 0xff, 0x9c, 0xa7, 0xfe, 0xf6, 0x54,
	0xcd, 0xa1, 0x58, 0xaa, 0xae, 0x8f, 0x1e, 0xda, 0xe7, 0xe6, 0x57, 0xd4, 0xea, 0x7e, 0x42, 0xfb,
	0x66, 0x07, 0x33, 0x7a, 0xfd, 0xa9, 0x37, 0x54, 0xf3, 0x1a, 0x80, 0xfa, 0xe3, 0xfd, 0xad, 0xfb,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x5d, 0xa1, 0x03, 0x2e, 0x01, 0x00, 0x00,
}


var _ context.Context
var _ grpc.ClientConn



const _ = grpc.SupportPackageIsVersion4


//

type QueryClient interface {

	TestQuery(ctx context.Context, in *TestQueryRequest, opts ...grpc.CallOption) (*TestQueryResponse, error)
}

type queryClient struct {
	cc *grpc.ClientConn
}

func NewQueryClient(cc *grpc.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) TestQuery(ctx context.Context, in *TestQueryRequest, opts ...grpc.CallOption) (*TestQueryResponse, error) {
	out := new(TestQueryResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.chat.v1.Query/TestQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


type QueryServer interface {

	TestQuery(context.Context, *TestQueryRequest) (*TestQueryResponse, error)
}


type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) TestQuery(ctx context.Context, req *TestQueryRequest) (*TestQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestQuery not implemented")
}

func RegisterQueryServer(s *grpc.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_TestQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TestQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freemasonry.chat.v1.Query/TestQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TestQuery(ctx, req.(*TestQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "freemasonry.chat.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestQuery",
			Handler:    _Query_TestQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "query.proto",
}
