package service

import "github.com/segmentio/kafka-go"

type ISendMailFactory interface {
	BuildMessage(data kafka.Message) (subject, body, recipient string, err error)
}
