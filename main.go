package main

import (
	"fmt"

	"./Funciones"
)

func main() {
	var a, b, c float64
	fmt.Println("Hola Mundo")
	fmt.Println("Dame un numero")
	fmt.Scan(&a)
	fmt.Println("Dame otro numero")
	fmt.Scan(&b)
	c = a + b
	fmt.Println("La suma de:", a, "+", b, " es:", c)
	c = a - b
	fmt.Println("La resta de:", a, "-", b, " es:", c)
	c = a * b
	fmt.Println("La multiplicacion de:", a, "*", b, " es:", c)
	c = a / b
	fmt.Println("La Division de:", a, "/", b, " es:", c)
	c = Funciones.Potencia(a, b)
	fmt.Println("la potencia de: ", a, " elevado a: ", b, " es:", c)
}
