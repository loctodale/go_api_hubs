package server

import (
	"context"
	"fmt"
	"github.com/loctodale/go_api_hubs_microservice/discord-logs-bot/internal/service"
	pb "github.com/loctodale/go_api_hubs_microservice/discord-logs-bot/pb/discord_logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"
)

type grpcServer struct {
	pb.UnimplementedDiscordLogsServiceServer
	service service.DiscordLogsService
}

func ListenGrpcServer(s service.DiscordLogsService, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterDiscordLogsServiceServer(serv, &grpcServer{service: s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) SendLogsMessage(ctx context.Context, req *pb.SendLogsMessageRequest) (*pb.SendLogsMessageResponse, error) {
	fmt.Println("Start logs message")
	err := s.service.SendMessage(req.Service, req.Message)
	if err != nil {
		fmt.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.SendLogsMessageResponse{
		BaseResponse: &pb.BaseResponse{
			Code:    200,
			Message: "Success",
		},
	}, nil
}

func (s *grpcServer) SendLogsTracking(ctx context.Context, req *pb.SendLogsTrackingRequest) (*pb.SendLogsTrackingResponse, error) {
	err := s.service.SendTrackingLogs(req.Service, req.Message, req.Code, req.Method, req.Path)
	if err != nil {
		fmt.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.SendLogsTrackingResponse{
		Message: "Success",
		Code:    200,
	}, nil
}
