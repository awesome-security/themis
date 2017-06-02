// Code generated by protoc-gen-go. DO NOT EDIT.
// source: control.proto

/*
Package control is a generated protocol buffer package.

It is generated from these files:
	control.proto

It has these top-level messages:
	Item
	Chunk
	Update
	Response
*/
package control

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Item_DataType int32

const (
	Item_POLICIES Item_DataType = 0
	Item_CONTENT  Item_DataType = 1
)

var Item_DataType_name = map[int32]string{
	0: "POLICIES",
	1: "CONTENT",
}
var Item_DataType_value = map[string]int32{
	"POLICIES": 0,
	"CONTENT":  1,
}

func (x Item_DataType) String() string {
	return proto.EnumName(Item_DataType_name, int32(x))
}
func (Item_DataType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Response_Status int32

const (
	Response_ACK           Response_Status = 0
	Response_ERROR         Response_Status = 1
	Response_VERSION_ERROR Response_Status = 2
)

var Response_Status_name = map[int32]string{
	0: "ACK",
	1: "ERROR",
	2: "VERSION_ERROR",
}
var Response_Status_value = map[string]int32{
	"ACK":           0,
	"ERROR":         1,
	"VERSION_ERROR": 2,
}

func (x Response_Status) String() string {
	return proto.EnumName(Response_Status_name, int32(x))
}
func (Response_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type Item struct {
	Type        Item_DataType `protobuf:"varint,1,opt,name=type,enum=control.Item_DataType" json:"type,omitempty"`
	FromVersion string        `protobuf:"bytes,2,opt,name=fromVersion" json:"fromVersion,omitempty"`
	ToVersion   string        `protobuf:"bytes,3,opt,name=toVersion" json:"toVersion,omitempty"`
	DataId      int32         `protobuf:"varint,4,opt,name=dataId" json:"dataId,omitempty"`
	Id          string        `protobuf:"bytes,5,opt,name=id" json:"id,omitempty"`
	Includes    []int32       `protobuf:"varint,6,rep,packed,name=includes" json:"includes,omitempty"`
}

func (m *Item) Reset()                    { *m = Item{} }
func (m *Item) String() string            { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()               {}
func (*Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Item) GetType() Item_DataType {
	if m != nil {
		return m.Type
	}
	return Item_POLICIES
}

func (m *Item) GetFromVersion() string {
	if m != nil {
		return m.FromVersion
	}
	return ""
}

func (m *Item) GetToVersion() string {
	if m != nil {
		return m.ToVersion
	}
	return ""
}

func (m *Item) GetDataId() int32 {
	if m != nil {
		return m.DataId
	}
	return 0
}

func (m *Item) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Item) GetIncludes() []int32 {
	if m != nil {
		return m.Includes
	}
	return nil
}

type Chunk struct {
	Data string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *Chunk) Reset()                    { *m = Chunk{} }
func (m *Chunk) String() string            { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()               {}
func (*Chunk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Chunk) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type Update struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *Update) Reset()                    { *m = Update{} }
func (m *Update) String() string            { return proto.CompactTextString(m) }
func (*Update) ProtoMessage()               {}
func (*Update) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Update) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Response struct {
	Status  Response_Status `protobuf:"varint,1,opt,name=status,enum=control.Response_Status" json:"status,omitempty"`
	Id      int32           `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Details string          `protobuf:"bytes,3,opt,name=details" json:"details,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetStatus() Response_Status {
	if m != nil {
		return m.Status
	}
	return Response_ACK
}

func (m *Response) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Response) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func init() {
	proto.RegisterType((*Item)(nil), "control.Item")
	proto.RegisterType((*Chunk)(nil), "control.Chunk")
	proto.RegisterType((*Update)(nil), "control.Update")
	proto.RegisterType((*Response)(nil), "control.Response")
	proto.RegisterEnum("control.Item_DataType", Item_DataType_name, Item_DataType_value)
	proto.RegisterEnum("control.Response_Status", Response_Status_name, Response_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PDPControl service

type PDPControlClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (PDPControl_UploadClient, error)
	Parse(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Response, error)
	Apply(ctx context.Context, in *Update, opts ...grpc.CallOption) (*Response, error)
}

type pDPControlClient struct {
	cc *grpc.ClientConn
}

func NewPDPControlClient(cc *grpc.ClientConn) PDPControlClient {
	return &pDPControlClient{cc}
}

func (c *pDPControlClient) Upload(ctx context.Context, opts ...grpc.CallOption) (PDPControl_UploadClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PDPControl_serviceDesc.Streams[0], c.cc, "/control.PDPControl/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &pDPControlUploadClient{stream}
	return x, nil
}

type PDPControl_UploadClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type pDPControlUploadClient struct {
	grpc.ClientStream
}

func (x *pDPControlUploadClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pDPControlUploadClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pDPControlClient) Parse(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/control.PDPControl/Parse", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pDPControlClient) Apply(ctx context.Context, in *Update, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/control.PDPControl/Apply", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PDPControl service

type PDPControlServer interface {
	Upload(PDPControl_UploadServer) error
	Parse(context.Context, *Item) (*Response, error)
	Apply(context.Context, *Update) (*Response, error)
}

func RegisterPDPControlServer(s *grpc.Server, srv PDPControlServer) {
	s.RegisterService(&_PDPControl_serviceDesc, srv)
}

func _PDPControl_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PDPControlServer).Upload(&pDPControlUploadServer{stream})
}

