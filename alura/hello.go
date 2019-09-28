package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	nome, _ := devolveNomeEIdade()
	exibeIntroducao(nome)
	leSitesDoArquivo()
	for {
		exibeMenu()
		switch leComando() {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido!")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao(nome string) {
	var versao float32 = 1.2
	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")
	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do Programa")
}

func devolveNomeEIdade() (string, int) {
	nome := "Saulo"
	idade := 31
	return nome, idade
}

func iniciarMonitoramento() {
	fmt.Println("Iniciando monitoramento...")
	sites := leSitesDoArquivo()
	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, site)
			testaSite(site)
		}
		fmt.Println("")
		time.Sleep(delay * time.Second)
	}
	fmt.Println("")
}

func testaSite(site string) {
	response, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro ao tentar acessar o site", err)
	}
	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso! StatusCode:", response.StatusCode)
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas! StatusCode:", response.StatusCode)
		registraLog(site, false)
	}

}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro ao tentar abrir o arquivo sites.txt!", err)
	}
	reader := bufio.NewReader(arquivo)
	for {
		linha, err := reader.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo log.txt", err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo log.txt", err)
	}
	fmt.Println(string(arquivo))
}
