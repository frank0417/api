// Code generated by protoc-gen-go. DO NOT EDIT.
// source: health.proto

/*
Package health is a generated protocol buffer package.

It is generated from these files:
	health.proto

It has these top-level messages:
	StatusResponse
	URLBase
	NetConfig
	Metadata
*/
package health

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "appscode.com/api/dtypes"
import appscode_version "appscode.com/api/version"

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

type StatusResponse struct {
	Version  *appscode_version.Version `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Metadata *Metadata                 `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StatusResponse) GetVersion() *appscode_version.Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *StatusResponse) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type URLBase struct {
	Scheme   string `protobuf:"bytes,1,opt,name=scheme" json:"scheme,omitempty"`
	BaseAddr string `protobuf:"bytes,2,opt,name=base_addr,json=baseAddr" json:"base_addr,omitempty"`
}

func (m *URLBase) Reset()                    { *m = URLBase{} }
func (m *URLBase) String() string            { return proto.CompactTextString(m) }
func (*URLBase) ProtoMessage()               {}
func (*URLBase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *URLBase) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *URLBase) GetBaseAddr() string {
	if m != nil {
		return m.BaseAddr
	}
	return ""
}

type NetConfig struct {
	TeamAddr         string   `protobuf:"bytes,1,opt,name=team_addr,json=teamAddr" json:"team_addr,omitempty"`
	PublicUrls       *URLBase `protobuf:"bytes,2,opt,name=public_urls,json=publicUrls" json:"public_urls,omitempty"`
	TeamUrls         *URLBase `protobuf:"bytes,3,opt,name=team_urls,json=teamUrls" json:"team_urls,omitempty"`
	ClusterUrls      *URLBase `protobuf:"bytes,4,opt,name=cluster_urls,json=clusterUrls" json:"cluster_urls,omitempty"`
	InClusterUrls    *URLBase `protobuf:"bytes,5,opt,name=in_cluster_urls,json=inClusterUrls" json:"in_cluster_urls,omitempty"`
	URLShortenerUrls *URLBase `protobuf:"bytes,6,opt,name=URL_shortener_urls,json=URLShortenerUrls" json:"URL_shortener_urls,omitempty"`
	FileUrls         *URLBase `protobuf:"bytes,7,opt,name=file_urls,json=fileUrls" json:"file_urls,omitempty"`
}

func (m *NetConfig) Reset()                    { *m = NetConfig{} }
func (m *NetConfig) String() string            { return proto.CompactTextString(m) }
func (*NetConfig) ProtoMessage()               {}
func (*NetConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NetConfig) GetTeamAddr() string {
	if m != nil {
		return m.TeamAddr
	}
	return ""
}

func (m *NetConfig) GetPublicUrls() *URLBase {
	if m != nil {
		return m.PublicUrls
	}
	return nil
}

func (m *NetConfig) GetTeamUrls() *URLBase {
	if m != nil {
		return m.TeamUrls
	}
	return nil
}

func (m *NetConfig) GetClusterUrls() *URLBase {
	if m != nil {
		return m.ClusterUrls
	}
	return nil
}

func (m *NetConfig) GetInClusterUrls() *URLBase {
	if m != nil {
		return m.InClusterUrls
	}
	return nil
}

func (m *NetConfig) GetURLShortenerUrls() *URLBase {
	if m != nil {
		return m.URLShortenerUrls
	}
	return nil
}

func (m *NetConfig) GetFileUrls() *URLBase {
	if m != nil {
		return m.FileUrls
	}
	return nil
}

