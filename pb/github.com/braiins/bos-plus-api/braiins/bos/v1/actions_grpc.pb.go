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
// source: bos/v1/actions.proto

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
	ActionsService_Start_FullMethodName                 = "/braiins.bos.v1.ActionsService/Start"
	ActionsService_Stop_FullMethodName                  = "/braiins.bos.v1.ActionsService/Stop"
	ActionsService_PauseMining_FullMethodName           = "/braiins.bos.v1.ActionsService/PauseMining"
	ActionsService_ResumeMining_FullMethodName          = "/braiins.bos.v1.ActionsService/ResumeMining"
	ActionsService_Restart_FullMethodName               = "/braiins.bos.v1.ActionsService/Restart"
	ActionsService_Reboot_FullMethodName                = "/braiins.bos.v1.ActionsService/Reboot"
	ActionsService_SetLocateDeviceStatus_FullMethodName = "/braiins.bos.v1.ActionsService/SetLocateDeviceStatus"
	ActionsService_GetLocateDeviceStatus_FullMethodName = "/braiins.bos.v1.ActionsService/GetLocateDeviceStatus"
)

// ActionsServiceClient is the client API for ActionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActionsServiceClient interface {
	// Method to start bosminer
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	// Method to stop bosminer
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
	// Method to pause mining
	PauseMining(ctx context.Context, in *PauseMiningRequest, opts ...grpc.CallOption) (*PauseMiningResponse, error)
	// Method to resume mining
	ResumeMining(ctx context.Context, in *ResumeMiningRequest, opts ...grpc.CallOption) (*ResumeMiningResponse, error)
	// Method to restart bosminer
	Restart(ctx context.Context, in *RestartRequest, opts ...grpc.CallOption) (*RestartResponse, error)
	// Method to reboot whole miner
	Reboot(ctx context.Context, in *RebootRequest, opts ...grpc.CallOption) (*RebootResponse, error)
	// Method to enable/disable locate device mode
	SetLocateDeviceStatus(ctx context.Context, in *SetLocateDeviceStatusRequest, opts ...grpc.CallOption) (*LocateDeviceStatusResponse, error)
	// Method to retrieve the locate device mode status
	GetLocateDeviceStatus(ctx context.Context, in *GetLocateDeviceStatusRequest, opts ...grpc.CallOption) (*LocateDeviceStatusResponse, error)
}

type actionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActionsServiceClient(cc grpc.ClientConnInterface) ActionsServiceClient {
	return &actionsServiceClient{cc}
}

