// Copyright 2020 Dolthub, Inc.
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

// WARNING: This file was is automatically generated. DO NOT EDIT BY HAND.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.26.0
// source: dolt/services/eventsapi/v1alpha1/event_constants.proto

package eventsapi

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

type Platform int32

const (
	Platform_PLATFORM_UNSPECIFIED Platform = 0
	Platform_LINUX                Platform = 1
	Platform_WINDOWS              Platform = 2
	Platform_DARWIN               Platform = 3
)

// Enum value maps for Platform.
var (
	Platform_name = map[int32]string{
		0: "PLATFORM_UNSPECIFIED",
		1: "LINUX",
		2: "WINDOWS",
		3: "DARWIN",
	}
	Platform_value = map[string]int32{
		"PLATFORM_UNSPECIFIED": 0,
		"LINUX":                1,
		"WINDOWS":              2,
		"DARWIN":               3,
	}
)

func (x Platform) Enum() *Platform {
	p := new(Platform)
	*p = x
	return p
}

func (x Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_dolt_services_eventsapi_v1alpha1_event_constants_proto_enumTypes[0].Descriptor()
}

func (Platform) Type() protoreflect.EnumType {
	return &file_dolt_services_eventsapi_v1alpha1_event_constants_proto_enumTypes[0]
}

