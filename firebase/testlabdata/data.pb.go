// Copyright 2020 Google LLC
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.6
// source: firebase/testlab/v1/data.proto

package testlabdata

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

// Possible test states for a test matrix.
type TestState int32

const (
	// The default value. This value is used if the state is omitted.
	TestState_TEST_STATE_UNSPECIFIED TestState = 0
	// The test matrix is being validated.
	TestState_VALIDATING TestState = 1
	// The test matrix is waiting for resources to become available.
	TestState_PENDING TestState = 2
	// The test matrix has completed normally.
	TestState_FINISHED TestState = 3
	// The test matrix has completed because of an infrastructure failure.
	TestState_ERROR TestState = 4
	// The test matrix was not run because the provided inputs are not valid.
	TestState_INVALID TestState = 5
)

// Enum value maps for TestState.
var (
	TestState_name = map[int32]string{
		0: "TEST_STATE_UNSPECIFIED",
		1: "VALIDATING",
		2: "PENDING",
		3: "FINISHED",
		4: "ERROR",
		5: "INVALID",
	}
	TestState_value = map[string]int32{
		"TEST_STATE_UNSPECIFIED": 0,
		"VALIDATING":             1,
		"PENDING":                2,
		"FINISHED":               3,
		"ERROR":                  4,
		"INVALID":                5,
	}
)

func (x TestState) Enum() *TestState {
	p := new(TestState)
	*p = x
	return p
}

func (x TestState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestState) Descriptor() protoreflect.EnumDescriptor {
	return file_firebase_testlab_v1_data_proto_enumTypes[0].Descriptor()
}

func (TestState) Type() protoreflect.EnumType {
	return &file_firebase_testlab_v1_data_proto_enumTypes[0]
}

func (x TestState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestState.Descriptor instead.
func (TestState) EnumDescriptor() ([]byte, []int) {
	return file_firebase_testlab_v1_data_proto_rawDescGZIP(), []int{0}
}

// Outcome summary for a finished test matrix.
type OutcomeSummary int32

const (
	// The default value. This value is used if the state is omitted.
	OutcomeSummary_OUTCOME_SUMMARY_UNSPECIFIED OutcomeSummary = 0
	// The test matrix run was successful, for instance:
	// - All test cases passed.
	// - No crash of the application under test was detected.
	OutcomeSummary_SUCCESS OutcomeSummary = 1
	// A run failed, for instance:
	// - One or more test case failed.
	// - A test timed out.
	// - The application under test crashed.
	OutcomeSummary_FAILURE OutcomeSummary = 2
	// Something unexpected happened. The test run should still be considered
	// unsuccessful but this is likely a transient problem and re-running the
	// test might be successful.
	OutcomeSummary_INCONCLUSIVE OutcomeSummary = 3
	// All tests were skipped.
	OutcomeSummary_SKIPPED OutcomeSummary = 4
)

// Enum value maps for OutcomeSummary.
var (
	OutcomeSummary_name = map[int32]string{
		0: "OUTCOME_SUMMARY_UNSPECIFIED",
		1: "SUCCESS",
		2: "FAILURE",
		3: "INCONCLUSIVE",
		4: "SKIPPED",
	}
	OutcomeSummary_value = map[string]int32{
		"OUTCOME_SUMMARY_UNSPECIFIED": 0,
		"SUCCESS":                     1,
		"FAILURE":                     2,
		"INCONCLUSIVE":                3,
		"SKIPPED":                     4,
	}
)

func (x OutcomeSummary) Enum() *OutcomeSummary {
	p := new(OutcomeSummary)
	*p = x
	return p
}

func (x OutcomeSummary) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OutcomeSummary) Descriptor() protoreflect.EnumDescriptor {
	return file_firebase_testlab_v1_data_proto_enumTypes[1].Descriptor()
}

func (OutcomeSummary) Type() protoreflect.EnumType {
	return &file_firebase_testlab_v1_data_proto_enumTypes[1]
}

