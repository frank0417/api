// Code generated by protoc-gen-go.
// source: incident.proto
// DO NOT EDIT!

package alert

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

type Incident struct {
	Phid                 string      `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	KubernetesCluster    string      `protobuf:"bytes,2,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
	KubernetesNamespace  string      `protobuf:"bytes,3,opt,name=kubernetes_namespace,json=kubernetesNamespace" json:"kubernetes_namespace,omitempty"`
	KubernetesObjectType string      `protobuf:"bytes,4,opt,name=kubernetes_objectType,json=kubernetesObjectType" json:"kubernetes_objectType,omitempty"`
	KubernetesObjectName string      `protobuf:"bytes,5,opt,name=kubernetes_objectName,json=kubernetesObjectName" json:"kubernetes_objectName,omitempty"`
	Alert                *dtypes.Uid `protobuf:"bytes,6,opt,name=alert" json:"alert,omitempty"`
	IcingaHost           string      `protobuf:"bytes,7,opt,name=icinga_host,json=icingaHost" json:"icinga_host,omitempty"`
	IcingaService        string      `protobuf:"bytes,8,opt,name=icinga_service,json=icingaService" json:"icinga_service,omitempty"`
	Type                 string      `protobuf:"bytes,9,opt,name=type" json:"type,omitempty"`
	State                string      `protobuf:"bytes,10,opt,name=state" json:"state,omitempty"`
	User                 *dtypes.Uid `protobuf:"bytes,11,opt,name=user" json:"user,omitempty"`
	// Timestamp of first reported event
	Reported int64 `protobuf:"varint,12,opt,name=reported" json:"reported,omitempty"`
	// Timestamp of first acknowledgement
	Acknowledged int64             `protobuf:"varint,13,opt,name=acknowledged" json:"acknowledged,omitempty"`
	Recovered    int64             `protobuf:"varint,14,opt,name=recovered" json:"recovered,omitempty"`
	IcingawebUrl string            `protobuf:"bytes,15,opt,name=icingaweb_url,json=icingawebUrl" json:"icingaweb_url,omitempty"`
	Events       []*Incident_Event `protobuf:"bytes,16,rep,name=events" json:"events,omitempty"`
}

func (m *Incident) Reset()                    { *m = Incident{} }
func (m *Incident) String() string            { return proto.CompactTextString(m) }
func (*Incident) ProtoMessage()               {}
func (*Incident) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Incident) GetAlert() *dtypes.Uid {
	if m != nil {
		return m.Alert
	}
	return nil
}

func (m *Incident) GetUser() *dtypes.Uid {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Incident) GetEvents() []*Incident_Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type Incident_Event struct {
	Type     string      `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	State    string      `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	Reported int64       `protobuf:"varint,3,opt,name=reported" json:"reported,omitempty"`
	User     *dtypes.Uid `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	Comment  string      `protobuf:"bytes,5,opt,name=comment" json:"comment,omitempty"`
}

func (m *Incident_Event) Reset()                    { *m = Incident_Event{} }
func (m *Incident_Event) String() string            { return proto.CompactTextString(m) }
func (*Incident_Event) ProtoMessage()               {}
func (*Incident_Event) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func (m *Incident_Event) GetUser() *dtypes.Uid {
	if m != nil {
		return m.User
	}
	return nil
}

// Next Id: 6
type IncidentListRequest struct {
	KubernetesCluster    string   `protobuf:"bytes,1,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
	KubernetesNamespace  string   `protobuf:"bytes,2,opt,name=kubernetes_namespace,json=kubernetesNamespace" json:"kubernetes_namespace,omitempty"`
	KubernetesObjectType string   `protobuf:"bytes,3,opt,name=kubernetes_objectType,json=kubernetesObjectType" json:"kubernetes_objectType,omitempty"`
	KubernetesObjectName string   `protobuf:"bytes,4,opt,name=kubernetes_objectName,json=kubernetesObjectName" json:"kubernetes_objectName,omitempty"`
	State                []string `protobuf:"bytes,5,rep,name=state" json:"state,omitempty"`
}

