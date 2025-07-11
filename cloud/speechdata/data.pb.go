// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.6
// source: cloud/speech/v1/data.proto

package speechdata

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Set of states that define the lifecycle of a CustomClass.
type CustomClass_State int32

const (
	// Unspecified state.  This is only used/useful for distinguishing
	// unset values.
	CustomClass_STATE_UNSPECIFIED CustomClass_State = 0
	// The normal and active state.
	CustomClass_ACTIVE CustomClass_State = 2
	// This CustomClass has been deleted.
	CustomClass_DELETED CustomClass_State = 4
)

// Enum value maps for CustomClass_State.
var (
	CustomClass_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		2: "ACTIVE",
		4: "DELETED",
	}
	CustomClass_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"ACTIVE":            2,
		"DELETED":           4,
	}
)

func (x CustomClass_State) Enum() *CustomClass_State {
	p := new(CustomClass_State)
	*p = x
	return p
}

func (x CustomClass_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustomClass_State) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_speech_v1_data_proto_enumTypes[0].Descriptor()
}

func (CustomClass_State) Type() protoreflect.EnumType {
	return &file_cloud_speech_v1_data_proto_enumTypes[0]
}

func (x CustomClass_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustomClass_State.Descriptor instead.
func (CustomClass_State) EnumDescriptor() ([]byte, []int) {
	return file_cloud_speech_v1_data_proto_rawDescGZIP(), []int{0, 0}
}

// Set of states that define the lifecycle of a CustomClass.
type PhraseSet_State int32

const (
	// Unspecified state.  This is only used/useful for distinguishing
	// unset values.
	PhraseSet_STATE_UNSPECIFIED PhraseSet_State = 0
	// The normal and active state.
	PhraseSet_ACTIVE PhraseSet_State = 2
	// This CustomClass has been deleted.
	PhraseSet_DELETED PhraseSet_State = 4
)

// Enum value maps for PhraseSet_State.
var (
	PhraseSet_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		2: "ACTIVE",
		4: "DELETED",
	}
	PhraseSet_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"ACTIVE":            2,
		"DELETED":           4,
	}
)

func (x PhraseSet_State) Enum() *PhraseSet_State {
	p := new(PhraseSet_State)
	*p = x
	return p
}

func (x PhraseSet_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PhraseSet_State) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_speech_v1_data_proto_enumTypes[1].Descriptor()
}

func (PhraseSet_State) Type() protoreflect.EnumType {
	return &file_cloud_speech_v1_data_proto_enumTypes[1]
}

func (x PhraseSet_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PhraseSet_State.Descriptor instead.
func (PhraseSet_State) EnumDescriptor() ([]byte, []int) {
	return file_cloud_speech_v1_data_proto_rawDescGZIP(), []int{1, 0}
}

// A set of words or phrases that represents a common concept likely to appear
// in your audio, for example a list of passenger ship names. CustomClass items
// can be substituted into placeholders that you set in PhraseSet phrases.
type CustomClass struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The resource name of the custom class.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// If this custom class is a resource, the custom_class_id is the resource id
	// of the CustomClass. Case sensitive.
	CustomClassId string `protobuf:"bytes,2,opt,name=custom_class_id,json=customClassId,proto3" json:"custom_class_id,omitempty"`
	// A collection of class items.
	Items []*CustomClass_ClassItem `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	// Output only. The [KMS key
	// name](https://cloud.google.com/kms/docs/resource-hierarchy#keys) with which
	// the content of the ClassItem is encrypted. The expected format is
	// `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}`.
	KmsKeyName string `protobuf:"bytes,6,opt,name=kms_key_name,json=kmsKeyName,proto3" json:"kms_key_name,omitempty"`
	// Output only. The [KMS key version
	// name](https://cloud.google.com/kms/docs/resource-hierarchy#key_versions)
	// with which content of the ClassItem is encrypted. The expected format is
	// `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}/cryptoKeyVersions/{crypto_key_version}`.
	KmsKeyVersionName string `protobuf:"bytes,7,opt,name=kms_key_version_name,json=kmsKeyVersionName,proto3" json:"kms_key_version_name,omitempty"`
	// Output only. System-assigned unique identifier for the CustomClass.
	// This field is not used.
	Uid string `protobuf:"bytes,8,opt,name=uid,proto3" json:"uid,omitempty"`
	// Output only. User-settable, human-readable name for the CustomClass. Must
	// be 63 characters or less. This field is not used.
	DisplayName string `protobuf:"bytes,9,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Output only. The CustomClass lifecycle state.
	// This field is not used.
	State CustomClass_State `protobuf:"varint,10,opt,name=state,proto3,enum=google.events.cloud.speech.v1.CustomClass_State" json:"state,omitempty"`
	// Output only. The time at which this resource was requested for deletion.
	// This field is not used.
	DeleteTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// Output only. The time at which this resource will be purged.
	// This field is not used.
	ExpireTime *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	// Output only. Allows users to store small amounts of arbitrary data.
	// Both the key and the value must be 63 characters or less each.
	// At most 100 annotations.
	// This field is not used.
	Annotations map[string]string `protobuf:"bytes,13,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Output only. This checksum is computed by the server based on the value of
	// other fields. This may be sent on update, undelete, and delete requests to
	// ensure the client has an up-to-date value before proceeding. This field is
	// not used.
	Etag string `protobuf:"bytes,14,opt,name=etag,proto3" json:"etag,omitempty"`
	// Output only. Whether or not this CustomClass is in the process of being
	// updated. This field is not used.
	Reconciling   bool `protobuf:"varint,15,opt,name=reconciling,proto3" json:"reconciling,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomClass) Reset() {
	*x = CustomClass{}
	mi := &file_cloud_speech_v1_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomClass) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomClass) ProtoMessage() {}

