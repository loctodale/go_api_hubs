package server

import (
	"context"
	"fmt"
	"github.com/loctodale/go_api_hubs_microservice/account/internal/service"
	pb "github.com/loctodale/go_api_hubs_microservice/account/pb/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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
	fmt.Println("Start register")
	result, err := s.service.RegisterAccount(req.UserAccount)
	if err != nil {
		return &result, nil
	}
	return &pb.BaseResponseMessage{
		Message: fmt.Sprintf("User Account %s is registered", req.UserAccount),
		Code:    200,
	}, nil
}

func (s *grpcServer) VerifyAccount(ctx context.Context, req *pb.VerifyAccountRequest) (*pb.BaseResponseMessage, error) {
	request, err := s.service.VerifyAccount(req.Account, req.Otp)
	if err != nil {
		fmt.Println(err)
		return &request, nil
	}
	return &request, nil
}
func (s *grpcServer) Login(ctx context.Context, req *pb.LoginModel) (*pb.LoginResponse, error) {
	result, err := s.service.LoginAccount(req.Username, req.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
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
	return nil, status.Errorf(codes.Unimplemented, "method VerifyProviderAccount not implemented")
}
