// ejemplo de uso de interfaces en go
package main

import "fmt"

// Definimos una interfaz "Animal" con el método "HacerSonido"
type Animal interface {
	HacerSonido()
}

// Definimos el tipo "Perro" que implementa la interfaz "Animal"
type Perro struct{}

func (p Perro) HacerSonido() {
	fmt.Println("Guau guau!")
}

// Definimos el tipo "Gato" que también implementa la interfaz "Animal"
type Gato struct{}

func (g Gato) HacerSonido() {
	fmt.Println("Miau miau!")
}

func main() {
	// Creamos una lista de animales usando la interfaz "Animal"
	animales := []Animal{
		Perro{},
		Gato{},
	}

	// Iteramos sobre la lista de animales y hacemos que cada animal haga su sonido
	for _, animal := range animales {
		animal.HacerSonido()
	}
}
