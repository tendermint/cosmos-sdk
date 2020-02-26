// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/evidence/types/types.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"

	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// MsgSubmitEvidenceBase defines an sdk.Msg type that supports submitting arbitrary
// Evidence.
//
// Note, this message type provides the basis for which a true MsgSubmitEvidence
// can be constructed. Since the evidence submitted in the message can be arbitrary,
// assuming it fulfills the Evidence interface, it must be defined at the
// application-level and extend MsgSubmitEvidenceBase.
type MsgSubmitEvidenceBase struct {
	Submitter github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=submitter,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"submitter,omitempty"`
}

func (m *MsgSubmitEvidenceBase) Reset()         { *m = MsgSubmitEvidenceBase{} }
func (m *MsgSubmitEvidenceBase) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitEvidenceBase) ProtoMessage()    {}
func (*MsgSubmitEvidenceBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_72113e6a7b2536ae, []int{0}
}
func (m *MsgSubmitEvidenceBase) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitEvidenceBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitEvidenceBase.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitEvidenceBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitEvidenceBase.Merge(m, src)
}
func (m *MsgSubmitEvidenceBase) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitEvidenceBase) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitEvidenceBase.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitEvidenceBase proto.InternalMessageInfo

func (m *MsgSubmitEvidenceBase) GetSubmitter() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Submitter
	}
	return nil
}

// Equivocation implements the Evidence interface and defines evidence of double
// signing misbehavior.
type Equivocation struct {
	Height           int64                                          `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Time             time.Time                                      `protobuf:"bytes,2,opt,name=time,proto3,stdtime" json:"time"`
	Power            int64                                          `protobuf:"varint,3,opt,name=power,proto3" json:"power,omitempty"`
	ConsensusAddress github_com_cosmos_cosmos_sdk_types.ConsAddress `protobuf:"bytes,4,opt,name=consensus_address,json=consensusAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ConsAddress" json:"consensus_address,omitempty" yaml:"consensus_address"`
}

func (m *Equivocation) Reset()      { *m = Equivocation{} }
func (*Equivocation) ProtoMessage() {}
func (*Equivocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_72113e6a7b2536ae, []int{1}
}
func (m *Equivocation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Equivocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Equivocation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Equivocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Equivocation.Merge(m, src)
}
func (m *Equivocation) XXX_Size() int {
	return m.Size()
}
func (m *Equivocation) XXX_DiscardUnknown() {
	xxx_messageInfo_Equivocation.DiscardUnknown(m)
}

var xxx_messageInfo_Equivocation proto.InternalMessageInfo

