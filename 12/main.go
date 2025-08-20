package main

import "fmt"

// Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.
// Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.

func main() {
	array := []string{"cat", "cat", "dog", "cat", "tree"}

	// реализуем множество на мапе
	set := make(map[string]struct{})
	for _, s := range array {
		set[s] = struct{}{}
	}

	res := []string{}
	for k := range set {
		res = append(res, k)
	}

	fmt.Println(res)
}