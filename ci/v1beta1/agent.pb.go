// Code generated by protoc-gen-go.
// source: agent.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	agent.proto
	build.proto
	job.proto
	master.proto

It has these top-level messages:
	AgentListRequest
	AgentListResponse
	Agent
	AgentCreateRequest
	PortInfo
	AgentDescribeRequest
	AgentDescribeResponse
	AgentDeleteRequest
	AgentRestartRequest
	AgentRestartResponse
	BuildDescribeRequest
	BuildDescribeResponse
	BuildListRequest
	BuildListResponse
	Build
	JobListRequest
	JobListResponse
	JobBuildRequest
	JobDescribeRequest
	JobDescribeResponse
	JobHealth
	JobDeleteRequest
	JobCreateRequest
	JobCopyRequest
	Job
	MasterCreateRequest
	MasterDeleteRequest
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

type AgentListRequest struct {
	// List of status to get the agent filterd on the status
	// values in
	//   PENDING
	//   FAILED
	//   ONLINE
	//   OFFLINE
	//   DELETED
	Status []string `protobuf:"bytes,1,rep,name=status" json:"status,omitempty"`
}

func (m *AgentListRequest) Reset()                    { *m = AgentListRequest{} }
func (m *AgentListRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentListRequest) ProtoMessage()               {}
func (*AgentListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AgentListResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Agents []*Agent                `protobuf:"bytes,2,rep,name=agents" json:"agents,omitempty"`
}

func (m *AgentListResponse) Reset()                    { *m = AgentListResponse{} }
func (m *AgentListResponse) String() string            { return proto.CompactTextString(m) }
func (*AgentListResponse) ProtoMessage()               {}
func (*AgentListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AgentListResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AgentListResponse) GetAgents() []*Agent {
	if m != nil {
		return m.Agents
	}
	return nil
}

type Agent struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Status      string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	AgentStatus string `protobuf:"bytes,3,opt,name=agent_status,json=agentStatus" json:"agent_status,omitempty"`
	IsRefreshed int32  `protobuf:"varint,4,opt,name=is_refreshed,json=isRefreshed" json:"is_refreshed,omitempty"`
	CreatedAt   int64  `protobuf:"varint,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt   int64  `protobuf:"varint,6,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *Agent) Reset()                    { *m = Agent{} }
