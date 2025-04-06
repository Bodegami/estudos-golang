package main

import (
	"fmt"
	math "github.com/Bodegami/estudos-golang/tree/course/linuxtips-go-essentials/main/day04/calc/main/day04/calc/operations"
	"os"
	"strconv"
)

func main() {

	// slice (array? mais ou menos) uma lista de valores
	args := os.Args
	op := args[1]

	s1 := args[2]
	s2 := args[3]
	fmt.Println(args)

	//Converte os valores dos argumentos como strings para numericos
	n1, err := strconv.Atoi(s1)
	if err != nil {
		panic("valor invalido")
	}
	
	n2, err := strconv.Atoi(s2)
	if err != nil {
		panic("valor invalido")
	}

	fmt.Printf("Executando a operacao %s com os valores %d e %d\n", op, n1, n2)
	fmt.Println("Operacao: ", op)
	fmt.Println("Numero 1: ", n1)
	fmt.Println("Numero 2: ", n2)

	var result int
	var erro error

	switch op {
	case "+":
		result = math.Sum(n1, n2)
	case "-":
		result = math.Subtract(n1, n2)
	case "*":
		result = math.Multiply(n1, n2)
	case "/":
		result, erro = math.Divide(n1, n2)
	}

	//sempre verificar se o erro existe antes de tentar usar o valor de result
	if erro != nil {
		//panic para a execucao do programa
		panic(erro)
	}

	fmt.Println("Resultado ===== ", result)

}
