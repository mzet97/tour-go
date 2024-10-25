package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Definição do tipo Image
type Image struct {
	width, height int
}

// Implementação do método ColorModel para Image
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Implementação do método Bounds para Image
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

// Implementação do método At para Image
// Aqui aplicamos uma função para gerar a cor de cada pixel
func (img Image) At(x, y int) color.Color {
	// Gera o valor v para os componentes R e G
	v := uint8((x + y) / 2)
	// Retorna uma cor RGBA onde R e G são iguais ao valor v, e B é fixo em 255
	return color.RGBA{v, v, 255, 255}
}

func main() {
	// Cria uma imagem com dimensões 256x256
	m := Image{width: 256, height: 256}
	// Exibe a imagem
	pic.ShowImage(m)
}