func (m *Agent) String() string            { return proto.CompactTextString(m) }
func (*Agent) ProtoMessage()               {}
func (*Agent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type AgentCreateRequest struct {
	Sku               string      `protobuf:"bytes,1,opt,name=sku" json:"sku,omitempty"`
	Executors         int32       `protobuf:"varint,2,opt,name=executors" json:"executors,omitempty"`
	Labels            string      `protobuf:"bytes,3,opt,name=labels" json:"labels,omitempty"`
	UserStartupScript string      `protobuf:"bytes,4,opt,name=user_startup_script,json=userStartupScript" json:"user_startup_script,omitempty"`
	SaltbaseVersion   string      `protobuf:"bytes,5,opt,name=saltbase_version,json=saltbaseVersion" json:"saltbase_version,omitempty"`
	CiStarterVersion  string      `protobuf:"bytes,6,opt,name=ci_starter_version,json=ciStarterVersion" json:"ci_starter_version,omitempty"`
	Ports             []*PortInfo `protobuf:"bytes,7,rep,name=ports" json:"ports,omitempty"`
	Role              string      `protobuf:"bytes,8,opt,name=role" json:"role,omitempty"`
}

func (m *AgentCreateRequest) Reset()                    { *m = AgentCreateRequest{} }
func (m *AgentCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentCreateRequest) ProtoMessage()               {}
func (*AgentCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AgentCreateRequest) GetPorts() []*PortInfo {
	if m != nil {
		return m.Ports
	}
	return nil
}

type PortInfo struct {
	Protocol  string `protobuf:"bytes,1,opt,name=protocol" json:"protocol,omitempty"`
	PortRange string `protobuf:"bytes,2,opt,name=port_range,json=portRange" json:"port_range,omitempty"`
}

func (m *PortInfo) Reset()                    { *m = PortInfo{} }
func (m *PortInfo) String() string            { return proto.CompactTextString(m) }
func (*PortInfo) ProtoMessage()               {}
func (*PortInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type AgentDescribeRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *AgentDescribeRequest) Reset()                    { *m = AgentDescribeRequest{} }
func (m *AgentDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentDescribeRequest) ProtoMessage()               {}
func (*AgentDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type AgentDescribeResponse struct {
	Status           *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Name             string                  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Executors        int64                   `protobuf:"varint,3,opt,name=executors" json:"executors,omitempty"`
	AgentStatus      string                  `protobuf:"bytes,4,opt,name=agent_status,json=agentStatus" json:"agent_status,omitempty"`
	AgentStatusCause string                  `protobuf:"bytes,5,opt,name=agent_status_cause,json=agentStatusCause" json:"agent_status_cause,omitempty"`
	Label            string                  `protobuf:"bytes,6,opt,name=label" json:"label,omitempty"`
	Provider         string                  `protobuf:"bytes,7,opt,name=provider" json:"provider,omitempty"`
	Sku              string                  `protobuf:"bytes,8,opt,name=sku" json:"sku,omitempty"`
	StartupScript    string                  `protobuf:"bytes,9,opt,name=startup_script,json=startupScript" json:"startup_script,omitempty"`
	CreatedAt        int64                   `protobuf:"varint,10,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt        int64                   `protobuf:"varint,11,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *AgentDescribeResponse) Reset()                    { *m = AgentDescribeResponse{} }
func (m *AgentDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*AgentDescribeResponse) ProtoMessage()               {}
func (*AgentDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AgentDescribeResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type AgentDeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *AgentDeleteRequest) Reset()                    { *m = AgentDeleteRequest{} }
func (m *AgentDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentDeleteRequest) ProtoMessage()               {}
func (*AgentDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type AgentRestartRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *AgentRestartRequest) Reset()                    { *m = AgentRestartRequest{} }
func (m *AgentRestartRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentRestartRequest) ProtoMessage()               {}
func (*AgentRestartRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type AgentRestartResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *AgentRestartResponse) Reset()                    { *m = AgentRestartResponse{} }
func (m *AgentRestartResponse) String() string            { return proto.CompactTextString(m) }
func (*AgentRestartResponse) ProtoMessage()               {}
func (*AgentRestartResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AgentRestartResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*AgentListRequest)(nil), "appscode.ci.v1beta1.AgentListRequest")
	proto.RegisterType((*AgentListResponse)(nil), "appscode.ci.v1beta1.AgentListResponse")
	proto.RegisterType((*Agent)(nil), "appscode.ci.v1beta1.Agent")
	proto.RegisterType((*AgentCreateRequest)(nil), "appscode.ci.v1beta1.AgentCreateRequest")
	proto.RegisterType((*PortInfo)(nil), "appscode.ci.v1beta1.PortInfo")
	proto.RegisterType((*AgentDescribeRequest)(nil), "appscode.ci.v1beta1.AgentDescribeRequest")
	proto.RegisterType((*AgentDescribeResponse)(nil), "appscode.ci.v1beta1.AgentDescribeResponse")
	proto.RegisterType((*AgentDeleteRequest)(nil), "appscode.ci.v1beta1.AgentDeleteRequest")
	proto.RegisterType((*AgentRestartRequest)(nil), "appscode.ci.v1beta1.AgentRestartRequest")
	proto.RegisterType((*AgentRestartResponse)(nil), "appscode.ci.v1beta1.AgentRestartResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Agents service

type AgentsClient interface {
	List(ctx context.Context, in *AgentListRequest, opts ...grpc.CallOption) (*AgentListResponse, error)
	Describe(ctx context.Context, in *AgentDescribeRequest, opts ...grpc.CallOption) (*AgentDescribeResponse, error)
	Create(ctx context.Context, in *AgentCreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.LongRunningResponse, error)
	Delete(ctx context.Context, in *AgentDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.LongRunningResponse, error)
	Restart(ctx context.Context, in *AgentRestartRequest, opts ...grpc.CallOption) (*AgentRestartResponse, error)
}

type agentsClient struct {
	cc *grpc.ClientConn
}

func NewAgentsClient(cc *grpc.ClientConn) AgentsClient {
	return &agentsClient{cc}
}

func (c *agentsClient) List(ctx context.Context, in *AgentListRequest, opts ...grpc.CallOption) (*AgentListResponse, error) {
	out := new(AgentListResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Agents/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) Describe(ctx context.Context, in *AgentDescribeRequest, opts ...grpc.CallOption) (*AgentDescribeResponse, error) {
	out := new(AgentDescribeResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Agents/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) Create(ctx context.Context, in *AgentCreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.LongRunningResponse, error) {
	out := new(appscode_dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Agents/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) Delete(ctx context.Context, in *AgentDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.LongRunningResponse, error) {
	out := new(appscode_dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Agents/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) Restart(ctx context.Context, in *AgentRestartRequest, opts ...grpc.CallOption) (*AgentRestartResponse, error) {
	out := new(AgentRestartResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Agents/Restart", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Agents service

type AgentsServer interface {
	List(context.Context, *AgentListRequest) (*AgentListResponse, error)
	Describe(context.Context, *AgentDescribeRequest) (*AgentDescribeResponse, error)
	Create(context.Context, *AgentCreateRequest) (*appscode_dtypes.LongRunningResponse, error)
	Delete(context.Context, *AgentDeleteRequest) (*appscode_dtypes.LongRunningResponse, error)
	Restart(context.Context, *AgentRestartRequest) (*AgentRestartResponse, error)
}

func RegisterAgentsServer(s *grpc.Server, srv AgentsServer) {
	s.RegisterService(&_Agents_serviceDesc, srv)
}

func _Agents_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Agents/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).List(ctx, req.(*AgentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Agents/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).Describe(ctx, req.(*AgentDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Agents/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).Create(ctx, req.(*AgentCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Agents/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).Delete(ctx, req.(*AgentDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentRestartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Agents/Restart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).Restart(ctx, req.(*AgentRestartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agents_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.ci.v1beta1.Agents",
	HandlerType: (*AgentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Agents_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Agents_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Agents_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Agents_Delete_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _Agents_Restart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("agent.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 821 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0xcd, 0x8e, 0xe3, 0x44,
	0x10, 0x96, 0xf3, 0xe3, 0xc4, 0x15, 0x7e, 0xb2, 0x3d, 0xcb, 0x62, 0x59, 0xbb, 0x90, 0xb5, 0x76,
	0x17, 0x27, 0x5a, 0x6c, 0x26, 0x03, 0x42, 0xe2, 0x36, 0xcc, 0x20, 0x84, 0x34, 0x07, 0xe4, 0x48,
	0x1c, 0xb8, 0x44, 0x1d, 0xa7, 0x27, 0x58, 0x04, 0xb7, 0xe9, 0x6e, 0x8f, 0x06, 0x21, 0x2e, 0x23,
	0x2e, 0x1c, 0x38, 0x71, 0xe0, 0x01, 0xb8, 0x73, 0x43, 0x3c, 0x08, 0xaf, 0xc0, 0x8d, 0x2b, 0x0f,
	0xb0, 0xea, 0x1f, 0x7b, 0xec, 0xcc, 0x24, 0x13, 0xcd, 0x25, 0xea, 0xae, 0xfa, 0xba, 0xeb, 0xab,
	0xaa, 0xaf, 0x2b, 0x86, 0x01, 0x5e, 0x91, 0x4c, 0x84, 0x39, 0xa3, 0x82, 0xa2, 0x03, 0x9c, 0xe7,
	0x3c, 0xa1, 0x4b, 0x12, 0x26, 0x69, 0x78, 0x71, 0xb8, 0x20, 0x02, 0x1f, 0x7a, 0x8f, 0x57, 0x94,
	0xae, 0xd6, 0x24, 0xc2, 0x79, 0x1a, 0xe1, 0x2c, 0xa3, 0x02, 0x8b, 0x94, 0x66, 0x5c, 0x1f, 0xf1,
	0xde, 0x29, 0x8f, 0x6c, 0xf1, 0x3f, 0x92, 0xe6, 0xa5, 0xf8, 0x21, 0x27, 0x3c, 0x52, 0xbf, 0xda,
	0xee, 0x4f, 0x60, 0x78, 0x2c, 0x23, 0x9f, 0xa5, 0x5c, 0xc4, 0xe4, 0xfb, 0x82, 0x70, 0x81, 0x1e,
	0x81, 0xcd, 0x05, 0x16, 0x05, 0x77, 0xad, 0x51, 0x3b, 0x70, 0x62, 0xb3, 0xf3, 0x2f, 0xe1, 0x41,
	0x0d, 0xcb, 0x73, 0x9a, 0x71, 0x82, 0xa2, 0x1a, 0xd8, 0x0a, 0x06, 0xd3, 0xb7, 0xc3, 0x8a, 0xbc,
	0x0e, 0x17, 0xce, 0x94, 0xbb, 0xbc, 0x05, 0x4d, 0xc1, 0x56, 0xb9, 0x72, 0xb7, 0x35, 0x6a, 0x07,
	0x83, 0xa9, 0x17, 0xde, 0x92, 0x6d, 0xa8, 0x02, 0xc5, 0x06, 0xe9, 0xff, 0x6d, 0x41, 0x57, 0x59,
	0x10, 0x82, 0x4e, 0x86, 0xbf, 0x23, 0x2a, 0x98, 0x13, 0xab, 0x75, 0x8d, 0x6f, 0x4b, 0x59, 0xcb,
	0x48, 0x4f, 0xe1, 0x35, 0x75, 0x7e, 0x6e, 0xbc, 0x6d, 0xe5, 0xd5, 0x95, 0x9e, 0x55, 0x90, 0x94,
	0xcf, 0x19, 0x39, 0x67, 0x84, 0x7f, 0x43, 0x96, 0x6e, 0x67, 0x64, 0x05, 0xdd, 0x78, 0x90, 0xf2,
	0xb8, 0x34, 0xa1, 0x27, 0x00, 0x09, 0x23, 0x58, 0x90, 0xe5, 0x1c, 0x0b, 0xb7, 0x3b, 0xb2, 0x82,
	0x76, 0xec, 0x18, 0xcb, 0xb1, 0x90, 0xee, 0x22, 0x5f, 0x96, 0x6e, 0x5b, 0xbb, 0x8d, 0xe5, 0x58,
	0xf8, 0x7f, 0xb6, 0x00, 0x29, 0xe6, 0x27, 0xea, 0x44, 0x59, 0xe2, 0x21, 0xb4, 0xf9, 0xb7, 0x85,
	0xc9, 0x42, 0x2e, 0xd1, 0x63, 0x70, 0xc8, 0x25, 0x49, 0x0a, 0x41, 0x99, 0xce, 0xa3, 0x1b, 0x5f,
	0x1b, 0x64, 0x8a, 0x6b, 0xbc, 0x20, 0xeb, 0x32, 0x09, 0xb3, 0x43, 0x21, 0x1c, 0x14, 0x9c, 0x30,
	0x99, 0x21, 0x13, 0x45, 0x3e, 0xe7, 0x09, 0x4b, 0x73, 0xa1, 0xd2, 0x70, 0xe2, 0x07, 0xd2, 0x35,
	0xd3, 0x9e, 0x99, 0x72, 0xa0, 0x31, 0x0c, 0x39, 0x5e, 0x8b, 0x05, 0xe6, 0x64, 0x7e, 0x41, 0x18,
	0x4f, 0x69, 0xa6, 0x52, 0x72, 0xe2, 0x37, 0x4b, 0xfb, 0x57, 0xda, 0x8c, 0x5e, 0x02, 0x4a, 0x52,
	0x7d, 0x31, 0x61, 0x15, 0xd8, 0x56, 0xe0, 0x61, 0x92, 0xce, 0xb4, 0xa3, 0x44, 0x1f, 0x41, 0x37,
	0xa7, 0x4c, 0x70, 0xb7, 0xa7, 0x9a, 0xfa, 0xe4, 0xd6, 0xa6, 0x7e, 0x49, 0x99, 0xf8, 0x22, 0x3b,
	0xa7, 0xb1, 0xc6, 0xca, 0x66, 0x32, 0xba, 0x26, 0x6e, 0x5f, 0x37, 0x53, 0xae, 0xfd, 0xcf, 0xa0,
	0x5f, 0xc2, 0x90, 0x07, 0x7d, 0xa5, 0xd2, 0x84, 0xae, 0x4d, 0xa9, 0xaa, 0xbd, 0xac, 0xbb, 0xbc,
	0x64, 0xce, 0x70, 0xb6, 0x22, 0xa6, 0xf1, 0x8e, 0xb4, 0xc4, 0xd2, 0xe0, 0x4f, 0xe0, 0xa1, 0x2a,
	0xfb, 0x29, 0x91, 0x25, 0x59, 0x54, 0x85, 0xbf, 0x45, 0x3f, 0xfe, 0xff, 0x2d, 0x78, 0x6b, 0x03,
	0x7c, 0x5f, 0x71, 0x97, 0xd7, 0xb7, 0x6a, 0xf2, 0x6c, 0x74, 0xb6, 0xad, 0x05, 0x72, 0xdd, 0xd9,
	0x4d, 0x91, 0x76, 0x6e, 0x8a, 0xf4, 0x25, 0xa0, 0x3a, 0x64, 0x9e, 0xe0, 0x82, 0x13, 0xd3, 0xb6,
	0x61, 0x0d, 0x78, 0x22, 0xed, 0xe8, 0x21, 0x74, 0x95, 0x38, 0x4c, 0xab, 0xf4, 0xc6, 0x94, 0xf2,
	0x22, 0x5d, 0x12, 0xe6, 0xf6, 0xaa, 0x52, 0xaa, 0x7d, 0x29, 0xc6, 0xfe, 0xb5, 0x18, 0x9f, 0xc3,
	0x1b, 0x1b, 0x8a, 0x72, 0x94, 0xf3, 0x75, 0xde, 0x50, 0x53, 0xf3, 0x69, 0xc0, 0xee, 0xa7, 0x31,
	0xd8, 0x7c, 0x1a, 0x81, 0x79, 0x19, 0xa7, 0x64, 0x4d, 0xc4, 0xce, 0x06, 0x8d, 0xe1, 0x40, 0xcf,
	0x03, 0xa2, 0xe2, 0xef, 0x82, 0x7e, 0x6e, 0xfa, 0x5e, 0x41, 0xef, 0xd9, 0xc9, 0xe9, 0x7f, 0x5d,
	0xb0, 0xd5, 0x4d, 0x1c, 0xfd, 0x6c, 0x41, 0x47, 0xce, 0x3c, 0xf4, 0x7c, 0xfb, 0xa8, 0xaa, 0xcd,
	0x4f, 0xef, 0xc5, 0x5d, 0x30, 0xcd, 0xc9, 0x7f, 0xff, 0xea, 0x2f, 0xb7, 0xd5, 0xb7, 0xae, 0xfe,
	0xf9, 0xf7, 0xb7, 0xd6, 0x53, 0xf4, 0x6e, 0xd4, 0x18, 0xe1, 0x49, 0x1a, 0x99, 0x93, 0x91, 0x1e,
	0x82, 0xe8, 0x77, 0x0b, 0xfa, 0xa5, 0x42, 0xd1, 0x78, 0x7b, 0x8c, 0x0d, 0xc9, 0x7b, 0x93, 0x7d,
	0xa0, 0x86, 0xd2, 0x87, 0x35, 0x4a, 0x01, 0x7a, 0x71, 0x07, 0xa5, 0xe8, 0x47, 0x59, 0xf3, 0x9f,
	0xd0, 0x2f, 0x16, 0xd8, 0x7a, 0xbe, 0xa1, 0xf7, 0xb6, 0x07, 0x6b, 0x4c, 0x40, 0xef, 0xd9, 0x8d,
	0x06, 0x9c, 0xd1, 0x6c, 0x15, 0x17, 0x59, 0x96, 0x66, 0xab, 0x8a, 0xcf, 0x07, 0x35, 0x3e, 0xcf,
	0xfc, 0xbb, 0x4a, 0xf4, 0x89, 0x35, 0x41, 0xbf, 0x5a, 0x60, 0x6b, 0x45, 0xed, 0xe2, 0xd2, 0xd0,
	0xdc, 0x9e, 0x5c, 0x1a, 0xb5, 0x99, 0xec, 0x5b, 0x9b, 0x3f, 0x2c, 0xe8, 0x19, 0x31, 0xa2, 0x60,
	0xc7, 0x5f, 0x5d, 0x43, 0xda, 0xde, 0x78, 0x0f, 0xa4, 0xa1, 0x75, 0x52, 0xa3, 0xf5, 0xb1, 0xff,
	0xd1, 0x7e, 0xb4, 0x22, 0x9c, 0xa8, 0xaf, 0x83, 0x88, 0x91, 0x05, 0xa5, 0xe2, 0x53, 0xe7, 0xeb,
	0x9e, 0xc1, 0x2d, 0x6c, 0x35, 0x62, 0x8f, 0x5e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x33, 0xf6,
	0xae, 0x92, 0x08, 0x00, 0x00,
}
