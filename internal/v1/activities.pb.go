// Copyright 2020 StrongDM Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.3
// source: activities.proto

package v1

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

// ActivityGetRequest specifies which Activity to retrieve.
type ActivityGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *GetRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The unique identifier of the Activity to retrieve.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ActivityGetRequest) Reset() {
	*x = ActivityGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityGetRequest) ProtoMessage() {}

func (x *ActivityGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_activities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityGetRequest.ProtoReflect.Descriptor instead.
func (*ActivityGetRequest) Descriptor() ([]byte, []int) {
	return file_activities_proto_rawDescGZIP(), []int{0}
}

func (x *ActivityGetRequest) GetMeta() *GetRequestMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ActivityGetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// ActivityGetResponse returns a requested Activity.
type ActivityGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *GetResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The requested Activity.
	Activity *Activity `protobuf:"bytes,2,opt,name=activity,proto3" json:"activity,omitempty"`
	// Rate limit information.
	RateLimit *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *ActivityGetResponse) Reset() {
	*x = ActivityGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityGetResponse) ProtoMessage() {}

func (x *ActivityGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_activities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityGetResponse.ProtoReflect.Descriptor instead.
func (*ActivityGetResponse) Descriptor() ([]byte, []int) {
	return file_activities_proto_rawDescGZIP(), []int{1}
}

func (x *ActivityGetResponse) GetMeta() *GetResponseMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ActivityGetResponse) GetActivity() *Activity {
	if x != nil {
		return x.Activity
	}
	return nil
}

func (x *ActivityGetResponse) GetRateLimit() *RateLimitMetadata {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

// ActivityListRequest specifies criteria for retrieving a list of Activities.
type ActivityListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Paging parameters for the query.
	Meta *ListRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A human-readable filter query string.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ActivityListRequest) Reset() {
	*x = ActivityListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activities_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityListRequest) ProtoMessage() {}

func (x *ActivityListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_activities_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityListRequest.ProtoReflect.Descriptor instead.
func (*ActivityListRequest) Descriptor() ([]byte, []int) {
	return file_activities_proto_rawDescGZIP(), []int{2}
}

func (x *ActivityListRequest) GetMeta() *ListRequestMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ActivityListRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// ActivityListResponse returns a list of Activities that meet the criteria of a
// ActivityListRequest.
type ActivityListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Paging information for the query.
	Meta *ListResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A single page of results matching the list request criteria.
	Activities []*Activity `protobuf:"bytes,2,rep,name=activities,proto3" json:"activities,omitempty"`
	// Rate limit information.
	RateLimit *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *ActivityListResponse) Reset() {
	*x = ActivityListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activities_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityListResponse) ProtoMessage() {}

func (x *ActivityListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_activities_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityListResponse.ProtoReflect.Descriptor instead.
func (*ActivityListResponse) Descriptor() ([]byte, []int) {
	return file_activities_proto_rawDescGZIP(), []int{3}
}

func (x *ActivityListResponse) GetMeta() *ListResponseMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ActivityListResponse) GetActivities() []*Activity {
	if x != nil {
		return x.Activities
	}
	return nil
}

func (x *ActivityListResponse) GetRateLimit() *RateLimitMetadata {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

// An Activity is a record of an action taken against a strongDM deployment, e.g.
// a user creation, resource deletion, sso configuration change, etc.
type Activity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier of the Activity.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The kind of activity which has taken place, one of the ActivityVerb constants.
	Verb string `protobuf:"bytes,2,opt,name=verb,proto3" json:"verb,omitempty"`
	// A humanized description of the activity.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// The time this activity took effect.
	CompletedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=completed_at,json=completedAt,proto3" json:"completed_at,omitempty"`
	// The entities involved in this activity. These entities can be any first class
	// entity in the strongDM system, eg. a user, a role, a node, an account grant. Not
	// every activity affects explicit entities.
	Entities []*ActivityEntity `protobuf:"bytes,5,rep,name=entities,proto3" json:"entities,omitempty"`
	// The IP from which this action was taken.
	IpAddress string `protobuf:"bytes,6,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	// The account who executed this activity. If the actor later has a name or email change,
	// that change is not reflected here. Actor is a snapshot of the executing account at
	// the time an activity took place.
	Actor *ActivityActor `protobuf:"bytes,7,opt,name=actor,proto3" json:"actor,omitempty"`
	// The User Agent present when this request was executed. Generally a client type and version
	// like strongdm-cli/55.66.77
	UserAgent string `protobuf:"bytes,8,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
}

func (x *Activity) Reset() {
	*x = Activity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activities_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity) ProtoMessage() {}

