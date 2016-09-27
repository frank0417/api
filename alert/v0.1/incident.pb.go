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
	KubernetesNode       string      `protobuf:"bytes,6,opt,name=kubernetes_node,json=kubernetesNode" json:"kubernetes_node,omitempty"`
	KubernetesPod        string      `protobuf:"bytes,7,opt,name=kubernetes_pod,json=kubernetesPod" json:"kubernetes_pod,omitempty"`
	Alert                *dtypes.Uid `protobuf:"bytes,8,opt,name=alert" json:"alert,omitempty"`
	IcingaHost           string      `protobuf:"bytes,9,opt,name=icinga_host,json=icingaHost" json:"icinga_host,omitempty"`
	IcingaService        string      `protobuf:"bytes,10,opt,name=icinga_service,json=icingaService" json:"icinga_service,omitempty"`
	Type                 string      `protobuf:"bytes,11,opt,name=type" json:"type,omitempty"`
	State                string      `protobuf:"bytes,12,opt,name=state" json:"state,omitempty"`
	User                 *dtypes.Uid `protobuf:"bytes,13,opt,name=user" json:"user,omitempty"`
	// Timestamp of first reported event
	Reported int64 `protobuf:"varint,14,opt,name=reported" json:"reported,omitempty"`
	// Timestamp of first acknowledgement
	Acknowledged int64             `protobuf:"varint,15,opt,name=acknowledged" json:"acknowledged,omitempty"`
	Recovered    int64             `protobuf:"varint,16,opt,name=recovered" json:"recovered,omitempty"`
	IcingawebUrl string            `protobuf:"bytes,17,opt,name=icingaweb_url,json=icingawebUrl" json:"icingaweb_url,omitempty"`
	Events       []*Incident_Event `protobuf:"bytes,18,rep,name=events" json:"events,omitempty"`
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

