// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: orkestraerp/orkestraerp.proto

package orkestraerp

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

type GetPageParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Period     int32    `protobuf:"varint,1,opt,name=Period,proto3" json:"Period,omitempty"`
	InputType  *string  `protobuf:"bytes,2,opt,name=InputType,proto3,oneof" json:"InputType,omitempty"`
	OutputType string   `protobuf:"bytes,3,opt,name=OutputType,proto3" json:"OutputType,omitempty"`
	EntityName string   `protobuf:"bytes,4,opt,name=EntityName,proto3" json:"EntityName,omitempty"`
	ClassType  int32    `protobuf:"varint,5,opt,name=ClassType,proto3" json:"ClassType,omitempty"`
	Fields     *string  `protobuf:"bytes,6,opt,name=Fields,proto3,oneof" json:"Fields,omitempty"`
	Reference  int32    `protobuf:"varint,7,opt,name=Reference,proto3" json:"Reference,omitempty"`
	PageSize   *int32   `protobuf:"varint,8,opt,name=PageSize,proto3,oneof" json:"PageSize,omitempty"`
	PageIndex  *int32   `protobuf:"varint,9,opt,name=PageIndex,proto3,oneof" json:"PageIndex,omitempty"`
	Filters    *Filters `protobuf:"bytes,10,opt,name=Filters,proto3,oneof" json:"Filters,omitempty"`
	OrderBy    *OrderBy `protobuf:"bytes,11,opt,name=OrderBy,proto3,oneof" json:"OrderBy,omitempty"`
}

func (x *GetPageParams) Reset() {
	*x = GetPageParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orkestraerp_orkestraerp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPageParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPageParams) ProtoMessage() {}

func (x *GetPageParams) ProtoReflect() protoreflect.Message {
	mi := &file_orkestraerp_orkestraerp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPageParams.ProtoReflect.Descriptor instead.
func (*GetPageParams) Descriptor() ([]byte, []int) {
	return file_orkestraerp_orkestraerp_proto_rawDescGZIP(), []int{0}
}

func (x *GetPageParams) GetPeriod() int32 {
	if x != nil {
		return x.Period
	}
	return 0
}

func (x *GetPageParams) GetInputType() string {
	if x != nil && x.InputType != nil {
		return *x.InputType
	}
	return ""
}

func (x *GetPageParams) GetOutputType() string {
	if x != nil {
		return x.OutputType
	}
	return ""
}

func (x *GetPageParams) GetEntityName() string {
	if x != nil {
		return x.EntityName
	}
	return ""
}

func (x *GetPageParams) GetClassType() int32 {
	if x != nil {
		return x.ClassType
	}
	return 0
}

func (x *GetPageParams) GetFields() string {
	if x != nil && x.Fields != nil {
		return *x.Fields
	}
	return ""
}

func (x *GetPageParams) GetReference() int32 {
	if x != nil {
		return x.Reference
	}
	return 0
}

func (x *GetPageParams) GetPageSize() int32 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

func (x *GetPageParams) GetPageIndex() int32 {
	if x != nil && x.PageIndex != nil {
		return *x.PageIndex
	}
	return 0
}

func (x *GetPageParams) GetFilters() *Filters {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *GetPageParams) GetOrderBy() *OrderBy {
	if x != nil {
		return x.OrderBy
	}
	return nil
}

type ResponsePage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Xml string `protobuf:"bytes,1,opt,name=xml,proto3" json:"xml,omitempty"`
}

func (x *ResponsePage) Reset() {
	*x = ResponsePage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orkestraerp_orkestraerp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponsePage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponsePage) ProtoMessage() {}

