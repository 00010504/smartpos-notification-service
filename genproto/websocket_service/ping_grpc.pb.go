// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: ping.proto

package websocket_service

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

// WebsocketServiceClient is the client API for WebsocketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebsocketServiceClient interface {
	Ping(ctx context.Context, in *PingPong, opts ...grpc.CallOption) (*PingPong, error)
}

type websocketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWebsocketServiceClient(cc grpc.ClientConnInterface) WebsocketServiceClient {
	return &websocketServiceClient{cc}
}

func (c *websocketServiceClient) Ping(ctx context.Context, in *PingPong, opts ...grpc.CallOption) (*PingPong, error) {
	out := new(PingPong)
	err := c.cc.Invoke(ctx, "/WebsocketService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebsocketServiceServer is the server API for WebsocketService service.
// All implementations should embed UnimplementedWebsocketServiceServer
// for forward compatibility
type WebsocketServiceServer interface {
	Ping(context.Context, *PingPong) (*PingPong, error)
}

// UnimplementedWebsocketServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWebsocketServiceServer struct {
}

func (UnimplementedWebsocketServiceServer) Ping(context.Context, *PingPong) (*PingPong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

// UnsafeWebsocketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebsocketServiceServer will
// result in compilation errors.
type UnsafeWebsocketServiceServer interface {
	mustEmbedUnimplementedWebsocketServiceServer()
}

func RegisterWebsocketServiceServer(s grpc.ServiceRegistrar, srv WebsocketServiceServer) {
	s.RegisterService(&WebsocketService_ServiceDesc, srv)
}

func _WebsocketService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingPong)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebsocketServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WebsocketService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebsocketServiceServer).Ping(ctx, req.(*PingPong))
	}
	return interceptor(ctx, in, info, handler)
}

// WebsocketService_ServiceDesc is the grpc.ServiceDesc for WebsocketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebsocketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "WebsocketService",
	HandlerType: (*WebsocketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _WebsocketService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ping.proto",
}
