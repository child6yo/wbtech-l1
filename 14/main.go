package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в runtime способна определить тип переменной, переданной в неё (на вход подаётся interface{}).
// Типы, которые нужно распознавать: int, string, bool, chan (канал).
// Подсказка: оператор типа switch v.(type) поможет в решении.

// через рефлексию
func typeChecker(in interface{}) {
    switch v := reflect.TypeOf(in); {
    case v.Kind() == reflect.Int: 
        fmt.Println("int")
    case v.Kind() == reflect.String:
        fmt.Println("string")
    case v.Kind() == reflect.Bool:
        fmt.Println("bool")
    case v.Kind() == reflect.Chan:
        fmt.Println("chan")
    default:
        fmt.Println("unknown")
    }
}

func main() {
	typeChecker(12)
	typeChecker("twenty")
	typeChecker(true)
	typeChecker(make(chan int))
	typeChecker(make(chan string))
}
