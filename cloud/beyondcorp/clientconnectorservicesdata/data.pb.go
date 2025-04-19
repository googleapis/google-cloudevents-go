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
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.6
// source: cloud/beyondcorp/clientconnectorservices/v1/data.proto

package clientconnectorservicesdata

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

// Represents the different states of a ClientConnectorService.
type ClientConnectorService_State int32

const (
	// Default value. This value is unused.
	ClientConnectorService_STATE_UNSPECIFIED ClientConnectorService_State = 0
	// ClientConnectorService is being created.
	ClientConnectorService_CREATING ClientConnectorService_State = 1
	// ClientConnectorService is being updated.
	ClientConnectorService_UPDATING ClientConnectorService_State = 2
	// ClientConnectorService is being deleted.
	ClientConnectorService_DELETING ClientConnectorService_State = 3
	// ClientConnectorService is running.
	ClientConnectorService_RUNNING ClientConnectorService_State = 4
	// ClientConnectorService is down and may be restored in the future.
	// This happens when CCFE sends ProjectState = OFF.
	ClientConnectorService_DOWN ClientConnectorService_State = 5
	// ClientConnectorService encountered an error and is in an indeterministic
	// state.
	ClientConnectorService_ERROR ClientConnectorService_State = 6
)

// Enum value maps for ClientConnectorService_State.
var (
	ClientConnectorService_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "UPDATING",
		3: "DELETING",
		4: "RUNNING",
		5: "DOWN",
		6: "ERROR",
	}
	ClientConnectorService_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          1,
		"UPDATING":          2,
		"DELETING":          3,
		"RUNNING":           4,
		"DOWN":              5,
		"ERROR":             6,
	}
)

func (x ClientConnectorService_State) Enum() *ClientConnectorService_State {
	p := new(ClientConnectorService_State)
	*p = x
	return p
}

func (x ClientConnectorService_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientConnectorService_State) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_enumTypes[0].Descriptor()
}

func (ClientConnectorService_State) Type() protoreflect.EnumType {
	return &file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_enumTypes[0]
}

func (x ClientConnectorService_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientConnectorService_State.Descriptor instead.
func (ClientConnectorService_State) EnumDescriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_rawDescGZIP(), []int{0, 0}
}

// The protocol used to connect to the server.
type ClientConnectorService_Ingress_Config_TransportProtocol int32

const (
	// Default value. This value is unused.
	ClientConnectorService_Ingress_Config_TRANSPORT_PROTOCOL_UNSPECIFIED ClientConnectorService_Ingress_Config_TransportProtocol = 0
	// TCP protocol.
	ClientConnectorService_Ingress_Config_TCP ClientConnectorService_Ingress_Config_TransportProtocol = 1
)

// Enum value maps for ClientConnectorService_Ingress_Config_TransportProtocol.
var (
	ClientConnectorService_Ingress_Config_TransportProtocol_name = map[int32]string{
		0: "TRANSPORT_PROTOCOL_UNSPECIFIED",
		1: "TCP",
	}
	ClientConnectorService_Ingress_Config_TransportProtocol_value = map[string]int32{
		"TRANSPORT_PROTOCOL_UNSPECIFIED": 0,
		"TCP":                            1,
	}
)

func (x ClientConnectorService_Ingress_Config_TransportProtocol) Enum() *ClientConnectorService_Ingress_Config_TransportProtocol {
	p := new(ClientConnectorService_Ingress_Config_TransportProtocol)
	*p = x
	return p
}

func (x ClientConnectorService_Ingress_Config_TransportProtocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientConnectorService_Ingress_Config_TransportProtocol) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_enumTypes[1].Descriptor()
}

func (ClientConnectorService_Ingress_Config_TransportProtocol) Type() protoreflect.EnumType {
	return &file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_enumTypes[1]
}

