package main

import (
	"github.com/loctodale/go_api_hubs_microservice/account/config"
	"github.com/loctodale/go_api_hubs_microservice/account/internal/repository"
	"github.com/loctodale/go_api_hubs_microservice/account/internal/server"
	"github.com/loctodale/go_api_hubs_microservice/account/internal/service"
	"github.com/loctodale/go_api_hubs_microservice/account/utils"
	"github.com/spf13/viper"
	"log"
)

func main() {

	config.LoadConfig()

	var r repository.Repository
	r, err := repository.NewPostgresRepository(viper.GetString("account_service.database.postgres"))
	if err != nil {
		log.Println(err)
	}

	defer r.Close()

	u := utils.NewUtils()
	if err != nil {
		log.Println(err)
	}
	port := viper.GetInt("account_service.port.local")
	log.Println("Listening on port: ", port)
	s := service.NewAccountService(r, u)
	log.Fatal(server.ListenGRPC(s, port))

}
