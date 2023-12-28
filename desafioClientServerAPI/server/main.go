package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

const apiURL = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

const sqlCreateTable = `
	CREATE TABLE IF NOT EXISTS quotation (
		id varchar(255) NOT NULL PRIMARY KEY,
		code varchar(255),
		codein varchar(255),
		name varchar(255),
		high varchar(255),
		low varchar(255),
		varBid varchar(255),
		pctChange varchar(255),
		bid varchar(255),
		ask varchar(255),
		timestamp varchar(255),
		create_date varchar(255)
	);
`

type Quotation struct {
	Usdbrl struct {
		ID         string `json:"id"`
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

type Response struct {
	Value float64 `json:"value"`
}

func main() {
	db, err := createConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = createTable(db)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		ctxRequest, cancelRequest := context.WithTimeout(r.Context(), time.Millisecond*100)
		defer cancelRequest()

		quotation, err := getQuotationApi(ctxRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		quotation.Usdbrl.ID = uuid.New().String()

		ctxSaveDB, cancelSaveDB := context.WithTimeout(r.Context(), time.Millisecond*10)
		defer cancelSaveDB()

		err = saveQuotation(ctxSaveDB, db, quotation)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		valueQuotation, err := strconv.ParseFloat(quotation.Usdbrl.Bid, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := Response{
			Value: valueQuotation,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8080", nil)
}

func createConnection() (*sql.DB, error) {
	return sql.Open("sqlite3", "quotation.db")
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(sqlCreateTable)
	if err != nil {
		return err
	}
	return nil
}

func getQuotationApi(ctx context.Context) (*Quotation, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response Quotation
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func saveQuotation(ctx context.Context, db *sql.DB, quotation *Quotation) error {
	insertQuotation := `
			insert into quotation(id, code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date)
			values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`
	stmt, err := db.Prepare(insertQuotation)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		quotation.Usdbrl.ID,
		quotation.Usdbrl.Code,
		quotation.Usdbrl.Codein,
		quotation.Usdbrl.Name,
		quotation.Usdbrl.High,
		quotation.Usdbrl.Low,
		quotation.Usdbrl.VarBid,
		quotation.Usdbrl.PctChange,
		quotation.Usdbrl.Bid,
		quotation.Usdbrl.Ask,
		quotation.Usdbrl.Timestamp,
		quotation.Usdbrl.CreateDate)
	if err != nil {
		return err
	}
	return nil
}