func (x ClientConnectorService_Ingress_Config_TransportProtocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientConnectorService_Ingress_Config_TransportProtocol.Descriptor instead.
func (ClientConnectorService_Ingress_Config_TransportProtocol) EnumDescriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

// Message describing ClientConnectorService object.
type ClientConnectorService struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. Name of resource. The name is ignored during creation.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. [Output only] Create time stamp.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. [Output only] Update time stamp.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Optional. User-provided name.
	// The display name should follow certain format.
	// * Must be 6 to 30 characters in length.
	// * Can only contain lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Required. The details of the ingress settings.
	Ingress *ClientConnectorService_Ingress `protobuf:"bytes,6,opt,name=ingress,proto3" json:"ingress,omitempty"`
	// Required. The details of the egress settings.
	Egress *ClientConnectorService_Egress `protobuf:"bytes,7,opt,name=egress,proto3" json:"egress,omitempty"`
	// Output only. The operational state of the ClientConnectorService.
	State         ClientConnectorService_State `protobuf:"varint,8,opt,name=state,proto3,enum=google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService_State" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientConnectorService) Reset() {
	*x = ClientConnectorService{}
	mi := &file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientConnectorService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConnectorService) ProtoMessage() {}

func (x *ClientConnectorService) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConnectorService.ProtoReflect.Descriptor instead.
func (*ClientConnectorService) Descriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_rawDescGZIP(), []int{0}
}

func (x *ClientConnectorService) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClientConnectorService) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ClientConnectorService) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *ClientConnectorService) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *ClientConnectorService) GetIngress() *ClientConnectorService_Ingress {
	if x != nil {
		return x.Ingress
	}
	return nil
}

func (x *ClientConnectorService) GetEgress() *ClientConnectorService_Egress {
	if x != nil {
		return x.Egress
	}
	return nil
}

func (x *ClientConnectorService) GetState() ClientConnectorService_State {
	if x != nil {
		return x.State
	}
	return ClientConnectorService_STATE_UNSPECIFIED
}

// The data within all ClientConnectorService events.
type ClientConnectorServiceEventData struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Optional. The ClientConnectorService event payload. Unset for deletion
	// events.
	Payload       *ClientConnectorService `protobuf:"bytes,1,opt,name=payload,proto3,oneof" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientConnectorServiceEventData) Reset() {
	*x = ClientConnectorServiceEventData{}
	mi := &file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientConnectorServiceEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConnectorServiceEventData) ProtoMessage() {}

func (x *ClientConnectorServiceEventData) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConnectorServiceEventData.ProtoReflect.Descriptor instead.
func (*ClientConnectorServiceEventData) Descriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_rawDescGZIP(), []int{1}
}

func (x *ClientConnectorServiceEventData) GetPayload() *ClientConnectorService {
	if x != nil {
		return x.Payload
	}
	return nil
}

// Settings of how to connect to the ClientGateway.
// One of the following options should be set.
type ClientConnectorService_Ingress struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to IngressConfig:
	//
	//	*ClientConnectorService_Ingress_Config_
	IngressConfig isClientConnectorService_Ingress_IngressConfig `protobuf_oneof:"ingress_config"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientConnectorService_Ingress) Reset() {
	*x = ClientConnectorService_Ingress{}
	mi := &file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientConnectorService_Ingress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConnectorService_Ingress) ProtoMessage() {}

func (x *ClientConnectorService_Ingress) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConnectorService_Ingress.ProtoReflect.Descriptor instead.
func (*ClientConnectorService_Ingress) Descriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ClientConnectorService_Ingress) GetIngressConfig() isClientConnectorService_Ingress_IngressConfig {
	if x != nil {
		return x.IngressConfig
	}
	return nil
}

func (x *ClientConnectorService_Ingress) GetConfig() *ClientConnectorService_Ingress_Config {
	if x != nil {
		if x, ok := x.IngressConfig.(*ClientConnectorService_Ingress_Config_); ok {
			return x.Config
		}
	}
	return nil
}

type isClientConnectorService_Ingress_IngressConfig interface {
	isClientConnectorService_Ingress_IngressConfig()
}

type ClientConnectorService_Ingress_Config_ struct {
	// The basic ingress config for ClientGateways.
	Config *ClientConnectorService_Ingress_Config `protobuf:"bytes,1,opt,name=config,proto3,oneof"`
}

func (*ClientConnectorService_Ingress_Config_) isClientConnectorService_Ingress_IngressConfig() {}

// The details of the egress info. One of the following options should be set.
type ClientConnectorService_Egress struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to DestinationType:
	//
	//	*ClientConnectorService_Egress_PeeredVpc_
	DestinationType isClientConnectorService_Egress_DestinationType `protobuf_oneof:"destination_type"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ClientConnectorService_Egress) Reset() {
	*x = ClientConnectorService_Egress{}
	mi := &file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientConnectorService_Egress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConnectorService_Egress) ProtoMessage() {}

