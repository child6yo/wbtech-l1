package main

import "fmt"

// Разработать программу, которая переворачивает порядок слов в строке.
// Пример: входная строка:
// «snow dog sun», выход: «sun dog snow».
// Считайте, что слова разделяются одиночным пробелом. Постарайтесь не использовать дополнительные срезы, а выполнять операцию «на месте».

func reverseWords(s string) string {
	chars := reverseString([]rune(s))

	start := 0
	for i := 0; i <= len(chars); i++ {
		if i == len(chars) || chars[i] == ' ' {
			reverseString(chars[start:i])
			start = i + 1
		}
	}

	return string(chars)
}

func reverseString(chars []rune) []rune {
	left, right := 0, len(chars)-1

	for left < right {
		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--
	}

	return chars
}

func main() {
	fmt.Println(reverseWords("sun dog snow cat"))
}
