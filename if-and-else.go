package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {

	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// ここではv使える
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// ここではv使えない
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 10),
	)
}
