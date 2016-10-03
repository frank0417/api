// Code generated by protoc-gen-go.
// source: client.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	client.proto
	server.proto

It has these top-level messages:
	ClientReconfigureRequest
	ServerCreateRequest
	ServerDeleteRequest
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClientReconfigureRequest struct {
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
}

func (m *ClientReconfigureRequest) Reset()                    { *m = ClientReconfigureRequest{} }
func (m *ClientReconfigureRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientReconfigureRequest) ProtoMessage()               {}
func (*ClientReconfigureRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*ClientReconfigureRequest)(nil), "backup.v1beta1.ClientReconfigureRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Clients service

type ClientsClient interface {
	Reconfigure(ctx context.Context, in *ClientReconfigureRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type clientsClient struct {
	cc *grpc.ClientConn
}

func NewClientsClient(cc *grpc.ClientConn) ClientsClient {
	return &clientsClient{cc}
}

func (c *clientsClient) Reconfigure(ctx context.Context, in *ClientReconfigureRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/backup.v1beta1.Clients/Reconfigure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Clients service

type ClientsServer interface {
	Reconfigure(context.Context, *ClientReconfigureRequest) (*dtypes.VoidResponse, error)
}

func RegisterClientsServer(s *grpc.Server, srv ClientsServer) {
	s.RegisterService(&_Clients_serviceDesc, srv)
}

func _Clients_Reconfigure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientReconfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServer).Reconfigure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup.v1beta1.Clients/Reconfigure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServer).Reconfigure(ctx, req.(*ClientReconfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Clients_serviceDesc = grpc.ServiceDesc{
	ServiceName: "backup.v1beta1.Clients",
	HandlerType: (*ClientsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reconfigure",
			Handler:    _Clients_Reconfigure_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("client.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4b, 0x4a, 0x4c, 0xce, 0x2e, 0x2d,
	0xd0, 0x2b, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0x94, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49,
	0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b,
	0x86, 0xa8, 0x96, 0x12, 0x03, 0x09, 0xa7, 0x94, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0x83, 0x49, 0x88,
	0xb8, 0x92, 0x2d, 0x97, 0x84, 0x33, 0xd8, 0xd4, 0xa0, 0xd4, 0xe4, 0xfc, 0xbc, 0xb4, 0xcc, 0xf4,
	0xd2, 0xa2, 0xd4, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x45, 0x90, 0x8d, 0xa5, 0xc5,
	0x25, 0xa9, 0x45, 0xf1, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xdc,
	0x50, 0x31, 0xbf, 0xc4, 0xdc, 0x54, 0xa3, 0x69, 0x8c, 0x5c, 0xec, 0x10, 0xfd, 0xc5, 0x42, 0x5d,
	0x8c, 0x5c, 0xdc, 0x48, 0xa6, 0x08, 0x69, 0xe8, 0xa1, 0xba, 0x50, 0x0f, 0x97, 0x45, 0x52, 0x22,
	0x7a, 0x10, 0x97, 0xe9, 0x85, 0xe5, 0x67, 0xa6, 0x04, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7,
	0x2a, 0x59, 0x37, 0x5d, 0x7e, 0x32, 0x99, 0xc9, 0x54, 0xca, 0x40, 0x3f, 0xb1, 0xa0, 0xa0, 0x38,
	0x39, 0x3f, 0x05, 0xe2, 0x37, 0x88, 0xa1, 0xfa, 0x50, 0x43, 0xf5, 0x21, 0x61, 0x52, 0xac, 0x5f,
	0x84, 0x30, 0xd6, 0x8a, 0x51, 0xcb, 0x89, 0x33, 0x8a, 0x1d, 0xaa, 0x22, 0x89, 0x0d, 0xec, 0x53,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xf1, 0xe3, 0xad, 0x3f, 0x01, 0x00, 0x00,
}
