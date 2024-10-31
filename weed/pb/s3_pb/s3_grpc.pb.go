// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: s3.proto

package s3_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SeaweedS3_Configure_FullMethodName = "/messaging_pb.SeaweedS3/Configure"
)

// SeaweedS3Client is the client API for SeaweedS3 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SeaweedS3Client interface {
	Configure(ctx context.Context, in *S3ConfigureRequest, opts ...grpc.CallOption) (*S3ConfigureResponse, error)
}

type seaweedS3Client struct {
	cc grpc.ClientConnInterface
}

func NewSeaweedS3Client(cc grpc.ClientConnInterface) SeaweedS3Client {
	return &seaweedS3Client{cc}
}

func (c *seaweedS3Client) Configure(ctx context.Context, in *S3ConfigureRequest, opts ...grpc.CallOption) (*S3ConfigureResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(S3ConfigureResponse)
	err := c.cc.Invoke(ctx, SeaweedS3_Configure_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SeaweedS3Server is the server API for SeaweedS3 service.
// All implementations must embed UnimplementedSeaweedS3Server
// for forward compatibility.
type SeaweedS3Server interface {
	Configure(context.Context, *S3ConfigureRequest) (*S3ConfigureResponse, error)
	mustEmbedUnimplementedSeaweedS3Server()
}

// UnimplementedSeaweedS3Server must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSeaweedS3Server struct{}

func (UnimplementedSeaweedS3Server) Configure(context.Context, *S3ConfigureRequest) (*S3ConfigureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (UnimplementedSeaweedS3Server) mustEmbedUnimplementedSeaweedS3Server() {}
func (UnimplementedSeaweedS3Server) testEmbeddedByValue()                   {}

// UnsafeSeaweedS3Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SeaweedS3Server will
// result in compilation errors.
type UnsafeSeaweedS3Server interface {
	mustEmbedUnimplementedSeaweedS3Server()
}

func RegisterSeaweedS3Server(s grpc.ServiceRegistrar, srv SeaweedS3Server) {
	// If the following call pancis, it indicates UnimplementedSeaweedS3Server was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SeaweedS3_ServiceDesc, srv)
}

func _SeaweedS3_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(S3ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedS3Server).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SeaweedS3_Configure_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedS3Server).Configure(ctx, req.(*S3ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SeaweedS3_ServiceDesc is the grpc.ServiceDesc for SeaweedS3 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SeaweedS3_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messaging_pb.SeaweedS3",
	HandlerType: (*SeaweedS3Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _SeaweedS3_Configure_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "s3.proto",
}
