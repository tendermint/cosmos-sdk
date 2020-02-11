// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/slashing/internal/types/types.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgUnjail - struct for unjailing jailed validator
type MsgUnjail struct {
	ValidatorAddr github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,1,opt,name=validator_addr,json=validatorAddr,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"validator_addr,omitempty" yaml:"validator_address"`
}

func (m *MsgUnjail) Reset()         { *m = MsgUnjail{} }
func (m *MsgUnjail) String() string { return proto.CompactTextString(m) }
func (*MsgUnjail) ProtoMessage()    {}
func (*MsgUnjail) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b882c3b0cdd6f57, []int{0}
}
func (m *MsgUnjail) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnjail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnjail.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnjail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnjail.Merge(m, src)
}
func (m *MsgUnjail) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnjail) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnjail.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnjail proto.InternalMessageInfo

func (m *MsgUnjail) GetValidatorAddr() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.ValidatorAddr
	}
	return nil
}

// ValidatorSigningInfo defines the signing info for a validator
type ValidatorSigningInfo struct {
	Address             github_com_cosmos_cosmos_sdk_types.ConsAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ConsAddress" json:"address,omitempty" yaml:"address"`
	StartHeight         int64                                          `protobuf:"varint,2,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty" yaml:"start_height"`
	IndexOffset         int64                                          `protobuf:"varint,3,opt,name=index_offset,json=indexOffset,proto3" json:"index_offset,omitempty" yaml:"index_offset"`
	JailedUntil         time.Time                                      `protobuf:"bytes,4,opt,name=jailed_until,json=jailedUntil,proto3,stdtime" json:"jailed_until"`
	Tombstoned          bool                                           `protobuf:"varint,5,opt,name=tombstoned,proto3" json:"tombstoned,omitempty" yaml:"tombstoned"`
	MissedBlocksCounter int64                                          `protobuf:"varint,6,opt,name=missed_blocks_counter,json=missedBlocksCounter,proto3" json:"missed_blocks_counter,omitempty" yaml:"missed_blocks_counter"`
}

func (m *ValidatorSigningInfo) Reset()      { *m = ValidatorSigningInfo{} }
func (*ValidatorSigningInfo) ProtoMessage() {}
func (*ValidatorSigningInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b882c3b0cdd6f57, []int{1}
}
func (m *ValidatorSigningInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSigningInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSigningInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSigningInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSigningInfo.Merge(m, src)
}
func (m *ValidatorSigningInfo) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSigningInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSigningInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSigningInfo proto.InternalMessageInfo

func (m *ValidatorSigningInfo) GetAddress() github_com_cosmos_cosmos_sdk_types.ConsAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ValidatorSigningInfo) GetStartHeight() int64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *ValidatorSigningInfo) GetIndexOffset() int64 {
	if m != nil {
		return m.IndexOffset
	}
	return 0
}

func (m *ValidatorSigningInfo) GetJailedUntil() time.Time {
	if m != nil {
		return m.JailedUntil
	}
	return time.Time{}
}

func (m *ValidatorSigningInfo) GetTombstoned() bool {
	if m != nil {
		return m.Tombstoned
	}
	return false
}

func (m *ValidatorSigningInfo) GetMissedBlocksCounter() int64 {
	if m != nil {
		return m.MissedBlocksCounter
	}
	return 0
}

// QuerySigningInfoParams defines the params for the following queries:
// - 'custom/slashing/signingInfo'
type QuerySigningInfoParams struct {
	ConsAddress github_com_cosmos_cosmos_sdk_types.ConsAddress `protobuf:"bytes,1,opt,name=cons_address,json=consAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ConsAddress" json:"cons_address,omitempty" yaml:"cons_address"`
}

func (m *QuerySigningInfoParams) Reset()         { *m = QuerySigningInfoParams{} }
func (m *QuerySigningInfoParams) String() string { return proto.CompactTextString(m) }
func (*QuerySigningInfoParams) ProtoMessage()    {}
func (*QuerySigningInfoParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b882c3b0cdd6f57, []int{2}
}
func (m *QuerySigningInfoParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySigningInfoParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySigningInfoParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySigningInfoParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySigningInfoParams.Merge(m, src)
}
func (m *QuerySigningInfoParams) XXX_Size() int {
	return m.Size()
}
func (m *QuerySigningInfoParams) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySigningInfoParams.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySigningInfoParams proto.InternalMessageInfo

