// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: api/v1/cat_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateCatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cat *Cat `protobuf:"bytes,1,opt,name=cat,proto3" json:"cat,omitempty"`
}

func (x *CreateCatRequest) Reset() {
	*x = CreateCatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cat_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCatRequest) ProtoMessage() {}

func (x *CreateCatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cat_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCatRequest.ProtoReflect.Descriptor instead.
func (*CreateCatRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_cat_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCatRequest) GetCat() *Cat {
	if x != nil {
		return x.Cat
	}
	return nil
}

type CreateCatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cat *Cat `protobuf:"bytes,1,opt,name=cat,proto3" json:"cat,omitempty"`
}

func (x *CreateCatResponse) Reset() {
	*x = CreateCatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cat_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCatResponse) ProtoMessage() {}

func (x *CreateCatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cat_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCatResponse.ProtoReflect.Descriptor instead.
func (*CreateCatResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_cat_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCatResponse) GetCat() *Cat {
	if x != nil {
		return x.Cat
	}
	return nil
}

type ListCatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCatsRequest) Reset() {
	*x = ListCatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cat_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCatsRequest) ProtoMessage() {}

func (x *ListCatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cat_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCatsRequest.ProtoReflect.Descriptor instead.
func (*ListCatsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_cat_service_proto_rawDescGZIP(), []int{2}
}

type ListCatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cats []*Cat `protobuf:"bytes,1,rep,name=cats,proto3" json:"cats,omitempty"`
}

func (x *ListCatsResponse) Reset() {
	*x = ListCatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cat_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCatsResponse) ProtoMessage() {}

func (x *ListCatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cat_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCatsResponse.ProtoReflect.Descriptor instead.
func (*ListCatsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_cat_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListCatsResponse) GetCats() []*Cat {
	if x != nil {
		return x.Cats
	}
	return nil
}

type Cat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Fact string `protobuf:"bytes,3,opt,name=fact,proto3" json:"fact,omitempty"`
}

func (x *Cat) Reset() {
	*x = Cat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cat_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cat) ProtoMessage() {}

func (x *Cat) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cat_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cat.ProtoReflect.Descriptor instead.
func (*Cat) Descriptor() ([]byte, []int) {
	return file_api_v1_cat_service_proto_rawDescGZIP(), []int{4}
}

func (x *Cat) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Cat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cat) GetFact() string {
	if x != nil {
		return x.Fact
	}
	return ""
}

var File_api_v1_cat_service_proto protoreflect.FileDescriptor

var file_api_v1_cat_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x31, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x03, 0x63, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x03,
	0x63, 0x61, 0x74, 0x22, 0x32, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x03, 0x63, 0x61, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x61, 0x74, 0x52, 0x03, 0x63, 0x61, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x04, 0x63, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x04, 0x63, 0x61, 0x74, 0x73, 0x22,
	0x3d, 0x0a, 0x03, 0x43, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x61,
	0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x61, 0x63, 0x74, 0x32, 0xb6,
	0x01, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74,
	0x73, 0x3a, 0x03, 0x63, 0x61, 0x74, 0x12, 0x4f, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61,
	0x74, 0x73, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x73, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x65, 0x6e, 0x72, 0x6f, 0x64, 0x2f, 0x6f, 0x70, 0x65,
	0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_cat_service_proto_rawDescOnce sync.Once
	file_api_v1_cat_service_proto_rawDescData = file_api_v1_cat_service_proto_rawDesc
)

func file_api_v1_cat_service_proto_rawDescGZIP() []byte {
	file_api_v1_cat_service_proto_rawDescOnce.Do(func() {
		file_api_v1_cat_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_cat_service_proto_rawDescData)
	})
	return file_api_v1_cat_service_proto_rawDescData
}

var file_api_v1_cat_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1_cat_service_proto_goTypes = []interface{}{
	(*CreateCatRequest)(nil),  // 0: api.v1.CreateCatRequest
	(*CreateCatResponse)(nil), // 1: api.v1.CreateCatResponse
	(*ListCatsRequest)(nil),   // 2: api.v1.ListCatsRequest
	(*ListCatsResponse)(nil),  // 3: api.v1.ListCatsResponse
	(*Cat)(nil),               // 4: api.v1.Cat
}
var file_api_v1_cat_service_proto_depIdxs = []int32{
	4, // 0: api.v1.CreateCatRequest.cat:type_name -> api.v1.Cat
	4, // 1: api.v1.CreateCatResponse.cat:type_name -> api.v1.Cat
	4, // 2: api.v1.ListCatsResponse.cats:type_name -> api.v1.Cat
	0, // 3: api.v1.CatService.CreateCat:input_type -> api.v1.CreateCatRequest
	2, // 4: api.v1.CatService.ListCats:input_type -> api.v1.ListCatsRequest
	1, // 5: api.v1.CatService.CreateCat:output_type -> api.v1.CreateCatResponse
	3, // 6: api.v1.CatService.ListCats:output_type -> api.v1.ListCatsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_v1_cat_service_proto_init() }
func file_api_v1_cat_service_proto_init() {
	if File_api_v1_cat_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_cat_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCatRequest); i {
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
		file_api_v1_cat_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCatResponse); i {
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
		file_api_v1_cat_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCatsRequest); i {
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
		file_api_v1_cat_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCatsResponse); i {
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
		file_api_v1_cat_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cat); i {
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
			RawDescriptor: file_api_v1_cat_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_cat_service_proto_goTypes,
		DependencyIndexes: file_api_v1_cat_service_proto_depIdxs,
		MessageInfos:      file_api_v1_cat_service_proto_msgTypes,
	}.Build()
	File_api_v1_cat_service_proto = out.File
	file_api_v1_cat_service_proto_rawDesc = nil
	file_api_v1_cat_service_proto_goTypes = nil
	file_api_v1_cat_service_proto_depIdxs = nil
}
