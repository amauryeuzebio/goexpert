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
	cursos := Cursos{
		{"Go", 20},
		{"Java", 40},
		{"Python", 30},
		{"Node", 60},
	}

	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := t.Execute(os.Stdout, cursos)

	if err != nil {
		panic(err)
	}

}
