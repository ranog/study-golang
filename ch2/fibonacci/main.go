package main

import "fmt"

func main() {
	const n = 50
	fmt.Printf("fibonacci(%d) = %d\n", n, fibonacci(n))
}

func fibonacci(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
