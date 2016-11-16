// Code generated by protoc-gen-go.
// source: loadbalancer.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	loadbalancer.proto

It has these top-level messages:
	ListRequest
	ListResponse
	DescribeRequest
	DescribeResponse
	CreateRequest
	UpdateRequest
	DeleteRequest
	LoadBalancer
	Spec
	Status
	LoadBalancerStatus
	LoadBalancerBackend
	LoadBalancerRule
	HTTPLoadBalancerRule
	TCPLoadBalancerRule
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

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

type ListRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ListRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

type ListResponse struct {
	Status        *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	LoadBalancers []*LoadBalancer         `protobuf:"bytes,2,rep,name=load_balancers,json=loadBalancers" json:"load_balancers,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListResponse) GetLoadBalancers() []*LoadBalancer {
	if m != nil {
		return m.LoadBalancers
	}
	return nil
}

type DescribeRequest struct {
	Kind      string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	Cluster   string `protobuf:"bytes,4,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *DescribeRequest) Reset()                    { *m = DescribeRequest{} }
func (m *DescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeRequest) ProtoMessage()               {}
func (*DescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DescribeRequest) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *DescribeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DescribeRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DescribeRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

type DescribeResponse struct {
	Status       *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	LoadBalancer *LoadBalancer           `protobuf:"bytes,2,opt,name=load_balancer,json=loadBalancer" json:"load_balancer,omitempty"`
}

func (m *DescribeResponse) Reset()                    { *m = DescribeResponse{} }
func (m *DescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeResponse) ProtoMessage()               {}
func (*DescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DescribeResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeResponse) GetLoadBalancer() *LoadBalancer {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

type CreateRequest struct {
	Name         string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Namespace    string        `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	Cluster      string        `protobuf:"bytes,3,opt,name=cluster" json:"cluster,omitempty"`
	LoadBalancer *LoadBalancer `protobuf:"bytes,4,opt,name=load_balancer,json=loadBalancer" json:"load_balancer,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *CreateRequest) GetLoadBalancer() *LoadBalancer {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

type UpdateRequest struct {
	Name         string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Cluster      string        `protobuf:"bytes,2,opt,name=cluster" json:"cluster,omitempty"`
	LoadBalancer *LoadBalancer `protobuf:"bytes,3,opt,name=load_balancer,json=loadBalancer" json:"load_balancer,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *UpdateRequest) GetLoadBalancer() *LoadBalancer {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

type DeleteRequest struct {
	Kind      string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	Cluster   string `protobuf:"bytes,4,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteRequest) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *DeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeleteRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DeleteRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

type LoadBalancer struct {
	// 'kind' defines is it the regular kubernetes instance or the
	// appscode superset called Extended Ingress. This field will
	// strictly contains only those two values
	// 'ingress' - default kubernetes ingress object.
	// 'extendedIngress' - appscode superset of ingress.
	// when creating a Loadbalancer from UI this field will always
	// be only 'extendedIngress.' List, Describe, Update and Delete
	// will support both two modes.
	// Create will support only extendedIngress.
	// For Creating or Updating an regular ingress one must use the
	// kubectl or direct API calls directly to kubernetes.
	Kind              string            `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Name              string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace         string            `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	CreationTimestamp int64             `protobuf:"varint,4,opt,name=creation_timestamp,json=creationTimestamp" json:"creation_timestamp,omitempty"`
	Options           map[string]string `protobuf:"bytes,5,rep,name=options" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Spec              *Spec             `protobuf:"bytes,6,opt,name=spec" json:"spec,omitempty"`
	Status            *Status           `protobuf:"bytes,7,opt,name=status" json:"status,omitempty"`
	Json              string            `protobuf:"bytes,8,opt,name=json" json:"json,omitempty"`
}

func (m *LoadBalancer) Reset()                    { *m = LoadBalancer{} }
func (m *LoadBalancer) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancer) ProtoMessage()               {}
func (*LoadBalancer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *LoadBalancer) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *LoadBalancer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoadBalancer) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *LoadBalancer) GetCreationTimestamp() int64 {
	if m != nil {
		return m.CreationTimestamp
	}
	return 0
}

func (m *LoadBalancer) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *LoadBalancer) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *LoadBalancer) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *LoadBalancer) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

type Spec struct {
	Backend *HTTPLoadBalancerRule `protobuf:"bytes,1,opt,name=backend" json:"backend,omitempty"`
	Rules   []*LoadBalancerRule   `protobuf:"bytes,2,rep,name=rules" json:"rules,omitempty"`
}

func (m *Spec) Reset()                    { *m = Spec{} }
func (m *Spec) String() string            { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()               {}
func (*Spec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Spec) GetBackend() *HTTPLoadBalancerRule {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (m *Spec) GetRules() []*LoadBalancerRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

type Status struct {
	Status []*LoadBalancerStatus `protobuf:"bytes,1,rep,name=status" json:"status,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Status) GetStatus() []*LoadBalancerStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type LoadBalancerStatus struct {
	IP   string `protobuf:"bytes,1,opt,name=IP" json:"IP,omitempty"`
	Host string `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
}

func (m *LoadBalancerStatus) Reset()                    { *m = LoadBalancerStatus{} }
func (m *LoadBalancerStatus) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancerStatus) ProtoMessage()               {}
func (*LoadBalancerStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *LoadBalancerStatus) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *LoadBalancerStatus) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type LoadBalancerBackend struct {
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	ServicePort string `protobuf:"bytes,2,opt,name=service_port,json=servicePort" json:"service_port,omitempty"`
}

func (m *LoadBalancerBackend) Reset()                    { *m = LoadBalancerBackend{} }
func (m *LoadBalancerBackend) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancerBackend) ProtoMessage()               {}
func (*LoadBalancerBackend) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *LoadBalancerBackend) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *LoadBalancerBackend) GetServicePort() string {
	if m != nil {
		return m.ServicePort
	}
	return ""
}

type LoadBalancerRule struct {
	Host string `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	// ssl secret name to enable https on the host.
	// ssl secret must contain data with the certs pem file.
	SSLSecretName string                  `protobuf:"bytes,5,opt,name=SSL_secret_name,json=SSLSecretName" json:"SSL_secret_name,omitempty"`
	Http          []*HTTPLoadBalancerRule `protobuf:"bytes,2,rep,name=http" json:"http,omitempty"`
	Tcp           []*TCPLoadBalancerRule  `protobuf:"bytes,3,rep,name=tcp" json:"tcp,omitempty"`
}

func (m *LoadBalancerRule) Reset()                    { *m = LoadBalancerRule{} }
func (m *LoadBalancerRule) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancerRule) ProtoMessage()               {}
func (*LoadBalancerRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *LoadBalancerRule) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *LoadBalancerRule) GetSSLSecretName() string {
	if m != nil {
		return m.SSLSecretName
	}
	return ""
}

func (m *LoadBalancerRule) GetHttp() []*HTTPLoadBalancerRule {
	if m != nil {
		return m.Http
	}
	return nil
}

func (m *LoadBalancerRule) GetTcp() []*TCPLoadBalancerRule {
	if m != nil {
		return m.Tcp
	}
	return nil
}

type HTTPLoadBalancerRule struct {
	Path         string               `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Backend      *LoadBalancerBackend `protobuf:"bytes,2,opt,name=backend" json:"backend,omitempty"`
	HeaderRules  []string             `protobuf:"bytes,3,rep,name=header_rules,json=headerRules" json:"header_rules,omitempty"`
	RewriteRules []string             `protobuf:"bytes,4,rep,name=rewrite_rules,json=rewriteRules" json:"rewrite_rules,omitempty"`
}

func (m *HTTPLoadBalancerRule) Reset()                    { *m = HTTPLoadBalancerRule{} }
func (m *HTTPLoadBalancerRule) String() string            { return proto.CompactTextString(m) }
func (*HTTPLoadBalancerRule) ProtoMessage()               {}
func (*HTTPLoadBalancerRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *HTTPLoadBalancerRule) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HTTPLoadBalancerRule) GetBackend() *LoadBalancerBackend {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (m *HTTPLoadBalancerRule) GetHeaderRules() []string {
	if m != nil {
		return m.HeaderRules
	}
	return nil
}

func (m *HTTPLoadBalancerRule) GetRewriteRules() []string {
	if m != nil {
		return m.RewriteRules
	}
	return nil
}

type TCPLoadBalancerRule struct {
	Port          string               `protobuf:"bytes,1,opt,name=port" json:"port,omitempty"`
	Backend       *LoadBalancerBackend `protobuf:"bytes,2,opt,name=backend" json:"backend,omitempty"`
	SSLSecretName string               `protobuf:"bytes,3,opt,name=SSL_secret_name,json=SSLSecretName" json:"SSL_secret_name,omitempty"`
	SecretPemName string               `protobuf:"bytes,4,opt,name=secret_pem_name,json=secretPemName" json:"secret_pem_name,omitempty"`
}

func (m *TCPLoadBalancerRule) Reset()                    { *m = TCPLoadBalancerRule{} }
func (m *TCPLoadBalancerRule) String() string            { return proto.CompactTextString(m) }
func (*TCPLoadBalancerRule) ProtoMessage()               {}
func (*TCPLoadBalancerRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *TCPLoadBalancerRule) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *TCPLoadBalancerRule) GetBackend() *LoadBalancerBackend {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (m *TCPLoadBalancerRule) GetSSLSecretName() string {
	if m != nil {
		return m.SSLSecretName
	}
	return ""
}

func (m *TCPLoadBalancerRule) GetSecretPemName() string {
	if m != nil {
		return m.SecretPemName
	}
	return ""
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "appscode.loadbalancer.v1beta1.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "appscode.loadbalancer.v1beta1.ListResponse")
	proto.RegisterType((*DescribeRequest)(nil), "appscode.loadbalancer.v1beta1.DescribeRequest")
	proto.RegisterType((*DescribeResponse)(nil), "appscode.loadbalancer.v1beta1.DescribeResponse")
	proto.RegisterType((*CreateRequest)(nil), "appscode.loadbalancer.v1beta1.CreateRequest")
	proto.RegisterType((*UpdateRequest)(nil), "appscode.loadbalancer.v1beta1.UpdateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "appscode.loadbalancer.v1beta1.DeleteRequest")
	proto.RegisterType((*LoadBalancer)(nil), "appscode.loadbalancer.v1beta1.LoadBalancer")
	proto.RegisterType((*Spec)(nil), "appscode.loadbalancer.v1beta1.Spec")
	proto.RegisterType((*Status)(nil), "appscode.loadbalancer.v1beta1.Status")
	proto.RegisterType((*LoadBalancerStatus)(nil), "appscode.loadbalancer.v1beta1.LoadBalancerStatus")
	proto.RegisterType((*LoadBalancerBackend)(nil), "appscode.loadbalancer.v1beta1.LoadBalancerBackend")
	proto.RegisterType((*LoadBalancerRule)(nil), "appscode.loadbalancer.v1beta1.LoadBalancerRule")
	proto.RegisterType((*HTTPLoadBalancerRule)(nil), "appscode.loadbalancer.v1beta1.HTTPLoadBalancerRule")
	proto.RegisterType((*TCPLoadBalancerRule)(nil), "appscode.loadbalancer.v1beta1.TCPLoadBalancerRule")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LoadBalancers service

type LoadBalancersClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*DescribeResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type loadBalancersClient struct {
	cc *grpc.ClientConn
}

func NewLoadBalancersClient(cc *grpc.ClientConn) LoadBalancersClient {
	return &loadBalancersClient{cc}
}

func (c *loadBalancersClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/appscode.loadbalancer.v1beta1.LoadBalancers/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*DescribeResponse, error) {
	out := new(DescribeResponse)
	err := grpc.Invoke(ctx, "/appscode.loadbalancer.v1beta1.LoadBalancers/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.loadbalancer.v1beta1.LoadBalancers/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.loadbalancer.v1beta1.LoadBalancers/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.loadbalancer.v1beta1.LoadBalancers/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LoadBalancers service

type LoadBalancersServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Describe(context.Context, *DescribeRequest) (*DescribeResponse, error)
	Create(context.Context, *CreateRequest) (*appscode_dtypes.VoidResponse, error)
	Update(context.Context, *UpdateRequest) (*appscode_dtypes.VoidResponse, error)
	Delete(context.Context, *DeleteRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterLoadBalancersServer(s *grpc.Server, srv LoadBalancersServer) {
	s.RegisterService(&_LoadBalancers_serviceDesc, srv)
}

func _LoadBalancers_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.loadbalancer.v1beta1.LoadBalancers/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancers_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.loadbalancer.v1beta1.LoadBalancers/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).Describe(ctx, req.(*DescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancers_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.loadbalancer.v1beta1.LoadBalancers/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancers_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.loadbalancer.v1beta1.LoadBalancers/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancers_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.loadbalancer.v1beta1.LoadBalancers/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoadBalancers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.loadbalancer.v1beta1.LoadBalancers",
	HandlerType: (*LoadBalancersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _LoadBalancers_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _LoadBalancers_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _LoadBalancers_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _LoadBalancers_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _LoadBalancers_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loadbalancer.proto",
}

func init() { proto.RegisterFile("loadbalancer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1012 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x57, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0xd7, 0x78, 0x1d, 0x3b, 0x79, 0xb6, 0x9b, 0x74, 0x5a, 0x89, 0x95, 0xd5, 0x42, 0xbb, 0x55,
	0x4b, 0x95, 0x82, 0x57, 0x71, 0x0f, 0x44, 0x91, 0x38, 0x90, 0x3f, 0xd0, 0xa0, 0x94, 0x5a, 0xeb,
	0x50, 0x09, 0x38, 0x58, 0xe3, 0xf5, 0x53, 0xb3, 0x64, 0xbd, 0xb3, 0xdd, 0x19, 0x07, 0x45, 0x55,
	0x2f, 0xfd, 0x02, 0x1c, 0x2a, 0xc4, 0x05, 0x89, 0x33, 0x87, 0x4a, 0x5c, 0xb8, 0x23, 0x71, 0xe0,
	0x03, 0x70, 0xe6, 0xc6, 0x07, 0x41, 0x33, 0xb3, 0x1b, 0xef, 0x36, 0x6e, 0xb7, 0x6e, 0xa2, 0x5e,
	0x92, 0xd9, 0xf7, 0xf7, 0xf7, 0x7e, 0xef, 0xcd, 0xfa, 0x2d, 0xd0, 0x90, 0xb3, 0xd1, 0x90, 0x85,
	0x2c, 0xf2, 0x31, 0xe9, 0xc4, 0x09, 0x97, 0x9c, 0x5e, 0x65, 0x71, 0x2c, 0x7c, 0x3e, 0xc2, 0x4e,
	0x41, 0x79, 0xb4, 0x36, 0x44, 0xc9, 0xd6, 0xda, 0x57, 0x1e, 0x71, 0xfe, 0x28, 0x44, 0x97, 0xc5,
	0x81, 0xcb, 0xa2, 0x88, 0x4b, 0x26, 0x03, 0x1e, 0x09, 0xe3, 0xdc, 0x7e, 0x3f, 0x73, 0x7e, 0x85,
	0xfe, 0x83, 0x82, 0x7e, 0x24, 0x8f, 0x63, 0x14, 0xae, 0xfe, 0x6b, 0x0c, 0x9c, 0x0f, 0xa1, 0xb1,
	0x17, 0x08, 0xe9, 0xe1, 0xe3, 0x09, 0x0a, 0x49, 0x6d, 0xa8, 0xfb, 0xe1, 0x44, 0x48, 0x4c, 0x6c,
	0x72, 0x8d, 0xdc, 0x5e, 0xf2, 0xb2, 0x47, 0xe7, 0x39, 0x81, 0xa6, 0xb1, 0x14, 0x31, 0x8f, 0x04,
	0x52, 0x17, 0x6a, 0x42, 0x32, 0x39, 0x11, 0xda, 0xb2, 0xd1, 0x7d, 0xaf, 0x73, 0x52, 0x88, 0xc9,
	0xd3, 0xe9, 0x6b, 0xb5, 0x97, 0x9a, 0x51, 0x0f, 0x2e, 0xa8, 0x0a, 0x07, 0x59, 0x89, 0xc2, 0xae,
	0x5c, 0xb3, 0x6e, 0x37, 0xba, 0x77, 0x3a, 0xaf, 0x65, 0xa0, 0xb3, 0xc7, 0xd9, 0x68, 0x33, 0x15,
	0x7a, 0xad, 0x30, 0xf7, 0x24, 0x9c, 0xc7, 0xb0, 0xbc, 0x8d, 0xc2, 0x4f, 0x82, 0x21, 0x66, 0x25,
	0x50, 0xa8, 0x1e, 0x06, 0xd1, 0x28, 0xc5, 0xaf, 0xcf, 0x4a, 0x16, 0xb1, 0x31, 0xda, 0x15, 0x23,
	0x53, 0x67, 0x7a, 0x05, 0x96, 0xd4, 0x7f, 0x11, 0x33, 0x1f, 0x6d, 0x4b, 0x2b, 0xa6, 0x82, 0x3c,
	0x11, 0xd5, 0x22, 0x11, 0x3f, 0x11, 0x58, 0x99, 0xe6, 0x7c, 0x5b, 0x32, 0x7a, 0xd0, 0x2a, 0x90,
	0xa1, 0xa1, 0xcd, 0xc9, 0x45, 0x33, 0xcf, 0x85, 0xf3, 0x82, 0x40, 0x6b, 0x2b, 0x41, 0x26, 0xf3,
	0x4c, 0xe8, 0xaa, 0xc9, 0xab, 0xaa, 0xae, 0xbc, 0xa6, 0x6a, 0xab, 0x50, 0xf5, 0x69, 0xbc, 0xd5,
	0xb3, 0xe2, 0xfd, 0x91, 0x40, 0xeb, 0xeb, 0x78, 0x54, 0x82, 0x37, 0x87, 0xa8, 0x52, 0x82, 0xc8,
	0x3a, 0x2b, 0x22, 0x0e, 0xad, 0x6d, 0x0c, 0x51, 0xbe, 0xb3, 0x51, 0xfa, 0xd5, 0x82, 0x66, 0x1e,
	0xcf, 0x39, 0x25, 0xfc, 0x18, 0xa8, 0xaf, 0x06, 0x21, 0xe0, 0xd1, 0x40, 0x06, 0x63, 0x14, 0x92,
	0x8d, 0x63, 0x9d, 0xdb, 0xf2, 0x2e, 0x66, 0x9a, 0xfd, 0x4c, 0x41, 0x3d, 0xa8, 0xf3, 0x58, 0xbf,
	0x34, 0xec, 0x05, 0x7d, 0x21, 0xd7, 0xe7, 0xa0, 0xb0, 0xf3, 0xc0, 0xb8, 0xee, 0x44, 0x32, 0x39,
	0xf6, 0xb2, 0x40, 0xf4, 0x13, 0xa8, 0x8a, 0x18, 0x7d, 0xbb, 0xa6, 0x7b, 0x72, 0xa3, 0x24, 0x60,
	0x3f, 0x46, 0xdf, 0xd3, 0x0e, 0xf4, 0xd3, 0x93, 0x8b, 0x54, 0xd7, 0xae, 0x37, 0xcb, 0x5c, 0x8b,
	0xd7, 0x8a, 0x42, 0xf5, 0x7b, 0xc1, 0x23, 0x7b, 0xd1, 0x90, 0xa5, 0xce, 0xed, 0x0d, 0x68, 0xe6,
	0x41, 0xd2, 0x15, 0xb0, 0x0e, 0xf1, 0x38, 0xe5, 0x58, 0x1d, 0xe9, 0x65, 0x58, 0x38, 0x62, 0xe1,
	0x24, 0xe3, 0xd8, 0x3c, 0x6c, 0x54, 0xd6, 0x89, 0xf3, 0x0b, 0x81, 0xaa, 0x42, 0x47, 0xef, 0x43,
	0x7d, 0xc8, 0xfc, 0x43, 0x4c, 0x9b, 0xd3, 0xe8, 0xde, 0x2d, 0x01, 0x76, 0x6f, 0x7f, 0xbf, 0x57,
	0x98, 0xb5, 0x49, 0x88, 0x5e, 0x16, 0x83, 0xee, 0xc0, 0x42, 0x32, 0x09, 0x31, 0x7b, 0x05, 0xba,
	0xf3, 0x0c, 0xad, 0x0a, 0x64, 0xbc, 0x9d, 0x3e, 0xd4, 0x0c, 0x01, 0x74, 0x37, 0xf7, 0x02, 0x52,
	0x11, 0xd7, 0xe6, 0x88, 0x58, 0xe4, 0xd0, 0x59, 0x07, 0x7a, 0x5a, 0x4b, 0x2f, 0x40, 0x65, 0xb7,
	0x97, 0x92, 0x56, 0xd9, 0xed, 0x29, 0xa6, 0x0f, 0xb8, 0x90, 0xd9, 0x58, 0xaa, 0xb3, 0xf3, 0x1d,
	0x5c, 0xca, 0x7b, 0x6e, 0xa6, 0xc5, 0x5e, 0x87, 0xa6, 0xc0, 0xe4, 0x28, 0xf0, 0x71, 0x90, 0xbb,
	0xdf, 0x8d, 0x54, 0xf6, 0x95, 0x1a, 0xe8, 0x9c, 0x49, 0xcc, 0x93, 0x2c, 0x6a, 0x66, 0xd2, 0xe3,
	0x89, 0x74, 0xfe, 0x25, 0xb0, 0xf2, 0x32, 0x0f, 0x27, 0x28, 0xc8, 0x14, 0x05, 0xbd, 0x05, 0xcb,
	0xfd, 0xfe, 0xde, 0x40, 0xa0, 0x9f, 0xa0, 0x34, 0x19, 0x17, 0xb4, 0xba, 0xd5, 0xef, 0xef, 0xf5,
	0xb5, 0x54, 0xe7, 0xfc, 0x02, 0xaa, 0x07, 0x52, 0xc6, 0x69, 0x0b, 0xde, 0xaa, 0x9f, 0x3a, 0x00,
	0xdd, 0x06, 0x4b, 0xfa, 0xb1, 0x6d, 0xe9, 0x38, 0xdd, 0x92, 0x38, 0xfb, 0x5b, 0xa7, 0xc3, 0x28,
	0x77, 0xe7, 0x4f, 0x02, 0x97, 0x67, 0x25, 0x51, 0x35, 0xc6, 0x4c, 0x1e, 0x64, 0x35, 0xaa, 0x33,
	0xdd, 0x9b, 0x8e, 0xa3, 0xf9, 0xe1, 0xe8, 0xce, 0xd1, 0xef, 0xb4, 0x2f, 0xd3, 0x69, 0xbc, 0x0e,
	0xcd, 0x03, 0x64, 0x23, 0x4c, 0x06, 0x66, 0x28, 0x55, 0x25, 0x4b, 0x5e, 0xc3, 0xc8, 0x14, 0x06,
	0x41, 0x6f, 0x40, 0x2b, 0xc1, 0x1f, 0x92, 0x40, 0x62, 0x6a, 0x53, 0xd5, 0x36, 0xcd, 0x54, 0xa8,
	0x8d, 0x9c, 0xbf, 0x09, 0x5c, 0x9a, 0x51, 0x9f, 0xae, 0x40, 0x75, 0x35, 0xab, 0x80, 0x27, 0xf2,
	0x9c, 0x2b, 0x98, 0xd1, 0x73, 0x6b, 0x56, 0xcf, 0x6f, 0xc1, 0x72, 0x6a, 0x13, 0xe3, 0xd8, 0xd8,
	0x99, 0x77, 0x72, 0xcb, 0x88, 0x7b, 0x38, 0x56, 0x76, 0xdd, 0x9f, 0xeb, 0xd0, 0xca, 0x27, 0x14,
	0xf4, 0x77, 0x02, 0x55, 0xb5, 0xff, 0xd0, 0xd5, 0x32, 0x9c, 0xd3, 0x75, 0xaa, 0x7d, 0xe7, 0x8d,
	0x6c, 0xcd, 0x0e, 0xe1, 0x3c, 0x78, 0xf6, 0x87, 0x5d, 0x59, 0x24, 0xcf, 0xfe, 0xf9, 0xef, 0x79,
	0x65, 0x8b, 0x7e, 0xe6, 0x16, 0x56, 0xb7, 0xc3, 0xc9, 0x10, 0x93, 0x08, 0x25, 0x0a, 0x37, 0x75,
	0x76, 0xd3, 0x5f, 0x11, 0xe1, 0x3e, 0x49, 0x4f, 0x4f, 0xdd, 0x7c, 0x12, 0x41, 0xff, 0x22, 0xb0,
	0x98, 0x6d, 0x2a, 0xb4, 0x53, 0x02, 0xe5, 0xa5, 0x35, 0xaa, 0xed, 0xbe, 0xb1, 0x7d, 0x0a, 0xff,
	0x61, 0x0e, 0xfe, 0x97, 0xf4, 0xde, 0x99, 0xe1, 0xbb, 0x4f, 0x54, 0x7b, 0x9e, 0xd2, 0xdf, 0x08,
	0xd4, 0xcc, 0x5e, 0x43, 0x3f, 0x2a, 0xc1, 0x54, 0x58, 0x7f, 0xda, 0x57, 0x4f, 0xed, 0x60, 0x0f,
	0x79, 0x30, 0x3a, 0xc1, 0xeb, 0xe5, 0xf0, 0x7e, 0xee, 0x9c, 0x9d, 0xee, 0x0d, 0xb2, 0xaa, 0x86,
	0xa4, 0x66, 0x76, 0x9a, 0x52, 0xac, 0x85, 0xd5, 0xa7, 0x0c, 0xeb, 0x37, 0x39, 0xac, 0xf7, 0xdb,
	0xe7, 0xc6, 0xad, 0x82, 0xfc, 0x82, 0x40, 0xcd, 0x6c, 0x3d, 0xa5, 0x90, 0x0b, 0xcb, 0x51, 0x19,
	0xe4, 0xc2, 0x38, 0xac, 0x9e, 0x1b, 0xe4, 0xcd, 0x1d, 0xb8, 0xe9, 0xf3, 0xf1, 0x34, 0x37, 0x8b,
	0x83, 0x99, 0x68, 0x37, 0x2f, 0xe6, 0xef, 0x6f, 0x4f, 0x7d, 0xec, 0xf4, 0xc8, 0xb7, 0xf5, 0x54,
	0x3b, 0xac, 0xe9, 0xcf, 0x9f, 0xbb, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xde, 0x17, 0x1c, 0x0b,
	0x92, 0x0d, 0x00, 0x00,
}
