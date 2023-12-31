// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: food.proto

package food

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

const (
	Food_Search_FullMethodName     = "/food.Food/Search"
	Food_AddFood_FullMethodName    = "/food.Food/AddFood"
	Food_DeleteFood_FullMethodName = "/food.Food/DeleteFood"
	Food_FoodList_FullMethodName   = "/food.Food/FoodList"
)

// FoodClient is the client API for Food service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FoodClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*FoodInfoResponse, error)
	AddFood(ctx context.Context, in *AddFoodRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	DeleteFood(ctx context.Context, in *DeleteFoodRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	FoodList(ctx context.Context, in *FoodListRequest, opts ...grpc.CallOption) (*FoodListResponse, error)
}

type foodClient struct {
	cc grpc.ClientConnInterface
}

func NewFoodClient(cc grpc.ClientConnInterface) FoodClient {
	return &foodClient{cc}
}

func (c *foodClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*FoodInfoResponse, error) {
	out := new(FoodInfoResponse)
	err := c.cc.Invoke(ctx, Food_Search_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodClient) AddFood(ctx context.Context, in *AddFoodRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, Food_AddFood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodClient) DeleteFood(ctx context.Context, in *DeleteFoodRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, Food_DeleteFood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodClient) FoodList(ctx context.Context, in *FoodListRequest, opts ...grpc.CallOption) (*FoodListResponse, error) {
	out := new(FoodListResponse)
	err := c.cc.Invoke(ctx, Food_FoodList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FoodServer is the server API for Food service.
// All implementations must embed UnimplementedFoodServer
// for forward compatibility
type FoodServer interface {
	Search(context.Context, *SearchRequest) (*FoodInfoResponse, error)
	AddFood(context.Context, *AddFoodRequest) (*StatusResponse, error)
	DeleteFood(context.Context, *DeleteFoodRequest) (*StatusResponse, error)
	FoodList(context.Context, *FoodListRequest) (*FoodListResponse, error)
	mustEmbedUnimplementedFoodServer()
}

// UnimplementedFoodServer must be embedded to have forward compatible implementations.
type UnimplementedFoodServer struct {
}

func (UnimplementedFoodServer) Search(context.Context, *SearchRequest) (*FoodInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedFoodServer) AddFood(context.Context, *AddFoodRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFood not implemented")
}
func (UnimplementedFoodServer) DeleteFood(context.Context, *DeleteFoodRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFood not implemented")
}
func (UnimplementedFoodServer) FoodList(context.Context, *FoodListRequest) (*FoodListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FoodList not implemented")
}
func (UnimplementedFoodServer) mustEmbedUnimplementedFoodServer() {}

// UnsafeFoodServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FoodServer will
// result in compilation errors.
type UnsafeFoodServer interface {
	mustEmbedUnimplementedFoodServer()
}

func RegisterFoodServer(s grpc.ServiceRegistrar, srv FoodServer) {
	s.RegisterService(&Food_ServiceDesc, srv)
}

func _Food_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Food_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Food_AddFood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServer).AddFood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Food_AddFood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServer).AddFood(ctx, req.(*AddFoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Food_DeleteFood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServer).DeleteFood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Food_DeleteFood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServer).DeleteFood(ctx, req.(*DeleteFoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Food_FoodList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FoodListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServer).FoodList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Food_FoodList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServer).FoodList(ctx, req.(*FoodListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Food_ServiceDesc is the grpc.ServiceDesc for Food service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Food_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "food.Food",
	HandlerType: (*FoodServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Food_Search_Handler,
		},
		{
			MethodName: "AddFood",
			Handler:    _Food_AddFood_Handler,
		},
		{
			MethodName: "DeleteFood",
			Handler:    _Food_DeleteFood_Handler,
		},
		{
			MethodName: "FoodList",
			Handler:    _Food_FoodList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "food.proto",
}
