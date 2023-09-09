/*En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje
“Error: el salario ingresado no alcanza el mínimo imponible"
y lanzalo en caso de que “salary” sea menor a 150.000.
De lo contrario, tendrás que imprimir por consola el mensaje “Debe pagar impuesto”.*/

package main

import (
	"fmt"
)

// ErrorSalarioBajo es una estructura que implementa la interfaz error.
type ErrorSalario struct{}

// Implementación del método Error para ErrorSalarioBajo.
func (e ErrorSalario) Error() string {
	return "Error: el salario ingresado no alcanza el mínimo imponible"
}

func calcularImpuesto(salario float64) (float64, error) {
	if salario < 150000 {
		return 0, ErrorSalario{}
	}
	// Realizar cálculos de impuestos aquí.
	return salario * 0.2, nil
}

func main() {
	salario := 180000.00

	impuesto, err := calcularImpuesto(salario)
	if err != nil {
		fmt.Println(err)
		return
	}
	// esto es %f  hace referencia donde se insertará el valor flotante en la cadena de formato.
	// el 2 hace referencia a la cantidad de decimales que va a tener despues del . (.00)
	fmt.Printf("Impuesto a pagar: %.2f\n", impuesto)
}
