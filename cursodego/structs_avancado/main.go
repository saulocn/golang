package main

import (
	"encoding/json"
	"fmt"

	"github.com/saulocn/golang/cursodego/structs_avancado/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa azul"
	casa.X = 123
	casa.Y = 132

	casa.SetValor(159999)

	fmt.Printf("Casa é %+v\r\n", casa)
	fmt.Printf("O valor da casa é %d\r\n", casa.GetValor())
	objJSON, error := json.Marshal(casa)
	if error != nil {
		fmt.Printf("Erro ao ler o Json: %s\r\n", error)
	}
	fmt.Println("A casa em JSON:", string(objJSON))

}
