// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.1
// source: proto/trips.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateTripRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trip *Trip `protobuf:"bytes,1,opt,name=trip,proto3" json:"trip,omitempty"`
}

func (x *CreateTripRequest) Reset() {
	*x = CreateTripRequest{}
	mi := &file_proto_trips_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTripRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTripRequest) ProtoMessage() {}

func (x *CreateTripRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trips_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTripRequest.ProtoReflect.Descriptor instead.
func (*CreateTripRequest) Descriptor() ([]byte, []int) {
	return file_proto_trips_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTripRequest) GetTrip() *Trip {
	if x != nil {
		return x.Trip
	}
	return nil
}

type UpdateTripRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trip *Trip `protobuf:"bytes,1,opt,name=trip,proto3" json:"trip,omitempty"`
}

func (x *UpdateTripRequest) Reset() {
	*x = UpdateTripRequest{}
	mi := &file_proto_trips_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTripRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTripRequest) ProtoMessage() {}

func (x *UpdateTripRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trips_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTripRequest.ProtoReflect.Descriptor instead.
func (*UpdateTripRequest) Descriptor() ([]byte, []int) {
	return file_proto_trips_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateTripRequest) GetTrip() *Trip {
	if x != nil {
		return x.Trip
	}
	return nil
}

type DeleteTripRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTripRequest) Reset() {
	*x = DeleteTripRequest{}
	mi := &file_proto_trips_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTripRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTripRequest) ProtoMessage() {}

