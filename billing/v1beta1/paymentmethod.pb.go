// Code generated by protoc-gen-go.
// source: paymentmethod.proto
// DO NOT EDIT!

package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for PaymentMethods service

type PaymentMethodsClient interface {
	Check(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type paymentMethodsClient struct {
	cc *grpc.ClientConn
}

func NewPaymentMethodsClient(cc *grpc.ClientConn) PaymentMethodsClient {
	return &paymentMethodsClient{cc}
}

func (c *paymentMethodsClient) Check(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/billing.v1beta1.PaymentMethods/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PaymentMethods service

type PaymentMethodsServer interface {
	Check(context.Context, *dtypes.VoidRequest) (*dtypes.VoidResponse, error)
}

func RegisterPaymentMethodsServer(s *grpc.Server, srv PaymentMethodsServer) {
	s.RegisterService(&_PaymentMethods_serviceDesc, srv)
}

func _PaymentMethods_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentMethodsServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/billing.v1beta1.PaymentMethods/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentMethodsServer).Check(ctx, req.(*dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaymentMethods_serviceDesc = grpc.ServiceDesc{
	ServiceName: "billing.v1beta1.PaymentMethods",
	HandlerType: (*PaymentMethodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _PaymentMethods_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("paymentmethod.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x48, 0xac, 0xcc,
	0x4d, 0xcd, 0x2b, 0xc9, 0x4d, 0x2d, 0xc9, 0xc8, 0x4f, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x4f, 0xca, 0xcc, 0xc9, 0xc9, 0xcc, 0x4b, 0xd7, 0x2b, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34,
	0x94, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb,
	0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x28, 0x97, 0x92, 0x4b, 0x2c, 0x28, 0x28,
	0x4e, 0xce, 0x4f, 0xc1, 0x25, 0x2f, 0x06, 0x12, 0x4e, 0x29, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x07,
	0x93, 0x10, 0x71, 0xa3, 0x06, 0x46, 0x2e, 0xbe, 0x00, 0x88, 0xf5, 0xbe, 0x60, 0xeb, 0x8b, 0x85,
	0xf2, 0xb8, 0x58, 0x9d, 0x33, 0x52, 0x93, 0xb3, 0x85, 0x84, 0xf5, 0x20, 0x1a, 0xf4, 0xc2, 0xf2,
	0x33, 0x53, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xa4, 0x44, 0x50, 0x05, 0x8b, 0x0b, 0xf2,
	0xf3, 0x8a, 0x53, 0x95, 0xac, 0x9a, 0xb6, 0x4a, 0x30, 0x71, 0x30, 0x36, 0x5d, 0x7e, 0x32, 0x99,
	0x49, 0x4f, 0x48, 0x47, 0x1f, 0xc5, 0x35, 0x50, 0x9f, 0xe8, 0x43, 0x7d, 0xa2, 0x8f, 0xe2, 0xdd,
	0x62, 0x27, 0xce, 0x28, 0x76, 0xa8, 0x4c, 0x12, 0x1b, 0xd8, 0x51, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x02, 0x22, 0xee, 0x5b, 0x12, 0x01, 0x00, 0x00,
}
