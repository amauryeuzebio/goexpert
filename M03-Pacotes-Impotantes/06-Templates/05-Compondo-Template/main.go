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
		"header.html",
		"content.html",
		"footer.html",
	}

	cursos := Cursos{
		{"Go", 20},
		{"Java", 40},
		{"Python", 30},
		{"Node", 60},
	}

	t := template.Must(template.New("content.html").ParseFiles(templates...))

	err := t.Execute(os.Stdout, cursos)

	if err != nil {
		panic(err)
	}

}
