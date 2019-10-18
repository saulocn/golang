package model

//Usuario é a estrutura que representa um usuário do sistema
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}
