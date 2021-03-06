// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: cart/v1/cart.proto

package v1

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

// CartClient is the client API for Cart service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartClient interface {
	CreateCart(ctx context.Context, in *CreateCartRequest, opts ...grpc.CallOption) (*CartInfoReply, error)
	UpdateCart(ctx context.Context, in *UpdateCartRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	DeleteCart(ctx context.Context, in *DeleteCartRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	ListCart(ctx context.Context, in *ListCartRequest, opts ...grpc.CallOption) (*CartListReply, error)
}

type cartClient struct {
	cc grpc.ClientConnInterface
}

func NewCartClient(cc grpc.ClientConnInterface) CartClient {
	return &cartClient{cc}
}

func (c *cartClient) CreateCart(ctx context.Context, in *CreateCartRequest, opts ...grpc.CallOption) (*CartInfoReply, error) {
	out := new(CartInfoReply)
	err := c.cc.Invoke(ctx, "/cart.v1.Cart/CreateCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) UpdateCart(ctx context.Context, in *UpdateCartRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/cart.v1.Cart/UpdateCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) DeleteCart(ctx context.Context, in *DeleteCartRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/cart.v1.Cart/DeleteCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) ListCart(ctx context.Context, in *ListCartRequest, opts ...grpc.CallOption) (*CartListReply, error) {
	out := new(CartListReply)
	err := c.cc.Invoke(ctx, "/cart.v1.Cart/ListCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServer is the server API for Cart service.
// All implementations must embed UnimplementedCartServer
// for forward compatibility
type CartServer interface {
	CreateCart(context.Context, *CreateCartRequest) (*CartInfoReply, error)
	UpdateCart(context.Context, *UpdateCartRequest) (*CheckResponse, error)
	DeleteCart(context.Context, *DeleteCartRequest) (*CheckResponse, error)
	ListCart(context.Context, *ListCartRequest) (*CartListReply, error)
	mustEmbedUnimplementedCartServer()
}

// UnimplementedCartServer must be embedded to have forward compatible implementations.
type UnimplementedCartServer struct {
}

func (UnimplementedCartServer) CreateCart(context.Context, *CreateCartRequest) (*CartInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCart not implemented")
}
func (UnimplementedCartServer) UpdateCart(context.Context, *UpdateCartRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCart not implemented")
}
func (UnimplementedCartServer) DeleteCart(context.Context, *DeleteCartRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCart not implemented")
}
func (UnimplementedCartServer) ListCart(context.Context, *ListCartRequest) (*CartListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCart not implemented")
}
func (UnimplementedCartServer) mustEmbedUnimplementedCartServer() {}

// UnsafeCartServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServer will
// result in compilation errors.
type UnsafeCartServer interface {
	mustEmbedUnimplementedCartServer()
}

func RegisterCartServer(s grpc.ServiceRegistrar, srv CartServer) {
	s.RegisterService(&Cart_ServiceDesc, srv)
}

func _Cart_CreateCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).CreateCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.v1.Cart/CreateCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).CreateCart(ctx, req.(*CreateCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_UpdateCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).UpdateCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.v1.Cart/UpdateCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).UpdateCart(ctx, req.(*UpdateCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_DeleteCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).DeleteCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.v1.Cart/DeleteCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).DeleteCart(ctx, req.(*DeleteCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_ListCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).ListCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.v1.Cart/ListCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).ListCart(ctx, req.(*ListCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cart_ServiceDesc is the grpc.ServiceDesc for Cart service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cart_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cart.v1.Cart",
	HandlerType: (*CartServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCart",
			Handler:    _Cart_CreateCart_Handler,
		},
		{
			MethodName: "UpdateCart",
			Handler:    _Cart_UpdateCart_Handler,
		},
		{
			MethodName: "DeleteCart",
			Handler:    _Cart_DeleteCart_Handler,
		},
		{
			MethodName: "ListCart",
			Handler:    _Cart_ListCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cart/v1/cart.proto",
}
