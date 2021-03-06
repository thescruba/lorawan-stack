// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/client_services.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("lorawan-stack/api/client_services.proto", fileDescriptor_80815ba053239a77)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/client_services.proto", fileDescriptor_80815ba053239a77)
}

var fileDescriptor_80815ba053239a77 = []byte{
	// 708 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x3f, 0x4c, 0x14, 0x4d,
	0x14, 0x9f, 0xe1, 0xfb, 0xbe, 0xcb, 0x97, 0x8d, 0x81, 0x38, 0x31, 0x98, 0x1c, 0xf0, 0x94, 0xc5,
	0x84, 0x84, 0xc0, 0xac, 0x81, 0x98, 0x18, 0x3b, 0x45, 0x45, 0xa3, 0x8d, 0x10, 0x9b, 0x6b, 0xc8,
	0xde, 0x31, 0xec, 0x4d, 0xee, 0xdc, 0x39, 0x76, 0x06, 0x08, 0x12, 0x12, 0x62, 0xa1, 0x34, 0x26,
	0x12, 0x1b, 0xed, 0x8c, 0x15, 0x25, 0x25, 0x25, 0x25, 0x9d, 0x24, 0x34, 0x94, 0xdc, 0xac, 0x05,
	0x25, 0x9d, 0x94, 0x66, 0x67, 0x77, 0x75, 0xef, 0x9f, 0x70, 0x89, 0xdd, 0xde, 0x7b, 0xbf, 0xf9,
	0xbd, 0x79, 0xbf, 0xf7, 0x7b, 0x73, 0xd6, 0x68, 0x55, 0x04, 0xee, 0xaa, 0xeb, 0x4f, 0x48, 0xe5,
	0x96, 0x2a, 0x8e, 0x5b, 0xe3, 0x4e, 0xa9, 0xca, 0x99, 0xaf, 0xe6, 0x25, 0x0b, 0x56, 0x78, 0x89,
	0x49, 0x5a, 0x0b, 0x84, 0x12, 0xa4, 0x57, 0x29, 0x9f, 0x26, 0x60, 0xba, 0x32, 0x95, 0x1f, 0xf4,
	0x84, 0xf0, 0xaa, 0xcc, 0x9c, 0x70, 0x7d, 0x5f, 0x28, 0x57, 0x71, 0xe1, 0x27, 0xe8, 0xfc, 0x40,
	0x92, 0x35, 0xbf, 0x8a, 0xcb, 0x8b, 0x0e, 0x7b, 0x55, 0x53, 0x6b, 0x49, 0x12, 0x3a, 0xd5, 0x4c,
	0xf2, 0x23, 0xad, 0x79, 0xbe, 0xc0, 0x7c, 0xc5, 0x17, 0x39, 0x0b, 0x64, 0x67, 0x92, 0x80, 0x7b,
	0x65, 0x95, 0xe4, 0x27, 0xcf, 0xff, 0xb3, 0x7a, 0xa7, 0x0d, 0xeb, 0x2c, 0xf3, 0xb8, 0x54, 0xc1,
	0x1a, 0xf9, 0x86, 0xad, 0xdc, 0x74, 0xc0, 0x5c, 0xc5, 0xc8, 0x08, 0x6d, 0x6c, 0x87, 0xc6, 0xf1,
	0xf4, 0xc0, 0xd2, 0x32, 0x93, 0x2a, 0xdf, 0xdf, 0x02, 0x32, 0x69, 0xfb, 0x1d, 0x7e, 0x73, 0xf4,
	0xfd, 0x63, 0xcf, 0x26, 0xb6, 0xa9, 0xb3, 0x2c, 0x59, 0x20, 0x9d, 0xf5, 0x92, 0xa8, 0x56, 0xdd,
	0xa2, 0x08, 0x5c, 0x25, 0x02, 0x1a, 0xc5, 0xe6, 0xf9, 0x82, 0x4c, 0x3f, 0x36, 0x92, 0xf6, 0xe4,
	0x3d, 0x3c, 0x56, 0x78, 0x66, 0x3f, 0x76, 0x44, 0xe0, 0xb9, 0x3e, 0x7f, 0x1d, 0x2b, 0xd6, 0x74,
	0x38, 0x9b, 0x33, 0x24, 0x4d, 0x81, 0x2c, 0x19, 0x29, 0x5b, 0xff, 0xcc, 0x30, 0x45, 0x6e, 0x36,
	0x5f, 0x74, 0x86, 0xa9, 0xcb, 0xb5, 0x32, 0x6a, 0x3a, 0x19, 0x26, 0x37, 0x52, 0x56, 0x67, 0x3d,
	0x19, 0x7f, 0x54, 0xfa, 0xd7, 0xe7, 0x06, 0x39, 0xc2, 0xd6, 0xbf, 0xcf, 0xb9, 0x54, 0xc4, 0x6e,
	0x66, 0x8a, 0xa2, 0x31, 0x9b, 0x4c, 0xab, 0x5d, 0x6f, 0x5f, 0x4d, 0xda, 0xef, 0x63, 0xe5, 0xde,
	0x62, 0xf2, 0x7f, 0x5a, 0xb0, 0x70, 0x9b, 0x74, 0xa9, 0x62, 0xe1, 0x09, 0xf9, 0x4b, 0x12, 0x92,
	0x25, 0x2b, 0xf7, 0xb2, 0xb6, 0xd0, 0xd6, 0x10, 0x71, 0xfc, 0x72, 0x2a, 0x8e, 0x99, 0xae, 0x6e,
	0xe5, 0x5b, 0x54, 0xa4, 0x8d, 0x2a, 0x46, 0x23, 0x73, 0xad, 0xdc, 0x43, 0x56, 0x65, 0x8a, 0x91,
	0xe1, 0xf6, 0x6c, 0x4f, 0x7f, 0x5b, 0x3d, 0xdf, 0x4f, 0xe3, 0x3d, 0xa2, 0xe9, 0x1e, 0xd1, 0x47,
	0xd1, 0x1e, 0xd9, 0x83, 0xa6, 0x60, 0xff, 0xd8, 0xb5, 0x36, 0x63, 0xdb, 0x98, 0xdc, 0xce, 0x59,
	0x57, 0x62, 0xae, 0xfb, 0xa5, 0x12, 0x93, 0x92, 0x54, 0x2d, 0x2b, 0x9a, 0xd2, 0xac, 0xd9, 0x8f,
	0xcb, 0xd5, 0x6d, 0x82, 0xc4, 0x47, 0xed, 0x11, 0x53, 0x77, 0x88, 0x0c, 0xb4, 0xab, 0x9b, 0xec,
	0x1f, 0xd1, 0x3d, 0x56, 0x5f, 0x64, 0xc0, 0xcc, 0x48, 0xc8, 0x78, 0x47, 0x87, 0x66, 0x61, 0xa9,
	0xce, 0xa3, 0xed, 0xd0, 0x0d, 0x38, 0x59, 0x13, 0xbe, 0x64, 0xf6, 0x8f, 0xd8, 0x4f, 0x67, 0x98,
	0x8c, 0x5f, 0x60, 0x60, 0x27, 0xeb, 0x90, 0xc2, 0x1c, 0x79, 0xd1, 0x0d, 0xde, 0xf8, 0xf3, 0x22,
	0x7b, 0x16, 0x2a, 0x84, 0x77, 0x45, 0x9a, 0x75, 0x65, 0xb7, 0x0e, 0x26, 0xdb, 0xd8, 0xea, 0x9b,
	0xbb, 0x48, 0xe4, 0xb9, 0x3f, 0x89, 0xdc, 0xc9, 0x5b, 0x77, 0x8d, 0xa4, 0x93, 0xf9, 0x89, 0x6e,
	0x9a, 0x31, 0xaf, 0xd1, 0x67, 0x6c, 0x5d, 0x35, 0xaf, 0x41, 0x36, 0x41, 0x68, 0xe7, 0x07, 0xa3,
	0x01, 0x98, 0xde, 0x6b, 0xa8, 0xc5, 0x9e, 0x59, 0x94, 0x7d, 0xc7, 0x5c, 0xcf, 0x21, 0xdd, 0x5d,
	0xef, 0xc1, 0x57, 0x7c, 0x50, 0x07, 0x7c, 0x58, 0x07, 0x7c, 0x5c, 0x07, 0x74, 0x52, 0x07, 0x74,
	0x5a, 0x07, 0x74, 0x56, 0x07, 0x74, 0x5e, 0x07, 0xbc, 0xa9, 0x01, 0x6f, 0x69, 0x40, 0x3b, 0x1a,
	0xf0, 0xae, 0x06, 0xb4, 0xa7, 0x01, 0xed, 0x6b, 0x40, 0x07, 0x1a, 0xf0, 0xa1, 0x06, 0x7c, 0xac,
	0x01, 0x9d, 0x68, 0xc0, 0xa7, 0x1a, 0xd0, 0x99, 0x06, 0x7c, 0xae, 0x01, 0x6d, 0x86, 0x80, 0xb6,
	0x42, 0xc0, 0x1f, 0x42, 0x40, 0x9f, 0x42, 0xc0, 0x5f, 0x42, 0x40, 0x3b, 0x21, 0xa0, 0xdd, 0x10,
	0xf0, 0x5e, 0x08, 0x78, 0x3f, 0x04, 0x5c, 0x18, 0xf7, 0x04, 0x55, 0x65, 0xa6, 0xca, 0xdc, 0xf7,
	0x24, 0xf5, 0x99, 0x5a, 0x15, 0x41, 0xc5, 0x69, 0xfc, 0xeb, 0xaa, 0x55, 0x3c, 0x47, 0x29, 0xbf,
	0x56, 0x2c, 0xe6, 0xcc, 0x28, 0xa6, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xfe, 0xaf, 0x34,
	0x95, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientRegistryClient is the client API for ClientRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientRegistryClient interface {
	// Create a new OAuth client. This also sets the given organization or user as
	// first collaborator with all possible rights.
	Create(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*Client, error)
	// Get the OAuth client with the given identifiers, selecting the fields given
	// by the field mask. The method may return more or less fields, depending on
	// the rights of the caller.
	Get(ctx context.Context, in *GetClientRequest, opts ...grpc.CallOption) (*Client, error)
	// List OAuth clients. See request message for details.
	List(ctx context.Context, in *ListClientsRequest, opts ...grpc.CallOption) (*Clients, error)
	Update(ctx context.Context, in *UpdateClientRequest, opts ...grpc.CallOption) (*Client, error)
	Delete(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
}

type clientRegistryClient struct {
	cc *grpc.ClientConn
}

func NewClientRegistryClient(cc *grpc.ClientConn) ClientRegistryClient {
	return &clientRegistryClient{cc}
}

func (c *clientRegistryClient) Create(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientRegistry/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientRegistryClient) Get(ctx context.Context, in *GetClientRequest, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientRegistry/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientRegistryClient) List(ctx context.Context, in *ListClientsRequest, opts ...grpc.CallOption) (*Clients, error) {
	out := new(Clients)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientRegistry/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientRegistryClient) Update(ctx context.Context, in *UpdateClientRequest, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientRegistry/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientRegistryClient) Delete(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientRegistry/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientRegistryServer is the server API for ClientRegistry service.
type ClientRegistryServer interface {
	// Create a new OAuth client. This also sets the given organization or user as
	// first collaborator with all possible rights.
	Create(context.Context, *CreateClientRequest) (*Client, error)
	// Get the OAuth client with the given identifiers, selecting the fields given
	// by the field mask. The method may return more or less fields, depending on
	// the rights of the caller.
	Get(context.Context, *GetClientRequest) (*Client, error)
	// List OAuth clients. See request message for details.
	List(context.Context, *ListClientsRequest) (*Clients, error)
	Update(context.Context, *UpdateClientRequest) (*Client, error)
	Delete(context.Context, *ClientIdentifiers) (*types.Empty, error)
}

// UnimplementedClientRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedClientRegistryServer struct {
}

func (*UnimplementedClientRegistryServer) Create(ctx context.Context, req *CreateClientRequest) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedClientRegistryServer) Get(ctx context.Context, req *GetClientRequest) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedClientRegistryServer) List(ctx context.Context, req *ListClientsRequest) (*Clients, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedClientRegistryServer) Update(ctx context.Context, req *UpdateClientRequest) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedClientRegistryServer) Delete(ctx context.Context, req *ClientIdentifiers) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterClientRegistryServer(s *grpc.Server, srv ClientRegistryServer) {
	s.RegisterService(&_ClientRegistry_serviceDesc, srv)
}

