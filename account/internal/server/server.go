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

func (s *grpcServer) GetAccount(ctx context.Context, request *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	result, err := s.service.GetAccount(request.UserAccount)
	if err != nil {
		return nil, err
	}
	userId := result.UserID.String()
	return &pb.GetAccountResponse{Account: &pb.Account{Id: userId, Name: result.UserAccount}}, nil
}
