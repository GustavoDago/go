package main

import (
	"errors"
	"fmt"
	"log"
)

// Ejercicio 1 - Calcular salario
// Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
// Categoría C: su salario es de $1.000 por hora.
// Categoría B: su salario es de $1.500 por hora, más un 20 % de su salario mensual.
// Categoría A: su salario es de $3.000 por hora, más un 50 % de su salario mensual.
// Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.

func main() {

  minutosTrabajados := 120
  categoria := "C"
  salario,err := calculoSalario(minutosTrabajados,categoria)

  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("El empleado con %d minutos trabajados y categoría %s cobrará $%.2f\n", minutosTrabajados, categoria, salario)


}

func calculoSalario(minutos int, cat string)(float64, error)  {
  horasTrabajadas := float64(minutos) / 60

  switch cat {
    case "C":
      return 1000 * horasTrabajadas,nil
    case "B":
      return 1500 * horasTrabajadas * 1.2,nil
    case "A":
      return 3000 * horasTrabajadas * 1.5, nil
    default:
      return 0, errors.New("La categoría no existe")
  }
}

