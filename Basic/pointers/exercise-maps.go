package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// Cria um map para armazenar as contagens
	wordMap := make(map[string]int)

	// Divide a string em palavras
	words := strings.Fields(s)

	// Para cada palavra, incrementa a contagem no map
	for _, word := range words {
		wordMap[word]++
	}

	return wordMap
}

func main() {
	// Executa os testes
	wc.Test(WordCount)
}
