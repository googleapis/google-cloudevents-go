// Copyright 2022 Google LLC
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
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.6
// source: cloud/beyondcorp/appgateways/v1/data.proto

package appgatewaysdata

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

// Enum containing list of all possible network connectivity options
// supported by BeyondCorp AppGateway.
type AppGateway_Type int32

const (
	// Default value. This value is unused.
	AppGateway_TYPE_UNSPECIFIED AppGateway_Type = 0
	// TCP Proxy based BeyondCorp Connection. API will default to this if unset.
	AppGateway_TCP_PROXY AppGateway_Type = 1
)

// Enum value maps for AppGateway_Type.
var (
	AppGateway_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TCP_PROXY",
	}
	AppGateway_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TCP_PROXY":        1,
	}
)

func (x AppGateway_Type) Enum() *AppGateway_Type {
	p := new(AppGateway_Type)
	*p = x
	return p
}

func (x AppGateway_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppGateway_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_beyondcorp_appgateways_v1_data_proto_enumTypes[0].Descriptor()
}

func (AppGateway_Type) Type() protoreflect.EnumType {
	return &file_cloud_beyondcorp_appgateways_v1_data_proto_enumTypes[0]
}

func (x AppGateway_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppGateway_Type.Descriptor instead.
func (AppGateway_Type) EnumDescriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_appgateways_v1_data_proto_rawDescGZIP(), []int{0, 0}
}

// Represents the different states of an AppGateway.
type AppGateway_State int32

const (
	// Default value. This value is unused.
	AppGateway_STATE_UNSPECIFIED AppGateway_State = 0
	// AppGateway is being created.
	AppGateway_CREATING AppGateway_State = 1
	// AppGateway has been created.
	AppGateway_CREATED AppGateway_State = 2
	// AppGateway's configuration is being updated.
	AppGateway_UPDATING AppGateway_State = 3
	// AppGateway is being deleted.
	AppGateway_DELETING AppGateway_State = 4
	// AppGateway is down and may be restored in the future.
	// This happens when CCFE sends ProjectState = OFF.
	AppGateway_DOWN AppGateway_State = 5
)

// Enum value maps for AppGateway_State.
var (
	AppGateway_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "CREATED",
		3: "UPDATING",
		4: "DELETING",
		5: "DOWN",
	}
	AppGateway_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          1,
		"CREATED":           2,
		"UPDATING":          3,
		"DELETING":          4,
		"DOWN":              5,
	}
)

func (x AppGateway_State) Enum() *AppGateway_State {
	p := new(AppGateway_State)
	*p = x
	return p
}

func (x AppGateway_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppGateway_State) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_beyondcorp_appgateways_v1_data_proto_enumTypes[1].Descriptor()
}

func (AppGateway_State) Type() protoreflect.EnumType {
	return &file_cloud_beyondcorp_appgateways_v1_data_proto_enumTypes[1]
}

func (x AppGateway_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppGateway_State.Descriptor instead.
func (AppGateway_State) EnumDescriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_appgateways_v1_data_proto_rawDescGZIP(), []int{0, 1}
}

// Enum containing list of all possible host types supported by BeyondCorp
// Connection.
type AppGateway_HostType int32

const (
	// Default value. This value is unused.
	AppGateway_HOST_TYPE_UNSPECIFIED AppGateway_HostType = 0
	// AppGateway hosted in a GCP regional managed instance group.
	AppGateway_GCP_REGIONAL_MIG AppGateway_HostType = 1
)

// Enum value maps for AppGateway_HostType.
var (
	AppGateway_HostType_name = map[int32]string{
		0: "HOST_TYPE_UNSPECIFIED",
		1: "GCP_REGIONAL_MIG",
	}
	AppGateway_HostType_value = map[string]int32{
		"HOST_TYPE_UNSPECIFIED": 0,
		"GCP_REGIONAL_MIG":      1,
	}
)

func (x AppGateway_HostType) Enum() *AppGateway_HostType {
	p := new(AppGateway_HostType)
	*p = x
	return p
}

func (x AppGateway_HostType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppGateway_HostType) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_beyondcorp_appgateways_v1_data_proto_enumTypes[2].Descriptor()
}

