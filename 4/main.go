package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).
// Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.

// обычный воркер, в который передается контекст
func worker(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("goroutine %d stoped\n", n)
			return
		default:
			fmt.Printf("goroutine %d working\n", n)
			time.Sleep(8 * time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	// контекст с отменой позволяет завершить целую цепочку горутин
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			worker(ctx, i)
		}()
	}

	// создание канала для сигналов ос
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT) // syscall.SIGINT = Ctrl+C
	<-quit

	cancel()
	wg.Wait()
}