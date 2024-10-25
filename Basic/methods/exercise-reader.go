package main

import "golang.org/x/tour/reader"

// Definição do tipo MyReader
type MyReader struct{}

// Implementação do método Read para MyReader
func (r MyReader) Read(b []byte) (int, error) {
	// Preenche o slice 'b' com o caractere ASCII 'A'
	for i := range b {
		b[i] = 'A'
	}
	// Retorna o número de bytes lidos e um erro nil
	return len(b), nil
}

func main() {
	// Valida a implementação do MyReader
	reader.Validate(MyReader{})
}
