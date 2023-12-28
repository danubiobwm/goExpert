package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	templates := []string{
		"content.html",
		"header.html",
		"footer.html",
	}

	t := template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Python", 40},
		{"Java", 40},
		{"JavaScript", 40},
		{"C#", 40},
	})

	if err != nil {
		panic(err)
	}
}
