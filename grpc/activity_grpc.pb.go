// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: activity.proto

package grpc

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

// Activity_LogClient is the client API for Activity_Log service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Activity_LogClient interface {
	Insert(ctx context.Context, in *Activity, opts ...grpc.CallOption) (*InsertResponse, error)
	Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*Activity, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Activities, error)
}

type activity_LogClient struct {
	cc grpc.ClientConnInterface
}

func NewActivity_LogClient(cc grpc.ClientConnInterface) Activity_LogClient {
	return &activity_LogClient{cc}
}

func (c *activity_LogClient) Insert(ctx context.Context, in *Activity, opts ...grpc.CallOption) (*InsertResponse, error) {
	out := new(InsertResponse)
	err := c.cc.Invoke(ctx, "/grpc.Activity_Log/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activity_LogClient) Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*Activity, error) {
	out := new(Activity)
	err := c.cc.Invoke(ctx, "/grpc.Activity_Log/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activity_LogClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Activities, error) {
	out := new(Activities)
	err := c.cc.Invoke(ctx, "/grpc.Activity_Log/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Activity_LogServer is the server API for Activity_Log service.
// All implementations must embed UnimplementedActivity_LogServer
// for forward compatibility
type Activity_LogServer interface {
	Insert(context.Context, *Activity) (*InsertResponse, error)
	Retrieve(context.Context, *RetrieveRequest) (*Activity, error)
	List(context.Context, *ListRequest) (*Activities, error)
	mustEmbedUnimplementedActivity_LogServer()
}

// UnimplementedActivity_LogServer must be embedded to have forward compatible implementations.
type UnimplementedActivity_LogServer struct {
}

func (UnimplementedActivity_LogServer) Insert(context.Context, *Activity) (*InsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedActivity_LogServer) Retrieve(context.Context, *RetrieveRequest) (*Activity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedActivity_LogServer) List(context.Context, *ListRequest) (*Activities, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedActivity_LogServer) mustEmbedUnimplementedActivity_LogServer() {}

// UnsafeActivity_LogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Activity_LogServer will
// result in compilation errors.
type UnsafeActivity_LogServer interface {
	mustEmbedUnimplementedActivity_LogServer()
}

func RegisterActivity_LogServer(s grpc.ServiceRegistrar, srv Activity_LogServer) {
	s.RegisterService(&Activity_Log_ServiceDesc, srv)
}

func _Activity_Log_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Activity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Activity_LogServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Activity_Log/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Activity_LogServer).Insert(ctx, req.(*Activity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_Log_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Activity_LogServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Activity_Log/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Activity_LogServer).Retrieve(ctx, req.(*RetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_Log_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Activity_LogServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Activity_Log/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Activity_LogServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Activity_Log_ServiceDesc is the grpc.ServiceDesc for Activity_Log service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Activity_Log_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Activity_Log",
	HandlerType: (*Activity_LogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _Activity_Log_Insert_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _Activity_Log_Retrieve_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Activity_Log_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "activity.proto",
}
