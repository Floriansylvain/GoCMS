package gateways

type IMailRepository interface {
	Send(receiverAddress string, templateName string, data interface{}) error
}
