package handler

import "html/template"

//Modelos armazana os modelos de página q serão utilizados pelos handlers
var Modelos = template.Must(template.ParseFiles("html/ola.html"))
