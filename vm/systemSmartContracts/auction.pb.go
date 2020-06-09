// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auction.proto

package systemSmartContracts

import (
	bytes "bytes"
	fmt "fmt"
	github_com_ElrondNetwork_elrond_go_data "github.com/ElrondNetwork/elrond-go/data"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_big "math/big"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type AuctionData struct {
	RegisterNonce   uint64        `protobuf:"varint,1,opt,name=RegisterNonce,proto3" json:"RegisterNonce"`
	Epoch           uint32        `protobuf:"varint,2,opt,name=Epoch,proto3" json:"Epoch"`
	RewardAddress   []byte        `protobuf:"bytes,3,opt,name=RewardAddress,proto3" json:"RewardAddress"`
	TotalStakeValue *math_big.Int `protobuf:"bytes,4,opt,name=TotalStakeValue,proto3,casttypewith=math/big.Int;github.com/ElrondNetwork/elrond-go/data.BigIntCaster" json:"TotalStakeValue"`
	LockedStake     *math_big.Int `protobuf:"bytes,5,opt,name=LockedStake,proto3,casttypewith=math/big.Int;github.com/ElrondNetwork/elrond-go/data.BigIntCaster" json:"LockedStake"`
	MaxStakePerNode *math_big.Int `protobuf:"bytes,6,opt,name=MaxStakePerNode,proto3,casttypewith=math/big.Int;github.com/ElrondNetwork/elrond-go/data.BigIntCaster" json:"MaxStakePerNode"`
	BlsPubKeys      [][]byte      `protobuf:"bytes,7,rep,name=BlsPubKeys,proto3" json:"BlsPubKeys"`
	NumStaked       uint32        `protobuf:"varint,8,opt,name=NumStaked,proto3" json:"NumStaked"`
}

func (m *AuctionData) Reset()      { *m = AuctionData{} }
func (*AuctionData) ProtoMessage() {}
func (*AuctionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_622f477c3a3f2896, []int{0}
}
func (m *AuctionData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuctionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *AuctionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuctionData.Merge(m, src)
}
func (m *AuctionData) XXX_Size() int {
	return m.Size()
}
func (m *AuctionData) XXX_DiscardUnknown() {
	xxx_messageInfo_AuctionData.DiscardUnknown(m)
}

var xxx_messageInfo_AuctionData proto.InternalMessageInfo

func (m *AuctionData) GetRegisterNonce() uint64 {
	if m != nil {
		return m.RegisterNonce
	}
	return 0
}

func (m *AuctionData) GetEpoch() uint32 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *AuctionData) GetRewardAddress() []byte {
	if m != nil {
		return m.RewardAddress
	}
	return nil
}

func (m *AuctionData) GetTotalStakeValue() *math_big.Int {
	if m != nil {
		return m.TotalStakeValue
	}
	return nil
}

func (m *AuctionData) GetLockedStake() *math_big.Int {
	if m != nil {
		return m.LockedStake
	}
	return nil
}

func (m *AuctionData) GetMaxStakePerNode() *math_big.Int {
	if m != nil {
		return m.MaxStakePerNode
	}
	return nil
}

func (m *AuctionData) GetBlsPubKeys() [][]byte {
	if m != nil {
		return m.BlsPubKeys
	}
	return nil
}

func (m *AuctionData) GetNumStaked() uint32 {
	if m != nil {
		return m.NumStaked
	}
	return 0
}

