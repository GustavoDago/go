package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 43, 6, 34, 65, 87}
	cpar := make(chan int)
	cimpar := make(chan int)

	go par(cpar)
	go impar(cimpar)

	for _, v := range s {
		if v%2 == 0 {
			cpar <- v
		} else {
			cimpar <- v
		}
	}

	close(cpar)
	close(cimpar)
}

func par(c chan int) {
	for num := range c {
		fmt.Println("Par: ", num)
	}
}
func impar(c chan int) {
	for num := range c {
		fmt.Println("Impar: ", num)
	}
}
