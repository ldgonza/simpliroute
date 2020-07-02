//
// (C) Copyright SimpliRoute.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.11.4
// source: simpli/commons/global/global.proto

package global

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Individual location
type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Geolocation latitude
	Latitude float32 `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// Geolocation longitude
	Longitude float32 `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simpli_commons_global_global_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_simpli_commons_global_global_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_simpli_commons_global_global_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Location) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

// A period model
type Period struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// period start
	Start *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	// period end
	End *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *Period) Reset() {
	*x = Period{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simpli_commons_global_global_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Period) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Period) ProtoMessage() {}

func (x *Period) ProtoReflect() protoreflect.Message {
	mi := &file_simpli_commons_global_global_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Period.ProtoReflect.Descriptor instead.
func (*Period) Descriptor() ([]byte, []int) {
	return file_simpli_commons_global_global_proto_rawDescGZIP(), []int{1}
}

func (x *Period) GetStart() *timestamp.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *Period) GetEnd() *timestamp.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

var File_simpli_commons_global_global_proto protoreflect.FileDescriptor

var file_simpli_commons_global_global_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x08,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x22, 0x68, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x30, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2c,
	0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x42, 0x34, 0x5a, 0x32,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x69, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_simpli_commons_global_global_proto_rawDescOnce sync.Once
	file_simpli_commons_global_global_proto_rawDescData = file_simpli_commons_global_global_proto_rawDesc
)

func file_simpli_commons_global_global_proto_rawDescGZIP() []byte {
	file_simpli_commons_global_global_proto_rawDescOnce.Do(func() {
		file_simpli_commons_global_global_proto_rawDescData = protoimpl.X.CompressGZIP(file_simpli_commons_global_global_proto_rawDescData)
	})
	return file_simpli_commons_global_global_proto_rawDescData
}

var file_simpli_commons_global_global_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_simpli_commons_global_global_proto_goTypes = []interface{}{
	(*Location)(nil),            // 0: simpli.commons.global.Location
	(*Period)(nil),              // 1: simpli.commons.global.Period
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_simpli_commons_global_global_proto_depIdxs = []int32{
	2, // 0: simpli.commons.global.Period.start:type_name -> google.protobuf.Timestamp
	2, // 1: simpli.commons.global.Period.end:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_simpli_commons_global_global_proto_init() }
func file_simpli_commons_global_global_proto_init() {
	if File_simpli_commons_global_global_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_simpli_commons_global_global_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_simpli_commons_global_global_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Period); i {
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
			RawDescriptor: file_simpli_commons_global_global_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_simpli_commons_global_global_proto_goTypes,
		DependencyIndexes: file_simpli_commons_global_global_proto_depIdxs,
		MessageInfos:      file_simpli_commons_global_global_proto_msgTypes,
	}.Build()
	File_simpli_commons_global_global_proto = out.File
	file_simpli_commons_global_global_proto_rawDesc = nil
	file_simpli_commons_global_global_proto_goTypes = nil
	file_simpli_commons_global_global_proto_depIdxs = nil
}
