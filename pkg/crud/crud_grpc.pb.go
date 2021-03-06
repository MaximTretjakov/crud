// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: crud.proto

package crud

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CRUDClient is the client API for CRUD service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CRUDClient interface {
	// create note
	CreateNote(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Status, error)
	// read note
	ReadNote(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*Response, error)
	// update note
	UpdateNote(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Status, error)
	// delete note
	DeleteNote(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Status, error)
	// produce to kafka
	KafkaSendMessage(ctx context.Context, in *KafkaRequest, opts ...grpc.CallOption) (*KafkaResponse, error)
	// consume from kafka
	KafkaReadMessage(ctx context.Context, in *KafkaRequest, opts ...grpc.CallOption) (*KafkaResponse, error)
}

type cRUDClient struct {
	cc grpc.ClientConnInterface
}

func NewCRUDClient(cc grpc.ClientConnInterface) CRUDClient {
	return &cRUDClient{cc}
}

func (c *cRUDClient) CreateNote(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/api.CRUD/CreateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) ReadNote(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.CRUD/ReadNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) UpdateNote(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/api.CRUD/UpdateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) DeleteNote(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/api.CRUD/DeleteNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) KafkaSendMessage(ctx context.Context, in *KafkaRequest, opts ...grpc.CallOption) (*KafkaResponse, error) {
	out := new(KafkaResponse)
	err := c.cc.Invoke(ctx, "/api.CRUD/KafkaSendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) KafkaReadMessage(ctx context.Context, in *KafkaRequest, opts ...grpc.CallOption) (*KafkaResponse, error) {
	out := new(KafkaResponse)
	err := c.cc.Invoke(ctx, "/api.CRUD/KafkaReadMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CRUDServer is the server API for CRUD service.
// All implementations must embed UnimplementedCRUDServer
// for forward compatibility
type CRUDServer interface {
	// create note
	CreateNote(context.Context, *CreateRequest) (*Status, error)
	// read note
	ReadNote(context.Context, *ReadRequest) (*Response, error)
	// update note
	UpdateNote(context.Context, *UpdateRequest) (*Status, error)
	// delete note
	DeleteNote(context.Context, *DeleteRequest) (*Status, error)
	// produce to kafka
	KafkaSendMessage(context.Context, *KafkaRequest) (*KafkaResponse, error)
	// consume from kafka
	KafkaReadMessage(context.Context, *KafkaRequest) (*KafkaResponse, error)
	mustEmbedUnimplementedCRUDServer()
}

// UnimplementedCRUDServer must be embedded to have forward compatible implementations.
type UnimplementedCRUDServer struct {
}

func (UnimplementedCRUDServer) CreateNote(context.Context, *CreateRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNote not implemented")
}
func (UnimplementedCRUDServer) ReadNote(context.Context, *ReadRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadNote not implemented")
}
func (UnimplementedCRUDServer) UpdateNote(context.Context, *UpdateRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNote not implemented")
}
func (UnimplementedCRUDServer) DeleteNote(context.Context, *DeleteRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNote not implemented")
}
func (UnimplementedCRUDServer) KafkaSendMessage(context.Context, *KafkaRequest) (*KafkaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KafkaSendMessage not implemented")
}
func (UnimplementedCRUDServer) KafkaReadMessage(context.Context, *KafkaRequest) (*KafkaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KafkaReadMessage not implemented")
}
func (UnimplementedCRUDServer) mustEmbedUnimplementedCRUDServer() {}

// UnsafeCRUDServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CRUDServer will
// result in compilation errors.
type UnsafeCRUDServer interface {
	mustEmbedUnimplementedCRUDServer()
}

func RegisterCRUDServer(s grpc.ServiceRegistrar, srv CRUDServer) {
	s.RegisterService(&CRUD_ServiceDesc, srv)
}

func _CRUD_CreateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).CreateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CRUD/CreateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).CreateNote(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_ReadNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).ReadNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CRUD/ReadNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).ReadNote(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_UpdateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).UpdateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CRUD/UpdateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).UpdateNote(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_DeleteNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).DeleteNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CRUD/DeleteNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).DeleteNote(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_KafkaSendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KafkaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).KafkaSendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CRUD/KafkaSendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).KafkaSendMessage(ctx, req.(*KafkaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_KafkaReadMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KafkaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).KafkaReadMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CRUD/KafkaReadMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).KafkaReadMessage(ctx, req.(*KafkaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CRUD_ServiceDesc is the grpc.ServiceDesc for CRUD service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CRUD_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.CRUD",
	HandlerType: (*CRUDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNote",
			Handler:    _CRUD_CreateNote_Handler,
		},
		{
			MethodName: "ReadNote",
			Handler:    _CRUD_ReadNote_Handler,
		},
		{
			MethodName: "UpdateNote",
			Handler:    _CRUD_UpdateNote_Handler,
		},
		{
			MethodName: "DeleteNote",
			Handler:    _CRUD_DeleteNote_Handler,
		},
		{
			MethodName: "KafkaSendMessage",
			Handler:    _CRUD_KafkaSendMessage_Handler,
		},
		{
			MethodName: "KafkaReadMessage",
			Handler:    _CRUD_KafkaReadMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crud.proto",
}
