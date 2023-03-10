// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: list.proto

package listServer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ListServiceClient is the client API for ListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListServiceClient interface {
	FindListItem(ctx context.Context, in *ItemReq, opts ...grpc.CallOption) (*Item, error)
	InsertListItem(ctx context.Context, in *Item, opts ...grpc.CallOption) (*InsertResponse, error)
	DeleteListItem(ctx context.Context, in *ItemReq, opts ...grpc.CallOption) (*DeleteResponse, error)
	GetList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ListService_GetListClient, error)
	UpdateListItem(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*Item, error)
}

type listServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewListServiceClient(cc grpc.ClientConnInterface) ListServiceClient {
	return &listServiceClient{cc}
}

func (c *listServiceClient) FindListItem(ctx context.Context, in *ItemReq, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/ListService/findListItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listServiceClient) InsertListItem(ctx context.Context, in *Item, opts ...grpc.CallOption) (*InsertResponse, error) {
	out := new(InsertResponse)
	err := c.cc.Invoke(ctx, "/ListService/insertListItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listServiceClient) DeleteListItem(ctx context.Context, in *ItemReq, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/ListService/deleteListItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listServiceClient) GetList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ListService_GetListClient, error) {
	stream, err := c.cc.NewStream(ctx, &ListService_ServiceDesc.Streams[0], "/ListService/getList", opts...)
	if err != nil {
		return nil, err
	}
	x := &listServiceGetListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ListService_GetListClient interface {
	Recv() (*Item, error)
	grpc.ClientStream
}

type listServiceGetListClient struct {
	grpc.ClientStream
}

func (x *listServiceGetListClient) Recv() (*Item, error) {
	m := new(Item)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *listServiceClient) UpdateListItem(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/ListService/updateListItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListServiceServer is the server API for ListService service.
// All implementations must embed UnimplementedListServiceServer
// for forward compatibility
type ListServiceServer interface {
	FindListItem(context.Context, *ItemReq) (*Item, error)
	InsertListItem(context.Context, *Item) (*InsertResponse, error)
	DeleteListItem(context.Context, *ItemReq) (*DeleteResponse, error)
	GetList(*Empty, ListService_GetListServer) error
	UpdateListItem(context.Context, *UpdateReq) (*Item, error)
	mustEmbedUnimplementedListServiceServer()
}

// UnimplementedListServiceServer must be embedded to have forward compatible implementations.
type UnimplementedListServiceServer struct {
}

func (UnimplementedListServiceServer) FindListItem(context.Context, *ItemReq) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindListItem not implemented")
}
func (UnimplementedListServiceServer) InsertListItem(context.Context, *Item) (*InsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertListItem not implemented")
}
func (UnimplementedListServiceServer) DeleteListItem(context.Context, *ItemReq) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteListItem not implemented")
}
func (UnimplementedListServiceServer) GetList(*Empty, ListService_GetListServer) error {
	return status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedListServiceServer) UpdateListItem(context.Context, *UpdateReq) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateListItem not implemented")
}
func (UnimplementedListServiceServer) mustEmbedUnimplementedListServiceServer() {}

// UnsafeListServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListServiceServer will
// result in compilation errors.
type UnsafeListServiceServer interface {
	mustEmbedUnimplementedListServiceServer()
}

func RegisterListServiceServer(s grpc.ServiceRegistrar, srv ListServiceServer) {
	s.RegisterService(&ListService_ServiceDesc, srv)
}

func _ListService_FindListItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServiceServer).FindListItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ListService/findListItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServiceServer).FindListItem(ctx, req.(*ItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListService_InsertListItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServiceServer).InsertListItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ListService/insertListItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServiceServer).InsertListItem(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListService_DeleteListItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServiceServer).DeleteListItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ListService/deleteListItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServiceServer).DeleteListItem(ctx, req.(*ItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListService_GetList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ListServiceServer).GetList(m, &listServiceGetListServer{stream})
}

type ListService_GetListServer interface {
	Send(*Item) error
	grpc.ServerStream
}

type listServiceGetListServer struct {
	grpc.ServerStream
}

func (x *listServiceGetListServer) Send(m *Item) error {
	return x.ServerStream.SendMsg(m)
}

func _ListService_UpdateListItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServiceServer).UpdateListItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ListService/updateListItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServiceServer).UpdateListItem(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ListService_ServiceDesc is the grpc.ServiceDesc for ListService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ListService",
	HandlerType: (*ListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findListItem",
			Handler:    _ListService_FindListItem_Handler,
		},
		{
			MethodName: "insertListItem",
			Handler:    _ListService_InsertListItem_Handler,
		},
		{
			MethodName: "deleteListItem",
			Handler:    _ListService_DeleteListItem_Handler,
		},
		{
			MethodName: "updateListItem",
			Handler:    _ListService_UpdateListItem_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getList",
			Handler:       _ListService_GetList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "list.proto",
}
