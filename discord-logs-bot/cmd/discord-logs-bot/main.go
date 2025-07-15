package main

import (
	"github.com/loctodale/go_api_hubs_microservice/discord-logs-bot/global"
	"github.com/loctodale/go_api_hubs_microservice/discord-logs-bot/internal/initialize"
	"github.com/loctodale/go_api_hubs_microservice/discord-logs-bot/internal/server"
	"github.com/loctodale/go_api_hubs_microservice/discord-logs-bot/internal/service"
	"log"
)

func main() {
	initialize.Run()
	port := global.Config.AccountService.Ports.Local
	log.Println("Listening on port: ", port)
	s := service.NewDiscordLogsService()
	log.Fatal(server.ListenGrpcServer(s, 8080))

}
