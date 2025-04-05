package main

// programa escopo 1 (inicial / maior escopo)

import "fmt"

const i float64 = 5.5
const pi float64 = 3.1415926

var x string

// main - escopo 2 (menor / mais limitado)
func main() {
	x := "Hello World!" // criação/declaração e atribuição a x
	fmt.Println(x)      // referencia a x

	var y string       // criação/declaração a y
	y = "Hello World!" // atribuição a y
	fmt.Println(y)     // referencia a y

	var z, a, b int = 1, 2, 3 // criação/declaração e atribuição de n variáveis
	fmt.Println(z)
	fmt.Println(a)
	fmt.Println(b)

	c, d := 1, 2      // criação/declaração e atribuição de n variáveis
	fmt.Println(c, d) // impressao de duas variáveis na mesma linha

	// estaticamente tipado + fortemente tipado
	var e, f int
	e = 1
	fmt.Println(e)
	fmt.Println(f)

	e = 2
	fmt.Println(e)

	var g bool
	g = true
	fmt.Println(g)

	// constante criadas e atribuidas fora da funcao
	fmt.Println(i)
	fmt.Println(pi)

	// const declarada dentro da funcao
	const constInsideBlock string = "Hello World constInsideBlock!"

	//variavel criada fora da funcao, mas atribuida dentro da funcao
	x = "Hello World X!"

	helloWorld()
}

func helloWorld() {
	x := "Hello World!" // criação/declaração e atribuição a x
	fmt.Println(x)      // referencia a x

	var y string       // criação/declaração a y
	y = "Hello World!" // atribuição a y
	fmt.Println(y)     // referencia a y

	//nao tem acesso à constante de outra funcao
	//fmt.Println(constInsideBlock)
}