func (x *Activity) ProtoReflect() protoreflect.Message {
	mi := &file_activities_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity.ProtoReflect.Descriptor instead.
func (*Activity) Descriptor() ([]byte, []int) {
	return file_activities_proto_rawDescGZIP(), []int{4}
}

func (x *Activity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Activity) GetVerb() string {
	if x != nil {
		return x.Verb
	}
	return ""
}

func (x *Activity) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Activity) GetCompletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CompletedAt
	}
	return nil
}

func (x *Activity) GetEntities() []*ActivityEntity {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *Activity) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *Activity) GetActor() *ActivityActor {
	if x != nil {
		return x.Actor
	}
	return nil
}

func (x *Activity) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

type ActivityEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of the entity this activity affected.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The type of entity affected, one of the ActivityEntityType constants.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// A display name representing the affected entity.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The email of the affected entity, if it has one (for example, if it is an account).
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	// The external ID of the affected entity, if it has one (for example, if it is an account).
	ExternalId string `protobuf:"bytes,5,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
}

func (x *ActivityEntity) Reset() {
	*x = ActivityEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activities_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityEntity) ProtoMessage() {}

func (x *ActivityEntity) ProtoReflect() protoreflect.Message {
	mi := &file_activities_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityEntity.ProtoReflect.Descriptor instead.
func (*ActivityEntity) Descriptor() ([]byte, []int) {
	return file_activities_proto_rawDescGZIP(), []int{5}
}

func (x *ActivityEntity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ActivityEntity) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ActivityEntity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ActivityEntity) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ActivityEntity) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

type ActivityActor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier of the actor. Immutable.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The email of the actor at the time this activity occurred.
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// The first name of the actor at the time this activity occurred.
	FirstName string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	// The last name of the actor at the time this activity occurred.
	LastName string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	// The external ID of the actor at the time this activity occurred.
	ActivityExternalId string `protobuf:"bytes,5,opt,name=activity_external_id,json=activityExternalId,proto3" json:"activity_external_id,omitempty"`
}

func (x *ActivityActor) Reset() {
	*x = ActivityActor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activities_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityActor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityActor) ProtoMessage() {}

func (x *ActivityActor) ProtoReflect() protoreflect.Message {
	mi := &file_activities_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityActor.ProtoReflect.Descriptor instead.
func (*ActivityActor) Descriptor() ([]byte, []int) {
	return file_activities_proto_rawDescGZIP(), []int{6}
}

func (x *ActivityActor) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ActivityActor) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ActivityActor) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *ActivityActor) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *ActivityActor) GetActivityExternalId() string {
	if x != nil {
		return x.ActivityExternalId
	}
	return ""
}

var File_activities_proto protoreflect.FileDescriptor

var file_activities_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x12, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x02, 0x69,
	0x64, 0x3a, 0x28, 0xfa, 0xf8, 0xb3, 0x07, 0x06, 0xd2, 0xf3, 0xb3, 0x07, 0x01, 0x2a, 0xfa, 0xf8,
	0xb3, 0x07, 0x18, 0xd2, 0xf3, 0xb3, 0x07, 0x13, 0x21, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f,
	0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x9c, 0x02, 0x0a, 0x13,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07,
	0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x08,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x42, 0x0a, 0xf2, 0xf8,
	0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x12, 0x62, 0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x2c, 0xf2,
	0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xf2, 0xf8, 0xb3, 0x07, 0x06, 0xb2, 0xf4,
	0xb3, 0x07, 0x01, 0x2a, 0xf2, 0xf8, 0xb3, 0x07, 0x12, 0xb2, 0xf4, 0xb3, 0x07, 0x0d, 0x21, 0x6a,
	0x73, 0x6f, 0x6e, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x09, 0x72, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x3a, 0x32, 0xfa, 0xf8, 0xb3, 0x07, 0x06, 0xd2, 0xf3, 0xb3,
	0x07, 0x01, 0x2a, 0xfa, 0xf8, 0xb3, 0x07, 0x18, 0xd2, 0xf3, 0xb3, 0x07, 0x13, 0x21, 0x74, 0x65,
	0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0xfa, 0xf8, 0xb3, 0x07, 0x05, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0x22, 0x90, 0x01, 0x0a, 0x13, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12,
	0x22, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x3a, 0x28, 0xfa, 0xf8, 0xb3, 0x07, 0x06, 0xd2, 0xf3, 0xb3, 0x07, 0x01, 0x2a,
	0xfa, 0xf8, 0xb3, 0x07, 0x18, 0xd2, 0xf3, 0xb3, 0x07, 0x13, 0x21, 0x74, 0x65, 0x72, 0x72, 0x61,
	0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x8c, 0x02,
	0x0a, 0x14, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb8, 0xf3, 0xb3,
	0x07, 0x01, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x62,
	0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x2c, 0xf2, 0xf8, 0xb3, 0x07, 0x05,
	0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xf2, 0xf8, 0xb3, 0x07, 0x06, 0xb2, 0xf4, 0xb3, 0x07, 0x01, 0x2a,
	0xf2, 0xf8, 0xb3, 0x07, 0x12, 0xb2, 0xf4, 0xb3, 0x07, 0x0d, 0x21, 0x6a, 0x73, 0x6f, 0x6e, 0x5f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x3a, 0x28, 0xfa, 0xf8, 0xb3, 0x07, 0x06, 0xd2, 0xf3, 0xb3, 0x07, 0x01, 0x2a, 0xfa,
	0xf8, 0xb3, 0x07, 0x18, 0xd2, 0xf3, 0xb3, 0x07, 0x13, 0x21, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66,
	0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0xd3, 0x03, 0x0a,
	0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x76, 0x65, 0x72, 0x62, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52,
	0x04, 0x76, 0x65, 0x72, 0x62, 0x12, 0x2c, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07,
	0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07,
	0x01, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3a,
	0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01,
	0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x0a, 0x69, 0x70,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23,
	0xf2, 0xf8, 0xb3, 0x07, 0x1e, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xca, 0xf3, 0xb3, 0x07, 0x14, 0xc2,
	0xf4, 0xb3, 0x07, 0x0f, 0x0a, 0x02, 0x67, 0x6f, 0x12, 0x09, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x33,
	0x0a, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x41, 0x63, 0x74, 0x6f, 0x72,
	0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x05, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x29, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3,
	0xb3, 0x07, 0x01, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x3a, 0x32,
	0xfa, 0xf8, 0xb3, 0x07, 0x05, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0xfa, 0xf8, 0xb3, 0x07, 0x06, 0xd2,
	0xf3, 0xb3, 0x07, 0x01, 0x2a, 0xfa, 0xf8, 0xb3, 0x07, 0x18, 0xd2, 0xf3, 0xb3, 0x07, 0x13, 0x21,
	0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x22, 0xef, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1e, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x2b, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0,
	0xf3, 0xb3, 0x07, 0x01, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64,
	0x3a, 0x32, 0xfa, 0xf8, 0xb3, 0x07, 0x05, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0xfa, 0xf8, 0xb3, 0x07,
	0x06, 0xd2, 0xf3, 0xb3, 0x07, 0x01, 0x2a, 0xfa, 0xf8, 0xb3, 0x07, 0x18, 0xd2, 0xf3, 0xb3, 0x07,
	0x13, 0x21, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x22, 0x93, 0x02, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x20, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x29, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0,
	0xf3, 0xb3, 0x07, 0x01, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x27, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x14, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3,
	0x07, 0x01, 0x52, 0x12, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x3a, 0x32, 0xfa, 0xf8, 0xb3, 0x07, 0x05, 0xa8, 0xf3, 0xb3,
	0x07, 0x01, 0xfa, 0xf8, 0xb3, 0x07, 0x06, 0xd2, 0xf3, 0xb3, 0x07, 0x01, 0x2a, 0xfa, 0xf8, 0xb3,
	0x07, 0x18, 0xd2, 0xf3, 0xb3, 0x07, 0x13, 0x21, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72,
	0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x32, 0xa5, 0x02, 0x0a, 0x0a, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x62, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2a, 0x82, 0xf9, 0xb3, 0x07, 0x08, 0xa2, 0xf3, 0xb3, 0x07, 0x03, 0x67, 0x65, 0x74,
	0x82, 0xf9, 0xb3, 0x07, 0x18, 0xaa, 0xf3, 0xb3, 0x07, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x60, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xf9, 0xb3, 0x07, 0x08, 0xa2,
	0xf3, 0xb3, 0x07, 0x03, 0x67, 0x65, 0x74, 0x82, 0xf9, 0xb3, 0x07, 0x13, 0xaa, 0xf3, 0xb3, 0x07,
	0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x1a,
	0x51, 0xca, 0xf9, 0xb3, 0x07, 0x0d, 0xc2, 0xf9, 0xb3, 0x07, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0xca, 0xf9, 0xb3, 0x07, 0x08, 0xd2, 0xf9, 0xb3, 0x07, 0x03, 0x61, 0x74, 0x2d,
	0xca, 0xf9, 0xb3, 0x07, 0x05, 0xd8, 0xf9, 0xb3, 0x07, 0x01, 0xca, 0xf9, 0xb3, 0x07, 0x06, 0xca,
	0xf9, 0xb3, 0x07, 0x01, 0x2a, 0xca, 0xf9, 0xb3, 0x07, 0x18, 0xca, 0xf9, 0xb3, 0x07, 0x13, 0x21,
	0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x42, 0x8e, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x6e,
	0x67, 0x64, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6e, 0x67,
	0x42, 0x12, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x50, 0x6c, 0x75, 0x6d,
	0x62, 0x69, 0x6e, 0x67, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67,
	0x64, 0x6d, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0xc2, 0x92, 0xb4, 0x07, 0x06,
	0xa2, 0x8c, 0xb4, 0x07, 0x01, 0x2a, 0xc2, 0x92, 0xb4, 0x07, 0x18, 0xa2, 0x8c, 0xb4, 0x07, 0x13,
	0x21, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_activities_proto_rawDescOnce sync.Once
	file_activities_proto_rawDescData = file_activities_proto_rawDesc
)

func file_activities_proto_rawDescGZIP() []byte {
	file_activities_proto_rawDescOnce.Do(func() {
		file_activities_proto_rawDescData = protoimpl.X.CompressGZIP(file_activities_proto_rawDescData)
	})
	return file_activities_proto_rawDescData
}

var file_activities_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_activities_proto_goTypes = []interface{}{
	(*ActivityGetRequest)(nil),    // 0: v1.ActivityGetRequest
	(*ActivityGetResponse)(nil),   // 1: v1.ActivityGetResponse
	(*ActivityListRequest)(nil),   // 2: v1.ActivityListRequest
	(*ActivityListResponse)(nil),  // 3: v1.ActivityListResponse
	(*Activity)(nil),              // 4: v1.Activity
	(*ActivityEntity)(nil),        // 5: v1.ActivityEntity
	(*ActivityActor)(nil),         // 6: v1.ActivityActor
	(*GetRequestMetadata)(nil),    // 7: v1.GetRequestMetadata
	(*GetResponseMetadata)(nil),   // 8: v1.GetResponseMetadata
	(*RateLimitMetadata)(nil),     // 9: v1.RateLimitMetadata
	(*ListRequestMetadata)(nil),   // 10: v1.ListRequestMetadata
	(*ListResponseMetadata)(nil),  // 11: v1.ListResponseMetadata
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
}
var file_activities_proto_depIdxs = []int32{
	7,  // 0: v1.ActivityGetRequest.meta:type_name -> v1.GetRequestMetadata
	8,  // 1: v1.ActivityGetResponse.meta:type_name -> v1.GetResponseMetadata
	4,  // 2: v1.ActivityGetResponse.activity:type_name -> v1.Activity
	9,  // 3: v1.ActivityGetResponse.rate_limit:type_name -> v1.RateLimitMetadata
	10, // 4: v1.ActivityListRequest.meta:type_name -> v1.ListRequestMetadata
	11, // 5: v1.ActivityListResponse.meta:type_name -> v1.ListResponseMetadata
	4,  // 6: v1.ActivityListResponse.activities:type_name -> v1.Activity
	9,  // 7: v1.ActivityListResponse.rate_limit:type_name -> v1.RateLimitMetadata
	12, // 8: v1.Activity.completed_at:type_name -> google.protobuf.Timestamp
	5,  // 9: v1.Activity.entities:type_name -> v1.ActivityEntity
	6,  // 10: v1.Activity.actor:type_name -> v1.ActivityActor
	0,  // 11: v1.Activities.Get:input_type -> v1.ActivityGetRequest
	2,  // 12: v1.Activities.List:input_type -> v1.ActivityListRequest
	1,  // 13: v1.Activities.Get:output_type -> v1.ActivityGetResponse
	3,  // 14: v1.Activities.List:output_type -> v1.ActivityListResponse
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_activities_proto_init() }
func file_activities_proto_init() {
	if File_activities_proto != nil {
		return
	}
	file_options_proto_init()
	file_spec_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_activities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityGetRequest); i {
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
		file_activities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityGetResponse); i {
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
		file_activities_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityListRequest); i {
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
		file_activities_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityListResponse); i {
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
		file_activities_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activity); i {
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
		file_activities_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityEntity); i {
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
		file_activities_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityActor); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_activities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_activities_proto_goTypes,
		DependencyIndexes: file_activities_proto_depIdxs,
		MessageInfos:      file_activities_proto_msgTypes,
	}.Build()
	File_activities_proto = out.File
	file_activities_proto_rawDesc = nil
	file_activities_proto_goTypes = nil
	file_activities_proto_depIdxs = nil
}
