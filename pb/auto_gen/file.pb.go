// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.3
// source: file.proto

package file

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

type FileRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *FileRegisterRequest) Reset() {
	*x = FileRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRegisterRequest) ProtoMessage() {}

func (x *FileRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRegisterRequest.ProtoReflect.Descriptor instead.
func (*FileRegisterRequest) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{0}
}

func (x *FileRegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileRegisterRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type FileRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Status     string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *FileRegisterResponse) Reset() {
	*x = FileRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRegisterResponse) ProtoMessage() {}

func (x *FileRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRegisterResponse.ProtoReflect.Descriptor instead.
func (*FileRegisterResponse) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{1}
}

func (x *FileRegisterResponse) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *FileRegisterResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type FileUploadStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *FileUploadStatusRequest) Reset() {
	*x = FileUploadStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadStatusRequest) ProtoMessage() {}

func (x *FileUploadStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadStatusRequest.ProtoReflect.Descriptor instead.
func (*FileUploadStatusRequest) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{2}
}

func (x *FileUploadStatusRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

type FileUploadStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *FileUploadStatusResponse) Reset() {
	*x = FileUploadStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadStatusResponse) ProtoMessage() {}

func (x *FileUploadStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadStatusResponse.ProtoReflect.Descriptor instead.
func (*FileUploadStatusResponse) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{3}
}

func (x *FileUploadStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_file_proto protoreflect.FileDescriptor

var file_file_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x13,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x4e, 0x0a, 0x14, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x39, 0x0a, 0x17, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x32, 0x0a, 0x18,
	0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x32, 0x45, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x35, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x14, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x52, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3e, 0x0a, 0x07, 0x49,
	0x6e, 0x71, 0x75, 0x69, 0x72, 0x65, 0x12, 0x18, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_file_proto_rawDescOnce sync.Once
	file_file_proto_rawDescData = file_file_proto_rawDesc
)

func file_file_proto_rawDescGZIP() []byte {
	file_file_proto_rawDescOnce.Do(func() {
		file_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_proto_rawDescData)
	})
	return file_file_proto_rawDescData
}

var file_file_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_file_proto_goTypes = []interface{}{
	(*FileRegisterRequest)(nil),      // 0: FileRegisterRequest
	(*FileRegisterResponse)(nil),     // 1: FileRegisterResponse
	(*FileUploadStatusRequest)(nil),  // 2: FileUploadStatusRequest
	(*FileUploadStatusResponse)(nil), // 3: FileUploadStatusResponse
}
var file_file_proto_depIdxs = []int32{
	0, // 0: FileUploader.Upload:input_type -> FileRegisterRequest
	2, // 1: FileUploadStatus.Inquire:input_type -> FileUploadStatusRequest
	1, // 2: FileUploader.Upload:output_type -> FileRegisterResponse
	3, // 3: FileUploadStatus.Inquire:output_type -> FileUploadStatusResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_file_proto_init() }
func file_file_proto_init() {
	if File_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRegisterRequest); i {
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
		file_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRegisterResponse); i {
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
		file_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadStatusRequest); i {
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
		file_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadStatusResponse); i {
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
			RawDescriptor: file_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_file_proto_goTypes,
		DependencyIndexes: file_file_proto_depIdxs,
		MessageInfos:      file_file_proto_msgTypes,
	}.Build()
	File_file_proto = out.File
	file_file_proto_rawDesc = nil
	file_file_proto_goTypes = nil
	file_file_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FileUploaderClient is the client API for FileUploader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileUploaderClient interface {
	Upload(ctx context.Context, in *FileRegisterRequest, opts ...grpc.CallOption) (*FileRegisterResponse, error)
}

type fileUploaderClient struct {
	cc grpc.ClientConnInterface
}

func NewFileUploaderClient(cc grpc.ClientConnInterface) FileUploaderClient {
	return &fileUploaderClient{cc}
}

func (c *fileUploaderClient) Upload(ctx context.Context, in *FileRegisterRequest, opts ...grpc.CallOption) (*FileRegisterResponse, error) {
	out := new(FileRegisterResponse)
	err := c.cc.Invoke(ctx, "/FileUploader/Upload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileUploaderServer is the server API for FileUploader service.
type FileUploaderServer interface {
	Upload(context.Context, *FileRegisterRequest) (*FileRegisterResponse, error)
}

// UnimplementedFileUploaderServer can be embedded to have forward compatible implementations.
type UnimplementedFileUploaderServer struct {
}

func (*UnimplementedFileUploaderServer) Upload(context.Context, *FileRegisterRequest) (*FileRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}

func RegisterFileUploaderServer(s *grpc.Server, srv FileUploaderServer) {
	s.RegisterService(&_FileUploader_serviceDesc, srv)
}

func _FileUploader_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileUploaderServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileUploader/Upload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileUploaderServer).Upload(ctx, req.(*FileRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileUploader_serviceDesc = grpc.ServiceDesc{
	ServiceName: "FileUploader",
	HandlerType: (*FileUploaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upload",
			Handler:    _FileUploader_Upload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "file.proto",
}

// FileUploadStatusClient is the client API for FileUploadStatus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileUploadStatusClient interface {
	Inquire(ctx context.Context, in *FileUploadStatusRequest, opts ...grpc.CallOption) (*FileUploadStatusResponse, error)
}

type fileUploadStatusClient struct {
	cc grpc.ClientConnInterface
}

func NewFileUploadStatusClient(cc grpc.ClientConnInterface) FileUploadStatusClient {
	return &fileUploadStatusClient{cc}
}

func (c *fileUploadStatusClient) Inquire(ctx context.Context, in *FileUploadStatusRequest, opts ...grpc.CallOption) (*FileUploadStatusResponse, error) {
	out := new(FileUploadStatusResponse)
	err := c.cc.Invoke(ctx, "/FileUploadStatus/Inquire", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileUploadStatusServer is the server API for FileUploadStatus service.
type FileUploadStatusServer interface {
	Inquire(context.Context, *FileUploadStatusRequest) (*FileUploadStatusResponse, error)
}

// UnimplementedFileUploadStatusServer can be embedded to have forward compatible implementations.
type UnimplementedFileUploadStatusServer struct {
}

func (*UnimplementedFileUploadStatusServer) Inquire(context.Context, *FileUploadStatusRequest) (*FileUploadStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Inquire not implemented")
}

func RegisterFileUploadStatusServer(s *grpc.Server, srv FileUploadStatusServer) {
	s.RegisterService(&_FileUploadStatus_serviceDesc, srv)
}

func _FileUploadStatus_Inquire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileUploadStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileUploadStatusServer).Inquire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileUploadStatus/Inquire",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileUploadStatusServer).Inquire(ctx, req.(*FileUploadStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileUploadStatus_serviceDesc = grpc.ServiceDesc{
	ServiceName: "FileUploadStatus",
	HandlerType: (*FileUploadStatusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Inquire",
			Handler:    _FileUploadStatus_Inquire_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "file.proto",
}