func (x *ClientConnectorService_Egress) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConnectorService_Egress.ProtoReflect.Descriptor instead.
func (*ClientConnectorService_Egress) Descriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ClientConnectorService_Egress) GetDestinationType() isClientConnectorService_Egress_DestinationType {
	if x != nil {
		return x.DestinationType
	}
	return nil
}

func (x *ClientConnectorService_Egress) GetPeeredVpc() *ClientConnectorService_Egress_PeeredVpc {
	if x != nil {
		if x, ok := x.DestinationType.(*ClientConnectorService_Egress_PeeredVpc_); ok {
			return x.PeeredVpc
		}
	}
	return nil
}

type isClientConnectorService_Egress_DestinationType interface {
	isClientConnectorService_Egress_DestinationType()
}

type ClientConnectorService_Egress_PeeredVpc_ struct {
	// A VPC from the consumer project.
	PeeredVpc *ClientConnectorService_Egress_PeeredVpc `protobuf:"bytes,1,opt,name=peered_vpc,json=peeredVpc,proto3,oneof"`
}

func (*ClientConnectorService_Egress_PeeredVpc_) isClientConnectorService_Egress_DestinationType() {}

// The basic ingress config for ClientGateways.
type ClientConnectorService_Ingress_Config struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. Immutable. The transport protocol used between the client and
	// the server.
	TransportProtocol ClientConnectorService_Ingress_Config_TransportProtocol `protobuf:"varint,1,opt,name=transport_protocol,json=transportProtocol,proto3,enum=google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService_Ingress_Config_TransportProtocol" json:"transport_protocol,omitempty"`
	// Required. The settings used to configure basic ClientGateways.
	DestinationRoutes []*ClientConnectorService_Ingress_Config_DestinationRoute `protobuf:"bytes,2,rep,name=destination_routes,json=destinationRoutes,proto3" json:"destination_routes,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *ClientConnectorService_Ingress_Config) Reset() {
	*x = ClientConnectorService_Ingress_Config{}
	mi := &file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientConnectorService_Ingress_Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConnectorService_Ingress_Config) ProtoMessage() {}

func (x *ClientConnectorService_Ingress_Config) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConnectorService_Ingress_Config.ProtoReflect.Descriptor instead.
func (*ClientConnectorService_Ingress_Config) Descriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ClientConnectorService_Ingress_Config) GetTransportProtocol() ClientConnectorService_Ingress_Config_TransportProtocol {
	if x != nil {
		return x.TransportProtocol
	}
	return ClientConnectorService_Ingress_Config_TRANSPORT_PROTOCOL_UNSPECIFIED
}

func (x *ClientConnectorService_Ingress_Config) GetDestinationRoutes() []*ClientConnectorService_Ingress_Config_DestinationRoute {
	if x != nil {
		return x.DestinationRoutes
	}
	return nil
}

// The setting used to configure ClientGateways.
// It is adding routes to the client's routing table
// after the connection is established.
type ClientConnectorService_Ingress_Config_DestinationRoute struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The network address of the subnet
	// for which the packet is routed to the ClientGateway.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Required. The network mask of the subnet
	// for which the packet is routed to the ClientGateway.
	Netmask       string `protobuf:"bytes,2,opt,name=netmask,proto3" json:"netmask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientConnectorService_Ingress_Config_DestinationRoute) Reset() {
	*x = ClientConnectorService_Ingress_Config_DestinationRoute{}
	mi := &file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientConnectorService_Ingress_Config_DestinationRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConnectorService_Ingress_Config_DestinationRoute) ProtoMessage() {}

