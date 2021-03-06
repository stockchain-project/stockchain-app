// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blockchain/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the blockchain module's genesis state.
type GenesisState struct {
	// this line is used by starport scaffolding # genesis/proto/state
	BoardCommentList     []*BoardComment     `protobuf:"bytes,6,rep,name=boardCommentList,proto3" json:"boardCommentList,omitempty"`
	BoardList            []*Board            `protobuf:"bytes,4,rep,name=boardList,proto3" json:"boardList,omitempty"`
	BoardCount           uint64              `protobuf:"varint,5,opt,name=boardCount,proto3" json:"boardCount,omitempty"`
	StockTransactionList []*StockTransaction `protobuf:"bytes,3,rep,name=stockTransactionList,proto3" json:"stockTransactionList,omitempty"`
	StockDataList        []*StockData        `protobuf:"bytes,2,rep,name=stockDataList,proto3" json:"stockDataList,omitempty"`
	UserList             []*User             `protobuf:"bytes,1,rep,name=userList,proto3" json:"userList,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d37ae782202dfbb, []int{0}
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

func (m *GenesisState) GetBoardCommentList() []*BoardComment {
	if m != nil {
		return m.BoardCommentList
	}
	return nil
}

func (m *GenesisState) GetBoardList() []*Board {
	if m != nil {
		return m.BoardList
	}
	return nil
}

func (m *GenesisState) GetBoardCount() uint64 {
	if m != nil {
		return m.BoardCount
	}
	return 0
}

func (m *GenesisState) GetStockTransactionList() []*StockTransaction {
	if m != nil {
		return m.StockTransactionList
	}
	return nil
}

func (m *GenesisState) GetStockDataList() []*StockData {
	if m != nil {
		return m.StockDataList
	}
	return nil
}

func (m *GenesisState) GetUserList() []*User {
	if m != nil {
		return m.UserList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "chainstockproject.blockchain.blockchain.GenesisState")
}

func init() { proto.RegisterFile("blockchain/genesis.proto", fileDescriptor_1d37ae782202dfbb) }

var fileDescriptor_1d37ae782202dfbb = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4b, 0x3a, 0x41,
	0x14, 0xc7, 0xdd, 0x9f, 0xfe, 0xa4, 0xa6, 0x82, 0x18, 0x2a, 0xc4, 0x60, 0x10, 0x2f, 0x79, 0x71,
	0x05, 0xa3, 0x43, 0x74, 0xb3, 0x20, 0x02, 0x4f, 0x6b, 0x42, 0x74, 0x91, 0xd9, 0x71, 0xd0, 0xcd,
	0x76, 0x46, 0x66, 0x9e, 0x50, 0xff, 0x45, 0x7f, 0x56, 0x47, 0x8f, 0x1d, 0x43, 0xff, 0x84, 0xfe,
	0x81, 0xf0, 0xed, 0xd2, 0x4e, 0xda, 0x61, 0xbd, 0xed, 0x7e, 0xdf, 0x9b, 0xcf, 0x67, 0xdf, 0xce,
	0x23, 0x95, 0xf0, 0x59, 0x8b, 0x89, 0x18, 0xf3, 0x48, 0xb5, 0x46, 0x52, 0x49, 0x1b, 0x59, 0x7f,
	0x6a, 0x34, 0x68, 0x7a, 0x86, 0xa1, 0x05, 0x2d, 0x26, 0x53, 0xa3, 0x9f, 0xa4, 0x00, 0x3f, 0xeb,
	0x75, 0x1e, 0xab, 0xcc, 0x41, 0x84, 0x9a, 0x9b, 0xe1, 0x40, 0xe8, 0x38, 0x96, 0x0a, 0x12, 0x50,
	0xf5, 0x64, 0xbd, 0x9e, 0xe6, 0x75, 0x27, 0x47, 0xcd, 0x00, 0x0c, 0x57, 0x96, 0x0b, 0x88, 0xb4,
	0x4a, 0x7b, 0x4e, 0x37, 0x7a, 0x86, 0x1c, 0x78, 0x5a, 0x3c, 0x76, 0x8a, 0x33, 0x2b, 0x4d, 0x12,
	0xd7, 0xbf, 0x8a, 0x64, 0xff, 0x36, 0x19, 0xa5, 0x07, 0x1c, 0x24, 0xe5, 0xe4, 0x10, 0xbd, 0xd7,
	0xc9, 0x67, 0x75, 0x23, 0x0b, 0x95, 0x72, 0xad, 0xd8, 0xd8, 0x6b, 0x5f, 0xf8, 0x39, 0x87, 0xf4,
	0x3b, 0x0e, 0x20, 0xd8, 0xc0, 0xd1, 0x2e, 0xd9, 0xc5, 0x0c, 0xd9, 0x25, 0x64, 0xfb, 0xdb, 0xb1,
	0x83, 0x0c, 0x40, 0x19, 0x21, 0xa9, 0x61, 0xa6, 0xa0, 0xf2, 0xbf, 0xe6, 0x35, 0x4a, 0x81, 0x93,
	0xd0, 0x98, 0x1c, 0x21, 0xf6, 0x3e, 0xfb, 0x5f, 0x28, 0x2e, 0xa2, 0xf8, 0x32, 0xb7, 0xb8, 0xb7,
	0x06, 0x09, 0xfe, 0xc4, 0xd2, 0x07, 0x72, 0x80, 0xf9, 0x0d, 0x07, 0x8e, 0x9e, 0x7f, 0xe8, 0x69,
	0x6f, 0xe7, 0x59, 0x9d, 0x0e, 0x7e, 0x83, 0xe8, 0x1d, 0xd9, 0x59, 0x5d, 0x1c, 0x42, 0x3d, 0x84,
	0x36, 0x73, 0x43, 0xfb, 0x56, 0x9a, 0xe0, 0xe7, 0x78, 0xa7, 0xff, 0xbe, 0x60, 0xde, 0x7c, 0xc1,
	0xbc, 0xcf, 0x05, 0xf3, 0xde, 0x96, 0xac, 0x30, 0x5f, 0xb2, 0xc2, 0xc7, 0x92, 0x15, 0x1e, 0xaf,
	0x46, 0x11, 0x8c, 0x67, 0xa1, 0x2f, 0x74, 0xdc, 0xca, 0xe0, 0xcd, 0x94, 0xde, 0x72, 0x96, 0xe8,
	0xc5, 0x7d, 0x81, 0xd7, 0xa9, 0xb4, 0x61, 0x19, 0x77, 0xea, 0xfc, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0xc4, 0x8e, 0x94, 0x40, 0x28, 0x03, 0x00, 0x00,
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
	if len(m.BoardCommentList) > 0 {
		for iNdEx := len(m.BoardCommentList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BoardCommentList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.BoardCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.BoardCount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.BoardList) > 0 {
		for iNdEx := len(m.BoardList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BoardList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.StockTransactionList) > 0 {
		for iNdEx := len(m.StockTransactionList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StockTransactionList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.StockDataList) > 0 {
		for iNdEx := len(m.StockDataList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StockDataList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.UserList) > 0 {
		for iNdEx := len(m.UserList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UserList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.UserList) > 0 {
		for _, e := range m.UserList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.StockDataList) > 0 {
		for _, e := range m.StockDataList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.StockTransactionList) > 0 {
		for _, e := range m.StockTransactionList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BoardList) > 0 {
		for _, e := range m.BoardList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.BoardCount != 0 {
		n += 1 + sovGenesis(uint64(m.BoardCount))
	}
	if len(m.BoardCommentList) > 0 {
		for _, e := range m.BoardCommentList {
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
				return fmt.Errorf("proto: wrong wireType = %d for field UserList", wireType)
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
			m.UserList = append(m.UserList, &User{})
			if err := m.UserList[len(m.UserList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StockDataList", wireType)
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
			m.StockDataList = append(m.StockDataList, &StockData{})
			if err := m.StockDataList[len(m.StockDataList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StockTransactionList", wireType)
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
			m.StockTransactionList = append(m.StockTransactionList, &StockTransaction{})
			if err := m.StockTransactionList[len(m.StockTransactionList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BoardList", wireType)
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
			m.BoardList = append(m.BoardList, &Board{})
			if err := m.BoardList[len(m.BoardList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BoardCount", wireType)
			}
			m.BoardCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BoardCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BoardCommentList", wireType)
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
			m.BoardCommentList = append(m.BoardCommentList, &BoardComment{})
			if err := m.BoardCommentList[len(m.BoardCommentList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
