package main

import (
	"github.com/loctodale/go_api_hubs_microservice/account/global"
	"github.com/loctodale/go_api_hubs_microservice/account/internal/initialize"
	"github.com/loctodale/go_api_hubs_microservice/account/internal/repository"
	"github.com/loctodale/go_api_hubs_microservice/account/internal/server"
	"github.com/loctodale/go_api_hubs_microservice/account/internal/service"
	"github.com/loctodale/go_api_hubs_microservice/account/utils"
	"log"
)

func main() {
	initialize.Run()
	var r repository.Repository
	r, err := repository.NewAccountRepository()
	if err != nil {
		log.Println(err)
	}

	defer r.Close()

	u := utils.NewUtils()
	if err != nil {
		log.Println(err)
	}
	port := global.Config.AccountService.Ports.Local
	log.Println("Listening on port: ", port)
	s := service.NewAccountService(r, u)
	log.Fatal(server.ListenGRPC(s, port))

}