func (x *ClientConnectorService_Ingress_Config_DestinationRoute) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConnectorService_Ingress_Config_DestinationRoute.ProtoReflect.Descriptor instead.
func (*ClientConnectorService_Ingress_Config_DestinationRoute) Descriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

func (x *ClientConnectorService_Ingress_Config_DestinationRoute) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ClientConnectorService_Ingress_Config_DestinationRoute) GetNetmask() string {
	if x != nil {
		return x.Netmask
	}
	return ""
}

// The peered VPC owned by the consumer project.
type ClientConnectorService_Egress_PeeredVpc struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The name of the peered VPC owned by the consumer project.
	NetworkVpc    string `protobuf:"bytes,1,opt,name=network_vpc,json=networkVpc,proto3" json:"network_vpc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientConnectorService_Egress_PeeredVpc) Reset() {
	*x = ClientConnectorService_Egress_PeeredVpc{}
	mi := &file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientConnectorService_Egress_PeeredVpc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConnectorService_Egress_PeeredVpc) ProtoMessage() {}

func (x *ClientConnectorService_Egress_PeeredVpc) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConnectorService_Egress_PeeredVpc.ProtoReflect.Descriptor instead.
func (*ClientConnectorService_Egress_PeeredVpc) Descriptor() ([]byte, []int) {
	return file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *ClientConnectorService_Egress_PeeredVpc) GetNetworkVpc() string {
	if x != nil {
		return x.NetworkVpc
	}
	return ""
}

var File_cloud_beyondcorp_clientconnectorservices_v1_data_proto protoreflect.FileDescriptor

const file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_rawDesc = "" +
	"\n" +
	"6cloud/beyondcorp/clientconnectorservices/v1/data.proto\x129google.events.cloud.beyondcorp.clientconnectorservices.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"\xd4\v\n" +
	"\x16ClientConnectorService\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12;\n" +
	"\vcreate_time\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"createTime\x12;\n" +
	"\vupdate_time\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"updateTime\x12!\n" +
	"\fdisplay_name\x18\x04 \x01(\tR\vdisplayName\x12s\n" +
	"\aingress\x18\x06 \x01(\v2Y.google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.IngressR\aingress\x12p\n" +
	"\x06egress\x18\a \x01(\v2X.google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.EgressR\x06egress\x12m\n" +
	"\x05state\x18\b \x01(\x0e2W.google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.StateR\x05state\x1a\xf3\x04\n" +
	"\aIngress\x12z\n" +
	"\x06config\x18\x01 \x01(\v2`.google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress.ConfigH\x00R\x06config\x1a\xd9\x03\n" +
	"\x06Config\x12\xa1\x01\n" +
	"\x12transport_protocol\x18\x01 \x01(\x0e2r.google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress.Config.TransportProtocolR\x11transportProtocol\x12\xa0\x01\n" +
	"\x12destination_routes\x18\x02 \x03(\v2q.google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress.Config.DestinationRouteR\x11destinationRoutes\x1aF\n" +
	"\x10DestinationRoute\x12\x18\n" +
	"\aaddress\x18\x01 \x01(\tR\aaddress\x12\x18\n" +
	"\anetmask\x18\x02 \x01(\tR\anetmask\"@\n" +
	"\x11TransportProtocol\x12\"\n" +
	"\x1eTRANSPORT_PROTOCOL_UNSPECIFIED\x10\x00\x12\a\n" +
	"\x03TCP\x10\x01B\x10\n" +
	"\x0eingress_config\x1a\xd0\x01\n" +
	"\x06Egress\x12\x83\x01\n" +
	"\n" +
	"peered_vpc\x18\x01 \x01(\v2b.google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Egress.PeeredVpcH\x00R\tpeeredVpc\x1a,\n" +
	"\tPeeredVpc\x12\x1f\n" +
	"\vnetwork_vpc\x18\x01 \x01(\tR\n" +
	"networkVpcB\x12\n" +
	"\x10destination_type\"j\n" +
	"\x05State\x12\x15\n" +
	"\x11STATE_UNSPECIFIED\x10\x00\x12\f\n" +
	"\bCREATING\x10\x01\x12\f\n" +
	"\bUPDATING\x10\x02\x12\f\n" +
	"\bDELETING\x10\x03\x12\v\n" +
	"\aRUNNING\x10\x04\x12\b\n" +
	"\x04DOWN\x10\x05\x12\t\n" +
	"\x05ERROR\x10\x06\"\x9f\x01\n" +
	"\x1fClientConnectorServiceEventData\x12p\n" +
	"\apayload\x18\x01 \x01(\v2Q.google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorServiceH\x00R\apayload\x88\x01\x01B\n" +
	"\n" +
	"\b_payloadB\xc2\x01\xaa\x02BGoogle.Events.Protobuf.Cloud.BeyondCorp.ClientConnectorServices.V1\xca\x029Google\\Events\\Cloud\\BeyondCorp\\ClientConnectorServices\\V1\xea\x02>Google::Events::Cloud::BeyondCorp::ClientConnectorServices::V1b\x06proto3"

