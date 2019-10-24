package handler

import (
	"fmt"
	"net/http"
)

//Funcao é um handler http
func Funcao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Handler utilizando função num package")
}
