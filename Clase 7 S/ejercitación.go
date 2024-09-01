package main

import (
	"fmt"
	"time"
)

func main() {

	devolucion := make(chan string)
	prestamo := make(chan string)

	go func() {
		defer close(devolucion)
		for {
			time.Sleep(time.Second)
			devolucion <- "Devolucion procesada"
		}
	}()
	go func() {
		defer close(prestamo)
		for {
			time.Sleep(time.Second / 2)
			prestamo <- "Préstamo procesado"
		}
	}()

	// go func() {
	// 	for devProcesada := range devolucion {
	// 		fmt.Println(devProcesada)
	// 	}
	// }()
	// go func() {
	// 	for presProcesada := range prestamo {
	// 		fmt.Println(presProcesada)
	// 	}
	// }()

	go func ()  {
		for{
			select {
			case msg:= <- devolucion:
				fmt.Println(msg)
			case msg:= <- prestamo:
				fmt.Println(msg)
			}
		}
	}()

	fmt.Scanln()
}
