// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: entpb/entpb.proto

package entpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DirectoryServiceClient is the client API for DirectoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DirectoryServiceClient interface {
	Create(ctx context.Context, in *CreateDirectoryRequest, opts ...grpc.CallOption) (*Directory, error)
	Get(ctx context.Context, in *GetDirectoryRequest, opts ...grpc.CallOption) (*Directory, error)
	Update(ctx context.Context, in *UpdateDirectoryRequest, opts ...grpc.CallOption) (*Directory, error)
	Delete(ctx context.Context, in *DeleteDirectoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	List(ctx context.Context, in *ListDirectoryRequest, opts ...grpc.CallOption) (*ListDirectoryResponse, error)
}

type directoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDirectoryServiceClient(cc grpc.ClientConnInterface) DirectoryServiceClient {
	return &directoryServiceClient{cc}
}

func (c *directoryServiceClient) Create(ctx context.Context, in *CreateDirectoryRequest, opts ...grpc.CallOption) (*Directory, error) {
	out := new(Directory)
	err := c.cc.Invoke(ctx, "/entpb.DirectoryService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryServiceClient) Get(ctx context.Context, in *GetDirectoryRequest, opts ...grpc.CallOption) (*Directory, error) {
	out := new(Directory)
	err := c.cc.Invoke(ctx, "/entpb.DirectoryService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryServiceClient) Update(ctx context.Context, in *UpdateDirectoryRequest, opts ...grpc.CallOption) (*Directory, error) {
	out := new(Directory)
	err := c.cc.Invoke(ctx, "/entpb.DirectoryService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryServiceClient) Delete(ctx context.Context, in *DeleteDirectoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/entpb.DirectoryService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryServiceClient) List(ctx context.Context, in *ListDirectoryRequest, opts ...grpc.CallOption) (*ListDirectoryResponse, error) {
	out := new(ListDirectoryResponse)
	err := c.cc.Invoke(ctx, "/entpb.DirectoryService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DirectoryServiceServer is the server API for DirectoryService service.
// All implementations must embed UnimplementedDirectoryServiceServer
// for forward compatibility
type DirectoryServiceServer interface {
	Create(context.Context, *CreateDirectoryRequest) (*Directory, error)
	Get(context.Context, *GetDirectoryRequest) (*Directory, error)
	Update(context.Context, *UpdateDirectoryRequest) (*Directory, error)
	Delete(context.Context, *DeleteDirectoryRequest) (*emptypb.Empty, error)
	List(context.Context, *ListDirectoryRequest) (*ListDirectoryResponse, error)
	mustEmbedUnimplementedDirectoryServiceServer()
}

// UnimplementedDirectoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDirectoryServiceServer struct {
}

func (UnimplementedDirectoryServiceServer) Create(context.Context, *CreateDirectoryRequest) (*Directory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDirectoryServiceServer) Get(context.Context, *GetDirectoryRequest) (*Directory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDirectoryServiceServer) Update(context.Context, *UpdateDirectoryRequest) (*Directory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDirectoryServiceServer) Delete(context.Context, *DeleteDirectoryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDirectoryServiceServer) List(context.Context, *ListDirectoryRequest) (*ListDirectoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDirectoryServiceServer) mustEmbedUnimplementedDirectoryServiceServer() {}

// UnsafeDirectoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DirectoryServiceServer will
// result in compilation errors.
type UnsafeDirectoryServiceServer interface {
	mustEmbedUnimplementedDirectoryServiceServer()
}

func RegisterDirectoryServiceServer(s grpc.ServiceRegistrar, srv DirectoryServiceServer) {
	s.RegisterService(&DirectoryService_ServiceDesc, srv)
}

func _DirectoryService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.DirectoryService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServiceServer).Create(ctx, req.(*CreateDirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.DirectoryService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServiceServer).Get(ctx, req.(*GetDirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.DirectoryService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServiceServer).Update(ctx, req.(*UpdateDirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.DirectoryService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServiceServer).Delete(ctx, req.(*DeleteDirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.DirectoryService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServiceServer).List(ctx, req.(*ListDirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DirectoryService_ServiceDesc is the grpc.ServiceDesc for DirectoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DirectoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entpb.DirectoryService",
	HandlerType: (*DirectoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DirectoryService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DirectoryService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DirectoryService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DirectoryService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DirectoryService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entpb/entpb.proto",
}
