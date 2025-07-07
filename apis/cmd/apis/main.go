package main

import (
	"github.com/loctodale/go_api_hubs_microservice/apis/global"
	"github.com/loctodale/go_api_hubs_microservice/apis/internal/initialize"
	"github.com/loctodale/go_api_hubs_microservice/apis/internal/repository"
	"github.com/loctodale/go_api_hubs_microservice/apis/internal/server"
	"github.com/loctodale/go_api_hubs_microservice/apis/internal/service"
	"log"
)

func main() {
	initialize.Run()

	port := global.Config.ApisService.Ports.Local
	log.Println("Listening on port: ", port)
	//s := service.NewAccountService(accountRepository, u, tokenRepository)
	//log.Fatal(server.ListenGRPC(s, port))
	apisRepository := repository.NewApiRepository()
	apisPlanRepository := repository.NewApiPlanRepository()
	apisKeyRepository := repository.NewApisKeyRepository()
	apisService := service.NewApiService(apisRepository, apisPlanRepository, apisKeyRepository)
	apisPlanService := service.NewApisPlanService(apisPlanRepository, apisRepository)
	apisKeyService := service.NewApisKeyService(apisKeyRepository, apisRepository)
	log.Fatal(server.ListenGRPC(apisService, apisPlanService, apisKeyService, port))
}
