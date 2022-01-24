// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package go_Parkinglot_grpc

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

// ParkinigLotClient is the client API for ParkinigLot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ParkinigLotClient interface {
	ParkNewCar(ctx context.Context, in *CarDetailsInfo, opts ...grpc.CallOption) (*SlotInfo, error)
	ParkingDetails(ctx context.Context, in *DetailsInfoParams, opts ...grpc.CallOption) (*CarDetails, error)
	UnparkCar(ctx context.Context, in *GetId, opts ...grpc.CallOption) (*CarDetails, error)
}

type parkinigLotClient struct {
	cc grpc.ClientConnInterface
}

func NewParkinigLotClient(cc grpc.ClientConnInterface) ParkinigLotClient {
	return &parkinigLotClient{cc}
}

func (c *parkinigLotClient) ParkNewCar(ctx context.Context, in *CarDetailsInfo, opts ...grpc.CallOption) (*SlotInfo, error) {
	out := new(SlotInfo)
	err := c.cc.Invoke(ctx, "/ParkingLot_Main.ParkinigLot/ParkNewCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parkinigLotClient) ParkingDetails(ctx context.Context, in *DetailsInfoParams, opts ...grpc.CallOption) (*CarDetails, error) {
	out := new(CarDetails)
	err := c.cc.Invoke(ctx, "/ParkingLot_Main.ParkinigLot/ParkingDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parkinigLotClient) UnparkCar(ctx context.Context, in *GetId, opts ...grpc.CallOption) (*CarDetails, error) {
	out := new(CarDetails)
	err := c.cc.Invoke(ctx, "/ParkingLot_Main.ParkinigLot/UnparkCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParkinigLotServer is the server API for ParkinigLot service.
// All implementations must embed UnimplementedParkinigLotServer
// for forward compatibility
type ParkinigLotServer interface {
	ParkNewCar(context.Context, *CarDetailsInfo) (*SlotInfo, error)
	ParkingDetails(context.Context, *DetailsInfoParams) (*CarDetails, error)
	UnparkCar(context.Context, *GetId) (*CarDetails, error)
	mustEmbedUnimplementedParkinigLotServer()
}

// UnimplementedParkinigLotServer must be embedded to have forward compatible implementations.
type UnimplementedParkinigLotServer struct {
}

func (UnimplementedParkinigLotServer) ParkNewCar(context.Context, *CarDetailsInfo) (*SlotInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParkNewCar not implemented")
}
func (UnimplementedParkinigLotServer) ParkingDetails(context.Context, *DetailsInfoParams) (*CarDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParkingDetails not implemented")
}
func (UnimplementedParkinigLotServer) UnparkCar(context.Context, *GetId) (*CarDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnparkCar not implemented")
}
func (UnimplementedParkinigLotServer) mustEmbedUnimplementedParkinigLotServer() {}

// UnsafeParkinigLotServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParkinigLotServer will
// result in compilation errors.
type UnsafeParkinigLotServer interface {
	mustEmbedUnimplementedParkinigLotServer()
}

func RegisterParkinigLotServer(s grpc.ServiceRegistrar, srv ParkinigLotServer) {
	s.RegisterService(&ParkinigLot_ServiceDesc, srv)
}

func _ParkinigLot_ParkNewCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CarDetailsInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParkinigLotServer).ParkNewCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ParkingLot_Main.ParkinigLot/ParkNewCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParkinigLotServer).ParkNewCar(ctx, req.(*CarDetailsInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParkinigLot_ParkingDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailsInfoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParkinigLotServer).ParkingDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ParkingLot_Main.ParkinigLot/ParkingDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParkinigLotServer).ParkingDetails(ctx, req.(*DetailsInfoParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParkinigLot_UnparkCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParkinigLotServer).UnparkCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ParkingLot_Main.ParkinigLot/UnparkCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParkinigLotServer).UnparkCar(ctx, req.(*GetId))
	}
	return interceptor(ctx, in, info, handler)
}

// ParkinigLot_ServiceDesc is the grpc.ServiceDesc for ParkinigLot service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ParkinigLot_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ParkingLot_Main.ParkinigLot",
	HandlerType: (*ParkinigLotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParkNewCar",
			Handler:    _ParkinigLot_ParkNewCar_Handler,
		},
		{
			MethodName: "ParkingDetails",
			Handler:    _ParkinigLot_ParkingDetails_Handler,
		},
		{
			MethodName: "UnparkCar",
			Handler:    _ParkinigLot_UnparkCar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Parkinglot/Parkinglot.proto",
}