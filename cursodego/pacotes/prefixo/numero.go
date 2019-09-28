package prefixo

import "strconv"

//Capital representa o número do prefixo da Capital de um Estado / Província
var Capital = 11

var teste = "teste"

//TesteComPrefixo é um teste de variável q só é exposta para outros pacotes quando pública (nome começa com letra maiúscula)
var TesteComPrefixo = teste + " " + strconv.Itoa(Capital)

// Dependência Cíclica não é permitida em Go
//NomeOperadoraPrefixoInterior é o nome da operadora e seu prefixo do Interior
//var NomeOperadoraPrefixoInterior = operadora.NomeOperadora + "33"
