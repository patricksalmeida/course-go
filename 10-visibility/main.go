package main

import (
	"fmt"

	"./contas"
)

func main() {
	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

	status := contaDaSilvia.Transferir(-10000, &contaDoGustavo)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
}