var (
	file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_rawDescOnce sync.Once
	file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_rawDescData []byte
)

func file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_rawDescGZIP() []byte {
	file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_rawDescOnce.Do(func() {
		file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_rawDesc), len(file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_rawDesc)))
	})
	return file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_rawDescData
}

var file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_goTypes = []any{
	(ClientConnectorService_State)(0),                              // 0: google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.State
	(ClientConnectorService_Ingress_Config_TransportProtocol)(0),   // 1: google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress.Config.TransportProtocol
	(*ClientConnectorService)(nil),                                 // 2: google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService
	(*ClientConnectorServiceEventData)(nil),                        // 3: google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorServiceEventData
	(*ClientConnectorService_Ingress)(nil),                         // 4: google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress
	(*ClientConnectorService_Egress)(nil),                          // 5: google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Egress
	(*ClientConnectorService_Ingress_Config)(nil),                  // 6: google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress.Config
	(*ClientConnectorService_Ingress_Config_DestinationRoute)(nil), // 7: google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress.Config.DestinationRoute
	(*ClientConnectorService_Egress_PeeredVpc)(nil),                // 8: google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Egress.PeeredVpc
	(*timestamppb.Timestamp)(nil),                                  // 9: google.protobuf.Timestamp
}
var file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_depIdxs = []int32{
	9,  // 0: google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.create_time:type_name -> google.protobuf.Timestamp
	9,  // 1: google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.update_time:type_name -> google.protobuf.Timestamp
	4,  // 2: google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.ingress:type_name -> google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress
	5,  // 3: google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.egress:type_name -> google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Egress
	0,  // 4: google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.state:type_name -> google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.State
	2,  // 5: google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorServiceEventData.payload:type_name -> google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService
	6,  // 6: google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress.config:type_name -> google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress.Config
	8,  // 7: google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Egress.peered_vpc:type_name -> google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Egress.PeeredVpc
	1,  // 8: google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress.Config.transport_protocol:type_name -> google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress.Config.TransportProtocol
	7,  // 9: google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress.Config.destination_routes:type_name -> google.events.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress.Config.DestinationRoute
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_init() }
func file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_init() {
	if File_cloud_beyondcorp_clientconnectorservices_v1_data_proto != nil {
		return
	}
	file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_msgTypes[1].OneofWrappers = []any{}
	file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_msgTypes[2].OneofWrappers = []any{
		(*ClientConnectorService_Ingress_Config_)(nil),
	}
	file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_msgTypes[3].OneofWrappers = []any{
		(*ClientConnectorService_Egress_PeeredVpc_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_rawDesc), len(file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_goTypes,
		DependencyIndexes: file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_depIdxs,
		EnumInfos:         file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_enumTypes,
		MessageInfos:      file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_msgTypes,
	}.Build()
	File_cloud_beyondcorp_clientconnectorservices_v1_data_proto = out.File
	file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_goTypes = nil
	file_cloud_beyondcorp_clientconnectorservices_v1_data_proto_depIdxs = nil
}
