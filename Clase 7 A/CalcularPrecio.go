package main

import (
	"fmt"
	"strconv"
)

type Productos struct {
	nombre   string
	precio   float64
	cantidad int
}
type Servicios struct {
	nombre          string
	precio          float64
	minutosTrabajos int
}

type Mantenimiento struct {
	nombre string
	precio float64
}

func main() {
	// Crear objetos de Productos
	productos := []Productos{
		{nombre: "Producto 1", precio: 10.5, cantidad: 2},
		{nombre: "Producto 2", precio: 15.0, cantidad: 1},
		{nombre: "Producto 3", precio: 8.75, cantidad: 4},
		{nombre: "Producto 4", precio: 20.25, cantidad: 3},
		{nombre: "Producto 5", precio: 12.99, cantidad: 2},
	}

	// Crear objetos de Servicios
	servicios := []Servicios{
		{nombre: "Servicio 1", precio: 50.0, minutosTrabajos: 30},
		{nombre: "Servicio 2", precio: 75.0, minutosTrabajos: 45},
		{nombre: "Servicio 3", precio: 80.5, minutosTrabajos: 60},
		{nombre: "Servicio 4", precio: 60.75, minutosTrabajos: 20},
		{nombre: "Servicio 5", precio: 90.99, minutosTrabajos: 75},
	}

	// Crear objetos de Mantenimiento
	mantenimientos := []Mantenimiento{
		{nombre: "Mantenimiento 1", precio: 100.0},
		{nombre: "Mantenimiento 2", precio: 150.0},
		{nombre: "Mantenimiento 3", precio: 125.5},
		{nombre: "Mantenimiento 4", precio: 200.75},
		{nombre: "Mantenimiento 5", precio: 175.99},
	}

	totalProductos := make(chan float64)
	totalServicios := make(chan float64)
	totalMantenimiento := make(chan float64)

	go SumarProductos(productos, totalProductos)
	go SumarServicios(servicios, totalServicios)
	go SumarMantenimiento(mantenimientos, totalMantenimiento)

	sumaProductos := <- totalProductos
	sumaServicios := <- totalServicios
	sumaMantenimiento := <- totalMantenimiento

	totalGeneral := sumaProductos + sumaServicios + sumaMantenimiento
	totalGeneralBase10 := strconv.FormatFloat(totalGeneral, 'f', 2, 64)
	fmt.Println("Total General en Base 10:", totalGeneralBase10)
}

func SumarProductos(listaProductos []Productos, total chan float64)  {
	suma := 0.0
	for _, lp := range listaProductos {
		suma += lp.precio * float64(lp.cantidad)
	}
	total <- suma

}

func SumarMantenimiento(listaMantenimiento []Mantenimiento, total chan float64)  {
	suma := 0.0
	for _, l := range listaMantenimiento {
		suma += l.precio
	}
	total <- suma

}

func SumarServicios(listaServicios []Servicios, total chan float64)  {
	suma := 0.0
	for _, l := range listaServicios {
		if l.minutosTrabajos < 30 {
			suma += l.precio
		} else {
			suma += l.precio * float64(l.minutosTrabajos) / 30
		}
	}
	total <- suma
}
