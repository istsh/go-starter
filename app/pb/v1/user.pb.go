// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateUserRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateUserResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateUserRequest)(nil), "v1.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "v1.CreateUserResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4b, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0xc9, 0xe8, 0x8c, 0x99, 0x16, 0x61, 0x6c, 0x5f, 0x21, 0xb8, 0x18, 0x02, 0x82, 0x0f,
	0x26, 0x4d, 0x74, 0xe7, 0x4a, 0xe2, 0x05, 0x24, 0xe2, 0x46, 0x04, 0xe9, 0x64, 0x8a, 0x4c, 0x43,
	0x92, 0x6a, 0xbb, 0x3a, 0xed, 0xde, 0x2b, 0x78, 0x34, 0xaf, 0xe0, 0x29, 0x66, 0x21, 0x32, 0x09,
	0x3e, 0xd0, 0x5d, 0x55, 0xfd, 0x1f, 0x3f, 0xc5, 0xc7, 0x58, 0x4b, 0x60, 0x62, 0x6d, 0xd0, 0x22,
	0x1f, 0xb8, 0x24, 0x3c, 0x2c, 0x11, 0xcb, 0x0a, 0x84, 0xd4, 0x4a, 0xc8, 0xa6, 0x41, 0x2b, 0xad,
	0xc2, 0x86, 0x7a, 0x22, 0x3c, 0x70, 0xb2, 0x52, 0x73, 0x69, 0x41, 0x7c, 0x0d, 0x7d, 0x10, 0x3d,
	0xb0, 0xed, 0x6b, 0x03, 0xd2, 0xc2, 0x1d, 0x81, 0xc9, 0xe0, 0xa9, 0x05, 0xb2, 0x7c, 0xca, 0x86,
	0x50, 0x4b, 0x55, 0x05, 0xde, 0xd4, 0x3b, 0x1e, 0xa7, 0x6c, 0x99, 0x6e, 0x98, 0xe1, 0x64, 0x2d,
	0xf8, 0xf0, 0xb2, 0x3e, 0xe0, 0x47, 0xcc, 0xd7, 0x92, 0xe8, 0x19, 0xcd, 0x3c, 0x18, 0x74, 0xd0,
	0x78, 0x99, 0x8e, 0xcc, 0xfa, 0xc4, 0x0f, 0xae, 0xb2, 0xef, 0x28, 0xda, 0x65, 0xfc, 0x77, 0x3b,
	0x69, 0x6c, 0x08, 0xce, 0x1f, 0xd9, 0xe6, 0x6a, 0xbf, 0x05, 0xe3, 0x54, 0x01, 0xfc, 0x86, 0xb1,
	0x1f, 0x88, 0xef, 0xc5, 0x2e, 0x89, 0xff, 0xbd, 0x14, 0xee, 0xff, 0x3d, 0xf7, 0x5d, 0xd1, 0xce,
	0xcb, 0xdb, 0xfb, 0xeb, 0x60, 0x2b, 0xf2, 0x85, 0x4b, 0xc4, 0x4a, 0xc9, 0xa5, 0x77, 0x9a, 0x9e,
	0xdd, 0x9f, 0x94, 0xca, 0x2e, 0xda, 0x3c, 0x2e, 0xb0, 0x16, 0x8a, 0x2c, 0x2d, 0x44, 0x89, 0xb3,
	0xd2, 0xe8, 0x62, 0x46, 0xb2, 0xd6, 0x9d, 0x26, 0x2d, 0x74, 0x2e, 0x5c, 0x92, 0x8f, 0x3a, 0x11,
	0x17, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x9f, 0x3c, 0xf4, 0x51, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
