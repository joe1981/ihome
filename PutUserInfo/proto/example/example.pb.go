// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package go_micro_srv_PutUserInfo

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
	Username             string   `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
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

func (m *Request) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type Response struct {
	Errno                string   `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=Errmsg,proto3" json:"Errmsg,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
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

func (m *Response) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.PutUserInfo.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.PutUserInfo.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.PutUserInfo.Response")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x3d, 0x4b, 0xc7, 0x30,
	0x10, 0xc6, 0xfd, 0x5b, 0xec, 0xcb, 0x89, 0x20, 0x87, 0x48, 0xa9, 0x0e, 0x9a, 0xc9, 0x29, 0x83,
	0x7e, 0x04, 0xe9, 0xe0, 0x20, 0x48, 0x7d, 0xd9, 0xa3, 0x9e, 0xa5, 0x60, 0x72, 0xf5, 0xae, 0x15,
	0xfd, 0xf6, 0xd2, 0x36, 0xa8, 0x1d, 0x74, 0xcb, 0x93, 0x3c, 0xf9, 0x71, 0xbf, 0x83, 0x3d, 0xfa,
	0x70, 0xbe, 0x7f, 0x25, 0xdb, 0x0b, 0x0f, 0x8c, 0x65, 0xcb, 0xd6, 0x77, 0x4f, 0xc2, 0x56, 0xe5,
	0xdd, 0xde, 0x8c, 0xc3, 0xbd, 0x92, 0x5c, 0x85, 0x17, 0x36, 0x47, 0x90, 0x5d, 0x93, 0xaa, 0x6b,
	0x09, 0xf7, 0x21, 0x51, 0xf7, 0x59, 0x6e, 0x4e, 0x36, 0x67, 0x45, 0x33, 0x1d, 0xcd, 0x25, 0x64,
	0x0d, 0xbd, 0x8d, 0xa4, 0x03, 0x1e, 0x43, 0x71, 0x4b, 0xaa, 0x1d, 0x87, 0xee, 0x39, 0x56, 0x7e,
	0x2e, 0xb0, 0x82, 0x7c, 0x22, 0x06, 0xe7, 0xa9, 0xdc, 0x9e, 0x1f, 0xbf, 0xb3, 0xb9, 0x83, 0xbc,
	0x21, 0xed, 0x39, 0x28, 0xe1, 0x01, 0xec, 0xd4, 0x22, 0x81, 0x23, 0x61, 0x09, 0x78, 0x08, 0x69,
	0x2d, 0xe2, 0xb5, 0x8d, 0x7f, 0x63, 0x5a, 0x51, 0x93, 0x35, 0xf5, 0xdc, 0x41, 0x56, 0x2f, 0x8a,
	0xf8, 0x00, 0xbb, 0xbf, 0x8c, 0xf0, 0xd4, 0xfe, 0x25, 0x6b, 0xa3, 0x4c, 0x65, 0xfe, 0xab, 0x2c,
	0xa3, 0x9a, 0xad, 0xc7, 0x74, 0xde, 0xdd, 0xc5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x25, 0x5f,
	0x65, 0xf2, 0x4c, 0x01, 0x00, 0x00,
}