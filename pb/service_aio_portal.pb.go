// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: service_aio_portal.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_aio_portal_proto protoreflect.FileDescriptor

var file_service_aio_portal_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x69, 0x6f, 0x5f, 0x70, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x70, 0x63, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x72,
	0x70, 0x63, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1a, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6f,
	0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x72,
	0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f,
	0x5f, 0x72, 0x6f, 0x6c, 0x6c, 0x5f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a,
	0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x66,
	0x6f, 0x6c, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x72, 0x70, 0x63, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65,
	0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1a, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x72,
	0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1a, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72,
	0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x72, 0x70, 0x63,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x63, 0x6f, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x72, 0x70, 0x63,
	0x5f, 0x61, 0x64, 0x64, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x63,
	0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xfe, 0x1d, 0x0a, 0x09, 0x41, 0x69,
	0x6f, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x12, 0x8e, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x92, 0x41, 0x34, 0x12, 0x0f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x21, 0x55, 0x73, 0x65,
	0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x75, 0x73, 0x65, 0x72, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x12, 0xa6, 0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x6c, 0x92, 0x41, 0x50, 0x12, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x20,
	0x75, 0x73, 0x65, 0x72, 0x1a, 0x42, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41,
	0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x20, 0x61, 0x6e, 0x64, 0x20, 0x67, 0x65, 0x74, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x26, 0x20, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01,
	0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x84, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x47, 0x92, 0x41, 0x2a, 0x12, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x75, 0x73, 0x65,
	0x72, 0x1a, 0x1b, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20,
	0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x32, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x8c, 0x01, 0x0a, 0x0b, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c, 0x92, 0x41, 0x31, 0x12, 0x0c,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x21, 0x55, 0x73,
	0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0xac, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x12, 0x1a, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60, 0x92, 0x41, 0x3e, 0x12, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x1a,
	0x26, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f,
	0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x70, 0x6f,
	0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a,
	0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x72,
	0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x12, 0xa6, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x5a, 0x92, 0x41, 0x38, 0x12, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x20, 0x61, 0x20, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x1a, 0x22, 0x55, 0x73,
	0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x12,
	0x91, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f,
	0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c,
	0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x4e, 0x92, 0x41, 0x32, 0x12, 0x0f, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20,
	0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x1a, 0x1f, 0x55, 0x73, 0x65, 0x20, 0x74,
	0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20, 0x61,
	0x20, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f,
	0x6c, 0x69, 0x6f, 0x12, 0xb1, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74,
	0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a,
	0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69,
	0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6b, 0x92, 0x41, 0x4e, 0x12,
	0x1d, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x75,
	0x73, 0x65, 0x72, 0x20, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x73, 0x1a, 0x2d,
	0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20,
	0x67, 0x65, 0x74, 0x20, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x75, 0x73,
	0x65, 0x72, 0x20, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x73, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x72,
	0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x73, 0x12, 0xa6, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x6c, 0x6c, 0x55, 0x70, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x55,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f,
	0x6c, 0x6c, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x72, 0x92, 0x41,
	0x59, 0x12, 0x24, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66,
	0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x20, 0x68,
	0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x31, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69,
	0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20, 0x61, 0x20, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x20, 0x68, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10,
	0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x6f, 0x6c, 0x6c, 0x75, 0x70,
	0x12, 0xa3, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x66,
	0x6f, 0x6c, 0x69, 0x6f, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74,
	0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x57, 0x92,
	0x41, 0x38, 0x12, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x20, 0x61, 0x20, 0x70, 0x6f, 0x72,
	0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x1a, 0x22, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x20, 0x61,
	0x20, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16,
	0x2a, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x72,
	0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x12, 0xb2, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60, 0x92, 0x41, 0x3c, 0x12, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x24, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41,
	0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b,
	0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x9d, 0x01, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x54, 0x92, 0x41, 0x36, 0x12, 0x11, 0x47, 0x65, 0x74, 0x20,
	0x61, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x21, 0x55,
	0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x67,
	0x65, 0x74, 0x20, 0x61, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xba, 0x01, 0x0a, 0x10,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6b, 0x92, 0x41, 0x4b,
	0x12, 0x19, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x2e, 0x55, 0x73, 0x65,
	0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74,
	0x20, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0xd0, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x43,
	0x6f, 0x69, 0x6e, 0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x75, 0x92, 0x41, 0x50, 0x12, 0x1e, 0x47, 0x65, 0x74, 0x20, 0x61,
	0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x63, 0x6f, 0x69, 0x6e, 0x20, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x2e, 0x55, 0x73, 0x65, 0x20, 0x74,
	0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20, 0x61,
	0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x63, 0x6f, 0x69, 0x6e, 0x20, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12,
	0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x12, 0xac, 0x01, 0x0a, 0x0f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60, 0x92, 0x41, 0x3e, 0x12, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73,
	0x74, 0x1a, 0x28, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20,
	0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20, 0x61, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20,
	0x61, 0x20, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xa6, 0x01, 0x0a, 0x0f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x1a,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5a, 0x92, 0x41, 0x38, 0x12, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74,
	0x1a, 0x22, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74,
	0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x6c, 0x69, 0x73, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x91, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x6c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4e, 0x92, 0x41, 0x32, 0x12, 0x0f, 0x47, 0x65,
	0x74, 0x20, 0x61, 0x20, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x1f, 0x55,
	0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x67,
	0x65, 0x74, 0x20, 0x61, 0x20, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xa6, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x6c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x61, 0x92,
	0x41, 0x44, 0x12, 0x18, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f,
	0x66, 0x20, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x1a, 0x28, 0x55, 0x73,
	0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65,
	0x74, 0x20, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x73,
	0x12, 0xa3, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x6c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x57, 0x92,
	0x41, 0x38, 0x12, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x20, 0x61, 0x20, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x22, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x20, 0x61,
	0x20, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16,
	0x2a, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x8e, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x92, 0x41, 0x34, 0x12, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x63, 0x6f, 0x69, 0x6e, 0x1a, 0x21, 0x55, 0x73, 0x65, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x63, 0x6f, 0x69, 0x6e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x12, 0xbe, 0x01, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x1b, 0x2e, 0x70,
	0x62, 0x2e, 0x41, 0x64, 0x64, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x41,
	0x64, 0x64, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6f, 0x92, 0x41, 0x4b, 0x12, 0x19, 0x41, 0x64,
	0x64, 0x20, 0x61, 0x20, 0x63, 0x6f, 0x69, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x20, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x2e, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69,
	0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x61, 0x64,
	0x64, 0x20, 0x61, 0x20, 0x63, 0x6f, 0x69, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x20, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a,
	0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2f,
	0x61, 0x64, 0x64, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x12, 0x88, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4b, 0x92, 0x41, 0x2e, 0x12, 0x0d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x63, 0x6f, 0x69, 0x6e, 0x1a, 0x1d, 0x55, 0x73, 0x65, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x20, 0x61, 0x20, 0x63, 0x6f, 0x69, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a,
	0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63,
	0x6f, 0x69, 0x6e, 0x12, 0xb4, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5f, 0x92, 0x41, 0x3c, 0x12, 0x14,
	0x4c, 0x69, 0x73, 0x74, 0x20, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x63,
	0x6f, 0x69, 0x6e, 0x73, 0x1a, 0x24, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41,
	0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x6c, 0x69, 0x73, 0x74, 0x20, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a,
	0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x42, 0x7a, 0x92, 0x41, 0x51, 0x12,
	0x4f, 0x0a, 0x0e, 0x41, 0x49, 0x4f, 0x20, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x20, 0x41, 0x50,
	0x49, 0x22, 0x38, 0x0a, 0x04, 0x4a, 0x61, 0x6b, 0x65, 0x12, 0x1a, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x61,
	0x6b, 0x75, 0x62, 0x6a, 0x65, 0x1a, 0x14, 0x6a, 0x61, 0x6b, 0x65, 0x74, 0x68, 0x65, 0x64, 0x65,
	0x76, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31, 0x2e, 0x31,
	0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x61, 0x6b,
	0x75, 0x62, 0x2f, 0x61, 0x69, 0x6f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_aio_portal_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),             // 0: pb.CreateUserRequest
	(*LoginUserRequest)(nil),              // 1: pb.LoginUserRequest
	(*UpdateUserRequest)(nil),             // 2: pb.UpdateUserRequest
	(*VerifyEmailRequest)(nil),            // 3: pb.VerifyEmailRequest
	(*CreatePortfolioRequest)(nil),        // 4: pb.CreatePortfolioRequest
	(*UpdatePortfolioRequest)(nil),        // 5: pb.UpdatePortfolioRequest
	(*GetPortfolioRequest)(nil),           // 6: pb.GetPortfolioRequest
	(*emptypb.Empty)(nil),                 // 7: google.protobuf.Empty
	(*RollUpRequest)(nil),                 // 8: pb.RollUpRequest
	(*DeletePortfolioRequest)(nil),        // 9: pb.DeletePortfolioRequest
	(*CreateTransactionRequest)(nil),      // 10: pb.CreateTransactionRequest
	(*GetTransactionRequest)(nil),         // 11: pb.GetTransactionRequest
	(*ListTransactionsRequest)(nil),       // 12: pb.ListTransactionsRequest
	(*ListTransactionsByCoinRequest)(nil), // 13: pb.ListTransactionsByCoinRequest
	(*CreateWatchlistRequest)(nil),        // 14: pb.CreateWatchlistRequest
	(*UpdateWatchlistRequest)(nil),        // 15: pb.UpdateWatchlistRequest
	(*GetWatchlistRequest)(nil),           // 16: pb.GetWatchlistRequest
	(*DeleteWatchlistRequest)(nil),        // 17: pb.DeleteWatchlistRequest
	(*CreateCoinRequest)(nil),             // 18: pb.CreateCoinRequest
	(*AddWatchlistCoinRequest)(nil),       // 19: pb.AddWatchlistCoinRequest
	(*UpdateCoinRequest)(nil),             // 20: pb.UpdateCoinRequest
	(*ListWatchlistCoinsRequest)(nil),     // 21: pb.ListWatchlistCoinsRequest
	(*CreateUserResponse)(nil),            // 22: pb.CreateUserResponse
	(*LoginUserResponse)(nil),             // 23: pb.LoginUserResponse
	(*UpdateUserResponse)(nil),            // 24: pb.UpdateUserResponse
	(*VerifyEmailResponse)(nil),           // 25: pb.VerifyEmailResponse
	(*CreatePortfolioResponse)(nil),       // 26: pb.CreatePortfolioResponse
	(*UpdatePortfolioResponse)(nil),       // 27: pb.UpdatePortfolioResponse
	(*GetPortfolioResponse)(nil),          // 28: pb.GetPortfolioResponse
	(*ListPortfoliosResponse)(nil),        // 29: pb.ListPortfoliosResponse
	(*RollUpResponse)(nil),                // 30: pb.RollUpResponse
	(*DeletePortfolioResponse)(nil),       // 31: pb.DeletePortfolioResponse
	(*CreateTransactionResponse)(nil),     // 32: pb.CreateTransactionResponse
	(*GetTransactionResponse)(nil),        // 33: pb.GetTransactionResponse
	(*ListTransactionsResponse)(nil),      // 34: pb.ListTransactionsResponse
	(*CreateWatchlistResponse)(nil),       // 35: pb.CreateWatchlistResponse
	(*UpdateWatchlistResponse)(nil),       // 36: pb.UpdateWatchlistResponse
	(*GetWatchlistResponse)(nil),          // 37: pb.GetWatchlistResponse
	(*GetWatchlistsResponse)(nil),         // 38: pb.GetWatchlistsResponse
	(*DeleteWatchlistResponse)(nil),       // 39: pb.DeleteWatchlistResponse
	(*CreateCoinResponse)(nil),            // 40: pb.CreateCoinResponse
	(*AddWatchlistCoinResponse)(nil),      // 41: pb.AddWatchlistCoinResponse
	(*UpdateCoinResponse)(nil),            // 42: pb.UpdateCoinResponse
	(*ListWatchlistCoinsResponse)(nil),    // 43: pb.ListWatchlistCoinsResponse
}
var file_service_aio_portal_proto_depIdxs = []int32{
	0,  // 0: pb.AioPortal.CreateUser:input_type -> pb.CreateUserRequest
	1,  // 1: pb.AioPortal.LoginUser:input_type -> pb.LoginUserRequest
	2,  // 2: pb.AioPortal.UpdateUser:input_type -> pb.UpdateUserRequest
	3,  // 3: pb.AioPortal.VerifyEmail:input_type -> pb.VerifyEmailRequest
	4,  // 4: pb.AioPortal.CreatePortfolio:input_type -> pb.CreatePortfolioRequest
	5,  // 5: pb.AioPortal.UpdatePortfolio:input_type -> pb.UpdatePortfolioRequest
	6,  // 6: pb.AioPortal.GetPortfolio:input_type -> pb.GetPortfolioRequest
	7,  // 7: pb.AioPortal.ListPortfolios:input_type -> google.protobuf.Empty
	8,  // 8: pb.AioPortal.GetRollUp:input_type -> pb.RollUpRequest
	9,  // 9: pb.AioPortal.DeletePortfolio:input_type -> pb.DeletePortfolioRequest
	10, // 10: pb.AioPortal.CreateTransaction:input_type -> pb.CreateTransactionRequest
	11, // 11: pb.AioPortal.GetTransaction:input_type -> pb.GetTransactionRequest
	12, // 12: pb.AioPortal.ListTransactions:input_type -> pb.ListTransactionsRequest
	13, // 13: pb.AioPortal.ListTransactionsByCoin:input_type -> pb.ListTransactionsByCoinRequest
	14, // 14: pb.AioPortal.CreateWatchlist:input_type -> pb.CreateWatchlistRequest
	15, // 15: pb.AioPortal.UpdateWatchlist:input_type -> pb.UpdateWatchlistRequest
	16, // 16: pb.AioPortal.GetWatchlist:input_type -> pb.GetWatchlistRequest
	7,  // 17: pb.AioPortal.ListWatchlists:input_type -> google.protobuf.Empty
	17, // 18: pb.AioPortal.DeleteWatchlist:input_type -> pb.DeleteWatchlistRequest
	18, // 19: pb.AioPortal.CreateCoin:input_type -> pb.CreateCoinRequest
	19, // 20: pb.AioPortal.AddWatchlistCoin:input_type -> pb.AddWatchlistCoinRequest
	20, // 21: pb.AioPortal.UpdateCoin:input_type -> pb.UpdateCoinRequest
	21, // 22: pb.AioPortal.ListWatchlistCoins:input_type -> pb.ListWatchlistCoinsRequest
	22, // 23: pb.AioPortal.CreateUser:output_type -> pb.CreateUserResponse
	23, // 24: pb.AioPortal.LoginUser:output_type -> pb.LoginUserResponse
	24, // 25: pb.AioPortal.UpdateUser:output_type -> pb.UpdateUserResponse
	25, // 26: pb.AioPortal.VerifyEmail:output_type -> pb.VerifyEmailResponse
	26, // 27: pb.AioPortal.CreatePortfolio:output_type -> pb.CreatePortfolioResponse
	27, // 28: pb.AioPortal.UpdatePortfolio:output_type -> pb.UpdatePortfolioResponse
	28, // 29: pb.AioPortal.GetPortfolio:output_type -> pb.GetPortfolioResponse
	29, // 30: pb.AioPortal.ListPortfolios:output_type -> pb.ListPortfoliosResponse
	30, // 31: pb.AioPortal.GetRollUp:output_type -> pb.RollUpResponse
	31, // 32: pb.AioPortal.DeletePortfolio:output_type -> pb.DeletePortfolioResponse
	32, // 33: pb.AioPortal.CreateTransaction:output_type -> pb.CreateTransactionResponse
	33, // 34: pb.AioPortal.GetTransaction:output_type -> pb.GetTransactionResponse
	34, // 35: pb.AioPortal.ListTransactions:output_type -> pb.ListTransactionsResponse
	34, // 36: pb.AioPortal.ListTransactionsByCoin:output_type -> pb.ListTransactionsResponse
	35, // 37: pb.AioPortal.CreateWatchlist:output_type -> pb.CreateWatchlistResponse
	36, // 38: pb.AioPortal.UpdateWatchlist:output_type -> pb.UpdateWatchlistResponse
	37, // 39: pb.AioPortal.GetWatchlist:output_type -> pb.GetWatchlistResponse
	38, // 40: pb.AioPortal.ListWatchlists:output_type -> pb.GetWatchlistsResponse
	39, // 41: pb.AioPortal.DeleteWatchlist:output_type -> pb.DeleteWatchlistResponse
	40, // 42: pb.AioPortal.CreateCoin:output_type -> pb.CreateCoinResponse
	41, // 43: pb.AioPortal.AddWatchlistCoin:output_type -> pb.AddWatchlistCoinResponse
	42, // 44: pb.AioPortal.UpdateCoin:output_type -> pb.UpdateCoinResponse
	43, // 45: pb.AioPortal.ListWatchlistCoins:output_type -> pb.ListWatchlistCoinsResponse
	23, // [23:46] is the sub-list for method output_type
	0,  // [0:23] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_aio_portal_proto_init() }
func file_service_aio_portal_proto_init() {
	if File_service_aio_portal_proto != nil {
		return
	}
	file_rpc_create_user_proto_init()
	file_rpc_login_user_proto_init()
	file_rpc_update_user_proto_init()
	file_rpc_verify_email_proto_init()
	file_rpc_create_portfolio_proto_init()
	file_rpc_update_portfolio_proto_init()
	file_rpc_get_portfolio_proto_init()
	file_rpc_list_portfolios_proto_init()
	file_rpc_portfolio_roll_up_proto_init()
	file_rpc_delete_portfolio_proto_init()
	file_rpc_create_transaction_proto_init()
	file_rpc_get_transaction_proto_init()
	file_rpc_list_transactions_proto_init()
	file_rpc_list_transactions_by_coin_proto_init()
	file_rpc_create_watchlist_proto_init()
	file_rpc_update_watchlist_proto_init()
	file_rpc_get_watchlist_proto_init()
	file_rpc_list_watchlists_proto_init()
	file_rpc_delete_watchlist_proto_init()
	file_rpc_create_coin_proto_init()
	file_rpc_update_coin_proto_init()
	file_rpc_list_watchlist_coins_proto_init()
	file_rpc_add_watchlist_coin_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_aio_portal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_aio_portal_proto_goTypes,
		DependencyIndexes: file_service_aio_portal_proto_depIdxs,
	}.Build()
	File_service_aio_portal_proto = out.File
	file_service_aio_portal_proto_rawDesc = nil
	file_service_aio_portal_proto_goTypes = nil
	file_service_aio_portal_proto_depIdxs = nil
}
