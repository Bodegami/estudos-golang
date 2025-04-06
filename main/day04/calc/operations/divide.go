package operations

import "fmt"

func divide(a int, b int) int {
	result := a / b
	fmt.Printf("A divisão de %d e %d é: %d\n", 5, 10, result)
	return result
}

func Divide(a int, b int) int {
	return divide(a, b)
}
