// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package go_micro_srv_PutOrders

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Sessionid            string   `protobuf:"bytes,1,opt,name=Sessionid,proto3" json:"Sessionid,omitempty"`
	Orderid              string   `protobuf:"bytes,2,opt,name=Orderid,proto3" json:"Orderid,omitempty"`
	Action               string   `protobuf:"bytes,3,opt,name=Action,proto3" json:"Action,omitempty"`
	Reason               string   `protobuf:"bytes,4,opt,name=Reason,proto3" json:"Reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetSessionid() string {
	if m != nil {
		return m.Sessionid
	}
	return ""
}

func (m *Request) GetOrderid() string {
	if m != nil {
		return m.Orderid
	}
	return ""
}

func (m *Request) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Request) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type Response struct {
	Errno                string   `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=Errmsg,proto3" json:"Errmsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.PutOrders.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.PutOrders.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.PutOrders.Response")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4f, 0xc4, 0x30,
	0x0c, 0x46, 0x39, 0x0e, 0xae, 0xd4, 0x12, 0x12, 0xb2, 0xd0, 0x29, 0x02, 0x24, 0x4e, 0x9d, 0x98,
	0x32, 0xc0, 0xc2, 0xca, 0xd0, 0x11, 0x71, 0x2a, 0x23, 0x53, 0xb9, 0x5a, 0x55, 0x24, 0x1a, 0xf7,
	0xec, 0x14, 0xc1, 0xbf, 0x47, 0x4d, 0x02, 0x2c, 0xb0, 0xe5, 0x7d, 0x76, 0xe2, 0x3c, 0xc3, 0x29,
	0x7d, 0xb4, 0xc3, 0xf8, 0x46, 0x76, 0x14, 0x0e, 0x8c, 0xeb, 0x9e, 0xed, 0xe0, 0x76, 0xc2, 0x56,
	0xe5, 0xdd, 0x6e, 0xa7, 0xf0, 0x24, 0x1d, 0x89, 0x56, 0x97, 0x50, 0x3c, 0x92, 0x6a, 0xdb, 0x13,
	0x9e, 0xc1, 0x52, 0xdb, 0x4f, 0xb3, 0xd8, 0x2c, 0x6e, 0xca, 0x66, 0x3e, 0x56, 0x7b, 0x28, 0x1a,
	0xda, 0x4f, 0xa4, 0x01, 0xaf, 0xa0, 0x7c, 0x26, 0x55, 0xc7, 0xde, 0x75, 0xb9, 0xe5, 0x37, 0x40,
	0x03, 0x45, 0x7c, 0xcf, 0x75, 0xe6, 0x30, 0xd6, 0xbe, 0x11, 0xd7, 0xb0, 0x7a, 0xd8, 0x05, 0xc7,
	0xde, 0x2c, 0x63, 0x21, 0xd3, 0x9c, 0x37, 0xd4, 0x2a, 0x7b, 0x73, 0x94, 0xf2, 0x44, 0xd5, 0x3d,
	0x9c, 0x34, 0xa4, 0x23, 0x7b, 0x25, 0x3c, 0x87, 0xe3, 0x5a, 0xc4, 0x73, 0x9e, 0x97, 0x60, 0xbe,
	0x59, 0x8b, 0x0c, 0xda, 0xe7, 0x51, 0x99, 0x6e, 0x5f, 0xa0, 0xa8, 0x93, 0x32, 0x6e, 0xa1, 0xfc,
	0x31, 0xc4, 0x6b, 0xfb, 0xb7, 0xba, 0xcd, 0x6a, 0x17, 0x9b, 0xff, 0x1b, 0xd2, 0x47, 0xaa, 0x83,
	0xd7, 0x55, 0xdc, 0xe2, 0xdd, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0xe0, 0xf0, 0x0e, 0x56,
	0x01, 0x00, 0x00,
}
