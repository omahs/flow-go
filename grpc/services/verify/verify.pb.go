// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/verify/verify.proto

package verify

import (
	fmt "fmt"
	shared "github.com/dapperlabs/bamboo-node/grpc/shared"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12065c2e72b67281, []int{0}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

type PingResponse struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12065c2e72b67281, []int{1}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

type SubmitExecutionReceiptRequest struct {
	ExecutionReceipt     *shared.ExecutionReceipt `protobuf:"bytes,1,opt,name=executionReceipt,proto3" json:"executionReceipt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SubmitExecutionReceiptRequest) Reset()         { *m = SubmitExecutionReceiptRequest{} }
func (m *SubmitExecutionReceiptRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitExecutionReceiptRequest) ProtoMessage()    {}
func (*SubmitExecutionReceiptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12065c2e72b67281, []int{2}
}

func (m *SubmitExecutionReceiptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitExecutionReceiptRequest.Unmarshal(m, b)
}
func (m *SubmitExecutionReceiptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitExecutionReceiptRequest.Marshal(b, m, deterministic)
}
func (m *SubmitExecutionReceiptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitExecutionReceiptRequest.Merge(m, src)
}
func (m *SubmitExecutionReceiptRequest) XXX_Size() int {
	return xxx_messageInfo_SubmitExecutionReceiptRequest.Size(m)
}
func (m *SubmitExecutionReceiptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitExecutionReceiptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitExecutionReceiptRequest proto.InternalMessageInfo

func (m *SubmitExecutionReceiptRequest) GetExecutionReceipt() *shared.ExecutionReceipt {
	if m != nil {
		return m.ExecutionReceipt
	}
	return nil
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "bamboo.services.verify.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "bamboo.services.verify.PingResponse")
	proto.RegisterType((*SubmitExecutionReceiptRequest)(nil), "bamboo.services.verify.SubmitExecutionReceiptRequest")
}

func init() { proto.RegisterFile("services/verify/verify.proto", fileDescriptor_12065c2e72b67281) }

var fileDescriptor_12065c2e72b67281 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0x09, 0xbc, 0xf4, 0x95, 0x69, 0x0b, 0xb2, 0x60, 0x28, 0x51, 0x51, 0xa2, 0x87, 0x9e,
	0x26, 0x50, 0xf1, 0x0b, 0x08, 0x3d, 0x79, 0xd1, 0x14, 0x3c, 0x78, 0xcb, 0x26, 0xd3, 0x75, 0xa1,
	0xc9, 0xc6, 0x9d, 0x4d, 0xd1, 0xcf, 0xe8, 0x97, 0x92, 0xee, 0x26, 0x20, 0xb6, 0xf6, 0xb4, 0x0c,
	0xcf, 0x1f, 0xf6, 0xf7, 0xc0, 0x05, 0x93, 0xdd, 0xea, 0x92, 0x38, 0xdb, 0x92, 0xd5, 0xeb, 0xcf,
	0xfe, 0xc1, 0xd6, 0x1a, 0x67, 0x44, 0x2c, 0x8b, 0x5a, 0x1a, 0x83, 0x83, 0x09, 0x83, 0x9a, 0x9c,
	0xf1, 0x5b, 0x61, 0xa9, 0xca, 0x6a, 0x62, 0x2e, 0x14, 0x71, 0xb0, 0x27, 0xe7, 0xca, 0x18, 0xb5,
	0xa1, 0xcc, 0x5f, 0xb2, 0x5b, 0x67, 0x54, 0xb7, 0xae, 0xef, 0x4a, 0xa7, 0x30, 0x7e, 0xd2, 0x8d,
	0xca, 0xe9, 0xbd, 0x23, 0x76, 0xe9, 0x1c, 0x26, 0xe1, 0xe4, 0xd6, 0x34, 0x4c, 0x62, 0x06, 0xff,
	0x8b, 0xaa, 0xb2, 0xc4, 0x3c, 0x8b, 0xae, 0xa3, 0xf9, 0x24, 0x1f, 0xce, 0x74, 0x03, 0x97, 0xab,
	0x4e, 0xd6, 0xda, 0x2d, 0x3f, 0xa8, 0xec, 0x9c, 0x36, 0x4d, 0x4e, 0x25, 0xe9, 0xd6, 0xf5, 0x55,
	0xe2, 0x11, 0x4e, 0xe9, 0x97, 0xe4, 0x3b, 0xc6, 0x8b, 0x2b, 0x1c, 0x00, 0xfc, 0x7f, 0x71, 0xaf,
	0x61, 0x2f, 0xb8, 0xf8, 0x8a, 0x60, 0xfa, 0xe2, 0x29, 0x57, 0x01, 0x5a, 0x3c, 0xc3, 0xbf, 0xdd,
	0x4f, 0xc5, 0x0d, 0x1e, 0x5e, 0x03, 0x7f, 0x60, 0x25, 0xb7, 0xc7, 0x4d, 0x3d, 0xac, 0x82, 0xf8,
	0x30, 0x92, 0xb8, 0xff, 0x2b, 0x7f, 0x74, 0x82, 0x24, 0xc6, 0x30, 0x3d, 0x0e, 0xd3, 0xe3, 0x72,
	0x37, 0xfd, 0xc3, 0xc9, 0xeb, 0x28, 0xe4, 0xe5, 0xc8, 0x2b, 0x77, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x84, 0x11, 0xc6, 0x16, 0xf1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VerifyServiceClient is the client API for VerifyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VerifyServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	SubmitExecutionReceipt(ctx context.Context, in *SubmitExecutionReceiptRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type verifyServiceClient struct {
	cc *grpc.ClientConn
}

func NewVerifyServiceClient(cc *grpc.ClientConn) VerifyServiceClient {
	return &verifyServiceClient{cc}
}

func (c *verifyServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/bamboo.services.verify.VerifyService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifyServiceClient) SubmitExecutionReceipt(ctx context.Context, in *SubmitExecutionReceiptRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/bamboo.services.verify.VerifyService/SubmitExecutionReceipt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VerifyServiceServer is the server API for VerifyService service.
type VerifyServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	SubmitExecutionReceipt(context.Context, *SubmitExecutionReceiptRequest) (*empty.Empty, error)
}

func RegisterVerifyServiceServer(s *grpc.Server, srv VerifyServiceServer) {
	s.RegisterService(&_VerifyService_serviceDesc, srv)
}

func _VerifyService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifyServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.services.verify.VerifyService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifyServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerifyService_SubmitExecutionReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitExecutionReceiptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifyServiceServer).SubmitExecutionReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.services.verify.VerifyService/SubmitExecutionReceipt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifyServiceServer).SubmitExecutionReceipt(ctx, req.(*SubmitExecutionReceiptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VerifyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bamboo.services.verify.VerifyService",
	HandlerType: (*VerifyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _VerifyService_Ping_Handler,
		},
		{
			MethodName: "SubmitExecutionReceipt",
			Handler:    _VerifyService_SubmitExecutionReceipt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/verify/verify.proto",
}
