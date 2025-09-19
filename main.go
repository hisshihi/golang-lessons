package main

import "fmt"

// создать обычный калькулятор, где iota - это операции + - * /

const (
	ADD = iota
	SUB
	MUL
	DIV
)

type Operation int

func (op Operation) Calculate(a, b float64) float64 {
	switch op {
	case ADD:
		return a + b
	case SUB:
		return a - b
	case MUL:
		return a * b
	case DIV:
		if b != 0 {
			return a / b
		}
		panic("division by zero")
	default:
		panic("unknown operation")
	}
}

func main() {
	var op Operation
	op = ADD
	fmt.Println(op.Calculate(5, 3))
	op = SUB
	fmt.Println(op.Calculate(5, 3))
	op = MUL
	fmt.Println(op.Calculate(5, 3))
	op = DIV
	fmt.Println(op.Calculate(5, 3))
}
