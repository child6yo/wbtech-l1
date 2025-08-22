package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).

type BigIntArithmetic struct {
	a *big.Int
	b *big.Int
}

func New(x, y string) *BigIntArithmetic {
	a, _ := new(big.Int).SetString(x, 10)
	b, _ := new(big.Int).SetString(y, 10)
	return &BigIntArithmetic{a: a, b: b}
}

func (b *BigIntArithmetic) Multiplie() {
	mul := new(big.Int)
	fmt.Println(mul.Mul(b.a, b.b))
}

func (b *BigIntArithmetic) Divide() {
	div := new(big.Int)
	fmt.Println(div.Div(b.a, b.b))
}

func (b *BigIntArithmetic) Add() {
	add := new(big.Int)
	fmt.Println(add.Add(b.a, b.b))
}

func (b *BigIntArithmetic) Subtract() {
	sub := new(big.Int)
	fmt.Println(sub.Sub(b.a, b.b))
}

func main() {
	bia := New("999999999999999999999999", "10000000000000000000000")
	bia.Multiplie()
	bia.Divide()
	bia.Add()
	bia.Subtract()
}
