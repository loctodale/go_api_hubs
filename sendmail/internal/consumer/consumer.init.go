package consumer

import "fmt"

func InitConsumerSendMail() {
	service := NewOTPConsumer()
	err := service.ConsumeOTP()
	if err != nil {
		fmt.Println(err)
	}
}
