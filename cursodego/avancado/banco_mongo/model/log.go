package model

//RegistroLog armazena os dados de vistita da página
type RegistroLog struct {
	DataVisita string `json:"dataVisita" db:"dataVisita" bson:"dataVisita"`
	Quem       string `json:"quem,omitempty" db:"quem" bson:"quem,omitempty"`
}
