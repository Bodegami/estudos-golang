package operations

import "fmt"

func sum(a int, b int) int {
	result := a + b
	fmt.Printf("A soma de %d e %d é: %d\n", 5, 10, result)
	return a + b
}

func Sum(a int, b int) int {
	return sum(a, b)
}
