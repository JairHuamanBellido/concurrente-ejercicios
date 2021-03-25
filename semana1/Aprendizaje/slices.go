package main

import "fmt"

func sliceTutorial() {
	// Declaración de un slice , no lleva ningun valor dentro de []
	var x = []int{1, 2, 34}

	fmt.Println("x", x)
	fmt.Println("is null", x == nil)
	fmt.Println("len", len(x))

	// Añadir un elemento a slice
	x = append(x, 5)

	// Añadir varios valores a slice
	x = append(x, 6, 7, 8, 9)

	fmt.Println("x", x)
	fmt.Println("len", len(x))

	var y = []int{110, 111, 112}
	x = append(x, y...)

	fmt.Println("x", x)
}
func capTutorial() {
	fmt.Println("CAP SECTION ============")
	var newArr = []int{1}
	fmt.Println(newArr, len(newArr), cap(newArr))

	newArr = append(newArr, 2)
	fmt.Println(newArr, len(newArr), cap(newArr))

	newArr = append(newArr, 10)
	fmt.Println(newArr, len(newArr), cap(newArr))

	newArr = append(newArr, 5)
	fmt.Println(newArr, len(newArr), cap(newArr))

	newArr = append(newArr, 8)
	fmt.Println(newArr, len(newArr), cap(newArr))

}

func makeTutorial() {

	// La sentencia "make", sirve para crear un slice, definiendo la capacidad y el tamaño
	x := make([]int, 0, 10)

	x = append(x, 5, 6, 7, 8)
	println("MAKE SECTION -------")
	fmt.Println(x)
}

func slicingArrayTutorial() {
	var x = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("SLICING SECTION -----------")

	// Imprime hasta la posicion 2
	fmt.Println(x[:2])

	// Imprime luego de la posicion 2
	fmt.Println(x[2:])

	// Imprime luego de la posicion 2 hasta la 4
	fmt.Println(x[2:4])

	// Imprime todo el arreglo
	fmt.Println(x[:])
}

func copyTutorial() {
	fmt.Println("COPY SLICE SECTION -----------")
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println(y, num)

}
func main() {

	sliceTutorial()
	capTutorial()
	makeTutorial()
	slicingArrayTutorial()
	copyTutorial()
}
