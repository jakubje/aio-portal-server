// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: roll_up.proto

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

type Roll_Up struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol       string  `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Type         int32   `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	TotalCost    float64 `protobuf:"fixed64,3,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	TotalCoins   float64 `protobuf:"fixed64,4,opt,name=total_coins,json=totalCoins,proto3" json:"total_coins,omitempty"`
	PricePerCoin float64 `protobuf:"fixed64,5,opt,name=price_per_coin,json=pricePerCoin,proto3" json:"price_per_coin,omitempty"`
}

func (x *Roll_Up) Reset() {
	*x = Roll_Up{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roll_up_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Roll_Up) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Roll_Up) ProtoMessage() {}

func (x *Roll_Up) ProtoReflect() protoreflect.Message {
	mi := &file_roll_up_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Roll_Up.ProtoReflect.Descriptor instead.
func (*Roll_Up) Descriptor() ([]byte, []int) {
	return file_roll_up_proto_rawDescGZIP(), []int{0}
}

func (x *Roll_Up) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Roll_Up) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Roll_Up) GetTotalCost() float64 {
	if x != nil {
		return x.TotalCost
	}
	return 0
}

func (x *Roll_Up) GetTotalCoins() float64 {
	if x != nil {
		return x.TotalCoins
	}
	return 0
}

func (x *Roll_Up) GetPricePerCoin() float64 {
	if x != nil {
		return x.PricePerCoin
	}
	return 0
}

var File_roll_up_proto protoreflect.FileDescriptor

var file_roll_up_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x6f, 0x6c, 0x6c, 0x5f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x9b, 0x01, 0x0a, 0x07, 0x52, 0x6f, 0x6c, 0x6c, 0x5f, 0x55, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x43, 0x6f, 0x69,
	0x6e, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6a, 0x61, 0x6b, 0x75, 0x62, 0x2f, 0x61, 0x69, 0x6f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_roll_up_proto_rawDescOnce sync.Once
	file_roll_up_proto_rawDescData = file_roll_up_proto_rawDesc
)

func file_roll_up_proto_rawDescGZIP() []byte {
	file_roll_up_proto_rawDescOnce.Do(func() {
		file_roll_up_proto_rawDescData = protoimpl.X.CompressGZIP(file_roll_up_proto_rawDescData)
	})
	return file_roll_up_proto_rawDescData
}

var file_roll_up_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_roll_up_proto_goTypes = []interface{}{
	(*Roll_Up)(nil), // 0: pb.Roll_Up
}
var file_roll_up_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_roll_up_proto_init() }
func file_roll_up_proto_init() {
	if File_roll_up_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_roll_up_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Roll_Up); i {
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
			RawDescriptor: file_roll_up_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_roll_up_proto_goTypes,
		DependencyIndexes: file_roll_up_proto_depIdxs,
		MessageInfos:      file_roll_up_proto_msgTypes,
	}.Build()
	File_roll_up_proto = out.File
	file_roll_up_proto_rawDesc = nil
	file_roll_up_proto_goTypes = nil
	file_roll_up_proto_depIdxs = nil
}
