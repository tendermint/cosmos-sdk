// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/bank/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// GenesisState defines the bank module's genesis state.
type GenesisState struct {
	// params defines all the paramaters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// balances is an array containing the balances of all the accounts.
	Balances []Balance `protobuf:"bytes,2,rep,name=balances,proto3" json:"balances"`
	// supply represents the total supply.
	Supply github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=supply,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"supply"`
	// denom_metadata defines the metadata of the differents coins.
	DenomMetadata []Metadata `protobuf:"bytes,4,rep,name=denom_metadata,json=denomMetadata,proto3" json:"denom_metadata" yaml:"denom_metadata"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f007de11b420c6e, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetBalances() []Balance {
	if m != nil {
		return m.Balances
	}
	return nil
}

func (m *GenesisState) GetSupply() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Supply
	}
	return nil
}

func (m *GenesisState) GetDenomMetadata() []Metadata {
	if m != nil {
		return m.DenomMetadata
	}
	return nil
}

// Balance defines an account address and balance pair used in the bank module's
// genesis state.
type Balance struct {
	// address is the address of the balance holder.
	Address github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty"`
	// coins defines the different coins this balance holds.
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *Balance) Reset()         { *m = Balance{} }
func (m *Balance) String() string { return proto.CompactTextString(m) }
func (*Balance) ProtoMessage()    {}
func (*Balance) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f007de11b420c6e, []int{1}
}
func (m *Balance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Balance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Balance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Balance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Balance.Merge(m, src)
}
func (m *Balance) XXX_Size() int {
	return m.Size()
}
func (m *Balance) XXX_DiscardUnknown() {
	xxx_messageInfo_Balance.DiscardUnknown(m)
}

var xxx_messageInfo_Balance proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "cosmos.bank.v1beta1.GenesisState")
	proto.RegisterType((*Balance)(nil), "cosmos.bank.v1beta1.Balance")
}

func init() { proto.RegisterFile("cosmos/bank/v1beta1/genesis.proto", fileDescriptor_8f007de11b420c6e) }

var fileDescriptor_8f007de11b420c6e = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x4f, 0x4a, 0xcc, 0xcb, 0xd6, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4,
	0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0xce, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x86,
	0x28, 0xd1, 0x03, 0x29, 0xd1, 0x83, 0x2a, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xcb, 0xeb,
	0x83, 0x58, 0x10, 0xa5, 0x52, 0x72, 0x70, 0xd3, 0x8a, 0x53, 0xe1, 0xa6, 0x25, 0xe7, 0x67, 0xe6,
	0x61, 0xc8, 0x23, 0xd9, 0x06, 0x36, 0x17, 0x2c, 0xaf, 0x74, 0x99, 0x89, 0x8b, 0xc7, 0x1d, 0x62,
	0x79, 0x70, 0x49, 0x62, 0x49, 0xaa, 0x90, 0x25, 0x17, 0x5b, 0x41, 0x62, 0x51, 0x62, 0x6e, 0xb1,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xb4, 0x1e, 0x16, 0xc7, 0xe8, 0x05, 0x80, 0x95, 0x38,
	0xb1, 0x9c, 0xb8, 0x27, 0xcf, 0x10, 0x04, 0xd5, 0x20, 0x64, 0xc7, 0xc5, 0x91, 0x94, 0x98, 0x93,
	0x98, 0x97, 0x9c, 0x5a, 0x2c, 0xc1, 0xa4, 0xc0, 0xac, 0xc1, 0x6d, 0x24, 0x83, 0x55, 0xb3, 0x13,
	0x44, 0x11, 0x54, 0x37, 0x5c, 0x8f, 0x50, 0x32, 0x17, 0x5b, 0x71, 0x69, 0x41, 0x41, 0x4e, 0xa5,
	0x04, 0x33, 0x58, 0xb7, 0x24, 0x42, 0x77, 0x71, 0x2a, 0x5c, 0xb7, 0x73, 0x7e, 0x66, 0x9e, 0x93,
	0x01, 0x48, 0xeb, 0xaa, 0xfb, 0xf2, 0x1a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9,
	0xb9, 0xfa, 0x50, 0x9f, 0x42, 0x28, 0xdd, 0xe2, 0x94, 0x6c, 0xfd, 0x92, 0xca, 0x82, 0xd4, 0x62,
	0xb0, 0x86, 0xe2, 0x20, 0xa8, 0xd1, 0x42, 0xc9, 0x5c, 0x7c, 0x29, 0xa9, 0x79, 0xf9, 0xb9, 0xf1,
	0xb9, 0xa9, 0x25, 0x89, 0x29, 0x89, 0x25, 0x89, 0x12, 0x2c, 0x60, 0xcb, 0x64, 0xb1, 0x3a, 0xd5,
	0x17, 0xaa, 0xc8, 0x49, 0x16, 0x64, 0xe1, 0xa7, 0x7b, 0xf2, 0xa2, 0x95, 0x89, 0xb9, 0x39, 0x56,
	0x4a, 0xa8, 0x46, 0x28, 0x05, 0xf1, 0x82, 0x05, 0x60, 0xaa, 0x95, 0x0e, 0x33, 0x72, 0xb1, 0x43,
	0x7d, 0x29, 0xe4, 0xcd, 0xc5, 0x9e, 0x98, 0x92, 0x52, 0x94, 0x5a, 0x0c, 0x09, 0x51, 0x4e, 0x27,
	0xc3, 0x5f, 0xf7, 0xe4, 0x75, 0x89, 0x70, 0xb7, 0x63, 0x72, 0xb2, 0x23, 0x44, 0x63, 0x10, 0xcc,
	0x04, 0xa1, 0x44, 0x2e, 0x56, 0x50, 0xe4, 0xc2, 0xc2, 0x97, 0xaa, 0x21, 0x04, 0x31, 0xd9, 0x8a,
	0xa3, 0x63, 0x81, 0x3c, 0xc3, 0x8b, 0x05, 0xf2, 0x0c, 0x4e, 0xce, 0x27, 0x1e, 0xc9, 0x31, 0x5e,
	0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31,
	0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x89, 0xd7, 0xd0, 0x0a, 0x48, 0x6a, 0x03, 0x9b, 0x9d, 0xc4,
	0x06, 0x4e, 0x67, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee, 0x90, 0x58, 0x4a, 0xf7, 0x02,
	0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DenomMetadata) > 0 {
		for iNdEx := len(m.DenomMetadata) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DenomMetadata[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Supply) > 0 {
		for iNdEx := len(m.Supply) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Supply[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Balances) > 0 {
		for iNdEx := len(m.Balances) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Balances[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Balance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Balance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Balance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Balances) > 0 {
		for _, e := range m.Balances {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Supply) > 0 {
		for _, e := range m.Supply {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DenomMetadata) > 0 {
		for _, e := range m.DenomMetadata {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Balance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Balances = append(m.Balances, Balance{})
			if err := m.Balances[len(m.Balances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supply", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Supply = append(m.Supply, types.Coin{})
			if err := m.Supply[len(m.Supply)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomMetadata = append(m.DenomMetadata, Metadata{})
			if err := m.DenomMetadata[len(m.DenomMetadata)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Balance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Balance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Balance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = github_com_cosmos_cosmos_sdk_types.AccAddress(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
