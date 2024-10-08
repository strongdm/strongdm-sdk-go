// Copyright 2020 StrongDM Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ProxyClusterKeysClient is the client API for ProxyClusterKeys service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProxyClusterKeysClient interface {
	// Create registers a new ProxyClusterKey.
	Create(ctx context.Context, in *ProxyClusterKeyCreateRequest, opts ...grpc.CallOption) (*ProxyClusterKeyCreateResponse, error)
	// Get reads one ProxyClusterKey by ID.
	Get(ctx context.Context, in *ProxyClusterKeyGetRequest, opts ...grpc.CallOption) (*ProxyClusterKeyGetResponse, error)
	// Delete removes a ProxyClusterKey by ID.
	Delete(ctx context.Context, in *ProxyClusterKeyDeleteRequest, opts ...grpc.CallOption) (*ProxyClusterKeyDeleteResponse, error)
	// List gets a list of ProxyClusterKeys matching a given set of criteria.
	List(ctx context.Context, in *ProxyClusterKeyListRequest, opts ...grpc.CallOption) (*ProxyClusterKeyListResponse, error)
}

type proxyClusterKeysClient struct {
	cc grpc.ClientConnInterface
}

func NewProxyClusterKeysClient(cc grpc.ClientConnInterface) ProxyClusterKeysClient {
	return &proxyClusterKeysClient{cc}
}

func (c *proxyClusterKeysClient) Create(ctx context.Context, in *ProxyClusterKeyCreateRequest, opts ...grpc.CallOption) (*ProxyClusterKeyCreateResponse, error) {
	out := new(ProxyClusterKeyCreateResponse)
	err := c.cc.Invoke(ctx, "/v1.ProxyClusterKeys/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClusterKeysClient) Get(ctx context.Context, in *ProxyClusterKeyGetRequest, opts ...grpc.CallOption) (*ProxyClusterKeyGetResponse, error) {
	out := new(ProxyClusterKeyGetResponse)
	err := c.cc.Invoke(ctx, "/v1.ProxyClusterKeys/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClusterKeysClient) Delete(ctx context.Context, in *ProxyClusterKeyDeleteRequest, opts ...grpc.CallOption) (*ProxyClusterKeyDeleteResponse, error) {
	out := new(ProxyClusterKeyDeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.ProxyClusterKeys/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClusterKeysClient) List(ctx context.Context, in *ProxyClusterKeyListRequest, opts ...grpc.CallOption) (*ProxyClusterKeyListResponse, error) {
	out := new(ProxyClusterKeyListResponse)
	err := c.cc.Invoke(ctx, "/v1.ProxyClusterKeys/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProxyClusterKeysServer is the server API for ProxyClusterKeys service.
// All implementations must embed UnimplementedProxyClusterKeysServer
// for forward compatibility
type ProxyClusterKeysServer interface {
	// Create registers a new ProxyClusterKey.
	Create(context.Context, *ProxyClusterKeyCreateRequest) (*ProxyClusterKeyCreateResponse, error)
	// Get reads one ProxyClusterKey by ID.
	Get(context.Context, *ProxyClusterKeyGetRequest) (*ProxyClusterKeyGetResponse, error)
	// Delete removes a ProxyClusterKey by ID.
	Delete(context.Context, *ProxyClusterKeyDeleteRequest) (*ProxyClusterKeyDeleteResponse, error)
	// List gets a list of ProxyClusterKeys matching a given set of criteria.
	List(context.Context, *ProxyClusterKeyListRequest) (*ProxyClusterKeyListResponse, error)
	mustEmbedUnimplementedProxyClusterKeysServer()
}

// UnimplementedProxyClusterKeysServer must be embedded to have forward compatible implementations.
type UnimplementedProxyClusterKeysServer struct {
}

func (UnimplementedProxyClusterKeysServer) Create(context.Context, *ProxyClusterKeyCreateRequest) (*ProxyClusterKeyCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProxyClusterKeysServer) Get(context.Context, *ProxyClusterKeyGetRequest) (*ProxyClusterKeyGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedProxyClusterKeysServer) Delete(context.Context, *ProxyClusterKeyDeleteRequest) (*ProxyClusterKeyDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedProxyClusterKeysServer) List(context.Context, *ProxyClusterKeyListRequest) (*ProxyClusterKeyListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedProxyClusterKeysServer) mustEmbedUnimplementedProxyClusterKeysServer() {}

// UnsafeProxyClusterKeysServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProxyClusterKeysServer will
// result in compilation errors.
type UnsafeProxyClusterKeysServer interface {
	mustEmbedUnimplementedProxyClusterKeysServer()
}

func RegisterProxyClusterKeysServer(s grpc.ServiceRegistrar, srv ProxyClusterKeysServer) {
	s.RegisterService(&_ProxyClusterKeys_serviceDesc, srv)
}

func _ProxyClusterKeys_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProxyClusterKeyCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyClusterKeysServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ProxyClusterKeys/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyClusterKeysServer).Create(ctx, req.(*ProxyClusterKeyCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProxyClusterKeys_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProxyClusterKeyGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyClusterKeysServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ProxyClusterKeys/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyClusterKeysServer).Get(ctx, req.(*ProxyClusterKeyGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProxyClusterKeys_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProxyClusterKeyDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyClusterKeysServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ProxyClusterKeys/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyClusterKeysServer).Delete(ctx, req.(*ProxyClusterKeyDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProxyClusterKeys_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProxyClusterKeyListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyClusterKeysServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ProxyClusterKeys/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyClusterKeysServer).List(ctx, req.(*ProxyClusterKeyListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProxyClusterKeys_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ProxyClusterKeys",
	HandlerType: (*ProxyClusterKeysServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProxyClusterKeys_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ProxyClusterKeys_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ProxyClusterKeys_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ProxyClusterKeys_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proxy_cluster_keys.proto",
}
