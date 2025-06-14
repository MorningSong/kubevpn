// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.29.3
// source: dhcpserver.proto

package rpc

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

type RentIPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodName      string `protobuf:"bytes,1,opt,name=PodName,proto3" json:"PodName,omitempty"`
	PodNamespace string `protobuf:"bytes,2,opt,name=PodNamespace,proto3" json:"PodNamespace,omitempty"`
}

func (x *RentIPRequest) Reset() {
	*x = RentIPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dhcpserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RentIPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RentIPRequest) ProtoMessage() {}

func (x *RentIPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dhcpserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RentIPRequest.ProtoReflect.Descriptor instead.
func (*RentIPRequest) Descriptor() ([]byte, []int) {
	return file_dhcpserver_proto_rawDescGZIP(), []int{0}
}

func (x *RentIPRequest) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *RentIPRequest) GetPodNamespace() string {
	if x != nil {
		return x.PodNamespace
	}
	return ""
}

type RentIPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IPv4CIDR string `protobuf:"bytes,1,opt,name=IPv4CIDR,proto3" json:"IPv4CIDR,omitempty"`
	IPv6CIDR string `protobuf:"bytes,2,opt,name=IPv6CIDR,proto3" json:"IPv6CIDR,omitempty"`
}

func (x *RentIPResponse) Reset() {
	*x = RentIPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dhcpserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RentIPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RentIPResponse) ProtoMessage() {}

func (x *RentIPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dhcpserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RentIPResponse.ProtoReflect.Descriptor instead.
func (*RentIPResponse) Descriptor() ([]byte, []int) {
	return file_dhcpserver_proto_rawDescGZIP(), []int{1}
}

func (x *RentIPResponse) GetIPv4CIDR() string {
	if x != nil {
		return x.IPv4CIDR
	}
	return ""
}

func (x *RentIPResponse) GetIPv6CIDR() string {
	if x != nil {
		return x.IPv6CIDR
	}
	return ""
}

type ReleaseIPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodName      string `protobuf:"bytes,1,opt,name=PodName,proto3" json:"PodName,omitempty"`
	PodNamespace string `protobuf:"bytes,2,opt,name=PodNamespace,proto3" json:"PodNamespace,omitempty"`
	IPv4CIDR     string `protobuf:"bytes,3,opt,name=IPv4CIDR,proto3" json:"IPv4CIDR,omitempty"`
	IPv6CIDR     string `protobuf:"bytes,4,opt,name=IPv6CIDR,proto3" json:"IPv6CIDR,omitempty"`
}

func (x *ReleaseIPRequest) Reset() {
	*x = ReleaseIPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dhcpserver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseIPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseIPRequest) ProtoMessage() {}

func (x *ReleaseIPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dhcpserver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseIPRequest.ProtoReflect.Descriptor instead.
func (*ReleaseIPRequest) Descriptor() ([]byte, []int) {
	return file_dhcpserver_proto_rawDescGZIP(), []int{2}
}

func (x *ReleaseIPRequest) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *ReleaseIPRequest) GetPodNamespace() string {
	if x != nil {
		return x.PodNamespace
	}
	return ""
}

func (x *ReleaseIPRequest) GetIPv4CIDR() string {
	if x != nil {
		return x.IPv4CIDR
	}
	return ""
}

func (x *ReleaseIPRequest) GetIPv6CIDR() string {
	if x != nil {
		return x.IPv6CIDR
	}
	return ""
}

type ReleaseIPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ReleaseIPResponse) Reset() {
	*x = ReleaseIPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dhcpserver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseIPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseIPResponse) ProtoMessage() {}

