// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package hzglexamplesvc

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

// DataTypeExperimentsClient is the client API for DataTypeExperiments service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataTypeExperimentsClient interface {
	GetSingleResponse(ctx context.Context, in *RequestOfManyTypes, opts ...grpc.CallOption) (*ResponseOfManyTypes, error)
	GetStreamResponse(ctx context.Context, in *RequestOfManyTypes, opts ...grpc.CallOption) (DataTypeExperiments_GetStreamResponseClient, error)
	GetSingleResponseFromStream(ctx context.Context, opts ...grpc.CallOption) (DataTypeExperiments_GetSingleResponseFromStreamClient, error)
	GetStreamResponseFromStream(ctx context.Context, opts ...grpc.CallOption) (DataTypeExperiments_GetStreamResponseFromStreamClient, error)
}

type dataTypeExperimentsClient struct {
	cc grpc.ClientConnInterface
}

func NewDataTypeExperimentsClient(cc grpc.ClientConnInterface) DataTypeExperimentsClient {
	return &dataTypeExperimentsClient{cc}
}

func (c *dataTypeExperimentsClient) GetSingleResponse(ctx context.Context, in *RequestOfManyTypes, opts ...grpc.CallOption) (*ResponseOfManyTypes, error) {
	out := new(ResponseOfManyTypes)
	err := c.cc.Invoke(ctx, "/hzglexamplesvc.DataTypeExperiments/GetSingleResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataTypeExperimentsClient) GetStreamResponse(ctx context.Context, in *RequestOfManyTypes, opts ...grpc.CallOption) (DataTypeExperiments_GetStreamResponseClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataTypeExperiments_ServiceDesc.Streams[0], "/hzglexamplesvc.DataTypeExperiments/GetStreamResponse", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataTypeExperimentsGetStreamResponseClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataTypeExperiments_GetStreamResponseClient interface {
	Recv() (*ResponseOfManyTypes, error)
	grpc.ClientStream
}

type dataTypeExperimentsGetStreamResponseClient struct {
	grpc.ClientStream
}

func (x *dataTypeExperimentsGetStreamResponseClient) Recv() (*ResponseOfManyTypes, error) {
	m := new(ResponseOfManyTypes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataTypeExperimentsClient) GetSingleResponseFromStream(ctx context.Context, opts ...grpc.CallOption) (DataTypeExperiments_GetSingleResponseFromStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataTypeExperiments_ServiceDesc.Streams[1], "/hzglexamplesvc.DataTypeExperiments/GetSingleResponseFromStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataTypeExperimentsGetSingleResponseFromStreamClient{stream}
	return x, nil
}

type DataTypeExperiments_GetSingleResponseFromStreamClient interface {
	Send(*RequestOfManyTypes) error
	CloseAndRecv() (*ResponseOfManyTypes, error)
	grpc.ClientStream
}

type dataTypeExperimentsGetSingleResponseFromStreamClient struct {
	grpc.ClientStream
}

func (x *dataTypeExperimentsGetSingleResponseFromStreamClient) Send(m *RequestOfManyTypes) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataTypeExperimentsGetSingleResponseFromStreamClient) CloseAndRecv() (*ResponseOfManyTypes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ResponseOfManyTypes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataTypeExperimentsClient) GetStreamResponseFromStream(ctx context.Context, opts ...grpc.CallOption) (DataTypeExperiments_GetStreamResponseFromStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataTypeExperiments_ServiceDesc.Streams[2], "/hzglexamplesvc.DataTypeExperiments/GetStreamResponseFromStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataTypeExperimentsGetStreamResponseFromStreamClient{stream}
	return x, nil
}

type DataTypeExperiments_GetStreamResponseFromStreamClient interface {
	Send(*RequestOfManyTypes) error
	Recv() (*ResponseOfManyTypes, error)
	grpc.ClientStream
}

type dataTypeExperimentsGetStreamResponseFromStreamClient struct {
	grpc.ClientStream
}

func (x *dataTypeExperimentsGetStreamResponseFromStreamClient) Send(m *RequestOfManyTypes) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataTypeExperimentsGetStreamResponseFromStreamClient) Recv() (*ResponseOfManyTypes, error) {
	m := new(ResponseOfManyTypes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataTypeExperimentsServer is the server API for DataTypeExperiments service.
// All implementations must embed UnimplementedDataTypeExperimentsServer
// for forward compatibility
type DataTypeExperimentsServer interface {
	GetSingleResponse(context.Context, *RequestOfManyTypes) (*ResponseOfManyTypes, error)
	GetStreamResponse(*RequestOfManyTypes, DataTypeExperiments_GetStreamResponseServer) error
	GetSingleResponseFromStream(DataTypeExperiments_GetSingleResponseFromStreamServer) error
	GetStreamResponseFromStream(DataTypeExperiments_GetStreamResponseFromStreamServer) error
	mustEmbedUnimplementedDataTypeExperimentsServer()
}

// UnimplementedDataTypeExperimentsServer must be embedded to have forward compatible implementations.
type UnimplementedDataTypeExperimentsServer struct {
}

func (UnimplementedDataTypeExperimentsServer) GetSingleResponse(context.Context, *RequestOfManyTypes) (*ResponseOfManyTypes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleResponse not implemented")
}
func (UnimplementedDataTypeExperimentsServer) GetStreamResponse(*RequestOfManyTypes, DataTypeExperiments_GetStreamResponseServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStreamResponse not implemented")
}
func (UnimplementedDataTypeExperimentsServer) GetSingleResponseFromStream(DataTypeExperiments_GetSingleResponseFromStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSingleResponseFromStream not implemented")
}
func (UnimplementedDataTypeExperimentsServer) GetStreamResponseFromStream(DataTypeExperiments_GetStreamResponseFromStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStreamResponseFromStream not implemented")
}
func (UnimplementedDataTypeExperimentsServer) mustEmbedUnimplementedDataTypeExperimentsServer() {}

// UnsafeDataTypeExperimentsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataTypeExperimentsServer will
// result in compilation errors.
type UnsafeDataTypeExperimentsServer interface {
	mustEmbedUnimplementedDataTypeExperimentsServer()
}

func RegisterDataTypeExperimentsServer(s grpc.ServiceRegistrar, srv DataTypeExperimentsServer) {
	s.RegisterService(&DataTypeExperiments_ServiceDesc, srv)
}

func _DataTypeExperiments_GetSingleResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestOfManyTypes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataTypeExperimentsServer).GetSingleResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hzglexamplesvc.DataTypeExperiments/GetSingleResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataTypeExperimentsServer).GetSingleResponse(ctx, req.(*RequestOfManyTypes))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataTypeExperiments_GetStreamResponse_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestOfManyTypes)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataTypeExperimentsServer).GetStreamResponse(m, &dataTypeExperimentsGetStreamResponseServer{stream})
}