func (x Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Platform.Descriptor instead.
func (Platform) EnumDescriptor() ([]byte, []int) {
	return file_dolt_services_eventsapi_v1alpha1_event_constants_proto_rawDescGZIP(), []int{0}
}

type ClientEventType int32

const (
	ClientEventType_TYPE_UNSPECIFIED                 ClientEventType = 0
	ClientEventType_INIT                             ClientEventType = 1
	ClientEventType_STATUS                           ClientEventType = 2
	ClientEventType_ADD                              ClientEventType = 3
	ClientEventType_RESET                            ClientEventType = 4
	ClientEventType_COMMIT                           ClientEventType = 5
	ClientEventType_SQL                              ClientEventType = 6
	ClientEventType_SQL_SERVER                       ClientEventType = 7
	ClientEventType_LOG                              ClientEventType = 8
	ClientEventType_DIFF                             ClientEventType = 9
	ClientEventType_MERGE                            ClientEventType = 10
	ClientEventType_BRANCH                           ClientEventType = 11
	ClientEventType_CHECKOUT                         ClientEventType = 12
	ClientEventType_REMOTE                           ClientEventType = 13
	ClientEventType_PUSH                             ClientEventType = 14
	ClientEventType_PULL                             ClientEventType = 15
	ClientEventType_FETCH                            ClientEventType = 16
	ClientEventType_CLONE                            ClientEventType = 17
	ClientEventType_LOGIN                            ClientEventType = 18
	ClientEventType_VERSION                          ClientEventType = 19
	ClientEventType_CONFIG                           ClientEventType = 20
	ClientEventType_LS                               ClientEventType = 21
	ClientEventType_SCHEMA                           ClientEventType = 22
	ClientEventType_TABLE_IMPORT                     ClientEventType = 23
	ClientEventType_TABLE_EXPORT                     ClientEventType = 24
	ClientEventType_TABLE_CREATE                     ClientEventType = 25
	ClientEventType_TABLE_RM                         ClientEventType = 26
	ClientEventType_TABLE_MV                         ClientEventType = 27
	ClientEventType_TABLE_CP                         ClientEventType = 28
	ClientEventType_TABLE_SELECT                     ClientEventType = 29
	ClientEventType_TABLE_PUT_ROW                    ClientEventType = 30
	ClientEventType_TABLE_RM_ROW                     ClientEventType = 31
	ClientEventType_CREDS_NEW                        ClientEventType = 32
	ClientEventType_CREDS_RM                         ClientEventType = 33
	ClientEventType_CREDS_LS                         ClientEventType = 34
	ClientEventType_CONF_CAT                         ClientEventType = 35
	ClientEventType_CONF_RESOLVE                     ClientEventType = 36
	ClientEventType_REMOTEAPI_GET_REPO_METADATA      ClientEventType = 37
	ClientEventType_REMOTEAPI_HAS_CHUNKS             ClientEventType = 38
	ClientEventType_REMOTEAPI_GET_DOWNLOAD_LOCATIONS ClientEventType = 39
	ClientEventType_REMOTEAPI_GET_UPLOAD_LOCATIONS   ClientEventType = 40
	ClientEventType_REMOTEAPI_REBASE                 ClientEventType = 41
	ClientEventType_REMOTEAPI_ROOT                   ClientEventType = 42
	ClientEventType_REMOTEAPI_COMMIT                 ClientEventType = 43
	ClientEventType_REMOTEAPI_LIST_TABLE_FILES       ClientEventType = 44
	ClientEventType_BLAME                            ClientEventType = 45
	ClientEventType_CREDS_CHECK                      ClientEventType = 46
	ClientEventType_CREDS_USE                        ClientEventType = 47
	ClientEventType_CREDS_IMPORT                     ClientEventType = 48
	ClientEventType_REMOTEAPI_ADD_TABLE_FILES        ClientEventType = 49
	ClientEventType_MIGRATE                          ClientEventType = 50
	ClientEventType_TAG                              ClientEventType = 51
	ClientEventType_GARBAGE_COLLECTION               ClientEventType = 52
	ClientEventType_FILTER_BRANCH                    ClientEventType = 53
	ClientEventType_DUMP                             ClientEventType = 54
	ClientEventType_CHERRY_PICK                      ClientEventType = 55
	ClientEventType_STASH                            ClientEventType = 56
	ClientEventType_STASH_CLEAR                      ClientEventType = 57
	ClientEventType_STASH_DROP                       ClientEventType = 58
	ClientEventType_STASH_LIST                       ClientEventType = 59
	ClientEventType_STASH_POP                        ClientEventType = 60
	ClientEventType_SHOW                             ClientEventType = 61
	ClientEventType_PROFILE                          ClientEventType = 62
	ClientEventType_REFLOG                           ClientEventType = 63
	ClientEventType_SQL_SERVER_HEARTBEAT             ClientEventType = 64
	ClientEventType_REBASE                           ClientEventType = 65
)

// Enum value maps for ClientEventType.
var (
	ClientEventType_name = map[int32]string{
		0:  "TYPE_UNSPECIFIED",
		1:  "INIT",
		2:  "STATUS",
		3:  "ADD",
		4:  "RESET",
		5:  "COMMIT",
		6:  "SQL",
		7:  "SQL_SERVER",
		8:  "LOG",
		9:  "DIFF",
		10: "MERGE",
		11: "BRANCH",
		12: "CHECKOUT",
		13: "REMOTE",
		14: "PUSH",
		15: "PULL",
		16: "FETCH",
		17: "CLONE",
		18: "LOGIN",
		19: "VERSION",
		20: "CONFIG",
		21: "LS",
		22: "SCHEMA",
		23: "TABLE_IMPORT",
		24: "TABLE_EXPORT",
		25: "TABLE_CREATE",
		26: "TABLE_RM",
		27: "TABLE_MV",
		28: "TABLE_CP",
		29: "TABLE_SELECT",
		30: "TABLE_PUT_ROW",
		31: "TABLE_RM_ROW",
		32: "CREDS_NEW",
		33: "CREDS_RM",
		34: "CREDS_LS",
		35: "CONF_CAT",
		36: "CONF_RESOLVE",
		37: "REMOTEAPI_GET_REPO_METADATA",
		38: "REMOTEAPI_HAS_CHUNKS",
		39: "REMOTEAPI_GET_DOWNLOAD_LOCATIONS",
		40: "REMOTEAPI_GET_UPLOAD_LOCATIONS",
		41: "REMOTEAPI_REBASE",
		42: "REMOTEAPI_ROOT",
		43: "REMOTEAPI_COMMIT",
		44: "REMOTEAPI_LIST_TABLE_FILES",
		45: "BLAME",
		46: "CREDS_CHECK",
		47: "CREDS_USE",
		48: "CREDS_IMPORT",
		49: "REMOTEAPI_ADD_TABLE_FILES",
		50: "MIGRATE",
		51: "TAG",
		52: "GARBAGE_COLLECTION",
		53: "FILTER_BRANCH",
		54: "DUMP",
		55: "CHERRY_PICK",
		56: "STASH",
		57: "STASH_CLEAR",
		58: "STASH_DROP",
		59: "STASH_LIST",
		60: "STASH_POP",
		61: "SHOW",
		62: "PROFILE",
		63: "REFLOG",
		64: "SQL_SERVER_HEARTBEAT",
		65: "REBASE",
	}
	ClientEventType_value = map[string]int32{
		"TYPE_UNSPECIFIED":                 0,
		"INIT":                             1,
		"STATUS":                           2,
		"ADD":                              3,
		"RESET":                            4,
		"COMMIT":                           5,
		"SQL":                              6,
		"SQL_SERVER":                       7,
		"LOG":                              8,
		"DIFF":                             9,
		"MERGE":                            10,
		"BRANCH":                           11,
		"CHECKOUT":                         12,
		"REMOTE":                           13,
		"PUSH":                             14,
		"PULL":                             15,
		"FETCH":                            16,
		"CLONE":                            17,
		"LOGIN":                            18,
		"VERSION":                          19,
		"CONFIG":                           20,
		"LS":                               21,
		"SCHEMA":                           22,
		"TABLE_IMPORT":                     23,
		"TABLE_EXPORT":                     24,
		"TABLE_CREATE":                     25,
		"TABLE_RM":                         26,
		"TABLE_MV":                         27,
		"TABLE_CP":                         28,
		"TABLE_SELECT":                     29,
		"TABLE_PUT_ROW":                    30,
		"TABLE_RM_ROW":                     31,
		"CREDS_NEW":                        32,
		"CREDS_RM":                         33,
		"CREDS_LS":                         34,
		"CONF_CAT":                         35,
		"CONF_RESOLVE":                     36,
		"REMOTEAPI_GET_REPO_METADATA":      37,
		"REMOTEAPI_HAS_CHUNKS":             38,
		"REMOTEAPI_GET_DOWNLOAD_LOCATIONS": 39,
		"REMOTEAPI_GET_UPLOAD_LOCATIONS":   40,
		"REMOTEAPI_REBASE":                 41,
		"REMOTEAPI_ROOT":                   42,
		"REMOTEAPI_COMMIT":                 43,
		"REMOTEAPI_LIST_TABLE_FILES":       44,
		"BLAME":                            45,
		"CREDS_CHECK":                      46,
		"CREDS_USE":                        47,
		"CREDS_IMPORT":                     48,
		"REMOTEAPI_ADD_TABLE_FILES":        49,
		"MIGRATE":                          50,
		"TAG":                              51,
		"GARBAGE_COLLECTION":               52,
		"FILTER_BRANCH":                    53,
		"DUMP":                             54,
		"CHERRY_PICK":                      55,
		"STASH":                            56,
		"STASH_CLEAR":                      57,
		"STASH_DROP":                       58,
		"STASH_LIST":                       59,
		"STASH_POP":                        60,
		"SHOW":                             61,
		"PROFILE":                          62,
		"REFLOG":                           63,
		"SQL_SERVER_HEARTBEAT":             64,
		"REBASE":                           65,
	}
)

func (x ClientEventType) Enum() *ClientEventType {
	p := new(ClientEventType)
	*p = x
	return p
}

func (x ClientEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_dolt_services_eventsapi_v1alpha1_event_constants_proto_enumTypes[1].Descriptor()
}

func (ClientEventType) Type() protoreflect.EnumType {
	return &file_dolt_services_eventsapi_v1alpha1_event_constants_proto_enumTypes[1]
}

func (x ClientEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientEventType.Descriptor instead.
func (ClientEventType) EnumDescriptor() ([]byte, []int) {
	return file_dolt_services_eventsapi_v1alpha1_event_constants_proto_rawDescGZIP(), []int{1}
}

type MetricID int32

const (
	MetricID_METRIC_UNSPECIFIED  MetricID = 0
	MetricID_BYTES_DOWNLOADED    MetricID = 1
	MetricID_DOWNLOAD_MS_ELAPSED MetricID = 2
	MetricID_REMOTEAPI_RPC_ERROR MetricID = 3
)

// Enum value maps for MetricID.
var (
	MetricID_name = map[int32]string{
		0: "METRIC_UNSPECIFIED",
		1: "BYTES_DOWNLOADED",
		2: "DOWNLOAD_MS_ELAPSED",
		3: "REMOTEAPI_RPC_ERROR",
	}
	MetricID_value = map[string]int32{
		"METRIC_UNSPECIFIED":  0,
		"BYTES_DOWNLOADED":    1,
		"DOWNLOAD_MS_ELAPSED": 2,
		"REMOTEAPI_RPC_ERROR": 3,
	}
)

func (x MetricID) Enum() *MetricID {
	p := new(MetricID)
	*p = x
	return p
}

func (x MetricID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricID) Descriptor() protoreflect.EnumDescriptor {
	return file_dolt_services_eventsapi_v1alpha1_event_constants_proto_enumTypes[2].Descriptor()
}

func (MetricID) Type() protoreflect.EnumType {
	return &file_dolt_services_eventsapi_v1alpha1_event_constants_proto_enumTypes[2]
}

func (x MetricID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricID.Descriptor instead.
func (MetricID) EnumDescriptor() ([]byte, []int) {
	return file_dolt_services_eventsapi_v1alpha1_event_constants_proto_rawDescGZIP(), []int{2}
}

type AttributeID int32

const (
	AttributeID_ATTRIBUTE_UNSPECIFIED AttributeID = 0
	AttributeID_REMOTE_URL_SCHEME     AttributeID = 2
)

// Enum value maps for AttributeID.
var (
	AttributeID_name = map[int32]string{
		0: "ATTRIBUTE_UNSPECIFIED",
		2: "REMOTE_URL_SCHEME",
	}
	AttributeID_value = map[string]int32{
		"ATTRIBUTE_UNSPECIFIED": 0,
		"REMOTE_URL_SCHEME":     2,
	}
)

func (x AttributeID) Enum() *AttributeID {
	p := new(AttributeID)
	*p = x
	return p
}

func (x AttributeID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AttributeID) Descriptor() protoreflect.EnumDescriptor {
	return file_dolt_services_eventsapi_v1alpha1_event_constants_proto_enumTypes[3].Descriptor()
}

func (AttributeID) Type() protoreflect.EnumType {
	return &file_dolt_services_eventsapi_v1alpha1_event_constants_proto_enumTypes[3]
}

func (x AttributeID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AttributeID.Descriptor instead.
func (AttributeID) EnumDescriptor() ([]byte, []int) {
	return file_dolt_services_eventsapi_v1alpha1_event_constants_proto_rawDescGZIP(), []int{3}
}

type AppID int32

const (
	AppID_APP_ID_UNSPECIFIED AppID = 0
	AppID_APP_DOLT           AppID = 1
	AppID_APP_DOLTGRES       AppID = 2
)

// Enum value maps for AppID.
var (
	AppID_name = map[int32]string{
		0: "APP_ID_UNSPECIFIED",
		1: "APP_DOLT",
		2: "APP_DOLTGRES",
	}
	AppID_value = map[string]int32{
		"APP_ID_UNSPECIFIED": 0,
		"APP_DOLT":           1,
		"APP_DOLTGRES":       2,
	}
)

func (x AppID) Enum() *AppID {
	p := new(AppID)
	*p = x
	return p
}

func (x AppID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppID) Descriptor() protoreflect.EnumDescriptor {
	return file_dolt_services_eventsapi_v1alpha1_event_constants_proto_enumTypes[4].Descriptor()
}

func (AppID) Type() protoreflect.EnumType {
	return &file_dolt_services_eventsapi_v1alpha1_event_constants_proto_enumTypes[4]
}

func (x AppID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppID.Descriptor instead.
func (AppID) EnumDescriptor() ([]byte, []int) {
	return file_dolt_services_eventsapi_v1alpha1_event_constants_proto_rawDescGZIP(), []int{4}
}

var File_dolt_services_eventsapi_v1alpha1_event_constants_proto protoreflect.FileDescriptor

var file_dolt_services_eventsapi_v1alpha1_event_constants_proto_rawDesc = []byte{
	0x0a, 0x36, 0x64, 0x6f, 0x6c, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x64, 0x6f, 0x6c, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2a, 0x48, 0x0a, 0x08, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f,
	0x52, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x4c, 0x49, 0x4e, 0x55, 0x58, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x57,
	0x49, 0x4e, 0x44, 0x4f, 0x57, 0x53, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x41, 0x52, 0x57,
	0x49, 0x4e, 0x10, 0x03, 0x2a, 0xaf, 0x08, 0x0a, 0x0f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x49, 0x4e, 0x49, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x44, 0x44, 0x10, 0x03, 0x12, 0x09, 0x0a,
	0x05, 0x52, 0x45, 0x53, 0x45, 0x54, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x4d, 0x4d,
	0x49, 0x54, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x51, 0x4c, 0x10, 0x06, 0x12, 0x0e, 0x0a,
	0x0a, 0x53, 0x51, 0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x07, 0x12, 0x07, 0x0a,
	0x03, 0x4c, 0x4f, 0x47, 0x10, 0x08, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x49, 0x46, 0x46, 0x10, 0x09,
	0x12, 0x09, 0x0a, 0x05, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06, 0x42,
	0x52, 0x41, 0x4e, 0x43, 0x48, 0x10, 0x0b, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x48, 0x45, 0x43, 0x4b,
	0x4f, 0x55, 0x54, 0x10, 0x0c, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4d, 0x4f, 0x54, 0x45, 0x10,
	0x0d, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x55, 0x53, 0x48, 0x10, 0x0e, 0x12, 0x08, 0x0a, 0x04, 0x50,
	0x55, 0x4c, 0x4c, 0x10, 0x0f, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x45, 0x54, 0x43, 0x48, 0x10, 0x10,
	0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x4f, 0x4e, 0x45, 0x10, 0x11, 0x12, 0x09, 0x0a, 0x05, 0x4c,
	0x4f, 0x47, 0x49, 0x4e, 0x10, 0x12, 0x12, 0x0b, 0x0a, 0x07, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f,
	0x4e, 0x10, 0x13, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10, 0x14, 0x12,
	0x06, 0x0a, 0x02, 0x4c, 0x53, 0x10, 0x15, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x43, 0x48, 0x45, 0x4d,
	0x41, 0x10, 0x16, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x49, 0x4d, 0x50,
	0x4f, 0x52, 0x54, 0x10, 0x17, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x45,
	0x58, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x18, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x41, 0x42, 0x4c, 0x45,
	0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x19, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x41, 0x42,
	0x4c, 0x45, 0x5f, 0x52, 0x4d, 0x10, 0x1a, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x41, 0x42, 0x4c, 0x45,
	0x5f, 0x4d, 0x56, 0x10, 0x1b, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x43,
	0x50, 0x10, 0x1c, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x53, 0x45, 0x4c,
	0x45, 0x43, 0x54, 0x10, 0x1d, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x50,
	0x55, 0x54, 0x5f, 0x52, 0x4f, 0x57, 0x10, 0x1e, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x41, 0x42, 0x4c,
	0x45, 0x5f, 0x52, 0x4d, 0x5f, 0x52, 0x4f, 0x57, 0x10, 0x1f, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x52,
	0x45, 0x44, 0x53, 0x5f, 0x4e, 0x45, 0x57, 0x10, 0x20, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45,
	0x44, 0x53, 0x5f, 0x52, 0x4d, 0x10, 0x21, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x44, 0x53,
	0x5f, 0x4c, 0x53, 0x10, 0x22, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4e, 0x46, 0x5f, 0x43, 0x41,
	0x54, 0x10, 0x23, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x4e, 0x46, 0x5f, 0x52, 0x45, 0x53, 0x4f,
	0x4c, 0x56, 0x45, 0x10, 0x24, 0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x45, 0x4d, 0x4f, 0x54, 0x45, 0x41,
	0x50, 0x49, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x5f, 0x4d, 0x45, 0x54, 0x41,
	0x44, 0x41, 0x54, 0x41, 0x10, 0x25, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x4d, 0x4f, 0x54, 0x45,
	0x41, 0x50, 0x49, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x43, 0x48, 0x55, 0x4e, 0x4b, 0x53, 0x10, 0x26,
	0x12, 0x24, 0x0a, 0x20, 0x52, 0x45, 0x4d, 0x4f, 0x54, 0x45, 0x41, 0x50, 0x49, 0x5f, 0x47, 0x45,
	0x54, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x53, 0x10, 0x27, 0x12, 0x22, 0x0a, 0x1e, 0x52, 0x45, 0x4d, 0x4f, 0x54, 0x45,
	0x41, 0x50, 0x49, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x4c,
	0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x28, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x45,
	0x4d, 0x4f, 0x54, 0x45, 0x41, 0x50, 0x49, 0x5f, 0x52, 0x45, 0x42, 0x41, 0x53, 0x45, 0x10, 0x29,
	0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x4d, 0x4f, 0x54, 0x45, 0x41, 0x50, 0x49, 0x5f, 0x52, 0x4f,
	0x4f, 0x54, 0x10, 0x2a, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x4d, 0x4f, 0x54, 0x45, 0x41, 0x50,
	0x49, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x10, 0x2b, 0x12, 0x1e, 0x0a, 0x1a, 0x52, 0x45,
	0x4d, 0x4f, 0x54, 0x45, 0x41, 0x50, 0x49, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x54, 0x41, 0x42,
	0x4c, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x53, 0x10, 0x2c, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x4c,
	0x41, 0x4d, 0x45, 0x10, 0x2d, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x52, 0x45, 0x44, 0x53, 0x5f, 0x43,
	0x48, 0x45, 0x43, 0x4b, 0x10, 0x2e, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x52, 0x45, 0x44, 0x53, 0x5f,
	0x55, 0x53, 0x45, 0x10, 0x2f, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x52, 0x45, 0x44, 0x53, 0x5f, 0x49,
	0x4d, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x30, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x4d, 0x4f, 0x54,
	0x45, 0x41, 0x50, 0x49, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x46,
	0x49, 0x4c, 0x45, 0x53, 0x10, 0x31, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x49, 0x47, 0x52, 0x41, 0x54,
	0x45, 0x10, 0x32, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x41, 0x47, 0x10, 0x33, 0x12, 0x16, 0x0a, 0x12,
	0x47, 0x41, 0x52, 0x42, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x34, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42,
	0x52, 0x41, 0x4e, 0x43, 0x48, 0x10, 0x35, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x55, 0x4d, 0x50, 0x10,
	0x36, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x48, 0x45, 0x52, 0x52, 0x59, 0x5f, 0x50, 0x49, 0x43, 0x4b,
	0x10, 0x37, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x53, 0x48, 0x10, 0x38, 0x12, 0x0f, 0x0a,
	0x0b, 0x53, 0x54, 0x41, 0x53, 0x48, 0x5f, 0x43, 0x4c, 0x45, 0x41, 0x52, 0x10, 0x39, 0x12, 0x0e,
	0x0a, 0x0a, 0x53, 0x54, 0x41, 0x53, 0x48, 0x5f, 0x44, 0x52, 0x4f, 0x50, 0x10, 0x3a, 0x12, 0x0e,
	0x0a, 0x0a, 0x53, 0x54, 0x41, 0x53, 0x48, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x3b, 0x12, 0x0d,
	0x0a, 0x09, 0x53, 0x54, 0x41, 0x53, 0x48, 0x5f, 0x50, 0x4f, 0x50, 0x10, 0x3c, 0x12, 0x08, 0x0a,
	0x04, 0x53, 0x48, 0x4f, 0x57, 0x10, 0x3d, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x4f, 0x46, 0x49,
	0x4c, 0x45, 0x10, 0x3e, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x46, 0x4c, 0x4f, 0x47, 0x10, 0x3f,
	0x12, 0x18, 0x0a, 0x14, 0x53, 0x51, 0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x48,
	0x45, 0x41, 0x52, 0x54, 0x42, 0x45, 0x41, 0x54, 0x10, 0x40, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45,
	0x42, 0x41, 0x53, 0x45, 0x10, 0x41, 0x2a, 0x6a, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x42, 0x59,
	0x54, 0x45, 0x53, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x17, 0x0a, 0x13, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x4d, 0x53, 0x5f,
	0x45, 0x4c, 0x41, 0x50, 0x53, 0x45, 0x44, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x4d,
	0x4f, 0x54, 0x45, 0x41, 0x50, 0x49, 0x5f, 0x52, 0x50, 0x43, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x03, 0x2a, 0x45, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49,
	0x44, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11,
	0x52, 0x45, 0x4d, 0x4f, 0x54, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d,
	0x45, 0x10, 0x02, 0x22, 0x04, 0x08, 0x01, 0x10, 0x01, 0x2a, 0x3f, 0x0a, 0x05, 0x41, 0x70, 0x70,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x50, 0x50, 0x5f, 0x49, 0x44, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x50,
	0x50, 0x5f, 0x44, 0x4f, 0x4c, 0x54, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x50, 0x50, 0x5f,
	0x44, 0x4f, 0x4c, 0x54, 0x47, 0x52, 0x45, 0x53, 0x10, 0x02, 0x42, 0x51, 0x5a, 0x4f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x6c, 0x74, 0x68, 0x75, 0x62,
	0x2f, 0x64, 0x6f, 0x6c, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x6c, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dolt_services_eventsapi_v1alpha1_event_constants_proto_rawDescOnce sync.Once
	file_dolt_services_eventsapi_v1alpha1_event_constants_proto_rawDescData = file_dolt_services_eventsapi_v1alpha1_event_constants_proto_rawDesc
)

func file_dolt_services_eventsapi_v1alpha1_event_constants_proto_rawDescGZIP() []byte {
	file_dolt_services_eventsapi_v1alpha1_event_constants_proto_rawDescOnce.Do(func() {
		file_dolt_services_eventsapi_v1alpha1_event_constants_proto_rawDescData = protoimpl.X.CompressGZIP(file_dolt_services_eventsapi_v1alpha1_event_constants_proto_rawDescData)
	})
	return file_dolt_services_eventsapi_v1alpha1_event_constants_proto_rawDescData
}

var file_dolt_services_eventsapi_v1alpha1_event_constants_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_dolt_services_eventsapi_v1alpha1_event_constants_proto_goTypes = []interface{}{
	(Platform)(0),        // 0: dolt.services.eventsapi.v1alpha1.Platform
	(ClientEventType)(0), // 1: dolt.services.eventsapi.v1alpha1.ClientEventType
	(MetricID)(0),        // 2: dolt.services.eventsapi.v1alpha1.MetricID
	(AttributeID)(0),     // 3: dolt.services.eventsapi.v1alpha1.AttributeID
	(AppID)(0),           // 4: dolt.services.eventsapi.v1alpha1.AppID
}
var file_dolt_services_eventsapi_v1alpha1_event_constants_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dolt_services_eventsapi_v1alpha1_event_constants_proto_init() }
func file_dolt_services_eventsapi_v1alpha1_event_constants_proto_init() {
	if File_dolt_services_eventsapi_v1alpha1_event_constants_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dolt_services_eventsapi_v1alpha1_event_constants_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dolt_services_eventsapi_v1alpha1_event_constants_proto_goTypes,
		DependencyIndexes: file_dolt_services_eventsapi_v1alpha1_event_constants_proto_depIdxs,
		EnumInfos:         file_dolt_services_eventsapi_v1alpha1_event_constants_proto_enumTypes,
	}.Build()
	File_dolt_services_eventsapi_v1alpha1_event_constants_proto = out.File
	file_dolt_services_eventsapi_v1alpha1_event_constants_proto_rawDesc = nil
	file_dolt_services_eventsapi_v1alpha1_event_constants_proto_goTypes = nil
	file_dolt_services_eventsapi_v1alpha1_event_constants_proto_depIdxs = nil
}
