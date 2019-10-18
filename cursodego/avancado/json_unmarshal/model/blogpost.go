package model

//BlogPost armazena dados de um Post de um blog
type BlogPost struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