type PDPControl_UploadServer interface {
	SendAndClose(*Response) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type pDPControlUploadServer struct {
	grpc.ServerStream
}

func (x *pDPControlUploadServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pDPControlUploadServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PDPControl_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PDPControlServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control.PDPControl/Parse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PDPControlServer).Parse(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _PDPControl_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Update)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PDPControlServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control.PDPControl/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PDPControlServer).Apply(ctx, req.(*Update))
	}
	return interceptor(ctx, in, info, handler)
}

var _PDPControl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "control.PDPControl",
	HandlerType: (*PDPControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Parse",
			Handler:    _PDPControl_Parse_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _PDPControl_Apply_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _PDPControl_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "control.proto",
}

func init() { proto.RegisterFile("control.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0xad, 0xd3, 0x26, 0x6d, 0x66, 0x69, 0xc9, 0xce, 0x61, 0x65, 0x2d, 0x1c, 0xa2, 0x48, 0x48,
	0x11, 0x88, 0x2c, 0x5a, 0x7e, 0xc1, 0x2a, 0x9b, 0x43, 0x04, 0x6a, 0x22, 0xb7, 0xec, 0x15, 0x99,
	0xb5, 0x11, 0x11, 0xd9, 0x38, 0x8a, 0xdd, 0x43, 0x7f, 0x09, 0xe2, 0xc7, 0xf1, 0x5f, 0x50, 0x9d,
	0x0f, 0x40, 0xf4, 0xe6, 0x37, 0xef, 0x79, 0xe6, 0xcd, 0xb3, 0x61, 0xfd, 0xa8, 0x1a, 0xd3, 0xa9,
	0x3a, 0x69, 0x3b, 0x65, 0x14, 0x2e, 0x07, 0x18, 0xfd, 0x22, 0xb0, 0xc8, 0x8d, 0x7c, 0xc2, 0xd7,
	0xb0, 0x30, 0xc7, 0x56, 0x52, 0x12, 0x92, 0x78, 0x73, 0x7b, 0x95, 0x8c, 0xfa, 0x13, 0x99, 0xdc,
	0x73, 0xc3, 0xf7, 0xc7, 0x56, 0x32, 0xab, 0xc1, 0x10, 0x2e, 0xbe, 0x76, 0xea, 0xe9, 0x41, 0x76,
	0xba, 0x52, 0x0d, 0x75, 0x42, 0x12, 0xfb, 0xec, 0xef, 0x12, 0xbe, 0x04, 0xdf, 0xa8, 0x91, 0x9f,
	0x5b, 0xfe, 0x4f, 0x01, 0xaf, 0xc0, 0x13, 0xdc, 0xf0, 0x5c, 0xd0, 0x45, 0x48, 0x62, 0x97, 0x0d,
	0x08, 0x37, 0xe0, 0x54, 0x82, 0xba, 0x56, 0xee, 0x54, 0x02, 0xaf, 0x61, 0x55, 0x35, 0x8f, 0xf5,
	0x41, 0x48, 0x4d, 0xbd, 0x70, 0x1e, 0xbb, 0x6c, 0xc2, 0xd1, 0x2b, 0x58, 0x8d, 0xae, 0xf0, 0x19,
	0xac, 0xca, 0xe2, 0x63, 0x9e, 0xe6, 0xd9, 0x2e, 0x98, 0xe1, 0x05, 0x2c, 0xd3, 0x62, 0xbb, 0xcf,
	0xb6, 0xfb, 0x80, 0x44, 0x2f, 0xc0, 0x4d, 0xbf, 0x1d, 0x9a, 0xef, 0x88, 0xb0, 0x38, 0x4d, 0xb1,
	0xfb, 0xf9, 0xcc, 0x9e, 0x23, 0x0a, 0xde, 0xa7, 0x56, 0x70, 0x23, 0x87, 0xc9, 0xc4, 0xba, 0x71,
	0x2a, 0x11, 0xfd, 0x20, 0xb0, 0x62, 0x52, 0xb7, 0xaa, 0xd1, 0x12, 0xdf, 0x81, 0xa7, 0x0d, 0x37,
	0x07, 0x3d, 0x84, 0x43, 0xa7, 0x70, 0x46, 0x49, 0xb2, 0xb3, 0x3c, 0x1b, 0x74, 0x43, 0x3b, 0x67,
	0x6c, 0x87, 0x14, 0x96, 0x42, 0x1a, 0x5e, 0xd5, 0x7a, 0x08, 0x63, 0x84, 0xd1, 0x0d, 0x78, 0xfd,
	0x5d, 0x5c, 0xc2, 0xfc, 0x2e, 0xfd, 0x10, 0xcc, 0xd0, 0x07, 0x37, 0x63, 0xac, 0x60, 0x01, 0xc1,
	0x4b, 0x58, 0x3f, 0x64, 0x6c, 0x97, 0x17, 0xdb, 0xcf, 0x7d, 0xc9, 0xb9, 0xfd, 0x49, 0x00, 0xca,
	0xfb, 0x32, 0xed, 0x1d, 0xe0, 0xcd, 0x69, 0x85, 0x5a, 0x71, 0x81, 0x9b, 0xc9, 0x95, 0x5d, 0xf8,
	0xfa, 0xf2, 0x3f, 0x97, 0xd1, 0x2c, 0x26, 0xf8, 0x06, 0xdc, 0x92, 0x77, 0x5a, 0xe2, 0xfa, 0x9f,
	0x27, 0x3e, 0x2b, 0xc7, 0xb7, 0xe0, 0xde, 0xb5, 0x6d, 0x7d, 0xc4, 0xe7, 0x13, 0xdb, 0x07, 0x76,
	0x56, 0xfe, 0xc5, 0xb3, 0x9f, 0xeb, 0xfd, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x4b, 0xf9,
	0x93, 0x6d, 0x02, 0x00, 0x00,
}
