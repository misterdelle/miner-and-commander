// Copyright (C) 2023  Braiins Systems s.r.o.
//
// This file is part of Braiins Open-Source Initiative (BOSI).
//
// BOSI is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//
// Please, keep in mind that we may also license BOSI or any part thereof
// under a proprietary license. For more information on the terms and conditions
// of such proprietary license or if you have any other questions, please
// contact us at opensource@braiins.com.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.1
// source: bos/v1/miner.proto

package bos_proto_v1

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
	MinerService_GetMinerStatus_FullMethodName    = "/braiins.bos.v1.MinerService/GetMinerStatus"
	MinerService_GetMinerDetails_FullMethodName   = "/braiins.bos.v1.MinerService/GetMinerDetails"
	MinerService_GetMinerStats_FullMethodName     = "/braiins.bos.v1.MinerService/GetMinerStats"
	MinerService_GetHashboards_FullMethodName     = "/braiins.bos.v1.MinerService/GetHashboards"
	MinerService_GetSupportArchive_FullMethodName = "/braiins.bos.v1.MinerService/GetSupportArchive"
	MinerService_EnableHashboards_FullMethodName  = "/braiins.bos.v1.MinerService/EnableHashboards"
	MinerService_DisableHashboards_FullMethodName = "/braiins.bos.v1.MinerService/DisableHashboards"
)

// MinerServiceClient is the client API for MinerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MinerServiceClient interface {
	// Method to fetch miner status
	GetMinerStatus(ctx context.Context, in *GetMinerStatusRequest, opts ...grpc.CallOption) (MinerService_GetMinerStatusClient, error)
	// Method to get miner details.
	GetMinerDetails(ctx context.Context, in *GetMinerDetailsRequest, opts ...grpc.CallOption) (*GetMinerDetailsResponse, error)
	// Method to get aggregated miner stats.
	GetMinerStats(ctx context.Context, in *GetMinerStatsRequest, opts ...grpc.CallOption) (*GetMinerStatsResponse, error)
	// Method to get miner hashboards state and statistics.
	GetHashboards(ctx context.Context, in *GetHashboardsRequest, opts ...grpc.CallOption) (*GetHashboardsResponse, error)
	// Method to download BOS support archive
	// Method returns stream of messages with binary chunks that needs to be concatenated on the caller side
	GetSupportArchive(ctx context.Context, in *GetSupportArchiveRequest, opts ...grpc.CallOption) (MinerService_GetSupportArchiveClient, error)
	// Method to enable hashboards
	EnableHashboards(ctx context.Context, in *EnableHashboardsRequest, opts ...grpc.CallOption) (*EnableHashboardsResponse, error)
	// Method to disable hashboards
	DisableHashboards(ctx context.Context, in *DisableHashboardsRequest, opts ...grpc.CallOption) (*DisableHashboardsResponse, error)
}

type minerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMinerServiceClient(cc grpc.ClientConnInterface) MinerServiceClient {
	return &minerServiceClient{cc}
}

