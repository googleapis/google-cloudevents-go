// Copyright 2023 Google LLC
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
// source: cloud/datastore/v1/data.proto

package datastoredata

import (
	latlng "google.golang.org/genproto/googleapis/type/latlng"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

// Specifies what data the 'entity' field contains.
// A `ResultType` is either implied (for example, in `LookupResponse.missing`
// from `datastore.proto`, it is always `KEY_ONLY`) or specified by context
// (for example, in message `QueryResultBatch`, field `entity_result_type`
// specifies a `ResultType` for all the values in field `entity_results`).
type EntityResult_ResultType int32

const (
	// Unspecified. This value is never used.
	EntityResult_RESULT_TYPE_UNSPECIFIED EntityResult_ResultType = 0
	// The key and properties.
	EntityResult_FULL EntityResult_ResultType = 1
	// A projected subset of properties. The entity may have no key.
	EntityResult_PROJECTION EntityResult_ResultType = 2
	// Only the key.
	EntityResult_KEY_ONLY EntityResult_ResultType = 3
)

// Enum value maps for EntityResult_ResultType.
var (
	EntityResult_ResultType_name = map[int32]string{
		0: "RESULT_TYPE_UNSPECIFIED",
		1: "FULL",
		2: "PROJECTION",
		3: "KEY_ONLY",
	}
	EntityResult_ResultType_value = map[string]int32{
		"RESULT_TYPE_UNSPECIFIED": 0,
		"FULL":                    1,
		"PROJECTION":              2,
		"KEY_ONLY":                3,
	}
)

func (x EntityResult_ResultType) Enum() *EntityResult_ResultType {
	p := new(EntityResult_ResultType)
	*p = x
	return p
}

func (x EntityResult_ResultType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntityResult_ResultType) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_datastore_v1_data_proto_enumTypes[0].Descriptor()
}

func (EntityResult_ResultType) Type() protoreflect.EnumType {
	return &file_cloud_datastore_v1_data_proto_enumTypes[0]
}

func (x EntityResult_ResultType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntityResult_ResultType.Descriptor instead.
func (EntityResult_ResultType) EnumDescriptor() ([]byte, []int) {
	return file_cloud_datastore_v1_data_proto_rawDescGZIP(), []int{2, 0}
}

// The data within all Firestore in Datastore Mode entity events.
type EntityEventData struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// An EntityResult object containing a post-operation entity snapshot.
	// This is not populated for delete events.
	Value *EntityResult `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// An EntityResult object containing a pre-operation entity snapshot.
	// This is only populated for update and delete events.
	OldValue *EntityResult `protobuf:"bytes,2,opt,name=old_value,json=oldValue,proto3" json:"old_value,omitempty"`
	// A PropertyMask object that lists changed properties.
	// This is only populated for update events..
	UpdateMask    *PropertyMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EntityEventData) Reset() {
	*x = EntityEventData{}
	mi := &file_cloud_datastore_v1_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntityEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityEventData) ProtoMessage() {}

func (x *EntityEventData) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_datastore_v1_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityEventData.ProtoReflect.Descriptor instead.
func (*EntityEventData) Descriptor() ([]byte, []int) {
	return file_cloud_datastore_v1_data_proto_rawDescGZIP(), []int{0}
}

func (x *EntityEventData) GetValue() *EntityResult {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *EntityEventData) GetOldValue() *EntityResult {
	if x != nil {
		return x.OldValue
	}
	return nil
}

func (x *EntityEventData) GetUpdateMask() *PropertyMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// A set of property paths on an entity.
type PropertyMask struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The list of property paths in the mask.
	// This is not populated for delete events.
	PropertyPaths []string `protobuf:"bytes,1,rep,name=property_paths,json=propertyPaths,proto3" json:"property_paths,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PropertyMask) Reset() {
	*x = PropertyMask{}
	mi := &file_cloud_datastore_v1_data_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PropertyMask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropertyMask) ProtoMessage() {}

func (x *PropertyMask) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_datastore_v1_data_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropertyMask.ProtoReflect.Descriptor instead.
func (*PropertyMask) Descriptor() ([]byte, []int) {
	return file_cloud_datastore_v1_data_proto_rawDescGZIP(), []int{1}
}