func (x *DeleteTripRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trips_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTripRequest.ProtoReflect.Descriptor instead.
func (*DeleteTripRequest) Descriptor() ([]byte, []int) {
	return file_proto_trips_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteTripRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetTripsByUserIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Limit  int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetTripsByUserIDRequest) Reset() {
	*x = GetTripsByUserIDRequest{}
	mi := &file_proto_trips_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTripsByUserIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripsByUserIDRequest) ProtoMessage() {}

func (x *GetTripsByUserIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trips_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripsByUserIDRequest.ProtoReflect.Descriptor instead.
func (*GetTripsByUserIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_trips_proto_rawDescGZIP(), []int{3}
}

func (x *GetTripsByUserIDRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetTripsByUserIDRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetTripsByUserIDRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetTripsByUserIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trips []*Trip `protobuf:"bytes,1,rep,name=trips,proto3" json:"trips,omitempty"`
}

func (x *GetTripsByUserIDResponse) Reset() {
	*x = GetTripsByUserIDResponse{}
	mi := &file_proto_trips_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTripsByUserIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripsByUserIDResponse) ProtoMessage() {}

func (x *GetTripsByUserIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trips_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripsByUserIDResponse.ProtoReflect.Descriptor instead.
func (*GetTripsByUserIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_trips_proto_rawDescGZIP(), []int{4}
}

func (x *GetTripsByUserIDResponse) GetTrips() []*Trip {
	if x != nil {
		return x.Trips
	}
	return nil
}

type GetTripRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TripId uint32 `protobuf:"varint,1,opt,name=trip_id,json=tripId,proto3" json:"trip_id,omitempty"`
}

func (x *GetTripRequest) Reset() {
	*x = GetTripRequest{}
	mi := &file_proto_trips_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTripRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripRequest) ProtoMessage() {}

func (x *GetTripRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trips_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripRequest.ProtoReflect.Descriptor instead.
func (*GetTripRequest) Descriptor() ([]byte, []int) {
	return file_proto_trips_proto_rawDescGZIP(), []int{5}
}

func (x *GetTripRequest) GetTripId() uint32 {
	if x != nil {
		return x.TripId
	}
	return 0
}

type GetTripResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trip *Trip `protobuf:"bytes,1,opt,name=trip,proto3" json:"trip,omitempty"`
}

func (x *GetTripResponse) Reset() {
	*x = GetTripResponse{}
	mi := &file_proto_trips_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTripResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripResponse) ProtoMessage() {}

func (x *GetTripResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trips_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripResponse.ProtoReflect.Descriptor instead.
func (*GetTripResponse) Descriptor() ([]byte, []int) {
	return file_proto_trips_proto_rawDescGZIP(), []int{6}
}

func (x *GetTripResponse) GetTrip() *Trip {
	if x != nil {
		return x.Trip
	}
	return nil
}

type AddPlaceToTripRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TripId  uint32 `protobuf:"varint,1,opt,name=trip_id,json=tripId,proto3" json:"trip_id,omitempty"`
	PlaceId uint32 `protobuf:"varint,2,opt,name=place_id,json=placeId,proto3" json:"place_id,omitempty"`
}

func (x *AddPlaceToTripRequest) Reset() {
	*x = AddPlaceToTripRequest{}
	mi := &file_proto_trips_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddPlaceToTripRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPlaceToTripRequest) ProtoMessage() {}

func (x *AddPlaceToTripRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trips_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPlaceToTripRequest.ProtoReflect.Descriptor instead.
func (*AddPlaceToTripRequest) Descriptor() ([]byte, []int) {
	return file_proto_trips_proto_rawDescGZIP(), []int{7}
}

func (x *AddPlaceToTripRequest) GetTripId() uint32 {
	if x != nil {
		return x.TripId
	}
	return 0
}

func (x *AddPlaceToTripRequest) GetPlaceId() uint32 {
	if x != nil {
		return x.PlaceId
	}
	return 0
}

type AddPhotosToTripRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TripId uint32   `protobuf:"varint,1,opt,name=trip_id,json=tripId,proto3" json:"trip_id,omitempty"`
	Photos []string `protobuf:"bytes,2,rep,name=photos,proto3" json:"photos,omitempty"`
}

func (x *AddPhotosToTripRequest) Reset() {
	*x = AddPhotosToTripRequest{}
	mi := &file_proto_trips_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddPhotosToTripRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPhotosToTripRequest) ProtoMessage() {}

func (x *AddPhotosToTripRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trips_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPhotosToTripRequest.ProtoReflect.Descriptor instead.
func (*AddPhotosToTripRequest) Descriptor() ([]byte, []int) {
	return file_proto_trips_proto_rawDescGZIP(), []int{8}
}

func (x *AddPhotosToTripRequest) GetTripId() uint32 {
	if x != nil {
		return x.TripId
	}
	return 0
}

func (x *AddPhotosToTripRequest) GetPhotos() []string {
	if x != nil {
		return x.Photos
	}
	return nil
}

type AddPhotosToTripResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Photos []*Photo `protobuf:"bytes,2,rep,name=photos,proto3" json:"photos,omitempty"`
}

func (x *AddPhotosToTripResponse) Reset() {
	*x = AddPhotosToTripResponse{}
	mi := &file_proto_trips_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddPhotosToTripResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPhotosToTripResponse) ProtoMessage() {}

func (x *AddPhotosToTripResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trips_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPhotosToTripResponse.ProtoReflect.Descriptor instead.
func (*AddPhotosToTripResponse) Descriptor() ([]byte, []int) {
	return file_proto_trips_proto_rawDescGZIP(), []int{9}
}

func (x *AddPhotosToTripResponse) GetPhotos() []*Photo {
	if x != nil {
		return x.Photos
	}
	return nil
}

type Photo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhotoPath string `protobuf:"bytes,1,opt,name=photoPath,proto3" json:"photoPath,omitempty"`
}

func (x *Photo) Reset() {
	*x = Photo{}
	mi := &file_proto_trips_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Photo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Photo) ProtoMessage() {}

func (x *Photo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trips_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Photo.ProtoReflect.Descriptor instead.
func (*Photo) Descriptor() ([]byte, []int) {
	return file_proto_trips_proto_rawDescGZIP(), []int{10}
}

func (x *Photo) GetPhotoPath() string {
	if x != nil {
		return x.PhotoPath
	}
	return ""
}

type DeletePhotoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TripId    uint32 `protobuf:"varint,1,opt,name=trip_id,json=tripId,proto3" json:"trip_id,omitempty"`
	PhotoPath string `protobuf:"bytes,2,opt,name=photo_path,json=photoPath,proto3" json:"photo_path,omitempty"`
}

func (x *DeletePhotoRequest) Reset() {
	*x = DeletePhotoRequest{}
	mi := &file_proto_trips_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePhotoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePhotoRequest) ProtoMessage() {}

func (x *DeletePhotoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trips_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePhotoRequest.ProtoReflect.Descriptor instead.
func (*DeletePhotoRequest) Descriptor() ([]byte, []int) {
	return file_proto_trips_proto_rawDescGZIP(), []int{11}
}

func (x *DeletePhotoRequest) GetTripId() uint32 {
	if x != nil {
		return x.TripId
	}
	return 0
}

func (x *DeletePhotoRequest) GetPhotoPath() string {
	if x != nil {
		return x.PhotoPath
	}
	return ""
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	mi := &file_proto_trips_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trips_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_proto_trips_proto_rawDescGZIP(), []int{12}
}

type Trip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      uint32                 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name        string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CityId      uint32                 `protobuf:"varint,5,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`
	StartDate   string                 `protobuf:"bytes,6,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate     string                 `protobuf:"bytes,7,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Private     bool                   `protobuf:"varint,8,opt,name=private,proto3" json:"private,omitempty"`
	Photos      []string               `protobuf:"bytes,9,rep,name=photos,proto3" json:"photos,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Trip) Reset() {
	*x = Trip{}
	mi := &file_proto_trips_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Trip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trip) ProtoMessage() {}

func (x *Trip) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trips_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trip.ProtoReflect.Descriptor instead.
func (*Trip) Descriptor() ([]byte, []int) {
	return file_proto_trips_proto_rawDescGZIP(), []int{13}
}

func (x *Trip) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Trip) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Trip) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Trip) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Trip) GetCityId() uint32 {
	if x != nil {
		return x.CityId
	}
	return 0
}

func (x *Trip) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *Trip) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *Trip) GetPrivate() bool {
	if x != nil {
		return x.Private
	}
	return false
}

func (x *Trip) GetPhotos() []string {
	if x != nil {
		return x.Photos
	}
	return nil
}

func (x *Trip) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_proto_trips_proto protoreflect.FileDescriptor

var file_proto_trips_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x69, 0x70, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x72, 0x69, 0x70, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x04, 0x74, 0x72, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x74, 0x72, 0x69, 0x70, 0x73, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x52, 0x04, 0x74, 0x72, 0x69,
	0x70, 0x22, 0x34, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x74, 0x72, 0x69, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x73, 0x2e, 0x54, 0x72, 0x69,
	0x70, 0x52, 0x04, 0x74, 0x72, 0x69, 0x70, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x60, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x3d,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x74, 0x72,
	0x69, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x72, 0x69, 0x70,
	0x73, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x52, 0x05, 0x74, 0x72, 0x69, 0x70, 0x73, 0x22, 0x29, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x74, 0x72, 0x69, 0x70, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x74,
	0x72, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x72, 0x69, 0x70,
	0x73, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x52, 0x04, 0x74, 0x72, 0x69, 0x70, 0x22, 0x4b, 0x0a, 0x15,
	0x41, 0x64, 0x64, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x54, 0x6f, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x72, 0x69, 0x70, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x16, 0x41, 0x64, 0x64,
	0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x54, 0x6f, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x72, 0x69, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x73, 0x22, 0x3f, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x50, 0x68, 0x6f, 0x74, 0x6f,
	0x73, 0x54, 0x6f, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x73, 0x2e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x06, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x25, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x50, 0x61, 0x74, 0x68, 0x22, 0x4c, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x72, 0x69, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x50, 0x61, 0x74, 0x68, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa5, 0x02, 0x0a, 0x04,
	0x54, 0x72, 0x69, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x32, 0xb0, 0x04, 0x0a, 0x05, 0x54, 0x72, 0x69, 0x70, 0x73, 0x12, 0x3c, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x70, 0x12, 0x18, 0x2e, 0x74, 0x72,
	0x69, 0x70, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x73, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x70, 0x12, 0x18, 0x2e, 0x74, 0x72, 0x69, 0x70,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x72, 0x69, 0x70, 0x12, 0x18, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x73, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x72,
	0x69, 0x70, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x2e, 0x74, 0x72,
	0x69, 0x70, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x73, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x72,
	0x69, 0x70, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x73, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x12, 0x15, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x74, 0x72, 0x69, 0x70, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x54, 0x6f, 0x54, 0x72, 0x69, 0x70, 0x12, 0x1c, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x73,
	0x2e, 0x41, 0x64, 0x64, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x54, 0x6f, 0x54, 0x72, 0x69, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x73, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0f,
	0x41, 0x64, 0x64, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x54, 0x6f, 0x54, 0x72, 0x69, 0x70, 0x12,
	0x1d, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x68, 0x6f, 0x74, 0x6f,
	0x73, 0x54, 0x6f, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x74, 0x72, 0x69, 0x70, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73,
	0x54, 0x6f, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x46, 0x72, 0x6f,
	0x6d, 0x54, 0x72, 0x69, 0x70, 0x12, 0x19, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x73, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x69, 0x70, 0x73, 0x2f, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x3b, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_trips_proto_rawDescOnce sync.Once
	file_proto_trips_proto_rawDescData = file_proto_trips_proto_rawDesc
)

func file_proto_trips_proto_rawDescGZIP() []byte {
	file_proto_trips_proto_rawDescOnce.Do(func() {
		file_proto_trips_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_trips_proto_rawDescData)
	})
	return file_proto_trips_proto_rawDescData
}

var file_proto_trips_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_proto_trips_proto_goTypes = []any{
	(*CreateTripRequest)(nil),        // 0: trips.CreateTripRequest
	(*UpdateTripRequest)(nil),        // 1: trips.UpdateTripRequest
	(*DeleteTripRequest)(nil),        // 2: trips.DeleteTripRequest
	(*GetTripsByUserIDRequest)(nil),  // 3: trips.GetTripsByUserIDRequest
	(*GetTripsByUserIDResponse)(nil), // 4: trips.GetTripsByUserIDResponse
	(*GetTripRequest)(nil),           // 5: trips.GetTripRequest
	(*GetTripResponse)(nil),          // 6: trips.GetTripResponse
	(*AddPlaceToTripRequest)(nil),    // 7: trips.AddPlaceToTripRequest
	(*AddPhotosToTripRequest)(nil),   // 8: trips.AddPhotosToTripRequest
	(*AddPhotosToTripResponse)(nil),  // 9: trips.AddPhotosToTripResponse
	(*Photo)(nil),                    // 10: trips.Photo
	(*DeletePhotoRequest)(nil),       // 11: trips.DeletePhotoRequest
	(*EmptyResponse)(nil),            // 12: trips.EmptyResponse
	(*Trip)(nil),                     // 13: trips.Trip
	(*timestamppb.Timestamp)(nil),    // 14: google.protobuf.Timestamp
}
var file_proto_trips_proto_depIdxs = []int32{
	13, // 0: trips.CreateTripRequest.trip:type_name -> trips.Trip
	13, // 1: trips.UpdateTripRequest.trip:type_name -> trips.Trip
	13, // 2: trips.GetTripsByUserIDResponse.trips:type_name -> trips.Trip
	13, // 3: trips.GetTripResponse.trip:type_name -> trips.Trip
	10, // 4: trips.AddPhotosToTripResponse.photos:type_name -> trips.Photo
	14, // 5: trips.Trip.created_at:type_name -> google.protobuf.Timestamp
	0,  // 6: trips.Trips.CreateTrip:input_type -> trips.CreateTripRequest
	1,  // 7: trips.Trips.UpdateTrip:input_type -> trips.UpdateTripRequest
	2,  // 8: trips.Trips.DeleteTrip:input_type -> trips.DeleteTripRequest
	3,  // 9: trips.Trips.GetTripsByUserID:input_type -> trips.GetTripsByUserIDRequest
	5,  // 10: trips.Trips.GetTrip:input_type -> trips.GetTripRequest
	7,  // 11: trips.Trips.AddPlaceToTrip:input_type -> trips.AddPlaceToTripRequest
	8,  // 12: trips.Trips.AddPhotosToTrip:input_type -> trips.AddPhotosToTripRequest
	11, // 13: trips.Trips.DeletePhotoFromTrip:input_type -> trips.DeletePhotoRequest
	12, // 14: trips.Trips.CreateTrip:output_type -> trips.EmptyResponse
	12, // 15: trips.Trips.UpdateTrip:output_type -> trips.EmptyResponse
	12, // 16: trips.Trips.DeleteTrip:output_type -> trips.EmptyResponse
	4,  // 17: trips.Trips.GetTripsByUserID:output_type -> trips.GetTripsByUserIDResponse
	6,  // 18: trips.Trips.GetTrip:output_type -> trips.GetTripResponse
	12, // 19: trips.Trips.AddPlaceToTrip:output_type -> trips.EmptyResponse
	9,  // 20: trips.Trips.AddPhotosToTrip:output_type -> trips.AddPhotosToTripResponse
	12, // 21: trips.Trips.DeletePhotoFromTrip:output_type -> trips.EmptyResponse
	14, // [14:22] is the sub-list for method output_type
	6,  // [6:14] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_trips_proto_init() }
func file_proto_trips_proto_init() {
	if File_proto_trips_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_trips_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_trips_proto_goTypes,
		DependencyIndexes: file_proto_trips_proto_depIdxs,
		MessageInfos:      file_proto_trips_proto_msgTypes,
	}.Build()
	File_proto_trips_proto = out.File
	file_proto_trips_proto_rawDesc = nil
	file_proto_trips_proto_goTypes = nil
	file_proto_trips_proto_depIdxs = nil
}
