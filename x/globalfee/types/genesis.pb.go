// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargaze/globalfee/v1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the globalfee module's genesis state.
type GenesisState struct {
	// Module params
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	// Authorizations configured by code id
	CodeAuthorizations []CodeAuthorization `protobuf:"bytes,2,rep,name=code_authorizations,json=codeAuthorizations,proto3" json:"code_authorizations,omitempty"`
	// Authorizations configured by contract addresses
	ContractAuthorizations []ContractAuthorization `protobuf:"bytes,3,rep,name=contract_authorizations,json=contractAuthorizations,proto3" json:"contract_authorizations,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_36828f28a9fbff57, []int{0}
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

func (m *GenesisState) GetCodeAuthorizations() []CodeAuthorization {
	if m != nil {
		return m.CodeAuthorizations
	}
	return nil
}

func (m *GenesisState) GetContractAuthorizations() []ContractAuthorization {
	if m != nil {
		return m.ContractAuthorizations
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "publicawesome.stargaze.globalfee.v1.GenesisState")
}

func init() {
	proto.RegisterFile("stargaze/globalfee/v1/genesis.proto", fileDescriptor_36828f28a9fbff57)
}

var fileDescriptor_36828f28a9fbff57 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x2e, 0x49, 0x2c,
	0x4a, 0x4f, 0xac, 0x4a, 0xd5, 0x4f, 0xcf, 0xc9, 0x4f, 0x4a, 0xcc, 0x49, 0x4b, 0x4d, 0xd5, 0x2f,
	0x33, 0xd4, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0xce, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x52, 0x2e, 0x28, 0x4d, 0xca, 0xc9, 0x4c, 0x4e, 0x2c, 0x4f, 0x2d, 0xce, 0xcf, 0x4d, 0xd5, 0x83,
	0x69, 0xd1, 0x83, 0x6b, 0xd1, 0x2b, 0x33, 0x94, 0x52, 0xc5, 0x61, 0x12, 0x5c, 0x0d, 0xd8, 0x2c,
	0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0x30, 0x53, 0x1f, 0xc4, 0x82, 0x88, 0x2a, 0xcd, 0x67, 0xe6,
	0xe2, 0x71, 0x87, 0xd8, 0x19, 0x5c, 0x92, 0x58, 0x92, 0x2a, 0x14, 0xcb, 0xc5, 0x56, 0x90, 0x58,
	0x94, 0x98, 0x5b, 0x2c, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xa4, 0xad, 0x47, 0x84, 0x1b, 0xf4,
	0x02, 0xc0, 0x5a, 0x9c, 0x24, 0x4e, 0xdc, 0x93, 0x67, 0x78, 0x75, 0x4f, 0x5e, 0x00, 0x62, 0x84,
	0x4e, 0x7e, 0x6e, 0x66, 0x49, 0x6a, 0x6e, 0x41, 0x49, 0x65, 0x10, 0xd4, 0x50, 0xa1, 0x3e, 0x46,
	0x2e, 0xe1, 0xe4, 0xfc, 0x94, 0xd4, 0xf8, 0xc4, 0xd2, 0x92, 0x8c, 0xfc, 0xa2, 0xcc, 0xaa, 0xc4,
	0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x09, 0x26, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x33, 0xa2, 0x2c, 0x73,
	0xce, 0x4f, 0x49, 0x75, 0x44, 0xd6, 0xee, 0xa4, 0x0a, 0xb5, 0x57, 0x16, 0x8b, 0xd1, 0x48, 0x8e,
	0x10, 0x4a, 0x46, 0xd7, 0x59, 0x2c, 0x34, 0x8f, 0x91, 0x4b, 0x3c, 0x39, 0x3f, 0xaf, 0xa4, 0x28,
	0x31, 0xb9, 0x04, 0xdd, 0x51, 0xcc, 0x60, 0x47, 0x59, 0x11, 0xe9, 0x28, 0x88, 0x19, 0xa8, 0x0e,
	0xd3, 0x84, 0x3a, 0x4c, 0x11, 0x87, 0x15, 0x48, 0x8e, 0x13, 0x4b, 0xc6, 0x66, 0x42, 0xb1, 0x53,
	0xd0, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1,
	0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x59, 0xa4, 0x67, 0x96, 0x64,
	0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x43, 0x9c, 0xa8, 0x0b, 0x75, 0xa3, 0x3e, 0x3c, 0x49,
	0x94, 0x19, 0x1a, 0xe8, 0x57, 0x20, 0x25, 0x8c, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70,
	0xe4, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x71, 0x63, 0xf7, 0xfe, 0x85, 0x02, 0x00, 0x00,
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
	if len(m.ContractAuthorizations) > 0 {
		for iNdEx := len(m.ContractAuthorizations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ContractAuthorizations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.CodeAuthorizations) > 0 {
		for iNdEx := len(m.CodeAuthorizations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CodeAuthorizations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.CodeAuthorizations) > 0 {
		for _, e := range m.CodeAuthorizations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ContractAuthorizations) > 0 {
		for _, e := range m.ContractAuthorizations {
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
				return fmt.Errorf("proto: wrong wireType = %d for field CodeAuthorizations", wireType)
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
			m.CodeAuthorizations = append(m.CodeAuthorizations, CodeAuthorization{})
			if err := m.CodeAuthorizations[len(m.CodeAuthorizations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAuthorizations", wireType)
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
			m.ContractAuthorizations = append(m.ContractAuthorizations, ContractAuthorization{})
			if err := m.ContractAuthorizations[len(m.ContractAuthorizations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
