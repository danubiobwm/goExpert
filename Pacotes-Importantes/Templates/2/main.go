package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string `json:"nome"`
	CargaHoraria int    `json:"cargaHoraria"`
}

func main() {
	curso := Curso{"Go", 40}
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horaria:{{.CargaHoraria}}"))

	err := t.Execute(os.Stdout, curso)

	if err != nil {
		panic(err)
	}
}
