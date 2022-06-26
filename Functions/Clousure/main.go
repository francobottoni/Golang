package main

import (
	"fmt"
	"strings"
)

//Clousures
func repeat(n int) func(cadena string) string {
	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}

func main() {

	repeat4 := repeat(4)
	fmt.Println(repeat4("Pepe"))
	fmt.Println(repeat4("Capo"))

}
