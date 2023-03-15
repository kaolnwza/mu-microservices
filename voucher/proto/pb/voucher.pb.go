// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: voucher.proto

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

type VoucherCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoucherCode string `protobuf:"bytes,1,opt,name=voucher_code,json=voucherCode,proto3" json:"voucher_code,omitempty"`
}

func (x *VoucherCodeRequest) Reset() {
	*x = VoucherCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voucher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoucherCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoucherCodeRequest) ProtoMessage() {}

func (x *VoucherCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_voucher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoucherCodeRequest.ProtoReflect.Descriptor instead.
func (*VoucherCodeRequest) Descriptor() ([]byte, []int) {
	return file_voucher_proto_rawDescGZIP(), []int{0}
}

func (x *VoucherCodeRequest) GetVoucherCode() string {
	if x != nil {
		return x.VoucherCode
	}
	return ""
}

type VoucherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoucherName     string `protobuf:"bytes,1,opt,name=voucher_name,json=voucherName,proto3" json:"voucher_name,omitempty"`
	DiscountType    string `protobuf:"bytes,2,opt,name=discount_type,json=discountType,proto3" json:"discount_type,omitempty"`
	Discount        int32  `protobuf:"varint,3,opt,name=discount,proto3" json:"discount,omitempty"`
	VoucherQuantity int32  `protobuf:"varint,4,opt,name=voucher_quantity,json=voucherQuantity,proto3" json:"voucher_quantity,omitempty"`
	ExpiredAt       string `protobuf:"bytes,5,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	Status          bool   `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *VoucherResponse) Reset() {
	*x = VoucherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voucher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoucherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoucherResponse) ProtoMessage() {}

func (x *VoucherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_voucher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoucherResponse.ProtoReflect.Descriptor instead.
func (*VoucherResponse) Descriptor() ([]byte, []int) {
	return file_voucher_proto_rawDescGZIP(), []int{1}
}

func (x *VoucherResponse) GetVoucherName() string {
	if x != nil {
		return x.VoucherName
	}
	return ""
}

func (x *VoucherResponse) GetDiscountType() string {
	if x != nil {
		return x.DiscountType
	}
	return ""
}

func (x *VoucherResponse) GetDiscount() int32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *VoucherResponse) GetVoucherQuantity() int32 {
	if x != nil {
		return x.VoucherQuantity
	}
	return 0
}

func (x *VoucherResponse) GetExpiredAt() string {
	if x != nil {
		return x.ExpiredAt
	}
	return ""
}

func (x *VoucherResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type VoucherStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *VoucherStatusResponse) Reset() {
	*x = VoucherStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voucher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoucherStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoucherStatusResponse) ProtoMessage() {}

func (x *VoucherStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_voucher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoucherStatusResponse.ProtoReflect.Descriptor instead.
func (*VoucherStatusResponse) Descriptor() ([]byte, []int) {
	return file_voucher_proto_rawDescGZIP(), []int{2}
}

func (x *VoucherStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_voucher_proto protoreflect.FileDescriptor

var file_voucher_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x22, 0x37, 0x0a, 0x12, 0x56, 0x6f, 0x75, 0x63,
	0x68, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0xd7, 0x01, 0x0a, 0x0f, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x6f, 0x75,
	0x63, 0x68, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x76, 0x6f, 0x75,
	0x63, 0x68, 0x65, 0x72, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0f, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2f, 0x0a, 0x15, 0x56,
	0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xaf, 0x01, 0x0a,
	0x0e, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x52, 0x0a, 0x13, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x75, 0x63, 0x68,
	0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x2e, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x56, 0x6f,
	0x75, 0x63, 0x68, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65,
	0x72, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x2e, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x56,
	0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30,
	0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x6f,
	0x6c, 0x6e, 0x77, 0x7a, 0x61, 0x2f, 0x6d, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2f,
	0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_voucher_proto_rawDescOnce sync.Once
	file_voucher_proto_rawDescData = file_voucher_proto_rawDesc
)

func file_voucher_proto_rawDescGZIP() []byte {
	file_voucher_proto_rawDescOnce.Do(func() {
		file_voucher_proto_rawDescData = protoimpl.X.CompressGZIP(file_voucher_proto_rawDescData)
	})
	return file_voucher_proto_rawDescData
}

var file_voucher_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_voucher_proto_goTypes = []interface{}{
	(*VoucherCodeRequest)(nil),    // 0: voucher.VoucherCodeRequest
	(*VoucherResponse)(nil),       // 1: voucher.VoucherResponse
	(*VoucherStatusResponse)(nil), // 2: voucher.VoucherStatusResponse
}
var file_voucher_proto_depIdxs = []int32{
	0, // 0: voucher.VoucherService.ValidateVoucherCode:input_type -> voucher.VoucherCodeRequest
	0, // 1: voucher.VoucherService.GetVoucherByCode:input_type -> voucher.VoucherCodeRequest
	2, // 2: voucher.VoucherService.ValidateVoucherCode:output_type -> voucher.VoucherStatusResponse
	1, // 3: voucher.VoucherService.GetVoucherByCode:output_type -> voucher.VoucherResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_voucher_proto_init() }
func file_voucher_proto_init() {
	if File_voucher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_voucher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoucherCodeRequest); i {
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
		file_voucher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoucherResponse); i {
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
		file_voucher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoucherStatusResponse); i {
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
			RawDescriptor: file_voucher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_voucher_proto_goTypes,
		DependencyIndexes: file_voucher_proto_depIdxs,
		MessageInfos:      file_voucher_proto_msgTypes,
	}.Build()
	File_voucher_proto = out.File
	file_voucher_proto_rawDesc = nil
	file_voucher_proto_goTypes = nil
	file_voucher_proto_depIdxs = nil
}