func _ClientRegistry_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientRegistry/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).Create(ctx, req.(*CreateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientRegistry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientRegistry/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).Get(ctx, req.(*GetClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientRegistry_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClientsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientRegistry/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).List(ctx, req.(*ListClientsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientRegistry_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientRegistry/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).Update(ctx, req.(*UpdateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientRegistry/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).Delete(ctx, req.(*ClientIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.ClientRegistry",
	HandlerType: (*ClientRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ClientRegistry_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ClientRegistry_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ClientRegistry_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ClientRegistry_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ClientRegistry_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/client_services.proto",
}

// ClientAccessClient is the client API for ClientAccess service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientAccessClient interface {
	ListRights(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*Rights, error)
	// Get the rights of a collaborator (member) of the client.
	// Pseudo-rights in the response (such as the "_ALL" right) are not expanded.
	GetCollaborator(ctx context.Context, in *GetClientCollaboratorRequest, opts ...grpc.CallOption) (*GetCollaboratorResponse, error)
	// Set the rights of a collaborator (member) on the client. It is required for the caller to
	// have all assigned or/and removed rights.
	// Setting a collaborator without rights, removes them.
	SetCollaborator(ctx context.Context, in *SetClientCollaboratorRequest, opts ...grpc.CallOption) (*types.Empty, error)
	ListCollaborators(ctx context.Context, in *ListClientCollaboratorsRequest, opts ...grpc.CallOption) (*Collaborators, error)
}

type clientAccessClient struct {
	cc *grpc.ClientConn
}

func NewClientAccessClient(cc *grpc.ClientConn) ClientAccessClient {
	return &clientAccessClient{cc}
}

func (c *clientAccessClient) ListRights(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*Rights, error) {
	out := new(Rights)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientAccess/ListRights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientAccessClient) GetCollaborator(ctx context.Context, in *GetClientCollaboratorRequest, opts ...grpc.CallOption) (*GetCollaboratorResponse, error) {
	out := new(GetCollaboratorResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientAccess/GetCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientAccessClient) SetCollaborator(ctx context.Context, in *SetClientCollaboratorRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientAccess/SetCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientAccessClient) ListCollaborators(ctx context.Context, in *ListClientCollaboratorsRequest, opts ...grpc.CallOption) (*Collaborators, error) {
	out := new(Collaborators)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientAccess/ListCollaborators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientAccessServer is the server API for ClientAccess service.
type ClientAccessServer interface {
	ListRights(context.Context, *ClientIdentifiers) (*Rights, error)
	// Get the rights of a collaborator (member) of the client.
	// Pseudo-rights in the response (such as the "_ALL" right) are not expanded.
	GetCollaborator(context.Context, *GetClientCollaboratorRequest) (*GetCollaboratorResponse, error)
	// Set the rights of a collaborator (member) on the client. It is required for the caller to
	// have all assigned or/and removed rights.
	// Setting a collaborator without rights, removes them.
	SetCollaborator(context.Context, *SetClientCollaboratorRequest) (*types.Empty, error)
	ListCollaborators(context.Context, *ListClientCollaboratorsRequest) (*Collaborators, error)
}

// UnimplementedClientAccessServer can be embedded to have forward compatible implementations.
type UnimplementedClientAccessServer struct {
}

func (*UnimplementedClientAccessServer) ListRights(ctx context.Context, req *ClientIdentifiers) (*Rights, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRights not implemented")
}
func (*UnimplementedClientAccessServer) GetCollaborator(ctx context.Context, req *GetClientCollaboratorRequest) (*GetCollaboratorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollaborator not implemented")
}
func (*UnimplementedClientAccessServer) SetCollaborator(ctx context.Context, req *SetClientCollaboratorRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCollaborator not implemented")
}
func (*UnimplementedClientAccessServer) ListCollaborators(ctx context.Context, req *ListClientCollaboratorsRequest) (*Collaborators, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollaborators not implemented")
}

func RegisterClientAccessServer(s *grpc.Server, srv ClientAccessServer) {
	s.RegisterService(&_ClientAccess_serviceDesc, srv)
}

func _ClientAccess_ListRights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAccessServer).ListRights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientAccess/ListRights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAccessServer).ListRights(ctx, req.(*ClientIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientAccess_GetCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAccessServer).GetCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientAccess/GetCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAccessServer).GetCollaborator(ctx, req.(*GetClientCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientAccess_SetCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetClientCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAccessServer).SetCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientAccess/SetCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAccessServer).SetCollaborator(ctx, req.(*SetClientCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientAccess_ListCollaborators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClientCollaboratorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAccessServer).ListCollaborators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientAccess/ListCollaborators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAccessServer).ListCollaborators(ctx, req.(*ListClientCollaboratorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientAccess_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.ClientAccess",
	HandlerType: (*ClientAccessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRights",
			Handler:    _ClientAccess_ListRights_Handler,
		},
		{
			MethodName: "GetCollaborator",
			Handler:    _ClientAccess_GetCollaborator_Handler,
		},
		{
			MethodName: "SetCollaborator",
			Handler:    _ClientAccess_SetCollaborator_Handler,
		},
		{
			MethodName: "ListCollaborators",
			Handler:    _ClientAccess_ListCollaborators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/client_services.proto",
}
