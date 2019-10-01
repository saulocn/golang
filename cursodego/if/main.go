package main

import "fmt"

func main() {
	meses := 0
	situacao := true
	cidade := "Teste"

	// > < != ==  <= >= && ||
	if meses <= 6 {
		fmt.Println("Esse credor deve a pouco tempo!")
	}

	if situacao {
		fmt.Println("Esse credor está devendo!")
	}

	if !situacao {
		fmt.Println("Esse credor está adimplente!")
	}

	if cidade == "Teste" {
		fmt.Println("O cliente vive no estado de Teste!")
	}

	// variáveis só valem para o escopo do if
	// é importante para economia de memória
	if descricao, status := haQuantoTempoEstaDevendo(meses); status {
		fmt.Println("Qual a situação do cliente:", descricao)
		return
	}
	// as variáveis acima valem para o escopo do if
	// a instrução abaixo não compila
	//fmt.Println("Me passa o status", descricao)
	fmt.Println("Obrigado por nos consultar!")
}

func haQuantoTempoEstaDevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente está devendo!"
		return
	}
	descricao = "O cliente está com o carnê em dia!"
	return
}
