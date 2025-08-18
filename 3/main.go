package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
)

// Реализовать постоянную запись данных в канал (в главной горутине).
// Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.
// Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.

// WorkerPool реализует пул воркеров
type WorkerPool struct {
	c  chan string
	wg sync.WaitGroup
}

// New создает новый WorkerPool.
func New(numWorkers int, c chan string) *WorkerPool {
	wp := WorkerPool{
		c: c,
	}

	// для ожидания завершения
	wp.wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		// создадим горутины по указанному количеству numWorkers
		go func() {
			// забираем задачи из канала
			for w := range wp.c {
				// обрабатываем
				fmt.Println(w)
			}
			// после закрытия канала нужно оповестить наш пул
			wp.wg.Done()
		}()
	}

	return &wp
}

// Shutdown завершает работу пула воркеров.
func (wp *WorkerPool) Shutdown() {
	// закроем канал с задачами
	close(wp.c)
	// дождемся завершения работы уже запущенных задач
	wp.wg.Wait()
}

func main() {
	var numWorkers int
	flag.IntVar(&numWorkers, "wrk", 5, "num workers in worker pool")
	flag.Parse()

	c := make(chan string, numWorkers)
	wp := New(numWorkers, c)

	log.Printf("worker pool started with %d goroutines", numWorkers)

	for {
		var s string
		fmt.Scanln(&s)

		if s == "exit" {
			wp.Shutdown()
			break
		}

		c <- s
	}
}
