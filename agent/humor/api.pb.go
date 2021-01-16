// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: agent/humor/api.proto

package humor

import (
	proto "github.com/golang/protobuf/proto"
	common "github.com/wilenceyao/humor-api/common"
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

type TtsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *common.BaseRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Text    string              `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *TtsRequest) Reset() {
	*x = TtsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_humor_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TtsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TtsRequest) ProtoMessage() {}

func (x *TtsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_humor_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TtsRequest.ProtoReflect.Descriptor instead.
func (*TtsRequest) Descriptor() ([]byte, []int) {
	return file_agent_humor_api_proto_rawDescGZIP(), []int{0}
}

func (x *TtsRequest) GetRequest() *common.BaseRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *TtsRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type TtsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *common.BaseResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *TtsResponse) Reset() {
	*x = TtsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_humor_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TtsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TtsResponse) ProtoMessage() {}

func (x *TtsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_humor_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TtsResponse.ProtoReflect.Descriptor instead.
func (*TtsResponse) Descriptor() ([]byte, []int) {
	return file_agent_humor_api_proto_rawDescGZIP(), []int{1}
}

func (x *TtsResponse) GetResponse() *common.BaseResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type WeatherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *common.BaseRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *WeatherRequest) Reset() {
	*x = WeatherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_humor_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeatherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherRequest) ProtoMessage() {}

func (x *WeatherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_humor_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherRequest.ProtoReflect.Descriptor instead.
func (*WeatherRequest) Descriptor() ([]byte, []int) {
	return file_agent_humor_api_proto_rawDescGZIP(), []int{2}
}

func (x *WeatherRequest) GetRequest() *common.BaseRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type WeatherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *common.BaseResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *WeatherResponse) Reset() {
	*x = WeatherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_humor_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeatherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherResponse) ProtoMessage() {}

func (x *WeatherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_humor_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherResponse.ProtoReflect.Descriptor instead.
func (*WeatherResponse) Descriptor() ([]byte, []int) {
	return file_agent_humor_api_proto_rawDescGZIP(), []int{3}
}

func (x *WeatherResponse) GetResponse() *common.BaseResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type TakePhotoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *common.BaseRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Id      string              `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TakePhotoRequest) Reset() {
	*x = TakePhotoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_humor_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakePhotoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakePhotoRequest) ProtoMessage() {}

func (x *TakePhotoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_humor_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakePhotoRequest.ProtoReflect.Descriptor instead.
func (*TakePhotoRequest) Descriptor() ([]byte, []int) {
	return file_agent_humor_api_proto_rawDescGZIP(), []int{4}
}

func (x *TakePhotoRequest) GetRequest() *common.BaseRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *TakePhotoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TakePhotoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *common.BaseResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *TakePhotoResponse) Reset() {
	*x = TakePhotoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_humor_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakePhotoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakePhotoResponse) ProtoMessage() {}

func (x *TakePhotoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_humor_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakePhotoResponse.ProtoReflect.Descriptor instead.
func (*TakePhotoResponse) Descriptor() ([]byte, []int) {
	return file_agent_humor_api_proto_rawDescGZIP(), []int{5}
}

func (x *TakePhotoResponse) GetResponse() *common.BaseResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_agent_humor_api_proto protoreflect.FileDescriptor

var file_agent_humor_api_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x68, 0x75, 0x6d, 0x6f, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68, 0x75, 0x6d, 0x6f, 0x72, 0x1a, 0x13,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x0a, 0x54, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2d, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x3f, 0x0a, 0x0b, 0x54, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0x0a, 0x0e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x43, 0x0a, 0x0f, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x0a, 0x10, 0x54,
	0x61, 0x6b, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2d, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x45,
	0x0a, 0x11, 0x54, 0x61, 0x6b, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb6, 0x01, 0x0a, 0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x54, 0x74, 0x73, 0x12, 0x11, 0x2e,
	0x68, 0x75, 0x6d, 0x6f, 0x72, 0x2e, 0x54, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x68, 0x75, 0x6d, 0x6f, 0x72, 0x2e, 0x54, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12,
	0x15, 0x2e, 0x68, 0x75, 0x6d, 0x6f, 0x72, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x68, 0x75, 0x6d, 0x6f, 0x72, 0x2e, 0x57,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e,
	0x0a, 0x09, 0x54, 0x61, 0x6b, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x2e, 0x68, 0x75,
	0x6d, 0x6f, 0x72, 0x2e, 0x54, 0x61, 0x6b, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x68, 0x75, 0x6d, 0x6f, 0x72, 0x2e, 0x54, 0x61, 0x6b,
	0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d,
	0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x69, 0x6c,
	0x65, 0x6e, 0x63, 0x65, 0x79, 0x61, 0x6f, 0x2f, 0x68, 0x75, 0x6d, 0x6f, 0x72, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x68, 0x75, 0x6d, 0x6f, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agent_humor_api_proto_rawDescOnce sync.Once
	file_agent_humor_api_proto_rawDescData = file_agent_humor_api_proto_rawDesc
)

