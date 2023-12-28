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
	tmp := template.New("CursoTemplate")

	tmp, _ = tmp.Parse("Curso: {{.Nome}} - Carga Horaria:{{.CargaHoraria}}")
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
