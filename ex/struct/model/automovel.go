package model

type Automovel struct {
	Ano    int
	Placa  string
	Modelo string
}

// Herança
type Moto struct {
	Automovel
	Cilindradas int
}

// Herança
type Carro struct {
	Automovel
	QuantidadeDePortas   int
	Potencia             int
	PossuiArCondicionado bool
}
