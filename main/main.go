package main

import "fmt"

func main() {
	fmt.Println("Hello World") //String

	fmt.Println(2) //Numericos

	fmt.Println(2 + 5) // Expressoes

	fmt.Println(false) //Booleanos

	//Para fazer a interpolacao, utilizamos o Printf
	fmt.Printf("Hello, I'm %v and I'm an Integer\n", 2)
	fmt.Printf("Hello, I'm %v and I'm an String\n", "texto")
	fmt.Printf("Hello, I'm %v and I'm an Boolean\n", true)

	fmt.Printf("Hello, I'm %d and I'm an Integer\n", 2)
}