func (x *CustomClass) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_speech_v1_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomClass.ProtoReflect.Descriptor instead.
func (*CustomClass) Descriptor() ([]byte, []int) {
	return file_cloud_speech_v1_data_proto_rawDescGZIP(), []int{0}
}

func (x *CustomClass) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomClass) GetCustomClassId() string {
	if x != nil {
		return x.CustomClassId
	}
	return ""
}

func (x *CustomClass) GetItems() []*CustomClass_ClassItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CustomClass) GetKmsKeyName() string {
	if x != nil {
		return x.KmsKeyName
	}
	return ""
}

func (x *CustomClass) GetKmsKeyVersionName() string {
	if x != nil {
		return x.KmsKeyVersionName
	}
	return ""
}

func (x *CustomClass) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *CustomClass) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *CustomClass) GetState() CustomClass_State {
	if x != nil {
		return x.State
	}
	return CustomClass_STATE_UNSPECIFIED
}

func (x *CustomClass) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

func (x *CustomClass) GetExpireTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireTime
	}
	return nil
}

func (x *CustomClass) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *CustomClass) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *CustomClass) GetReconciling() bool {
	if x != nil {
		return x.Reconciling
	}
	return false
}

// Provides "hints" to the speech recognizer to favor specific words and phrases
// in the results.
type PhraseSet struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The resource name of the phrase set.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A list of word and phrases.
	Phrases []*PhraseSet_Phrase `protobuf:"bytes,2,rep,name=phrases,proto3" json:"phrases,omitempty"`
	// Hint Boost. Positive value will increase the probability that a specific
	// phrase will be recognized over other similar sounding phrases. The higher
	// the boost, the higher the chance of false positive recognition as well.
	// Negative boost values would correspond to anti-biasing. Anti-biasing is not
	// enabled, so negative boost will simply be ignored. Though `boost` can
	// accept a wide range of positive values, most use cases are best served with
	// values between 0 (exclusive) and 20. We recommend using a binary search
	// approach to finding the optimal value for your use case as well as adding
	// phrases both with and without boost to your requests.
	Boost float32 `protobuf:"fixed32,4,opt,name=boost,proto3" json:"boost,omitempty"`
	// Output only. The [KMS key
	// name](https://cloud.google.com/kms/docs/resource-hierarchy#keys) with which
	// the content of the PhraseSet is encrypted. The expected format is
	// `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}`.
	KmsKeyName string `protobuf:"bytes,7,opt,name=kms_key_name,json=kmsKeyName,proto3" json:"kms_key_name,omitempty"`
	// Output only. The [KMS key version
	// name](https://cloud.google.com/kms/docs/resource-hierarchy#key_versions)
	// with which content of the PhraseSet is encrypted. The expected format is
	// `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}/cryptoKeyVersions/{crypto_key_version}`.
	KmsKeyVersionName string `protobuf:"bytes,8,opt,name=kms_key_version_name,json=kmsKeyVersionName,proto3" json:"kms_key_version_name,omitempty"`
	// Output only. System-assigned unique identifier for the PhraseSet.
	// This field is not used.
	Uid string `protobuf:"bytes,9,opt,name=uid,proto3" json:"uid,omitempty"`
	// Output only. User-settable, human-readable name for the PhraseSet. Must be
	// 63 characters or less. This field is not used.
	DisplayName string `protobuf:"bytes,10,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Output only. The CustomClass lifecycle state.
	// This field is not used.
	State PhraseSet_State `protobuf:"varint,11,opt,name=state,proto3,enum=google.events.cloud.speech.v1.PhraseSet_State" json:"state,omitempty"`
	// Output only. The time at which this resource was requested for deletion.
	// This field is not used.
	DeleteTime *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// Output only. The time at which this resource will be purged.
	// This field is not used.
	ExpireTime *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	// Output only. Allows users to store small amounts of arbitrary data.
	// Both the key and the value must be 63 characters or less each.
	// At most 100 annotations.
	// This field is not used.
	Annotations map[string]string `protobuf:"bytes,14,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Output only. This checksum is computed by the server based on the value of
	// other fields. This may be sent on update, undelete, and delete requests to
	// ensure the client has an up-to-date value before proceeding. This field is
	// not used.
	Etag string `protobuf:"bytes,15,opt,name=etag,proto3" json:"etag,omitempty"`
	// Output only. Whether or not this PhraseSet is in the process of being
	// updated. This field is not used.
	Reconciling   bool `protobuf:"varint,16,opt,name=reconciling,proto3" json:"reconciling,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PhraseSet) Reset() {
	*x = PhraseSet{}
	mi := &file_cloud_speech_v1_data_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PhraseSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhraseSet) ProtoMessage() {}

func (x *PhraseSet) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_speech_v1_data_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhraseSet.ProtoReflect.Descriptor instead.
func (*PhraseSet) Descriptor() ([]byte, []int) {
	return file_cloud_speech_v1_data_proto_rawDescGZIP(), []int{1}
}

func (x *PhraseSet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PhraseSet) GetPhrases() []*PhraseSet_Phrase {
	if x != nil {
		return x.Phrases
	}
	return nil
}

func (x *PhraseSet) GetBoost() float32 {
	if x != nil {
		return x.Boost
	}
	return 0
}

func (x *PhraseSet) GetKmsKeyName() string {
	if x != nil {
		return x.KmsKeyName
	}
	return ""
}

func (x *PhraseSet) GetKmsKeyVersionName() string {
	if x != nil {
		return x.KmsKeyVersionName
	}
	return ""
}

func (x *PhraseSet) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *PhraseSet) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *PhraseSet) GetState() PhraseSet_State {
	if x != nil {
		return x.State
	}
	return PhraseSet_STATE_UNSPECIFIED
}

func (x *PhraseSet) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

func (x *PhraseSet) GetExpireTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireTime
	}
	return nil
}

func (x *PhraseSet) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *PhraseSet) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *PhraseSet) GetReconciling() bool {
	if x != nil {
		return x.Reconciling
	}
	return false
}

// The data within all PhraseSet events.
type PhraseSetEventData struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Optional. The PhraseSet event payload. Unset for deletion events.
	Payload       *PhraseSet `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PhraseSetEventData) Reset() {
	*x = PhraseSetEventData{}
	mi := &file_cloud_speech_v1_data_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PhraseSetEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhraseSetEventData) ProtoMessage() {}

