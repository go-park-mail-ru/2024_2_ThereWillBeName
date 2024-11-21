// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--rc2
// source: proto/attractions.proto

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
	Attractions_GetPlaces_FullMethodName           = "/attractions.Attractions/GetPlaces"
	Attractions_GetPlace_FullMethodName            = "/attractions.Attractions/GetPlace"
	Attractions_SearchPlaces_FullMethodName        = "/attractions.Attractions/SearchPlaces"
	Attractions_GetPlacesByCategory_FullMethodName = "/attractions.Attractions/GetPlacesByCategory"
	Attractions_GetCategories_FullMethodName       = "/attractions.Attractions/GetCategories"
	Attractions_SearchCitiesByName_FullMethodName  = "/attractions.Attractions/SearchCitiesByName"
	Attractions_SearchCityByID_FullMethodName      = "/attractions.Attractions/SearchCityByID"
	Attractions_CreateReview_FullMethodName        = "/attractions.Attractions/CreateReview"
	Attractions_UpdateReview_FullMethodName        = "/attractions.Attractions/UpdateReview"
	Attractions_DeleteReview_FullMethodName        = "/attractions.Attractions/DeleteReview"
	Attractions_GetReviewsByPlaceID_FullMethodName = "/attractions.Attractions/GetReviewsByPlaceID"
	Attractions_GetReviewsByUserID_FullMethodName  = "/attractions.Attractions/GetReviewsByUserID"
	Attractions_GetReview_FullMethodName           = "/attractions.Attractions/GetReview"
)

// AttractionsClient is the client API for Attractions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttractionsClient interface {
	GetPlaces(ctx context.Context, in *GetPlacesRequest, opts ...grpc.CallOption) (*GetPlacesResponse, error)
	// rpc CreatePlace(CreatePlaceRequest) returns (CreatePlaceResponse) {}
	GetPlace(ctx context.Context, in *GetPlaceRequest, opts ...grpc.CallOption) (*GetPlaceResponse, error)
	// rpc UpdatePlace(UpdatePlaceRequest) returns (UpdatePlaceResponse) {}
	// rpc DeletePlace(DeletePlaceRequest) returns (DeletePlaceResponse) {}
	SearchPlaces(ctx context.Context, in *SearchPlacesRequest, opts ...grpc.CallOption) (*SearchPlacesResponse, error)
	GetPlacesByCategory(ctx context.Context, in *GetPlacesByCategoryRequest, opts ...grpc.CallOption) (*GetPlacesByCategoryResponse, error)
	GetCategories(ctx context.Context, in *GetCategoriesRequest, opts ...grpc.CallOption) (*GetCategoriesResponse, error)
	SearchCitiesByName(ctx context.Context, in *SearchCitiesByNameRequest, opts ...grpc.CallOption) (*SearchCitiesByNameResponse, error)
	SearchCityByID(ctx context.Context, in *SearchCityByIDRequest, opts ...grpc.CallOption) (*SearchCityByIDResponse, error)
	CreateReview(ctx context.Context, in *CreateReviewRequest, opts ...grpc.CallOption) (*CreateReviewResponse, error)
	UpdateReview(ctx context.Context, in *UpdateReviewRequest, opts ...grpc.CallOption) (*UpdateReviewResponse, error)
	DeleteReview(ctx context.Context, in *DeleteReviewRequest, opts ...grpc.CallOption) (*DeleteReviewResponse, error)
	GetReviewsByPlaceID(ctx context.Context, in *GetReviewsByPlaceIDRequest, opts ...grpc.CallOption) (*GetReviewsByPlaceIDResponse, error)
	GetReviewsByUserID(ctx context.Context, in *GetReviewsByUserIDRequest, opts ...grpc.CallOption) (*GetReviewsByUserIDResponse, error)
	GetReview(ctx context.Context, in *GetReviewRequest, opts ...grpc.CallOption) (*GetReviewResponse, error)
}

type attractionsClient struct {
	cc grpc.ClientConnInterface
}

