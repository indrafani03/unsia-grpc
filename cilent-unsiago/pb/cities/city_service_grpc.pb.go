// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: city_service.proto

package cities

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	CitiesService_GetCity_FullMethodName   = "/cities.CitiesService/GetCity"
	CitiesService_GetCities_FullMethodName = "/cities.CitiesService/GetCities"
)

// CitiesServiceClient is the client API for CitiesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CitiesServiceClient interface {
	GetCity(ctx context.Context, in *City, opts ...grpc.CallOption) (*City, error)
	GetCities(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (CitiesService_GetCitiesClient, error)
}

type citiesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCitiesServiceClient(cc grpc.ClientConnInterface) CitiesServiceClient {
	return &citiesServiceClient{cc}
}

func (c *citiesServiceClient) GetCity(ctx context.Context, in *City, opts ...grpc.CallOption) (*City, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(City)
	err := c.cc.Invoke(ctx, CitiesService_GetCity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *citiesServiceClient) GetCities(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (CitiesService_GetCitiesClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CitiesService_ServiceDesc.Streams[0], CitiesService_GetCities_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &citiesServiceGetCitiesClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CitiesService_GetCitiesClient interface {
	Recv() (*CitiesStream, error)
	grpc.ClientStream
}

type citiesServiceGetCitiesClient struct {
	grpc.ClientStream
}

func (x *citiesServiceGetCitiesClient) Recv() (*CitiesStream, error) {
	m := new(CitiesStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CitiesServiceServer is the server API for CitiesService service.
// All implementations should embed UnimplementedCitiesServiceServer
// for forward compatibility
type CitiesServiceServer interface {
	GetCity(context.Context, *City) (*City, error)
	GetCities(*EmptyMessage, CitiesService_GetCitiesServer) error
}

// UnimplementedCitiesServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCitiesServiceServer struct {
}

func (UnimplementedCitiesServiceServer) GetCity(context.Context, *City) (*City, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCity not implemented")
}
func (UnimplementedCitiesServiceServer) GetCities(*EmptyMessage, CitiesService_GetCitiesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCities not implemented")
}

// UnsafeCitiesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CitiesServiceServer will
// result in compilation errors.
type UnsafeCitiesServiceServer interface {
	mustEmbedUnimplementedCitiesServiceServer()
}

func RegisterCitiesServiceServer(s grpc.ServiceRegistrar, srv CitiesServiceServer) {
	s.RegisterService(&CitiesService_ServiceDesc, srv)
}

func _CitiesService_GetCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(City)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CitiesServiceServer).GetCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CitiesService_GetCity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CitiesServiceServer).GetCity(ctx, req.(*City))
	}
	return interceptor(ctx, in, info, handler)
}

func _CitiesService_GetCities_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CitiesServiceServer).GetCities(m, &citiesServiceGetCitiesServer{ServerStream: stream})
}

type CitiesService_GetCitiesServer interface {
	Send(*CitiesStream) error
	grpc.ServerStream
}

type citiesServiceGetCitiesServer struct {
	grpc.ServerStream
}

func (x *citiesServiceGetCitiesServer) Send(m *CitiesStream) error {
	return x.ServerStream.SendMsg(m)
}

// CitiesService_ServiceDesc is the grpc.ServiceDesc for CitiesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CitiesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cities.CitiesService",
	HandlerType: (*CitiesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCity",
			Handler:    _CitiesService_GetCity_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCities",
			Handler:       _CitiesService_GetCities_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "city_service.proto",
}
