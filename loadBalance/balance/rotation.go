package balance

import "github.com/pkg/errors"

type RotateBalance struct {
	currentIndex int
}

func init() {
	Register("rotation", &RotateBalance{})
}

func (r *RotateBalance) GetInstance(instances []*Instance, keys ...string) (*Instance, error) {
	size := len(instances)

	if size < 1 {
		return nil, errors.New("no instance")
	}

	r.currentIndex = (r.currentIndex + 1) % size
	return instances[r.currentIndex], nil
}
