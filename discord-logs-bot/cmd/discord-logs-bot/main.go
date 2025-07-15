package main

import (
	"github.com/loctodale/go_api_hubs_microservice/discord-logs-bot/internal/initialize"
	"github.com/loctodale/go_api_hubs_microservice/discord-logs-bot/internal/server"
	"github.com/loctodale/go_api_hubs_microservice/discord-logs-bot/internal/service"
	"log"
)

func main() {
	initialize.Run()
	s := service.NewDiscordLogsService()
	log.Println("Discord logs service is Running")

	log.Fatal(server.ListenGrpcServer(s, 8080))
}
