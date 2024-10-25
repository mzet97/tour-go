package main

import (
	"fmt"
)

// Definição do tipo ErrNegativeSqrt
type ErrNegativeSqrt float64

// Implementação do método Error() para ErrNegativeSqrt
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Função Sqrt modificada para lidar com números negativos
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		// Retorna o erro quando o número é negativo
		return 0, ErrNegativeSqrt(x)
	}
	// Implementação básica de Sqrt usando o método de Newton
	z := x
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	// Testando com números positivos e negativos
	fmt.Println(Sqrt(2))  // Deve calcular a raiz quadrada de 2
	fmt.Println(Sqrt(-2)) // Deve retornar um erro
}
