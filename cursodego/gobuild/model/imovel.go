package model

import "errors"

//Imovel representa as informações de imóvel/
type Imovel struct {
	X     int    `json:"coordenadaX"`
	Y     int    `json:"coordenadaY"`
	Nome  string `json:"Nome"`
	valor int
}

//ErrValorDeImovelMuitoAlto é um erro que é apresentado quando é atribuído ao imóvel um valor muito alto e fora do mercado
var ErrValorDeImovelMuitoAlto = errors.New("O valor informado é muito alto!")

//ErrValorDeImovelInvalido é um erro que é apresentado com é atribuído ao Imóvel um valor muito baixo
var ErrValorDeImovelInvalido = errors.New("O valor informado não é válido para o Imóvel")

//SetValor é a função utilizada para manipular o valor da estrutura de Imóvel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 10000000 {
		err = ErrValorDeImovelMuitoAlto
		return
	}
	i.valor = valor
	return
}

//GetValor é a função que retorna o valor do Imóvel
func (i *Imovel) GetValor() int {
	return i.valor
}