func (x *PropertyMask) GetPropertyPaths() []string {
	if x != nil {
		return x.PropertyPaths
	}
	return nil
}

// The result of fetching an entity from Datastore.
type EntityResult struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The resulting entity.
	Entity *Entity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	// The version of the entity, a strictly positive number that monotonically
	// increases with changes to the entity.
	//
	// This field is set for
	// [`FULL`][google.datastore.v1.EntityResult.ResultType.FULL] entity results.
	//
	// For [missing][google.datastore.v1.LookupResponse.missing] entities in
	// `LookupResponse`, this is the version of the snapshot that was used to look
	// up the entity, and it is always set except for eventually consistent reads.
	Version int64 `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	// The time at which the entity was created.
	// This field is set for
	// [`FULL`][google.datastore.v1.EntityResult.ResultType.FULL] entity results.
	// If this entity is missing, this field will not be set.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The time at which the entity was last changed.
	// This field is set for
	// [`FULL`][google.datastore.v1.EntityResult.ResultType.FULL] entity results.
	// If this entity is missing, this field will not be set.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// A cursor that points to the position after the result entity.
	// Set only when the `EntityResult` is part of a `QueryResultBatch` message.
	Cursor        []byte `protobuf:"bytes,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EntityResult) Reset() {
	*x = EntityResult{}
	mi := &file_cloud_datastore_v1_data_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntityResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityResult) ProtoMessage() {}

func (x *EntityResult) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_datastore_v1_data_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityResult.ProtoReflect.Descriptor instead.
func (*EntityResult) Descriptor() ([]byte, []int) {
	return file_cloud_datastore_v1_data_proto_rawDescGZIP(), []int{2}
}

func (x *EntityResult) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *EntityResult) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *EntityResult) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *EntityResult) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *EntityResult) GetCursor() []byte {
	if x != nil {
		return x.Cursor
	}
	return nil
}

// A partition ID identifies a grouping of entities. The grouping is always
// by project and namespace, however the namespace ID may be empty.
//
// A partition ID contains several dimensions:
// project ID and namespace ID.
//
// Partition dimensions:
//
// - May be `""`.
// - Must be valid UTF-8 bytes.
// - Must have values that match regex `[A-Za-z\d\.\-_]{1,100}`
// If the value of any dimension matches regex `__.*__`, the partition is
// reserved/read-only.
// A reserved/read-only partition ID is forbidden in certain documented
// contexts.
//
// Foreign partition IDs (in which the project ID does
// not match the context project ID ) are discouraged.
// Reads and writes of foreign partition IDs may fail if the project is not in
// an active state.
type PartitionId struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The ID of the project to which the entities belong.
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// If not empty, the ID of the database to which the entities
	// belong.
	DatabaseId string `protobuf:"bytes,3,opt,name=database_id,json=databaseId,proto3" json:"database_id,omitempty"`
	// If not empty, the ID of the namespace to which the entities belong.
	NamespaceId   string `protobuf:"bytes,4,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PartitionId) Reset() {
	*x = PartitionId{}
	mi := &file_cloud_datastore_v1_data_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PartitionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionId) ProtoMessage() {}

func (x *PartitionId) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_datastore_v1_data_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartitionId.ProtoReflect.Descriptor instead.
func (*PartitionId) Descriptor() ([]byte, []int) {
	return file_cloud_datastore_v1_data_proto_rawDescGZIP(), []int{3}
}

func (x *PartitionId) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *PartitionId) GetDatabaseId() string {
	if x != nil {
		return x.DatabaseId
	}
	return ""
}

func (x *PartitionId) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

// A unique identifier for an entity.
// If a key's partition ID or any of its path kinds or names are
// reserved/read-only, the key is reserved/read-only.
// A reserved/read-only key is forbidden in certain documented contexts.
type Key struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Entities are partitioned into subsets, currently identified by a project
	// ID and namespace ID.
	// Queries are scoped to a single partition.
	PartitionId *PartitionId `protobuf:"bytes,1,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	// The entity path.
	// An entity path consists of one or more elements composed of a kind and a
	// string or numerical identifier, which identify entities. The first
	// element identifies a _root entity_, the second element identifies
	// a _child_ of the root entity, the third element identifies a child of the
	// second entity, and so forth. The entities identified by all prefixes of
	// the path are called the element's _ancestors_.
	//
	// An entity path is always fully complete: *all* of the entity's ancestors
	// are required to be in the path along with the entity identifier itself.
	// The only exception is that in some documented cases, the identifier in the
	// last path element (for the entity) itself may be omitted. For example,
	// the last path element of the key of `Mutation.insert` may have no
	// identifier.
	//
	// A path can never be empty, and a path can have at most 100 elements.
	Path          []*Key_PathElement `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Key) Reset() {
	*x = Key{}
	mi := &file_cloud_datastore_v1_data_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_datastore_v1_data_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_cloud_datastore_v1_data_proto_rawDescGZIP(), []int{4}
}

func (x *Key) GetPartitionId() *PartitionId {
	if x != nil {
		return x.PartitionId
	}
	return nil
}

func (x *Key) GetPath() []*Key_PathElement {
	if x != nil {
		return x.Path
	}
	return nil
}

// An array value.
type ArrayValue struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Values in the array.
	// The order of values in an array is preserved as long as all values have
	// identical settings for 'exclude_from_indexes'.
	Values        []*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ArrayValue) Reset() {
	*x = ArrayValue{}
	mi := &file_cloud_datastore_v1_data_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ArrayValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrayValue) ProtoMessage() {}

func (x *ArrayValue) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_datastore_v1_data_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrayValue.ProtoReflect.Descriptor instead.
func (*ArrayValue) Descriptor() ([]byte, []int) {
	return file_cloud_datastore_v1_data_proto_rawDescGZIP(), []int{5}
}

func (x *ArrayValue) GetValues() []*Value {
	if x != nil {
		return x.Values
	}
	return nil
}

// A message that can hold any of the supported value types and associated
// metadata.
type Value struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Must have a value set.
	//
	// Types that are valid to be assigned to ValueType:
	//
	//	*Value_NullValue
	//	*Value_BooleanValue
	//	*Value_IntegerValue
	//	*Value_DoubleValue
	//	*Value_TimestampValue
	//	*Value_KeyValue
	//	*Value_StringValue
	//	*Value_BlobValue
	//	*Value_GeoPointValue
	//	*Value_EntityValue
	//	*Value_ArrayValue
	ValueType isValue_ValueType `protobuf_oneof:"value_type"`
	// The `meaning` field should only be populated for backwards compatibility.
	Meaning int32 `protobuf:"varint,14,opt,name=meaning,proto3" json:"meaning,omitempty"`
	// If the value should be excluded from all indexes including those defined
	// explicitly.
	ExcludeFromIndexes bool `protobuf:"varint,19,opt,name=exclude_from_indexes,json=excludeFromIndexes,proto3" json:"exclude_from_indexes,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *Value) Reset() {
	*x = Value{}
	mi := &file_cloud_datastore_v1_data_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_datastore_v1_data_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_cloud_datastore_v1_data_proto_rawDescGZIP(), []int{6}
}

func (x *Value) GetValueType() isValue_ValueType {
	if x != nil {
		return x.ValueType
	}
	return nil
}

func (x *Value) GetNullValue() structpb.NullValue {
	if x != nil {
		if x, ok := x.ValueType.(*Value_NullValue); ok {
			return x.NullValue
		}
	}
	return structpb.NullValue(0)
}

func (x *Value) GetBooleanValue() bool {
	if x != nil {
		if x, ok := x.ValueType.(*Value_BooleanValue); ok {
			return x.BooleanValue
		}
	}
	return false
}

func (x *Value) GetIntegerValue() int64 {
	if x != nil {
		if x, ok := x.ValueType.(*Value_IntegerValue); ok {
			return x.IntegerValue
		}
	}
	return 0
}

func (x *Value) GetDoubleValue() float64 {
	if x != nil {
		if x, ok := x.ValueType.(*Value_DoubleValue); ok {
			return x.DoubleValue
		}
	}
	return 0
}

