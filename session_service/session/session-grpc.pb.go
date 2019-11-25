// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session-grpc.proto

package session

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

type SessionId struct {
	SessionID            string   `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionId) Reset()         { *m = SessionId{} }
func (m *SessionId) String() string { return proto.CompactTextString(m) }
func (*SessionId) ProtoMessage()    {}
func (*SessionId) Descriptor() ([]byte, []int) {
	return fileDescriptor_33bdf83d440e1b7e, []int{0}
}

func (m *SessionId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionId.Unmarshal(m, b)
}
func (m *SessionId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionId.Marshal(b, m, deterministic)
}
func (m *SessionId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionId.Merge(m, src)
}
func (m *SessionId) XXX_Size() int {
	return xxx_messageInfo_SessionId.Size(m)
}
func (m *SessionId) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionId.DiscardUnknown(m)
}

var xxx_messageInfo_SessionId proto.InternalMessageInfo

func (m *SessionId) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

type UserID struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserID) Reset()         { *m = UserID{} }
func (m *UserID) String() string { return proto.CompactTextString(m) }
func (*UserID) ProtoMessage()    {}
func (*UserID) Descriptor() ([]byte, []int) {
	return fileDescriptor_33bdf83d440e1b7e, []int{1}
}

func (m *UserID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserID.Unmarshal(m, b)
}
func (m *UserID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserID.Marshal(b, m, deterministic)
}
func (m *UserID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserID.Merge(m, src)
}
func (m *UserID) XXX_Size() int {
	return xxx_messageInfo_UserID.Size(m)
}
func (m *UserID) XXX_DiscardUnknown() {
	xxx_messageInfo_UserID.DiscardUnknown(m)
}

var xxx_messageInfo_UserID proto.InternalMessageInfo

func (m *UserID) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type Session struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID               string   `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	CsrfToken            string   `protobuf:"bytes,3,opt,name=CsrfToken,proto3" json:"CsrfToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_33bdf83d440e1b7e, []int{2}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Session) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Session) GetCsrfToken() string {
	if m != nil {
		return m.CsrfToken
	}
	return ""
}

type SessionInfo struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID               string   `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionInfo) Reset()         { *m = SessionInfo{} }
func (m *SessionInfo) String() string { return proto.CompactTextString(m) }
func (*SessionInfo) ProtoMessage()    {}
func (*SessionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_33bdf83d440e1b7e, []int{3}
}

func (m *SessionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionInfo.Unmarshal(m, b)
}
func (m *SessionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionInfo.Marshal(b, m, deterministic)
}
func (m *SessionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionInfo.Merge(m, src)
}
func (m *SessionInfo) XXX_Size() int {
	return xxx_messageInfo_SessionInfo.Size(m)
}
func (m *SessionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SessionInfo proto.InternalMessageInfo

func (m *SessionInfo) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *SessionInfo) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type Nothing struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nothing) Reset()         { *m = Nothing{} }
func (m *Nothing) String() string { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()    {}
func (*Nothing) Descriptor() ([]byte, []int) {
	return fileDescriptor_33bdf83d440e1b7e, []int{4}
}

func (m *Nothing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nothing.Unmarshal(m, b)
}
func (m *Nothing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nothing.Marshal(b, m, deterministic)
}
func (m *Nothing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nothing.Merge(m, src)
}
func (m *Nothing) XXX_Size() int {
	return xxx_messageInfo_Nothing.Size(m)
}
func (m *Nothing) XXX_DiscardUnknown() {
	xxx_messageInfo_Nothing.DiscardUnknown(m)
}

var xxx_messageInfo_Nothing proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SessionId)(nil), "main.SessionId")
	proto.RegisterType((*UserID)(nil), "main.UserID")
	proto.RegisterType((*Session)(nil), "main.Session")
	proto.RegisterType((*SessionInfo)(nil), "main.SessionInfo")
	proto.RegisterType((*Nothing)(nil), "main.Nothing")
}

func init() { proto.RegisterFile("session-grpc.proto", fileDescriptor_33bdf83d440e1b7e) }

