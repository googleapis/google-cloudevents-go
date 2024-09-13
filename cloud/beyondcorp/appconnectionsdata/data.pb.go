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
// source: cloud/beyondcorp/appconnections/v1/data.proto

package appconnectionsdata

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
// supported by BeyondCorp AppConnection.
type AppConnection_Type int32

const (
	// Default value. This value is unused.
	AppConnection_TYPE_UNSPECIFIED AppConnection_Type = 0
	// TCP Proxy based BeyondCorp AppConnection. API will default to this if
	// unset.
	AppConnection_TCP_PROXY AppConnection_Type = 1
)

// Enum value maps for AppConnection_Type.
var (
	AppConnection_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TCP_PROXY",
	}
	AppConnection_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TCP_PROXY":        1,
	}
)

func (x AppConnection_Type) Enum() *AppConnection_Type {
	p := new(AppConnection_Type)
	*p = x
	return p
}

func (x AppConnection_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppConnection_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_beyondcorp_appconnections_v1_data_proto_enumTypes[0].Descriptor()
}

func (AppConnection_Type) Type() protoreflect.EnumType {
	return &file_cloud_beyondcorp_appconnections_v1_data_proto_enumTypes[0]
}

func (x AppConnection_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppConnection_Type.Descriptor instead.
func (AppConnection_Type) EnumDescriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_appconnections_v1_data_proto_rawDescGZIP(), []int{0, 0}
}

// Represents the different states of a AppConnection.
type AppConnection_State int32

const (
	// Default value. This value is unused.
	AppConnection_STATE_UNSPECIFIED AppConnection_State = 0
	// AppConnection is being created.
	AppConnection_CREATING AppConnection_State = 1
	// AppConnection has been created.
	AppConnection_CREATED AppConnection_State = 2
	// AppConnection's configuration is being updated.
	AppConnection_UPDATING AppConnection_State = 3
	// AppConnection is being deleted.
	AppConnection_DELETING AppConnection_State = 4
	// AppConnection is down and may be restored in the future.
	// This happens when CCFE sends ProjectState = OFF.
	AppConnection_DOWN AppConnection_State = 5
)

// Enum value maps for AppConnection_State.
var (
	AppConnection_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "CREATED",
		3: "UPDATING",
		4: "DELETING",
		5: "DOWN",
	}
	AppConnection_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          1,
		"CREATED":           2,
		"UPDATING":          3,
		"DELETING":          4,
		"DOWN":              5,
	}
)

func (x AppConnection_State) Enum() *AppConnection_State {
	p := new(AppConnection_State)
	*p = x
	return p
}

func (x AppConnection_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppConnection_State) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_beyondcorp_appconnections_v1_data_proto_enumTypes[1].Descriptor()
}

func (AppConnection_State) Type() protoreflect.EnumType {
	return &file_cloud_beyondcorp_appconnections_v1_data_proto_enumTypes[1]
}

func (x AppConnection_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppConnection_State.Descriptor instead.
func (AppConnection_State) EnumDescriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_appconnections_v1_data_proto_rawDescGZIP(), []int{0, 1}
}

// Enum listing possible gateway hosting options.
type AppConnection_Gateway_Type int32

const (
	// Default value. This value is unused.
	AppConnection_Gateway_TYPE_UNSPECIFIED AppConnection_Gateway_Type = 0
	// Gateway hosted in a GCP regional managed instance group.
	AppConnection_Gateway_GCP_REGIONAL_MIG AppConnection_Gateway_Type = 1
)

// Enum value maps for AppConnection_Gateway_Type.
var (
	AppConnection_Gateway_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "GCP_REGIONAL_MIG",
	}
	AppConnection_Gateway_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"GCP_REGIONAL_MIG": 1,
	}
)

func (x AppConnection_Gateway_Type) Enum() *AppConnection_Gateway_Type {
	p := new(AppConnection_Gateway_Type)
	*p = x
	return p
}

func (x AppConnection_Gateway_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppConnection_Gateway_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_beyondcorp_appconnections_v1_data_proto_enumTypes[2].Descriptor()
}

func (AppConnection_Gateway_Type) Type() protoreflect.EnumType {
	return &file_cloud_beyondcorp_appconnections_v1_data_proto_enumTypes[2]
}

func (x AppConnection_Gateway_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppConnection_Gateway_Type.Descriptor instead.
func (AppConnection_Gateway_Type) EnumDescriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_appconnections_v1_data_proto_rawDescGZIP(), []int{0, 1, 0}
}

