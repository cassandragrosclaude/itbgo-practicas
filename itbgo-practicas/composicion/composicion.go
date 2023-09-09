package composicion

import "fmt"

type Autor struct {
	Nombre   string
	Apellido string
}

func main() {

}
func (a *Autor) nombreCompleto() string {
	return fmt.Sprintf("%s %s", a.Nombre, a.Apellido)
}
