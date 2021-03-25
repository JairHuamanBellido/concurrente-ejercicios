package main

import "fmt"

func main() {

	// Declaración de arreglos
	var foo [3]int
	var bar = [3]int{1, 2, 3}

	// Completar los espacios con  (hasta la posicion 1)
	var zer = [12]int{1, 5: 4, 6, 10: 100, 15}

	// Declaración de un arreglo sin conocer el tamaño
	var fis = [...]int{1, 2}

	// Declaración de arreglos bidimensionales
	var laf [2][3]int

	fmt.Println("foo", foo)
	fmt.Println("bar", bar)
	fmt.Println("zer", zer)
	fmt.Println("fis", fis)
	fmt.Println("laf", laf)

	// Imprimir el tamaño de un arreglo
	fmt.Println("bar", len(bar))
}
