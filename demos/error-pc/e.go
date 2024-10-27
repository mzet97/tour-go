package main

import (
	"fmt"
	"math"
)

func main() {
	// Iniciando com um valor grande e diminuindo até não ser mais detectável
	epsilon := 1.0

	// Loop até que 1 + epsilon não seja mais distinguível de 1
	for (1.0 + epsilon) != 1.0 {
		epsilon /= 2.0
	}

	// O último epsilon ainda distinguível de 1 será o dobro do valor atual
	epsilon *= 2.0

	fmt.Printf("Epsilon da máquina (float64): %g\n", epsilon)
	fmt.Printf("Comparando com math.SmallestNonzeroFloat64: %g\n", math.SmallestNonzeroFloat64)
}
