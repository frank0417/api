// Code generated by protoc-gen-go.
// source: cloudcredential.proto
// DO NOT EDIT!

/*
Package credential is a generated protocol buffer package.

It is generated from these files:
	cloudcredential.proto

It has these top-level messages:
	CloudCredentialCreateRequest
	CloudCredentialUpdateRequest
	CloudCredentialDeleteRequest
	CloudCredentialListRequest
	CloudCredentialListResponse
	Credential
*/
package credential

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

type CloudCredentialCreateRequest struct {
	Name     string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Provider string            `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	Data     map[string]string `protobuf:"bytes,3,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CloudCredentialCreateRequest) Reset()                    { *m = CloudCredentialCreateRequest{} }
func (m *CloudCredentialCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CloudCredentialCreateRequest) ProtoMessage()               {}
func (*CloudCredentialCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CloudCredentialCreateRequest) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type CloudCredentialUpdateRequest struct {
	Name     string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Provider string            `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	Data     map[string]string `protobuf:"bytes,3,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CloudCredentialUpdateRequest) Reset()                    { *m = CloudCredentialUpdateRequest{} }
func (m *CloudCredentialUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*CloudCredentialUpdateRequest) ProtoMessage()               {}
func (*CloudCredentialUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CloudCredentialUpdateRequest) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type CloudCredentialDeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CloudCredentialDeleteRequest) Reset()                    { *m = CloudCredentialDeleteRequest{} }
func (m *CloudCredentialDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*CloudCredentialDeleteRequest) ProtoMessage()               {}
func (*CloudCredentialDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type CloudCredentialListRequest struct {
}

func (m *CloudCredentialListRequest) Reset()                    { *m = CloudCredentialListRequest{} }
func (m *CloudCredentialListRequest) String() string            { return proto.CompactTextString(m) }
func (*CloudCredentialListRequest) ProtoMessage()               {}
func (*CloudCredentialListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type CloudCredentialListResponse struct {
	Status      *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Credentials []*Credential  `protobuf:"bytes,2,rep,name=credentials" json:"credentials,omitempty"`
}

func (m *CloudCredentialListResponse) Reset()                    { *m = CloudCredentialListResponse{} }
func (m *CloudCredentialListResponse) String() string            { return proto.CompactTextString(m) }
func (*CloudCredentialListResponse) ProtoMessage()               {}
func (*CloudCredentialListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CloudCredentialListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CloudCredentialListResponse) GetCredentials() []*Credential {
	if m != nil {
		return m.Credentials
	}
	return nil
}

type Credential struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Provider    string `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	Information string `protobuf:"bytes,3,opt,name=information" json:"information,omitempty"`
}

func (m *Credential) Reset()                    { *m = Credential{} }
func (m *Credential) String() string            { return proto.CompactTextString(m) }
func (*Credential) ProtoMessage()               {}
func (*Credential) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*CloudCredentialCreateRequest)(nil), "credential.CloudCredentialCreateRequest")
	proto.RegisterType((*CloudCredentialUpdateRequest)(nil), "credential.CloudCredentialUpdateRequest")
	proto.RegisterType((*CloudCredentialDeleteRequest)(nil), "credential.CloudCredentialDeleteRequest")
	proto.RegisterType((*CloudCredentialListRequest)(nil), "credential.CloudCredentialListRequest")
	proto.RegisterType((*CloudCredentialListResponse)(nil), "credential.CloudCredentialListResponse")
	proto.RegisterType((*Credential)(nil), "credential.Credential")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for CloudCredential service

type CloudCredentialClient interface {
	List(ctx context.Context, in *CloudCredentialListRequest, opts ...grpc.CallOption) (*CloudCredentialListResponse, error)
	Create(ctx context.Context, in *CloudCredentialCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Update(ctx context.Context, in *CloudCredentialUpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *CloudCredentialDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type cloudCredentialClient struct {
	cc *grpc.ClientConn
}

func NewCloudCredentialClient(cc *grpc.ClientConn) CloudCredentialClient {
	return &cloudCredentialClient{cc}
}

func (c *cloudCredentialClient) List(ctx context.Context, in *CloudCredentialListRequest, opts ...grpc.CallOption) (*CloudCredentialListResponse, error) {
	out := new(CloudCredentialListResponse)
	err := grpc.Invoke(ctx, "/credential.CloudCredential/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudCredentialClient) Create(ctx context.Context, in *CloudCredentialCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/credential.CloudCredential/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudCredentialClient) Update(ctx context.Context, in *CloudCredentialUpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/credential.CloudCredential/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudCredentialClient) Delete(ctx context.Context, in *CloudCredentialDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/credential.CloudCredential/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CloudCredential service

type CloudCredentialServer interface {
	List(context.Context, *CloudCredentialListRequest) (*CloudCredentialListResponse, error)
	Create(context.Context, *CloudCredentialCreateRequest) (*dtypes.VoidResponse, error)
	Update(context.Context, *CloudCredentialUpdateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *CloudCredentialDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterCloudCredentialServer(s *grpc.Server, srv CloudCredentialServer) {
	s.RegisterService(&_CloudCredential_serviceDesc, srv)
}

func _CloudCredential_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudCredentialListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCredentialServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.CloudCredential/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCredentialServer).List(ctx, req.(*CloudCredentialListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudCredential_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudCredentialCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCredentialServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.CloudCredential/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCredentialServer).Create(ctx, req.(*CloudCredentialCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudCredential_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudCredentialUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCredentialServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.CloudCredential/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCredentialServer).Update(ctx, req.(*CloudCredentialUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudCredential_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudCredentialDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCredentialServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.CloudCredential/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCredentialServer).Delete(ctx, req.(*CloudCredentialDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CloudCredential_serviceDesc = grpc.ServiceDesc{
	ServiceName: "credential.CloudCredential",
	HandlerType: (*CloudCredentialServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _CloudCredential_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CloudCredential_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CloudCredential_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CloudCredential_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("cloudcredential.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x94, 0x41, 0x8b, 0xd4, 0x30,
	0x18, 0x86, 0xe9, 0xcc, 0x38, 0xb8, 0x5f, 0x41, 0x25, 0xac, 0x4b, 0x89, 0x73, 0x18, 0x8a, 0x8c,
	0xe3, 0xaa, 0x8d, 0x76, 0x0f, 0x2e, 0x7b, 0x9d, 0xd5, 0x93, 0xa7, 0x8a, 0x1e, 0x85, 0x38, 0x89,
	0x4b, 0xb1, 0x9b, 0xc4, 0x26, 0x1d, 0x28, 0xb2, 0x08, 0x82, 0x57, 0x2f, 0xfe, 0x30, 0x0f, 0xfe,
	0x05, 0x7f, 0x88, 0x34, 0xa9, 0x9d, 0x76, 0x1d, 0x6b, 0x05, 0xf1, 0x32, 0xa4, 0xef, 0x97, 0xe4,
	0xc9, 0xfb, 0xe5, 0xcd, 0xc0, 0xcd, 0x75, 0x26, 0x0b, 0xb6, 0xce, 0x39, 0xe3, 0xc2, 0xa4, 0x34,
	0x8b, 0x54, 0x2e, 0x8d, 0x44, 0xb0, 0x55, 0xf0, 0xec, 0x4c, 0xca, 0xb3, 0x8c, 0x13, 0xaa, 0x52,
	0x42, 0x85, 0x90, 0x86, 0x9a, 0x54, 0x0a, 0xed, 0x66, 0xe2, 0x83, 0x4a, 0x66, 0xa6, 0x54, 0x5c,
	0x13, 0xfb, 0xeb, 0xf4, 0xf0, 0xab, 0x07, 0xb3, 0x55, 0xb5, 0xf7, 0xaa, 0xd9, 0x69, 0x95, 0x73,
	0x6a, 0x78, 0xc2, 0xdf, 0x15, 0x5c, 0x1b, 0x84, 0x60, 0x22, 0xe8, 0x39, 0x0f, 0xbc, 0xb9, 0xb7,
	0xdc, 0x4b, 0xec, 0x18, 0x61, 0xb8, 0xaa, 0x72, 0xb9, 0x49, 0x19, 0xcf, 0x83, 0x91, 0xd5, 0x9b,
	0x6f, 0xf4, 0x14, 0x26, 0x8c, 0x1a, 0x1a, 0x8c, 0xe7, 0xe3, 0xa5, 0x1f, 0xc7, 0x51, 0xeb, 0xcc,
	0x7d, 0x9c, 0xe8, 0x94, 0x1a, 0xfa, 0x44, 0x98, 0xbc, 0x4c, 0xec, 0x7a, 0xfc, 0x18, 0xf6, 0x1a,
	0x09, 0xdd, 0x80, 0xf1, 0x5b, 0x5e, 0xd6, 0x67, 0xa8, 0x86, 0x68, 0x1f, 0xae, 0x6c, 0x68, 0x56,
	0xf0, 0x9a, 0xef, 0x3e, 0x4e, 0x46, 0xc7, 0xde, 0x2e, 0x47, 0x2f, 0x14, 0xfb, 0x2f, 0x8e, 0x3a,
	0x9c, 0x7f, 0xe7, 0x28, 0xfe, 0xc5, 0xd0, 0x29, 0xcf, 0x78, 0xaf, 0xa1, 0x70, 0x06, 0xf8, 0xd2,
	0x9a, 0x67, 0xa9, 0x36, 0xf5, 0x8a, 0xf0, 0x03, 0xdc, 0xda, 0x59, 0xd5, 0x4a, 0x0a, 0xcd, 0xd1,
	0x02, 0xa6, 0xda, 0x50, 0x53, 0x68, 0xbb, 0xa5, 0x1f, 0x5f, 0x8b, 0x5c, 0x72, 0xa2, 0xe7, 0x56,
	0x4d, 0xea, 0x2a, 0x3a, 0x06, 0x7f, 0xdb, 0x0c, 0x1d, 0x8c, 0x6c, 0x83, 0x0e, 0x3a, 0x0d, 0x6a,
	0x86, 0x49, 0x7b, 0x6a, 0xf8, 0x0a, 0x60, 0x5b, 0xfa, 0xeb, 0x1b, 0x99, 0x83, 0x9f, 0x8a, 0x37,
	0x32, 0x3f, 0xb7, 0x11, 0x0f, 0xc6, 0xb6, 0xdc, 0x96, 0xe2, 0xcf, 0x13, 0xb8, 0x7e, 0xc9, 0x21,
	0xfa, 0xe4, 0xc1, 0xa4, 0xb2, 0x89, 0x16, 0x3d, 0x57, 0xd8, 0xea, 0x12, 0xbe, 0xf3, 0xc7, 0x79,
	0xae, 0x5f, 0xe1, 0xfd, 0x8f, 0xdf, 0xbe, 0x7f, 0x19, 0x2d, 0xd0, 0x6d, 0x42, 0x95, 0xd2, 0x6b,
	0xc9, 0xdc, 0x2b, 0x6c, 0x19, 0x26, 0x9b, 0x87, 0xd1, 0x23, 0x62, 0x1f, 0x31, 0x2a, 0x61, 0xea,
	0xa2, 0x8f, 0x96, 0x43, 0x5f, 0x07, 0xde, 0xff, 0x79, 0x03, 0x2f, 0x65, 0xca, 0x1a, 0x2e, 0xb1,
	0xdc, 0xbb, 0x78, 0x10, 0xf7, 0xc4, 0x3b, 0xac, 0xd0, 0x2e, 0xa3, 0xbd, 0xe8, 0x4e, 0x8c, 0xfb,
	0xd1, 0xe1, 0x60, 0xf4, 0x05, 0x4c, 0x5d, 0x6a, 0x7b, 0xd1, 0x9d, 0x60, 0xff, 0x06, 0x7d, 0x64,
	0xd1, 0x0f, 0x0e, 0xef, 0x0d, 0x41, 0x93, 0xf7, 0x55, 0x9a, 0x2e, 0x5e, 0x4f, 0xed, 0xdf, 0xdd,
	0xd1, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x02, 0x25, 0xd5, 0x49, 0x05, 0x00, 0x00,
}
