package main

import "fmt"

//Imovel é uma struct que armazena dados de um imóvel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	casa := new(Imovel)
	fmt.Printf("A casa é: %p - %+v \r\n", &casa, casa)

	chacara := Imovel{12, 23, "Chácara bonita!", 290000}
	fmt.Printf("A casa é: %p - %+v \r\n", &chacara, chacara)
	//Nesse caso é passado a posição de memória do objeto chácara que será alterado para a função
	//caso não seja passada com &, será passada uma cópia do objeto não alterando o antigo objeto
	mudaImovel(&chacara)
	fmt.Printf("A casa é: %p - %+v \r\n", &chacara, chacara)

}

func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 15
}
