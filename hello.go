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

const monitoramento = 3
const delay = 5

func main() {

	nome, num := mostraDoisRetornos()
	fmt.Println(nome, " ", num)

	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Log...")
			imprimeLog()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando desconhecido.")
			//os.Exit(-1)
		}

	}

	//	if comando == 1 {
	//		fmt.Println("Monitorando...")
	//	} else if comando == 2 {
	//		fmt.Println("Exibindo Log...")
	//	} else if comando == 0 {
	//		fmt.Println("Saindo do programa...")
	//	} else {
	//		fmt.Println("Comando desconhecido.")
	//	}

}

func mostraDoisRetornos() (string, int) {
	nome := "Taís"
	num := 42
	return nome, num
}

func exibeIntroducao() {
	var nome string = "Taís"
	versao := 1.1
	fmt.Println("Hello Mundo!", nome)
	fmt.Println("Este programa esta na versao: ", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sai do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi: ", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	//sites := []string{"https://www.alura.com.br", "https://www.caelum.com.br"}
	sites := lerTextoDoArquivo()

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site: ", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}

}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, " foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, " está com problemas. Status: ", resp.StatusCode)
		registraLog(site, false)
	}
}

func lerTextoDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
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
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site +
		" - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLog() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
