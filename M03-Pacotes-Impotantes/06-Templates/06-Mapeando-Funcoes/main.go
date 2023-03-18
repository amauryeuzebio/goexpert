package main

import (
	"html/template"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

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

	t := template.New("content.html")

	t.Funcs(template.FuncMap{"ToUpper": ToUpper})

	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(os.Stdout, cursos)

	if err != nil {
		panic(err)
	}

}
