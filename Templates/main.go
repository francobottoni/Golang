package main

import (
	"log"
	"net/http"
	"text/template"
)

type Usuarios struct {
	Name   string
	Edad   int
}


func Saludar() string {
	return "Hola desde una funcion"
}

// Handler
func Index(rw http.ResponseWriter, r *http.Request) {

	//fmt.Fprintln(rw, "Hola mundo!")
	template, err := template.New("index.html").ParseFiles("index.html", "base.html")
	usuario := Usuarios{Edad: 23, Name: "franco"}

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, usuario)
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