type AuctionConfig struct {
	NumNodes      uint32        `protobuf:"varint,1,opt,name=NumNodes,proto3" json:"NumNodes"`
	MinStakeValue *math_big.Int `protobuf:"bytes,2,opt,name=MinStakeValue,proto3,casttypewith=math/big.Int;github.com/ElrondNetwork/elrond-go/data.BigIntCaster" json:"MinStakeValue"`
	TotalSupply   *math_big.Int `protobuf:"bytes,3,opt,name=TotalSupply,proto3,casttypewith=math/big.Int;github.com/ElrondNetwork/elrond-go/data.BigIntCaster" json:"TotalSupply"`
	MinStep       *math_big.Int `protobuf:"bytes,4,opt,name=MinStep,proto3,casttypewith=math/big.Int;github.com/ElrondNetwork/elrond-go/data.BigIntCaster" json:"MinStep"`
	NodePrice     *math_big.Int `protobuf:"bytes,5,opt,name=NodePrice,proto3,casttypewith=math/big.Int;github.com/ElrondNetwork/elrond-go/data.BigIntCaster" json:"NodePrice"`
	UnJailPrice   *math_big.Int `protobuf:"bytes,6,opt,name=UnJailPrice,proto3,casttypewith=math/big.Int;github.com/ElrondNetwork/elrond-go/data.BigIntCaster" json:"UnJailPrice"`
}

