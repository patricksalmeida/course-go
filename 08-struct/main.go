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

	contaDaSabrina := ContaCorrente{
		"Sabrina Almeida",
		589,
		123457,
		130.5,
	}

	fmt.Println(contaDoPatrick, contaDaSabrina)
}
