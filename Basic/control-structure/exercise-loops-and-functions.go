package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	precision := 0.0000001
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteração %d: z = %v\n", i+1, z)
		if math.Abs(z*z-x) < precision {
			break
		}
	}
	return z
}

func main() {

	for x := 1; x <= 10; x++ {
		result := Sqrt(float64(x))
		fmt.Printf("Raiz aproximada de %d: %v\n", x, result)
		fmt.Printf("Raiz exata de %d: %v\n", x, math.Sqrt(float64(x)))
	}
}
