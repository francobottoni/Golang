package main

import (
	"log"
	"net/http"
	"text/template"
)

type Usuarios struct {
	Name string
	Edad int
}

var templates = template.Must(template.New("T").ParseGlob("templates/**/*.html"))
var errorTemplate = template.Must(template.ParseFiles("templates/error/error.html"))

// Handler
func Error(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	errorTemplate.Execute(rw, nil)
}

func readerTemplate(rw http.ResponseWriter, name string, data interface{}) {
	err := templates.ExecuteTemplate(rw, name, data)

	if err != nil {
		//http.Error(rw, "No es posible retornar el template", http.StatusInternalServerError)
		Error(rw, http.StatusInternalServerError)
	}
}

// Handler
func Index(rw http.ResponseWriter, r *http.Request) {

	usuario := Usuarios{Name: "franco", Edad: 23}
	//template := template.Must(template.New("index.html").ParseFiles("index.html", "base.html"))

	//template.Execute(rw, usuario)
	readerTemplate(rw, "index.html", usuario)
}

// Handler
func Registro(rw http.ResponseWriter, r *http.Request) {
	readerTemplate(rw, "registro.html", nil)
}

func main() {
	// Archivos estaticos
	staticFile := http.FileServer(http.Dir("static"))

	// Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/registro", Registro)

	// Mux de staticFile
	mux.Handle("/static/", http.StripPrefix("/static/", staticFile))

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe(), mux)
}
