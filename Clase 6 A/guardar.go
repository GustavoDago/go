package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	type producto struct{
		idProducto int;
		precio float32;
		cantidad int;
	}
	f,err := os.Create("ProductosComprados.csv")
	check(err)

	defer f.Close()

	producto1 := producto{idProducto: 112,precio: 45.6,cantidad: 52}
	var st = "ID Producto: "+ strconv.Itoa(producto1.idProducto) + ", Precio: " +
		strconv.Itoa(int(producto1.precio)) + ", Cantidad: " + strconv.Itoa(producto1.cantidad)
	fmt.Println( st)
	n2, err := f.WriteString(st)
	check(err)
	fmt.Printf("Escrito %d bytes", n2)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}