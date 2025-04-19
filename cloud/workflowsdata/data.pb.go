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
// source: cloud/workflows/v1/data.proto

package workflowsdata

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

// Describes the current state of workflow deployment. More states may be
// added in the future.
type Workflow_State int32

const (
	// Invalid state.
	Workflow_STATE_UNSPECIFIED Workflow_State = 0
	// The workflow has been deployed successfully and is serving.
	Workflow_ACTIVE Workflow_State = 1
)

// Enum value maps for Workflow_State.
var (
	Workflow_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "ACTIVE",
	}
	Workflow_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"ACTIVE":            1,
	}
)

func (x Workflow_State) Enum() *Workflow_State {
	p := new(Workflow_State)
	*p = x
	return p
}

func (x Workflow_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Workflow_State) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_workflows_v1_data_proto_enumTypes[0].Descriptor()
}

func (Workflow_State) Type() protoreflect.EnumType {
	return &file_cloud_workflows_v1_data_proto_enumTypes[0]
}

func (x Workflow_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Workflow_State.Descriptor instead.
func (Workflow_State) EnumDescriptor() ([]byte, []int) {
	return file_cloud_workflows_v1_data_proto_rawDescGZIP(), []int{0, 0}
}

// Workflow program to be executed by Workflows.
type Workflow struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The resource name of the workflow.
	// Format: projects/{project}/locations/{location}/workflows/{workflow}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the workflow provided by the user.
	// Must be at most 1000 unicode characters long.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. State of the workflow deployment.
	State Workflow_State `protobuf:"varint,3,opt,name=state,proto3,enum=google.events.cloud.workflows.v1.Workflow_State" json:"state,omitempty"`
	// Output only. The revision of the workflow.
	// A new revision of a workflow is created as a result of updating the
	// following properties of a workflow:
	//
	// - [Service account][google.cloud.workflows.v1.Workflow.service_account]
	// - [Workflow code to be
	// executed][google.cloud.workflows.v1.Workflow.source_contents]
	//
	// The format is "000001-a4d", where the first 6 characters define
	// the zero-padded revision ordinal number. They are followed by a hyphen and
	// 3 hexadecimal random characters.
	RevisionId string `protobuf:"bytes,4,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty"`
	// Output only. The timestamp of when the workflow was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The last update timestamp of the workflow.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. The timestamp that the latest revision of the workflow
	// was created.
	RevisionCreateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=revision_create_time,json=revisionCreateTime,proto3" json:"revision_create_time,omitempty"`
	// Labels associated with this workflow.
	// Labels can contain at most 64 entries. Keys and values can be no longer
	// than 63 characters and can only contain lowercase letters, numeric
	// characters, underscores and dashes. Label keys must start with a letter.
	// International characters are allowed.
	Labels map[string]string `protobuf:"bytes,8,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// The service account associated with the latest workflow version.
	// This service account represents the identity of the workflow and determines
	// what permissions the workflow has.
	// Format: projects/{project}/serviceAccounts/{account} or {account}
	//
	// Using `-` as a wildcard for the `{project}` or not providing one at all
	// will infer the project from the account. The `{account}` value can be the
	// `email` address or the `unique_id` of the service account.
	//
	// If not provided, workflow will use the project's default service account.
	// Modifying this field for an existing workflow results in a new workflow
	// revision.
	ServiceAccount string `protobuf:"bytes,9,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
	// Required. Location of the workflow source code.
	// Modifying this field for an existing workflow results in a new workflow
	// revision.
	//
	// Types that are valid to be assigned to SourceCode:
	//
	//	*Workflow_SourceContents
	SourceCode    isWorkflow_SourceCode `protobuf_oneof:"source_code"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Workflow) Reset() {
	*x = Workflow{}
	mi := &file_cloud_workflows_v1_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Workflow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workflow) ProtoMessage() {}

func (x *Workflow) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_workflows_v1_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workflow.ProtoReflect.Descriptor instead.
func (*Workflow) Descriptor() ([]byte, []int) {
	return file_cloud_workflows_v1_data_proto_rawDescGZIP(), []int{0}
}

func (x *Workflow) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Workflow) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Workflow) GetState() Workflow_State {
	if x != nil {
		return x.State
	}
	return Workflow_STATE_UNSPECIFIED
}

func (x *Workflow) GetRevisionId() string {
	if x != nil {
		return x.RevisionId
	}
	return ""
}

func (x *Workflow) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Workflow) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Workflow) GetRevisionCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RevisionCreateTime
	}
	return nil
}

func (x *Workflow) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Workflow) GetServiceAccount() string {
	if x != nil {
		return x.ServiceAccount
	}
	return ""
}

func (x *Workflow) GetSourceCode() isWorkflow_SourceCode {
	if x != nil {
		return x.SourceCode
	}
	return nil
}

func (x *Workflow) GetSourceContents() string {
	if x != nil {
		if x, ok := x.SourceCode.(*Workflow_SourceContents); ok {
			return x.SourceContents
		}
	}
	return ""
}

type isWorkflow_SourceCode interface {
	isWorkflow_SourceCode()
}

type Workflow_SourceContents struct {
	// Workflow code to be executed. The size limit is 128KB.
	SourceContents string `protobuf:"bytes,10,opt,name=source_contents,json=sourceContents,proto3,oneof"`
}

func (*Workflow_SourceContents) isWorkflow_SourceCode() {}

// The data within all Workflow events.
type WorkflowEventData struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Optional. The Workflow event payload. Unset for deletion events.
	Payload       *Workflow `protobuf:"bytes,1,opt,name=payload,proto3,oneof" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkflowEventData) Reset() {
	*x = WorkflowEventData{}
	mi := &file_cloud_workflows_v1_data_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkflowEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowEventData) ProtoMessage() {}

