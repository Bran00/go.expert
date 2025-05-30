package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("template.html"))
		err := t.Execute(w, Cursos{
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
	})
	http.ListenAndServe(":8282", nil)
}
