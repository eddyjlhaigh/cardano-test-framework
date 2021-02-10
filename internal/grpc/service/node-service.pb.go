// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: internal/proto-files/service/node-service.proto

package service

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_files_service_node_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_files_service_node_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_internal_proto_files_service_node_service_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_files_service_node_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_files_service_node_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_internal_proto_files_service_node_service_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_internal_proto_files_service_node_service_proto protoreflect.FileDescriptor

var file_internal_proto_files_service_node_service_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x1e, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x32, 0x7a, 0x0a, 0x0b, 0x4e, 0x6f, 0x64,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6e,
	0x74, 0x44, 0x69, 0x72, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0d, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_files_service_node_service_proto_rawDescOnce sync.Once
	file_internal_proto_files_service_node_service_proto_rawDescData = file_internal_proto_files_service_node_service_proto_rawDesc
)

func file_internal_proto_files_service_node_service_proto_rawDescGZIP() []byte {
	file_internal_proto_files_service_node_service_proto_rawDescOnce.Do(func() {
		file_internal_proto_files_service_node_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_files_service_node_service_proto_rawDescData)
	})
	return file_internal_proto_files_service_node_service_proto_rawDescData
}

var file_internal_proto_files_service_node_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_proto_files_service_node_service_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: service.Request
	(*Response)(nil), // 1: service.Response
}
var file_internal_proto_files_service_node_service_proto_depIdxs = []int32{
	0, // 0: service.NodeService.PrintDir:input_type -> service.Request
	0, // 1: service.NodeService.FetchResponse:input_type -> service.Request
	1, // 2: service.NodeService.PrintDir:output_type -> service.Response
	1, // 3: service.NodeService.FetchResponse:output_type -> service.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_proto_files_service_node_service_proto_init() }
func file_internal_proto_files_service_node_service_proto_init() {
	if File_internal_proto_files_service_node_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_files_service_node_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_internal_proto_files_service_node_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_internal_proto_files_service_node_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_files_service_node_service_proto_goTypes,
		DependencyIndexes: file_internal_proto_files_service_node_service_proto_depIdxs,
		MessageInfos:      file_internal_proto_files_service_node_service_proto_msgTypes,
	}.Build()
	File_internal_proto_files_service_node_service_proto = out.File
	file_internal_proto_files_service_node_service_proto_rawDesc = nil
	file_internal_proto_files_service_node_service_proto_goTypes = nil
	file_internal_proto_files_service_node_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NodeServiceClient is the client API for NodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeServiceClient interface {
	PrintDir(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	FetchResponse(ctx context.Context, in *Request, opts ...grpc.CallOption) (NodeService_FetchResponseClient, error)
}

type nodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeServiceClient(cc grpc.ClientConnInterface) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) PrintDir(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/service.NodeService/PrintDir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) FetchResponse(ctx context.Context, in *Request, opts ...grpc.CallOption) (NodeService_FetchResponseClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NodeService_serviceDesc.Streams[0], "/service.NodeService/FetchResponse", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceFetchResponseClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_FetchResponseClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type nodeServiceFetchResponseClient struct {
	grpc.ClientStream
}

func (x *nodeServiceFetchResponseClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NodeServiceServer is the server API for NodeService service.
type NodeServiceServer interface {
	PrintDir(context.Context, *Request) (*Response, error)
	FetchResponse(*Request, NodeService_FetchResponseServer) error
}

// UnimplementedNodeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNodeServiceServer struct {
}

func (*UnimplementedNodeServiceServer) PrintDir(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrintDir not implemented")
}
func (*UnimplementedNodeServiceServer) FetchResponse(*Request, NodeService_FetchResponseServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchResponse not implemented")
}

func RegisterNodeServiceServer(s *grpc.Server, srv NodeServiceServer) {
	s.RegisterService(&_NodeService_serviceDesc, srv)
}

func _NodeService_PrintDir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).PrintDir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NodeService/PrintDir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).PrintDir(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_FetchResponse_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).FetchResponse(m, &nodeServiceFetchResponseServer{stream})
}

type NodeService_FetchResponseServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type nodeServiceFetchResponseServer struct {
	grpc.ServerStream
}

func (x *nodeServiceFetchResponseServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

var _NodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PrintDir",
			Handler:    _NodeService_PrintDir_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchResponse",
			Handler:       _NodeService_FetchResponse_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/proto-files/service/node-service.proto",
}
