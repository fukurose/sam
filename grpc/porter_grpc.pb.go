// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: porter.proto

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

const (
	PorterService_ListSegmentStream_FullMethodName = "/grpc.PorterService/ListSegmentStream"
	PorterService_OrderStream_FullMethodName       = "/grpc.PorterService/OrderStream"
)

// PorterServiceClient is the client API for PorterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PorterServiceClient interface {
	ListSegmentStream(ctx context.Context, in *LSRequest, opts ...grpc.CallOption) (PorterService_ListSegmentStreamClient, error)
	OrderStream(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (PorterService_OrderStreamClient, error)
}

type porterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPorterServiceClient(cc grpc.ClientConnInterface) PorterServiceClient {
	return &porterServiceClient{cc}
}

func (c *porterServiceClient) ListSegmentStream(ctx context.Context, in *LSRequest, opts ...grpc.CallOption) (PorterService_ListSegmentStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &PorterService_ServiceDesc.Streams[0], PorterService_ListSegmentStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &porterServiceListSegmentStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PorterService_ListSegmentStreamClient interface {
	Recv() (*LSResponse, error)
	grpc.ClientStream
}

type porterServiceListSegmentStreamClient struct {
	grpc.ClientStream
}

func (x *porterServiceListSegmentStreamClient) Recv() (*LSResponse, error) {
	m := new(LSResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *porterServiceClient) OrderStream(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (PorterService_OrderStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &PorterService_ServiceDesc.Streams[1], PorterService_OrderStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &porterServiceOrderStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PorterService_OrderStreamClient interface {
	Recv() (*OrderResponse, error)
	grpc.ClientStream
}

type porterServiceOrderStreamClient struct {
	grpc.ClientStream
}

func (x *porterServiceOrderStreamClient) Recv() (*OrderResponse, error) {
	m := new(OrderResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PorterServiceServer is the server API for PorterService service.
// All implementations must embed UnimplementedPorterServiceServer
// for forward compatibility
type PorterServiceServer interface {
	ListSegmentStream(*LSRequest, PorterService_ListSegmentStreamServer) error
	OrderStream(*OrderRequest, PorterService_OrderStreamServer) error
	mustEmbedUnimplementedPorterServiceServer()
}

// UnimplementedPorterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPorterServiceServer struct {
}

func (UnimplementedPorterServiceServer) ListSegmentStream(*LSRequest, PorterService_ListSegmentStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ListSegmentStream not implemented")
}
func (UnimplementedPorterServiceServer) OrderStream(*OrderRequest, PorterService_OrderStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method OrderStream not implemented")
}
func (UnimplementedPorterServiceServer) mustEmbedUnimplementedPorterServiceServer() {}

// UnsafePorterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PorterServiceServer will
// result in compilation errors.
type UnsafePorterServiceServer interface {
	mustEmbedUnimplementedPorterServiceServer()
}

func RegisterPorterServiceServer(s grpc.ServiceRegistrar, srv PorterServiceServer) {
	s.RegisterService(&PorterService_ServiceDesc, srv)
}

func _PorterService_ListSegmentStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LSRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PorterServiceServer).ListSegmentStream(m, &porterServiceListSegmentStreamServer{stream})
}

type PorterService_ListSegmentStreamServer interface {
	Send(*LSResponse) error
	grpc.ServerStream
}

type porterServiceListSegmentStreamServer struct {
	grpc.ServerStream
}

func (x *porterServiceListSegmentStreamServer) Send(m *LSResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PorterService_OrderStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OrderRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PorterServiceServer).OrderStream(m, &porterServiceOrderStreamServer{stream})
}

type PorterService_OrderStreamServer interface {
	Send(*OrderResponse) error
	grpc.ServerStream
}

type porterServiceOrderStreamServer struct {
	grpc.ServerStream
}

func (x *porterServiceOrderStreamServer) Send(m *OrderResponse) error {
	return x.ServerStream.SendMsg(m)
}

// PorterService_ServiceDesc is the grpc.ServiceDesc for PorterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PorterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.PorterService",
	HandlerType: (*PorterServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListSegmentStream",
			Handler:       _PorterService_ListSegmentStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OrderStream",
			Handler:       _PorterService_OrderStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "porter.proto",
}