func (AppGateway_HostType) Type() protoreflect.EnumType {
	return &file_cloud_beyondcorp_appgateways_v1_data_proto_enumTypes[2]
}

func (x AppGateway_HostType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppGateway_HostType.Descriptor instead.
func (AppGateway_HostType) EnumDescriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_appgateways_v1_data_proto_rawDescGZIP(), []int{0, 2}
}

// A BeyondCorp AppGateway resource represents a BeyondCorp protected AppGateway
// to a remote application. It creates all the necessary GCP components needed
// for creating a BeyondCorp protected AppGateway. Multiple connectors can be
// authorised for a single AppGateway.
type AppGateway struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Unique resource name of the AppGateway.
	// The name is ignored when creating an AppGateway.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Timestamp when the resource was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp when the resource was last modified.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Optional. Resource labels to represent user provided metadata.
	Labels map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Optional. An arbitrary user-provided name for the AppGateway. Cannot exceed
	// 64 characters.
	DisplayName string `protobuf:"bytes,5,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Output only. A unique identifier for the instance generated by the
	// system.
	Uid string `protobuf:"bytes,6,opt,name=uid,proto3" json:"uid,omitempty"`
	// Required. The type of network connectivity used by the AppGateway.
	Type AppGateway_Type `protobuf:"varint,7,opt,name=type,proto3,enum=google.events.cloud.beyondcorp.appgateways.v1.AppGateway_Type" json:"type,omitempty"`
	// Output only. The current state of the AppGateway.
	State AppGateway_State `protobuf:"varint,8,opt,name=state,proto3,enum=google.events.cloud.beyondcorp.appgateways.v1.AppGateway_State" json:"state,omitempty"`
	// Output only. Server-defined URI for this resource.
	Uri string `protobuf:"bytes,9,opt,name=uri,proto3" json:"uri,omitempty"`
	// Output only. A list of connections allocated for the Gateway
	AllocatedConnections []*AppGateway_AllocatedConnection `protobuf:"bytes,10,rep,name=allocated_connections,json=allocatedConnections,proto3" json:"allocated_connections,omitempty"`
	// Required. The type of hosting used by the AppGateway.
	HostType AppGateway_HostType `protobuf:"varint,11,opt,name=host_type,json=hostType,proto3,enum=google.events.cloud.beyondcorp.appgateways.v1.AppGateway_HostType" json:"host_type,omitempty"`
}

func (x *AppGateway) Reset() {
	*x = AppGateway{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_beyondcorp_appgateways_v1_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppGateway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppGateway) ProtoMessage() {}

func (x *AppGateway) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_beyondcorp_appgateways_v1_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppGateway.ProtoReflect.Descriptor instead.
func (*AppGateway) Descriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_appgateways_v1_data_proto_rawDescGZIP(), []int{0}
}

func (x *AppGateway) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppGateway) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *AppGateway) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *AppGateway) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *AppGateway) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *AppGateway) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *AppGateway) GetType() AppGateway_Type {
	if x != nil {
		return x.Type
	}
	return AppGateway_TYPE_UNSPECIFIED
}

func (x *AppGateway) GetState() AppGateway_State {
	if x != nil {
		return x.State
	}
	return AppGateway_STATE_UNSPECIFIED
}

func (x *AppGateway) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *AppGateway) GetAllocatedConnections() []*AppGateway_AllocatedConnection {
	if x != nil {
		return x.AllocatedConnections
	}
	return nil
}

func (x *AppGateway) GetHostType() AppGateway_HostType {
	if x != nil {
		return x.HostType
	}
	return AppGateway_HOST_TYPE_UNSPECIFIED
}

// The data within all AppGateway events.
type AppGatewayEventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The AppGateway event payload. Unset for deletion events.
	Payload *AppGateway `protobuf:"bytes,1,opt,name=payload,proto3,oneof" json:"payload,omitempty"`
}

func (x *AppGatewayEventData) Reset() {
	*x = AppGatewayEventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_beyondcorp_appgateways_v1_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppGatewayEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppGatewayEventData) ProtoMessage() {}

func (x *AppGatewayEventData) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_beyondcorp_appgateways_v1_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppGatewayEventData.ProtoReflect.Descriptor instead.
func (*AppGatewayEventData) Descriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_appgateways_v1_data_proto_rawDescGZIP(), []int{1}
}

func (x *AppGatewayEventData) GetPayload() *AppGateway {
	if x != nil {
		return x.Payload
	}
	return nil
}

// Allocated connection of the AppGateway.
type AppGateway_AllocatedConnection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The PSC uri of an allocated connection
	PscUri string `protobuf:"bytes,1,opt,name=psc_uri,json=pscUri,proto3" json:"psc_uri,omitempty"`
	// Required. The ingress port of an allocated connection
	IngressPort int32 `protobuf:"varint,2,opt,name=ingress_port,json=ingressPort,proto3" json:"ingress_port,omitempty"`
}

func (x *AppGateway_AllocatedConnection) Reset() {
	*x = AppGateway_AllocatedConnection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_beyondcorp_appgateways_v1_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppGateway_AllocatedConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppGateway_AllocatedConnection) ProtoMessage() {}

func (x *AppGateway_AllocatedConnection) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_beyondcorp_appgateways_v1_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppGateway_AllocatedConnection.ProtoReflect.Descriptor instead.
func (*AppGateway_AllocatedConnection) Descriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_appgateways_v1_data_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AppGateway_AllocatedConnection) GetPscUri() string {
	if x != nil {
		return x.PscUri
	}
	return ""
}

func (x *AppGateway_AllocatedConnection) GetIngressPort() int32 {
	if x != nil {
		return x.IngressPort
	}
	return 0
}

var File_cloud_beyondcorp_appgateways_v1_data_proto protoreflect.FileDescriptor

var file_cloud_beyondcorp_appgateways_v1_data_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f,
	0x72, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x61, 0x70, 0x70,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x08, 0x0a,
	0x0a, 0x41, 0x70, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x5d, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x52, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x61, 0x70, 0x70,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x55, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x61, 0x70, 0x70, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x70, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x82, 0x01, 0x0a, 0x15, 0x61,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x14, 0x61, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x5f, 0x0a, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63,
	0x6f, 0x72, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x48, 0x6f,
	0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x1a, 0x51, 0x0a, 0x13, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x73, 0x63, 0x5f, 0x75,
	0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x73, 0x63, 0x55, 0x72, 0x69,
	0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x50,
	0x6f, 0x72, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2b,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x54, 0x43, 0x50, 0x5f, 0x50, 0x52, 0x4f, 0x58, 0x59, 0x10, 0x01, 0x22, 0x5f, 0x0a, 0x05, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x50, 0x44, 0x41, 0x54, 0x49,
	0x4e, 0x47, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47,
	0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x05, 0x22, 0x3b, 0x0a, 0x08,
	0x48, 0x6f, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x48, 0x4f, 0x53, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x43, 0x50, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f,
	0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x49, 0x47, 0x10, 0x01, 0x22, 0x7b, 0x0a, 0x13, 0x41, 0x70, 0x70,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x58, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f,
	0x72, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x70, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x48, 0x00, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x9e, 0x01, 0xaa, 0x02, 0x36, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x42, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x43,
	0x6f, 0x72, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x2d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x42, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x43,
	0x6f, 0x72, 0x70, 0x5c, 0x41, 0x70, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x5c,
	0x56, 0x31, 0xea, 0x02, 0x32, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x42, 0x65, 0x79, 0x6f,
	0x6e, 0x64, 0x43, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x41, 0x70, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_beyondcorp_appgateways_v1_data_proto_rawDescOnce sync.Once
	file_cloud_beyondcorp_appgateways_v1_data_proto_rawDescData = file_cloud_beyondcorp_appgateways_v1_data_proto_rawDesc
)

func file_cloud_beyondcorp_appgateways_v1_data_proto_rawDescGZIP() []byte {
	file_cloud_beyondcorp_appgateways_v1_data_proto_rawDescOnce.Do(func() {
		file_cloud_beyondcorp_appgateways_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_beyondcorp_appgateways_v1_data_proto_rawDescData)
	})
	return file_cloud_beyondcorp_appgateways_v1_data_proto_rawDescData
}

var file_cloud_beyondcorp_appgateways_v1_data_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_cloud_beyondcorp_appgateways_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cloud_beyondcorp_appgateways_v1_data_proto_goTypes = []any{
	(AppGateway_Type)(0),                   // 0: google.events.cloud.beyondcorp.appgateways.v1.AppGateway.Type
	(AppGateway_State)(0),                  // 1: google.events.cloud.beyondcorp.appgateways.v1.AppGateway.State
	(AppGateway_HostType)(0),               // 2: google.events.cloud.beyondcorp.appgateways.v1.AppGateway.HostType
	(*AppGateway)(nil),                     // 3: google.events.cloud.beyondcorp.appgateways.v1.AppGateway
	(*AppGatewayEventData)(nil),            // 4: google.events.cloud.beyondcorp.appgateways.v1.AppGatewayEventData
	(*AppGateway_AllocatedConnection)(nil), // 5: google.events.cloud.beyondcorp.appgateways.v1.AppGateway.AllocatedConnection
	nil,                                    // 6: google.events.cloud.beyondcorp.appgateways.v1.AppGateway.LabelsEntry
	(*timestamppb.Timestamp)(nil),          // 7: google.protobuf.Timestamp
}
var file_cloud_beyondcorp_appgateways_v1_data_proto_depIdxs = []int32{
	7, // 0: google.events.cloud.beyondcorp.appgateways.v1.AppGateway.create_time:type_name -> google.protobuf.Timestamp
	7, // 1: google.events.cloud.beyondcorp.appgateways.v1.AppGateway.update_time:type_name -> google.protobuf.Timestamp
	6, // 2: google.events.cloud.beyondcorp.appgateways.v1.AppGateway.labels:type_name -> google.events.cloud.beyondcorp.appgateways.v1.AppGateway.LabelsEntry
	0, // 3: google.events.cloud.beyondcorp.appgateways.v1.AppGateway.type:type_name -> google.events.cloud.beyondcorp.appgateways.v1.AppGateway.Type
	1, // 4: google.events.cloud.beyondcorp.appgateways.v1.AppGateway.state:type_name -> google.events.cloud.beyondcorp.appgateways.v1.AppGateway.State
	5, // 5: google.events.cloud.beyondcorp.appgateways.v1.AppGateway.allocated_connections:type_name -> google.events.cloud.beyondcorp.appgateways.v1.AppGateway.AllocatedConnection
	2, // 6: google.events.cloud.beyondcorp.appgateways.v1.AppGateway.host_type:type_name -> google.events.cloud.beyondcorp.appgateways.v1.AppGateway.HostType
	3, // 7: google.events.cloud.beyondcorp.appgateways.v1.AppGatewayEventData.payload:type_name -> google.events.cloud.beyondcorp.appgateways.v1.AppGateway
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_cloud_beyondcorp_appgateways_v1_data_proto_init() }
func file_cloud_beyondcorp_appgateways_v1_data_proto_init() {
	if File_cloud_beyondcorp_appgateways_v1_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_beyondcorp_appgateways_v1_data_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AppGateway); i {
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
		file_cloud_beyondcorp_appgateways_v1_data_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AppGatewayEventData); i {
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
		file_cloud_beyondcorp_appgateways_v1_data_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AppGateway_AllocatedConnection); i {
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
	file_cloud_beyondcorp_appgateways_v1_data_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_beyondcorp_appgateways_v1_data_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_beyondcorp_appgateways_v1_data_proto_goTypes,
		DependencyIndexes: file_cloud_beyondcorp_appgateways_v1_data_proto_depIdxs,
		EnumInfos:         file_cloud_beyondcorp_appgateways_v1_data_proto_enumTypes,
		MessageInfos:      file_cloud_beyondcorp_appgateways_v1_data_proto_msgTypes,
	}.Build()
	File_cloud_beyondcorp_appgateways_v1_data_proto = out.File
	file_cloud_beyondcorp_appgateways_v1_data_proto_rawDesc = nil
	file_cloud_beyondcorp_appgateways_v1_data_proto_goTypes = nil
	file_cloud_beyondcorp_appgateways_v1_data_proto_depIdxs = nil
}
