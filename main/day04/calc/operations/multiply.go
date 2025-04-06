package operations

import "fmt"

// funcoes que comecam com a letra minuscula == privada
// so sao visiveis dentro do pacote
func multiply(a int, b int) int {
	result := a * b
	fmt.Printf("A multiplicação de %d e %d é: %d\n", a, b, result)
	return result
}

// funcoes que comecam com a letra maiuscula == publica
func Multiply(a int, b int) int {
	return multiply(a, b)
}
