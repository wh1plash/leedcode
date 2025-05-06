package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(i int) int {
	fib := make(map[int]int, 10)
	fib[0] = 0
	fib[1] = 1
	return func(i int) int {
		if _, ok := fib[i]; !ok {
			fib[i] = fib[i-2] + fib[i-1]
			return fib[i]
		}
		return fib[i]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
