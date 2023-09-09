/*Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros
y devuelva el promedio. No se pueden introducir notas negativas. */

package main

import (
	"fmt"
)

func main() {
	notasEstudiante1 := []int{6, 2, 7, 9}
	promedioEstudiante1 := calcularPromedio(notasEstudiante1...)
	fmt.Printf("El promedio del estudiante 1 es: %.2f\n", promedioEstudiante1)

	notasEstudiante2 := []int{7, 5, 8, 7}
	promedioEstudiante2 := calcularPromedio(notasEstudiante2...)
	fmt.Printf("El promedio del estudiante 2 es: %.2f\n", promedioEstudiante2)

	notasEstudiante3 := []int{2, 2, 2, 2}
	promedioEstudiante3 := calcularPromedio(notasEstudiante3...)
	fmt.Printf("El promedio del estudiante 2 es: %.2f\n", promedioEstudiante3)
}

func calcularPromedio(notas ...int) float64 {

	total := 0
	cantidad := 0

	for _, nota := range notas {
		if nota < 0 {
			fmt.Println("No se permiten notas negativas")
			return 0.0
		}
		total += nota
		cantidad++
	}

	if cantidad == 0 {
		fmt.Println("No hay notas")
		return 0.0
	}

	promedio := float64(total) / float64(cantidad)
	return promedio
}