func (x *Value) GetTimestampValue() *timestamppb.Timestamp {
	if x != nil {
		if x, ok := x.ValueType.(*Value_TimestampValue); ok {
			return x.TimestampValue
		}
	}
	return nil
}

func (x *Value) GetKeyValue() *Key {
	if x != nil {
		if x, ok := x.ValueType.(*Value_KeyValue); ok {
			return x.KeyValue
		}
	}
	return nil
}

func (x *Value) GetStringValue() string {
	if x != nil {
		if x, ok := x.ValueType.(*Value_StringValue); ok {
			return x.StringValue
		}
	}
	return ""
}

func (x *Value) GetBlobValue() []byte {
	if x != nil {
		if x, ok := x.ValueType.(*Value_BlobValue); ok {
			return x.BlobValue
		}
	}
	return nil
}

func (x *Value) GetGeoPointValue() *latlng.LatLng {
	if x != nil {
		if x, ok := x.ValueType.(*Value_GeoPointValue); ok {
			return x.GeoPointValue
		}
	}
	return nil
}

func (x *Value) GetEntityValue() *Entity {
	if x != nil {
		if x, ok := x.ValueType.(*Value_EntityValue); ok {
			return x.EntityValue
		}
	}
	return nil
}

func (x *Value) GetArrayValue() *ArrayValue {
	if x != nil {
		if x, ok := x.ValueType.(*Value_ArrayValue); ok {
			return x.ArrayValue
		}
	}
	return nil
}

func (x *Value) GetMeaning() int32 {
	if x != nil {
		return x.Meaning
	}
	return 0
}

func (x *Value) GetExcludeFromIndexes() bool {
	if x != nil {
		return x.ExcludeFromIndexes
	}
	return false
}

type isValue_ValueType interface {
	isValue_ValueType()
}

type Value_NullValue struct {
	// A null value.
	NullValue structpb.NullValue `protobuf:"varint,11,opt,name=null_value,json=nullValue,proto3,enum=google.protobuf.NullValue,oneof"`
}

type Value_BooleanValue struct {
	// A boolean value.
	BooleanValue bool `protobuf:"varint,1,opt,name=boolean_value,json=booleanValue,proto3,oneof"`
}

type Value_IntegerValue struct {
	// An integer value.
	IntegerValue int64 `protobuf:"varint,2,opt,name=integer_value,json=integerValue,proto3,oneof"`
}

