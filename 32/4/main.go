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

		t := template.Must(template.New("content.html").ParseFiles(templates...))
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