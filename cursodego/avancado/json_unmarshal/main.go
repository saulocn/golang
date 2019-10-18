package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/saulocn/golang/cursodego/avancado/json_unmarshal/model"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	request, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para o servidor", err.Error())
		return
	}
	request.SetBasicAuth("teste", "Teste")
	resposta, err := client.Do(request)

	if err != nil {
		fmt.Println("[main] Erro ao realizar uma request para o servidor", err.Error())
		return
	}
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteúdo do servidor", err.Error())
			return
		}
		fmt.Println(" ")

		post := model.BlogPost{}

		err = json.Unmarshal(corpo, &post)
		if err != nil {
			fmt.Println("[main] Erro ao realizar o Unmarshall do objeto", err.Error())
			return
		}
		fmt.Printf("%+v", post)
	}

}
