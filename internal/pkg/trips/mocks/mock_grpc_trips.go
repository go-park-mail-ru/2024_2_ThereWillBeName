// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/trips/delivery/grpc/gen/trips_grpc.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gen "2024_2_ThereWillBeName/internal/pkg/trips/delivery/grpc/gen"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockTripsClient is a mock of TripsClient interface.
type MockTripsClient struct {
	ctrl     *gomock.Controller
	recorder *MockTripsClientMockRecorder
}

// MockTripsClientMockRecorder is the mock recorder for MockTripsClient.
type MockTripsClientMockRecorder struct {
	mock *MockTripsClient
}

// NewMockTripsClient creates a new mock instance.
func NewMockTripsClient(ctrl *gomock.Controller) *MockTripsClient {
	mock := &MockTripsClient{ctrl: ctrl}
	mock.recorder = &MockTripsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTripsClient) EXPECT() *MockTripsClientMockRecorder {
	return m.recorder
}

// AddPhotosToTrip mocks base method.
func (m *MockTripsClient) AddPhotosToTrip(ctx context.Context, in *gen.AddPhotosToTripRequest, opts ...grpc.CallOption) (*gen.AddPhotosToTripResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddPhotosToTrip", varargs...)
	ret0, _ := ret[0].(*gen.AddPhotosToTripResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPhotosToTrip indicates an expected call of AddPhotosToTrip.
func (mr *MockTripsClientMockRecorder) AddPhotosToTrip(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPhotosToTrip", reflect.TypeOf((*MockTripsClient)(nil).AddPhotosToTrip), varargs...)
}

// AddPlaceToTrip mocks base method.
func (m *MockTripsClient) AddPlaceToTrip(ctx context.Context, in *gen.AddPlaceToTripRequest, opts ...grpc.CallOption) (*gen.EmptyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddPlaceToTrip", varargs...)
	ret0, _ := ret[0].(*gen.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPlaceToTrip indicates an expected call of AddPlaceToTrip.
func (mr *MockTripsClientMockRecorder) AddPlaceToTrip(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPlaceToTrip", reflect.TypeOf((*MockTripsClient)(nil).AddPlaceToTrip), varargs...)
}

// CreateTrip mocks base method.
func (m *MockTripsClient) CreateTrip(ctx context.Context, in *gen.CreateTripRequest, opts ...grpc.CallOption) (*gen.EmptyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTrip", varargs...)
	ret0, _ := ret[0].(*gen.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTrip indicates an expected call of CreateTrip.
func (mr *MockTripsClientMockRecorder) CreateTrip(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrip", reflect.TypeOf((*MockTripsClient)(nil).CreateTrip), varargs...)
}

// DeletePhotoFromTrip mocks base method.
func (m *MockTripsClient) DeletePhotoFromTrip(ctx context.Context, in *gen.DeletePhotoRequest, opts ...grpc.CallOption) (*gen.EmptyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeletePhotoFromTrip", varargs...)
	ret0, _ := ret[0].(*gen.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePhotoFromTrip indicates an expected call of DeletePhotoFromTrip.
func (mr *MockTripsClientMockRecorder) DeletePhotoFromTrip(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePhotoFromTrip", reflect.TypeOf((*MockTripsClient)(nil).DeletePhotoFromTrip), varargs...)
}

// DeleteTrip mocks base method.
func (m *MockTripsClient) DeleteTrip(ctx context.Context, in *gen.DeleteTripRequest, opts ...grpc.CallOption) (*gen.EmptyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteTrip", varargs...)
	ret0, _ := ret[0].(*gen.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTrip indicates an expected call of DeleteTrip.
func (mr *MockTripsClientMockRecorder) DeleteTrip(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrip", reflect.TypeOf((*MockTripsClient)(nil).DeleteTrip), varargs...)
}

// GetTrip mocks base method.
func (m *MockTripsClient) GetTrip(ctx context.Context, in *gen.GetTripRequest, opts ...grpc.CallOption) (*gen.GetTripResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTrip", varargs...)
	ret0, _ := ret[0].(*gen.GetTripResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrip indicates an expected call of GetTrip.
func (mr *MockTripsClientMockRecorder) GetTrip(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrip", reflect.TypeOf((*MockTripsClient)(nil).GetTrip), varargs...)
}

// GetTripsByUserID mocks base method.
func (m *MockTripsClient) GetTripsByUserID(ctx context.Context, in *gen.GetTripsByUserIDRequest, opts ...grpc.CallOption) (*gen.GetTripsByUserIDResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTripsByUserID", varargs...)
	ret0, _ := ret[0].(*gen.GetTripsByUserIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTripsByUserID indicates an expected call of GetTripsByUserID.
func (mr *MockTripsClientMockRecorder) GetTripsByUserID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTripsByUserID", reflect.TypeOf((*MockTripsClient)(nil).GetTripsByUserID), varargs...)
}

// UpdateTrip mocks base method.
func (m *MockTripsClient) UpdateTrip(ctx context.Context, in *gen.UpdateTripRequest, opts ...grpc.CallOption) (*gen.EmptyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateTrip", varargs...)
	ret0, _ := ret[0].(*gen.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTrip indicates an expected call of UpdateTrip.
func (mr *MockTripsClientMockRecorder) UpdateTrip(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrip", reflect.TypeOf((*MockTripsClient)(nil).UpdateTrip), varargs...)
}

// MockTripsServer is a mock of TripsServer interface.
type MockTripsServer struct {
	ctrl     *gomock.Controller
	recorder *MockTripsServerMockRecorder
}

// MockTripsServerMockRecorder is the mock recorder for MockTripsServer.
type MockTripsServerMockRecorder struct {
	mock *MockTripsServer
}

// NewMockTripsServer creates a new mock instance.
func NewMockTripsServer(ctrl *gomock.Controller) *MockTripsServer {
	mock := &MockTripsServer{ctrl: ctrl}
	mock.recorder = &MockTripsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTripsServer) EXPECT() *MockTripsServerMockRecorder {
	return m.recorder
}

// AddPhotosToTrip mocks base method.
func (m *MockTripsServer) AddPhotosToTrip(arg0 context.Context, arg1 *gen.AddPhotosToTripRequest) (*gen.AddPhotosToTripResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPhotosToTrip", arg0, arg1)
	ret0, _ := ret[0].(*gen.AddPhotosToTripResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPhotosToTrip indicates an expected call of AddPhotosToTrip.
func (mr *MockTripsServerMockRecorder) AddPhotosToTrip(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPhotosToTrip", reflect.TypeOf((*MockTripsServer)(nil).AddPhotosToTrip), arg0, arg1)
}

// AddPlaceToTrip mocks base method.
func (m *MockTripsServer) AddPlaceToTrip(arg0 context.Context, arg1 *gen.AddPlaceToTripRequest) (*gen.EmptyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPlaceToTrip", arg0, arg1)
	ret0, _ := ret[0].(*gen.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPlaceToTrip indicates an expected call of AddPlaceToTrip.
func (mr *MockTripsServerMockRecorder) AddPlaceToTrip(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPlaceToTrip", reflect.TypeOf((*MockTripsServer)(nil).AddPlaceToTrip), arg0, arg1)
}

// CreateTrip mocks base method.
func (m *MockTripsServer) CreateTrip(arg0 context.Context, arg1 *gen.CreateTripRequest) (*gen.EmptyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTrip", arg0, arg1)
	ret0, _ := ret[0].(*gen.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTrip indicates an expected call of CreateTrip.
func (mr *MockTripsServerMockRecorder) CreateTrip(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrip", reflect.TypeOf((*MockTripsServer)(nil).CreateTrip), arg0, arg1)
}

// DeletePhotoFromTrip mocks base method.
func (m *MockTripsServer) DeletePhotoFromTrip(arg0 context.Context, arg1 *gen.DeletePhotoRequest) (*gen.EmptyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePhotoFromTrip", arg0, arg1)
	ret0, _ := ret[0].(*gen.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePhotoFromTrip indicates an expected call of DeletePhotoFromTrip.
func (mr *MockTripsServerMockRecorder) DeletePhotoFromTrip(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePhotoFromTrip", reflect.TypeOf((*MockTripsServer)(nil).DeletePhotoFromTrip), arg0, arg1)
}

// DeleteTrip mocks base method.
func (m *MockTripsServer) DeleteTrip(arg0 context.Context, arg1 *gen.DeleteTripRequest) (*gen.EmptyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrip", arg0, arg1)
	ret0, _ := ret[0].(*gen.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTrip indicates an expected call of DeleteTrip.
func (mr *MockTripsServerMockRecorder) DeleteTrip(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrip", reflect.TypeOf((*MockTripsServer)(nil).DeleteTrip), arg0, arg1)
}

// GetTrip mocks base method.
func (m *MockTripsServer) GetTrip(arg0 context.Context, arg1 *gen.GetTripRequest) (*gen.GetTripResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrip", arg0, arg1)
	ret0, _ := ret[0].(*gen.GetTripResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrip indicates an expected call of GetTrip.
func (mr *MockTripsServerMockRecorder) GetTrip(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrip", reflect.TypeOf((*MockTripsServer)(nil).GetTrip), arg0, arg1)
}

// GetTripsByUserID mocks base method.
func (m *MockTripsServer) GetTripsByUserID(arg0 context.Context, arg1 *gen.GetTripsByUserIDRequest) (*gen.GetTripsByUserIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTripsByUserID", arg0, arg1)
	ret0, _ := ret[0].(*gen.GetTripsByUserIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTripsByUserID indicates an expected call of GetTripsByUserID.
func (mr *MockTripsServerMockRecorder) GetTripsByUserID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTripsByUserID", reflect.TypeOf((*MockTripsServer)(nil).GetTripsByUserID), arg0, arg1)
}

// UpdateTrip mocks base method.
func (m *MockTripsServer) UpdateTrip(arg0 context.Context, arg1 *gen.UpdateTripRequest) (*gen.EmptyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTrip", arg0, arg1)
	ret0, _ := ret[0].(*gen.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTrip indicates an expected call of UpdateTrip.
func (mr *MockTripsServerMockRecorder) UpdateTrip(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrip", reflect.TypeOf((*MockTripsServer)(nil).UpdateTrip), arg0, arg1)
}

// mustEmbedUnimplementedTripsServer mocks base method.
func (m *MockTripsServer) mustEmbedUnimplementedTripsServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedTripsServer")
}

// mustEmbedUnimplementedTripsServer indicates an expected call of mustEmbedUnimplementedTripsServer.
func (mr *MockTripsServerMockRecorder) mustEmbedUnimplementedTripsServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedTripsServer", reflect.TypeOf((*MockTripsServer)(nil).mustEmbedUnimplementedTripsServer))
}

// MockUnsafeTripsServer is a mock of UnsafeTripsServer interface.
type MockUnsafeTripsServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeTripsServerMockRecorder
}

// MockUnsafeTripsServerMockRecorder is the mock recorder for MockUnsafeTripsServer.
type MockUnsafeTripsServerMockRecorder struct {
	mock *MockUnsafeTripsServer
}

// NewMockUnsafeTripsServer creates a new mock instance.
func NewMockUnsafeTripsServer(ctrl *gomock.Controller) *MockUnsafeTripsServer {
	mock := &MockUnsafeTripsServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeTripsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeTripsServer) EXPECT() *MockUnsafeTripsServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedTripsServer mocks base method.
func (m *MockUnsafeTripsServer) mustEmbedUnimplementedTripsServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedTripsServer")
}

// mustEmbedUnimplementedTripsServer indicates an expected call of mustEmbedUnimplementedTripsServer.
func (mr *MockUnsafeTripsServerMockRecorder) mustEmbedUnimplementedTripsServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedTripsServer", reflect.TypeOf((*MockUnsafeTripsServer)(nil).mustEmbedUnimplementedTripsServer))
}