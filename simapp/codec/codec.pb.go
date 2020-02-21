// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: simapp/codec/codec.proto

package codec

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_x_auth_exported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	types "github.com/cosmos/cosmos-sdk/x/auth/types"
	types1 "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	github_com_cosmos_cosmos_sdk_x_evidence_exported "github.com/cosmos/cosmos-sdk/x/evidence/exported"
	types3 "github.com/cosmos/cosmos-sdk/x/evidence/types"
	github_com_cosmos_cosmos_sdk_x_supply_exported "github.com/cosmos/cosmos-sdk/x/supply/exported"
	types2 "github.com/cosmos/cosmos-sdk/x/supply/types"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// Account defines the application-level Account type.
type Account struct {
	// sum defines a list of all acceptable concrete Account implementations.
	//
	// Types that are valid to be assigned to Sum:
	//	*Account_BaseAccount
	//	*Account_ContinuousVestingAccount
	//	*Account_DelayedVestingAccount
	//	*Account_PeriodicVestingAccount
	//	*Account_ModuleAccount
	Sum isAccount_Sum `protobuf_oneof:"sum"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c6d4085e4065f5a, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Account.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return m.Size()
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

type isAccount_Sum interface {
	isAccount_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Account_BaseAccount struct {
	BaseAccount *types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,oneof" json:"base_account,omitempty"`
}
type Account_ContinuousVestingAccount struct {
	ContinuousVestingAccount *types1.ContinuousVestingAccount `protobuf:"bytes,2,opt,name=continuous_vesting_account,json=continuousVestingAccount,proto3,oneof" json:"continuous_vesting_account,omitempty"`
}
type Account_DelayedVestingAccount struct {
	DelayedVestingAccount *types1.DelayedVestingAccount `protobuf:"bytes,3,opt,name=delayed_vesting_account,json=delayedVestingAccount,proto3,oneof" json:"delayed_vesting_account,omitempty"`
}
type Account_PeriodicVestingAccount struct {
	PeriodicVestingAccount *types1.PeriodicVestingAccount `protobuf:"bytes,4,opt,name=periodic_vesting_account,json=periodicVestingAccount,proto3,oneof" json:"periodic_vesting_account,omitempty"`
}
type Account_ModuleAccount struct {
	ModuleAccount *types2.ModuleAccount `protobuf:"bytes,5,opt,name=module_account,json=moduleAccount,proto3,oneof" json:"module_account,omitempty"`
}

func (*Account_BaseAccount) isAccount_Sum()              {}
func (*Account_ContinuousVestingAccount) isAccount_Sum() {}
func (*Account_DelayedVestingAccount) isAccount_Sum()    {}
func (*Account_PeriodicVestingAccount) isAccount_Sum()   {}
func (*Account_ModuleAccount) isAccount_Sum()            {}

func (m *Account) GetSum() isAccount_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Account) GetBaseAccount() *types.BaseAccount {
	if x, ok := m.GetSum().(*Account_BaseAccount); ok {
		return x.BaseAccount
	}
	return nil
}

func (m *Account) GetContinuousVestingAccount() *types1.ContinuousVestingAccount {
	if x, ok := m.GetSum().(*Account_ContinuousVestingAccount); ok {
		return x.ContinuousVestingAccount
	}
	return nil
}

func (m *Account) GetDelayedVestingAccount() *types1.DelayedVestingAccount {
	if x, ok := m.GetSum().(*Account_DelayedVestingAccount); ok {
		return x.DelayedVestingAccount
	}
	return nil
}

func (m *Account) GetPeriodicVestingAccount() *types1.PeriodicVestingAccount {
	if x, ok := m.GetSum().(*Account_PeriodicVestingAccount); ok {
		return x.PeriodicVestingAccount
	}
	return nil
}

