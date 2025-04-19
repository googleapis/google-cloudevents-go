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
// source: firebase/firebasealerts/v1/data.proto

package firebasealertsdata

import (
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

// The data within all Firebase Alerts.
type AlertData struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Time that the event has created
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Time that the event has ended. Optional, only present for alertsthat are
	// ongoing
	EndTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Payload of the event, which includes the details of the specific alert.
	// It's a map of keys of String type and values of various types
	Payload       *structpb.Struct `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlertData) Reset() {
	*x = AlertData{}
	mi := &file_firebase_firebasealerts_v1_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertData) ProtoMessage() {}

func (x *AlertData) ProtoReflect() protoreflect.Message {
	mi := &file_firebase_firebasealerts_v1_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertData.ProtoReflect.Descriptor instead.
func (*AlertData) Descriptor() ([]byte, []int) {
	return file_firebase_firebasealerts_v1_data_proto_rawDescGZIP(), []int{0}
}

func (x *AlertData) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *AlertData) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *AlertData) GetPayload() *structpb.Struct {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_firebase_firebasealerts_v1_data_proto protoreflect.FileDescriptor

const file_firebase_firebasealerts_v1_data_proto_rawDesc = "" +
	"\n" +
	"%firebase/firebasealerts/v1/data.proto\x12(google.events.firebase.firebasealerts.v1\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xb2\x01\n" +
	"\tAlertData\x12;\n" +
	"\vcreate_time\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"createTime\x125\n" +
	"\bend_time\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\aendTime\x121\n" +
	"\apayload\x18\x03 \x01(\v2\x17.google.protobuf.StructR\apayloadB\x8e\x01\xaa\x021Google.Events.Protobuf.Firebase.FirebaseAlerts.V1\xca\x02(Google\\Events\\Firebase\\FirebaseAlerts\\V1\xea\x02,Google::Events::Firebase::FirebaseAlerts::V1b\x06proto3"

var (
	file_firebase_firebasealerts_v1_data_proto_rawDescOnce sync.Once
	file_firebase_firebasealerts_v1_data_proto_rawDescData []byte
)

func file_firebase_firebasealerts_v1_data_proto_rawDescGZIP() []byte {
	file_firebase_firebasealerts_v1_data_proto_rawDescOnce.Do(func() {
		file_firebase_firebasealerts_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_firebase_firebasealerts_v1_data_proto_rawDesc), len(file_firebase_firebasealerts_v1_data_proto_rawDesc)))
	})
	return file_firebase_firebasealerts_v1_data_proto_rawDescData
}

var file_firebase_firebasealerts_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_firebase_firebasealerts_v1_data_proto_goTypes = []any{
	(*AlertData)(nil),             // 0: google.events.firebase.firebasealerts.v1.AlertData
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*structpb.Struct)(nil),       // 2: google.protobuf.Struct
}
var file_firebase_firebasealerts_v1_data_proto_depIdxs = []int32{
	1, // 0: google.events.firebase.firebasealerts.v1.AlertData.create_time:type_name -> google.protobuf.Timestamp
	1, // 1: google.events.firebase.firebasealerts.v1.AlertData.end_time:type_name -> google.protobuf.Timestamp
	2, // 2: google.events.firebase.firebasealerts.v1.AlertData.payload:type_name -> google.protobuf.Struct
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_firebase_firebasealerts_v1_data_proto_init() }
func file_firebase_firebasealerts_v1_data_proto_init() {
	if File_firebase_firebasealerts_v1_data_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_firebase_firebasealerts_v1_data_proto_rawDesc), len(file_firebase_firebasealerts_v1_data_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_firebase_firebasealerts_v1_data_proto_goTypes,
		DependencyIndexes: file_firebase_firebasealerts_v1_data_proto_depIdxs,
		MessageInfos:      file_firebase_firebasealerts_v1_data_proto_msgTypes,
	}.Build()
	File_firebase_firebasealerts_v1_data_proto = out.File
	file_firebase_firebasealerts_v1_data_proto_goTypes = nil
	file_firebase_firebasealerts_v1_data_proto_depIdxs = nil
}
