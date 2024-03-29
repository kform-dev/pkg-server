// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: pkg/pkginventory/pkg/pkg.proto

package pkg

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PackageResourcesClient is the client API for PackageResources service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PackageResourcesClient interface {
	Get(ctx context.Context, in *Get_Request, opts ...grpc.CallOption) (*Get_Response, error)
}

type packageResourcesClient struct {
	cc grpc.ClientConnInterface
}

func NewPackageResourcesClient(cc grpc.ClientConnInterface) PackageResourcesClient {
	return &packageResourcesClient{cc}
}

func (c *packageResourcesClient) Get(ctx context.Context, in *Get_Request, opts ...grpc.CallOption) (*Get_Response, error) {
	out := new(Get_Response)
	err := c.cc.Invoke(ctx, "/pkg.PackageResources/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PackageResourcesServer is the server API for PackageResources service.
// All implementations must embed UnimplementedPackageResourcesServer
// for forward compatibility
type PackageResourcesServer interface {
	Get(context.Context, *Get_Request) (*Get_Response, error)
	mustEmbedUnimplementedPackageResourcesServer()
}

// UnimplementedPackageResourcesServer must be embedded to have forward compatible implementations.
type UnimplementedPackageResourcesServer struct {
}

func (UnimplementedPackageResourcesServer) Get(context.Context, *Get_Request) (*Get_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPackageResourcesServer) mustEmbedUnimplementedPackageResourcesServer() {}

// UnsafePackageResourcesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PackageResourcesServer will
// result in compilation errors.
type UnsafePackageResourcesServer interface {
	mustEmbedUnimplementedPackageResourcesServer()
}

func RegisterPackageResourcesServer(s grpc.ServiceRegistrar, srv PackageResourcesServer) {
	s.RegisterService(&PackageResources_ServiceDesc, srv)
}

func _PackageResources_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Get_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageResourcesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.PackageResources/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageResourcesServer).Get(ctx, req.(*Get_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// PackageResources_ServiceDesc is the grpc.ServiceDesc for PackageResources service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PackageResources_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pkg.PackageResources",
	HandlerType: (*PackageResourcesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PackageResources_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pkginventory/pkg/pkg.proto",
}
