package main

import (
	"fmt"

	"github.com/saulocn/golang/cursodego/intermediario/interface/model"
)

func main() {
	marilu := model.Ave{}
	marilu.Nome = "Marilu"
	queroAcordarComUmCacarejo(marilu)
	queroOuvirUmaPataNoLago(marilu)
}

func queroAcordarComUmCacarejo(galinha model.Galinha) {
	fmt.Println(galinha.Cacarejar())
}

func queroOuvirUmaPataNoLago(pata model.Pata) {
	fmt.Println(pata.Quack())
}
