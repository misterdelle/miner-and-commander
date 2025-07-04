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
// - protoc             v6.31.1
// source: bos/v1/performance.proto

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
	PerformanceService_GetTunerState_FullMethodName            = "/braiins.bos.v1.PerformanceService/GetTunerState"
	PerformanceService_ListTargetProfiles_FullMethodName       = "/braiins.bos.v1.PerformanceService/ListTargetProfiles"
	PerformanceService_SetDefaultPowerTarget_FullMethodName    = "/braiins.bos.v1.PerformanceService/SetDefaultPowerTarget"
	PerformanceService_SetPowerTarget_FullMethodName           = "/braiins.bos.v1.PerformanceService/SetPowerTarget"
	PerformanceService_IncrementPowerTarget_FullMethodName     = "/braiins.bos.v1.PerformanceService/IncrementPowerTarget"
	PerformanceService_DecrementPowerTarget_FullMethodName     = "/braiins.bos.v1.PerformanceService/DecrementPowerTarget"
	PerformanceService_SetDefaultHashrateTarget_FullMethodName = "/braiins.bos.v1.PerformanceService/SetDefaultHashrateTarget"
	PerformanceService_SetHashrateTarget_FullMethodName        = "/braiins.bos.v1.PerformanceService/SetHashrateTarget"
	PerformanceService_IncrementHashrateTarget_FullMethodName  = "/braiins.bos.v1.PerformanceService/IncrementHashrateTarget"
	PerformanceService_DecrementHashrateTarget_FullMethodName  = "/braiins.bos.v1.PerformanceService/DecrementHashrateTarget"
	PerformanceService_SetDPS_FullMethodName                   = "/braiins.bos.v1.PerformanceService/SetDPS"
	PerformanceService_SetPerformanceMode_FullMethodName       = "/braiins.bos.v1.PerformanceService/SetPerformanceMode"
	PerformanceService_GetActivePerformanceMode_FullMethodName = "/braiins.bos.v1.PerformanceService/GetActivePerformanceMode"
	PerformanceService_RemoveTunedProfiles_FullMethodName      = "/braiins.bos.v1.PerformanceService/RemoveTunedProfiles"
)

// PerformanceServiceClient is the client API for PerformanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PerformanceServiceClient interface {
	// Method to get the current tuner details
	GetTunerState(ctx context.Context, in *GetTunerStateRequest, opts ...grpc.CallOption) (*GetTunerStateResponse, error)
	// Method to get the current tuner details
	ListTargetProfiles(ctx context.Context, in *ListTargetProfilesRequest, opts ...grpc.CallOption) (*ListTargetProfilesResponse, error)
	// Method to set default power target for tuner
	SetDefaultPowerTarget(ctx context.Context, in *SetDefaultPowerTargetRequest, opts ...grpc.CallOption) (*SetPowerTargetResponse, error)
	// Method to set absolute power target for tuner
	SetPowerTarget(ctx context.Context, in *SetPowerTargetRequest, opts ...grpc.CallOption) (*SetPowerTargetResponse, error)
	// Method to increment power target for tuner
	IncrementPowerTarget(ctx context.Context, in *IncrementPowerTargetRequest, opts ...grpc.CallOption) (*SetPowerTargetResponse, error)
	// Method to decrement power target for tuner
	DecrementPowerTarget(ctx context.Context, in *DecrementPowerTargetRequest, opts ...grpc.CallOption) (*SetPowerTargetResponse, error)
	// Method to set default hashrate target for tuner
	SetDefaultHashrateTarget(ctx context.Context, in *SetDefaultHashrateTargetRequest, opts ...grpc.CallOption) (*SetHashrateTargetResponse, error)
	// Method to set absolute hashrate target for tuner
	SetHashrateTarget(ctx context.Context, in *SetHashrateTargetRequest, opts ...grpc.CallOption) (*SetHashrateTargetResponse, error)
	// Method to increment hashrate target for tuner
	IncrementHashrateTarget(ctx context.Context, in *IncrementHashrateTargetRequest, opts ...grpc.CallOption) (*SetHashrateTargetResponse, error)
	// Method to decrement hashrate target for tuner
	DecrementHashrateTarget(ctx context.Context, in *DecrementHashrateTargetRequest, opts ...grpc.CallOption) (*SetHashrateTargetResponse, error)
	// Method to set Dynamic Performance Scaling
	SetDPS(ctx context.Context, in *SetDPSRequest, opts ...grpc.CallOption) (*SetDPSResponse, error)
	// Method to set performance mode
	SetPerformanceMode(ctx context.Context, in *SetPerformanceModeRequest, opts ...grpc.CallOption) (*PerformanceMode, error)
	// Method to read active(runtime) performance mode
	GetActivePerformanceMode(ctx context.Context, in *GetPerformanceModeRequest, opts ...grpc.CallOption) (*PerformanceMode, error)
	// Method to remove tuned profiles
	RemoveTunedProfiles(ctx context.Context, in *RemoveTunedProfilesRequest, opts ...grpc.CallOption) (*RemoveTunedProfilesResponse, error)
}

type performanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPerformanceServiceClient(cc grpc.ClientConnInterface) PerformanceServiceClient {
	return &performanceServiceClient{cc}
}

func (c *performanceServiceClient) GetTunerState(ctx context.Context, in *GetTunerStateRequest, opts ...grpc.CallOption) (*GetTunerStateResponse, error) {
	out := new(GetTunerStateResponse)
	err := c.cc.Invoke(ctx, PerformanceService_GetTunerState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *performanceServiceClient) ListTargetProfiles(ctx context.Context, in *ListTargetProfilesRequest, opts ...grpc.CallOption) (*ListTargetProfilesResponse, error) {
	out := new(ListTargetProfilesResponse)
	err := c.cc.Invoke(ctx, PerformanceService_ListTargetProfiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *performanceServiceClient) SetDefaultPowerTarget(ctx context.Context, in *SetDefaultPowerTargetRequest, opts ...grpc.CallOption) (*SetPowerTargetResponse, error) {
	out := new(SetPowerTargetResponse)
	err := c.cc.Invoke(ctx, PerformanceService_SetDefaultPowerTarget_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *performanceServiceClient) SetPowerTarget(ctx context.Context, in *SetPowerTargetRequest, opts ...grpc.CallOption) (*SetPowerTargetResponse, error) {
	out := new(SetPowerTargetResponse)
	err := c.cc.Invoke(ctx, PerformanceService_SetPowerTarget_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *performanceServiceClient) IncrementPowerTarget(ctx context.Context, in *IncrementPowerTargetRequest, opts ...grpc.CallOption) (*SetPowerTargetResponse, error) {
	out := new(SetPowerTargetResponse)
	err := c.cc.Invoke(ctx, PerformanceService_IncrementPowerTarget_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *performanceServiceClient) DecrementPowerTarget(ctx context.Context, in *DecrementPowerTargetRequest, opts ...grpc.CallOption) (*SetPowerTargetResponse, error) {
	out := new(SetPowerTargetResponse)
	err := c.cc.Invoke(ctx, PerformanceService_DecrementPowerTarget_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *performanceServiceClient) SetDefaultHashrateTarget(ctx context.Context, in *SetDefaultHashrateTargetRequest, opts ...grpc.CallOption) (*SetHashrateTargetResponse, error) {
	out := new(SetHashrateTargetResponse)
	err := c.cc.Invoke(ctx, PerformanceService_SetDefaultHashrateTarget_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *performanceServiceClient) SetHashrateTarget(ctx context.Context, in *SetHashrateTargetRequest, opts ...grpc.CallOption) (*SetHashrateTargetResponse, error) {
	out := new(SetHashrateTargetResponse)
	err := c.cc.Invoke(ctx, PerformanceService_SetHashrateTarget_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *performanceServiceClient) IncrementHashrateTarget(ctx context.Context, in *IncrementHashrateTargetRequest, opts ...grpc.CallOption) (*SetHashrateTargetResponse, error) {
	out := new(SetHashrateTargetResponse)
	err := c.cc.Invoke(ctx, PerformanceService_IncrementHashrateTarget_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *performanceServiceClient) DecrementHashrateTarget(ctx context.Context, in *DecrementHashrateTargetRequest, opts ...grpc.CallOption) (*SetHashrateTargetResponse, error) {
	out := new(SetHashrateTargetResponse)
	err := c.cc.Invoke(ctx, PerformanceService_DecrementHashrateTarget_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *performanceServiceClient) SetDPS(ctx context.Context, in *SetDPSRequest, opts ...grpc.CallOption) (*SetDPSResponse, error) {
	out := new(SetDPSResponse)
	err := c.cc.Invoke(ctx, PerformanceService_SetDPS_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *performanceServiceClient) SetPerformanceMode(ctx context.Context, in *SetPerformanceModeRequest, opts ...grpc.CallOption) (*PerformanceMode, error) {
	out := new(PerformanceMode)
	err := c.cc.Invoke(ctx, PerformanceService_SetPerformanceMode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *performanceServiceClient) GetActivePerformanceMode(ctx context.Context, in *GetPerformanceModeRequest, opts ...grpc.CallOption) (*PerformanceMode, error) {
	out := new(PerformanceMode)
	err := c.cc.Invoke(ctx, PerformanceService_GetActivePerformanceMode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *performanceServiceClient) RemoveTunedProfiles(ctx context.Context, in *RemoveTunedProfilesRequest, opts ...grpc.CallOption) (*RemoveTunedProfilesResponse, error) {
	out := new(RemoveTunedProfilesResponse)
	err := c.cc.Invoke(ctx, PerformanceService_RemoveTunedProfiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PerformanceServiceServer is the server API for PerformanceService service.
// All implementations must embed UnimplementedPerformanceServiceServer
// for forward compatibility
type PerformanceServiceServer interface {
	// Method to get the current tuner details
	GetTunerState(context.Context, *GetTunerStateRequest) (*GetTunerStateResponse, error)
	// Method to get the current tuner details
	ListTargetProfiles(context.Context, *ListTargetProfilesRequest) (*ListTargetProfilesResponse, error)
	// Method to set default power target for tuner
	SetDefaultPowerTarget(context.Context, *SetDefaultPowerTargetRequest) (*SetPowerTargetResponse, error)
	// Method to set absolute power target for tuner
	SetPowerTarget(context.Context, *SetPowerTargetRequest) (*SetPowerTargetResponse, error)
	// Method to increment power target for tuner
	IncrementPowerTarget(context.Context, *IncrementPowerTargetRequest) (*SetPowerTargetResponse, error)
	// Method to decrement power target for tuner
	DecrementPowerTarget(context.Context, *DecrementPowerTargetRequest) (*SetPowerTargetResponse, error)
	// Method to set default hashrate target for tuner
	SetDefaultHashrateTarget(context.Context, *SetDefaultHashrateTargetRequest) (*SetHashrateTargetResponse, error)
	// Method to set absolute hashrate target for tuner
	SetHashrateTarget(context.Context, *SetHashrateTargetRequest) (*SetHashrateTargetResponse, error)
	// Method to increment hashrate target for tuner
	IncrementHashrateTarget(context.Context, *IncrementHashrateTargetRequest) (*SetHashrateTargetResponse, error)
	// Method to decrement hashrate target for tuner
	DecrementHashrateTarget(context.Context, *DecrementHashrateTargetRequest) (*SetHashrateTargetResponse, error)
	// Method to set Dynamic Performance Scaling
	SetDPS(context.Context, *SetDPSRequest) (*SetDPSResponse, error)
	// Method to set performance mode
	SetPerformanceMode(context.Context, *SetPerformanceModeRequest) (*PerformanceMode, error)
	// Method to read active(runtime) performance mode
	GetActivePerformanceMode(context.Context, *GetPerformanceModeRequest) (*PerformanceMode, error)
	// Method to remove tuned profiles
	RemoveTunedProfiles(context.Context, *RemoveTunedProfilesRequest) (*RemoveTunedProfilesResponse, error)
	mustEmbedUnimplementedPerformanceServiceServer()
}

// UnimplementedPerformanceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPerformanceServiceServer struct {
}

func (UnimplementedPerformanceServiceServer) GetTunerState(context.Context, *GetTunerStateRequest) (*GetTunerStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTunerState not implemented")
}
func (UnimplementedPerformanceServiceServer) ListTargetProfiles(context.Context, *ListTargetProfilesRequest) (*ListTargetProfilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTargetProfiles not implemented")
}
func (UnimplementedPerformanceServiceServer) SetDefaultPowerTarget(context.Context, *SetDefaultPowerTargetRequest) (*SetPowerTargetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDefaultPowerTarget not implemented")
}
func (UnimplementedPerformanceServiceServer) SetPowerTarget(context.Context, *SetPowerTargetRequest) (*SetPowerTargetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPowerTarget not implemented")
}
func (UnimplementedPerformanceServiceServer) IncrementPowerTarget(context.Context, *IncrementPowerTargetRequest) (*SetPowerTargetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrementPowerTarget not implemented")
}
func (UnimplementedPerformanceServiceServer) DecrementPowerTarget(context.Context, *DecrementPowerTargetRequest) (*SetPowerTargetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecrementPowerTarget not implemented")
}
func (UnimplementedPerformanceServiceServer) SetDefaultHashrateTarget(context.Context, *SetDefaultHashrateTargetRequest) (*SetHashrateTargetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDefaultHashrateTarget not implemented")
}
func (UnimplementedPerformanceServiceServer) SetHashrateTarget(context.Context, *SetHashrateTargetRequest) (*SetHashrateTargetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetHashrateTarget not implemented")
}
func (UnimplementedPerformanceServiceServer) IncrementHashrateTarget(context.Context, *IncrementHashrateTargetRequest) (*SetHashrateTargetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrementHashrateTarget not implemented")
}
func (UnimplementedPerformanceServiceServer) DecrementHashrateTarget(context.Context, *DecrementHashrateTargetRequest) (*SetHashrateTargetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecrementHashrateTarget not implemented")
}
func (UnimplementedPerformanceServiceServer) SetDPS(context.Context, *SetDPSRequest) (*SetDPSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDPS not implemented")
}
func (UnimplementedPerformanceServiceServer) SetPerformanceMode(context.Context, *SetPerformanceModeRequest) (*PerformanceMode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPerformanceMode not implemented")
}
func (UnimplementedPerformanceServiceServer) GetActivePerformanceMode(context.Context, *GetPerformanceModeRequest) (*PerformanceMode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivePerformanceMode not implemented")
}
func (UnimplementedPerformanceServiceServer) RemoveTunedProfiles(context.Context, *RemoveTunedProfilesRequest) (*RemoveTunedProfilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTunedProfiles not implemented")
}
func (UnimplementedPerformanceServiceServer) mustEmbedUnimplementedPerformanceServiceServer() {}

// UnsafePerformanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PerformanceServiceServer will
// result in compilation errors.
type UnsafePerformanceServiceServer interface {
	mustEmbedUnimplementedPerformanceServiceServer()
}

func RegisterPerformanceServiceServer(s grpc.ServiceRegistrar, srv PerformanceServiceServer) {
	s.RegisterService(&PerformanceService_ServiceDesc, srv)
}

func _PerformanceService_GetTunerState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTunerStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerformanceServiceServer).GetTunerState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerformanceService_GetTunerState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerformanceServiceServer).GetTunerState(ctx, req.(*GetTunerStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PerformanceService_ListTargetProfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTargetProfilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerformanceServiceServer).ListTargetProfiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerformanceService_ListTargetProfiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerformanceServiceServer).ListTargetProfiles(ctx, req.(*ListTargetProfilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PerformanceService_SetDefaultPowerTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDefaultPowerTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerformanceServiceServer).SetDefaultPowerTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerformanceService_SetDefaultPowerTarget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerformanceServiceServer).SetDefaultPowerTarget(ctx, req.(*SetDefaultPowerTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PerformanceService_SetPowerTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPowerTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerformanceServiceServer).SetPowerTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerformanceService_SetPowerTarget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerformanceServiceServer).SetPowerTarget(ctx, req.(*SetPowerTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PerformanceService_IncrementPowerTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrementPowerTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerformanceServiceServer).IncrementPowerTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerformanceService_IncrementPowerTarget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerformanceServiceServer).IncrementPowerTarget(ctx, req.(*IncrementPowerTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PerformanceService_DecrementPowerTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecrementPowerTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerformanceServiceServer).DecrementPowerTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerformanceService_DecrementPowerTarget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerformanceServiceServer).DecrementPowerTarget(ctx, req.(*DecrementPowerTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PerformanceService_SetDefaultHashrateTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDefaultHashrateTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerformanceServiceServer).SetDefaultHashrateTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerformanceService_SetDefaultHashrateTarget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerformanceServiceServer).SetDefaultHashrateTarget(ctx, req.(*SetDefaultHashrateTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PerformanceService_SetHashrateTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetHashrateTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerformanceServiceServer).SetHashrateTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerformanceService_SetHashrateTarget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerformanceServiceServer).SetHashrateTarget(ctx, req.(*SetHashrateTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PerformanceService_IncrementHashrateTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrementHashrateTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerformanceServiceServer).IncrementHashrateTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerformanceService_IncrementHashrateTarget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerformanceServiceServer).IncrementHashrateTarget(ctx, req.(*IncrementHashrateTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PerformanceService_DecrementHashrateTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecrementHashrateTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerformanceServiceServer).DecrementHashrateTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerformanceService_DecrementHashrateTarget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerformanceServiceServer).DecrementHashrateTarget(ctx, req.(*DecrementHashrateTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PerformanceService_SetDPS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDPSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerformanceServiceServer).SetDPS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerformanceService_SetDPS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerformanceServiceServer).SetDPS(ctx, req.(*SetDPSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PerformanceService_SetPerformanceMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPerformanceModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerformanceServiceServer).SetPerformanceMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerformanceService_SetPerformanceMode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerformanceServiceServer).SetPerformanceMode(ctx, req.(*SetPerformanceModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PerformanceService_GetActivePerformanceMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPerformanceModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerformanceServiceServer).GetActivePerformanceMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerformanceService_GetActivePerformanceMode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerformanceServiceServer).GetActivePerformanceMode(ctx, req.(*GetPerformanceModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PerformanceService_RemoveTunedProfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveTunedProfilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerformanceServiceServer).RemoveTunedProfiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerformanceService_RemoveTunedProfiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerformanceServiceServer).RemoveTunedProfiles(ctx, req.(*RemoveTunedProfilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PerformanceService_ServiceDesc is the grpc.ServiceDesc for PerformanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PerformanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "braiins.bos.v1.PerformanceService",
	HandlerType: (*PerformanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTunerState",
			Handler:    _PerformanceService_GetTunerState_Handler,
		},
		{
			MethodName: "ListTargetProfiles",
			Handler:    _PerformanceService_ListTargetProfiles_Handler,
		},
		{
			MethodName: "SetDefaultPowerTarget",
			Handler:    _PerformanceService_SetDefaultPowerTarget_Handler,
		},
		{
			MethodName: "SetPowerTarget",
			Handler:    _PerformanceService_SetPowerTarget_Handler,
		},
		{
			MethodName: "IncrementPowerTarget",
			Handler:    _PerformanceService_IncrementPowerTarget_Handler,
		},
		{
			MethodName: "DecrementPowerTarget",
			Handler:    _PerformanceService_DecrementPowerTarget_Handler,
		},
		{
			MethodName: "SetDefaultHashrateTarget",
			Handler:    _PerformanceService_SetDefaultHashrateTarget_Handler,
		},
		{
			MethodName: "SetHashrateTarget",
			Handler:    _PerformanceService_SetHashrateTarget_Handler,
		},
		{
			MethodName: "IncrementHashrateTarget",
			Handler:    _PerformanceService_IncrementHashrateTarget_Handler,
		},
		{
			MethodName: "DecrementHashrateTarget",
			Handler:    _PerformanceService_DecrementHashrateTarget_Handler,
		},
		{
			MethodName: "SetDPS",
			Handler:    _PerformanceService_SetDPS_Handler,
		},
		{
			MethodName: "SetPerformanceMode",
			Handler:    _PerformanceService_SetPerformanceMode_Handler,
		},
		{
			MethodName: "GetActivePerformanceMode",
			Handler:    _PerformanceService_GetActivePerformanceMode_Handler,
		},
		{
			MethodName: "RemoveTunedProfiles",
			Handler:    _PerformanceService_RemoveTunedProfiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bos/v1/performance.proto",
}
