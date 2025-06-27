package initialize

import "github.com/loctodale/go_api_hubs_microservice/sendmail/internal/service/implements"

func Run() {
	LoadConfig()
	InitRedisServer()
	InitKafkaServer()
	implements.InitSendMailService()
}
