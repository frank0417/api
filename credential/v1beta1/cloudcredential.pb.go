// Code generated by protoc-gen-go.
// source: cloudcredential.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

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
	proto.RegisterType((*CloudCredentialCreateRequest)(nil), "credential.v1beta1.CloudCredentialCreateRequest")
	proto.RegisterType((*CloudCredentialUpdateRequest)(nil), "credential.v1beta1.CloudCredentialUpdateRequest")
	proto.RegisterType((*CloudCredentialDeleteRequest)(nil), "credential.v1beta1.CloudCredentialDeleteRequest")
	proto.RegisterType((*CloudCredentialListRequest)(nil), "credential.v1beta1.CloudCredentialListRequest")
	proto.RegisterType((*CloudCredentialListResponse)(nil), "credential.v1beta1.CloudCredentialListResponse")
	proto.RegisterType((*Credential)(nil), "credential.v1beta1.Credential")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for CloudCredentials service

type CloudCredentialsClient interface {
	List(ctx context.Context, in *CloudCredentialListRequest, opts ...grpc.CallOption) (*CloudCredentialListResponse, error)
	Create(ctx context.Context, in *CloudCredentialCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Update(ctx context.Context, in *CloudCredentialUpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *CloudCredentialDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type cloudCredentialsClient struct {
	cc *grpc.ClientConn
}

func NewCloudCredentialsClient(cc *grpc.ClientConn) CloudCredentialsClient {
	return &cloudCredentialsClient{cc}
}

func (c *cloudCredentialsClient) List(ctx context.Context, in *CloudCredentialListRequest, opts ...grpc.CallOption) (*CloudCredentialListResponse, error) {
	out := new(CloudCredentialListResponse)
	err := grpc.Invoke(ctx, "/credential.v1beta1.CloudCredentials/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudCredentialsClient) Create(ctx context.Context, in *CloudCredentialCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/credential.v1beta1.CloudCredentials/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudCredentialsClient) Update(ctx context.Context, in *CloudCredentialUpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/credential.v1beta1.CloudCredentials/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudCredentialsClient) Delete(ctx context.Context, in *CloudCredentialDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/credential.v1beta1.CloudCredentials/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CloudCredentials service

type CloudCredentialsServer interface {
	List(context.Context, *CloudCredentialListRequest) (*CloudCredentialListResponse, error)
	Create(context.Context, *CloudCredentialCreateRequest) (*dtypes.VoidResponse, error)
	Update(context.Context, *CloudCredentialUpdateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *CloudCredentialDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterCloudCredentialsServer(s *grpc.Server, srv CloudCredentialsServer) {
	s.RegisterService(&_CloudCredentials_serviceDesc, srv)
}

func _CloudCredentials_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudCredentialListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCredentialsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.v1beta1.CloudCredentials/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCredentialsServer).List(ctx, req.(*CloudCredentialListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudCredentials_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudCredentialCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCredentialsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.v1beta1.CloudCredentials/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCredentialsServer).Create(ctx, req.(*CloudCredentialCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudCredentials_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudCredentialUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCredentialsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.v1beta1.CloudCredentials/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCredentialsServer).Update(ctx, req.(*CloudCredentialUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudCredentials_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudCredentialDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCredentialsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.v1beta1.CloudCredentials/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCredentialsServer).Delete(ctx, req.(*CloudCredentialDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CloudCredentials_serviceDesc = grpc.ServiceDesc{
	ServiceName: "credential.v1beta1.CloudCredentials",
	HandlerType: (*CloudCredentialsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _CloudCredentials_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CloudCredentials_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CloudCredentials_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CloudCredentials_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("cloudcredential.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x94, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x99, 0xb6, 0x56, 0xfb, 0x0a, 0xb2, 0x0c, 0xab, 0x84, 0x58, 0xa4, 0xe4, 0x20, 0xcb,
	0x82, 0x89, 0x1b, 0x85, 0xd5, 0x8a, 0xa0, 0x76, 0xbd, 0x89, 0x87, 0x88, 0x1e, 0x3c, 0x08, 0xb3,
	0x9d, 0xe7, 0x12, 0xcc, 0xce, 0x8c, 0x99, 0x69, 0xa1, 0x88, 0x17, 0x2f, 0x1e, 0x3d, 0xec, 0x47,
	0xf0, 0x23, 0x79, 0xf0, 0x0b, 0xf8, 0x09, 0xfc, 0x04, 0x92, 0x99, 0x98, 0xa6, 0xbb, 0xb5, 0xa4,
	0x45, 0xbc, 0x84, 0x99, 0xff, 0x9b, 0xf7, 0xe6, 0xf7, 0xde, 0xbc, 0x17, 0xb8, 0x36, 0xc9, 0xe4,
	0x94, 0x4f, 0x72, 0xe4, 0x28, 0x4c, 0xca, 0xb2, 0x50, 0xe5, 0xd2, 0x48, 0x4a, 0x6b, 0xca, 0xec,
	0xe0, 0x18, 0x0d, 0x3b, 0xf0, 0x07, 0x27, 0x52, 0x9e, 0x64, 0x18, 0x31, 0x95, 0x46, 0x4c, 0x08,
	0x69, 0x98, 0x49, 0xa5, 0xd0, 0xce, 0xc3, 0xbf, 0x5e, 0xc8, 0xdc, 0xcc, 0x15, 0xea, 0xc8, 0x7e,
	0x9d, 0x1e, 0xfc, 0x20, 0x30, 0x18, 0x17, 0x77, 0x8c, 0xab, 0x88, 0xe3, 0x1c, 0x99, 0xc1, 0x04,
	0x3f, 0x4c, 0x51, 0x1b, 0x4a, 0xa1, 0x23, 0xd8, 0x29, 0x7a, 0x64, 0x48, 0xf6, 0x7a, 0x89, 0x5d,
	0x53, 0x1f, 0xae, 0xa8, 0x5c, 0xce, 0x52, 0x8e, 0xb9, 0xd7, 0xb2, 0x7a, 0xb5, 0xa7, 0x2f, 0xa0,
	0xc3, 0x99, 0x61, 0x5e, 0x7b, 0xd8, 0xde, 0xeb, 0xc7, 0xa3, 0xf0, 0x22, 0x69, 0xb8, 0xee, 0xbe,
	0xf0, 0x88, 0x19, 0xf6, 0x4c, 0x98, 0x7c, 0x9e, 0xd8, 0x38, 0xfe, 0x21, 0xf4, 0x2a, 0x89, 0xee,
	0x40, 0xfb, 0x3d, 0xce, 0x4b, 0x96, 0x62, 0x49, 0x77, 0xe1, 0xd2, 0x8c, 0x65, 0x53, 0x2c, 0x39,
	0xdc, 0x66, 0xd4, 0xba, 0x4f, 0x56, 0x65, 0xf6, 0x4a, 0xf1, 0xff, 0x9a, 0xd9, 0xd2, 0x7d, 0xff,
	0x2e, 0xb3, 0xf8, 0x42, 0x62, 0x47, 0x98, 0xe1, 0xda, 0xc4, 0x82, 0x01, 0xf8, 0xe7, 0x7c, 0x9e,
	0xa7, 0xda, 0x94, 0x1e, 0xc1, 0x17, 0x02, 0x37, 0x56, 0x9a, 0xb5, 0x92, 0x42, 0x23, 0xbd, 0x05,
	0x5d, 0x6d, 0x98, 0x99, 0x6a, 0x1b, 0xb3, 0x1f, 0x5f, 0x0d, 0x5d, 0x2b, 0x85, 0x2f, 0xad, 0x9a,
	0x94, 0x56, 0xfa, 0x18, 0xfa, 0x8b, 0xaa, 0x68, 0xaf, 0x65, 0x2b, 0x75, 0x73, 0x65, 0xa5, 0x2a,
	0x29, 0xa9, 0xbb, 0x04, 0x6f, 0x01, 0x16, 0xa6, 0x8d, 0x9f, 0x68, 0x08, 0xfd, 0x54, 0xbc, 0x93,
	0xf9, 0xa9, 0xed, 0x7d, 0xaf, 0x6d, 0xcd, 0x75, 0x29, 0xfe, 0xd5, 0x81, 0x9d, 0x73, 0x99, 0x6a,
	0xfa, 0x8d, 0x40, 0xa7, 0xc8, 0x97, 0x86, 0x0d, 0x1e, 0xb5, 0x56, 0x37, 0x3f, 0x6a, 0x7c, 0xde,
	0x15, 0x32, 0x18, 0x7d, 0xfe, 0xfe, 0xf3, 0xac, 0x75, 0x8f, 0xc6, 0x11, 0x53, 0x4a, 0x4f, 0x24,
	0x77, 0xf3, 0xba, 0x88, 0x12, 0x95, 0x51, 0x22, 0x3b, 0xf8, 0xb7, 0x6b, 0xa5, 0xa1, 0x5f, 0x09,
	0x74, 0xdd, 0xac, 0xd0, 0x3b, 0x9b, 0x8e, 0x95, 0xbf, 0xfb, 0xe7, 0xc5, 0x5e, 0xcb, 0x94, 0x57,
	0x38, 0x8f, 0x2c, 0xce, 0x61, 0xb0, 0x05, 0xce, 0x88, 0xec, 0x5b, 0x22, 0xd7, 0xe3, 0x8d, 0x88,
	0x96, 0xc6, 0x61, 0x3d, 0x91, 0xbf, 0x25, 0xd1, 0x19, 0x81, 0xae, 0x1b, 0x86, 0x46, 0x44, 0x4b,
	0x73, 0xf3, 0x17, 0xa2, 0x27, 0x96, 0xe8, 0xe1, 0xfe, 0x83, 0xcd, 0x89, 0xa2, 0x8f, 0x45, 0xc7,
	0x7e, 0x7a, 0xda, 0x7b, 0x73, 0xb9, 0x3c, 0x73, 0xdc, 0xb5, 0xbf, 0xdd, 0xbb, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x5d, 0x87, 0xde, 0x03, 0xd9, 0x05, 0x00, 0x00,
}