func NewAttractionsClient(cc grpc.ClientConnInterface) AttractionsClient {
	return &attractionsClient{cc}
}

func (c *attractionsClient) GetPlaces(ctx context.Context, in *GetPlacesRequest, opts ...grpc.CallOption) (*GetPlacesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPlacesResponse)
	err := c.cc.Invoke(ctx, Attractions_GetPlaces_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attractionsClient) GetPlace(ctx context.Context, in *GetPlaceRequest, opts ...grpc.CallOption) (*GetPlaceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPlaceResponse)
	err := c.cc.Invoke(ctx, Attractions_GetPlace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attractionsClient) SearchPlaces(ctx context.Context, in *SearchPlacesRequest, opts ...grpc.CallOption) (*SearchPlacesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchPlacesResponse)
	err := c.cc.Invoke(ctx, Attractions_SearchPlaces_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attractionsClient) GetPlacesByCategory(ctx context.Context, in *GetPlacesByCategoryRequest, opts ...grpc.CallOption) (*GetPlacesByCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPlacesByCategoryResponse)
	err := c.cc.Invoke(ctx, Attractions_GetPlacesByCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attractionsClient) GetCategories(ctx context.Context, in *GetCategoriesRequest, opts ...grpc.CallOption) (*GetCategoriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCategoriesResponse)
	err := c.cc.Invoke(ctx, Attractions_GetCategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attractionsClient) SearchCitiesByName(ctx context.Context, in *SearchCitiesByNameRequest, opts ...grpc.CallOption) (*SearchCitiesByNameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchCitiesByNameResponse)
	err := c.cc.Invoke(ctx, Attractions_SearchCitiesByName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attractionsClient) SearchCityByID(ctx context.Context, in *SearchCityByIDRequest, opts ...grpc.CallOption) (*SearchCityByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchCityByIDResponse)
	err := c.cc.Invoke(ctx, Attractions_SearchCityByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attractionsClient) CreateReview(ctx context.Context, in *CreateReviewRequest, opts ...grpc.CallOption) (*CreateReviewResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateReviewResponse)
	err := c.cc.Invoke(ctx, Attractions_CreateReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attractionsClient) UpdateReview(ctx context.Context, in *UpdateReviewRequest, opts ...grpc.CallOption) (*UpdateReviewResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateReviewResponse)
	err := c.cc.Invoke(ctx, Attractions_UpdateReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attractionsClient) DeleteReview(ctx context.Context, in *DeleteReviewRequest, opts ...grpc.CallOption) (*DeleteReviewResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteReviewResponse)
	err := c.cc.Invoke(ctx, Attractions_DeleteReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attractionsClient) GetReviewsByPlaceID(ctx context.Context, in *GetReviewsByPlaceIDRequest, opts ...grpc.CallOption) (*GetReviewsByPlaceIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReviewsByPlaceIDResponse)
	err := c.cc.Invoke(ctx, Attractions_GetReviewsByPlaceID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attractionsClient) GetReviewsByUserID(ctx context.Context, in *GetReviewsByUserIDRequest, opts ...grpc.CallOption) (*GetReviewsByUserIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReviewsByUserIDResponse)
	err := c.cc.Invoke(ctx, Attractions_GetReviewsByUserID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attractionsClient) GetReview(ctx context.Context, in *GetReviewRequest, opts ...grpc.CallOption) (*GetReviewResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReviewResponse)
	err := c.cc.Invoke(ctx, Attractions_GetReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttractionsServer is the server API for Attractions service.
// All implementations must embed UnimplementedAttractionsServer
// for forward compatibility.
type AttractionsServer interface {
	GetPlaces(context.Context, *GetPlacesRequest) (*GetPlacesResponse, error)
	// rpc CreatePlace(CreatePlaceRequest) returns (CreatePlaceResponse) {}
	GetPlace(context.Context, *GetPlaceRequest) (*GetPlaceResponse, error)
	// rpc UpdatePlace(UpdatePlaceRequest) returns (UpdatePlaceResponse) {}
	// rpc DeletePlace(DeletePlaceRequest) returns (DeletePlaceResponse) {}
	SearchPlaces(context.Context, *SearchPlacesRequest) (*SearchPlacesResponse, error)
	GetPlacesByCategory(context.Context, *GetPlacesByCategoryRequest) (*GetPlacesByCategoryResponse, error)
	GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error)
	SearchCitiesByName(context.Context, *SearchCitiesByNameRequest) (*SearchCitiesByNameResponse, error)
	SearchCityByID(context.Context, *SearchCityByIDRequest) (*SearchCityByIDResponse, error)
	CreateReview(context.Context, *CreateReviewRequest) (*CreateReviewResponse, error)
	UpdateReview(context.Context, *UpdateReviewRequest) (*UpdateReviewResponse, error)
	DeleteReview(context.Context, *DeleteReviewRequest) (*DeleteReviewResponse, error)
	GetReviewsByPlaceID(context.Context, *GetReviewsByPlaceIDRequest) (*GetReviewsByPlaceIDResponse, error)
	GetReviewsByUserID(context.Context, *GetReviewsByUserIDRequest) (*GetReviewsByUserIDResponse, error)
	GetReview(context.Context, *GetReviewRequest) (*GetReviewResponse, error)
	mustEmbedUnimplementedAttractionsServer()
}

// UnimplementedAttractionsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAttractionsServer struct{}

func (UnimplementedAttractionsServer) GetPlaces(context.Context, *GetPlacesRequest) (*GetPlacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaces not implemented")
}
func (UnimplementedAttractionsServer) GetPlace(context.Context, *GetPlaceRequest) (*GetPlaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlace not implemented")
}
func (UnimplementedAttractionsServer) SearchPlaces(context.Context, *SearchPlacesRequest) (*SearchPlacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPlaces not implemented")
}
func (UnimplementedAttractionsServer) GetPlacesByCategory(context.Context, *GetPlacesByCategoryRequest) (*GetPlacesByCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlacesByCategory not implemented")
}
func (UnimplementedAttractionsServer) GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategories not implemented")
}
func (UnimplementedAttractionsServer) SearchCitiesByName(context.Context, *SearchCitiesByNameRequest) (*SearchCitiesByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCitiesByName not implemented")
}
func (UnimplementedAttractionsServer) SearchCityByID(context.Context, *SearchCityByIDRequest) (*SearchCityByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCityByID not implemented")
}
func (UnimplementedAttractionsServer) CreateReview(context.Context, *CreateReviewRequest) (*CreateReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReview not implemented")
}
func (UnimplementedAttractionsServer) UpdateReview(context.Context, *UpdateReviewRequest) (*UpdateReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReview not implemented")
}
func (UnimplementedAttractionsServer) DeleteReview(context.Context, *DeleteReviewRequest) (*DeleteReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReview not implemented")
}
func (UnimplementedAttractionsServer) GetReviewsByPlaceID(context.Context, *GetReviewsByPlaceIDRequest) (*GetReviewsByPlaceIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReviewsByPlaceID not implemented")
}
func (UnimplementedAttractionsServer) GetReviewsByUserID(context.Context, *GetReviewsByUserIDRequest) (*GetReviewsByUserIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReviewsByUserID not implemented")
}
func (UnimplementedAttractionsServer) GetReview(context.Context, *GetReviewRequest) (*GetReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReview not implemented")
}
func (UnimplementedAttractionsServer) mustEmbedUnimplementedAttractionsServer() {}
func (UnimplementedAttractionsServer) testEmbeddedByValue()                     {}

// UnsafeAttractionsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttractionsServer will
// result in compilation errors.
type UnsafeAttractionsServer interface {
	mustEmbedUnimplementedAttractionsServer()
}

func RegisterAttractionsServer(s grpc.ServiceRegistrar, srv AttractionsServer) {
	// If the following call pancis, it indicates UnimplementedAttractionsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Attractions_ServiceDesc, srv)
}

