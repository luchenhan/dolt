// Copyright 2023 Dolthub, Inc.
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
// 	protoc-gen-go v1.28.1
// 	protoc        v4.26.0
// source: dolt/services/replicationapi/v1alpha1/replication.proto

package replicationapi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpdateUsersAndGrantsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The contents of the *MySQLDb instance, as seen by a Persister
	// implementation.
	SerializedContents []byte `protobuf:"bytes,1,opt,name=serialized_contents,json=serializedContents,proto3" json:"serialized_contents,omitempty"`
}

func (x *UpdateUsersAndGrantsRequest) Reset() {
	*x = UpdateUsersAndGrantsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dolt_services_replicationapi_v1alpha1_replication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUsersAndGrantsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUsersAndGrantsRequest) ProtoMessage() {}

func (x *UpdateUsersAndGrantsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dolt_services_replicationapi_v1alpha1_replication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUsersAndGrantsRequest.ProtoReflect.Descriptor instead.
func (*UpdateUsersAndGrantsRequest) Descriptor() ([]byte, []int) {
	return file_dolt_services_replicationapi_v1alpha1_replication_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateUsersAndGrantsRequest) GetSerializedContents() []byte {
	if x != nil {
		return x.SerializedContents
	}
	return nil
}

type UpdateUsersAndGrantsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateUsersAndGrantsResponse) Reset() {
	*x = UpdateUsersAndGrantsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dolt_services_replicationapi_v1alpha1_replication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUsersAndGrantsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUsersAndGrantsResponse) ProtoMessage() {}

func (x *UpdateUsersAndGrantsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dolt_services_replicationapi_v1alpha1_replication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUsersAndGrantsResponse.ProtoReflect.Descriptor instead.
func (*UpdateUsersAndGrantsResponse) Descriptor() ([]byte, []int) {
	return file_dolt_services_replicationapi_v1alpha1_replication_proto_rawDescGZIP(), []int{1}
}

type UpdateBranchControlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The serialized contents of the branch_control.Controller.
	SerializedContents []byte `protobuf:"bytes,1,opt,name=serialized_contents,json=serializedContents,proto3" json:"serialized_contents,omitempty"`
}

func (x *UpdateBranchControlRequest) Reset() {
	*x = UpdateBranchControlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dolt_services_replicationapi_v1alpha1_replication_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBranchControlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBranchControlRequest) ProtoMessage() {}

func (x *UpdateBranchControlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dolt_services_replicationapi_v1alpha1_replication_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBranchControlRequest.ProtoReflect.Descriptor instead.
func (*UpdateBranchControlRequest) Descriptor() ([]byte, []int) {
	return file_dolt_services_replicationapi_v1alpha1_replication_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateBranchControlRequest) GetSerializedContents() []byte {
	if x != nil {
		return x.SerializedContents
	}
	return nil
}

type UpdateBranchControlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBranchControlResponse) Reset() {
	*x = UpdateBranchControlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dolt_services_replicationapi_v1alpha1_replication_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBranchControlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBranchControlResponse) ProtoMessage() {}

func (x *UpdateBranchControlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dolt_services_replicationapi_v1alpha1_replication_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBranchControlResponse.ProtoReflect.Descriptor instead.
func (*UpdateBranchControlResponse) Descriptor() ([]byte, []int) {
	return file_dolt_services_replicationapi_v1alpha1_replication_proto_rawDescGZIP(), []int{3}
}

type DropDatabaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the database to be dropped.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DropDatabaseRequest) Reset() {
	*x = DropDatabaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dolt_services_replicationapi_v1alpha1_replication_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropDatabaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropDatabaseRequest) ProtoMessage() {}

