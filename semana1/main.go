package main

import "fmt"

func main(){

	// Declaración de variables tipadas var name string
	var name string = "Golang"
	var version float32 = 1.62

	fmt.Println(name," -> ",version)

	// Declaración de variables no tipadas x:= 12
	_name := "Golang"
	_version := 1.62

	fmt.Println(_name, " -> ", _version);


}