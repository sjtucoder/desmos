// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v3/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the profiles module's genesis state.
type GenesisState struct {
	DTagTransferRequests     []DTagTransferRequest         `protobuf:"bytes,1,rep,name=dtag_transfer_requests,json=dtagTransferRequests,proto3" json:"dtag_transfer_requests" yaml:"dtag_transfer_requests"`
	ChainLinks               []ChainLink                   `protobuf:"bytes,2,rep,name=chain_links,json=chainLinks,proto3" json:"chain_links" yaml:"chain_links"`
	ApplicationLinks         []ApplicationLink             `protobuf:"bytes,3,rep,name=application_links,json=applicationLinks,proto3" json:"application_links" yaml:"application_links"`
	DefaultExternalAddresses []DefaultExternalAddressEntry `protobuf:"bytes,4,rep,name=default_external_addresses,json=defaultExternalAddresses,proto3" json:"default_external_addresses" yaml:"default_external_addresses"`
	IBCPortID                string                        `protobuf:"bytes,5,opt,name=ibc_port_id,json=ibcPortId,proto3" json:"ibc_port_id,omitempty" yaml:"ibc_port_id"`
	Params                   Params                        `protobuf:"bytes,6,opt,name=params,proto3" json:"params" yaml:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd22d098f73f0a1c, []int{0}
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

// DefaultExternalAddressEntry contains the data of a default extnernal address
type DefaultExternalAddressEntry struct {
	Owner     string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	ChainName string `protobuf:"bytes,2,opt,name=chain_name,json=chainName,proto3" json:"chain_name,omitempty"`
	Target    string `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
}

func (m *DefaultExternalAddressEntry) Reset()         { *m = DefaultExternalAddressEntry{} }
func (m *DefaultExternalAddressEntry) String() string { return proto.CompactTextString(m) }
func (*DefaultExternalAddressEntry) ProtoMessage()    {}
func (*DefaultExternalAddressEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd22d098f73f0a1c, []int{1}
}
func (m *DefaultExternalAddressEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DefaultExternalAddressEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DefaultExternalAddressEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DefaultExternalAddressEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefaultExternalAddressEntry.Merge(m, src)
}
func (m *DefaultExternalAddressEntry) XXX_Size() int {
	return m.Size()
}
func (m *DefaultExternalAddressEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_DefaultExternalAddressEntry.DiscardUnknown(m)
}

var xxx_messageInfo_DefaultExternalAddressEntry proto.InternalMessageInfo

func (m *DefaultExternalAddressEntry) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *DefaultExternalAddressEntry) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *DefaultExternalAddressEntry) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "desmos.profiles.v3.GenesisState")
	proto.RegisterType((*DefaultExternalAddressEntry)(nil), "desmos.profiles.v3.DefaultExternalAddressEntry")
}

func init() { proto.RegisterFile("desmos/profiles/v3/genesis.proto", fileDescriptor_bd22d098f73f0a1c) }

