package implements

import "github.com/loctodale/go_api_hubs_microservice/sendmail/global"

func InitSendMailService() {
	// Initialize OTP Mail service
	otpMail := NewOTPMail(*global.KafkaReader)
	factory := NewSendmailFactory()
	err := factory.SendMailRegistryCreate("otp", otpMail)
	if err != nil {
		panic(err)
	}

}
