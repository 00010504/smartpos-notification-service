// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: sms.proto

package notification_service

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

// SMSClient is the client API for SMS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SMSClient interface {
	SendSms(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*SendSmsRespone, error)
}

type sMSClient struct {
	cc grpc.ClientConnInterface
}

func NewSMSClient(cc grpc.ClientConnInterface) SMSClient {
	return &sMSClient{cc}
}

func (c *sMSClient) SendSms(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*SendSmsRespone, error) {
	out := new(SendSmsRespone)
	err := c.cc.Invoke(ctx, "/SMS/SendSms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SMSServer is the server API for SMS service.
// All implementations should embed UnimplementedSMSServer
// for forward compatibility
type SMSServer interface {
	SendSms(context.Context, *SendSmsRequest) (*SendSmsRespone, error)
}

// UnimplementedSMSServer should be embedded to have forward compatible implementations.
type UnimplementedSMSServer struct {
}

func (UnimplementedSMSServer) SendSms(context.Context, *SendSmsRequest) (*SendSmsRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSms not implemented")
}

// UnsafeSMSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SMSServer will
// result in compilation errors.
type UnsafeSMSServer interface {
	mustEmbedUnimplementedSMSServer()
}

func RegisterSMSServer(s grpc.ServiceRegistrar, srv SMSServer) {
	s.RegisterService(&SMS_ServiceDesc, srv)
}

func _SMS_SendSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSServer).SendSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SMS/SendSms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSServer).SendSms(ctx, req.(*SendSmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SMS_ServiceDesc is the grpc.ServiceDesc for SMS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SMS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SMS",
	HandlerType: (*SMSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSms",
			Handler:    _SMS_SendSms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sms.proto",
}