func file_agent_humor_api_proto_rawDescGZIP() []byte {
	file_agent_humor_api_proto_rawDescOnce.Do(func() {
		file_agent_humor_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_humor_api_proto_rawDescData)
	})
	return file_agent_humor_api_proto_rawDescData
}

var file_agent_humor_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_agent_humor_api_proto_goTypes = []interface{}{
	(*TtsRequest)(nil),          // 0: humor.TtsRequest
	(*TtsResponse)(nil),         // 1: humor.TtsResponse
	(*WeatherRequest)(nil),      // 2: humor.WeatherRequest
	(*WeatherResponse)(nil),     // 3: humor.WeatherResponse
	(*TakePhotoRequest)(nil),    // 4: humor.TakePhotoRequest
	(*TakePhotoResponse)(nil),   // 5: humor.TakePhotoResponse
	(*common.BaseRequest)(nil),  // 6: common.BaseRequest
	(*common.BaseResponse)(nil), // 7: common.BaseResponse
}
var file_agent_humor_api_proto_depIdxs = []int32{
	6, // 0: humor.TtsRequest.request:type_name -> common.BaseRequest
	7, // 1: humor.TtsResponse.response:type_name -> common.BaseResponse
	6, // 2: humor.WeatherRequest.request:type_name -> common.BaseRequest
	7, // 3: humor.WeatherResponse.response:type_name -> common.BaseResponse
	6, // 4: humor.TakePhotoRequest.request:type_name -> common.BaseRequest
	7, // 5: humor.TakePhotoResponse.response:type_name -> common.BaseResponse
	0, // 6: humor.AgentService.Tts:input_type -> humor.TtsRequest
	2, // 7: humor.AgentService.Weather:input_type -> humor.WeatherRequest
	4, // 8: humor.AgentService.TakePhoto:input_type -> humor.TakePhotoRequest
	1, // 9: humor.AgentService.Tts:output_type -> humor.TtsResponse
	3, // 10: humor.AgentService.Weather:output_type -> humor.WeatherResponse
	5, // 11: humor.AgentService.TakePhoto:output_type -> humor.TakePhotoResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_agent_humor_api_proto_init() }
func file_agent_humor_api_proto_init() {
	if File_agent_humor_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agent_humor_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TtsRequest); i {
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
		file_agent_humor_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TtsResponse); i {
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
		file_agent_humor_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeatherRequest); i {
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
		file_agent_humor_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeatherResponse); i {
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
		file_agent_humor_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakePhotoRequest); i {
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
		file_agent_humor_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakePhotoResponse); i {
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
			RawDescriptor: file_agent_humor_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_humor_api_proto_goTypes,
		DependencyIndexes: file_agent_humor_api_proto_depIdxs,
		MessageInfos:      file_agent_humor_api_proto_msgTypes,
	}.Build()
	File_agent_humor_api_proto = out.File
	file_agent_humor_api_proto_rawDesc = nil
	file_agent_humor_api_proto_goTypes = nil
	file_agent_humor_api_proto_depIdxs = nil
}
