// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/execute/execute.proto

package execute

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
	return fileDescriptor_04325f5bc58d3094, []int{0}
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
	return fileDescriptor_04325f5bc58d3094, []int{1}
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

type ExecuteBlockRequest struct {
	Block                *shared.Block `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ExecuteBlockRequest) Reset()         { *m = ExecuteBlockRequest{} }
func (m *ExecuteBlockRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteBlockRequest) ProtoMessage()    {}
func (*ExecuteBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_04325f5bc58d3094, []int{2}
}

func (m *ExecuteBlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteBlockRequest.Unmarshal(m, b)
}
func (m *ExecuteBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteBlockRequest.Marshal(b, m, deterministic)
}
func (m *ExecuteBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteBlockRequest.Merge(m, src)
}
func (m *ExecuteBlockRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteBlockRequest.Size(m)
}
func (m *ExecuteBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteBlockRequest proto.InternalMessageInfo

func (m *ExecuteBlockRequest) GetBlock() *shared.Block {
	if m != nil {
		return m.Block
	}
	return nil
}

type NotifyBlockExecutedRequest struct {
	Block                *shared.Block `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NotifyBlockExecutedRequest) Reset()         { *m = NotifyBlockExecutedRequest{} }
func (m *NotifyBlockExecutedRequest) String() string { return proto.CompactTextString(m) }
func (*NotifyBlockExecutedRequest) ProtoMessage()    {}
func (*NotifyBlockExecutedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_04325f5bc58d3094, []int{3}
}

