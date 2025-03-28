package gateways

type IPageRepository interface {
	Get(name string, data any) ([]byte, error)
}