func (x *PhraseSetEventData) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_speech_v1_data_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhraseSetEventData.ProtoReflect.Descriptor instead.
func (*PhraseSetEventData) Descriptor() ([]byte, []int) {
	return file_cloud_speech_v1_data_proto_rawDescGZIP(), []int{2}
}

func (x *PhraseSetEventData) GetPayload() *PhraseSet {
	if x != nil {
		return x.Payload
	}
	return nil
}

// The data within all CustomClass events.
type CustomClassEventData struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Optional. The CustomClass event payload. Unset for deletion events.
	Payload       *CustomClass `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomClassEventData) Reset() {
	*x = CustomClassEventData{}
	mi := &file_cloud_speech_v1_data_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomClassEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomClassEventData) ProtoMessage() {}

func (x *CustomClassEventData) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_speech_v1_data_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomClassEventData.ProtoReflect.Descriptor instead.
func (*CustomClassEventData) Descriptor() ([]byte, []int) {
	return file_cloud_speech_v1_data_proto_rawDescGZIP(), []int{3}
}

func (x *CustomClassEventData) GetPayload() *CustomClass {
	if x != nil {
		return x.Payload
	}
	return nil
}

// An item of the class.
type CustomClass_ClassItem struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The class item's value.
	Value         string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomClass_ClassItem) Reset() {
	*x = CustomClass_ClassItem{}
	mi := &file_cloud_speech_v1_data_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomClass_ClassItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomClass_ClassItem) ProtoMessage() {}

func (x *CustomClass_ClassItem) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_speech_v1_data_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomClass_ClassItem.ProtoReflect.Descriptor instead.
func (*CustomClass_ClassItem) Descriptor() ([]byte, []int) {
	return file_cloud_speech_v1_data_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CustomClass_ClassItem) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// A phrases containing words and phrase "hints" so that
// the speech recognition is more likely to recognize them. This can be used
// to improve the accuracy for specific words and phrases, for example, if
// specific commands are typically spoken by the user. This can also be used
// to add additional words to the vocabulary of the recognizer. See
// [usage limits](https://cloud.google.com/speech-to-text/quotas#content).
//
// List items can also include pre-built or custom classes containing groups
// of words that represent common concepts that occur in natural language. For
// example, rather than providing a phrase hint for every month of the
// year (e.g. "i was born in january", "i was born in febuary", ...), use the
// pre-built `$MONTH` class improves the likelihood of correctly transcribing
// audio that includes months (e.g. "i was born in $month").
// To refer to pre-built classes, use the class' symbol prepended with `$`
// e.g. `$MONTH`. To refer to custom classes that were defined inline in the
// request, set the class's `custom_class_id` to a string unique to all class
// resources and inline classes. Then use the class' id wrapped in $`{...}`
// e.g. "${my-months}". To refer to custom classes resources, use the class'
// id wrapped in `${}` (e.g. `${my-months}`).
//
// Speech-to-Text supports three locations: `global`, `us` (US North America),
// and `eu` (Europe). If you are calling the `speech.googleapis.com`
// endpoint, use the `global` location. To specify a region, use a
// [regional endpoint](https://cloud.google.com/speech-to-text/docs/endpoints)
// with matching `us` or `eu` location value.
type PhraseSet_Phrase struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The phrase itself.
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// Hint Boost. Overrides the boost set at the phrase set level.
	// Positive value will increase the probability that a specific phrase will
	// be recognized over other similar sounding phrases. The higher the boost,
	// the higher the chance of false positive recognition as well. Negative
	// boost will simply be ignored. Though `boost` can accept a wide range of
	// positive values, most use cases are best served
	// with values between 0 and 20. We recommend using a binary search approach
	// to finding the optimal value for your use case as well as adding
	// phrases both with and without boost to your requests.
	Boost         float32 `protobuf:"fixed32,2,opt,name=boost,proto3" json:"boost,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PhraseSet_Phrase) Reset() {
	*x = PhraseSet_Phrase{}
	mi := &file_cloud_speech_v1_data_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PhraseSet_Phrase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhraseSet_Phrase) ProtoMessage() {}

