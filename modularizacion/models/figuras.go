package models

type Geometria interface {
	Area() float32
}

type Cuadrado struct {
	Ancho  float32
	Altura float32
}
type Circulo struct {
	Pi    float32
	Radio float32
}
