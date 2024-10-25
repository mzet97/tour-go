package main

import "fmt"

// List representa uma lista encadeada simples que pode armazenar valores de qualquer tipo comparável.
type List[T comparable] struct {
	next *List[T]
	val  T
}

// Adiciona um novo valor ao final da lista.
func (l *List[T]) Append(value T) {
	if l == nil {
		return
	}

	// Se o próximo nó for nulo, cria um novo nó.
	if l.next == nil {
		l.next = &List[T]{val: value}
	} else {
		// Caso contrário, avança para o próximo nó e chama Append recursivamente.
		l.next.Append(value)
	}
}

// Retorna o comprimento da lista.
func (l *List[T]) Length() int {
	if l == nil {
		return 0
	}
	count := 1
	current := l
	for current.next != nil {
		count++
		current = current.next
	}
	return count
}

// Exibe todos os valores na lista.
func (l *List[T]) Display() {
	current := l
	for current != nil {
		fmt.Printf("%v -> ", current.val)
		current = current.next
	}
	fmt.Println("nil")
}

// Remove o primeiro valor correspondente encontrado na lista.
func (l *List[T]) Remove(value T) {
	if l == nil || l.next == nil {
		return
	}

	// Remove o próximo valor se ele corresponder.
	if l.next.val == value {
		l.next = l.next.next
	} else {
		l.next.Remove(value)
	}
}

func main() {
	// Criando uma lista encadeada de inteiros
	list := &List[int]{val: 1}
	list.Append(2)
	list.Append(3)
	list.Append(4)

	// Exibe a lista
	fmt.Println("Lista antes da remoção:")
	list.Display()

	// Exibe o comprimento da lista
	fmt.Printf("Comprimento da lista: %d\n", list.Length())

	// Removendo um valor
	list.Remove(3)

	// Exibe a lista após remoção
	fmt.Println("Lista após remoção do valor 3:")
	list.Display()

	// Exibe o comprimento da lista
	fmt.Printf("Comprimento da lista: %d\n", list.Length())
}