func (m *AuctionConfig) Reset()      { *m = AuctionConfig{} }
func (*AuctionConfig) ProtoMessage() {}
func (*AuctionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_622f477c3a3f2896, []int{1}
}
func (m *AuctionConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuctionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *AuctionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuctionConfig.Merge(m, src)
}
func (m *AuctionConfig) XXX_Size() int {
	return m.Size()
}
func (m *AuctionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AuctionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AuctionConfig proto.InternalMessageInfo

func (m *AuctionConfig) GetNumNodes() uint32 {
	if m != nil {
		return m.NumNodes
	}
	return 0
}

func (m *AuctionConfig) GetMinStakeValue() *math_big.Int {
	if m != nil {
		return m.MinStakeValue
	}
	return nil
}

func (m *AuctionConfig) GetTotalSupply() *math_big.Int {
	if m != nil {
		return m.TotalSupply
	}
	return nil
}

func (m *AuctionConfig) GetMinStep() *math_big.Int {
	if m != nil {
		return m.MinStep
	}
	return nil
}

func (m *AuctionConfig) GetNodePrice() *math_big.Int {
	if m != nil {
		return m.NodePrice
	}
	return nil
}

func (m *AuctionConfig) GetUnJailPrice() *math_big.Int {
	if m != nil {
		return m.UnJailPrice
	}
	return nil
}

func init() {
	proto.RegisterType((*AuctionData)(nil), "proto.AuctionData")
	proto.RegisterType((*AuctionConfig)(nil), "proto.AuctionConfig")
}

func init() { proto.RegisterFile("auction.proto", fileDescriptor_622f477c3a3f2896) }

var fileDescriptor_622f477c3a3f2896 = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x31, 0x6f, 0xd3, 0x40,
	0x18, 0xf5, 0xb5, 0x4d, 0xda, 0x5c, 0x63, 0x10, 0x27, 0x06, 0x8b, 0xe1, 0x1c, 0x75, 0x8a, 0x84,
	0x9a, 0x0c, 0x0c, 0x0c, 0x4c, 0x71, 0xe8, 0x50, 0xa0, 0x51, 0xe4, 0x96, 0x0a, 0xb1, 0x5d, 0xec,
	0xab, 0x63, 0x62, 0xfb, 0x2c, 0xfb, 0x4c, 0x89, 0xc4, 0x80, 0x90, 0xd8, 0xf9, 0x19, 0x88, 0x5f,
	0xc2, 0x98, 0x31, 0x93, 0x21, 0xce, 0x82, 0x3c, 0xf5, 0x1f, 0x80, 0x7c, 0x4e, 0x9a, 0x4b, 0x66,
	0x4f, 0xf7, 0x7d, 0xef, 0xd3, 0xbd, 0xe7, 0xef, 0xf9, 0xd9, 0x50, 0x25, 0x89, 0xc5, 0x5d, 0x16,
	0x74, 0xc2, 0x88, 0x71, 0x86, 0x6a, 0xe2, 0x78, 0x72, 0xea, 0xb8, 0x7c, 0x9c, 0x8c, 0x3a, 0x16,
	0xf3, 0xbb, 0x0e, 0x73, 0x58, 0x57, 0xc0, 0xa3, 0xe4, 0x46, 0x74, 0xa2, 0x11, 0x55, 0x79, 0xeb,
	0xe4, 0xdf, 0x01, 0x3c, 0xee, 0x95, 0x3c, 0x2f, 0x09, 0x27, 0xe8, 0x39, 0x54, 0x4d, 0xea, 0xb8,
	0x31, 0xa7, 0xd1, 0x80, 0x05, 0x16, 0xd5, 0x40, 0x0b, 0xb4, 0x0f, 0x8c, 0x47, 0x79, 0xaa, 0x6f,
	0x0f, 0xcc, 0xed, 0x16, 0xe9, 0xb0, 0x76, 0x16, 0x32, 0x6b, 0xac, 0xed, 0xb5, 0x40, 0x5b, 0x35,
	0x1a, 0x79, 0xaa, 0x97, 0x80, 0x59, 0x1e, 0x25, 0xf3, 0x2d, 0x89, 0xec, 0x9e, 0x6d, 0x47, 0x34,
	0x8e, 0xb5, 0xfd, 0x16, 0x68, 0x37, 0xd7, 0xcc, 0xd2, 0xc0, 0xdc, 0x6e, 0xd1, 0x57, 0x00, 0x1f,
	0x5e, 0x31, 0x4e, 0xbc, 0x4b, 0x4e, 0x26, 0xf4, 0x9a, 0x78, 0x09, 0xd5, 0x0e, 0xc4, 0xdd, 0x77,
	0x79, 0xaa, 0xef, 0x8e, 0x7e, 0xfe, 0xd6, 0x7b, 0x3e, 0xe1, 0xe3, 0xee, 0xc8, 0x75, 0x3a, 0xe7,
	0x01, 0x7f, 0x21, 0xf9, 0x71, 0xe6, 0x45, 0x2c, 0xb0, 0x07, 0x94, 0xdf, 0xb2, 0x68, 0xd2, 0xa5,
	0xa2, 0x3b, 0x75, 0x58, 0xd7, 0x26, 0x9c, 0x74, 0x0c, 0xd7, 0x39, 0x0f, 0x78, 0x9f, 0x14, 0x2b,
	0x99, 0xbb, 0xac, 0xe8, 0x23, 0x3c, 0x7e, 0xc3, 0xac, 0x09, 0xb5, 0x05, 0xa6, 0xd5, 0x84, 0xfe,
	0x55, 0x9e, 0xea, 0x32, 0x5c, 0x8d, 0xb6, 0xcc, 0x28, 0x96, 0xbf, 0x20, 0x9f, 0x44, 0x33, 0x2c,
	0xbc, 0xb6, 0xa9, 0x56, 0xdf, 0x2c, 0xbf, 0x33, 0xaa, 0x68, 0xf9, 0x1d, 0x56, 0xd4, 0x81, 0xd0,
	0xf0, 0xe2, 0x61, 0x32, 0x7a, 0x4d, 0xa7, 0xb1, 0x76, 0xd8, 0xda, 0x6f, 0x37, 0x8d, 0x07, 0x79,
	0xaa, 0x4b, 0xa8, 0x29, 0xd5, 0xe8, 0x29, 0x6c, 0x0c, 0x12, 0x5f, 0x50, 0xd8, 0xda, 0x91, 0xc8,
	0x83, 0x9a, 0xa7, 0xfa, 0x06, 0x34, 0x37, 0xe5, 0xc9, 0xb7, 0x1a, 0x54, 0x57, 0x09, 0xec, 0xb3,
	0xe0, 0xc6, 0x75, 0x50, 0x1b, 0x1e, 0x0d, 0x12, 0xbf, 0x50, 0x8e, 0x45, 0xfc, 0x54, 0xa3, 0x99,
	0xa7, 0xfa, 0x3d, 0x66, 0xde, 0x57, 0xe8, 0x33, 0x54, 0x2f, 0xdc, 0x40, 0xca, 0xc5, 0x9e, 0xb0,
	0xe6, 0xba, 0xc8, 0xd4, 0xd6, 0xa0, 0x1a, 0x63, 0xb6, 0x39, 0x8b, 0x4c, 0x94, 0x31, 0x49, 0xc2,
	0xd0, 0x9b, 0xae, 0xf2, 0x2c, 0x32, 0x21, 0xc1, 0x15, 0x65, 0x42, 0x62, 0x44, 0x1f, 0xe0, 0xa1,
	0x78, 0x10, 0x1a, 0xae, 0xbe, 0x83, 0x61, 0x9e, 0xea, 0x6b, 0xa8, 0x1a, 0xbd, 0x35, 0x1b, 0x0a,
	0x61, 0xa3, 0xb0, 0x7a, 0x18, 0xb9, 0xd6, 0x3a, 0xf5, 0xa6, 0x78, 0x95, 0x6b, 0xb0, 0x1a, 0xbd,
	0x0d, 0x5f, 0xe1, 0xea, 0xdb, 0xe0, 0x15, 0x71, 0xbd, 0x52, 0xb3, 0xbe, 0x71, 0x55, 0x82, 0x2b,
	0x72, 0x55, 0x62, 0x34, 0x06, 0xb3, 0x05, 0x56, 0xe6, 0x0b, 0xac, 0xdc, 0x2d, 0x30, 0xf8, 0x92,
	0x61, 0xf0, 0x23, 0xc3, 0xe0, 0x57, 0x86, 0xc1, 0x2c, 0xc3, 0x60, 0x9e, 0x61, 0xf0, 0x27, 0xc3,
	0xe0, 0x6f, 0x86, 0x95, 0xbb, 0x0c, 0x83, 0xef, 0x4b, 0xac, 0xcc, 0x96, 0x58, 0x99, 0x2f, 0xb1,
	0xf2, 0xfe, 0x71, 0x3c, 0x8d, 0x39, 0xf5, 0x2f, 0x7d, 0x12, 0xf1, 0x3e, 0x0b, 0x78, 0x44, 0x2c,
	0x1e, 0x8f, 0xea, 0xe2, 0x07, 0xfb, 0xec, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45, 0xe7, 0xd6,
	0x45, 0xa7, 0x05, 0x00, 0x00,
}

