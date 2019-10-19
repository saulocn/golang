package main

import (
	"fmt"
	"net/http"

	"github.com/saulocn/golang/cursodego/avancado/servidor_web/handler"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá mundo!")
	})
	http.HandleFunc("/funcao", handler.Funcao)
	http.HandleFunc("/ola", handler.Ola)
	fmt.Println("Servidor disponível ...")
	http.ListenAndServe(":8081", nil)
}
