package main

import (
	"sync"
)

type concMap struct {
	mu sync.RWMutex
	m  map[int]string
}

func (cm *concMap) put(id int, value string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.m[id] = value
}

func (cm *concMap) get(id int) (string, bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	v, ok := cm.m[id]

	return v, ok
}

func main() {
	m := &concMap{m: make(map[int]string)}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				key := (i * 100) + j%100
				m.put(key, "value-from-"+string(rune('0'+i)))
				m.get(key)
			}
		}(i)
	}

	wg.Wait()
}
