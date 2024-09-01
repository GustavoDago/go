package main

import "fmt"

func main() {
	var (
		nombreCliente      = "Marcos"
		edad          uint = 25
		empleado           = true
		antiguedad    uint = 6
		sueldo        uint = 54000
	)
	fmt.Println(nombreCliente + " "+	aplicarReglas(edad, empleado, antiguedad, sueldo))
}
func aplicarReglas(edad uint, empleado bool, antiguedad uint, sueldo uint) string {
	mensaje := "No cumple las condiciones"
	if edad >22 && empleado == true && antiguedad > 1{
		if sueldo > 100000{
			mensaje = "Obtiene el préstamo sin interés"
		}else{
			mensaje = "Obtiene el préstamo con interés"
		}
	}
	return mensaje
}