package main

import "fmt"

func main() {

	miFuncion := func(nombre string) string {
		return fmt.Sprintf("El nombre que ingresaste es: %s", nombre)
	}

	fmt.Println(miFuncion("Franco"))
}
