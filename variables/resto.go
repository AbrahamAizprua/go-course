package variables

import (
	"fmt"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Sueldo2 float64
var Fecha time.Time

func RestoVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 123.45
	Sueldo2 = 123.4567
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Sueldo2)
	fmt.Println(Fecha)

}
