package operations

import "fmt"

func sum(a int, b int) int {
	result := a + b
	fmt.Printf("A soma de %d e %d é: %d\n", a, b, result)
	return result
}

func Sum(a int, b int) int {
	return sum(a, b)
}