var fileDescriptor_33bdf83d440e1b7e = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x4e, 0x2d, 0x2e,
	0xce, 0xcc, 0xcf, 0xd3, 0x4d, 0x2f, 0x2a, 0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0xd2, 0xe4, 0xe2, 0x0c, 0x86, 0xc8, 0x79, 0xa6, 0x08, 0xc9, 0x70,
	0x71, 0x42, 0x15, 0x7a, 0xba, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x21, 0x04, 0x94, 0x14,
	0xb8, 0xd8, 0x42, 0x8b, 0x53, 0x8b, 0x3c, 0x5d, 0x84, 0xc4, 0xb8, 0xd8, 0x4a, 0xc1, 0x2c, 0xa8,
	0x22, 0x28, 0x4f, 0xc9, 0x9f, 0x8b, 0x1d, 0x6a, 0x98, 0x10, 0x1f, 0x17, 0x13, 0x5c, 0x9a, 0x09,
	0xa2, 0x05, 0xa2, 0x59, 0x82, 0x09, 0xa2, 0x05, 0x6a, 0x94, 0x0c, 0x17, 0xa7, 0x73, 0x71, 0x51,
	0x5a, 0x48, 0x7e, 0x76, 0x6a, 0x9e, 0x04, 0x33, 0xc4, 0x4a, 0xb8, 0x80, 0x92, 0x29, 0x17, 0x37,
	0xcc, 0x75, 0x79, 0x69, 0xf9, 0xc4, 0x1a, 0xaa, 0xc4, 0xc9, 0xc5, 0xee, 0x97, 0x5f, 0x92, 0x91,
	0x99, 0x97, 0x6e, 0xd4, 0xc3, 0xc8, 0xc5, 0x11, 0x9c, 0x9a, 0x5c, 0x5a, 0x94, 0x59, 0x52, 0x29,
	0xa4, 0xce, 0xc5, 0xe6, 0x5c, 0x94, 0x9a, 0x58, 0x92, 0x2a, 0xc4, 0xa3, 0x07, 0xf2, 0xbd, 0x1e,
	0x44, 0xbd, 0x14, 0x2f, 0x84, 0x07, 0xb5, 0x4a, 0x89, 0x41, 0x48, 0x9b, 0x8b, 0xd5, 0x39, 0x23,
	0x35, 0x39, 0x5b, 0x48, 0x10, 0x45, 0x06, 0xe4, 0x08, 0x4c, 0xc5, 0x5a, 0x5c, 0x6c, 0x2e, 0xa9,
	0x39, 0xa9, 0x25, 0xa9, 0x42, 0xfc, 0xa8, 0xaa, 0x53, 0x60, 0x6a, 0xa1, 0x8e, 0x51, 0x62, 0x48,
	0x62, 0x03, 0x87, 0xbd, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x08, 0xd5, 0x20, 0x26, 0x91, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SecurityClient is the client API for Security service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecurityClient interface {
	Create(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Session, error)
	Check(ctx context.Context, in *SessionInfo, opts ...grpc.CallOption) (*Session, error)
	Delete(ctx context.Context, in *SessionId, opts ...grpc.CallOption) (*Nothing, error)
}

type securityClient struct {
	cc *grpc.ClientConn
}

func NewSecurityClient(cc *grpc.ClientConn) SecurityClient {
	return &securityClient{cc}
}

func (c *securityClient) Create(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/main.Security/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityClient) Check(ctx context.Context, in *SessionInfo, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/main.Security/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityClient) Delete(ctx context.Context, in *SessionId, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/main.Security/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecurityServer is the server API for Security service.
type SecurityServer interface {
	Create(context.Context, *UserID) (*Session, error)
	Check(context.Context, *SessionInfo) (*Session, error)
	Delete(context.Context, *SessionId) (*Nothing, error)
}

// UnimplementedSecurityServer can be embedded to have forward compatible implementations.
type UnimplementedSecurityServer struct {
}

func (*UnimplementedSecurityServer) Create(ctx context.Context, req *UserID) (*Session, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedSecurityServer) Check(ctx context.Context, req *SessionInfo) (*Session, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (*UnimplementedSecurityServer) Delete(ctx context.Context, req *SessionId) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterSecurityServer(s *grpc.Server, srv SecurityServer) {
	s.RegisterService(&_Security_serviceDesc, srv)
}

func _Security_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Security/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServer).Create(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Security_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Security/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServer).Check(ctx, req.(*SessionInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Security_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Security/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServer).Delete(ctx, req.(*SessionId))
	}
	return interceptor(ctx, in, info, handler)
}

var _Security_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.Security",
	HandlerType: (*SecurityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Security_Create_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _Security_Check_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Security_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session-grpc.proto",
}
