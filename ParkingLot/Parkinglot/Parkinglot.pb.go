// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: Parkinglot/Parkinglot.proto

package go_Parkinglot_grpc

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

type CarDetailsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarName  string `protobuf:"bytes,1,opt,name=CarName,proto3" json:"CarName,omitempty"`
	CarId    string `protobuf:"bytes,2,opt,name=CarId,proto3" json:"CarId,omitempty"`
	CarColor string `protobuf:"bytes,3,opt,name=CarColor,proto3" json:"CarColor,omitempty"`
}

func (x *CarDetailsInfo) Reset() {
	*x = CarDetailsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Parkinglot_Parkinglot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarDetailsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarDetailsInfo) ProtoMessage() {}

func (x *CarDetailsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Parkinglot_Parkinglot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarDetailsInfo.ProtoReflect.Descriptor instead.
func (*CarDetailsInfo) Descriptor() ([]byte, []int) {
	return file_Parkinglot_Parkinglot_proto_rawDescGZIP(), []int{0}
}

func (x *CarDetailsInfo) GetCarName() string {
	if x != nil {
		return x.CarName
	}
	return ""
}

func (x *CarDetailsInfo) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

func (x *CarDetailsInfo) GetCarColor() string {
	if x != nil {
		return x.CarColor
	}
	return ""
}

type SlotInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarName  string `protobuf:"bytes,1,opt,name=CarName,proto3" json:"CarName,omitempty"`
	CarId    string `protobuf:"bytes,2,opt,name=CarId,proto3" json:"CarId,omitempty"`
	CarColor string `protobuf:"bytes,3,opt,name=CarColor,proto3" json:"CarColor,omitempty"`
	Slot     int32  `protobuf:"varint,4,opt,name=Slot,proto3" json:"Slot,omitempty"`
}

func (x *SlotInfo) Reset() {
	*x = SlotInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Parkinglot_Parkinglot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlotInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlotInfo) ProtoMessage() {}

func (x *SlotInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Parkinglot_Parkinglot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlotInfo.ProtoReflect.Descriptor instead.
func (*SlotInfo) Descriptor() ([]byte, []int) {
	return file_Parkinglot_Parkinglot_proto_rawDescGZIP(), []int{1}
}

func (x *SlotInfo) GetCarName() string {
	if x != nil {
		return x.CarName
	}
	return ""
}

func (x *SlotInfo) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

func (x *SlotInfo) GetCarColor() string {
	if x != nil {
		return x.CarColor
	}
	return ""
}

func (x *SlotInfo) GetSlot() int32 {
	if x != nil {
		return x.Slot
	}
	return 0
}

type DetailsInfoParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DetailsInfoParams) Reset() {
	*x = DetailsInfoParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Parkinglot_Parkinglot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailsInfoParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailsInfoParams) ProtoMessage() {}

func (x *DetailsInfoParams) ProtoReflect() protoreflect.Message {
	mi := &file_Parkinglot_Parkinglot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailsInfoParams.ProtoReflect.Descriptor instead.
func (*DetailsInfoParams) Descriptor() ([]byte, []int) {
	return file_Parkinglot_Parkinglot_proto_rawDescGZIP(), []int{2}
}

type CarDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details []*SlotInfo `protobuf:"bytes,1,rep,name=Details,proto3" json:"Details,omitempty"`
}

func (x *CarDetails) Reset() {
	*x = CarDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Parkinglot_Parkinglot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarDetails) ProtoMessage() {}

func (x *CarDetails) ProtoReflect() protoreflect.Message {
	mi := &file_Parkinglot_Parkinglot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarDetails.ProtoReflect.Descriptor instead.
func (*CarDetails) Descriptor() ([]byte, []int) {
	return file_Parkinglot_Parkinglot_proto_rawDescGZIP(), []int{3}
}

func (x *CarDetails) GetDetails() []*SlotInfo {
	if x != nil {
		return x.Details
	}
	return nil
}

type GetId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarId string `protobuf:"bytes,1,opt,name=CarId,proto3" json:"CarId,omitempty"`
}

func (x *GetId) Reset() {
	*x = GetId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Parkinglot_Parkinglot_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetId) ProtoMessage() {}

