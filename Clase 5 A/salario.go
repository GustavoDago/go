package main

import (
	"fmt"
)

func main() {
	salary := 120000
	txt := "Error: el mínimo imponible es de 150.000 y el salario ingresado es de %d"
	if salary < 150000 {
		err := fmt.Errorf(txt,salary)
		fmt.Println(err)
		return
	}
	fmt.Println("Debe pagar impuesto")
}