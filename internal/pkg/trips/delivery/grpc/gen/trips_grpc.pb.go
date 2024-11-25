// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
// source: proto/trips.proto

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
	Trips_CreateTrip_FullMethodName       = "/trips.Trips/CreateTrip"
	Trips_UpdateTrip_FullMethodName       = "/trips.Trips/UpdateTrip"
	Trips_DeleteTrip_FullMethodName       = "/trips.Trips/DeleteTrip"
	Trips_GetTripsByUserID_FullMethodName = "/trips.Trips/GetTripsByUserID"
	Trips_GetTrip_FullMethodName          = "/trips.Trips/GetTrip"
	Trips_AddPlaceToTrip_FullMethodName   = "/trips.Trips/AddPlaceToTrip"
	Trips_AddPhotosToTrip_FullMethodName  = "/trips.Trips/AddPhotosToTrip"
)

// TripsClient is the client API for Trips service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TripsClient interface {
	CreateTrip(ctx context.Context, in *CreateTripRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	UpdateTrip(ctx context.Context, in *UpdateTripRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	DeleteTrip(ctx context.Context, in *DeleteTripRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetTripsByUserID(ctx context.Context, in *GetTripsByUserIDRequest, opts ...grpc.CallOption) (*GetTripsByUserIDResponse, error)
	GetTrip(ctx context.Context, in *GetTripRequest, opts ...grpc.CallOption) (*GetTripResponse, error)
	AddPlaceToTrip(ctx context.Context, in *AddPlaceToTripRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	AddPhotosToTrip(ctx context.Context, in *AddPhotosToTripRequest, opts ...grpc.CallOption) (*AddPhotosToTripResponse, error)
}

type tripsClient struct {
	cc grpc.ClientConnInterface
}

func NewTripsClient(cc grpc.ClientConnInterface) TripsClient {
	return &tripsClient{cc}
}

func (c *tripsClient) CreateTrip(ctx context.Context, in *CreateTripRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, Trips_CreateTrip_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripsClient) UpdateTrip(ctx context.Context, in *UpdateTripRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, Trips_UpdateTrip_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripsClient) DeleteTrip(ctx context.Context, in *DeleteTripRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, Trips_DeleteTrip_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripsClient) GetTripsByUserID(ctx context.Context, in *GetTripsByUserIDRequest, opts ...grpc.CallOption) (*GetTripsByUserIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTripsByUserIDResponse)
	err := c.cc.Invoke(ctx, Trips_GetTripsByUserID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripsClient) GetTrip(ctx context.Context, in *GetTripRequest, opts ...grpc.CallOption) (*GetTripResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTripResponse)
	err := c.cc.Invoke(ctx, Trips_GetTrip_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripsClient) AddPlaceToTrip(ctx context.Context, in *AddPlaceToTripRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, Trips_AddPlaceToTrip_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripsClient) AddPhotosToTrip(ctx context.Context, in *AddPhotosToTripRequest, opts ...grpc.CallOption) (*AddPhotosToTripResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddPhotosToTripResponse)
	err := c.cc.Invoke(ctx, Trips_AddPhotosToTrip_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TripsServer is the server API for Trips service.
// All implementations must embed UnimplementedTripsServer
// for forward compatibility.
type TripsServer interface {
	CreateTrip(context.Context, *CreateTripRequest) (*EmptyResponse, error)
	UpdateTrip(context.Context, *UpdateTripRequest) (*EmptyResponse, error)
	DeleteTrip(context.Context, *DeleteTripRequest) (*EmptyResponse, error)
	GetTripsByUserID(context.Context, *GetTripsByUserIDRequest) (*GetTripsByUserIDResponse, error)
	GetTrip(context.Context, *GetTripRequest) (*GetTripResponse, error)
	AddPlaceToTrip(context.Context, *AddPlaceToTripRequest) (*EmptyResponse, error)
	AddPhotosToTrip(context.Context, *AddPhotosToTripRequest) (*AddPhotosToTripResponse, error)
	mustEmbedUnimplementedTripsServer()
}

// UnimplementedTripsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTripsServer struct{}

func (UnimplementedTripsServer) CreateTrip(context.Context, *CreateTripRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrip not implemented")
}
func (UnimplementedTripsServer) UpdateTrip(context.Context, *UpdateTripRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTrip not implemented")
}
func (UnimplementedTripsServer) DeleteTrip(context.Context, *DeleteTripRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTrip not implemented")
}
func (UnimplementedTripsServer) GetTripsByUserID(context.Context, *GetTripsByUserIDRequest) (*GetTripsByUserIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTripsByUserID not implemented")
}
func (UnimplementedTripsServer) GetTrip(context.Context, *GetTripRequest) (*GetTripResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrip not implemented")
}
func (UnimplementedTripsServer) AddPlaceToTrip(context.Context, *AddPlaceToTripRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPlaceToTrip not implemented")
}
func (UnimplementedTripsServer) AddPhotosToTrip(context.Context, *AddPhotosToTripRequest) (*AddPhotosToTripResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPhotosToTrip not implemented")
}
func (UnimplementedTripsServer) mustEmbedUnimplementedTripsServer() {}
func (UnimplementedTripsServer) testEmbeddedByValue()               {}

// UnsafeTripsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TripsServer will
// result in compilation errors.
type UnsafeTripsServer interface {
	mustEmbedUnimplementedTripsServer()
}

func RegisterTripsServer(s grpc.ServiceRegistrar, srv TripsServer) {
	// If the following call pancis, it indicates UnimplementedTripsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Trips_ServiceDesc, srv)
}

func _Trips_CreateTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripsServer).CreateTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trips_CreateTrip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripsServer).CreateTrip(ctx, req.(*CreateTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trips_UpdateTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripsServer).UpdateTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trips_UpdateTrip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripsServer).UpdateTrip(ctx, req.(*UpdateTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trips_DeleteTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripsServer).DeleteTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trips_DeleteTrip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripsServer).DeleteTrip(ctx, req.(*DeleteTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trips_GetTripsByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTripsByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripsServer).GetTripsByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trips_GetTripsByUserID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripsServer).GetTripsByUserID(ctx, req.(*GetTripsByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trips_GetTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripsServer).GetTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trips_GetTrip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripsServer).GetTrip(ctx, req.(*GetTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trips_AddPlaceToTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPlaceToTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripsServer).AddPlaceToTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trips_AddPlaceToTrip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripsServer).AddPlaceToTrip(ctx, req.(*AddPlaceToTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trips_AddPhotosToTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPhotosToTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripsServer).AddPhotosToTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trips_AddPhotosToTrip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripsServer).AddPhotosToTrip(ctx, req.(*AddPhotosToTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Trips_ServiceDesc is the grpc.ServiceDesc for Trips service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Trips_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trips.Trips",
	HandlerType: (*TripsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTrip",
			Handler:    _Trips_CreateTrip_Handler,
		},
		{
			MethodName: "UpdateTrip",
			Handler:    _Trips_UpdateTrip_Handler,
		},
		{
			MethodName: "DeleteTrip",
			Handler:    _Trips_DeleteTrip_Handler,
		},
		{
			MethodName: "GetTripsByUserID",
			Handler:    _Trips_GetTripsByUserID_Handler,
		},
		{
			MethodName: "GetTrip",
			Handler:    _Trips_GetTrip_Handler,
		},
		{
			MethodName: "AddPlaceToTrip",
			Handler:    _Trips_AddPlaceToTrip_Handler,
		},
		{
			MethodName: "AddPhotosToTrip",
			Handler:    _Trips_AddPhotosToTrip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/trips.proto",
}
