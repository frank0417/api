// Code generated by protoc-gen-go.
// source: pvc.proto
// DO NOT EDIT!

package pv

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"
import dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PVCRegisterRequest struct {
	Cluster   string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Size      int64  `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	Namespace string `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *PVCRegisterRequest) Reset()                    { *m = PVCRegisterRequest{} }
func (m *PVCRegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*PVCRegisterRequest) ProtoMessage()               {}
func (*PVCRegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type PVCUnregisterRequest struct {
	Cluster   string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *PVCUnregisterRequest) Reset()                    { *m = PVCUnregisterRequest{} }
func (m *PVCUnregisterRequest) String() string            { return proto.CompactTextString(m) }
func (*PVCUnregisterRequest) ProtoMessage()               {}
func (*PVCUnregisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type PVCDescribeRequest struct {
	Cluster   string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *PVCDescribeRequest) Reset()                    { *m = PVCDescribeRequest{} }
func (m *PVCDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*PVCDescribeRequest) ProtoMessage()               {}
func (*PVCDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type PVCInfo struct {
	Name        string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Size        int64    `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Namespace   string   `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	Status      string   `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	Volume      string   `protobuf:"bytes,5,opt,name=volume" json:"volume,omitempty"`
	AccessModes []string `protobuf:"bytes,6,rep,name=accessModes" json:"accessModes,omitempty"`
}

func (m *PVCInfo) Reset()                    { *m = PVCInfo{} }
func (m *PVCInfo) String() string            { return proto.CompactTextString(m) }
func (*PVCInfo) ProtoMessage()               {}
func (*PVCInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type PVCDescribeResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Pvc    *PVCInfo       `protobuf:"bytes,2,opt,name=pvc" json:"pvc,omitempty"`
}

func (m *PVCDescribeResponse) Reset()                    { *m = PVCDescribeResponse{} }
func (m *PVCDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*PVCDescribeResponse) ProtoMessage()               {}
func (*PVCDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *PVCDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *PVCDescribeResponse) GetPvc() *PVCInfo {
	if m != nil {
		return m.Pvc
	}
	return nil
}

func init() {
	proto.RegisterType((*PVCRegisterRequest)(nil), "pv.PVCRegisterRequest")
	proto.RegisterType((*PVCUnregisterRequest)(nil), "pv.PVCUnregisterRequest")
	proto.RegisterType((*PVCDescribeRequest)(nil), "pv.PVCDescribeRequest")
	proto.RegisterType((*PVCInfo)(nil), "pv.PVCInfo")
	proto.RegisterType((*PVCDescribeResponse)(nil), "pv.PVCDescribeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for PersistentVolumeClaims service

type PersistentVolumeClaimsClient interface {
	Describe(ctx context.Context, in *PVCDescribeRequest, opts ...grpc.CallOption) (*PVCDescribeResponse, error)
	Register(ctx context.Context, in *PVCRegisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Unregister(ctx context.Context, in *PVCUnregisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type persistentVolumeClaimsClient struct {
	cc *grpc.ClientConn
}

func NewPersistentVolumeClaimsClient(cc *grpc.ClientConn) PersistentVolumeClaimsClient {
	return &persistentVolumeClaimsClient{cc}
}

func (c *persistentVolumeClaimsClient) Describe(ctx context.Context, in *PVCDescribeRequest, opts ...grpc.CallOption) (*PVCDescribeResponse, error) {
	out := new(PVCDescribeResponse)
	err := grpc.Invoke(ctx, "/pv.PersistentVolumeClaims/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumeClaimsClient) Register(ctx context.Context, in *PVCRegisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/pv.PersistentVolumeClaims/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumeClaimsClient) Unregister(ctx context.Context, in *PVCUnregisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/pv.PersistentVolumeClaims/Unregister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PersistentVolumeClaims service

type PersistentVolumeClaimsServer interface {
	Describe(context.Context, *PVCDescribeRequest) (*PVCDescribeResponse, error)
	Register(context.Context, *PVCRegisterRequest) (*dtypes.VoidResponse, error)
	Unregister(context.Context, *PVCUnregisterRequest) (*dtypes.VoidResponse, error)
}

func RegisterPersistentVolumeClaimsServer(s *grpc.Server, srv PersistentVolumeClaimsServer) {
	s.RegisterService(&_PersistentVolumeClaims_serviceDesc, srv)
}

func _PersistentVolumeClaims_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PVCDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumeClaimsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.PersistentVolumeClaims/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumeClaimsServer).Describe(ctx, req.(*PVCDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersistentVolumeClaims_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PVCRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumeClaimsServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.PersistentVolumeClaims/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumeClaimsServer).Register(ctx, req.(*PVCRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersistentVolumeClaims_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PVCUnregisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumeClaimsServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.PersistentVolumeClaims/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumeClaimsServer).Unregister(ctx, req.(*PVCUnregisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PersistentVolumeClaims_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pv.PersistentVolumeClaims",
	HandlerType: (*PersistentVolumeClaimsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Describe",
			Handler:    _PersistentVolumeClaims_Describe_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _PersistentVolumeClaims_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _PersistentVolumeClaims_Unregister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("pvc.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x93, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x49, 0xb2, 0x76, 0xb7, 0x6f, 0xc0, 0xc3, 0xb8, 0xc4, 0x10, 0x56, 0x28, 0x39, 0x88,
	0xf4, 0x90, 0x68, 0x45, 0x11, 0x0f, 0x5e, 0xd6, 0x8b, 0x07, 0x61, 0x89, 0xd8, 0x93, 0x07, 0xd3,
	0xe9, 0x58, 0x02, 0xd9, 0xcc, 0x98, 0x77, 0x12, 0xd0, 0xba, 0x17, 0xf1, 0x1b, 0xf8, 0x09, 0xfc,
	0x4c, 0x7e, 0x05, 0x3f, 0x88, 0x33, 0x93, 0x49, 0xd6, 0xa6, 0x16, 0x44, 0xf1, 0xd2, 0xcc, 0x3c,
	0xf3, 0xe7, 0x37, 0xcf, 0xfb, 0x3e, 0x85, 0xa9, 0x68, 0x69, 0x22, 0x6a, 0x2e, 0x39, 0x71, 0x45,
	0x1b, 0x9d, 0x6d, 0x38, 0xdf, 0x94, 0x2c, 0xcd, 0x45, 0x91, 0xe6, 0x55, 0xc5, 0x65, 0x2e, 0x0b,
	0x5e, 0x61, 0xb7, 0x23, 0x0a, 0xb4, 0xbc, 0x96, 0x1f, 0x04, 0xc3, 0xd4, 0xfc, 0x76, 0x7a, 0x2c,
	0x81, 0x5c, 0x2c, 0xcf, 0x33, 0xb6, 0x29, 0x50, 0xb2, 0x3a, 0x63, 0xef, 0x1b, 0x86, 0x92, 0x84,
	0x70, 0x4c, 0xcb, 0x46, 0x2b, 0xa1, 0x33, 0x73, 0xee, 0x4d, 0xb3, 0x7e, 0x4a, 0x08, 0x1c, 0x55,
	0xf9, 0x25, 0x0b, 0x5d, 0x23, 0x9b, 0xb1, 0xd6, 0xb0, 0xf8, 0xc8, 0x42, 0x4f, 0x69, 0x5e, 0x66,
	0xc6, 0xe4, 0x0c, 0xa6, 0x7a, 0x0d, 0x45, 0x4e, 0x59, 0x78, 0x64, 0x36, 0x5f, 0x0b, 0xf1, 0x0a,
	0x4e, 0x15, 0xf5, 0x75, 0x55, 0xff, 0x13, 0x77, 0x87, 0xe1, 0x8d, 0x19, 0x6f, 0x8d, 0xb3, 0xe7,
	0x0c, 0x69, 0x5d, 0xac, 0xd8, 0xff, 0x20, 0x7c, 0x73, 0xe0, 0x58, 0x21, 0x5e, 0x54, 0xef, 0xf8,
	0x70, 0xda, 0xf9, 0x4d, 0x5d, 0xdc, 0x43, 0x75, 0x19, 0xdf, 0x48, 0x02, 0x98, 0xa0, 0xea, 0x5b,
	0x83, 0xb6, 0x64, 0x76, 0xa6, 0xf5, 0x96, 0x97, 0x8d, 0xba, 0xff, 0x46, 0xa7, 0x77, 0x33, 0x32,
	0x03, 0x3f, 0xa7, 0x94, 0x21, 0xbe, 0xe4, 0x6b, 0x86, 0xe1, 0x64, 0xe6, 0xa9, 0xc5, 0x5f, 0xa5,
	0xf8, 0x0d, 0xdc, 0xda, 0xa9, 0x02, 0x0a, 0x95, 0x09, 0x46, 0xee, 0x0e, 0x20, 0xfd, 0x60, 0x7f,
	0x71, 0x33, 0xe9, 0xb2, 0x91, 0xbc, 0x32, 0xea, 0x00, 0xbe, 0x03, 0x9e, 0x4a, 0x99, 0x71, 0xe0,
	0x2f, 0xfc, 0x44, 0xb4, 0x89, 0x35, 0x9c, 0x69, 0x7d, 0xf1, 0xc5, 0x83, 0xe0, 0x82, 0xd5, 0xa8,
	0x9b, 0x58, 0xc9, 0xa5, 0x79, 0xd4, 0x79, 0x99, 0x17, 0x97, 0x48, 0xb6, 0x70, 0xd2, 0x53, 0x49,
	0x60, 0x0f, 0x8e, 0x9a, 0x11, 0xdd, 0xde, 0xd3, 0xbb, 0xe7, 0xc5, 0xcf, 0x3e, 0x7f, 0xff, 0xf1,
	0xd5, 0x7d, 0x42, 0x1e, 0xab, 0x34, 0x0b, 0xa4, 0xca, 0x89, 0x89, 0xb5, 0x68, 0xd3, 0xf6, 0x7e,
	0xf2, 0x40, 0x7d, 0x69, 0xba, 0xb5, 0x7d, 0xbb, 0x4a, 0xb7, 0xba, 0x7e, 0xf6, 0x63, 0xca, 0x78,
	0x45, 0x38, 0x9c, 0xf4, 0x91, 0x1e, 0xe0, 0xa3, 0x8c, 0x47, 0xa7, 0xbd, 0xe5, 0x25, 0x2f, 0xd6,
	0x03, 0xf9, 0x91, 0x21, 0xa7, 0xd1, 0xfc, 0xcf, 0xc9, 0x4f, 0x9d, 0x39, 0xf9, 0x04, 0x70, 0x9d,
	0x66, 0x12, 0x5a, 0xe4, 0x5e, 0xc0, 0x0f, 0x40, 0xad, 0xdd, 0xf9, 0x5f, 0xda, 0x5d, 0x4d, 0xcc,
	0x7f, 0xf9, 0xe1, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x3f, 0xcf, 0xfc, 0x12, 0x04, 0x00,
	0x00,
}
