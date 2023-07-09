package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoPatrick := ContaCorrente{}
	contaDoPatrick.titular = "Patrick Almeida"
	contaDoPatrick.numeroAgencia = 589
	contaDoPatrick.numeroConta = 123456
	contaDoPatrick.saldo = 500

	fmt.Println(contaDoPatrick.saldo)
	fmt.Println(contaDoPatrick.Sacar(-100))
	fmt.Println(contaDoPatrick.saldo)
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}
