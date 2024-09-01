package main

import (
	"fmt"
)

type ErrSalary struct{
	message string
}

func main() {

	salary := 120000
	err := validarSalario(salary)
	if err != nil {
		fmt.Printf("%v",err.Error())
	}else{
		fmt.Println("Debe pagar impuestos")
	}


}

func (e *ErrSalary) Error() string  {
	return e.message
}

func validarSalario(salary int) error  {
	if salary <= 150000{
		return &ErrSalary{message: "Error: el salario ingresado no alcanza el mínimo imponible"}
	}
	return nil
}