type Value_DoubleValue struct {
	// A double value.
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Value_TimestampValue struct {
	// A timestamp value.
	// When stored in the Datastore, precise only to microseconds;
	// any additional precision is rounded down.
	TimestampValue *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=timestamp_value,json=timestampValue,proto3,oneof"`
}

type Value_KeyValue struct {
	// A key value.
	KeyValue *Key `protobuf:"bytes,5,opt,name=key_value,json=keyValue,proto3,oneof"`
}

type Value_StringValue struct {
	// A UTF-8 encoded string value.
	// When `exclude_from_indexes` is false (it is indexed) , may have at most
	// 1500 bytes. Otherwise, may be set to at most 1,000,000 bytes.
	StringValue string `protobuf:"bytes,17,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Value_BlobValue struct {
	// A blob value.
	// May have at most 1,000,000 bytes.
	// When `exclude_from_indexes` is false, may have at most 1500 bytes.
	// In JSON requests, must be base64-encoded.
	BlobValue []byte `protobuf:"bytes,18,opt,name=blob_value,json=blobValue,proto3,oneof"`
}

type Value_GeoPointValue struct {
	// A geo point value representing a point on the surface of Earth.
	GeoPointValue *latlng.LatLng `protobuf:"bytes,8,opt,name=geo_point_value,json=geoPointValue,proto3,oneof"`
}

type Value_EntityValue struct {
	// An entity value.
	//
	// - May have no key.
	// - May have a key with an incomplete key path.
	// - May have a reserved/read-only key.
	EntityValue *Entity `protobuf:"bytes,6,opt,name=entity_value,json=entityValue,proto3,oneof"`
}

type Value_ArrayValue struct {
	// An array value.
	// Cannot contain another array value.
	// A `Value` instance that sets field `array_value` must not set fields
	// `meaning` or `exclude_from_indexes`.
	ArrayValue *ArrayValue `protobuf:"bytes,9,opt,name=array_value,json=arrayValue,proto3,oneof"`
}

func (*Value_NullValue) isValue_ValueType() {}

func (*Value_BooleanValue) isValue_ValueType() {}

func (*Value_IntegerValue) isValue_ValueType() {}

func (*Value_DoubleValue) isValue_ValueType() {}

func (*Value_TimestampValue) isValue_ValueType() {}

func (*Value_KeyValue) isValue_ValueType() {}

func (*Value_StringValue) isValue_ValueType() {}

func (*Value_BlobValue) isValue_ValueType() {}

func (*Value_GeoPointValue) isValue_ValueType() {}

func (*Value_EntityValue) isValue_ValueType() {}

func (*Value_ArrayValue) isValue_ValueType() {}

// A Datastore data object.
//
// Must not exceed 1 MiB - 4 bytes.
type Entity struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The entity's key.
	//
	// An entity must have a key, unless otherwise documented (for example,
	// an entity in `Value.entity_value` may have no key).
	// An entity's kind is its key path's last element's kind,
	// or null if it has no key.
	Key *Key `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The entity's properties.
	// The map's keys are property names.
	// A property name matching regex `__.*__` is reserved.
	// A reserved property name is forbidden in certain documented contexts.
	// The map keys, represented as UTF-8, must not exceed 1,500 bytes and cannot
	// be empty.
	Properties    map[string]*Value `protobuf:"bytes,3,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Entity) Reset() {
	*x = Entity{}
	mi := &file_cloud_datastore_v1_data_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_datastore_v1_data_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_cloud_datastore_v1_data_proto_rawDescGZIP(), []int{7}
}

func (x *Entity) GetKey() *Key {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Entity) GetProperties() map[string]*Value {
	if x != nil {
		return x.Properties
	}
	return nil
}

// A (kind, ID/name) pair used to construct a key path.
//
// If either name or ID is set, the element is complete.
// If neither is set, the element is incomplete.
type Key_PathElement struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The kind of the entity.
	//
	// A kind matching regex `__.*__` is reserved/read-only.
	// A kind must not contain more than 1500 bytes when UTF-8 encoded.
	// Cannot be `""`.
	//
	// Must be valid UTF-8 bytes. Legacy values that are not valid UTF-8 are
	// encoded as `__bytes<X>__` where `<X>` is the base-64 encoding of the
	// bytes.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// The type of ID.
	//
	// Types that are valid to be assigned to IdType:
	//
	//	*Key_PathElement_Id
	//	*Key_PathElement_Name
	IdType        isKey_PathElement_IdType `protobuf_oneof:"id_type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Key_PathElement) Reset() {
	*x = Key_PathElement{}
	mi := &file_cloud_datastore_v1_data_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Key_PathElement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key_PathElement) ProtoMessage() {}

func (x *Key_PathElement) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_datastore_v1_data_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key_PathElement.ProtoReflect.Descriptor instead.
func (*Key_PathElement) Descriptor() ([]byte, []int) {
	return file_cloud_datastore_v1_data_proto_rawDescGZIP(), []int{4, 0}
}

func (x *Key_PathElement) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Key_PathElement) GetIdType() isKey_PathElement_IdType {
	if x != nil {
		return x.IdType
	}
	return nil
}

func (x *Key_PathElement) GetId() int64 {
	if x != nil {
		if x, ok := x.IdType.(*Key_PathElement_Id); ok {
			return x.Id
		}
	}
	return 0
}

func (x *Key_PathElement) GetName() string {
	if x != nil {
		if x, ok := x.IdType.(*Key_PathElement_Name); ok {
			return x.Name
		}
	}
	return ""
}

type isKey_PathElement_IdType interface {
	isKey_PathElement_IdType()
}

type Key_PathElement_Id struct {
	// The auto-allocated ID of the entity.
	//
	// Never equal to zero. Values less than zero are discouraged and may not
	// be supported in the future.
	Id int64 `protobuf:"varint,2,opt,name=id,proto3,oneof"`
}