func (m *QuerySigningInfoParams) GetConsAddress() github_com_cosmos_cosmos_sdk_types.ConsAddress {
	if m != nil {
		return m.ConsAddress
	}
	return nil
}

// QuerySigningInfosParams defines the params for the following queries:
// - 'custom/slashing/signingInfos'
type QuerySigningInfosParams struct {
	Page  int32 `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	Limit int32 `protobuf:"varint,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (m *QuerySigningInfosParams) Reset()         { *m = QuerySigningInfosParams{} }
func (m *QuerySigningInfosParams) String() string { return proto.CompactTextString(m) }
func (*QuerySigningInfosParams) ProtoMessage()    {}
func (*QuerySigningInfosParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b882c3b0cdd6f57, []int{3}
}
func (m *QuerySigningInfosParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySigningInfosParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySigningInfosParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySigningInfosParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySigningInfosParams.Merge(m, src)
}
func (m *QuerySigningInfosParams) XXX_Size() int {
	return m.Size()
}
func (m *QuerySigningInfosParams) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySigningInfosParams.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySigningInfosParams proto.InternalMessageInfo

func (m *QuerySigningInfosParams) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *QuerySigningInfosParams) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgUnjail)(nil), "cosmos_sdk.x.slashing.v1.MsgUnjail")
	proto.RegisterType((*ValidatorSigningInfo)(nil), "cosmos_sdk.x.slashing.v1.ValidatorSigningInfo")
	proto.RegisterType((*QuerySigningInfoParams)(nil), "cosmos_sdk.x.slashing.v1.QuerySigningInfoParams")
	proto.RegisterType((*QuerySigningInfosParams)(nil), "cosmos_sdk.x.slashing.v1.QuerySigningInfosParams")
}

func init() {
	proto.RegisterFile("x/slashing/internal/types/types.proto", fileDescriptor_2b882c3b0cdd6f57)
}