func (m *IncidentListRequest) Reset()                    { *m = IncidentListRequest{} }
func (m *IncidentListRequest) String() string            { return proto.CompactTextString(m) }
func (*IncidentListRequest) ProtoMessage()               {}
func (*IncidentListRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type IncidentListResponse struct {
	Status    *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Incidents []*Incident    `protobuf:"bytes,2,rep,name=incidents" json:"incidents,omitempty"`
}

func (m *IncidentListResponse) Reset()                    { *m = IncidentListResponse{} }
func (m *IncidentListResponse) String() string            { return proto.CompactTextString(m) }
func (*IncidentListResponse) ProtoMessage()               {}
func (*IncidentListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *IncidentListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *IncidentListResponse) GetIncidents() []*Incident {
	if m != nil {
		return m.Incidents
	}
	return nil
}

// Next Id: 2
type IncidentDescribeRequest struct {
	Phid string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
}

func (m *IncidentDescribeRequest) Reset()                    { *m = IncidentDescribeRequest{} }
func (m *IncidentDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*IncidentDescribeRequest) ProtoMessage()               {}
func (*IncidentDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type IncidentDescribeResponse struct {
	Status   *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Incident *Incident      `protobuf:"bytes,2,opt,name=incident" json:"incident,omitempty"`
}

func (m *IncidentDescribeResponse) Reset()                    { *m = IncidentDescribeResponse{} }
func (m *IncidentDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*IncidentDescribeResponse) ProtoMessage()               {}
func (*IncidentDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *IncidentDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *IncidentDescribeResponse) GetIncident() *Incident {
	if m != nil {
		return m.Incident
	}
	return nil
}

// Next Id: 9
type IncidentNotifyRequest struct {
	AlertPhid string `protobuf:"bytes,1,opt,name=alert_phid,json=alertPhid" json:"alert_phid,omitempty"`
	HostName  string `protobuf:"bytes,2,opt,name=host_name,json=hostName" json:"host_name,omitempty"`
	Type      string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	State     string `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
	Output    string `protobuf:"bytes,5,opt,name=output" json:"output,omitempty"`
	Time      string `protobuf:"bytes,6,opt,name=time" json:"time,omitempty"`
	Author    string `protobuf:"bytes,7,opt,name=author" json:"author,omitempty"`
	Comment   string `protobuf:"bytes,8,opt,name=comment" json:"comment,omitempty"`
}

func (m *IncidentNotifyRequest) Reset()                    { *m = IncidentNotifyRequest{} }
func (m *IncidentNotifyRequest) String() string            { return proto.CompactTextString(m) }
func (*IncidentNotifyRequest) ProtoMessage()               {}
func (*IncidentNotifyRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

// Next Id: 4
type IncidentEventCreateRequest struct {
	// Incident PHID
	Phid        string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Comment     string `protobuf:"bytes,2,opt,name=comment" json:"comment,omitempty"`
	Acknowledge bool   `protobuf:"varint,3,opt,name=acknowledge" json:"acknowledge,omitempty"`
}

func (m *IncidentEventCreateRequest) Reset()                    { *m = IncidentEventCreateRequest{} }
func (m *IncidentEventCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*IncidentEventCreateRequest) ProtoMessage()               {}
func (*IncidentEventCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func init() {
	proto.RegisterType((*Incident)(nil), "alert.Incident")
	proto.RegisterType((*Incident_Event)(nil), "alert.Incident.Event")
	proto.RegisterType((*IncidentListRequest)(nil), "alert.IncidentListRequest")
	proto.RegisterType((*IncidentListResponse)(nil), "alert.IncidentListResponse")
	proto.RegisterType((*IncidentDescribeRequest)(nil), "alert.IncidentDescribeRequest")
	proto.RegisterType((*IncidentDescribeResponse)(nil), "alert.IncidentDescribeResponse")
	proto.RegisterType((*IncidentNotifyRequest)(nil), "alert.IncidentNotifyRequest")
	proto.RegisterType((*IncidentEventCreateRequest)(nil), "alert.IncidentEventCreateRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Incidents service

type IncidentsClient interface {
	List(ctx context.Context, in *IncidentListRequest, opts ...grpc.CallOption) (*IncidentListResponse, error)
	Describe(ctx context.Context, in *IncidentDescribeRequest, opts ...grpc.CallOption) (*IncidentDescribeResponse, error)
	Notify(ctx context.Context, in *IncidentNotifyRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	CreateEvent(ctx context.Context, in *IncidentEventCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type incidentsClient struct {
	cc *grpc.ClientConn
}

func NewIncidentsClient(cc *grpc.ClientConn) IncidentsClient {
	return &incidentsClient{cc}
}

func (c *incidentsClient) List(ctx context.Context, in *IncidentListRequest, opts ...grpc.CallOption) (*IncidentListResponse, error) {
	out := new(IncidentListResponse)
	err := grpc.Invoke(ctx, "/alert.Incidents/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsClient) Describe(ctx context.Context, in *IncidentDescribeRequest, opts ...grpc.CallOption) (*IncidentDescribeResponse, error) {
	out := new(IncidentDescribeResponse)
	err := grpc.Invoke(ctx, "/alert.Incidents/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsClient) Notify(ctx context.Context, in *IncidentNotifyRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/alert.Incidents/Notify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsClient) CreateEvent(ctx context.Context, in *IncidentEventCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/alert.Incidents/CreateEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Incidents service

type IncidentsServer interface {
	List(context.Context, *IncidentListRequest) (*IncidentListResponse, error)
	Describe(context.Context, *IncidentDescribeRequest) (*IncidentDescribeResponse, error)
	Notify(context.Context, *IncidentNotifyRequest) (*dtypes.VoidResponse, error)
	CreateEvent(context.Context, *IncidentEventCreateRequest) (*dtypes.VoidResponse, error)
}

func RegisterIncidentsServer(s *grpc.Server, srv IncidentsServer) {
	s.RegisterService(&_Incidents_serviceDesc, srv)
}

func _Incidents_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncidentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.Incidents/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServer).List(ctx, req.(*IncidentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Incidents_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncidentDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.Incidents/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServer).Describe(ctx, req.(*IncidentDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Incidents_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncidentNotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.Incidents/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServer).Notify(ctx, req.(*IncidentNotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Incidents_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncidentEventCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.Incidents/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServer).CreateEvent(ctx, req.(*IncidentEventCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Incidents_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alert.Incidents",
	HandlerType: (*IncidentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Incidents_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Incidents_Describe_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _Incidents_Notify_Handler,
		},
		{
			MethodName: "CreateEvent",
			Handler:    _Incidents_CreateEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("incident.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 827 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0xcf, 0x6f, 0xe3, 0x44,
	0x14, 0x96, 0xe3, 0x24, 0xeb, 0x3c, 0x77, 0xb3, 0x30, 0x9b, 0x2e, 0x23, 0x6f, 0xa1, 0x59, 0xf3,
	0x2b, 0x74, 0x95, 0x78, 0x9b, 0x15, 0x97, 0xe5, 0xb8, 0x20, 0x81, 0x84, 0x16, 0xe4, 0xa5, 0x5c,
	0x23, 0xc7, 0x7e, 0xa4, 0x26, 0x89, 0xc7, 0x78, 0xc6, 0xa9, 0x2a, 0x04, 0x07, 0xb8, 0x70, 0xe7,
	0x4f, 0xe3, 0x88, 0xb8, 0x71, 0xe3, 0x0f, 0xe0, 0x8a, 0x3c, 0xe3, 0x71, 0x9c, 0x34, 0x29, 0x6d,
	0x2f, 0x55, 0xe6, 0x7b, 0xbf, 0xbf, 0xf7, 0xe5, 0x35, 0xd0, 0x8d, 0x93, 0x30, 0x8e, 0x30, 0x11,
	0xa3, 0x34, 0x63, 0x82, 0x91, 0x56, 0xb0, 0xc0, 0x4c, 0x38, 0x47, 0x33, 0xc6, 0x66, 0x0b, 0xf4,
	0x82, 0x34, 0xf6, 0x82, 0x24, 0x61, 0x22, 0x10, 0x31, 0x4b, 0xb8, 0x72, 0x72, 0x1e, 0x15, 0x70,
	0x24, 0x2e, 0x53, 0xe4, 0x9e, 0xfc, 0xab, 0x70, 0xf7, 0xaf, 0x16, 0x58, 0x5f, 0x94, 0xf9, 0x08,
	0x81, 0x66, 0x7a, 0x1e, 0x47, 0xd4, 0xe8, 0x1b, 0x83, 0x8e, 0x2f, 0x3f, 0x93, 0x21, 0x90, 0x79,
	0x3e, 0xc5, 0x2c, 0x41, 0x81, 0x7c, 0x12, 0x2e, 0x72, 0x2e, 0x30, 0xa3, 0x0d, 0xe9, 0xf1, 0xe6,
	0xda, 0xf2, 0x52, 0x19, 0xc8, 0x29, 0xf4, 0x6a, 0xee, 0x49, 0xb0, 0x44, 0x9e, 0x06, 0x21, 0x52,
	0x53, 0x06, 0x3c, 0x5c, 0xdb, 0x5e, 0x69, 0x13, 0x79, 0x0e, 0x87, 0xb5, 0x10, 0x36, 0xfd, 0x1e,
	0x43, 0xf1, 0xcd, 0x65, 0x8a, 0xb4, 0x29, 0x63, 0x6a, 0xf9, 0xbe, 0xaa, 0x6c, 0x3b, 0x83, 0x8a,
	0x94, 0xb4, 0xb5, 0x3b, 0xa8, 0xb0, 0x91, 0x27, 0xa0, 0xb8, 0xa2, 0xed, 0xbe, 0x31, 0xb0, 0xc7,
	0xf6, 0x48, 0x11, 0x32, 0x3a, 0x8b, 0x23, 0x5f, 0x59, 0xc8, 0x31, 0xd8, 0x71, 0x18, 0x27, 0xb3,
	0x60, 0x72, 0xce, 0xb8, 0xa0, 0xf7, 0x64, 0x36, 0x50, 0xd0, 0xe7, 0x8c, 0x0b, 0xf2, 0x3e, 0x74,
	0x4b, 0x07, 0x8e, 0xd9, 0x2a, 0x0e, 0x91, 0x5a, 0xd2, 0xe7, 0xbe, 0x42, 0x5f, 0x2b, 0xb0, 0xa0,
	0xb2, 0xc8, 0x4d, 0x3b, 0x8a, 0xca, 0xe2, 0x33, 0xe9, 0x41, 0x8b, 0x8b, 0x40, 0x20, 0x05, 0x09,
	0xaa, 0x07, 0x39, 0x86, 0x66, 0xce, 0x31, 0xa3, 0xf6, 0xd5, 0x9e, 0xa4, 0x81, 0x38, 0x60, 0x65,
	0x98, 0xb2, 0x4c, 0x60, 0x44, 0x0f, 0xfa, 0xc6, 0xc0, 0xf4, 0xab, 0x37, 0x71, 0xe1, 0x20, 0x08,
	0xe7, 0x09, 0xbb, 0x58, 0x60, 0x34, 0xc3, 0x88, 0xde, 0x97, 0xf6, 0x0d, 0x8c, 0x1c, 0x41, 0x27,
	0xc3, 0x90, 0xad, 0x30, 0xc3, 0x88, 0x76, 0xa5, 0xc3, 0x1a, 0x20, 0xef, 0x42, 0xd9, 0xf9, 0x05,
	0x4e, 0x27, 0x79, 0xb6, 0xa0, 0x0f, 0x64, 0x73, 0x07, 0x15, 0x78, 0x96, 0x2d, 0xc8, 0x10, 0xda,
	0xb8, 0xc2, 0x44, 0x70, 0xfa, 0x46, 0xdf, 0x1c, 0xd8, 0xe3, 0xc3, 0x91, 0x64, 0x6b, 0xa4, 0x95,
	0x33, 0xfa, 0xac, 0xb0, 0xfa, 0xa5, 0x93, 0xf3, 0x9b, 0x01, 0x2d, 0x89, 0x54, 0x34, 0x18, 0xbb,
	0x68, 0x68, 0xd4, 0x69, 0xa8, 0x4f, 0x69, 0x6e, 0x4d, 0xa9, 0x29, 0x6a, 0xee, 0xa3, 0x88, 0xc2,
	0xbd, 0x90, 0x2d, 0x97, 0x98, 0x88, 0x72, 0xff, 0xfa, 0xe9, 0xfe, 0x6b, 0xc0, 0x43, 0xdd, 0xe5,
	0x97, 0x31, 0x17, 0x3e, 0xfe, 0x90, 0x23, 0x17, 0x7b, 0x64, 0x6d, 0xdc, 0x56, 0xd6, 0x8d, 0x3b,
	0xc8, 0xda, 0xbc, 0x8b, 0xac, 0x9b, 0xd7, 0xc8, 0xba, 0x22, 0xb4, 0xd5, 0x37, 0x2b, 0x42, 0xdd,
	0x25, 0xf4, 0x36, 0x07, 0xe7, 0x29, 0x4b, 0x38, 0x92, 0x0f, 0xa0, 0x5d, 0x38, 0xe4, 0x5c, 0x4e,
	0x6b, 0x8f, 0xbb, 0x9a, 0xce, 0xd7, 0x12, 0xf5, 0x4b, 0x2b, 0x19, 0x42, 0x47, 0x1f, 0x1a, 0x4e,
	0x1b, 0x72, 0xed, 0x0f, 0xb6, 0xd6, 0xee, 0xaf, 0x3d, 0xdc, 0x21, 0xbc, 0xa5, 0xe1, 0x4f, 0x91,
	0x87, 0x59, 0x3c, 0x45, 0xcd, 0xf5, 0x8e, 0xb3, 0xe2, 0x32, 0xa0, 0x57, 0xdd, 0x6f, 0xd9, 0xe1,
	0x53, 0xb0, 0x74, 0x7d, 0xb9, 0x88, 0x1d, 0x0d, 0x56, 0x0e, 0xee, 0x9f, 0x06, 0x1c, 0x6a, 0xf8,
	0x15, 0x13, 0xf1, 0x77, 0x97, 0xba, 0xbd, 0xb7, 0x01, 0x64, 0xd4, 0xa4, 0xd6, 0x64, 0x47, 0x22,
	0x5f, 0x17, 0x07, 0xf0, 0x31, 0x74, 0x8a, 0x53, 0x20, 0x97, 0x5e, 0xee, 0xdb, 0x2a, 0x00, 0x49,
	0xbd, 0xd6, 0xb7, 0xb9, 0x4b, 0xdf, 0xcd, 0xba, 0xbe, 0x1f, 0x41, 0x9b, 0xe5, 0x22, 0xcd, 0xb5,
	0x42, 0xcb, 0x97, 0xcc, 0x10, 0x2f, 0x51, 0x9e, 0xa4, 0x22, 0x43, 0xbc, 0x94, 0xbe, 0x41, 0x2e,
	0xce, 0x59, 0x56, 0xde, 0x9f, 0xf2, 0x55, 0x97, 0xb9, 0xb5, 0x29, 0xf3, 0x05, 0x38, 0x7a, 0x38,
	0xf9, 0xc5, 0x7b, 0x99, 0x61, 0x20, 0xae, 0x5b, 0x40, 0x3d, 0x57, 0x63, 0x23, 0x17, 0xe9, 0x83,
	0x5d, 0xbb, 0x1f, 0x72, 0x34, 0xcb, 0xaf, 0x43, 0xe3, 0x7f, 0x4c, 0xe8, 0xe8, 0x72, 0x9c, 0xcc,
	0xa1, 0x59, 0x08, 0x8c, 0x38, 0x5b, 0xe4, 0xd7, 0xbe, 0x6e, 0xce, 0xe3, 0x9d, 0x36, 0xb5, 0x6f,
	0xf7, 0xe4, 0x97, 0x3f, 0xfe, 0xfe, 0xbd, 0xf1, 0x1e, 0x71, 0xbd, 0x20, 0x4d, 0x79, 0xc8, 0xa2,
	0xf2, 0x9f, 0x58, 0x11, 0xe1, 0xad, 0x9e, 0x8d, 0x4e, 0xbd, 0x4a, 0x66, 0xe4, 0x67, 0xb0, 0xb4,
	0x5e, 0xc8, 0x3b, 0x5b, 0x49, 0xb7, 0x74, 0xe7, 0x1c, 0xef, 0xb5, 0x97, 0x85, 0x4f, 0x65, 0xe1,
	0xa7, 0xe4, 0xa3, 0xff, 0x2f, 0xec, 0xfd, 0x58, 0xb0, 0xf6, 0x13, 0x99, 0x43, 0x5b, 0xa9, 0x87,
	0x1c, 0x6d, 0x65, 0xdf, 0x10, 0x95, 0xd3, 0xd3, 0x9a, 0xfd, 0x96, 0xc5, 0x51, 0x55, 0x70, 0x28,
	0x0b, 0x7e, 0xe8, 0xde, 0x60, 0xd2, 0x17, 0xc6, 0x09, 0xf9, 0xd5, 0x00, 0x5b, 0x6d, 0x52, 0x5d,
	0xd3, 0x27, 0x5b, 0x25, 0xaf, 0xae, 0x7a, 0x4f, 0xdd, 0x4f, 0x64, 0xdd, 0x8f, 0xdd, 0x67, 0x37,
	0x1e, 0xd4, 0x53, 0xa7, 0xfc, 0x85, 0x71, 0x32, 0x6d, 0xcb, 0x5f, 0x0a, 0xcf, 0xff, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0x03, 0x4b, 0x5f, 0x57, 0x78, 0x08, 0x00, 0x00,
}