package main

import "fmt"

func main() {
	fmt.Println("counting")

	// 先入れ後出しですよ
	for i := 0; i < 10; i += 1 {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
