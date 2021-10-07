package main

import (
	"sync"
	"time"
)

func main() {
	var lock sync.RWMutex

	m := make(map[string]int)

	go func() {
		for {
			lock.Lock()
			m["a"] += 1
			lock.Unlock()

			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for {
			lock.RLock()
			_ = m["b"]
			lock.RUnlock()

			time.Sleep(time.Microsecond)
		}
	}()

	select {}
}