func (x *PhraseSet_Phrase) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_speech_v1_data_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhraseSet_Phrase.ProtoReflect.Descriptor instead.
func (*PhraseSet_Phrase) Descriptor() ([]byte, []int) {
	return file_cloud_speech_v1_data_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PhraseSet_Phrase) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *PhraseSet_Phrase) GetBoost() float32 {
	if x != nil {
		return x.Boost
	}
	return 0
}

var File_cloud_speech_v1_data_proto protoreflect.FileDescriptor

const file_cloud_speech_v1_data_proto_rawDesc = "" +
	"\n" +
	"\x1acloud/speech/v1/data.proto\x12\x1dgoogle.events.cloud.speech.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"\x90\x06\n" +
	"\vCustomClass\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12&\n" +
	"\x0fcustom_class_id\x18\x02 \x01(\tR\rcustomClassId\x12J\n" +
	"\x05items\x18\x03 \x03(\v24.google.events.cloud.speech.v1.CustomClass.ClassItemR\x05items\x12 \n" +
	"\fkms_key_name\x18\x06 \x01(\tR\n" +
	"kmsKeyName\x12/\n" +
	"\x14kms_key_version_name\x18\a \x01(\tR\x11kmsKeyVersionName\x12\x10\n" +
	"\x03uid\x18\b \x01(\tR\x03uid\x12!\n" +
	"\fdisplay_name\x18\t \x01(\tR\vdisplayName\x12F\n" +
	"\x05state\x18\n" +
	" \x01(\x0e20.google.events.cloud.speech.v1.CustomClass.StateR\x05state\x12;\n" +
	"\vdelete_time\x18\v \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"deleteTime\x12;\n" +
	"\vexpire_time\x18\f \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"expireTime\x12]\n" +
	"\vannotations\x18\r \x03(\v2;.google.events.cloud.speech.v1.CustomClass.AnnotationsEntryR\vannotations\x12\x12\n" +
	"\x04etag\x18\x0e \x01(\tR\x04etag\x12 \n" +
	"\vreconciling\x18\x0f \x01(\bR\vreconciling\x1a!\n" +
	"\tClassItem\x12\x14\n" +
	"\x05value\x18\x01 \x01(\tR\x05value\x1a>\n" +
	"\x10AnnotationsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"7\n" +
	"\x05State\x12\x15\n" +
	"\x11STATE_UNSPECIFIED\x10\x00\x12\n" +
	"\n" +
	"\x06ACTIVE\x10\x02\x12\v\n" +
	"\aDELETED\x10\x04\"\x8a\x06\n" +
	"\tPhraseSet\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12I\n" +
	"\aphrases\x18\x02 \x03(\v2/.google.events.cloud.speech.v1.PhraseSet.PhraseR\aphrases\x12\x14\n" +
	"\x05boost\x18\x04 \x01(\x02R\x05boost\x12 \n" +
	"\fkms_key_name\x18\a \x01(\tR\n" +
	"kmsKeyName\x12/\n" +
	"\x14kms_key_version_name\x18\b \x01(\tR\x11kmsKeyVersionName\x12\x10\n" +
	"\x03uid\x18\t \x01(\tR\x03uid\x12!\n" +
	"\fdisplay_name\x18\n" +
	" \x01(\tR\vdisplayName\x12D\n" +
	"\x05state\x18\v \x01(\x0e2..google.events.cloud.speech.v1.PhraseSet.StateR\x05state\x12;\n" +
	"\vdelete_time\x18\f \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"deleteTime\x12;\n" +
	"\vexpire_time\x18\r \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"expireTime\x12[\n" +
	"\vannotations\x18\x0e \x03(\v29.google.events.cloud.speech.v1.PhraseSet.AnnotationsEntryR\vannotations\x12\x12\n" +
	"\x04etag\x18\x0f \x01(\tR\x04etag\x12 \n" +
	"\vreconciling\x18\x10 \x01(\bR\vreconciling\x1a4\n" +
	"\x06Phrase\x12\x14\n" +
	"\x05value\x18\x01 \x01(\tR\x05value\x12\x14\n" +
	"\x05boost\x18\x02 \x01(\x02R\x05boost\x1a>\n" +
	"\x10AnnotationsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"7\n" +
	"\x05State\x12\x15\n" +
	"\x11STATE_UNSPECIFIED\x10\x00\x12\n" +
	"\n" +
	"\x06ACTIVE\x10\x02\x12\v\n" +
	"\aDELETED\x10\x04\"X\n" +
	"\x12PhraseSetEventData\x12B\n" +
	"\apayload\x18\x01 \x01(\v2(.google.events.cloud.speech.v1.PhraseSetR\apayload\"\\\n" +
	"\x14CustomClassEventData\x12D\n" +
	"\apayload\x18\x01 \x01(\v2*.google.events.cloud.speech.v1.CustomClassR\apayloadBm\xaa\x02&Google.Events.Protobuf.Cloud.Speech.V1\xca\x02\x1dGoogle\\Events\\Cloud\\Speech\\V1\xea\x02!Google::Events::Cloud::Speech::V1b\x06proto3"

var (
	file_cloud_speech_v1_data_proto_rawDescOnce sync.Once
	file_cloud_speech_v1_data_proto_rawDescData []byte
)

func file_cloud_speech_v1_data_proto_rawDescGZIP() []byte {
	file_cloud_speech_v1_data_proto_rawDescOnce.Do(func() {
		file_cloud_speech_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_cloud_speech_v1_data_proto_rawDesc), len(file_cloud_speech_v1_data_proto_rawDesc)))
	})
	return file_cloud_speech_v1_data_proto_rawDescData
}

