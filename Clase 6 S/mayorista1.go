package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type producto struct {
	Codigo         string
	Nombre         string
	PrecioActual   float32
	CantidadActual int
}

type categoria struct {
	ID        	string
	Nombre 		string
	Lista       string
}

func generarArchivoCategorias() {
	file, err := os.Create("categorias.csv")
	if err != nil {
		log.Fatal("Error al crear el archivo CSV:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	categorias := [][]string{
		{"ID", "Nombre", "Lista"},
		{"001", "Electrodomésticos", "ListaProductos"},
		{"002", "Muebles", "ListaProductos"},
		{"003", "Herramientas", "ListaProductos"},
		{"004", "Pinturas", "ListaProductos"},
		{"005", "Aberturas", "ListaProductos"},
		{"006", "Construcción", "ListaProductos"},
		{"007", "Automotor", "ListaProductos"},
	}

	err = writer.WriteAll(categorias)
	if err != nil {
		log.Fatal("Error de escritura")
	}
}

func generarProductos() map[string][]producto {
	productosPorCategoria := map[string][]producto{
		"Electrodomésticos": {
			{Codigo: "0012", Nombre: "Lavadora", PrecioActual: 5999.6, CantidadActual: 5},
			{Codigo: "0542", Nombre: "Refrigerador", PrecioActual: 59900.67, CantidadActual: 15},
			{Codigo: "1420", Nombre: "Horno microondas", PrecioActual: 2541.15, CantidadActual: 11},
			{Codigo: "012", Nombre: "Licuadora", PrecioActual: 745.8, CantidadActual: 25},
		},
		"Muebles": {
			{Codigo: "4012", Nombre: "Sofá", PrecioActual: 5999.6, CantidadActual: 5},
			{Codigo: "054", Nombre: "Mesa de comedor", PrecioActual: 500.67, CantidadActual: 15},
			{Codigo: "11420", Nombre: "Silla", PrecioActual: 241.15, CantidadActual: 11},
			{Codigo: "019", Nombre: "Cama", PrecioActual: 8745.8, CantidadActual: 2},
		},
		"Herramientas": {
			{Codigo: "00312", Nombre: "Destornillador", PrecioActual: 599.6, CantidadActual: 45},
			{Codigo: "042", Nombre: "Martillo", PrecioActual: 5990.0, CantidadActual: 11},
			{Codigo: "140", Nombre: "Sierra eléctrica", PrecioActual: 8141.15, CantidadActual: 1},
			{Codigo: "0812", Nombre: "Taladro", PrecioActual: 24745.8, CantidadActual: 55},
		},
	}

	return productosPorCategoria
}

func generarArchivoEstimaciones(productosPorCategoria map[string][]producto) {
	fileEst, err := os.Create("estimaciones.csv")
	if err != nil {
		log.Fatal("Error al crear el archivo CSV:", err)
	}
	defer fileEst.Close()

	writer2 := csv.NewWriter(fileEst)
	defer writer2.Flush()

	err = writer2.Write([]string{"Categoría","Suma productos"})
	for categoria, productos := range productosPorCategoria {
		total := 0
		for _, producto := range productos {
			total += producto.CantidadActual * int(producto.PrecioActual)
		}
		registro := []string{categoria, strconv.Itoa(total)}
		err = writer2.Write(registro)
		if err != nil {
			log.Fatal("Error de escritura")
		}

		fmt.Printf("Categoría: %s - Total estimado: $%d\n", categoria, total)
	}
}

func main() {
	generarArchivoCategorias()
	productosPorCategoria := generarProductos()
	generarArchivoEstimaciones(productosPorCategoria)
}