// Code generated by protoc-gen-go.
// source: volume.proto
// DO NOT EDIT!

package glusterfs

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

type VolumeCreateRequest struct {
	KubeCluster      string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	GlusterfsCluster string `protobuf:"bytes,2,opt,name=glusterfs_cluster,json=glusterfsCluster" json:"glusterfs_cluster,omitempty"`
	Volume           string `protobuf:"bytes,3,opt,name=volume" json:"volume,omitempty"`
	EnableBacula     int32  `protobuf:"varint,4,opt,name=enable_bacula,json=enableBacula" json:"enable_bacula,omitempty"`
}

func (m *VolumeCreateRequest) Reset()                    { *m = VolumeCreateRequest{} }
func (m *VolumeCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*VolumeCreateRequest) ProtoMessage()               {}
func (*VolumeCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type VolumeDeleteRequest struct {
	KubeCluster      string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	GlusterfsCluster string `protobuf:"bytes,2,opt,name=glusterfs_cluster,json=glusterfsCluster" json:"glusterfs_cluster,omitempty"`
	Volume           string `protobuf:"bytes,3,opt,name=volume" json:"volume,omitempty"`
}

func (m *VolumeDeleteRequest) Reset()                    { *m = VolumeDeleteRequest{} }
func (m *VolumeDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*VolumeDeleteRequest) ProtoMessage()               {}
func (*VolumeDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type VolumeListRequest struct {
	KubeCluster      string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	GlusterfsCluster string `protobuf:"bytes,2,opt,name=glusterfs_cluster,json=glusterfsCluster" json:"glusterfs_cluster,omitempty"`
}

func (m *VolumeListRequest) Reset()                    { *m = VolumeListRequest{} }
func (m *VolumeListRequest) String() string            { return proto.CompactTextString(m) }
func (*VolumeListRequest) ProtoMessage()               {}
func (*VolumeListRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type VolumeListResponse struct {
	Status  *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Volumes []*Volume      `protobuf:"bytes,2,rep,name=volumes" json:"volumes,omitempty"`
}

func (m *VolumeListResponse) Reset()                    { *m = VolumeListResponse{} }
func (m *VolumeListResponse) String() string            { return proto.CompactTextString(m) }
func (*VolumeListResponse) ProtoMessage()               {}
func (*VolumeListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *VolumeListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VolumeListResponse) GetVolumes() []*Volume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

type Volume struct {
	Path     string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Bricks   string `protobuf:"bytes,2,opt,name=bricks" json:"bricks,omitempty"`
	Endpoint string `protobuf:"bytes,3,opt,name=endpoint" json:"endpoint,omitempty"`
	Replica  int64  `protobuf:"varint,4,opt,name=replica" json:"replica,omitempty"`
}

func (m *Volume) Reset()                    { *m = Volume{} }
func (m *Volume) String() string            { return proto.CompactTextString(m) }
func (*Volume) ProtoMessage()               {}
func (*Volume) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func init() {
	proto.RegisterType((*VolumeCreateRequest)(nil), "glusterfs.VolumeCreateRequest")
	proto.RegisterType((*VolumeDeleteRequest)(nil), "glusterfs.VolumeDeleteRequest")
	proto.RegisterType((*VolumeListRequest)(nil), "glusterfs.VolumeListRequest")
	proto.RegisterType((*VolumeListResponse)(nil), "glusterfs.VolumeListResponse")
	proto.RegisterType((*Volume)(nil), "glusterfs.Volume")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Volumes service

type VolumesClient interface {
	List(ctx context.Context, in *VolumeListRequest, opts ...grpc.CallOption) (*VolumeListResponse, error)
	Create(ctx context.Context, in *VolumeCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *VolumeDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type volumesClient struct {
	cc *grpc.ClientConn
}

func NewVolumesClient(cc *grpc.ClientConn) VolumesClient {
	return &volumesClient{cc}
}

func (c *volumesClient) List(ctx context.Context, in *VolumeListRequest, opts ...grpc.CallOption) (*VolumeListResponse, error) {
	out := new(VolumeListResponse)
	err := grpc.Invoke(ctx, "/glusterfs.Volumes/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumesClient) Create(ctx context.Context, in *VolumeCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/glusterfs.Volumes/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumesClient) Delete(ctx context.Context, in *VolumeDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/glusterfs.Volumes/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Volumes service

type VolumesServer interface {
	List(context.Context, *VolumeListRequest) (*VolumeListResponse, error)
	Create(context.Context, *VolumeCreateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *VolumeDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterVolumesServer(s *grpc.Server, srv VolumesServer) {
	s.RegisterService(&_Volumes_serviceDesc, srv)
}

func _Volumes_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glusterfs.Volumes/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumesServer).List(ctx, req.(*VolumeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Volumes_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glusterfs.Volumes/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumesServer).Create(ctx, req.(*VolumeCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Volumes_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glusterfs.Volumes/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumesServer).Delete(ctx, req.(*VolumeDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Volumes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "glusterfs.Volumes",
	HandlerType: (*VolumesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Volumes_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Volumes_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Volumes_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("volume.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x94, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x49, 0x5b, 0x53, 0xf7, 0xb5, 0x8a, 0x7d, 0xca, 0x12, 0xc2, 0x2a, 0x35, 0x82, 0x94,
	0x5d, 0xc8, 0x68, 0xf7, 0x22, 0x1e, 0x3c, 0xb8, 0x82, 0x08, 0x82, 0x12, 0x61, 0xf1, 0xb6, 0x4c,
	0xd2, 0xb1, 0x8e, 0x1b, 0x33, 0x63, 0x66, 0xb2, 0x20, 0xa5, 0x1e, 0xfc, 0x0a, 0x7a, 0xf0, 0xec,
	0xd7, 0xf0, 0x63, 0xf8, 0x15, 0xfc, 0x20, 0xd2, 0x79, 0x49, 0xe8, 0x52, 0x7b, 0x11, 0xdd, 0x4b,
	0xe9, 0xfb, 0xcf, 0xcb, 0xfb, 0xff, 0xe6, 0xbd, 0x97, 0xc0, 0xf0, 0x4c, 0xe5, 0xd5, 0x7b, 0x11,
	0xeb, 0x52, 0x59, 0x85, 0x3b, 0xf3, 0xbc, 0x32, 0x56, 0x94, 0x6f, 0x4c, 0xb8, 0x37, 0x57, 0x6a,
	0x9e, 0x0b, 0xc6, 0xb5, 0x64, 0xbc, 0x28, 0x94, 0xe5, 0x56, 0xaa, 0xc2, 0x50, 0x62, 0xb8, 0xbb,
	0x92, 0x67, 0xf6, 0xa3, 0x16, 0x86, 0xb9, 0x5f, 0xd2, 0xa3, 0xef, 0x1e, 0x5c, 0x3f, 0x76, 0x15,
	0x8f, 0x4a, 0xc1, 0xad, 0x48, 0xc4, 0x87, 0x4a, 0x18, 0x8b, 0xb7, 0x61, 0x78, 0x5a, 0xa5, 0xe2,
	0x24, 0xa3, 0xfa, 0x81, 0x37, 0xf6, 0x26, 0x3b, 0xc9, 0x60, 0xa5, 0x1d, 0x91, 0x84, 0x07, 0x30,
	0x6a, 0xdd, 0xdb, 0xbc, 0x8e, 0xcb, 0xbb, 0xd6, 0x1e, 0x34, 0xc9, 0xbb, 0xe0, 0x13, 0x78, 0xd0,
	0x75, 0x19, 0x75, 0x84, 0x77, 0xe0, 0x8a, 0x28, 0x78, 0x9a, 0x8b, 0x93, 0x94, 0x67, 0x55, 0xce,
	0x83, 0xde, 0xd8, 0x9b, 0x5c, 0x4a, 0x86, 0x24, 0x3e, 0x76, 0x5a, 0xb4, 0x6c, 0x18, 0x9f, 0x88,
	0x5c, 0x5c, 0x38, 0x63, 0x94, 0xc1, 0x88, 0xec, 0x9f, 0x4b, 0x63, 0xff, 0x93, 0x79, 0x24, 0x01,
	0xd7, 0x4d, 0x8c, 0x56, 0x85, 0x11, 0x78, 0x17, 0x7c, 0x63, 0xb9, 0xad, 0x8c, 0xab, 0x3f, 0x98,
	0x5e, 0x8d, 0x69, 0x86, 0xf1, 0x2b, 0xa7, 0x26, 0xf5, 0x29, 0x1e, 0x40, 0x9f, 0x60, 0x4d, 0xd0,
	0x19, 0x77, 0x27, 0x83, 0xe9, 0x28, 0x6e, 0x1d, 0x62, 0xaa, 0x9b, 0x34, 0x19, 0xd1, 0x3b, 0xf0,
	0x49, 0x42, 0x84, 0x9e, 0xe6, 0xf6, 0x6d, 0x0d, 0xef, 0xfe, 0xaf, 0xba, 0x90, 0x96, 0x32, 0x3b,
	0x35, 0x35, 0x6a, 0x1d, 0x61, 0x08, 0x97, 0x45, 0x31, 0xd3, 0x4a, 0x16, 0xb6, 0xee, 0x4f, 0x1b,
	0x63, 0x00, 0xfd, 0x52, 0xe8, 0x5c, 0x66, 0x34, 0xbf, 0x6e, 0xd2, 0x84, 0xd3, 0x1f, 0x5d, 0xe8,
	0x93, 0x99, 0xc1, 0xaf, 0x1e, 0xf4, 0x56, 0xb7, 0xc3, 0xbd, 0x0d, 0xb8, 0xb5, 0xce, 0x86, 0x37,
	0xb7, 0x9c, 0x52, 0x4b, 0xa2, 0x17, 0x9f, 0x7f, 0xfe, 0xfa, 0xd2, 0x79, 0x86, 0x4f, 0x19, 0xd7,
	0xda, 0x64, 0x6a, 0x46, 0x2b, 0xdf, 0x3e, 0xc3, 0xce, 0xee, 0xc5, 0xf7, 0x59, 0x7d, 0x57, 0xb6,
	0x58, 0x1f, 0xd2, 0x92, 0x2d, 0x36, 0x06, 0xb2, 0xc4, 0x4f, 0xe0, 0xd3, 0xee, 0xe3, 0xad, 0x0d,
	0xe7, 0x73, 0x2f, 0x45, 0x78, 0xa3, 0xe9, 0xfe, 0xb1, 0x92, 0xb3, 0x16, 0xe8, 0x91, 0x03, 0x7a,
	0x10, 0x1e, 0xfe, 0x05, 0xd0, 0x43, 0x6f, 0x1f, 0xbf, 0x79, 0xe0, 0xd3, 0x62, 0xff, 0x01, 0xe0,
	0xdc, 0xc6, 0x6f, 0x01, 0x78, 0xed, 0x00, 0x92, 0xfd, 0x97, 0xff, 0xa8, 0x23, 0x6c, 0x41, 0xb9,
	0xcb, 0xd4, 0x77, 0x1f, 0x89, 0xc3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x48, 0x07, 0x2f, 0xef,
	0x75, 0x04, 0x00, 0x00,
}
