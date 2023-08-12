package gateways

type IPageRepository interface {
	Get(name string, data interface{}) ([]byte, error)
}
