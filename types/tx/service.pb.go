// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/tx/v1beta1/service.proto

package tx

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// SimulateRequest is the request type for the Service.Simulate
// RPC method.
type SimulateRequest struct {
	// tx is the transaction to simulate.
	Tx *Tx `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (m *SimulateRequest) Reset()         { *m = SimulateRequest{} }
func (m *SimulateRequest) String() string { return proto.CompactTextString(m) }
func (*SimulateRequest) ProtoMessage()    {}
func (*SimulateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b00a618705eca7, []int{0}
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

// SimulateResponse is the response type for the
// Service.SimulateRPC method.
type SimulateResponse struct {
	// gas_info is the information about gas used in the simulation.
	GasInfo *types.GasInfo `protobuf:"bytes,1,opt,name=gas_info,json=gasInfo,proto3" json:"gas_info,omitempty"`
	// result is the result of the simulation.
	Result *types.Result `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *SimulateResponse) Reset()         { *m = SimulateResponse{} }
func (m *SimulateResponse) String() string { return proto.CompactTextString(m) }
func (*SimulateResponse) ProtoMessage()    {}
func (*SimulateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b00a618705eca7, []int{1}
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

// GetTx is the request type for the Service.GetTx
// RPC method.
type GetTxRequest struct {
	// hash is the tx hash to query, encoded as a hex string.
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *GetTxRequest) Reset()         { *m = GetTxRequest{} }
func (m *GetTxRequest) String() string { return proto.CompactTextString(m) }
func (*GetTxRequest) ProtoMessage()    {}
func (*GetTxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b00a618705eca7, []int{2}
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

// GetTxResponse is the response type for the Service.GetTx method.
type GetTxResponse struct {
	// tx is the queried transaction.
	Tx *Tx `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (m *GetTxResponse) Reset()         { *m = GetTxResponse{} }
func (m *GetTxResponse) String() string { return proto.CompactTextString(m) }
func (*GetTxResponse) ProtoMessage()    {}
func (*GetTxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b00a618705eca7, []int{3}
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

func init() {
	proto.RegisterType((*SimulateRequest)(nil), "cosmos.tx.v1beta1.SimulateRequest")
	proto.RegisterType((*SimulateResponse)(nil), "cosmos.tx.v1beta1.SimulateResponse")
	proto.RegisterType((*GetTxRequest)(nil), "cosmos.tx.v1beta1.GetTxRequest")
	proto.RegisterType((*GetTxResponse)(nil), "cosmos.tx.v1beta1.GetTxResponse")
}

func init() { proto.RegisterFile("cosmos/tx/v1beta1/service.proto", fileDescriptor_e0b00a618705eca7) }

var fileDescriptor_e0b00a618705eca7 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x31, 0x6b, 0xdb, 0x40,
	0x1c, 0xc5, 0x2d, 0xd1, 0xda, 0xee, 0xb5, 0xa5, 0xed, 0x41, 0xc1, 0xa8, 0xae, 0xec, 0x9e, 0x6b,
	0xe8, 0x52, 0x1d, 0x76, 0xa1, 0x78, 0x08, 0x04, 0xb2, 0x98, 0xac, 0xb2, 0xa7, 0x2c, 0xe1, 0xa4,
	0x9c, 0x65, 0x11, 0x5b, 0xa7, 0xf8, 0x4e, 0xe6, 0x42, 0xc8, 0x92, 0x31, 0x53, 0x20, 0x5f, 0x2a,
	0xa3, 0x21, 0x4b, 0xc6, 0x60, 0xe7, 0x2b, 0x64, 0x0f, 0x3a, 0x9d, 0x92, 0x10, 0xcb, 0x21, 0x93,
	0x4e, 0xe8, 0xbd, 0xdf, 0xff, 0xbd, 0xd3, 0x1f, 0x34, 0x7c, 0xc6, 0xa7, 0x8c, 0x63, 0x21, 0xf1,
	0xbc, 0xe3, 0x51, 0x41, 0x3a, 0x98, 0xd3, 0xd9, 0x3c, 0xf4, 0xa9, 0x13, 0xcf, 0x98, 0x60, 0xf0,
	0x5b, 0x26, 0x70, 0x84, 0x74, 0xb4, 0xc0, 0xaa, 0x07, 0x8c, 0x05, 0x13, 0x8a, 0x49, 0x1c, 0x62,
	0x12, 0x45, 0x4c, 0x10, 0x11, 0xb2, 0x88, 0x67, 0x06, 0xab, 0xa5, 0x89, 0x1e, 0xe1, 0x14, 0x13,
	0xcf, 0x0f, 0x1f, 0xc1, 0xe9, 0x8b, 0x16, 0x59, 0xeb, 0x63, 0x85, 0xcc, 0xbe, 0xa1, 0x1e, 0xf8,
	0x32, 0x08, 0xa7, 0xc9, 0x84, 0x08, 0xea, 0xd2, 0xa3, 0x84, 0x72, 0x01, 0xdb, 0xc0, 0x14, 0xb2,
	0x66, 0x34, 0x8d, 0x3f, 0x1f, 0xbb, 0xdf, 0x9d, 0xb5, 0x44, 0xce, 0x50, 0xba, 0xa6, 0x90, 0xe8,
	0xdc, 0x00, 0x5f, 0x9f, 0xac, 0x3c, 0x66, 0x11, 0xa7, 0x70, 0x0b, 0x54, 0x03, 0xc2, 0xf7, 0xc3,
	0x68, 0xc4, 0x34, 0xe1, 0x57, 0x4e, 0x48, 0x23, 0x3a, 0x2a, 0x55, 0x0e, 0xea, 0x13, 0xbe, 0x1b,
	0x8d, 0x98, 0x5b, 0x09, 0xb2, 0x03, 0xec, 0x81, 0xf2, 0x8c, 0xf2, 0x64, 0x22, 0x6a, 0xa6, 0xf2,
	0x36, 0x37, 0x7b, 0x5d, 0xa5, 0x73, 0xb5, 0x1e, 0x21, 0xf0, 0xa9, 0x4f, 0xc5, 0x50, 0xe6, 0x1d,
	0x20, 0x78, 0x37, 0x26, 0x7c, 0xac, 0x32, 0x7c, 0x70, 0xd5, 0x19, 0xfd, 0x07, 0x9f, 0xb5, 0x46,
	0x87, 0x7d, 0x5b, 0xd1, 0xee, 0xbd, 0x01, 0x2a, 0x83, 0xec, 0x37, 0x41, 0x09, 0xaa, 0x79, 0x67,
	0x88, 0x0a, 0x2c, 0x2f, 0xee, 0xd2, 0x6a, 0xbd, 0xaa, 0xc9, 0x72, 0xa0, 0xd6, 0xd9, 0xf5, 0xdd,
	0xa5, 0xf9, 0x13, 0xfd, 0xc0, 0x05, 0xfb, 0x91, 0x4f, 0x8b, 0xc1, 0x7b, 0x95, 0x1e, 0x36, 0x0a,
	0x90, 0xcf, 0xbb, 0x5b, 0xcd, 0xcd, 0x02, 0x3d, 0xf0, 0xb7, 0x1a, 0x68, 0xc3, 0x3a, 0x2e, 0xda,
	0x0c, 0x7c, 0x92, 0x5e, 0xd7, 0xe9, 0xce, 0xf6, 0xd5, 0xd2, 0x36, 0x16, 0x4b, 0xdb, 0xb8, 0x5d,
	0xda, 0xc6, 0xc5, 0xca, 0x2e, 0x2d, 0x56, 0x76, 0xe9, 0x66, 0x65, 0x97, 0xf6, 0xda, 0x41, 0x28,
	0xc6, 0x89, 0xe7, 0xf8, 0x6c, 0x9a, 0x13, 0xb2, 0xc7, 0x5f, 0x7e, 0x70, 0x88, 0xc5, 0x71, 0x4c,
	0x53, 0xa4, 0x57, 0x56, 0x2b, 0xf6, 0xef, 0x21, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x4c, 0x40, 0xff,
	0xf7, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	// Simulate simulates executing a transaction for estimating gas usage.
	Simulate(ctx context.Context, in *SimulateRequest, opts ...grpc.CallOption) (*SimulateResponse, error)
	// GetTx fetches a tx by hash.
	GetTx(ctx context.Context, in *GetTxRequest, opts ...grpc.CallOption) (*GetTxResponse, error)
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

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	// Simulate simulates executing a transaction for estimating gas usage.
	Simulate(context.Context, *SimulateRequest) (*SimulateResponse, error)
	// GetTx fetches a tx by hash.
	GetTx(context.Context, *GetTxRequest) (*GetTxResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Simulate(ctx context.Context, req *SimulateRequest) (*SimulateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Simulate not implemented")
}
func (*UnimplementedServiceServer) GetTx(ctx context.Context, req *GetTxRequest) (*GetTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTx not implemented")
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/tx/v1beta1/service.proto",
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
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
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
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
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
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
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
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
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
