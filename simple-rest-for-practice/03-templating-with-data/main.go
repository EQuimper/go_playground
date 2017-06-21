package main

import (
	"html/template"
	"net/http"
)

type person struct {
	Firstname string
	Lastname  string
	Email     string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	p1 := person{"James", "Bond", "jamesbond@gmail.com"}

	tpl.ExecuteTemplate(w, "index.gohtml", p1)
}
