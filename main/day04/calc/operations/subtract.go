package operations

import "fmt"

func subtract(a int, b int) int {
	result := a - b
	fmt.Printf("A subtração de %d e %d é: %d\n", 5, 10, result)
	return result
}

func Subtract(a int, b int) int {
	return subtract(a, b)
}
