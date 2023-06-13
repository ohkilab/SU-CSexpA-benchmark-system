// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: backend/resources.proto

package backend

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

type Status int32

const (
	Status_WAITING           Status = 0 // waiting for benchmark
	Status_IN_PROGRESS       Status = 1 // in progress
	Status_SUCCESS           Status = 2 // benchmark succeeded
	Status_CONNECTION_FAILED Status = 3 // failed to connect
	Status_VALIDATION_ERROR  Status = 4 // validation error
	Status_INTERNAL_ERROR    Status = 5 // backend error
	Status_TIMEOUT           Status = 6 // timeout
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "WAITING",
		1: "IN_PROGRESS",
		2: "SUCCESS",
		3: "CONNECTION_FAILED",
		4: "VALIDATION_ERROR",
		5: "INTERNAL_ERROR",
		6: "TIMEOUT",
	}
	Status_value = map[string]int32{
		"WAITING":           0,
		"IN_PROGRESS":       1,
		"SUCCESS":           2,
		"CONNECTION_FAILED": 3,
		"VALIDATION_ERROR":  4,
		"INTERNAL_ERROR":    5,
		"TIMEOUT":           6,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_backend_resources_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_backend_resources_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_backend_resources_proto_rawDescGZIP(), []int{0}
}

type Language int32

const (
	Language_PHP        Language = 0
	Language_GO         Language = 1
	Language_RUST       Language = 2
	Language_JAVASCRIPT Language = 3
	Language_CSHARP     Language = 4
	Language_CPP        Language = 5
	Language_RUBY       Language = 6
	Language_PYTHON     Language = 7
)

// Enum value maps for Language.
var (
	Language_name = map[int32]string{
		0: "PHP",
		1: "GO",
		2: "RUST",
		3: "JAVASCRIPT",
		4: "CSHARP",
		5: "CPP",
		6: "RUBY",
		7: "PYTHON",
	}
	Language_value = map[string]int32{
		"PHP":        0,
		"GO":         1,
		"RUST":       2,
		"JAVASCRIPT": 3,
		"CSHARP":     4,
		"CPP":        5,
		"RUBY":       6,
		"PYTHON":     7,
	}
)

func (x Language) Enum() *Language {
	p := new(Language)
	*p = x
	return p
}

func (x Language) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language) Descriptor() protoreflect.EnumDescriptor {
	return file_backend_resources_proto_enumTypes[1].Descriptor()
}

func (Language) Type() protoreflect.EnumType {
	return &file_backend_resources_proto_enumTypes[1]
}

func (x Language) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language.Descriptor instead.
func (Language) EnumDescriptor() ([]byte, []int) {
	return file_backend_resources_proto_rawDescGZIP(), []int{1}
}

type Role int32

const (
	Role_CONTESTANT Role = 0
	Role_GUEST      Role = 1
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "CONTESTANT",
		1: "GUEST",
	}
	Role_value = map[string]int32{
		"CONTESTANT": 0,
		"GUEST":      1,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_backend_resources_proto_enumTypes[2].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_backend_resources_proto_enumTypes[2]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_backend_resources_proto_rawDescGZIP(), []int{2}
}

type Contest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	StartAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	EndAt       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	SubmitLimit int32                  `protobuf:"varint,6,opt,name=submit_limit,json=submitLimit,proto3" json:"submit_limit,omitempty"`
	Slug        string                 `protobuf:"bytes,8,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (x *Contest) Reset() {
	*x = Contest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contest) ProtoMessage() {}

func (x *Contest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contest.ProtoReflect.Descriptor instead.
func (*Contest) Descriptor() ([]byte, []int) {
	return file_backend_resources_proto_rawDescGZIP(), []int{0}
}

func (x *Contest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Contest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Contest) GetStartAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartAt
	}
	return nil
}

func (x *Contest) GetEndAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndAt
	}
	return nil
}

func (x *Contest) GetSubmitLimit() int32 {
	if x != nil {
		return x.SubmitLimit
	}
	return 0
}

func (x *Contest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Role Role   `protobuf:"varint,4,opt,name=role,proto3,enum=backend.Role" json:"role,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_backend_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_backend_resources_proto_rawDescGZIP(), []int{1}
}

func (x *Group) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Group) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_CONTESTANT
}

type Submit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	GroupName    string                 `protobuf:"bytes,2,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	Score        int32                  `protobuf:"varint,4,opt,name=score,proto3" json:"score,omitempty"`
	Language     Language               `protobuf:"varint,5,opt,name=language,proto3,enum=backend.Language" json:"language,omitempty"`
	SubmitedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=submited_at,json=submitedAt,proto3" json:"submited_at,omitempty"`
	CompletedAt  *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=completed_at,json=completedAt,proto3,oneof" json:"completed_at,omitempty"` // it this field is not null, this submit is completed
	TaskResults  []*TaskResult          `protobuf:"bytes,8,rep,name=task_results,json=taskResults,proto3" json:"task_results,omitempty"`
	Status       Status                 `protobuf:"varint,9,opt,name=status,proto3,enum=backend.Status" json:"status,omitempty"`
	ErrorMessage *string                `protobuf:"bytes,10,opt,name=error_message,json=errorMessage,proto3,oneof" json:"error_message,omitempty"` // if the connection error occurs, then this field is filled
	TagCount     int32                  `protobuf:"varint,11,opt,name=tag_count,json=tagCount,proto3" json:"tag_count,omitempty"`
}

