// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: schedule.proto

package models

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ScheduleService_GetSchedule_FullMethodName    = "/schedule.ScheduleService/GetSchedule"
	ScheduleService_CreateSchedule_FullMethodName = "/schedule.ScheduleService/CreateSchedule"
)

// ScheduleServiceClient is the client API for ScheduleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScheduleServiceClient interface {
	GetSchedule(ctx context.Context, in *GetScheduleRequest, opts ...grpc.CallOption) (*GetScheduleResponse, error)
	CreateSchedule(ctx context.Context, in *CreateScheduleRequest, opts ...grpc.CallOption) (*CreateScheduleResponse, error)
}

type scheduleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScheduleServiceClient(cc grpc.ClientConnInterface) ScheduleServiceClient {
	return &scheduleServiceClient{cc}
}

func (c *scheduleServiceClient) GetSchedule(ctx context.Context, in *GetScheduleRequest, opts ...grpc.CallOption) (*GetScheduleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetScheduleResponse)
	err := c.cc.Invoke(ctx, ScheduleService_GetSchedule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) CreateSchedule(ctx context.Context, in *CreateScheduleRequest, opts ...grpc.CallOption) (*CreateScheduleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateScheduleResponse)
	err := c.cc.Invoke(ctx, ScheduleService_CreateSchedule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScheduleServiceServer is the server API for ScheduleService service.
// All implementations must embed UnimplementedScheduleServiceServer
// for forward compatibility.
type ScheduleServiceServer interface {
	GetSchedule(context.Context, *GetScheduleRequest) (*GetScheduleResponse, error)
	CreateSchedule(context.Context, *CreateScheduleRequest) (*CreateScheduleResponse, error)
	mustEmbedUnimplementedScheduleServiceServer()
}

// UnimplementedScheduleServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedScheduleServiceServer struct{}

func (UnimplementedScheduleServiceServer) GetSchedule(context.Context, *GetScheduleRequest) (*GetScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchedule not implemented")
}
func (UnimplementedScheduleServiceServer) CreateSchedule(context.Context, *CreateScheduleRequest) (*CreateScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSchedule not implemented")
}
func (UnimplementedScheduleServiceServer) mustEmbedUnimplementedScheduleServiceServer() {}
func (UnimplementedScheduleServiceServer) testEmbeddedByValue()                         {}

// UnsafeScheduleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScheduleServiceServer will
// result in compilation errors.
type UnsafeScheduleServiceServer interface {
	mustEmbedUnimplementedScheduleServiceServer()
}

func RegisterScheduleServiceServer(s grpc.ServiceRegistrar, srv ScheduleServiceServer) {
	// If the following call pancis, it indicates UnimplementedScheduleServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ScheduleService_ServiceDesc, srv)
}

func _ScheduleService_GetSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).GetSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScheduleService_GetSchedule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).GetSchedule(ctx, req.(*GetScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_CreateSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).CreateSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScheduleService_CreateSchedule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).CreateSchedule(ctx, req.(*CreateScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ScheduleService_ServiceDesc is the grpc.ServiceDesc for ScheduleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScheduleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "schedule.ScheduleService",
	HandlerType: (*ScheduleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSchedule",
			Handler:    _ScheduleService_GetSchedule_Handler,
		},
		{
			MethodName: "CreateSchedule",
			Handler:    _ScheduleService_CreateSchedule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schedule.proto",
}
