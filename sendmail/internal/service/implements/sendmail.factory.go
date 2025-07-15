package implements

import (
	"fmt"
	"github.com/loctodale/go_api_hubs_microservice/sendmail/internal/service"
)

type SendmailFactory struct {
}

var strategyRegistry = map[string]service.ISendMailFactory{}

func NewSendmailFactory() *SendmailFactory {
	return &SendmailFactory{}
}
func (s *SendmailFactory) GetRegistry(mailType string) (service.ISendMailFactory, error) {
	strategy, ok := strategyRegistry[mailType]
	if !ok {
		return nil, fmt.Errorf("mail type not found: %s", mailType)
	}
	return strategy, nil
}

func (s *SendmailFactory) SendMailRegistryCreate(mailType string, factory service.ISendMailFactory) error {
	found, _ := strategyRegistry[mailType]
	if found != nil {
		return fmt.Errorf("mail type already exists: %s", mailType)
	}
	strategyRegistry[mailType] = factory
	return nil
}
