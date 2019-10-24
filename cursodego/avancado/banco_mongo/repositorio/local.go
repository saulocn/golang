package repositorio

import (
	"github.com/saulocn/golang/cursodego/avancado/banco_mongo/model"
	"gopkg.in/mgo.v2/bson"
)

//GetLocal obtém um local do MongoDB
func GetLocal(codigoTelefone int) (local model.Local, err error) {
	sessao := MongoSession.Copy()
	defer sessao.Close()
	collection := sessao.DB("cursodego").C("local")
	err = collection.Find(bson.M{"telcode": codigoTelefone}).One(&local)
	return
}

//SaveLog registra a consulta ao local
func SaveLog(registro model.RegistroLog) (err error) {
	sessao := MongoSession.Copy()
	defer sessao.Close()
	collection := sessao.DB("cursodego").C("logvisitas")
	err = collection.Insert(registro)
	return
}
