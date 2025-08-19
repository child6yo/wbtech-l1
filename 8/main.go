package main

import "fmt"

// Дана переменная типа int64. Разработать программу, которая устанавливает i-й бит этого числа в 1 или 0.
// Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).
// Подсказка: используйте битовые операции (|, &^).

// setBit устанавливает i-й бит числа n в значение 1 или 0
// bitIndex — номер бита (считается с 0, справа)
// value — 0 или 1
func setBit(n int64, bitIndex int, value int) int64 {
	if value == 1 {
		return n | (1 << bitIndex)
	} else {
		return n &^ (1 << bitIndex)
	}
}

func main() {
	var n int64 = 5
	ix := 0
	val := 0

	fmt.Println(setBit(n, ix, val))
}
