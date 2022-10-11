// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: pop.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SkyClient is the client API for Sky service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SkyClient interface {
	//? COMING FROM Sky
	//? TO Opencore
	Challenge(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*P2PChallenge, error)
	Notification(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Sky_NotificationClient, error)
	Sign(ctx context.Context, in *Web3WalletPermission, opts ...grpc.CallOption) (Sky_SignClient, error)
	Interaction(ctx context.Context, in *Web3WalletPermission, opts ...grpc.CallOption) (Sky_InteractionClient, error)
	//? GOING TO Sky
	//? FROM Opencore
	Validate(ctx context.Context, in *Web3WalletPermission, opts ...grpc.CallOption) (*P2PConnectionStatus, error)
	Invalidate(ctx context.Context, in *Web3WalletPermission, opts ...grpc.CallOption) (*P2PConnectionStatus, error)
	Sync(ctx context.Context, in *Web3WalletPermission, opts ...grpc.CallOption) (*RubixWalletData, error)
	// Sync updates node about signed challenge codes and initiated transactions.
	Find(ctx context.Context, in *Web3WalletPermission, opts ...grpc.CallOption) (*RubixWalletData, error)
	Host(ctx context.Context, in *RubixWalletData, opts ...grpc.CallOption) (*Web3WalletPermission, error)
	RbtPay(ctx context.Context, in *TxnPayload, opts ...grpc.CallOption) (Sky_RbtPayClient, error)
	RbtPOS(ctx context.Context, in *TxnPayload, opts ...grpc.CallOption) (Sky_RbtPOSClient, error)
}

type skyClient struct {
	cc grpc.ClientConnInterface
}

func NewSkyClient(cc grpc.ClientConnInterface) SkyClient {
	return &skyClient{cc}
}

