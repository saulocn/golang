package main

import "fmt"

var endereco string // endereco = ""
var telefone = "9999-9999"
var quantidade int // quantidade = 0
var comprou bool   // comprou = false
var valor float64  // valor = 0.00
var palavras rune

func main() {
	fmt.Printf("endereco: %s\r\n", endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
}
