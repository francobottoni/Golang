package main

import (
	"figuras/modularizacion/handlers"
	"figuras/modularizacion/models"
)

func main() {
	circulo := models.Circulo{
		Pi:    3.14,
		Radio: 7,
	}

	cuadrado := models.Cuadrado{
		Ancho:  1.5,
		Altura: 2,
	}

	handlers.Calculador(&circulo)
	handlers.Calculador(&cuadrado)

}