func (x *Submit) Reset() {
	*x = Submit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_resources_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Submit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Submit) ProtoMessage() {}

func (x *Submit) ProtoReflect() protoreflect.Message {
	mi := &file_backend_resources_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Submit.ProtoReflect.Descriptor instead.
func (*Submit) Descriptor() ([]byte, []int) {
	return file_backend_resources_proto_rawDescGZIP(), []int{2}
}

func (x *Submit) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Submit) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *Submit) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Submit) GetLanguage() Language {
	if x != nil {
		return x.Language
	}
	return Language_PHP
}

func (x *Submit) GetSubmitedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SubmitedAt
	}
	return nil
}

func (x *Submit) GetCompletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CompletedAt
	}
	return nil
}

func (x *Submit) GetTaskResults() []*TaskResult {
	if x != nil {
		return x.TaskResults
	}
	return nil
}

func (x *Submit) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_WAITING
}

func (x *Submit) GetErrorMessage() string {
	if x != nil && x.ErrorMessage != nil {
		return *x.ErrorMessage
	}
	return ""
}

func (x *Submit) GetTagCount() int32 {
	if x != nil {
		return x.TagCount
	}
	return 0
}

type TaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RequestPerSec       int32                  `protobuf:"varint,2,opt,name=request_per_sec,json=requestPerSec,proto3" json:"request_per_sec,omitempty"`
	Url                 string                 `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Method              string                 `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
	RequestContentType  string                 `protobuf:"bytes,5,opt,name=request_content_type,json=requestContentType,proto3" json:"request_content_type,omitempty"`
	RequestBody         *string                `protobuf:"bytes,6,opt,name=request_body,json=requestBody,proto3,oneof" json:"request_body,omitempty"`
	ResponseCode        string                 `protobuf:"bytes,7,opt,name=response_code,json=responseCode,proto3" json:"response_code,omitempty"`
	ResponseContentType string                 `protobuf:"bytes,8,opt,name=response_content_type,json=responseContentType,proto3" json:"response_content_type,omitempty"`
	ResponseBody        string                 `protobuf:"bytes,9,opt,name=response_body,json=responseBody,proto3" json:"response_body,omitempty"`
	ThreadNum           int32                  `protobuf:"varint,10,opt,name=thread_num,json=threadNum,proto3" json:"thread_num,omitempty"`
	AttemptCount        int32                  `protobuf:"varint,11,opt,name=attempt_count,json=attemptCount,proto3" json:"attempt_count,omitempty"`
	AttemptTime         int32                  `protobuf:"varint,12,opt,name=attempt_time,json=attemptTime,proto3" json:"attempt_time,omitempty"`
	CreatedAt           *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	DeletedAt           *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=deleted_at,json=deletedAt,proto3,oneof" json:"deleted_at,omitempty"`
	ErrorMessage        *string                `protobuf:"bytes,15,opt,name=error_message,json=errorMessage,proto3,oneof" json:"error_message,omitempty"`
	Status              Status                 `protobuf:"varint,16,opt,name=status,proto3,enum=backend.Status" json:"status,omitempty"`
}

func (x *TaskResult) Reset() {
	*x = TaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_resources_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResult) ProtoMessage() {}

func (x *TaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_backend_resources_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResult.ProtoReflect.Descriptor instead.
func (*TaskResult) Descriptor() ([]byte, []int) {
	return file_backend_resources_proto_rawDescGZIP(), []int{3}
}

func (x *TaskResult) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskResult) GetRequestPerSec() int32 {
	if x != nil {
		return x.RequestPerSec
	}
	return 0
}

func (x *TaskResult) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *TaskResult) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *TaskResult) GetRequestContentType() string {
	if x != nil {
		return x.RequestContentType
	}
	return ""
}

func (x *TaskResult) GetRequestBody() string {
	if x != nil && x.RequestBody != nil {
		return *x.RequestBody
	}
	return ""
}

func (x *TaskResult) GetResponseCode() string {
	if x != nil {
		return x.ResponseCode
	}
	return ""
}

func (x *TaskResult) GetResponseContentType() string {
	if x != nil {
		return x.ResponseContentType
	}
	return ""
}

func (x *TaskResult) GetResponseBody() string {
	if x != nil {
		return x.ResponseBody
	}
	return ""
}

func (x *TaskResult) GetThreadNum() int32 {
	if x != nil {
		return x.ThreadNum
	}
	return 0
}

func (x *TaskResult) GetAttemptCount() int32 {
	if x != nil {
		return x.AttemptCount
	}
	return 0
}

