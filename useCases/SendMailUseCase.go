package useCases

import (
	"GohCMS2/adapters/secondary/gateways"
)

type SendMailUseCase struct {
	mailRepository gateways.MailRepository
}

func NewSendMailUseCase() *SendMailUseCase {
	return &SendMailUseCase{}
}

func (g *SendMailUseCase) SendMail(receiverAddress string, templateName string, data interface{}) error {
	return g.mailRepository.Send(receiverAddress, templateName, data)
}