func (x *ResponsePage) ProtoReflect() protoreflect.Message {
	mi := &file_orkestraerp_orkestraerp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponsePage.ProtoReflect.Descriptor instead.
func (*ResponsePage) Descriptor() ([]byte, []int) {
	return file_orkestraerp_orkestraerp_proto_rawDescGZIP(), []int{1}
}

func (x *ResponsePage) GetXml() string {
	if x != nil {
		return x.Xml
	}
	return ""
}

type Filters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        *string `protobuf:"bytes,1,opt,name=Type,proto3,oneof" json:"Type,omitempty"`
	NumValue    *int32  `protobuf:"varint,2,opt,name=NumValue,proto3,oneof" json:"NumValue,omitempty"`
	StringValue *string `protobuf:"bytes,3,opt,name=StringValue,proto3,oneof" json:"StringValue,omitempty"`
	Property    *string `protobuf:"bytes,4,opt,name=Property,proto3,oneof" json:"Property,omitempty"`
	Operator    *string `protobuf:"bytes,5,opt,name=Operator,proto3,oneof" json:"Operator,omitempty"`
}

func (x *Filters) Reset() {
	*x = Filters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orkestraerp_orkestraerp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filters) ProtoMessage() {}

func (x *Filters) ProtoReflect() protoreflect.Message {
	mi := &file_orkestraerp_orkestraerp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filters.ProtoReflect.Descriptor instead.
func (*Filters) Descriptor() ([]byte, []int) {
	return file_orkestraerp_orkestraerp_proto_rawDescGZIP(), []int{2}
}

func (x *Filters) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

func (x *Filters) GetNumValue() int32 {
	if x != nil && x.NumValue != nil {
		return *x.NumValue
	}
	return 0
}

func (x *Filters) GetStringValue() string {
	if x != nil && x.StringValue != nil {
		return *x.StringValue
	}
	return ""
}

func (x *Filters) GetProperty() string {
	if x != nil && x.Property != nil {
		return *x.Property
	}
	return ""
}

func (x *Filters) GetOperator() string {
	if x != nil && x.Operator != nil {
		return *x.Operator
	}
	return ""
}

type OrderBy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Property *string `protobuf:"bytes,1,opt,name=Property,proto3,oneof" json:"Property,omitempty"`
	Desc     bool    `protobuf:"varint,2,opt,name=Desc,proto3" json:"Desc,omitempty"`
}

func (x *OrderBy) Reset() {
	*x = OrderBy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orkestraerp_orkestraerp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderBy) ProtoMessage() {}

func (x *OrderBy) ProtoReflect() protoreflect.Message {
	mi := &file_orkestraerp_orkestraerp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderBy.ProtoReflect.Descriptor instead.
func (*OrderBy) Descriptor() ([]byte, []int) {
	return file_orkestraerp_orkestraerp_proto_rawDescGZIP(), []int{3}
}

func (x *OrderBy) GetProperty() string {
	if x != nil && x.Property != nil {
		return *x.Property
	}
	return ""
}

func (x *OrderBy) GetDesc() bool {
	if x != nil {
		return x.Desc
	}
	return false
}

var File_orkestraerp_orkestraerp_proto protoreflect.FileDescriptor

var file_orkestraerp_orkestraerp_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6f, 0x72, 0x6b, 0x65, 0x73, 0x74, 0x72, 0x61, 0x65, 0x72, 0x70, 0x2f, 0x6f, 0x72,
	0x6b, 0x65, 0x73, 0x74, 0x72, 0x61, 0x65, 0x72, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x6f, 0x72, 0x6b, 0x65, 0x73, 0x74, 0x72, 0x61, 0x65, 0x72, 0x70, 0x22, 0xdd, 0x03, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x21, 0x0a, 0x09, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x0a, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x09, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x07, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x72, 0x6b, 0x65, 0x73, 0x74,
	0x72, 0x61, 0x65, 0x72, 0x70, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x48, 0x04, 0x52,
	0x07, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x07, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f,
	0x72, 0x6b, 0x65, 0x73, 0x74, 0x72, 0x61, 0x65, 0x72, 0x70, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x42, 0x79, 0x48, 0x05, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x88, 0x01, 0x01,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x50, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x50, 0x61, 0x67, 0x65, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x22, 0x20, 0x0a, 0x0c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x78, 0x6d, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x78, 0x6d, 0x6c, 0x22, 0xec,
	0x01, 0x0a, 0x07, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x17, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x4e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x08, 0x4e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0b, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x50,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52,
	0x08, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04,
	0x52, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x4e, 0x75, 0x6d, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x4b, 0x0a,
	0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x1f, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x65, 0x73,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x44, 0x65, 0x73, 0x63, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x32, 0x51, 0x0a, 0x0b, 0x4f, 0x72,
	0x6b, 0x65, 0x73, 0x74, 0x72, 0x61, 0x45, 0x72, 0x70, 0x12, 0x42, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x2e, 0x6f, 0x72, 0x6b, 0x65, 0x73, 0x74, 0x72, 0x61, 0x65,
	0x72, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x1a, 0x19, 0x2e, 0x6f, 0x72, 0x6b, 0x65, 0x73, 0x74, 0x72, 0x61, 0x65, 0x72, 0x70, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x34, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x6d, 0x65,
	0x72, 0x70, 0x73, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x6b, 0x65, 0x73, 0x74, 0x72, 0x61,
	0x65, 0x72, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orkestraerp_orkestraerp_proto_rawDescOnce sync.Once
	file_orkestraerp_orkestraerp_proto_rawDescData = file_orkestraerp_orkestraerp_proto_rawDesc
)

