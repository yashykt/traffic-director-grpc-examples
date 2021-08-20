// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package wallet

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

// WalletClient is the client API for Wallet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletClient interface {
	// Unary RPC to fetch balancer of the account.
	FetchBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
	// Server streaming RPC to watch the balance of the account.
	WatchBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (Wallet_WatchBalanceClient, error)
}

type walletClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletClient(cc grpc.ClientConnInterface) WalletClient {
	return &walletClient{cc}
}

func (c *walletClient) FetchBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, "/grpc.examples.wallet.Wallet/FetchBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) WatchBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (Wallet_WatchBalanceClient, error) {
	stream, err := c.cc.NewStream(ctx, &Wallet_ServiceDesc.Streams[0], "/grpc.examples.wallet.Wallet/WatchBalance", opts...)
	if err != nil {
		return nil, err
	}
	x := &walletWatchBalanceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Wallet_WatchBalanceClient interface {
	Recv() (*BalanceResponse, error)
	grpc.ClientStream
}

type walletWatchBalanceClient struct {
	grpc.ClientStream
}

func (x *walletWatchBalanceClient) Recv() (*BalanceResponse, error) {
	m := new(BalanceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WalletServer is the server API for Wallet service.
// All implementations must embed UnimplementedWalletServer
// for forward compatibility
type WalletServer interface {
	// Unary RPC to fetch balancer of the account.
	FetchBalance(context.Context, *BalanceRequest) (*BalanceResponse, error)
	// Server streaming RPC to watch the balance of the account.
	WatchBalance(*BalanceRequest, Wallet_WatchBalanceServer) error
	mustEmbedUnimplementedWalletServer()
}

// UnimplementedWalletServer must be embedded to have forward compatible implementations.
type UnimplementedWalletServer struct {
}

func (UnimplementedWalletServer) FetchBalance(context.Context, *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchBalance not implemented")
}
func (UnimplementedWalletServer) WatchBalance(*BalanceRequest, Wallet_WatchBalanceServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchBalance not implemented")
}
func (UnimplementedWalletServer) mustEmbedUnimplementedWalletServer() {}

// UnsafeWalletServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServer will
// result in compilation errors.
type UnsafeWalletServer interface {
	mustEmbedUnimplementedWalletServer()
}

func RegisterWalletServer(s grpc.ServiceRegistrar, srv WalletServer) {
	s.RegisterService(&Wallet_ServiceDesc, srv)
}

func _Wallet_FetchBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).FetchBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.examples.wallet.Wallet/FetchBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).FetchBalance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_WatchBalance_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BalanceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WalletServer).WatchBalance(m, &walletWatchBalanceServer{stream})
}

type Wallet_WatchBalanceServer interface {
	Send(*BalanceResponse) error
	grpc.ServerStream
}

type walletWatchBalanceServer struct {
	grpc.ServerStream
}

func (x *walletWatchBalanceServer) Send(m *BalanceResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Wallet_ServiceDesc is the grpc.ServiceDesc for Wallet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wallet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.examples.wallet.Wallet",
	HandlerType: (*WalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchBalance",
			Handler:    _Wallet_FetchBalance_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchBalance",
			Handler:       _Wallet_WatchBalance_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/grpc/examples/wallet/wallet.proto",
}