var fileDescriptor_bd22d098f73f0a1c = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0xe3, 0xfe, 0x89, 0x94, 0x0b, 0x48, 0xf4, 0x14, 0x55, 0x26, 0xa8, 0x76, 0x48, 0x05,
	0x14, 0x50, 0x6d, 0xd1, 0x8a, 0xa5, 0x5b, 0xdd, 0x56, 0xa8, 0x02, 0x55, 0x55, 0xda, 0x89, 0xc5,
	0xba, 0xd8, 0x57, 0xf7, 0x54, 0xdb, 0x67, 0xee, 0x2e, 0x69, 0xb3, 0x33, 0x30, 0xf2, 0x05, 0x90,
	0x60, 0x63, 0x64, 0xe0, 0x3b, 0xd0, 0xb1, 0x62, 0x62, 0x8a, 0x50, 0x32, 0xb0, 0xf7, 0x13, 0x20,
	0xdf, 0x9d, 0x49, 0xa0, 0x0e, 0x4b, 0x74, 0xf6, 0xfb, 0x7b, 0xde, 0xe7, 0xb9, 0x37, 0x7e, 0x41,
	0x2b, 0xc4, 0x3c, 0xa1, 0xdc, 0xcd, 0x18, 0x3d, 0x21, 0x31, 0xe6, 0x6e, 0x7f, 0xd3, 0x8d, 0x70,
	0x8a, 0x39, 0xe1, 0x4e, 0xc6, 0xa8, 0xa0, 0x10, 0x2a, 0xc2, 0x29, 0x08, 0xa7, 0xbf, 0xd9, 0x5c,
	0x42, 0x09, 0x49, 0xa9, 0x2b, 0x7f, 0x15, 0xd6, 0x6c, 0x44, 0x34, 0xa2, 0xf2, 0xe8, 0xe6, 0x27,
	0xfd, 0xf6, 0x6e, 0x40, 0x73, 0xb1, 0xaf, 0x0a, 0xea, 0x41, 0x97, 0x1e, 0x96, 0x38, 0x27, 0x34,
	0xc4, 0x31, 0xf7, 0x33, 0xc4, 0x50, 0x52, 0x70, 0xeb, 0xb3, 0xb9, 0x50, 0xa0, 0xc8, 0x67, 0xf8,
	0x4d, 0x0f, 0x73, 0x51, 0xe0, 0x4f, 0x67, 0xe3, 0xc1, 0x29, 0x22, 0xa9, 0x1f, 0x93, 0xf4, 0xac,
	0x80, 0x1f, 0xcf, 0x86, 0x51, 0x96, 0x4d, 0xa3, 0xed, 0x6f, 0x8b, 0xe0, 0xd6, 0x0b, 0x35, 0x98,
	0x23, 0x81, 0x04, 0x86, 0x9f, 0x0c, 0xb0, 0x2c, 0x03, 0x08, 0x86, 0x52, 0x7e, 0x82, 0xd9, 0x9f,
	0x24, 0xa6, 0xd1, 0x9a, 0x5f, 0xab, 0x6f, 0x3c, 0x72, 0x6e, 0x4e, 0xce, 0xd9, 0x3d, 0x46, 0xd1,
	0xb1, 0x16, 0x74, 0x14, 0xef, 0x79, 0x97, 0x43, 0xbb, 0x32, 0x1a, 0xda, 0x8d, 0x92, 0x22, 0xbf,
	0x1e, 0xda, 0x2b, 0x03, 0x94, 0xc4, 0x5b, 0xed, 0x72, 0xb3, 0xf6, 0xe7, 0x5f, 0x5f, 0x9e, 0x18,
	0x9d, 0x46, 0x5e, 0xfd, 0x57, 0x0b, 0x7d, 0x50, 0x9f, 0xba, 0xb4, 0x39, 0x27, 0x73, 0xad, 0x94,
	0xe5, 0xda, 0xc9, 0xb1, 0x57, 0x24, 0x3d, 0xf3, 0xec, 0x3c, 0xcd, 0xf5, 0xd0, 0x86, 0xca, 0x75,
	0x4a, 0xaf, 0xad, 0x40, 0x50, 0xb0, 0x1c, 0x9e, 0x83, 0x25, 0x94, 0x65, 0x31, 0x09, 0x90, 0x20,
	0xb4, 0xb0, 0x99, 0x97, 0x36, 0xab, 0x65, 0x36, 0xdb, 0x13, 0x58, 0x9a, 0x3d, 0xd0, 0x66, 0xa6,
	0x32, 0xbb, 0xd1, 0x4b, 0x5b, 0xde, 0x41, 0x7f, 0xeb, 0x38, 0xfc, 0x60, 0x80, 0x66, 0x88, 0x4f,
	0x50, 0x2f, 0x16, 0x3e, 0xbe, 0x10, 0x98, 0xa5, 0x28, 0xf6, 0x51, 0x18, 0x32, 0xcc, 0x39, 0xe6,
	0xe6, 0x82, 0x8c, 0xe0, 0x96, 0xfe, 0x03, 0x4a, 0xb5, 0xa7, 0x45, 0xdb, 0x4a, 0xb3, 0x97, 0x0a,
	0x36, 0xf0, 0x1c, 0x1d, 0xe7, 0xbe, 0x9e, 0xf8, 0x4c, 0x03, 0x9d, 0xcb, 0x0c, 0x4b, 0x9b, 0x61,
	0x0e, 0x77, 0x40, 0x9d, 0x74, 0x03, 0x3f, 0xa3, 0x4c, 0xf8, 0x24, 0x34, 0x17, 0x5b, 0xc6, 0x5a,
	0xcd, 0x5b, 0x1d, 0x0d, 0xed, 0xda, 0xbe, 0xb7, 0x73, 0x48, 0x99, 0xd8, 0xdf, 0x9d, 0xcc, 0x78,
	0x8a, 0x6c, 0x77, 0x6a, 0xa4, 0x1b, 0x48, 0x20, 0x84, 0x07, 0xa0, 0xaa, 0x56, 0xc1, 0xac, 0xb6,
	0x8c, 0xb5, 0xfa, 0x46, 0xb3, 0xec, 0x3e, 0x87, 0x92, 0xf0, 0x9a, 0x3a, 0xfa, 0x6d, 0xd5, 0x52,
	0xe9, 0x74, 0x4c, 0xdd, 0x65, 0x6b, 0xe1, 0xdd, 0x47, 0xbb, 0xd2, 0x7e, 0x6b, 0x80, 0x7b, 0xff,
	0x19, 0x02, 0x74, 0xc0, 0x22, 0x3d, 0x4f, 0x31, 0x33, 0x0d, 0x19, 0xda, 0xfc, 0xfe, 0x75, 0xbd,
	0xa1, 0x37, 0x57, 0x73, 0x47, 0x82, 0x91, 0x34, 0xea, 0x28, 0x0c, 0xae, 0x00, 0xf5, 0x45, 0xf8,
	0x29, 0x4a, 0xb0, 0x39, 0x97, 0x8b, 0x3a, 0x35, 0xf9, 0xe6, 0x00, 0x25, 0x18, 0x2e, 0x83, 0xaa,
	0x40, 0x2c, 0xc2, 0xc2, 0x9c, 0x97, 0x25, 0xfd, 0xe4, 0xbd, 0xbc, 0x1c, 0x59, 0xc6, 0xd5, 0xc8,
	0x32, 0x7e, 0x8e, 0x2c, 0xe3, 0xfd, 0xd8, 0xaa, 0x5c, 0x8d, 0xad, 0xca, 0x8f, 0xb1, 0x55, 0x79,
	0xfd, 0x2c, 0x22, 0xe2, 0xb4, 0xd7, 0x75, 0x02, 0x9a, 0xb8, 0xea, 0xc2, 0xeb, 0x31, 0xea, 0x72,
	0x7d, 0x76, 0xfb, 0xcf, 0xdd, 0x8b, 0xc9, 0xc6, 0x8a, 0x41, 0x86, 0x79, 0xb7, 0x2a, 0x97, 0x74,
	0xf3, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0x0b, 0x97, 0x85, 0xcf, 0x04, 0x00, 0x00,
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.IBCPortID) > 0 {
		i -= len(m.IBCPortID)
		copy(dAtA[i:], m.IBCPortID)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.IBCPortID)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DefaultExternalAddresses) > 0 {
		for iNdEx := len(m.DefaultExternalAddresses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DefaultExternalAddresses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ApplicationLinks) > 0 {
		for iNdEx := len(m.ApplicationLinks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ApplicationLinks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ChainLinks) > 0 {
		for iNdEx := len(m.ChainLinks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChainLinks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.DTagTransferRequests) > 0 {
		for iNdEx := len(m.DTagTransferRequests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DTagTransferRequests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *DefaultExternalAddressEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DefaultExternalAddressEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DefaultExternalAddressEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Target) > 0 {
		i -= len(m.Target)
		copy(dAtA[i:], m.Target)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Target)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainName) > 0 {
		i -= len(m.ChainName)
		copy(dAtA[i:], m.ChainName)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChainName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Owner)))
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
	if len(m.DTagTransferRequests) > 0 {
		for _, e := range m.DTagTransferRequests {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ChainLinks) > 0 {
		for _, e := range m.ChainLinks {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ApplicationLinks) > 0 {
		for _, e := range m.ApplicationLinks {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DefaultExternalAddresses) > 0 {
		for _, e := range m.DefaultExternalAddresses {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = len(m.IBCPortID)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *DefaultExternalAddressEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ChainName)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Target)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field DTagTransferRequests", wireType)
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
			m.DTagTransferRequests = append(m.DTagTransferRequests, DTagTransferRequest{})
			if err := m.DTagTransferRequests[len(m.DTagTransferRequests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainLinks", wireType)
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
			m.ChainLinks = append(m.ChainLinks, ChainLink{})
			if err := m.ChainLinks[len(m.ChainLinks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationLinks", wireType)
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
			m.ApplicationLinks = append(m.ApplicationLinks, ApplicationLink{})
			if err := m.ApplicationLinks[len(m.ApplicationLinks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultExternalAddresses", wireType)
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
			m.DefaultExternalAddresses = append(m.DefaultExternalAddresses, DefaultExternalAddressEntry{})
			if err := m.DefaultExternalAddresses[len(m.DefaultExternalAddresses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IBCPortID", wireType)
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
			m.IBCPortID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
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
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *DefaultExternalAddressEntry) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DefaultExternalAddressEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DefaultExternalAddressEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainName", wireType)
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
			m.ChainName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
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
			m.Target = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
