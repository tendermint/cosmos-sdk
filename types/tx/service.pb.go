// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/tx/service.proto

package tx

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// GenerateRequest is the request type for the TxService.Generate RPC method.
type SimulateRequest struct {
	Tx *Tx `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (m *SimulateRequest) Reset()         { *m = SimulateRequest{} }
func (m *SimulateRequest) String() string { return proto.CompactTextString(m) }
func (*SimulateRequest) ProtoMessage()    {}
func (*SimulateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3815655f4ee03b11, []int{0}
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

// GenerateResponse is the response type for the TxService.Generate RPC method.
type SimulateResponse struct {
	SimulationResponse *types.SimulationResponse `protobuf:"bytes,1,opt,name=simulation_response,json=simulationResponse,proto3" json:"simulation_response,omitempty"`
}

func (m *SimulateResponse) Reset()         { *m = SimulateResponse{} }
func (m *SimulateResponse) String() string { return proto.CompactTextString(m) }
func (*SimulateResponse) ProtoMessage()    {}
func (*SimulateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3815655f4ee03b11, []int{1}
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

func (m *SimulateResponse) GetSimulationResponse() *types.SimulationResponse {
	if m != nil {
		return m.SimulationResponse
	}
	return nil
}

func init() {
	proto.RegisterType((*SimulateRequest)(nil), "cosmos.tx.SimulateRequest")
	proto.RegisterType((*SimulateResponse)(nil), "cosmos.tx.SimulateResponse")
}

func init() { proto.RegisterFile("cosmos/tx/service.proto", fileDescriptor_3815655f4ee03b11) }

var fileDescriptor_3815655f4ee03b11 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x2f, 0xa9, 0xd0, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x48, 0xe8, 0x95, 0x54, 0x48, 0x09, 0x43, 0xd5, 0x40, 0x45,
	0xc0, 0xf2, 0x52, 0x42, 0x08, 0x8d, 0x25, 0x15, 0x10, 0x31, 0x25, 0x03, 0x2e, 0xfe, 0xe0, 0xcc,
	0xdc, 0xd2, 0x9c, 0xc4, 0x92, 0xd4, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x59, 0x2e,
	0xa6, 0x92, 0x0a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x5e, 0x3d, 0xb8, 0x99, 0x7a, 0x21,
	0x15, 0x41, 0x4c, 0x25, 0x15, 0x4a, 0xf1, 0x5c, 0x02, 0x08, 0x1d, 0xc5, 0x05, 0xf9, 0x79, 0xc5,
	0xa9, 0x42, 0xde, 0x5c, 0xc2, 0xc5, 0x10, 0xb1, 0xcc, 0xfc, 0xbc, 0xf8, 0x22, 0xa8, 0x30, 0xd4,
	0x0c, 0x29, 0x98, 0x19, 0xc1, 0x70, 0x25, 0x30, 0x8d, 0x41, 0x42, 0xc5, 0x18, 0x62, 0x46, 0x01,
	0x5c, 0xec, 0xc1, 0x10, 0x7f, 0x09, 0xb9, 0x72, 0x71, 0xc0, 0xec, 0x12, 0x92, 0x42, 0x72, 0x0a,
	0x9a, 0x93, 0xa5, 0xa4, 0xb1, 0xca, 0x41, 0xcc, 0x53, 0x62, 0x70, 0xb2, 0x3f, 0xf1, 0x48, 0x8e,
	0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58,
	0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xd5, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4,
	0xfc, 0x5c, 0x7d, 0x94, 0x20, 0xd3, 0x2d, 0x4e, 0xc9, 0xd6, 0x2f, 0xa9, 0x2c, 0x48, 0x05, 0x05,
	0x57, 0x12, 0x1b, 0x38, 0xb0, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x67, 0xd4, 0x8a,
	0x7b, 0x01, 0x00, 0x00,
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
}

type serviceClient struct {
	cc grpc1.ClientConn
}

func NewServiceClient(cc grpc1.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Simulate(ctx context.Context, in *SimulateRequest, opts ...grpc.CallOption) (*SimulateResponse, error) {
	out := new(SimulateResponse)
	err := c.cc.Invoke(ctx, "/cosmos.tx.Service/Simulate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	// Simulate simulates executing a transaction for estimating gas usage.
	Simulate(context.Context, *SimulateRequest) (*SimulateResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Simulate(ctx context.Context, req *SimulateRequest) (*SimulateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Simulate not implemented")
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
		FullMethod: "/cosmos.tx.Service/Simulate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Simulate(ctx, req.(*SimulateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.tx.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Simulate",
			Handler:    _Service_Simulate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/tx/service.proto",
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
	if m.SimulationResponse != nil {
		{
			size, err := m.SimulationResponse.MarshalToSizedBuffer(dAtA[:i])
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
	if m.SimulationResponse != nil {
		l = m.SimulationResponse.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field SimulationResponse", wireType)
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
			if m.SimulationResponse == nil {
				m.SimulationResponse = &types.SimulationResponse{}
			}
			if err := m.SimulationResponse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
