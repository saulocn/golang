package main

import (
	"fmt"
	"net/http"

	"github.com/saulocn/golang/cursodego/avancado/banco_mongo/handler"
	"github.com/saulocn/golang/cursodego/avancado/banco_mongo/repositorio"
)

func init() {
	fmt.Println("Vamos iniciar o servidor!")
}

func main() {
	err := repositorio.OpenSQLConnection()
	if err != nil {
		fmt.Println("Servidor não inicializado.\nErro ao se conectar ao banco de dados MariaDB", err.Error())
		return
	}

	err = repositorio.OpenMongoSession()
	if err != nil {
		fmt.Println("Servidor não inicializado.\nErro ao se conectar ao banco de dados MongoDB", err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá mundo!")
	})
	http.HandleFunc("/funcao", handler.Funcao)
	http.HandleFunc("/ola", handler.Ola)
	http.HandleFunc("/local/", handler.Local)
	fmt.Println("Servidor disponível ...")
	http.ListenAndServe(":8081", nil)
}
