package main

import (
	"fmt"

	"github.com.br/patricksalmeida/course-go/12-interfaces/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	fmt.Println(contaDoDenis.ObterSaldo())

	PagarBoleto(&contaDoDenis, 60)
	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaLuiza := contas.ContaCorrente{}
	contaDaLuiza.Depositar(500)

	PagarBoleto(&contaDaLuiza, -250)
	fmt.Println(contaDaLuiza.ObterSaldo())
}
