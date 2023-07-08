package main

import (
	"fmt"
)

func main() {
	name := "Patrick"
	version := 1.1

	fmt.Println("Olá sr.", name)
	fmt.Println("Este programa está na versão", version)

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	var command int

	fmt.Scan(&command)
	fmt.Println("O comando escolhido foi", command)

	if command == 1 {
		fmt.Println("Monitorando...")
	} else if command == 2 {
		fmt.Println("Exibindo logs...")
	} else if command == 0 {
		fmt.Println("Saindo do programa")
	} else {
		fmt.Println("Não conheço este comando")
	}
}
