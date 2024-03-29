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

// AccountGrantsHistoryClient is the client API for AccountGrantsHistory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountGrantsHistoryClient interface {
	// List gets a list of AccountGrantHistory records matching a given set of criteria.
	List(ctx context.Context, in *AccountGrantHistoryListRequest, opts ...grpc.CallOption) (*AccountGrantHistoryListResponse, error)
}

type accountGrantsHistoryClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountGrantsHistoryClient(cc grpc.ClientConnInterface) AccountGrantsHistoryClient {
	return &accountGrantsHistoryClient{cc}
}

func (c *accountGrantsHistoryClient) List(ctx context.Context, in *AccountGrantHistoryListRequest, opts ...grpc.CallOption) (*AccountGrantHistoryListResponse, error) {
	out := new(AccountGrantHistoryListResponse)
	err := c.cc.Invoke(ctx, "/v1.AccountGrantsHistory/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountGrantsHistoryServer is the server API for AccountGrantsHistory service.
// All implementations must embed UnimplementedAccountGrantsHistoryServer
// for forward compatibility
type AccountGrantsHistoryServer interface {
	// List gets a list of AccountGrantHistory records matching a given set of criteria.
	List(context.Context, *AccountGrantHistoryListRequest) (*AccountGrantHistoryListResponse, error)
	mustEmbedUnimplementedAccountGrantsHistoryServer()
}

// UnimplementedAccountGrantsHistoryServer must be embedded to have forward compatible implementations.
type UnimplementedAccountGrantsHistoryServer struct {
}

func (UnimplementedAccountGrantsHistoryServer) List(context.Context, *AccountGrantHistoryListRequest) (*AccountGrantHistoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAccountGrantsHistoryServer) mustEmbedUnimplementedAccountGrantsHistoryServer() {}

// UnsafeAccountGrantsHistoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountGrantsHistoryServer will
// result in compilation errors.
type UnsafeAccountGrantsHistoryServer interface {
	mustEmbedUnimplementedAccountGrantsHistoryServer()
}

func RegisterAccountGrantsHistoryServer(s grpc.ServiceRegistrar, srv AccountGrantsHistoryServer) {
	s.RegisterService(&_AccountGrantsHistory_serviceDesc, srv)
}

func _AccountGrantsHistory_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountGrantHistoryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountGrantsHistoryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccountGrantsHistory/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountGrantsHistoryServer).List(ctx, req.(*AccountGrantHistoryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountGrantsHistory_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.AccountGrantsHistory",
	HandlerType: (*AccountGrantsHistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _AccountGrantsHistory_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account_grants_history.proto",
}
