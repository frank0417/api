// Code generated by protoc-gen-go.
// source: build.proto
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

type BuildDescribeRequest struct {
	JobName string `protobuf:"bytes,1,opt,name=job_name,json=jobName" json:"job_name,omitempty"`
	Parents string `protobuf:"bytes,2,opt,name=parents" json:"parents,omitempty"`
	Number  int64  `protobuf:"varint,3,opt,name=number" json:"number,omitempty"`
}

func (m *BuildDescribeRequest) Reset()                    { *m = BuildDescribeRequest{} }
func (m *BuildDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*BuildDescribeRequest) ProtoMessage()               {}
func (*BuildDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type BuildDescribeResponse struct {
	Status            *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	JobName           string         `protobuf:"bytes,2,opt,name=job_name,json=jobName" json:"job_name,omitempty"`
	Parents           string         `protobuf:"bytes,3,opt,name=parents" json:"parents,omitempty"`
	BuildNumber       int64          `protobuf:"varint,4,opt,name=build_number,json=buildNumber" json:"build_number,omitempty"`
	BaseUrl           string         `protobuf:"bytes,5,opt,name=base_url,json=baseUrl" json:"base_url,omitempty"`
	Building          bool           `protobuf:"varint,6,opt,name=building" json:"building,omitempty"`
	BuiltOn           string         `protobuf:"bytes,7,opt,name=built_on,json=builtOn" json:"built_on,omitempty"`
	Duration          int64          `protobuf:"varint,8,opt,name=duration" json:"duration,omitempty"`
	EstimatedDuration int64          `protobuf:"varint,9,opt,name=estimated_duration,json=estimatedDuration" json:"estimated_duration,omitempty"`
	FullDisplayName   string         `protobuf:"bytes,10,opt,name=full_display_name,json=fullDisplayName" json:"full_display_name,omitempty"`
	Result            string         `protobuf:"bytes,11,opt,name=result" json:"result,omitempty"`
	ConsoleOutput     string         `protobuf:"bytes,12,opt,name=console_output,json=consoleOutput" json:"console_output,omitempty"`
}

func (m *BuildDescribeResponse) Reset()                    { *m = BuildDescribeResponse{} }
func (m *BuildDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*BuildDescribeResponse) ProtoMessage()               {}
func (*BuildDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *BuildDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type BuildListRequest struct {
	JobName string `protobuf:"bytes,1,opt,name=job_name,json=jobName" json:"job_name,omitempty"`
	Parents string `protobuf:"bytes,2,opt,name=parents" json:"parents,omitempty"`
}

func (m *BuildListRequest) Reset()                    { *m = BuildListRequest{} }
func (m *BuildListRequest) String() string            { return proto.CompactTextString(m) }
func (*BuildListRequest) ProtoMessage()               {}
func (*BuildListRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type BuildListResponse struct {
	Status  *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	JobName string         `protobuf:"bytes,2,opt,name=job_name,json=jobName" json:"job_name,omitempty"`
	Parents string         `protobuf:"bytes,3,opt,name=parents" json:"parents,omitempty"`
	Builds  []*Build       `protobuf:"bytes,4,rep,name=builds" json:"builds,omitempty"`
}

func (m *BuildListResponse) Reset()                    { *m = BuildListResponse{} }
func (m *BuildListResponse) String() string            { return proto.CompactTextString(m) }
func (*BuildListResponse) ProtoMessage()               {}
func (*BuildListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *BuildListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *BuildListResponse) GetBuilds() []*Build {
	if m != nil {
		return m.Builds
	}
	return nil
}

type Build struct {
	Number    int64  `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
	Status    string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *Build) Reset()                    { *m = Build{} }
func (m *Build) String() string            { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()               {}
func (*Build) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func init() {
	proto.RegisterType((*BuildDescribeRequest)(nil), "ci.v1beta1.BuildDescribeRequest")
	proto.RegisterType((*BuildDescribeResponse)(nil), "ci.v1beta1.BuildDescribeResponse")
	proto.RegisterType((*BuildListRequest)(nil), "ci.v1beta1.BuildListRequest")
	proto.RegisterType((*BuildListResponse)(nil), "ci.v1beta1.BuildListResponse")
	proto.RegisterType((*Build)(nil), "ci.v1beta1.Build")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Builds service

type BuildsClient interface {
	Describe(ctx context.Context, in *BuildDescribeRequest, opts ...grpc.CallOption) (*BuildDescribeResponse, error)
	List(ctx context.Context, in *BuildListRequest, opts ...grpc.CallOption) (*BuildListResponse, error)
}

type buildsClient struct {
	cc *grpc.ClientConn
}

func NewBuildsClient(cc *grpc.ClientConn) BuildsClient {
	return &buildsClient{cc}
}

func (c *buildsClient) Describe(ctx context.Context, in *BuildDescribeRequest, opts ...grpc.CallOption) (*BuildDescribeResponse, error) {
	out := new(BuildDescribeResponse)
	err := grpc.Invoke(ctx, "/ci.v1beta1.Builds/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildsClient) List(ctx context.Context, in *BuildListRequest, opts ...grpc.CallOption) (*BuildListResponse, error) {
	out := new(BuildListResponse)
	err := grpc.Invoke(ctx, "/ci.v1beta1.Builds/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Builds service

type BuildsServer interface {
	Describe(context.Context, *BuildDescribeRequest) (*BuildDescribeResponse, error)
	List(context.Context, *BuildListRequest) (*BuildListResponse, error)
}

func RegisterBuildsServer(s *grpc.Server, srv BuildsServer) {
	s.RegisterService(&_Builds_serviceDesc, srv)
}

func _Builds_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.v1beta1.Builds/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildsServer).Describe(ctx, req.(*BuildDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Builds_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.v1beta1.Builds/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildsServer).List(ctx, req.(*BuildListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Builds_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ci.v1beta1.Builds",
	HandlerType: (*BuildsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Describe",
			Handler:    _Builds_Describe_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Builds_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("build.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 570 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x54, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0x26, 0xed, 0x6e, 0x9a, 0xbe, 0xae, 0xab, 0x1d, 0x74, 0x19, 0xcb, 0x2a, 0xdd, 0x80, 0x52,
	0x05, 0x13, 0x5a, 0x2f, 0x22, 0x78, 0x59, 0x8a, 0x5e, 0x64, 0x17, 0x22, 0x7b, 0xf1, 0x12, 0x26,
	0xc9, 0x58, 0xa6, 0x24, 0x99, 0x98, 0x99, 0x11, 0x96, 0x65, 0x2f, 0xfd, 0x0b, 0x1e, 0x3c, 0x78,
	0xf0, 0xe8, 0x2f, 0xf0, 0x97, 0xf8, 0x17, 0xfc, 0x21, 0x92, 0x99, 0x69, 0xb7, 0xdd, 0xa5, 0x82,
	0x08, 0x5e, 0x02, 0xef, 0x7b, 0xdf, 0x9b, 0xef, 0x9b, 0xf7, 0xde, 0x04, 0x7a, 0x89, 0x62, 0x79,
	0x16, 0x54, 0x35, 0x97, 0x1c, 0x41, 0xca, 0x82, 0x4f, 0xe3, 0x84, 0x4a, 0x32, 0x1e, 0x1c, 0xce,
	0x38, 0x9f, 0xe5, 0x34, 0x24, 0x15, 0x0b, 0x49, 0x59, 0x72, 0x49, 0x24, 0xe3, 0xa5, 0x30, 0xcc,
	0xc1, 0x43, 0x52, 0x55, 0x22, 0xe5, 0xd9, 0xb6, 0xfc, 0x41, 0x03, 0x67, 0xf2, 0xbc, 0xa2, 0x22,
	0xd4, 0x5f, 0x83, 0xfb, 0x29, 0xdc, 0x3d, 0x6e, 0x04, 0xa7, 0x54, 0xa4, 0x35, 0x4b, 0x68, 0x44,
	0x3f, 0x2a, 0x2a, 0x24, 0xba, 0x0f, 0xde, 0x9c, 0x27, 0x71, 0x49, 0x0a, 0x8a, 0x9d, 0xa1, 0x33,
	0xea, 0x46, 0x9d, 0x39, 0x4f, 0x4e, 0x48, 0x41, 0x11, 0x86, 0x4e, 0x45, 0x6a, 0x5a, 0x4a, 0x81,
	0x5b, 0x26, 0x63, 0x43, 0x74, 0x00, 0x6e, 0xa9, 0x8a, 0x84, 0xd6, 0xb8, 0x3d, 0x74, 0x46, 0xed,
	0xc8, 0x46, 0xfe, 0xb7, 0x36, 0xdc, 0xbb, 0xa6, 0x22, 0x2a, 0x5e, 0x0a, 0x8a, 0x1e, 0x83, 0x2b,
	0x24, 0x91, 0x4a, 0x68, 0x91, 0xde, 0x64, 0x3f, 0x30, 0x1e, 0x83, 0x77, 0x1a, 0x8d, 0x6c, 0x76,
	0xc3, 0x4e, 0x6b, 0xab, 0x9d, 0xf6, 0xa6, 0x9d, 0x23, 0xd8, 0xd3, 0xcd, 0x8c, 0xad, 0xa9, 0x1d,
	0x6d, 0xca, 0x34, 0xf8, 0x44, 0x43, 0xcd, 0xb9, 0x09, 0x11, 0x34, 0x56, 0x75, 0x8e, 0x77, 0x4d,
	0x75, 0x13, 0x9f, 0xd5, 0x39, 0x1a, 0x80, 0xa7, 0x99, 0xac, 0x9c, 0x61, 0x77, 0xe8, 0x8c, 0xbc,
	0x68, 0x15, 0xeb, 0x32, 0xc5, 0x72, 0x19, 0xf3, 0x12, 0x77, 0x6c, 0x59, 0x13, 0x9f, 0x96, 0x4d,
	0x59, 0xa6, 0x6a, 0xdd, 0x7b, 0xec, 0x69, 0xc1, 0x55, 0x8c, 0x9e, 0x01, 0xa2, 0x42, 0xb2, 0x82,
	0x48, 0x9a, 0xc5, 0x2b, 0x56, 0x57, 0xb3, 0xfa, 0xab, 0xcc, 0x74, 0x49, 0x7f, 0x0a, 0xfd, 0x0f,
	0x2a, 0xcf, 0xe3, 0x8c, 0x89, 0x2a, 0x27, 0xe7, 0xe6, 0xf6, 0xa0, 0xe5, 0x6e, 0x37, 0x89, 0xa9,
	0xc1, 0x75, 0x17, 0x0e, 0xc0, 0xad, 0xa9, 0x50, 0xb9, 0xc4, 0x3d, 0x4d, 0xb0, 0x11, 0x7a, 0x04,
	0xfb, 0x29, 0x2f, 0x05, 0xcf, 0x69, 0xcc, 0x95, 0xac, 0x94, 0xc4, 0x7b, 0x3a, 0x7f, 0xcb, 0xa2,
	0xa7, 0x1a, 0xf4, 0xdf, 0xc0, 0x1d, 0x3d, 0xa0, 0xb7, 0x4c, 0xc8, 0x7f, 0x59, 0x01, 0xff, 0xab,
	0x03, 0xfd, 0xb5, 0x93, 0xfe, 0xc7, 0x98, 0x9f, 0x80, 0xab, 0x07, 0x23, 0xf0, 0xce, 0xb0, 0x3d,
	0xea, 0x4d, 0xfa, 0xc1, 0xd5, 0xab, 0x09, 0xb4, 0x97, 0xc8, 0x12, 0xfc, 0x33, 0xd8, 0xd5, 0xc0,
	0xda, 0xa6, 0x3a, 0xeb, 0x9b, 0xda, 0xe0, 0xd6, 0xa8, 0x91, 0x5f, 0x1a, 0x3b, 0x84, 0xae, 0x64,
	0x05, 0x15, 0x92, 0x14, 0x95, 0x5d, 0xee, 0x2b, 0x60, 0xf2, 0xbd, 0x05, 0xae, 0x3e, 0x57, 0xa0,
	0x2f, 0x0e, 0x78, 0xcb, 0x2d, 0x47, 0xc3, 0x1b, 0x4e, 0xae, 0x3d, 0xb3, 0xc1, 0xd1, 0x1f, 0x18,
	0xa6, 0x77, 0xfe, 0xeb, 0xc5, 0x0f, 0xdc, 0xf2, 0x9c, 0xc5, 0xcf, 0x5f, 0x9f, 0x5b, 0x2f, 0xd1,
	0x8b, 0x70, 0xe3, 0xa1, 0xa7, 0x2c, 0xb4, 0xd5, 0xe1, 0x9c, 0x27, 0x22, 0xbc, 0x58, 0xb6, 0xef,
	0x32, 0x34, 0x57, 0x0e, 0x2f, 0xcc, 0xcd, 0x2e, 0xd1, 0xc2, 0x81, 0x9d, 0x66, 0x28, 0xe8, 0xf0,
	0x86, 0xe6, 0xda, 0xd4, 0x07, 0x0f, 0xb6, 0x64, 0xad, 0x9b, 0x57, 0x6b, 0x6e, 0xc6, 0x28, 0xfc,
	0x4b, 0x37, 0xc7, 0xdd, 0xf7, 0x1d, 0xcb, 0x48, 0x5c, 0xfd, 0x03, 0x7a, 0xfe, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0xe9, 0x66, 0xb3, 0xa3, 0xf1, 0x04, 0x00, 0x00,
}
