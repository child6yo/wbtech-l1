package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на вход строку.
// Например: при вводе строки «главрыба» вывод должен быть «абырвалг».
// Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.),
// то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).

func reverseString(s string) string {
	chars := []rune(s)
	left, right := 0, len(chars)-1

	for left < right {
		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--
	}

	return string(chars)
}

func main() {
	fmt.Println(reverseString("главрыба"))
}
