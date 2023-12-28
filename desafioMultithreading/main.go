package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Address struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
}

func main() {
	// Substitua pelo CEP desejado
	cep := "63031-330"

	// Crie dois canais para receber os resultados das goroutines
	result1 := make(chan Result)
	result2 := make(chan Result)

	// Execute as duas goroutines simultaneamente
	go fetchAPI("https://cdn.apicep.com/file/apicep/"+cep+".json", result1)
	go fetchAPI("https://viacep.com.br/ws/"+cep+"/json/", result2)

	// Use select para aguardar o resultado da goroutine mais rápida
	select {
	case res := <-result1:
		displayResult(res, "API 1")
	case res := <-result2:
		displayResult(res, "API 2")
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: A solicitação excedeu o tempo limite de 1 segundo.")
	}
}

type Result struct {
	Address Address
	API     string
	Err     error
}

func fetchAPI(url string, resultChan chan Result) {
	startTime := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		resultChan <- Result{Err: err}
		return
	}
	defer resp.Body.Close()

	// Verifique o código de status da resposta
	if resp.StatusCode != http.StatusOK {
		resultChan <- Result{Err: fmt.Errorf("Erro na resposta da API %s. Código de status: %d", url, resp.StatusCode)}
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resultChan <- Result{Err: err}
		return
	}

	var address Address
	err = json.Unmarshal(body, &address)
	if err != nil {
		resultChan <- Result{Err: err}
		return
	}

	duration := time.Since(startTime)
	resultChan <- Result{Address: address, API: url, Err: nil}

	fmt.Printf("Tempo de resposta da %s: %v\n", url, duration)
}

func displayResult(result Result, apiName string) {
	if result.Err != nil {
		fmt.Printf("Erro ao obter dados da %s: %v\n", apiName, result.Err)
		return
	}

	fmt.Printf("Dados do endereço obtidos da %s:\n", apiName)
	fmt.Printf("CEP: %s\n", result.Address.Cep)
	fmt.Printf("Logradouro: %s\n", result.Address.Logradouro)
	fmt.Printf("Bairro: %s\n", result.Address.Bairro)
	fmt.Printf("Localidade: %s\n", result.Address.Localidade)
	fmt.Printf("UF: %s\n", result.Address.UF)
	fmt.Printf("API: %s\n", result.API)
}
