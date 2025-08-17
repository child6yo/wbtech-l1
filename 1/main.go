package main

import "fmt"

// Human структура с произвольным набором полей и методов.
type Human struct {
	Age    int
	Weight float64
	Height float64
}

// PrintInfo произвольных метод.
func (h *Human) PrintInfo() {
	fmt.Printf("human age=%d, weight=%.2f, height=%.2f\n", h.Age, h.Weight, h.Height)
}

// Action встраивает Human с его методами.
type Action struct {
	Human
}

func main() {
	h := Human{32, 71.123, 180.123}
	a := Action{h}

	a.PrintInfo()
}