func (x *ReleaseIPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dhcpserver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseIPResponse.ProtoReflect.Descriptor instead.
func (*ReleaseIPResponse) Descriptor() ([]byte, []int) {
	return file_dhcpserver_proto_rawDescGZIP(), []int{3}
}

func (x *ReleaseIPResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_dhcpserver_proto protoreflect.FileDescriptor

var file_dhcpserver_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x68, 0x63, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x22, 0x4d, 0x0a, 0x0d, 0x52, 0x65, 0x6e, 0x74, 0x49,
	0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6f, 0x64, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x6f, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x50, 0x6f, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x48, 0x0a, 0x0e, 0x52, 0x65, 0x6e, 0x74, 0x49, 0x50,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x50, 0x76, 0x34,
	0x43, 0x49, 0x44, 0x52, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x50, 0x76, 0x34,
	0x43, 0x49, 0x44, 0x52, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x50, 0x76, 0x36, 0x43, 0x49, 0x44, 0x52,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x50, 0x76, 0x36, 0x43, 0x49, 0x44, 0x52,
	0x22, 0x88, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x50, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x50, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x50, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x50, 0x76, 0x34, 0x43, 0x49, 0x44, 0x52, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x50, 0x76, 0x34, 0x43, 0x49, 0x44, 0x52, 0x12,
	0x1a, 0x0a, 0x08, 0x49, 0x50, 0x76, 0x36, 0x43, 0x49, 0x44, 0x52, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x49, 0x50, 0x76, 0x36, 0x43, 0x49, 0x44, 0x52, 0x22, 0x2d, 0x0a, 0x11, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x79, 0x0a, 0x04, 0x44, 0x48,
	0x43, 0x50, 0x12, 0x33, 0x0a, 0x06, 0x52, 0x65, 0x6e, 0x74, 0x49, 0x50, 0x12, 0x12, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x52, 0x65, 0x6e, 0x74, 0x49, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x6e, 0x74, 0x49, 0x50, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x09, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x49, 0x50, 0x12, 0x15, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x49, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dhcpserver_proto_rawDescOnce sync.Once
	file_dhcpserver_proto_rawDescData = file_dhcpserver_proto_rawDesc
)

func file_dhcpserver_proto_rawDescGZIP() []byte {
	file_dhcpserver_proto_rawDescOnce.Do(func() {
		file_dhcpserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_dhcpserver_proto_rawDescData)
	})
	return file_dhcpserver_proto_rawDescData
}

var file_dhcpserver_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_dhcpserver_proto_goTypes = []interface{}{
	(*RentIPRequest)(nil),     // 0: rpc.RentIPRequest
	(*RentIPResponse)(nil),    // 1: rpc.RentIPResponse
	(*ReleaseIPRequest)(nil),  // 2: rpc.ReleaseIPRequest
	(*ReleaseIPResponse)(nil), // 3: rpc.ReleaseIPResponse
}
var file_dhcpserver_proto_depIdxs = []int32{
	0, // 0: rpc.DHCP.RentIP:input_type -> rpc.RentIPRequest
	2, // 1: rpc.DHCP.ReleaseIP:input_type -> rpc.ReleaseIPRequest
	1, // 2: rpc.DHCP.RentIP:output_type -> rpc.RentIPResponse
	3, // 3: rpc.DHCP.ReleaseIP:output_type -> rpc.ReleaseIPResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dhcpserver_proto_init() }
func file_dhcpserver_proto_init() {
	if File_dhcpserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dhcpserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RentIPRequest); i {
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
		file_dhcpserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RentIPResponse); i {
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
		file_dhcpserver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseIPRequest); i {
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
		file_dhcpserver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseIPResponse); i {
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
			RawDescriptor: file_dhcpserver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dhcpserver_proto_goTypes,
		DependencyIndexes: file_dhcpserver_proto_depIdxs,
		MessageInfos:      file_dhcpserver_proto_msgTypes,
	}.Build()
	File_dhcpserver_proto = out.File
	file_dhcpserver_proto_rawDesc = nil
	file_dhcpserver_proto_goTypes = nil
	file_dhcpserver_proto_depIdxs = nil
}
