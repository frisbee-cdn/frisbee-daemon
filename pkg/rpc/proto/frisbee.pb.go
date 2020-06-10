// Code generated by protoc-gen-go. DO NOT EDIT.
// source: frisbee.proto

package models

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Node cotains ID and address
type Node struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Addr                 string   `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_ede8bc4112172419, []int{0}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Node) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Node) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type NodeResponse struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeResponse) Reset()         { *m = NodeResponse{} }
func (m *NodeResponse) String() string { return proto.CompactTextString(m) }
func (*NodeResponse) ProtoMessage()    {}
func (*NodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ede8bc4112172419, []int{1}
}

func (m *NodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeResponse.Unmarshal(m, b)
}
func (m *NodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeResponse.Marshal(b, m, deterministic)
}
func (m *NodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeResponse.Merge(m, src)
}
func (m *NodeResponse) XXX_Size() int {
	return xxx_messageInfo_NodeResponse.Size(m)
}
func (m *NodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodeResponse proto.InternalMessageInfo

func (m *NodeResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Error struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_ede8bc4112172419, []int{2}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

type ID struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_ede8bc4112172419, []int{3}
}

func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type CheckStatusRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckStatusRequest) Reset()         { *m = CheckStatusRequest{} }
func (m *CheckStatusRequest) String() string { return proto.CompactTextString(m) }
func (*CheckStatusRequest) ProtoMessage()    {}
func (*CheckStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ede8bc4112172419, []int{4}
}

func (m *CheckStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckStatusRequest.Unmarshal(m, b)
}
func (m *CheckStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckStatusRequest.Marshal(b, m, deterministic)
}
func (m *CheckStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckStatusRequest.Merge(m, src)
}
func (m *CheckStatusRequest) XXX_Size() int {
	return xxx_messageInfo_CheckStatusRequest.Size(m)
}
func (m *CheckStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckStatusRequest proto.InternalMessageInfo

func (m *CheckStatusRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CheckStatusReply struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckStatusReply) Reset()         { *m = CheckStatusReply{} }
func (m *CheckStatusReply) String() string { return proto.CompactTextString(m) }
func (*CheckStatusReply) ProtoMessage()    {}
func (*CheckStatusReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ede8bc4112172419, []int{5}
}

func (m *CheckStatusReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckStatusReply.Unmarshal(m, b)
}
func (m *CheckStatusReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckStatusReply.Marshal(b, m, deterministic)
}
func (m *CheckStatusReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckStatusReply.Merge(m, src)
}
func (m *CheckStatusReply) XXX_Size() int {
	return xxx_messageInfo_CheckStatusReply.Size(m)
}
func (m *CheckStatusReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckStatusReply.DiscardUnknown(m)
}

var xxx_messageInfo_CheckStatusReply proto.InternalMessageInfo

func (m *CheckStatusReply) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type StoreRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Content              []byte   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreRequest) Reset()         { *m = StoreRequest{} }
func (m *StoreRequest) String() string { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()    {}
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ede8bc4112172419, []int{6}
}

func (m *StoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreRequest.Unmarshal(m, b)
}
func (m *StoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreRequest.Marshal(b, m, deterministic)
}
func (m *StoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreRequest.Merge(m, src)
}
func (m *StoreRequest) XXX_Size() int {
	return xxx_messageInfo_StoreRequest.Size(m)
}
func (m *StoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreRequest proto.InternalMessageInfo

func (m *StoreRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StoreRequest) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type StorageResponse struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StorageResponse) Reset()         { *m = StorageResponse{} }
func (m *StorageResponse) String() string { return proto.CompactTextString(m) }
func (*StorageResponse) ProtoMessage()    {}
func (*StorageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ede8bc4112172419, []int{7}
}

func (m *StorageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageResponse.Unmarshal(m, b)
}
func (m *StorageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageResponse.Marshal(b, m, deterministic)
}
func (m *StorageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageResponse.Merge(m, src)
}
func (m *StorageResponse) XXX_Size() int {
	return xxx_messageInfo_StorageResponse.Size(m)
}
func (m *StorageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StorageResponse proto.InternalMessageInfo

func (m *StorageResponse) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "models.Node")
	proto.RegisterType((*NodeResponse)(nil), "models.NodeResponse")
	proto.RegisterType((*Error)(nil), "models.Error")
	proto.RegisterType((*ID)(nil), "models.ID")
	proto.RegisterType((*CheckStatusRequest)(nil), "models.CheckStatusRequest")
	proto.RegisterType((*CheckStatusReply)(nil), "models.CheckStatusReply")
	proto.RegisterType((*StoreRequest)(nil), "models.StoreRequest")
	proto.RegisterType((*StorageResponse)(nil), "models.StorageResponse")
}

func init() {
	proto.RegisterFile("frisbee.proto", fileDescriptor_ede8bc4112172419)
}

var fileDescriptor_ede8bc4112172419 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x9b, 0xf4, 0x1f, 0x3d, 0x52, 0xa8, 0xac, 0x0a, 0xa2, 0x4c, 0x91, 0xa7, 0x08, 0xa4,
	0x0c, 0x61, 0x63, 0xe8, 0x42, 0xa9, 0xd4, 0x05, 0xa1, 0x54, 0x62, 0x4f, 0xeb, 0xa3, 0x44, 0x4d,
	0xe3, 0x60, 0xbb, 0x43, 0x3f, 0x2c, 0xdf, 0x05, 0xd9, 0x89, 0x45, 0x02, 0x6c, 0xf7, 0x7c, 0xbf,
	0xcb, 0xcb, 0x3b, 0x1b, 0xa6, 0xef, 0x22, 0x97, 0x5b, 0xc4, 0xb8, 0x12, 0x5c, 0x71, 0x32, 0x3a,
	0x72, 0x86, 0x85, 0xa4, 0x0b, 0x18, 0xbc, 0x70, 0x86, 0xe4, 0x0a, 0xdc, 0x9c, 0xf9, 0x4e, 0xe8,
	0x44, 0x5e, 0xea, 0xe6, 0x8c, 0x10, 0x18, 0x54, 0x5c, 0x28, 0xdf, 0x0d, 0x9d, 0x68, 0x9a, 0x9a,
	0x5a, 0x9f, 0x65, 0x8c, 0x09, 0xbf, 0x1f, 0x3a, 0xd1, 0x24, 0x35, 0x35, 0x4d, 0xc0, 0xd3, 0xf3,
	0x29, 0xca, 0x8a, 0x97, 0x12, 0x09, 0x85, 0x61, 0xc9, 0x19, 0x4a, 0xdf, 0x09, 0xfb, 0xd1, 0x65,
	0xe2, 0xc5, 0xb5, 0x4f, 0x6c, 0xa0, 0xba, 0x45, 0xc7, 0x30, 0x7c, 0x16, 0x82, 0x0b, 0x3a, 0x07,
	0x77, 0xbd, 0xfc, 0x6d, 0x4d, 0x63, 0x20, 0x4f, 0x1f, 0xb8, 0x3b, 0x6c, 0x54, 0xa6, 0x4e, 0x32,
	0xc5, 0xcf, 0x13, 0x4a, 0x45, 0x7c, 0x18, 0x1f, 0x51, 0xca, 0x6c, 0x8f, 0x06, 0x9d, 0xa4, 0x56,
	0xd2, 0x3b, 0x98, 0x75, 0xf8, 0xaa, 0x38, 0x93, 0x1b, 0x18, 0x49, 0x23, 0x1b, 0xb8, 0x51, 0xf4,
	0x11, 0xbc, 0x8d, 0xe2, 0x02, 0xed, 0x57, 0x67, 0xd0, 0x3f, 0xe0, 0xb9, 0x81, 0x74, 0xa9, 0x7d,
	0x76, 0xbc, 0x54, 0x58, 0xd6, 0xd9, 0xbd, 0xd4, 0x4a, 0x7a, 0x0f, 0xd7, 0x7a, 0x36, 0xdb, 0xff,
	0xa4, 0x6d, 0xc1, 0x4e, 0x07, 0x4e, 0xbe, 0x1c, 0x18, 0xaf, 0xea, 0x8d, 0x93, 0x05, 0x0c, 0x5e,
	0xf3, 0x72, 0x4f, 0x02, 0xbb, 0x8c, 0xbf, 0xf1, 0x02, 0xff, 0xdf, 0x5e, 0x55, 0x9c, 0x69, 0x8f,
	0xc4, 0x30, 0x34, 0x3f, 0x4d, 0xe6, 0x16, 0x6a, 0x67, 0x08, 0xa6, 0xf6, 0xb4, 0x5e, 0xaa, 0xe6,
	0x2f, 0x56, 0x79, 0xc9, 0xcc, 0xbd, 0x82, 0x6d, 0xae, 0x97, 0xc1, 0xbc, 0x73, 0x19, 0x4d, 0x06,
	0xda, 0x23, 0x09, 0x4c, 0x34, 0xff, 0x96, 0x15, 0xa7, 0xee, 0xc0, 0x6d, 0xdb, 0xaf, 0x95, 0x9b,
	0xf6, 0xb6, 0x23, 0xf3, 0x8c, 0x1e, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x90, 0xcc, 0x55, 0xbd,
	0x57, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FrisbeeClient is the client API for Frisbee service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FrisbeeClient interface {
	Ping(ctx context.Context, in *CheckStatusRequest, opts ...grpc.CallOption) (*CheckStatusReply, error)
	Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*Error, error)
	FindNode(ctx context.Context, in *ID, opts ...grpc.CallOption) (*NodeResponse, error)
	FindValue(ctx context.Context, in *ID, opts ...grpc.CallOption) (*StorageResponse, error)
}

type frisbeeClient struct {
	cc grpc.ClientConnInterface
}

func NewFrisbeeClient(cc grpc.ClientConnInterface) FrisbeeClient {
	return &frisbeeClient{cc}
}

func (c *frisbeeClient) Ping(ctx context.Context, in *CheckStatusRequest, opts ...grpc.CallOption) (*CheckStatusReply, error) {
	out := new(CheckStatusReply)
	err := c.cc.Invoke(ctx, "/models.Frisbee/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frisbeeClient) Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/models.Frisbee/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frisbeeClient) FindNode(ctx context.Context, in *ID, opts ...grpc.CallOption) (*NodeResponse, error) {
	out := new(NodeResponse)
	err := c.cc.Invoke(ctx, "/models.Frisbee/FindNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frisbeeClient) FindValue(ctx context.Context, in *ID, opts ...grpc.CallOption) (*StorageResponse, error) {
	out := new(StorageResponse)
	err := c.cc.Invoke(ctx, "/models.Frisbee/FindValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FrisbeeServer is the server API for Frisbee service.
type FrisbeeServer interface {
	Ping(context.Context, *CheckStatusRequest) (*CheckStatusReply, error)
	Store(context.Context, *StoreRequest) (*Error, error)
	FindNode(context.Context, *ID) (*NodeResponse, error)
	FindValue(context.Context, *ID) (*StorageResponse, error)
}

// UnimplementedFrisbeeServer can be embedded to have forward compatible implementations.
type UnimplementedFrisbeeServer struct {
}

func (*UnimplementedFrisbeeServer) Ping(ctx context.Context, req *CheckStatusRequest) (*CheckStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedFrisbeeServer) Store(ctx context.Context, req *StoreRequest) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (*UnimplementedFrisbeeServer) FindNode(ctx context.Context, req *ID) (*NodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNode not implemented")
}
func (*UnimplementedFrisbeeServer) FindValue(ctx context.Context, req *ID) (*StorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindValue not implemented")
}

func RegisterFrisbeeServer(s *grpc.Server, srv FrisbeeServer) {
	s.RegisterService(&_Frisbee_serviceDesc, srv)
}

func _Frisbee_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrisbeeServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Frisbee/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrisbeeServer).Ping(ctx, req.(*CheckStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Frisbee_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrisbeeServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Frisbee/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrisbeeServer).Store(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Frisbee_FindNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrisbeeServer).FindNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Frisbee/FindNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrisbeeServer).FindNode(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Frisbee_FindValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrisbeeServer).FindValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Frisbee/FindValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrisbeeServer).FindValue(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Frisbee_serviceDesc = grpc.ServiceDesc{
	ServiceName: "models.Frisbee",
	HandlerType: (*FrisbeeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Frisbee_Ping_Handler,
		},
		{
			MethodName: "Store",
			Handler:    _Frisbee_Store_Handler,
		},
		{
			MethodName: "FindNode",
			Handler:    _Frisbee_FindNode_Handler,
		},
		{
			MethodName: "FindValue",
			Handler:    _Frisbee_FindValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "frisbee.proto",
}