type DataTypeExperiments_GetStreamResponseServer interface {
	Send(*ResponseOfManyTypes) error
	grpc.ServerStream
}

type dataTypeExperimentsGetStreamResponseServer struct {
	grpc.ServerStream
}

func (x *dataTypeExperimentsGetStreamResponseServer) Send(m *ResponseOfManyTypes) error {
	return x.ServerStream.SendMsg(m)
}

func _DataTypeExperiments_GetSingleResponseFromStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataTypeExperimentsServer).GetSingleResponseFromStream(&dataTypeExperimentsGetSingleResponseFromStreamServer{stream})
}

type DataTypeExperiments_GetSingleResponseFromStreamServer interface {
	SendAndClose(*ResponseOfManyTypes) error
	Recv() (*RequestOfManyTypes, error)
	grpc.ServerStream
}

type dataTypeExperimentsGetSingleResponseFromStreamServer struct {
	grpc.ServerStream
}

func (x *dataTypeExperimentsGetSingleResponseFromStreamServer) SendAndClose(m *ResponseOfManyTypes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataTypeExperimentsGetSingleResponseFromStreamServer) Recv() (*RequestOfManyTypes, error) {
	m := new(RequestOfManyTypes)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DataTypeExperiments_GetStreamResponseFromStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataTypeExperimentsServer).GetStreamResponseFromStream(&dataTypeExperimentsGetStreamResponseFromStreamServer{stream})
}

type DataTypeExperiments_GetStreamResponseFromStreamServer interface {
	Send(*ResponseOfManyTypes) error
	Recv() (*RequestOfManyTypes, error)
	grpc.ServerStream
}

type dataTypeExperimentsGetStreamResponseFromStreamServer struct {
	grpc.ServerStream
}

func (x *dataTypeExperimentsGetStreamResponseFromStreamServer) Send(m *ResponseOfManyTypes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataTypeExperimentsGetStreamResponseFromStreamServer) Recv() (*RequestOfManyTypes, error) {
	m := new(RequestOfManyTypes)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataTypeExperiments_ServiceDesc is the grpc.ServiceDesc for DataTypeExperiments service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataTypeExperiments_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hzglexamplesvc.DataTypeExperiments",
	HandlerType: (*DataTypeExperimentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSingleResponse",
			Handler:    _DataTypeExperiments_GetSingleResponse_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStreamResponse",
			Handler:       _DataTypeExperiments_GetStreamResponse_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetSingleResponseFromStream",
			Handler:       _DataTypeExperiments_GetSingleResponseFromStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetStreamResponseFromStream",
			Handler:       _DataTypeExperiments_GetStreamResponseFromStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hzglexamplesvc/myservice.proto",
}