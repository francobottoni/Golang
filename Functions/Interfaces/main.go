package main

import "fmt"

type Geometria interface {
	area() float32
}

type Circulo struct {
	ancho  float32
	altura float32
}
type Cuadrado struct {
	pi    float32
	radio float32
}

func (cir *Circulo) area() float32 {
	return cir.altura * cir.ancho
}

func (cua *Cuadrado) area() float32 {
	return cua.pi * (cua.radio * cua.radio)
}

func calculador(g Geometria) {
	fmt.Println(g)
	fmt.Println(g.area())
}

func main() {
	circulo := Circulo{
		ancho:  1.5,
		altura: 2,
	}

	cuadrado := Cuadrado{
		pi:    3.14,
		radio: 7,
	}

	calculador(&circulo)
	calculador(&cuadrado)

}
