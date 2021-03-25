package main

import "fmt"

type person struct {
	name string
	age  uint8
}

func main() {
	// Declaracion de variables con structs
	var me person = person{name: "Jair", age: 20}
	you := person{name: "You", age: 21}

	fmt.Println(me)
	fmt.Println(you)
}
