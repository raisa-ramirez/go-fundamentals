package main

import (
	"fmt"
	"math"
)

func main() {
	const texto = "Hola Srita."
	fmt.Println(texto)

	const pi = 3.1416
	fmt.Println(pi)

	const number = 9.80665
	fmt.Println(math.Round(number))

	const raiz = 4
	fmt.Printf("La raíz de %v es: %v", raiz, math.Sqrt(raiz))
}
