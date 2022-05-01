package main

import "fmt"

func main() {
	// Hello world
	fmt.Println("Hello World!")

	/*
		Variables: declarar tipo de la variable es una buena práctica, especialmente
		cuando se quiere aprovechar el consumo de memoria (bits).
		Una vez que se asigna un tipo a una variable ya no podemos cambiar el tipo.
	*/
	var number int
	number = 25
	fmt.Println(number)
	number = 45
	fmt.Println(number)
	var name string
	name = "Alejandra Avila"
	fmt.Println(name)
	name = "Ale Avila"
	fmt.Println(name)

	// Asignación de varias variables
	var nombre string
	var numero int
	nombre = "Esteban"
	fmt.Println(nombre)
	fmt.Println(numero)
}
