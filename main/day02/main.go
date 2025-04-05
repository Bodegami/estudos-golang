package main

import (
	"fmt"
	"github.com/Bodegami/estudos-golang/tree/course/linuxtips-go-essentials/main/day02/printer"
)

func main() {
	//criando um modulo:
	// go mod init "github???"

	//criando um executavel:
	//go build . -o "name"

	fmt.Println("Hello World from main!")
	printer.PrintSomething()
}