type Key_PathElement_Name struct {
	// The name of the entity.
	//
	// A name matching regex `__.*__` is reserved/read-only.
	// A name must not be more than 1500 bytes when UTF-8 encoded.
	// Cannot be `""`.
	//
	// Must be valid UTF-8 bytes. Legacy values that are not valid UTF-8 are
	// encoded as `__bytes<X>__` where `<X>` is the base-64 encoding of the
	// bytes.
	Name string `protobuf:"bytes,3,opt,name=name,proto3,oneof"`
}

func (*Key_PathElement_Id) isKey_PathElement_IdType() {}

func (*Key_PathElement_Name) isKey_PathElement_IdType() {}

var File_cloud_datastore_v1_data_proto protoreflect.FileDescriptor

const file_cloud_datastore_v1_data_proto_rawDesc = "" +
	"\n" +
	"\x1dcloud/datastore/v1/data.proto\x12 google.events.cloud.datastore.v1\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x18google/type/latlng.proto\"\xf5\x01\n" +
	"\x0fEntityEventData\x12D\n" +
	"\x05value\x18\x01 \x01(\v2..google.events.cloud.datastore.v1.EntityResultR\x05value\x12K\n" +
	"\told_value\x18\x02 \x01(\v2..google.events.cloud.datastore.v1.EntityResultR\boldValue\x12O\n" +
	"\vupdate_mask\x18\x03 \x01(\v2..google.events.cloud.datastore.v1.PropertyMaskR\n" +
	"updateMask\"5\n" +
	"\fPropertyMask\x12%\n" +
	"\x0eproperty_paths\x18\x01 \x03(\tR\rpropertyPaths\"\xcf\x02\n" +
	"\fEntityResult\x12@\n" +
	"\x06entity\x18\x01 \x01(\v2(.google.events.cloud.datastore.v1.EntityR\x06entity\x12\x18\n" +
	"\aversion\x18\x04 \x01(\x03R\aversion\x12;\n" +
	"\vcreate_time\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"createTime\x12;\n" +
	"\vupdate_time\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"updateTime\x12\x16\n" +
	"\x06cursor\x18\x03 \x01(\fR\x06cursor\"Q\n" +
	"\n" +
	"ResultType\x12\x1b\n" +
	"\x17RESULT_TYPE_UNSPECIFIED\x10\x00\x12\b\n" +
	"\x04FULL\x10\x01\x12\x0e\n" +
	"\n" +
	"PROJECTION\x10\x02\x12\f\n" +
	"\bKEY_ONLY\x10\x03\"p\n" +
	"\vPartitionId\x12\x1d\n" +
	"\n" +
	"project_id\x18\x02 \x01(\tR\tprojectId\x12\x1f\n" +
	"\vdatabase_id\x18\x03 \x01(\tR\n" +
	"databaseId\x12!\n" +
	"\fnamespace_id\x18\x04 \x01(\tR\vnamespaceId\"\xf4\x01\n" +
	"\x03Key\x12P\n" +
	"\fpartition_id\x18\x01 \x01(\v2-.google.events.cloud.datastore.v1.PartitionIdR\vpartitionId\x12E\n" +
	"\x04path\x18\x02 \x03(\v21.google.events.cloud.datastore.v1.Key.PathElementR\x04path\x1aT\n" +
	"\vPathElement\x12\x12\n" +
	"\x04kind\x18\x01 \x01(\tR\x04kind\x12\x10\n" +
	"\x02id\x18\x02 \x01(\x03H\x00R\x02id\x12\x14\n" +
	"\x04name\x18\x03 \x01(\tH\x00R\x04nameB\t\n" +
	"\aid_type\"M\n" +
	"\n" +
	"ArrayValue\x12?\n" +
	"\x06values\x18\x01 \x03(\v2'.google.events.cloud.datastore.v1.ValueR\x06values\"\xc3\x05\n" +
	"\x05Value\x12;\n" +
	"\n" +
	"null_value\x18\v \x01(\x0e2\x1a.google.protobuf.NullValueH\x00R\tnullValue\x12%\n" +
	"\rboolean_value\x18\x01 \x01(\bH\x00R\fbooleanValue\x12%\n" +
	"\rinteger_value\x18\x02 \x01(\x03H\x00R\fintegerValue\x12#\n" +
	"\fdouble_value\x18\x03 \x01(\x01H\x00R\vdoubleValue\x12E\n" +
	"\x0ftimestamp_value\x18\n" +
	" \x01(\v2\x1a.google.protobuf.TimestampH\x00R\x0etimestampValue\x12D\n" +
	"\tkey_value\x18\x05 \x01(\v2%.google.events.cloud.datastore.v1.KeyH\x00R\bkeyValue\x12#\n" +
	"\fstring_value\x18\x11 \x01(\tH\x00R\vstringValue\x12\x1f\n" +
	"\n" +
	"blob_value\x18\x12 \x01(\fH\x00R\tblobValue\x12=\n" +
	"\x0fgeo_point_value\x18\b \x01(\v2\x13.google.type.LatLngH\x00R\rgeoPointValue\x12M\n" +
	"\fentity_value\x18\x06 \x01(\v2(.google.events.cloud.datastore.v1.EntityH\x00R\ventityValue\x12O\n" +
	"\varray_value\x18\t \x01(\v2,.google.events.cloud.datastore.v1.ArrayValueH\x00R\n" +
	"arrayValue\x12\x18\n" +
	"\ameaning\x18\x0e \x01(\x05R\ameaning\x120\n" +
	"\x14exclude_from_indexes\x18\x13 \x01(\bR\x12excludeFromIndexesB\f\n" +
	"\n" +
	"value_type\"\x83\x02\n" +
	"\x06Entity\x127\n" +
	"\x03key\x18\x01 \x01(\v2%.google.events.cloud.datastore.v1.KeyR\x03key\x12X\n" +
	"\n" +
	"properties\x18\x03 \x03(\v28.google.events.cloud.datastore.v1.Entity.PropertiesEntryR\n" +
	"properties\x1af\n" +
	"\x0fPropertiesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12=\n" +
	"\x05value\x18\x02 \x01(\v2'.google.events.cloud.datastore.v1.ValueR\x05value:\x028\x01BxP\x01\xaa\x02)Google.Events.Protobuf.Cloud.Datastore.V1\xca\x02 Google\\Events\\Cloud\\Datastore\\V1\xea\x02$Google::Events::Cloud::Datastore::V1b\x06proto3"

