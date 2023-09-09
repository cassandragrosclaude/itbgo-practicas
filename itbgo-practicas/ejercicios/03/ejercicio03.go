/*Realizar una aplicación que reciba  una variable con el número del mes.
Según el número, imprimir el mes que corresponda en texto.
¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
Ej: 7, Julio.
*/

package main

import (
	"fmt"
)

func ejercicios() {
	meses := map[int]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
	}

	var numeroMes int
	fmt.Print("Ingresa el número del mes: ")
	fmt.Scan(&numeroMes)

	nombreMes, existe := meses[numeroMes]
	if !existe {
		fmt.Println("Número de mes inválido")
	} else {
		fmt.Println("El mes correspondiente al número", numeroMes, "es", nombreMes)
	}
}
