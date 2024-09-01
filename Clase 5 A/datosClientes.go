package main

import (
	"fmt"
	"os"
)

func main() {
	defer func(){
		fmt.Println("Ejecución finalizada")
	}()
	fmt.Println("Iniciando")
	_,err:= os.Open("customers.txt")
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")		
	}
}