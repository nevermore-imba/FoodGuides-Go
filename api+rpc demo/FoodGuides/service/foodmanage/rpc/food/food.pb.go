// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: food.proto

package food

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

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SearchKey string `protobuf:"bytes,1,opt,name=searchKey,proto3" json:"searchKey,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{0}
}

func (x *SearchRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

type AddFoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid string `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
	FoodId string `protobuf:"bytes,2,opt,name=foodId,proto3" json:"foodId,omitempty"`
}

func (x *AddFoodRequest) Reset() {
	*x = AddFoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFoodRequest) ProtoMessage() {}

func (x *AddFoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFoodRequest.ProtoReflect.Descriptor instead.
func (*AddFoodRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{1}
}

func (x *AddFoodRequest) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

func (x *AddFoodRequest) GetFoodId() string {
	if x != nil {
		return x.FoodId
	}
	return ""
}

type DeleteFoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid string `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
	FoodId string `protobuf:"bytes,2,opt,name=foodId,proto3" json:"foodId,omitempty"`
}

func (x *DeleteFoodRequest) Reset() {
	*x = DeleteFoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFoodRequest) ProtoMessage() {}

func (x *DeleteFoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFoodRequest.ProtoReflect.Descriptor instead.
func (*DeleteFoodRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteFoodRequest) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

func (x *DeleteFoodRequest) GetFoodId() string {
	if x != nil {
		return x.FoodId
	}
	return ""
}

type FoodListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid string `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
}

func (x *FoodListRequest) Reset() {
	*x = FoodListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodListRequest) ProtoMessage() {}

func (x *FoodListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodListRequest.ProtoReflect.Descriptor instead.
func (*FoodListRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{3}
}

func (x *FoodListRequest) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

type FoodInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protein      string `protobuf:"bytes,1,opt,name=protein,proto3" json:"protein,omitempty"`
	Fat          string `protobuf:"bytes,2,opt,name=fat,proto3" json:"fat,omitempty"`
	Carbohydrate string `protobuf:"bytes,3,opt,name=carbohydrate,proto3" json:"carbohydrate,omitempty"`
	Calorie      string `protobuf:"bytes,4,opt,name=calorie,proto3" json:"calorie,omitempty"`
	Minerals     string `protobuf:"bytes,5,opt,name=minerals,proto3" json:"minerals,omitempty"`
	Calcium      string `protobuf:"bytes,6,opt,name=calcium,proto3" json:"calcium,omitempty"`
	Phosphorus   string `protobuf:"bytes,7,opt,name=phosphorus,proto3" json:"phosphorus,omitempty"`
	Iron         string `protobuf:"bytes,8,opt,name=iron,proto3" json:"iron,omitempty"`
	Purine       string `protobuf:"bytes,9,opt,name=purine,proto3" json:"purine,omitempty"`
	Id           string `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FoodInfoResponse) Reset() {
	*x = FoodInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodInfoResponse) ProtoMessage() {}

func (x *FoodInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodInfoResponse.ProtoReflect.Descriptor instead.
func (*FoodInfoResponse) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{4}
}

func (x *FoodInfoResponse) GetProtein() string {
	if x != nil {
		return x.Protein
	}
	return ""
}

func (x *FoodInfoResponse) GetFat() string {
	if x != nil {
		return x.Fat
	}
	return ""
}

func (x *FoodInfoResponse) GetCarbohydrate() string {
	if x != nil {
		return x.Carbohydrate
	}
	return ""
}

func (x *FoodInfoResponse) GetCalorie() string {
	if x != nil {
		return x.Calorie
	}
	return ""
}

func (x *FoodInfoResponse) GetMinerals() string {
	if x != nil {
		return x.Minerals
	}
	return ""
}

func (x *FoodInfoResponse) GetCalcium() string {
	if x != nil {
		return x.Calcium
	}
	return ""
}

func (x *FoodInfoResponse) GetPhosphorus() string {
	if x != nil {
		return x.Phosphorus
	}
	return ""
}

func (x *FoodInfoResponse) GetIron() string {
	if x != nil {
		return x.Iron
	}
	return ""
}

func (x *FoodInfoResponse) GetPurine() string {
	if x != nil {
		return x.Purine
	}
	return ""
}

func (x *FoodInfoResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FoodInfoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success int32 `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{5}
}

func (x *StatusResponse) GetSuccess() int32 {
	if x != nil {
		return x.Success
	}
	return 0
}

type FoodListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*FoodInfoResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *FoodListResponse) Reset() {
	*x = FoodListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodListResponse) ProtoMessage() {}

func (x *FoodListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodListResponse.ProtoReflect.Descriptor instead.
func (*FoodListResponse) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{6}
}

func (x *FoodListResponse) GetData() []*FoodInfoResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_food_proto protoreflect.FileDescriptor

var file_food_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x6f,
	0x6f, 0x64, 0x22, 0x2d, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65,
	0x79, 0x22, 0x40, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x6f, 0x6f, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x6f,
	0x64, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x6f, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x66, 0x6f, 0x6f, 0x64, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x0f, 0x46, 0x6f, 0x6f, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x69, 0x64, 0x22, 0xa2, 0x02, 0x0a, 0x10, 0x46, 0x6f, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x74,
	0x65, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x74, 0x65,
	0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x66, 0x61, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x68, 0x79, 0x64,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x72, 0x62,
	0x6f, 0x68, 0x79, 0x64, 0x72, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x6c, 0x6f,
	0x72, 0x69, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x6c, 0x6f, 0x72,
	0x69, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x61, 0x6c, 0x63, 0x69, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x61, 0x6c, 0x63, 0x69, 0x75, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x68, 0x6f, 0x73,
	0x70, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x68,
	0x6f, 0x73, 0x70, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x72, 0x6f, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x72, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x75, 0x72, 0x69, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x75,
	0x72, 0x69, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x3e, 0x0a, 0x10, 0x46, 0x6f, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f,
	0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x32, 0xec, 0x01, 0x0a, 0x04, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x35, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x13, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x66,
	0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6f, 0x64, 0x12,
	0x14, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x17, 0x2e, 0x66, 0x6f, 0x6f, 0x64,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x46, 0x6f, 0x6f, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x66, 0x6f,
	0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x66, 0x6f, 0x6f, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_food_proto_rawDescOnce sync.Once
	file_food_proto_rawDescData = file_food_proto_rawDesc
)

func file_food_proto_rawDescGZIP() []byte {
	file_food_proto_rawDescOnce.Do(func() {
		file_food_proto_rawDescData = protoimpl.X.CompressGZIP(file_food_proto_rawDescData)
	})
	return file_food_proto_rawDescData
}

var file_food_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_food_proto_goTypes = []interface{}{
	(*SearchRequest)(nil),     // 0: food.SearchRequest
	(*AddFoodRequest)(nil),    // 1: food.AddFoodRequest
	(*DeleteFoodRequest)(nil), // 2: food.DeleteFoodRequest
	(*FoodListRequest)(nil),   // 3: food.FoodListRequest
	(*FoodInfoResponse)(nil),  // 4: food.FoodInfoResponse
	(*StatusResponse)(nil),    // 5: food.StatusResponse
	(*FoodListResponse)(nil),  // 6: food.FoodListResponse
}
var file_food_proto_depIdxs = []int32{
	4, // 0: food.FoodListResponse.data:type_name -> food.FoodInfoResponse
	0, // 1: food.Food.Search:input_type -> food.SearchRequest
	1, // 2: food.Food.AddFood:input_type -> food.AddFoodRequest
	2, // 3: food.Food.DeleteFood:input_type -> food.DeleteFoodRequest
	3, // 4: food.Food.FoodList:input_type -> food.FoodListRequest
	4, // 5: food.Food.Search:output_type -> food.FoodInfoResponse
	5, // 6: food.Food.AddFood:output_type -> food.StatusResponse
	5, // 7: food.Food.DeleteFood:output_type -> food.StatusResponse
	6, // 8: food.Food.FoodList:output_type -> food.FoodListResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_food_proto_init() }
func file_food_proto_init() {
	if File_food_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_food_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_food_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFoodRequest); i {
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
		file_food_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFoodRequest); i {
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
		file_food_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodListRequest); i {
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
		file_food_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodInfoResponse); i {
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
		file_food_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_food_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodListResponse); i {
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
			RawDescriptor: file_food_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_food_proto_goTypes,
		DependencyIndexes: file_food_proto_depIdxs,
		MessageInfos:      file_food_proto_msgTypes,
	}.Build()
	File_food_proto = out.File
	file_food_proto_rawDesc = nil
	file_food_proto_goTypes = nil
	file_food_proto_depIdxs = nil
}