func _Attractions_GetPlaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttractionsServer).GetPlaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Attractions_GetPlaces_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttractionsServer).GetPlaces(ctx, req.(*GetPlacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attractions_GetPlace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttractionsServer).GetPlace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Attractions_GetPlace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttractionsServer).GetPlace(ctx, req.(*GetPlaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attractions_SearchPlaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPlacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttractionsServer).SearchPlaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Attractions_SearchPlaces_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttractionsServer).SearchPlaces(ctx, req.(*SearchPlacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attractions_GetPlacesByCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlacesByCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttractionsServer).GetPlacesByCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Attractions_GetPlacesByCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttractionsServer).GetPlacesByCategory(ctx, req.(*GetPlacesByCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attractions_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttractionsServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Attractions_GetCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttractionsServer).GetCategories(ctx, req.(*GetCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attractions_SearchCitiesByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCitiesByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttractionsServer).SearchCitiesByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Attractions_SearchCitiesByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttractionsServer).SearchCitiesByName(ctx, req.(*SearchCitiesByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attractions_SearchCityByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCityByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttractionsServer).SearchCityByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Attractions_SearchCityByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttractionsServer).SearchCityByID(ctx, req.(*SearchCityByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attractions_CreateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttractionsServer).CreateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Attractions_CreateReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttractionsServer).CreateReview(ctx, req.(*CreateReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attractions_UpdateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttractionsServer).UpdateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Attractions_UpdateReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttractionsServer).UpdateReview(ctx, req.(*UpdateReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attractions_DeleteReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttractionsServer).DeleteReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Attractions_DeleteReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttractionsServer).DeleteReview(ctx, req.(*DeleteReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attractions_GetReviewsByPlaceID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReviewsByPlaceIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttractionsServer).GetReviewsByPlaceID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Attractions_GetReviewsByPlaceID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttractionsServer).GetReviewsByPlaceID(ctx, req.(*GetReviewsByPlaceIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attractions_GetReviewsByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReviewsByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttractionsServer).GetReviewsByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Attractions_GetReviewsByUserID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttractionsServer).GetReviewsByUserID(ctx, req.(*GetReviewsByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attractions_GetReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttractionsServer).GetReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Attractions_GetReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttractionsServer).GetReview(ctx, req.(*GetReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Attractions_ServiceDesc is the grpc.ServiceDesc for Attractions service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Attractions_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "attractions.Attractions",
	HandlerType: (*AttractionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlaces",
			Handler:    _Attractions_GetPlaces_Handler,
		},
		{
			MethodName: "GetPlace",
			Handler:    _Attractions_GetPlace_Handler,
		},
		{
			MethodName: "SearchPlaces",
			Handler:    _Attractions_SearchPlaces_Handler,
		},
		{
			MethodName: "GetPlacesByCategory",
			Handler:    _Attractions_GetPlacesByCategory_Handler,
		},
		{
			MethodName: "GetCategories",
			Handler:    _Attractions_GetCategories_Handler,
		},
		{
			MethodName: "SearchCitiesByName",
			Handler:    _Attractions_SearchCitiesByName_Handler,
		},
		{
			MethodName: "SearchCityByID",
			Handler:    _Attractions_SearchCityByID_Handler,
		},
		{
			MethodName: "CreateReview",
			Handler:    _Attractions_CreateReview_Handler,
		},
		{
			MethodName: "UpdateReview",
			Handler:    _Attractions_UpdateReview_Handler,
		},
		{
			MethodName: "DeleteReview",
			Handler:    _Attractions_DeleteReview_Handler,
		},
		{
			MethodName: "GetReviewsByPlaceID",
			Handler:    _Attractions_GetReviewsByPlaceID_Handler,
		},
		{
			MethodName: "GetReviewsByUserID",
			Handler:    _Attractions_GetReviewsByUserID_Handler,
		},
		{
			MethodName: "GetReview",
			Handler:    _Attractions_GetReview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/attractions.proto",
}
