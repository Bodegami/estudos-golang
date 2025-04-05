package printer

import "fmt"

// funcoes privadas comecam com a letra minuscula
func printSomething() {
	fmt.Println("Hello World from printer")
}

// funcoes publicas comecam com a letra maiscula
func PrintSomething() {
	printSomething()
}