func (m *Account) GetModuleAccount() *types2.ModuleAccount {
	if x, ok := m.GetSum().(*Account_ModuleAccount); ok {
		return x.ModuleAccount
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Account) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Account_BaseAccount)(nil),
		(*Account_ContinuousVestingAccount)(nil),
		(*Account_DelayedVestingAccount)(nil),
		(*Account_PeriodicVestingAccount)(nil),
		(*Account_ModuleAccount)(nil),
	}
}

// Supply defines the application-level Supply type.
type Supply struct {
	// sum defines a list of all acceptable concrete Supply implementations.
	//
	// Types that are valid to be assigned to Sum:
	//	*Supply_Supply
	Sum isSupply_Sum `protobuf_oneof:"sum"`
}

func (m *Supply) Reset()         { *m = Supply{} }
func (m *Supply) String() string { return proto.CompactTextString(m) }
func (*Supply) ProtoMessage()    {}
func (*Supply) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c6d4085e4065f5a, []int{1}
}
func (m *Supply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Supply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Supply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Supply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Supply.Merge(m, src)
}
func (m *Supply) XXX_Size() int {
	return m.Size()
}
func (m *Supply) XXX_DiscardUnknown() {
	xxx_messageInfo_Supply.DiscardUnknown(m)
}

var xxx_messageInfo_Supply proto.InternalMessageInfo

type isSupply_Sum interface {
	isSupply_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Supply_Supply struct {
	Supply *types2.Supply `protobuf:"bytes,1,opt,name=supply,proto3,oneof" json:"supply,omitempty"`
}

func (*Supply_Supply) isSupply_Sum() {}

func (m *Supply) GetSum() isSupply_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Supply) GetSupply() *types2.Supply {
	if x, ok := m.GetSum().(*Supply_Supply); ok {
		return x.Supply
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Supply) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Supply_Supply)(nil),
	}
}

type Evidence struct {
	// Types that are valid to be assigned to Sum:
	//	*Evidence_Equivocation
	Sum isEvidence_Sum `protobuf_oneof:"sum"`
}

func (m *Evidence) Reset()         { *m = Evidence{} }
func (m *Evidence) String() string { return proto.CompactTextString(m) }
func (*Evidence) ProtoMessage()    {}
func (*Evidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c6d4085e4065f5a, []int{2}
}
func (m *Evidence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Evidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Evidence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Evidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Evidence.Merge(m, src)
}
func (m *Evidence) XXX_Size() int {
	return m.Size()
}
func (m *Evidence) XXX_DiscardUnknown() {
	xxx_messageInfo_Evidence.DiscardUnknown(m)
}

var xxx_messageInfo_Evidence proto.InternalMessageInfo

type isEvidence_Sum interface {
	isEvidence_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Evidence_Equivocation struct {
	Equivocation *types3.Equivocation `protobuf:"bytes,1,opt,name=equivocation,proto3,oneof" json:"equivocation,omitempty"`
}

func (*Evidence_Equivocation) isEvidence_Sum() {}

func (m *Evidence) GetSum() isEvidence_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Evidence) GetEquivocation() *types3.Equivocation {
	if x, ok := m.GetSum().(*Evidence_Equivocation); ok {
		return x.Equivocation
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Evidence) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Evidence_Equivocation)(nil),
	}
}

func init() {
	proto.RegisterType((*Account)(nil), "cosmos_sdk.simapp.codec.v1.Account")
	proto.RegisterType((*Supply)(nil), "cosmos_sdk.simapp.codec.v1.Supply")
	proto.RegisterType((*Evidence)(nil), "cosmos_sdk.simapp.codec.v1.Evidence")
}

func init() { proto.RegisterFile("simapp/codec/codec.proto", fileDescriptor_3c6d4085e4065f5a) }

