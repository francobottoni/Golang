package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		//El recover obtiene los panics que hubo en el programa y nos sirve para controlarlo
		if error := recover(); error != nil {
			fmt.Println("Al parecer el programa no finalizo de forma correcta")
		}
	}()

	if file, err := os.Open("hola.txt"); err != nil {
		panic("No fue posible obtener el archivo")
	} else {
		defer func() {
			fmt.Println("Cerrando el archivo..")
			file.Close()
		}()

		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:long])
		fmt.Println(contenidoArchivo)
	}
}
