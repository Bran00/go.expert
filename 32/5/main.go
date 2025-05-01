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

	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))
	
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Python", 50},
		{"Java", 60},
		{"JavaScript", 50},
		{"C++", 60},
		{"C#", 50},
		{"PHP", 20},
		{"Ruby", 30},
	})

	if err != nil {
		panic(err)
	}
}
