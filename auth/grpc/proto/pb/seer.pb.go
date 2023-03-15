// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: seer.proto

package pb

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

type SeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid string `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
}

func (x *SeerRequest) Reset() {
	*x = SeerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeerRequest) ProtoMessage() {}

func (x *SeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_seer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeerRequest.ProtoReflect.Descriptor instead.
func (*SeerRequest) Descriptor() ([]byte, []int) {
	return file_seer_proto_rawDescGZIP(), []int{0}
}

func (x *SeerRequest) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

type SeerUUIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeerUuid string `protobuf:"bytes,1,opt,name=seer_uuid,json=seerUuid,proto3" json:"seer_uuid,omitempty"`
}

func (x *SeerUUIDRequest) Reset() {
	*x = SeerUUIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeerUUIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeerUUIDRequest) ProtoMessage() {}

func (x *SeerUUIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_seer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeerUUIDRequest.ProtoReflect.Descriptor instead.
func (*SeerUUIDRequest) Descriptor() ([]byte, []int) {
	return file_seer_proto_rawDescGZIP(), []int{1}
}

func (x *SeerUUIDRequest) GetSeerUuid() string {
	if x != nil {
		return x.SeerUuid
	}
	return ""
}

type SeerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid               string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	OnsiteAvailable    bool   `protobuf:"varint,2,opt,name=onsite_available,json=onsiteAvailable,proto3" json:"onsite_available,omitempty"`
	ChatAvailable      bool   `protobuf:"varint,3,opt,name=chat_available,json=chatAvailable,proto3" json:"chat_available,omitempty"`
	CallAvailable      bool   `protobuf:"varint,4,opt,name=call_available,json=callAvailable,proto3" json:"call_available,omitempty"`
	VideoCallAvailable bool   `protobuf:"varint,5,opt,name=video_call_available,json=videoCallAvailable,proto3" json:"video_call_available,omitempty"`
	Major              string `protobuf:"bytes,6,opt,name=major,proto3" json:"major,omitempty"`
	MajorDescription   string `protobuf:"bytes,7,opt,name=major_description,json=majorDescription,proto3" json:"major_description,omitempty"`
	DescriptionProfile string `protobuf:"bytes,8,opt,name=description_profile,json=descriptionProfile,proto3" json:"description_profile,omitempty"`
	MapCoordinate      string `protobuf:"bytes,9,opt,name=map_coordinate,json=mapCoordinate,proto3" json:"map_coordinate,omitempty"`
	ImageUrl           string `protobuf:"bytes,10,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	DisplayName        string `protobuf:"bytes,11,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *SeerResponse) Reset() {
	*x = SeerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeerResponse) ProtoMessage() {}

func (x *SeerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_seer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeerResponse.ProtoReflect.Descriptor instead.
func (*SeerResponse) Descriptor() ([]byte, []int) {
	return file_seer_proto_rawDescGZIP(), []int{2}
}

func (x *SeerResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *SeerResponse) GetOnsiteAvailable() bool {
	if x != nil {
		return x.OnsiteAvailable
	}
	return false
}

func (x *SeerResponse) GetChatAvailable() bool {
	if x != nil {
		return x.ChatAvailable
	}
	return false
}

func (x *SeerResponse) GetCallAvailable() bool {
	if x != nil {
		return x.CallAvailable
	}
	return false
}

func (x *SeerResponse) GetVideoCallAvailable() bool {
	if x != nil {
		return x.VideoCallAvailable
	}
	return false
}

func (x *SeerResponse) GetMajor() string {
	if x != nil {
		return x.Major
	}
	return ""
}

func (x *SeerResponse) GetMajorDescription() string {
	if x != nil {
		return x.MajorDescription
	}
	return ""
}

func (x *SeerResponse) GetDescriptionProfile() string {
	if x != nil {
		return x.DescriptionProfile
	}
	return ""
}

func (x *SeerResponse) GetMapCoordinate() string {
	if x != nil {
		return x.MapCoordinate
	}
	return ""
}

func (x *SeerResponse) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *SeerResponse) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

type UserUUIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid string `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
}

func (x *UserUUIDResponse) Reset() {
	*x = UserUUIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUUIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUUIDResponse) ProtoMessage() {}

func (x *UserUUIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_seer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUUIDResponse.ProtoReflect.Descriptor instead.
func (*UserUUIDResponse) Descriptor() ([]byte, []int) {
	return file_seer_proto_rawDescGZIP(), []int{3}
}

func (x *UserUUIDResponse) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

var File_seer_proto protoreflect.FileDescriptor

var file_seer_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x65,
	0x65, 0x72, 0x22, 0x2a, 0x0a, 0x0b, 0x53, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x22, 0x2e,
	0x0a, 0x0f, 0x53, 0x65, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x22, 0xa8,
	0x03, 0x0a, 0x0c, 0x53, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x6f, 0x6e, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x6f,
	0x6e, 0x73, 0x69, 0x74, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x74, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x63,
	0x61, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x14,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x43, 0x61, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d,
	0x61, 0x6a, 0x6f, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2f, 0x0a, 0x13, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x61, 0x70, 0x5f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x61, 0x70, 0x43,
	0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x10, 0x55, 0x73, 0x65,
	0x72, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x32, 0x91, 0x01, 0x0a, 0x0b, 0x53,
	0x65, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x12,
	0x11, 0x2e, 0x73, 0x65, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73, 0x65, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x55, 0x55, 0x49, 0x44, 0x42, 0x79, 0x53, 0x65, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x12,
	0x15, 0x2e, 0x73, 0x65, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x44,
	0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x6f,
	0x6c, 0x6e, 0x77, 0x7a, 0x61, 0x2f, 0x6d, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2f,
	0x73, 0x65, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x64,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_seer_proto_rawDescOnce sync.Once
	file_seer_proto_rawDescData = file_seer_proto_rawDesc
)

func file_seer_proto_rawDescGZIP() []byte {
	file_seer_proto_rawDescOnce.Do(func() {
		file_seer_proto_rawDescData = protoimpl.X.CompressGZIP(file_seer_proto_rawDescData)
	})
	return file_seer_proto_rawDescData
}

var file_seer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_seer_proto_goTypes = []interface{}{
	(*SeerRequest)(nil),      // 0: seer.SeerRequest
	(*SeerUUIDRequest)(nil),  // 1: seer.SeerUUIDRequest
	(*SeerResponse)(nil),     // 2: seer.SeerResponse
	(*UserUUIDResponse)(nil), // 3: seer.UserUUIDResponse
}
var file_seer_proto_depIdxs = []int32{
	0, // 0: seer.SeerService.GetSeerByUserUUID:input_type -> seer.SeerRequest
	1, // 1: seer.SeerService.GetUserUUIDBySeerUUID:input_type -> seer.SeerUUIDRequest
	2, // 2: seer.SeerService.GetSeerByUserUUID:output_type -> seer.SeerResponse
	3, // 3: seer.SeerService.GetUserUUIDBySeerUUID:output_type -> seer.UserUUIDResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_seer_proto_init() }
func file_seer_proto_init() {
	if File_seer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_seer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeerRequest); i {
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
		file_seer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeerUUIDRequest); i {
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
		file_seer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeerResponse); i {
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
		file_seer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUUIDResponse); i {
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
			RawDescriptor: file_seer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_seer_proto_goTypes,
		DependencyIndexes: file_seer_proto_depIdxs,
		MessageInfos:      file_seer_proto_msgTypes,
	}.Build()
	File_seer_proto = out.File
	file_seer_proto_rawDesc = nil
	file_seer_proto_goTypes = nil
	file_seer_proto_depIdxs = nil
}
