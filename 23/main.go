package main

import "fmt"

// Удалить i-ый элемент из слайса.
// Продемонстрируйте корректное удаление без утечки памяти.
// Подсказка: можно сдвинуть хвост слайса на место удаляемого элемента (copy(slice[i:], slice[i+1:])) и уменьшить длину слайса на 1.

func main() {
	i := 3

	slice := []int{0,1,2,3,4,5,6,7,8,9}
	copy(slice[i:], slice[i+1:])
	slice = slice[:len(slice)-1]
	fmt.Println(slice, cap(slice), len(slice))
}