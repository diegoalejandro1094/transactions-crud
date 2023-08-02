// ejemplo de uso de punteros en go
package main

import "fmt"

func main() {
	// Declarar una variable y asignarle un valor
	x := 42

	// Declarar un puntero que apunte a la dirección de memoria de 'x'
	var ptr *int
	ptr = &x

	fmt.Println("Valor de x:", *ptr)

	*ptr = 10
	fmt.Println("Nuevo valor de x:", x)
}
