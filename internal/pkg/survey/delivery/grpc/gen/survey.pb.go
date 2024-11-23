// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: proto/survey.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetSurveyByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetSurveyByIdRequest) Reset() {
	*x = GetSurveyByIdRequest{}
	mi := &file_proto_survey_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSurveyByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSurveyByIdRequest) ProtoMessage() {}

func (x *GetSurveyByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_survey_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSurveyByIdRequest.ProtoReflect.Descriptor instead.
func (*GetSurveyByIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_survey_proto_rawDescGZIP(), []int{0}
}

func (x *GetSurveyByIdRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Survey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SurveyText string `protobuf:"bytes,2,opt,name=survey_text,json=surveyText,proto3" json:"survey_text,omitempty"`
	MaxRating  uint32 `protobuf:"varint,3,opt,name=max_rating,json=maxRating,proto3" json:"max_rating,omitempty"`
}

func (x *Survey) Reset() {
	*x = Survey{}
	mi := &file_proto_survey_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Survey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Survey) ProtoMessage() {}

func (x *Survey) ProtoReflect() protoreflect.Message {
	mi := &file_proto_survey_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Survey.ProtoReflect.Descriptor instead.
func (*Survey) Descriptor() ([]byte, []int) {
	return file_proto_survey_proto_rawDescGZIP(), []int{1}
}

func (x *Survey) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Survey) GetSurveyText() string {
	if x != nil {
		return x.SurveyText
	}
	return ""
}

func (x *Survey) GetMaxRating() uint32 {
	if x != nil {
		return x.MaxRating
	}
	return 0
}

type SurveyResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SurveyId    uint32 `protobuf:"varint,1,opt,name=survey_id,json=surveyId,proto3" json:"survey_id,omitempty"`
	UserId      uint32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Rating      uint32 `protobuf:"varint,4,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *SurveyResponce) Reset() {
	*x = SurveyResponce{}
	mi := &file_proto_survey_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SurveyResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyResponce) ProtoMessage() {}

func (x *SurveyResponce) ProtoReflect() protoreflect.Message {
	mi := &file_proto_survey_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyResponce.ProtoReflect.Descriptor instead.
func (*SurveyResponce) Descriptor() ([]byte, []int) {
	return file_proto_survey_proto_rawDescGZIP(), []int{2}
}

func (x *SurveyResponce) GetSurveyId() uint32 {
	if x != nil {
		return x.SurveyId
	}
	return 0
}

func (x *SurveyResponce) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SurveyResponce) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SurveyResponce) GetRating() uint32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type GetSurveyByIdResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Survey *Survey `protobuf:"bytes,1,opt,name=survey,proto3" json:"survey,omitempty"`
}

func (x *GetSurveyByIdResponce) Reset() {
	*x = GetSurveyByIdResponce{}
	mi := &file_proto_survey_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSurveyByIdResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSurveyByIdResponce) ProtoMessage() {}

func (x *GetSurveyByIdResponce) ProtoReflect() protoreflect.Message {
	mi := &file_proto_survey_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSurveyByIdResponce.ProtoReflect.Descriptor instead.
func (*GetSurveyByIdResponce) Descriptor() ([]byte, []int) {
	return file_proto_survey_proto_rawDescGZIP(), []int{3}
}

func (x *GetSurveyByIdResponce) GetSurvey() *Survey {
	if x != nil {
		return x.Survey
	}
	return nil
}

type CreateSurveyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServeyResponce *SurveyResponce `protobuf:"bytes,1,opt,name=serveyResponce,proto3" json:"serveyResponce,omitempty"`
}

func (x *CreateSurveyRequest) Reset() {
	*x = CreateSurveyRequest{}
	mi := &file_proto_survey_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSurveyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSurveyRequest) ProtoMessage() {}

func (x *CreateSurveyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_survey_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSurveyRequest.ProtoReflect.Descriptor instead.
func (*CreateSurveyRequest) Descriptor() ([]byte, []int) {
	return file_proto_survey_proto_rawDescGZIP(), []int{4}
}

func (x *CreateSurveyRequest) GetServeyResponce() *SurveyResponce {
	if x != nil {
		return x.ServeyResponce
	}
	return nil
}

type CreateSurveyResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CreateSurveyResponce) Reset() {
	*x = CreateSurveyResponce{}
	mi := &file_proto_survey_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSurveyResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSurveyResponce) ProtoMessage() {}

func (x *CreateSurveyResponce) ProtoReflect() protoreflect.Message {
	mi := &file_proto_survey_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSurveyResponce.ProtoReflect.Descriptor instead.
func (*CreateSurveyResponce) Descriptor() ([]byte, []int) {
	return file_proto_survey_proto_rawDescGZIP(), []int{5}
}

func (x *CreateSurveyResponce) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type SurveyStatsBySurvey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServeyId     uint32          `protobuf:"varint,1,opt,name=servey_id,json=serveyId,proto3" json:"servey_id,omitempty"`
	ServeyText   string          `protobuf:"bytes,2,opt,name=servey_text,json=serveyText,proto3" json:"servey_text,omitempty"`
	AvgRating    float32         `protobuf:"fixed32,3,opt,name=avg_rating,json=avgRating,proto3" json:"avg_rating,omitempty"`
	RatingsCount map[int32]int32 `protobuf:"bytes,4,rep,name=ratings_count,json=ratingsCount,proto3" json:"ratings_count,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *SurveyStatsBySurvey) Reset() {
	*x = SurveyStatsBySurvey{}
	mi := &file_proto_survey_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SurveyStatsBySurvey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyStatsBySurvey) ProtoMessage() {}

