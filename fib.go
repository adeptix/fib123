package fib123

import (
	"fmt"
)

// Fib returns n fibonacci number
func Fib(n int) int {
	a := 0
	b := 1

	for range n {
		a, b = b, a+b
	}

	return a
}

// PrintFibs prints fibonacci numbers
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

// IsNFib checks if number is n-th fibonacci number. See Fib
func IsNFib(n int, number int) bool {
	return Fib(n) == number
}
