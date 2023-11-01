// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: rpc_create_coin.proto

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

type CreateCoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinId            string   `protobuf:"bytes,1,opt,name=coin_id,json=coinId,proto3" json:"coin_id,omitempty"`
	CoinUuid          string   `protobuf:"bytes,2,opt,name=coin_uuid,json=coinUuid,proto3" json:"coin_uuid,omitempty"`
	Name              string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Price             string   `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	MarketCap         string   `protobuf:"bytes,5,opt,name=market_cap,json=marketCap,proto3" json:"market_cap,omitempty"`
	NumberOfMarkets   int32    `protobuf:"varint,6,opt,name=number_of_markets,json=numberOfMarkets,proto3" json:"number_of_markets,omitempty"`
	NumberOfExchanges int32    `protobuf:"varint,7,opt,name=number_of_exchanges,json=numberOfExchanges,proto3" json:"number_of_exchanges,omitempty"`
	ApprovedSupply    bool     `protobuf:"varint,8,opt,name=approved_supply,json=approvedSupply,proto3" json:"approved_supply,omitempty"`
	CirculatingSupply string   `protobuf:"bytes,9,opt,name=circulating_supply,json=circulatingSupply,proto3" json:"circulating_supply,omitempty"`
	TotalSupply       string   `protobuf:"bytes,10,opt,name=total_supply,json=totalSupply,proto3" json:"total_supply,omitempty"`
	MaxSupply         string   `protobuf:"bytes,11,opt,name=max_supply,json=maxSupply,proto3" json:"max_supply,omitempty"`
	Rank              int32    `protobuf:"varint,12,opt,name=rank,proto3" json:"rank,omitempty"`
	Volume            string   `protobuf:"bytes,13,opt,name=volume,proto3" json:"volume,omitempty"`
	DailyChange       string   `protobuf:"bytes,14,opt,name=daily_change,json=dailyChange,proto3" json:"daily_change,omitempty"`
	ImageUrl          string   `protobuf:"bytes,15,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Description       string   `protobuf:"bytes,16,opt,name=description,proto3" json:"description,omitempty"`
	AllTimeHigh       string   `protobuf:"bytes,17,opt,name=all_time_high,json=allTimeHigh,proto3" json:"all_time_high,omitempty"`
	Tags              []string `protobuf:"bytes,18,rep,name=tags,proto3" json:"tags,omitempty"`
	Website           string   `protobuf:"bytes,19,opt,name=website,proto3" json:"website,omitempty"`
	SocialMediaLinks  []string `protobuf:"bytes,20,rep,name=social_media_links,json=socialMediaLinks,proto3" json:"social_media_links,omitempty"`
}

func (x *CreateCoinRequest) Reset() {
	*x = CreateCoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_coin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCoinRequest) ProtoMessage() {}

func (x *CreateCoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_coin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCoinRequest.ProtoReflect.Descriptor instead.
func (*CreateCoinRequest) Descriptor() ([]byte, []int) {
	return file_rpc_create_coin_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCoinRequest) GetCoinId() string {
	if x != nil {
		return x.CoinId
	}
	return ""
}

func (x *CreateCoinRequest) GetCoinUuid() string {
	if x != nil {
		return x.CoinUuid
	}
	return ""
}

func (x *CreateCoinRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCoinRequest) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *CreateCoinRequest) GetMarketCap() string {
	if x != nil {
		return x.MarketCap
	}
	return ""
}

func (x *CreateCoinRequest) GetNumberOfMarkets() int32 {
	if x != nil {
		return x.NumberOfMarkets
	}
	return 0
}

func (x *CreateCoinRequest) GetNumberOfExchanges() int32 {
	if x != nil {
		return x.NumberOfExchanges
	}
	return 0
}

func (x *CreateCoinRequest) GetApprovedSupply() bool {
	if x != nil {
		return x.ApprovedSupply
	}
	return false
}

func (x *CreateCoinRequest) GetCirculatingSupply() string {
	if x != nil {
		return x.CirculatingSupply
	}
	return ""
}