func (x *TaskResult) GetAttemptTime() int32 {
	if x != nil {
		return x.AttemptTime
	}
	return 0
}

func (x *TaskResult) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *TaskResult) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *TaskResult) GetErrorMessage() string {
	if x != nil && x.ErrorMessage != nil {
		return *x.ErrorMessage
	}
	return ""
}

func (x *TaskResult) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_WAITING
}

var File_backend_resources_proto protoreflect.FileDescriptor

var file_backend_resources_proto_rawDesc = []byte{
	0x0a, 0x17, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x31, 0x0a, 0x06,
	0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x22, 0x3a, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x21, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x22, 0xc8, 0x03, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x42,
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x48, 0x00, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x36, 0x0a, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0b, 0x74,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x61, 0x67, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x74, 0x61, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x10, 0x0a, 0x0e, 0x5f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xad, 0x05,
	0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x65,
	0x72, 0x53, 0x65, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x30,
	0x0a, 0x14, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x26, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x42, 0x6f, 0x64, 0x79, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a,
	0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x61, 0x74,
	0x74, 0x65, 0x6d, 0x70, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x74,
	0x74, 0x65, 0x6d, 0x70, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0f, 0x0a, 0x0d, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x10, 0x0a, 0x0e, 0x5f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x81, 0x01,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x49, 0x54,
	0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47,
	0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04,
	0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10,
	0x06, 0x2a, 0x60, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x07, 0x0a,
	0x03, 0x50, 0x48, 0x50, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x4f, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x52, 0x55, 0x53, 0x54, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4a, 0x41, 0x56, 0x41,
	0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x53, 0x48, 0x41,
	0x52, 0x50, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x50, 0x50, 0x10, 0x05, 0x12, 0x08, 0x0a,
	0x04, 0x52, 0x55, 0x42, 0x59, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x59, 0x54, 0x48, 0x4f,
	0x4e, 0x10, 0x07, 0x2a, 0x21, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x43,
	0x4f, 0x4e, 0x54, 0x45, 0x53, 0x54, 0x41, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x47,
	0x55, 0x45, 0x53, 0x54, 0x10, 0x01, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x68, 0x6b, 0x69, 0x6c, 0x61, 0x62, 0x2f, 0x53, 0x55, 0x2d,
	0x43, 0x53, 0x65, 0x78, 0x70, 0x41, 0x2d, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b,
	0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backend_resources_proto_rawDescOnce sync.Once
	file_backend_resources_proto_rawDescData = file_backend_resources_proto_rawDesc
)

func file_backend_resources_proto_rawDescGZIP() []byte {
	file_backend_resources_proto_rawDescOnce.Do(func() {
		file_backend_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_backend_resources_proto_rawDescData)
	})
	return file_backend_resources_proto_rawDescData
}

var file_backend_resources_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_backend_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_backend_resources_proto_goTypes = []interface{}{
	(Status)(0),                   // 0: backend.Status
	(Language)(0),                 // 1: backend.Language
	(Role)(0),                     // 2: backend.Role
	(*Contest)(nil),               // 3: backend.Contest
	(*Group)(nil),                 // 4: backend.Group
	(*Submit)(nil),                // 5: backend.Submit
	(*TaskResult)(nil),            // 6: backend.TaskResult
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_backend_resources_proto_depIdxs = []int32{
	7,  // 0: backend.Contest.start_at:type_name -> google.protobuf.Timestamp
	7,  // 1: backend.Contest.end_at:type_name -> google.protobuf.Timestamp
	2,  // 2: backend.Group.role:type_name -> backend.Role
	1,  // 3: backend.Submit.language:type_name -> backend.Language
	7,  // 4: backend.Submit.submited_at:type_name -> google.protobuf.Timestamp
	7,  // 5: backend.Submit.completed_at:type_name -> google.protobuf.Timestamp
	6,  // 6: backend.Submit.task_results:type_name -> backend.TaskResult
	0,  // 7: backend.Submit.status:type_name -> backend.Status
	7,  // 8: backend.TaskResult.created_at:type_name -> google.protobuf.Timestamp
	7,  // 9: backend.TaskResult.deleted_at:type_name -> google.protobuf.Timestamp
	0,  // 10: backend.TaskResult.status:type_name -> backend.Status
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_backend_resources_proto_init() }
func file_backend_resources_proto_init() {
	if File_backend_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_backend_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_backend_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_backend_resources_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Submit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_backend_resources_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_backend_resources_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_backend_resources_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_backend_resources_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_backend_resources_proto_goTypes,
		DependencyIndexes: file_backend_resources_proto_depIdxs,
		EnumInfos:         file_backend_resources_proto_enumTypes,
		MessageInfos:      file_backend_resources_proto_msgTypes,
	}.Build()
	File_backend_resources_proto = out.File
	file_backend_resources_proto_rawDesc = nil
	file_backend_resources_proto_goTypes = nil
	file_backend_resources_proto_depIdxs = nil
}
