package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/saulocn/golang/cursodego/avancado/web_post/model"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	usuario := model.Usuario{}

	usuario.ID = 1
	usuario.Nome = "Saulo"

	conteudoAEnviar, err := json.Marshal(usuario)

	if err != nil {
		fmt.Println("Erro ao realizar o parse do usuário", err.Error())
		return
	}

	request, err := http.NewRequest("POST", "https://en805174n9qa9.x.pipedream.net/", bytes.NewBuffer(conteudoAEnviar))
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para o Request Bin", err.Error())
		return
	}
	request.SetBasicAuth("fizz", "buzz")
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	resposta, err := client.Do(request)

	if err != nil {
		fmt.Println("[main] Erro ao o serviço do Request Bin", err.Error())
		return
	}
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteúdo retornado pelo Request Bin", err.Error())
			return
		}
		fmt.Println(" ")
		fmt.Println(string(corpo))
	}

}
