package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Patrick"
	age := 26
	version := 1.1
	fmt.Println("Olá sr.", name, "sua idade é", age)
	fmt.Println("Este programa está na versão", version)
	fmt.Println("O tipo da variável version é", reflect.TypeOf(version))
}
