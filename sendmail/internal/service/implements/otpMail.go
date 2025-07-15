package implements

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"strconv"
)

type OTPMail struct {
	reader kafka.Reader
	ctx    context.Context
}

type OTPMessage struct {
	OTP   int    `json:"otp"`
	Email string `json:"email"`
}

func NewOTPMail(reader kafka.Reader) *OTPMail {
	return &OTPMail{
		reader: reader,
		ctx:    context.Background(),
	}
}
func (o *OTPMail) BuildMessage(data kafka.Message) (subject, body, recipient string, err error) {

	//defer func() {
	//	err := o.reader.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}()
	//
	//m, err := o.reader.ReadMessage(o.ctx)
	//if err != nil {
	//	return "", "", "", nil
	//}
	//var payload OTPMessage
	var payload OTPMessage
	err = json.Unmarshal(data.Value, &payload)
	fmt.Println("Data::", payload)
	if err != nil {
		log.Println("err:: ", err)
		return "", "", "", nil
	}

	subject = "OTP Verify Account"
	body = "Your verification code is " + strconv.Itoa(payload.OTP)

	return subject, body, payload.Email, nil
}
