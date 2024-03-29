package main

import "fmt"

var (
	//Endereco é um valor importante e tem de ser público
	Endereco string
	//Telefone é um valor importante para a aula
	Telefone            string  // Variável Endereço é pública enquanto telefone é privada
	quantidade, estoque int     // quantidade = 0
	comprou             bool    // comprou = false
	valor               float64 // valor = 0.00
	palavras            rune
)

func main() {
	teste := "Valor de Teste"
	fmt.Printf("endereco: %s\r\n", Endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("palavras: %v\r\n", palavras)
	fmt.Printf("O valor de teste é: %s\r\n", teste)
}
