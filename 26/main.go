package main

import "fmt"

// Разработать программу, которая проверяет, что все символы в строке встречаются один раз (т.е. строка состоит из уникальных символов).
// Вывод: true, если все символы уникальны, false, если есть повторения. Проверка должна быть регистронезависимой, т.е. символы в разных регистрах считать одинаковыми.
// Например: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false.
// Подумайте, какой структурой данных удобно воспользоваться для проверки условия.

func isUnique(s string) bool {
	// лучше всего реализовать проверку через множество
	set := make(map[rune]struct{})
	chars := []rune(s)

	for _, c := range chars {
		if _, ok := set[c]; ok {
			return false
		}
		set[c] = struct{}{}
	}

	return true
}

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"

	fmt.Println(isUnique(s1))
	fmt.Println(isUnique(s2))
	fmt.Println(isUnique(s3))
}
