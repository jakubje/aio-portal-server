// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: service_aio_portal.proto

package pb

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

const (
	AioPortal_CreateUser_FullMethodName             = "/pb.AioPortal/CreateUser"
	AioPortal_LoginUser_FullMethodName              = "/pb.AioPortal/LoginUser"
	AioPortal_UpdateUser_FullMethodName             = "/pb.AioPortal/UpdateUser"
	AioPortal_VerifyEmail_FullMethodName            = "/pb.AioPortal/VerifyEmail"
	AioPortal_CreatePortfolio_FullMethodName        = "/pb.AioPortal/CreatePortfolio"
	AioPortal_UpdatePortfolio_FullMethodName        = "/pb.AioPortal/UpdatePortfolio"
	AioPortal_GetPortfolio_FullMethodName           = "/pb.AioPortal/GetPortfolio"
	AioPortal_ListPortfolios_FullMethodName         = "/pb.AioPortal/ListPortfolios"
	AioPortal_GetRollUp_FullMethodName              = "/pb.AioPortal/GetRollUp"
	AioPortal_DeletePortfolio_FullMethodName        = "/pb.AioPortal/DeletePortfolio"
	AioPortal_CreateTransaction_FullMethodName      = "/pb.AioPortal/CreateTransaction"
	AioPortal_GetTransaction_FullMethodName         = "/pb.AioPortal/GetTransaction"
	AioPortal_ListTransactions_FullMethodName       = "/pb.AioPortal/ListTransactions"
	AioPortal_ListTransactionsByCoin_FullMethodName = "/pb.AioPortal/ListTransactionsByCoin"
	AioPortal_CreateWatchlist_FullMethodName        = "/pb.AioPortal/CreateWatchlist"
	AioPortal_UpdateWatchlist_FullMethodName        = "/pb.AioPortal/UpdateWatchlist"
	AioPortal_GetWatchlist_FullMethodName           = "/pb.AioPortal/GetWatchlist"
	AioPortal_ListWatchlists_FullMethodName         = "/pb.AioPortal/ListWatchlists"
	AioPortal_DeleteWatchlist_FullMethodName        = "/pb.AioPortal/DeleteWatchlist"
	AioPortal_CreateCoin_FullMethodName             = "/pb.AioPortal/CreateCoin"
	AioPortal_AddWatchlistCoin_FullMethodName       = "/pb.AioPortal/AddWatchlistCoin"
	AioPortal_UpdateCoin_FullMethodName             = "/pb.AioPortal/UpdateCoin"
	AioPortal_ListWatchlistCoins_FullMethodName     = "/pb.AioPortal/ListWatchlistCoins"
)

// AioPortalClient is the client API for AioPortal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AioPortalClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error)
	CreatePortfolio(ctx context.Context, in *CreatePortfolioRequest, opts ...grpc.CallOption) (*CreatePortfolioResponse, error)
	UpdatePortfolio(ctx context.Context, in *UpdatePortfolioRequest, opts ...grpc.CallOption) (*UpdatePortfolioResponse, error)
	GetPortfolio(ctx context.Context, in *GetPortfolioRequest, opts ...grpc.CallOption) (*GetPortfolioResponse, error)
	ListPortfolios(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListPortfoliosResponse, error)
	GetRollUp(ctx context.Context, in *RollUpRequest, opts ...grpc.CallOption) (*RollUpResponse, error)
	DeletePortfolio(ctx context.Context, in *DeletePortfolioRequest, opts ...grpc.CallOption) (*DeletePortfolioResponse, error)
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error)
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error)
	ListTransactions(ctx context.Context, in *ListTransactionsRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error)
	ListTransactionsByCoin(ctx context.Context, in *ListTransactionsByCoinRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error)
	CreateWatchlist(ctx context.Context, in *CreateWatchlistRequest, opts ...grpc.CallOption) (*CreateWatchlistResponse, error)
	UpdateWatchlist(ctx context.Context, in *UpdateWatchlistRequest, opts ...grpc.CallOption) (*UpdateWatchlistResponse, error)
	GetWatchlist(ctx context.Context, in *GetWatchlistRequest, opts ...grpc.CallOption) (*GetWatchlistResponse, error)
	ListWatchlists(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetWatchlistsResponse, error)
	DeleteWatchlist(ctx context.Context, in *DeleteWatchlistRequest, opts ...grpc.CallOption) (*DeleteWatchlistResponse, error)
	CreateCoin(ctx context.Context, in *CreateCoinRequest, opts ...grpc.CallOption) (*CreateCoinResponse, error)
	AddWatchlistCoin(ctx context.Context, in *AddWatchlistCoinRequest, opts ...grpc.CallOption) (*AddWatchlistCoinResponse, error)
	UpdateCoin(ctx context.Context, in *UpdateCoinRequest, opts ...grpc.CallOption) (*UpdateCoinResponse, error)
	ListWatchlistCoins(ctx context.Context, in *ListWatchlistCoinsRequest, opts ...grpc.CallOption) (*ListWatchlistCoinsResponse, error)
}

