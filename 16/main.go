package main

import "fmt"

// Реализовать алгоритм быстрой сортировки массива встроенными средствами языка.
// Можно использовать рекурсию.

// Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел.
// Для выбора опорного элемента можно взять середину или первый элемент.

func quickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	left, right := 0, len(array)-1
	pivot := (left + right) / 2

	array[pivot], array[right] = array[right], array[pivot]

	for i := range array {
		if array[i] < array[right] {
			array[i], array[left] = array[left], array[i]
			left++
		}
	}

	array[left], array[right] = array[right], array[left]

	quickSort(array[:left])
	quickSort(array[left+1:])

	return array
}

func main() {
	array := []int{2,1,4,5,3,6,4,2,1,10,7}
	fmt.Println(quickSort(array))
}
