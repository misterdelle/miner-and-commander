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
// source: bos/v1/cooling.proto

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
	CoolingService_GetCoolingState_FullMethodName  = "/braiins.bos.v1.CoolingService/GetCoolingState"
	CoolingService_SetImmersionMode_FullMethodName = "/braiins.bos.v1.CoolingService/SetImmersionMode"
)

// CoolingServiceClient is the client API for CoolingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoolingServiceClient interface {
	// Method to get current temperature measurements and fans states
	GetCoolingState(ctx context.Context, in *GetCoolingStateRequest, opts ...grpc.CallOption) (*GetCoolingStateResponse, error)
	// Method to set/toggle immersion mode
	SetImmersionMode(ctx context.Context, in *SetImmersionModeRequest, opts ...grpc.CallOption) (*SetImmersionModeResponse, error)
}

type coolingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoolingServiceClient(cc grpc.ClientConnInterface) CoolingServiceClient {
	return &coolingServiceClient{cc}
}

func (c *coolingServiceClient) GetCoolingState(ctx context.Context, in *GetCoolingStateRequest, opts ...grpc.CallOption) (*GetCoolingStateResponse, error) {
	out := new(GetCoolingStateResponse)
	err := c.cc.Invoke(ctx, CoolingService_GetCoolingState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coolingServiceClient) SetImmersionMode(ctx context.Context, in *SetImmersionModeRequest, opts ...grpc.CallOption) (*SetImmersionModeResponse, error) {
	out := new(SetImmersionModeResponse)
	err := c.cc.Invoke(ctx, CoolingService_SetImmersionMode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoolingServiceServer is the server API for CoolingService service.
// All implementations must embed UnimplementedCoolingServiceServer
// for forward compatibility
type CoolingServiceServer interface {
	// Method to get current temperature measurements and fans states
	GetCoolingState(context.Context, *GetCoolingStateRequest) (*GetCoolingStateResponse, error)
	// Method to set/toggle immersion mode
	SetImmersionMode(context.Context, *SetImmersionModeRequest) (*SetImmersionModeResponse, error)
	mustEmbedUnimplementedCoolingServiceServer()
}

// UnimplementedCoolingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCoolingServiceServer struct {
}

func (UnimplementedCoolingServiceServer) GetCoolingState(context.Context, *GetCoolingStateRequest) (*GetCoolingStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoolingState not implemented")
}
func (UnimplementedCoolingServiceServer) SetImmersionMode(context.Context, *SetImmersionModeRequest) (*SetImmersionModeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetImmersionMode not implemented")
}
func (UnimplementedCoolingServiceServer) mustEmbedUnimplementedCoolingServiceServer() {}

// UnsafeCoolingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoolingServiceServer will
// result in compilation errors.
type UnsafeCoolingServiceServer interface {
	mustEmbedUnimplementedCoolingServiceServer()
}

func RegisterCoolingServiceServer(s grpc.ServiceRegistrar, srv CoolingServiceServer) {
	s.RegisterService(&CoolingService_ServiceDesc, srv)
}

func _CoolingService_GetCoolingState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoolingStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoolingServiceServer).GetCoolingState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoolingService_GetCoolingState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoolingServiceServer).GetCoolingState(ctx, req.(*GetCoolingStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoolingService_SetImmersionMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetImmersionModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoolingServiceServer).SetImmersionMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoolingService_SetImmersionMode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoolingServiceServer).SetImmersionMode(ctx, req.(*SetImmersionModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CoolingService_ServiceDesc is the grpc.ServiceDesc for CoolingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoolingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "braiins.bos.v1.CoolingService",
	HandlerType: (*CoolingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCoolingState",
			Handler:    _CoolingService_GetCoolingState_Handler,
		},
		{
			MethodName: "SetImmersionMode",
			Handler:    _CoolingService_SetImmersionMode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bos/v1/cooling.proto",
}