func (m *NotifyBlockExecutedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyBlockExecutedRequest.Unmarshal(m, b)
}
func (m *NotifyBlockExecutedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyBlockExecutedRequest.Marshal(b, m, deterministic)
}
func (m *NotifyBlockExecutedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyBlockExecutedRequest.Merge(m, src)
}
func (m *NotifyBlockExecutedRequest) XXX_Size() int {
	return xxx_messageInfo_NotifyBlockExecutedRequest.Size(m)
}
func (m *NotifyBlockExecutedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyBlockExecutedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyBlockExecutedRequest proto.InternalMessageInfo

func (m *NotifyBlockExecutedRequest) GetBlock() *shared.Block {
	if m != nil {
		return m.Block
	}
	return nil
}

type ReceiptsByBlockHeightRequest struct {
	BlockHeight          uint64   `protobuf:"varint,1,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiptsByBlockHeightRequest) Reset()         { *m = ReceiptsByBlockHeightRequest{} }
func (m *ReceiptsByBlockHeightRequest) String() string { return proto.CompactTextString(m) }
func (*ReceiptsByBlockHeightRequest) ProtoMessage()    {}
func (*ReceiptsByBlockHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_04325f5bc58d3094, []int{4}
}

func (m *ReceiptsByBlockHeightRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiptsByBlockHeightRequest.Unmarshal(m, b)
}
func (m *ReceiptsByBlockHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiptsByBlockHeightRequest.Marshal(b, m, deterministic)
}
func (m *ReceiptsByBlockHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiptsByBlockHeightRequest.Merge(m, src)
}
func (m *ReceiptsByBlockHeightRequest) XXX_Size() int {
	return xxx_messageInfo_ReceiptsByBlockHeightRequest.Size(m)
}
func (m *ReceiptsByBlockHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiptsByBlockHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiptsByBlockHeightRequest proto.InternalMessageInfo

func (m *ReceiptsByBlockHeightRequest) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

type ReceiptsByBlockHeightResponse struct {
	ExecutionReceipts    []*shared.ExecutionReceipt `protobuf:"bytes,1,rep,name=executionReceipts,proto3" json:"executionReceipts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ReceiptsByBlockHeightResponse) Reset()         { *m = ReceiptsByBlockHeightResponse{} }
func (m *ReceiptsByBlockHeightResponse) String() string { return proto.CompactTextString(m) }
func (*ReceiptsByBlockHeightResponse) ProtoMessage()    {}
func (*ReceiptsByBlockHeightResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_04325f5bc58d3094, []int{5}
}

func (m *ReceiptsByBlockHeightResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiptsByBlockHeightResponse.Unmarshal(m, b)
}
func (m *ReceiptsByBlockHeightResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiptsByBlockHeightResponse.Marshal(b, m, deterministic)
}
func (m *ReceiptsByBlockHeightResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiptsByBlockHeightResponse.Merge(m, src)
}
func (m *ReceiptsByBlockHeightResponse) XXX_Size() int {
	return xxx_messageInfo_ReceiptsByBlockHeightResponse.Size(m)
}
func (m *ReceiptsByBlockHeightResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiptsByBlockHeightResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiptsByBlockHeightResponse proto.InternalMessageInfo

func (m *ReceiptsByBlockHeightResponse) GetExecutionReceipts() []*shared.ExecutionReceipt {
	if m != nil {
		return m.ExecutionReceipts
	}
	return nil
}

type GetRegistersRequest struct {
	Ids                  [][]byte `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRegistersRequest) Reset()         { *m = GetRegistersRequest{} }
func (m *GetRegistersRequest) String() string { return proto.CompactTextString(m) }
func (*GetRegistersRequest) ProtoMessage()    {}
func (*GetRegistersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_04325f5bc58d3094, []int{6}
}

func (m *GetRegistersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRegistersRequest.Unmarshal(m, b)
}
func (m *GetRegistersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRegistersRequest.Marshal(b, m, deterministic)
}
func (m *GetRegistersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRegistersRequest.Merge(m, src)
}
func (m *GetRegistersRequest) XXX_Size() int {
	return xxx_messageInfo_GetRegistersRequest.Size(m)
}
func (m *GetRegistersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRegistersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRegistersRequest proto.InternalMessageInfo

func (m *GetRegistersRequest) GetIds() [][]byte {
	if m != nil {
		return m.Ids
	}
	return nil
}

type GetRegistersAtBlockHeightRequest struct {
	Ids                  [][]byte `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	BlockHeight          uint64   `protobuf:"varint,2,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRegistersAtBlockHeightRequest) Reset()         { *m = GetRegistersAtBlockHeightRequest{} }
func (m *GetRegistersAtBlockHeightRequest) String() string { return proto.CompactTextString(m) }
func (*GetRegistersAtBlockHeightRequest) ProtoMessage()    {}
func (*GetRegistersAtBlockHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_04325f5bc58d3094, []int{7}
}

func (m *GetRegistersAtBlockHeightRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRegistersAtBlockHeightRequest.Unmarshal(m, b)
}
func (m *GetRegistersAtBlockHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRegistersAtBlockHeightRequest.Marshal(b, m, deterministic)
}
func (m *GetRegistersAtBlockHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRegistersAtBlockHeightRequest.Merge(m, src)
}
func (m *GetRegistersAtBlockHeightRequest) XXX_Size() int {
	return xxx_messageInfo_GetRegistersAtBlockHeightRequest.Size(m)
}
func (m *GetRegistersAtBlockHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRegistersAtBlockHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRegistersAtBlockHeightRequest proto.InternalMessageInfo

func (m *GetRegistersAtBlockHeightRequest) GetIds() [][]byte {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *GetRegistersAtBlockHeightRequest) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

type GetRegistersResponse struct {
	Registers            []*shared.Register `protobuf:"bytes,1,rep,name=registers,proto3" json:"registers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetRegistersResponse) Reset()         { *m = GetRegistersResponse{} }
func (m *GetRegistersResponse) String() string { return proto.CompactTextString(m) }
func (*GetRegistersResponse) ProtoMessage()    {}
func (*GetRegistersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_04325f5bc58d3094, []int{8}
}

func (m *GetRegistersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRegistersResponse.Unmarshal(m, b)
}
func (m *GetRegistersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRegistersResponse.Marshal(b, m, deterministic)
}
func (m *GetRegistersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRegistersResponse.Merge(m, src)
}
func (m *GetRegistersResponse) XXX_Size() int {
	return xxx_messageInfo_GetRegistersResponse.Size(m)
}
func (m *GetRegistersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRegistersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRegistersResponse proto.InternalMessageInfo

func (m *GetRegistersResponse) GetRegisters() []*shared.Register {
	if m != nil {
		return m.Registers
	}
	return nil
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "bamboo.services.execute.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "bamboo.services.execute.PingResponse")
	proto.RegisterType((*ExecuteBlockRequest)(nil), "bamboo.services.execute.ExecuteBlockRequest")
	proto.RegisterType((*NotifyBlockExecutedRequest)(nil), "bamboo.services.execute.NotifyBlockExecutedRequest")
	proto.RegisterType((*ReceiptsByBlockHeightRequest)(nil), "bamboo.services.execute.ReceiptsByBlockHeightRequest")
	proto.RegisterType((*ReceiptsByBlockHeightResponse)(nil), "bamboo.services.execute.ReceiptsByBlockHeightResponse")
	proto.RegisterType((*GetRegistersRequest)(nil), "bamboo.services.execute.GetRegistersRequest")
	proto.RegisterType((*GetRegistersAtBlockHeightRequest)(nil), "bamboo.services.execute.GetRegistersAtBlockHeightRequest")
	proto.RegisterType((*GetRegistersResponse)(nil), "bamboo.services.execute.GetRegistersResponse")
}

func init() { proto.RegisterFile("services/execute/execute.proto", fileDescriptor_04325f5bc58d3094) }

var fileDescriptor_04325f5bc58d3094 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x55, 0x58, 0x61, 0xda, 0x4d, 0x40, 0xe0, 0x0e, 0x56, 0xc2, 0x57, 0x14, 0x81, 0xa8, 0x10,
	0x38, 0x52, 0x27, 0x1e, 0x78, 0x63, 0x95, 0x2a, 0xf6, 0x32, 0x84, 0x3c, 0xc4, 0x03, 0x6f, 0xf9,
	0xb8, 0xcb, 0xac, 0xad, 0x71, 0x88, 0x5d, 0xc4, 0x7e, 0x00, 0xff, 0x80, 0x1f, 0x8c, 0x12, 0xdb,
	0x6a, 0x48, 0x13, 0x28, 0x7b, 0x6a, 0xef, 0xf5, 0xb9, 0xc7, 0xe7, 0xdc, 0x1c, 0xc3, 0x53, 0x89,
	0xd5, 0x77, 0x9e, 0xa2, 0x8c, 0xf0, 0x07, 0xa6, 0x2b, 0x85, 0xf6, 0x97, 0x96, 0x95, 0x50, 0x82,
	0x1c, 0x24, 0xf1, 0x32, 0x11, 0x82, 0x5a, 0x18, 0x35, 0xc7, 0xfe, 0x7d, 0x79, 0x1e, 0x57, 0x98,
	0x45, 0x4b, 0x94, 0x32, 0xce, 0x51, 0x6a, 0xbc, 0xff, 0x28, 0x17, 0x22, 0xbf, 0xc4, 0xa8, 0xa9,
	0x92, 0xd5, 0x59, 0x84, 0xcb, 0x52, 0x5d, 0xe9, 0xc3, 0xf0, 0x36, 0xb8, 0x9f, 0x78, 0x91, 0x33,
	0xfc, 0xb6, 0x42, 0xa9, 0xc2, 0x29, 0x78, 0xba, 0x94, 0xa5, 0x28, 0x24, 0x92, 0x09, 0xec, 0xc6,
	0x59, 0x56, 0xa1, 0x94, 0x13, 0x27, 0x70, 0xa6, 0x1e, 0xb3, 0x65, 0x78, 0x04, 0xe3, 0x85, 0xbe,
	0x77, 0x7e, 0x29, 0xd2, 0x0b, 0x43, 0x40, 0x5e, 0xc1, 0xcd, 0xa4, 0xae, 0x1b, 0xb8, 0x3b, 0xdb,
	0xa7, 0x56, 0x6c, 0x23, 0x8d, 0x6a, 0xac, 0x86, 0x84, 0xc7, 0xe0, 0x7f, 0x14, 0x8a, 0x9f, 0x5d,
	0x35, 0x5d, 0xc3, 0x96, 0x5d, 0x87, 0xe9, 0x3d, 0x3c, 0x66, 0x98, 0x22, 0x2f, 0x95, 0x9c, 0x6b,
	0xb6, 0x63, 0xe4, 0xf9, 0xb9, 0xb2, 0x5c, 0x01, 0xb8, 0xc9, 0xba, 0xdb, 0x30, 0x8e, 0x58, 0xbb,
	0x15, 0x16, 0xf0, 0x64, 0x80, 0xc1, 0x6c, 0xe2, 0x04, 0xee, 0xe9, 0x3d, 0x73, 0x51, 0x58, 0xe4,
	0xc4, 0x09, 0x76, 0xa6, 0xee, 0xec, 0x59, 0x47, 0xda, 0xa2, 0x83, 0x63, 0x9b, 0x93, 0xe1, 0x4b,
	0x18, 0x7f, 0x40, 0xc5, 0x30, 0xe7, 0x52, 0x61, 0x25, 0xad, 0xd0, 0xbb, 0xb0, 0xc3, 0x33, 0xcd,
	0xeb, 0xb1, 0xfa, 0x6f, 0xf8, 0x05, 0x82, 0x36, 0xf0, 0x48, 0xf5, 0xd8, 0xdb, 0x98, 0xea, 0x1a,
	0xbe, 0xb1, 0x69, 0xf8, 0x04, 0xf6, 0xff, 0x14, 0x60, 0x7c, 0xbe, 0x85, 0xbd, 0xca, 0x36, 0x8d,
	0xbf, 0x83, 0x8e, 0x3f, 0x3b, 0xc4, 0xd6, 0xc8, 0xd9, 0xaf, 0x11, 0xdc, 0x31, 0x5f, 0xf0, 0x54,
	0xe7, 0x92, 0x9c, 0xc2, 0xa8, 0xce, 0x12, 0x79, 0x4e, 0x07, 0x02, 0x4b, 0x5b, 0xc9, 0xf3, 0x5f,
	0xfc, 0x03, 0x65, 0xe4, 0x7d, 0x06, 0xaf, 0x1d, 0x3b, 0xf2, 0x7a, 0x70, 0xac, 0x27, 0x9d, 0xfe,
	0x03, 0xaa, 0xdf, 0x02, 0xb5, 0x6f, 0x81, 0x2e, 0xea, 0xb7, 0x40, 0x12, 0x18, 0xf7, 0x24, 0x91,
	0x1c, 0x0e, 0x92, 0x0f, 0xe7, 0x76, 0xf0, 0x8e, 0x0b, 0xf0, 0xda, 0x0b, 0xff, 0x8b, 0xf2, 0x9e,
	0x60, 0xf8, 0x6f, 0xb6, 0x44, 0x9b, 0x35, 0xfd, 0x74, 0xe0, 0xe1, 0x60, 0x6c, 0xc8, 0xbb, 0xad,
	0xc8, 0xfa, 0xa2, 0xf6, 0x9f, 0x3a, 0xe6, 0x7b, 0x5f, 0x77, 0xcd, 0x79, 0x72, 0xab, 0xd9, 0xc7,
	0xe1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0x52, 0x41, 0xa2, 0xdf, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExecuteServiceClient is the client API for ExecuteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExecuteServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	ExecuteBlock(ctx context.Context, in *ExecuteBlockRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	NotifyBlockExecuted(ctx context.Context, in *NotifyBlockExecutedRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetRegisters(ctx context.Context, in *GetRegistersRequest, opts ...grpc.CallOption) (*GetRegistersResponse, error)
	GetRegistersAtBlockHeight(ctx context.Context, in *GetRegistersAtBlockHeightRequest, opts ...grpc.CallOption) (*GetRegistersResponse, error)
}

type executeServiceClient struct {
	cc *grpc.ClientConn
}

func NewExecuteServiceClient(cc *grpc.ClientConn) ExecuteServiceClient {
	return &executeServiceClient{cc}
}

func (c *executeServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/bamboo.services.execute.ExecuteService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executeServiceClient) ExecuteBlock(ctx context.Context, in *ExecuteBlockRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/bamboo.services.execute.ExecuteService/ExecuteBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executeServiceClient) NotifyBlockExecuted(ctx context.Context, in *NotifyBlockExecutedRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/bamboo.services.execute.ExecuteService/NotifyBlockExecuted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executeServiceClient) GetRegisters(ctx context.Context, in *GetRegistersRequest, opts ...grpc.CallOption) (*GetRegistersResponse, error) {
	out := new(GetRegistersResponse)
	err := c.cc.Invoke(ctx, "/bamboo.services.execute.ExecuteService/GetRegisters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executeServiceClient) GetRegistersAtBlockHeight(ctx context.Context, in *GetRegistersAtBlockHeightRequest, opts ...grpc.CallOption) (*GetRegistersResponse, error) {
	out := new(GetRegistersResponse)
	err := c.cc.Invoke(ctx, "/bamboo.services.execute.ExecuteService/GetRegistersAtBlockHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecuteServiceServer is the server API for ExecuteService service.
type ExecuteServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	ExecuteBlock(context.Context, *ExecuteBlockRequest) (*empty.Empty, error)
	NotifyBlockExecuted(context.Context, *NotifyBlockExecutedRequest) (*empty.Empty, error)
	GetRegisters(context.Context, *GetRegistersRequest) (*GetRegistersResponse, error)
	GetRegistersAtBlockHeight(context.Context, *GetRegistersAtBlockHeightRequest) (*GetRegistersResponse, error)
}

func RegisterExecuteServiceServer(s *grpc.Server, srv ExecuteServiceServer) {
	s.RegisterService(&_ExecuteService_serviceDesc, srv)
}

func _ExecuteService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecuteServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.services.execute.ExecuteService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecuteServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecuteService_ExecuteBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecuteServiceServer).ExecuteBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.services.execute.ExecuteService/ExecuteBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecuteServiceServer).ExecuteBlock(ctx, req.(*ExecuteBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecuteService_NotifyBlockExecuted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyBlockExecutedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecuteServiceServer).NotifyBlockExecuted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.services.execute.ExecuteService/NotifyBlockExecuted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecuteServiceServer).NotifyBlockExecuted(ctx, req.(*NotifyBlockExecutedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecuteService_GetRegisters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegistersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecuteServiceServer).GetRegisters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.services.execute.ExecuteService/GetRegisters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecuteServiceServer).GetRegisters(ctx, req.(*GetRegistersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecuteService_GetRegistersAtBlockHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegistersAtBlockHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecuteServiceServer).GetRegistersAtBlockHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.services.execute.ExecuteService/GetRegistersAtBlockHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecuteServiceServer).GetRegistersAtBlockHeight(ctx, req.(*GetRegistersAtBlockHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExecuteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bamboo.services.execute.ExecuteService",
	HandlerType: (*ExecuteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ExecuteService_Ping_Handler,
		},
		{
			MethodName: "ExecuteBlock",
			Handler:    _ExecuteService_ExecuteBlock_Handler,
		},
		{
			MethodName: "NotifyBlockExecuted",
			Handler:    _ExecuteService_NotifyBlockExecuted_Handler,
		},
		{
			MethodName: "GetRegisters",
			Handler:    _ExecuteService_GetRegisters_Handler,
		},
		{
			MethodName: "GetRegistersAtBlockHeight",
			Handler:    _ExecuteService_GetRegistersAtBlockHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/execute/execute.proto",
}
