package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	numero := 3

	fmt.Println("O número", numero, "se escreve assim")

	switch numero {
	case 1:
		fmt.Println("Um")
	case 2:
		fmt.Println("Dois")
	case 3:
		fmt.Println("Três")
	}

	fmt.Println("Você é da familia do Unix!?")
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("Sim, sou da família Unix!")
	default:
		fmt.Println("Deixa essa pergunta pra lá!")
	}

	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Ok, você pode ficar até mais tarde!")
	default:
		fmt.Println("Vamos lá q é dia de trabalho!")
	}

	numero = 10
	fmt.Println("Este número cabe em um dígito?")

	switch {
	case numero < 10:
		fmt.Println("Sim!")
	case numero >= 10 && numero < 100:
		fmt.Println("Serve 2 dígitos?")
	case numero > 100:
		fmt.Println("Não dá, o número é grande demais!")
	}

}