var fileDescriptor_2b882c3b0cdd6f57 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x31, 0x6f, 0xd3, 0x40,
	0x18, 0xcd, 0xd1, 0xa6, 0x94, 0x4b, 0xa8, 0x84, 0xdb, 0x52, 0x2b, 0x42, 0xbe, 0xc8, 0x12, 0x28,
	0x4b, 0x6d, 0x01, 0x62, 0xc9, 0x86, 0x83, 0x04, 0x95, 0xa8, 0x28, 0xa6, 0xed, 0xc0, 0x80, 0x75,
	0xb1, 0x2f, 0xce, 0x11, 0xfb, 0x2e, 0xf2, 0x5d, 0xaa, 0x64, 0x64, 0x61, 0xee, 0xc8, 0xd8, 0x9f,
	0xd3, 0x09, 0x75, 0x64, 0x32, 0x28, 0x59, 0x98, 0x33, 0x32, 0xa1, 0xdc, 0xc5, 0xc4, 0xad, 0x18,
	0x10, 0x4b, 0xe2, 0xf7, 0x7d, 0xdf, 0xfb, 0xde, 0xbb, 0x7b, 0x07, 0x1f, 0x8e, 0x5d, 0x91, 0x60,
	0xd1, 0xa7, 0x2c, 0x76, 0x29, 0x93, 0x24, 0x63, 0x38, 0x71, 0xe5, 0x64, 0x48, 0x84, 0xfe, 0x75,
	0x86, 0x19, 0x97, 0xdc, 0x30, 0x43, 0x2e, 0x52, 0x2e, 0x02, 0x11, 0x0d, 0x9c, 0xb1, 0x53, 0x30,
	0x9c, 0xb3, 0xc7, 0x8d, 0x47, 0xb2, 0x4f, 0xb3, 0x28, 0x18, 0xe2, 0x4c, 0x4e, 0x5c, 0x35, 0xec,
	0xc6, 0x3c, 0xe6, 0xab, 0x2f, 0xbd, 0xa1, 0x81, 0x62, 0xce, 0xe3, 0x84, 0xe8, 0x91, 0xee, 0xa8,
	0xe7, 0x4a, 0x9a, 0x12, 0x21, 0x71, 0x3a, 0xd4, 0x03, 0xf6, 0x27, 0x00, 0xef, 0x1c, 0x8a, 0xf8,
	0x84, 0x7d, 0xc4, 0x34, 0x31, 0x24, 0xdc, 0x3a, 0xc3, 0x09, 0x8d, 0xb0, 0xe4, 0x59, 0x80, 0xa3,
	0x28, 0x33, 0x41, 0x13, 0xb4, 0xea, 0xde, 0xe1, 0x3c, 0x47, 0xe6, 0x04, 0xa7, 0x49, 0xdb, 0xbe,
	0xde, 0x27, 0x42, 0xd8, 0xbf, 0x72, 0xb4, 0x1f, 0x53, 0xd9, 0x1f, 0x75, 0x9d, 0x90, 0xa7, 0xae,
	0xf6, 0xbc, 0xfc, 0xdb, 0x17, 0xd1, 0x60, 0x79, 0xa4, 0x53, 0x9c, 0x3c, 0xd7, 0x0c, 0xff, 0xee,
	0x9f, 0x25, 0x8b, 0x8a, 0xfd, 0x75, 0x0d, 0xee, 0x9c, 0x16, 0x95, 0x77, 0x34, 0x66, 0x94, 0xc5,
	0x07, 0xac, 0xc7, 0x8d, 0x0f, 0xf0, 0xf6, 0x52, 0x64, 0xe9, 0xe3, 0xc5, 0x3c, 0x47, 0x5b, 0xda,
	0x47, 0x49, 0xdd, 0xf9, 0x07, 0xf5, 0x0e, 0x67, 0xa2, 0x90, 0x2f, 0x96, 0x1a, 0x6d, 0x58, 0x17,
	0x12, 0x67, 0x32, 0xe8, 0x13, 0x1a, 0xf7, 0xa5, 0x79, 0xab, 0x09, 0x5a, 0x6b, 0xde, 0xde, 0x3c,
	0x47, 0xdb, 0x5a, 0xa4, 0xdc, 0xb5, 0xfd, 0x9a, 0x82, 0xaf, 0x14, 0x5a, 0x70, 0x29, 0x8b, 0xc8,
	0x38, 0xe0, 0xbd, 0x9e, 0x20, 0xd2, 0x5c, 0xbb, 0xc9, 0x2d, 0x77, 0x6d, 0xbf, 0xa6, 0xe0, 0x1b,
	0x85, 0x8c, 0x97, 0xb0, 0xbe, 0xb8, 0x6e, 0x12, 0x05, 0x23, 0x26, 0x69, 0x62, 0xae, 0x37, 0x41,
	0xab, 0xf6, 0xa4, 0xe1, 0xe8, 0xb0, 0x9c, 0x22, 0x2c, 0xe7, 0xb8, 0x08, 0xcb, 0xdb, 0xbc, 0xcc,
	0x51, 0xe5, 0xfc, 0x3b, 0x02, 0x7e, 0x4d, 0x33, 0x4f, 0x16, 0x44, 0xe3, 0x19, 0x84, 0x92, 0xa7,
	0x5d, 0x21, 0x39, 0x23, 0x91, 0x59, 0x6d, 0x82, 0xd6, 0xa6, 0xb7, 0x3b, 0xcf, 0xd1, 0x3d, 0x6d,
	0x61, 0xd5, 0xb3, 0xfd, 0xd2, 0xa0, 0x71, 0x0c, 0x77, 0x53, 0x2a, 0x04, 0x89, 0x82, 0x6e, 0xc2,
	0xc3, 0x81, 0x08, 0x42, 0x3e, 0x5a, 0xbc, 0x42, 0x73, 0x43, 0x1d, 0xa2, 0x39, 0xcf, 0xd1, 0x03,
	0xbd, 0xe1, 0xaf, 0x63, 0xb6, 0xbf, 0xad, 0xeb, 0x9e, 0x2a, 0x77, 0x74, 0xb5, 0xbd, 0xf9, 0xe5,
	0x02, 0x55, 0x7e, 0x5e, 0x20, 0x60, 0x7f, 0x06, 0xf0, 0xfe, 0xdb, 0x11, 0xc9, 0x26, 0xa5, 0x30,
	0x8f, 0x70, 0x86, 0x53, 0x61, 0x24, 0xb0, 0x1e, 0x72, 0x26, 0x82, 0xeb, 0xb9, 0x1e, 0xac, 0xae,
	0xad, 0xdc, 0xfd, 0x9f, 0x70, 0x6b, 0xe1, 0x0a, 0xd8, 0x1d, 0xb8, 0x77, 0xd3, 0x87, 0x58, 0x1a,
	0x31, 0xe0, 0xfa, 0x11, 0x8e, 0x89, 0x32, 0x50, 0xf5, 0xd5, 0xb7, 0xb1, 0x03, 0xab, 0xaf, 0x69,
	0x4a, 0xf5, 0x43, 0xa8, 0xfa, 0x1a, 0x78, 0xe8, 0x72, 0x6a, 0x81, 0xab, 0xa9, 0x05, 0x7e, 0x4c,
	0x2d, 0x70, 0x3e, 0xb3, 0x2a, 0x57, 0x33, 0xab, 0xf2, 0x6d, 0x66, 0x55, 0xde, 0x57, 0x95, 0x7c,
	0x77, 0x43, 0x05, 0xf6, 0xf4, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x12, 0xe8, 0xb0, 0x88, 0xd6,
	0x03, 0x00, 0x00,
}

