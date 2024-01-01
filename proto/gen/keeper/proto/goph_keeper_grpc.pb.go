// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: proto/goph_keeper.proto

package pbkeeper

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
	GophKeeper_Register_FullMethodName    = "/gophKeeper.GophKeeper/Register"
	GophKeeper_Login_FullMethodName       = "/gophKeeper.GophKeeper/Login"
	GophKeeper_SaveLogPass_FullMethodName = "/gophKeeper.GophKeeper/SaveLogPass"
	GophKeeper_SaveText_FullMethodName    = "/gophKeeper.GophKeeper/SaveText"
	GophKeeper_SaveBinary_FullMethodName  = "/gophKeeper.GophKeeper/SaveBinary"
	GophKeeper_ShowData_FullMethodName    = "/gophKeeper.GophKeeper/ShowData"
)

// GophKeeperClient is the client API for GophKeeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GophKeeperClient interface {
	Register(ctx context.Context, in *AuthInfo, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *AuthInfo, opts ...grpc.CallOption) (*LoginResponse, error)
	SaveLogPass(ctx context.Context, in *LogPassData, opts ...grpc.CallOption) (*Empty, error)
	SaveText(ctx context.Context, in *TextData, opts ...grpc.CallOption) (*Empty, error)
	SaveBinary(ctx context.Context, in *BinaryData, opts ...grpc.CallOption) (*Empty, error)
	ShowData(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Data, error)
}

type gophKeeperClient struct {
	cc grpc.ClientConnInterface
}

func NewGophKeeperClient(cc grpc.ClientConnInterface) GophKeeperClient {
	return &gophKeeperClient{cc}
}

func (c *gophKeeperClient) Register(ctx context.Context, in *AuthInfo, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, GophKeeper_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) Login(ctx context.Context, in *AuthInfo, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, GophKeeper_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) SaveLogPass(ctx context.Context, in *LogPassData, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, GophKeeper_SaveLogPass_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) SaveText(ctx context.Context, in *TextData, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, GophKeeper_SaveText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) SaveBinary(ctx context.Context, in *BinaryData, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, GophKeeper_SaveBinary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) ShowData(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, GophKeeper_ShowData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GophKeeperServer is the server API for GophKeeper service.
// All implementations must embed UnimplementedGophKeeperServer
// for forward compatibility
type GophKeeperServer interface {
	Register(context.Context, *AuthInfo) (*RegisterResponse, error)
	Login(context.Context, *AuthInfo) (*LoginResponse, error)
	SaveLogPass(context.Context, *LogPassData) (*Empty, error)
	SaveText(context.Context, *TextData) (*Empty, error)
	SaveBinary(context.Context, *BinaryData) (*Empty, error)
	ShowData(context.Context, *Empty) (*Data, error)
	mustEmbedUnimplementedGophKeeperServer()
}

// UnimplementedGophKeeperServer must be embedded to have forward compatible implementations.
type UnimplementedGophKeeperServer struct {
}

func (UnimplementedGophKeeperServer) Register(context.Context, *AuthInfo) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedGophKeeperServer) Login(context.Context, *AuthInfo) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedGophKeeperServer) SaveLogPass(context.Context, *LogPassData) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveLogPass not implemented")
}
func (UnimplementedGophKeeperServer) SaveText(context.Context, *TextData) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveText not implemented")
}
func (UnimplementedGophKeeperServer) SaveBinary(context.Context, *BinaryData) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveBinary not implemented")
}
func (UnimplementedGophKeeperServer) ShowData(context.Context, *Empty) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowData not implemented")
}
func (UnimplementedGophKeeperServer) mustEmbedUnimplementedGophKeeperServer() {}

// UnsafeGophKeeperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GophKeeperServer will
// result in compilation errors.
type UnsafeGophKeeperServer interface {
	mustEmbedUnimplementedGophKeeperServer()
}

func RegisterGophKeeperServer(s grpc.ServiceRegistrar, srv GophKeeperServer) {
	s.RegisterService(&GophKeeper_ServiceDesc, srv)
}

func _GophKeeper_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).Register(ctx, req.(*AuthInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).Login(ctx, req.(*AuthInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_SaveLogPass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogPassData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).SaveLogPass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_SaveLogPass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).SaveLogPass(ctx, req.(*LogPassData))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_SaveText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).SaveText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_SaveText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).SaveText(ctx, req.(*TextData))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_SaveBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BinaryData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).SaveBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_SaveBinary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).SaveBinary(ctx, req.(*BinaryData))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_ShowData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).ShowData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_ShowData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).ShowData(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// GophKeeper_ServiceDesc is the grpc.ServiceDesc for GophKeeper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GophKeeper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gophKeeper.GophKeeper",
	HandlerType: (*GophKeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _GophKeeper_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _GophKeeper_Login_Handler,
		},
		{
			MethodName: "SaveLogPass",
			Handler:    _GophKeeper_SaveLogPass_Handler,
		},
		{
			MethodName: "SaveText",
			Handler:    _GophKeeper_SaveText_Handler,
		},
		{
			MethodName: "SaveBinary",
			Handler:    _GophKeeper_SaveBinary_Handler,
		},
		{
			MethodName: "ShowData",
			Handler:    _GophKeeper_ShowData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/goph_keeper.proto",
}
