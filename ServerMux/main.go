package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/books/{id}", GetBookHandler)
	mux.HandleFunc("/books/dir/{d...}", BooksPathHandler)
	mux.HandleFunc("/books/precedence/latest", BooksPrecedenceHandler)
	mux.HandleFunc("/books/precedence/{x}", BooksPrecedence2Handler)
	mux.HandleFunc("/books/precedence/other/{s}", BooksPrecedenceHandler)
	mux.HandleFunc("/categories/{category}/latest", BooksPrecedence2Handler)
	mux.HandleFunc("/books/", BooksHandler)
	http.ListenAndServe(":9000", mux)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("Book " + id))
}

func BooksPathHandler(w http.ResponseWriter, r *http.Request) {
	dirpath := r.PathValue("d")
	fmt.Fprintf(w, "Accessing directory path: %s\n", dirpath)
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books"))
}

func BooksPrecedenceHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books Precedence"))
}

func BooksPrecedence2Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books Precedence 2"))
}
