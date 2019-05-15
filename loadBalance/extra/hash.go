package extra

import (
	"fmt"
	"github.com/pkg/errors"
	"go_dev/loadBalance/balance"
	"hash/crc32"
	"math/rand"
)

type HashBalance struct{}

func init() {
	balance.Register("hash", &HashBalance{})
}

func (h *HashBalance) GetInstance(instances []*balance.Instance, keys ...string) (*balance.Instance, error) {
	var key string
	if len(keys) == 0 {
		key = fmt.Sprintf("%d", rand.Int())
	}

	size := len(instances)

	if size < 1 {
		return nil, errors.New("no instance")
	}

	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(key), crcTable)
	index := int(hashVal) % size
	return instances[index], nil
}
