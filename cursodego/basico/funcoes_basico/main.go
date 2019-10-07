package main

import (
	"fmt"

	"github.com/saulocn/golang/cursodego/funcoes_basico/matematica"
)

func main() {
	resultado := matematica.Multiplicador(2, 3)
	fmt.Printf("O resultado da multiplicação é %d\r\n", resultado)
	resultado = matematica.Soma(2, 3)
	fmt.Printf("O resultado da soma é %d\r\n", resultado)
	resultado = matematica.Calculo(matematica.Multiplicador, 2, 3)
	fmt.Printf("O resultado da multiplicação é %d\r\n", resultado)
	resultado = matematica.Calculo(matematica.Soma, 2, 3)
	fmt.Printf("O resultado da soma é %d\r\n", resultado)
	resultado, resto := matematica.Divisor(14, 3)
	fmt.Printf("O resultado da divisão é %d e o resto é %d\r\n", resultado, resto)
}