var (
	file_cloud_datastore_v1_data_proto_rawDescOnce sync.Once
	file_cloud_datastore_v1_data_proto_rawDescData []byte
)

func file_cloud_datastore_v1_data_proto_rawDescGZIP() []byte {
	file_cloud_datastore_v1_data_proto_rawDescOnce.Do(func() {
		file_cloud_datastore_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_cloud_datastore_v1_data_proto_rawDesc), len(file_cloud_datastore_v1_data_proto_rawDesc)))
	})
	return file_cloud_datastore_v1_data_proto_rawDescData
}

var file_cloud_datastore_v1_data_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_datastore_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_cloud_datastore_v1_data_proto_goTypes = []any{
	(EntityResult_ResultType)(0),  // 0: google.events.cloud.datastore.v1.EntityResult.ResultType
	(*EntityEventData)(nil),       // 1: google.events.cloud.datastore.v1.EntityEventData
	(*PropertyMask)(nil),          // 2: google.events.cloud.datastore.v1.PropertyMask
	(*EntityResult)(nil),          // 3: google.events.cloud.datastore.v1.EntityResult
	(*PartitionId)(nil),           // 4: google.events.cloud.datastore.v1.PartitionId
	(*Key)(nil),                   // 5: google.events.cloud.datastore.v1.Key
	(*ArrayValue)(nil),            // 6: google.events.cloud.datastore.v1.ArrayValue
	(*Value)(nil),                 // 7: google.events.cloud.datastore.v1.Value
	(*Entity)(nil),                // 8: google.events.cloud.datastore.v1.Entity
	(*Key_PathElement)(nil),       // 9: google.events.cloud.datastore.v1.Key.PathElement
	nil,                           // 10: google.events.cloud.datastore.v1.Entity.PropertiesEntry
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
	(structpb.NullValue)(0),       // 12: google.protobuf.NullValue
	(*latlng.LatLng)(nil),         // 13: google.type.LatLng
}
var file_cloud_datastore_v1_data_proto_depIdxs = []int32{
	3,  // 0: google.events.cloud.datastore.v1.EntityEventData.value:type_name -> google.events.cloud.datastore.v1.EntityResult
	3,  // 1: google.events.cloud.datastore.v1.EntityEventData.old_value:type_name -> google.events.cloud.datastore.v1.EntityResult
	2,  // 2: google.events.cloud.datastore.v1.EntityEventData.update_mask:type_name -> google.events.cloud.datastore.v1.PropertyMask
	8,  // 3: google.events.cloud.datastore.v1.EntityResult.entity:type_name -> google.events.cloud.datastore.v1.Entity
	11, // 4: google.events.cloud.datastore.v1.EntityResult.create_time:type_name -> google.protobuf.Timestamp
	11, // 5: google.events.cloud.datastore.v1.EntityResult.update_time:type_name -> google.protobuf.Timestamp
	4,  // 6: google.events.cloud.datastore.v1.Key.partition_id:type_name -> google.events.cloud.datastore.v1.PartitionId
	9,  // 7: google.events.cloud.datastore.v1.Key.path:type_name -> google.events.cloud.datastore.v1.Key.PathElement
	7,  // 8: google.events.cloud.datastore.v1.ArrayValue.values:type_name -> google.events.cloud.datastore.v1.Value
	12, // 9: google.events.cloud.datastore.v1.Value.null_value:type_name -> google.protobuf.NullValue
	11, // 10: google.events.cloud.datastore.v1.Value.timestamp_value:type_name -> google.protobuf.Timestamp
	5,  // 11: google.events.cloud.datastore.v1.Value.key_value:type_name -> google.events.cloud.datastore.v1.Key
	13, // 12: google.events.cloud.datastore.v1.Value.geo_point_value:type_name -> google.type.LatLng
	8,  // 13: google.events.cloud.datastore.v1.Value.entity_value:type_name -> google.events.cloud.datastore.v1.Entity
	6,  // 14: google.events.cloud.datastore.v1.Value.array_value:type_name -> google.events.cloud.datastore.v1.ArrayValue
	5,  // 15: google.events.cloud.datastore.v1.Entity.key:type_name -> google.events.cloud.datastore.v1.Key
	10, // 16: google.events.cloud.datastore.v1.Entity.properties:type_name -> google.events.cloud.datastore.v1.Entity.PropertiesEntry
	7,  // 17: google.events.cloud.datastore.v1.Entity.PropertiesEntry.value:type_name -> google.events.cloud.datastore.v1.Value
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_cloud_datastore_v1_data_proto_init() }
func file_cloud_datastore_v1_data_proto_init() {
	if File_cloud_datastore_v1_data_proto != nil {
		return
	}
	file_cloud_datastore_v1_data_proto_msgTypes[6].OneofWrappers = []any{
		(*Value_NullValue)(nil),
		(*Value_BooleanValue)(nil),
		(*Value_IntegerValue)(nil),
		(*Value_DoubleValue)(nil),
		(*Value_TimestampValue)(nil),
		(*Value_KeyValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BlobValue)(nil),
		(*Value_GeoPointValue)(nil),
		(*Value_EntityValue)(nil),
		(*Value_ArrayValue)(nil),
	}
	file_cloud_datastore_v1_data_proto_msgTypes[8].OneofWrappers = []any{
		(*Key_PathElement_Id)(nil),
		(*Key_PathElement_Name)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cloud_datastore_v1_data_proto_rawDesc), len(file_cloud_datastore_v1_data_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_datastore_v1_data_proto_goTypes,
		DependencyIndexes: file_cloud_datastore_v1_data_proto_depIdxs,
		EnumInfos:         file_cloud_datastore_v1_data_proto_enumTypes,
		MessageInfos:      file_cloud_datastore_v1_data_proto_msgTypes,
	}.Build()
	File_cloud_datastore_v1_data_proto = out.File
	file_cloud_datastore_v1_data_proto_goTypes = nil
	file_cloud_datastore_v1_data_proto_depIdxs = nil
}