// Params defines the total set of parameters for the evidence module
type Params struct {
	MaxEvidenceAge time.Duration `protobuf:"bytes,1,opt,name=max_evidence_age,json=maxEvidenceAge,proto3,stdduration" json:"max_evidence_age" yaml:"max_evidence_age"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_72113e6a7b2536ae, []int{2}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMaxEvidenceAge() time.Duration {
	if m != nil {
		return m.MaxEvidenceAge
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgSubmitEvidenceBase)(nil), "cosmos_sdk.x.evidence.v1.MsgSubmitEvidenceBase")
	proto.RegisterType((*Equivocation)(nil), "cosmos_sdk.x.evidence.v1.Equivocation")
	proto.RegisterType((*Params)(nil), "cosmos_sdk.x.evidence.v1.Params")
}

func init() { proto.RegisterFile("x/evidence/types/types.proto", fileDescriptor_72113e6a7b2536ae) }

var fileDescriptor_72113e6a7b2536ae = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x3d, 0x6f, 0xd4, 0x30,
	0x1c, 0xc6, 0x63, 0x7a, 0x9c, 0x8a, 0x5b, 0xa1, 0x12, 0xf1, 0x12, 0x4e, 0x28, 0xae, 0x82, 0x84,
	0xba, 0xd4, 0x51, 0xcb, 0x82, 0x6e, 0xbb, 0x40, 0x27, 0xc4, 0x8b, 0x0e, 0x26, 0x96, 0xc8, 0x97,
	0x18, 0x27, 0x6a, 0x1d, 0x07, 0xdb, 0x39, 0x72, 0xe2, 0x0b, 0x30, 0x76, 0xec, 0xd8, 0x91, 0x8f,
	0xd2, 0xb1, 0x23, 0x53, 0x40, 0x77, 0xdf, 0xa0, 0x63, 0x25, 0x24, 0x74, 0x76, 0x42, 0xa5, 0xab,
	0x84, 0xba, 0x24, 0xb1, 0xf3, 0xe4, 0xc9, 0xf3, 0xfb, 0x3f, 0x86, 0x4f, 0xea, 0x90, 0x4e, 0xf3,
	0x94, 0x16, 0x09, 0x0d, 0xf5, 0xac, 0xa4, 0xca, 0x5e, 0x71, 0x29, 0x85, 0x16, 0xae, 0x97, 0x08,
	0xc5, 0x85, 0x8a, 0x55, 0x7a, 0x88, 0x6b, 0xdc, 0x09, 0xf1, 0x74, 0x6f, 0xf0, 0x4c, 0x67, 0xb9,
	0x4c, 0xe3, 0x92, 0x48, 0x3d, 0x0b, 0x8d, 0x38, 0x64, 0x82, 0x89, 0xab, 0x27, 0xeb, 0x30, 0x40,
	0x4c, 0x08, 0x76, 0x44, 0xad, 0x64, 0x52, 0x7d, 0x0e, 0x75, 0xce, 0xa9, 0xd2, 0x84, 0x97, 0xad,
	0xc0, 0x5f, 0x15, 0xa4, 0x95, 0x24, 0x3a, 0x17, 0x85, 0x7d, 0x1f, 0x64, 0xf0, 0xc1, 0x1b, 0xc5,
	0x3e, 0x54, 0x13, 0x9e, 0xeb, 0x83, 0x36, 0x40, 0x44, 0x14, 0x75, 0xdf, 0xc1, 0x3b, 0xca, 0xec,
	0x6a, 0x2a, 0x3d, 0xb0, 0x0d, 0x76, 0x36, 0xa3, 0xbd, 0xcb, 0x06, 0xed, 0xb2, 0x5c, 0x67, 0xd5,
	0x04, 0x27, 0x82, 0x87, 0x36, 0x7d, 0x7b, 0xdb, 0x55, 0xe9, 0x61, 0x0b, 0x37, 0x4a, 0x92, 0x51,
	0x9a, 0x4a, 0xaa, 0xd4, 0xf8, 0xca, 0x23, 0xf8, 0x03, 0xe0, 0xe6, 0xc1, 0x97, 0x2a, 0x9f, 0x8a,
	0xc4, 0x04, 0x70, 0x1f, 0xc2, 0x7e, 0x46, 0x73, 0x96, 0x69, 0x63, 0xbf, 0x36, 0x6e, 0x57, 0xee,
	0x0b, 0xd8, 0x5b, 0x52, 0x78, 0xb7, 0xb6, 0xc1, 0xce, 0xc6, 0xfe, 0x00, 0x5b, 0x02, 0xdc, 0x11,
	0xe0, 0x8f, 0x1d, 0x62, 0xb4, 0x7e, 0xd6, 0x20, 0xe7, 0xf8, 0x17, 0x02, 0x63, 0xf3, 0x85, 0x7b,
	0x1f, 0xde, 0x2e, 0xc5, 0x57, 0x2a, 0xbd, 0x35, 0x63, 0x68, 0x17, 0xee, 0x37, 0x78, 0x2f, 0x11,
	0x85, 0xa2, 0x85, 0xaa, 0x54, 0x4c, 0x6c, 0x30, 0xaf, 0x67, 0x88, 0xde, 0x5e, 0x34, 0xc8, 0x9b,
	0x11, 0x7e, 0x34, 0x0c, 0xae, 0x49, 0x82, 0xcb, 0x06, 0xe1, 0x1b, 0xd0, 0xbe, 0x14, 0x85, 0xea,
	0x70, 0xb7, 0xfe, 0xb9, 0xb4, 0x3b, 0xc3, 0xf5, 0xef, 0xa7, 0xc8, 0x39, 0x39, 0x45, 0x4e, 0x50,
	0xc3, 0xfe, 0x7b, 0x22, 0x09, 0x57, 0x6e, 0x06, 0xb7, 0x38, 0xa9, 0xe3, 0xae, 0xef, 0x98, 0x30,
	0x6a, 0x46, 0xb0, 0xb1, 0xff, 0xf8, 0x1a, 0xec, 0xab, 0xb6, 0xae, 0xe8, 0xe9, 0x92, 0xf5, 0xa2,
	0x41, 0x8f, 0x6c, 0xdc, 0x55, 0x83, 0xe0, 0x64, 0x39, 0x86, 0xbb, 0x9c, 0xd4, 0x5d, 0x8b, 0x23,
	0x46, 0x87, 0xbd, 0xe5, 0x9f, 0xa3, 0xd7, 0x3f, 0xe6, 0x3e, 0x38, 0x9b, 0xfb, 0xe0, 0x7c, 0xee,
	0x83, 0xdf, 0x73, 0x1f, 0x1c, 0x2f, 0x7c, 0xe7, 0x7c, 0xe1, 0x3b, 0x3f, 0x17, 0xbe, 0xf3, 0xe9,
	0xff, 0x8d, 0xae, 0x9e, 0xdf, 0x49, 0xdf, 0x44, 0x7b, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x35,
	0xfb, 0xe6, 0x44, 0xda, 0x02, 0x00, 0x00,
}

func (this *MsgSubmitEvidenceBase) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgSubmitEvidenceBase)
	if !ok {
		that2, ok := that.(MsgSubmitEvidenceBase)
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
	if !bytes.Equal(this.Submitter, that1.Submitter) {
		return false
	}
	return true
}
func (this *Equivocation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Equivocation)
	if !ok {
		that2, ok := that.(Equivocation)
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
	if this.Height != that1.Height {
		return false
	}
	if !this.Time.Equal(that1.Time) {
		return false
	}
	if this.Power != that1.Power {
		return false
	}
	if !bytes.Equal(this.ConsensusAddress, that1.ConsensusAddress) {
		return false
	}
	return true
}
func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if this.MaxEvidenceAge != that1.MaxEvidenceAge {
		return false
	}
	return true
}
func (m *MsgSubmitEvidenceBase) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitEvidenceBase) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitEvidenceBase) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Submitter) > 0 {
		i -= len(m.Submitter)
		copy(dAtA[i:], m.Submitter)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Submitter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Equivocation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Equivocation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Equivocation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ConsensusAddress) > 0 {
		i -= len(m.ConsensusAddress)
		copy(dAtA[i:], m.ConsensusAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ConsensusAddress)))
		i--
		dAtA[i] = 0x22
	}
	if m.Power != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x18
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTypes(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.MaxEvidenceAge, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.MaxEvidenceAge):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintTypes(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0xa
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
func (m *MsgSubmitEvidenceBase) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Submitter)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Equivocation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovTypes(uint64(l))
	if m.Power != 0 {
		n += 1 + sovTypes(uint64(m.Power))
	}
	l = len(m.ConsensusAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.MaxEvidenceAge)
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSubmitEvidenceBase) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSubmitEvidenceBase: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitEvidenceBase: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Submitter", wireType)
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
			m.Submitter = append(m.Submitter[:0], dAtA[iNdEx:postIndex]...)
			if m.Submitter == nil {
				m.Submitter = []byte{}
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
func (m *Equivocation) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Equivocation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Equivocation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusAddress", wireType)
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
			m.ConsensusAddress = append(m.ConsensusAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.ConsensusAddress == nil {
				m.ConsensusAddress = []byte{}
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
func (m *Params) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxEvidenceAge", wireType)
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
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.MaxEvidenceAge, dAtA[iNdEx:postIndex]); err != nil {
				return err
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
