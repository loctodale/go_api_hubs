package main

import (
	"github.com/loctodale/go_api_hubs_microservice/sendmail/internal/consumer"
	"github.com/loctodale/go_api_hubs_microservice/sendmail/internal/initialize"
)

func main() {
	initialize.Run()
	//s := service.ISendMailFactory
	//port := global.Config.SendmailService.Ports.Local
	consumer.InitConsumerSendMail()
}
