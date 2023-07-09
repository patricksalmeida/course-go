package main

import (
	"fmt"

	"github.com.br/patricksalmeida/course-go/11-composition/clientes"
	"github.com.br/patricksalmeida/course-go/11-composition/contas"
)

func main() {
	clienteBruno := clientes.Titular{
		Nome:      "Bruno",
		CPF:       "123.123.123-00",
		Profissao: "Desenvolvedor",
	}

	contaDoBruno := contas.ContaCorrente{
		Titular:       clienteBruno,
		NumeroAgencia: 123,
		NumeroConta:   123456,
	}

	contaDoBruno.Depositar(100.0)

	fmt.Println(contaDoBruno.ObterSaldo())

	contaDoBruno.Sacar(50.)

	fmt.Println(contaDoBruno.ObterSaldo())
}
