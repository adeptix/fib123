package fib123

import (
	"fmt"
)

func Fib(n int) int {
	a := 0
	b := 1

	for range n {
		a, b = b, a+b
	}

	return a
}

func PrintFibs(count int) {
	if count < 1 {
		return
	}

	a := 0
	b := 1

	for range count {
		fmt.Println(a)
		a, b = b, a+b
	}
}