func (c *minerServiceClient) GetMinerStatus(ctx context.Context, in *GetMinerStatusRequest, opts ...grpc.CallOption) (MinerService_GetMinerStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &MinerService_ServiceDesc.Streams[0], MinerService_GetMinerStatus_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &minerServiceGetMinerStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MinerService_GetMinerStatusClient interface {
	Recv() (*GetMinerStatusResponse, error)
	grpc.ClientStream
}

type minerServiceGetMinerStatusClient struct {
	grpc.ClientStream
}

func (x *minerServiceGetMinerStatusClient) Recv() (*GetMinerStatusResponse, error) {
	m := new(GetMinerStatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *minerServiceClient) GetMinerDetails(ctx context.Context, in *GetMinerDetailsRequest, opts ...grpc.CallOption) (*GetMinerDetailsResponse, error) {
	out := new(GetMinerDetailsResponse)
	err := c.cc.Invoke(ctx, MinerService_GetMinerDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerServiceClient) GetMinerStats(ctx context.Context, in *GetMinerStatsRequest, opts ...grpc.CallOption) (*GetMinerStatsResponse, error) {
	out := new(GetMinerStatsResponse)
	err := c.cc.Invoke(ctx, MinerService_GetMinerStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerServiceClient) GetHashboards(ctx context.Context, in *GetHashboardsRequest, opts ...grpc.CallOption) (*GetHashboardsResponse, error) {
	out := new(GetHashboardsResponse)
	err := c.cc.Invoke(ctx, MinerService_GetHashboards_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerServiceClient) GetSupportArchive(ctx context.Context, in *GetSupportArchiveRequest, opts ...grpc.CallOption) (MinerService_GetSupportArchiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &MinerService_ServiceDesc.Streams[1], MinerService_GetSupportArchive_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &minerServiceGetSupportArchiveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MinerService_GetSupportArchiveClient interface {
	Recv() (*GetSupportArchiveResponse, error)
	grpc.ClientStream
}

type minerServiceGetSupportArchiveClient struct {
	grpc.ClientStream
}

func (x *minerServiceGetSupportArchiveClient) Recv() (*GetSupportArchiveResponse, error) {
	m := new(GetSupportArchiveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *minerServiceClient) EnableHashboards(ctx context.Context, in *EnableHashboardsRequest, opts ...grpc.CallOption) (*EnableHashboardsResponse, error) {
	out := new(EnableHashboardsResponse)
	err := c.cc.Invoke(ctx, MinerService_EnableHashboards_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerServiceClient) DisableHashboards(ctx context.Context, in *DisableHashboardsRequest, opts ...grpc.CallOption) (*DisableHashboardsResponse, error) {
	out := new(DisableHashboardsResponse)
	err := c.cc.Invoke(ctx, MinerService_DisableHashboards_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MinerServiceServer is the server API for MinerService service.
// All implementations must embed UnimplementedMinerServiceServer
// for forward compatibility
type MinerServiceServer interface {
	// Method to fetch miner status
	GetMinerStatus(*GetMinerStatusRequest, MinerService_GetMinerStatusServer) error
	// Method to get miner details.
	GetMinerDetails(context.Context, *GetMinerDetailsRequest) (*GetMinerDetailsResponse, error)
	// Method to get aggregated miner stats.
	GetMinerStats(context.Context, *GetMinerStatsRequest) (*GetMinerStatsResponse, error)
	// Method to get miner hashboards state and statistics.
	GetHashboards(context.Context, *GetHashboardsRequest) (*GetHashboardsResponse, error)
	// Method to download BOS support archive
	// Method returns stream of messages with binary chunks that needs to be concatenated on the caller side
	GetSupportArchive(*GetSupportArchiveRequest, MinerService_GetSupportArchiveServer) error
	// Method to enable hashboards
	EnableHashboards(context.Context, *EnableHashboardsRequest) (*EnableHashboardsResponse, error)
	// Method to disable hashboards
	DisableHashboards(context.Context, *DisableHashboardsRequest) (*DisableHashboardsResponse, error)
	mustEmbedUnimplementedMinerServiceServer()
}

// UnimplementedMinerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMinerServiceServer struct {
}

func (UnimplementedMinerServiceServer) GetMinerStatus(*GetMinerStatusRequest, MinerService_GetMinerStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMinerStatus not implemented")
}
func (UnimplementedMinerServiceServer) GetMinerDetails(context.Context, *GetMinerDetailsRequest) (*GetMinerDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMinerDetails not implemented")
}
func (UnimplementedMinerServiceServer) GetMinerStats(context.Context, *GetMinerStatsRequest) (*GetMinerStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMinerStats not implemented")
}
func (UnimplementedMinerServiceServer) GetHashboards(context.Context, *GetHashboardsRequest) (*GetHashboardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHashboards not implemented")
}
func (UnimplementedMinerServiceServer) GetSupportArchive(*GetSupportArchiveRequest, MinerService_GetSupportArchiveServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSupportArchive not implemented")
}
func (UnimplementedMinerServiceServer) EnableHashboards(context.Context, *EnableHashboardsRequest) (*EnableHashboardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableHashboards not implemented")
}
func (UnimplementedMinerServiceServer) DisableHashboards(context.Context, *DisableHashboardsRequest) (*DisableHashboardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableHashboards not implemented")
}
func (UnimplementedMinerServiceServer) mustEmbedUnimplementedMinerServiceServer() {}

// UnsafeMinerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MinerServiceServer will
// result in compilation errors.
type UnsafeMinerServiceServer interface {
	mustEmbedUnimplementedMinerServiceServer()
}

func RegisterMinerServiceServer(s grpc.ServiceRegistrar, srv MinerServiceServer) {
	s.RegisterService(&MinerService_ServiceDesc, srv)
}

func _MinerService_GetMinerStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetMinerStatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MinerServiceServer).GetMinerStatus(m, &minerServiceGetMinerStatusServer{stream})
}

type MinerService_GetMinerStatusServer interface {
	Send(*GetMinerStatusResponse) error
	grpc.ServerStream
}

type minerServiceGetMinerStatusServer struct {
	grpc.ServerStream
}

func (x *minerServiceGetMinerStatusServer) Send(m *GetMinerStatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MinerService_GetMinerDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMinerDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServiceServer).GetMinerDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MinerService_GetMinerDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServiceServer).GetMinerDetails(ctx, req.(*GetMinerDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MinerService_GetMinerStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMinerStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServiceServer).GetMinerStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MinerService_GetMinerStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServiceServer).GetMinerStats(ctx, req.(*GetMinerStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MinerService_GetHashboards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHashboardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServiceServer).GetHashboards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MinerService_GetHashboards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServiceServer).GetHashboards(ctx, req.(*GetHashboardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MinerService_GetSupportArchive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetSupportArchiveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MinerServiceServer).GetSupportArchive(m, &minerServiceGetSupportArchiveServer{stream})
}

type MinerService_GetSupportArchiveServer interface {
	Send(*GetSupportArchiveResponse) error
	grpc.ServerStream
}

type minerServiceGetSupportArchiveServer struct {
	grpc.ServerStream
}

func (x *minerServiceGetSupportArchiveServer) Send(m *GetSupportArchiveResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MinerService_EnableHashboards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableHashboardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServiceServer).EnableHashboards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MinerService_EnableHashboards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServiceServer).EnableHashboards(ctx, req.(*EnableHashboardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MinerService_DisableHashboards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableHashboardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServiceServer).DisableHashboards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MinerService_DisableHashboards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServiceServer).DisableHashboards(ctx, req.(*DisableHashboardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MinerService_ServiceDesc is the grpc.ServiceDesc for MinerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MinerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "braiins.bos.v1.MinerService",
	HandlerType: (*MinerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMinerDetails",
			Handler:    _MinerService_GetMinerDetails_Handler,
		},
		{
			MethodName: "GetMinerStats",
			Handler:    _MinerService_GetMinerStats_Handler,
		},
		{
			MethodName: "GetHashboards",
			Handler:    _MinerService_GetHashboards_Handler,
		},
		{
			MethodName: "EnableHashboards",
			Handler:    _MinerService_EnableHashboards_Handler,
		},
		{
			MethodName: "DisableHashboards",
			Handler:    _MinerService_DisableHashboards_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMinerStatus",
			Handler:       _MinerService_GetMinerStatus_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetSupportArchive",
			Handler:       _MinerService_GetSupportArchive_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bos/v1/miner.proto",
}