var fileDescriptor_3c6d4085e4065f5a = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x4d, 0x8b, 0xd3, 0x40,
	0x1c, 0xc6, 0x13, 0xf7, 0x45, 0x99, 0x5d, 0x3d, 0x04, 0xd4, 0x10, 0x24, 0xac, 0x0b, 0x8a, 0x2f,
	0x74, 0xc2, 0xba, 0xbe, 0x76, 0x0f, 0x62, 0xd7, 0x95, 0x0a, 0x2a, 0x52, 0xc1, 0x83, 0x97, 0x92,
	0xce, 0x0c, 0xdb, 0x61, 0x9b, 0xcc, 0x98, 0x99, 0x09, 0xe9, 0x17, 0xf0, 0xec, 0x47, 0xf0, 0x43,
	0xf4, 0xe8, 0x07, 0x90, 0x9e, 0xf6, 0xe8, 0x51, 0xda, 0x2f, 0x22, 0xcd, 0x4c, 0xd2, 0x94, 0x64,
	0xbb, 0x97, 0xc0, 0xe4, 0x79, 0xfe, 0xcf, 0xef, 0x29, 0xfd, 0x4f, 0x80, 0x2b, 0x68, 0x14, 0x72,
	0x1e, 0x20, 0x86, 0x09, 0xd2, 0x4f, 0xc8, 0x13, 0x26, 0x99, 0xe3, 0x21, 0x26, 0x22, 0x26, 0xfa,
	0x02, 0x9f, 0x41, 0x6d, 0x82, 0x5a, 0x4e, 0x0f, 0xbc, 0xc7, 0x72, 0x48, 0x13, 0xdc, 0xe7, 0x61,
	0x22, 0xc7, 0x41, 0x6e, 0x0f, 0xb4, 0xbb, 0x55, 0x3d, 0xe8, 0x20, 0xcf, 0xcd, 0x82, 0x50, 0xc9,
	0x61, 0x20, 0xc7, 0x9c, 0x08, 0xfd, 0x34, 0xca, 0x9e, 0x51, 0x52, 0x22, 0x24, 0x8d, 0x4f, 0x1b,
	0x1c, 0x5e, 0x16, 0x08, 0xc5, 0xf9, 0x68, 0xdc, 0xa0, 0xdd, 0xc9, 0x02, 0x92, 0x52, 0x4c, 0x62,
	0x44, 0xea, 0xea, 0xfe, 0xef, 0x4d, 0x70, 0xf5, 0x0d, 0x42, 0x4c, 0xc5, 0xd2, 0x79, 0x07, 0x76,
	0x07, 0xa1, 0x20, 0xfd, 0x50, 0x9f, 0x5d, 0x7b, 0xcf, 0x7e, 0xb0, 0xf3, 0xe4, 0x2e, 0xac, 0xfc,
	0xc2, 0x0c, 0x2e, 0x9a, 0xc0, 0xf4, 0x00, 0x76, 0x42, 0x41, 0xcc, 0x60, 0xd7, 0xea, 0xed, 0x0c,
	0x96, 0x47, 0x27, 0x05, 0x1e, 0x62, 0xb1, 0xa4, 0xb1, 0x62, 0x4a, 0xf4, 0x4d, 0xeb, 0x32, 0xf5,
	0x4a, 0x9e, 0xfa, 0xbc, 0x29, 0x55, 0x3b, 0x17, 0xe9, 0xc7, 0xe5, 0xfc, 0x57, 0xfd, 0x72, 0x89,
	0x72, 0xd1, 0x05, 0x9a, 0x13, 0x81, 0xdb, 0x98, 0x8c, 0xc2, 0x31, 0xc1, 0x35, 0xe8, 0x46, 0x0e,
	0x3d, 0x5c, 0x0f, 0x7d, 0xab, 0x87, 0x6b, 0xc4, 0x9b, 0xb8, 0x49, 0x70, 0x38, 0x70, 0x39, 0x49,
	0x28, 0xc3, 0x14, 0xd5, 0x78, 0x9b, 0x39, 0xef, 0xe9, 0x7a, 0xde, 0x67, 0x33, 0x5d, 0x03, 0xde,
	0xe2, 0x8d, 0x8a, 0xf3, 0x09, 0xdc, 0x88, 0x18, 0x56, 0xa3, 0xe5, 0x5f, 0xb4, 0x95, 0x73, 0xee,
	0xad, 0x72, 0xf4, 0x2a, 0x2c, 0x08, 0x1f, 0x73, 0xf7, 0x32, 0xf8, 0x7a, 0x54, 0x7d, 0xd1, 0x7e,
	0x35, 0x9d, 0xb4, 0x9e, 0x3d, 0x3a, 0xa5, 0x72, 0xa8, 0x06, 0x10, 0xb1, 0xc8, 0x2c, 0x64, 0xb1,
	0xa4, 0x02, 0x9f, 0x05, 0x66, 0xf5, 0x48, 0xc6, 0x59, 0x22, 0x09, 0x86, 0x66, 0xb4, 0xb3, 0x05,
	0x36, 0x84, 0x8a, 0xf6, 0x7f, 0xd8, 0x60, 0xfb, 0x4b, 0x8e, 0x73, 0x5e, 0x82, 0x6d, 0x0d, 0x36,
	0x7b, 0xe3, 0x5f, 0x54, 0x4a, 0xfb, 0xbb, 0x56, 0xcf, 0xf8, 0xdb, 0x47, 0xd3, 0x49, 0xeb, 0xc5,
	0x65, 0x35, 0xcc, 0x7e, 0x97, 0x45, 0x74, 0xca, 0xfb, 0xa2, 0xc8, 0x2f, 0x1b, 0x5c, 0x3b, 0x31,
	0x6b, 0xee, 0x7c, 0x00, 0xbb, 0xe4, 0xbb, 0xa2, 0x29, 0x43, 0xa1, 0xa4, 0x2c, 0x36, 0x85, 0xee,
	0xaf, 0x16, 0x2a, 0x2e, 0xc5, 0xa2, 0xd2, 0x49, 0xc5, 0xdd, 0xb5, 0x7a, 0x2b, 0xd3, 0xed, 0xd7,
	0xd3, 0x49, 0xeb, 0xe8, 0xb2, 0x7a, 0xe5, 0x15, 0x2b, 0x0b, 0x16, 0x6d, 0x8a, 0x8a, 0x9d, 0xe3,
	0x3f, 0x33, 0xdf, 0x3e, 0x9f, 0xf9, 0xf6, 0xbf, 0x99, 0x6f, 0xff, 0x9c, 0xfb, 0xd6, 0xf9, 0xdc,
	0xb7, 0xfe, 0xce, 0x7d, 0xeb, 0xdb, 0xc3, 0xb5, 0xe1, 0xd5, 0x4f, 0xcf, 0x60, 0x3b, 0xbf, 0xb6,
	0x87, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xae, 0xe2, 0xbd, 0x10, 0x91, 0x04, 0x00, 0x00,
}

