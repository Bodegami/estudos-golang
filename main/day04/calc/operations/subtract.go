package operations

import "fmt"

func subtract(a int, b int) int {
	result := a - b
	fmt.Printf("A subtração de %d e %d é: %d\n", a, b, result)
	return result
}

func Subtract(a int, b int) int {
	return subtract(a, b)
}
