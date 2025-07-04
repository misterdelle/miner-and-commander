// Copyright (C) 2023  Braiins Systems s.r.o.
//
// This file is part of Braiins Open-Source Initiative (BOSI).
//
// BOSI is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//
// Please, keep in mind that we may also license BOSI or any part thereof
// under a proprietary license. For more information on the terms and conditions
// of such proprietary license or if you have any other questions, please
// contact us at opensource@braiins.com.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v6.31.1
// source: bos/version.proto

package bos_proto

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

// LATEST_API_VERSION=1.5.0
type ApiVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Major uint64 `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor uint64 `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch uint64 `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
	Pre   string `protobuf:"bytes,4,opt,name=pre,proto3" json:"pre,omitempty"`
	Build string `protobuf:"bytes,5,opt,name=build,proto3" json:"build,omitempty"`
}

func (x *ApiVersion) Reset() {
	*x = ApiVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bos_version_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiVersion) ProtoMessage() {}

func (x *ApiVersion) ProtoReflect() protoreflect.Message {
	mi := &file_bos_version_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiVersion.ProtoReflect.Descriptor instead.
func (*ApiVersion) Descriptor() ([]byte, []int) {
	return file_bos_version_proto_rawDescGZIP(), []int{0}
}

func (x *ApiVersion) GetMajor() uint64 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *ApiVersion) GetMinor() uint64 {
	if x != nil {
		return x.Minor
	}
	return 0
}

func (x *ApiVersion) GetPatch() uint64 {
	if x != nil {
		return x.Patch
	}
	return 0
}

func (x *ApiVersion) GetPre() string {
	if x != nil {
		return x.Pre
	}
	return ""
}

func (x *ApiVersion) GetBuild() string {
	if x != nil {
		return x.Build
	}
	return ""
}

type ApiVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ApiVersionRequest) Reset() {
	*x = ApiVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bos_version_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiVersionRequest) ProtoMessage() {}

func (x *ApiVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bos_version_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiVersionRequest.ProtoReflect.Descriptor instead.
func (*ApiVersionRequest) Descriptor() ([]byte, []int) {
	return file_bos_version_proto_rawDescGZIP(), []int{1}
}

var File_bos_version_proto protoreflect.FileDescriptor

var file_bos_version_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x6f, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x72, 0x61, 0x69, 0x69, 0x6e, 0x73, 0x2e, 0x62, 0x6f, 0x73,
	0x22, 0x76, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d,
	0x61, 0x6a, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70,
	0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x41, 0x70, 0x69, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x5d, 0x0a,
	0x11, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x62, 0x72, 0x61, 0x69, 0x69, 0x6e, 0x73, 0x2e, 0x62, 0x6f,
	0x73, 0x2e, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x62, 0x72, 0x61, 0x69, 0x69, 0x6e, 0x73, 0x2e, 0x62, 0x6f,
	0x73, 0x2e, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x37, 0x5a, 0x35,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x61, 0x69, 0x69,
	0x6e, 0x73, 0x2f, 0x62, 0x6f, 0x73, 0x2d, 0x70, 0x6c, 0x75, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x62, 0x72, 0x61, 0x69, 0x69, 0x6e, 0x73, 0x2f, 0x62, 0x6f, 0x73, 0x3b, 0x62, 0x6f, 0x73, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bos_version_proto_rawDescOnce sync.Once
	file_bos_version_proto_rawDescData = file_bos_version_proto_rawDesc
)

func file_bos_version_proto_rawDescGZIP() []byte {
	file_bos_version_proto_rawDescOnce.Do(func() {
		file_bos_version_proto_rawDescData = protoimpl.X.CompressGZIP(file_bos_version_proto_rawDescData)
	})
	return file_bos_version_proto_rawDescData
}

var file_bos_version_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bos_version_proto_goTypes = []interface{}{
	(*ApiVersion)(nil),        // 0: braiins.bos.ApiVersion
	(*ApiVersionRequest)(nil), // 1: braiins.bos.ApiVersionRequest
}
var file_bos_version_proto_depIdxs = []int32{
	1, // 0: braiins.bos.ApiVersionService.GetApiVersion:input_type -> braiins.bos.ApiVersionRequest
	0, // 1: braiins.bos.ApiVersionService.GetApiVersion:output_type -> braiins.bos.ApiVersion
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bos_version_proto_init() }
func file_bos_version_proto_init() {
	if File_bos_version_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bos_version_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiVersion); i {
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
		file_bos_version_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiVersionRequest); i {
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
			RawDescriptor: file_bos_version_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bos_version_proto_goTypes,
		DependencyIndexes: file_bos_version_proto_depIdxs,
		MessageInfos:      file_bos_version_proto_msgTypes,
	}.Build()
	File_bos_version_proto = out.File
	file_bos_version_proto_rawDesc = nil
	file_bos_version_proto_goTypes = nil
	file_bos_version_proto_depIdxs = nil
}
