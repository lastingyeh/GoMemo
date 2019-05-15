package balance

type LoadBalancer interface {
	GetInstance([]*Instance,...string) (*Instance, error)
}
