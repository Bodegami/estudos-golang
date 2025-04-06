package main

import "fmt"

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

}
