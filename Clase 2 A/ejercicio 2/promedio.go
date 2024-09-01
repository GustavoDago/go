package main

import "fmt"

func main() {
	var values []float64 = []float64{10, 9, 7, 8.4, 6.2, 9.1, 2, 4.6, 7.9}
	fmt.Println(promediar(values...))
}

func promediar(values ...float64) float64 {
	var sumatoria float64
	for _, valor := range values {
		sumatoria += valor
	}
	return sumatoria / float64(len(values))
}