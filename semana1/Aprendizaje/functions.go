package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type Person struct {
	name string
	age  uint8
}

func sum(p1 int, p2 int) int {
	return p1 + p2
}

// Función que retorna múltiples valores
func division(numerador int, denominador int) (float64, error) {
	if denominador == 0 {
		return 0, errors.New("No puede ser denominador 0")
	}
	return float64(numerador) / float64(denominador), nil
}

func greeting(person Person) {
	fmt.Println("Hola soy ", person.name)
}

func main() {

	var x int = 2
	var y int = 3

	fmt.Println(sum(x, y))

	var me Person = Person{name: "Jair", age: 12}

	var result, err = division(3, 0)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println("Resultado de división", result)

	greeting(me)

}
