package model

//Galinha representa uma ave do tipo Galinha
type Galinha interface {
	Cacarejar() string
}

//Pata representa uma ave do tipo Pato
type Pata interface {
	Quack() string
}

//Ave representa um animal
type Ave struct {
	Nome string
}

//Cacarejar retorna um som emitido por uma Galinha
func (a Ave) Cacarejar() string {
	return "cocoricó"
}

//Quack retorna um som emitido por uma Pata
func (a Ave) Quack() string {
	return "Quack, quack.."
}
