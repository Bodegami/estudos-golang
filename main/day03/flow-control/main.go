package main

import (
	"fmt"
)

const MAX int = 100

func main() {
	fmt.Println("1 - Oi, estou em uma execucao 100% linear")
	fmt.Println("2 - Oi, estou em uma execucao 100% linear")
	fmt.Println("3 - Oi, estou em uma execucao 100% linear")

	var n int
	n = 0
	expr := n < MAX

	if expr {
		fmt.Printf("%d menor que MAX == %d\n", n, MAX)
	} else {
		fmt.Printf("%d maior que MAX == %d\n", n, MAX)
	}

	for n < MAX {
		expr := n%2 == 0 //expressao que retorna verdadeiro se for par e falso se for impar
		if expr {
			fmt.Printf("%d numero par\n", n)
		} else {
			fmt.Printf("%d numero impar\n", n)
		}
		n += 1
	}

	// for (inicialicacao); (condicao que precisa ser verdadeira); (incremento)
	for i := 0; i < MAX; i++ {
	}

	n = 0
	for {
		n = n + 1
		fmt.Println(n)
		if n > MAX {
			break
		}
	} // loop infinito (nao existe while em Go, apenas for infinito)

	switch n < MAX {
	case true:
		fmt.Println("n é menor do que max")
	case false:
		fmt.Println("n nao é menor do que max")
	}

	var j interface{} = "42" // você pode mudar o valor para testar diferentes tipos

	switch j.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64, complex64, complex128:
		fmt.Println("j é do tipo numérico")
	default:
		fmt.Println("j não é do tipo numérico")
	}

}
