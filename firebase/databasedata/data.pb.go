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
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.6
// source: firebase/database/v1/data.proto

package databasedata

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The data within all Firebase Real Time Database reference events.
type ReferenceEventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The original data for the reference.
	Data *structpb.Value `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// The change in the reference data.
	Delta *structpb.Value `protobuf:"bytes,2,opt,name=delta,proto3" json:"delta,omitempty"`
}

func (x *ReferenceEventData) Reset() {
	*x = ReferenceEventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_firebase_database_v1_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReferenceEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReferenceEventData) ProtoMessage() {}

func (x *ReferenceEventData) ProtoReflect() protoreflect.Message {
	mi := &file_firebase_database_v1_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReferenceEventData.ProtoReflect.Descriptor instead.
func (*ReferenceEventData) Descriptor() ([]byte, []int) {
	return file_firebase_database_v1_data_proto_rawDescGZIP(), []int{0}
}

func (x *ReferenceEventData) GetData() *structpb.Value {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ReferenceEventData) GetDelta() *structpb.Value {
	if x != nil {
		return x.Delta
	}
	return nil
}

var File_firebase_database_v1_data_proto protoreflect.FileDescriptor

var file_firebase_database_v1_data_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x12, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x64, 0x65,
	0x6c, 0x74, 0x61, 0x42, 0x7c, 0xaa, 0x02, 0x2b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x5c, 0x46, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x5c, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x3a, 0x46, 0x69, 0x72, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x3a, 0x3a, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_firebase_database_v1_data_proto_rawDescOnce sync.Once
	file_firebase_database_v1_data_proto_rawDescData = file_firebase_database_v1_data_proto_rawDesc
)

func file_firebase_database_v1_data_proto_rawDescGZIP() []byte {
	file_firebase_database_v1_data_proto_rawDescOnce.Do(func() {
		file_firebase_database_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_firebase_database_v1_data_proto_rawDescData)
	})
	return file_firebase_database_v1_data_proto_rawDescData
}

var file_firebase_database_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_firebase_database_v1_data_proto_goTypes = []interface{}{
	(*ReferenceEventData)(nil), // 0: google.events.firebase.database.v1.ReferenceEventData
	(*structpb.Value)(nil),     // 1: google.protobuf.Value
}
var file_firebase_database_v1_data_proto_depIdxs = []int32{
	1, // 0: google.events.firebase.database.v1.ReferenceEventData.data:type_name -> google.protobuf.Value
	1, // 1: google.events.firebase.database.v1.ReferenceEventData.delta:type_name -> google.protobuf.Value
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_firebase_database_v1_data_proto_init() }
func file_firebase_database_v1_data_proto_init() {
	if File_firebase_database_v1_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_firebase_database_v1_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReferenceEventData); i {
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
			RawDescriptor: file_firebase_database_v1_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_firebase_database_v1_data_proto_goTypes,
		DependencyIndexes: file_firebase_database_v1_data_proto_depIdxs,
		MessageInfos:      file_firebase_database_v1_data_proto_msgTypes,
	}.Build()
	File_firebase_database_v1_data_proto = out.File
	file_firebase_database_v1_data_proto_rawDesc = nil
	file_firebase_database_v1_data_proto_goTypes = nil
	file_firebase_database_v1_data_proto_depIdxs = nil
}
