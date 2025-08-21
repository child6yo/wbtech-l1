package main

import (
	"sync"
	"testing"
)

func runConcurrentIncrement(inc func(), goroutines int) {
	var wg sync.WaitGroup
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				inc()
			}
		}()
	}
	wg.Wait()
}

func Benchmark_Concurrent_Mutex_1Goroutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		counter := &muInc{}
		runConcurrentIncrement(counter.increment, 1)
	}
}

func Benchmark_Concurrent_Mutex_4Goroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		counter := &muInc{}
		runConcurrentIncrement(counter.increment, 4)
	}
}

func Benchmark_Concurrent_Mutex_8Goroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		counter := &muInc{}
		runConcurrentIncrement(counter.increment, 8)
	}
}

func Benchmark_Concurrent_Atomic_1Goroutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		counter := &atInc{}
		runConcurrentIncrement(counter.increment, 1)
	}
}

func Benchmark_Concurrent_Atomic_4Goroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		counter := &atInc{}
		runConcurrentIncrement(counter.increment, 4)
	}
}

func Benchmark_Concurrent_Atomic_8Goroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		counter := &atInc{}
		runConcurrentIncrement(counter.increment, 8)
	}
}