func (x *SurveyStatsBySurvey) ProtoReflect() protoreflect.Message {
	mi := &file_proto_survey_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyStatsBySurvey.ProtoReflect.Descriptor instead.
func (*SurveyStatsBySurvey) Descriptor() ([]byte, []int) {
	return file_proto_survey_proto_rawDescGZIP(), []int{6}
}

func (x *SurveyStatsBySurvey) GetServeyId() uint32 {
	if x != nil {
		return x.ServeyId
	}
	return 0
}

func (x *SurveyStatsBySurvey) GetServeyText() string {
	if x != nil {
		return x.ServeyText
	}
	return ""
}

func (x *SurveyStatsBySurvey) GetAvgRating() float32 {
	if x != nil {
		return x.AvgRating
	}
	return 0
}

func (x *SurveyStatsBySurvey) GetRatingsCount() map[int32]int32 {
	if x != nil {
		return x.RatingsCount
	}
	return nil
}

type GetSurveyStatsBySurveyIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetSurveyStatsBySurveyIdRequest) Reset() {
	*x = GetSurveyStatsBySurveyIdRequest{}
	mi := &file_proto_survey_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSurveyStatsBySurveyIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSurveyStatsBySurveyIdRequest) ProtoMessage() {}

func (x *GetSurveyStatsBySurveyIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_survey_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSurveyStatsBySurveyIdRequest.ProtoReflect.Descriptor instead.
func (*GetSurveyStatsBySurveyIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_survey_proto_rawDescGZIP(), []int{7}
}

func (x *GetSurveyStatsBySurveyIdRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetSurveyStatsBySurveyIdResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SurveyStatsBySurvey *SurveyStatsBySurvey `protobuf:"bytes,1,opt,name=surveyStatsBySurvey,proto3" json:"surveyStatsBySurvey,omitempty"`
}

func (x *GetSurveyStatsBySurveyIdResponce) Reset() {
	*x = GetSurveyStatsBySurveyIdResponce{}
	mi := &file_proto_survey_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSurveyStatsBySurveyIdResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSurveyStatsBySurveyIdResponce) ProtoMessage() {}

func (x *GetSurveyStatsBySurveyIdResponce) ProtoReflect() protoreflect.Message {
	mi := &file_proto_survey_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSurveyStatsBySurveyIdResponce.ProtoReflect.Descriptor instead.
func (*GetSurveyStatsBySurveyIdResponce) Descriptor() ([]byte, []int) {
	return file_proto_survey_proto_rawDescGZIP(), []int{8}
}

func (x *GetSurveyStatsBySurveyIdResponce) GetSurveyStatsBySurvey() *SurveyStatsBySurvey {
	if x != nil {
		return x.SurveyStatsBySurvey
	}
	return nil
}

type UserSurveyStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServeyId   uint32 `protobuf:"varint,1,opt,name=servey_id,json=serveyId,proto3" json:"servey_id,omitempty"`
	ServeyText string `protobuf:"bytes,2,opt,name=servey_text,json=serveyText,proto3" json:"servey_text,omitempty"`
	Answered   bool   `protobuf:"varint,3,opt,name=answered,proto3" json:"answered,omitempty"`
}

func (x *UserSurveyStats) Reset() {
	*x = UserSurveyStats{}
	mi := &file_proto_survey_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserSurveyStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSurveyStats) ProtoMessage() {}

func (x *UserSurveyStats) ProtoReflect() protoreflect.Message {
	mi := &file_proto_survey_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSurveyStats.ProtoReflect.Descriptor instead.
func (*UserSurveyStats) Descriptor() ([]byte, []int) {
	return file_proto_survey_proto_rawDescGZIP(), []int{9}
}

func (x *UserSurveyStats) GetServeyId() uint32 {
	if x != nil {
		return x.ServeyId
	}
	return 0
}

func (x *UserSurveyStats) GetServeyText() string {
	if x != nil {
		return x.ServeyText
	}
	return ""
}

func (x *UserSurveyStats) GetAnswered() bool {
	if x != nil {
		return x.Answered
	}
	return false
}

type GetSurveyStatsByUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetSurveyStatsByUserIdRequest) Reset() {
	*x = GetSurveyStatsByUserIdRequest{}
	mi := &file_proto_survey_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSurveyStatsByUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSurveyStatsByUserIdRequest) ProtoMessage() {}

func (x *GetSurveyStatsByUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_survey_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSurveyStatsByUserIdRequest.ProtoReflect.Descriptor instead.
func (*GetSurveyStatsByUserIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_survey_proto_rawDescGZIP(), []int{10}
}

func (x *GetSurveyStatsByUserIdRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetSurveyStatsByUserIdResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserServeyStats []*UserSurveyStats `protobuf:"bytes,1,rep,name=userServeyStats,proto3" json:"userServeyStats,omitempty"`
}

func (x *GetSurveyStatsByUserIdResponce) Reset() {
	*x = GetSurveyStatsByUserIdResponce{}
	mi := &file_proto_survey_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSurveyStatsByUserIdResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSurveyStatsByUserIdResponce) ProtoMessage() {}

func (x *GetSurveyStatsByUserIdResponce) ProtoReflect() protoreflect.Message {
	mi := &file_proto_survey_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSurveyStatsByUserIdResponce.ProtoReflect.Descriptor instead.
func (*GetSurveyStatsByUserIdResponce) Descriptor() ([]byte, []int) {
	return file_proto_survey_proto_rawDescGZIP(), []int{11}
}

func (x *GetSurveyStatsByUserIdResponce) GetUserServeyStats() []*UserSurveyStats {
	if x != nil {
		return x.UserServeyStats
	}
	return nil
}

var File_proto_survey_proto protoreflect.FileDescriptor

var file_proto_survey_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x22, 0x26, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x58, 0x0a, 0x06, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x54, 0x65, 0x78, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x80,
	0x01, 0x0a, 0x0e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x22, 0x3f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x06, 0x73, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x22, 0x55, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0x30, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x87, 0x02, 0x0a, 0x13,
	0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x42, 0x79, 0x53, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x79, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x79, 0x54, 0x65, 0x78,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x67, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x61, 0x76, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x52, 0x0a, 0x0d, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x42, 0x79, 0x53, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x3f, 0x0a, 0x11, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x31, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x42, 0x79, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x71, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x42, 0x79, 0x53, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x13,
	0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x42, 0x79, 0x53, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x42, 0x79,
	0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x13, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x42, 0x79, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x22, 0x6b, 0x0a, 0x0f, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x79, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x64, 0x22, 0x37, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x63, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x32, 0x88, 0x03, 0x0a, 0x0d, 0x53, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x2e, 0x73, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x12, 0x1b, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x63, 0x65, 0x22, 0x00, 0x12, 0x6f, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x42, 0x79, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x49,
	0x64, 0x12, 0x27, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x42, 0x79, 0x53, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x42, 0x79, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x63, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x25, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x22,
	0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_survey_proto_rawDescOnce sync.Once
	file_proto_survey_proto_rawDescData = file_proto_survey_proto_rawDesc
)

func file_proto_survey_proto_rawDescGZIP() []byte {
	file_proto_survey_proto_rawDescOnce.Do(func() {
		file_proto_survey_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_survey_proto_rawDescData)
	})
	return file_proto_survey_proto_rawDescData
}

var file_proto_survey_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_survey_proto_goTypes = []any{
	(*GetSurveyByIdRequest)(nil),             // 0: survey.GetSurveyByIdRequest
	(*Survey)(nil),                           // 1: survey.Survey
	(*SurveyResponce)(nil),                   // 2: survey.SurveyResponce
	(*GetSurveyByIdResponce)(nil),            // 3: survey.GetSurveyByIdResponce
	(*CreateSurveyRequest)(nil),              // 4: survey.CreateSurveyRequest
	(*CreateSurveyResponce)(nil),             // 5: survey.CreateSurveyResponce
	(*SurveyStatsBySurvey)(nil),              // 6: survey.SurveyStatsBySurvey
	(*GetSurveyStatsBySurveyIdRequest)(nil),  // 7: survey.GetSurveyStatsBySurveyIdRequest
	(*GetSurveyStatsBySurveyIdResponce)(nil), // 8: survey.GetSurveyStatsBySurveyIdResponce
	(*UserSurveyStats)(nil),                  // 9: survey.UserSurveyStats
	(*GetSurveyStatsByUserIdRequest)(nil),    // 10: survey.GetSurveyStatsByUserIdRequest
	(*GetSurveyStatsByUserIdResponce)(nil),   // 11: survey.GetSurveyStatsByUserIdResponce
	nil,                                      // 12: survey.SurveyStatsBySurvey.RatingsCountEntry
}
var file_proto_survey_proto_depIdxs = []int32{
	1,  // 0: survey.GetSurveyByIdResponce.survey:type_name -> survey.Survey
	2,  // 1: survey.CreateSurveyRequest.serveyResponce:type_name -> survey.SurveyResponce
	12, // 2: survey.SurveyStatsBySurvey.ratings_count:type_name -> survey.SurveyStatsBySurvey.RatingsCountEntry
	6,  // 3: survey.GetSurveyStatsBySurveyIdResponce.surveyStatsBySurvey:type_name -> survey.SurveyStatsBySurvey
	9,  // 4: survey.GetSurveyStatsByUserIdResponce.userServeyStats:type_name -> survey.UserSurveyStats
	0,  // 5: survey.SurveyService.GetSurveyById:input_type -> survey.GetSurveyByIdRequest
	4,  // 6: survey.SurveyService.CreateSurvey:input_type -> survey.CreateSurveyRequest
	7,  // 7: survey.SurveyService.GetSurveyStatsBySurveyId:input_type -> survey.GetSurveyStatsBySurveyIdRequest
	10, // 8: survey.SurveyService.GetSurveyStatsByUserId:input_type -> survey.GetSurveyStatsByUserIdRequest
	3,  // 9: survey.SurveyService.GetSurveyById:output_type -> survey.GetSurveyByIdResponce
	5,  // 10: survey.SurveyService.CreateSurvey:output_type -> survey.CreateSurveyResponce
	8,  // 11: survey.SurveyService.GetSurveyStatsBySurveyId:output_type -> survey.GetSurveyStatsBySurveyIdResponce
	11, // 12: survey.SurveyService.GetSurveyStatsByUserId:output_type -> survey.GetSurveyStatsByUserIdResponce
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_survey_proto_init() }
func file_proto_survey_proto_init() {
	if File_proto_survey_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_survey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_survey_proto_goTypes,
		DependencyIndexes: file_proto_survey_proto_depIdxs,
		MessageInfos:      file_proto_survey_proto_msgTypes,
	}.Build()
	File_proto_survey_proto = out.File
	file_proto_survey_proto_rawDesc = nil
	file_proto_survey_proto_goTypes = nil
	file_proto_survey_proto_depIdxs = nil
}
