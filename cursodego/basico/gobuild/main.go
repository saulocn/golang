package main

import (
	"encoding/json"
	"fmt"

	"github.com/saulocn/golang/cursodego/gobuild/model"
)

// Go build para buildar o projeto
// Para executá-lo em outras plataformas como Windows, basta digitar:
// GOOS=windows GOARCH=386 go build -o meuapp_windows.exe
// Para buildar para Linux, no caso raspberry, basta digitar:
// GOOS=linux GOARCH=arm go build -o meuapp_raspberrypi -v github.com/saulocn/golang/cursodego/gobuild
// O -v significa que se trata de modo verboso
func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa azul"
	casa.X = 123
	casa.Y = 132

	if err := casa.SetValor(1000000); err != nil {
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
