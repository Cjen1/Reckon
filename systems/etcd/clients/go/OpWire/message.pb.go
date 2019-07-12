// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package OpWire

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

type Operation struct {
	// Types that are valid to be assigned to OpType:
	//	*Operation_Put
	//	*Operation_Get
	//	*Operation_Quit
	OpType               isOperation_OpType `protobuf_oneof:"op_type"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}
func (*Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operation.Unmarshal(m, b)
}
func (m *Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operation.Marshal(b, m, deterministic)
}
func (m *Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation.Merge(m, src)
}
func (m *Operation) XXX_Size() int {
	return xxx_messageInfo_Operation.Size(m)
}
func (m *Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Operation proto.InternalMessageInfo

type isOperation_OpType interface {
	isOperation_OpType()
}

type Operation_Put struct {
	Put *OperationOpPut `protobuf:"bytes,1,opt,name=put,proto3,oneof"`
}

type Operation_Get struct {
	Get *OperationOpGet `protobuf:"bytes,2,opt,name=get,proto3,oneof"`
}

type Operation_Quit struct {
	Quit *OperationOpQuit `protobuf:"bytes,3,opt,name=quit,proto3,oneof"`
}

func (*Operation_Put) isOperation_OpType() {}

func (*Operation_Get) isOperation_OpType() {}

func (*Operation_Quit) isOperation_OpType() {}

func (m *Operation) GetOpType() isOperation_OpType {
	if m != nil {
		return m.OpType
	}
	return nil
}

func (m *Operation) GetPut() *OperationOpPut {
	if x, ok := m.GetOpType().(*Operation_Put); ok {
		return x.Put
	}
	return nil
}

func (m *Operation) GetGet() *OperationOpGet {
	if x, ok := m.GetOpType().(*Operation_Get); ok {
		return x.Get
	}
	return nil
}

func (m *Operation) GetQuit() *OperationOpQuit {
	if x, ok := m.GetOpType().(*Operation_Quit); ok {
		return x.Quit
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Operation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Operation_Put)(nil),
		(*Operation_Get)(nil),
		(*Operation_Quit)(nil),
	}
}

type OperationOpPut struct {
	Key                  uint64   `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Opid                 uint64   `protobuf:"varint,3,opt,name=opid,proto3" json:"opid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationOpPut) Reset()         { *m = OperationOpPut{} }
func (m *OperationOpPut) String() string { return proto.CompactTextString(m) }
func (*OperationOpPut) ProtoMessage()    {}
func (*OperationOpPut) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0}
}

func (m *OperationOpPut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationOpPut.Unmarshal(m, b)
}
func (m *OperationOpPut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationOpPut.Marshal(b, m, deterministic)
}
func (m *OperationOpPut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationOpPut.Merge(m, src)
}
func (m *OperationOpPut) XXX_Size() int {
	return xxx_messageInfo_OperationOpPut.Size(m)
}
func (m *OperationOpPut) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationOpPut.DiscardUnknown(m)
}

var xxx_messageInfo_OperationOpPut proto.InternalMessageInfo

func (m *OperationOpPut) GetKey() uint64 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *OperationOpPut) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *OperationOpPut) GetOpid() uint64 {
	if m != nil {
		return m.Opid
	}
	return 0
}

type OperationOpGet struct {
	Key                  uint64   `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Opid                 uint64   `protobuf:"varint,3,opt,name=opid,proto3" json:"opid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationOpGet) Reset()         { *m = OperationOpGet{} }
func (m *OperationOpGet) String() string { return proto.CompactTextString(m) }
func (*OperationOpGet) ProtoMessage()    {}
func (*OperationOpGet) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 1}
}

func (m *OperationOpGet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationOpGet.Unmarshal(m, b)
}
func (m *OperationOpGet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationOpGet.Marshal(b, m, deterministic)
}
func (m *OperationOpGet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationOpGet.Merge(m, src)
}
func (m *OperationOpGet) XXX_Size() int {
	return xxx_messageInfo_OperationOpGet.Size(m)
}
func (m *OperationOpGet) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationOpGet.DiscardUnknown(m)
}

var xxx_messageInfo_OperationOpGet proto.InternalMessageInfo

func (m *OperationOpGet) GetKey() uint64 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *OperationOpGet) GetOpid() uint64 {
	if m != nil {
		return m.Opid
	}
	return 0
}

type OperationOpQuit struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationOpQuit) Reset()         { *m = OperationOpQuit{} }
func (m *OperationOpQuit) String() string { return proto.CompactTextString(m) }
func (*OperationOpQuit) ProtoMessage()    {}
func (*OperationOpQuit) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 2}
}

func (m *OperationOpQuit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationOpQuit.Unmarshal(m, b)
}
func (m *OperationOpQuit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationOpQuit.Marshal(b, m, deterministic)
}
func (m *OperationOpQuit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationOpQuit.Merge(m, src)
}
func (m *OperationOpQuit) XXX_Size() int {
	return xxx_messageInfo_OperationOpQuit.Size(m)
}
func (m *OperationOpQuit) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationOpQuit.DiscardUnknown(m)
}

var xxx_messageInfo_OperationOpQuit proto.InternalMessageInfo

func (m *OperationOpQuit) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Response struct {
	ResponseTime         float64  `protobuf:"fixed64,1,opt,name=response_time,json=responseTime,proto3" json:"response_time,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	Start                float64  `protobuf:"fixed64,3,opt,name=start,proto3" json:"start,omitempty"`
	End                  float64  `protobuf:"fixed64,4,opt,name=end,proto3" json:"end,omitempty"`
	Clientid             uint32   `protobuf:"varint,5,opt,name=clientid,proto3" json:"clientid,omitempty"`
	Opid                 uint64   `protobuf:"varint,6,opt,name=opid,proto3" json:"opid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
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

func (m *Response) GetResponseTime() float64 {
	if m != nil {
		return m.ResponseTime
	}
	return 0
}

func (m *Response) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func (m *Response) GetStart() float64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Response) GetEnd() float64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *Response) GetClientid() uint32 {
	if m != nil {
		return m.Clientid
	}
	return 0
}

func (m *Response) GetOpid() uint64 {
	if m != nil {
		return m.Opid
	}
	return 0
}

func init() {
	proto.RegisterType((*Operation)(nil), "OpWire.Operation")
	proto.RegisterType((*OperationOpPut)(nil), "OpWire.Operation.op_put")
	proto.RegisterType((*OperationOpGet)(nil), "OpWire.Operation.op_get")
	proto.RegisterType((*OperationOpQuit)(nil), "OpWire.Operation.op_quit")
	proto.RegisterType((*Response)(nil), "OpWire.Response")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0x9b, 0xc6, 0x66, 0x6c, 0x40, 0x16, 0x0f, 0x31, 0x5e, 0xa4, 0x5e, 0x3c, 0x48,
	0x04, 0x7d, 0x83, 0xe2, 0xc1, 0x5b, 0x61, 0x11, 0x3c, 0x96, 0x68, 0x87, 0xb0, 0xd8, 0x64, 0xd7,
	0xdd, 0x8d, 0xd0, 0x47, 0xf1, 0x45, 0x7c, 0x3e, 0x77, 0x26, 0x31, 0x1e, 0x2c, 0x9e, 0xf2, 0xcf,
	0xcc, 0x37, 0x33, 0xff, 0x64, 0x21, 0x6b, 0xd0, 0xb9, 0xaa, 0xc6, 0xd2, 0x58, 0xed, 0xb5, 0x48,
	0xd6, 0xe6, 0x59, 0x59, 0x5c, 0x7e, 0x4d, 0x20, 0x5d, 0x1b, 0xb4, 0x95, 0x57, 0xba, 0x15, 0x37,
	0x30, 0x35, 0x9d, 0xcf, 0xa3, 0xcb, 0xe8, 0xfa, 0xe4, 0x2e, 0x2f, 0x7b, 0xa6, 0x1c, 0xeb, 0xa5,
	0x36, 0x9b, 0x50, 0x7f, 0x3c, 0x92, 0x84, 0x11, 0x5d, 0xa3, 0xcf, 0x27, 0xff, 0xd0, 0xa1, 0x4e,
	0x74, 0xf8, 0x88, 0x5b, 0x88, 0xdf, 0x3b, 0xe5, 0xf3, 0x29, 0xe3, 0xe7, 0x07, 0x71, 0x02, 0x02,
	0xcf, 0x60, 0xf1, 0x00, 0x49, 0xbf, 0x4f, 0x9c, 0xc2, 0xf4, 0x0d, 0xf7, 0x6c, 0x2b, 0x96, 0x24,
	0xc5, 0x19, 0xcc, 0x3e, 0xaa, 0x5d, 0x87, 0xbc, 0x7c, 0x21, 0xfb, 0x40, 0x08, 0x88, 0xb5, 0x51,
	0x5b, 0x5e, 0x11, 0x4b, 0xd6, 0x45, 0xc9, 0x53, 0xc8, 0xc0, 0xdf, 0x29, 0x87, 0xf8, 0x0b, 0x38,
	0x1e, 0x8c, 0x50, 0x43, 0xe3, 0x6a, 0x6e, 0x48, 0x25, 0xc9, 0x55, 0xca, 0x45, 0xbf, 0x37, 0xb8,
	0xfc, 0x8c, 0x60, 0x2e, 0xd1, 0x19, 0xdd, 0x3a, 0x14, 0x57, 0x90, 0xd9, 0x41, 0x6f, 0xbc, 0x6a,
	0x90, 0x7b, 0x22, 0xb9, 0xf8, 0x49, 0x3e, 0x85, 0x1c, 0x8d, 0x43, 0x6b, 0xd9, 0x71, 0x18, 0x17,
	0x24, 0x5d, 0xe1, 0x7c, 0x65, 0xfb, 0x7f, 0x12, 0xc9, 0x3e, 0x60, 0xae, 0xdd, 0xe6, 0x31, 0xe7,
	0x48, 0x8a, 0x02, 0xe6, 0xaf, 0x3b, 0x85, 0xad, 0x0f, 0x5e, 0x67, 0x21, 0x9d, 0xc9, 0x31, 0x1e,
	0x6f, 0x48, 0x7e, 0x6f, 0x58, 0xcd, 0x61, 0x78, 0xde, 0x97, 0x84, 0x5f, 0xfb, 0xfe, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0xdc, 0x55, 0xa6, 0x46, 0xfe, 0x01, 0x00, 0x00,
}
