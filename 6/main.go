package main

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.
// Классические подходы: выход по условию, через канал уведомления, через контекст, прекращение работы runtime.Goexit() и др.
// Продемонстрируйте каждый способ в отдельном фрагменте кода.

// по условию
func first(n int) {
	// любое условие
	if n+1 > n {
		return
	}
}

// через канал уведомления
func second(done <-chan struct{}) {
	select {
	case <-done:
		return
	default:
		// блок
	}
}

// через контекст. выход либо через его отмену, либо через таймаут.
func third(ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	default:
		// блок
	}
}

// runtime.Goexit()
func fourth() {
	// останавливает конкретно эту горутину, применимо в особых случаях
	// для пула горутин, чтобы не завершать весь пул
	runtime.Goexit()
}

// по сигналу ос
func fifth() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT)
	<-quit
}

// через time.After
func sixth(timeLimit time.Duration) {
	select {
	case <- time.After(timeLimit):
		return
	default:
		// блок
	}
}

// через таймер
func seventh(timer *time.Timer) {
	select {
	case <- timer.C:
		return
	default:
		// блок
	}
}

func main() {

}
