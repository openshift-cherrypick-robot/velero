// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RestoreItemAction.proto

package generated

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RestoreItemActionExecuteRequest struct {
	Plugin         string `protobuf:"bytes,1,opt,name=plugin" json:"plugin,omitempty"`
	Item           []byte `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	Restore        []byte `protobuf:"bytes,3,opt,name=restore,proto3" json:"restore,omitempty"`
	ItemFromBackup []byte `protobuf:"bytes,4,opt,name=itemFromBackup,proto3" json:"itemFromBackup,omitempty"`
}

func (m *RestoreItemActionExecuteRequest) Reset()                    { *m = RestoreItemActionExecuteRequest{} }
func (m *RestoreItemActionExecuteRequest) String() string            { return proto.CompactTextString(m) }
func (*RestoreItemActionExecuteRequest) ProtoMessage()               {}
func (*RestoreItemActionExecuteRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *RestoreItemActionExecuteRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *RestoreItemActionExecuteRequest) GetItem() []byte {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *RestoreItemActionExecuteRequest) GetRestore() []byte {
	if m != nil {
		return m.Restore
	}
	return nil
}

func (m *RestoreItemActionExecuteRequest) GetItemFromBackup() []byte {
	if m != nil {
		return m.ItemFromBackup
	}
	return nil
}

type RestoreItemActionExecuteResponse struct {
	Item            []byte                `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	AdditionalItems []*ResourceIdentifier `protobuf:"bytes,2,rep,name=additionalItems" json:"additionalItems,omitempty"`
	SkipRestore     bool                  `protobuf:"varint,3,opt,name=skipRestore,proto3" json:"skipRestore,omitempty"`
}

func (m *RestoreItemActionExecuteResponse) Reset()         { *m = RestoreItemActionExecuteResponse{} }
func (m *RestoreItemActionExecuteResponse) String() string { return proto.CompactTextString(m) }
func (*RestoreItemActionExecuteResponse) ProtoMessage()    {}
func (*RestoreItemActionExecuteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{1}
}

func (m *RestoreItemActionExecuteResponse) GetItem() []byte {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *RestoreItemActionExecuteResponse) GetAdditionalItems() []*ResourceIdentifier {
	if m != nil {
		return m.AdditionalItems
	}
	return nil
}

func (m *RestoreItemActionExecuteResponse) GetSkipRestore() bool {
	if m != nil {
		return m.SkipRestore
	}
	return false
}

type RestoreItemActionAppliesToRequest struct {
	Plugin string `protobuf:"bytes,1,opt,name=plugin" json:"plugin,omitempty"`
}

func (m *RestoreItemActionAppliesToRequest) Reset()         { *m = RestoreItemActionAppliesToRequest{} }
func (m *RestoreItemActionAppliesToRequest) String() string { return proto.CompactTextString(m) }
func (*RestoreItemActionAppliesToRequest) ProtoMessage()    {}
func (*RestoreItemActionAppliesToRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{2}
}

func (m *RestoreItemActionAppliesToRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

type RestoreItemActionAppliesToResponse struct {
	ResourceSelector *ResourceSelector `protobuf:"bytes,1,opt,name=ResourceSelector" json:"ResourceSelector,omitempty"`
}

func (m *RestoreItemActionAppliesToResponse) Reset()         { *m = RestoreItemActionAppliesToResponse{} }
func (m *RestoreItemActionAppliesToResponse) String() string { return proto.CompactTextString(m) }
func (*RestoreItemActionAppliesToResponse) ProtoMessage()    {}
func (*RestoreItemActionAppliesToResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{3}
}

func (m *RestoreItemActionAppliesToResponse) GetResourceSelector() *ResourceSelector {
	if m != nil {
		return m.ResourceSelector
	}
	return nil
}

func init() {
	proto.RegisterType((*RestoreItemActionExecuteRequest)(nil), "generated.RestoreItemActionExecuteRequest")
	proto.RegisterType((*RestoreItemActionExecuteResponse)(nil), "generated.RestoreItemActionExecuteResponse")
	proto.RegisterType((*RestoreItemActionAppliesToRequest)(nil), "generated.RestoreItemActionAppliesToRequest")
	proto.RegisterType((*RestoreItemActionAppliesToResponse)(nil), "generated.RestoreItemActionAppliesToResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RestoreItemAction service

type RestoreItemActionClient interface {
	AppliesTo(ctx context.Context, in *RestoreItemActionAppliesToRequest, opts ...grpc.CallOption) (*RestoreItemActionAppliesToResponse, error)
	Execute(ctx context.Context, in *RestoreItemActionExecuteRequest, opts ...grpc.CallOption) (*RestoreItemActionExecuteResponse, error)
}

type restoreItemActionClient struct {
	cc *grpc.ClientConn
}

func NewRestoreItemActionClient(cc *grpc.ClientConn) RestoreItemActionClient {
	return &restoreItemActionClient{cc}
}

func (c *restoreItemActionClient) AppliesTo(ctx context.Context, in *RestoreItemActionAppliesToRequest, opts ...grpc.CallOption) (*RestoreItemActionAppliesToResponse, error) {
	out := new(RestoreItemActionAppliesToResponse)
	err := grpc.Invoke(ctx, "/generated.RestoreItemAction/AppliesTo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restoreItemActionClient) Execute(ctx context.Context, in *RestoreItemActionExecuteRequest, opts ...grpc.CallOption) (*RestoreItemActionExecuteResponse, error) {
	out := new(RestoreItemActionExecuteResponse)
	err := grpc.Invoke(ctx, "/generated.RestoreItemAction/Execute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RestoreItemAction service

type RestoreItemActionServer interface {
	AppliesTo(context.Context, *RestoreItemActionAppliesToRequest) (*RestoreItemActionAppliesToResponse, error)
	Execute(context.Context, *RestoreItemActionExecuteRequest) (*RestoreItemActionExecuteResponse, error)
}

func RegisterRestoreItemActionServer(s *grpc.Server, srv RestoreItemActionServer) {
	s.RegisterService(&_RestoreItemAction_serviceDesc, srv)
}

func _RestoreItemAction_AppliesTo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreItemActionAppliesToRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestoreItemActionServer).AppliesTo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.RestoreItemAction/AppliesTo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestoreItemActionServer).AppliesTo(ctx, req.(*RestoreItemActionAppliesToRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestoreItemAction_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreItemActionExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestoreItemActionServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.RestoreItemAction/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestoreItemActionServer).Execute(ctx, req.(*RestoreItemActionExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RestoreItemAction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "generated.RestoreItemAction",
	HandlerType: (*RestoreItemActionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppliesTo",
			Handler:    _RestoreItemAction_AppliesTo_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _RestoreItemAction_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "RestoreItemAction.proto",
}

func init() { proto.RegisterFile("RestoreItemAction.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4e, 0xf3, 0x30,
	0x10, 0x94, 0xdb, 0xaa, 0x55, 0xb7, 0xd5, 0xf7, 0x81, 0x0f, 0x60, 0x15, 0x21, 0x42, 0x0e, 0xa8,
	0xe2, 0xa7, 0x87, 0x72, 0xe4, 0x54, 0x24, 0xa8, 0x7a, 0x75, 0x79, 0x81, 0x34, 0x59, 0x5a, 0x8b,
	0x24, 0x36, 0xb6, 0x23, 0x71, 0xe3, 0x11, 0x78, 0x46, 0xde, 0x04, 0xe5, 0x57, 0xb4, 0x81, 0xd0,
	0x9b, 0x77, 0x3d, 0xb3, 0x33, 0x63, 0x2f, 0x1c, 0x73, 0x34, 0x56, 0x6a, 0x5c, 0x58, 0x8c, 0x66,
	0xbe, 0x15, 0x32, 0x9e, 0x28, 0x2d, 0xad, 0xa4, 0xfd, 0x35, 0xc6, 0xa8, 0x3d, 0x8b, 0xc1, 0x68,
	0xb8, 0xdc, 0x78, 0x1a, 0x83, 0xfc, 0xc2, 0xfd, 0x20, 0x70, 0x56, 0x23, 0x3d, 0xbc, 0xa1, 0x9f,
	0x58, 0xe4, 0xf8, 0x9a, 0xa0, 0xb1, 0xf4, 0x08, 0xba, 0x2a, 0x4c, 0xd6, 0x22, 0x66, 0xc4, 0x21,
	0xe3, 0x3e, 0x2f, 0x2a, 0x4a, 0xa1, 0x23, 0x2c, 0x46, 0xac, 0xe5, 0x90, 0xf1, 0x90, 0x67, 0x67,
	0xca, 0xa0, 0xa7, 0xf3, 0x71, 0xac, 0x9d, 0xb5, 0xcb, 0x92, 0x5e, 0xc0, 0xbf, 0x14, 0xf1, 0xa8,
	0x65, 0x74, 0xef, 0xf9, 0x2f, 0x89, 0x62, 0x9d, 0x0c, 0xb0, 0xd3, 0x75, 0xdf, 0xc1, 0xf9, 0xdd,
	0x90, 0x51, 0x32, 0x36, 0x58, 0x29, 0x93, 0x6f, 0xca, 0x73, 0xf8, 0xef, 0x05, 0x81, 0x48, 0xe1,
	0x5e, 0x98, 0x52, 0x0d, 0x6b, 0x39, 0xed, 0xf1, 0x60, 0x7a, 0x3a, 0xa9, 0xc2, 0x4f, 0x38, 0x1a,
	0x99, 0x68, 0x1f, 0x17, 0x01, 0xc6, 0x56, 0x3c, 0x0b, 0xd4, 0x7c, 0x97, 0xe5, 0xde, 0xc1, 0x79,
	0xcd, 0xc0, 0x4c, 0xa9, 0x50, 0xa0, 0x79, 0x92, 0x7f, 0xbc, 0x89, 0x1b, 0x81, 0xdb, 0x44, 0x2e,
	0xfc, 0xcf, 0xe1, 0xa0, 0x74, 0xb2, 0xc4, 0x10, 0x7d, 0x2b, 0x75, 0x36, 0x67, 0x30, 0x3d, 0xf9,
	0xc1, 0x6c, 0x09, 0xe1, 0x35, 0xd2, 0xf4, 0x93, 0xc0, 0x61, 0x4d, 0x8f, 0x6e, 0xa0, 0x5f, 0x69,
	0xd2, 0xeb, 0xed, 0x89, 0xcd, 0xb9, 0x46, 0x37, 0x7b, 0xa2, 0x8b, 0x20, 0x2b, 0xe8, 0x15, 0x7f,
	0x43, 0x2f, 0x9b, 0x98, 0xdb, 0x1b, 0x35, 0xba, 0xda, 0x0b, 0x9b, 0x6b, 0xac, 0xba, 0xd9, 0xa6,
	0xde, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x96, 0x3a, 0x66, 0xfa, 0xdd, 0x02, 0x00, 0x00,
}
