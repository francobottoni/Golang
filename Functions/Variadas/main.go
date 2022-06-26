package main

import "fmt"

func sumar(nombre string, numeros ...int) (string, int) {
	var total int

	mensaje := fmt.Sprintf("El valor de la suma de %s es:", nombre)

	for _, num := range numeros {
		total += num
	}

	return mensaje, total
}

func main() {

	msj, result := sumar("Franco", 1, 45, 6, 2, 6, 8)
	fmt.Println(msj, result)

}
