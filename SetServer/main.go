package main

import (
	"fmt"
	"log"
	"net/http"
)

func HandlerFuncHola(rw http.ResponseWriter, r *http.Request){
	fmt.Fprintln(rw,"Hola mundo!")
}

func main(){

	mux := http.NewServeMux()
	mux.HandleFunc("/Hola",HandlerFuncHola)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe(),mux)
}
