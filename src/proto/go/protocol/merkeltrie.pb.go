// Code generated by protoc-gen-go. DO NOT EDIT.
// source: merkeltrie.proto

package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CHILDTYPE int32

const (
	CHILDTYPE_NONE  CHILDTYPE = 0
	CHILDTYPE_INNER CHILDTYPE = 1
	CHILDTYPE_LEAF  CHILDTYPE = 2
)

var CHILDTYPE_name = map[int32]string{
	0: "NONE",
	1: "INNER",
	2: "LEAF",
}
var CHILDTYPE_value = map[string]int32{
	"NONE":  0,
	"INNER": 1,
	"LEAF":  2,
}

func (x CHILDTYPE) String() string {
	return proto.EnumName(CHILDTYPE_name, int32(x))
}
func (CHILDTYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_merkeltrie_6c222c6531e1ca81, []int{0}
}

type Child struct {
	Sublocation          []byte    `protobuf:"bytes,1,opt,name=sublocation,proto3" json:"sublocation,omitempty"`
	Hash                 []byte    `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Childtype            CHILDTYPE `protobuf:"varint,3,opt,name=childtype,enum=protocol.CHILDTYPE" json:"childtype,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Child) Reset()         { *m = Child{} }
func (m *Child) String() string { return proto.CompactTextString(m) }
func (*Child) ProtoMessage()    {}
func (*Child) Descriptor() ([]byte, []int) {
	return fileDescriptor_merkeltrie_6c222c6531e1ca81, []int{0}
}
func (m *Child) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Child.Unmarshal(m, b)
}
func (m *Child) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Child.Marshal(b, m, deterministic)
}
func (dst *Child) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Child.Merge(dst, src)
}
func (m *Child) XXX_Size() int {
	return xxx_messageInfo_Child.Size(m)
}
func (m *Child) XXX_DiscardUnknown() {
	xxx_messageInfo_Child.DiscardUnknown(m)
}

var xxx_messageInfo_Child proto.InternalMessageInfo

func (m *Child) GetSublocation() []byte {
	if m != nil {
		return m.Sublocation
	}
	return nil
}

func (m *Child) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Child) GetChildtype() CHILDTYPE {
	if m != nil {
		return m.Childtype
	}
	return CHILDTYPE_NONE
}

type Node struct {
	Children             []*Child `protobuf:"bytes,1,rep,name=children" json:"children,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_merkeltrie_6c222c6531e1ca81, []int{1}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetChildren() []*Child {
	if m != nil {
		return m.Children
	}
	return nil
}

func init() {
	proto.RegisterType((*Child)(nil), "protocol.Child")
	proto.RegisterType((*Node)(nil), "protocol.Node")
	proto.RegisterEnum("protocol.CHILDTYPE", CHILDTYPE_name, CHILDTYPE_value)
}

func init() { proto.RegisterFile("merkeltrie.proto", fileDescriptor_merkeltrie_6c222c6531e1ca81) }

var fileDescriptor_merkeltrie_6c222c6531e1ca81 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8d, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x4d, 0x9b, 0x4a, 0x32, 0x15, 0x0d, 0xe3, 0x25, 0xc7, 0x90, 0x53, 0xa8, 0xb0, 0x60,
	0xfb, 0x04, 0x5a, 0x23, 0x16, 0xca, 0x2a, 0xc1, 0x8b, 0xc7, 0x26, 0x19, 0xc9, 0xd2, 0x6d, 0x27,
	0x6c, 0x36, 0xa0, 0x6f, 0x2f, 0xbb, 0x62, 0xed, 0x69, 0x3e, 0xbe, 0x3f, 0xbf, 0x81, 0xe4, 0x40,
	0x66, 0x4f, 0xda, 0x1a, 0x45, 0xa2, 0x37, 0x6c, 0x19, 0x23, 0x7f, 0x1a, 0xd6, 0x79, 0x0f, 0xb3,
	0x75, 0xa7, 0x74, 0x8b, 0x19, 0xcc, 0x87, 0xb1, 0xd6, 0xdc, 0xec, 0xac, 0xe2, 0x63, 0x1a, 0x64,
	0x41, 0x71, 0x55, 0x9d, 0x5b, 0x88, 0x10, 0x76, 0xbb, 0xa1, 0x4b, 0x27, 0x3e, 0xf2, 0x1a, 0xef,
	0x21, 0x6e, 0xdc, 0xdc, 0x7e, 0xf7, 0x94, 0x4e, 0xb3, 0xa0, 0xb8, 0x5e, 0xde, 0x8a, 0x3f, 0xb8,
	0x58, 0xbf, 0x6c, 0xb6, 0x4f, 0xef, 0x1f, 0x6f, 0x65, 0xf5, 0xdf, 0xca, 0x57, 0x10, 0x4a, 0x6e,
	0x09, 0xef, 0x20, 0xf2, 0xa6, 0x21, 0xf7, 0x6d, 0x5a, 0xcc, 0x97, 0x37, 0x67, 0x4b, 0x97, 0x54,
	0xa7, 0xc2, 0x62, 0x01, 0xf1, 0x09, 0x86, 0x11, 0x84, 0xf2, 0x55, 0x96, 0xc9, 0x05, 0xc6, 0x30,
	0xdb, 0x48, 0x59, 0x56, 0x49, 0xe0, 0xcc, 0x6d, 0xf9, 0xf0, 0x9c, 0x4c, 0x1e, 0x73, 0xc8, 0x14,
	0x8b, 0x7a, 0x3c, 0xb0, 0x18, 0xda, 0xbd, 0x68, 0xd8, 0x90, 0xa0, 0x2f, 0x4b, 0xc7, 0xf6, 0x97,
	0x5f, 0x8f, 0x9f, 0xf5, 0xa5, 0x57, 0xab, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0xab, 0x16,
	0x9e, 0x1b, 0x01, 0x00, 0x00,
}