type Metadata struct {
	Env       string     `protobuf:"bytes,1,opt,name=env" json:"env,omitempty"`
	TeamId    string     `protobuf:"bytes,2,opt,name=team_id,json=teamId" json:"team_id,omitempty"`
	NetConfig *NetConfig `protobuf:"bytes,3,opt,name=net_config,json=netConfig" json:"net_config,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Metadata) GetEnv() string {
	if m != nil {
		return m.Env
	}
	return ""
}

func (m *Metadata) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *Metadata) GetNetConfig() *NetConfig {
	if m != nil {
		return m.NetConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*StatusResponse)(nil), "appscode.health.StatusResponse")
	proto.RegisterType((*URLBase)(nil), "appscode.health.URLBase")
	proto.RegisterType((*NetConfig)(nil), "appscode.health.NetConfig")
	proto.RegisterType((*Metadata)(nil), "appscode.health.Metadata")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Health service

type HealthClient interface {
	Status(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Status(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/appscode.health.Health/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Health service

type HealthServer interface {
	Status(context.Context, *appscode_dtypes.VoidRequest) (*StatusResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.health.Health/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Status(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.health.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Health_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "health.proto",
}

func init() { proto.RegisterFile("health.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xd5, 0x16, 0xd2, 0xf6, 0x75, 0xb0, 0xc9, 0x07, 0x56, 0x4a, 0x05, 0x28, 0x48, 0x68,
	0xa7, 0x06, 0x6d, 0xda, 0x61, 0x42, 0x42, 0xd0, 0x49, 0x08, 0xa4, 0x82, 0x26, 0x4f, 0xdd, 0x81,
	0x4b, 0xe4, 0xc6, 0x6f, 0xad, 0x51, 0x62, 0x67, 0xb1, 0x33, 0x09, 0xc1, 0x69, 0x5f, 0x81, 0x0b,
	0x5f, 0x84, 0x4f, 0xc2, 0x57, 0xe0, 0x83, 0xa0, 0xd8, 0x4e, 0x47, 0xa9, 0x50, 0x2e, 0xb1, 0xfd,
	0xfc, 0xfb, 0xff, 0xfd, 0xec, 0x97, 0x07, 0x3b, 0x2b, 0x64, 0xa9, 0x59, 0x4d, 0xf2, 0x42, 0x19,
	0x45, 0x76, 0x59, 0x9e, 0xeb, 0x44, 0x71, 0x9c, 0xb8, 0xf0, 0x68, 0xbc, 0x54, 0x6a, 0x99, 0x62,
	0xc4, 0x72, 0x11, 0x31, 0x29, 0x95, 0x61, 0x46, 0x28, 0xa9, 0x1d, 0x3e, 0x7a, 0x5c, 0xe3, 0xff,
	0xd9, 0x7f, 0xb6, 0xb6, 0x4b, 0x54, 0x66, 0x19, 0x6e, 0xbe, 0xe4, 0xa8, 0x23, 0xfb, 0xf5, 0xd0,
	0xf3, 0x2d, 0xe8, 0x1a, 0x0b, 0x2d, 0x94, 0xac, 0x47, 0xc7, 0x85, 0xdf, 0xe0, 0xfe, 0xb9, 0x61,
	0xa6, 0xd4, 0x14, 0x75, 0xae, 0xa4, 0x46, 0x72, 0x04, 0x5d, 0x8f, 0x0c, 0x5b, 0x4f, 0x5b, 0x07,
	0x83, 0xc3, 0x87, 0x93, 0xb5, 0x57, 0xad, 0xbd, 0x70, 0x23, 0xad, 0x49, 0x72, 0x0c, 0xbd, 0x0c,
	0x0d, 0xe3, 0xcc, 0xb0, 0x61, 0xfb, 0x5f, 0x95, 0x7f, 0x8c, 0x0f, 0x1e, 0xa0, 0x6b, 0x34, 0x7c,
	0x05, 0xdd, 0x39, 0x9d, 0x4d, 0x99, 0x46, 0xf2, 0x00, 0x02, 0x9d, 0xac, 0x30, 0x43, 0x7b, 0x6a,
	0x9f, 0xfa, 0x15, 0x79, 0x04, 0xfd, 0x05, 0xd3, 0x18, 0x33, 0xce, 0x0b, 0x6b, 0xdd, 0xa7, 0xbd,
	0x2a, 0xf0, 0x86, 0xf3, 0x22, 0xfc, 0xd1, 0x81, 0xfe, 0x47, 0x34, 0xa7, 0x4a, 0x5e, 0x8a, 0x65,
	0x85, 0x1a, 0x64, 0x99, 0x43, 0x9d, 0x4b, 0xaf, 0x0a, 0x54, 0x28, 0x39, 0x81, 0x41, 0x5e, 0x2e,
	0x52, 0x91, 0xc4, 0x65, 0x91, 0x6a, 0x9f, 0xe4, 0x70, 0x2b, 0x49, 0x9f, 0x0e, 0x05, 0x07, 0xcf,
	0x8b, 0x54, 0x93, 0x63, 0xef, 0x6b, 0x85, 0x9d, 0x06, 0xa1, 0x3d, 0xd1, 0xca, 0x5e, 0xc2, 0x4e,
	0x92, 0x96, 0xda, 0x60, 0xe1, 0x94, 0x77, 0x1a, 0x94, 0x03, 0x4f, 0x5b, 0xf1, 0x6b, 0xd8, 0x15,
	0x32, 0xde, 0xd0, 0xdf, 0x6d, 0xd0, 0xdf, 0x13, 0xf2, 0xf4, 0x2f, 0x87, 0xb7, 0x40, 0xe6, 0x74,
	0x16, 0xeb, 0x95, 0x2a, 0x0c, 0xca, 0xda, 0x24, 0x68, 0x30, 0xd9, 0x9b, 0xd3, 0xd9, 0x79, 0x2d,
	0xa9, 0x6f, 0x7f, 0x29, 0x52, 0x74, 0xf2, 0x6e, 0xd3, 0xed, 0x2b, 0xb4, 0x92, 0x85, 0x39, 0xf4,
	0xea, 0x82, 0x93, 0x3d, 0xe8, 0xa0, 0xbc, 0xf6, 0x25, 0xa9, 0xa6, 0x64, 0x1f, 0xba, 0xf6, 0x49,
	0x05, 0xf7, 0x35, 0x0d, 0xaa, 0xe5, 0x7b, 0x4e, 0x4e, 0x00, 0x24, 0x9a, 0x38, 0xb1, 0x15, 0xf5,
	0x8f, 0x3d, 0xda, 0x3a, 0x6e, 0x5d, 0x73, 0xda, 0x97, 0xf5, 0xf4, 0xf0, 0x2b, 0x04, 0xef, 0xec,
	0x36, 0xb9, 0x82, 0xc0, 0xfd, 0xd4, 0x64, 0x7c, 0x2b, 0x75, 0x4d, 0x32, 0xb9, 0x50, 0x82, 0x53,
	0xbc, 0x2a, 0x51, 0x9b, 0xd1, 0x93, 0x2d, 0xe3, 0xcd, 0x5e, 0x08, 0x0f, 0x6e, 0x7e, 0x0e, 0xdb,
	0xbd, 0xd6, 0xcd, 0xaf, 0xdf, 0xdf, 0xdb, 0x63, 0x32, 0x8a, 0xe2, 0x8d, 0xd6, 0x74, 0x9a, 0xe8,
	0xb3, 0x56, 0x72, 0xfa, 0x02, 0xf6, 0x13, 0x95, 0xdd, 0xfa, 0xb1, 0x5c, 0x78, 0xcf, 0xe9, 0xc0,
	0x65, 0x75, 0x56, 0xf5, 0xdb, 0x59, 0xeb, 0x53, 0xe0, 0xc2, 0x8b, 0xc0, 0x36, 0xe0, 0xd1, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x06, 0xdb, 0x43, 0xff, 0x2c, 0x04, 0x00, 0x00,
}
