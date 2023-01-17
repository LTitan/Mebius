// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/protos/server.proto

package protos

import (
	context "context"
	fmt "fmt"
	v1alpha1 "github.com/LTitan/Mebius/pkg/apis/v1alpha1"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("pkg/protos/server.proto", fileDescriptor_bfa55bbbe0c38e8f) }
func init() { golang_proto.RegisterFile("pkg/protos/server.proto", fileDescriptor_bfa55bbbe0c38e8f) }

var fileDescriptor_bfa55bbbe0c38e8f = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0xc8, 0x4e, 0xd7,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x2f, 0xd6, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x03, 0xf3,
	0x84, 0x58, 0xc1, 0x94, 0x94, 0x6e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae,
	0x7e, 0x7a, 0x7e, 0x7a, 0x3e, 0x44, 0x6d, 0x52, 0x69, 0x1a, 0x98, 0x07, 0xe6, 0x80, 0x59, 0x10,
	0x5d, 0x52, 0x56, 0x48, 0xca, 0x7d, 0x42, 0x32, 0x4b, 0x12, 0xf3, 0xf4, 0x7d, 0x53, 0x93, 0x32,
	0x4b, 0x8b, 0xf5, 0x41, 0xf6, 0x24, 0x16, 0x64, 0x16, 0xeb, 0x97, 0x19, 0x26, 0xe6, 0x14, 0x64,
	0x24, 0x1a, 0xea, 0xa7, 0xa7, 0xe6, 0xa5, 0x16, 0x25, 0x96, 0xa4, 0xa6, 0x40, 0xf5, 0x9a, 0xa3,
	0x5b, 0x95, 0x9e, 0x9f, 0x9f, 0x9e, 0x93, 0x0a, 0xd6, 0x04, 0x61, 0x82, 0x0c, 0xd0, 0x4f, 0xcc,
	0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x68, 0x34, 0x5a, 0xc0, 0xc8, 0xc5,
	0x16, 0x0c, 0x76, 0xbb, 0xd0, 0x54, 0x46, 0x2e, 0x2e, 0xf7, 0xd4, 0x12, 0xdf, 0xc4, 0xe4, 0x8c,
	0xcc, 0xbc, 0x54, 0x21, 0x63, 0x3d, 0x84, 0x99, 0x7a, 0x10, 0xf7, 0xe8, 0x41, 0xdc, 0xa3, 0x57,
	0x90, 0x9d, 0xae, 0x07, 0x32, 0x5a, 0x0f, 0xe6, 0x1e, 0x3d, 0xa8, 0x26, 0x29, 0x72, 0x34, 0x29,
	0xc9, 0x36, 0x5d, 0x7e, 0x32, 0x99, 0x49, 0x5c, 0x48, 0x14, 0xec, 0x46, 0xb8, 0x1f, 0x73, 0x21,
	0xd2, 0x4e, 0x0e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3,
	0x81, 0xc7, 0x72, 0x8c, 0x27, 0x1e, 0xcb, 0x31, 0x46, 0x69, 0xe1, 0x0d, 0x2d, 0x48, 0xac, 0x58,
	0x43, 0xa8, 0x24, 0x36, 0x30, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x7a, 0x82, 0xb2,
	0xb1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerClient interface {
	GetMachine(ctx context.Context, in *v1alpha1.Machine, opts ...grpc.CallOption) (*v1alpha1.Machine, error)
}

type serverClient struct {
	cc *grpc.ClientConn
}

func NewServerClient(cc *grpc.ClientConn) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) GetMachine(ctx context.Context, in *v1alpha1.Machine, opts ...grpc.CallOption) (*v1alpha1.Machine, error) {
	out := new(v1alpha1.Machine)
	err := c.cc.Invoke(ctx, "/proto.Server/GetMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerServer is the server API for Server service.
type ServerServer interface {
	GetMachine(context.Context, *v1alpha1.Machine) (*v1alpha1.Machine, error)
}

// UnimplementedServerServer can be embedded to have forward compatible implementations.
type UnimplementedServerServer struct {
}

func (*UnimplementedServerServer) GetMachine(ctx context.Context, req *v1alpha1.Machine) (*v1alpha1.Machine, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMachine not implemented")
}

func RegisterServerServer(s *grpc.Server, srv ServerServer) {
	s.RegisterService(&_Server_serviceDesc, srv)
}

func _Server_GetMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.Machine)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).GetMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Server/GetMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).GetMachine(ctx, req.(*v1alpha1.Machine))
	}
	return interceptor(ctx, in, info, handler)
}

var _Server_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMachine",
			Handler:    _Server_GetMachine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/protos/server.proto",
}
