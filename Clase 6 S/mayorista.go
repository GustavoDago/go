package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Se necesita un software desarrollado en Go que imprima las estimaciones de lo que generarían,
// los productos de cada categoría en un mayorista, en ventas de productos para el hogar.
// Para eso se detallarán los pasos para conseguirlo:

// 1 - Funcionalidad para generar un archivo CSV, llamado categorías.csv.

func main() {
	// Abrir el archivo CSV para escritura
	file, err := os.Create("categorias.csv")
	if err != nil {
		log.Fatal("Error al crear el archivo CSV:", err)
	}
	defer file.Close()

	// crear un escritor CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 2 -Cargar el archivo con las categorías
	// Crear la estructura de datos
	type categorias struct {
		ID          string
		Descripcion string
		Lista       string
	}
	var contenido [][]string
	contenido = append(contenido,
		[]string{"ID", "Nombre", "Lista"},
		[]string{"001", "Electrodomésticos", "ListaProductos"},
		[]string{"002", "Muebles", "ListaProductos"},
		[]string{"003", "Herramientas", "ListaProductos"},
		[]string{"004", "Pinturas", "ListaProductos"},
		[]string{"005", "Aberturas", "ListaProductos"},
		[]string{"006", "Construcción", "ListaProductos"},
		[]string{"007", "Automotor", "ListaProductos"})

	// escribir filas en el archivo csv
	err = writer.WriteAll(contenido)
	if err != nil {
		log.Fatal("Error de escritura")
	}

	// 3 - Generar para cada una de estas categorías los productos. Estos tendrán como información:
	// Código
	// Nombre
	// PrecioActual
	// CantidadActual
	// Insertar al menos cuatro productos.

	type producto struct {
		Codigo         string
		Nombre         string
		PrecioActual   float32
		CantidadActual int
	}

	productosPorCategoria := map[string][]producto{
		"Electrodomésticos": {{Codigo: "0012", Nombre: "Lavadora", PrecioActual: 5999.6, CantidadActual: 5},
			{Codigo: "0542", Nombre: "Refrigerador", PrecioActual: 59900.67, CantidadActual: 15},
			{Codigo: "1420", Nombre: "Horno microondas", PrecioActual: 2541.15, CantidadActual: 11},
			{Codigo: "012", Nombre: "Licuadora", PrecioActual: 745.8, CantidadActual: 25}},
		"Muebles": {{Codigo: "4012", Nombre: "Sofá", PrecioActual: 5999.6, CantidadActual: 5},
			{Codigo: "054", Nombre: "Mesa de comedor", PrecioActual: 500.67, CantidadActual: 15},
			{Codigo: "11420", Nombre: "Silla", PrecioActual: 241.15, CantidadActual: 11},
			{Codigo: "019", Nombre: "Cama", PrecioActual: 8745.8, CantidadActual: 2}},
		"Herramientas": {{Codigo: "00312", Nombre: "Destornillador", PrecioActual: 599.6, CantidadActual: 45},
			{Codigo: "042", Nombre: "Martillo", PrecioActual: 5990.0, CantidadActual: 11},
			{Codigo: "140", Nombre: "Sierra eléctrica", PrecioActual: 8141.15, CantidadActual: 1},
			{Codigo: "0812", Nombre: "Taladro", PrecioActual: 24745.8, CantidadActual: 55}}}
	// 4 - Generar otro archivo CSV, llamado estimaciones.csv.
	// Este tendrá los resultados de la suma de todos los productos de cada una de las categorías.

	fileEst, err := os.Create("estimaciones.csv")
	if err != nil {
		log.Fatal("Error al crear el archivo CSV:", err)
	}
	defer fileEst.Close()
	// crear un segundo escritor CSV
	writer2 := csv.NewWriter(fileEst)
	defer writer2.Flush()
	for categoría, productos := range productosPorCategoria {
		total := 0
		for _, producto := range productos {
			total += producto.CantidadActual * int(producto.PrecioActual)
		}
		registro := []string{categoría, strconv.Itoa(total)}
		err = writer2.Write(registro)
		// 5 - Imprimir todos los estimativos por consola.
		fmt.Printf("Categoría: %s - Total estimado: $%d\n", categoría, total)
	}

}