type aioPortalClient struct {
	cc grpc.ClientConnInterface
}

func NewAioPortalClient(cc grpc.ClientConnInterface) AioPortalClient {
	return &aioPortalClient{cc}
}

func (c *aioPortalClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, AioPortal_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, AioPortal_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, AioPortal_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error) {
	out := new(VerifyEmailResponse)
	err := c.cc.Invoke(ctx, AioPortal_VerifyEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) CreatePortfolio(ctx context.Context, in *CreatePortfolioRequest, opts ...grpc.CallOption) (*CreatePortfolioResponse, error) {
	out := new(CreatePortfolioResponse)
	err := c.cc.Invoke(ctx, AioPortal_CreatePortfolio_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) UpdatePortfolio(ctx context.Context, in *UpdatePortfolioRequest, opts ...grpc.CallOption) (*UpdatePortfolioResponse, error) {
	out := new(UpdatePortfolioResponse)
	err := c.cc.Invoke(ctx, AioPortal_UpdatePortfolio_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) GetPortfolio(ctx context.Context, in *GetPortfolioRequest, opts ...grpc.CallOption) (*GetPortfolioResponse, error) {
	out := new(GetPortfolioResponse)
	err := c.cc.Invoke(ctx, AioPortal_GetPortfolio_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) ListPortfolios(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListPortfoliosResponse, error) {
	out := new(ListPortfoliosResponse)
	err := c.cc.Invoke(ctx, AioPortal_ListPortfolios_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) GetRollUp(ctx context.Context, in *RollUpRequest, opts ...grpc.CallOption) (*RollUpResponse, error) {
	out := new(RollUpResponse)
	err := c.cc.Invoke(ctx, AioPortal_GetRollUp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) DeletePortfolio(ctx context.Context, in *DeletePortfolioRequest, opts ...grpc.CallOption) (*DeletePortfolioResponse, error) {
	out := new(DeletePortfolioResponse)
	err := c.cc.Invoke(ctx, AioPortal_DeletePortfolio_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error) {
	out := new(CreateTransactionResponse)
	err := c.cc.Invoke(ctx, AioPortal_CreateTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error) {
	out := new(GetTransactionResponse)
	err := c.cc.Invoke(ctx, AioPortal_GetTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) ListTransactions(ctx context.Context, in *ListTransactionsRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error) {
	out := new(ListTransactionsResponse)
	err := c.cc.Invoke(ctx, AioPortal_ListTransactions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) ListTransactionsByCoin(ctx context.Context, in *ListTransactionsByCoinRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error) {
	out := new(ListTransactionsResponse)
	err := c.cc.Invoke(ctx, AioPortal_ListTransactionsByCoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) CreateWatchlist(ctx context.Context, in *CreateWatchlistRequest, opts ...grpc.CallOption) (*CreateWatchlistResponse, error) {
	out := new(CreateWatchlistResponse)
	err := c.cc.Invoke(ctx, AioPortal_CreateWatchlist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) UpdateWatchlist(ctx context.Context, in *UpdateWatchlistRequest, opts ...grpc.CallOption) (*UpdateWatchlistResponse, error) {
	out := new(UpdateWatchlistResponse)
	err := c.cc.Invoke(ctx, AioPortal_UpdateWatchlist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) GetWatchlist(ctx context.Context, in *GetWatchlistRequest, opts ...grpc.CallOption) (*GetWatchlistResponse, error) {
	out := new(GetWatchlistResponse)
	err := c.cc.Invoke(ctx, AioPortal_GetWatchlist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) ListWatchlists(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetWatchlistsResponse, error) {
	out := new(GetWatchlistsResponse)
	err := c.cc.Invoke(ctx, AioPortal_ListWatchlists_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) DeleteWatchlist(ctx context.Context, in *DeleteWatchlistRequest, opts ...grpc.CallOption) (*DeleteWatchlistResponse, error) {
	out := new(DeleteWatchlistResponse)
	err := c.cc.Invoke(ctx, AioPortal_DeleteWatchlist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) CreateCoin(ctx context.Context, in *CreateCoinRequest, opts ...grpc.CallOption) (*CreateCoinResponse, error) {
	out := new(CreateCoinResponse)
	err := c.cc.Invoke(ctx, AioPortal_CreateCoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) AddWatchlistCoin(ctx context.Context, in *AddWatchlistCoinRequest, opts ...grpc.CallOption) (*AddWatchlistCoinResponse, error) {
	out := new(AddWatchlistCoinResponse)
	err := c.cc.Invoke(ctx, AioPortal_AddWatchlistCoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) UpdateCoin(ctx context.Context, in *UpdateCoinRequest, opts ...grpc.CallOption) (*UpdateCoinResponse, error) {
	out := new(UpdateCoinResponse)
	err := c.cc.Invoke(ctx, AioPortal_UpdateCoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aioPortalClient) ListWatchlistCoins(ctx context.Context, in *ListWatchlistCoinsRequest, opts ...grpc.CallOption) (*ListWatchlistCoinsResponse, error) {
	out := new(ListWatchlistCoinsResponse)
	err := c.cc.Invoke(ctx, AioPortal_ListWatchlistCoins_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AioPortalServer is the server API for AioPortal service.
// All implementations must embed UnimplementedAioPortalServer
// for forward compatibility
type AioPortalServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error)
	CreatePortfolio(context.Context, *CreatePortfolioRequest) (*CreatePortfolioResponse, error)
	UpdatePortfolio(context.Context, *UpdatePortfolioRequest) (*UpdatePortfolioResponse, error)
	GetPortfolio(context.Context, *GetPortfolioRequest) (*GetPortfolioResponse, error)
	ListPortfolios(context.Context, *emptypb.Empty) (*ListPortfoliosResponse, error)
	GetRollUp(context.Context, *RollUpRequest) (*RollUpResponse, error)
	DeletePortfolio(context.Context, *DeletePortfolioRequest) (*DeletePortfolioResponse, error)
	CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error)
	GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error)
	ListTransactions(context.Context, *ListTransactionsRequest) (*ListTransactionsResponse, error)
	ListTransactionsByCoin(context.Context, *ListTransactionsByCoinRequest) (*ListTransactionsResponse, error)
	CreateWatchlist(context.Context, *CreateWatchlistRequest) (*CreateWatchlistResponse, error)
	UpdateWatchlist(context.Context, *UpdateWatchlistRequest) (*UpdateWatchlistResponse, error)
	GetWatchlist(context.Context, *GetWatchlistRequest) (*GetWatchlistResponse, error)
	ListWatchlists(context.Context, *emptypb.Empty) (*GetWatchlistsResponse, error)
	DeleteWatchlist(context.Context, *DeleteWatchlistRequest) (*DeleteWatchlistResponse, error)
	CreateCoin(context.Context, *CreateCoinRequest) (*CreateCoinResponse, error)
	AddWatchlistCoin(context.Context, *AddWatchlistCoinRequest) (*AddWatchlistCoinResponse, error)
	UpdateCoin(context.Context, *UpdateCoinRequest) (*UpdateCoinResponse, error)
	ListWatchlistCoins(context.Context, *ListWatchlistCoinsRequest) (*ListWatchlistCoinsResponse, error)
	mustEmbedUnimplementedAioPortalServer()
}

// UnimplementedAioPortalServer must be embedded to have forward compatible implementations.
type UnimplementedAioPortalServer struct {
}

func (UnimplementedAioPortalServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedAioPortalServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedAioPortalServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedAioPortalServer) VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyEmail not implemented")
}
func (UnimplementedAioPortalServer) CreatePortfolio(context.Context, *CreatePortfolioRequest) (*CreatePortfolioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePortfolio not implemented")
}
func (UnimplementedAioPortalServer) UpdatePortfolio(context.Context, *UpdatePortfolioRequest) (*UpdatePortfolioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePortfolio not implemented")
}
func (UnimplementedAioPortalServer) GetPortfolio(context.Context, *GetPortfolioRequest) (*GetPortfolioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPortfolio not implemented")
}
func (UnimplementedAioPortalServer) ListPortfolios(context.Context, *emptypb.Empty) (*ListPortfoliosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPortfolios not implemented")
}
func (UnimplementedAioPortalServer) GetRollUp(context.Context, *RollUpRequest) (*RollUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRollUp not implemented")
}
func (UnimplementedAioPortalServer) DeletePortfolio(context.Context, *DeletePortfolioRequest) (*DeletePortfolioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePortfolio not implemented")
}
func (UnimplementedAioPortalServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedAioPortalServer) GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedAioPortalServer) ListTransactions(context.Context, *ListTransactionsRequest) (*ListTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransactions not implemented")
}
func (UnimplementedAioPortalServer) ListTransactionsByCoin(context.Context, *ListTransactionsByCoinRequest) (*ListTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransactionsByCoin not implemented")
}
func (UnimplementedAioPortalServer) CreateWatchlist(context.Context, *CreateWatchlistRequest) (*CreateWatchlistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWatchlist not implemented")
}
func (UnimplementedAioPortalServer) UpdateWatchlist(context.Context, *UpdateWatchlistRequest) (*UpdateWatchlistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWatchlist not implemented")
}
func (UnimplementedAioPortalServer) GetWatchlist(context.Context, *GetWatchlistRequest) (*GetWatchlistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWatchlist not implemented")
}
func (UnimplementedAioPortalServer) ListWatchlists(context.Context, *emptypb.Empty) (*GetWatchlistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWatchlists not implemented")
}
func (UnimplementedAioPortalServer) DeleteWatchlist(context.Context, *DeleteWatchlistRequest) (*DeleteWatchlistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWatchlist not implemented")
}
func (UnimplementedAioPortalServer) CreateCoin(context.Context, *CreateCoinRequest) (*CreateCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoin not implemented")
}
func (UnimplementedAioPortalServer) AddWatchlistCoin(context.Context, *AddWatchlistCoinRequest) (*AddWatchlistCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWatchlistCoin not implemented")
}
func (UnimplementedAioPortalServer) UpdateCoin(context.Context, *UpdateCoinRequest) (*UpdateCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCoin not implemented")
}
func (UnimplementedAioPortalServer) ListWatchlistCoins(context.Context, *ListWatchlistCoinsRequest) (*ListWatchlistCoinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWatchlistCoins not implemented")
}
func (UnimplementedAioPortalServer) mustEmbedUnimplementedAioPortalServer() {}

// UnsafeAioPortalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AioPortalServer will
// result in compilation errors.
type UnsafeAioPortalServer interface {
	mustEmbedUnimplementedAioPortalServer()
}

func RegisterAioPortalServer(s grpc.ServiceRegistrar, srv AioPortalServer) {
	s.RegisterService(&AioPortal_ServiceDesc, srv)
}

func _AioPortal_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_VerifyEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).VerifyEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_VerifyEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).VerifyEmail(ctx, req.(*VerifyEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_CreatePortfolio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePortfolioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).CreatePortfolio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_CreatePortfolio_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).CreatePortfolio(ctx, req.(*CreatePortfolioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_UpdatePortfolio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePortfolioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).UpdatePortfolio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_UpdatePortfolio_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).UpdatePortfolio(ctx, req.(*UpdatePortfolioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_GetPortfolio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPortfolioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).GetPortfolio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_GetPortfolio_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).GetPortfolio(ctx, req.(*GetPortfolioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_ListPortfolios_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).ListPortfolios(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_ListPortfolios_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).ListPortfolios(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_GetRollUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).GetRollUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_GetRollUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).GetRollUp(ctx, req.(*RollUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_DeletePortfolio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePortfolioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).DeletePortfolio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_DeletePortfolio_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).DeletePortfolio(ctx, req.(*DeletePortfolioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_CreateTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_GetTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_ListTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).ListTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_ListTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).ListTransactions(ctx, req.(*ListTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_ListTransactionsByCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransactionsByCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).ListTransactionsByCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_ListTransactionsByCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).ListTransactionsByCoin(ctx, req.(*ListTransactionsByCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_CreateWatchlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWatchlistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).CreateWatchlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_CreateWatchlist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).CreateWatchlist(ctx, req.(*CreateWatchlistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_UpdateWatchlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWatchlistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).UpdateWatchlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_UpdateWatchlist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).UpdateWatchlist(ctx, req.(*UpdateWatchlistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_GetWatchlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWatchlistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).GetWatchlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_GetWatchlist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).GetWatchlist(ctx, req.(*GetWatchlistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_ListWatchlists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).ListWatchlists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_ListWatchlists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).ListWatchlists(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_DeleteWatchlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWatchlistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).DeleteWatchlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_DeleteWatchlist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).DeleteWatchlist(ctx, req.(*DeleteWatchlistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_CreateCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).CreateCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_CreateCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).CreateCoin(ctx, req.(*CreateCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_AddWatchlistCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddWatchlistCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).AddWatchlistCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_AddWatchlistCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).AddWatchlistCoin(ctx, req.(*AddWatchlistCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_UpdateCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).UpdateCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_UpdateCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).UpdateCoin(ctx, req.(*UpdateCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AioPortal_ListWatchlistCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWatchlistCoinsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AioPortalServer).ListWatchlistCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AioPortal_ListWatchlistCoins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AioPortalServer).ListWatchlistCoins(ctx, req.(*ListWatchlistCoinsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AioPortal_ServiceDesc is the grpc.ServiceDesc for AioPortal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AioPortal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AioPortal",
	HandlerType: (*AioPortalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _AioPortal_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _AioPortal_LoginUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _AioPortal_UpdateUser_Handler,
		},
		{
			MethodName: "VerifyEmail",
			Handler:    _AioPortal_VerifyEmail_Handler,
		},
		{
			MethodName: "CreatePortfolio",
			Handler:    _AioPortal_CreatePortfolio_Handler,
		},
		{
			MethodName: "UpdatePortfolio",
			Handler:    _AioPortal_UpdatePortfolio_Handler,
		},
		{
			MethodName: "GetPortfolio",
			Handler:    _AioPortal_GetPortfolio_Handler,
		},
		{
			MethodName: "ListPortfolios",
			Handler:    _AioPortal_ListPortfolios_Handler,
		},
		{
			MethodName: "GetRollUp",
			Handler:    _AioPortal_GetRollUp_Handler,
		},
		{
			MethodName: "DeletePortfolio",
			Handler:    _AioPortal_DeletePortfolio_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _AioPortal_CreateTransaction_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _AioPortal_GetTransaction_Handler,
		},
		{
			MethodName: "ListTransactions",
			Handler:    _AioPortal_ListTransactions_Handler,
		},
		{
			MethodName: "ListTransactionsByCoin",
			Handler:    _AioPortal_ListTransactionsByCoin_Handler,
		},
		{
			MethodName: "CreateWatchlist",
			Handler:    _AioPortal_CreateWatchlist_Handler,
		},
		{
			MethodName: "UpdateWatchlist",
			Handler:    _AioPortal_UpdateWatchlist_Handler,
		},
		{
			MethodName: "GetWatchlist",
			Handler:    _AioPortal_GetWatchlist_Handler,
		},
		{
			MethodName: "ListWatchlists",
			Handler:    _AioPortal_ListWatchlists_Handler,
		},
		{
			MethodName: "DeleteWatchlist",
			Handler:    _AioPortal_DeleteWatchlist_Handler,
		},
		{
			MethodName: "CreateCoin",
			Handler:    _AioPortal_CreateCoin_Handler,
		},
		{
			MethodName: "AddWatchlistCoin",
			Handler:    _AioPortal_AddWatchlistCoin_Handler,
		},
		{
			MethodName: "UpdateCoin",
			Handler:    _AioPortal_UpdateCoin_Handler,
		},
		{
			MethodName: "ListWatchlistCoins",
			Handler:    _AioPortal_ListWatchlistCoins_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_aio_portal.proto",
}
