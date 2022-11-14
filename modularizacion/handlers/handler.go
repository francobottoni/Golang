package handlers

import (
	"fmt"
)

func (cir *Cuadrado) Area() float32 {
	return cir.Altura * cir.Ancho
}

func (cua *Circulo) Area() float32 {
	return cua.Pi * (cua.Radio * cua.Radio)
}

func Calculador(g Geometria) {
	fmt.Println(g)
	fmt.Println(g.area())
}
