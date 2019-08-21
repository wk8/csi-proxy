// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package main

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ComputeDoubleRequest struct {
	Input32              int32    `protobuf:"varint,1,opt,name=input32,proto3" json:"input32,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeDoubleRequest) Reset()         { *m = ComputeDoubleRequest{} }
func (m *ComputeDoubleRequest) String() string { return proto.CompactTextString(m) }
func (*ComputeDoubleRequest) ProtoMessage()    {}
func (*ComputeDoubleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *ComputeDoubleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeDoubleRequest.Unmarshal(m, b)
}
func (m *ComputeDoubleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeDoubleRequest.Marshal(b, m, deterministic)
}
func (m *ComputeDoubleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeDoubleRequest.Merge(m, src)
}
func (m *ComputeDoubleRequest) XXX_Size() int {
	return xxx_messageInfo_ComputeDoubleRequest.Size(m)
}
func (m *ComputeDoubleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeDoubleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeDoubleRequest proto.InternalMessageInfo

func (m *ComputeDoubleRequest) GetInput32() int32 {
	if m != nil {
		return m.Input32
	}
	return 0
}

type ComputeDoubleResponse struct {
	Response32           int32    `protobuf:"varint,1,opt,name=response32,proto3" json:"response32,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeDoubleResponse) Reset()         { *m = ComputeDoubleResponse{} }
func (m *ComputeDoubleResponse) String() string { return proto.CompactTextString(m) }
func (*ComputeDoubleResponse) ProtoMessage()    {}
func (*ComputeDoubleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *ComputeDoubleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeDoubleResponse.Unmarshal(m, b)
}
func (m *ComputeDoubleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeDoubleResponse.Marshal(b, m, deterministic)
}
func (m *ComputeDoubleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeDoubleResponse.Merge(m, src)
}
func (m *ComputeDoubleResponse) XXX_Size() int {
	return xxx_messageInfo_ComputeDoubleResponse.Size(m)
}
func (m *ComputeDoubleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeDoubleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeDoubleResponse proto.InternalMessageInfo

func (m *ComputeDoubleResponse) GetResponse32() int32 {
	if m != nil {
		return m.Response32
	}
	return 0
}

func init() {
	proto.RegisterType((*ComputeDoubleRequest)(nil), "v0alpha1.ComputeDoubleRequest")
	proto.RegisterType((*ComputeDoubleResponse)(nil), "v0alpha1.ComputeDoubleResponse")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x33, 0x48, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0x54,
	0x32, 0xe0, 0x12, 0x71, 0xce, 0xcf, 0x2d, 0x28, 0x2d, 0x49, 0x75, 0xc9, 0x2f, 0x4d, 0xca, 0x49,
	0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe0, 0x62, 0xcf, 0xcc, 0x2b, 0x28, 0x2d,
	0x31, 0x36, 0x92, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0x71, 0x95, 0xcc, 0xb9, 0x44, 0xd1,
	0x74, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0xc9, 0x71, 0x71, 0x15, 0x41, 0xd9, 0x70, 0x5d,
	0x48, 0x22, 0x46, 0x99, 0x5c, 0xc2, 0x21, 0xa9, 0xc5, 0x25, 0xce, 0xc1, 0x9e, 0x01, 0x45, 0xf9,
	0x15, 0x95, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x41, 0x5c, 0xbc, 0x28, 0xe6, 0x09,
	0xc9, 0xe9, 0xc1, 0x5c, 0xa7, 0x87, 0xcd, 0x69, 0x52, 0xf2, 0x38, 0xe5, 0x21, 0x56, 0x29, 0x31,
	0x24, 0xb1, 0x81, 0xbd, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x85, 0xae, 0x36, 0xfb, 0xf3,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestCSIProxyServiceClient is the client API for TestCSIProxyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestCSIProxyServiceClient interface {
	// ComputeDouble computes the double of the input. Real smart stuff!
	ComputeDouble(ctx context.Context, in *ComputeDoubleRequest, opts ...grpc.CallOption) (*ComputeDoubleResponse, error)
}

type testCSIProxyServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestCSIProxyServiceClient(cc *grpc.ClientConn) TestCSIProxyServiceClient {
	return &testCSIProxyServiceClient{cc}
}

func (c *testCSIProxyServiceClient) ComputeDouble(ctx context.Context, in *ComputeDoubleRequest, opts ...grpc.CallOption) (*ComputeDoubleResponse, error) {
	out := new(ComputeDoubleResponse)
	err := c.cc.Invoke(ctx, "/v0alpha1.TestCSIProxyService/ComputeDouble", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestCSIProxyServiceServer is the server API for TestCSIProxyService service.
type TestCSIProxyServiceServer interface {
	// ComputeDouble computes the double of the input. Real smart stuff!
	ComputeDouble(context.Context, *ComputeDoubleRequest) (*ComputeDoubleResponse, error)
}

// UnimplementedTestCSIProxyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTestCSIProxyServiceServer struct {
}

func (*UnimplementedTestCSIProxyServiceServer) ComputeDouble(ctx context.Context, req *ComputeDoubleRequest) (*ComputeDoubleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeDouble not implemented")
}

func RegisterTestCSIProxyServiceServer(s *grpc.Server, srv TestCSIProxyServiceServer) {
	s.RegisterService(&_TestCSIProxyService_serviceDesc, srv)
}

func _TestCSIProxyService_ComputeDouble_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComputeDoubleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestCSIProxyServiceServer).ComputeDouble(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v0alpha1.TestCSIProxyService/ComputeDouble",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestCSIProxyServiceServer).ComputeDouble(ctx, req.(*ComputeDoubleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestCSIProxyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v0alpha1.TestCSIProxyService",
	HandlerType: (*TestCSIProxyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ComputeDouble",
			Handler:    _TestCSIProxyService_ComputeDouble_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
