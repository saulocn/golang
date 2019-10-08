package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	arquivo, err := os.Open("cidades.csv")
	if err != nil {
		fmt.Println("Houve um erro ao abrir o arquivo:", err.Error())
	}
	// scanner := bufio.NewScanner(arquivo)
	// for scanner.Scan() {
	// 	linha := scanner.Text()
	// 	fmt.Println("O conteúdo da linha é:", linha)
	// }
	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()
	if err != nil {
		fmt.Println("Houve um erro ao ler o arquivo com leitor CSV:", err.Error())
	}

	for indice, linha := range conteudo {
		fmt.Printf("A linha[%d] contém: %s\r\n", indice, linha)
		for indice, item := range linha {
			fmt.Printf("O item[%d] é: %s\r\n", indice, item)
		}
	}
	arquivo.Close()
}
