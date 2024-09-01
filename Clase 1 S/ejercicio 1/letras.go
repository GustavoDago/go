package main

import "fmt"

func main() {
	palabra := "AlumnosCTD"
	var cant int = len(palabra)

	fmt.Printf("La palabra %s tiene %d letras\n",palabra,cant)
	for _, letra := range palabra {
		fmt.Println(string( letra))
	}
}