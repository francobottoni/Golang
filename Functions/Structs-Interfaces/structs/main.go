package main

import "fmt"

//Struct Persona
type Persona struct {
	nombre string
	edad   int
}

//Metodo
func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s , Edad: %d ", p.nombre, p.edad)
}

//'Herencia'
type Empleado struct {
	persona Persona
	sueldo  float32
}

func main() {
	p1 := Persona{nombre: "Franco", edad: 26}

	p1.imprimir()

	emp := Empleado{}
	emp.sueldo = 230000
	emp.persona.edad = 6
	emp.persona.nombre = "Boca"

	emp.persona.imprimir()
	fmt.Println(emp)

}
