package main

import (
	"encoding/json"
	"fmt"

	"github.com/saulocn/golang/cursodego/erro/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa azul"
	casa.X = 123
	casa.Y = 132

	if err := casa.SetValor(10000001); err != nil {
		fmt.Println("[main] Houve um erro na atribuíção de valor:", err.Error())
		if err == model.ErrValorDeImovelMuitoAlto {
			fmt.Println("Corretor verifique sua avaliação!", err.Error())
		}
		return
	}

	fmt.Printf("Casa é %+v\r\n", casa)
	fmt.Printf("O valor da casa é %d\r\n", casa.GetValor())
	objJSON, err := json.Marshal(casa)
	if err != nil {
		fmt.Println("[main] Houve um erro na geração do objeto JSON", err.Error())
		return
	}
	fmt.Println("A casa em JSON:", string(objJSON))

}
