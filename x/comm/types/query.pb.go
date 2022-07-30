


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

func init() {
	proto.RegisterType((*TestQueryRequest)(nil), "freemasonry.comm.v1.TestQueryRequest")
	proto.RegisterType((*TestQueryResponse)(nil), "freemasonry.comm.v1.TestQueryResponse")
}

func init() { proto.RegisterFile("query.proto", fileDescriptor_5c6ac9b241082464) }

var fileDescriptor_5c6ac9b241082464 = []byte{
	
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0x31, 0x0b, 0xc2, 0x30,
	0x10, 0x85, 0x5d, 0x14, 0x8c, 0x8b, 0xa6, 0x4e, 0xc5, 0xa9, 0x60, 0x71, 0x4a, 0x50, 0xff, 0x81,
	0xff, 0x40, 0x71, 0x12, 0x97, 0x34, 0x9c, 0xb1, 0xd8, 0xe6, 0xd2, 0x5c, 0x2a, 0xf6, 0xdf, 0x4b,
	0x5b, 0x10, 0x15, 0xc1, 0xed, 0xf1, 0xdd, 0x7b, 0xef, 0xee, 0xd8, 0xa4, 0xaa, 0xc1, 0x37, 0xc2,
	0x79, 0x0c, 0xc8, 0xa3, 0x8b, 0x07, 0x28, 0x15, 0xa1, 0xf5, 0x8d, 0xd0, 0x58, 0x96, 0xe2, 0xbe,
	0x8e, 0x17, 0x06, 0xd1, 0x14, 0x20, 0x95, 0xcb, 0xa5, 0xb2, 0x16, 0x83, 0x0a, 0x39, 0x5a, 0xea,
	0x23, 0xf1, 0xdc, 0xa0, 0xc1, 0x4e, 0xca, 0x56, 0xf5, 0x34, 0xe1, 0x6c, 0x7a, 0x04, 0x0a, 0xfb,
	0xb6, 0xfb, 0x00, 0x55, 0x0d, 0x14, 0x92, 0x88, 0xcd, 0xde, 0x18, 0x39, 0xb4, 0x04, 0x1b, 0x60,
	0xc3, 0x0e, 0xf0, 0x33, 0x1b, 0xbf, 0xa6, 0x7c, 0x29, 0x7e, 0x1c, 0x22, 0xbe, 0x1b, 0xe3, 0xf4,
	0x9f, 0xad, 0x5f, 0x92, 0x0c, 0x76, 0xab, 0x53, 0xfa, 0x61, 0xd5, 0x32, 0x2b, 0x50, 0xdf, 0xf4,
	0x55, 0xe5, 0x56, 0x3e, 0x64, 0x1b, 0x95, 0xa1, 0x71, 0x40, 0xd9, 0xa8, 0x7b, 0x60, 0xfb, 0x0c,
	0x00, 0x00, 0xff, 0xff, 0xce, 0x5b, 0x7d, 0x23, 0x18, 0x01, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/freemasonry.comm.v1.Query/TestQuery", in, out, opts...)
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
		FullMethod: "/freemasonry.comm.v1.Query/TestQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TestQuery(ctx, req.(*TestQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "freemasonry.comm.v1.Query",
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
