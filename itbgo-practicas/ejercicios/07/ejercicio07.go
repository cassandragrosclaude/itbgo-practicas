/*Una empresa marinera necesita calcular el salario de sus empleados basándose
en la cantidad de horas trabajadas por mes y la categoría.
Categoría C, su salario es de $1.000 por hora.
Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes,
la categoría y que devuelva su salario.
*/

package main

import (
	"fmt"
)

func calcularSalario(minutosTrabajados int, categoria string) float64 {
	var tarifaHora float64
	var porcentajeAdicional float64

	switch categoria {
	case "C":
		tarifaHora = 1000
		porcentajeAdicional = 0
	case "B":
		tarifaHora = 1500
		porcentajeAdicional = 0.2
	case "A":
		tarifaHora = 3000
		porcentajeAdicional = 0.5
	default:
		fmt.Println("Categoría inválida")
		return 0
	}

	horasTrabajadas := float64(minutosTrabajados) / 60
	salarioBase := tarifaHora * horasTrabajadas
	salarioAdicional := salarioBase * porcentajeAdicional
	salarioTotal := salarioBase + salarioAdicional

	return salarioTotal
}

func main() {
	minutosTrabajados := 1800 // Ejemplo: 30 horas trabajadas en un mes
	categoria := "B"          // Ejemplo: Categoría B

	salario := calcularSalario(minutosTrabajados, categoria)
	fmt.Printf("El salario es: $%.2f\n", salario)
}
