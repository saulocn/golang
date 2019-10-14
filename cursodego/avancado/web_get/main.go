package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	resposta, err := client.Get("http://www.google.com.br")
	if err != nil {
		fmt.Println("[main] Erro ao abrir a página do Google", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteúdo da página do Google", err.Error())
			return
		}
		fmt.Println(string(corpo))
	}

	request, err := http.NewRequest("GET", "http://www.google.com.br", nil)
	if err != nil {
		fmt.Println("[main] Erro ao criar um request na página do Google", err.Error())
		return
	}
	request.SetBasicAuth("teste", "Teste")
	resposta, err = client.Do(request)

	if err != nil {
		fmt.Println("[main] Erro ao realizar uma request na página do Google", err.Error())
		return
	}
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteúdo da página do Google", err.Error())
			return
		}
		fmt.Println(" ")
		fmt.Println(string(corpo))
	}

}
