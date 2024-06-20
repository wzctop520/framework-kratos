// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0
// source: test.proto

package v1

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
	Test_CreateTest_FullMethodName = "/api.test.service.v1.Test/CreateTest"
	Test_UpdateTest_FullMethodName = "/api.test.service.v1.Test/UpdateTest"
	Test_DeleteTest_FullMethodName = "/api.test.service.v1.Test/DeleteTest"
	Test_GetTest_FullMethodName    = "/api.test.service.v1.Test/GetTest"
	Test_ListTest_FullMethodName   = "/api.test.service.v1.Test/ListTest"
)

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestClient interface {
	CreateTest(ctx context.Context, in *CreateTestRequest, opts ...grpc.CallOption) (*CreateTestReply, error)
	UpdateTest(ctx context.Context, in *UpdateTestRequest, opts ...grpc.CallOption) (*UpdateTestReply, error)
	DeleteTest(ctx context.Context, in *DeleteTestRequest, opts ...grpc.CallOption) (*DeleteTestReply, error)
	GetTest(ctx context.Context, in *GetTestRequest, opts ...grpc.CallOption) (*GetTestReply, error)
	ListTest(ctx context.Context, in *ListTestRequest, opts ...grpc.CallOption) (*ListTestReply, error)
}

type testClient struct {
	cc grpc.ClientConnInterface
}

func NewTestClient(cc grpc.ClientConnInterface) TestClient {
	return &testClient{cc}
}

func (c *testClient) CreateTest(ctx context.Context, in *CreateTestRequest, opts ...grpc.CallOption) (*CreateTestReply, error) {
	out := new(CreateTestReply)
	err := c.cc.Invoke(ctx, Test_CreateTest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) UpdateTest(ctx context.Context, in *UpdateTestRequest, opts ...grpc.CallOption) (*UpdateTestReply, error) {
	out := new(UpdateTestReply)
	err := c.cc.Invoke(ctx, Test_UpdateTest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) DeleteTest(ctx context.Context, in *DeleteTestRequest, opts ...grpc.CallOption) (*DeleteTestReply, error) {
	out := new(DeleteTestReply)
	err := c.cc.Invoke(ctx, Test_DeleteTest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) GetTest(ctx context.Context, in *GetTestRequest, opts ...grpc.CallOption) (*GetTestReply, error) {
	out := new(GetTestReply)
	err := c.cc.Invoke(ctx, Test_GetTest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) ListTest(ctx context.Context, in *ListTestRequest, opts ...grpc.CallOption) (*ListTestReply, error) {
	out := new(ListTestReply)
	err := c.cc.Invoke(ctx, Test_ListTest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServer is the server API for Test service.
// All implementations must embed UnimplementedTestServer
// for forward compatibility
type TestServer interface {
	CreateTest(context.Context, *CreateTestRequest) (*CreateTestReply, error)
	UpdateTest(context.Context, *UpdateTestRequest) (*UpdateTestReply, error)
	DeleteTest(context.Context, *DeleteTestRequest) (*DeleteTestReply, error)
	GetTest(context.Context, *GetTestRequest) (*GetTestReply, error)
	ListTest(context.Context, *ListTestRequest) (*ListTestReply, error)
	mustEmbedUnimplementedTestServer()
}

// UnimplementedTestServer must be embedded to have forward compatible implementations.
type UnimplementedTestServer struct {
}

func (UnimplementedTestServer) CreateTest(context.Context, *CreateTestRequest) (*CreateTestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTest not implemented")
}
func (UnimplementedTestServer) UpdateTest(context.Context, *UpdateTestRequest) (*UpdateTestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTest not implemented")
}
func (UnimplementedTestServer) DeleteTest(context.Context, *DeleteTestRequest) (*DeleteTestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTest not implemented")
}
func (UnimplementedTestServer) GetTest(context.Context, *GetTestRequest) (*GetTestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTest not implemented")
}
func (UnimplementedTestServer) ListTest(context.Context, *ListTestRequest) (*ListTestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTest not implemented")
}
func (UnimplementedTestServer) mustEmbedUnimplementedTestServer() {}

// UnsafeTestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServer will
// result in compilation errors.
type UnsafeTestServer interface {
	mustEmbedUnimplementedTestServer()
}

func RegisterTestServer(s grpc.ServiceRegistrar, srv TestServer) {
	s.RegisterService(&Test_ServiceDesc, srv)
}

func _Test_CreateTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).CreateTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Test_CreateTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).CreateTest(ctx, req.(*CreateTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_UpdateTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).UpdateTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Test_UpdateTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).UpdateTest(ctx, req.(*UpdateTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_DeleteTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).DeleteTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Test_DeleteTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).DeleteTest(ctx, req.(*DeleteTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_GetTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).GetTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Test_GetTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).GetTest(ctx, req.(*GetTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_ListTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).ListTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Test_ListTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).ListTest(ctx, req.(*ListTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Test_ServiceDesc is the grpc.ServiceDesc for Test service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Test_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.test.service.v1.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTest",
			Handler:    _Test_CreateTest_Handler,
		},
		{
			MethodName: "UpdateTest",
			Handler:    _Test_UpdateTest_Handler,
		},
		{
			MethodName: "DeleteTest",
			Handler:    _Test_DeleteTest_Handler,
		},
		{
			MethodName: "GetTest",
			Handler:    _Test_GetTest_Handler,
		},
		{
			MethodName: "ListTest",
			Handler:    _Test_ListTest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}