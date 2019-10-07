package main

import "fmt"

func main() {
	var teste [3]int
	teste[0] = 3
	teste[1] = 2
	teste[2] = 1
	fmt.Println("Qual a capacidade desse Array?", len(teste))

	cantores := [2]string{"Michael Jackson", "John Lennon"}

	fmt.Printf("O que há nesse array? \n\r%v\n\r", cantores)

	capitais := [...]string{"Lisboa", "Brasilia", "Luanda", "Maputo"}
	fmt.Println("Qual a capacidade desse Array?", len(teste))

	for indice, cidade := range capitais {
		fmt.Printf("Capital[%d]: %s\r\n", indice, cidade)
	}
}
