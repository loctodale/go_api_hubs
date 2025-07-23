package server

import (
	"context"
	"fmt"
	"github.com/loctodale/go_api_hubs_microservice/account/internal/logs"
	"github.com/loctodale/go_api_hubs_microservice/account/internal/service"
	pb "github.com/loctodale/go_api_hubs_microservice/account/pb/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type grpcServer struct {
	pb.UnimplementedAccountServiceServer
	pb.UnimplementedPrivateAccountServiceServer
	service service.Service
}

func ListenGRPC(s service.Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()

	pb.RegisterAccountServiceServer(serv, &grpcServer{service: s})
	pb.RegisterPrivateAccountServiceServer(serv, &grpcServer{service: s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostAccount(ctx context.Context, r *pb.PostAccountRequest) (*pb.PostAccountResponse, error) {
	err := s.service.PostAccount(r.UserAccount, r.UserPassword)
	if err != nil {
		return nil, err
	}
	return &pb.PostAccountResponse{}, nil
}

func (s *grpcServer) GetAccounts(ctx context.Context, request *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {
	res, err := s.service.GetAccounts()
	if err != nil {
		return nil, err
	}
	accounts := []*pb.Account{}
	for _, p := range res {
		accounts = append(accounts, &pb.Account{
			Id:   p.UserID.String(),
			Name: p.UserAccount,
		})
	}
	return &pb.GetAccountsResponse{Account: accounts}, nil
}

func (s *grpcServer) RegisterAccount(ctx context.Context, req *pb.RegisterAccountRequest) (*pb.BaseResponseMessage, error) {
	result, err := s.service.RegisterAccount(req.UserAccount)
	if err != nil {
		go logs.LogsAccountService(req, "POST", "/account/register", 500)
		return &result, nil
	}
	go logs.LogsAccountService(req, "POST", "/account/register", 200)
	return &pb.BaseResponseMessage{
		Message: fmt.Sprintf("User Account %s is registered", req.UserAccount),
		Code:    200,
	}, nil
}

func (s *grpcServer) VerifyAccount(ctx context.Context, req *pb.VerifyAccountRequest) (*pb.BaseResponseMessage, error) {
	request, err := s.service.VerifyAccount(req.Account, req.Otp)
	if err != nil {
		fmt.Println(err)
		go logs.LogsAccountService(req, "POST", "/account/verify", 500)
		return &request, nil
	}
	go logs.LogsAccountService(req, "POST", "/account/verify", 200)
	return &request, nil
}
func (s *grpcServer) Login(ctx context.Context, req *pb.LoginModel) (*pb.LoginResponse, error) {
	result, err := s.service.LoginAccount(req.Username, req.Password)
	if err != nil {
		fmt.Println(err)
		go logs.LogsAccountService(req, "POST", "/account/login", 500)
		return nil, err
	}
	go logs.LogsAccountService(req, "POST", "/account/login", 200)
	return &result, nil
}

func (s *grpcServer) GetAccount(ctx context.Context, req *pb.GetAccountByIdRequest) (*pb.Profile, error) {
	result, err := s.service.GetByUserId(req.GetId())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *grpcServer) VerifyProviderAccount(ctx context.Context, req *pb.VerifyAccountRequest) (*pb.BaseResponseMessage, error) {
	result, err := s.service.VerifyProviderAccount(req.Account, req.Otp)
	if err != nil {
		fmt.Println(err)
		go logs.LogsAccountService(req, "POST", "/account/verify/provider", 500)
		return &result, nil
	}
	go logs.LogsAccountService(req, "POST", "/account/verify/provider", 200)
	return &result, nil
}
