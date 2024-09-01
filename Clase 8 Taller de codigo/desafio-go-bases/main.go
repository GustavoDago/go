package main

import (
	"desafio-go-bases/desafio-go-bases/internal/tickets"
	"fmt"
)

const (
	madrugada = "madrugada"
	mañana    = "mañana"
	tarde     = "tarde"
	noche     = "noche"
)

func main() {
	//creamos un canal para cada gorutina
	canal1 := make(chan string)
	canal2 := make(chan string)
	canal3 := make(chan string)
	gesTickets := tickets.GestorTickets{}

	// Primer requerimiento: Una función que calcule cuántas personas viajan a un país determinado.
	go func(pais string) {
		req1, err := gesTickets.GetTotalTickets(pais)
		if err != nil {
			canal1 <- fmt.Sprintf("Error en GetTotalTickets para %s: %v", pais, err)
			return
		}
		canal1 <- fmt.Sprintf("A %s viajaron %d personas", pais, req1)
	}("Brazil")

	// Segundo requerimiento: Una función que calcule cuántas personas viajan en un horario determinado.
	go func(horario string) {
		req2, err := gesTickets.GetCountByPeriod(horario)
		if err != nil {
			canal2 <- fmt.Sprintf("Error en GetCountByPeriod para %s: %v", horario, err)
			return
		}
		canal2 <- fmt.Sprintf("Viajaron %d personas durante la %s", req2, horario)
	}(mañana)

	// Tercer requerimiento: Una función que calcule el porcentaje de personas que viajaron a un país respecto del total
	go func(paisPorcentaje string) {
		req3, err := gesTickets.PercentageDestination(paisPorcentaje)
		if err != nil {
			canal3 <- fmt.Sprintf("Error en PercentageDestination para %s: %v", paisPorcentaje, err)
			return
		}
		canal3 <- fmt.Sprintf("El porcentaje de personas que viajaron a %s es %.2f", paisPorcentaje, req3)
	}("Indonesia")

	// imprimimos el resultado de las gorutinas
	fmt.Println(<-canal1)
	fmt.Println(<-canal2)
	fmt.Println(<-canal3)

}