func (this *ValidatorSigningInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ValidatorSigningInfo)
	if !ok {
		that2, ok := that.(ValidatorSigningInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Address, that1.Address) {
		return false
	}
	if this.StartHeight != that1.StartHeight {
		return false
	}
	if this.IndexOffset != that1.IndexOffset {
		return false
	}
	if !this.JailedUntil.Equal(that1.JailedUntil) {
		return false
	}
	if this.Tombstoned != that1.Tombstoned {
		return false
	}
	if this.MissedBlocksCounter != that1.MissedBlocksCounter {
		return false
	}
	return true
}
func (m *MsgUnjail) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnjail) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnjail) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddr) > 0 {
		i -= len(m.ValidatorAddr)
		copy(dAtA[i:], m.ValidatorAddr)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ValidatorAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorSigningInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSigningInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSigningInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MissedBlocksCounter != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MissedBlocksCounter))
		i--
		dAtA[i] = 0x30
	}
	if m.Tombstoned {
		i--
		if m.Tombstoned {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.JailedUntil, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.JailedUntil):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTypes(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if m.IndexOffset != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.IndexOffset))
		i--
		dAtA[i] = 0x18
	}
	if m.StartHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.StartHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySigningInfoParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySigningInfoParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySigningInfoParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ConsAddress) > 0 {
		i -= len(m.ConsAddress)
		copy(dAtA[i:], m.ConsAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ConsAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySigningInfosParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySigningInfosParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySigningInfosParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Limit != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x10
	}
	if m.Page != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Page))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUnjail) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddr)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ValidatorSigningInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.StartHeight != 0 {
		n += 1 + sovTypes(uint64(m.StartHeight))
	}
	if m.IndexOffset != 0 {
		n += 1 + sovTypes(uint64(m.IndexOffset))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.JailedUntil)
	n += 1 + l + sovTypes(uint64(l))
	if m.Tombstoned {
		n += 2
	}
	if m.MissedBlocksCounter != 0 {
		n += 1 + sovTypes(uint64(m.MissedBlocksCounter))
	}
	return n
}

func (m *QuerySigningInfoParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ConsAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *QuerySigningInfosParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Page != 0 {
		n += 1 + sovTypes(uint64(m.Page))
	}
	if m.Limit != 0 {
		n += 1 + sovTypes(uint64(m.Limit))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUnjail) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: MsgUnjail: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnjail: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddr = append(m.ValidatorAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.ValidatorAddr == nil {
				m.ValidatorAddr = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ValidatorSigningInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ValidatorSigningInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSigningInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartHeight", wireType)
			}
			m.StartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexOffset", wireType)
			}
			m.IndexOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IndexOffset |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JailedUntil", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.JailedUntil, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tombstoned", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Tombstoned = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissedBlocksCounter", wireType)
			}
			m.MissedBlocksCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MissedBlocksCounter |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *QuerySigningInfoParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: QuerySigningInfoParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySigningInfoParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsAddress = append(m.ConsAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.ConsAddress == nil {
				m.ConsAddress = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *QuerySigningInfosParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: QuerySigningInfosParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySigningInfosParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Page", wireType)
			}
			m.Page = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Page |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
