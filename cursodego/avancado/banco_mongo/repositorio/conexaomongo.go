package repositorio

import "gopkg.in/mgo.v2"

//MongoSession armazena a sessão de conexão com o MongoDB
var MongoSession *mgo.Session

//OpenMongoSession é a função que abre a sessão com o MongoDB
func OpenMongoSession() (err error) {
	err = nil
	MongoSession, err = mgo.Dial("mongodb://localhost:27017/cursodego")
	return
}
