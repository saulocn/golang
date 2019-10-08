package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/saulocn/golang/cursodego/intermediario/defer/model"
)

func main() {
	arquivo, err := os.Open("cidades.csv")
	errorHandler("Houve um erro ao abrir o arquivo:", err)
	defer arquivo.Close()

	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()
	errorHandler("Houve um erro ao ler o arquivo com leitor CSV:", err)

	arquivoJson, err := os.Create("cidades.json")
	errorHandler("Houve um erro ao criar o arquivo JSON:", err)
	defer arquivoJson.Close()

	writter := bufio.NewWriter(arquivoJson)
	writter.WriteString("[\n\r")

	for _, linha := range conteudo {
		for indice, item := range linha {
			dados := strings.Split(item, "/")
			cidade := model.Cidade{}
			cidade.Nome = dados[0]
			cidade.Estado = dados[1]
			cidadeJSON, err := json.Marshal(cidade)
			errorHandler("Houve um erro ao gerar o JSON do item:"+string(indice), err)
			writter.WriteString("  " + string(cidadeJSON))
			if indice+1 < len(linha) {
				writter.WriteString(",\n")
			}

			fmt.Printf("O item[%d] é: %s\r\n", indice, item)
		}
	}
	writter.WriteString("\r\n]")
	writter.Flush()
}

func errorHandler(message string, err error) {
	if err != nil {
		fmt.Println(message, err.Error())
	}
}
