package balance

import (
	"errors"
	"math/rand"
)

type RandomBalance struct{}

func init() {
	Register("random", &RandomBalance{})
}

func (r *RandomBalance) GetInstance(instances []*Instance, keys ...string) (*Instance, error) {
	if len(instances) < 1 {
		return nil, errors.New("no instance")
	}

	inst := rand.Intn(len(instances))
	return instances[inst], nil
}
