package model

//Imovel representa as informações de imóvel/
type Imovel struct {
	X     int    `json:"coordenadaX"`
	Y     int    `json:"coordenadaY"`
	Nome  string `json:"Nome"`
	valor int
}

//SetValor é a função utilizada para manipular o valor da estrutura de Imóvel
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

//GetValor é a função que retorna o valor do Imóvel
func (i *Imovel) GetValor() int {
	return i.valor
}
