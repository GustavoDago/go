package main

import "fmt"

func main() {
	est:= Estudiante{
		Nombre: "Martín",
		Apellido: "Foris",
		DNI: "5255444",
		Fecha: "20/05/2002",
	}
	est.Detalle()
}

type Estudiante struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (e Estudiante) Detalle() {
	fmt.Printf("%s %s \nDNI:%s\nFecha de ingreso:%s",e.Nombre,e.Apellido,e.DNI,e.Fecha)
}