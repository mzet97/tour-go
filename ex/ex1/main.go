package main

import (
	"bufio"
	"ex1/model"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	no := model.NewNo()

	for {
		clearScreen()
		fmt.Println("Bem-vindo ao lista de compras: ")
		fmt.Println("1 - Adicionar item")
		fmt.Println("2 - Remover item")
		fmt.Println("3 - Listar itens")
		fmt.Println("4 - Pesquisar item")
		fmt.Println("5 - Sair")
		fmt.Print("Escolha uma opção: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		fmt.Println("Sua resposta:", text)

		switch text {
		case "1":
			fmt.Print("Digite o nome do item: ")
			name, _ := reader.ReadString('\n')
			fmt.Print("Digite o preço do item: ")
			price, _ := reader.ReadString('\n')
			fmt.Print("Digite a quantidade do item: ")
			quantity, _ := reader.ReadString('\n')
			price = strings.TrimSpace(price)
			quantity = strings.TrimSpace(quantity)
			priceValue, priceErr := strconv.ParseFloat(price, 64)
			if priceErr != nil {
				fmt.Println("Erro ao converter para float64:", priceErr)
				return
			}
			quantityValue, quantityErr := strconv.ParseInt(quantity, 10, 32)
			if quantityErr != nil {
				fmt.Println("Erro ao converter para int:", quantityErr)
				return
			}
			quantityValueInt := int(quantityValue)
			item := model.Item{
				Name:     name,
				Price:    priceValue,
				Quantity: quantityValueInt,
			}
			no.Add(item)
		case "2":
			fmt.Print("Digite o nome do item: ")
			name, _ := reader.ReadString('\n')
			no.RemovByName(name)
		case "3":
			no.Display()
		case "4":
			fmt.Print("Digite o nome do item: ")
			name, _ := reader.ReadString('\n')
			item := no.FindByName(name)
			fmt.Println(item)
		case "5":
			os.Exit(0)
		}
	}
}
func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // Comando para Windows
	} else {
		cmd = exec.Command("clear") // Comando para Linux/macOS
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
