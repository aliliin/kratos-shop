// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: order/v1/order.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartId int64 `protobuf:"varint,1,opt,name=cartId,proto3" json:"cartId,omitempty"`
	SkuId  int64 `protobuf:"varint,2,opt,name=skuId,proto3" json:"skuId,omitempty"`
	SkuNum int32 `protobuf:"varint,3,opt,name=skuNum,proto3" json:"skuNum,omitempty"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_v1_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_order_v1_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_order_v1_order_proto_rawDescGZIP(), []int{0}
}

func (x *CartItem) GetCartId() int64 {
	if x != nil {
		return x.CartId
	}
	return 0
}

func (x *CartItem) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *CartItem) GetSkuNum() int32 {
	if x != nil {
		return x.SkuNum
	}
	return 0
}

type OrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId   int64       `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Address  int64       `protobuf:"varint,3,opt,name=address,proto3" json:"address,omitempty"`
	CartItem []*CartItem `protobuf:"bytes,4,rep,name=cartItem,proto3" json:"cartItem,omitempty"` //  购物车购买和直接购买
}

func (x *OrderRequest) Reset() {
	*x = OrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_v1_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRequest) ProtoMessage() {}

func (x *OrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_v1_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRequest.ProtoReflect.Descriptor instead.
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return file_order_v1_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OrderRequest) GetAddress() int64 {
	if x != nil {
		return x.Address
	}
	return 0
}

func (x *OrderRequest) GetCartItem() []*CartItem {
	if x != nil {
		return x.CartItem
	}
	return nil
}

type OrderInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId  int32   `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	OrderSn string  `protobuf:"bytes,3,opt,name=orderSn,proto3" json:"orderSn,omitempty"`
	PayType string  `protobuf:"bytes,4,opt,name=payType,proto3" json:"payType,omitempty"`
	Status  string  `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Post    string  `protobuf:"bytes,6,opt,name=post,proto3" json:"post,omitempty"`
	Total   float32 `protobuf:"fixed32,7,opt,name=total,proto3" json:"total,omitempty"`
	Address string  `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	Name    string  `protobuf:"bytes,9,opt,name=name,proto3" json:"name,omitempty"`
	Mobile  string  `protobuf:"bytes,10,opt,name=mobile,proto3" json:"mobile,omitempty"`
	AddTime string  `protobuf:"bytes,11,opt,name=addTime,proto3" json:"addTime,omitempty"`
}

func (x *OrderInfoResponse) Reset() {
	*x = OrderInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_v1_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderInfoResponse) ProtoMessage() {}

func (x *OrderInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_v1_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderInfoResponse.ProtoReflect.Descriptor instead.
func (*OrderInfoResponse) Descriptor() ([]byte, []int) {
	return file_order_v1_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderInfoResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderInfoResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OrderInfoResponse) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

func (x *OrderInfoResponse) GetPayType() string {
	if x != nil {
		return x.PayType
	}
	return ""
}

func (x *OrderInfoResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderInfoResponse) GetPost() string {
	if x != nil {
		return x.Post
	}
	return ""
}

func (x *OrderInfoResponse) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *OrderInfoResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *OrderInfoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrderInfoResponse) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *OrderInfoResponse) GetAddTime() string {
	if x != nil {
		return x.AddTime
	}
	return ""
}

var File_order_v1_order_proto protoreflect.FileDescriptor

var file_order_v1_order_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x08, 0x43, 0x61, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x6b,
	0x75, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x6b, 0x75, 0x4e, 0x75, 0x6d, 0x22, 0x92, 0x01, 0x0a, 0x0c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x2e, 0x0a, 0x08, 0x63, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x63, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x22, 0x91, 0x02, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x32, 0x4b, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x42, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x17, 0x5a, 0x15, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_order_v1_order_proto_rawDescOnce sync.Once
	file_order_v1_order_proto_rawDescData = file_order_v1_order_proto_rawDesc
)

func file_order_v1_order_proto_rawDescGZIP() []byte {
	file_order_v1_order_proto_rawDescOnce.Do(func() {
		file_order_v1_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_v1_order_proto_rawDescData)
	})
	return file_order_v1_order_proto_rawDescData
}

var file_order_v1_order_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_order_v1_order_proto_goTypes = []interface{}{
	(*CartItem)(nil),          // 0: order.v1.CartItem
	(*OrderRequest)(nil),      // 1: order.v1.OrderRequest
	(*OrderInfoResponse)(nil), // 2: order.v1.OrderInfoResponse
}
var file_order_v1_order_proto_depIdxs = []int32{
	0, // 0: order.v1.OrderRequest.cartItem:type_name -> order.v1.CartItem
	1, // 1: order.v1.Order.CreateOrder:input_type -> order.v1.OrderRequest
	2, // 2: order.v1.Order.CreateOrder:output_type -> order.v1.OrderInfoResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_order_v1_order_proto_init() }
func file_order_v1_order_proto_init() {
	if File_order_v1_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_v1_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItem); i {
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
		file_order_v1_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderRequest); i {
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
		file_order_v1_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderInfoResponse); i {
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
			RawDescriptor: file_order_v1_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_v1_order_proto_goTypes,
		DependencyIndexes: file_order_v1_order_proto_depIdxs,
		MessageInfos:      file_order_v1_order_proto_msgTypes,
	}.Build()
	File_order_v1_order_proto = out.File
	file_order_v1_order_proto_rawDesc = nil
	file_order_v1_order_proto_goTypes = nil
	file_order_v1_order_proto_depIdxs = nil
}
