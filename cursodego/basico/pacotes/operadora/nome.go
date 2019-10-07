package operadora

import (
	"strconv"

	"github.com/saulocn/golang/cursodego/pacotes/prefixo"
)

//NomeOperadora é o Nome da Operadora utilizada
var NomeOperadora = "Blablabla Telecom"

//PrefixoOperadora é o Prefixo mais o nome da Operadora
var PrefixoOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora
