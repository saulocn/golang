package main

import (
	"fmt"

	"github.com/saulocn/golang/cursodego/pacotes/operadora"
	"github.com/saulocn/golang/cursodego/pacotes/prefixo"
)

//NomeDoUsuario é o nome do usuário do Sistema
var NomeDoUsuario = "Jeff"

func main() {
	fmt.Printf("O nome do usuário é %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital %d\r\n", prefixo.Capital)
	fmt.Printf("Operadora %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Prefixo Operadora %s\r\n", operadora.PrefixoOperadora)
	fmt.Printf("Teste com Prefixo %s\r\n", prefixo.TesteComPrefixo)
}
