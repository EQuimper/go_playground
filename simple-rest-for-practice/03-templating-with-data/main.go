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

// type persons struct {
// 	persons []person
// }

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
	p2 := person{"Jane", "Doe", "janedoe@gmail.com"}

	p := []persons{p1, p2}

	tpl.ExecuteTemplate(w, "index.gohtml", p)
}
