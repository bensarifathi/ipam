// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.6.1
// source: v1/pb/ipam.proto

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

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subnet      string `protobuf:"bytes,1,opt,name=subnet,proto3" json:"subnet,omitempty"`
	Gateway     string `protobuf:"bytes,2,opt,name=gateway,proto3" json:"gateway,omitempty"`
	ContainerId string `protobuf:"bytes,3,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pb_ipam_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pb_ipam_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_v1_pb_ipam_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetSubnet() string {
	if x != nil {
		return x.Subnet
	}
	return ""
}

func (x *AddRequest) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

func (x *AddRequest) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodIp   string `protobuf:"bytes,1,opt,name=pod_ip,json=podIp,proto3" json:"pod_ip,omitempty"`
	Gateway string `protobuf:"bytes,2,opt,name=gateway,proto3" json:"gateway,omitempty"`
	NetMask int32  `protobuf:"varint,3,opt,name=net_mask,json=netMask,proto3" json:"net_mask,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pb_ipam_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pb_ipam_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_v1_pb_ipam_proto_rawDescGZIP(), []int{1}
}

func (x *AddResponse) GetPodIp() string {
	if x != nil {
		return x.PodIp
	}
	return ""
}

func (x *AddResponse) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

func (x *AddResponse) GetNetMask() int32 {
	if x != nil {
		return x.NetMask
	}
	return 0
}

type DelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerId string `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
}

func (x *DelRequest) Reset() {
	*x = DelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pb_ipam_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelRequest) ProtoMessage() {}

func (x *DelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pb_ipam_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelRequest.ProtoReflect.Descriptor instead.
func (*DelRequest) Descriptor() ([]byte, []int) {
	return file_v1_pb_ipam_proto_rawDescGZIP(), []int{2}
}

func (x *DelRequest) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

type DelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DelResponse) Reset() {
	*x = DelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pb_ipam_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelResponse) ProtoMessage() {}

func (x *DelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pb_ipam_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelResponse.ProtoReflect.Descriptor instead.
func (*DelResponse) Descriptor() ([]byte, []int) {
	return file_v1_pb_ipam_proto_rawDescGZIP(), []int{3}
}

var File_v1_pb_ipam_proto protoreflect.FileDescriptor

var file_v1_pb_ipam_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x70, 0x62, 0x2f, 0x69, 0x70, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x69, 0x70, 0x61, 0x6d, 0x22, 0x61, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x59, 0x0a, 0x0b, 0x61,
	0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x6f,
	0x64, 0x5f, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6f, 0x64, 0x49,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x6e,
	0x65, 0x74, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6e,
	0x65, 0x74, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x2f, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x5e, 0x0a, 0x04, 0x69, 0x70, 0x61, 0x6d, 0x12, 0x2a,
	0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x10, 0x2e, 0x69, 0x70, 0x61, 0x6d, 0x2e, 0x61, 0x64, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x69, 0x70, 0x61, 0x6d, 0x2e, 0x61,
	0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x03, 0x44, 0x65,
	0x6c, 0x12, 0x10, 0x2e, 0x69, 0x70, 0x61, 0x6d, 0x2e, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x69, 0x70, 0x61, 0x6d, 0x2e, 0x64, 0x65, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_pb_ipam_proto_rawDescOnce sync.Once
	file_v1_pb_ipam_proto_rawDescData = file_v1_pb_ipam_proto_rawDesc
)

func file_v1_pb_ipam_proto_rawDescGZIP() []byte {
	file_v1_pb_ipam_proto_rawDescOnce.Do(func() {
		file_v1_pb_ipam_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_pb_ipam_proto_rawDescData)
	})
	return file_v1_pb_ipam_proto_rawDescData
}

var file_v1_pb_ipam_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_pb_ipam_proto_goTypes = []interface{}{
	(*AddRequest)(nil),  // 0: ipam.addRequest
	(*AddResponse)(nil), // 1: ipam.addResponse
	(*DelRequest)(nil),  // 2: ipam.delRequest
	(*DelResponse)(nil), // 3: ipam.delResponse
}
var file_v1_pb_ipam_proto_depIdxs = []int32{
	0, // 0: ipam.ipam.Add:input_type -> ipam.addRequest
	2, // 1: ipam.ipam.Del:input_type -> ipam.delRequest
	1, // 2: ipam.ipam.Add:output_type -> ipam.addResponse
	3, // 3: ipam.ipam.Del:output_type -> ipam.delResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_pb_ipam_proto_init() }
func file_v1_pb_ipam_proto_init() {
	if File_v1_pb_ipam_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_pb_ipam_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_v1_pb_ipam_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResponse); i {
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
		file_v1_pb_ipam_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelRequest); i {
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
		file_v1_pb_ipam_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelResponse); i {
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
			RawDescriptor: file_v1_pb_ipam_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_pb_ipam_proto_goTypes,
		DependencyIndexes: file_v1_pb_ipam_proto_depIdxs,
		MessageInfos:      file_v1_pb_ipam_proto_msgTypes,
	}.Build()
	File_v1_pb_ipam_proto = out.File
	file_v1_pb_ipam_proto_rawDesc = nil
	file_v1_pb_ipam_proto_goTypes = nil
	file_v1_pb_ipam_proto_depIdxs = nil
}
