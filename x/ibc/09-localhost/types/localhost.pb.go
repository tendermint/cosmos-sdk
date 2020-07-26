// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/localhost/localhost.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// MsgCreateClient defines a message to create an IBC client
type MsgCreateClient struct {
	Signer github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=signer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"signer,omitempty"`
}

func (m *MsgCreateClient) Reset()         { *m = MsgCreateClient{} }
func (m *MsgCreateClient) String() string { return proto.CompactTextString(m) }
func (*MsgCreateClient) ProtoMessage()    {}
func (*MsgCreateClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a04d924e6f8a88e, []int{0}
}
func (m *MsgCreateClient) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateClient.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateClient.Merge(m, src)
}
func (m *MsgCreateClient) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateClient) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateClient.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateClient proto.InternalMessageInfo

func (m *MsgCreateClient) GetSigner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Signer
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgCreateClient)(nil), "ibc.localhost.MsgCreateClient")
}

func init() { proto.RegisterFile("ibc/localhost/localhost.proto", fileDescriptor_6a04d924e6f8a88e) }

var fileDescriptor_6a04d924e6f8a88e = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0x4c, 0x4a, 0xd6,
	0xcf, 0xc9, 0x4f, 0x4e, 0xcc, 0xc9, 0xc8, 0x2f, 0x2e, 0x41, 0xb0, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x78, 0x33, 0x93, 0x92, 0xf5, 0xe0, 0x82, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60,
	0x19, 0x7d, 0x10, 0x0b, 0xa2, 0x48, 0x29, 0x86, 0x8b, 0xdf, 0xb7, 0x38, 0xdd, 0xb9, 0x28, 0x35,
	0xb1, 0x24, 0xd5, 0x39, 0x27, 0x33, 0x35, 0xaf, 0x44, 0xc8, 0x93, 0x8b, 0xad, 0x38, 0x33, 0x3d,
	0x2f, 0xb5, 0x48, 0x82, 0x51, 0x81, 0x51, 0x83, 0xc7, 0xc9, 0xf0, 0xd7, 0x3d, 0x79, 0xdd, 0xf4,
	0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xe4, 0xfc, 0xe2, 0xdc, 0xfc, 0x62,
	0x28, 0xa5, 0x5b, 0x9c, 0x92, 0xad, 0x5f, 0x52, 0x59, 0x90, 0x5a, 0xac, 0xe7, 0x98, 0x9c, 0xec,
	0x98, 0x92, 0x52, 0x94, 0x5a, 0x5c, 0x1c, 0x04, 0x35, 0xc0, 0xc9, 0xff, 0xc4, 0x23, 0x39, 0xc6,
	0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39,
	0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x4c, 0xf1, 0x1a, 0x58, 0xa1, 0x0f, 0xf2, 0x9a, 0x81, 0xa5,
	0x2e, 0xc2, 0x77, 0x60, 0x3b, 0x92, 0xd8, 0xc0, 0xae, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x0c, 0x82, 0x3a, 0x52, 0xfb, 0x00, 0x00, 0x00,
}

func (m *MsgCreateClient) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateClient) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateClient) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintLocalhost(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLocalhost(dAtA []byte, offset int, v uint64) int {
	offset -= sovLocalhost(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateClient) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovLocalhost(uint64(l))
	}
	return n
}

func sovLocalhost(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLocalhost(x uint64) (n int) {
	return sovLocalhost(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateClient) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLocalhost
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
			return fmt.Errorf("proto: MsgCreateClient: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateClient: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocalhost
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
				return ErrInvalidLengthLocalhost
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLocalhost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = append(m.Signer[:0], dAtA[iNdEx:postIndex]...)
			if m.Signer == nil {
				m.Signer = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLocalhost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLocalhost
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLocalhost
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
func skipLocalhost(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLocalhost
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
					return 0, ErrIntOverflowLocalhost
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
					return 0, ErrIntOverflowLocalhost
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
				return 0, ErrInvalidLengthLocalhost
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLocalhost
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLocalhost
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLocalhost        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLocalhost          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLocalhost = fmt.Errorf("proto: unexpected end of group")
)