func (this *Account) GetAccount() github_com_cosmos_cosmos_sdk_x_auth_exported.Account {
	if x := this.GetBaseAccount(); x != nil {
		return x
	}
	if x := this.GetContinuousVestingAccount(); x != nil {
		return x
	}
	if x := this.GetDelayedVestingAccount(); x != nil {
		return x
	}
	if x := this.GetPeriodicVestingAccount(); x != nil {
		return x
	}
	if x := this.GetModuleAccount(); x != nil {
		return x
	}
	return nil
}

func (this *Account) SetAccount(value github_com_cosmos_cosmos_sdk_x_auth_exported.Account) error {
	if value == nil {
		this.Sum = nil
		return nil
	}
	switch vt := value.(type) {
	case *types.BaseAccount:
		this.Sum = &Account_BaseAccount{vt}
		return nil
	case *types1.ContinuousVestingAccount:
		this.Sum = &Account_ContinuousVestingAccount{vt}
		return nil
	case *types1.DelayedVestingAccount:
		this.Sum = &Account_DelayedVestingAccount{vt}
		return nil
	case *types1.PeriodicVestingAccount:
		this.Sum = &Account_PeriodicVestingAccount{vt}
		return nil
	case *types2.ModuleAccount:
		this.Sum = &Account_ModuleAccount{vt}
		return nil
	}
	return fmt.Errorf("can't encode value of type %T as message Account", value)
}

