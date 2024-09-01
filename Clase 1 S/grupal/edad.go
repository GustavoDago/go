package main

import "fmt"

// Qué edad tiene...
// Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente map, debemos imprimir la edad de Benjamin.
//   var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

// Por otro lado, también es necesario:
// Saber cuántos de sus empleados son mayores de 21 años.
// Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
// Eliminar a Pedro del map.

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Printf("La edad de %s es %d\n", "Benjamín", employees["Benjamin"])
	mayores := 0
	for _, edad := range employees {

		if edad > 20 {
			mayores++
		}
	}
	fmt.Printf("hay %d empleados mayores de 21 años\n", mayores)
	employees["Federico"] = 25
	fmt.Println(employees)
	delete(employees, "Pedro")
	fmt.Println(employees)

}
