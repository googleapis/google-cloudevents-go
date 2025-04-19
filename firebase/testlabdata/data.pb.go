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
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.6
// source: firebase/testlab/v1/data.proto

package testlabdata

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
	state protoimpl.MessageState `protogen:"open.v1"`
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
	TestMatrixId  string `protobuf:"bytes,7,opt,name=test_matrix_id,json=testMatrixId,proto3" json:"test_matrix_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TestMatrixEventData) Reset() {
	*x = TestMatrixEventData{}
	mi := &file_firebase_testlab_v1_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestMatrixEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMatrixEventData) ProtoMessage() {}

func (x *TestMatrixEventData) ProtoReflect() protoreflect.Message {
	mi := &file_firebase_testlab_v1_data_proto_msgTypes[0]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// Client name, such as "gcloud".
	Client string `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
	// Map of detailed information about the client.
	Details       map[string]string `protobuf:"bytes,2,rep,name=details,proto3" json:"details,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientInfo) Reset() {
	*x = ClientInfo{}
	mi := &file_firebase_testlab_v1_data_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInfo) ProtoMessage() {}

func (x *ClientInfo) ProtoReflect() protoreflect.Message {
	mi := &file_firebase_testlab_v1_data_proto_msgTypes[1]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
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
	GcsPath       string `protobuf:"bytes,4,opt,name=gcs_path,json=gcsPath,proto3" json:"gcs_path,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResultStorage) Reset() {
	*x = ResultStorage{}
	mi := &file_firebase_testlab_v1_data_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResultStorage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultStorage) ProtoMessage() {}

func (x *ResultStorage) ProtoReflect() protoreflect.Message {
	mi := &file_firebase_testlab_v1_data_proto_msgTypes[2]
	if x != nil {
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

const file_firebase_testlab_v1_data_proto_rawDesc = "" +
	"\n" +
	"\x1efirebase/testlab/v1/data.proto\x12!google.events.firebase.testlab.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"\xf7\x03\n" +
	"\x13TestMatrixEventData\x12;\n" +
	"\vcreate_time\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"createTime\x12B\n" +
	"\x05state\x18\x02 \x01(\x0e2,.google.events.firebase.testlab.v1.TestStateR\x05state\x124\n" +
	"\x16invalid_matrix_details\x18\x03 \x01(\tR\x14invalidMatrixDetails\x12Z\n" +
	"\x0foutcome_summary\x18\x04 \x01(\x0e21.google.events.firebase.testlab.v1.OutcomeSummaryR\x0eoutcomeSummary\x12W\n" +
	"\x0eresult_storage\x18\x05 \x01(\v20.google.events.firebase.testlab.v1.ResultStorageR\rresultStorage\x12N\n" +
	"\vclient_info\x18\x06 \x01(\v2-.google.events.firebase.testlab.v1.ClientInfoR\n" +
	"clientInfo\x12$\n" +
	"\x0etest_matrix_id\x18\a \x01(\tR\ftestMatrixId\"\xb6\x01\n" +
	"\n" +
	"ClientInfo\x12\x16\n" +
	"\x06client\x18\x01 \x01(\tR\x06client\x12T\n" +
	"\adetails\x18\x02 \x03(\v2:.google.events.firebase.testlab.v1.ClientInfo.DetailsEntryR\adetails\x1a:\n" +
	"\fDetailsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\xb3\x01\n" +
	"\rResultStorage\x120\n" +
	"\x14tool_results_history\x18\x01 \x01(\tR\x12toolResultsHistory\x124\n" +
	"\x16tool_results_execution\x18\x02 \x01(\tR\x14toolResultsExecution\x12\x1f\n" +
	"\vresults_uri\x18\x03 \x01(\tR\n" +
	"resultsUri\x12\x19\n" +
	"\bgcs_path\x18\x04 \x01(\tR\agcsPath*j\n" +
	"\tTestState\x12\x1a\n" +
	"\x16TEST_STATE_UNSPECIFIED\x10\x00\x12\x0e\n" +
	"\n" +
	"VALIDATING\x10\x01\x12\v\n" +
	"\aPENDING\x10\x02\x12\f\n" +
	"\bFINISHED\x10\x03\x12\t\n" +
	"\x05ERROR\x10\x04\x12\v\n" +
	"\aINVALID\x10\x05*j\n" +
	"\x0eOutcomeSummary\x12\x1f\n" +
	"\x1bOUTCOME_SUMMARY_UNSPECIFIED\x10\x00\x12\v\n" +
	"\aSUCCESS\x10\x01\x12\v\n" +
	"\aFAILURE\x10\x02\x12\x10\n" +
	"\fINCONCLUSIVE\x10\x03\x12\v\n" +
	"\aSKIPPED\x10\x04B\xad\x01\n" +
	"%com.google.events.firebase.testlab.v1B\tDataProtoP\x01\xaa\x02*Google.Events.Protobuf.Firebase.TestLab.V1\xca\x02!Google\\Events\\Firebase\\TestLab\\V1\xea\x02%Google::Events::Firebase::TestLab::V1b\x06proto3"

var (
	file_firebase_testlab_v1_data_proto_rawDescOnce sync.Once
	file_firebase_testlab_v1_data_proto_rawDescData []byte
)

func file_firebase_testlab_v1_data_proto_rawDescGZIP() []byte {
	file_firebase_testlab_v1_data_proto_rawDescOnce.Do(func() {
		file_firebase_testlab_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_firebase_testlab_v1_data_proto_rawDesc), len(file_firebase_testlab_v1_data_proto_rawDesc)))
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_firebase_testlab_v1_data_proto_rawDesc), len(file_firebase_testlab_v1_data_proto_rawDesc)),
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
	file_firebase_testlab_v1_data_proto_goTypes = nil
	file_firebase_testlab_v1_data_proto_depIdxs = nil
}
