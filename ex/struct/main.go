package main

import (
	"fmt"
	model "struct/model"
	"time"
)

func main() {
	fmt.Println("Iniciando...")

	endereco := model.Endereco{
		Rua:    "Rua 1",
		Numero: 123,
		Cidade: "Cidade 1",
	}

	pessoa := model.Pessoa{
		Nome:             "Luiz",
		DataDeNascimento: time.Date(1990, 11, 25, 0, 0, 0, 0, time.Local),
		Endereco:         endereco,
	}
	pessoa.CalculaIdade()
	fmt.Println(pessoa)

	automovelMoto := model.Automovel{
		Ano:    2020,
		Placa:  "ABC-1234",
		Modelo: "Veiculo 1",
	}

	fmt.Println("Automovel Moto: ", automovelMoto)

	moto := model.Moto{
		Automovel:   automovelMoto,
		Cilindradas: 125,
	}

	fmt.Println("Moto: ", moto)
}
