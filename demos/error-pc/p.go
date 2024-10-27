package main

import (
	"fmt"
	"math"
)

func main() {
	// Soma repetida de 0.1 pode gerar erros de ponto flutuante perceptíveis
	valorExato := 1.0 // Valor que esperamos alcançar
	valorCalculado := 0.0

	// Somar 0.1 repetidamente
	for i := 0; i < 10; i++ {
		valorCalculado += 0.1
	}

	// Calcula o erro absoluto
	erroAbsoluto := math.Abs(valorCalculado - valorExato)

	// Calcula o erro relativo (evitando divisão por zero)
	var erroRelativo float64
	if valorExato != 0 {
		erroRelativo = erroAbsoluto / math.Abs(valorExato)
	} else {
		erroRelativo = math.Inf(1)
	}

	// Exibe os resultados
	fmt.Printf("Valor exato: %v\n", valorExato)
	fmt.Printf("Valor calculado: %v\n", valorCalculado)
	fmt.Printf("Erro absoluto: %v\n", erroAbsoluto)
	fmt.Printf("Erro relativo: %v\n", erroRelativo)
}
