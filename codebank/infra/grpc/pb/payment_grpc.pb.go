// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: protofile/payment.proto

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PaymnetServiceClient is the client API for PaymnetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymnetServiceClient interface {
	Payment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type paymnetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymnetServiceClient(cc grpc.ClientConnInterface) PaymnetServiceClient {
	return &paymnetServiceClient{cc}
}

func (c *paymnetServiceClient) Payment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/payment.PaymnetService/Payment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymnetServiceServer is the server API for PaymnetService service.
// All implementations must embed UnimplementedPaymnetServiceServer
// for forward compatibility
type PaymnetServiceServer interface {
	Payment(context.Context, *PaymentRequest) (*empty.Empty, error)
	mustEmbedUnimplementedPaymnetServiceServer()
}

// UnimplementedPaymnetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPaymnetServiceServer struct {
}

func (UnimplementedPaymnetServiceServer) Payment(context.Context, *PaymentRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Payment not implemented")
}
func (UnimplementedPaymnetServiceServer) mustEmbedUnimplementedPaymnetServiceServer() {}

// UnsafePaymnetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymnetServiceServer will
// result in compilation errors.
type UnsafePaymnetServiceServer interface {
	mustEmbedUnimplementedPaymnetServiceServer()
}

func RegisterPaymnetServiceServer(s grpc.ServiceRegistrar, srv PaymnetServiceServer) {
	s.RegisterService(&PaymnetService_ServiceDesc, srv)
}

func _PaymnetService_Payment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymnetServiceServer).Payment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.PaymnetService/Payment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymnetServiceServer).Payment(ctx, req.(*PaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymnetService_ServiceDesc is the grpc.ServiceDesc for PaymnetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymnetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.PaymnetService",
	HandlerType: (*PaymnetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Payment",
			Handler:    _PaymnetService_Payment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protofile/payment.proto",
}