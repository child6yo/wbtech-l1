package main

import "fmt"

// Реализовать алгоритм бинарного поиска встроенными методами языка.
// Функция должна принимать отсортированный слайс и искомый элемент, возвращать индекс элемента или -1, если элемент не найден.
// Подсказка: можно реализовать рекурсивно или итеративно, используя цикл for.

func binarySearch(array []int, target int) int {
	left := 0
	right := len(array) - 1

	if target < array[left] || target > array[right] {
		return -1
	}

	for left <= right {
		middle := (left + right) / 2
		if target == array[middle] {
			return middle
		}
		if target < array[middle] {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return -1
}

func main() {
	array := []int{1, 1, 2, 2, 3, 4, 4, 5, 6, 7, 10}
	fmt.Println(binarySearch(array, 5))
	fmt.Println(binarySearch(array, 11))
}
