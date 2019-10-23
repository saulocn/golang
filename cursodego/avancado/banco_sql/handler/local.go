package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/saulocn/golang/cursodego/avancado/banco_sql/model"
	"github.com/saulocn/golang/cursodego/avancado/banco_sql/repositorio"
)

//Local é a função de handler para a página local.html
func Local(w http.ResponseWriter, r *http.Request) {
	local := model.Local{}
	codigoTelefone, err := strconv.Atoi(r.URL.Path[7:])
	if err != nil {
		http.Error(w, "Não foi enviado um número válido na requisição. Favor verificar!", http.StatusBadRequest)
		fmt.Println("[Local] - Erro na conversão do número", err.Error())
		return
	}

	sql := "SELECT country, city, telcode FROM cursodego.place WHERE telcode = ?;"
	rows, err := repositorio.Db.Queryx(sql, codigoTelefone)
	if err != nil {
		http.Error(w, "Não foi possível pesquisar este código!", http.StatusInternalServerError)
		fmt.Println("[Local] - Erro na busca do local", sql, err.Error())
		return
	}
	for rows.Next() {
		err = rows.StructScan(&local)
		if err != nil {
			http.Error(w, "Não foi possível pesquisar este código!", http.StatusInternalServerError)
			fmt.Println("[Local] - Erro no binding do local", err.Error())
			return
		}
	}
	if err := ModeloLocal.ExecuteTemplate(w, "local.html", local); err != nil {
		http.Error(w, "Houve um erro na renderização da página!", http.StatusInternalServerError)
		fmt.Println("[Local] - Erro na execução do modelo", err.Error())
	}
	sql = "INSERT INTO logquery(daterequest) VALUES (?)"

	resultado, err := repositorio.Db.Exec(sql, time.Now().Format("02/01/2006 15:04:05"))
	if err != nil {
		fmt.Println("[Local] - Erro na inclusão do Log", err.Error(), " - ", sql)
	}

	linhasAfetadas, err := resultado.RowsAffected()
	if err != nil {
		fmt.Println("[Local] - Erro ao pegar o número de linhas afetadas", err.Error(), " - ", sql)
	}
	fmt.Println("Sucesso! Linhas afetadas: ", linhasAfetadas)
}
