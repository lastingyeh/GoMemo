package main

import (
	"fmt"
	"go_dev/loadBalance/balance"
	_ "go_dev/loadBalance/extra"
	"math/rand"
	"time"
)

// load-balance concept
// server start -> register to center -> notify client -> dynamic server instance selected.
// https://segmentfault.com/a/1190000016097418
var size = 10
var instances []*balance.Instance

func init() {
	for i := 0; i < size; i++ {
		instances = append(instances, balance.NewInstance(
			fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255)),
			rand.Intn(9000),
		))
	}
}

func main() {
	for {
		instance, err := balance.UseBalance("hash", instances)
		if err != nil {
			fmt.Println("getInstance err: ", err)
			continue
		}
		fmt.Println(instance)
		time.Sleep(time.Second)
	}
}
