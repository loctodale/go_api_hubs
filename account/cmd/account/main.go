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
	//var r repository.Repository
	accountRepository, err := repository.NewAccountRepository()
	tokenRepository := repository.NewTokenRepository()
	if err != nil {
		log.Println(err)
	}

	//defer r.Close()

	u := utils.NewUtilsConfig()
	if err != nil {
		log.Println(err)
	}
	port := global.Config.AccountService.Ports.Local
	log.Println("Listening on port: ", port)
	
	s := service.NewAccountService(accountRepository, u, tokenRepository)
	log.Fatal(server.ListenGRPC(s, port))

}