func (x OutcomeSummary) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OutcomeSummary.Descriptor instead.
func (OutcomeSummary) EnumDescriptor() ([]byte, []int) {
	return file_firebase_testlab_v1_data_proto_rawDescGZIP(), []int{1}
}

// The data within all Firebase test matrix events.
type TestMatrixEventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Time the test matrix was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// State of the test matrix.
	State TestState `protobuf:"varint,2,opt,name=state,proto3,enum=google.events.firebase.testlab.v1.TestState" json:"state,omitempty"`
	// Code that describes why the test matrix is considered invalid. Only set for
	// matrices in the INVALID state.
	InvalidMatrixDetails string `protobuf:"bytes,3,opt,name=invalid_matrix_details,json=invalidMatrixDetails,proto3" json:"invalid_matrix_details,omitempty"`
	// Outcome summary of the test matrix.
	OutcomeSummary OutcomeSummary `protobuf:"varint,4,opt,name=outcome_summary,json=outcomeSummary,proto3,enum=google.events.firebase.testlab.v1.OutcomeSummary" json:"outcome_summary,omitempty"`
	// Locations where test results are stored.
	ResultStorage *ResultStorage `protobuf:"bytes,5,opt,name=result_storage,json=resultStorage,proto3" json:"result_storage,omitempty"`
	// Information provided by the client that created the test matrix.
	ClientInfo *ClientInfo `protobuf:"bytes,6,opt,name=client_info,json=clientInfo,proto3" json:"client_info,omitempty"`
	// ID of the test matrix this event belongs to.
	TestMatrixId string `protobuf:"bytes,7,opt,name=test_matrix_id,json=testMatrixId,proto3" json:"test_matrix_id,omitempty"`
}

func (x *TestMatrixEventData) Reset() {
	*x = TestMatrixEventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_firebase_testlab_v1_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMatrixEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMatrixEventData) ProtoMessage() {}

func (x *TestMatrixEventData) ProtoReflect() protoreflect.Message {
	mi := &file_firebase_testlab_v1_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMatrixEventData.ProtoReflect.Descriptor instead.
func (*TestMatrixEventData) Descriptor() ([]byte, []int) {
	return file_firebase_testlab_v1_data_proto_rawDescGZIP(), []int{0}
}

func (x *TestMatrixEventData) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *TestMatrixEventData) GetState() TestState {
	if x != nil {
		return x.State
	}
	return TestState_TEST_STATE_UNSPECIFIED
}

func (x *TestMatrixEventData) GetInvalidMatrixDetails() string {
	if x != nil {
		return x.InvalidMatrixDetails
	}
	return ""
}

func (x *TestMatrixEventData) GetOutcomeSummary() OutcomeSummary {
	if x != nil {
		return x.OutcomeSummary
	}
	return OutcomeSummary_OUTCOME_SUMMARY_UNSPECIFIED
}

func (x *TestMatrixEventData) GetResultStorage() *ResultStorage {
	if x != nil {
		return x.ResultStorage
	}
	return nil
}

func (x *TestMatrixEventData) GetClientInfo() *ClientInfo {
	if x != nil {
		return x.ClientInfo
	}
	return nil
}

func (x *TestMatrixEventData) GetTestMatrixId() string {
	if x != nil {
		return x.TestMatrixId
	}
	return ""
}

// Information about the client which invoked the test.
type ClientInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Client name, such as "gcloud".
	Client string `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
	// Map of detailed information about the client.
	Details map[string]string `protobuf:"bytes,2,rep,name=details,proto3" json:"details,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ClientInfo) Reset() {
	*x = ClientInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_firebase_testlab_v1_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInfo) ProtoMessage() {}

func (x *ClientInfo) ProtoReflect() protoreflect.Message {
	mi := &file_firebase_testlab_v1_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInfo.ProtoReflect.Descriptor instead.
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return file_firebase_testlab_v1_data_proto_rawDescGZIP(), []int{1}
}

func (x *ClientInfo) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

