package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Println("Hola mundo\nDame un numero")
	fmt.Scan(&a)
	fmt.Println("Dame otro numero")
	c = a + b
	fmt.Println("La suma de:", a, "+", b, " es:", c)
}