func (c *actionsServiceClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, ActionsService_Start_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, ActionsService_Stop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) PauseMining(ctx context.Context, in *PauseMiningRequest, opts ...grpc.CallOption) (*PauseMiningResponse, error) {
	out := new(PauseMiningResponse)
	err := c.cc.Invoke(ctx, ActionsService_PauseMining_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) ResumeMining(ctx context.Context, in *ResumeMiningRequest, opts ...grpc.CallOption) (*ResumeMiningResponse, error) {
	out := new(ResumeMiningResponse)
	err := c.cc.Invoke(ctx, ActionsService_ResumeMining_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) Restart(ctx context.Context, in *RestartRequest, opts ...grpc.CallOption) (*RestartResponse, error) {
	out := new(RestartResponse)
	err := c.cc.Invoke(ctx, ActionsService_Restart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) Reboot(ctx context.Context, in *RebootRequest, opts ...grpc.CallOption) (*RebootResponse, error) {
	out := new(RebootResponse)
	err := c.cc.Invoke(ctx, ActionsService_Reboot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) SetLocateDeviceStatus(ctx context.Context, in *SetLocateDeviceStatusRequest, opts ...grpc.CallOption) (*LocateDeviceStatusResponse, error) {
	out := new(LocateDeviceStatusResponse)
	err := c.cc.Invoke(ctx, ActionsService_SetLocateDeviceStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) GetLocateDeviceStatus(ctx context.Context, in *GetLocateDeviceStatusRequest, opts ...grpc.CallOption) (*LocateDeviceStatusResponse, error) {
	out := new(LocateDeviceStatusResponse)
	err := c.cc.Invoke(ctx, ActionsService_GetLocateDeviceStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActionsServiceServer is the server API for ActionsService service.
// All implementations must embed UnimplementedActionsServiceServer
// for forward compatibility
type ActionsServiceServer interface {
	// Method to start bosminer
	Start(context.Context, *StartRequest) (*StartResponse, error)
	// Method to stop bosminer
	Stop(context.Context, *StopRequest) (*StopResponse, error)
	// Method to pause mining
	PauseMining(context.Context, *PauseMiningRequest) (*PauseMiningResponse, error)
	// Method to resume mining
	ResumeMining(context.Context, *ResumeMiningRequest) (*ResumeMiningResponse, error)
	// Method to restart bosminer
	Restart(context.Context, *RestartRequest) (*RestartResponse, error)
	// Method to reboot whole miner
	Reboot(context.Context, *RebootRequest) (*RebootResponse, error)
	// Method to enable/disable locate device mode
	SetLocateDeviceStatus(context.Context, *SetLocateDeviceStatusRequest) (*LocateDeviceStatusResponse, error)
	// Method to retrieve the locate device mode status
	GetLocateDeviceStatus(context.Context, *GetLocateDeviceStatusRequest) (*LocateDeviceStatusResponse, error)
	mustEmbedUnimplementedActionsServiceServer()
}

// UnimplementedActionsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedActionsServiceServer struct {
}

func (UnimplementedActionsServiceServer) Start(context.Context, *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedActionsServiceServer) Stop(context.Context, *StopRequest) (*StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedActionsServiceServer) PauseMining(context.Context, *PauseMiningRequest) (*PauseMiningResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PauseMining not implemented")
}
func (UnimplementedActionsServiceServer) ResumeMining(context.Context, *ResumeMiningRequest) (*ResumeMiningResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeMining not implemented")
}
func (UnimplementedActionsServiceServer) Restart(context.Context, *RestartRequest) (*RestartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedActionsServiceServer) Reboot(context.Context, *RebootRequest) (*RebootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reboot not implemented")
}
func (UnimplementedActionsServiceServer) SetLocateDeviceStatus(context.Context, *SetLocateDeviceStatusRequest) (*LocateDeviceStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLocateDeviceStatus not implemented")
}
func (UnimplementedActionsServiceServer) GetLocateDeviceStatus(context.Context, *GetLocateDeviceStatusRequest) (*LocateDeviceStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocateDeviceStatus not implemented")
}
func (UnimplementedActionsServiceServer) mustEmbedUnimplementedActionsServiceServer() {}

// UnsafeActionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActionsServiceServer will
// result in compilation errors.
type UnsafeActionsServiceServer interface {
	mustEmbedUnimplementedActionsServiceServer()
}

func RegisterActionsServiceServer(s grpc.ServiceRegistrar, srv ActionsServiceServer) {
	s.RegisterService(&ActionsService_ServiceDesc, srv)
}

func _ActionsService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_PauseMining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PauseMiningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).PauseMining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_PauseMining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).PauseMining(ctx, req.(*PauseMiningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_ResumeMining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeMiningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).ResumeMining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_ResumeMining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).ResumeMining(ctx, req.(*ResumeMiningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_Restart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).Restart(ctx, req.(*RestartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_Reboot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RebootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).Reboot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_Reboot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).Reboot(ctx, req.(*RebootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_SetLocateDeviceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetLocateDeviceStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).SetLocateDeviceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_SetLocateDeviceStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).SetLocateDeviceStatus(ctx, req.(*SetLocateDeviceStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_GetLocateDeviceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocateDeviceStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).GetLocateDeviceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_GetLocateDeviceStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).GetLocateDeviceStatus(ctx, req.(*GetLocateDeviceStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ActionsService_ServiceDesc is the grpc.ServiceDesc for ActionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "braiins.bos.v1.ActionsService",
	HandlerType: (*ActionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _ActionsService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _ActionsService_Stop_Handler,
		},
		{
			MethodName: "PauseMining",
			Handler:    _ActionsService_PauseMining_Handler,
		},
		{
			MethodName: "ResumeMining",
			Handler:    _ActionsService_ResumeMining_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _ActionsService_Restart_Handler,
		},
		{
			MethodName: "Reboot",
			Handler:    _ActionsService_Reboot_Handler,
		},
		{
			MethodName: "SetLocateDeviceStatus",
			Handler:    _ActionsService_SetLocateDeviceStatus_Handler,
		},
		{
			MethodName: "GetLocateDeviceStatus",
			Handler:    _ActionsService_GetLocateDeviceStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bos/v1/actions.proto",
}