type IncidentNotifyRequest struct {
	AlertPhid string `protobuf:"bytes,1,opt,name=alert_phid,json=alertPhid" json:"alert_phid,omitempty"`
	HostName  string `protobuf:"bytes,2,opt,name=host_name,json=hostName" json:"host_name,omitempty"`
	Type      string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	State     string `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
	Output    string `protobuf:"bytes,5,opt,name=output" json:"output,omitempty"`
	Time      string `protobuf:"bytes,6,opt,name=time" json:"time,omitempty"`
}

func (m *IncidentNotifyRequest) Reset()                    { *m = IncidentNotifyRequest{} }
func (m *IncidentNotifyRequest) String() string            { return proto.CompactTextString(m) }
func (*IncidentNotifyRequest) ProtoMessage()               {}
func (*IncidentNotifyRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

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
	// 845 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0x4f, 0x73, 0xdb, 0x44,
	0x14, 0x1f, 0x59, 0xb6, 0x6b, 0x3f, 0xa5, 0x0e, 0xdd, 0x3a, 0x65, 0x47, 0x0d, 0xc4, 0x15, 0x7f,
	0x6a, 0xd2, 0xb1, 0xd5, 0xb8, 0xc3, 0xa5, 0x1c, 0x0b, 0x33, 0x30, 0xc3, 0x94, 0x8e, 0x42, 0xb8,
	0x7a, 0x64, 0xe9, 0xe1, 0x08, 0xdb, 0x5a, 0xa1, 0x5d, 0x39, 0x93, 0x61, 0xe0, 0x00, 0x17, 0xee,
	0x7c, 0x0a, 0x0e, 0x7c, 0x1a, 0xbe, 0x02, 0x37, 0x3e, 0x00, 0x57, 0x46, 0xbb, 0x5a, 0x59, 0x76,
	0xec, 0x34, 0xc9, 0x25, 0xa3, 0xfd, 0xfd, 0xde, 0xbe, 0xf7, 0xf6, 0xf7, 0xfe, 0xc4, 0xd0, 0x89,
	0xe2, 0x20, 0x0a, 0x31, 0x16, 0xc3, 0x24, 0x65, 0x82, 0x91, 0x86, 0x3f, 0xc7, 0x54, 0xd8, 0x87,
	0x53, 0xc6, 0xa6, 0x73, 0x74, 0xfd, 0x24, 0x72, 0xfd, 0x38, 0x66, 0xc2, 0x17, 0x11, 0x8b, 0xb9,
	0x32, 0xb2, 0x1f, 0xe5, 0x70, 0x28, 0x2e, 0x13, 0xe4, 0xae, 0xfc, 0xab, 0x70, 0xe7, 0xaf, 0x26,
	0xb4, 0xbe, 0x2a, 0xfc, 0x11, 0x02, 0xf5, 0xe4, 0x3c, 0x0a, 0xa9, 0xd1, 0x33, 0xfa, 0x6d, 0x4f,
	0x7e, 0x93, 0x01, 0x90, 0x59, 0x36, 0xc1, 0x34, 0x46, 0x81, 0x7c, 0x1c, 0xcc, 0x33, 0x2e, 0x30,
	0xa5, 0x35, 0x69, 0xf1, 0x60, 0xc5, 0xbc, 0x52, 0x04, 0x39, 0x81, 0x6e, 0xc5, 0x3c, 0xf6, 0x17,
	0xc8, 0x13, 0x3f, 0x40, 0x6a, 0xca, 0x0b, 0x0f, 0x57, 0xdc, 0x6b, 0x4d, 0x91, 0x17, 0x70, 0x50,
	0xb9, 0xc2, 0x26, 0x3f, 0x60, 0x20, 0xbe, 0xbd, 0x4c, 0x90, 0xd6, 0xe5, 0x9d, 0x8a, 0xbf, 0x6f,
	0x4a, 0x6e, 0xeb, 0xa5, 0xdc, 0x25, 0x6d, 0x6c, 0xbf, 0x94, 0x73, 0xe4, 0x29, 0xec, 0x57, 0x93,
	0x63, 0x21, 0xd2, 0xa6, 0x34, 0xef, 0x54, 0xf2, 0x62, 0x21, 0x92, 0x8f, 0xa0, 0x82, 0x8c, 0x13,
	0x16, 0xd2, 0x7b, 0xd2, 0xee, 0xfe, 0x0a, 0x7d, 0xc3, 0x42, 0xf2, 0x04, 0x94, 0xf6, 0xb4, 0xd5,
	0x33, 0xfa, 0xd6, 0xc8, 0x1a, 0x2a, 0x81, 0x87, 0x67, 0x51, 0xe8, 0x29, 0x86, 0x1c, 0x81, 0x15,
	0x05, 0x51, 0x3c, 0xf5, 0xc7, 0xe7, 0x8c, 0x0b, 0xda, 0x96, 0x6e, 0x40, 0x41, 0x5f, 0x32, 0x2e,
	0xf2, 0x50, 0x85, 0x01, 0xc7, 0x74, 0x19, 0x05, 0x48, 0x41, 0x85, 0x52, 0xe8, 0xa9, 0x02, 0xf3,
	0xd2, 0xe4, 0xbe, 0xa9, 0xa5, 0x4a, 0x93, 0x7f, 0x93, 0x2e, 0x34, 0xb8, 0xf0, 0x05, 0xd2, 0x3d,
	0x09, 0xaa, 0x03, 0x39, 0x82, 0x7a, 0xc6, 0x31, 0xa5, 0xf7, 0xaf, 0xe6, 0x24, 0x09, 0x62, 0x43,
	0x2b, 0xc5, 0x84, 0xa5, 0x02, 0x43, 0xda, 0xe9, 0x19, 0x7d, 0xd3, 0x2b, 0xcf, 0xc4, 0x81, 0x3d,
	0x3f, 0x98, 0xc5, 0xec, 0x62, 0x8e, 0xe1, 0x14, 0x43, 0xba, 0x2f, 0xf9, 0x35, 0x8c, 0x1c, 0x42,
	0x3b, 0xc5, 0x80, 0x2d, 0x31, 0xc5, 0x90, 0xbe, 0x23, 0x0d, 0x56, 0x00, 0xf9, 0x00, 0x8a, 0xcc,
	0x2f, 0x70, 0x32, 0xce, 0xd2, 0x39, 0x7d, 0x20, 0x93, 0xdb, 0x2b, 0xc1, 0xb3, 0x74, 0x4e, 0x06,
	0xd0, 0xc4, 0x25, 0xc6, 0x82, 0x53, 0xd2, 0x33, 0xfb, 0xd6, 0xe8, 0x60, 0x28, 0xd5, 0x1a, 0xea,
	0x4e, 0x1c, 0x7e, 0x91, 0xb3, 0x5e, 0x61, 0x64, 0xff, 0x6e, 0x40, 0x43, 0x22, 0xa5, 0x0c, 0xc6,
	0x36, 0x19, 0x6a, 0x55, 0x19, 0xaa, 0xaf, 0x34, 0x37, 0x5e, 0xa9, 0x25, 0xaa, 0xef, 0x92, 0x88,
	0xc2, 0xbd, 0x80, 0x2d, 0x16, 0x18, 0x8b, 0xa2, 0x9f, 0xf4, 0xd1, 0xf9, 0xcf, 0x80, 0x87, 0x3a,
	0xcb, 0xaf, 0x23, 0x2e, 0x3c, 0xfc, 0x31, 0x43, 0x2e, 0x76, 0x8c, 0x89, 0x71, 0xdb, 0x31, 0xa9,
	0xdd, 0x61, 0x4c, 0xcc, 0xbb, 0x8c, 0x49, 0xfd, 0x9a, 0x31, 0x29, 0x05, 0x6d, 0xf4, 0xcc, 0x52,
	0x50, 0x67, 0x01, 0xdd, 0xf5, 0x87, 0xf3, 0x84, 0xc5, 0x1c, 0xc9, 0xc7, 0xd0, 0xcc, 0x0d, 0x32,
	0x2e, 0x5f, 0x6b, 0x8d, 0x3a, 0x5a, 0xce, 0x53, 0x89, 0x7a, 0x05, 0x4b, 0x06, 0xd0, 0xd6, 0x8b,
	0x8b, 0xd3, 0x9a, 0x2c, 0xfb, 0xfe, 0x46, 0xd9, 0xbd, 0x95, 0x85, 0x33, 0x80, 0x77, 0x35, 0xfc,
	0x39, 0xf2, 0x20, 0x8d, 0x26, 0xa8, 0xb5, 0xde, 0xb2, 0xa6, 0x1c, 0x06, 0xf4, 0xaa, 0xf9, 0x2d,
	0x33, 0x7c, 0x06, 0x2d, 0x1d, 0x5f, 0x16, 0x62, 0x4b, 0x82, 0xa5, 0x81, 0xf3, 0xa7, 0x01, 0x07,
	0x1a, 0x7e, 0xcd, 0x44, 0xf4, 0xfd, 0xa5, 0x4e, 0xef, 0x3d, 0x00, 0x79, 0x6b, 0x5c, 0x49, 0xb2,
	0x2d, 0x91, 0x37, 0xf9, 0x42, 0x7d, 0x0c, 0xed, 0x7c, 0x15, 0xc8, 0xa2, 0x17, 0xf5, 0x6e, 0xe5,
	0x80, 0x94, 0x5e, 0xf7, 0xb7, 0xb9, 0xad, 0xbf, 0xeb, 0xd5, 0xfe, 0x7e, 0x04, 0x4d, 0x96, 0x89,
	0x24, 0xd3, 0x1d, 0x5a, 0x9c, 0xa4, 0x87, 0x68, 0xa1, 0x17, 0x9b, 0xfc, 0x76, 0xe6, 0x60, 0xeb,
	0x54, 0xe5, 0x18, 0xbd, 0x4a, 0xd1, 0x17, 0xd7, 0xc9, 0x59, 0x1d, 0x80, 0xda, 0xda, 0x00, 0x90,
	0x1e, 0x58, 0x95, 0x6d, 0x20, 0x13, 0x6d, 0x79, 0x55, 0x68, 0xf4, 0xaf, 0x09, 0x6d, 0x1d, 0x8e,
	0x93, 0x19, 0xd4, 0xf3, 0x76, 0x21, 0xf6, 0x86, 0x94, 0x95, 0xe1, 0xb1, 0x1f, 0x6f, 0xe5, 0x54,
	0xf5, 0x9c, 0xe3, 0x5f, 0xff, 0xfe, 0xe7, 0x8f, 0xda, 0x87, 0xc4, 0x71, 0xfd, 0x24, 0xe1, 0x01,
	0x0b, 0x8b, 0x7f, 0x71, 0xf9, 0x0d, 0x77, 0xf9, 0x7c, 0x78, 0xe2, 0x96, 0x4d, 0x43, 0x7e, 0x81,
	0x96, 0xae, 0x3e, 0x79, 0x7f, 0xc3, 0xe9, 0x46, 0x17, 0xd9, 0x47, 0x3b, 0xf9, 0x22, 0xf0, 0x89,
	0x0c, 0xfc, 0x8c, 0x7c, 0xf2, 0xf6, 0xc0, 0xee, 0x4f, 0xb9, 0x6a, 0x3f, 0x93, 0x19, 0x34, 0x55,
	0x2f, 0x90, 0xc3, 0x0d, 0xef, 0x6b, 0x2d, 0x62, 0x77, 0x75, 0x07, 0x7e, 0xc7, 0xa2, 0xb0, 0x0c,
	0x38, 0x90, 0x01, 0x9f, 0x3a, 0x37, 0x78, 0xe9, 0x4b, 0xe3, 0x98, 0xfc, 0x66, 0x80, 0xa5, 0x2a,
	0xa9, 0x76, 0xe3, 0x93, 0x8d, 0x90, 0x57, 0x4b, 0xbd, 0x23, 0xee, 0x67, 0x32, 0xee, 0xa7, 0xce,
	0xf3, 0x1b, 0x3f, 0xd4, 0x55, 0x8b, 0xf9, 0xa5, 0x71, 0x3c, 0x69, 0xca, 0xdf, 0x11, 0x2f, 0xfe,
	0x0f, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xab, 0xd3, 0x6d, 0x96, 0x08, 0x00, 0x00,
}
