package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoPatrick := ContaCorrente{
		titular:       "Patrick Almeida",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.5,
	}

	var contaDaSabrina *ContaCorrente
	contaDaSabrina = new(ContaCorrente)
	contaDaSabrina.titular = "Sabrina Almeida"
	contaDaSabrina.numeroAgencia = 589
	contaDaSabrina.numeroConta = 123457
	contaDaSabrina.saldo = 130.5

	fmt.Println(contaDoPatrick, contaDaSabrina)
}
