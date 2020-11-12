// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/stack-labs/stack-rpc-plugins/service/gateway/api/proto/api.proto

package stack_rpc_api

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

type Pair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_254013c1e2f5907a, []int{0}
}

func (m *Pair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pair.Unmarshal(m, b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
}
func (m *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(m, src)
}
func (m *Pair) XXX_Size() int {
	return xxx_messageInfo_Pair.Size(m)
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Pair) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type Request struct {
	Method               string           `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Path                 string           `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Header               map[string]*Pair `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Get                  map[string]*Pair `protobuf:"bytes,4,rep,name=get,proto3" json:"get,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Post                 map[string]*Pair `protobuf:"bytes,5,rep,name=post,proto3" json:"post,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 string           `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
	Url                  string           `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_254013c1e2f5907a, []int{1}
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

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Request) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetGet() map[string]*Pair {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Request) GetPost() map[string]*Pair {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *Request) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Request) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Response struct {
	StatusCode           int32            `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Header               map[string]*Pair `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 string           `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_254013c1e2f5907a, []int{2}
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

func (m *Response) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Response) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Pair)(nil), "stack.rpc.api.Pair")
	proto.RegisterType((*Request)(nil), "stack.rpc.api.Request")
	proto.RegisterMapType((map[string]*Pair)(nil), "stack.rpc.api.Request.GetEntry")
	proto.RegisterMapType((map[string]*Pair)(nil), "stack.rpc.api.Request.HeaderEntry")
	proto.RegisterMapType((map[string]*Pair)(nil), "stack.rpc.api.Request.PostEntry")
	proto.RegisterType((*Response)(nil), "stack.rpc.api.Response")
	proto.RegisterMapType((map[string]*Pair)(nil), "stack.rpc.api.Response.HeaderEntry")
}

func init() {
	proto.RegisterFile("github.com/stack-labs/stack-rpc-plugins/service/gateway/api/proto/api.proto", fileDescriptor_254013c1e2f5907a)
}

var fileDescriptor_254013c1e2f5907a = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xdd, 0x6e, 0xda, 0x30,
	0x1c, 0xc5, 0x95, 0x0f, 0x02, 0xfc, 0xd1, 0xa4, 0xc9, 0x93, 0x26, 0x8b, 0x8b, 0x2d, 0xca, 0x6e,
	0xd8, 0x05, 0xc9, 0xc6, 0x76, 0x51, 0xd1, 0xcb, 0xaa, 0x6a, 0x25, 0xaa, 0x0a, 0xe5, 0x0d, 0x9c,
	0xc4, 0x4a, 0x22, 0x02, 0x76, 0x6d, 0x87, 0x2a, 0xcf, 0xd8, 0x47, 0xe9, 0x4b, 0x54, 0x76, 0x02,
	0xa5, 0x2d, 0x5c, 0xd1, 0xbb, 0x63, 0xeb, 0x9c, 0xc3, 0xe1, 0x67, 0x05, 0x16, 0x79, 0xa9, 0x8a,
	0x3a, 0x09, 0x53, 0xb6, 0x8e, 0xa4, 0x22, 0xe9, 0x6a, 0x5a, 0x91, 0x44, 0x76, 0x52, 0xf0, 0x74,
	0xca, 0xab, 0x3a, 0x2f, 0x37, 0x32, 0x92, 0x54, 0x6c, 0xcb, 0x94, 0x46, 0x39, 0x51, 0xf4, 0x91,
	0x34, 0x11, 0xe1, 0x65, 0xc4, 0x05, 0x53, 0x4c, 0xab, 0xd0, 0x28, 0xf4, 0xc5, 0xc4, 0x42, 0xc1,
	0xd3, 0x90, 0xf0, 0x32, 0xf8, 0x03, 0xee, 0x92, 0x94, 0x02, 0x7d, 0x05, 0x67, 0x45, 0x1b, 0x6c,
	0xf9, 0xd6, 0x64, 0x18, 0x6b, 0x89, 0xbe, 0x83, 0xb7, 0x25, 0x55, 0x4d, 0x25, 0xb6, 0x7d, 0x67,
	0x32, 0x8c, 0xbb, 0x53, 0xf0, 0xec, 0x40, 0x3f, 0xa6, 0x0f, 0x35, 0x95, 0x4a, 0x7b, 0xd6, 0x54,
	0x15, 0x2c, 0xeb, 0x82, 0xdd, 0x09, 0x21, 0x70, 0x39, 0x51, 0x05, 0xb6, 0xcd, 0xad, 0xd1, 0x68,
	0x0e, 0x5e, 0x41, 0x49, 0x46, 0x05, 0x76, 0x7c, 0x67, 0x32, 0x9a, 0x05, 0xe1, 0x9b, 0x25, 0x61,
	0xd7, 0x19, 0xde, 0x1a, 0xd3, 0xf5, 0x46, 0x89, 0x26, 0xee, 0x12, 0xe8, 0x2f, 0x38, 0x39, 0x55,
	0xd8, 0x35, 0xc1, 0x9f, 0x27, 0x82, 0x37, 0x54, 0xb5, 0x29, 0xed, 0x45, 0xff, 0xc1, 0xe5, 0x4c,
	0x2a, 0xdc, 0x33, 0x19, 0xff, 0x44, 0x66, 0xc9, 0x64, 0x17, 0x32, 0x6e, 0x3d, 0x3c, 0x61, 0x59,
	0x83, 0xbd, 0x76, 0xb8, 0xd6, 0x1a, 0x4d, 0x2d, 0x2a, 0xdc, 0x6f, 0xd1, 0xd4, 0xa2, 0x1a, 0xdf,
	0xc3, 0xe8, 0x60, 0xe5, 0x11, 0x76, 0xbf, 0xa1, 0x67, 0x68, 0x19, 0x00, 0xa3, 0xd9, 0xb7, 0x77,
	0xbf, 0xae, 0x89, 0xc7, 0xad, 0x63, 0x6e, 0x5f, 0x58, 0xe3, 0x05, 0x0c, 0x76, 0xe3, 0xcf, 0x2f,
	0xbb, 0x83, 0xe1, 0xfe, 0x5f, 0x9d, 0xdd, 0x16, 0x3c, 0x59, 0x30, 0x88, 0xa9, 0xe4, 0x6c, 0x23,
	0x29, 0xfa, 0x01, 0x20, 0x15, 0x51, 0xb5, 0xbc, 0x62, 0x19, 0x35, 0xa5, 0xbd, 0xf8, 0xe0, 0x06,
	0x5d, 0xee, 0x9f, 0xd8, 0x36, 0xd4, 0x7f, 0x7d, 0xa0, 0xde, 0x16, 0x1d, 0x7d, 0xe3, 0x1d, 0x7a,
	0xe7, 0x15, 0xfd, 0x67, 0x83, 0x4e, 0x3c, 0xf3, 0x0d, 0xfc, 0x7b, 0x09, 0x00, 0x00, 0xff, 0xff,
	0xe3, 0xe1, 0x9a, 0xa8, 0x52, 0x03, 0x00, 0x00,
}
