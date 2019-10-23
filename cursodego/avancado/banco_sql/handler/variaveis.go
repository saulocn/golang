package handler

import "html/template"

//ModeloOla armazana os modelos de página q serão utilizados pelos handlers na página ola
var ModeloOla = template.Must(template.ParseFiles("html/ola.html"))

//ModeloLocal armazana os modelos de página q serão utilizados pelos handlers na página local
var ModeloLocal = template.Must(template.ParseFiles("html/local.html"))
