package main

import "fmt"

func main() {
	var mapaEmpleados = map[string]uint{"Héctor": 115000, "Juanjo": 48000, "Estefanía": 160000}
	CalcularImpuestos(mapaEmpleados)
}

func CalcularImpuestos(mapaEmpleados map[string]uint) {
	for nombre, salario := range mapaEmpleados {

		respuesta := "Al empleado " + nombre + " no se le realizan descuentos"
		if salario > 50000 {respuesta="Al empleado " + nombre + " se le descuentan " + fmt.Sprintf("%d", salario * 17 /100) +" pesos" }
		if salario > 150000 {respuesta="Al empleado " + nombre + " se le descuentan " + fmt.Sprintf("%d", salario * 27 /100) +" pesos" }

		fmt.Println(respuesta)
	}
}