func (x *GetId) ProtoReflect() protoreflect.Message {
	mi := &file_Parkinglot_Parkinglot_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetId.ProtoReflect.Descriptor instead.
func (*GetId) Descriptor() ([]byte, []int) {
	return file_Parkinglot_Parkinglot_proto_rawDescGZIP(), []int{4}
}

func (x *GetId) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

var File_Parkinglot_Parkinglot_proto protoreflect.FileDescriptor

var file_Parkinglot_Parkinglot_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x6c, 0x6f, 0x74, 0x2f, 0x50, 0x61, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x6c, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x50,
	0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x74, 0x5f, 0x4d, 0x61, 0x69, 0x6e, 0x22, 0x5c,
	0x0a, 0x0e, 0x43, 0x61, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x61,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x61, 0x72, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x43, 0x61, 0x72, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x6a, 0x0a, 0x08,
	0x53, 0x6c, 0x6f, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x61, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x61, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x43, 0x61, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x61, 0x72, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x6c, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x53, 0x6c, 0x6f, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x41, 0x0a,
	0x0a, 0x43, 0x61, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x50,
	0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x74, 0x5f, 0x4d, 0x61, 0x69, 0x6e, 0x2e, 0x53,
	0x6c, 0x6f, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x22, 0x1d, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x61, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x61, 0x72, 0x49, 0x64, 0x32,
	0xf2, 0x01, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x69, 0x67, 0x4c, 0x6f, 0x74, 0x12,
	0x4a, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x6b, 0x4e, 0x65, 0x77, 0x43, 0x61, 0x72, 0x12, 0x1f, 0x2e,
	0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x74, 0x5f, 0x4d, 0x61, 0x69, 0x6e, 0x2e,
	0x43, 0x61, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x19,
	0x2e, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x74, 0x5f, 0x4d, 0x61, 0x69, 0x6e,
	0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0e, 0x50,
	0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x22, 0x2e,
	0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x74, 0x5f, 0x4d, 0x61, 0x69, 0x6e, 0x2e,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x1a, 0x1b, 0x2e, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x74, 0x5f, 0x4d,
	0x61, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00,
	0x12, 0x42, 0x0a, 0x09, 0x55, 0x6e, 0x70, 0x61, 0x72, 0x6b, 0x43, 0x61, 0x72, 0x12, 0x16, 0x2e,
	0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x74, 0x5f, 0x4d, 0x61, 0x69, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x64, 0x1a, 0x1b, 0x2e, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x4c,
	0x6f, 0x74, 0x5f, 0x4d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x22, 0x00, 0x42, 0x33, 0x5a, 0x31, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x6c, 0x6f,
	0x74, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x67, 0x6f, 0x5f, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x6c, 0x6f, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_Parkinglot_Parkinglot_proto_rawDescOnce sync.Once
	file_Parkinglot_Parkinglot_proto_rawDescData = file_Parkinglot_Parkinglot_proto_rawDesc
)

func file_Parkinglot_Parkinglot_proto_rawDescGZIP() []byte {
	file_Parkinglot_Parkinglot_proto_rawDescOnce.Do(func() {
		file_Parkinglot_Parkinglot_proto_rawDescData = protoimpl.X.CompressGZIP(file_Parkinglot_Parkinglot_proto_rawDescData)
	})
	return file_Parkinglot_Parkinglot_proto_rawDescData
}

var file_Parkinglot_Parkinglot_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Parkinglot_Parkinglot_proto_goTypes = []interface{}{
	(*CarDetailsInfo)(nil),    // 0: ParkingLot_Main.CarDetailsInfo
	(*SlotInfo)(nil),          // 1: ParkingLot_Main.SlotInfo
	(*DetailsInfoParams)(nil), // 2: ParkingLot_Main.DetailsInfoParams
	(*CarDetails)(nil),        // 3: ParkingLot_Main.CarDetails
	(*GetId)(nil),             // 4: ParkingLot_Main.GetId
}
var file_Parkinglot_Parkinglot_proto_depIdxs = []int32{
	1, // 0: ParkingLot_Main.CarDetails.Details:type_name -> ParkingLot_Main.SlotInfo
	0, // 1: ParkingLot_Main.ParkinigLot.ParkNewCar:input_type -> ParkingLot_Main.CarDetailsInfo
	2, // 2: ParkingLot_Main.ParkinigLot.ParkingDetails:input_type -> ParkingLot_Main.DetailsInfoParams
	4, // 3: ParkingLot_Main.ParkinigLot.UnparkCar:input_type -> ParkingLot_Main.GetId
	1, // 4: ParkingLot_Main.ParkinigLot.ParkNewCar:output_type -> ParkingLot_Main.SlotInfo
	3, // 5: ParkingLot_Main.ParkinigLot.ParkingDetails:output_type -> ParkingLot_Main.CarDetails
	3, // 6: ParkingLot_Main.ParkinigLot.UnparkCar:output_type -> ParkingLot_Main.CarDetails
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Parkinglot_Parkinglot_proto_init() }
func file_Parkinglot_Parkinglot_proto_init() {
	if File_Parkinglot_Parkinglot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Parkinglot_Parkinglot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarDetailsInfo); i {
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
		file_Parkinglot_Parkinglot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlotInfo); i {
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
		file_Parkinglot_Parkinglot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailsInfoParams); i {
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
		file_Parkinglot_Parkinglot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarDetails); i {
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
		file_Parkinglot_Parkinglot_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetId); i {
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
			RawDescriptor: file_Parkinglot_Parkinglot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Parkinglot_Parkinglot_proto_goTypes,
		DependencyIndexes: file_Parkinglot_Parkinglot_proto_depIdxs,
		MessageInfos:      file_Parkinglot_Parkinglot_proto_msgTypes,
	}.Build()
	File_Parkinglot_Parkinglot_proto = out.File
	file_Parkinglot_Parkinglot_proto_rawDesc = nil
	file_Parkinglot_Parkinglot_proto_goTypes = nil
	file_Parkinglot_Parkinglot_proto_depIdxs = nil
}
