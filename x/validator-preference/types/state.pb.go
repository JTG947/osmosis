// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/validator-preference/v1beta1/state.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// ValidatorPreference defines the message structure for
// CreateValidatorSetPreference. It allows a user to set {val_addr, weight} in
// state. If a user does not have a validator set preference list set, and has
// staked, make their preference list default to their current staking
// distribution.
type ValidatorPreference struct {
	// val_oper_address holds the validator address the user wants to delegate
	// funds to.
	ValOperAddress string `protobuf:"bytes,1,opt,name=val_oper_address,json=valOperAddress,proto3" json:"val_oper_address,omitempty" yaml:"val_oper_address"`
	// weight is decimal between 0 and 1, and they all sum to 1.
	Weight github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=weight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"weight"`
}

func (m *ValidatorPreference) Reset()         { *m = ValidatorPreference{} }
func (m *ValidatorPreference) String() string { return proto.CompactTextString(m) }
func (*ValidatorPreference) ProtoMessage()    {}
func (*ValidatorPreference) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f4199f1be974865, []int{0}
}
func (m *ValidatorPreference) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorPreference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorPreference.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorPreference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorPreference.Merge(m, src)
}
func (m *ValidatorPreference) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorPreference) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorPreference.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorPreference proto.InternalMessageInfo

// ValidatorSetPreferences defines a delegator's validator set preference.
// It contains a list of (validator, percent_allocation) pairs.
// The percent allocation are arranged in decimal notation from 0 to 1 and must
// add up to 1.
type ValidatorSetPreferences struct {
	// preference holds {valAddr, weight} for the user who created it.
	Preferences []ValidatorPreference `protobuf:"bytes,2,rep,name=preferences,proto3" json:"preferences" yaml:"preferences"`
}

func (m *ValidatorSetPreferences) Reset()         { *m = ValidatorSetPreferences{} }
func (m *ValidatorSetPreferences) String() string { return proto.CompactTextString(m) }
func (*ValidatorSetPreferences) ProtoMessage()    {}
func (*ValidatorSetPreferences) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f4199f1be974865, []int{1}
}
func (m *ValidatorSetPreferences) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSetPreferences) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSetPreferences.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSetPreferences) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSetPreferences.Merge(m, src)
}
func (m *ValidatorSetPreferences) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSetPreferences) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSetPreferences.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSetPreferences proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ValidatorPreference)(nil), "osmosis.validatorpreference.v1beta1.ValidatorPreference")
	proto.RegisterType((*ValidatorSetPreferences)(nil), "osmosis.validatorpreference.v1beta1.ValidatorSetPreferences")
}

func init() {
	proto.RegisterFile("osmosis/validator-preference/v1beta1/state.proto", fileDescriptor_2f4199f1be974865)
}

var fileDescriptor_2f4199f1be974865 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0x9b, 0xfd, 0x61, 0xf0, 0xef, 0x40, 0xa4, 0x0a, 0x1b, 0x53, 0xd2, 0x51, 0x41, 0x76,
	0x59, 0xe2, 0xe6, 0x45, 0x3c, 0xe9, 0x50, 0xaf, 0xca, 0x04, 0x0f, 0x1e, 0x1c, 0x69, 0xfb, 0xda,
	0x15, 0xbb, 0xa6, 0x24, 0xb1, 0xba, 0x6f, 0xa1, 0xdf, 0xc1, 0x0f, 0xb3, 0xe3, 0x8e, 0xe2, 0xa1,
	0xe8, 0xf6, 0x0d, 0xf6, 0x09, 0x64, 0x6d, 0xed, 0x86, 0xec, 0xe0, 0x29, 0x21, 0x79, 0x7f, 0x4f,
	0x9e, 0xe7, 0x7d, 0xa3, 0x1f, 0x70, 0x39, 0xe4, 0xd2, 0x97, 0x34, 0x66, 0x81, 0xef, 0x32, 0xc5,
	0x45, 0x2b, 0x12, 0x70, 0x0f, 0x02, 0x42, 0x07, 0x68, 0xdc, 0xb6, 0x41, 0xb1, 0x36, 0x95, 0x8a,
	0x29, 0x20, 0x91, 0xe0, 0x8a, 0x1b, 0x7b, 0x39, 0x41, 0x0a, 0x62, 0x09, 0x90, 0x1c, 0xa8, 0x6f,
	0x7b, 0xdc, 0xe3, 0x69, 0x3d, 0x5d, 0xec, 0x32, 0xb4, 0xbe, 0xeb, 0x71, 0xee, 0x05, 0x40, 0x59,
	0xe4, 0x53, 0x16, 0x86, 0x5c, 0x31, 0xe5, 0xf3, 0x50, 0x66, 0xb7, 0xd6, 0x1b, 0xd2, 0xb7, 0x6e,
	0x7e, 0x34, 0xaf, 0x0a, 0x4d, 0xe3, 0x5c, 0xdf, 0x8c, 0x59, 0xd0, 0xe7, 0x11, 0x88, 0x3e, 0x73,
	0x5d, 0x01, 0x52, 0xd6, 0x50, 0x03, 0x35, 0xff, 0x77, 0x77, 0xe6, 0x89, 0x59, 0x1d, 0xb1, 0x61,
	0x70, 0x6c, 0xfd, 0xae, 0xb0, 0x7a, 0x1b, 0x31, 0x0b, 0x2e, 0x23, 0x10, 0xa7, 0xd9, 0x81, 0x71,
	0xa1, 0x97, 0x9f, 0xc0, 0xf7, 0x06, 0xaa, 0x56, 0x4a, 0x61, 0x32, 0x4e, 0x4c, 0xed, 0x23, 0x31,
	0xf7, 0x3d, 0x5f, 0x0d, 0x1e, 0x6d, 0xe2, 0xf0, 0x21, 0x75, 0xd2, 0x6c, 0xf9, 0xd2, 0x92, 0xee,
	0x03, 0x55, 0xa3, 0x08, 0x24, 0x39, 0x03, 0xa7, 0x97, 0xd3, 0xd6, 0x2b, 0xd2, 0xab, 0x85, 0xcd,
	0x6b, 0x50, 0x4b, 0xa7, 0xd2, 0x88, 0xf5, 0xca, 0xb2, 0x19, 0xb2, 0x56, 0x6a, 0xfc, 0x6b, 0x56,
	0x3a, 0x47, 0xe4, 0x0f, 0x1d, 0x23, 0x6b, 0x92, 0x77, 0xeb, 0x0b, 0x8b, 0xf3, 0xc4, 0x34, 0xb2,
	0x8c, 0x2b, 0xd2, 0x56, 0x6f, 0xf5, 0xa1, 0xee, 0xdd, 0xf8, 0x0b, 0x6b, 0xe3, 0x29, 0x46, 0x93,
	0x29, 0x46, 0x9f, 0x53, 0x8c, 0x5e, 0x66, 0x58, 0x9b, 0xcc, 0xb0, 0xf6, 0x3e, 0xc3, 0xda, 0xed,
	0xc9, 0x4a, 0xc2, 0xdc, 0x4a, 0x2b, 0x60, 0xb6, 0xa4, 0xc5, 0xec, 0xdb, 0x1d, 0xfa, 0xbc, 0xfe,
	0x07, 0xa4, 0xf9, 0xed, 0x72, 0x3a, 0xa1, 0xc3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0xdc,
	0x3d, 0x1c, 0x2e, 0x02, 0x00, 0x00,
}

func (m *ValidatorPreference) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorPreference) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorPreference) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ValOperAddress) > 0 {
		i -= len(m.ValOperAddress)
		copy(dAtA[i:], m.ValOperAddress)
		i = encodeVarintState(dAtA, i, uint64(len(m.ValOperAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorSetPreferences) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSetPreferences) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSetPreferences) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Preferences) > 0 {
		for iNdEx := len(m.Preferences) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Preferences[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintState(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	offset -= sovState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorPreference) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValOperAddress)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = m.Weight.Size()
	n += 1 + l + sovState(uint64(l))
	return n
}

func (m *ValidatorSetPreferences) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Preferences) > 0 {
		for _, e := range m.Preferences {
			l = e.Size()
			n += 1 + l + sovState(uint64(l))
		}
	}
	return n
}

func sovState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorPreference) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: ValidatorPreference: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorPreference: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValOperAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValOperAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func (m *ValidatorSetPreferences) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: ValidatorSetPreferences: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSetPreferences: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Preferences", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Preferences = append(m.Preferences, ValidatorPreference{})
			if err := m.Preferences[len(m.Preferences)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
				return 0, ErrInvalidLengthState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupState = fmt.Errorf("proto: unexpected end of group")
)