package repositorio

import (
	// "github.com/go-sql-driver/mysql" não é usado diretamente pela aplicação
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//Db é a variável q armazena a conexão com o banco de dados
var Db *sqlx.DB

//OpenSQLConnection é a função que inicializa a conexão com o banco de dados
func OpenSQLConnection() (err error) {
	err = nil
	Db, err = sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/cursodego")
	if err != nil {
		return
	}
	err = Db.Ping()
	return
}
