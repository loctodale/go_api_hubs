package initialize

import (
	"fmt"
	"github.com/loctodale/go_api_hubs_microservice/account/global"
	"github.com/segmentio/kafka-go"
)

func InitKafkaServer() {
	config := global.Config.AccountService
	fmt.Println("kafka address: ", config.Kafka.Address)
	k := &kafka.Writer{
		Addr:  kafka.TCP("kafka:9092"),
		Topic: config.Kafka.Topic,
	}
	if k == nil {
		panic(fmt.Sprintf("kafka writer is nil"))
	}
	fmt.Println("kafka writer is ready")
	global.KafkaProducer = k
}
