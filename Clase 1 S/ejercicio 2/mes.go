package main

import "fmt"

/*
Ejercicio 2 - A qué mes corresponde
Realizar una aplicación que contenga una variable con el número del mes.
1. Según el número, imprimir el mes que corresponda en texto.
2. ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?
*/
func main() {
	meses  := [12]string {"enero","febrero","marzo","abril","mayo","junio","julio","agosto","septiembre","octubre","noviembre","diciembre"}
	mes := 6
	if mes >0 || mes < 13{
	fmt.Println(meses[mes-1])}
	
}