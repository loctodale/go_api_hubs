package initialize

import (
	"fmt"
	"github.com/loctodale/go_api_hubs_microservice/sendmail/global"
	"github.com/segmentio/kafka-go"
	"time"
)

func InitKafkaServer() {
	config := global.Config.SendmailService
	fmt.Println("init kafka server: ", config.Kafka.Address)
	k := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"kafka:9092"},
		GroupID:        "group-verify-otp",
		Topic:          "otp-auth-topic",
		CommitInterval: time.Second,
	})

	global.KafkaReader = k
}
