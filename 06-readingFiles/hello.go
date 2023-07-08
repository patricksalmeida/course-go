package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitorTimes = 3
const delayInSeconds = 1 * time.Second

func main() {
	showIntroduction()
	readSitesOfFile()

	for {
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

	// sites := []string{
	// 	"https://random-status-code.herokuapp.com",
	// 	"https://www.alura.com.br",
	// 	"https://www.caelum.com.br",
	// }

	sites := readSitesOfFile()

	for i := 0; i < monitorTimes; i++ {
		fmt.Println("Test", i+1)
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(delayInSeconds)
		fmt.Println("")
	}

}

func readSitesOfFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")
	// file, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)

	line, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(line)

	return sites
}

func testSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code:", response.StatusCode)
	}
}
