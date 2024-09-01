package main

import (
	"errors"
	"fmt"
	"log"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {
	minFun, err := operacion(minimum)
	if err != nil {
		log.Fatal(err)
	}
	maxFun, err := operacion(maximum)
	if err != nil {
		log.Fatal(err)
	}
	proFun, err := operacion(average)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("El mínimo es ", minFun(3,4,7,1,9,5,6))
	fmt.Println("El máximo es ", maxFun(3,4,7,1,9,5,6))
	fmt.Println("El promedio es ", proFun(3,4,7,1,9,5,6))
	minFun2, maxFun2, proFun2,err := operacionMulti()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("El mínimo es %d, | El máximo es %d | El promedio es %d\n", 
	minFun2(3,4,7,1,9,5,6),
	maxFun2(3,4,7,1,9,5,6),
	proFun2(3,4,7,1,9,5,6))
	errorFunc, err := operacion("error")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("El error es ", errorFunc(3,4,7,1,9,5,6))
}

func operacion(calculo string) (func(...int) int, error) {
	switch calculo {
	case minimum:
		return calculoMinimo, nil
	case maximum:
		return calculoMaximo, nil
	case average:
		return calculoAverage, nil
	default:
		return nil, errors.New("El cálculo no existe")
	}
}

func operacionMulti() (func(...int) int,func(...int) int,func(...int) int, error) {

		return calculoMinimo,calculoMaximo,calculoAverage, nil

}

func calculoMinimo(values ...int) int {
	min := values[0]
	for _, val := range values {
		if val < min {
			min = val
		}
	}
	return min
}
func calculoMaximo(values ...int) int {
	var max int
	for _, val := range values {
		if val > max {
			max = val
		}
	}
	return max
}
func calculoAverage(values ...int) int {
	var suma int
	for _, val := range values {
		suma += val
	}
	return suma/ len(values)
}