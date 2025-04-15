package server

import (
	"context"
	"fmt"
	"github.com/loctodale/go_api_hubs_microservice/account/internal/service"
	"github.com/loctodale/go_api_hubs_microservice/account/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type grpcServer struct {
	pb.UnimplementedAccountServiceServer
	service service.Service
}

func ListenGRPC(s service.Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterAccountServiceServer(serv, &grpcServer{service: s})
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

func (s *grpcServer) GetAccounts(ctx context.Context, request *pb.Empty) (*pb.GetAccountsResponse, error) {
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
