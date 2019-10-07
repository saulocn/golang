package main

import (
	"fmt"

	"github.com/saulocn/golang/cursodego/structs_avancado/model"
)

var cache map[string]model.Imovel

func main() {
	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa Azul"
	casa.X = 123
	casa.Y = 132

	casa.SetValor(159999)
	cache["Casa Azul"] = casa

	fmt.Println("Há uma casa Azul no cache?")

	imovel, achou := cache["Casa Azul"]
	if achou {
		fmt.Printf("Sim, e o que achei foi %+v", imovel)
	}

	apartamento := model.Imovel{}
	apartamento.Nome = "Apartamento Amarelo"
	apartamento.X = 132
	apartamento.Y = 165

	apartamento.SetValor(72999)

	cache[apartamento.Nome] = apartamento
	fmt.Println("Quantos itens há no cache?!", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave {%s} = %+v\r\n", chave, imovel)
	}

	imovel, achou = cache["Casa Azul"]
	if achou {
		delete(cache, imovel.Nome)
	}
	fmt.Println("Quantos itens há no cache?!", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave {%s} = %+v\r\n", chave, imovel)
	}
}
