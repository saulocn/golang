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
	casa := Imovel{}
	fmt.Printf("A casa é: %+v\r\n", casa)

	apartamento := Imovel{17, 123, "Apartamento Local", 731004}
	fmt.Printf("O apartamento é: %+v\r\n", apartamento)

	chacara := Imovel{
		Y:     56,
		Nome:  "Chácara Interior",
		valor: 56123,
		X:     123,
	}
	fmt.Printf("A chácara é: %+v\r\n", chacara)

	casa.Nome = "Lar doce lar"
	casa.valor = 50123
	casa.X = 10
	casa.Y = 30

	fmt.Printf("A casa é: %+v\r\n", casa)

}