func file_orkestraerp_orkestraerp_proto_rawDescGZIP() []byte {
	file_orkestraerp_orkestraerp_proto_rawDescOnce.Do(func() {
		file_orkestraerp_orkestraerp_proto_rawDescData = protoimpl.X.CompressGZIP(file_orkestraerp_orkestraerp_proto_rawDescData)
	})
	return file_orkestraerp_orkestraerp_proto_rawDescData
}

var file_orkestraerp_orkestraerp_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_orkestraerp_orkestraerp_proto_goTypes = []interface{}{
	(*GetPageParams)(nil), // 0: orkestraerp.GetPageParams
	(*ResponsePage)(nil),  // 1: orkestraerp.ResponsePage
	(*Filters)(nil),       // 2: orkestraerp.Filters
	(*OrderBy)(nil),       // 3: orkestraerp.OrderBy
}
var file_orkestraerp_orkestraerp_proto_depIdxs = []int32{
	2, // 0: orkestraerp.GetPageParams.Filters:type_name -> orkestraerp.Filters
	3, // 1: orkestraerp.GetPageParams.OrderBy:type_name -> orkestraerp.OrderBy
	0, // 2: orkestraerp.OrkestraErp.GetPage:input_type -> orkestraerp.GetPageParams
	1, // 3: orkestraerp.OrkestraErp.GetPage:output_type -> orkestraerp.ResponsePage
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_orkestraerp_orkestraerp_proto_init() }
func file_orkestraerp_orkestraerp_proto_init() {
	if File_orkestraerp_orkestraerp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orkestraerp_orkestraerp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPageParams); i {
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
		file_orkestraerp_orkestraerp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponsePage); i {
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
		file_orkestraerp_orkestraerp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filters); i {
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
		file_orkestraerp_orkestraerp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderBy); i {
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
	file_orkestraerp_orkestraerp_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_orkestraerp_orkestraerp_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_orkestraerp_orkestraerp_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orkestraerp_orkestraerp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orkestraerp_orkestraerp_proto_goTypes,
		DependencyIndexes: file_orkestraerp_orkestraerp_proto_depIdxs,
		MessageInfos:      file_orkestraerp_orkestraerp_proto_msgTypes,
	}.Build()
	File_orkestraerp_orkestraerp_proto = out.File
	file_orkestraerp_orkestraerp_proto_rawDesc = nil
	file_orkestraerp_orkestraerp_proto_goTypes = nil
	file_orkestraerp_orkestraerp_proto_depIdxs = nil
}
