package main

import (
	"log"
	"net/http"
	"text/template"
)

type Usuarios struct {
	Name   string
	Edad   int
	Activo bool
	Admin  bool
	Cursos []Curso
}

type Curso struct {
	Nombre string
}

func Saludar() string {
	return "Hola desde una funcion"
}

// Handler
func Index(rw http.ResponseWriter, r *http.Request) {

	//fmt.Fprintln(rw, "Hola mundo!")
	template, err := template.ParseFiles("index.html")
	//usuario := Usuarios{Edad: 23, Name: "franco", Activo: true, Admin: false, Cursos: cursos}

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, nil)
	}
}

func main() {

	// Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe(), mux)
}
