package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Реализовать структуру-счётчик, которая будет инкрементироваться в конкурентной среде (т.е. из нескольких горутин).
// По завершению программы структура должна выводить итоговое значение счётчика.
// Подсказка: вам понадобится механизм синхронизации, например, sync.Mutex или sync/Atomic для безопасного инкремента.

// на мьютексах
type muInc struct {
	x  int
	mu sync.Mutex
}

func (mi *muInc) increment() {
	mi.mu.Lock()
	defer mi.mu.Unlock()
	mi.x++
}

// на чистых атомиках
type atInc struct {
	x atomic.Int32
}

func (ai *atInc) increment() {
	ai.x.Add(1)
}

func main() {
	mi := muInc{x: 0}
	ai := atInc{}

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			mi.increment()
		}()
		go func() {
			defer wg.Done()
			ai.increment()
		}()
	}

	wg.Wait()

	fmt.Println(mi.x)
	fmt.Println(ai.x.Load())
}
