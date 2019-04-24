package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func testLockPerf(lock sync.Locker) {
	var count int32
	m := make(map[string]int)
	m["A"] = 123
	m["B"] = 456
	m["C"] = 789

	// write
	for i := 0; i < 2; i++ {
		go func(mp map[string]int) {
			lock.Lock()
			mp["C"] = rand.Intn(100)
			time.Sleep(10 * time.Millisecond)
			lock.Unlock()
		}(m)
	}

	// read
	for i := 0; i < 100; i++ {
		go func(mp map[string]int) {
			for {
				switch v := lock.(type) {
				case *sync.Mutex:
					v.Lock()
					//fmt.Println(m["C"])
					time.Sleep(time.Millisecond)
					lock.Unlock()
				case *sync.RWMutex:
					v.RLock()
					//fmt.Println(m["C"])
					time.Sleep(time.Millisecond)
					v.RUnlock()
				}
				atomic.AddInt32(&count, 1)
			}
		}(m)
	}
	time.Sleep(3 * time.Second)
	// read count
	fmt.Println(atomic.LoadInt32(&count))
}

func main() {
	var lock sync.Mutex
	var rwLock sync.RWMutex
	// during 3 secs
	testLockPerf(&lock) // 2132 times exec
	testLockPerf(&rwLock) // 213350 times exec
}
