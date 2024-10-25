package main

import "fmt"

// Função fibonacci retorna um closure que gera os números de Fibonacci
func fibonacci() func() int {
	// Variáveis para armazenar os dois últimos números de Fibonacci
	a, b := 0, 1

	// Retorna uma função que calcula o próximo número de Fibonacci
	return func() int {
		// Retorna o valor atual de 'a', que é o próximo número de Fibonacci
		next := a
		// Atualiza 'a' e 'b' para os próximos números da sequência
		a, b = b, a+b
		// Retorna o próximo número de Fibonacci
		return next
	}
}

func main() {
	// Cria um closure de Fibonacci
	fib := fibonacci()

	// Exibe os primeiros 10 números da sequência de Fibonacci
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}
