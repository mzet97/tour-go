package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	// Cria uma slice de slices com comprimento dy
	image := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		// Para cada linha, cria uma slice com comprimento dx
		image[y] = make([]uint8, dx)

		for x := 0; x < dx; x++ {
			// Atribui um valor à imagem. Neste caso, escolhemos a função x^y
			image[y][x] = uint8(x ^ y)
		}
	}

	return image
}

func main() {
	// Chama a função Pic para exibir a imagem
	pic.Show(Pic)
}