var file_cloud_speech_v1_data_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cloud_speech_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cloud_speech_v1_data_proto_goTypes = []any{
	(CustomClass_State)(0),        // 0: google.events.cloud.speech.v1.CustomClass.State
	(PhraseSet_State)(0),          // 1: google.events.cloud.speech.v1.PhraseSet.State
	(*CustomClass)(nil),           // 2: google.events.cloud.speech.v1.CustomClass
	(*PhraseSet)(nil),             // 3: google.events.cloud.speech.v1.PhraseSet
	(*PhraseSetEventData)(nil),    // 4: google.events.cloud.speech.v1.PhraseSetEventData
	(*CustomClassEventData)(nil),  // 5: google.events.cloud.speech.v1.CustomClassEventData
	(*CustomClass_ClassItem)(nil), // 6: google.events.cloud.speech.v1.CustomClass.ClassItem
	nil,                           // 7: google.events.cloud.speech.v1.CustomClass.AnnotationsEntry
	(*PhraseSet_Phrase)(nil),      // 8: google.events.cloud.speech.v1.PhraseSet.Phrase
	nil,                           // 9: google.events.cloud.speech.v1.PhraseSet.AnnotationsEntry
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_cloud_speech_v1_data_proto_depIdxs = []int32{
	6,  // 0: google.events.cloud.speech.v1.CustomClass.items:type_name -> google.events.cloud.speech.v1.CustomClass.ClassItem
	0,  // 1: google.events.cloud.speech.v1.CustomClass.state:type_name -> google.events.cloud.speech.v1.CustomClass.State
	10, // 2: google.events.cloud.speech.v1.CustomClass.delete_time:type_name -> google.protobuf.Timestamp
	10, // 3: google.events.cloud.speech.v1.CustomClass.expire_time:type_name -> google.protobuf.Timestamp
	7,  // 4: google.events.cloud.speech.v1.CustomClass.annotations:type_name -> google.events.cloud.speech.v1.CustomClass.AnnotationsEntry
	8,  // 5: google.events.cloud.speech.v1.PhraseSet.phrases:type_name -> google.events.cloud.speech.v1.PhraseSet.Phrase
	1,  // 6: google.events.cloud.speech.v1.PhraseSet.state:type_name -> google.events.cloud.speech.v1.PhraseSet.State
	10, // 7: google.events.cloud.speech.v1.PhraseSet.delete_time:type_name -> google.protobuf.Timestamp
	10, // 8: google.events.cloud.speech.v1.PhraseSet.expire_time:type_name -> google.protobuf.Timestamp
	9,  // 9: google.events.cloud.speech.v1.PhraseSet.annotations:type_name -> google.events.cloud.speech.v1.PhraseSet.AnnotationsEntry
	3,  // 10: google.events.cloud.speech.v1.PhraseSetEventData.payload:type_name -> google.events.cloud.speech.v1.PhraseSet
	2,  // 11: google.events.cloud.speech.v1.CustomClassEventData.payload:type_name -> google.events.cloud.speech.v1.CustomClass
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_cloud_speech_v1_data_proto_init() }
func file_cloud_speech_v1_data_proto_init() {
	if File_cloud_speech_v1_data_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cloud_speech_v1_data_proto_rawDesc), len(file_cloud_speech_v1_data_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_speech_v1_data_proto_goTypes,
		DependencyIndexes: file_cloud_speech_v1_data_proto_depIdxs,
		EnumInfos:         file_cloud_speech_v1_data_proto_enumTypes,
		MessageInfos:      file_cloud_speech_v1_data_proto_msgTypes,
	}.Build()
	File_cloud_speech_v1_data_proto = out.File
	file_cloud_speech_v1_data_proto_goTypes = nil
	file_cloud_speech_v1_data_proto_depIdxs = nil
}