func (x *CreateCoinRequest) GetTotalSupply() string {
	if x != nil {
		return x.TotalSupply
	}
	return ""
}

func (x *CreateCoinRequest) GetMaxSupply() string {
	if x != nil {
		return x.MaxSupply
	}
	return ""
}

func (x *CreateCoinRequest) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *CreateCoinRequest) GetVolume() string {
	if x != nil {
		return x.Volume
	}
	return ""
}

func (x *CreateCoinRequest) GetDailyChange() string {
	if x != nil {
		return x.DailyChange
	}
	return ""
}

func (x *CreateCoinRequest) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *CreateCoinRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateCoinRequest) GetAllTimeHigh() string {
	if x != nil {
		return x.AllTimeHigh
	}
	return ""
}

func (x *CreateCoinRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CreateCoinRequest) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *CreateCoinRequest) GetSocialMediaLinks() []string {
	if x != nil {
		return x.SocialMediaLinks
	}
	return nil
}

type CreateCoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coin *Coin `protobuf:"bytes,1,opt,name=coin,proto3" json:"coin,omitempty"`
}

func (x *CreateCoinResponse) Reset() {
	*x = CreateCoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_coin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCoinResponse) ProtoMessage() {}

func (x *CreateCoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_coin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCoinResponse.ProtoReflect.Descriptor instead.
func (*CreateCoinResponse) Descriptor() ([]byte, []int) {
	return file_rpc_create_coin_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCoinResponse) GetCoin() *Coin {
	if x != nil {
		return x.Coin
	}
	return nil
}

var File_rpc_create_coin_proto protoreflect.FileDescriptor

var file_rpc_create_coin_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0a, 0x63, 0x6f, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x05, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x6f, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x55,
	0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x61, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x70, 0x12, 0x2a, 0x0a, 0x11,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f,
	0x66, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x64, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x79, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x69, 0x72, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63,
	0x69, 0x72, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79,
	0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x22, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x68, 0x69, 0x67,
	0x68, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65,
	0x48, 0x69, 0x67, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x12, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73,
	0x69, 0x74, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4c, 0x69, 0x6e, 0x6b, 0x73,
	0x22, 0x32, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x63, 0x6f, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x04,
	0x63, 0x6f, 0x69, 0x6e, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6a, 0x61, 0x6b, 0x75, 0x62, 0x2f, 0x61, 0x69, 0x6f, 0x70, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_create_coin_proto_rawDescOnce sync.Once
	file_rpc_create_coin_proto_rawDescData = file_rpc_create_coin_proto_rawDesc
)

func file_rpc_create_coin_proto_rawDescGZIP() []byte {
	file_rpc_create_coin_proto_rawDescOnce.Do(func() {
		file_rpc_create_coin_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_create_coin_proto_rawDescData)
	})
	return file_rpc_create_coin_proto_rawDescData
}

var file_rpc_create_coin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_create_coin_proto_goTypes = []interface{}{
	(*CreateCoinRequest)(nil),  // 0: pb.CreateCoinRequest
	(*CreateCoinResponse)(nil), // 1: pb.CreateCoinResponse
	(*Coin)(nil),               // 2: pb.Coin
}
var file_rpc_create_coin_proto_depIdxs = []int32{
	2, // 0: pb.CreateCoinResponse.coin:type_name -> pb.Coin
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_create_coin_proto_init() }
func file_rpc_create_coin_proto_init() {
	if File_rpc_create_coin_proto != nil {
		return
	}
	file_coin_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_create_coin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCoinRequest); i {
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
		file_rpc_create_coin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCoinResponse); i {
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
			RawDescriptor: file_rpc_create_coin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_create_coin_proto_goTypes,
		DependencyIndexes: file_rpc_create_coin_proto_depIdxs,
		MessageInfos:      file_rpc_create_coin_proto_msgTypes,
	}.Build()
	File_rpc_create_coin_proto = out.File
	file_rpc_create_coin_proto_rawDesc = nil
	file_rpc_create_coin_proto_goTypes = nil
	file_rpc_create_coin_proto_depIdxs = nil
}