func (x *ClientInfo) GetDetails() map[string]string {
	if x != nil {
		return x.Details
	}
	return nil
}

// Locations where test results are stored.
type ResultStorage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tool Results history resource containing test results. Format is
	// `projects/{project_id}/histories/{history_id}`.
	// See https://firebase.google.com/docs/test-lab/reference/toolresults/rest
	// for more information.
	ToolResultsHistory string `protobuf:"bytes,1,opt,name=tool_results_history,json=toolResultsHistory,proto3" json:"tool_results_history,omitempty"`
	// Tool Results execution resource containing test results. Format is
	// `projects/{project_id}/histories/{history_id}/executions/{execution_id}`.
	// Optional, can be omitted in erroneous test states.
	// See https://firebase.google.com/docs/test-lab/reference/toolresults/rest
	// for more information.
	ToolResultsExecution string `protobuf:"bytes,2,opt,name=tool_results_execution,json=toolResultsExecution,proto3" json:"tool_results_execution,omitempty"`
	// URI to the test results in the Firebase Web Console.
	ResultsUri string `protobuf:"bytes,3,opt,name=results_uri,json=resultsUri,proto3" json:"results_uri,omitempty"`
	// Location in Google Cloud Storage where test results are written to.
	// In the form "gs://bucket/path/to/somewhere".
	GcsPath string `protobuf:"bytes,4,opt,name=gcs_path,json=gcsPath,proto3" json:"gcs_path,omitempty"`
}

func (x *ResultStorage) Reset() {
	*x = ResultStorage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_firebase_testlab_v1_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultStorage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultStorage) ProtoMessage() {}

func (x *ResultStorage) ProtoReflect() protoreflect.Message {
	mi := &file_firebase_testlab_v1_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultStorage.ProtoReflect.Descriptor instead.
func (*ResultStorage) Descriptor() ([]byte, []int) {
	return file_firebase_testlab_v1_data_proto_rawDescGZIP(), []int{2}
}

func (x *ResultStorage) GetToolResultsHistory() string {
	if x != nil {
		return x.ToolResultsHistory
	}
	return ""
}

func (x *ResultStorage) GetToolResultsExecution() string {
	if x != nil {
		return x.ToolResultsExecution
	}
	return ""
}

func (x *ResultStorage) GetResultsUri() string {
	if x != nil {
		return x.ResultsUri
	}
	return ""
}

func (x *ResultStorage) GetGcsPath() string {
	if x != nil {
		return x.GcsPath
	}
	return ""
}

var File_firebase_testlab_v1_data_proto protoreflect.FileDescriptor

var file_firebase_testlab_v1_data_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x6c,
	0x61, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x03, 0x0a, 0x13, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x74,
	0x72, 0x69, 0x78, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a,
	0x16, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x5f,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x69,
	0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x5a, 0x0a, 0x0f, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x66, 0x69, 0x72,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52,
	0x0e, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12,
	0x57, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x66, 0x69,
	0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x65, 0x73, 0x74,
	0x5f, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x49, 0x64, 0x22, 0xb6,
	0x01, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x54, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb3, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x6f, 0x6f,
	0x6c, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x74, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x34, 0x0a, 0x16, 0x74,
	0x6f, 0x6f, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x5f, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x74, 0x6f, 0x6f,
	0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x5f, 0x75, 0x72, 0x69,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x55,
	0x72, 0x69, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x63, 0x73, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x63, 0x73, 0x50, 0x61, 0x74, 0x68, 0x2a, 0x6a, 0x0a,
	0x09, 0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x45,
	0x53, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10,
	0x03, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x05, 0x2a, 0x6a, 0x0a, 0x0e, 0x4f, 0x75, 0x74,
	0x63, 0x6f, 0x6d, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x1b, 0x4f,
	0x55, 0x54, 0x43, 0x4f, 0x4d, 0x45, 0x5f, 0x53, 0x55, 0x4d, 0x4d, 0x41, 0x52, 0x59, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49,
	0x4c, 0x55, 0x52, 0x45, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x43, 0x4f, 0x4e, 0x43,
	0x4c, 0x55, 0x53, 0x49, 0x56, 0x45, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x4b, 0x49, 0x50,
	0x50, 0x45, 0x44, 0x10, 0x04, 0x42, 0xad, 0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x66, 0x69, 0x72, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x76, 0x31, 0x42,
	0x09, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xaa, 0x02, 0x2a, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x46, 0x69, 0x72, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x5c, 0x54, 0x65, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x25,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x3a,
	0x46, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x3a, 0x3a, 0x54, 0x65, 0x73, 0x74, 0x4c, 0x61,
	0x62, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_firebase_testlab_v1_data_proto_rawDescOnce sync.Once
	file_firebase_testlab_v1_data_proto_rawDescData = file_firebase_testlab_v1_data_proto_rawDesc
)

