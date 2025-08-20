package main

import "fmt"

// Рассмотреть следующий код и ответить на вопросы: к каким негативным последствиям он может привести и как это исправить?
// Вопрос: что происходит с переменной justString?

// var justString string

// func someFunc() {
//   v := createHugeString(1 << 10) - создается новая строка в 1024 байт
//   justString = v[:100] - начинает ссылаться на v. таким образом происходит утечка памяти.
// }

// func main() {
//   someFunc()
// }

// Приведите корректный пример реализации.

var justString string

func someFunc() {
    v := createHugeString(1 << 10)
    justString = string([]byte(v[:100])) // копируем через слайс байт
}

func main() {
    someFunc()
}