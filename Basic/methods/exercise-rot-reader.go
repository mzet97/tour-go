package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// Implementação do método Read para rot13Reader
func (r *rot13Reader) Read(b []byte) (int, error) {
	// Lê o conteúdo do io.Reader subjacente
	n, err := r.r.Read(b)
	if err != nil {
		return n, err
	}

	// Aplica a cifra ROT13 aos caracteres lidos
	for i := 0; i < n; i++ {
		// Se o caractere for uma letra minúscula
		if b[i] >= 'a' && b[i] <= 'z' {
			// Aplica o ROT13, com wrap-around
			b[i] = 'a' + (b[i]-'a'+13)%26
		} else if b[i] >= 'A' && b[i] <= 'Z' { // Se o caractere for uma letra maiúscula
			// Aplica o ROT13, com wrap-around
			b[i] = 'A' + (b[i]-'A'+13)%26
		}
	}

	return n, nil
}

func main() {
	// String codificada com ROT13
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	// Cria um rot13Reader que vai decodificar o fluxo
	r := rot13Reader{s}
	// Copia a saída decodificada para o stdout
	io.Copy(os.Stdout, &r)
}