func file_firebase_testlab_v1_data_proto_rawDescGZIP() []byte {
	file_firebase_testlab_v1_data_proto_rawDescOnce.Do(func() {
		file_firebase_testlab_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_firebase_testlab_v1_data_proto_rawDescData)
	})
	return file_firebase_testlab_v1_data_proto_rawDescData
}

var file_firebase_testlab_v1_data_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_firebase_testlab_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_firebase_testlab_v1_data_proto_goTypes = []any{
	(TestState)(0),                // 0: google.events.firebase.testlab.v1.TestState
	(OutcomeSummary)(0),           // 1: google.events.firebase.testlab.v1.OutcomeSummary
	(*TestMatrixEventData)(nil),   // 2: google.events.firebase.testlab.v1.TestMatrixEventData
	(*ClientInfo)(nil),            // 3: google.events.firebase.testlab.v1.ClientInfo
	(*ResultStorage)(nil),         // 4: google.events.firebase.testlab.v1.ResultStorage
	nil,                           // 5: google.events.firebase.testlab.v1.ClientInfo.DetailsEntry
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_firebase_testlab_v1_data_proto_depIdxs = []int32{
	6, // 0: google.events.firebase.testlab.v1.TestMatrixEventData.create_time:type_name -> google.protobuf.Timestamp
	0, // 1: google.events.firebase.testlab.v1.TestMatrixEventData.state:type_name -> google.events.firebase.testlab.v1.TestState
	1, // 2: google.events.firebase.testlab.v1.TestMatrixEventData.outcome_summary:type_name -> google.events.firebase.testlab.v1.OutcomeSummary
	4, // 3: google.events.firebase.testlab.v1.TestMatrixEventData.result_storage:type_name -> google.events.firebase.testlab.v1.ResultStorage
	3, // 4: google.events.firebase.testlab.v1.TestMatrixEventData.client_info:type_name -> google.events.firebase.testlab.v1.ClientInfo
	5, // 5: google.events.firebase.testlab.v1.ClientInfo.details:type_name -> google.events.firebase.testlab.v1.ClientInfo.DetailsEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_firebase_testlab_v1_data_proto_init() }
func file_firebase_testlab_v1_data_proto_init() {
	if File_firebase_testlab_v1_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_firebase_testlab_v1_data_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TestMatrixEventData); i {
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
		file_firebase_testlab_v1_data_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ClientInfo); i {
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
		file_firebase_testlab_v1_data_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ResultStorage); i {
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
			RawDescriptor: file_firebase_testlab_v1_data_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_firebase_testlab_v1_data_proto_goTypes,
		DependencyIndexes: file_firebase_testlab_v1_data_proto_depIdxs,
		EnumInfos:         file_firebase_testlab_v1_data_proto_enumTypes,
		MessageInfos:      file_firebase_testlab_v1_data_proto_msgTypes,
	}.Build()
	File_firebase_testlab_v1_data_proto = out.File
	file_firebase_testlab_v1_data_proto_rawDesc = nil
	file_firebase_testlab_v1_data_proto_goTypes = nil
	file_firebase_testlab_v1_data_proto_depIdxs = nil
}
