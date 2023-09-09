package structmetodos

import "fmt"

/*Crear un programa que cumpla los siguiente puntos:
Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
Tener un slice global de Product llamado Products instanciado con valores.
2 métodos asociados a la estructura Product: Save(), GetAll().
El método Save() deberá tomar el slice de Products y añadir el producto desde el cual se llama al método.
El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
Ejecutar al menos una vez cada método y función definido desde main().*/

type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
	Category    string
}

//Declaro el slcice global

var Products = []Product{
	{ID: 1, Name: "Mouse", Price: 10000, Description: "Mouse inalámbrico Logitech", Category: "Categoría A"},
	{ID: 2, Name: "Monitor", Price: 119000, Description: "Monitor Samsung F27T350FHL led 27", Category: "Categoría B"},
	{ID: 3, Name: "Gabinete", Price: 55500, Description: "Gabinete Level Up Cassiopeia Atx 1 Cooler", Category: "Categoría A"},
}

// Metodo SAVE()
func (p Product) Save() {
	Products = append(Products, p)
}

// Metodo GET()
func GetAll() []Product {
	return Products
}

func getById(id int) Product {
	for _, p := range Products {
		if p.ID == id {
			return p
		}
	}
	return Product{}

}

func main() {

	// CreO productos y guardarlos usando el método Save()
	producto := Product{ID: 4, Name: "Cooler", Price: 20000, Description: "Cpu Cooler Xigmatek Windpower Wp964 Rgb Amd Intel", Category: "Categoría A"}

	producto.Save()

	idBuscado := 1
	productoEncontrado := getById(idBuscado)
	if productoEncontrado.ID != 0 {
		fmt.Printf("\nProducto encontrado con ID %d:\n", idBuscado)
		fmt.Printf("Nombre: %s, Precio: %.2f, Descripción: %s, Categoría: %s\n", productoEncontrado.Name, productoEncontrado.Price, productoEncontrado.Description, productoEncontrado.Category)
	} else {
		fmt.Printf("\nNo se encontró ningún producto con ID %d.\n", idBuscado)
	}

	allproducts := GetAll()

	fmt.Println(allproducts)
}
