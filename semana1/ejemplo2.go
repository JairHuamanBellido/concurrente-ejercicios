package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Funciones  func name(var tipado) tipado
func sumatoria(tope int) int {
	if tope == 1 {
		return tope
	} else {
		return tope + sumatoria(tope-1)
	}
}

func main() {

	fmt.Print("Sumatoria de número \n")

	// Lectura de entrada
	bufferIn := bufio.NewReader(os.Stdin)
	ingreso, _ := bufferIn.ReadString('\n')

	// Captura el dato ingreso
	dato := strings.TrimSpace(ingreso)

	// Conversión del texto a entero
	numero, _ := strconv.Atoi(dato)

	// Imprimir resultado
	fmt.Printf("La sumartoria es hasta %d es %d", numero, sumatoria(numero))

}