func (c *skyClient) Challenge(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*P2PChallenge, error) {
	out := new(P2PChallenge)
	err := c.cc.Invoke(ctx, "/protos.Sky/Challenge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skyClient) Notification(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Sky_NotificationClient, error) {
	stream, err := c.cc.NewStream(ctx, &Sky_ServiceDesc.Streams[0], "/protos.Sky/Notification", opts...)
	if err != nil {
		return nil, err
	}
	x := &skyNotificationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Sky_NotificationClient interface {
	Recv() (*P2PNotification, error)
	grpc.ClientStream
}

type skyNotificationClient struct {
	grpc.ClientStream
}

func (x *skyNotificationClient) Recv() (*P2PNotification, error) {
	m := new(P2PNotification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *skyClient) Sign(ctx context.Context, in *Web3WalletPermission, opts ...grpc.CallOption) (Sky_SignClient, error) {
	stream, err := c.cc.NewStream(ctx, &Sky_ServiceDesc.Streams[1], "/protos.Sky/Sign", opts...)
	if err != nil {
		return nil, err
	}
	x := &skySignClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Sky_SignClient interface {
	Recv() (*P2PConnectionStatus, error)
	grpc.ClientStream
}

type skySignClient struct {
	grpc.ClientStream
}

func (x *skySignClient) Recv() (*P2PConnectionStatus, error) {
	m := new(P2PConnectionStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *skyClient) Interaction(ctx context.Context, in *Web3WalletPermission, opts ...grpc.CallOption) (Sky_InteractionClient, error) {
	stream, err := c.cc.NewStream(ctx, &Sky_ServiceDesc.Streams[2], "/protos.Sky/Interaction", opts...)
	if err != nil {
		return nil, err
	}
	x := &skyInteractionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Sky_InteractionClient interface {
	Recv() (*P2PConnectionStatus, error)
	grpc.ClientStream
}

type skyInteractionClient struct {
	grpc.ClientStream
}

func (x *skyInteractionClient) Recv() (*P2PConnectionStatus, error) {
	m := new(P2PConnectionStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *skyClient) Validate(ctx context.Context, in *Web3WalletPermission, opts ...grpc.CallOption) (*P2PConnectionStatus, error) {
	out := new(P2PConnectionStatus)
	err := c.cc.Invoke(ctx, "/protos.Sky/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skyClient) Invalidate(ctx context.Context, in *Web3WalletPermission, opts ...grpc.CallOption) (*P2PConnectionStatus, error) {
	out := new(P2PConnectionStatus)
	err := c.cc.Invoke(ctx, "/protos.Sky/Invalidate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skyClient) Sync(ctx context.Context, in *Web3WalletPermission, opts ...grpc.CallOption) (*RubixWalletData, error) {
	out := new(RubixWalletData)
	err := c.cc.Invoke(ctx, "/protos.Sky/Sync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skyClient) Find(ctx context.Context, in *Web3WalletPermission, opts ...grpc.CallOption) (*RubixWalletData, error) {
	out := new(RubixWalletData)
	err := c.cc.Invoke(ctx, "/protos.Sky/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skyClient) Host(ctx context.Context, in *RubixWalletData, opts ...grpc.CallOption) (*Web3WalletPermission, error) {
	out := new(Web3WalletPermission)
	err := c.cc.Invoke(ctx, "/protos.Sky/Host", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skyClient) RbtPay(ctx context.Context, in *TxnPayload, opts ...grpc.CallOption) (Sky_RbtPayClient, error) {
	stream, err := c.cc.NewStream(ctx, &Sky_ServiceDesc.Streams[3], "/protos.Sky/RbtPay", opts...)
	if err != nil {
		return nil, err
	}
	x := &skyRbtPayClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Sky_RbtPayClient interface {
	Recv() (*TxnStatus, error)
	grpc.ClientStream
}

type skyRbtPayClient struct {
	grpc.ClientStream
}

func (x *skyRbtPayClient) Recv() (*TxnStatus, error) {
	m := new(TxnStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *skyClient) RbtPOS(ctx context.Context, in *TxnPayload, opts ...grpc.CallOption) (Sky_RbtPOSClient, error) {
	stream, err := c.cc.NewStream(ctx, &Sky_ServiceDesc.Streams[4], "/protos.Sky/RbtPOS", opts...)
	if err != nil {
		return nil, err
	}
	x := &skyRbtPOSClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Sky_RbtPOSClient interface {
	Recv() (*TxnStatus, error)
	grpc.ClientStream
}

type skyRbtPOSClient struct {
	grpc.ClientStream
}

func (x *skyRbtPOSClient) Recv() (*TxnStatus, error) {
	m := new(TxnStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SkyServer is the server API for Sky service.
// All implementations must embed UnimplementedSkyServer
// for forward compatibility
type SkyServer interface {
	//? COMING FROM Sky
	//? TO Opencore
	Challenge(context.Context, *emptypb.Empty) (*P2PChallenge, error)
	Notification(*emptypb.Empty, Sky_NotificationServer) error
	Sign(*Web3WalletPermission, Sky_SignServer) error
	Interaction(*Web3WalletPermission, Sky_InteractionServer) error
	//? GOING TO Sky
	//? FROM Opencore
	Validate(context.Context, *Web3WalletPermission) (*P2PConnectionStatus, error)
	Invalidate(context.Context, *Web3WalletPermission) (*P2PConnectionStatus, error)
	Sync(context.Context, *Web3WalletPermission) (*RubixWalletData, error)
	// Sync updates node about signed challenge codes and initiated transactions.
	Find(context.Context, *Web3WalletPermission) (*RubixWalletData, error)
	Host(context.Context, *RubixWalletData) (*Web3WalletPermission, error)
	RbtPay(*TxnPayload, Sky_RbtPayServer) error
	RbtPOS(*TxnPayload, Sky_RbtPOSServer) error
	mustEmbedUnimplementedSkyServer()
}

// UnimplementedSkyServer must be embedded to have forward compatible implementations.
type UnimplementedSkyServer struct {
}

func (UnimplementedSkyServer) Challenge(context.Context, *emptypb.Empty) (*P2PChallenge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Challenge not implemented")
}
func (UnimplementedSkyServer) Notification(*emptypb.Empty, Sky_NotificationServer) error {
	return status.Errorf(codes.Unimplemented, "method Notification not implemented")
}
func (UnimplementedSkyServer) Sign(*Web3WalletPermission, Sky_SignServer) error {
	return status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
func (UnimplementedSkyServer) Interaction(*Web3WalletPermission, Sky_InteractionServer) error {
	return status.Errorf(codes.Unimplemented, "method Interaction not implemented")
}
func (UnimplementedSkyServer) Validate(context.Context, *Web3WalletPermission) (*P2PConnectionStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedSkyServer) Invalidate(context.Context, *Web3WalletPermission) (*P2PConnectionStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invalidate not implemented")
}
func (UnimplementedSkyServer) Sync(context.Context, *Web3WalletPermission) (*RubixWalletData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedSkyServer) Find(context.Context, *Web3WalletPermission) (*RubixWalletData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedSkyServer) Host(context.Context, *RubixWalletData) (*Web3WalletPermission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Host not implemented")
}
func (UnimplementedSkyServer) RbtPay(*TxnPayload, Sky_RbtPayServer) error {
	return status.Errorf(codes.Unimplemented, "method RbtPay not implemented")
}
func (UnimplementedSkyServer) RbtPOS(*TxnPayload, Sky_RbtPOSServer) error {
	return status.Errorf(codes.Unimplemented, "method RbtPOS not implemented")
}
func (UnimplementedSkyServer) mustEmbedUnimplementedSkyServer() {}

// UnsafeSkyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SkyServer will
// result in compilation errors.
type UnsafeSkyServer interface {
	mustEmbedUnimplementedSkyServer()
}

func RegisterSkyServer(s grpc.ServiceRegistrar, srv SkyServer) {
	s.RegisterService(&Sky_ServiceDesc, srv)
}

func _Sky_Challenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkyServer).Challenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Sky/Challenge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkyServer).Challenge(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sky_Notification_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SkyServer).Notification(m, &skyNotificationServer{stream})
}

type Sky_NotificationServer interface {
	Send(*P2PNotification) error
	grpc.ServerStream
}

type skyNotificationServer struct {
	grpc.ServerStream
}

func (x *skyNotificationServer) Send(m *P2PNotification) error {
	return x.ServerStream.SendMsg(m)
}

func _Sky_Sign_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Web3WalletPermission)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SkyServer).Sign(m, &skySignServer{stream})
}

type Sky_SignServer interface {
	Send(*P2PConnectionStatus) error
	grpc.ServerStream
}

type skySignServer struct {
	grpc.ServerStream
}

func (x *skySignServer) Send(m *P2PConnectionStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _Sky_Interaction_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Web3WalletPermission)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SkyServer).Interaction(m, &skyInteractionServer{stream})
}

type Sky_InteractionServer interface {
	Send(*P2PConnectionStatus) error
	grpc.ServerStream
}

type skyInteractionServer struct {
	grpc.ServerStream
}

func (x *skyInteractionServer) Send(m *P2PConnectionStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _Sky_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Web3WalletPermission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkyServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Sky/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkyServer).Validate(ctx, req.(*Web3WalletPermission))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sky_Invalidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Web3WalletPermission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkyServer).Invalidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Sky/Invalidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkyServer).Invalidate(ctx, req.(*Web3WalletPermission))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sky_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Web3WalletPermission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkyServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Sky/Sync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkyServer).Sync(ctx, req.(*Web3WalletPermission))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sky_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Web3WalletPermission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkyServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Sky/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkyServer).Find(ctx, req.(*Web3WalletPermission))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sky_Host_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RubixWalletData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkyServer).Host(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Sky/Host",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkyServer).Host(ctx, req.(*RubixWalletData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sky_RbtPay_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TxnPayload)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SkyServer).RbtPay(m, &skyRbtPayServer{stream})
}

type Sky_RbtPayServer interface {
	Send(*TxnStatus) error
	grpc.ServerStream
}

type skyRbtPayServer struct {
	grpc.ServerStream
}

func (x *skyRbtPayServer) Send(m *TxnStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _Sky_RbtPOS_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TxnPayload)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SkyServer).RbtPOS(m, &skyRbtPOSServer{stream})
}

type Sky_RbtPOSServer interface {
	Send(*TxnStatus) error
	grpc.ServerStream
}

type skyRbtPOSServer struct {
	grpc.ServerStream
}

func (x *skyRbtPOSServer) Send(m *TxnStatus) error {
	return x.ServerStream.SendMsg(m)
}

// Sky_ServiceDesc is the grpc.ServiceDesc for Sky service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sky_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Sky",
	HandlerType: (*SkyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Challenge",
			Handler:    _Sky_Challenge_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _Sky_Validate_Handler,
		},
		{
			MethodName: "Invalidate",
			Handler:    _Sky_Invalidate_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _Sky_Sync_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _Sky_Find_Handler,
		},
		{
			MethodName: "Host",
			Handler:    _Sky_Host_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Notification",
			Handler:       _Sky_Notification_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Sign",
			Handler:       _Sky_Sign_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Interaction",
			Handler:       _Sky_Interaction_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RbtPay",
			Handler:       _Sky_RbtPay_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RbtPOS",
			Handler:       _Sky_RbtPOS_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pop.proto",
}