func (x *WorkflowEventData) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_workflows_v1_data_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowEventData.ProtoReflect.Descriptor instead.
func (*WorkflowEventData) Descriptor() ([]byte, []int) {
	return file_cloud_workflows_v1_data_proto_rawDescGZIP(), []int{1}
}

func (x *WorkflowEventData) GetPayload() *Workflow {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_cloud_workflows_v1_data_proto protoreflect.FileDescriptor

const file_cloud_workflows_v1_data_proto_rawDesc = "" +
	"\n" +
	"\x1dcloud/workflows/v1/data.proto\x12 google.events.cloud.workflows.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"\x8b\x05\n" +
	"\bWorkflow\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12F\n" +
	"\x05state\x18\x03 \x01(\x0e20.google.events.cloud.workflows.v1.Workflow.StateR\x05state\x12\x1f\n" +
	"\vrevision_id\x18\x04 \x01(\tR\n" +
	"revisionId\x12;\n" +
	"\vcreate_time\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"createTime\x12;\n" +
	"\vupdate_time\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"updateTime\x12L\n" +
	"\x14revision_create_time\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\x12revisionCreateTime\x12N\n" +
	"\x06labels\x18\b \x03(\v26.google.events.cloud.workflows.v1.Workflow.LabelsEntryR\x06labels\x12'\n" +
	"\x0fservice_account\x18\t \x01(\tR\x0eserviceAccount\x12)\n" +
	"\x0fsource_contents\x18\n" +
	" \x01(\tH\x00R\x0esourceContents\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"*\n" +
	"\x05State\x12\x15\n" +
	"\x11STATE_UNSPECIFIED\x10\x00\x12\n" +
	"\n" +
	"\x06ACTIVE\x10\x01B\r\n" +
	"\vsource_code\"j\n" +
	"\x11WorkflowEventData\x12I\n" +
	"\apayload\x18\x01 \x01(\v2*.google.events.cloud.workflows.v1.WorkflowH\x00R\apayload\x88\x01\x01B\n" +
	"\n" +
	"\b_payloadBv\xaa\x02)Google.Events.Protobuf.Cloud.Workflows.V1\xca\x02 Google\\Events\\Cloud\\Workflows\\V1\xea\x02$Google::Events::Cloud::Workflows::V1b\x06proto3"

var (
	file_cloud_workflows_v1_data_proto_rawDescOnce sync.Once
	file_cloud_workflows_v1_data_proto_rawDescData []byte
)

func file_cloud_workflows_v1_data_proto_rawDescGZIP() []byte {
	file_cloud_workflows_v1_data_proto_rawDescOnce.Do(func() {
		file_cloud_workflows_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_cloud_workflows_v1_data_proto_rawDesc), len(file_cloud_workflows_v1_data_proto_rawDesc)))
	})
	return file_cloud_workflows_v1_data_proto_rawDescData
}

var file_cloud_workflows_v1_data_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_workflows_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cloud_workflows_v1_data_proto_goTypes = []any{
	(Workflow_State)(0),           // 0: google.events.cloud.workflows.v1.Workflow.State
	(*Workflow)(nil),              // 1: google.events.cloud.workflows.v1.Workflow
	(*WorkflowEventData)(nil),     // 2: google.events.cloud.workflows.v1.WorkflowEventData
	nil,                           // 3: google.events.cloud.workflows.v1.Workflow.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_cloud_workflows_v1_data_proto_depIdxs = []int32{
	0, // 0: google.events.cloud.workflows.v1.Workflow.state:type_name -> google.events.cloud.workflows.v1.Workflow.State
	4, // 1: google.events.cloud.workflows.v1.Workflow.create_time:type_name -> google.protobuf.Timestamp
	4, // 2: google.events.cloud.workflows.v1.Workflow.update_time:type_name -> google.protobuf.Timestamp
	4, // 3: google.events.cloud.workflows.v1.Workflow.revision_create_time:type_name -> google.protobuf.Timestamp
	3, // 4: google.events.cloud.workflows.v1.Workflow.labels:type_name -> google.events.cloud.workflows.v1.Workflow.LabelsEntry
	1, // 5: google.events.cloud.workflows.v1.WorkflowEventData.payload:type_name -> google.events.cloud.workflows.v1.Workflow
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cloud_workflows_v1_data_proto_init() }
func file_cloud_workflows_v1_data_proto_init() {
	if File_cloud_workflows_v1_data_proto != nil {
		return
	}
	file_cloud_workflows_v1_data_proto_msgTypes[0].OneofWrappers = []any{
		(*Workflow_SourceContents)(nil),
	}
	file_cloud_workflows_v1_data_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cloud_workflows_v1_data_proto_rawDesc), len(file_cloud_workflows_v1_data_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_workflows_v1_data_proto_goTypes,
		DependencyIndexes: file_cloud_workflows_v1_data_proto_depIdxs,
		EnumInfos:         file_cloud_workflows_v1_data_proto_enumTypes,
		MessageInfos:      file_cloud_workflows_v1_data_proto_msgTypes,
	}.Build()
	File_cloud_workflows_v1_data_proto = out.File
	file_cloud_workflows_v1_data_proto_goTypes = nil
	file_cloud_workflows_v1_data_proto_depIdxs = nil
}
