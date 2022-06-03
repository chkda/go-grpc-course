// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: orderMgmnt/orderMgmnt.proto

package orderMgmnt

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

type OrderID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *OrderID) Reset() {
	*x = OrderID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderMgmnt_orderMgmnt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderID) ProtoMessage() {}

func (x *OrderID) ProtoReflect() protoreflect.Message {
	mi := &file_orderMgmnt_orderMgmnt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderID.ProtoReflect.Descriptor instead.
func (*OrderID) Descriptor() ([]byte, []int) {
	return file_orderMgmnt_orderMgmnt_proto_rawDescGZIP(), []int{0}
}

func (x *OrderID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Items       []string `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       float32  `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Destination string   `protobuf:"bytes,5,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderMgmnt_orderMgmnt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_orderMgmnt_orderMgmnt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_orderMgmnt_orderMgmnt_proto_rawDescGZIP(), []int{1}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetItems() []string {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Order) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Order) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

var File_orderMgmnt_orderMgmnt_proto protoreflect.FileDescriptor

var file_orderMgmnt_orderMgmnt_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x67, 0x6d, 0x6e, 0x74, 0x2f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x4d, 0x67, 0x6d, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x4d, 0x67, 0x6d, 0x6e, 0x74, 0x22, 0x1f, 0x0a, 0x07, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x05, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x32, 0x79, 0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x67, 0x6d, 0x6e, 0x74,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x67,
	0x6d, 0x6e, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x32, 0x0a, 0x08, 0x67,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d,
	0x67, 0x6d, 0x6e, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x4d, 0x67, 0x6d, 0x6e, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68,
	0x6b, 0x64, 0x61, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x67, 0x6d, 0x6e, 0x74, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x4d, 0x67, 0x6d, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orderMgmnt_orderMgmnt_proto_rawDescOnce sync.Once
	file_orderMgmnt_orderMgmnt_proto_rawDescData = file_orderMgmnt_orderMgmnt_proto_rawDesc
)

func file_orderMgmnt_orderMgmnt_proto_rawDescGZIP() []byte {
	file_orderMgmnt_orderMgmnt_proto_rawDescOnce.Do(func() {
		file_orderMgmnt_orderMgmnt_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderMgmnt_orderMgmnt_proto_rawDescData)
	})
	return file_orderMgmnt_orderMgmnt_proto_rawDescData
}

var file_orderMgmnt_orderMgmnt_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_orderMgmnt_orderMgmnt_proto_goTypes = []interface{}{
	(*OrderID)(nil), // 0: orderMgmnt.OrderID
	(*Order)(nil),   // 1: orderMgmnt.Order
}
var file_orderMgmnt_orderMgmnt_proto_depIdxs = []int32{
	1, // 0: orderMgmnt.OrderManagement.addOrder:input_type -> orderMgmnt.Order
	0, // 1: orderMgmnt.OrderManagement.getOrder:input_type -> orderMgmnt.OrderID
	0, // 2: orderMgmnt.OrderManagement.addOrder:output_type -> orderMgmnt.OrderID
	1, // 3: orderMgmnt.OrderManagement.getOrder:output_type -> orderMgmnt.Order
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_orderMgmnt_orderMgmnt_proto_init() }
func file_orderMgmnt_orderMgmnt_proto_init() {
	if File_orderMgmnt_orderMgmnt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orderMgmnt_orderMgmnt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderID); i {
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
		file_orderMgmnt_orderMgmnt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
			RawDescriptor: file_orderMgmnt_orderMgmnt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orderMgmnt_orderMgmnt_proto_goTypes,
		DependencyIndexes: file_orderMgmnt_orderMgmnt_proto_depIdxs,
		MessageInfos:      file_orderMgmnt_orderMgmnt_proto_msgTypes,
	}.Build()
	File_orderMgmnt_orderMgmnt_proto = out.File
	file_orderMgmnt_orderMgmnt_proto_rawDesc = nil
	file_orderMgmnt_orderMgmnt_proto_goTypes = nil
	file_orderMgmnt_orderMgmnt_proto_depIdxs = nil
}
