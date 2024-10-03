// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: trainer.proto

package models

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

type Trainer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Specialization string `protobuf:"bytes,3,opt,name=specialization,proto3" json:"specialization,omitempty"`
}

func (x *Trainer) Reset() {
	*x = Trainer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trainer) ProtoMessage() {}

func (x *Trainer) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trainer.ProtoReflect.Descriptor instead.
func (*Trainer) Descriptor() ([]byte, []int) {
	return file_trainer_proto_rawDescGZIP(), []int{0}
}

func (x *Trainer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Trainer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Trainer) GetSpecialization() string {
	if x != nil {
		return x.Specialization
	}
	return ""
}

type CreateTrainerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Specialization string `protobuf:"bytes,2,opt,name=specialization,proto3" json:"specialization,omitempty"`
}

func (x *CreateTrainerRequest) Reset() {
	*x = CreateTrainerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTrainerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTrainerRequest) ProtoMessage() {}

func (x *CreateTrainerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTrainerRequest.ProtoReflect.Descriptor instead.
func (*CreateTrainerRequest) Descriptor() ([]byte, []int) {
	return file_trainer_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTrainerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTrainerRequest) GetSpecialization() string {
	if x != nil {
		return x.Specialization
	}
	return ""
}

type GetTrainerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTrainerRequest) Reset() {
	*x = GetTrainerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrainerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrainerRequest) ProtoMessage() {}

func (x *GetTrainerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrainerRequest.ProtoReflect.Descriptor instead.
func (*GetTrainerRequest) Descriptor() ([]byte, []int) {
	return file_trainer_proto_rawDescGZIP(), []int{2}
}

func (x *GetTrainerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListTrainersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListTrainersRequest) Reset() {
	*x = ListTrainersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTrainersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTrainersRequest) ProtoMessage() {}

func (x *ListTrainersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTrainersRequest.ProtoReflect.Descriptor instead.
func (*ListTrainersRequest) Descriptor() ([]byte, []int) {
	return file_trainer_proto_rawDescGZIP(), []int{3}
}

func (x *ListTrainersRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListTrainersRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListTrainersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trainers []*Trainer `protobuf:"bytes,1,rep,name=trainers,proto3" json:"trainers,omitempty"`
	Total    int32      `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListTrainersResponse) Reset() {
	*x = ListTrainersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTrainersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTrainersResponse) ProtoMessage() {}

func (x *ListTrainersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTrainersResponse.ProtoReflect.Descriptor instead.
func (*ListTrainersResponse) Descriptor() ([]byte, []int) {
	return file_trainer_proto_rawDescGZIP(), []int{4}
}

func (x *ListTrainersResponse) GetTrainers() []*Trainer {
	if x != nil {
		return x.Trainers
	}
	return nil
}

func (x *ListTrainersResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_trainer_proto protoreflect.FileDescriptor

var file_trainer_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x22, 0x55, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x52, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x22, 0x5a, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x08, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xdb, 0x01, 0x0a,
	0x0e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x40, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x12, 0x1d, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x12, 0x3a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12,
	0x1a, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x4b, 0x0a,
	0x0c, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x64, 0x61, 0x72, 0x6c, 0x69, 0x2f,
	0x67, 0x79, 0x6d, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_trainer_proto_rawDescOnce sync.Once
	file_trainer_proto_rawDescData = file_trainer_proto_rawDesc
)

func file_trainer_proto_rawDescGZIP() []byte {
	file_trainer_proto_rawDescOnce.Do(func() {
		file_trainer_proto_rawDescData = protoimpl.X.CompressGZIP(file_trainer_proto_rawDescData)
	})
	return file_trainer_proto_rawDescData
}

var file_trainer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_trainer_proto_goTypes = []any{
	(*Trainer)(nil),              // 0: trainer.Trainer
	(*CreateTrainerRequest)(nil), // 1: trainer.CreateTrainerRequest
	(*GetTrainerRequest)(nil),    // 2: trainer.GetTrainerRequest
	(*ListTrainersRequest)(nil),  // 3: trainer.ListTrainersRequest
	(*ListTrainersResponse)(nil), // 4: trainer.ListTrainersResponse
}
var file_trainer_proto_depIdxs = []int32{
	0, // 0: trainer.ListTrainersResponse.trainers:type_name -> trainer.Trainer
	1, // 1: trainer.TrainerService.CreateTrainer:input_type -> trainer.CreateTrainerRequest
	2, // 2: trainer.TrainerService.GetTrainer:input_type -> trainer.GetTrainerRequest
	3, // 3: trainer.TrainerService.ListTrainers:input_type -> trainer.ListTrainersRequest
	0, // 4: trainer.TrainerService.CreateTrainer:output_type -> trainer.Trainer
	0, // 5: trainer.TrainerService.GetTrainer:output_type -> trainer.Trainer
	4, // 6: trainer.TrainerService.ListTrainers:output_type -> trainer.ListTrainersResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_trainer_proto_init() }
func file_trainer_proto_init() {
	if File_trainer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trainer_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Trainer); i {
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
		file_trainer_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTrainerRequest); i {
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
		file_trainer_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetTrainerRequest); i {
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
		file_trainer_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListTrainersRequest); i {
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
		file_trainer_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListTrainersResponse); i {
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
			RawDescriptor: file_trainer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trainer_proto_goTypes,
		DependencyIndexes: file_trainer_proto_depIdxs,
		MessageInfos:      file_trainer_proto_msgTypes,
	}.Build()
	File_trainer_proto = out.File
	file_trainer_proto_rawDesc = nil
	file_trainer_proto_goTypes = nil
	file_trainer_proto_depIdxs = nil
}