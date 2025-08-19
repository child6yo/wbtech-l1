package main

import (
	"fmt"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала – читать эти значения.
// По истечении N секунд программа должна завершаться.

// функция чтения значений из канала
func reader(c chan int) {
	for {
		v, ok := <-c
		if !ok {
			return
		}
		fmt.Println(v)
	}

}

// функция записи в канал
func writer(values []int, c chan int) {
	defer close(c)
	for _, v := range values {
		c <- v
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// создаем таймер
	timer := time.NewTimer(3 * time.Second)

	c := make(chan int)

	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	go func ()  {
		writer(values, c)
	}()
	go func() {
		reader(c)
	}()

	<-timer.C
}