// A BeyondCorp AppConnection resource represents a BeyondCorp protected
// AppConnection to a remote application. It creates all the necessary GCP
// components needed for creating a BeyondCorp protected AppConnection. Multiple
// connectors can be authorised for a single AppConnection.
type AppConnection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Unique resource name of the AppConnection.
	// The name is ignored when creating a AppConnection.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Timestamp when the resource was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp when the resource was last modified.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Optional. Resource labels to represent user provided metadata.
	Labels map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Optional. An arbitrary user-provided name for the AppConnection. Cannot
	// exceed 64 characters.
	DisplayName string `protobuf:"bytes,5,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Output only. A unique identifier for the instance generated by the
	// system.
	Uid string `protobuf:"bytes,6,opt,name=uid,proto3" json:"uid,omitempty"`
	// Required. The type of network connectivity used by the AppConnection.
	Type AppConnection_Type `protobuf:"varint,7,opt,name=type,proto3,enum=google.events.cloud.beyondcorp.appconnections.v1.AppConnection_Type" json:"type,omitempty"`
	// Required. Address of the remote application endpoint for the BeyondCorp
	// AppConnection.
	ApplicationEndpoint *AppConnection_ApplicationEndpoint `protobuf:"bytes,8,opt,name=application_endpoint,json=applicationEndpoint,proto3" json:"application_endpoint,omitempty"`
	// Optional. List of [google.cloud.beyondcorp.v1main.Connector.name] that are
	// authorised to be associated with this AppConnection.
	Connectors []string `protobuf:"bytes,9,rep,name=connectors,proto3" json:"connectors,omitempty"`
	// Output only. The current state of the AppConnection.
	State AppConnection_State `protobuf:"varint,10,opt,name=state,proto3,enum=google.events.cloud.beyondcorp.appconnections.v1.AppConnection_State" json:"state,omitempty"`
	// Optional. Gateway used by the AppConnection.
	Gateway *AppConnection_Gateway `protobuf:"bytes,11,opt,name=gateway,proto3" json:"gateway,omitempty"`
}

func (x *AppConnection) Reset() {
	*x = AppConnection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_beyondcorp_appconnections_v1_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConnection) ProtoMessage() {}

func (x *AppConnection) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_beyondcorp_appconnections_v1_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConnection.ProtoReflect.Descriptor instead.
func (*AppConnection) Descriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_appconnections_v1_data_proto_rawDescGZIP(), []int{0}
}

func (x *AppConnection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppConnection) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *AppConnection) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *AppConnection) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *AppConnection) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *AppConnection) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *AppConnection) GetType() AppConnection_Type {
	if x != nil {
		return x.Type
	}
	return AppConnection_TYPE_UNSPECIFIED
}

func (x *AppConnection) GetApplicationEndpoint() *AppConnection_ApplicationEndpoint {
	if x != nil {
		return x.ApplicationEndpoint
	}
	return nil
}

func (x *AppConnection) GetConnectors() []string {
	if x != nil {
		return x.Connectors
	}
	return nil
}

func (x *AppConnection) GetState() AppConnection_State {
	if x != nil {
		return x.State
	}
	return AppConnection_STATE_UNSPECIFIED
}

func (x *AppConnection) GetGateway() *AppConnection_Gateway {
	if x != nil {
		return x.Gateway
	}
	return nil
}

// The data within all AppConnection events.
type AppConnectionEventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The AppConnection event payload. Unset for deletion events.
	Payload *AppConnection `protobuf:"bytes,1,opt,name=payload,proto3,oneof" json:"payload,omitempty"`
}

func (x *AppConnectionEventData) Reset() {
	*x = AppConnectionEventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_beyondcorp_appconnections_v1_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConnectionEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConnectionEventData) ProtoMessage() {}

func (x *AppConnectionEventData) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_beyondcorp_appconnections_v1_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConnectionEventData.ProtoReflect.Descriptor instead.
func (*AppConnectionEventData) Descriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_appconnections_v1_data_proto_rawDescGZIP(), []int{1}
}

func (x *AppConnectionEventData) GetPayload() *AppConnection {
	if x != nil {
		return x.Payload
	}
	return nil
}

// ApplicationEndpoint represents a remote application endpoint.
type AppConnection_ApplicationEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Hostname or IP address of the remote application endpoint.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Required. Port of the remote application endpoint.
	Port int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *AppConnection_ApplicationEndpoint) Reset() {
	*x = AppConnection_ApplicationEndpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_beyondcorp_appconnections_v1_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConnection_ApplicationEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConnection_ApplicationEndpoint) ProtoMessage() {}

func (x *AppConnection_ApplicationEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_beyondcorp_appconnections_v1_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConnection_ApplicationEndpoint.ProtoReflect.Descriptor instead.
func (*AppConnection_ApplicationEndpoint) Descriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_appconnections_v1_data_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AppConnection_ApplicationEndpoint) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *AppConnection_ApplicationEndpoint) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

// Gateway represents a user facing component that serves as an entrance to
// enable connectivity.
type AppConnection_Gateway struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The type of hosting used by the gateway.
	Type AppConnection_Gateway_Type `protobuf:"varint,2,opt,name=type,proto3,enum=google.events.cloud.beyondcorp.appconnections.v1.AppConnection_Gateway_Type" json:"type,omitempty"`
	// Output only. Server-defined URI for this resource.
	Uri string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	// Output only. Ingress port reserved on the gateways for this
	// AppConnection, if not specified or zero, the default port is 19443.
	IngressPort int32 `protobuf:"varint,4,opt,name=ingress_port,json=ingressPort,proto3" json:"ingress_port,omitempty"`
	// Required. AppGateway name in following format:
	// `projects/{project_id}/locations/{location_id}/appgateways/{gateway_id}`
	AppGateway string `protobuf:"bytes,5,opt,name=app_gateway,json=appGateway,proto3" json:"app_gateway,omitempty"`
	// Output only. L7 private service connection for this resource.
	L7Psc string `protobuf:"bytes,6,opt,name=l7psc,proto3" json:"l7psc,omitempty"`
}

func (x *AppConnection_Gateway) Reset() {
	*x = AppConnection_Gateway{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_beyondcorp_appconnections_v1_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConnection_Gateway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConnection_Gateway) ProtoMessage() {}

func (x *AppConnection_Gateway) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_beyondcorp_appconnections_v1_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConnection_Gateway.ProtoReflect.Descriptor instead.
func (*AppConnection_Gateway) Descriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_appconnections_v1_data_proto_rawDescGZIP(), []int{0, 1}
}

func (x *AppConnection_Gateway) GetType() AppConnection_Gateway_Type {
	if x != nil {
		return x.Type
	}
	return AppConnection_Gateway_TYPE_UNSPECIFIED
}

func (x *AppConnection_Gateway) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *AppConnection_Gateway) GetIngressPort() int32 {
	if x != nil {
		return x.IngressPort
	}
	return 0
}

func (x *AppConnection_Gateway) GetAppGateway() string {
	if x != nil {
		return x.AppGateway
	}
	return ""
}

func (x *AppConnection_Gateway) GetL7Psc() string {
	if x != nil {
		return x.L7Psc
	}
	return ""
}

var File_cloud_beyondcorp_appconnections_v1_data_proto protoreflect.FileDescriptor

var file_cloud_beyondcorp_appconnections_v1_data_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f,
	0x72, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x90, 0x0a, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x63, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63,
	0x6f, 0x72, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x58, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x44, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70,
	0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x14, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x53, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f,
	0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x13, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12,
	0x5b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x45,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x61, 0x0a, 0x07,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x61,
	0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x1a,
	0x3d, 0x0a, 0x13, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x8b,
	0x02, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x60, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62,
	0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x21,
	0x0a, 0x0c, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x72,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x37, 0x70, 0x73, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x37, 0x70, 0x73, 0x63, 0x22, 0x32, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x43, 0x50, 0x5f, 0x52, 0x45,
	0x47, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x49, 0x47, 0x10, 0x01, 0x1a, 0x39, 0x0a, 0x0b,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x43, 0x50, 0x5f, 0x50, 0x52, 0x4f,
	0x58, 0x59, 0x10, 0x01, 0x22, 0x5f, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a,
	0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x0c, 0x0a, 0x08, 0x55, 0x50, 0x44, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0c, 0x0a,
	0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x44,
	0x4f, 0x57, 0x4e, 0x10, 0x05, 0x22, 0x84, 0x01, 0x0a, 0x16, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x5e, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x63, 0x6f,
	0x72, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x88, 0x01, 0x01,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0xa7, 0x01, 0xaa,
	0x02, 0x39, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x42,
	0x65, 0x79, 0x6f, 0x6e, 0x64, 0x43, 0x6f, 0x72, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x30, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x42, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x43, 0x6f, 0x72, 0x70, 0x5c, 0x41, 0x70, 0x70,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x56, 0x31, 0xea, 0x02,
	0x35, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a,
	0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x42, 0x65, 0x79, 0x6f, 0x6e, 0x64, 0x43, 0x6f,
	0x72, 0x70, 0x3a, 0x3a, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_beyondcorp_appconnections_v1_data_proto_rawDescOnce sync.Once
	file_cloud_beyondcorp_appconnections_v1_data_proto_rawDescData = file_cloud_beyondcorp_appconnections_v1_data_proto_rawDesc
)

func file_cloud_beyondcorp_appconnections_v1_data_proto_rawDescGZIP() []byte {
	file_cloud_beyondcorp_appconnections_v1_data_proto_rawDescOnce.Do(func() {
		file_cloud_beyondcorp_appconnections_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_beyondcorp_appconnections_v1_data_proto_rawDescData)
	})
	return file_cloud_beyondcorp_appconnections_v1_data_proto_rawDescData
}

var file_cloud_beyondcorp_appconnections_v1_data_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_cloud_beyondcorp_appconnections_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_beyondcorp_appconnections_v1_data_proto_goTypes = []any{
	(AppConnection_Type)(0),                   // 0: google.events.cloud.beyondcorp.appconnections.v1.AppConnection.Type
	(AppConnection_State)(0),                  // 1: google.events.cloud.beyondcorp.appconnections.v1.AppConnection.State
	(AppConnection_Gateway_Type)(0),           // 2: google.events.cloud.beyondcorp.appconnections.v1.AppConnection.Gateway.Type
	(*AppConnection)(nil),                     // 3: google.events.cloud.beyondcorp.appconnections.v1.AppConnection
	(*AppConnectionEventData)(nil),            // 4: google.events.cloud.beyondcorp.appconnections.v1.AppConnectionEventData
	(*AppConnection_ApplicationEndpoint)(nil), // 5: google.events.cloud.beyondcorp.appconnections.v1.AppConnection.ApplicationEndpoint
	(*AppConnection_Gateway)(nil),             // 6: google.events.cloud.beyondcorp.appconnections.v1.AppConnection.Gateway
	nil,                                       // 7: google.events.cloud.beyondcorp.appconnections.v1.AppConnection.LabelsEntry
	(*timestamppb.Timestamp)(nil),             // 8: google.protobuf.Timestamp
}
var file_cloud_beyondcorp_appconnections_v1_data_proto_depIdxs = []int32{
	8, // 0: google.events.cloud.beyondcorp.appconnections.v1.AppConnection.create_time:type_name -> google.protobuf.Timestamp
	8, // 1: google.events.cloud.beyondcorp.appconnections.v1.AppConnection.update_time:type_name -> google.protobuf.Timestamp
	7, // 2: google.events.cloud.beyondcorp.appconnections.v1.AppConnection.labels:type_name -> google.events.cloud.beyondcorp.appconnections.v1.AppConnection.LabelsEntry
	0, // 3: google.events.cloud.beyondcorp.appconnections.v1.AppConnection.type:type_name -> google.events.cloud.beyondcorp.appconnections.v1.AppConnection.Type
	5, // 4: google.events.cloud.beyondcorp.appconnections.v1.AppConnection.application_endpoint:type_name -> google.events.cloud.beyondcorp.appconnections.v1.AppConnection.ApplicationEndpoint
	1, // 5: google.events.cloud.beyondcorp.appconnections.v1.AppConnection.state:type_name -> google.events.cloud.beyondcorp.appconnections.v1.AppConnection.State
	6, // 6: google.events.cloud.beyondcorp.appconnections.v1.AppConnection.gateway:type_name -> google.events.cloud.beyondcorp.appconnections.v1.AppConnection.Gateway
	3, // 7: google.events.cloud.beyondcorp.appconnections.v1.AppConnectionEventData.payload:type_name -> google.events.cloud.beyondcorp.appconnections.v1.AppConnection
	2, // 8: google.events.cloud.beyondcorp.appconnections.v1.AppConnection.Gateway.type:type_name -> google.events.cloud.beyondcorp.appconnections.v1.AppConnection.Gateway.Type
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_cloud_beyondcorp_appconnections_v1_data_proto_init() }
func file_cloud_beyondcorp_appconnections_v1_data_proto_init() {
	if File_cloud_beyondcorp_appconnections_v1_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_beyondcorp_appconnections_v1_data_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AppConnection); i {
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
		file_cloud_beyondcorp_appconnections_v1_data_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AppConnectionEventData); i {
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
		file_cloud_beyondcorp_appconnections_v1_data_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AppConnection_ApplicationEndpoint); i {
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
		file_cloud_beyondcorp_appconnections_v1_data_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AppConnection_Gateway); i {
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
	file_cloud_beyondcorp_appconnections_v1_data_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_beyondcorp_appconnections_v1_data_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_beyondcorp_appconnections_v1_data_proto_goTypes,
		DependencyIndexes: file_cloud_beyondcorp_appconnections_v1_data_proto_depIdxs,
		EnumInfos:         file_cloud_beyondcorp_appconnections_v1_data_proto_enumTypes,
		MessageInfos:      file_cloud_beyondcorp_appconnections_v1_data_proto_msgTypes,
	}.Build()
	File_cloud_beyondcorp_appconnections_v1_data_proto = out.File
	file_cloud_beyondcorp_appconnections_v1_data_proto_rawDesc = nil
	file_cloud_beyondcorp_appconnections_v1_data_proto_goTypes = nil
	file_cloud_beyondcorp_appconnections_v1_data_proto_depIdxs = nil
}