func (this *AuctionData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AuctionData)
	if !ok {
		that2, ok := that.(AuctionData)
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
	if this.RegisterNonce != that1.RegisterNonce {
		return false
	}
	if this.Epoch != that1.Epoch {
		return false
	}
	if !bytes.Equal(this.RewardAddress, that1.RewardAddress) {
		return false
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		if !__caster.Equal(this.TotalStakeValue, that1.TotalStakeValue) {
			return false
		}
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		if !__caster.Equal(this.LockedStake, that1.LockedStake) {
			return false
		}
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		if !__caster.Equal(this.MaxStakePerNode, that1.MaxStakePerNode) {
			return false
		}
	}
	if len(this.BlsPubKeys) != len(that1.BlsPubKeys) {
		return false
	}
	for i := range this.BlsPubKeys {
		if !bytes.Equal(this.BlsPubKeys[i], that1.BlsPubKeys[i]) {
			return false
		}
	}
	if this.NumStaked != that1.NumStaked {
		return false
	}
	return true
}
func (this *AuctionConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AuctionConfig)
	if !ok {
		that2, ok := that.(AuctionConfig)
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
	if this.NumNodes != that1.NumNodes {
		return false
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		if !__caster.Equal(this.MinStakeValue, that1.MinStakeValue) {
			return false
		}
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		if !__caster.Equal(this.TotalSupply, that1.TotalSupply) {
			return false
		}
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		if !__caster.Equal(this.MinStep, that1.MinStep) {
			return false
		}
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		if !__caster.Equal(this.NodePrice, that1.NodePrice) {
			return false
		}
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		if !__caster.Equal(this.UnJailPrice, that1.UnJailPrice) {
			return false
		}
	}
	return true
}
func (this *AuctionData) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 12)
	s = append(s, "&systemSmartContracts.AuctionData{")
	s = append(s, "RegisterNonce: "+fmt.Sprintf("%#v", this.RegisterNonce)+",\n")
	s = append(s, "Epoch: "+fmt.Sprintf("%#v", this.Epoch)+",\n")
	s = append(s, "RewardAddress: "+fmt.Sprintf("%#v", this.RewardAddress)+",\n")
	s = append(s, "TotalStakeValue: "+fmt.Sprintf("%#v", this.TotalStakeValue)+",\n")
	s = append(s, "LockedStake: "+fmt.Sprintf("%#v", this.LockedStake)+",\n")
	s = append(s, "MaxStakePerNode: "+fmt.Sprintf("%#v", this.MaxStakePerNode)+",\n")
	s = append(s, "BlsPubKeys: "+fmt.Sprintf("%#v", this.BlsPubKeys)+",\n")
	s = append(s, "NumStaked: "+fmt.Sprintf("%#v", this.NumStaked)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *AuctionConfig) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&systemSmartContracts.AuctionConfig{")
	s = append(s, "NumNodes: "+fmt.Sprintf("%#v", this.NumNodes)+",\n")
	s = append(s, "MinStakeValue: "+fmt.Sprintf("%#v", this.MinStakeValue)+",\n")
	s = append(s, "TotalSupply: "+fmt.Sprintf("%#v", this.TotalSupply)+",\n")
	s = append(s, "MinStep: "+fmt.Sprintf("%#v", this.MinStep)+",\n")
	s = append(s, "NodePrice: "+fmt.Sprintf("%#v", this.NodePrice)+",\n")
	s = append(s, "UnJailPrice: "+fmt.Sprintf("%#v", this.UnJailPrice)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAuction(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *AuctionData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuctionData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuctionData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NumStaked != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.NumStaked))
		i--
		dAtA[i] = 0x40
	}
	if len(m.BlsPubKeys) > 0 {
		for iNdEx := len(m.BlsPubKeys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.BlsPubKeys[iNdEx])
			copy(dAtA[i:], m.BlsPubKeys[iNdEx])
			i = encodeVarintAuction(dAtA, i, uint64(len(m.BlsPubKeys[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		size := __caster.Size(m.MaxStakePerNode)
		i -= size
		if _, err := __caster.MarshalTo(m.MaxStakePerNode, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		size := __caster.Size(m.LockedStake)
		i -= size
		if _, err := __caster.MarshalTo(m.LockedStake, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		size := __caster.Size(m.TotalStakeValue)
		i -= size
		if _, err := __caster.MarshalTo(m.TotalStakeValue, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.RewardAddress) > 0 {
		i -= len(m.RewardAddress)
		copy(dAtA[i:], m.RewardAddress)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.RewardAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Epoch != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x10
	}
	if m.RegisterNonce != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.RegisterNonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AuctionConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuctionConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuctionConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		size := __caster.Size(m.UnJailPrice)
		i -= size
		if _, err := __caster.MarshalTo(m.UnJailPrice, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		size := __caster.Size(m.NodePrice)
		i -= size
		if _, err := __caster.MarshalTo(m.NodePrice, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		size := __caster.Size(m.MinStep)
		i -= size
		if _, err := __caster.MarshalTo(m.MinStep, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		size := __caster.Size(m.TotalSupply)
		i -= size
		if _, err := __caster.MarshalTo(m.TotalSupply, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		size := __caster.Size(m.MinStakeValue)
		i -= size
		if _, err := __caster.MarshalTo(m.MinStakeValue, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.NumNodes != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.NumNodes))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuction(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AuctionData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RegisterNonce != 0 {
		n += 1 + sovAuction(uint64(m.RegisterNonce))
	}
	if m.Epoch != 0 {
		n += 1 + sovAuction(uint64(m.Epoch))
	}
	l = len(m.RewardAddress)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		l = __caster.Size(m.TotalStakeValue)
		n += 1 + l + sovAuction(uint64(l))
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		l = __caster.Size(m.LockedStake)
		n += 1 + l + sovAuction(uint64(l))
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		l = __caster.Size(m.MaxStakePerNode)
		n += 1 + l + sovAuction(uint64(l))
	}
	if len(m.BlsPubKeys) > 0 {
		for _, b := range m.BlsPubKeys {
			l = len(b)
			n += 1 + l + sovAuction(uint64(l))
		}
	}
	if m.NumStaked != 0 {
		n += 1 + sovAuction(uint64(m.NumStaked))
	}
	return n
}

func (m *AuctionConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NumNodes != 0 {
		n += 1 + sovAuction(uint64(m.NumNodes))
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		l = __caster.Size(m.MinStakeValue)
		n += 1 + l + sovAuction(uint64(l))
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		l = __caster.Size(m.TotalSupply)
		n += 1 + l + sovAuction(uint64(l))
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		l = __caster.Size(m.MinStep)
		n += 1 + l + sovAuction(uint64(l))
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		l = __caster.Size(m.NodePrice)
		n += 1 + l + sovAuction(uint64(l))
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		l = __caster.Size(m.UnJailPrice)
		n += 1 + l + sovAuction(uint64(l))
	}
	return n
}

func sovAuction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuction(x uint64) (n int) {
	return sovAuction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *AuctionData) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AuctionData{`,
		`RegisterNonce:` + fmt.Sprintf("%v", this.RegisterNonce) + `,`,
		`Epoch:` + fmt.Sprintf("%v", this.Epoch) + `,`,
		`RewardAddress:` + fmt.Sprintf("%v", this.RewardAddress) + `,`,
		`TotalStakeValue:` + fmt.Sprintf("%v", this.TotalStakeValue) + `,`,
		`LockedStake:` + fmt.Sprintf("%v", this.LockedStake) + `,`,
		`MaxStakePerNode:` + fmt.Sprintf("%v", this.MaxStakePerNode) + `,`,
		`BlsPubKeys:` + fmt.Sprintf("%v", this.BlsPubKeys) + `,`,
		`NumStaked:` + fmt.Sprintf("%v", this.NumStaked) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AuctionConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AuctionConfig{`,
		`NumNodes:` + fmt.Sprintf("%v", this.NumNodes) + `,`,
		`MinStakeValue:` + fmt.Sprintf("%v", this.MinStakeValue) + `,`,
		`TotalSupply:` + fmt.Sprintf("%v", this.TotalSupply) + `,`,
		`MinStep:` + fmt.Sprintf("%v", this.MinStep) + `,`,
		`NodePrice:` + fmt.Sprintf("%v", this.NodePrice) + `,`,
		`UnJailPrice:` + fmt.Sprintf("%v", this.UnJailPrice) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAuction(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *AuctionData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: AuctionData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuctionData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegisterNonce", wireType)
			}
			m.RegisterNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegisterNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardAddress = append(m.RewardAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.RewardAddress == nil {
				m.RewardAddress = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalStakeValue", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.TotalStakeValue = tmp
				}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockedStake", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.LockedStake = tmp
				}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxStakePerNode", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.MaxStakePerNode = tmp
				}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlsPubKeys", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlsPubKeys = append(m.BlsPubKeys, make([]byte, postIndex-iNdEx))
			copy(m.BlsPubKeys[len(m.BlsPubKeys)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumStaked", wireType)
			}
			m.NumStaked = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumStaked |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuction
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAuction
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
func (m *AuctionConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: AuctionConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuctionConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumNodes", wireType)
			}
			m.NumNodes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumNodes |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinStakeValue", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.MinStakeValue = tmp
				}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalSupply", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.TotalSupply = tmp
				}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinStep", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.MinStep = tmp
				}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodePrice", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.NodePrice = tmp
				}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnJailPrice", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.UnJailPrice = tmp
				}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuction
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAuction
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
func skipAuction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
				return 0, ErrInvalidLengthAuction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuction = fmt.Errorf("proto: unexpected end of group")
)