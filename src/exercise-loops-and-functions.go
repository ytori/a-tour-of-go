package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {

	z, threshold := 1.0, 0.000001

	for {
		if d := (z*z - x) / (2 * z); d*d < threshold*threshold {
			return z
		} else {
			z -= d
		}
	}
}

func main() {
	fmt.Println(Sqrt(1000))
}
