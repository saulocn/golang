package model

//Cidade representa cidade e estado do Brasil
type Cidade struct {
	Nome   string `json:"Nome"`
	Estado string `json:"Estado"`
}
