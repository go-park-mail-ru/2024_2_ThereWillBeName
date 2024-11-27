// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--rc2
// source: proto/categories.proto

package gen

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
	Categories_GetCategories_FullMethodName = "/categories.Categories/GetCategories"
)

// CategoriesClient is the client API for Categories service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoriesClient interface {
	GetCategories(ctx context.Context, in *GetCategoriesRequest, opts ...grpc.CallOption) (*GetCategoriesResponse, error)
}

type categoriesClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoriesClient(cc grpc.ClientConnInterface) CategoriesClient {
	return &categoriesClient{cc}
}

func (c *categoriesClient) GetCategories(ctx context.Context, in *GetCategoriesRequest, opts ...grpc.CallOption) (*GetCategoriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCategoriesResponse)
	err := c.cc.Invoke(ctx, Categories_GetCategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoriesServer is the server API for Categories service.
// All implementations must embed UnimplementedCategoriesServer
// for forward compatibility.
type CategoriesServer interface {
	GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error)
	mustEmbedUnimplementedCategoriesServer()
}

// UnimplementedCategoriesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCategoriesServer struct{}

func (UnimplementedCategoriesServer) GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategories not implemented")
}
func (UnimplementedCategoriesServer) mustEmbedUnimplementedCategoriesServer() {}
func (UnimplementedCategoriesServer) testEmbeddedByValue()                    {}

// UnsafeCategoriesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoriesServer will
// result in compilation errors.
type UnsafeCategoriesServer interface {
	mustEmbedUnimplementedCategoriesServer()
}

func RegisterCategoriesServer(s grpc.ServiceRegistrar, srv CategoriesServer) {
	// If the following call pancis, it indicates UnimplementedCategoriesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Categories_ServiceDesc, srv)
}

func _Categories_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoriesServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Categories_GetCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoriesServer).GetCategories(ctx, req.(*GetCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Categories_ServiceDesc is the grpc.ServiceDesc for Categories service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Categories_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "categories.Categories",
	HandlerType: (*CategoriesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCategories",
			Handler:    _Categories_GetCategories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/categories.proto",
}
