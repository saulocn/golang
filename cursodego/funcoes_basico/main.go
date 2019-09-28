package main

import (
	"fmt"

	"github.com/saulocn/golang/cursodego/funcoes_basico/matematica"
)

func main() {
	resultado := multiplicador(2, 3)
	fmt.Printf("O resultado da multiplicação é %d\r\n", resultado)
	resultado = matematica.Soma(2, 3)
	fmt.Printf("O resultado da soma é %d\r\n", resultado)
	resultado = matematica.Calculo(multiplicador, 2, 3)
	fmt.Printf("O resultado da multiplicação é %d\r\n", resultado)
	resultado = matematica.Calculo(matematica.Soma, 2, 3)
	fmt.Printf("O resultado da soma é %d\r\n", resultado)
}

func multiplicador(x int, y int) int {
	return x * y
}
