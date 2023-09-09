/*Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
Según el siguiente mapa, ayudá a imprimir la edad de Benjamin.

  var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

Por otro lado también es necesario:
Saber cuántos de sus empleados son mayores de 21 años.
Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
Eliminar a Pedro del mapa.
*/

package main

import "fmt"

func ejercicio04() {

	var employees = map[string]int{"Benjamin": 20,
		"Nahuel": 26,
		"Brenda": 19,
		"Darío":  44,
		"Pedro":  30}

	delete(employees, "Pedro")
	employees["Federico"] = 25

	mayoresDe21 := 0
	for _, edad := range employees {
		if edad > 21 {
			mayoresDe21++
		}
	}
	fmt.Println("Cantidad de empleados mayores de 21 años:", mayoresDe21)

	fmt.Println("Lista de empleados:")
	for nombre, edad := range employees {
		fmt.Printf("%s: %d años\n", nombre, edad)
	}

}
