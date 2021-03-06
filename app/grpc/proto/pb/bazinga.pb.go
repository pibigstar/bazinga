// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: app/grpc/proto/bazinga.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RandomStoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keywords string `protobuf:"bytes,1,opt,name=keywords,proto3" json:"keywords,omitempty"`
}

func (x *RandomStoryReq) Reset() {
	*x = RandomStoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_grpc_proto_bazinga_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomStoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomStoryReq) ProtoMessage() {}

func (x *RandomStoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_grpc_proto_bazinga_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomStoryReq.ProtoReflect.Descriptor instead.
func (*RandomStoryReq) Descriptor() ([]byte, []int) {
	return file_app_grpc_proto_bazinga_proto_rawDescGZIP(), []int{0}
}

func (x *RandomStoryReq) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

type RandomStoryResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *RandomStoryResp) Reset() {
	*x = RandomStoryResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_grpc_proto_bazinga_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomStoryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomStoryResp) ProtoMessage() {}

func (x *RandomStoryResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_grpc_proto_bazinga_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomStoryResp.ProtoReflect.Descriptor instead.
func (*RandomStoryResp) Descriptor() ([]byte, []int) {
	return file_app_grpc_proto_bazinga_proto_rawDescGZIP(), []int{1}
}

func (x *RandomStoryResp) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_app_grpc_proto_bazinga_proto protoreflect.FileDescriptor

var file_app_grpc_proto_bazinga_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x62, 0x61, 0x7a, 0x69, 0x6e, 0x67, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c,
	0x0a, 0x0e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x2b, 0x0a, 0x0f,
	0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x42, 0x0a, 0x0e, 0x42, 0x61, 0x7a,
	0x69, 0x6e, 0x67, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x52,
	0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0f, 0x2e, 0x52, 0x61, 0x6e,
	0x64, 0x6f, 0x6d, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x52, 0x61,
	0x6e, 0x64, 0x6f, 0x6d, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x42, 0x16, 0x5a,
	0x14, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_grpc_proto_bazinga_proto_rawDescOnce sync.Once
	file_app_grpc_proto_bazinga_proto_rawDescData = file_app_grpc_proto_bazinga_proto_rawDesc
)

func file_app_grpc_proto_bazinga_proto_rawDescGZIP() []byte {
	file_app_grpc_proto_bazinga_proto_rawDescOnce.Do(func() {
		file_app_grpc_proto_bazinga_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_grpc_proto_bazinga_proto_rawDescData)
	})
	return file_app_grpc_proto_bazinga_proto_rawDescData
}

var file_app_grpc_proto_bazinga_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_grpc_proto_bazinga_proto_goTypes = []interface{}{
	(*RandomStoryReq)(nil),  // 0: RandomStoryReq
	(*RandomStoryResp)(nil), // 1: RandomStoryResp
}
var file_app_grpc_proto_bazinga_proto_depIdxs = []int32{
	0, // 0: BazingaService.RandomStory:input_type -> RandomStoryReq
	1, // 1: BazingaService.RandomStory:output_type -> RandomStoryResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_grpc_proto_bazinga_proto_init() }
func file_app_grpc_proto_bazinga_proto_init() {
	if File_app_grpc_proto_bazinga_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_grpc_proto_bazinga_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandomStoryReq); i {
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
		file_app_grpc_proto_bazinga_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandomStoryResp); i {
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
			RawDescriptor: file_app_grpc_proto_bazinga_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_grpc_proto_bazinga_proto_goTypes,
		DependencyIndexes: file_app_grpc_proto_bazinga_proto_depIdxs,
		MessageInfos:      file_app_grpc_proto_bazinga_proto_msgTypes,
	}.Build()
	File_app_grpc_proto_bazinga_proto = out.File
	file_app_grpc_proto_bazinga_proto_rawDesc = nil
	file_app_grpc_proto_bazinga_proto_goTypes = nil
	file_app_grpc_proto_bazinga_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BazingaServiceClient is the client API for BazingaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BazingaServiceClient interface {
	RandomStory(ctx context.Context, in *RandomStoryReq, opts ...grpc.CallOption) (*RandomStoryResp, error)
}

type bazingaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBazingaServiceClient(cc grpc.ClientConnInterface) BazingaServiceClient {
	return &bazingaServiceClient{cc}
}

func (c *bazingaServiceClient) RandomStory(ctx context.Context, in *RandomStoryReq, opts ...grpc.CallOption) (*RandomStoryResp, error) {
	out := new(RandomStoryResp)
	err := c.cc.Invoke(ctx, "/BazingaService/RandomStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BazingaServiceServer is the server API for BazingaService service.
type BazingaServiceServer interface {
	RandomStory(context.Context, *RandomStoryReq) (*RandomStoryResp, error)
}

// UnimplementedBazingaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBazingaServiceServer struct {
}

func (*UnimplementedBazingaServiceServer) RandomStory(context.Context, *RandomStoryReq) (*RandomStoryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RandomStory not implemented")
}

func RegisterBazingaServiceServer(s *grpc.Server, srv BazingaServiceServer) {
	s.RegisterService(&_BazingaService_serviceDesc, srv)
}

func _BazingaService_RandomStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandomStoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BazingaServiceServer).RandomStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BazingaService/RandomStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BazingaServiceServer).RandomStory(ctx, req.(*RandomStoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _BazingaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "BazingaService",
	HandlerType: (*BazingaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RandomStory",
			Handler:    _BazingaService_RandomStory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/grpc/proto/bazinga.proto",
}
