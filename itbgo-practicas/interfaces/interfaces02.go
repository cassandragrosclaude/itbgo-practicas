/*Ejercicio 2 - Productos
Algunas tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
La empresa tiene 3 tipos de productos: Pequeño, Mediano y Grande. (Se espera que sean muchos más)

Y los costos adicionales son:
Pequeño: solo tiene el costo del producto
Mediano: el precio del producto + un 3% de mantenerlo en la tienda
Grande: el precio del producto + un 6% de mantenerlo en la tienda, y adicional a eso $2500 de costo de envío.

El porcentaje de mantenerlo en la tienda es en base al precio del producto.
El costo de mantener el producto en stock en la tienda es un porcentaje del precio del producto.

Se requiere una función factory que reciba el tipo de producto y el precio y retorne una interfaz Producto que tenga el método Precio.

Se debe poder ejecutar el método Precio y que el método me devuelva el precio total en base al costo del producto y los adicionales en caso que los tenga
*/

package main

import (
	"fmt"
)

// Define la interfaz Producto con el método Precio
type Producto interface {
	Precio() float64
}

// Define el tipo Pequeño que implementa la interfaz Producto
type Pequeño struct {
	precio float64
}

func (p Pequeño) Precio() float64 {
	return p.precio
}

// Define el tipo Mediano que implementa la interfaz Producto
type Mediano struct {
	precio float64
}

func (m Mediano) Precio() float64 {
	return m.precio * 1.03 // 3% de costo adicional
}

// Define el tipo Grande que implementa la interfaz Producto
type Grande struct {
	precio float64
}

func (g Grande) Precio() float64 {
	return (g.precio * 1.06) + 2500.00 // 6% de costo adicional y $2500 de costo de envío
}

// Función factory para crear productos
func CrearProducto(tipo string, precio float64) Producto {
	switch tipo {
	case "Pequeño":
		return Pequeño{precio: precio}
	case "Mediano":
		return Mediano{precio: precio}
	case "Grande":
		return Grande{precio: precio}
	default:
		return nil // Manejar un tipo desconocido
	}
}

func main() {
	// Crear productos
	producto1 := CrearProducto("Pequeño", 100.00)
	producto2 := CrearProducto("Mediano", 200.00)
	producto3 := CrearProducto("Grande", 300.00)

	// Calcular y mostrar los precios totales
	fmt.Printf("Producto 1 (Pequeño): %.2f\n", producto1.Precio())
	fmt.Printf("Producto 2 (Mediano): %.2f\n", producto2.Precio())
	fmt.Printf("Producto 3 (Grande): %.2f\n", producto3.Precio())
}
