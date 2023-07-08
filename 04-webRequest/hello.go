package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	showIntroduction()

	showMenu()

	command := readUserCommand()

	switch command {
	case 1:
		startMonitoring()
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}

func showIntroduction() {
	name := "Patrick"
	version := 1.1

	fmt.Println("Olá sr.", name)
	fmt.Println("Este programa está na versão", version)
}

func showMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func readUserCommand() int {
	var command int

	fmt.Scan(&command)
	fmt.Println("O comando escolhido foi", command)

	return command
}

func startMonitoring() {
	fmt.Println("Monitorando...")

	site := "https://www.alura.com.br"

	response, _ := http.Get(site)

	fmt.Println(response)
}
