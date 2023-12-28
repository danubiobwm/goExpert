package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const apiURL = "http://localhost:8080/cotacao"

type Response struct {
	Value float64 `json:"value"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*200)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		panic(err)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var bodyResponse Response
	err = json.Unmarshal(body, &bodyResponse)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write([]byte(fmt.Sprintf("Dólar: %.2f", bodyResponse.Value)))
	if err != nil {
		panic(err)
	}
	fmt.Println("Cotação gravada com sucesso.")
}