func (this *Supply) GetSupplyI() github_com_cosmos_cosmos_sdk_x_supply_exported.SupplyI {
	if x := this.GetSupply(); x != nil {
		return x
	}
	return nil
}

func (this *Supply) SetSupplyI(value github_com_cosmos_cosmos_sdk_x_supply_exported.SupplyI) error {
	if value == nil {
		this.Sum = nil
		return nil
	}
	switch vt := value.(type) {
	case *types2.Supply:
		this.Sum = &Supply_Supply{vt}
		return nil
	}
	return fmt.Errorf("can't encode value of type %T as message Supply", value)
}

func (this *Evidence) GetEvidenceI() github_com_cosmos_cosmos_sdk_x_evidence_exported.EvidenceI {
	if x := this.GetEquivocation(); x != nil {
		return x
	}
	return nil
}

func (this *Evidence) SetEvidenceI(value github_com_cosmos_cosmos_sdk_x_evidence_exported.EvidenceI) error {
	if value == nil {
		this.Sum = nil
		return nil
	}
	switch vt := value.(type) {
	case *types3.Equivocation:
		this.Sum = &Evidence_Equivocation{vt}
		return nil
	}
	return fmt.Errorf("can't encode value of type %T as message Evidence", value)
}

func (m *Account) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Account) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Account_BaseAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account_BaseAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BaseAccount != nil {
		{
			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCodec(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Account_ContinuousVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account_ContinuousVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ContinuousVestingAccount != nil {
		{
			size, err := m.ContinuousVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCodec(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Account_DelayedVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account_DelayedVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DelayedVestingAccount != nil {
		{
			size, err := m.DelayedVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCodec(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *Account_PeriodicVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account_PeriodicVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PeriodicVestingAccount != nil {
		{
			size, err := m.PeriodicVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCodec(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *Account_ModuleAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account_ModuleAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ModuleAccount != nil {
		{
			size, err := m.ModuleAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCodec(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *Supply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Supply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Supply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Supply_Supply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Supply_Supply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Supply != nil {
		{
			size, err := m.Supply.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCodec(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Evidence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Evidence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Evidence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Evidence_Equivocation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Evidence_Equivocation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Equivocation != nil {
		{
			size, err := m.Equivocation.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCodec(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func encodeVarintCodec(dAtA []byte, offset int, v uint64) int {
	offset -= sovCodec(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Account) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *Account_BaseAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Account_ContinuousVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ContinuousVestingAccount != nil {
		l = m.ContinuousVestingAccount.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Account_DelayedVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DelayedVestingAccount != nil {
		l = m.DelayedVestingAccount.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Account_PeriodicVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PeriodicVestingAccount != nil {
		l = m.PeriodicVestingAccount.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Account_ModuleAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ModuleAccount != nil {
		l = m.ModuleAccount.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Supply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *Supply_Supply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Supply != nil {
		l = m.Supply.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Evidence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *Evidence_Equivocation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Equivocation != nil {
		l = m.Equivocation.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func sovCodec(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCodec(x uint64) (n int) {
	return sovCodec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Account) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: Account: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Account: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types.BaseAccount{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Account_BaseAccount{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContinuousVestingAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types1.ContinuousVestingAccount{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Account_ContinuousVestingAccount{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelayedVestingAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types1.DelayedVestingAccount{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Account_DelayedVestingAccount{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodicVestingAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types1.PeriodicVestingAccount{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Account_PeriodicVestingAccount{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types2.ModuleAccount{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Account_ModuleAccount{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
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
func (m *Supply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: Supply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Supply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supply", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types2.Supply{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Supply_Supply{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
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
func (m *Evidence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: Evidence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Evidence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Equivocation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types3.Equivocation{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Evidence_Equivocation{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
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
func skipCodec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
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
				return 0, ErrInvalidLengthCodec
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCodec
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCodec
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCodec        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCodec          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCodec = fmt.Errorf("proto: unexpected end of group")
)
