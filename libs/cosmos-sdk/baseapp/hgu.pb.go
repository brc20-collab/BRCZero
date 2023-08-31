// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hgu.proto

package baseapp

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

type HguRecord struct {
	MaxGas               int64    `protobuf:"varint,1,opt,name=MaxGas,proto3" json:"MaxGas,omitempty"`
	MinGas               int64    `protobuf:"varint,2,opt,name=MinGas,proto3" json:"MinGas,omitempty"`
	MovingAverageGas     int64    `protobuf:"varint,3,opt,name=MovingAverageGas,proto3" json:"MovingAverageGas,omitempty"`
	BlockNum             int64    `protobuf:"varint,4,opt,name=BlockNum,proto3" json:"BlockNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HguRecord) Reset()         { *m = HguRecord{} }
func (m *HguRecord) String() string { return proto.CompactTextString(m) }
func (*HguRecord) ProtoMessage()    {}
func (*HguRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_e71af30b0c14e8b0, []int{0}
}
func (m *HguRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HguRecord.Unmarshal(m, b)
}
func (m *HguRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HguRecord.Marshal(b, m, deterministic)
}
func (m *HguRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HguRecord.Merge(m, src)
}
func (m *HguRecord) XXX_Size() int {
	return xxx_messageInfo_HguRecord.Size(m)
}
func (m *HguRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_HguRecord.DiscardUnknown(m)
}

var xxx_messageInfo_HguRecord proto.InternalMessageInfo

func (m *HguRecord) GetMaxGas() int64 {
	if m != nil {
		return m.MaxGas
	}
	return 0
}

func (m *HguRecord) GetMinGas() int64 {
	if m != nil {
		return m.MinGas
	}
	return 0
}

func (m *HguRecord) GetMovingAverageGas() int64 {
	if m != nil {
		return m.MovingAverageGas
	}
	return 0
}

func (m *HguRecord) GetBlockNum() int64 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

func init() {
	proto.RegisterType((*HguRecord)(nil), "HguRecord")
}

func init() { proto.RegisterFile("hgu.proto", fileDescriptor_e71af30b0c14e8b0) }

var fileDescriptor_e71af30b0c14e8b0 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x48, 0x2f, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x6a, 0x66, 0xe4, 0xe2, 0xf4, 0x48, 0x2f, 0x0d, 0x4a, 0x4d,
	0xce, 0x2f, 0x4a, 0x11, 0x12, 0xe3, 0x62, 0xf3, 0x4d, 0xac, 0x70, 0x4f, 0x2c, 0x96, 0x60, 0x54,
	0x60, 0xd4, 0x60, 0x0e, 0x82, 0xf2, 0xc0, 0xe2, 0x99, 0x79, 0x20, 0x71, 0x26, 0xa8, 0x38, 0x98,
	0x27, 0xa4, 0xc5, 0x25, 0xe0, 0x9b, 0x5f, 0x96, 0x99, 0x97, 0xee, 0x58, 0x96, 0x5a, 0x94, 0x98,
	0x9e, 0x0a, 0x52, 0xc1, 0x0c, 0x56, 0x81, 0x21, 0x2e, 0x24, 0xc5, 0xc5, 0xe1, 0x94, 0x93, 0x9f,
	0x9c, 0xed, 0x57, 0x9a, 0x2b, 0xc1, 0x02, 0x56, 0x03, 0xe7, 0x3b, 0x71, 0x46, 0xb1, 0x27, 0x25,
	0x16, 0xa7, 0x26, 0x16, 0x14, 0x24, 0xb1, 0x81, 0xdd, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x82, 0xd0, 0xc9, 0x2f, 0xa4, 0x00, 0x00, 0x00,
}