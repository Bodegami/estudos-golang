package main

import (
	"fmt"
	math "github.com/Bodegami/estudos-golang/tree/course/linuxtips-go-essentials/main/day04/calc/main/day04/calc/operations"
)

func main() {
	fmt.Println("Calculadora Go Essentials")

	result := math.Sum(5, 10)
	fmt.Println(result)

	result = math.Multiply(5, 10)
	fmt.Println(result)

	var err error
	result, err = math.Divide(5, 0)
	if err != nil {
		fmt.Printf("Nao consegui fazer a divisao. Error: %v\n", err)
	}
	fmt.Println(result)

	result = math.Subtract(5, 10)
	fmt.Println(result)

}
