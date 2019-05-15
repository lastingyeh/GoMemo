package balance

import "fmt"

type BalanceManager struct {
	BalanceMap map[string]LoadBalancer
}

var mgr BalanceManager

func init() {
	mgr.BalanceMap = make(map[string]LoadBalancer)
}

func (b *BalanceManager) registerBalancer(name string, balancer LoadBalancer) {
	b.BalanceMap[name] = balancer
}

func Register(name string, balancer LoadBalancer) {
	mgr.registerBalancer(name, balancer)
}

func UseBalance(name string, instances []*Instance) (*Instance, error) {
	balancer, ok := mgr.BalanceMap[name]

	if !ok {
		return nil, fmt.Errorf("not found %s", name)
	}
	fmt.Printf("use %s balance: ", name)
	return balancer.GetInstance(instances)
}
