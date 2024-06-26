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

// RemoteIdentitiesClient is the client API for RemoteIdentities service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Deprecated: Do not use.
type RemoteIdentitiesClient interface {
	// Create registers a new RemoteIdentity.
	Create(ctx context.Context, in *RemoteIdentityCreateRequest, opts ...grpc.CallOption) (*RemoteIdentityCreateResponse, error)
	// Get reads one RemoteIdentity by ID.
	Get(ctx context.Context, in *RemoteIdentityGetRequest, opts ...grpc.CallOption) (*RemoteIdentityGetResponse, error)
	// Update replaces all the fields of a RemoteIdentity by ID.
	Update(ctx context.Context, in *RemoteIdentityUpdateRequest, opts ...grpc.CallOption) (*RemoteIdentityUpdateResponse, error)
	// Delete removes a RemoteIdentity by ID.
	Delete(ctx context.Context, in *RemoteIdentityDeleteRequest, opts ...grpc.CallOption) (*RemoteIdentityDeleteResponse, error)
	// List gets a list of RemoteIdentities matching a given set of criteria.
	List(ctx context.Context, in *RemoteIdentityListRequest, opts ...grpc.CallOption) (*RemoteIdentityListResponse, error)
}

type remoteIdentitiesClient struct {
	cc grpc.ClientConnInterface
}

// Deprecated: Do not use.
func NewRemoteIdentitiesClient(cc grpc.ClientConnInterface) RemoteIdentitiesClient {
	return &remoteIdentitiesClient{cc}
}

func (c *remoteIdentitiesClient) Create(ctx context.Context, in *RemoteIdentityCreateRequest, opts ...grpc.CallOption) (*RemoteIdentityCreateResponse, error) {
	out := new(RemoteIdentityCreateResponse)
	err := c.cc.Invoke(ctx, "/v1.RemoteIdentities/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteIdentitiesClient) Get(ctx context.Context, in *RemoteIdentityGetRequest, opts ...grpc.CallOption) (*RemoteIdentityGetResponse, error) {
	out := new(RemoteIdentityGetResponse)
	err := c.cc.Invoke(ctx, "/v1.RemoteIdentities/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteIdentitiesClient) Update(ctx context.Context, in *RemoteIdentityUpdateRequest, opts ...grpc.CallOption) (*RemoteIdentityUpdateResponse, error) {
	out := new(RemoteIdentityUpdateResponse)
	err := c.cc.Invoke(ctx, "/v1.RemoteIdentities/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteIdentitiesClient) Delete(ctx context.Context, in *RemoteIdentityDeleteRequest, opts ...grpc.CallOption) (*RemoteIdentityDeleteResponse, error) {
	out := new(RemoteIdentityDeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.RemoteIdentities/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteIdentitiesClient) List(ctx context.Context, in *RemoteIdentityListRequest, opts ...grpc.CallOption) (*RemoteIdentityListResponse, error) {
	out := new(RemoteIdentityListResponse)
	err := c.cc.Invoke(ctx, "/v1.RemoteIdentities/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteIdentitiesServer is the server API for RemoteIdentities service.
// All implementations must embed UnimplementedRemoteIdentitiesServer
// for forward compatibility
//
// Deprecated: Do not use.
type RemoteIdentitiesServer interface {
	// Create registers a new RemoteIdentity.
	Create(context.Context, *RemoteIdentityCreateRequest) (*RemoteIdentityCreateResponse, error)
	// Get reads one RemoteIdentity by ID.
	Get(context.Context, *RemoteIdentityGetRequest) (*RemoteIdentityGetResponse, error)
	// Update replaces all the fields of a RemoteIdentity by ID.
	Update(context.Context, *RemoteIdentityUpdateRequest) (*RemoteIdentityUpdateResponse, error)
	// Delete removes a RemoteIdentity by ID.
	Delete(context.Context, *RemoteIdentityDeleteRequest) (*RemoteIdentityDeleteResponse, error)
	// List gets a list of RemoteIdentities matching a given set of criteria.
	List(context.Context, *RemoteIdentityListRequest) (*RemoteIdentityListResponse, error)
	mustEmbedUnimplementedRemoteIdentitiesServer()
}

// UnimplementedRemoteIdentitiesServer must be embedded to have forward compatible implementations.
type UnimplementedRemoteIdentitiesServer struct {
}

func (UnimplementedRemoteIdentitiesServer) Create(context.Context, *RemoteIdentityCreateRequest) (*RemoteIdentityCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRemoteIdentitiesServer) Get(context.Context, *RemoteIdentityGetRequest) (*RemoteIdentityGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRemoteIdentitiesServer) Update(context.Context, *RemoteIdentityUpdateRequest) (*RemoteIdentityUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRemoteIdentitiesServer) Delete(context.Context, *RemoteIdentityDeleteRequest) (*RemoteIdentityDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRemoteIdentitiesServer) List(context.Context, *RemoteIdentityListRequest) (*RemoteIdentityListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRemoteIdentitiesServer) mustEmbedUnimplementedRemoteIdentitiesServer() {}

// UnsafeRemoteIdentitiesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoteIdentitiesServer will
// result in compilation errors.
type UnsafeRemoteIdentitiesServer interface {
	mustEmbedUnimplementedRemoteIdentitiesServer()
}

// Deprecated: Do not use.
func RegisterRemoteIdentitiesServer(s grpc.ServiceRegistrar, srv RemoteIdentitiesServer) {
	s.RegisterService(&_RemoteIdentities_serviceDesc, srv)
}

func _RemoteIdentities_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteIdentityCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteIdentitiesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RemoteIdentities/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteIdentitiesServer).Create(ctx, req.(*RemoteIdentityCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteIdentities_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteIdentityGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteIdentitiesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RemoteIdentities/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteIdentitiesServer).Get(ctx, req.(*RemoteIdentityGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteIdentities_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteIdentityUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteIdentitiesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RemoteIdentities/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteIdentitiesServer).Update(ctx, req.(*RemoteIdentityUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteIdentities_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteIdentityDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteIdentitiesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RemoteIdentities/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteIdentitiesServer).Delete(ctx, req.(*RemoteIdentityDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteIdentities_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteIdentityListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteIdentitiesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RemoteIdentities/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteIdentitiesServer).List(ctx, req.(*RemoteIdentityListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RemoteIdentities_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.RemoteIdentities",
	HandlerType: (*RemoteIdentitiesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RemoteIdentities_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RemoteIdentities_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RemoteIdentities_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RemoteIdentities_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RemoteIdentities_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "remote_identities.proto",
}
