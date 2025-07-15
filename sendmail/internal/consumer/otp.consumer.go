package consumer

import (
	"context"
	"fmt"
	"github.com/loctodale/go_api_hubs_microservice/sendmail/global"
	"github.com/loctodale/go_api_hubs_microservice/sendmail/internal/service/implements"
	"github.com/segmentio/kafka-go"
	"log"
	"net/smtp"
	"os"
	"os/signal"
	"syscall"
)

type OTPConsumer struct {
	reader *kafka.Reader
	ctx    context.Context
}

func NewOTPConsumer() *OTPConsumer {
	return &OTPConsumer{
		reader: global.KafkaReader,
		ctx:    global.Ctx,
	}
}

func (c *OTPConsumer) ConsumeOTP() error {
	log.Println("[OTPConsumer] Starting OTP Consumer...")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	consumer := global.KafkaReader
	ctx, cancel := context.WithCancel(c.ctx)
	defer cancel()
	defer consumer.Close()

	sendMailFactory := implements.NewSendmailFactory()
	config := global.Config.SendmailService.Email

	for {
		select {
		case <-sigChan:
			log.Println("Shutting down Kafka consumer...")
			return nil

		default:
			msg, err := consumer.ReadMessage(ctx)
			if err != nil {
				log.Println("OTP Consumer ReadMessage error:", err)
				continue
			}

			registry, err := sendMailFactory.GetRegistry("otp")
			if err != nil {
				log.Println("OTP Factory GetRegistry error:", err)
				continue
			}

			subject, body, recipient, err := registry.BuildMessage(msg)
			if err != nil {
				log.Println("OTP BuildMessage error:", err)
				continue
			}
			fmt.Println("From: " + global.Config.SendmailService.Email.Account + "\r\n" +
				"To: " + recipient + "\r\n" +
				"Subject: " + subject + "\r\n" +
				"MIME-Version: 1.0\r\n" +
				"Content-Type: text/plain; charset=\"utf-8\"\r\n" +
				"\r\n" +
				body + "\r\n")
			message := []byte(fmt.Sprintf("From: " + global.Config.SendmailService.Email.Account + "\r\n" +
				"To: " + recipient + "\r\n" +
				"Subject: " + subject + "\r\n" +
				"MIME-Version: 1.0\r\n" +
				"Content-Type: text/plain; charset=\"utf-8\"\r\n" +
				"\r\n" +
				body + "\r\n"))

			auth := smtp.PlainAuth("", config.Account, config.Password, config.Host)
			addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
			err = smtp.SendMail(addr, auth, config.Account, []string{recipient}, message)
			if err != nil {
				log.Println("OTP SendMail error:", err)
			} else {
				log.Printf("OTP email sent to %s", recipient)
			}
		}
	}
}
