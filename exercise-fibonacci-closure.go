package main

import "fmt"

// fibonacci is a function that returns
// a function that return as int.
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		f := a
		a, b = b, f+b
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
