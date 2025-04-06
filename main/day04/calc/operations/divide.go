package operations

import (
	"errors"
	"fmt"
)

func divide(a int, b int) int {
	result := a / b
	fmt.Printf("A divisão de %d e %d é: %d\n", a, b, result)
	return result
}

//	panic e para a execucao completamente
//func Divide(a int, b int) (int, error) {
//	if b == 0 {
//		fmt.Println("Divisão por zero não é permitida.")
//		panic("Divisão por zero nao permitida")
//	}
//
//	return divide(a, b)
//}

func Divide(a int, b int) (int, error) {
	if b == 0 {
		err := errors.New("divisão por zero não é permitida")
		return 0, err
	}

	return divide(a, b), nil
}
