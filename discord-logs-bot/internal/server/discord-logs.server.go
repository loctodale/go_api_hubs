package server

import (
	"fmt"
	"github.com/loctodale/go_api_hubs_microservice/discord-logs-bot/internal/service"
	pb "github.com/loctodale/go_api_hubs_microservice/discord-logs-bot/pb/discord_logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	pb.RegisterDiscordLogsServiceServer(serv, grpcServer{service: s})
	reflection.Register(serv)
	return serv.Serve(lis)
}