func (x *DropDatabaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dolt_services_replicationapi_v1alpha1_replication_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropDatabaseRequest.ProtoReflect.Descriptor instead.
func (*DropDatabaseRequest) Descriptor() ([]byte, []int) {
	return file_dolt_services_replicationapi_v1alpha1_replication_proto_rawDescGZIP(), []int{4}
}

func (x *DropDatabaseRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DropDatabaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DropDatabaseResponse) Reset() {
	*x = DropDatabaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dolt_services_replicationapi_v1alpha1_replication_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropDatabaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropDatabaseResponse) ProtoMessage() {}

func (x *DropDatabaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dolt_services_replicationapi_v1alpha1_replication_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropDatabaseResponse.ProtoReflect.Descriptor instead.
func (*DropDatabaseResponse) Descriptor() ([]byte, []int) {
	return file_dolt_services_replicationapi_v1alpha1_replication_proto_rawDescGZIP(), []int{5}
}

var File_dolt_services_replicationapi_v1alpha1_replication_proto protoreflect.FileDescriptor

var file_dolt_services_replicationapi_v1alpha1_replication_proto_rawDesc = []byte{
	0x0a, 0x37, 0x64, 0x6f, 0x6c, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x64, 0x6f, 0x6c, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x22, 0x4e, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41,
	0x6e, 0x64, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2f, 0x0a, 0x13, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0x1e, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41,
	0x6e, 0x64, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x4d, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f,
	0x0a, 0x13, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x1d, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29,
	0x0a, 0x13, 0x44, 0x72, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x72, 0x6f,
	0x70, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0xdf, 0x03, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9f, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x6e, 0x64, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x73, 0x12, 0x42, 0x2e, 0x64, 0x6f, 0x6c, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x6e, 0x64, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x64, 0x6f, 0x6c, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x6e, 0x64, 0x47, 0x72, 0x61, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x9c, 0x01, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x12, 0x41, 0x2e, 0x64, 0x6f, 0x6c, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x64, 0x6f, 0x6c, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x87, 0x01, 0x0a, 0x0c, 0x44, 0x72,
	0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x3a, 0x2e, 0x64, 0x6f, 0x6c,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x64, 0x6f, 0x6c, 0x74, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44,
	0x72, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x5b, 0x5a, 0x59, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x6f, 0x6c, 0x74, 0x68, 0x75, 0x62, 0x2f, 0x64, 0x6f, 0x6c, 0x74, 0x2f, 0x67,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x6c, 0x74,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x3b, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dolt_services_replicationapi_v1alpha1_replication_proto_rawDescOnce sync.Once
	file_dolt_services_replicationapi_v1alpha1_replication_proto_rawDescData = file_dolt_services_replicationapi_v1alpha1_replication_proto_rawDesc
)

func file_dolt_services_replicationapi_v1alpha1_replication_proto_rawDescGZIP() []byte {
	file_dolt_services_replicationapi_v1alpha1_replication_proto_rawDescOnce.Do(func() {
		file_dolt_services_replicationapi_v1alpha1_replication_proto_rawDescData = protoimpl.X.CompressGZIP(file_dolt_services_replicationapi_v1alpha1_replication_proto_rawDescData)
	})
	return file_dolt_services_replicationapi_v1alpha1_replication_proto_rawDescData
}

var file_dolt_services_replicationapi_v1alpha1_replication_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_dolt_services_replicationapi_v1alpha1_replication_proto_goTypes = []interface{}{
	(*UpdateUsersAndGrantsRequest)(nil),  // 0: dolt.services.replicationapi.v1alpha1.UpdateUsersAndGrantsRequest
	(*UpdateUsersAndGrantsResponse)(nil), // 1: dolt.services.replicationapi.v1alpha1.UpdateUsersAndGrantsResponse
	(*UpdateBranchControlRequest)(nil),   // 2: dolt.services.replicationapi.v1alpha1.UpdateBranchControlRequest
	(*UpdateBranchControlResponse)(nil),  // 3: dolt.services.replicationapi.v1alpha1.UpdateBranchControlResponse
	(*DropDatabaseRequest)(nil),          // 4: dolt.services.replicationapi.v1alpha1.DropDatabaseRequest
	(*DropDatabaseResponse)(nil),         // 5: dolt.services.replicationapi.v1alpha1.DropDatabaseResponse
}
var file_dolt_services_replicationapi_v1alpha1_replication_proto_depIdxs = []int32{
	0, // 0: dolt.services.replicationapi.v1alpha1.ReplicationService.UpdateUsersAndGrants:input_type -> dolt.services.replicationapi.v1alpha1.UpdateUsersAndGrantsRequest
	2, // 1: dolt.services.replicationapi.v1alpha1.ReplicationService.UpdateBranchControl:input_type -> dolt.services.replicationapi.v1alpha1.UpdateBranchControlRequest
	4, // 2: dolt.services.replicationapi.v1alpha1.ReplicationService.DropDatabase:input_type -> dolt.services.replicationapi.v1alpha1.DropDatabaseRequest
	1, // 3: dolt.services.replicationapi.v1alpha1.ReplicationService.UpdateUsersAndGrants:output_type -> dolt.services.replicationapi.v1alpha1.UpdateUsersAndGrantsResponse
	3, // 4: dolt.services.replicationapi.v1alpha1.ReplicationService.UpdateBranchControl:output_type -> dolt.services.replicationapi.v1alpha1.UpdateBranchControlResponse
	5, // 5: dolt.services.replicationapi.v1alpha1.ReplicationService.DropDatabase:output_type -> dolt.services.replicationapi.v1alpha1.DropDatabaseResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dolt_services_replicationapi_v1alpha1_replication_proto_init() }
func file_dolt_services_replicationapi_v1alpha1_replication_proto_init() {
	if File_dolt_services_replicationapi_v1alpha1_replication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dolt_services_replicationapi_v1alpha1_replication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUsersAndGrantsRequest); i {
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
		file_dolt_services_replicationapi_v1alpha1_replication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUsersAndGrantsResponse); i {
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
		file_dolt_services_replicationapi_v1alpha1_replication_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBranchControlRequest); i {
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
		file_dolt_services_replicationapi_v1alpha1_replication_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBranchControlResponse); i {
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
		file_dolt_services_replicationapi_v1alpha1_replication_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropDatabaseRequest); i {
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
		file_dolt_services_replicationapi_v1alpha1_replication_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropDatabaseResponse); i {
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
			RawDescriptor: file_dolt_services_replicationapi_v1alpha1_replication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dolt_services_replicationapi_v1alpha1_replication_proto_goTypes,
		DependencyIndexes: file_dolt_services_replicationapi_v1alpha1_replication_proto_depIdxs,
		MessageInfos:      file_dolt_services_replicationapi_v1alpha1_replication_proto_msgTypes,
	}.Build()
	File_dolt_services_replicationapi_v1alpha1_replication_proto = out.File
	file_dolt_services_replicationapi_v1alpha1_replication_proto_rawDesc = nil
	file_dolt_services_replicationapi_v1alpha1_replication_proto_goTypes = nil
	file_dolt_services_replicationapi_v1alpha1_replication_proto_depIdxs = nil
}
